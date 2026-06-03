package queue

import (
	"APT/internal/consts"
	"APT/internal/library/cache"
	"APT/internal/library/h5FxPay"
	_ "APT/internal/logic"
	"APT/utility/rabbitmq"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	FxChangeOrderLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameFxChangeOrderPush)
)

func FxChangeOrderPush() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		PushParams   *h5FxPay.ChangeOrderPushRequest
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameFxChangeOrderPush
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(FxChangeOrderLogger)
	// 设置缓存适配器
	cache.SetAdapter(ctx)
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
	FxChangeOrderLogger.Info(ctx, QueueName+" Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		FxChangeOrderLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		if err = gjson.New(msg.Body).Scan(&PushParams); err != nil {
			FxChangeOrderLogger.Error(ctx, err.Error())
			_ = msg.Ack(false)
			continue
		}
		FxChangeOrderLogger.Info(ctx, "--[start]----------------------------------------------")
		if _, err = h5FxPay.ChangeOrderPush(ctx, PushParams, FxChangeOrderLogger); err != nil {
			FxChangeOrderLogger.Error(ctx, err.Error())
		}
		_ = msg.Ack(false)
	}
ERR:
	FxChangeOrderLogger.Error(ctx, err)
	panic(err)
}
