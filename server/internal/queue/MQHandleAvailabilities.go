package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	_ "APT/internal/logic"
	"APT/internal/model/entity"
	"APT/utility/convert"
	"APT/utility/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type AvailabilitiesJSONData struct {
	EndDate  string `json:"end_date"`
	Property struct {
		ID  string `json:"id"`
		UID string `json:"uid"`
	} `json:"property"`
	RoomType struct {
		ID  string `json:"id"`
		UID string `json:"uid"`
	} `json:"room_type"`
	StartDate         string `json:"start_date"`
	ObjectType        string `json:"object_type"`
	ConnectedRoomType struct {
		ID  string `json:"id"`
		UID string `json:"uid"`
	} `json:"connected_room_type"`
}

var (
	AvailabilitiesLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameAvailabilities)
)

func Availabilities() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameAvailabilities
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(AvailabilitiesLogger)
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
	AvailabilitiesLogger.Info(ctx, "AvailabilitiesSync Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		AvailabilitiesLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		AvailabilitiesLogger.Info(ctx, "--[start]----------------------------------------------")
		AvailabilitiesLogger.Info(ctx, msg.Body)
		if err = HandleAvailabilitiesMQ(ctx, msg); err != nil {
			AvailabilitiesLogger.Error(ctx, err)
		}
		_ = msg.Ack(false)
	}
	return
ERR:
	AvailabilitiesLogger.Error(ctx, err)
	panic(err)
}

func HandleAvailabilitiesMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		AvailabilitiesData *AvailabilitiesJSONData
		AvailabilitiesInfo *airhousePublicApi.AvailabilitiesJSONDataResponse
		RoomType           *entity.PmsRoomType
		Dates              []*convert.DateRange
	)
	if err = json.Unmarshal(msg.Body, &AvailabilitiesData); err != nil {
		err = gerror.New("解析消息体失败")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			AvailabilitiesLogger.Error(ctx, r)
		}
	}()
	if err = dao.PmsRoomType.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsRoomType.Columns().Uid: AvailabilitiesData.RoomType.ID,
	}).Scan(&RoomType); err != nil {
		err = gerror.New("查询房型信息失败")
		return
	}
	if g.IsEmpty(RoomType) {
		err = gerror.New("房型信息为空")
		return
	}
	if g.IsEmpty(RoomType.RatePlanId) {
		err = gerror.New("房型未设置价格计划")
		return
	}
	AvailabilitiesData.EndDate = gtime.Now().Add(time.Hour * 24 * 180).Format("Y-m-d")
	if Dates, err = convert.GenerateDateRanges(AvailabilitiesData.StartDate, AvailabilitiesData.EndDate, 25); err != nil {
		AvailabilitiesLogger.Error(ctx, err)
		return
	}
	for _, DateItem := range Dates {
		ItemStartDate := DateItem.StartDate.Format("Y-m-d")
		ItemEndDate := DateItem.EndDate.Format("Y-m-d")
		// 查询最新的库存信息
		if AvailabilitiesInfo, err = airhousePublicApi.GetAvailabilities(ctx, &airhousePublicApi.AvailabilitiesJSONDataRequest{
			StartDate:  ItemStartDate,
			EndDate:    ItemEndDate,
			RoomTypeId: AvailabilitiesData.RoomType.UID,
		}); err != nil {
			return
		}
		// 更新库存信息
		for _, v := range AvailabilitiesInfo.Data.Inventories {
			if _, err = dao.PmsAvailabilities.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsAvailabilities.Columns().Puid: AvailabilitiesData.Property.UID,
				dao.PmsAvailabilities.Columns().Tuid: AvailabilitiesData.RoomType.UID,
			}).WhereBetween(dao.PmsAvailabilities.Columns().Date, v.StartDate, v.EndDate).
				Data(g.MapStrAny{
					dao.PmsAvailabilities.Columns().Allotment: v.InventoryCount,
				}).Update(); err != nil {
				AvailabilitiesLogger.Error(ctx, err)
			}
		}
		// 更新价格信息
		for _, v := range AvailabilitiesInfo.Data.Rates {
			if RoomType.RatePlanId == v.RatePlan.ID {
				if _, err = dao.PmsAvailabilities.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsAvailabilities.Columns().Puid: AvailabilitiesData.Property.UID,
					dao.PmsAvailabilities.Columns().Tuid: AvailabilitiesData.RoomType.UID,
				}).WhereBetween(dao.PmsAvailabilities.Columns().Date, v.StartDate, v.EndDate).
					Data(g.MapStrAny{
						dao.PmsAvailabilities.Columns().Price: v.PerDayBaseRate,
					}).Update(); err != nil {
					AvailabilitiesLogger.Error(ctx, err)
				}
			}
			if v.Closed == "true" && RoomType.RatePlanId == v.RatePlan.ID {

				if _, err = dao.PmsAvailabilities.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsAvailabilities.Columns().Puid: AvailabilitiesData.Property.UID,
					dao.PmsAvailabilities.Columns().Tuid: AvailabilitiesData.RoomType.UID,
				}).WhereBetween(dao.PmsAvailabilities.Columns().Date, v.StartDate, v.EndDate).
					Data(g.MapStrAny{
						dao.PmsAvailabilities.Columns().Allotment: 0,
					}).Update(); err != nil {
					AvailabilitiesLogger.Error(ctx, err)
				}
			}
		}
		// 更新锁房信息
		// 批量更新房态信息
		//for _, Blocks := range AvailabilitiesInfo.Data.Blocks {
		//	if _, err = dao.PmsRoomStatus.Ctx(ctx).
		//		WhereBetween(dao.PmsRoomStatus.Columns().Date, Blocks.StartDate, Blocks.EndDate).
		//		Where(dao.PmsRoomStatus.Columns().Ruid, Blocks.RoomUnit.ID).
		//		OmitEmptyData().
		//		Data(&entity.PmsRoomStatus{
		//			Status: "LOCK",
		//		}).UpdateAndGetAffected(); err != nil {
		//		AvailabilitiesLogger.Error(ctx, err)
		//		continue
		//	}
		//}
	}
	AvailabilitiesLogger.Info(ctx, "RUN_SUCCESS")
	return
}
