package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) OrderChangeOptions(ctx context.Context, req *hotel.OrderChangeOptionsReq) (res *hotel.OrderChangeOptionsRes, err error) {
	var (
		PreRefundOut *input_hotel.PreRefundOut
	)
	res = new(hotel.OrderChangeOptionsRes)
	if err = dao.PmsAppReservation.Ctx(ctx).WherePri(req.OrderId).WithAll().Scan(&res); err != nil {
		return
	}
	if g.IsEmpty(res) {
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}

	// 如果订单状态为已取消那么不可变更订单
	if res.OrderStatus == "CANCEL" {
		return
	}

	if res.CheckinStatus == "before_checkin" {
		// 可以进行变更
		res.IsNextChange = true
		// 可以取消订单
		res.IsCancelOrder = true
	}

	if res.CheckinStatus != "checked_out" {
		// 可以续住
		res.IsStayOn = true
	}

	// 该入住单是否在免费取消内
	PreRefundOut, err = service.HotelService().PreRefundOrderDetail(ctx, &input_hotel.PreRefundIn{OrderSn: res.OrderSn})
	if err != nil {
		err = nil
		res.IsNextDateChange = false
	} else {
		if PreRefundOut.CancelFee == 0 {
			res.IsNextDateChange = true
		}
	}

	return
}

func (c *ControllerHotel) OrderChangeGuestEdit(ctx context.Context, req *hotel.OrderChangeGuestEditReq) (res *hotel.OrderChangeGuestEditRes, err error) {
	if err = service.HotelService().OrderChangeGuestEdit(ctx, &req.OrderChangeGuestReq); err != nil {
		return
	}
	return
}

func (c *ControllerHotel) OrderChangeGuestInfo(ctx context.Context, req *hotel.OrderChangeGuestInfoReq) (res *hotel.OrderChangeGuestInfoRes, err error) {
	var (
		appReservation  *entity.PmsAppReservation
		PmsGuestProfile *entity.PmsGuestProfile
	)
	res = new(hotel.OrderChangeGuestInfoRes)
	if err = dao.PmsAppReservation.Ctx(ctx).WherePri(req.OrderId).Scan(&appReservation); err != nil {
		return
	}
	if g.IsEmpty(appReservation) {
		// 不存在入住单
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}
	if err = dao.PmsGuestProfile.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsGuestProfile.Columns().Uid: appReservation.MainGuest,
	}).Scan(&PmsGuestProfile); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(PmsGuestProfile) {
		res.Phone = PmsGuestProfile.Phone
		res.Email = PmsGuestProfile.Email
		res.FirstName = PmsGuestProfile.FirstName
		res.LastName = PmsGuestProfile.LastName
		res.AreaNo = PmsGuestProfile.AreaNo
	}
	res.AirhostOrderUuid = appReservation.Uuid
	res.OrderId = appReservation.Id
	return
}

