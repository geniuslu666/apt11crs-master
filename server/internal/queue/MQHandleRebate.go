package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	_ "APT/internal/logic"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shopspring/decimal"
)

var (
	RebateLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameRebate)
)

func Rebate() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameRebate
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(RebateLogger)
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
	RebateLogger.Info(ctx, "Rebate Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		RebateLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		RebateLogger.Info(ctx, "--[start]----------------------------------------------")
		RebateLogger.Info(ctx, msg.Body)
		if !g.IsEmpty(msg.Body) {
			if string(msg.Body)[:1] == "H" {
				if err = HandleReturnBrokerageMQ(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
				if err = HandleHotelReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
				if err = HandleInviteAward(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "F" {
				if err = HandleFoodReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "C" {
				if err = HandleCarReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "S" {
				if err = HandleSpaReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}

			if string(msg.Body)[:1] == "B" {
				if err = HandleCabinetReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}

			if string(msg.Body)[:1] == "T" {
				if err = HandleTravelReturnScore(ctx, msg); err != nil {
					RebateLogger.Error(ctx, err.Error())
				}
			}
		}

		_ = msg.Ack(false)
	}
ERR:
	RebateLogger.Error(ctx, err)
	panic(err)
}

// HandleReturnBrokerageMQ 用户下单推荐人返佣金
func HandleReturnBrokerageMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body             = msg.Body
		OrderSn          string
		StayInfo         *entity.PmsAppStay
		CheckStayInfo    *entity.PmsAppStay
		checkinStatusVar *gvar.Var
		ReferrerInfo     *entity.PmsMember
		MemberInfo       *entity.PmsMember
		// ChannelInfo      *entity.PmsChannel
		// StaffInfo        *entity.PmsStaff
		// YYConfig         *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
			Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
			Scan(&StayInfo); err != nil {
			return
		}
		if g.IsEmpty(StayInfo) {
			err = gerror.New("住宿订单数据为空")
			return
		}
		if StayInfo.RebateStatus != "WAIT" {
			err = gerror.New("订单状态不为佣金处理中")
			return
		}
		if g.IsEmpty(StayInfo.Referrer) {
			err = gerror.New("没有推荐人无需处理佣金")
			return
		}
		// 查询下单人信息
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, StayInfo.MemberId).Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(MemberInfo) {
			err = gerror.New("用户已注销")
			return
		}
		// 判断是否用户首单
		if err = dao.PmsAppStay.Ctx(ctx).
			Where(dao.PmsAppStay.Columns().MemberId, StayInfo.MemberId).
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			OrderAsc(dao.PmsAppStay.Columns().Id).
			Scan(&CheckStayInfo); err != nil {
			return
		}
		if g.IsEmpty(CheckStayInfo) {
			RebateLogger.Info(ctx, "无订单不需要处理")
		}
		// if CheckStayInfo.OrderSn == OrderSn && MemberInfo.Referrer != StayInfo.Referrer {
		// 	// 首单第一次推荐人是否存在
		// 	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, StayInfo.Referrer).Scan(&ReferrerInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 		return
		// 	}
		// 	if g.IsEmpty(ReferrerInfo) {
		// 		err = gerror.New("推荐人不存在")
		// 		return
		// 	}
		// 	// 是首单的情况下  结算给初次推荐人 并更新
		// 	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		// 		return err
		// 	}
		// 	// 计算新推荐人的佣金比例
		// 	if !g.IsEmpty(ReferrerInfo) {
		// 		PmsAppStayData := &entity.PmsAppStay{
		// 			RebateStatus: "WAIT",
		// 			RebateRate:   0,
		// 		}
		// 		PmsAppStayData.RebateStatus = "WAIT"
		// 		if ReferrerInfo.RebateMode == "CHANNEL" {
		// 			if err = dao.PmsChannel.Ctx(ctx).WherePri(ReferrerInfo.ChannelId).Scan(&ChannelInfo); err != nil {
		// 				return err
		// 			}
		// 			if !g.IsEmpty(ChannelInfo) {
		// 				PmsAppStayData.RebateRate = ChannelInfo.Rate
		// 			}
		// 		} else if ReferrerInfo.RebateMode == "STAFF" {
		// 			if err = dao.PmsStaff.Ctx(ctx).WherePri(ReferrerInfo.StaffId).Scan(&StaffInfo); err != nil {
		// 				return err
		// 			}
		// 			if !g.IsEmpty(StaffInfo) {
		// 				PmsAppStayData.RebateRate = StaffInfo.Rate
		// 			}
		// 		} else if ReferrerInfo.RebateMode == "MEMBER" {
		// 			PmsAppStayData.RebateRate = YYConfig.MemberBrokerageRate
		// 		} else {
		// 			return gerror.New("不存在该返佣模式")
		// 		}

		// 		// 更新订单返佣信息
		// 		_, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).Data(PmsAppStayData).Update()
		// 	}
		// 	// 重新查询订单信息
		// 	// 查询 app 订单信息
		// 	if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
		// 		Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
		// 		Scan(&StayInfo); err != nil {
		// 		return
		// 	}
		// 	if g.IsEmpty(StayInfo) {
		// 		err = gerror.New("住宿订单数据为空")
		// 		return
		// 	}
		// }
		if checkinStatusVar, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
			Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).
			Group(dao.PmsAppReservation.Columns().CheckinStatus).
			Fields("GROUP_CONCAT(checkin_status) as checkin_status").
			Value(); err != nil {
			return
		}
		if checkinStatusVar.String() != "checked_out" {
			err = gerror.New("订单未完结无法计算佣金")
			return
		}
		RebateLogger.Info(ctx, "结算佣金")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间转换导致的精度损失
		netAmount := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		RebatePrice := netAmount.
			Mul(decimal.NewFromFloat(StayInfo.RebateRate)).
			Div(decimal.NewFromFloat(100)).
			Round(0). // 四舍五入到整数
			InexactFloat64()
		if RebatePrice == 0 {
			err = gerror.New("订单金额不足无法进行结算佣金")
			return
		}
		if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Data(g.Map{
			dao.PmsAppStay.Columns().RebateStatus: "SUCCESS",
			dao.PmsAppStay.Columns().RebateAmount: RebatePrice,
			dao.PmsAppStay.Columns().RebateTime:   gtime.Now(),
		}).Where(dao.PmsAppStay.Columns().Id, StayInfo.Id).Update(); err != nil {
			return
		}
		// 查看推荐人信息
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, StayInfo.Referrer).Scan(&ReferrerInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(ReferrerInfo) {
			err = gerror.New("推荐人不存在无法结算佣金")
			return
		}
		if ReferrerInfo.RebateMode == "STAFF" {
			// 增加结算账户佣金
			if _, err = dao.PmsStaff.Ctx(ctx).TX(tx).
				Where(dao.PmsStaff.Columns().Id, ReferrerInfo.StaffId).
				Increment(dao.PmsStaff.Columns().Balance, RebatePrice); err != nil {
				return
			}
			// 增加历史账户佣金
			if _, err = dao.PmsStaff.Ctx(ctx).TX(tx).
				Where(dao.PmsStaff.Columns().Id, ReferrerInfo.StaffId).
				Increment(dao.PmsStaff.Columns().AllBalance, RebatePrice); err != nil {
				return
			}
			// 增加结算账户佣金变动记录
			if _, err = dao.PmsBrokerage.Ctx(ctx).Data(&entity.PmsBrokerage{
				Identity:  "STAFF",
				Type:      "REBATE",
				Balance:   RebatePrice,
				StaffId:   ReferrerInfo.StaffId,
				ChannelId: 0,
				MemberId:  ReferrerInfo.Id,
			}).OmitEmptyData().InsertAndGetId(); err != nil {
				return
			}
		}
		if ReferrerInfo.RebateMode == "CHANNEL" {
			// 增加结算账户佣金
			if _, err = dao.PmsChannel.Ctx(ctx).TX(tx).Where(dao.PmsChannel.Columns().Id, ReferrerInfo.ChannelId).Increment(dao.PmsChannel.Columns().Balance, RebatePrice); err != nil {
				return
			}
			// 增加结算账户佣金
			if _, err = dao.PmsChannel.Ctx(ctx).TX(tx).Where(dao.PmsChannel.Columns().Id, ReferrerInfo.ChannelId).Increment(dao.PmsChannel.Columns().AllBalance, RebatePrice); err != nil {
				return
			}
			// 增加结算账户佣金变动记录
			if _, err = dao.PmsBrokerage.Ctx(ctx).Data(&entity.PmsBrokerage{
				Identity:  "CHANNEL",
				Type:      "REBATE",
				Balance:   RebatePrice,
				StaffId:   0,
				ChannelId: ReferrerInfo.ChannelId,
				MemberId:  ReferrerInfo.Id,
			}).OmitEmptyData().InsertAndGetId(); err != nil {
				return
			}
		}
		if ReferrerInfo.RebateMode == "MEMBER" {
			if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
				MemberId:      ReferrerInfo.Id,
				Scene:         "HOTEL",
				Type:          "AWARD",
				ChangeBalance: RebatePrice,
				OrderSn:       StayInfo.OrderSn,
				Reason:        "推荐人下单返现",
			}, tx); err != nil {
				return err
			}
		}

		if RebatePrice > 0 {
			// 发消息
			var systemMessageTitle map[string]string
			systemMessageTitle = map[string]string{
				"zh":    "推荐人下单奖励",
				"en":    "Referrer Order Reward",
				"ja":    "紹介者注文報酬",
				"ko":    "추천인 주문 보상",
				"zh_CN": "推薦人下單獎勵",
			}
			if ReferrerInfo.RebateMode == "MEMBER" {
				systemMessageContent := map[string]string{
					"zh":    fmt.Sprintf("您的推荐人下单奖励您 %d 积分", gvar.New(RebatePrice).Int()),
					"en":    fmt.Sprintf("You earned %d points from a referral order", gvar.New(RebatePrice).Int()),
					"ja":    fmt.Sprintf("紹介注文で %d ポイント獲得", gvar.New(RebatePrice).Int()),
					"ko":    fmt.Sprintf("추천인 주문으로 %d포인트 획득", gvar.New(RebatePrice).Int()),
					"zh_CN": fmt.Sprintf("推薦人下單成功，獲得 %d 積分", gvar.New(RebatePrice).Int()),
				}
				appPushData := g.MapStrStr{
					"type": "0",
				}

				service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
					SystemMessageTitle:   systemMessageTitle,
					SystemMessageContent: systemMessageContent,
					Scene:                "system",
					Type:                 "im",
					MemberId:             ReferrerInfo.Id,
					Language:             contexts.GetLanguage(ctx),
					AppPushData:          appPushData,
					AppLink:              "/account/score",
					WxLink:               "/pages/user/points",
					EnablePush:           true,
					EnableSms:            false,
					PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
					PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
					ShowIndex:            true,
					OperatorId:           0,
					OperatorRole:         "SYSTEM",
				})
			} else {
				systemMessageContent := map[string]string{
					"zh":    fmt.Sprintf("您的推荐人下单奖励您 %d 佣金", gvar.New(RebatePrice).Int()),
					"en":    fmt.Sprintf("You earned %d commission from a referral order", gvar.New(RebatePrice).Int()),
					"ja":    fmt.Sprintf("紹介注文で %d の手数料を獲得", gvar.New(RebatePrice).Int()),
					"ko":    fmt.Sprintf("추천인 주문으로 %d 커미션 획득", gvar.New(RebatePrice).Int()),
					"zh_CN": fmt.Sprintf("推薦人下單成功，獲得 %d 佣金", gvar.New(RebatePrice).Int()),
				}
				appPushData := g.MapStrStr{
					"type": "0",
				}

				service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
					SystemMessageTitle:   systemMessageTitle,
					SystemMessageContent: systemMessageContent,
					Scene:                "system",
					Type:                 "im",
					MemberId:             ReferrerInfo.Id,
					Language:             contexts.GetLanguage(ctx),
					AppPushData:          appPushData,
					AppLink:              "/distribution/commission-rebate",
					WxLink:               "/pages/distribution/order",
					EnablePush:           true,
					EnableSms:            false,
					PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
					PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
					ShowIndex:            true,
					OperatorId:           0,
					OperatorRole:         "SYSTEM",
				})
			}
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnBrokerage")
		return err
	}
	return
}

