package logic_cabinet

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/cabinetApi"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateOrder 创建订单
func (s *sCabinetService) CreateOrder(ctx context.Context, in *input_cabinet.CreateOrderInp) (out *input_cabinet.CreateOrderModel, err error) {
	var (
		InsertId          int64
		OrderSn           = uuid.CreateOrderCode("B")
		PreOrderDetailVar *gvar.Var
		PreOrderDetail    *input_cabinet.PreOrderDetailModel
		MemberInfo        *model.MemberIdentity
		PmsMemberInfo     *entity.PmsMember
		ReferrerInfo      *entity.PmsMember
		ChannelInfo       *entity.PmsChannel
		StaffInfo         *entity.PmsStaff
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
	PreOrderDetail = new(input_cabinet.PreOrderDetailModel)
	if err = PreOrderDetailVar.Struct(&PreOrderDetail); err != nil {
		// 订单错误
		err = gerror.New(gi18n.T(ctx, "order_error"))
		return
	}
	// 校验
	if g.IsEmpty(PreOrderDetail.BoxTypeInfo) || g.IsEmpty(PreOrderDetail.BoxTypeInfo.Id) {
		err = gerror.New(gi18n.T(ctx, "please_select_the_format"))
		return
	}

	if g.IsEmpty(PreOrderDetail.ChooseHours) {
		err = gerror.New(gi18n.T(ctx, "please_select_the_rental_duration"))
		return
	}

	switch {
	case PreOrderDetail.CabinetInfo.BoxTypeA.ID == PreOrderDetail.BoxTypeInfo.Id:
		if PreOrderDetail.CabinetInfo.BoxTypeA.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}

	case PreOrderDetail.CabinetInfo.BoxTypeB.ID == PreOrderDetail.BoxTypeInfo.Id:
		if PreOrderDetail.CabinetInfo.BoxTypeB.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}

	case PreOrderDetail.CabinetInfo.BoxTypeC.ID == PreOrderDetail.BoxTypeInfo.Id:
		if PreOrderDetail.CabinetInfo.BoxTypeC.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}

	default:
		err = gerror.New(gi18n.T(ctx, "please_select_the_format"))
		return
	}

	if PreOrderDetail.ChooseHours < PreOrderDetail.CabinetInfo.MinHours {
		err = gerror.Newf(gi18n.T(ctx, "minimum_rental_time"), PreOrderDetail.CabinetInfo.MinHours)
		return
	}

	// 初始化响应变量
	out = new(input_cabinet.CreateOrderModel)
	out.OrderSn = OrderSn
	// 开启闭包事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 读取支付配置
		if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
			return
		}
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

		// 查询场景信息
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 5).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// 创建住宿预定
		CreateOrderTime := gtime.Now()

		BoxTypeJson := make([]*cabinetApi.BoxTypeItem, 0, 3)
		BoxTypeJson = append(BoxTypeJson, &cabinetApi.BoxTypeItem{
			ID:     PreOrderDetail.CabinetInfo.BoxTypeA.ID,
			Name:   PreOrderDetail.CabinetInfo.BoxTypeA.NameZh,
			NameZh: PreOrderDetail.CabinetInfo.BoxTypeA.NameZh,
			NameEn: PreOrderDetail.CabinetInfo.BoxTypeA.NameEn,
			NameJa: PreOrderDetail.CabinetInfo.BoxTypeA.NameJa,
			NameKo: PreOrderDetail.CabinetInfo.BoxTypeA.NameKo,
			NameTw: PreOrderDetail.CabinetInfo.BoxTypeA.NameTw,
			Price:  PreOrderDetail.CabinetInfo.BoxTypeA.Price,
			Num:    PreOrderDetail.CabinetInfo.BoxTypeA.Num,
		})
		BoxTypeJson = append(BoxTypeJson, &cabinetApi.BoxTypeItem{
			ID:     PreOrderDetail.CabinetInfo.BoxTypeB.ID,
			Name:   PreOrderDetail.CabinetInfo.BoxTypeB.NameZh,
			NameZh: PreOrderDetail.CabinetInfo.BoxTypeB.NameZh,
			NameEn: PreOrderDetail.CabinetInfo.BoxTypeB.NameEn,
			NameJa: PreOrderDetail.CabinetInfo.BoxTypeB.NameJa,
			NameKo: PreOrderDetail.CabinetInfo.BoxTypeB.NameKo,
			NameTw: PreOrderDetail.CabinetInfo.BoxTypeB.NameTw,
			Price:  PreOrderDetail.CabinetInfo.BoxTypeB.Price,
			Num:    PreOrderDetail.CabinetInfo.BoxTypeB.Num,
		})
		BoxTypeJson = append(BoxTypeJson, &cabinetApi.BoxTypeItem{
			ID:     PreOrderDetail.CabinetInfo.BoxTypeC.ID,
			Name:   PreOrderDetail.CabinetInfo.BoxTypeC.NameZh,
			NameZh: PreOrderDetail.CabinetInfo.BoxTypeC.NameZh,
			NameEn: PreOrderDetail.CabinetInfo.BoxTypeC.NameEn,
			NameJa: PreOrderDetail.CabinetInfo.BoxTypeC.NameJa,
			NameKo: PreOrderDetail.CabinetInfo.BoxTypeC.NameKo,
			NameTw: PreOrderDetail.CabinetInfo.BoxTypeC.NameTw,
			Price:  PreOrderDetail.CabinetInfo.BoxTypeC.Price,
			Num:    PreOrderDetail.CabinetInfo.BoxTypeC.Num,
		})

		var (
			CabinetNameJson   *input_cabinet.LanguageJson
			CityNameJson      *input_cabinet.LanguageJson
			MchNameJson       *input_cabinet.LanguageJson
			MchBranchNameJson *input_cabinet.LanguageJson
			BoxTypeNameJson   *input_cabinet.LanguageJson
		)

		CabinetNameJson = &input_cabinet.LanguageJson{
			Zh: PreOrderDetail.CabinetInfo.NameZh,
			En: PreOrderDetail.CabinetInfo.NameEn,
			Ja: PreOrderDetail.CabinetInfo.NameJa,
			Ko: PreOrderDetail.CabinetInfo.NameKo,
			Tw: PreOrderDetail.CabinetInfo.NameTw,
		}
		CityNameJson = &input_cabinet.LanguageJson{
			Zh: PreOrderDetail.CabinetInfo.CityNameZh,
			En: PreOrderDetail.CabinetInfo.CityNameEn,
			Ja: PreOrderDetail.CabinetInfo.CityNameJa,
			Ko: PreOrderDetail.CabinetInfo.CityNameKo,
			Tw: PreOrderDetail.CabinetInfo.CityNameTw,
		}
		MchNameJson = &input_cabinet.LanguageJson{
			Zh: PreOrderDetail.CabinetInfo.MchNameZh,
			En: PreOrderDetail.CabinetInfo.MchNameEn,
			Ja: PreOrderDetail.CabinetInfo.MchNameJa,
			Ko: PreOrderDetail.CabinetInfo.MchNameKo,
			Tw: PreOrderDetail.CabinetInfo.MchNameTw,
		}
		MchBranchNameJson = &input_cabinet.LanguageJson{
			Zh: PreOrderDetail.CabinetInfo.MchBranchNameZh,
			En: PreOrderDetail.CabinetInfo.MchBranchNameEn,
			Ja: PreOrderDetail.CabinetInfo.MchBranchNameJa,
			Ko: PreOrderDetail.CabinetInfo.MchBranchNameKo,
			Tw: PreOrderDetail.CabinetInfo.MchBranchNameTw,
		}
		BoxTypeNameJson = &input_cabinet.LanguageJson{
			Zh: PreOrderDetail.BoxTypeInfo.NameZh,
			En: PreOrderDetail.BoxTypeInfo.NameEn,
			Ja: PreOrderDetail.BoxTypeInfo.NameJa,
			Ko: PreOrderDetail.BoxTypeInfo.NameKo,
			Tw: PreOrderDetail.BoxTypeInfo.NameTw,
		}

		OrderData := &entity.CabinetOrder{
			OrderSn:           OrderSn,
			MemberId:          uint(MemberInfo.Id),
			MinHours:          PreOrderDetail.CabinetInfo.MinHours,
			CabinetId:         PreOrderDetail.CabinetInfo.ID,
			CabinetName:       PreOrderDetail.CabinetInfo.NameZh,
			CabinetNameJson:   gjson.New(CabinetNameJson),
			CityId:            PreOrderDetail.CabinetInfo.CityId,
			CityName:          PreOrderDetail.CabinetInfo.CityNameZh,
			CityNameJson:      gjson.New(CityNameJson),
			MchId:             PreOrderDetail.CabinetInfo.MchId,
			MchName:           PreOrderDetail.CabinetInfo.MchNameZh,
			MchNameJson:       gjson.New(MchNameJson),
			MchBranchId:       PreOrderDetail.CabinetInfo.MchBranchId,
			MchBranchName:     PreOrderDetail.CabinetInfo.MchBranchNameZh,
			MchBranchNameJson: gjson.New(MchBranchNameJson),
			MchBranchLat:      PreOrderDetail.CabinetInfo.MchBranchLat,
			MchBranchLgt:      PreOrderDetail.CabinetInfo.MchBranchLgt,
			BoxTypeJson:       gjson.New(BoxTypeJson),
			BoxTypeId:         PreOrderDetail.BoxTypeInfo.Id,
			BoxTypeName:       PreOrderDetail.BoxTypeInfo.NameZh,
			BoxTypeNameJson:   gjson.New(BoxTypeNameJson),
			BoxTypePrice:      PreOrderDetail.BoxTypeInfo.Price,
			OrderFirstFeeRate: PreOrderDetail.CabinetInfo.OrderFirstFeeRate,
			BuyHours:          PreOrderDetail.ChooseHours,
			OrderAmount:       PreOrderDetail.PayInfo.AllAmount,
			BaseAmount:        PreOrderDetail.PayInfo.AllAmount,
			CouponAmount:      PreOrderDetail.PayInfo.Coupon.CouponAmount,
			BalAmount:         PreOrderDetail.PayInfo.Balance.BalanceAmount,
			PayModel:          PreOrderDetail.PayInfo.PayModel,
			ExpirationTime:    ExpirationTime,
			//StartTime:         CreateOrderTime,
			//EndTime:           CreateOrderTime.Add(time.Duration(PreOrderDetail.ChooseHours) * time.Hour),
			CreatedAt:   CreateOrderTime,
			IsPayOpen:   MemberScene.IsPayOpen,
			OrderStatus: "WAIT_PAY",
			IsFx:        PreOrderDetail.IsFx,
		}
		if g.IsEmpty(OrderData.IsFx) {
			OrderData.IsFx = "N"
			if MemberInfo.IsFx {
				OrderData.IsFx = "Y"
			}
		}
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return err
		}
		if YYConfig.RecommendModel == "FIRST" {
			OrderData.Referrer = PmsMemberInfo.LastReferrer
		} else {
			OrderData.Referrer = PmsMemberInfo.Referrer
		}
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, OrderData.Referrer).Scan(&ReferrerInfo); err != nil {
			return err
		}
		// 记录返佣比例
		if !g.IsEmpty(ReferrerInfo) {
			OrderData.RebateStatus = "WAIT"
			if ReferrerInfo.RebateMode == "CHANNEL" {
				if err = dao.PmsChannel.Ctx(ctx).WherePri(ReferrerInfo.ChannelId).Scan(&ChannelInfo); err != nil {
					return err
				}
				if !g.IsEmpty(ChannelInfo) {
					OrderData.RebateRate = ChannelInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "STAFF" {
				if err = dao.PmsStaff.Ctx(ctx).WherePri(ReferrerInfo.StaffId).Scan(&StaffInfo); err != nil {
					return err
				}
				if !g.IsEmpty(StaffInfo) {
					OrderData.RebateRate = StaffInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "MEMBER" {
				OrderData.RebateRate = YYConfig.MemberBrokerageRate
			} else {
				// 不存在该返佣模式
				return gerror.New(gi18n.T(ctx, "the_rebate_mode_does_not_exist"))
			}
		}
		// 计算返现比例
		if !g.IsEmpty(MemberScene) {
			OrderData.IsGetOpen = MemberScene.IsGetOpen
			if MemberScene.IsGetOpen == "Y" {
				//记录场景比例
				OrderData.CabinetGetRateScene = MemberScene.GetRate
				//记录会员等级倍率
				if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				if !g.IsEmpty(MemberLevel) {
					OrderData.CabinetGetRateVip = MemberLevel.CabinetGetRate
				}
			}
		}
		if InsertId, err = dao.CabinetOrder.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(OrderData); err != nil {
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
		out.Balance.BalanceConfig = new(input_cabinet.BalanceConfig)
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
				OrderType:     "CABINET",
				Scene:         "CABINET",
				TransactionSn: out.Balance.BalancePayOrderSn,
				PayType:       "BAL",
				Amount:        PreOrderDetail.PayInfo.Balance.BalanceAmount,
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
			}); err != nil {
				return err
			}
		}
		if out.ThirdPay.ThirdAmount > 0 {
			// 插入三方支付信息
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:          OrderSn,
				OrderType:        "CABINET",
				Scene:            "CABINET",
				TransactionSn:    out.ThirdPay.ThirdPayOrderSn,
				PaymentRequestId: "",
				PayChannel:       "paycloud",
				PayType:          "",
				PriceCurrency:    "JPY",
				Amount:           PreOrderDetail.PayInfo.ThirdPay.ThirdAmount,
				PayStatus:        "WAIT",
				ExpiredTime:      gtime.New(ExpirationTime),
			}); err != nil {
				return err
			}
		}
		if !g.IsEmpty(out.Coupon.CouponId) {
			// 插入优惠券参与支付详情
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:       OrderSn,
				OrderType:     "CABINET",
				Scene:         "CABINET",
				TransactionSn: out.Coupon.CouponPayOrderSn,
				PayChannel:    "SYSTEM",
				PayType:       "COUPON",
				PriceCurrency: "JPY",
				Amount:        PreOrderDetail.PayInfo.Coupon.CouponAmount,
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
				CouponId:      PreOrderDetail.PayInfo.Coupon.CouponId,
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
	out.Countdown = ExpirationTime - gvar.New(gtime.New(out.CreateOrderTime).Unix()).Int()

	// 订单下单日志
	if _, err = dao.CabinetOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
		OrderId:     int(InsertId),
		ActionWay:   "CREATE",
		Remark:      "订单创建",
		OperateType: "USER",
		OperateId:   MemberInfo.Id,
	}); err != nil {
		return
	}

	_, err = cache.Instance().Remove(ctx, "PreOrder_"+in.PreOrderSn)
	return
}
