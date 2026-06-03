package logic_cabinet

import (
	"APT/internal/dao"
	"APT/internal/library/cabinetApi"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"APT/utility/uuid"
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

// PreOrder 预下单
func (s *sCabinetService) PreOrder(ctx context.Context, in *input_cabinet.PreCreateOrderInp) (out *input_cabinet.PreCreateOrderModel, err error) {
	var (
		preOrderDetail   = new(input_cabinet.PreOrderDetailModel)
		PayInfo          *input_pay.PayInfoModel
		MemberInfo       *model.MemberIdentity
		OrderAmount      float64
		cabinetRequest   *cabinetApi.CabinetInfoParams
		cabinetResponse  *cabinetApi.CabinetInfoResponse
		CabinetApiConfig *model.CabinetApiConfig
	)

	out = new(input_cabinet.PreCreateOrderModel)
	out.PreOrderSn = uuid.CreatePayCode("PRECAB")
	// 生成预订单详情
	// 会员ID
	MemberInfo = contexts.GetMemberUser(ctx)
	preOrderDetail.MemberId = MemberInfo.Id

	preOrderDetail.IsFx = "N"
	if MemberInfo.IsFx {
		preOrderDetail.IsFx = "Y"
	}

	// 预订单号
	preOrderDetail.PreOrderSn = out.PreOrderSn

	// 请求储物柜详情接口
	cabinetRequest = new(cabinetApi.CabinetInfoParams)
	cabinetRequest.CabinetId = in.CabinetId
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).Cabinet(ctx, cabinetRequest); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(gi18n.T(ctx, "failed_to_obtain_locker_information"))
		return
	}

	preOrderDetail.CabinetInfo = &cabinetResponse.Data

	var language = contexts.GetLanguage(ctx)
	switch language {
	case "zh_CN":
		preOrderDetail.CabinetInfo.Name = preOrderDetail.CabinetInfo.NameTw
		preOrderDetail.CabinetInfo.CityName = preOrderDetail.CabinetInfo.CityNameTw
		preOrderDetail.CabinetInfo.MchName = preOrderDetail.CabinetInfo.MchNameTw
		preOrderDetail.CabinetInfo.MchBranchName = preOrderDetail.CabinetInfo.MchBranchNameTw
		preOrderDetail.CabinetInfo.BoxTypeA.Name = preOrderDetail.CabinetInfo.BoxTypeA.NameTw
		preOrderDetail.CabinetInfo.BoxTypeB.Name = preOrderDetail.CabinetInfo.BoxTypeB.NameTw
		preOrderDetail.CabinetInfo.BoxTypeC.Name = preOrderDetail.CabinetInfo.BoxTypeC.NameTw
	case "en":
		preOrderDetail.CabinetInfo.Name = preOrderDetail.CabinetInfo.NameEn
		preOrderDetail.CabinetInfo.CityName = preOrderDetail.CabinetInfo.CityNameEn
		preOrderDetail.CabinetInfo.MchName = preOrderDetail.CabinetInfo.MchNameEn
		preOrderDetail.CabinetInfo.MchBranchName = preOrderDetail.CabinetInfo.MchBranchNameEn
		preOrderDetail.CabinetInfo.BoxTypeA.Name = preOrderDetail.CabinetInfo.BoxTypeA.NameEn
		preOrderDetail.CabinetInfo.BoxTypeB.Name = preOrderDetail.CabinetInfo.BoxTypeB.NameEn
		preOrderDetail.CabinetInfo.BoxTypeC.Name = preOrderDetail.CabinetInfo.BoxTypeC.NameEn
	case "ja":
		preOrderDetail.CabinetInfo.Name = preOrderDetail.CabinetInfo.NameJa
		preOrderDetail.CabinetInfo.CityName = preOrderDetail.CabinetInfo.CityNameJa
		preOrderDetail.CabinetInfo.MchName = preOrderDetail.CabinetInfo.MchNameJa
		preOrderDetail.CabinetInfo.MchBranchName = preOrderDetail.CabinetInfo.MchBranchNameJa
		preOrderDetail.CabinetInfo.BoxTypeA.Name = preOrderDetail.CabinetInfo.BoxTypeA.NameJa
		preOrderDetail.CabinetInfo.BoxTypeB.Name = preOrderDetail.CabinetInfo.BoxTypeB.NameJa
		preOrderDetail.CabinetInfo.BoxTypeC.Name = preOrderDetail.CabinetInfo.BoxTypeC.NameJa
	case "ko":
		preOrderDetail.CabinetInfo.Name = preOrderDetail.CabinetInfo.NameKo
		preOrderDetail.CabinetInfo.CityName = preOrderDetail.CabinetInfo.CityNameKo
		preOrderDetail.CabinetInfo.MchName = preOrderDetail.CabinetInfo.MchNameKo
		preOrderDetail.CabinetInfo.MchBranchName = preOrderDetail.CabinetInfo.MchBranchNameKo
		preOrderDetail.CabinetInfo.BoxTypeA.Name = preOrderDetail.CabinetInfo.BoxTypeA.NameKo
		preOrderDetail.CabinetInfo.BoxTypeB.Name = preOrderDetail.CabinetInfo.BoxTypeB.NameKo
		preOrderDetail.CabinetInfo.BoxTypeC.Name = preOrderDetail.CabinetInfo.BoxTypeC.NameKo
	case "zh":
		preOrderDetail.CabinetInfo.Name = preOrderDetail.CabinetInfo.NameZh
		preOrderDetail.CabinetInfo.CityName = preOrderDetail.CabinetInfo.CityNameZh
		preOrderDetail.CabinetInfo.MchName = preOrderDetail.CabinetInfo.MchNameZh
		preOrderDetail.CabinetInfo.MchBranchName = preOrderDetail.CabinetInfo.MchBranchNameZh
		preOrderDetail.CabinetInfo.BoxTypeA.Name = preOrderDetail.CabinetInfo.BoxTypeA.NameZh
		preOrderDetail.CabinetInfo.BoxTypeB.Name = preOrderDetail.CabinetInfo.BoxTypeB.NameZh
		preOrderDetail.CabinetInfo.BoxTypeC.Name = preOrderDetail.CabinetInfo.BoxTypeC.NameZh
	}

	// 支付信息 默认不使用积分进行支付
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		Scene:       "CABINET",
		OrderAmount: OrderAmount,
		MemberId:    MemberInfo.Id,
		IsBalance:   false,
	}); err != nil {
		return
	}
	preOrderDetail.PayInfo = new(input_cabinet.OrderPayInfoModel)
	preOrderDetail.PayInfo.Balance.BalanceConfig = new(input_cabinet.BalanceConfig)
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

	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, preOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}

