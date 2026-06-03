package queue

import (
	"APT/internal/consts"
	_ "APT/internal/logic"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"fmt"

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
	EmployeeActivityEffectLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameEmployeeActivityEffect)
)

func EmployeeActivityEffect() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		MQConnStruct *rabbitmq.MQConnection
		exchangeName = consts.RabbitMQExchangeDelayedName
		QueueName    = consts.RabbitMQQueueNameEmployeeActivityEffect
		err          error
	)
	g.DB().SetLogger(EmployeeActivityEffectLogger)
	MQConnStruct = &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: exchangeName,
		QueueName:    QueueName,
		RouteKey:     fmt.Sprintf("%s.%s", exchangeName, QueueName),
	}
	if err = MQConnStruct.Channel(); err != nil {
		err = gerror.New("create channel fail")
		goto ERR
	}
	if err = MQConnStruct.Exchange("x-delayed-message", true, false); err != nil {
		goto ERR
	}
	if err = MQConnStruct.Queue(true, false, nil); err != nil {
		goto ERR
	}
	if err = MQConnStruct.Bind(); err != nil {
		goto ERR
	}
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		goto ERR
	}
	EmployeeActivityEffectLogger.Info(ctx, consts.RabbitMQQueueNameEmployeeActivityEffect+" Queue START SUCCESSFUL")
	for msg := range MQMsg {
		ctx = gctx.New()
		EmployeeActivityEffectLogger.Info(ctx, "--[start]----------------------------------------------")
		EmployeeActivityEffectLogger.Info(ctx, msg.Body)
		if err = HandleEmployeeActivityEffectMQ(ctx, msg); err != nil {
			EmployeeActivityEffectLogger.Error(ctx, err)
		}
		_ = msg.Ack(true)
	}
	return
ERR:
	EmployeeActivityEffectLogger.Error(ctx, err)
	panic(err)
}

func HandleEmployeeActivityEffectMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	EmployeeActivityEffectLogger.Info(ctx, msg.Expiration)
	defer func() {
		if r := recover(); r != nil {
			EmployeeActivityEffectLogger.Error(ctx, r)
		}
	}()
	// TODO 活动生效业务
	// 预订单
	if err = service.EmployeeActivity().ActivityEffect(ctx, gvar.New(msg.Body).Uint64()); err != nil {
		err = gerror.New("活动生效处理失败")
		return
	}

	EmployeeActivityEffectLogger.Info(ctx, "RUN_SUCCESS")
	return
}