func (c *ControllerHotel) OrderChangeDatePreInfo(ctx context.Context, req *hotel.OrderChangeDatePreInfoReq) (res *hotel.OrderChangeDatePreInfoRes, err error) {
	var (
		appReservation  *entity.PmsAppReservation
		appReservations []*entity.PmsAppReservation
		PmsRoomType     *entity.PmsRoomType
		Availabilities  []*entity.PmsAvailabilities
		Charge          []*entity.PmsCharge
		NewAllAmount    float64
		CancelRate      []*struct {
			Name      string `json:"name"      dc:"规则名"`
			Mode      string `json:"mode"      dc:"规则模式"`
			StartDays int    `json:"startDays" dc:"开始天数"`
			EndDays   int    `json:"endDays"   dc:"结束天数"`
			Rate      int    `json:"rate"      dc:"取消费率"`
			Date      string `json:"date"      dc:"规则解析日期"`
			Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
		}
		ChangeOrderSn  = uuid.CreateOrderCode("HC")
		PayConfig      *model.PayConfig
		ExpirationTime int
		PmsPirceConfig *model.PmsPriceConfig
		MemberUser     = contexts.GetMemberUser(ctx)
		IsFx           = "N"
	)
	if MemberUser.IsFx {
		IsFx = "Y"
	}
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	res = new(hotel.OrderChangeDatePreInfoRes)
	// 查询订单信息
	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.Map{
		dao.PmsAppReservation.Columns().Id: req.OrderId,
	}).Scan(&appReservation); err != nil {
		return
	}

	if g.IsEmpty(appReservation) {
		// 不存在入住单
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}

	// 入住日期和离店日期没变动不可修改
	if appReservation.CheckinDate.Format("Y-m-d") == req.CheckInDate && appReservation.CheckoutDate.Format("Y-m-d") == req.CheckOutDate {
		err = gerror.New(gi18n.T(ctx, "checkin_checkout_date_unchanged"))
		return
	}

	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.Map{
		dao.PmsAppReservation.Columns().Id: req.OrderId,
	}).Scan(&appReservations); err != nil {
		return
	}
	if len(appReservations) > 1 {
		// 当前订单不允许变更日期
		err = gerror.New(gi18n.T(ctx, "the_current_order_cannot_change_date"))
		return
	}

	// 查询房间名称
	if err = dao.PmsRoomType.Ctx(ctx).Where(g.Map{
		dao.PmsRoomType.Columns().Uid: appReservation.RoomType,
	}).Hook(hook.PmsFindLanguageValueHook).Scan(&PmsRoomType); err != nil {
		return
	}

	if g.IsEmpty(PmsRoomType) {
		// 房间不存在
		err = gerror.New(gi18n.T(ctx, "the_room_does_not_exist"))
		return
	}

	res.RoomName = PmsRoomType.Name
	res.RoomNumber = 1

	// 查询变更单是否存在
	if err = service.HotelService().OrderChangeCheckIsExist(ctx, req.OrderId, "DATE"); err != nil {
		return
	}
	res.OldCheckInDate = appReservation.CheckinDate.Format("Y-m-d")
	res.OldCheckOutDate = appReservation.CheckoutDate.Format("Y-m-d")
	res.NewCheckInDate = req.CheckInDate
	res.NewCheckOutDate = req.CheckOutDate

	// 查询历史订单价格
	if err = dao.PmsCharge.Ctx(ctx).Where(g.Map{
		dao.PmsCharge.Columns().Uid: appReservation.Charges,
	}).Scan(&Charge); err != nil {
		return
	}

	if g.IsEmpty(Charge) {
		// 不存在价格信息
		err = gerror.New(gi18n.T(ctx, "no_price_info"))
		return
	}

	// 查询价格
	Availabilities = nil
	if err = dao.PmsAvailabilities.Ctx(ctx).
		Where(dao.PmsAvailabilities.Columns().Tuid, appReservation.RoomType).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, res.NewCheckInDate).
		WhereLT(dao.PmsAvailabilities.Columns().Date, res.NewCheckOutDate).
		Scan(&Availabilities); err != nil {
		return
	}

	// 读取支付配置
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()
	// 计算入住总人数（成人+儿童，婴儿不计入）
	totalGuests := appReservation.AdultCount + appReservation.ChildCount
	for _, v := range Availabilities {
		if v.Allotment < 1 {
			// 您选择的日期已经没有满足条件的房间库存了
			err = gerror.NewCode(gcode.New(1111, gi18n.T(ctx, "the_date_no_stock"), nil))
			return
		}

		// 检查该日期是否在原订单的Charge中存在
		var price float64
		isExistingDate := false
		for _, v1 := range Charge {
			if v.Date == v1.Date.Format("Y-m-d") {
				// 已存在的日期，累加所有费用（基础费+超员费等）
				price += v1.Amount
				isExistingDate = true
			}
		}

		if !isExistingDate {
			// 新日期：基础价格 + 超员费
			price = decimal.NewFromFloat(v.Price).Mul(PricePercent).Round(0).InexactFloat64()
			// 计算超员费
			if totalGuests > PmsRoomType.OccupantsForBaseRate {
				extraGuests := totalGuests - PmsRoomType.OccupantsForBaseRate
				extraFee := decimal.NewFromInt(int64(extraGuests)).Mul(decimal.NewFromFloat(PmsRoomType.AdditionalGuestAmounts)).InexactFloat64()
				price += extraFee
			}
		}
		NewAllAmount += price
	}

	res.NewOrderPrice = NewAllAmount
	res.OldOrderPrice = appReservation.BookingFee
	DiffOrderPrice := decimal.NewFromFloat(res.NewOrderPrice).Sub(decimal.NewFromFloat(appReservation.BookingFee)).Round(0).InexactFloat64()
	res.DiffOrderPrice = DiffOrderPrice
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = cache.Instance().Set(ctx, fmt.Sprintf("changeOrder:%s", ChangeOrderSn), &entity.PmsAppReservationChange{
			ChangeOrderSn:   ChangeOrderSn,
			OrderId:         appReservation.Id,
			OrderSn:         appReservation.OrderSn,
			OutOrderSn:      appReservation.OutOrderSn,
			OldCheckinDate:  appReservation.CheckinDate,
			OldCheckoutDate: appReservation.CheckoutDate,
			NewCheckinDate:  gtime.New(req.CheckInDate),
			NewCheckoutDate: gtime.New(req.CheckOutDate),
			ChangeType:      "DATE",
			ChangeStatus:    "ING",
			ChangeAmount:    DiffOrderPrice,
			SubmitDate:      gtime.Now(),
			PricePercent:    PmsPirceConfig.PricePercent,
			ExpirationTime:  ExpirationTime,
			OldOrderPrice:   appReservation.BookingFee,
			NewOrderPrice:   NewAllAmount,
			IsFx:            IsFx,
		}, time.Second*60*15); err != nil {
			return
		}
		return
	}); err != nil {
		return
	}
	res.ChangeOrderSn = ChangeOrderSn
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	var IsCancel = "Y"
	if !g.IsEmpty(appReservations[0].PricePlanInfo) {
		var PricePlanInfo *entity.PmsPricePlan
		if err = json.Unmarshal([]byte(gvar.New(appReservations[0].PricePlanInfo).String()), &PricePlanInfo); err != nil {
			// 解析路径失败
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
		IsCancel = PricePlanInfo.IsCancel
	}

	res.IsCancel = IsCancel

	for k, v := range CancelRate {
		NowTime := gtime.Now().Format("Y-m-d")
		switch v.Mode {
		case "after":
			startDate := gtime.New(req.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			//res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			if len(res.CancelDate) >= 1 {
				res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			} else {
				res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s", v.Name))
			}

			if startDate <= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = startDate

			break
		case "middle":
			startDay := gtime.New(req.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(req.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if NowTime <= endDate {
				res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s %s", fmt.Sprintf("%s ~ %s", startDay, endDate), v.Name))
			}
			if startDay <= NowTime && endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
			break
		case "before":
			endDate := gtime.New(req.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s %s", endDate, v.Name))
			}
			if endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = endDate
			break
		}
	}

	res.CancelRate = CancelRate
	return
}

func (c *ControllerHotel) OrderChangeBookingPeople(ctx context.Context, req *hotel.OrderChangeBookingPeopleReq) (res *hotel.OrderChangeBookingPeopleRes, err error) {
	var (
		AppReservation   *entity.PmsAppReservation
		RoomType         *entity.PmsRoomType
		ChangeOrderSn    = uuid.CreateOrderCode("HC")
		OldAllNum        int
		OldAllNumPrice   float64
		NewAllNum        int
		NewAllNumPrice   float64
		ChangeSubmitDate = gtime.Now()
		ExpirationTime   int
		PayConfig        *model.PayConfig
		MemberUser       = contexts.GetMemberUser(ctx)
		IsFx             = "N"
	)
	if MemberUser.IsFx {
		IsFx = "Y"
	}
	// 查询变更单是否存在
	if err = service.HotelService().OrderChangeCheckIsExist(ctx, req.OrderId, "PEOPLE"); err != nil {
		return
	}
	res = new(hotel.OrderChangeBookingPeopleRes)
	if err = dao.PmsAppReservation.Ctx(ctx).WherePri(req.OrderId).WithAll().Scan(&AppReservation); err != nil {
		return
	}
	if g.IsEmpty(AppReservation) {
		// 不存在入住单
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}
	if err = dao.PmsRoomType.Ctx(ctx).
		Where(dao.PmsRoomType.Columns().Uid, AppReservation.RoomType).
		Scan(&RoomType); err != nil {
		return
	}
	if g.IsEmpty(RoomType) {
		// 该房型异常
		err = gerror.New(gi18n.T(ctx, "room_type_abnormal"))
		return
	}
	if RoomType.Occupancy < req.Adult+req.Child {
		// 入住人数不符合要求
		err = gerror.New(gi18n.T(ctx, "check_in_people_does_not_meet_the_requirements"))
		return
	}
	OldAllNum = AppReservation.AdultCount + AppReservation.ChildCount
	g.Log().Debug(ctx, "OldAllNum", OldAllNum)
	OldAllNumPrice = decimal.NewFromInt(gvar.New(OldAllNum).Int64()).
		Sub(decimal.NewFromInt(gvar.New(RoomType.OccupantsForBaseRate).Int64())).
		Mul(decimal.NewFromFloat(RoomType.AdditionalGuestAmounts)).
		Round(0).InexactFloat64()
	if OldAllNumPrice < 0 {
		OldAllNumPrice = 0
	}
	g.Log().Debug(ctx, "OldAllNumPrice", OldAllNumPrice)
	NewAllNum = req.Adult + req.Child
	g.Log().Debug(ctx, "NewAllNum", NewAllNum)
	NewAllNumPrice = decimal.NewFromInt(gvar.New(NewAllNum).Int64()).
		Sub(decimal.NewFromInt(gvar.New(RoomType.OccupantsForBaseRate).Int64())).
		Mul(decimal.NewFromFloat(RoomType.AdditionalGuestAmounts)).
		Round(0).InexactFloat64()
	if NewAllNumPrice < 0 {
		NewAllNumPrice = 0
	}
	g.Log().Debug(ctx, "NewAllNumPrice", NewAllNumPrice)
	// 根据OldAllNum重新计算价格
	if NewAllNumPrice > OldAllNumPrice {
		res.ChangePrice = decimal.NewFromFloat(NewAllNumPrice).
			Sub(decimal.NewFromFloat(OldAllNumPrice)).
			Round(0).InexactFloat64()
	}
	res.OldOrderPrice = AppReservation.BookingFee
	res.NewOrderPrice = AppReservation.BookingFee + res.ChangePrice
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 读取支付配置
		if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
			return
		}
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

		// 插入变更单
		if _, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Data(&entity.PmsAppReservationChange{
			OrderId:        AppReservation.Id,
			OrderSn:        AppReservation.OrderSn,
			OutOrderSn:     AppReservation.OutOrderSn,
			ChangeType:     "PEOPLE",
			OldAdultCount:  AppReservation.AdultCount,
			NewAdultCount:  req.Adult,
			OldChildCount:  AppReservation.ChildCount,
			NewChildCount:  req.Child,
			OldInfantCount: AppReservation.InfantCount,
			NewInfantCount: req.Infant,
			ChangeStatus:   "ING",
			ChangeAmount:   res.ChangePrice,
			ChangeOrderSn:  ChangeOrderSn,
			SubmitDate:     ChangeSubmitDate,
			OldOrderPrice:  res.OldOrderPrice,
			NewOrderPrice:  res.NewOrderPrice,
			ExpirationTime: ExpirationTime,
			IsFx:           IsFx,
		}).InsertAndGetId(); err != nil {
			return
		}

		return
	}); err != nil {
		return
	}
	res.ChangeOrderSn = ChangeOrderSn
	return
}

