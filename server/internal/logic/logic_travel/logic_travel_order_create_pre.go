package logic_travel

import (
	"APT/internal/dao"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_pay"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTravelOrderCreatePreService struct{}

func NewTravelOrderCreatePreService() *sTravelOrderCreatePreService {
	return &sTravelOrderCreatePreService{}
}

func init() {
	service.RegisterTravelOrderCreatePreService(NewTravelOrderCreatePreService())
}

// PreOrder 预下单
func (s *sTravelOrderCreatePreService) PreOrder(ctx context.Context, in *input_travel.PreCreateOrderInp) (out *input_travel.PreCreateOrderModel, err error) {
	var (
		preOrderDetail = new(input_travel.PreOrderDetailModel)
		MemberInfo     *model.MemberIdentity
		PayInfo        *input_pay.PayInfoModel
		ProductInfo    *entity.TravelProduct
		ProductSkuInfo *entity.TravelProductSku
		OrderAmount    float64
	)

	// 验证产品信息
	if err = dao.TravelProduct.Ctx(ctx).Where(dao.TravelProduct.Columns().Id, in.ProductId).Hook(hook2.PmsFindLanguageValueHook).Scan(&ProductInfo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 产品不存在
			err = gerror.New(gi18n.T(ctx, "product_not_exist"))
		} else {
			// 获取产品信息失败
			err = gerror.Wrap(err, gi18n.T(ctx, "get_product_info_failed"))
		}
		return
	}
	if ProductInfo == nil {
		err = gerror.New(gi18n.T(ctx, "product_not_exist"))
		return
	}

	// 验证产品状态
	if ProductInfo.Status != 1 {
		// 产品不可用
		err = gerror.New(gi18n.T(ctx, "product_not_available"))
		return
	}

	// 验证产品SKU信息
	if err = dao.TravelProductSku.Ctx(ctx).Where(dao.TravelProductSku.Columns().Id, in.SkuId).Hook(hook2.PmsFindLanguageValueHook).Scan(&ProductSkuInfo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 车型不存在
			err = gerror.New(gi18n.T(ctx, "car_type_not_exist"))
		} else {
			// 获取车型信息失败
			err = gerror.Wrap(err, gi18n.T(ctx, "get_car_type_info_failed"))
		}
		return
	}
	if ProductSkuInfo == nil {
		// 车型不存在
		err = gerror.New(gi18n.T(ctx, "car_type_not_exist"))
		return
	}
	if ProductSkuInfo.ProductId != uint(in.ProductId) {
		// 车型不属于该产品
		err = gerror.New(gi18n.T(ctx, "car_type_not_belong_to_product"))
		return
	}

	// 验证预约日期
	bookDate, err := gtime.StrToTime(in.BookDate)
	if err != nil {
		// 日期格式错误
		err = gerror.New(gi18n.T(ctx, "invalid_book_date"))
		return
	}

	// 统一截断到日粒度，避免时分秒干扰比较结果
	now := gtime.Now()
	today := gtime.NewFromStr(now.Format("Y-m-d"))
	bookDay := gtime.NewFromStr(bookDate.Format("Y-m-d"))

	// 检查预约日期是否在允许范围内（今天可以预约）
	if bookDay.Before(today) {
		// 预约日期不能早于今天
		err = gerror.New(gi18n.T(ctx, "book_date_cannot_be_past"))
		return
	}

	// 检查最大可预约天数
	if ProductInfo.MaxBookDays > 0 {
		maxBookDate := today.AddDate(0, 0, int(ProductInfo.MaxBookDays))
		if bookDay.After(maxBookDate) {
			err = gerror.New(gi18n.T(ctx, "book_date_exceeds_limit"))
			return
		}
	}

	// 检查至少提前预订天数
	if ProductInfo.AdvanceBookDays > 0 {
		minBookDate := today.AddDate(0, 0, int(ProductInfo.AdvanceBookDays))
		if bookDay.Before(minBookDate) {
			err = gerror.New(gi18n.T(ctx, "this_date_is_not_available_for_booking"))
			return
		}
	}

	// 检查库存 - 总库存限制
	var totalBookingNum float64
	totalBookingNum, err = dao.TravelOrder.Ctx(ctx).
		Where(dao.TravelOrder.Columns().ProductId, in.ProductId).
		WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "WAIT_VERIFY", "DONE"}).
		Sum(dao.TravelOrder.Columns().BookingNum)
	if err != nil {
		//
		err = gerror.Wrap(err, gi18n.T(ctx, "check_stock_failed"))
		return
	}

	remainingStock := int(ProductInfo.Stock) - int(totalBookingNum)
	if remainingStock < in.BookNum {
		// 库存不足
		err = gerror.New(gi18n.T(ctx, "insufficient_stock"))
		return
	}

	// 检查每日最大接待人数限制
	if ProductSkuInfo.DailyCapacity > 0 {
		var dailyBookingNum float64
		dailyBookingNum, err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().ProductId, in.ProductId).
			Where(dao.TravelOrder.Columns().SkuId, in.SkuId).
			Where(dao.TravelOrder.Columns().BookDate, in.BookDate).
			WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "WAIT_VERIFY", "DONE"}).
			Sum(dao.TravelOrder.Columns().BookingNum)
		if err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "check_daily_capacity_failed"))
			return
		}

		remainingDailyCapacity := int(ProductSkuInfo.DailyCapacity) - int(dailyBookingNum)
		if remainingDailyCapacity < in.BookNum {
			err = gerror.New(gi18n.T(ctx, "daily_capacity_insufficient"))
			return
		}
	}

	out = new(input_travel.PreCreateOrderModel)
	out.PreOrderSn = uuid.CreatePayCode("PRETRAVEL")

	// 查询取消政策配置, 直接根据sku 获取
	allowCancel := ProductSkuInfo.AllowCancel
	if allowCancel == 2 {
		// 不可取消
		preOrderDetail.CancelFreeDate = ""
		preOrderDetail.CancelPolicy = gi18n.T(ctx, "cancel_policy_no_cancel")
	} else if allowCancel == 1 {
		// 可免费取消，计算免费取消截止时间
		// freeCancelDays 实际存的是小时数
		freeCancelHours := gconv.Int(ProductSkuInfo.FreeCancelHours)
		cancelFeePercent := ProductSkuInfo.CancelFeePercent

		// 以预定日期 + SKU 集合时间作为服务开始时间，再减去免费取消小时数
		meetingTimeStr := ProductSkuInfo.MeetingTime
		// 默认服务开始时间为预定日期 00:00
		serviceStartStr := bookDate.Format("Y-m-d") + " 00:00:00"
		if meetingTimeStr != "" {
			serviceStartStr = bookDate.Format("Y-m-d") + " " + meetingTimeStr + ":00"
		}
		serviceStart, parseErr := gtime.StrToTime(serviceStartStr)
		var cancelFreeDateStr string
		if parseErr == nil {
			cancelFreeTime := serviceStart.Add(-time.Duration(freeCancelHours) * time.Hour)
			cancelFreeDateStr = cancelFreeTime.Format("Y-m-d H:i")
		} else {
			cancelFreeDateStr = bookDate.Format("Y-m-d")
		}
		preOrderDetail.CancelFreeDate = cancelFreeDateStr
		preOrderDetail.CancelPolicy = gi18n.Tf(ctx, "cancel_policy_free_before", gvar.New(freeCancelHours).String(), cancelFeePercent)
	}

	// 获取预定须知
	// 获取当前语言
	currentLang := contexts.GetLanguage(ctx)
	// 映射语言到配置key（参考logic_travel_product.go）
	configLanguageKey := "zh" // 默认中文
	switch currentLang {
	case "zh", "zh_CN":
		configLanguageKey = "zh"
	case "en":
		configLanguageKey = "en"
	case "ja":
		configLanguageKey = "ja"
	case "ko":
		configLanguageKey = "ko"
	case "tw", "zh_TW":
		configLanguageKey = "zh_CN"
	default:
		// 默认使用中文
		configLanguageKey = "zh"
	}

	// 从sys_config表获取预定须知（参考app_travel_order_info.go）
	config, err := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "travelothersetting",
	})

	// 设置预定须知内容
	preOrderDetail.BookingNotice = "" // 默认空值
	if err == nil && config != nil && config.List != nil {
		// 获取预定须知
		bookingNoticeKey := "bookingNotice_" + configLanguageKey
		if bookingNoticeValue := config.List[bookingNoticeKey]; bookingNoticeValue != nil {
			preOrderDetail.BookingNotice = gvar.New(bookingNoticeValue).String()
		} else {
			// 如果当前语言配置不存在，尝试使用中文版本
			if bookingNoticeValue := config.List["bookingNotice_zh"]; bookingNoticeValue != nil {
				preOrderDetail.BookingNotice = gvar.New(bookingNoticeValue).String()
			}
		}
	}

	// 如果仍然没有获取到内容，使用默认文本
	if preOrderDetail.BookingNotice == "" {
		preOrderDetail.BookingNotice = gi18n.T(ctx, "default_booking_notice")
	}

	// 计算订单金额
	OrderAmount = ProductSkuInfo.Price * float64(in.BookNum)

	// 生成预订单详情
	preOrderDetail.PreOrderSn = out.PreOrderSn
	// 获取会员信息
	MemberInfo = contexts.GetMemberUser(ctx)
	preOrderDetail.MemberId = uint64(MemberInfo.Id)
	preOrderDetail.IsFx = "N"
	if MemberInfo.IsFx {
		preOrderDetail.IsFx = "Y"
	}
	// 预定信息
	preOrderDetail.BookingInfo.BookDate = in.BookDate
	preOrderDetail.BookingInfo.BookNum = in.BookNum
	// 产品信息
	preOrderDetail.ProductInfo.Id = ProductInfo.Id
	preOrderDetail.ProductInfo.Title = ProductInfo.Title
	preOrderDetail.ProductInfo.SubTitle = ProductInfo.SubTitle
	preOrderDetail.ProductInfo.Price = ProductSkuInfo.Price
	preOrderDetail.ProductInfo.MeetingPlace = ProductSkuInfo.MeetingPlace
	preOrderDetail.ProductInfo.MeetingTime = ProductSkuInfo.MeetingTime
	preOrderDetail.ProductInfo.GgLat = ProductSkuInfo.GgLat
	preOrderDetail.ProductInfo.GgLng = ProductSkuInfo.GgLng
	// 车型信息
	preOrderDetail.SkuInfo.Id = ProductSkuInfo.Id
	preOrderDetail.SkuInfo.Name = ProductSkuInfo.Name
	preOrderDetail.SkuInfo.Price = ProductSkuInfo.Price
	preOrderDetail.SkuInfo.MeetingPlace = ProductSkuInfo.MeetingPlace
	preOrderDetail.SkuInfo.MeetingTime = ProductSkuInfo.MeetingTime
	preOrderDetail.SkuInfo.GgLat = ProductSkuInfo.GgLat
	preOrderDetail.SkuInfo.GgLng = ProductSkuInfo.GgLng
	preOrderDetail.SkuInfo.ContactMobile = ProductSkuInfo.ContactMobile
	// 入住人信息
	preOrderDetail.PreUser.FullName = MemberInfo.FullName
	preOrderDetail.PreUser.Phone = MemberInfo.Phone
	preOrderDetail.PreUser.Mail = MemberInfo.Mail
	preOrderDetail.PreUser.PhoneArea = MemberInfo.PhoneArea

	// 支付信息 默认不使用积分进行支付
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		Scene:       "TRAVEL",
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

	if err = cache.Instance().Set(ctx, "PreOrder_"+out.PreOrderSn, preOrderDetail, gtime.M*30); err != nil {
		return
	}
	return
}

