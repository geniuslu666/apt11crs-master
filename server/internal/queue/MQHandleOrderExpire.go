package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	_ "APT/internal/logic"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	OrderExpireLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameOrderExpire)
)

func OrderExpire() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		MQConnStruct *rabbitmq.MQConnection
		exchangeName = consts.RabbitMQExchangeDelayedName
		QueueName    = consts.RabbitMQQueueNameOrderExpire
		err          error
	)
	g.DB().SetLogger(OrderExpireLogger)
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
	OrderExpireLogger.Info(ctx, consts.RabbitMQQueueNameOrderExpire+" Queue START SUCCESSFUL")
	for msg := range MQMsg {
		ctx = gctx.New()
		OrderExpireLogger.Info(ctx, "--[start]----------------------------------------------")
		OrderExpireLogger.Info(ctx, msg.Body)
		if err = HandleOrderExpireMQ(ctx, msg); err != nil {
			OrderExpireLogger.Error(ctx, err)
		}
		_ = msg.Ack(true)
	}
	return
ERR:
	OrderExpireLogger.Error(ctx, err)
	panic(err)
}

func HandleOrderExpireMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		gHttpClient = g.Client()
		response    *gclient.Response
	)
	g.Log().Info(ctx, msg.Expiration)
	defer func() {
		if r := recover(); r != nil {
			OrderExpireLogger.Error(ctx, r)
		}
	}()
	// TODO 订单过期业务
	if gvar.New(msg.Body).String()[:1] == "C" {
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/car/order/cancelOrder", g.MapStrAny{
			"orderSn": string(msg.Body),
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:1] == "S" {
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/spa/order/cancelOrder", g.MapStrAny{
			"orderSn": string(msg.Body),
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:1] == "F" {
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/food/order/cancelOrder", g.MapStrAny{
			"orderSn": string(msg.Body),
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:3] == "R-F" {
		// 尾款支付超时
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/food/order/refundDeposit", g.MapStrAny{
			"orderSn": gvar.New(msg.Body).String()[2:],
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:1] == "B" {
		if err = service.CabinetService().OrderExpiration(ctx, gvar.New(msg.Body).String()); err != nil {
			err = gerror.New("订单过期处理失败")
			return
		}
	} else if gvar.New(msg.Body).String()[:1] == "T" {
		// 一日游
		if err = service.TravelOrder().OrderExpiration(ctx, gvar.New(msg.Body).String()); err != nil {
			err = gerror.New("订单过期处理失败")
			return
		}
	} else if gvar.New(msg.Body).String()[:2] == "HC" {
		// 预订变更单
		if _, err = dao.PmsAppReservationChange.Ctx(ctx).
			Where(dao.PmsAppReservationChange.Columns().ChangeOrderSn, gvar.New(msg.Body).String()).
			Where(dao.PmsAppReservationChange.Columns().ChangeStatus, "ING").
			Data(g.Map{
				dao.PmsAppReservationChange.Columns().ChangeStatus: "FAIL",
			}).Update(); err != nil {
			return
		}

	} else if gvar.New(msg.Body).String()[:3] == "Y-C" {
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/car/order/overdueOrder", g.MapStrAny{
			"orderSn": gvar.New(msg.Body).String()[2:],
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:3] == "Y-S" {
		gHttpClient.SetHeader("Content-Type", "application/json")
		if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/spa/order/overdueOrder", g.MapStrAny{
			"orderSn": gvar.New(msg.Body).String()[2:],
		}); err != nil {
			return
		}
		OrderExpireLogger.Info(ctx, response.Raw())
		if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
			return
		}
	} else if gvar.New(msg.Body).String()[:3] == "Y-T" {
		// 一日游-逾期
		if err = service.TravelOrder().OrderOverdue(ctx, gvar.New(msg.Body).String()[2:]); err != nil {
			err = gerror.New("订单逾期处理失败")
			return
		}
	} else {
		// 预订单
		if err = service.HotelService().OrderExpiration(ctx, gvar.New(msg.Body).String()); err != nil {
			err = gerror.New("订单过期处理失败")
			return
		}
	}

	OrderExpireLogger.Info(ctx, "RUN_SUCCESS")
	return
}