func (c *ControllerHotel) OrderChangeBookingPeopleSubmit(ctx context.Context, req *hotel.OrderChangeBookingPeopleSubmitReq) (res *hotel.OrderChangeBookingPeopleSubmitRes, err error) {
	var (
		PmsAppReservationChange *entity.PmsAppReservationChange
		TransactionSn           = uuid.CreatePayCode("PT")
		PayConfig               *model.PayConfig
		ExpirationTime          int
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()
	res = new(hotel.OrderChangeBookingPeopleSubmitRes)
	if err = dao.PmsAppReservationChange.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservationChange.Columns().ChangeOrderSn: req.ChangeOrderSn,
	}).Scan(&PmsAppReservationChange); err != nil {
		return
	}
	if PmsAppReservationChange.ChangeStatus != "ING" {
		// 该变更单状态异常
		err = gerror.New(gi18n.T(ctx, "change_order_status_abnormal"))
		return
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if PmsAppReservationChange.ChangeAmount > 0 {
			// 插入支付订单流水
			IsFx := PmsAppReservationChange.IsFx
			if g.IsEmpty(IsFx) {
				IsFx = "N"
			}
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsTransaction{
				OrderSn:       PmsAppReservationChange.OrderSn,
				ChangeOrderSn: req.ChangeOrderSn,
				TransactionSn: TransactionSn,
				PayChannel:    "PAYCLOUD",
				PayType:       "WeChatPay",
				OrderType:     "BOOKING_CHANGE",
				Scene:         "HOTEL",
				Amount:        PmsAppReservationChange.ChangeAmount,
				PriceCurrency: "JPY",
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
				IsFx:          IsFx,
			}).InsertAndGetId(); err != nil {
				return
			}
			if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeDelayedName,
				QueueName:    consts.RabbitMQQueueNameOrderExpire,
				DataByte:     gvar.New(req.ChangeOrderSn).Bytes(),
			}); err != nil {
				g.Log().Error(ctx, "发送过期自动取消订单MQ失败", err)
			}
			res.TransactionSn = TransactionSn
			res.Countdown = ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
		} else {
			if _, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.PmsAppReservationChange.Columns().ChangeOrderSn: req.ChangeOrderSn,
			}).Data(g.MapStrAny{
				dao.PmsAppReservationChange.Columns().ChangeStatus: "DONE",
				dao.PmsAppReservationChange.Columns().DoneDate:     gtime.Now().Format("Y-m-d H:i:s"),
			}).Update(); err != nil {
				return
			}
			if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.PmsAppReservation.Columns().Id: PmsAppReservationChange.OrderId,
			}).Data(g.MapStrAny{
				dao.PmsAppReservation.Columns().IsChangePeople:   "Y",
				dao.PmsAppReservation.Columns().IsChangePeopleId: PmsAppReservationChange.Id,
				dao.PmsAppReservation.Columns().AdultCount:       PmsAppReservationChange.NewAdultCount,
				dao.PmsAppReservation.Columns().ChildCount:       PmsAppReservationChange.NewChildCount,
				dao.PmsAppReservation.Columns().InfantCount:      PmsAppReservationChange.NewInfantCount,
			}).Update(); err != nil {
				return
			}

			// 酒店订单日志
			var (
				PmsAppStay entity.PmsAppStay
			)
			if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsAppStay.Columns().OrderSn: PmsAppReservationChange.OrderSn,
			}).Scan(&PmsAppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			if g.IsEmpty(PmsAppStay) {
				// 该订单无需处理
				err = gerror.New(gi18n.T(ctx, "order_does_not_need_handle"))
				return
			}
			if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
				OrderId:     PmsAppStay.Id,
				ActionWay:   "CHANGE_PEOPLE_NUM",
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
				"zh":    "订单入住人数变更成功。",
				"en":    "The number of guests in the order has been successfully changed.",
				"ja":    "注文内のゲストの人数が正常に変更されました。",
				"ko":    "주문에 포함된 손님 수가 성공적으로 변경되었습니다。",
				"zh_CN": "訂單入住人數變更成功。",
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
		}
		return
	}); err != nil {
		return
	}
	res.CreateOrderTime = PmsAppReservationChange.CreatedAt.Format("Y-m-d H:i:s")
	res.ChangeOrderSn = req.ChangeOrderSn
	res.ChangePrice = PmsAppReservationChange.ChangeAmount
	res.IsFx = PmsAppReservationChange.IsFx
	return
}

