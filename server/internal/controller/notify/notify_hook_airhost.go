package notify

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/hook"
)

func (c *ControllerHook) AirhostHook(ctx context.Context, req *hook.AirhostHookReq) (res *hook.AirhostHookRes, err error) {
	var (
		r                 = ghttp.RequestFromCtx(ctx)
		InsertId          int64
		DataByte          []byte
		PmsAirhostMessage *entity.PmsAirhostMessage
	)
	g.Log().Path("logs/HOOK/AIRHOST_HOOK").Info(ctx, r.Header)
	g.Log().Path("logs/HOOK/AIRHOST_HOOK").Info(ctx, gvar.New(r.GetBody()).String())

	if DataByte, err = json.Marshal(req.Body); err != nil {
		goto ERR
	}
	if err = dao.PmsAirhostMessage.Ctx(ctx).Where(dao.PmsAirhostMessage.Columns().MessageId, req.Id).Scan(&PmsAirhostMessage); err != nil && !errors.Is(err, sql.ErrNoRows) {
		goto ERR
	}
	if g.IsEmpty(PmsAirhostMessage) {
		if InsertId, err = dao.PmsAirhostMessage.Ctx(ctx).Data(&entity.PmsAirhostMessage{
			ObjectType:       req.ObjectType,
			MessageId:        req.Id,
			MessageEventCode: req.Event,
			MessageContent:   string(DataByte),
		}).InsertAndGetId(); err != nil {
			goto ERR
		}
		if g.IsEmpty(InsertId) {
			err = gerror.New("解析失败")
			goto ERR
		}
	}
	if strings.Contains(req.Event, "stay") {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderStay,
			DataByte:     gjson.New(req.Body).Get("id").Bytes(),
			Header:       nil,
		}); err != nil {
			goto ERR
		}
	} else if strings.Contains(req.Event, "reservation") {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderStay,
			DataByte:     gjson.New(req.Body).Get("stay.id").Bytes(),
			Header:       nil,
		}); err != nil {
			goto ERR
		}
	} else if strings.Contains(req.Event, "availabilities.update") {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameAvailabilities,
			DataByte:     DataByte,
			Header:       nil,
		}); err != nil {
			goto ERR
		}
	}

	// webhook 生产环境需要转发请求
	if g.Cfg().MustGet(ctx, "system.mode").String() == "produce" {
		HttpHook(r)
	}
	r.Response.WriteStatusExit(200, "200")
	return
ERR:
	g.Log().Path("logs/HOOK/AIRHOST_HOOK").Error(ctx, err)
	r.Response.WriteStatusExit(http.StatusInternalServerError)
	_ = service.WarningWorkWx().SendWarningWorkWx(ctx, `
### AIRHOST_HOOK_ERROR
*  链路ID：`+gctx.CtxId(r.Context())+`
>`+err.Error())

	return
}

func HttpHook(r *ghttp.Request) {
	var (
		buffParams   = bytes.NewBuffer(r.GetBody())
		ghttpReq     *http.Request
		targetPmsURL = "https://crsdev.yeebok.cn"
		client       = &http.Client{}
	)
	// 创建新的请求
	ghttpReq, _ = http.NewRequest(r.Method, targetPmsURL+r.URL.Path, buffParams)
	// 复制请求头
	ghttpReq.Header = r.Header
	// 发送请求到目标服务器
	_, _ = client.Do(ghttpReq)
}
