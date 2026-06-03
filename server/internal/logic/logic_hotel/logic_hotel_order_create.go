package logic_hotel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"APT/utility/validate"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/shopspring/decimal"
)

// CreateHotelOrder 创建酒店订单
func (s *sHotelService) CreateHotelOrder(ctx context.Context, in *input_hotel.CreateOrderInp) (out *input_hotel.CreateOrderModel, err error) {
	var (
		InsertId          int64
		Booker            = uuid.GetUuid() // 新增预定人
		OrderSn           = uuid.CreateOrderCode("H")
		OutOrderSn        = uuid.CreateOrderCode("APP")
		PreOrderDetailVar *gvar.Var
		PreOrderDetail    *input_hotel.PreOrderDetailModel
		PmsProperty       *entity.PmsProperty
		PmsRoomType       *entity.PmsRoomType
		MemberInfo        *model.MemberIdentity
		PmsMemberInfo     *entity.PmsMember
		ReferrerInfo      *entity.PmsMember
		ChannelInfo       *entity.PmsChannel
		StaffInfo         *entity.PmsStaff
		ChargesInfo       []*entity.PmsCharge
		MemberScene       *entity.PmsMemberScene
		MemberLevel       *entity.PmsMemberLevel
		YYConfig          *model.YYConfig
		ExpirationTime    int
		PayConfig         *model.PayConfig
	)
	// 获取授权用户信息
	MemberInfo = contexts.GetMemberUser(ctx)
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberInfo.Id).Scan(&PmsMemberInfo); err != nil {
		return
	}
	if g.IsEmpty(PmsMemberInfo) {
		// 用户信息错误
		err = gerror.New(gi18n.T(ctx, "user_information_error"))
		return
	}
	// 读取预下单信息
	if PreOrderDetailVar, err = cache.Instance().Get(ctx, "PreOrder_"+in.PreOrderSn); err != nil || PreOrderDetailVar.IsEmpty() {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	PreOrderDetail = new(input_hotel.PreOrderDetailModel)
	if err = PreOrderDetailVar.Struct(&PreOrderDetail); err != nil {
		// 订单错误
		err = gerror.New(gi18n.T(ctx, "order_error"))
		return
	}
	g.Log().Info(ctx, "校验今日下单时间限制")
	g.Log().Info(ctx, PreOrderDetail.PreDate.StartDate)
	g.Log().Info(ctx, gtime.Now().Format("Y-m-d"))
	g.Log().Info(ctx, gtime.Now().Hour())
	// 校验今日下单时间限制
	if PreOrderDetail.PreDate.StartDate == gtime.Now().Format("Y-m-d") && gtime.Now().Hour() > 23 {
		// 今日下单时间限制,已无法下单
		err = gerror.New(gi18n.T(ctx, "today_order_time_limit"))
		return
	}

	if PreOrderDetail.PreDate.StartDate < gtime.Now().Format("Y-m-d") {
		// 下单时间限制,已无法下单
		err = gerror.New(gi18n.T(ctx, "order_time_limit"))
		return
	}

	// 校验手机号
	if !g.IsEmpty(in.User.Phone) && !g.IsEmpty(in.User.AreaNo) {
		// 使用手机号验证工具进行验证和标准化
		phoneResult := validate.IsPhoneNumberWithCountryCode(in.User.AreaNo, in.User.Phone)
		if !phoneResult.IsValid {
			// 手机号格式错误
			err = gerror.New(gi18n.T(ctx, "phone_format_error"))
			return
		}
	}

	// 初始化响应变量
	out = new(input_hotel.CreateOrderModel)
	out.ThirdPay.IsFxOrder = PreOrderDetail.IsFx == "Y"
	out.OrderSn = OrderSn
	out.RoomItems = PreOrderDetail.Rooms
	// 开启闭包事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 创建预订人信息
		if _, err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).OmitEmptyData().Insert(entity.PmsGuestProfile{
			Uid:         Booker,
			FirstName:   in.User.FirstName,
			LastName:    in.User.LastName,
			FullName:    fmt.Sprintf("%s %s", in.User.LastName, in.User.FirstName),
			Email:       in.User.Email,
			AreaNo:      in.User.AreaNo,
			Phone:       in.User.Phone,
			Nationality: MemberInfo.Nationality,
			MemberId:    MemberInfo.Id,
		}); err != nil {
			return err
		}
		// 是否初始化会员信息进行回写
		// 查询当前物业信息
		if err = dao.PmsProperty.Ctx(ctx).TX(tx).Where(dao.PmsProperty.Columns().Uid, PreOrderDetail.Property.Uid).Scan(&PmsProperty); err != nil && !errors.Is(err, sql.ErrNoRows) {
			// 不存在该物业
			return gerror.New(gi18n.T(ctx, "the_property_not_exist"))
		}
		ChargesAmount := gvar.New(0).Float64()
		// 先按房型汇总数量，再统一校验库存
		roomNumMap := make(map[string]int)
		for _, RoomItem := range PreOrderDetail.Rooms {
			roomNumMap[RoomItem.RoomId] += RoomItem.RoomNum
		}
		for roomId, totalNum := range roomNumMap {
			if err = s.CheckHotelInventory(ctx, &input_hotel.CheckHotelInventoryInp{
				StartDate: PreOrderDetail.PreDate.StartDate,
				EndDate:   PreOrderDetail.PreDate.EndDate,
				PUID:      PreOrderDetail.Property.Uid,
				TUID:      roomId,
				Num:       totalNum,
			}); err != nil {
				return err
			}
		}
		for RoomIndex, RoomItem := range PreOrderDetail.Rooms {
			PmsRoomType = new(entity.PmsRoomType)
			if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, RoomItem.RoomId).Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			appReservation := &entity.PmsAppReservation{
				OrderIndex:   RoomIndex + 1,
				Uid:          uuid.GetUuid(),
				MemberId:     MemberInfo.Id,
				Source:       "APP",
				Puid:         PreOrderDetail.Property.Uid,
				OrderSn:      OrderSn,
				OutOrderSn:   OutOrderSn,
				RoomType:     RoomItem.RoomId,
				RoomUnit:     RoomItem.RoomNoId,
				MainGuest:    Booker,
				CheckinDate:  gtime.New(PreOrderDetail.PreDate.StartDate),
				CheckoutDate: gtime.New(PreOrderDetail.PreDate.EndDate),
				CheckinTime:  PmsRoomType.CheckinAt,
				CheckoutTime: PmsRoomType.CheckoutAt,
				AdultCount:   RoomItem.Adult,
				ChildCount:   RoomItem.Child,
				InfantCount:  RoomItem.Infant,
				BookingFee:   RoomItem.OrderAmount,
				Charges:      uuid.GetUuid(),
				OrderStatus:  "WAIT_PAY",
				GuestRemarks: in.Remark,
				IsFx:         PreOrderDetail.IsFx,
			}
			if !g.IsEmpty(RoomItem.PricePlan) {
				appReservation.PricePlanId = RoomItem.PricePlan.Id
				appReservation.ChangeAmount = RoomItem.ChangeAmount
				appReservation.PricePlanInfo = gjson.New(RoomItem.PricePlan)
			}
			// 创建房间预定信息
			if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(appReservation); err != nil {
				return err
			}
			ChargesInfo = []*entity.PmsCharge{}
			// 插入费用明细表中
			for _, Charges := range RoomItem.Charges {
				ChargesAmount = decimal.NewFromInt(gvar.New(ChargesAmount).Int64()).Add(decimal.NewFromFloat(Charges.Amount)).Round(0).InexactFloat64()
				ChargesInfo = append(ChargesInfo, &entity.PmsCharge{
					Uid:            appReservation.Charges,
					AirUid:         guid.S([]byte("AirUid")),
					Date:           gtime.New(Charges.Date),
					Name:           consts.Charges[Charges.FeeType],
					FeeType:        Charges.FeeType,
					Amount:         Charges.Amount,
					OriginalAmount: Charges.Amount,
					Description:    gvar.New(Charges.Description).String(),
				})
			}

			if _, err = dao.PmsCharge.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ChargesInfo); err != nil {
				return err
			}
		}
		if ChargesAmount != PreOrderDetail.TotalAmount {
			g.Log().Info(ctx, gjson.New(PreOrderDetail).String())
			g.Log().Debugf(ctx, "订单金额异常 %f - %f", ChargesAmount, PreOrderDetail.TotalAmount)
			err = gerror.NewCode(gcode.New(gcode.CodeValidationFailed.Code(), "订单金额异常", nil))
			return
		}
		// 读取支付配置
		if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
			return
		}
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

		// 创建住宿预定
		CreateOrderTime := gtime.Now()
		PmsAppStayData := &entity.PmsAppStay{
			Uid:            uuid.GetUuid(),
			Source:         "APP",
			MemberId:       MemberInfo.Id,
			OrderSn:        OrderSn,
			OutOrderSn:     OutOrderSn,
			PayModel:       PreOrderDetail.PayInfo.PayModel,
			OrderAmount:    PreOrderDetail.TotalAmount,
			Puid:           PreOrderDetail.Property.Uid,
			ExpirationTime: ExpirationTime,
			Booker:         Booker,
			CheckInDate:    PreOrderDetail.PreDate.StartDate,
			CheckOutDate:   PreOrderDetail.PreDate.EndDate,
			CreatedAt:      CreateOrderTime,
			IsPayOpen:      PreOrderDetail.PayInfo.Balance.BalanceConfig.IsPayOpen,
			OrderStatus:    "WAIT_PAY",
			PricePercent:   PreOrderDetail.PricePercent,
			TotalAmount:    decimal.NewFromFloat(PreOrderDetail.PayInfo.AllAmount).Add(decimal.NewFromFloat(PreOrderDetail.ChangeAmount)).Round(0).InexactFloat64(),
			ChangeAmount:   PreOrderDetail.ChangeAmount,
			IsFx:           PreOrderDetail.IsFx,
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return err
		}
		if YYConfig.RecommendModel == "FIRST" {
			PmsAppStayData.Referrer = PmsMemberInfo.LastReferrer
		} else {
			PmsAppStayData.Referrer = PmsMemberInfo.Referrer
		}
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, PmsAppStayData.Referrer).Scan(&ReferrerInfo); err != nil {
			return err
		}
		// 记录返佣比例
		if !g.IsEmpty(ReferrerInfo) {
			PmsAppStayData.RebateStatus = "WAIT"
			if ReferrerInfo.RebateMode == "CHANNEL" {
				if err = dao.PmsChannel.Ctx(ctx).WherePri(ReferrerInfo.ChannelId).Scan(&ChannelInfo); err != nil {
					return err
				}
				if !g.IsEmpty(ChannelInfo) {
					PmsAppStayData.RebateRate = ChannelInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "STAFF" {
				if err = dao.PmsStaff.Ctx(ctx).WherePri(ReferrerInfo.StaffId).Scan(&StaffInfo); err != nil {
					return err
				}
				if !g.IsEmpty(StaffInfo) {
					PmsAppStayData.RebateRate = StaffInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "MEMBER" {
				PmsAppStayData.RebateRate = YYConfig.MemberBrokerageRate
			} else {
				// 不存在该返佣模式
				return gerror.New(gi18n.T(ctx, "the_rebate_mode_does_not_exist"))
			}
		}
		// 计算返现比例
		// 查询场景信息
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 1).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if !g.IsEmpty(MemberScene) {
			PmsAppStayData.IsGetOpen = MemberScene.IsGetOpen
			if MemberScene.IsGetOpen == "Y" {
				//记录酒店场景比例
				PmsAppStayData.HotelGetRateScene = MemberScene.GetRate
				//记录会员等级倍率
				if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				if !g.IsEmpty(MemberLevel) {
					PmsAppStayData.HotelGetRateVip = MemberLevel.HotelGetRate
				}
			}
		}
		if InsertId, err = dao.PmsAppStay.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(PmsAppStayData); err != nil {
			return err
		}
		out.CreateOrderTime = CreateOrderTime.Format("Y-m-d H:i:s")
		if g.IsEmpty(PreOrderDetail.PayInfo.Balance.BalancePayOrderSn) {
			PreOrderDetail.PayInfo.Balance.BalancePayOrderSn = uuid.CreateOrderCode("BL")
		}
		if g.IsEmpty(PreOrderDetail.PayInfo.Coupon.CouponPayOrderSn) {
			PreOrderDetail.PayInfo.Coupon.CouponPayOrderSn = uuid.CreateOrderCode("CP")
		}
		// 读取支付明细
		out.PayModel = PreOrderDetail.PayInfo.PayModel
		// 读取余额支付信息
		out.Balance.BalancePayOrderSn = PreOrderDetail.PayInfo.Balance.BalancePayOrderSn
		out.Balance.BalanceAmount = PreOrderDetail.PayInfo.Balance.BalanceAmount
		out.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
		out.Balance.BalanceConfig.ExchangeRate = PreOrderDetail.PayInfo.Balance.BalanceConfig.ExchangeRate
		out.Balance.BalanceConfig.ScenePayRate = PreOrderDetail.PayInfo.Balance.BalanceConfig.ScenePayRate
		out.Balance.BalanceConfig.Level = PreOrderDetail.PayInfo.Balance.BalanceConfig.Level
		out.Balance.BalanceConfig.LevelName = PreOrderDetail.PayInfo.Balance.BalanceConfig.LevelName
		out.Balance.BalanceConfig.IsPayOpen = PreOrderDetail.PayInfo.Balance.BalanceConfig.IsPayOpen
		// 读取第三方支付信息
		out.ThirdPay.ThirdPayOrderSn = PreOrderDetail.PayInfo.ThirdPay.ThirdPayOrderSn
		out.ThirdPay.ThirdAmount = PreOrderDetail.PayInfo.ThirdPay.ThirdAmount
		// 读取优惠券信息
		out.Coupon.CouponId = PreOrderDetail.PayInfo.Coupon.CouponId
		out.Coupon.CouponAmount = PreOrderDetail.PayInfo.Coupon.CouponAmount
		out.Coupon.CouponName = PreOrderDetail.PayInfo.Coupon.CouponName
		out.Coupon.CouponPayOrderSn = PreOrderDetail.PayInfo.Coupon.CouponPayOrderSn
		// 插入数据库支付信息内容
		if out.Balance.BalanceAmount > 0 {
			// 插入余额支付信息
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:       OrderSn,
				TransactionSn: out.Balance.BalancePayOrderSn,
				PayType:       "BAL",
				Amount:        PreOrderDetail.PayInfo.Balance.BalanceAmount,
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
				IsFx:          PreOrderDetail.IsFx,
			}); err != nil {
				return err
			}
		}
		if out.ThirdPay.ThirdAmount > 0 {
			// 插入三方支付信息
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:          OrderSn,
				TransactionSn:    out.ThirdPay.ThirdPayOrderSn,
				PaymentRequestId: "",
				PayChannel:       "paycloud",
				PayType:          "",
				PriceCurrency:    "JPY",
				Amount:           PreOrderDetail.PayInfo.ThirdPay.ThirdAmount,
				PayStatus:        "WAIT",
				ExpiredTime:      gtime.New(ExpirationTime),
				IsFx:             PreOrderDetail.IsFx,
			}); err != nil {
				return err
			}
		}
		if !g.IsEmpty(out.Coupon.CouponId) {
			// 插入优惠券参与支付详情
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:       OrderSn,
				TransactionSn: out.Coupon.CouponPayOrderSn,
				PayChannel:    "SYSTEM",
				PayType:       "COUPON",
				PriceCurrency: "JPY",
				Amount:        PreOrderDetail.PayInfo.Coupon.CouponAmount,
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
				CouponId:      PreOrderDetail.PayInfo.Coupon.CouponId,
				IsFx:          PreOrderDetail.IsFx,
			}); err != nil {
				return err
			}
			// 锁定优惠券 - 直接使用优惠券
			// 核销优惠券
			if _, err = dao.PmsCoupon.Ctx(ctx).TX(tx).
				Where(dao.PmsCoupon.Columns().Id, PreOrderDetail.PayInfo.Coupon.CouponId).
				Update(g.MapStrAny{
					dao.PmsCoupon.Columns().State:   2,
					dao.PmsCoupon.Columns().UseTime: gtime.Now(),
				}); err != nil {
				return
			}
			// 修改支付订单状态
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
				Where(dao.PmsTransaction.Columns().TransactionSn, out.Coupon.CouponPayOrderSn).
				Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
				Update(entity.PmsTransaction{
					PayAmount: PreOrderDetail.PayInfo.Coupon.CouponAmount,
					PayStatus: "DONE",
					PayTime:   gtime.Now(),
				}); err != nil {
				return
			}
		}
		return err
	}); err != nil {
		return
	}
	// 投递自动过期订单队列 如果订单在规定时间内未支付 必须取消订单
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeDelayedName,
		QueueName:    consts.RabbitMQQueueNameOrderExpire,
		DataByte:     gvar.New(out.OrderSn).Bytes(),
	}); err != nil {
		g.Log().Error(ctx, "发送过期自动取消订单MQ失败", err)
	}
	out.Countdown = gvar.New(gtime.New(out.CreateOrderTime).Unix()).Int() - ExpirationTime

	// 酒店订单下单日志
	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     int(InsertId),
		ActionWay:   "CREATE",
		Remark:      "订单创建",
		OperateType: "USER",
		OperateId:   MemberInfo.Id,
	}); err != nil {
		return
	}

	// 发送到消息队列
	systemMessageTitle := map[string]string{
		"zh":    "订单创建成功",
		"en":    "Order created successfully",
		"ja":    "注文が正常に作成されました",
		"ko":    "주문이 성공적으로 생성되었습니다.",
		"zh_CN": "訂單創建成功",
	}
	systemMessageContent := map[string]string{
		"zh":    "您已成功下单，订单号为：" + OrderSn,
		"en":    "You have successfully placed an order, the order number is: " + OrderSn,
		"ja":    "注文が完了しました。注文番号は：" + OrderSn,
		"ko":    "주문이 완료되었습니다. 주문번호는：" + OrderSn,
		"zh_CN": "您已成功下單，訂單編號為：" + OrderSn,
	}
	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(MemberInfo.Id).Value()
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
		"orderSn": OrderSn,
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
		MemberId:             int(MemberInfo.Id),
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/order/order-detail",
		WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", OrderSn),
		EnablePush:           true,
		EnableSms:            false,
		PushTitle:            systemMessageTitle[memberLanguage],
		PushContent:          systemMessageContent[memberLanguage],
		OperatorId:           MemberInfo.Id,
		OperatorRole:         "MEMBER",
		OrderSn:              OrderSn,
	})
	return
}

