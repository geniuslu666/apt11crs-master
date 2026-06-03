package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	_ "APT/internal/logic"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	OrderExportLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameOrderExport)
)

func OrderExport() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameOrderExport
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(OrderExportLogger)
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
	OrderExportLogger.Info(ctx, "Order Export Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		OrderExportLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		OrderExportLogger.Info(ctx, "--[start]----------------------------------------------")
		OrderExportLogger.Info(ctx, msg.Body)

		if err = HandleOrderExportMQ(ctx, msg); err != nil {
			OrderExportLogger.Error(ctx, err.Error())
		}

		_ = msg.Ack(false)
	}
ERR:
	OrderExportLogger.Error(ctx, err)
	panic(err)
}

func HandleOrderExportMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		id              int64
		orderExportInfo entity.OrderExport
		carOut          *input_car.CarOrderExportInp
		spaOut          *input_spa.SpaOrderExportInp
		foodOut         *input_food.FoodOrderExportInp
		cabinetOut      *input_cabinet.OrderExportInp
		filePath        string
	)
	id = gvar.New(msg.Body).Int64()
	// 查询导出信息
	if err = dao.OrderExport.Ctx(ctx).WherePri(id).Scan(&orderExportInfo); err != nil {
		err = gerror.New("查询导出记录失败")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			OrderExportLogger.Error(ctx, r)
		}
	}()
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.OrderExport.Ctx(ctx).WherePri(id).Data(g.Map{
			"status": 0,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if orderExportInfo.Scene == 5 {
			err = json.Unmarshal([]byte(orderExportInfo.Condition), &cabinetOut)
			if err != nil {
				panic(err)
			}

			// 储物柜
			if filePath, err = service.CabinetService().StartExport(ctx, cabinetOut); err != nil {
				err = gerror.New("导出失败")
				return
			}
		} else if orderExportInfo.Scene == 4 {
			err = json.Unmarshal([]byte(orderExportInfo.Condition), &carOut)
			if err != nil {
				panic(err)
			}

			//接送机
			if filePath, err = service.CarOrder().StartExport(ctx, carOut); err != nil {
				err = gerror.New("导出失败")
				return
			}
		} else if orderExportInfo.Scene == 3 {
			err = json.Unmarshal([]byte(orderExportInfo.Condition), &spaOut)
			if err != nil {
				panic(err)
			}

			//接送机
			if filePath, err = service.SpaOrder().StartExport(ctx, spaOut); err != nil {
				err = gerror.New("导出失败")
				return
			}
		} else if orderExportInfo.Scene == 2 {
			err = json.Unmarshal([]byte(orderExportInfo.Condition), &foodOut)
			if err != nil {
				panic(err)
			}

			//接送机
			if filePath, err = service.FoodOrder().StartExport(ctx, foodOut); err != nil {
				err = gerror.New("导出失败")
				return
			}
		}
		if !g.IsEmpty(filePath) {
			if _, err = dao.OrderExport.Ctx(ctx).WherePri(id).Data(g.Map{
				"path":   filePath,
				"status": 1,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
		} else {
			err = gerror.Wrap(err, "导出失败，请稍后重试！")
			return
		}

		OrderExportLogger.Info(ctx, "RUN_SUCCESS")
		return
	}); err != nil {
		if _, err = dao.OrderExport.Ctx(ctx).WherePri(id).Data(g.Map{
			"status": 2,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}
		OrderExportLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
