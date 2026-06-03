package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// PreOrder 预下单
func (s *sHotelService) PreOrder(ctx context.Context, in *input_hotel.PreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error) {
	var (
		preOrderDetail = new(input_hotel.PreOrderDetailModel)
		PropertyView   *entity.PmsProperty
		MemberInfo     *model.MemberIdentity
		PayInfo        *input_pay.PayInfoModel
		CancelRate     []*entity.PmsCancelRate
		basePrice      float64
		OrderAmount    float64
		PmsRoomType    *entity.PmsRoomType
		Availabilities []*entity.PmsAvailabilities
		PmsPirceConfig *model.PmsPriceConfig
	)
	g.Log().Debug(ctx, "overdueoverdueoverdue")
	g.Log().Debug(ctx, convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days"))
	if convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days") > 179 {
		err = gerror.New(gi18n.T(ctx, "stay_over_179_days_cannot_book"))
		return
	}

	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	out = new(input_hotel.PreCreateOrderModel)
	out.PreOrderSn = uuid.CreatePayCode("PRE")
	// 校验如果是短租是否满足当前短租最短入住天数
	// 物业信息
	if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsProperty.Columns().Uid, in.Puid).Scan(&PropertyView); err != nil {
		return
	}
	if g.IsEmpty(PropertyView) {
		// 物业不存在
		err = gerror.New(gi18n.T(ctx, "property_does_not_exist"))
		return
	}
	// 判断最小预定区间是否在预定时间区间内
	// 当前最小预定区间  如果小与2 不做校验  如果大于2 才做校验
	if PropertyView.MinDaysNotice >= 2 && !in.IsCheckDaysNotice {
		if convert.Diff(in.CheckinAt, in.CheckoutAt, "days") < PropertyView.MinDaysNotice {
			err = gerror.Newf(gi18n.T(ctx, "min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}
	// 生成预订单详情
	for k, v := range in.RoomItems {
		if g.IsEmpty(in.RoomItems[k].Adult) {
			in.RoomItems[k].Adult = 1
		}
		// 校验库存
		if err = service.HotelService().CheckHotelInventory(ctx, &input_hotel.CheckHotelInventoryInp{
			TUID:      v.RoomId,
			StartDate: in.CheckinAt,
			EndDate:   in.CheckoutAt,
			Num:       v.RoomNum,
			PUID:      in.Puid,
		}); err != nil {
			return
		}
		RoomsItems := &input_hotel.OrderRooms{
			RoomId:   v.RoomId,
			RoomNum:  v.RoomNum,
			RoomNoId: v.RoomNoId,
			Adult:    v.Adult,
			Child:    v.Child,
			Infant:   v.Infant,
		}

		// 查询房型信息
		PmsRoomType = nil
		if err = dao.PmsRoomType.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsRoomType.Columns().Uid, v.RoomId).Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(PmsRoomType) {
			// 房型不存在
			err = gerror.New(gi18n.T(ctx, "room_type_does_not_exist"))
			return
		}
		RoomsItems.Tid = PmsRoomType.Id
		RoomsItems.Puid = PmsRoomType.Puid
		RoomsItems.Cover = PmsRoomType.Cover
		RoomsItems.Name = PmsRoomType.Name

		Availabilities = nil
		if err = dao.PmsAvailabilities.Ctx(ctx).
			Where(dao.PmsAvailabilities.Columns().Tuid, v.RoomId).
			WhereGTE(dao.PmsAvailabilities.Columns().Date, in.CheckinAt).
			WhereLT(dao.PmsAvailabilities.Columns().Date, in.CheckoutAt).
			Scan(&Availabilities); err != nil {
			return
		}

		// 创建费率计划
		for _, Availabilitie := range Availabilities {
			if !g.IsEmpty(Availabilitie.Price) {

				//basePrice = Availabilitie.Price
				basePrice = decimal.NewFromFloat(Availabilitie.Price).Mul(PricePercent).Round(0).InexactFloat64()
			} else {
				basePrice = PmsRoomType.AdditionalGuestAmounts
			}

			RoomItemNum := 0
			RoomItemNum = v.Child + v.Adult + v.Infant
			OccupantsAmountRes := gvar.New(0).Float64()
			if RoomItemNum > PmsRoomType.OccupantsForBaseRate {
				OccupantsAmount := decimal.NewFromInt(gvar.New(RoomItemNum - PmsRoomType.OccupantsForBaseRate).Int64())
				OccupantsAmountRes = OccupantsAmount.Mul(decimal.NewFromFloat(PmsRoomType.AdditionalGuestAmounts)).Round(0).InexactFloat64()
				//RoomsItems.Charges = append(RoomsItems.Charges, &input_hotel.Charges{
				//	UID:         uuid.GetUuid(),
				//	Date:        Availabilitie.Date,
				//	FeeType:     "booking_fee",
				//	Amount:      OccupantsAmountRes,
				//	Description: "每日费用",
				//})
				RoomsItems.BookingFee += OccupantsAmountRes
			}
			RoomsItems.Charges = append(RoomsItems.Charges, &input_hotel.Charges{
				UID:         uuid.GetUuid(),
				Date:        Availabilitie.Date,
				FeeType:     "booking_fee",
				Amount:      basePrice + OccupantsAmountRes,
				Description: "每日费用",
			})
			RoomsItems.BookingFee += basePrice
		}
		OrderAmount = decimal.NewFromFloat(RoomsItems.BookingFee).Add(decimal.NewFromFloat(OrderAmount)).Round(0).InexactFloat64()

		preOrderDetail.Rooms = append(preOrderDetail.Rooms, RoomsItems)
	}
	// 会员ID
	MemberInfo = contexts.GetMemberUser(ctx)
	preOrderDetail.MemberId = MemberInfo.Id
	// 预订单号
	preOrderDetail.PreOrderSn = out.PreOrderSn

	if in.IsLease == 2 && !in.IsCheckDaysNotice {
		if PropertyView.LeaseClose != 1 {
			// 当前物业未开启短租模式
			err = gerror.New(gi18n.T(ctx, "has_not_activated_short_rent_mode"))
			return
		}

		days := convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
		if days < PropertyView.MinDaysNotice {
			// 当前短租最短入住天数为%d天
			err = gerror.Newf(gi18n.T(ctx, "short_rent_min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}

	preOrderDetail.Property.Id = PropertyView.Id
	preOrderDetail.Property.Uid = PropertyView.Uid
	preOrderDetail.Property.Name = PropertyView.Name
	preOrderDetail.Property.Address = PropertyView.Address
	preOrderDetail.Property.Cover = PropertyView.Cover
	preOrderDetail.Property.RequiredBook = PropertyView.RequiredBook
	// 入住时间信息
	preOrderDetail.PreDate.StartDate = in.CheckinAt
	preOrderDetail.PreDate.StartTime = PropertyView.CheckinAt
	preOrderDetail.PreDate.EndDate = in.CheckoutAt
	preOrderDetail.PreDate.EndTime = PropertyView.CheckoutAt
	preOrderDetail.PreDate.Days = convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
	// 入住人信息
	preOrderDetail.PreUser.FullName = MemberInfo.FullName
	preOrderDetail.PreUser.Phone = MemberInfo.Phone
	preOrderDetail.PreUser.Mail = MemberInfo.Mail
	preOrderDetail.PreUser.PhoneArea = MemberInfo.PhoneArea
	// 支付信息 默认不使用积分进行支付
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		OrderAmount: OrderAmount,
		MemberId:    MemberInfo.Id,
		IsBalance:   false,
	}); err != nil {
		return
	}
	preOrderDetail.PayInfo = new(input_hotel.OrderPayInfoModel)
	preOrderDetail.PayInfo.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
	preOrderDetail.PayInfo.PreOrderSn = out.PreOrderSn
	preOrderDetail.PayInfo.AllAmount = OrderAmount
	preOrderDetail.PayInfo.Score = PayInfo.Score
	preOrderDetail.PayInfo.MemberBalance = PayInfo.MemberBalance
	preOrderDetail.PayInfo.PayModel = PayInfo.PayModel
	preOrderDetail.PayInfo.Balance.BalanceAmount = PayInfo.BalanceAmount
	preOrderDetail.PayInfo.ThirdPay.ThirdAmount = PayInfo.ThirdAmount
	preOrderDetail.PayInfo.Balance.BalanceConfig.ScenePayRate = PayInfo.BalanceConfig.ScenePayRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.ExchangeRate = PayInfo.BalanceConfig.ExchangeRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.LevelName = PayInfo.BalanceConfig.LevelName
	preOrderDetail.PayInfo.Balance.BalanceConfig.Level = PayInfo.BalanceConfig.Level
	NowTime := gtime.Now().Format("Y-m-d")
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	for _, v := range CancelRate {
		switch v.Mode {
		case "after":
			startDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			//res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			if len(preOrderDetail.CancelDate) >= 1 {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			} else {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s", v.Name))
			}
		case "middle":
			startDay := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if NowTime <= endDate {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", fmt.Sprintf("%s ~ %s", startDay, endDate), v.Name))
			}
			break
		case "before":
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", endDate, v.Name))
			}
			break
		}
	}

	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, preOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}

// PrePricePlanOrder 价格计划预下单
func (s *sHotelService) PrePricePlanOrder(ctx context.Context, in *input_hotel.PreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error) {
	var (
		preOrderDetail = new(input_hotel.PreOrderDetailModel)
		PropertyView   *entity.PmsProperty
		MemberInfo     *model.MemberIdentity
		PayInfo        *input_pay.PayInfoModel
		CancelRate     []*struct {
			Name      string `json:"name"      dc:"规则名"`
			Mode      string `json:"mode"      dc:"规则模式"`
			StartDays int    `json:"startDays" dc:"开始天数"`
			EndDays   int    `json:"endDays"   dc:"结束天数"`
			Rate      int    `json:"rate"      dc:"取消费率"`
			Date      string `json:"date"      dc:"规则解析日期"`
			Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
		}
		OrderAmount       float64
		BookingPriceModel *input_pay.BookingPriceModel
	)
	out = new(input_hotel.PreCreateOrderModel)
	g.Log().Debug(ctx, "overdueoverdueoverdue")
	g.Log().Debug(ctx, convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days"))
	if convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days") > 179 {
		err = gerror.New(gi18n.T(ctx, "stay_over_179_days_cannot_book"))
		return
	}
	out.PreOrderSn = uuid.CreatePayCode("PRE")
	// 物业信息
	if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsProperty.Columns().Uid, in.Puid).Scan(&PropertyView); err != nil {
		return
	}

	if g.IsEmpty(PropertyView) {
		// 物业不存在
		err = gerror.New(gi18n.T(ctx, "property_does_not_exist"))
		return
	}
	if in.IsLease == 2 && !in.IsCheckDaysNotice {
		if PropertyView.LeaseClose != 1 {
			// 当前物业未开启短租模式
			err = gerror.New(gi18n.T(ctx, "has_not_activated_short_rent_mode"))
			return
		}

		days := convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
		if days < PropertyView.MinDaysNotice {
			// 当前短租最短入住天数为%d天
			err = gerror.Newf(gi18n.T(ctx, "short_rent_min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}
	// 校验如果是短租是否满足当前短租最短入住天数
	// 判断最小预定区间是否在预定时间区间内
	// 当前最小预定区间  如果小与2 不做校验  如果大于2 才做校验
	if PropertyView.MinDaysNotice >= 2 {
		if convert.Diff(in.CheckinAt, in.CheckoutAt, "days") < PropertyView.MinDaysNotice {
			err = gerror.Newf(gi18n.T(ctx, "min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}
	// 生成预订单详情 支持多房间下单
	for k, v := range in.RoomItems {
		var PmsRoomType *entity.PmsRoomType
		// 获取房型信息
		if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, v.RoomId).Scan(&PmsRoomType); err != nil {
			return
		}
		if PmsRoomType.IsShow != 1 {
			err = gerror.New(gi18n.T(ctx, "room_type_is_stopped"))
			return
		}

		// 判断总人数是否超过Occupancy
		if PmsRoomType.Occupancy > 0 && (v.Adult+v.Child) > PmsRoomType.Occupancy {
			err = gerror.Newf(gi18n.T(ctx, "room_occupancy_exceeded"), PmsRoomType.Occupancy)
			return
		}

		if v.PricePlanId > 0 {
			in.PricePlanId = v.PricePlanId
		}
		// 更具价格plan计算价格
		if BookingPriceModel, err = service.PayService().BookingPrice(ctx, &input_pay.BookingPriceInp{
			RoomInfos:   v,
			CheckinAt:   in.CheckinAt,
			CheckoutAt:  in.CheckoutAt,
			Puid:        in.Puid,
			PricePlanId: in.PricePlanId,
			IsPricePlan: true,
		}); err != nil {
			return
		}
		g.Log().Debug(ctx, "BookingPriceModel", BookingPriceModel)
		// 统计计算结果
		preOrderDetail.Rooms = append(preOrderDetail.Rooms, BookingPriceModel.OrderRooms)
		// 如果是默认的价格plan
		if len(BookingPriceModel.PlanPrice) > 0 {

			if g.IsEmpty(BookingPriceModel.PlanPrice[0].PmsPricePlan) {
				itemPricePlan := &entity.PmsPricePlan{
					PlanName:     "default",
					PlanShowName: "灵活取消",
					PropertyId:   in.Puid,
					RoomTypeId:   v.RoomId,
					IsCancel:     "Y",
				}
				preOrderDetail.Rooms[k].PricePlan = itemPricePlan
				preOrderDetail.PricePlan = itemPricePlan
			} else {
				// 计算单间总价
				preOrderDetail.Rooms[k].PricePlan = BookingPriceModel.PlanPrice[0].PmsPricePlan
				preOrderDetail.PricePlan = BookingPriceModel.PlanPrice[0].PmsPricePlan
			}

			// 叠加订单总价
			preOrderDetail.TotalAmount += BookingPriceModel.OrderRooms.OrderAmount
			// 叠加订单优惠等信息
			preOrderDetail.ChangeAmount += BookingPriceModel.OrderRooms.ChangeAmount
			// 计算订单总价
			OrderAmount = BookingPriceModel.OrderRooms.OrderAmount
		} else {
			// 叠加订单总价
			preOrderDetail.TotalAmount += BookingPriceModel.TotalPrice
			// 计算订单总价
			OrderAmount = decimal.NewFromFloat(BookingPriceModel.TotalPrice).Add(decimal.NewFromFloat(OrderAmount)).Round(0).InexactFloat64()
			preOrderDetail.Rooms[k].BookingFee = BookingPriceModel.TotalPrice
			itemPricePlan := &entity.PmsPricePlan{
				PlanName:     "default",
				PlanShowName: "灵活取消",
				PropertyId:   in.Puid,
				RoomTypeId:   v.RoomId,
				IsCancel:     "Y",
			}

			preOrderDetail.Rooms[k].PricePlan = itemPricePlan
			preOrderDetail.PricePlan = itemPricePlan
		}

	}
	// 会员ID
	MemberInfo = contexts.GetMemberUser(ctx)
	preOrderDetail.MemberId = MemberInfo.Id
	preOrderDetail.IsFx = "N"
	if MemberInfo.IsFx {
		preOrderDetail.IsFx = "Y"
	}

	// 预订单号
	preOrderDetail.PreOrderSn = out.PreOrderSn

	preOrderDetail.Property.Id = PropertyView.Id
	preOrderDetail.Property.Uid = PropertyView.Uid
	preOrderDetail.Property.Name = PropertyView.Name
	preOrderDetail.Property.Address = PropertyView.Address
	preOrderDetail.Property.Cover = PropertyView.Cover
	preOrderDetail.Property.RequiredBook = PropertyView.RequiredBook
	preOrderDetail.Property.GgLat = PropertyView.GgLat
	preOrderDetail.Property.GgLng = PropertyView.GgLng
	// 入住时间信息
	preOrderDetail.PreDate.StartDate = in.CheckinAt
	preOrderDetail.PreDate.StartTime = PropertyView.CheckinAt
	preOrderDetail.PreDate.EndDate = in.CheckoutAt
	preOrderDetail.PreDate.EndTime = PropertyView.CheckoutAt
	preOrderDetail.PreDate.Days = convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
	// 入住人信息
	preOrderDetail.PreUser.FullName = MemberInfo.FullName
	preOrderDetail.PreUser.Phone = MemberInfo.Phone
	preOrderDetail.PreUser.Mail = MemberInfo.Mail
	preOrderDetail.PreUser.PhoneArea = MemberInfo.PhoneArea
	// 支付信息 默认不使用积分进行支付
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		OrderAmount: OrderAmount,
		MemberId:    MemberInfo.Id,
		IsBalance:   false,
	}); err != nil {
		return
	}
	preOrderDetail.PayInfo = new(input_hotel.OrderPayInfoModel)
	preOrderDetail.PayInfo.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
	preOrderDetail.PayInfo.PreOrderSn = out.PreOrderSn
	preOrderDetail.PayInfo.AllAmount = OrderAmount
	preOrderDetail.PayInfo.Score = PayInfo.Score
	preOrderDetail.PayInfo.MemberBalance = PayInfo.MemberBalance
	preOrderDetail.PayInfo.PayModel = PayInfo.PayModel
	preOrderDetail.PayInfo.Balance.BalanceAmount = PayInfo.BalanceAmount
	preOrderDetail.PayInfo.ThirdPay.ThirdAmount = PayInfo.ThirdAmount
	preOrderDetail.PayInfo.Balance.BalanceConfig.ScenePayRate = PayInfo.BalanceConfig.ScenePayRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.ExchangeRate = PayInfo.BalanceConfig.ExchangeRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.LevelName = PayInfo.BalanceConfig.LevelName
	preOrderDetail.PayInfo.Balance.BalanceConfig.Level = PayInfo.BalanceConfig.Level
	NowTime := gtime.Now().Format("Y-m-d")
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	for k, v := range CancelRate {
		switch v.Mode {
		case "after":
			startDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			//res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			if len(preOrderDetail.CancelDate) >= 1 {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			} else {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s", v.Name))
			}
			if startDate <= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = startDate
			break
		case "middle":
			startDay := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if NowTime <= endDate {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", fmt.Sprintf("%s ~ %s", startDay, endDate), v.Name))
			}
			if startDay <= NowTime && endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
			break
		case "before":
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", endDate, v.Name))
			}
			if endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = endDate
			break
		}
	}
	preOrderDetail.CancelRate = CancelRate

	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, preOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}

// PreMorePricePlanOrder 多房间价格计划预下单
func (s *sHotelService) PreMorePricePlanOrder(ctx context.Context, in *input_hotel.PreMoreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error) {
	var (
		preOrderDetail = new(input_hotel.PreOrderDetailModel)
		PropertyView   *entity.PmsProperty
		MemberInfo     *model.MemberIdentity
		PayInfo        *input_pay.PayInfoModel
		CancelRate     []*struct {
			Name      string `json:"name"      dc:"规则名"`
			Mode      string `json:"mode"      dc:"规则模式"`
			StartDays int    `json:"startDays" dc:"开始天数"`
			EndDays   int    `json:"endDays"   dc:"结束天数"`
			Rate      int    `json:"rate"      dc:"取消费率"`
			Date      string `json:"date"      dc:"规则解析日期"`
			Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
		}
		OrderAmount       float64
		BookingPriceModel *input_pay.BookingPriceModel
	)
	out = new(input_hotel.PreCreateOrderModel)
	g.Log().Debug(ctx, "overdueoverdueoverdue")
	g.Log().Debug(ctx, convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days"))
	if convert.Diff(gtime.Now().Format("Y-m-d"), in.CheckoutAt, "days") > 179 {
		err = gerror.New(gi18n.T(ctx, "stay_over_179_days_cannot_book"))
		return
	}
	defer func() {
		g.Log().Debug(ctx, "PreOrderDetailModel", preOrderDetail)
	}()
	out.PreOrderSn = uuid.CreatePayCode("PRE")
	// 物业信息
	if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsProperty.Columns().Uid, in.Puid).Scan(&PropertyView); err != nil {
		return
	}

	if g.IsEmpty(PropertyView) {
		// 物业不存在
		err = gerror.New(gi18n.T(ctx, "property_does_not_exist"))
		return
	}
	if PropertyView.BatchReservation == "N" && !in.IsCheckDaysNotice {
		// 当前物业未开启多间预订模式
		err = gerror.New(gi18n.T(ctx, "has_not_activated_multiple_room_reservation_mode"))
		return
	}
	if in.IsLease == 2 && !in.IsCheckDaysNotice {
		if PropertyView.LeaseClose != 1 {
			// 当前物业未开启短租模式
			err = gerror.New(gi18n.T(ctx, "has_not_activated_short_rent_mode"))
			return
		}

		days := convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
		if days < PropertyView.MinDaysNotice {
			// 当前短租最短入住天数为%d天
			err = gerror.Newf(gi18n.T(ctx, "short_rent_min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}
	// 校验如果是短租是否满足当前短租最短入住天数
	// 判断最小预定区间是否在预定时间区间内
	// 当前最小预定区间  如果小与2 不做校验  如果大于2 才做校验
	if PropertyView.MinDaysNotice >= 2 && !in.IsCheckDaysNotice {
		if convert.Diff(in.CheckinAt, in.CheckoutAt, "days") < PropertyView.MinDaysNotice {
			err = gerror.Newf(gi18n.T(ctx, "min_stay_days"), PropertyView.MinDaysNotice)
			return
		}
	}
	// 生成预订单详情 支持多房间下单
	for k, v := range in.RoomItems {
		var PmsRoomType *entity.PmsRoomType
		// 获取房型信息
		if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, v.RoomId).Scan(&PmsRoomType); err != nil {
			return
		}
		if PmsRoomType.IsShow != 1 {
			err = gerror.New(gi18n.T(ctx, "room_type_is_stopped"))
			return
		}

		// 判断总人数是否超过Occupancy
		if PmsRoomType.Occupancy > 0 && (v.Adult+v.Child) > PmsRoomType.Occupancy {
			err = gerror.Newf(gi18n.T(ctx, "room_occupancy_exceeded_with_index"), k+1, PmsRoomType.Occupancy)
			return
		}

		// 更具价格plan计算价格
		if BookingPriceModel, err = service.PayService().BookingPrice(ctx, &input_pay.BookingPriceInp{
			RoomInfos:   v,
			CheckinAt:   in.CheckinAt,
			CheckoutAt:  in.CheckoutAt,
			Puid:        in.Puid,
			PricePlanId: v.PricePlanId,
			IsPricePlan: !in.IsCheckDaysNotice,
		}); err != nil {
			return
		}
		g.Log().Debug(ctx, "BookingPriceModel", BookingPriceModel)
		// 统计计算结果
		preOrderDetail.Rooms = append(preOrderDetail.Rooms, BookingPriceModel.OrderRooms)
		// 如果是默认的价格plan
		if len(BookingPriceModel.PlanPrice) > 0 {
			// 计算单间总价
			preOrderDetail.Rooms[k].PricePlan = BookingPriceModel.PlanPrice[0].PmsPricePlan
			preOrderDetail.PricePlan = BookingPriceModel.PlanPrice[0].PmsPricePlan
			// 叠加订单总价
			preOrderDetail.TotalAmount += BookingPriceModel.PlanPrice[0].TotalAmount
			// 叠加订单优惠等信息
			preOrderDetail.Rooms[k].ChangeAmount = BookingPriceModel.OrderRooms.ChangeAmount
			preOrderDetail.ChangeAmount += BookingPriceModel.OrderRooms.ChangeAmount
			// 计算订单总价
			OrderAmount = decimal.NewFromFloat(BookingPriceModel.PlanPrice[0].TotalAmount).Add(decimal.NewFromFloat(OrderAmount)).Round(0).InexactFloat64()
		} else {
			// 叠加订单总价
			preOrderDetail.TotalAmount += BookingPriceModel.TotalPrice
			// 计算订单总价
			OrderAmount = decimal.NewFromFloat(BookingPriceModel.TotalPrice).Add(decimal.NewFromFloat(OrderAmount)).Round(0).InexactFloat64()
			preOrderDetail.Rooms[k].BookingFee = BookingPriceModel.TotalPrice

			itemPricePlan := &entity.PmsPricePlan{
				PlanName:     "default",
				PlanShowName: "灵活取消",
				PropertyId:   in.Puid,
				RoomTypeId:   v.RoomId,
				IsCancel:     "Y",
			}

			preOrderDetail.Rooms[k].PricePlan = itemPricePlan
			preOrderDetail.PricePlan = itemPricePlan
		}

	}
	// 会员ID
	MemberInfo = contexts.GetMemberUser(ctx)
	preOrderDetail.MemberId = MemberInfo.Id
	// 预订单号
	preOrderDetail.PreOrderSn = out.PreOrderSn

	preOrderDetail.Property.Id = PropertyView.Id
	preOrderDetail.Property.Uid = PropertyView.Uid
	preOrderDetail.Property.Name = PropertyView.Name
	preOrderDetail.Property.Address = PropertyView.Address
	preOrderDetail.Property.Cover = PropertyView.Cover
	preOrderDetail.Property.RequiredBook = PropertyView.RequiredBook
	preOrderDetail.Property.GgLat = PropertyView.GgLat
	preOrderDetail.Property.GgLng = PropertyView.GgLng
	// 入住时间信息
	preOrderDetail.PreDate.StartDate = in.CheckinAt
	preOrderDetail.PreDate.StartTime = PropertyView.CheckinAt
	preOrderDetail.PreDate.EndDate = in.CheckoutAt
	preOrderDetail.PreDate.EndTime = PropertyView.CheckoutAt
	preOrderDetail.PreDate.Days = convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
	// 入住人信息
	preOrderDetail.PreUser.FullName = MemberInfo.FullName
	preOrderDetail.PreUser.Phone = MemberInfo.Phone
	preOrderDetail.PreUser.Mail = MemberInfo.Mail
	preOrderDetail.PreUser.PhoneArea = MemberInfo.PhoneArea
	// 支付信息 默认不使用积分进行支付
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		OrderAmount: OrderAmount,
		MemberId:    MemberInfo.Id,
		IsBalance:   false,
	}); err != nil {
		return
	}
	preOrderDetail.PayInfo = new(input_hotel.OrderPayInfoModel)
	preOrderDetail.PayInfo.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
	preOrderDetail.PayInfo.PreOrderSn = out.PreOrderSn
	preOrderDetail.PayInfo.AllAmount = OrderAmount
	preOrderDetail.PayInfo.Score = PayInfo.Score
	preOrderDetail.PayInfo.MemberBalance = PayInfo.MemberBalance
	preOrderDetail.PayInfo.PayModel = PayInfo.PayModel
	preOrderDetail.PayInfo.Balance.BalanceAmount = PayInfo.BalanceAmount
	preOrderDetail.PayInfo.ThirdPay.ThirdAmount = PayInfo.ThirdAmount
	preOrderDetail.PayInfo.Balance.BalanceConfig.ScenePayRate = PayInfo.BalanceConfig.ScenePayRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.ExchangeRate = PayInfo.BalanceConfig.ExchangeRate
	preOrderDetail.PayInfo.Balance.BalanceConfig.LevelName = PayInfo.BalanceConfig.LevelName
	preOrderDetail.PayInfo.Balance.BalanceConfig.Level = PayInfo.BalanceConfig.Level
	NowTime := gtime.Now().Format("Y-m-d")
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	for k, v := range CancelRate {
		switch v.Mode {
		case "after":
			startDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			//res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			if len(preOrderDetail.CancelDate) >= 1 {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
			} else {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s", v.Name))
			}
			if startDate <= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = startDate
			break
		case "middle":
			startDay := gtime.New(in.CheckinAt).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if NowTime <= endDate {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", fmt.Sprintf("%s ~ %s", startDay, endDate), v.Name))
			}
			if startDay <= NowTime && endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
			break
		case "before":
			endDate := gtime.New(in.CheckinAt).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				preOrderDetail.CancelDate = append(preOrderDetail.CancelDate, fmt.Sprintf("%s%s", endDate, v.Name))
			}
			if endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = endDate
			break
		}
	}
	preOrderDetail.CancelRate = CancelRate

	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, preOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}

// PreOrderDetail 预下单详情
func (s *sHotelService) PreOrderDetail(ctx context.Context, in *input_hotel.PreOrderDetailInp) (out *input_hotel.PreOrderDetailModel, err error) {
	var (
		PreOrderDetail *gvar.Var
	)
	if PreOrderDetail, err = cache.Instance().Get(ctx, "PreOrder_"+in.PreOrderSn); err != nil || PreOrderDetail.IsEmpty() {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	out = new(input_hotel.PreOrderDetailModel)
	if err = PreOrderDetail.Struct(&out); err != nil {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	return
}

// PrePayInfo 预下单重新计算支付金额支付模式
func (s *sHotelService) PrePayInfo(ctx context.Context, in *input_pay.PrePayInfoInp) (out *input_pay.PrePayInfoModel, err error) {
	var (
		PreOrderDetail *input_hotel.PreOrderDetailModel
		PayInfo        *input_pay.PayInfoModel
		IsBalance      bool
	)
	PreOrderDetail = new(input_hotel.PreOrderDetailModel)
	if PreOrderDetail, err = s.PreOrderDetail(ctx, &input_hotel.PreOrderDetailInp{PreOrderSn: in.PreOrderSn}); err != nil {
		return
	}
	PayInfo = new(input_pay.PayInfoModel)
	if in.IsBalance == 1 {
		IsBalance = true
	}
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		OrderAmount: PreOrderDetail.TotalAmount,
		MemberId:    PreOrderDetail.MemberId,
		IsBalance:   IsBalance,
		CouponId:    in.CouponId,
	}); err != nil {
		return
	}
	out = new(input_pay.PrePayInfoModel)
	out.OrderPayInfoModel = new(input_hotel.OrderPayInfoModel)
	out.OrderPayInfoModel.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
	out.OrderPayInfoModel.PreOrderSn = PreOrderDetail.PreOrderSn
	out.OrderPayInfoModel.AllAmount = PreOrderDetail.TotalAmount
	out.OrderPayInfoModel.Score = PayInfo.Score
	out.OrderPayInfoModel.MemberBalance = PayInfo.MemberBalance
	out.OrderPayInfoModel.PayModel = PayInfo.PayModel
	out.OrderPayInfoModel.Balance.BalanceAmount = PayInfo.BalanceAmount
	out.OrderPayInfoModel.ThirdPay.ThirdAmount = PayInfo.ThirdAmount
	out.OrderPayInfoModel.Balance.BalanceConfig.ScenePayRate = PayInfo.BalanceConfig.ScenePayRate
	out.OrderPayInfoModel.Balance.BalanceConfig.ExchangeRate = PayInfo.BalanceConfig.ExchangeRate
	out.OrderPayInfoModel.Balance.BalanceConfig.LevelName = PayInfo.BalanceConfig.LevelName
	out.OrderPayInfoModel.Balance.BalanceConfig.Level = PayInfo.BalanceConfig.Level
	out.OrderPayInfoModel.Coupon.CouponId = PayInfo.CouponId
	out.OrderPayInfoModel.Coupon.CouponAmount = PayInfo.CouponAmount
	out.OrderPayInfoModel.Coupon.CouponName = PayInfo.CouponName
	PreOrderDetail.PayInfo = out.OrderPayInfoModel
	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, PreOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}
