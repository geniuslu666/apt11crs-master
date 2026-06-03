package queue

import (
	"APT/internal/consts"
	_ "APT/internal/logic"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageJSONData struct {
	Type        string `json:"type"`
	Id          int    `json:"id"`
	CheckInDate string `json:"checkInDate"`
	MemberId    int    `json:"memberId"`
}

var (
	OrderAwardLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameOrderAward)
)

func OrderAward() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		MQConnStruct *rabbitmq.MQConnection
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameOrderAward
		err          error
	)
	g.DB().SetLogger(OrderAwardLogger)
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
	if err = MQConnStruct.Exchange("topic", true, false); err != nil {
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
	OrderAwardLogger.Info(ctx, consts.RabbitMQQueueNameOrderAward+" Queue START SUCCESSFUL")
	for msg := range MQMsg {
		ctx = gctx.New()
		OrderAwardLogger.Info(ctx, "--[start]----------------------------------------------")
		OrderAwardLogger.Info(ctx, msg.Body)
		if err = HandleOrderAwardMQ(ctx, msg); err != nil {
			OrderAwardLogger.Error(ctx, err)
		}
		_ = msg.Ack(true)
	}
	return
ERR:
	OrderAwardLogger.Error(ctx, err)
	panic(err)
}

func HandleOrderAwardMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var MessageData *MessageJSONData

	if err = json.Unmarshal(msg.Body, &MessageData); err != nil {
		err = gerror.New("解析消息体失败")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			OrderAwardLogger.Error(ctx, r)
		}
	}()
	if MessageData.Type == "AWARD" {
		// 发放
		if err = DoAward(ctx, MessageData); err != nil {
			OrderAwardLogger.Error(ctx, err)
		}
	} else {
		// 失效
		if err = DoInvalid(ctx, MessageData); err != nil {
			OrderAwardLogger.Error(ctx, err)
		}
	}
	return
}

func DoAward(ctx context.Context, MessageData *MessageJSONData) (err error) {
	if err = service.HotelService().HotelOrderAward(ctx, MessageData.Id, MessageData.CheckInDate, MessageData.MemberId); err != nil {
		return
	}

	OrderAwardLogger.Info(ctx, "AWARD_SUCCESS")
	return
}

func DoInvalid(ctx context.Context, MessageData *MessageJSONData) (err error) {
	if err = service.HotelService().HotelOrderAwardInvalid(ctx, MessageData.Id); err != nil {
		return
	}

	OrderAwardLogger.Info(ctx, "INVALID_SUCCESS")
	return
}
