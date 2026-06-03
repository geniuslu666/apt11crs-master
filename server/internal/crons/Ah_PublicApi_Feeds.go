package crons

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/model/entity"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

//stay.create						预订单创建
//stay.update						预订单更新
//room_reservation.create			订单创建
//room_reservation.update			订单更新
//room_reservation.checked_in       订单签入
//room_reservation.checked_out      订单签出
//room_guest.create                 房间来宾创建
//room_guest.update                 房间来宾更新
//room_guest.destroy                房间来宾销毁
//payment.create                    支付信息创建
//payment.update                    支付信息更新
//availabilities.update             可用性更新

// AhPublicApiFeeds 执行同步
func AhPublicApiFeeds(ctx context.Context) (err error) {
	var (
		Feeds             airhousePublicApi.FeedJSONData
		PmsAirhostMessage *entity.PmsAirhostMessage
		InsertId          int64
		DataByte          []byte
		Logger            = g.Log().Path("logs/CRON/AhPublicApiFeeds")
	)
	Logger.Info(ctx, "--[执行开始]----------------------------------------")
	g.DB().SetLogger(Logger)
	if Feeds, err = airhousePublicApi.GetFeed(ctx); err != nil {
		return err
	}
	Logger.Info(ctx, Feeds)

	for _, Data := range Feeds.Data {
		// 插入消息到message表中去
		if err = dao.PmsAirhostMessage.Ctx(ctx).Where(dao.PmsAirhostMessage.Columns().MessageId, Data.ID).Scan(&PmsAirhostMessage); err != nil && !errors.Is(err, sql.ErrNoRows) {
			Logger.Error(ctx, err)
			return
		}
		if g.IsEmpty(PmsAirhostMessage) {
			if InsertId, err = dao.PmsAirhostMessage.Ctx(ctx).Data(&entity.PmsAirhostMessage{
				ObjectType:       Data.ObjectType,
				MessageId:        Data.ID,
				MessageEventCode: Data.Event,
				MessageContent:   string(DataByte),
			}).InsertAndGetId(); err != nil {
				Logger.Error(ctx, err)
				return
			}
			if g.IsEmpty(InsertId) {
				err = gerror.New("解析失败")
				Logger.Error(ctx, err)
				return
			}
		}
		if DataByte, err = json.Marshal(Data.Body); err != nil {
			Logger.Error(ctx, err)
			return
		}
		if strings.Contains(Data.Event, "stay") {
			if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameOrderStay,
				DataByte:     gjson.New(DataByte).Get("id").Bytes(),
				Header:       nil,
			}); err != nil {
				Logger.Error(ctx, err)
				return
			}
		} else if strings.Contains(Data.Event, "reservation") {
			if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameOrderStay,
				DataByte:     gjson.New(DataByte).Get("stay.id").Bytes(),
				Header:       nil,
			}); err != nil {
				Logger.Error(ctx, err)
				return
			}
		} else {
			Logger.Error(ctx, "不支持的消息类型"+Data.Event)
		}
		if err = airhousePublicApi.CommitFeed(ctx, Data.ID); err != nil {
			err = gerror.New("完成消息处理通知失败")
		}
	}
	Logger.Info(ctx, "执行成功")
	return err
}