// PreOrderDetail 预下单详情
func (s *sCabinetService) PreOrderDetail(ctx context.Context, in *input_cabinet.PreOrderDetailInp) (out *input_cabinet.PreOrderDetailModel, err error) {
	var (
		PreOrderDetail *gvar.Var
	)
	if PreOrderDetail, err = cache.Instance().Get(ctx, "PreOrder_"+in.PreOrderSn); err != nil || PreOrderDetail.IsEmpty() {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	out = new(input_cabinet.PreOrderDetailModel)
	if err = PreOrderDetail.Struct(&out); err != nil {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	return
}

// PrePayInfo 预下单重新计算支付金额支付模式
func (s *sCabinetService) PrePayInfo(ctx context.Context, in *input_cabinet.PrePayInfoInp) (out *input_cabinet.OrderPayInfoModel, err error) {
	var (
		PreOrderDetail      *input_cabinet.PreOrderDetailModel
		PayInfo             *input_pay.PayInfoModel
		FirstOrder          *entity.CabinetOrder
		IsBalance           bool
		ChooseBoxTypeName   string
		ChooseBoxTypeNameZh string
		ChooseBoxTypeNameEn string
		ChooseBoxTypeNameJa string
		ChooseBoxTypeNameKo string
		ChooseBoxTypeNameTw string
		ChooseBoxTypePrice  int
		AllAmount           float64
	)
	PreOrderDetail = new(input_cabinet.PreOrderDetailModel)
	if PreOrderDetail, err = s.PreOrderDetail(ctx, &input_cabinet.PreOrderDetailInp{PreOrderSn: in.PreOrderSn}); err != nil {
		return
	}
	switch {
	case PreOrderDetail.CabinetInfo.BoxTypeA.ID == in.BoxTypeId:
		if PreOrderDetail.CabinetInfo.BoxTypeA.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}
		ChooseBoxTypeName = PreOrderDetail.CabinetInfo.BoxTypeA.Name
		ChooseBoxTypeNameZh = PreOrderDetail.CabinetInfo.BoxTypeA.NameZh
		ChooseBoxTypeNameEn = PreOrderDetail.CabinetInfo.BoxTypeA.NameEn
		ChooseBoxTypeNameJa = PreOrderDetail.CabinetInfo.BoxTypeA.NameJa
		ChooseBoxTypeNameKo = PreOrderDetail.CabinetInfo.BoxTypeA.NameKo
		ChooseBoxTypeNameTw = PreOrderDetail.CabinetInfo.BoxTypeA.NameTw
		ChooseBoxTypePrice = PreOrderDetail.CabinetInfo.BoxTypeA.Price

	case PreOrderDetail.CabinetInfo.BoxTypeB.ID == in.BoxTypeId:
		if PreOrderDetail.CabinetInfo.BoxTypeB.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}
		ChooseBoxTypeName = PreOrderDetail.CabinetInfo.BoxTypeB.Name
		ChooseBoxTypeNameZh = PreOrderDetail.CabinetInfo.BoxTypeB.NameZh
		ChooseBoxTypeNameEn = PreOrderDetail.CabinetInfo.BoxTypeB.NameEn
		ChooseBoxTypeNameJa = PreOrderDetail.CabinetInfo.BoxTypeB.NameJa
		ChooseBoxTypeNameKo = PreOrderDetail.CabinetInfo.BoxTypeB.NameKo
		ChooseBoxTypeNameTw = PreOrderDetail.CabinetInfo.BoxTypeB.NameTw
		ChooseBoxTypePrice = PreOrderDetail.CabinetInfo.BoxTypeB.Price

	case PreOrderDetail.CabinetInfo.BoxTypeC.ID == in.BoxTypeId:
		if PreOrderDetail.CabinetInfo.BoxTypeC.Num == 0 {
			err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
			return
		}
		ChooseBoxTypeName = PreOrderDetail.CabinetInfo.BoxTypeC.Name
		ChooseBoxTypeNameZh = PreOrderDetail.CabinetInfo.BoxTypeC.NameZh
		ChooseBoxTypeNameEn = PreOrderDetail.CabinetInfo.BoxTypeC.NameEn
		ChooseBoxTypeNameJa = PreOrderDetail.CabinetInfo.BoxTypeC.NameJa
		ChooseBoxTypeNameKo = PreOrderDetail.CabinetInfo.BoxTypeC.NameKo
		ChooseBoxTypeNameTw = PreOrderDetail.CabinetInfo.BoxTypeC.NameTw
		ChooseBoxTypePrice = PreOrderDetail.CabinetInfo.BoxTypeC.Price

	default:
		err = gerror.New(gi18n.T(ctx, "please_select_the_format"))
		return
	}

	if g.IsEmpty(in.Hours) {
		err = gerror.New(gi18n.T(ctx, "please_select_the_rental_duration"))
		return
	}

	if in.Hours < PreOrderDetail.CabinetInfo.MinHours {
		err = gerror.Newf(gi18n.T(ctx, "minimum_rental_time"), PreOrderDetail.CabinetInfo.MinHours)
		return
	}

	PreOrderDetail.BoxTypeInfo = new(input_cabinet.BoxTypeInfoItem)
	PreOrderDetail.BoxTypeInfo.Id = in.BoxTypeId
	PreOrderDetail.BoxTypeInfo.Name = ChooseBoxTypeName
	PreOrderDetail.BoxTypeInfo.NameZh = ChooseBoxTypeNameZh
	PreOrderDetail.BoxTypeInfo.NameEn = ChooseBoxTypeNameEn
	PreOrderDetail.BoxTypeInfo.NameJa = ChooseBoxTypeNameJa
	PreOrderDetail.BoxTypeInfo.NameKo = ChooseBoxTypeNameKo
	PreOrderDetail.BoxTypeInfo.NameTw = ChooseBoxTypeNameTw
	PreOrderDetail.BoxTypeInfo.Price = ChooseBoxTypePrice

	PreOrderDetail.ChooseHours = in.Hours

	// 查询是否是第一单
	_ = dao.CabinetOrder.Ctx(ctx).Where(dao.CabinetOrder.Columns().MemberId, PreOrderDetail.MemberId).Where(dao.CabinetOrder.Columns().PayStatus, "HAVE_PAID").Scan(&FirstOrder)

	if g.IsEmpty(FirstOrder) {
		AllAmount = float64(ChooseBoxTypePrice) * float64(in.Hours) * PreOrderDetail.CabinetInfo.OrderFirstFeeRate
	} else {
		AllAmount = float64(ChooseBoxTypePrice) * float64(in.Hours)
	}

	PayInfo = new(input_pay.PayInfoModel)
	if in.IsBalance == 1 {
		IsBalance = true
	}
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		Scene:       "CABINET",
		OrderAmount: AllAmount,
		MemberId:    PreOrderDetail.MemberId,
		IsBalance:   IsBalance,
		CouponId:    in.CouponId,
	}); err != nil {
		return
	}
	out = new(input_cabinet.OrderPayInfoModel)
	out.Balance.BalanceConfig = new(input_cabinet.BalanceConfig)
	out.PreOrderSn = PreOrderDetail.PreOrderSn
	out.AllAmount = AllAmount
	out.Score = PayInfo.Score
	out.MemberBalance = PayInfo.MemberBalance
	out.PayModel = PayInfo.PayModel
	out.Balance.BalanceAmount = PayInfo.BalanceAmount
	out.ThirdPay.ThirdAmount = PayInfo.ThirdAmount
	out.Balance.BalanceConfig.ScenePayRate = PayInfo.BalanceConfig.ScenePayRate
	out.Balance.BalanceConfig.ExchangeRate = PayInfo.BalanceConfig.ExchangeRate
	out.Balance.BalanceConfig.LevelName = PayInfo.BalanceConfig.LevelName
	out.Balance.BalanceConfig.Level = PayInfo.BalanceConfig.Level
	out.Coupon.CouponId = PayInfo.CouponId
	out.Coupon.CouponAmount = PayInfo.CouponAmount
	out.Coupon.CouponName = PayInfo.CouponName
	PreOrderDetail.PayInfo = out
	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, PreOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}
