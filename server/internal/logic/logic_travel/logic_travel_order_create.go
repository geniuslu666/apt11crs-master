package logic_travel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"APT/utility/validate"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

type sTravelOrderCreateService struct{}

func NewTravelOrderCreateService() *sTravelOrderCreateService {
	return &sTravelOrderCreateService{}
}

func init() {
	service.RegisterTravelOrderCreateService(NewTravelOrderCreateService())
}

// CreateOrder 创建订单
func (s *sTravelOrderCreateService) CreateOrder(ctx context.Context, in *input_travel.CreateOrderInp) (out *input_travel.CreateOrderModel, err error) {
	var (
		InsertId          int64
		OrderSn           = uuid.CreateOrderCode("T")
		PreOrderDetailVar *gvar.Var
		PreOrderDetail    *input_travel.PreOrderDetailModel
		TravelProduct     *entity.TravelProduct
		TravelProductSku  *entity.TravelProductSku
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
		travelOrder       *entity.TravelOrder
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
	PreOrderDetail = new(input_travel.PreOrderDetailModel)
	if err = PreOrderDetailVar.Struct(&PreOrderDetail); err != nil {
		// 订单错误
		err = gerror.New(gi18n.T(ctx, "order_error"))
		return
	}
	g.Log().Info(ctx, "校验今日下单时间限制")
	g.Log().Info(ctx, PreOrderDetail.BookingInfo.BookDate)
	g.Log().Info(ctx, gtime.Now().Format("Y-m-d"))

	if PreOrderDetail.BookingInfo.BookDate < gtime.Now().Format("Y-m-d") {
		// 下单时间限制,已无法下单
		err = gerror.New(gi18n.T(ctx, "order_time_limit"))
		return
	}

	// 校验手机号
	if !g.IsEmpty(in.BookingMobile) && !g.IsEmpty(in.PhoneArea) {
		// 使用手机号验证工具进行验证和标准化
		phoneResult := validate.IsPhoneNumberWithCountryCode(in.PhoneArea, in.BookingMobile)
		if !phoneResult.IsValid {
			// 手机号格式错误
			err = gerror.New(gi18n.T(ctx, "phone_format_error"))
			return
		}
	}

	// 查询当前产品信息
	if err = dao.TravelProduct.Ctx(ctx).Where(dao.TravelProduct.Columns().Id, PreOrderDetail.ProductInfo.Id).Scan(&TravelProduct); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 不存在该产品
		err = gerror.New(gi18n.T(ctx, "the_product_not_exist"))
		return
	}

	// 判断库存（每日库存&总库存）
	if TravelProduct.Stock > 0 {
		// 检查总库存
		totalBooked, totalErr := dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().ProductId, PreOrderDetail.ProductInfo.Id).
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"CANCEL", "REFUND"}).
			Sum(dao.TravelOrder.Columns().BookingNum)
		if totalErr != nil {
			// 查询已预订数量失败
			err = gerror.Wrap(totalErr, gi18n.T(ctx, "failed_to_query_the_number_of_reservations"))
			return
		}

		if uint(totalBooked)+uint(PreOrderDetail.BookingInfo.BookNum) > TravelProduct.Stock {
			err = gerror.New(gi18n.T(ctx, "insufficient_stock"))
			return
		}
	}

	// 查询当前产品Sku信息
	if err = dao.TravelProductSku.Ctx(ctx).Where(dao.TravelProductSku.Columns().Id, PreOrderDetail.SkuInfo.Id).Scan(&TravelProductSku); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 不存在该产品
		err = gerror.New(gi18n.T(ctx, "car_type_not_exist"))
		return
	}

	if TravelProductSku.DailyCapacity > 0 {
		// 检查每日库存
		dailyBooked, dailyErr := dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().ProductId, PreOrderDetail.ProductInfo.Id).
			Where(dao.TravelOrder.Columns().SkuId, PreOrderDetail.SkuInfo.Id).
			Where(dao.TravelOrder.Columns().BookDate, PreOrderDetail.BookingInfo.BookDate).
			WhereNotIn(dao.TravelOrder.Columns().OrderStatus, []string{"CANCEL", "REFUND"}).
			Sum(dao.TravelOrder.Columns().BookingNum)
		if dailyErr != nil {
			// 查询当日已预订数量失败
			err = gerror.Wrap(dailyErr, gi18n.T(ctx, "failed_to_query_the_number_of_reservations"))
			return
		}

		if uint(dailyBooked)+uint(PreOrderDetail.BookingInfo.BookNum) > TravelProductSku.DailyCapacity {
			// 日库存不足
			err = gerror.New(gi18n.T(ctx, "daily_capacity_insufficient"))
			return
		}
	}

	// 初始化响应变量
	out = new(input_travel.CreateOrderModel)
	out.ThirdPay.IsFxOrder = PreOrderDetail.IsFx == "Y"
	out.OrderSn = OrderSn
	// 开启闭包事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 读取支付配置
		if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
			return
		}
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

		// 创建一日游订单
		CreateOrderTime := gtime.Now()

		// 获取推荐人信息
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return err
		}
		referrerId := PmsMemberInfo.Referrer
		if YYConfig.RecommendModel == "FIRST" {
			referrerId = PmsMemberInfo.LastReferrer
		}
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, referrerId).Scan(&ReferrerInfo); err != nil {
			return err
		}

		// 设置积分获取信息
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 6).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		// 创建旅行订单
		travelOrder = &entity.TravelOrder{
			OrderSn:        OrderSn,
			MemberId:       uint64(MemberInfo.Id),
			ProductId:      PreOrderDetail.ProductInfo.Id,
			SkuId:          PreOrderDetail.SkuInfo.Id,
			BookDate:       gtime.New(PreOrderDetail.BookingInfo.BookDate),
			BookingNum:     uint(PreOrderDetail.BookingInfo.BookNum),
			OrderAmount:    PreOrderDetail.PayInfo.AllAmount,
			OrderStatus:    "WAIT_PAY",
			BookingName:    in.LastName + " " + in.FirstName,
			FirstName:      in.FirstName,
			LastName:       in.LastName,
			BookingMobile:  in.BookingMobile,
			PhoneArea:      in.PhoneArea,
			BookingEmail:   in.BookingEmail,
			Referrer:       referrerId,
			CreatedAt:      CreateOrderTime,
			CouponAmount:   PreOrderDetail.PayInfo.Coupon.CouponAmount,
			BalAmount:      PreOrderDetail.PayInfo.Balance.BalanceAmount,
			PayModel:       PreOrderDetail.PayInfo.PayModel,
			ExpirationTime: ExpirationTime,
			IsPayOpen:      MemberScene.IsPayOpen,
		}

		// 设置返佣信息
		if !g.IsEmpty(ReferrerInfo) {
			travelOrder.RebateStatus = "WAIT"
			if ReferrerInfo.RebateMode == "CHANNEL" {
				if err = dao.PmsChannel.Ctx(ctx).WherePri(ReferrerInfo.ChannelId).Scan(&ChannelInfo); err != nil {
					return err
				}
				if !g.IsEmpty(ChannelInfo) {
					travelOrder.RebateRate = ChannelInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "STAFF" {
				if err = dao.PmsStaff.Ctx(ctx).WherePri(ReferrerInfo.StaffId).Scan(&StaffInfo); err != nil {
					return err
				}
				if !g.IsEmpty(StaffInfo) {
					travelOrder.RebateRate = StaffInfo.Rate
				}
			} else if ReferrerInfo.RebateMode == "MEMBER" {
				travelOrder.RebateRate = YYConfig.MemberBrokerageRate
			} else {
				return gerror.New(gi18n.T(ctx, "the_rebate_mode_does_not_exist"))
			}
		}

		if !g.IsEmpty(MemberScene) {
			travelOrder.IsGetOpen = MemberScene.IsGetOpen
			if MemberScene.IsGetOpen == "Y" {
				travelOrder.TravelGetRateScene = MemberScene.GetRate
				if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				if !g.IsEmpty(MemberLevel) {
					travelOrder.TravelGetRateVip = MemberLevel.TravelGetRate
				}
			}
		}

		// 插入旅行订单
		if InsertId, err = dao.TravelOrder.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(travelOrder); err != nil {
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
				OrderType:     "TRAVEL",
				Scene:         "TRAVEL",
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
				OrderType:        "TRAVEL",
				Scene:            "TRAVEL",
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
				OrderType:     "TRAVEL",
				Scene:         "TRAVEL",
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

	// 旅行订单下单日志
	if _, err = dao.TravelOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.TravelOrderLog{
		OrderId:     int(InsertId),
		ActionWay:   "CREATE",
		Remark:      "订单创建",
		OperateType: "USER",
		OrderStatus: "WAIT_PAY",
		OperateId:   MemberInfo.Id,
	}); err != nil {
		return
	}

	// 发送到消息队列
	//systemMessageTitle := map[string]string{
	//	"zh":    "订单创建成功",
	//	"en":    "Order created successfully",
	//	"ja":    "注文が正常に作成されました",
	//	"ko":    "주문이 성공적으로 생성되었습니다.",
	//	"zh_CN": "訂單創建成功",
	//}
	//systemMessageContent := map[string]string{
	//	"zh":    "您已成功下单，订单号为：" + OrderSn,
	//	"en":    "You have successfully placed an order, the order number is: " + OrderSn,
	//	"ja":    "注文が完了しました。注文番号は：" + OrderSn,
	//	"ko":    "주문이 완료되었습니다. 주문번호는：" + OrderSn,
	//	"zh_CN": "您已成功下單，訂單編號為：" + OrderSn,
	//}
	//// 查询用户的手机号区号 来判断用户语言
	//phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(MemberInfo.Id).Value()
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
	//	"orderSn": OrderSn,
	//}
	//pushDataJson, _ := json.Marshal(pushData)
	//appPushData := g.MapStrStr{
	//	"type":  "2",
	//	"param": string(pushDataJson),
	//}
	//service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
	//	SystemMessageTitle:   systemMessageTitle,
	//	SystemMessageContent: systemMessageContent,
	//	Scene:                "travel",
	//	Type:                 "order",
	//	MemberId:             int(MemberInfo.Id),
	//	Language:             memberLanguage,
	//	AppPushData:          appPushData,
	//	AppLink:              "/order/order-detail",
	//	WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", OrderSn),
	//	EnablePush:           true,
	//	EnableSms:            false,
	//	PushTitle:            systemMessageTitle[memberLanguage],
	//	PushContent:          systemMessageContent[memberLanguage],
	//	OperatorId:           MemberInfo.Id,
	//	OperatorRole:         "MEMBER",
	//	OrderSn:              OrderSn,
	//})
	return
}
