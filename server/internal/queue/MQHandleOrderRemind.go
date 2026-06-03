package queue

import (
	"APT/internal/consts"
	_ "APT/internal/logic"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MsgJSONData struct {
	OrderSn string `json:"orderSn"`
	Event   string `json:"event"`
}

var (
	OrderRemindLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameOrderRemind)
)

func OrderRemind() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameOrderRemind
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(OrderRemindLogger)
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
	OrderRemindLogger.Info(ctx, "Order Remind Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		OrderRemindLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		OrderRemindLogger.Info(ctx, "--[start]----------------------------------------------")
		OrderRemindLogger.Info(ctx, msg.Body)
		var MsgBody *MsgJSONData

		if err = json.Unmarshal(msg.Body, &MsgBody); err != nil {
			err = gerror.New("解析消息体失败")
			return
		}

		if err = HandleOrderRemindMQ(ctx, MsgBody.OrderSn, MsgBody.Event); err != nil {
			OrderRemindLogger.Error(ctx, err.Error())
		}

		_ = msg.Ack(false)
	}
ERR:
	OrderRemindLogger.Error(ctx, err)
	panic(err)
}

func HandleOrderRemindMQ(ctx context.Context, orderSn string, event string) (err error) {
	var (
		in      *input_basics.SendTemplateInp
		OrderSn string
	)
	in = new(input_basics.SendTemplateInp)
	OrderSn = orderSn
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			OrderRemindLogger.Error(ctx, r)
		}
	}()
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		in.OrderSn = orderSn
		in.Event = event
		if err = service.BasicsSmsTemplate().SendTemplate(ctx, in); err != nil {
			err = gerror.New("短信发送失败")
			return
		}

		OrderRemindLogger.Info(ctx, "RUN_SUCCESS")
		return
	}); err != nil {
		OrderRemindLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