func (c *ControllerHotel) OrderChangeDatePreInfoSubmit(ctx context.Context, req *hotel.OrderChangeDatePreInfoSubmitReq) (res *hotel.OrderChangeDatePreInfoSubmitRes, err error) {
	var (
		PmsAppReservationChangeGvar *gvar.Var
		PmsAppReservationChange     *entity.PmsAppReservationChange
		PmsAppReservationChangeInfo *entity.PmsAppReservationChange
		PmsTransaction              []*entity.PmsTransaction
		TransactionSn               = uuid.CreatePayCode("PT")
		ChangeStatus                string
		DoneDate                    *gtime.Time
		ChangeDateId                int64
		PayConfig                   *model.PayConfig
		ExpirationTime              int
		PmsAppReservation           *entity.PmsAppReservation
		BookerResponse              *airhousePublicApi.RetrieveRoomReservationJSONDataResponse
		PmsRoomType                 *entity.PmsRoomType
		OldCharges                  []*entity.PmsCharge
		Availabilities              []*entity.PmsAvailabilities
		PmsPirceConfig              *model.PmsPriceConfig
	)
	res = new(hotel.OrderChangeDatePreInfoSubmitRes)
	if PmsAppReservationChangeGvar, err = cache.Instance().Get(ctx, fmt.Sprintf("changeOrder:%s", req.ChangeOrderSn)); err != nil {
		return
	}
	if err = gvar.New(PmsAppReservationChangeGvar).Scan(&PmsAppReservationChange); err != nil {
		return
	}
	if g.IsEmpty(PmsAppReservationChange) {
		// 不存在变更单
		err = gerror.New(gi18n.T(ctx, "no_change_order"))
		return
	}
	// 查询变更单是否存在
	if err = dao.PmsAppReservationChange.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservationChange.Columns().OrderId:      PmsAppReservationChange.OrderId,
		dao.PmsAppReservationChange.Columns().ChangeType:   "DATE",
		dao.PmsAppReservationChange.Columns().ChangeStatus: "DONE",
	}).Scan(&PmsAppReservationChangeInfo); err != nil {
		return
	}
	if !g.IsEmpty(PmsAppReservationChangeInfo) {
		// 已经进行过变更，无法继续变更
		err = gerror.New(gi18n.T(ctx, "the_change_has_already_been_made_and_cannot_continue"))
		return
	}

	res.ChangeOrderSn = req.ChangeOrderSn
	res.DiffOrderPrice = PmsAppReservationChange.ChangeAmount
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

	// 查询原订单信息，用于更新 Charge 表
	if err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().Id, PmsAppReservationChange.OrderId).Scan(&PmsAppReservation); err != nil {
		return
	}
	if g.IsEmpty(PmsAppReservation) {
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}

	// 查询房型信息
	if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, PmsAppReservation.RoomType).Scan(&PmsRoomType); err != nil {
		return
	}
	if g.IsEmpty(PmsRoomType) {
		err = gerror.New(gi18n.T(ctx, "the_room_does_not_exist"))
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

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if res.DiffOrderPrice < 0 {
			// 写入入住变更单
			if ChangeDateId, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Data(&entity.PmsAppReservationChange{
				ChangeOrderSn:   req.ChangeOrderSn,
				OrderId:         PmsAppReservationChange.OrderId,
				OrderSn:         PmsAppReservationChange.OrderSn,
				OutOrderSn:      PmsAppReservationChange.OutOrderSn,
				OldCheckinDate:  gtime.New(PmsAppReservationChange.OldCheckinDate),
				OldCheckoutDate: gtime.New(PmsAppReservationChange.OldCheckoutDate),
				NewCheckinDate:  gtime.New(PmsAppReservationChange.NewCheckinDate),
				NewCheckoutDate: gtime.New(PmsAppReservationChange.NewCheckoutDate),
				ChangeType:      "DATE",
				ChangeStatus:    "DONE",
				ChangeAmount:    PmsAppReservationChange.ChangeAmount,
				SubmitDate:      gtime.Now(),
				DoneDate:        gtime.Now(),
				NewOrderPrice:   PmsAppReservationChange.NewOrderPrice,
				OldOrderPrice:   PmsAppReservationChange.OldOrderPrice,
				ExpirationTime:  ExpirationTime,
				PricePercent:    PmsAppReservationChange.PricePercent,
				IsFx:            PmsAppReservationChange.IsFx,
			}).InsertAndGetId(); err != nil {
				// 已经进行过变更，无法继续变更
				err = gerror.New(gi18n.T(ctx, "the_change_has_already_been_made_and_cannot_continue"))
				return
			}

			// 回写住宿单变更信息
			if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Id, PmsAppReservationChange.OrderId).Update(g.Map{
				dao.PmsAppReservation.Columns().IsChangeDate:   "Y",
				dao.PmsAppReservation.Columns().IsChangeDateId: ChangeDateId,
				dao.PmsAppReservation.Columns().CheckinDate:    PmsAppReservationChange.NewCheckinDate,
				dao.PmsAppReservation.Columns().CheckoutDate:   PmsAppReservationChange.NewCheckoutDate,
				dao.PmsAppReservation.Columns().BookingFee:     PmsAppReservationChange.NewOrderPrice,
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
			if err = updateChargesForDateChange(ctx, tx, PmsAppReservation, PmsRoomType, OldCharges, Availabilities, PmsAppReservationChange, PmsPirceConfig); err != nil {
				return
			}

			// 查询支付信息
			if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.Map{
				dao.PmsTransaction.Columns().OrderSn:   PmsAppReservationChange.OrderSn,
				dao.PmsTransaction.Columns().PayStatus: "DONE",
			}).WhereNot(dao.PmsTransaction.Columns().RefundStatus, "DONE").
				OrderDesc(dao.PmsTransaction.Columns().Id).Scan(&PmsTransaction); err != nil {
				return
			}

			if len(PmsTransaction) == 0 {
				return
			}

			if err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
				OrderSn:      PmsAppReservationChange.OrderSn,
				RefundAmount: math.Abs(PmsAppReservationChange.ChangeAmount),
			}, tx); err != nil {
				return
			}
			// 修正原订单价格

			defer func() {
				if err = dao.PmsAppReservation.Ctx(ctx).Where(g.Map{
					dao.PmsAppReservation.Columns().OrderSn: PmsAppReservationChange.OrderSn,
				}).Scan(&PmsAppReservation); err != nil {
					return
				}
				if g.IsEmpty(PmsAppReservation) {
					// 入住单不存在，无法进行退款
					err = gerror.New(gi18n.T(ctx, "checkin_order_does_not_exist_cannot_refund"))
					return
				}
				if BookerResponse, err = airhousePublicApi.UpdateRoomReservationPost(ctx, PmsAppReservation.Uuid, g.Map{
					"checkin_date":   PmsAppReservationChange.NewCheckinDate.Format("Y-m-d"),
					"checkout_date":  PmsAppReservationChange.NewCheckoutDate.Format("Y-m-d"),
					"prepaid_amount": PmsAppReservationChange.NewOrderPrice,
					"booking_fee":    PmsAppReservationChange.NewOrderPrice,
				}); err != nil {
					// 变更修改失败
					err = gerror.New(gi18n.T(ctx, "order_change_fail"))
					return
				}
				g.Log().Debug(ctx, BookerResponse)
			}()

			// 酒店订单日志
			var (
				PmsAppStay entity.PmsAppStay
			)
			if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsAppStay.Columns().OrderSn: PmsAppReservationChange.OrderSn,
			}).Scan(&PmsAppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			if g.IsEmpty(PmsAppStay) {
				// 该订单无需处理
				err = gerror.New(gi18n.T(ctx, "order_does_not_need_handle"))
				return
			}

			if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
				OrderId:     PmsAppStay.Id,
				ActionWay:   "REFUND",
				Remark:      "订单退款",
				OperateType: "USER",
				OperateId:   PmsAppStay.MemberId,
			}); err != nil {
				return
			}

			if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
				OrderId:     PmsAppStay.Id,
				ActionWay:   "CHANGE_DATE",
				Remark:      gvar.New(ChangeDateId).String(),
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
				"zh":    "订单入住时间成功变更为" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "，已成功退还差价。",
				"en":    "The check-in date has been successfully changed to " + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + ", and the price difference has been successfully refunded.",
				"ja":    "チェックイン日を" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "に変更し、差額を返金いたしました。",
				"ko":    "체크인 날짜가 " + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "으로 성공적으로 변경되었으며, 가격 차액은 성공적으로 환불되었습니다.",
				"zh_CN": "訂單入住時間成功變更為" + gtime.New(PmsAppReservationChange.NewCheckinDate).Format("Y-m-d") + "~" + gtime.New(PmsAppReservationChange.NewCheckoutDate).Format("Y-m-d") + "，已成功退還差價",
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

			if gtime.New(PmsAppReservationChange.OldCheckinDate) != gtime.New(PmsAppReservationChange.NewCheckinDate) {

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

		} else {
			ChangeStatus = "ING"
			DoneDate = nil
			if PmsAppReservationChange.ChangeAmount == 0 {
				ChangeStatus = "DONE"
				DoneDate = gtime.Now()
			}
			if ChangeDateId, err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).Data(&entity.PmsAppReservationChange{
				ChangeOrderSn:   req.ChangeOrderSn,
				OrderId:         PmsAppReservationChange.OrderId,
				OrderSn:         PmsAppReservationChange.OrderSn,
				OutOrderSn:      PmsAppReservationChange.OutOrderSn,
				OldCheckinDate:  gtime.New(PmsAppReservationChange.OldCheckinDate),
				OldCheckoutDate: gtime.New(PmsAppReservationChange.OldCheckoutDate),
				NewCheckinDate:  gtime.New(PmsAppReservationChange.NewCheckinDate),
				NewCheckoutDate: gtime.New(PmsAppReservationChange.NewCheckoutDate),
				ChangeType:      "DATE",
				ChangeStatus:    ChangeStatus,
				ChangeAmount:    PmsAppReservationChange.ChangeAmount,
				SubmitDate:      gtime.Now(),
				DoneDate:        DoneDate,
				ExpirationTime:  ExpirationTime,
				OldOrderPrice:   PmsAppReservationChange.OldOrderPrice,
				NewOrderPrice:   PmsAppReservationChange.NewOrderPrice,
				IsFx:            PmsAppReservationChange.IsFx,
			}).InsertAndGetId(); err != nil {
				// 已经进行过变更，无法继续变更
				err = gerror.New(gi18n.T(ctx, "the_change_has_already_been_made_and_cannot_continue"))
				return
			}

			if PmsAppReservationChange.ChangeAmount == 0 {
				if _, err = dao.PmsAppReservation.Ctx(ctx).Where(g.Map{
					dao.PmsAppReservation.Columns().Id: PmsAppReservationChange.OrderId,
				}).Data(g.Map{
					dao.PmsAppReservation.Columns().IsChangeDate:   "Y",
					dao.PmsAppReservation.Columns().IsChangeDateId: ChangeDateId,
					dao.PmsAppReservation.Columns().CheckinDate:    PmsAppReservationChange.NewCheckinDate,
					dao.PmsAppReservation.Columns().CheckoutDate:   PmsAppReservationChange.NewCheckoutDate,
				}).Update(); err != nil {
					return
				}

				if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, PmsAppReservationChange.OrderSn).Update(g.Map{
					dao.PmsAppStay.Columns().CheckInDate:  PmsAppReservationChange.NewCheckinDate,
					dao.PmsAppStay.Columns().CheckOutDate: PmsAppReservationChange.NewCheckoutDate,
				}); err != nil {
					return
				}

				// 更新 Charge 表：删除不在新日期范围内的记录，添加新日期的记录
				if err = updateChargesForDateChange(ctx, tx, PmsAppReservation, PmsRoomType, OldCharges, Availabilities, PmsAppReservationChange, PmsPirceConfig); err != nil {
					return
				}

				// 酒店订单日志
				var (
					PmsAppStay entity.PmsAppStay
				)
				if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsAppStay.Columns().OrderSn: PmsAppReservationChange.OrderSn,
				}).Scan(&PmsAppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
					return
				}
				if g.IsEmpty(PmsAppStay) {
					// 该订单无需处理
					err = gerror.New(gi18n.T(ctx, "order_does_not_need_handle"))
					return
				}
				if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
					OrderId:     PmsAppStay.Id,
					ActionWay:   "CHANGE_DATE",
					Remark:      gvar.New(ChangeDateId).String(),
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

				defer func() {
					if err = dao.PmsAppReservation.Ctx(ctx).Where(g.Map{
						dao.PmsAppReservation.Columns().OrderSn: PmsAppReservationChange.OrderSn,
					}).Scan(&PmsAppReservation); err != nil {
						return
					}
					if g.IsEmpty(PmsAppReservation) {
						// 入住单不存在，无法进行退款
						err = gerror.New(gi18n.T(ctx, "checkin_order_does_not_exist_cannot_refund"))
						return
					}
					if BookerResponse, err = airhousePublicApi.UpdateRoomReservationPost(ctx, PmsAppReservation.Uuid, g.Map{
						"checkin_date":  PmsAppReservationChange.NewCheckinDate.Format("Y-m-d"),
						"checkout_date": PmsAppReservationChange.NewCheckoutDate.Format("Y-m-d"),
					}); err != nil {
						// 变更修改失败
						err = gerror.New(gi18n.T(ctx, "order_change_fail"))
						return
					}
					g.Log().Debug(ctx, BookerResponse)
				}()
				//if gtime.New(PmsAppReservationChange.OldCheckinDate) != gtime.New(PmsAppReservationChange.NewCheckinDate) {
				//	// 失效
				//	err = service.HotelService().HotelOrderAwardInvalid(ctx, tx, gvar.New(PmsAppStay.Id).Int())
				//	// 发放下单奖励 优惠券和礼品券
				//	err = service.HotelService().HotelOrderAward(ctx, tx, gvar.New(PmsAppStay.Id).Int(), PmsAppReservationChange.NewCheckinDate.String(), PmsAppStay.MemberId)
				//}

				return
			}
			IsFx := PmsAppReservationChange.IsFx
			if g.IsEmpty(IsFx) {
				IsFx = "N"
			}

			// 插入支付订单流水
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsTransaction{
				OrderSn:       PmsAppReservationChange.OrderSn,
				ChangeOrderSn: req.ChangeOrderSn,
				TransactionSn: TransactionSn,
				PayChannel:    "PAYCLOUD",
				PayType:       "WeChatPay",
				OrderType:     "BOOKING_CHANGE",
				Scene:         "HOTEL",
				Amount:        PmsAppReservationChange.ChangeAmount,
				PriceCurrency: "JPY",
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
				IsFx:          IsFx,
			}).InsertAndGetId(); err != nil {
				return
			}
			res.TransactionSn = TransactionSn
			res.CreateOrderTime = gtime.Now().Format("Y-m-d H:i:s")
			res.Countdown = PmsAppReservationChange.ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
			if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeDelayedName,
				QueueName:    consts.RabbitMQQueueNameOrderExpire,
				DataByte:     gvar.New(req.ChangeOrderSn).Bytes(),
			}); err != nil {
				g.Log().Error(ctx, "发送过期自动取消订单MQ失败", err)
			}
		}
		return
	}); err != nil {
		return
	}
	return
}