// HandleHotelReturnScore 用户下单返积分
func HandleHotelReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body             = msg.Body
		OrderSn          string
		StayInfo         *entity.PmsAppStay
		YYConfig         *model.YYConfig
		PayAmount        float64
		RefundAmount     float64
		checkinStatusVar *gvar.Var
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
			Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
			Scan(&StayInfo); err != nil {
			return
		}
		if g.IsEmpty(StayInfo) {
			err = gerror.New("住宿订单数据为空")
			return
		}

		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, StayInfo.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}

		if StayInfo.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if StayInfo.HotelGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")

		if checkinStatusVar, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
			Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).
			Group(dao.PmsAppReservation.Columns().CheckinStatus).
			Fields("GROUP_CONCAT(checkin_status) as checkin_status").
			Value(); err != nil {
			return
		}
		if checkinStatusVar.String() != "checked_out" {
			err = gerror.New("订单未完结")
			return
		}

		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		HotelGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "HotelGetAmountDecimal:%f", HotelGetAmountDecimal.InexactFloat64())
		HotelGetAmount := HotelGetAmountDecimal.
			Mul(decimal.NewFromFloat(StayInfo.HotelGetRateScene).Div(decimal.NewFromFloat(100))).
			Mul(decimal.NewFromFloat(StayInfo.HotelGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0). // 最后四舍五入到整数
			InexactFloat64()
		g.Log().Debugf(ctx, "HotelGetAmount:%f", HotelGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      StayInfo.MemberId,
			Scene:         "HOTEL",
			Type:          "AWARD",
			ChangeBalance: HotelGetAmount,
			OrderSn:       StayInfo.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).WherePri(StayInfo.Id).Data(g.Map{
			dao.PmsAppStay.Columns().HotelGetScoreStatus: "SUCCESS",
			dao.PmsAppStay.Columns().HotelGetAmount:      HotelGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if HotelGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("民宿订单下单成功奖励您 %d 积分", gvar.New(HotelGetAmount).Int()),
				"en":    fmt.Sprintf("Homestay order placed: +%d points", gvar.New(HotelGetAmount).Int()),
				"ja":    fmt.Sprintf("民宿注文成功：+%d ポイント", gvar.New(HotelGetAmount).Int()),
				"ko":    fmt.Sprintf("숙소 주문 완료: +%d포인트", gvar.New(HotelGetAmount).Int()),
				"zh_CN": fmt.Sprintf("民宿訂單下單成功，+%d 積分", gvar.New(HotelGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             StayInfo.MemberId,
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleSpaReturnScore 用户下单返积分
func HandleSpaReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		OrderSn      string
		SpaOrder     *entity.SpaOrder
		YYConfig     *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.SpaOrder.Ctx(ctx).TX(tx).
			Where(dao.SpaOrder.Columns().OrderSn, OrderSn).
			Scan(&SpaOrder); err != nil {
			return
		}
		if g.IsEmpty(SpaOrder) {
			err = gerror.New("按摩订单数据为空")
			return
		}

		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, SpaOrder.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}

		if SpaOrder.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if SpaOrder.SpaGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		SpaGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "HotelGetAmountDecimal:%f", SpaGetAmountDecimal.InexactFloat64())
		SpaGetAmount := SpaGetAmountDecimal.
			Mul(decimal.NewFromFloat(SpaOrder.SpaGetRateScene)).
			Div(decimal.NewFromFloat(100)).
			Mul(decimal.NewFromFloat(SpaOrder.SpaGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0).InexactFloat64()
		g.Log().Debugf(ctx, "HotelGetAmount:%f", SpaGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      gvar.New(SpaOrder.MemberId).Int(),
			Scene:         "SPA",
			Type:          "AWARD",
			ChangeBalance: SpaGetAmount,
			OrderSn:       SpaOrder.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.SpaOrder.Ctx(ctx).TX(tx).WherePri(SpaOrder.Id).Data(g.Map{
			dao.SpaOrder.Columns().SpaGetScoreStatus: "SUCCESS",
			dao.SpaOrder.Columns().SpaGetAmount:      SpaGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if SpaGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("按摩订单下单成功奖励您 %d 积分", gvar.New(SpaGetAmount).Int()),
				"en":    fmt.Sprintf("Massage order placed: +%d points", gvar.New(SpaGetAmount).Int()),
				"ja":    fmt.Sprintf("マッサージ注文成功：+%d ポイント", gvar.New(SpaGetAmount).Int()),
				"ko":    fmt.Sprintf("마사지 주문 완료: +%d포인트", gvar.New(SpaGetAmount).Int()),
				"zh_CN": fmt.Sprintf("按摩訂單下單成功，+%d 積分", gvar.New(SpaGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             int(SpaOrder.MemberId),
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleCarReturnScore 用户下单返积分
func HandleCarReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		OrderSn      string
		CarOrder     *entity.CarOrder
		YYConfig     *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.CarOrder.Ctx(ctx).TX(tx).
			Where(dao.CarOrder.Columns().OrderSn, OrderSn).
			Scan(&CarOrder); err != nil {
			return
		}
		if g.IsEmpty(CarOrder) {
			err = gerror.New("接送机订单数据为空")
			return
		}
		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, CarOrder.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}
		if CarOrder.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if CarOrder.CarGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		CarGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "HotelGetAmountDecimal:%f", CarGetAmountDecimal.InexactFloat64())
		CarGetAmount := CarGetAmountDecimal.
			Mul(decimal.NewFromFloat(CarOrder.CarGetRateScene)).
			Div(decimal.NewFromFloat(100)).
			Mul(decimal.NewFromFloat(CarOrder.CarGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0).InexactFloat64()
		g.Log().Debugf(ctx, "HotelGetAmount:%f", CarGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      gvar.New(CarOrder.MemberId).Int(),
			Scene:         "CAR",
			Type:          "AWARD",
			ChangeBalance: CarGetAmount,
			OrderSn:       CarOrder.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.CarOrder.Ctx(ctx).TX(tx).WherePri(CarOrder.Id).Data(g.Map{
			dao.CarOrder.Columns().CarGetScoreStatus: "SUCCESS",
			dao.CarOrder.Columns().CarGetAmount:      CarGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if CarGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("接送机订单下单成功奖励您 %d 积分", gvar.New(CarGetAmount).Int()),
				"en":    fmt.Sprintf("Transfer order placed: +%d points", gvar.New(CarGetAmount).Int()),
				"ja":    fmt.Sprintf("送迎注文成功：+%d ポイント", gvar.New(CarGetAmount).Int()),
				"ko":    fmt.Sprintf("공항 픽업 주문 완료: +%d포인트", gvar.New(CarGetAmount).Int()),
				"zh_CN": fmt.Sprintf("接送機訂單下單成功，+%d 積分", gvar.New(CarGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             int(CarOrder.MemberId),
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleCabinetReturnScore 用户下单返积分-储物柜
func HandleCabinetReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		OrderSn      string
		CabinetOrder *entity.CabinetOrder
		YYConfig     *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.CabinetOrder.Ctx(ctx).TX(tx).
			Where(dao.CabinetOrder.Columns().OrderSn, OrderSn).
			Scan(&CabinetOrder); err != nil {
			return
		}
		if g.IsEmpty(CabinetOrder) {
			err = gerror.New("储物柜订单数据为空")
			return
		}
		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, CabinetOrder.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}
		if CabinetOrder.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if CabinetOrder.CabinetGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		CabinetGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "CabinetGetAmountDecimal:%f", CabinetGetAmountDecimal.InexactFloat64())
		CabinetGetAmount := CabinetGetAmountDecimal.
			Mul(decimal.NewFromFloat(CabinetOrder.CabinetGetRateScene)).
			Div(decimal.NewFromFloat(100)).
			Mul(decimal.NewFromFloat(CabinetOrder.CabinetGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0).InexactFloat64()
		g.Log().Debugf(ctx, "CabinetGetAmount:%f", CabinetGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      gvar.New(CabinetOrder.MemberId).Int(),
			Scene:         "CABINET",
			Type:          "AWARD",
			ChangeBalance: CabinetGetAmount,
			OrderSn:       CabinetOrder.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).WherePri(CabinetOrder.Id).Data(g.Map{
			dao.CabinetOrder.Columns().CabinetGetScoreStatus: "SUCCESS",
			dao.CabinetOrder.Columns().CabinetGetAmount:      CabinetGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if CabinetGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("储物柜订单下单成功奖励您 %d 积分", gvar.New(CabinetGetAmount).Int()),
				"en":    fmt.Sprintf("Locker order placed: +%d points", gvar.New(CabinetGetAmount).Int()),
				"ja":    fmt.Sprintf("ロッカー注文成功：+%d ポイント", gvar.New(CabinetGetAmount).Int()),
				"ko":    fmt.Sprintf("보관함 주문 완료: +%d포인트", gvar.New(CabinetGetAmount).Int()),
				"zh_CN": fmt.Sprintf("儲物櫃訂單下單成功，+%d 積分", gvar.New(CabinetGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             int(CabinetOrder.MemberId),
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleTravelReturnScore 用户下单返积分
func HandleTravelReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		OrderSn      string
		TravelOrder  *entity.TravelOrder
		YYConfig     *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.TravelOrder.Ctx(ctx).TX(tx).
			Where(dao.TravelOrder.Columns().OrderSn, OrderSn).
			Scan(&TravelOrder); err != nil {
			return
		}
		if g.IsEmpty(TravelOrder) {
			err = gerror.New("一日游订单数据为空")
			return
		}
		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, TravelOrder.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}
		if TravelOrder.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if TravelOrder.TravelGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		CarGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "HotelGetAmountDecimal:%f", CarGetAmountDecimal.InexactFloat64())
		CarGetAmount := CarGetAmountDecimal.
			Mul(decimal.NewFromFloat(TravelOrder.TravelGetRateScene)).
			Div(decimal.NewFromFloat(100)).
			Mul(decimal.NewFromFloat(TravelOrder.TravelGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0).InexactFloat64()
		g.Log().Debugf(ctx, "TravelGetAmount:%f", CarGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      gvar.New(TravelOrder.MemberId).Int(),
			Scene:         "TRAVEL",
			Type:          "AWARD",
			ChangeBalance: CarGetAmount,
			OrderSn:       TravelOrder.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).WherePri(TravelOrder.Id).Data(g.Map{
			dao.TravelOrder.Columns().TravelGetScoreStatus: "SUCCESS",
			dao.TravelOrder.Columns().TravelGetAmount:      CarGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		// TODO 一日游-发送消息
		if CarGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("一日游订单下单成功奖励您 %d 积分", gvar.New(CarGetAmount).Int()),
				"en":    fmt.Sprintf("Travel order placed: +%d points", gvar.New(CarGetAmount).Int()),
				"ja":    fmt.Sprintf("旅行予約が成功しました：+%d ポイント", gvar.New(CarGetAmount).Int()),
				"ko":    fmt.Sprintf("공항 픽업 주문 완료: +%d포인트", gvar.New(CarGetAmount).Int()),
				"zh_CN": fmt.Sprintf("一日遊訂單下單成功，+%d 積分", gvar.New(CarGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             int(TravelOrder.MemberId),
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleFoodReturnScore 用户下单返积分
func HandleFoodReturnScore(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body         = msg.Body
		OrderSn      string
		FoodOrder    *entity.FoodOrder
		YYConfig     *model.YYConfig
		PayAmount    float64
		RefundAmount float64
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.FoodOrder.Ctx(ctx).TX(tx).
			Where(dao.FoodOrder.Columns().OrderSn, OrderSn).
			Scan(&FoodOrder); err != nil {
			return
		}
		if g.IsEmpty(FoodOrder) {
			err = gerror.New("餐厅订单数据为空")
			return
		}
		// 查询下单人信息
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, FoodOrder.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销")
			return
		}
		if FoodOrder.IsGetOpen != "Y" {
			err = gerror.New("未开启订单返现")
			return
		}
		if FoodOrder.FoodGetScoreStatus != "WAIT" {
			err = gerror.New("订单状态不为返现处理中")
			return
		}
		RebateLogger.Info(ctx, "结算返现")
		// 查询订单外部支付退款和外部支付信息
		if PayAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").
			WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Group(dao.PmsTransaction.Columns().OrderSn).
			Sum(dao.PmsTransaction.Columns().PayAmount); err != nil {
			return
		}
		if PayAmount <= 0 {
			err = gerror.New("未产生除积分外的支付信息无法进行结算佣金")
			return
		}
		if RefundAmount, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).
			Where(dao.PmsTransactionRefund.Columns().OrderSn, OrderSn).
			WhereNot(dao.PmsTransactionRefund.Columns().RefundType, "BAL").
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Group(dao.PmsTransactionRefund.Columns().OrderSn).
			Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
			return
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		// 使用decimal进行精确计算，避免中间四舍五入导致的精度损失
		FoodGetAmountDecimal := decimal.NewFromFloat(PayAmount).Sub(decimal.NewFromFloat(RefundAmount))
		g.Log().Debugf(ctx, "HotelGetAmountDecimal:%f", FoodGetAmountDecimal.InexactFloat64())
		FoodGetAmount := FoodGetAmountDecimal.
			Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateScene)).
			Div(decimal.NewFromFloat(100)).
			Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateVip)).
			Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).
			Round(0).InexactFloat64()
		g.Log().Debugf(ctx, "FoodGetRateScene:%f, FoodGetRateVip:%f, ExchangeRate:%f", FoodOrder.FoodGetRateScene, FoodOrder.FoodGetRateVip, YYConfig.ExchangeRate)
		g.Log().Debugf(ctx, "FoodGetAmount1:%f", FoodGetAmountDecimal.Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateScene)).Div(decimal.NewFromFloat(100)).InexactFloat64())
		g.Log().Debugf(ctx, "FoodGetAmount2:%f", FoodGetAmountDecimal.Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateScene)).Div(decimal.NewFromFloat(100)).Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateVip)).InexactFloat64())
		g.Log().Debugf(ctx, "FoodGetAmount3:%f", FoodGetAmountDecimal.Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateScene)).Div(decimal.NewFromFloat(100)).Mul(decimal.NewFromFloat(FoodOrder.FoodGetRateVip)).Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).InexactFloat64())
		g.Log().Debugf(ctx, "HotelGetAmount:%f", FoodGetAmount)
		// 结算给用户
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      gvar.New(FoodOrder.MemberId).Int(),
			Scene:         "FOOD",
			Type:          "AWARD",
			ChangeBalance: FoodGetAmount,
			OrderSn:       FoodOrder.OrderSn,
			Reason:        "下单返积分",
		}, tx); err != nil {
			return
		}
		// 记录结算结果
		if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).WherePri(FoodOrder.Id).Data(g.Map{
			dao.FoodOrder.Columns().FoodGetScoreStatus: "SUCCESS",
			dao.FoodOrder.Columns().FoodGetAmount:      FoodGetAmount,
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if FoodGetAmount > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "下单奖励积分",
				"en":    "Order Points Reward",
				"ja":    "注文ポイント報酬",
				"ko":    "주문 포인트 보상",
				"zh_CN": "下單獎勵積分",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("餐饮订单下单成功奖励您 %d 积分", gvar.New(FoodGetAmount).Int()),
				"en":    fmt.Sprintf("Dining order placed: +%d points", gvar.New(FoodGetAmount).Int()),
				"ja":    fmt.Sprintf("飲食注文成功：+%d ポイント", gvar.New(FoodGetAmount).Int()),
				"ko":    fmt.Sprintf("식당 주문 완료: +%d포인트", gvar.New(FoodGetAmount).Int()),
				"zh_CN": fmt.Sprintf("餐飲訂單下單成功，+%d 積分", gvar.New(FoodGetAmount).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             int(FoodOrder.MemberId),
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_ReturnBrokerage")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_ReturnScore")
		return err
	}
	return
}

// HandleInviteAward 邀新首单推荐人奖励
func HandleInviteAward(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body                  = msg.Body
		OrderSn               string
		StayInfo              *entity.PmsAppStay
		CheckStayInfo         *entity.PmsAppStay
		checkinStatusVar      *gvar.Var
		ReferrerInfo          *entity.PmsMember
		InviteNewRewardConfig *model.InviteNewRewardConfig
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			RebateLogger.Error(ctx, r)
		}
	}()
	// 获取奖励配置
	if InviteNewRewardConfig, err = service.BasicsConfig().GetInviteNewRewardConfig(ctx); err != nil {
		return
	}
	if InviteNewRewardConfig.IsInviteRewardOpen != 1 {
		err = gerror.New("未开启邀新奖励")
		return
	}
	// 获取奖励
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		RebateLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
			Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
			Scan(&StayInfo); err != nil {
			return
		}
		if g.IsEmpty(StayInfo) {
			err = gerror.New("住宿订单数据为空")
			return
		}
		if g.IsEmpty(StayInfo.Referrer) {
			err = gerror.New("没有推荐人无需处理推荐人推荐奖励")
			return
		}
		if checkinStatusVar, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
			Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).
			Group(dao.PmsAppReservation.Columns().CheckinStatus).
			Fields("GROUP_CONCAT(checkin_status) as checkin_status").
			Value(); err != nil {
			return
		}
		if checkinStatusVar.String() != "checked_out" {
			err = gerror.New("订单未完结无法处理推荐人推荐奖励")
			return
		}
		CheckStayInfo = new(entity.PmsAppStay)
		// 是否是首次下单
		if err = dao.PmsAppStay.Ctx(ctx).
			Where(dao.PmsAppStay.Columns().MemberId, StayInfo.MemberId).
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			OrderAsc(dao.PmsAppStay.Columns().Id).
			Scan(&CheckStayInfo); err != nil {
			return
		}
		if g.IsEmpty(CheckStayInfo) {
			// RebateLogger.Info(ctx, "无订单不需要处理下单之后邀请人获得奖励")
			err = gerror.New("无订单不需要处理下单之后邀请人获得奖励")
			return
		}
		if CheckStayInfo.OrderSn != OrderSn {
			// RebateLogger.Info(ctx, "非首次下单不进行结算")
			err = gerror.New("非首次下单不进行结算")
			return
		}
		RebateLogger.Info(ctx, "结算首次被邀请人下单之后邀请人获得奖励")
		// 查看推荐人信息
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, StayInfo.Referrer).Scan(&ReferrerInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(ReferrerInfo) {
			err = gerror.New("推荐人不存在无法处理推荐人推荐奖励")
			return
		}

		// 查看订单用户信息
		MemberInfo := new(entity.PmsMember)
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, StayInfo.MemberId).Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(MemberInfo) {
			err = gerror.New("订单用户不存在无法处理推荐人推荐奖励")
			return
		}

		if MemberInfo.IsInviteRewards == "Y" {
			// RebateLogger.Info(ctx, "推荐人已获得奖励无法处理推荐人推荐奖励")
			err = gerror.New("推荐人已获得奖励无法处理推荐人推荐奖励")
			return
		}
		ChangeBalance := 0.0
		if ReferrerInfo.RebateMode == "STAFF" {
			ChangeBalance = InviteNewRewardConfig.StaffRewardBalance
		} else if ReferrerInfo.RebateMode == "CHANNEL" {
			ChangeBalance = InviteNewRewardConfig.ChannelRewardBalance
		} else if ReferrerInfo.RebateMode == "MEMBER" {
			ChangeBalance = InviteNewRewardConfig.MemberRewardBalance
		} else {
			err = gerror.New("推荐人推荐奖励模式异常")
			return
		}
		// 发放积分奖励
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      StayInfo.Referrer,
			Scene:         "SYSTEM",
			Type:          "AWARD",
			ChangeBalance: ChangeBalance,
			OrderSn:       OrderSn,
			Reason:        "邀请人首单奖励",
		}, tx); err != nil {
			return
		}
		// 修改注册人发放过推荐人奖励
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).WherePri(StayInfo.MemberId).Data(g.Map{
			dao.PmsMember.Columns().IsInviteRewards: "Y",
		}).UpdateAndGetAffected(); err != nil {
			return
		}

		if ChangeBalance > 0 {
			// 发消息
			systemMessageTitle := map[string]string{
				"zh":    "推荐人首单奖励",
				"en":    "Referrer First Order Reward",
				"ja":    "紹介者初回注文報酬",
				"ko":    "추천인 첫 주문 보상",
				"zh_CN": "推薦人首單獎勵",
			}
			systemMessageContent := map[string]string{
				"zh":    fmt.Sprintf("您的推荐人首单下单成功奖励您 %d 积分", gvar.New(ChangeBalance).Int()),
				"en":    fmt.Sprintf("You earned %d points from a first referral order", gvar.New(ChangeBalance).Int()),
				"ja":    fmt.Sprintf("紹介者の初回注文で %d ポイント獲得", gvar.New(ChangeBalance).Int()),
				"ko":    fmt.Sprintf("추천인 첫 주문으로 %d포인트 획득", gvar.New(ChangeBalance).Int()),
				"zh_CN": fmt.Sprintf("推薦人首單成功，獲得 %d 積分", gvar.New(ChangeBalance).Int()),
			}
			appPushData := g.MapStrStr{
				"type": "0",
			}

			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "system",
				Type:                 "im",
				MemberId:             StayInfo.Referrer,
				Language:             contexts.GetLanguage(ctx),
				AppPushData:          appPushData,
				AppLink:              "/account/score",
				WxLink:               "/pages/user/points",
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
				PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
				ShowIndex:            true,
				OperatorId:           0,
				OperatorRole:         "SYSTEM",
			})
		}

		RebateLogger.Info(ctx, "RUN_SUCCESS_HandleInviteAward")
		return err
	}); err != nil {
		RebateLogger.Info(ctx, "RUN_FAIL_HandleInviteAward")
		return err
	}
	return
}
