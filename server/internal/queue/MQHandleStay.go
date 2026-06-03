package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	_ "APT/internal/logic"
	"APT/internal/model/entity"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"fmt"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	OrderStayLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameOrderStay)
)

// OrderStay 同步AirhostOrder
func OrderStay() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameOrderStay
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(OrderStayLogger)
	MQConnStruct = &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: exchangeName,
		QueueName:    QueueName,
		RouteKey:     fmt.Sprintf("%s.%s", exchangeName, QueueName),
	}
	if err = MQConnStruct.Channel(); err != nil {
		err = gerror.New("创建队列失败")
		goto ERR
	}
	if err = MQConnStruct.Exchange("topic", true, false); err != nil {
		err = gerror.New("创建交换机失败")
		goto ERR
	}
	if err = MQConnStruct.Queue(true, false, nil); err != nil {
		err = gerror.New("创建通道失败")
		goto ERR
	}
	if err = MQConnStruct.Bind(); err != nil {
		err = gerror.New("交换机绑定队列失败")
		goto ERR
	}
	OrderStayLogger.Info(ctx, "OrderSync Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		OrderStayLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		OrderStayLogger.Info(ctx, "--[start]----------------------------------------------")
		OrderStayLogger.Info(ctx, msg.Body)
		if err = HandleOrderStayMQ(ctx, msg); err != nil {
			OrderStayLogger.Error(ctx, err)
		}
		_ = msg.Ack(false)
	}
ERR:
	OrderStayLogger.Error(ctx, err)
	panic(err)
}

func HandleOrderStayMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		StayInfo     *entity.PmsAppStay
		StayResponse *airhousePublicApi.RetrieveStayJSONDataRequest
	)
	OrderStayLogger.Info(ctx, "更新APP订单")
	defer func() {
		if r := recover(); r != nil {
			OrderStayLogger.Error(ctx, r)
		}
	}()
	time.Sleep(time.Second)
AAA:
	if StayResponse, err = airhousePublicApi.GetRetrieveStay(ctx, gvar.New(body).String()); err != nil {
		if err.Error()[:3] == "429" {
			OrderStayLogger.Error(ctx, err.Error)
			time.Sleep(time.Second * 5)
			goto AAA
		}
		err = gerror.New(err.Error())
		return
	}
	// 非APP订单
	if StayResponse.Data.ReservationSourceName != "AppAirHost" {
		OrderStayLogger.Info(ctx, "非APP订单")
		if err = service.HotelService().SyncOtaStay(ctx, &StayResponse.Data); err != nil {
			return
		}
		return
	}
	OrderStayLogger.Info(ctx, "APP订单")
	// APP订单
	if err = service.HotelService().SyncAppStay(ctx, &StayResponse.Data); err != nil {
		return
	}
	// 查询订单信息
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().Uid, StayResponse.Data.UID).Scan(&StayInfo); err != nil {
		return
	}
	// 转发到返利队列
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameRebate,
		DataByte:     gvar.New(StayInfo.OrderSn).Bytes(),
		Header:       nil,
	}); err != nil {
		OrderStayLogger.Error(ctx, err)
	}
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameExp,
		DataByte:     gvar.New(StayInfo.OrderSn).Bytes(),
		Header:       nil,
	}); err != nil {
		OrderStayLogger.Error(ctx, err)
	}
	OrderStayLogger.Info(ctx, "RUN_SUCCESS")
	return
}