// CreateStayParams 创建预订参数
func (s *sHotelService) CreateStayParams(ctx context.Context, OrderSn string) (out *airhousePublicApi.CreateStayJSONDataRequest, err error) {
	var (
		ReservationInfo []*entity.PmsAppReservation
		StayInfo        *entity.PmsAppStay
		PmsGuestProfile *entity.PmsGuestProfile
		ReservationItem *airhousePublicApi.RoomReservations
		ChargesItem     []*entity.PmsCharge
		ChargesItems    []airhousePublicApi.Charges
		Charges         = garray.NewStrArray()
	)
	g.Log().Info(ctx, "AirHouseAPISumberBookieParams     START")
	g.DB().SetLogger(g.Log())
	// 查询支付订单信息
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).Scan(&StayInfo); err != nil {
		return
	}
	if g.IsEmpty(StayInfo) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	// 查询房间订单信息
	if err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).Scan(&ReservationInfo); err != nil {
		return
	}
	if g.IsEmpty(ReservationInfo) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	out = new(airhousePublicApi.CreateStayJSONDataRequest)
	out.UID = StayInfo.Uid
	out.ReservationNumber = StayInfo.OutOrderSn

	for _, v := range ReservationInfo {
		Charges.Append(v.Charges)
	}

	// 查询charges
	if err = dao.PmsCharge.Ctx(ctx).
		WhereIn(dao.PmsCharge.Columns().Uid, Charges.Slice()).
		WhereGT(dao.PmsCharge.Columns().Amount, 0).
		//Where(dao.PmsCharge.Columns().FeeType, "booking_fee").
		Scan(&ChargesItem); err != nil {
		return
	}
	for _, ChargesValue := range ChargesItem {
		ChargesItems = append(ChargesItems, airhousePublicApi.Charges{
			UID:  gvar.New(ChargesValue.AirUid).String(),
			Date: gvar.New(ChargesValue.Date).String(),
			//FeeType:     gvar.New(ChargesValue.FeeType).String(),
			FeeType:     "booking_fee",
			Amount:      gvar.New(ChargesValue.Amount).Float64(),
			Description: gvar.New(ChargesValue.Description).String(),
		})
	}

	for index, v := range ReservationInfo {
		out.Property.ID = v.Puid
		if err = dao.PmsGuestProfile.Ctx(ctx).Where(dao.PmsGuestProfile.Columns().Uid, v.MainGuest).Scan(&PmsGuestProfile); err != nil {
			return
		}
		if g.IsEmpty(PmsGuestProfile) {
			// 预订人信息不存在
			err = gerror.New(gi18n.T(ctx, "booking_information_does_not_exist"))
			return
		}
		ReservationItem = new(airhousePublicApi.RoomReservations)
		ReservationItem.UID = v.Uid
		ReservationItem.Property.ID = v.Puid
		ReservationItem.RoomType.ID = v.RoomType
		if !g.IsEmpty(v.RoomUnit) {
			// 查询房间号
			//if err = dao.PmsRoomUnit.Ctx(ctx).Where(dao.PmsRoomUnit.Columns().Uid, v.RoomUnit).Scan(&RoomUnit); err != nil {
			//	return
			//}
			//if !g.IsEmpty(RoomUnit) {
			//
			//}
			ReservationItem.RoomUnit.ID = v.RoomUnit
		}

		ReservationItem.MainGuest.UID = v.MainGuest
		ReservationItem.MainGuest.FirstName = PmsGuestProfile.FirstName
		ReservationItem.MainGuest.LastName = PmsGuestProfile.LastName
		ReservationItem.MainGuest.FullName = PmsGuestProfile.FullName
		ReservationItem.MainGuest.Email = PmsGuestProfile.Email
		ReservationItem.MainGuest.Phone = fmt.Sprintf("%s%s", PmsGuestProfile.AreaNo, PmsGuestProfile.Phone)
		ReservationItem.MainGuest.Nationality = nil
		ReservationItem.MainGuest.Language = nil
		ReservationItem.MainGuest.Address = nil

		ReservationItem.CheckinDate = v.CheckinDate.Format("Y-m-d")
		ReservationItem.CheckoutDate = v.CheckoutDate.Format("Y-m-d")
		ReservationItem.CheckinTime = v.CheckinTime
		ReservationItem.CheckoutTime = v.CheckoutTime
		ReservationItem.Status = "confirmed"
		ReservationItem.CheckinStatus = "before_checkin"
		ReservationItem.AdultCount = v.AdultCount
		ReservationItem.ChildCount = v.ChildCount
		ReservationItem.InfantCount = v.InfantCount
		ReservationItem.GuestRemarks = v.GuestRemarks
		ReservationItem.BookingFee = v.BookingFee
		ReservationItem.ChannelFee = v.ChannelFee
		ReservationItem.CleaningFee = v.CleaningFee
		ReservationItem.TotalPaymentAmount = v.BookingFee
		ReservationItem.TotalChargeAmount = v.BookingFee
		if index == 0 {
			ReservationItem.PrepaidAmount = StayInfo.OrderAmount
			ReservationItem.Charges = ChargesItems
		} else {
			ReservationItem.Charges = make([]airhousePublicApi.Charges, 0)
		}

		ReservationItem.CancellationFee = v.CancellationFee

		out.RoomReservations = append(out.RoomReservations, ReservationItem)
	}

	out.Booker.UID = StayInfo.Booker
	out.Booker.FirstName = PmsGuestProfile.FirstName
	out.Booker.LastName = PmsGuestProfile.LastName
	out.Booker.FullName = PmsGuestProfile.FullName
	out.Booker.Email = PmsGuestProfile.Email
	out.Booker.Language = nil
	out.Booker.Phone = fmt.Sprintf("%s%s", PmsGuestProfile.AreaNo, PmsGuestProfile.Phone)
	out.Booker.Nationality = nil
	out.Booker.Address = nil
	g.Log().Info(ctx, "AirHouseAPISumberBookieParams", out)
	return
}

