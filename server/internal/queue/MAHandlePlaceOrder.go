package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/cabinetApi"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	_ "APT/internal/logic"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shopspring/decimal"
)

var (
	PlaceOrderLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNamePlaceOrder)
)

func PlaceOrder() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNamePlaceOrder
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(PlaceOrderLogger)
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
	PlaceOrderLogger.Info(ctx, QueueName+" Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		PlaceOrderLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		orderSn := gvar.New(msg.Body).String()
		PlaceOrderLogger.Info(ctx, "--[start]----------------------------------------------")
		PlaceOrderLogger.Info(ctx, orderSn)
		if g.IsEmpty(orderSn) {
			PlaceOrderLogger.Error(ctx, "订单号错误")
			_ = msg.Ack(false)
			continue
		}

		if orderSn[:1] == "C" {
			// TODO -出行订单支付完成
			if err = service.PayService().HandleCarOrderMq(ctx, string(msg.Body), PlaceOrderLogger); err != nil {
				PlaceOrderLogger.Error(ctx, err.Error())
			}
		} else if orderSn[:1] == "S" {
			// TODO -按摩订单支付完成
			if err = service.PayService().HandleSpaOrderMq(ctx, string(msg.Body), PlaceOrderLogger); err != nil {
				PlaceOrderLogger.Error(ctx, err.Error())
			}
		} else if orderSn[:1] == "F" {
			// TODO 餐饮订单支付完成
			if err = service.PayService().HandleFoodOrderMq(ctx, string(msg.Body), PlaceOrderLogger); err != nil {
				PlaceOrderLogger.Error(ctx, err.Error())
			}
		} else if orderSn[:1] == "B" {
			// TODO 储物柜订单支付完成
			if err = HandleCabinetOrderMQ(ctx, msg); err != nil {
				PlaceOrderLogger.Error(ctx, err.Error())
			}
		} else if orderSn[:1] == "T" {
			// TODO 一日游订单支付完成
			if err = HandleTravelOrderMQ(ctx, msg); err != nil {
				PlaceOrderLogger.Error(ctx, err.Error())
			}
		} else if orderSn[:1] == "H" {
			if orderSn[:2] == "HC" {
				// TODO 酒店变更完成
				if err = HandleChangeMQ(ctx, msg); err != nil {
					PlaceOrderLogger.Error(ctx, err.Error())
				}
			} else {
				// TODO 酒店订单支付完成
				if err = HandleStayMQ(ctx, msg); err != nil {
					PlaceOrderLogger.Error(ctx, err.Error())
				}
			}
		}
		if err == nil && !g.IsEmpty(orderSn) {
			if err = cache.Instance().Set(ctx, fmt.Sprintf("PayOrderSn_%s", string(msg.Body)), "SUCCESS", gtime.M*10); err != nil {
				PlaceOrderLogger.Error(ctx, err)
			}
		}
		_ = msg.Ack(false)
	}
ERR:
	PlaceOrderLogger.Error(ctx, err)
	panic(err)
}

func HandleStayMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		PmsBalTransaction  entity.PmsTransaction
		PmsAppStay         *entity.PmsAppStay
		MemberInfo         *entity.PmsMember
		cookieBookerParams *airhousePublicApi.CreateStayJSONDataRequest
		BookerResponse     *airhousePublicApi.CreateStayJSONDataResponse
		tx                 gdb.TX
		OrderSn            string
	)
	defer func() {
		if r := recover(); r != nil {
			PlaceOrderLogger.Error(ctx, r)
		}
	}()
	OrderSn = gvar.New(msg.Body).String()
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
			for _, v := range BookerResponse.Data.RoomReservations {
				if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
					ExchangeName: consts.RabbitMQExchangeName,
					QueueName:    consts.RabbitMQQueueNameOrderStay,
					DataByte:     gvar.New(v.Stay.ID).Bytes(),
					Header:       nil,
				}); err != nil {
					return
				}
			}
		}
	}()
	// 查询住宿单信息
	if err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).Scan(&PmsAppStay); err != nil {
		err = gerror.New("查询住宿单信息失败")
		return
	}
	if PmsAppStay.OrderStatus != "WAIT_PAY" {
		err = gerror.New("订单状态错误")
		return
	}
	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, PmsAppStay.MemberId).Scan(&MemberInfo); err != nil {
		err = gerror.New("查询用户信息失败")
		return
	}
	// 查询是否存在未支付的余额支付
	PlaceOrderLogger.Info(ctx, "查询是否存在未支付的余额支付")
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Scan(&PmsBalTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	contexts.SetMemberUser(ctx, &model.MemberIdentity{
		PmsMember: MemberInfo,
		App:       "",
		LoginAt:   nil,
	})
	// 是否存在未支付的余额支付信息
	if !g.IsEmpty(PmsBalTransaction) {
		// 扣除余额
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      MemberInfo.Id,
			Scene:         "HOTEL",
			Type:          "CONSUME",
			ChangeBalance: PmsBalTransaction.Amount * -1,
			OrderSn:       PmsBalTransaction.OrderSn,
			Reason:        "支付订房费用",
		}, tx); err != nil {
			// 退三方支付金额
			return
		}

		// 修改支付订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
			Where(dao.PmsTransaction.Columns().TransactionSn, PmsBalTransaction.TransactionSn).
			Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
			Update(entity.PmsTransaction{
				PayAmount: PmsBalTransaction.Amount,
				PayStatus: "DONE",
				PayTime:   gtime.Now(),
			}); err != nil {
			return
		}
	}

	// 获取AIRHOST下单参数
	if cookieBookerParams, err = service.HotelService().CreateStayParams(ctx, OrderSn); err != nil {
		return
	}
	// 提交订单
	if BookerResponse, err = service.HotelService().CreateStay(ctx, cookieBookerParams, tx); err != nil {
		return
	}
	// 修改住宿订单为支付完成
	if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).Update(g.Map{
		dao.PmsAppStay.Columns().OrderStatus: "HAVE_PAID",
		dao.PmsAppStay.Columns().Uuid:        BookerResponse.Data.ID,
		dao.PmsAppStay.Columns().Booker:      BookerResponse.Data.Booker.ID,
	}); err != nil {
		return
	}

	// 修改入住人ID
	if _, err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).Where(g.Map{
		dao.PmsGuestProfile.Columns().Uid: PmsAppStay.Booker,
	}).Data(g.Map{
		dao.PmsGuestProfile.Columns().Uid: BookerResponse.Data.Booker.ID,
	}).Update(); err != nil {
		return
	}

	// 修改房间订单为支付完成
	if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).Update(g.Map{
		dao.PmsAppReservation.Columns().OrderStatus: "HAVE_PAID",
		dao.PmsAppReservation.Columns().MainGuest:   BookerResponse.Data.Booker.ID,
	}); err != nil {
		return
	}

	if err = cache.Instance().Set(ctx, "PayOrderSn_"+OrderSn, "HAVE_PAID", gtime.M*10); err != nil {
		return
	}

	// 酒店订单支付成功
	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   "HAVE_PAID",
		Remark:      "订单支付",
		OperateType: "USER",
		OperateId:   PmsAppStay.MemberId,
	}); err != nil {
		return
	}

	// 发放下单奖励 优惠券和礼品券
	// 发放下单奖励 优惠券和礼品券
	awardSendMsg, _ := json.Marshal(g.Map{
		"type":        "AWARD",
		"id":          gvar.New(PmsAppStay.Id).Int(),
		"checkInDate": PmsAppStay.CheckInDate,
		"memberId":    PmsAppStay.MemberId,
	})
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderAward,
		DataByte:     awardSendMsg,
		Header:       nil,
	}); err != nil {
		g.Log().Error(ctx, "发送下单奖励MQ失败", err)
	}

	// 发送短信(队列)
	_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderRemind,
		DataByte: gjson.New(g.Map{
			"orderSn": PmsAppStay.OrderSn,
			"event":   "hotel_order_pay",
		}).MustToJson(),
		Header: nil,
	})

	// 发送到消息队列
	systemMessageTitle := map[string]string{
		"zh":    "订单支付成功",
		"en":    "Order payment successful",
		"ja":    "注文の支払いが完了しました",
		"ko":    "주문 결제 성공",
		"zh_CN": "訂單支付成功",
	}
	systemMessageContent := map[string]string{
		"zh":    PmsAppStay.OrderSn + "订单支付成功",
		"en":    "Order " + PmsAppStay.OrderSn + " payment successful",
		"ja":    "注文" + PmsAppStay.OrderSn + "の支払いが完了しました",
		"ko":    "주문 " + PmsAppStay.OrderSn + " 결제 성공",
		"zh_CN": PmsAppStay.OrderSn + "訂單支付成功",
	}
	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(PmsAppStay.MemberId).Value()
	var memberLanguage string
	if phoneArea.String() == "+86" {
		memberLanguage = "zh"
	} else if phoneArea.String() == "+81" {
		memberLanguage = "ja"
	} else if phoneArea.String() == "+82" {
		memberLanguage = "ko"
	} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
		memberLanguage = "zh_CN"
	} else {
		memberLanguage = "en"
	}
	pushData := g.MapStrAny{
		"type":    0,
		"orderSn": PmsAppStay.OrderSn,
	}
	pushDataJson, _ := json.Marshal(pushData)
	appPushData := g.MapStrStr{
		"type":  "2",
		"param": string(pushDataJson),
	}
	service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle:   systemMessageTitle,
		SystemMessageContent: systemMessageContent,
		Scene:                "hotel",
		Type:                 "order",
		MemberId:             int(PmsAppStay.MemberId),
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/order/order-detail",
		WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", PmsAppStay.OrderSn),
		EnablePush:           true,
		EnableSms:            false,
		PushTitle:            systemMessageTitle[memberLanguage],
		PushContent:          systemMessageContent[memberLanguage],
		OperatorId:           int(PmsAppStay.MemberId),
		OperatorRole:         "MEMBER",
		OrderSn:              PmsAppStay.OrderSn,
	})

	return
}
func HandleTravelOrderMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		PmsBalTransaction entity.PmsTransaction
		TravelOrder       *entity.TravelOrder
		MemberInfo        *entity.PmsMember
		tx                gdb.TX
		OrderSn           string
	)
	defer func() {
		if r := recover(); r != nil {
			PlaceOrderLogger.Error(ctx, r)
		}
	}()
	OrderSn = gvar.New(msg.Body).String()
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	// 查询住宿单信息
	if err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().OrderSn, OrderSn).Scan(&TravelOrder); err != nil {
		err = gerror.New("查询一日游订单失败")
		return
	}
	if TravelOrder.OrderStatus != "WAIT_PAY" {
		err = gerror.New("订单状态错误")
		return
	}
	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, TravelOrder.MemberId).Scan(&MemberInfo); err != nil {
		err = gerror.New("查询用户信息失败")
		return
	}
	// 查询是否存在未支付的余额支付
	PlaceOrderLogger.Info(ctx, "查询是否存在未支付的余额支付")
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Scan(&PmsBalTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	contexts.SetMemberUser(ctx, &model.MemberIdentity{
		PmsMember: MemberInfo,
		App:       "",
		LoginAt:   nil,
	})
	// 是否存在未支付的余额支付信息
	if !g.IsEmpty(PmsBalTransaction) {
		// 扣除余额
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      MemberInfo.Id,
			Scene:         "TRAVEL",
			Type:          "CONSUME",
			ChangeBalance: PmsBalTransaction.Amount * -1,
			OrderSn:       PmsBalTransaction.OrderSn,
			Reason:        "支付一日游费用",
		}, tx); err != nil {
			// 退三方支付金额
			return
		}

		// 修改支付订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
			Where(dao.PmsTransaction.Columns().TransactionSn, PmsBalTransaction.TransactionSn).
			Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
			Update(entity.PmsTransaction{
				PayAmount: PmsBalTransaction.Amount,
				PayStatus: "DONE",
				PayTime:   gtime.Now(),
			}); err != nil {
			return
		}
	}

	// 修改订单为支付完成
	if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().OrderSn, OrderSn).Update(g.Map{
		dao.TravelOrder.Columns().OrderStatus: "WAIT_VERIFY",
		dao.TravelOrder.Columns().PayTime:     gtime.Now().Format("Y-m-d H:i:s"),
		dao.TravelOrder.Columns().PayStatus:   "HAVE_PAID",
	}); err != nil {
		return
	}

	// 支付成功增加产品销量字段
	if _, err = dao.TravelProduct.Ctx(ctx).TX(tx).
		Where(dao.TravelProduct.Columns().Id, TravelOrder.ProductId).
		Increment(dao.TravelProduct.Columns().SalesNum, TravelOrder.BookingNum); err != nil {
		PlaceOrderLogger.Error(ctx, "增加产品销量失败", err)
		return
	}

	// 支付成功增加产品Sku销量字段
	if _, err = dao.TravelProductSku.Ctx(ctx).TX(tx).
		Where(dao.TravelProductSku.Columns().Id, TravelOrder.SkuId).
		Increment(dao.TravelProductSku.Columns().SalesNum, TravelOrder.BookingNum); err != nil {
		PlaceOrderLogger.Error(ctx, "增加产品Sku销量失败", err)
		return
	}

	if err = cache.Instance().Set(ctx, "PayOrderSn_"+OrderSn, "HAVE_PAID", gtime.M*10); err != nil {
		return
	}

	// 订单支付成功日志
	if _, err = dao.TravelOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.TravelOrderLog{
		OrderId:     int(TravelOrder.Id),
		ActionWay:   "HAVE_PAID",
		Remark:      "订单支付",
		OperateType: "USER",
		OrderStatus: "WAIT_VERIFY",
		OperateId:   int(TravelOrder.MemberId),
	}); err != nil {
		return
	}

	// 发送超时队列
	var TimeDuration int64
	// 获取订单对应产品的集合时间
	var TravelProduct *entity.TravelProduct
	if err = dao.TravelProduct.Ctx(ctx).WherePri(TravelOrder.ProductId).Scan(&TravelProduct); err != nil {
		return
	}

	// 解析集合时间，格式为 HH:MM
	meetingTimeStr := TravelProduct.MeetingTime
	var meetingTime *gtime.Time
	if meetingTimeStr != "" {
		// 构造预约日期的集合时间
		meetingDateTime := TravelOrder.BookDate.Format("Y-m-d") + " " + meetingTimeStr
		meetingTime = gtime.New(meetingDateTime)
	} else {
		// 如果没有集合时间，使用预约日期的开始时间
		meetingTime = gtime.New(TravelOrder.BookDate.Format("Y-m-d") + " 10:00:00")
	}

	// 延时发送时间：集合时间 + 最大等待时间(30分钟) - 当前时间
	maxWaitMinutes := 30 // 默认30分钟最大等待时间
	TimeDuration = meetingTime.Unix() + int64(maxWaitMinutes*60) - gtime.Now().Unix()

	// 确保TimeDuration不为负数
	if TimeDuration < 0 {
		TimeDuration = 0
	}
	if TimeDuration > 0 {
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeDelayedName,
			QueueName:    consts.RabbitMQQueueNameOrderExpire,
			DataByte:     gvar.New("Y-" + TravelOrder.OrderSn).Bytes(),
			Header: amqp.Table{
				"x-delay": gvar.New(TimeDuration * 1000).String(),
			},
		})
	} else {
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderExpire,
			DataByte:     gvar.New("Y-" + TravelOrder.OrderSn).Bytes(),
			Header:       nil,
		})
	}

	//// 发送到消息队列
	//systemMessageTitle := map[string]string{
	//	"zh":    "订单支付成功",
	//	"en":    "Order payment successful",
	//	"ja":    "注文の支払いが完了しました",
	//	"ko":    "주문 결제 성공",
	//	"zh_CN": "訂單支付成功",
	//}
	//systemMessageContent := map[string]string{
	//	"zh":    TravelOrder.OrderSn + "订单支付成功",
	//	"en":    "Order " + TravelOrder.OrderSn + " payment successful",
	//	"ja":    "注文" + TravelOrder.OrderSn + "の支払いが完了しました",
	//	"ko":    "주문 " + TravelOrder.OrderSn + " 결제 성공",
	//	"zh_CN": TravelOrder.OrderSn + "訂單支付成功",
	//}
	//// 查询用户的手机号区号 来判断用户语言
	//phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(TravelOrder.MemberId).Value()
	//var memberLanguage string
	//if phoneArea.String() == "+86" {
	//	memberLanguage = "zh"
	//} else if phoneArea.String() == "+81" {
	//	memberLanguage = "ja"
	//} else if phoneArea.String() == "+82" {
	//	memberLanguage = "ko"
	//} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
	//	memberLanguage = "zh_CN"
	//} else {
	//	memberLanguage = "en"
	//}
	//pushData := g.MapStrAny{
	//	"type":    0,
	//	"orderSn": TravelOrder.OrderSn,
	//}
	//pushDataJson, _ := json.Marshal(pushData)
	//appPushData := g.MapStrStr{
	//	"type":  "2",
	//	"param": string(pushDataJson),
	//}
	//service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
	//	SystemMessageTitle:   systemMessageTitle,
	//	SystemMessageContent: systemMessageContent,
	//	Scene:                "hotel",
	//	Type:                 "order",
	//	MemberId:             int(TravelOrder.MemberId),
	//	Language:             memberLanguage,
	//	AppPushData:          appPushData,
	//	AppLink:              "/order/order-detail",
	//	WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", TravelOrder.OrderSn),
	//	EnablePush:           true,
	//	EnableSms:            false,
	//	PushTitle:            systemMessageTitle[memberLanguage],
	//	PushContent:          systemMessageContent[memberLanguage],
	//	OperatorId:           int(TravelOrder.MemberId),
	//	OperatorRole:         "MEMBER",
	//	OrderSn:              TravelOrder.OrderSn,
	//})

	return
}
func HandleChangeMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		PmsAppReservationChange *entity.PmsAppReservationChange
		PmsAppReservation       *entity.PmsAppReservation
		BookerResponse          *airhousePublicApi.RetrieveRoomReservationJSONDataResponse
		tx                      gdb.TX
		OrderSn                 string
		ChangeType              string
		PmsRoomType             *entity.PmsRoomType
		OldCharges              []*entity.PmsCharge
		Availabilities          []*entity.PmsAvailabilities
		PmsPirceConfig          *model.PmsPriceConfig
	)
	OrderSn = gvar.New(msg.Body).String()
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	// 查询住宿变更单信息
	if err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Where(dao.PmsAppReservationChange.Columns().ChangeOrderSn, OrderSn).Scan(&PmsAppReservationChange); err != nil {
		err = gerror.New("查询住宿单信息失败")
		return
	}
	if PmsAppReservationChange.ChangeStatus != "ING" {
		err = gerror.New("住宿变更单状态错误")
		return
	}
	// 查询住宿单信息
	if err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Id, PmsAppReservationChange.OrderId).Scan(&PmsAppReservation); err != nil {
		return
	}
	if g.IsEmpty(PmsAppReservation) {
		err = gerror.New("住宿单不存在")
		return
	}

	OldCheckedInDate := PmsAppReservation.CheckinDate

	// 变更人数
	if PmsAppReservationChange.ChangeType == "PEOPLE" {
		if PmsAppReservation.IsChangePeople == "Y" {
			err = gerror.New("该订单已变更过人数")
			return
		}
		ChangeType = "CHANGE_PEOPLE_NUM"
		// 人数变更
		if _, err = airhousePublicApi.UpdateRoomReservationPost(ctx, PmsAppReservation.Uuid, g.Map{
			"adult_count":    PmsAppReservationChange.NewAdultCount,
			"child_count":    PmsAppReservationChange.NewChildCount,
			"infant_count":   PmsAppReservationChange.NewInfantCount,
			"prepaid_amount": PmsAppReservationChange.ChangeAmount + PmsAppReservation.BookingFee,
		}); err != nil {
			return
		}
		// 回写住宿单变更信息
		if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Id, PmsAppReservationChange.OrderId).Update(g.Map{
			dao.PmsAppReservation.Columns().IsChangePeople:   "Y",
			dao.PmsAppReservation.Columns().IsChangePeopleId: PmsAppReservationChange.Id,
			dao.PmsAppReservation.Columns().AdultCount:       PmsAppReservationChange.NewAdultCount,
			dao.PmsAppReservation.Columns().ChildCount:       PmsAppReservationChange.NewChildCount,
			dao.PmsAppReservation.Columns().InfantCount:      PmsAppReservationChange.NewInfantCount,
		}); err != nil {
			return
		}

		// 更新 Charge 表中的超员费
		// 查询房型信息
		if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, PmsAppReservation.RoomType).Scan(&PmsRoomType); err != nil {
			return
		}
		if g.IsEmpty(PmsRoomType) {
			err = gerror.New("房型不存在")
			return
		}

		// 查询原订单的 Charge 记录
		if err = dao.PmsCharge.Ctx(ctx).Where(dao.PmsCharge.Columns().Uid, PmsAppReservation.Charges).Scan(&OldCharges); err != nil {
			return
		}

		// 更新超员费
		if err = updateSurchargeForPeopleChange(ctx, tx, PmsAppReservation, PmsRoomType, OldCharges, PmsAppReservationChange); err != nil {
			return
		}
	}
	// 变更日期
	if PmsAppReservationChange.ChangeType == "DATE" {
		if PmsAppReservation.IsChangeDate == "Y" {
			err = gerror.New("该订单已变更过日期")
			return
		}
		ChangeType = "CHANGE_DATE"
		// 日期变更
		if BookerResponse, err = airhousePublicApi.UpdateRoomReservationPost(ctx, PmsAppReservation.Uuid, g.Map{
			"checkin_date":   PmsAppReservationChange.NewCheckinDate.Format("Y-m-d"),
			"checkout_date":  PmsAppReservationChange.NewCheckoutDate.Format("Y-m-d"),
			"booking_fee":    PmsAppReservationChange.ChangeAmount + PmsAppReservation.BookingFee,
			"prepaid_amount": PmsAppReservationChange.ChangeAmount + PmsAppReservation.BookingFee,
		}); err != nil {
			// 续住失败
			if err = cache.Instance().Set(ctx, "PayOrderSn_"+OrderSn, "HC_DATE_FAIL", gtime.S*60); err != nil {
				return
			}
			return
		}
		if g.IsEmpty(BookerResponse.Data.RoomUnit.ID) {
			// 续住失败
			if err = cache.Instance().Set(ctx, "PayOrderSn_"+OrderSn, "HC_DATE_FAIL", gtime.S*60); err != nil {
				return
			}
		}
		// 回写住宿单变更信息
		if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Id, PmsAppReservationChange.OrderId).Update(g.Map{
			dao.PmsAppReservation.Columns().IsChangeDate:   "Y",
			dao.PmsAppReservation.Columns().IsChangeDateId: PmsAppReservationChange.Id,
			dao.PmsAppReservation.Columns().CheckinDate:    PmsAppReservationChange.NewCheckinDate,
			dao.PmsAppReservation.Columns().CheckoutDate:   PmsAppReservationChange.NewCheckoutDate,
			dao.PmsAppReservation.Columns().BookingFee:     gdb.Raw(fmt.Sprintf("booking_fee+%f", PmsAppReservationChange.ChangeAmount)),
		}); err != nil {
			return
		}

		if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, PmsAppReservationChange.OrderSn).Update(g.Map{
			dao.PmsAppStay.Columns().CheckInDate:  PmsAppReservationChange.NewCheckinDate,
			dao.PmsAppStay.Columns().CheckOutDate: PmsAppReservationChange.NewCheckoutDate,
		}); err != nil {
			return
		}

		// 更新 Charge 表：删除不在新日期范围内的记录，添加新日期的记录
		// 查询房型信息
		if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, PmsAppReservation.RoomType).Scan(&PmsRoomType); err != nil {
			return
		}
		if g.IsEmpty(PmsRoomType) {
			err = gerror.New("房型不存在")
			return
		}

		// 查询原订单的 Charge 记录
		if err = dao.PmsCharge.Ctx(ctx).Where(dao.PmsCharge.Columns().Uid, PmsAppReservation.Charges).Scan(&OldCharges); err != nil {
			return
		}

		// 查询新日期范围的价格信息
		if err = dao.PmsAvailabilities.Ctx(ctx).
			Where(dao.PmsAvailabilities.Columns().Tuid, PmsAppReservation.RoomType).
			WhereGTE(dao.PmsAvailabilities.Columns().Date, PmsAppReservationChange.NewCheckinDate.Format("Y-m-d")).
			WhereLT(dao.PmsAvailabilities.Columns().Date, PmsAppReservationChange.NewCheckoutDate.Format("Y-m-d")).
			Scan(&Availabilities); err != nil {
			return
		}

		// 获取价格配置
		if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
			return
		}

		// 更新 Charge 表
		if err = updateChargesForDateChangeInQueue(ctx, tx, PmsAppReservation, PmsRoomType, OldCharges, Availabilities, PmsPirceConfig); err != nil {
			return
		}
	}
	// 修改住宿订单为支付完成
	if _, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Where(dao.PmsAppReservationChange.Columns().ChangeOrderSn, OrderSn).Update(g.Map{
		dao.PmsAppReservationChange.Columns().ChangeStatus: "DONE",
		dao.PmsAppReservationChange.Columns().DoneDate:     gtime.Now(),
	}); err != nil {
		return
	}
	// 更新主订单的总订单金额
	if err = service.PayService().PayAmountChange(ctx, tx, PmsAppReservationChange.OrderSn); err != nil {
		return
	}

	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: PmsAppReservationChange.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		err = gerror.New("该订单无需处理")
		return
	}
	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   ChangeType,
		Remark:      gvar.New(PmsAppReservationChange.Id).String(),
		OperateType: "USER",
		OperateId:   PmsAppStay.MemberId,
	}); err != nil {
		return
	}

	// 发送到消息队列
	systemMessageTitle := map[string]string{
		"zh":    "订单变更成功",
		"en":    "Order change successful",
		"ja":    "注文の変更が完了しました",
		"ko":    "주문 변경 성공",
		"zh_CN": "訂單變更成功",
	}
	systemMessageContent := map[string]string{
		"zh":    "订单入住时间成功变更为" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "。",
		"en":    "The check-in date has been successfully changed to " + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + ".",
		"ja":    "チェックイン時間は" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "に変更されました。",
		"ko":    "체크인 시간이 " + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "으로 변경되었습니다.",
		"zh_CN": "訂單入住時間成功變更為" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "。",
	}
	if PmsAppReservationChange.ChangeType == "PEOPLE" {
		systemMessageContent = map[string]string{
			"zh":    "订单入住人数变更成功。",
			"en":    "The number of guests in the order has been successfully changed.",
			"ja":    "注文内のゲストの人数が正常に変更されました。",
			"ko":    "주문에 포함된 손님 수가 성공적으로 변경되었습니다.",
			"zh_CN": "訂單入住人數變更成功。",
		}
	}

	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(PmsAppStay.MemberId).Value()
	var memberLanguage string
	if phoneArea.String() == "+86" {
		memberLanguage = "zh"
	} else if phoneArea.String() == "+81" {
		memberLanguage = "ja"
	} else if phoneArea.String() == "+82" {
		memberLanguage = "ko"
	} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
		memberLanguage = "zh_CN"
	} else {
		memberLanguage = "en"
	}
	pushData := g.MapStrAny{
		"type":    0,
		"orderSn": PmsAppStay.OrderSn,
	}
	pushDataJson, _ := json.Marshal(pushData)
	appPushData := g.MapStrStr{
		"type":  "2",
		"param": string(pushDataJson),
	}
	service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle:   systemMessageTitle,
		SystemMessageContent: systemMessageContent,
		Scene:                "hotel",
		Type:                 "order",
		MemberId:             int(PmsAppStay.MemberId),
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/order/order-detail",
		WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", PmsAppStay.OrderSn),
		EnablePush:           true,
		EnableSms:            false,
		PushTitle:            systemMessageTitle[memberLanguage],
		PushContent:          systemMessageContent[memberLanguage],
		OperatorId:           int(PmsAppStay.MemberId),
		OperatorRole:         "MEMBER",
		OrderSn:              PmsAppStay.OrderSn,
	})

	if gtime.New(OldCheckedInDate) != gtime.New(PmsAppReservationChange.NewCheckinDate) {

		// 失效
		invalidSendMsg, _ := json.Marshal(g.Map{
			"type": "INVALID",
			"id":   gvar.New(PmsAppStay.Id).Int(),
		})
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderAward,
			DataByte:     invalidSendMsg,
			Header:       nil,
		}); err != nil {
			g.Log().Error(ctx, "发送下单奖励失效MQ失败", err)
		}

		// 发放下单奖励 优惠券和礼品券
		awardSendMsg, _ := json.Marshal(g.Map{
			"type":        "AWARD",
			"id":          gvar.New(PmsAppStay.Id).Int(),
			"checkInDate": PmsAppReservationChange.NewCheckinDate.String(),
			"memberId":    PmsAppStay.MemberId,
		})
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderAward,
			DataByte:     awardSendMsg,
			Header:       nil,
		}); err != nil {
			g.Log().Error(ctx, "发送下单奖励MQ失败", err)
		}
	}

	return
}
func HandleCabinetOrderMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		PmsBalTransaction entity.PmsTransaction
		CabinetOrder      *entity.CabinetOrder
		MemberInfo        *entity.PmsMember
		tx                gdb.TX
		OrderSn           string
	)
	OrderSn = gvar.New(msg.Body).String()
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()

			if CabinetOrder.PayStep == "BASE" {
				// 请求mch那边下单接口，接口返回成功则继续，失败则进行退款
				var (
					cabinetRequest   *cabinetApi.CabinetCreateOrderParams
					cabinetResponse  *cabinetApi.CabinetCreateOrderResponse
					CabinetApiConfig *model.CabinetApiConfig
				)
				if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
					return
				}
				cabinetRequest = new(cabinetApi.CabinetCreateOrderParams)
				cabinetRequest.CabinetId = CabinetOrder.CabinetId
				cabinetRequest.BoxTypeId = CabinetOrder.BoxTypeId
				cabinetRequest.BuyHours = CabinetOrder.BuyHours
				cabinetRequest.OutTradeNo = CabinetOrder.OrderSn
				cabinetRequest.OutTradeVipid = strconv.Itoa(int(CabinetOrder.MemberId))
				if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).CreateOrder(ctx, cabinetRequest); err != nil {
				}
				if cabinetResponse.Code != 0 {
					// 执行退款
					err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
						// 执行退款
						// 计算退款金额
						var (
							Transaction       []*entity.PmsTransaction
							TransactionRefund []*entity.PmsTransactionRefund
							RefundBalance     float64 // 可退款积分
							RefundFee         float64 // 第三方支付
						)
						if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().OrderSn, CabinetOrder.OrderSn).Scan(&Transaction); err != nil {
							return
						}
						if err = dao.PmsTransactionRefund.Ctx(ctx).Where(dao.PmsTransactionRefund.Columns().OrderSn, CabinetOrder.OrderSn).Scan(&TransactionRefund); err != nil {
							return
						}

						for _, v := range Transaction {
							// 可退款金额
							Refundable := v.Amount - v.RefundAmount
							if v.PayType == "BAL" {
								RefundBalance = RefundBalance + Refundable
							} else if v.PayType == "COUPON" {

							} else {
								RefundFee = RefundFee + Refundable
							}
						}
						// 退款金额
						RefundAmount := RefundBalance + RefundFee

						// 修改订单状态
						RefundStatus := "DONE"
						if CabinetOrder.OrderAmount > RefundAmount {
							RefundStatus = "PART"
						}
						if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
							WherePri(CabinetOrder.Id).Data(g.MapStrAny{
							dao.CabinetOrder.Columns().RefundAmount:       RefundAmount,
							dao.CabinetOrder.Columns().RefundBalAmount:    RefundBalance,
							dao.CabinetOrder.Columns().RefundCouponAmount: 0,
							dao.CabinetOrder.Columns().RefundStatus:       RefundStatus,
							dao.CabinetOrder.Columns().RefundTime:         gtime.Now(),
							dao.CabinetOrder.Columns().OrderStatus:        "CANCEL",
							dao.CabinetOrder.Columns().PayStatus:          "REFUND",
							dao.CabinetOrder.Columns().CancelTime:         gtime.Now(),
						}).Update(); err != nil {
							return
						}

						// 订单日志
						if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
							OrderId:     int(CabinetOrder.Id),
							OrderStatus: "CANCEL",
							ActionWay:   "REFUND",
							Remark:      "订单已退款",
							OperateType: "SYSTEM",
						}); err != nil {
							return
						}

						// 全额退款
						err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
							OrderSn:      CabinetOrder.OrderSn,
							RefundAmount: RefundAmount,
							OperateType:  "SYSTEM",
						}, tx)
						if err != nil {
							return
						}

						// 删除支付成功cache
						if _, err = cache.Instance().Remove(ctx, fmt.Sprintf("PayOrderSn_%s", string(msg.Body))); err != nil {
							return
						}

						return
					})

					if err != nil {
						if _, err = dao.CabinetOrder.Ctx(ctx).
							WherePri(CabinetOrder.Id).Data(g.MapStrAny{
							dao.CabinetOrder.Columns().IsAbnormal: 1,
						}).Update(); err != nil {
						}
					}
				} else {
					AddressJson := &input_cabinet.LanguageJson{
						Zh: cabinetResponse.Data.AddressZh,
						En: cabinetResponse.Data.AddressEn,
						Ja: cabinetResponse.Data.AddressJa,
						Ko: cabinetResponse.Data.AddressKo,
						Tw: cabinetResponse.Data.AddressTw,
					}
					if _, err = dao.CabinetOrder.Ctx(ctx).
						WherePri(CabinetOrder.Id).Data(g.MapStrAny{
						dao.CabinetOrder.Columns().OutOrderSn:   cabinetResponse.Data.OrderNo,
						dao.CabinetOrder.Columns().BoxId:        cabinetResponse.Data.BoxId,
						dao.CabinetOrder.Columns().BoxNo:        cabinetResponse.Data.BoxNo,
						dao.CabinetOrder.Columns().BoxAlias:     cabinetResponse.Data.BoxAlias,
						dao.CabinetOrder.Columns().Pin:          cabinetResponse.Data.Pin,
						dao.CabinetOrder.Columns().Address:      cabinetResponse.Data.AddressZh,
						dao.CabinetOrder.Columns().AddressJson:  gjson.New(AddressJson),
						dao.CabinetOrder.Columns().StartTime:    cabinetResponse.Data.StartTime,
						dao.CabinetOrder.Columns().EndTime:      cabinetResponse.Data.EndTime,
						dao.CabinetOrder.Columns().GraceSeconds: cabinetResponse.Data.GraceSeconds,
						dao.CabinetOrder.Columns().OrderStatus:  "ING",
					}).Update(); err != nil {

					}

					// 发送到消息队列
					systemMessageTitle := map[string]string{
						"zh":    "订单支付成功",
						"en":    "Order payment successful",
						"ja":    "注文の支払いが完了しました",
						"ko":    "주문 결제 성공",
						"zh_CN": "訂單支付成功",
					}
					systemMessageContent := map[string]string{
						"zh":    CabinetOrder.OrderSn + "订单支付成功",
						"en":    "Order " + CabinetOrder.OrderSn + " payment successful",
						"ja":    "注文" + CabinetOrder.OrderSn + "の支払いが完了しました",
						"ko":    "주문 " + CabinetOrder.OrderSn + " 결제 성공",
						"zh_CN": CabinetOrder.OrderSn + "訂單支付成功",
					}
					// 查询用户的手机号区号 来判断用户语言
					phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(CabinetOrder.MemberId).Value()
					var memberLanguage string
					if phoneArea.String() == "+86" {
						memberLanguage = "zh"
					} else if phoneArea.String() == "+81" {
						memberLanguage = "ja"
					} else if phoneArea.String() == "+82" {
						memberLanguage = "ko"
					} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
						memberLanguage = "zh_CN"
					} else {
						memberLanguage = "en"
					}
					appPushData := g.MapStrStr{
						"type":   "1",
						"string": CabinetOrder.OrderSn,
					}
					service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
						SystemMessageTitle:   systemMessageTitle,
						SystemMessageContent: systemMessageContent,
						Scene:                "cabinet",
						Type:                 "order",
						MemberId:             int(CabinetOrder.MemberId),
						Language:             memberLanguage,
						AppPushData:          appPushData,
						AppLink:              "/smart_locker_order_detail",
						WxLink:               fmt.Sprintf("/subpackages/lockers/pages/order-detail?orderSn=%s", CabinetOrder.OrderSn),
						EnablePush:           true,
						EnableSms:            false,
						PushTitle:            systemMessageTitle[memberLanguage],
						PushContent:          systemMessageContent[memberLanguage],
						OperatorId:           int(CabinetOrder.MemberId),
						OperatorRole:         "MEMBER",
						OrderSn:              CabinetOrder.OrderSn,
					})
				}
			} else {
				// 请求mch超时支付成功
				var (
					cabinetRequest   *cabinetApi.CabinetPayOvertimeParams
					cabinetResponse  *cabinetApi.CabinetPayOvertimeResponse
					CabinetApiConfig *model.CabinetApiConfig
				)
				if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
					return
				}
				cabinetRequest = new(cabinetApi.CabinetPayOvertimeParams)
				cabinetRequest.OutTradeNo = CabinetOrder.OrderSn
				if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).PayOvertime(ctx, cabinetRequest); err != nil {
				}
				if _, err = dao.CabinetOrder.Ctx(ctx).
					WherePri(CabinetOrder.Id).Data(g.MapStrAny{
					dao.CabinetOrder.Columns().OrderStatus: "GRACE",
				}).Update(); err != nil {

				}
				if cabinetResponse.Code == 0 {
					if _, err = dao.CabinetOrder.Ctx(ctx).
						WherePri(CabinetOrder.Id).Data(g.MapStrAny{
						dao.CabinetOrder.Columns().GraceEndTime: cabinetResponse.Data.GraceEndTime,
					}).Update(); err != nil {
						return
					}

					// 发送到消息队列
					systemMessageTitle := map[string]string{
						"zh":    "超时费支付成功",
						"en":    "Overtime fee paid successfully",
						"ja":    "残業代は正常に支払われました",
						"ko":    "초과 근무 수당이 성공적으로 지불되었습니다.",
						"zh_CN": "超時費支付成功",
					}
					systemMessageContent := map[string]string{
						"zh":    CabinetOrder.OrderSn + "订单超时费支付成功",
						"en":    "Order " + CabinetOrder.OrderSn + " Overtime Fee Payment Successful",
						"ja":    "注文番号" + CabinetOrder.OrderSn + " 残業料金の支払いが成功しました",
						"ko":    "주문 " + CabinetOrder.OrderSn + " 초과 근무 수수료 지불 성공",
						"zh_CN": CabinetOrder.OrderSn + "訂單超時費支付成功",
					}
					// 查询用户的手机号区号 来判断用户语言
					phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(CabinetOrder.MemberId).Value()
					var memberLanguage string
					if phoneArea.String() == "+86" {
						memberLanguage = "zh"
					} else if phoneArea.String() == "+81" {
						memberLanguage = "ja"
					} else if phoneArea.String() == "+82" {
						memberLanguage = "ko"
					} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
						memberLanguage = "zh_CN"
					} else {
						memberLanguage = "en"
					}
					appPushData := g.MapStrStr{
						"type":   "1",
						"string": CabinetOrder.OrderSn,
					}
					service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
						SystemMessageTitle:   systemMessageTitle,
						SystemMessageContent: systemMessageContent,
						Scene:                "cabinet",
						Type:                 "order",
						MemberId:             int(CabinetOrder.MemberId),
						Language:             memberLanguage,
						AppPushData:          appPushData,
						AppLink:              "/smart_locker_order_detail",
						WxLink:               fmt.Sprintf("/subpackages/lockers/pages/order-detail?orderSn=%s", CabinetOrder.OrderSn),
						EnablePush:           true,
						EnableSms:            false,
						PushTitle:            systemMessageTitle[memberLanguage],
						PushContent:          systemMessageContent[memberLanguage],
						OperatorId:           int(CabinetOrder.MemberId),
						OperatorRole:         "MEMBER",
						OrderSn:              CabinetOrder.OrderSn,
					})
				} else {
					if _, err = dao.CabinetOrder.Ctx(ctx).
						WherePri(CabinetOrder.Id).Data(g.MapStrAny{
						dao.CabinetOrder.Columns().PayOvertimeAbnormal: 1,
					}).Update(); err != nil {
					}
				}
			}

		}
	}()
	// 查询订单信息
	if err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, OrderSn).Scan(&CabinetOrder); err != nil {
		err = gerror.New("查询订单信息失败")
		return
	}
	if CabinetOrder.PayStep == "BASE" && CabinetOrder.PayStatus != "WAIT_PAY" {
		err = gerror.New("订单状态错误")
		return
	}
	if CabinetOrder.PayStep == "OVERTIME" && CabinetOrder.OvertimePayStatus != "WAIT_PAY" {
		err = gerror.New("订单状态错误")
		return
	}
	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, CabinetOrder.MemberId).Scan(&MemberInfo); err != nil {
		err = gerror.New("查询用户信息失败")
		return
	}
	// 查询是否存在未支付的余额支付
	PlaceOrderLogger.Info(ctx, "查询是否存在未支付的余额支付")
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Scan(&PmsBalTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	contexts.SetMemberUser(ctx, &model.MemberIdentity{
		PmsMember: MemberInfo,
		App:       "",
		LoginAt:   nil,
	})
	// 是否存在未支付的余额支付信息
	if !g.IsEmpty(PmsBalTransaction) {
		// 扣除余额
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      MemberInfo.Id,
			Scene:         "CABINET",
			Type:          "CONSUME",
			ChangeBalance: PmsBalTransaction.Amount * -1,
			OrderSn:       PmsBalTransaction.OrderSn,
			Reason:        "支付储物柜费用",
		}, tx); err != nil {
			// 退三方支付金额
			return
		}

		// 修改支付订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
			Where(dao.PmsTransaction.Columns().TransactionSn, PmsBalTransaction.TransactionSn).
			Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
			Update(entity.PmsTransaction{
				PayAmount: PmsBalTransaction.Amount,
				PayStatus: "DONE",
				PayTime:   gtime.Now(),
			}); err != nil {
			return
		}
	}

	if CabinetOrder.PayStep == "BASE" {
		// 修改订单为支付完成
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, OrderSn).Update(g.Map{
			dao.CabinetOrder.Columns().PayTime:     gtime.Now().Format("Y-m-d H:i:s"),
			dao.CabinetOrder.Columns().PayStatus:   "HAVE_PAID",
			dao.CabinetOrder.Columns().OrderStatus: "HAVE_PAID",
		}); err != nil {
			return
		}

		// 订单支付成功日志
		if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
			OrderId:     int(CabinetOrder.Id),
			OrderStatus: "HAVE_PAID",
			ActionWay:   "HAVE_PAID",
			Remark:      "订单支付",
			OperateType: "USER",
			OperateId:   int(CabinetOrder.MemberId),
		}); err != nil {
			return
		}
	} else {
		// 超时费支付成功
		// 修改订单为支付完成
		OrderAmount := decimal.NewFromFloat(CabinetOrder.OrderAmount).Add(decimal.NewFromFloat(float64(CabinetOrder.OvertimeFee))).Round(2)
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, OrderSn).Update(g.Map{
			dao.CabinetOrder.Columns().OvertimePayTime:   gtime.Now().Format("Y-m-d H:i:s"),
			dao.CabinetOrder.Columns().OvertimePayStatus: "HAVE_PAID",
			dao.CabinetOrder.Columns().OrderAmount:       OrderAmount,
		}); err != nil {
			return
		}

		// 订单支付成功日志
		if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
			OrderId:     int(CabinetOrder.Id),
			OrderStatus: "HAVE_PAID",
			ActionWay:   "HAVE_PAID",
			Remark:      "订单超时费支付",
			OperateType: "USER",
			OperateId:   int(CabinetOrder.MemberId),
		}); err != nil {
			return
		}
	}

	return
}

