package rabbitmq

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendMqMessageParams struct {
	ExchangeName string
	QueueName    string
	DataByte     []byte
	Header       amqp.Table
}

func SendMqMessage(ctx context.Context, Params *SendMqMessageParams) (err error) {
	var (
		MQConnStruct *MQConnection
		PayConfig    *model.PayConfig
	)
	MQConnStruct = &MQConnection{
		Conn:         Conn,
		Chan:         nil,
		ExchangeName: Params.ExchangeName,
		QueueName:    Params.QueueName,
		RouteKey:     fmt.Sprintf("%s.%s", Params.ExchangeName, Params.QueueName),
	}
	if err = MQConnStruct.Channel(); err != nil {
		return
	}
	defer func() {
		_ = MQConnStruct.Chan.Close()
	}()
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	if g.IsEmpty(Params.Header) {
		Params.Header = amqp.Table{
			"x-delay": gvar.New(PayConfig.HotelStayExp * 1000).String(),
		}
	}
	if err = MQConnStruct.Publish(Params.DataByte, Params.Header, gctx.CtxId(ctx)); err != nil {
		return
	}
	g.Log().Infof(ctx, "EXCHANGE %s     QUEUENAME   %s     MESSAGE %s", Params.ExchangeName, Params.QueueName, Params.DataByte)
	return
}