// CreateStay 提交预订
func (s *sHotelService) CreateStay(ctx context.Context, in *airhousePublicApi.CreateStayJSONDataRequest, tx gdb.TX) (BookerResponse *airhousePublicApi.CreateStayJSONDataResponse, err error) {
	g.DB().SetLogger(g.Log())
	if BookerResponse, err = airhousePublicApi.CreateStayPost(ctx, in); err != nil {
		g.Log().Error(ctx, err)
		// 订单提交失败
		err = gerror.New(gi18n.T(ctx, "order_submit_failed"))
		return
	}
	g.Log().Info(ctx, "airhost response", BookerResponse)
	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().Uid: BookerResponse.Data.UID,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		// 该订单无需处理
		err = gerror.New(gi18n.T(ctx, "order_need_not_handle"))
		return
	}
	for _, v := range BookerResponse.Data.RoomReservations {
		// 是否没有分配房间
		if g.IsEmpty(v.RoomUnit.UID) {
			g.Log().Info(ctx, "未分配房间")
		} else {
			// 回写房间信息
			if _, err = dao.PmsRoomStatus.Ctx(ctx).TX(tx).OmitEmptyWhere().
				Where(&entity.PmsRoomStatus{
					Puid: gvar.New(v.Property.UID).String(),
					Tuid: gvar.New(v.RoomType.UID).String(),
					Ruid: gvar.New(v.RoomUnit.UID).String(),
				}).
				WhereBetween(dao.PmsRoomStatus.Columns().Date, v.CheckinDate, v.CheckoutDate).
				Update(g.MapStrAny{
					dao.PmsRoomStatus.Columns().ReserveId: v.UID,
				}); err != nil {
				return
			}
		}

		if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsAppReservation.Columns().Uid: v.UID,
		}).Data(g.MapStrAny{
			dao.PmsAppReservation.Columns().Uuid:            v.ID,
			dao.PmsAppReservation.Columns().RoomUnit:        v.RoomUnit.ID,
			dao.PmsAppReservation.Columns().Status:          v.Status,
			dao.PmsAppReservation.Columns().AdultCount:      v.AdultCount,
			dao.PmsAppReservation.Columns().ChildCount:      v.ChildCount,
			dao.PmsAppReservation.Columns().InfantCount:     v.InfantCount,
			dao.PmsAppReservation.Columns().CancellationFee: v.CancellationFee,
			dao.PmsAppReservation.Columns().CheckinStatus:   v.CheckinStatus,
			dao.PmsAppReservation.Columns().RatePlanId:      v.Folio.ID,
		}).Update(); err != nil {
			return
		}

		var (
			ActionWay string
			Remark    string
		)
		if v.CheckinStatus == "checked_in" {
			ActionWay = "ORDER_CHECKED_IN"
			Remark = "客人已入住"
		} else if v.CheckinStatus == "checked_out" {
			ActionWay = "ORDER_CHECKED_OUT"
			Remark = "客人已退房"
		} else if v.Status == "blocked" {
			ActionWay = "ORDER_BLOCKED"
			Remark = "订单已阻止"
		} else if v.Status == "user_cancelled" {
			ActionWay = "ORDER_USER_CANCELLED"
			Remark = "订单已取消"
		} else if v.Status == "pending" {
			ActionWay = "ORDER_PENDING"
			Remark = "订单待确认"
		} else if v.Status == "overlapped" {
			ActionWay = "ORDER_OVERLAPPED"
			Remark = "订单重叠"
		} else if v.Status == "cancelled" {
			ActionWay = "ORDER_CANCELLED"
			Remark = "订单已取消"
		} else if v.Status == "confirmed" {
			ActionWay = "ORDER_CONFIRMED"
			Remark = "订单已确认"
		}

		var (
			PmsAppReservation entity.PmsAppReservation
		)
		if err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsAppReservation.Columns().Uid: v.UID,
		}).Scan(&PmsAppReservation); err != nil && errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(ActionWay) {
			if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
				OrderId:            PmsAppStay.Id,
				ReservationOrderId: PmsAppReservation.Id,
				ActionWay:          ActionWay,
				Remark:             Remark,
				OperateType:        "SYSTEM",
			}); err != nil {
				return
			}
		}
	}
	return
}