// updateChargesForDateChangeInQueue 更新 Charge 表：删除不在新日期范围内的记录，添加新日期的记录
func updateChargesForDateChangeInQueue(
	ctx context.Context,
	tx gdb.TX,
	pmsAppReservation *entity.PmsAppReservation,
	pmsRoomType *entity.PmsRoomType,
	oldCharges []*entity.PmsCharge,
	availabilities []*entity.PmsAvailabilities,
	priceConfig *model.PmsPriceConfig,
) (err error) {
	// 构建新日期范围的日期集合
	newDateSet := make(map[string]bool)
	for _, v := range availabilities {
		newDateSet[v.Date] = true
	}

	// 构建旧日期的 Charge 记录映射（按日期分组）
	oldChargeDateSet := make(map[string]bool)
	for _, charge := range oldCharges {
		oldChargeDateSet[charge.Date.Format("Y-m-d")] = true
	}

	// 1. 删除不在新日期范围内的 Charge 记录
	for _, charge := range oldCharges {
		chargeDate := charge.Date.Format("Y-m-d")
		if !newDateSet[chargeDate] {
			// 该日期不在新日期范围内，删除
			if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).Where(dao.PmsCharge.Columns().Id, charge.Id).Delete(); err != nil {
				return
			}
		}
	}

	// 2. 添加新日期的 Charge 记录（不在旧日期范围内的）
	pricePercent := decimal.NewFromInt(int64(priceConfig.PricePercent)).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	totalGuests := pmsAppReservation.AdultCount + pmsAppReservation.ChildCount

	for _, availability := range availabilities {
		if !oldChargeDateSet[availability.Date] {
			// 该日期是新增的，需要添加 Charge 记录
			// 基础费
			basePrice := decimal.NewFromFloat(availability.Price).Mul(pricePercent).Round(0).InexactFloat64()
			if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsCharge{
				Uid:            pmsAppReservation.Charges,
				AirUid:         guid.S([]byte("AirUid")),
				Date:           gtime.New(availability.Date),
				Name:           consts.Charges["booking_fee"],
				FeeType:        "booking_fee",
				Amount:         basePrice,
				OriginalAmount: basePrice,
				Description:    "每日费用",
			}).Insert(); err != nil {
				return
			}

			// 超员费（如果有）
			if totalGuests > pmsRoomType.OccupantsForBaseRate {
				extraGuests := totalGuests - pmsRoomType.OccupantsForBaseRate
				extraFee := decimal.NewFromInt(int64(extraGuests)).Mul(decimal.NewFromFloat(pmsRoomType.AdditionalGuestAmounts)).InexactFloat64()
				if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsCharge{
					Uid:            pmsAppReservation.Charges,
					AirUid:         guid.S([]byte("AirUid")),
					Date:           gtime.New(availability.Date),
					Name:           consts.Charges["booking_fee_people"],
					FeeType:        "booking_fee_people",
					Amount:         extraFee,
					OriginalAmount: extraFee,
					Description:    "超员费",
				}).Insert(); err != nil {
					return
				}
			}
		}
	}

	return nil
}