// PreOrderDetail 预下单详情
func (s *sTravelOrderCreatePreService) PreOrderDetail(ctx context.Context, in *input_travel.PreOrderDetailInp) (out *input_travel.PreOrderDetailModel, err error) {
	var (
		PreOrderDetail *gvar.Var
	)
	if PreOrderDetail, err = cache.Instance().Get(ctx, "PreOrder_"+in.PreOrderSn); err != nil || PreOrderDetail.IsEmpty() {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	out = new(input_travel.PreOrderDetailModel)
	if err = PreOrderDetail.Struct(&out); err != nil {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	return
}

// PrePayInfo 预下单重新计算支付金额支付模式
func (s *sTravelOrderCreatePreService) PrePayInfo(ctx context.Context, in *input_pay.PrePayInfoInp) (out *input_pay.PrePayInfoModel, err error) {
	var (
		PreOrderDetail *input_travel.PreOrderDetailModel
		PayInfo        *input_pay.PayInfoModel
		IsBalance      bool
	)
	PreOrderDetail = new(input_travel.PreOrderDetailModel)
	if PreOrderDetail, err = s.PreOrderDetail(ctx, &input_travel.PreOrderDetailInp{PreOrderSn: in.PreOrderSn}); err != nil {
		return
	}
	PayInfo = new(input_pay.PayInfoModel)
	if in.IsBalance == 1 {
		IsBalance = true
	}
	if PayInfo, err = service.PayService().PayInfo(ctx, &input_pay.PayInfoInp{
		Scene:       "TRAVEL",
		OrderAmount: PreOrderDetail.PayInfo.AllAmount,
		MemberId:    int(PreOrderDetail.MemberId),
		IsBalance:   IsBalance,
		CouponId:    in.CouponId,
	}); err != nil {
		return
	}
	out = new(input_pay.PrePayInfoModel)
	out.OrderPayInfoModel = new(input_hotel.OrderPayInfoModel)
	out.OrderPayInfoModel.Balance.BalanceConfig = new(input_hotel.BalanceConfig)
	out.OrderPayInfoModel.PreOrderSn = PreOrderDetail.PreOrderSn
	out.OrderPayInfoModel.AllAmount = PreOrderDetail.PayInfo.AllAmount
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