func (c *ControllerHotel) OrderChangeBookingPeopleInfo(ctx context.Context, req *hotel.OrderChangeBookingPeopleInfoReq) (res *hotel.OrderChangeBookingPeopleInfoRes, err error) {
	var (
		AppReservation *entity.PmsAppReservation
		RoomType       *entity.PmsRoomType
	)
	res = new(hotel.OrderChangeBookingPeopleInfoRes)
	if err = dao.PmsAppReservation.Ctx(ctx).WherePri(req.OrderId).WithAll().Scan(&AppReservation); err != nil {
		return
	}
	res.Child = AppReservation.ChildCount
	res.Adult = AppReservation.AdultCount
	res.Infant = AppReservation.InfantCount

	if err = dao.PmsRoomType.Ctx(ctx).
		Where(dao.PmsRoomType.Columns().Uid, AppReservation.RoomType).
		Scan(&RoomType); err != nil {
		return
	}

	if g.IsEmpty(RoomType) {
		// 暂不支持变更入住人信息
		err = gerror.New(gi18n.T(ctx, "not_support_change_checkin_info"))
		return
	}

	res.Occupancy = RoomType.Occupancy

	return
}

func (c *ControllerHotel) OrderChangeStayOnInfo(ctx context.Context, req *hotel.OrderChangeStayOnInfoReq) (res *hotel.OrderChangeStayOnInfoRes, err error) {
	var (
		AppReservation      *entity.PmsAppReservation
		PmsAvailabilities   *entity.PmsAvailabilities
		RoomType            *entity.PmsRoomType
		PreCreateOrderInp   = new(input_hotel.PreMoreCreateOrderInp)
		PreCreateOrderModel = new(input_hotel.PreCreateOrderModel)
		StartDate           *gtime.Time
		EndDate             *gtime.Time
		RoomItems           *input_hotel.RoomItems
		PmsPirceConfig      *model.PmsPriceConfig
	)
	res = new(hotel.OrderChangeStayOnInfoRes)
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	if err = dao.PmsAppReservation.Ctx(ctx).WherePri(req.OrderId).WithAll().Scan(&AppReservation); err != nil {
		return
	}
	if g.IsEmpty(AppReservation) {
		// 不存在入住单
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}
	if err = dao.PmsRoomType.Ctx(ctx).
		Where(dao.PmsRoomType.Columns().Uid, AppReservation.RoomType).
		Hook(hook.PmsFindLanguageValueHook).
		Scan(&RoomType); err != nil {
		return
	}
	if g.IsEmpty(RoomType) {
		// 不存在房型
		err = gerror.New(gi18n.T(ctx, "no_room_type"))
		return
	}

	res.Cover = RoomType.Cover
	res.Name = RoomType.Name
	StartDate = AppReservation.CheckoutDate
	EndDate = StartDate.Add(time.Hour * 24 * time.Duration(req.NightNum))
	res.CheckOutDate = EndDate.Format("Y-m-d")

	// 批量查询物业在时间段中的最低价格
	if err = dao.PmsAvailabilities.Ctx(ctx).
		Fields(dao.PmsAvailabilities.Columns().Tuid, "min(price) as price", "min(allotment) as allotment").
		Where(dao.PmsAvailabilities.Columns().Tuid, AppReservation.RoomType).
		WhereGT(dao.PmsAvailabilities.Columns().Price, 0).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, StartDate).
		WhereLT(dao.PmsAvailabilities.Columns().Date, EndDate).
		Group(dao.PmsAvailabilities.Columns().Tuid).Scan(&PmsAvailabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAvailabilities) {
		// 暂不支持续住
		err = gerror.New(gi18n.T(ctx, "not_supported_stay_on"))
		return
	}

	if PmsAvailabilities.Allotment < 1 {
		// 您选择的日期已经没有满足条件的房间库存了
		err = gerror.NewCode(gcode.New(1111, gi18n.T(ctx, "the_date_no_stock"), nil))
		return
	}
	res.BasePrice = decimal.NewFromFloat(PmsAvailabilities.Price).Mul(PricePercent).Round(0).InexactFloat64()
	res.Inventory = PmsAvailabilities.Allotment

	PreCreateOrderInp.Puid = AppReservation.Puid
	RoomItems = new(input_hotel.RoomItems)
	RoomItems = &input_hotel.RoomItems{
		RoomId:      AppReservation.RoomType,
		RoomNoId:    AppReservation.RoomUnit,
		RoomNum:     1,
		Adult:       AppReservation.AdultCount,
		Child:       AppReservation.ChildCount,
		Infant:      AppReservation.InfantCount,
		BookingFee:  0,
		Remark:      "",
		PricePlanId: 0,
	}
	PreCreateOrderInp = &input_hotel.PreMoreCreateOrderInp{
		Puid:              AppReservation.Puid,
		CheckinAt:         StartDate.Format("Y-m-d"),
		CheckoutAt:        EndDate.Format("Y-m-d"),
		RoomItems:         []*input_hotel.RoomItems{RoomItems},
		IsCheckDaysNotice: true,
	}
	if PreCreateOrderModel, err = service.HotelService().PreMorePricePlanOrder(ctx, PreCreateOrderInp); err != nil {
		return
	}
	g.Log().Info(ctx, gjson.New(PreCreateOrderModel).String())
	res.PreOrderSn = PreCreateOrderModel.PreOrderSn
	// 查询最长可续住天数
	PmsAvailabilities = nil
	if err = dao.PmsAvailabilities.Ctx(ctx).
		Where(dao.PmsAvailabilities.Columns().Tuid, AppReservation.RoomType).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, StartDate).
		WhereGTE(dao.PmsAvailabilities.Columns().Allotment, 1).
		OrderDesc(dao.PmsAvailabilities.Columns().Date).
		Scan(&PmsAvailabilities); err != nil {
		return
	}
	if g.IsEmpty(PmsAvailabilities) {
		if err = dao.PmsAvailabilities.Ctx(ctx).
			Where(dao.PmsAvailabilities.Columns().Tuid, AppReservation.RoomType).
			WhereGTE(dao.PmsAvailabilities.Columns().Date, StartDate).
			OrderDesc(dao.PmsAvailabilities.Columns().Date).
			Scan(&PmsAvailabilities); err != nil {
			return
		}
		if g.IsEmpty(PmsAvailabilities) {
			// 暂不支持续住
			err = gerror.New(gi18n.T(ctx, "not_supported_stay_on"))
			return
		}
	}
	PmsAvailabilitiesDate := gtime.NewFromStr(PmsAvailabilities.Date)
	res.MaxStayNoDays = int(PmsAvailabilitiesDate.Sub(StartDate).Hours() / 24)

	//res.MaxStayNoDays = convert.Diff(StartDate.Format("Y-m-d"), PmsAvailabilities.Date, "day")
	return
}

func (c *ControllerHotel) OrderChangeInfoDetail(ctx context.Context, req *hotel.OrderChangeInfoDetailReq) (res *hotel.OrderChangeInfoDetailRes, err error) {
	res = new(hotel.OrderChangeInfoDetailRes)

	if err = dao.PmsAppReservationChange.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservationChange.Columns().ChangeOrderSn: req.ChangeOrderSn,
	}).Scan(&res); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	res.Countdown = res.ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
	return
}

// updateChargesForDateChange 更新 Charge 表：删除不在新日期范围内的记录，添加新日期的记录
func updateChargesForDateChange(
	ctx context.Context,
	tx gdb.TX,
	pmsAppReservation *entity.PmsAppReservation,
	pmsRoomType *entity.PmsRoomType,
	oldCharges []*entity.PmsCharge,
	availabilities []*entity.PmsAvailabilities,
	changeInfo *entity.PmsAppReservationChange,
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