// updateSurchargeForPeopleChange 更新 Charge 表中的超员费（人数变更时）
func updateSurchargeForPeopleChange(
	ctx context.Context,
	tx gdb.TX,
	pmsAppReservation *entity.PmsAppReservation,
	pmsRoomType *entity.PmsRoomType,
	oldCharges []*entity.PmsCharge,
	changeInfo *entity.PmsAppReservationChange,
) (err error) {
	// 1. 删除原有的超员费记录
	for _, charge := range oldCharges {
		if charge.FeeType == "booking_fee_people" {
			if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).Where(dao.PmsCharge.Columns().Id, charge.Id).Delete(); err != nil {
				return
			}
		}
	}

	// 2. 计算新的超员费并添加
	newTotalGuests := changeInfo.NewAdultCount + changeInfo.NewChildCount
	if newTotalGuests > pmsRoomType.OccupantsForBaseRate {
		extraGuests := newTotalGuests - pmsRoomType.OccupantsForBaseRate
		extraFee := decimal.NewFromInt(int64(extraGuests)).Mul(decimal.NewFromFloat(pmsRoomType.AdditionalGuestAmounts)).InexactFloat64()

		// 获取所有日期（从原有的 booking_fee 记录中提取）
		for _, charge := range oldCharges {
			if charge.FeeType == "booking_fee" {
				if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsCharge{
					Uid:            pmsAppReservation.Charges,
					AirUid:         guid.S([]byte("AirUid")),
					Date:           charge.Date,
					Name:           consts.Charges["booking_fee_people"],
					FeeType:        "booking_fee_people",
					Amount:         extraFee,
					OriginalAmount: extraFee,
					Description:    "超员费",
				}).Insert(); err != nil {
					return
				}
			}
		}
	}

	return nil
}
