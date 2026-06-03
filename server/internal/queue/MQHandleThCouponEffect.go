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
	ThCouponEffectLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameThCouponEffect)
)

func ThCouponEffect() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		MQConnStruct *rabbitmq.MQConnection
		exchangeName = consts.RabbitMQExchangeDelayedName
		QueueName    = consts.RabbitMQQueueNameThCouponEffect
		err          error
	)
	g.DB().SetLogger(ThCouponEffectLogger)
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
	ThCouponEffectLogger.Info(ctx, consts.RabbitMQQueueNameThCouponEffect+" Queue START SUCCESSFUL")
	for msg := range MQMsg {
		ctx = gctx.New()
		ThCouponEffectLogger.Info(ctx, "--[start]----------------------------------------------")
		ThCouponEffectLogger.Info(ctx, msg.Body)
		if err = HandleThCouponEffectMQ(ctx, msg); err != nil {
			ThCouponEffectLogger.Error(ctx, err)
		}
		_ = msg.Ack(true)
	}
	return
ERR:
	ThCouponEffectLogger.Error(ctx, err)
	panic(err)
}

func HandleThCouponEffectMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	defer func() {
		if r := recover(); r != nil {
			ThCouponEffectLogger.Error(ctx, r)
		}
	}()
	ThCouponEffectLogger.Info(ctx, msg.Expiration)
	// TODO 礼品券生效业务
	// 预订单
	if err = service.ThMemberCoupon().CouponEffect(ctx, gvar.New(msg.Body).String()); err != nil {
		err = gerror.New("礼品券生效处理失败")
		return
	}

	ThCouponEffectLogger.Info(ctx, "RUN_SUCCESS")
	return
}
