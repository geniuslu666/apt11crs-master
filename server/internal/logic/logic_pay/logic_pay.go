package logic_pay

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/Paypal"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/cabinetApi"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/mlilifeWxPay"
	"APT/internal/library/paycloud"
	"APT/internal/library/stripePay"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_pay"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/go-pay/gopay/paypal"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"github.com/stripe/stripe-go/v79"
)

// SelectThirdPay 选择第三方支付方式
func (s *sPayService) SelectThirdPay(ctx context.Context, in *input_pay.SelectThirdPayInp) (out *input_pay.SelectThirdPayModel, err error) {
	var (
		PmsTransactionInfo           *entity.PmsTransaction
		SysPmsTransaction            *entity.PmsTransaction
		paycloudClient               *paycloud.Client
		CreateAppOrder               paycloud.CreateAppPayOrderResponse
		CreateWebOrder               paycloud.CreateWebPayOrderResponse
		CreatePayPayOrder            paycloud.CreatePayPayOrderResponse
		CreateWxMiniPayOrder         paycloud.CreateWxMiniPayOrderResponse
		mlilifeWxPayWxappPayResponse *mlilifeWxPay.WxappPayResponse
		PaypalClient                 *Paypal.Client
		StripeClient                 *stripePay.Client
		PaypalCardRespBody           *paypal.OrderDetail
		StripeCardRespBody           *stripe.CheckoutSession
		PaypalRespBody               *paypal.CreateOrderRsp
		AppConfig                    *model.AppConfig
		AppVersionInfo               []*model.AppVersionInfo
		WxminiPaychannel             string
		PayParams                    string
		TransNo                      string
		PayChannel                   string
		MemberUser                   = contexts.GetMemberUser(ctx)
		MemberAuth                   *entity.PmsMemberAuth
		PmsMember                    *entity.PmsMember
		h5FxPayCreateResponse        *h5FxPay.CreateOrderResponse
	)

	out = new(input_pay.SelectThirdPayModel)
	mod := dao.PmsTransaction.Ctx(ctx)
	// TODO 支付
	if g.IsEmpty(in) || len(in.OrderSn) < 1 {
		err = gerror.New("支付类型错误")
		return
	}
	if in.OrderSn[:1] == "C" {
		mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
		mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "CAR")
	} else if in.OrderSn[:1] == "S" {
		mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
		mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "SPA")
	} else if in.OrderSn[:1] == "F" {
		mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
		mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "FOOD")
		mod = mod.Where(dao.PmsTransaction.Columns().PayStatus, "WAIT")
	} else if in.OrderSn[:1] == "B" {
		mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
		mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "CABINET")
		mod = mod.Where(dao.PmsTransaction.Columns().PayStatus, "WAIT")
	} else if in.OrderSn[:1] == "T" {
		mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
		mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "TRAVEL")
		mod = mod.Where(dao.PmsTransaction.Columns().PayStatus, "WAIT")
	} else if in.OrderSn[:1] == "H" {
		if len(in.OrderSn) >= 2 && in.OrderSn[:2] == "HC" {
			mod = mod.Where(dao.PmsTransaction.Columns().ChangeOrderSn, in.OrderSn)
			mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "BOOKING_CHANGE")
		} else {
			mod = mod.Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
			mod = mod.Where(dao.PmsTransaction.Columns().OrderType, "BOOKING")
		}
	} else {
		err = gerror.New("支付类型错误")
		return
	}
	if err = mod.WhereNot(dao.PmsTransaction.Columns().PayChannel, "SYSTEM").Scan(&PmsTransactionInfo); err != nil {
		return
	}
	if g.IsEmpty(PmsTransactionInfo) {
		err = gerror.New("订单不存在")
		return
	}

	orderIsFx, err := getOrderIsFxByOrderSn(ctx, in.OrderSn)
	if err != nil {
		return
	}
	if orderIsFx == "Y" {
		if !MemberUser.IsFx {
			err = gerror.New("分销订单禁止非分销账号支付")
			return
		}
		if in.PayChannel != "H5_FX" {
			err = gerror.New("分销订单仅支持分销渠道支付")
			return
		}
	}
	if MemberUser.IsFx && orderIsFx != "Y" {
		err = gerror.New("非分销订单禁止分销账号支付")
		return
	}
	if PmsTransactionInfo.PayStatus != "WAIT" {
		err = gerror.New("订单超时无法支付")
		return
	}
	// 校验是否存在积分支付   积分是否足够
	if err = dao.PmsTransaction.Ctx(ctx).Where(g.Map{
		dao.PmsTransaction.Columns().PayChannel: "SYSTEM",
		dao.PmsTransaction.Columns().PayType:    "BAL",
		dao.PmsTransaction.Columns().OrderSn:    in.OrderSn,
		dao.PmsTransaction.Columns().PayStatus:  "WAIT",
	}).Scan(&SysPmsTransaction); err != nil {
		return
	}
	if !g.IsEmpty(SysPmsTransaction) {
		// 是存在积分支付的情况
		//查询当前用户积分余额
		if err = dao.PmsMember.Ctx(ctx).Where(g.Map{
			dao.PmsMember.Columns().Id: contexts.GetMemberUser(ctx).Id,
		}).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			err = gerror.New("会员不存在")
			return
		}
		g.Log().Debug(ctx, PmsMember.Balance)
		g.Log().Debug(ctx, SysPmsTransaction.Amount)
		g.Log().Debug(ctx, PmsMember.Balance < SysPmsTransaction.Amount)
		if PmsMember.Balance < SysPmsTransaction.Amount {
			err = gerror.New("积分不足")
			return
		}
	}

	PmsTransactionInfo.TransactionSn = uuid.CreatePayCode("PT")
	DefaultAmount := PmsTransactionInfo.Amount
	if paycloudClient, err = paycloud.NewClient(ctx, in.PaySuccessCallbackUrl); err != nil {
		return
	}
	if MemberUser.IsFx && orderIsFx == "Y" {
		in.PayChannel = "H5_FX"
	}

	switch in.PayChannel {
	case "Paypal":
		err = gerror.New("paypal payment is not supported.")
		return
		PayChannel = "PAYPAL"
		if PaypalClient, err = Paypal.GetClient(ctx, in.PaySuccessCallbackUrl, in.PayFailCallbackUrl); err != nil {
			return
		}
		if PaypalRespBody, err = PaypalClient.CreatOrder(ctx, contexts.GetLanguage(ctx), DefaultAmount); err != nil {
			return
		}
		PayParams = PaypalRespBody.Response.Links[1].Href
		TransNo = PaypalRespBody.Response.Id
		out.WebPayUrl = PayParams
		out.PaypalPayUrl = PayParams
	case "PaypalCard":
		PayChannel = "PAYPAL"
		return
		if PaypalClient, err = Paypal.GetClient(ctx, in.PaySuccessCallbackUrl, in.PayFailCallbackUrl); err != nil {
			return
		}
		if PaypalCardRespBody, err = PaypalClient.CreatCardOrder(ctx, &input_pay.PaypalCardPayParams{
			CardName:    in.CardName,
			CardNumber:  in.CardNumber,
			CardExp:     in.CardExp,
			CardCvv:     in.CardCvv,
			Address:     in.Address,
			PostalCode:  in.PostalCode,
			CountryCode: in.CountryCode,
			Amount:      DefaultAmount,
		}); err != nil {
			return
		}
		if _, err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().Id: PmsTransactionInfo.Id,
		}).Update(g.MapStrAny{
			dao.PmsTransaction.Columns().CaptureId:        PaypalCardRespBody.PurchaseUnits[0].Payments.Captures[0].Id,
			dao.PmsTransaction.Columns().PaymentRequestId: PaypalCardRespBody.PurchaseUnits[0].Payments.Captures[0].Id,
			dao.PmsTransaction.Columns().PayAmount:        PaypalCardRespBody.PurchaseUnits[0].Payments.Captures[0].SellerReceivableBreakdown.NetAmount.Value,
			dao.PmsTransaction.Columns().PayCharge:        PaypalCardRespBody.PurchaseUnits[0].Payments.Captures[0].SellerReceivableBreakdown.PaypalFee.Value,
		}); err != nil {
			return
		}
	case "StripeCard":
		PayChannel = "STRIPE"
		if StripeClient, err = stripePay.GetClient(ctx, in.PaySuccessCallbackUrl, in.PayFailCallbackUrl); err != nil {
			return
		}
		if StripeCardRespBody, err = StripeClient.Pay(ctx, gvar.New(DefaultAmount).Int64(), "JPY"); err != nil {
			return
		}
		if _, err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().Id: PmsTransactionInfo.Id,
		}).Update(g.MapStrAny{
			dao.PmsTransaction.Columns().PaymentRequestId: StripeCardRespBody.ClientReferenceID,
			dao.PmsTransaction.Columns().PayAmount:        StripeCardRespBody.AmountSubtotal,
			dao.PmsTransaction.Columns().PayCharge:        StripeCardRespBody.TotalDetails.AmountTax,
		}); err != nil {
			return
		}
		PayParams = StripeCardRespBody.URL
		TransNo = StripeCardRespBody.ClientReferenceID
		out.StripePayUrl = PayParams
	case "H5_FX":
		PayChannel = "H5_FX"
		if len(in.OrderSn) >= 2 && in.OrderSn[:2] == "HC" {
			in.OrderSn = PmsTransactionInfo.OrderSn
		}
		if h5FxPayCreateResponse, err = h5FxPay.CreateOrder(ctx, &h5FxPay.CreateOrderRequest{
			UserId:          MemberUser.PmsMemberAuth.AuthId,
			OrderNo:         in.OrderSn,
			TradeNo:         PmsTransactionInfo.TransactionSn,
			TotalAmount:     int64(DefaultAmount),
			OrderExpireTime: PmsTransactionInfo.ExpiredTime,
			OrderDetailUrl:  fmt.Sprintf("https://crsdev-api.yeebok.net/#/pages/payment/checkResult?isFxOrder=true&orderSn=%s", in.OrderSn),
			OrderResultUrl:  fmt.Sprintf("https://crsdev-api.yeebok.net/#/pages/payment/checkResult?isFxOrder=true&orderSn=%s", in.OrderSn),
		}); err != nil {
			return
		}
		TransNo = h5FxPayCreateResponse.Data.ThirdOrderNo
		out.WebPayUrl = h5FxPayCreateResponse.Data.CashierUrl
		PayParams = out.WebPayUrl
	default:
		TransNo = PmsTransactionInfo.TransactionSn
		PayChannel = "PAYCLOUD"
		if in.PayChannel == "CreditLinkPay" {
			in.PayChannel = "CreditLinkPay"
			paycloudClient.Endpoint = "https://open.n-age.co.jp/api/entry"
			paycloudClient.AppID = "wz715fc0d10ee9d156"
			paycloudClient.MerchantNo = "312100007235"
			paycloudClient.StoreNo = "4122000030"
			paycloudClient.SubAppid = ""
			paycloudClient.PrivateKey = os.Getenv("PAYCLOUD_PRIVATE_KEY")
			if CreateWebOrder, err = paycloudClient.WebOrder(ctx, &paycloud.CreateWebPayOrderParams{
				MerchantOrderNo: PmsTransactionInfo.TransactionSn,
				OrderAmount:     DefaultAmount,
			}); err != nil {
				return
			}
			out.AppPayParams = CreateWebOrder.Data.PayUrl
			PayParams = out.WebPayUrl
		} else if in.PayChannel == "Alipay+" {
			if CreateAppOrder, err = paycloudClient.AppOrder(ctx, &paycloud.CreateAppPayOrderParams{
				MerchantOrderNo: PmsTransactionInfo.TransactionSn,
				OrderAmount:     DefaultAmount,
				PayMethodID:     in.PayChannel,
			}); err != nil {
				return
			}
			out.AppPayParams = CreateAppOrder.Data.PayParams.PaymentUrl
			PayParams = out.AppPayParams
		} else if in.PayChannel == "WeChatPay" {
			if CreateAppOrder, err = paycloudClient.AppOrder(ctx, &paycloud.CreateAppPayOrderParams{
				MerchantOrderNo: PmsTransactionInfo.TransactionSn,
				OrderAmount:     DefaultAmount,
				PayMethodID:     in.PayChannel,
			}); err != nil {
				return
			}
			out.AppWechatPayParams.Appid = CreateAppOrder.Data.PayParams.Appid
			out.AppWechatPayParams.Package = CreateAppOrder.Data.PayParams.Package
			out.AppWechatPayParams.Partnerid = CreateAppOrder.Data.PayParams.Partnerid
			out.AppWechatPayParams.Prepayid = CreateAppOrder.Data.PayParams.Prepayid
			out.AppWechatPayParams.Noncestr = CreateAppOrder.Data.PayParams.Noncestr
			out.AppWechatPayParams.Timestamp = CreateAppOrder.Data.PayParams.Timestamp
			out.AppWechatPayParams.Sign = CreateAppOrder.Data.PayParams.Sign
			out.AppPayParams = gjson.New(CreateAppOrder.Data.PayParams).String()
			PayParams = out.AppPayParams
		} else if in.PayChannel == "WeChatMiniPay" {
			if AppConfig, err = service.BasicsConfig().GetApp(ctx); err != nil {
				return
			}
			g.Log().Error(ctx, "AppConfig", AppConfig)
			if err = gjson.New(AppConfig.AppVersion).Scan(&AppVersionInfo); err != nil {
				return
			}
			g.Log().Error(ctx, "AppVersionInfo", AppVersionInfo)
			for _, v := range AppVersionInfo {
				if v.Name == "miniapp" {
					g.Log().Error(ctx, "v", v)
					WxminiPaychannel = v.Info.MiniPayChannel
				}
			}
			if g.IsEmpty(WxminiPaychannel) {
				err = gerror.New("微信小程序支付渠道未配置")
				return
			}
			if WxminiPaychannel == "paycloud" {
				if err = dao.PmsMemberAuth.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsMemberAuth.Columns().MemberId: MemberUser.Id,
					dao.PmsMemberAuth.Columns().Channel:  "WX_MINI",
				}).Scan(&MemberAuth); err != nil {
					return
				}
				if g.IsEmpty(MemberAuth) {
					err = gerror.New("请先绑定微信小程序")
					return
				}
				if g.IsEmpty(MemberAuth.AuthId) {
					err = gerror.New("请先绑定微信小程序")
					return
				}
				if CreateWxMiniPayOrder, err = paycloudClient.WxMiniOrder(ctx, &paycloud.CreateWxMiniPayOrderParams{
					MerchantOrderNo: PmsTransactionInfo.TransactionSn,
					OrderAmount:     DefaultAmount,
					SubOpenid:       MemberAuth.AuthId,
				}); err != nil {
					return
				}
				out.WxMiniPayParams = gjson.New(CreateWxMiniPayOrder.Data.PayParams).String()
				PayParams = out.WxMiniPayParams
			}
			if WxminiPaychannel == "mlilife" {
				PayChannel = "MLILIFE"
				if err = dao.PmsMemberAuth.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsMemberAuth.Columns().MemberId: MemberUser.Id,
					dao.PmsMemberAuth.Columns().Channel:  "WX_MINI",
				}).Scan(&MemberAuth); err != nil {
					return
				}
				if g.IsEmpty(MemberAuth) {
					err = gerror.New("请先绑定微信小程序")
					return
				}
				if g.IsEmpty(MemberAuth.AuthId) {
					err = gerror.New("请先绑定微信小程序")
					return
				}
				if mlilifeWxPayWxappPayResponse, err = mlilifeWxPay.WxappPay(ctx, &mlilifeWxPay.WxappPayRequest{
					Openid:     MemberAuth.AuthId,
					Body:       "Apartment 11",
					TotalFee:   gvar.New(DefaultAmount).Int(),
					OutTradeNo: PmsTransactionInfo.TransactionSn,
				}); err != nil {
					return
				}
				out.WxMiniPayParams = gjson.New(mlilifeWxPayWxappPayResponse.Data).String()
				PayParams = out.WxMiniPayParams
				TransNo = ""
			}

		} else if in.PayChannel == "Paypay_h5" || in.PayChannel == "Paypay_app" {
			if CreatePayPayOrder, err = paycloudClient.PayPayOrder(ctx, &paycloud.CreatePayPayOrderParams{
				MerchantOrderNo: PmsTransactionInfo.TransactionSn,
				OrderAmount:     DefaultAmount,
			}); err != nil {
				return
			}
			out.AppPayParams = CreatePayPayOrder.Data.PayUrl
			PayParams = out.WebPayUrl
		}
	}

	// 更新三方支付信息

	PmsTransactionUpdate := &entity.PmsTransaction{
		TransactionSn:    PmsTransactionInfo.TransactionSn,
		PayChannel:       PayChannel,
		PayType:          in.PayChannel,
		PayParams:        PayParams,
		PaymentRequestId: TransNo,
		ExpiredTime:      gtime.Now().Add(gtime.M * 10),
	}

	if PayChannel == "H5_FX" {
		PmsTransactionUpdate.IsFx = "Y"
	}

	if _, err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().Id: PmsTransactionInfo.Id,
	}).OmitEmptyData().Data(PmsTransactionUpdate).Update(); err != nil {
		return
	}
	return
}

func getOrderIsFxByOrderSn(ctx context.Context, orderSn string) (isFx string, err error) {
	isFx = "N"
	if g.IsEmpty(orderSn) {
		return
	}
	switch orderSn[:1] {
	case "C":
		var carOrder *entity.CarOrder
		if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().OrderSn, orderSn).Scan(&carOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(carOrder) && !g.IsEmpty(carOrder.IsFx) {
			isFx = carOrder.IsFx
		}
	case "S":
		var spaOrder *entity.SpaOrder
		if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().OrderSn, orderSn).Scan(&spaOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(spaOrder) && !g.IsEmpty(spaOrder.IsFx) {
			isFx = spaOrder.IsFx
		}
	case "F":
		var foodOrder *entity.FoodOrder
		if err = dao.FoodOrder.Ctx(ctx).Where(dao.FoodOrder.Columns().OrderSn, orderSn).Scan(&foodOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(foodOrder) && !g.IsEmpty(foodOrder.IsFx) {
			isFx = foodOrder.IsFx
		}
	case "B":
		var cabinetOrder *entity.CabinetOrder
		if err = dao.CabinetOrder.Ctx(ctx).Where(dao.CabinetOrder.Columns().OrderSn, orderSn).Scan(&cabinetOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(cabinetOrder) && !g.IsEmpty(cabinetOrder.IsFx) {
			isFx = cabinetOrder.IsFx
		}
	case "H":
		if len(orderSn) >= 2 && orderSn[:2] == "HC" {
			var changeOrder *entity.PmsAppReservationChange
			if err = dao.PmsAppReservationChange.Ctx(ctx).Where(dao.PmsAppReservationChange.Columns().ChangeOrderSn, orderSn).Scan(&changeOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			if !g.IsEmpty(changeOrder) && !g.IsEmpty(changeOrder.IsFx) {
				isFx = changeOrder.IsFx
			}
			return
		}
		var stayOrder *entity.PmsAppStay
		if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, orderSn).Scan(&stayOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if !g.IsEmpty(stayOrder) && !g.IsEmpty(stayOrder.IsFx) {
			isFx = stayOrder.IsFx
		}
	}
	if g.IsEmpty(isFx) {
		isFx = "N"
	}
	return
}

// BalancePay 纯余额支付
func (s *sPayService) BalancePay(ctx context.Context, in *input_pay.BalancePayInp) (out *input_pay.BalancePayModel, err error) {
	var (
		PmsAppStayInfo              *entity.PmsAppStay
		CarOrder                    *entity.CarOrder
		FoodOrder                   *entity.FoodOrder
		SpaOrder                    *entity.SpaOrder
		CabinetOrder                *entity.CabinetOrder
		TravelOrder                 *entity.TravelOrder
		orderTransaction            *entity.PmsTransaction
		PmsTransactionDoneBalAmount float64
		PmsTransactionThirdAmount   float64
		PmsTransactionBalAmount     float64
		PmsTransactionCouponAmount  float64
		tx                          gdb.TX
		OrderAmount                 float64
		MemberId                    int
		PayType                     string
	)
	// 纯余额支付
	g.Log().Info(ctx, "MemberBalancePay", in)
	g.Log().Info(ctx, "纯余额支付")
	if g.IsEmpty(in) || len(in.OrderSn) < 1 {
		err = gerror.New("订单错误,无法进行支付")
		return
	}
	if err = dao.PmsTransaction.Ctx(ctx).
		Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
		OrderDesc(dao.PmsTransaction.Columns().Id).
		Limit(1).
		Scan(&orderTransaction); err != nil {
		return
	}
	balanceOrderIsFx, err := getOrderIsFxByOrderSn(ctx, in.OrderSn)
	if err != nil {
		return
	}
	if contexts.GetMemberUser(ctx).IsFx && balanceOrderIsFx != "Y" {
		err = gerror.New("非分销订单禁止分销账号支付")
		return
	}
	// 开启事务
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()

			if in.OrderSn[:1] == "F" {
				// TODO -餐饮订单支付完成
				if err = service.PayService().HandleFoodOrderMq(ctx, in.OrderSn, g.Log()); err != nil {
					g.Log().Error(ctx, err.Error())
				}
			}
		}
	}()
	if in.OrderSn[:1] == "H" {
		// 查询住宿订单
		if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).Scan(&PmsAppStayInfo); err != nil {
			return
		}
		if g.IsEmpty(PmsAppStayInfo) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if PmsAppStayInfo.PayModel != 1 {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if PmsAppStayInfo.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
			err = gerror.New("订单已过期,无法进行支付")
			return
		}
		OrderAmount = PmsAppStayInfo.OrderAmount
		MemberId = PmsAppStayInfo.MemberId
		PayType = "住宿"
	} else if in.OrderSn[:1] == "C" {
		// 查询出现订单
		if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).Scan(&CarOrder); err != nil {
			return
		}
		if g.IsEmpty(CarOrder) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if CarOrder.PayModel != 1 {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if CarOrder.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
			err = gerror.New("订单已过期,无法进行支付")
			return
		}
		OrderAmount = CarOrder.OrderAmount
		MemberId = gvar.New(CarOrder.MemberId).Int()
		PayType = "出行"
	} else if in.OrderSn[:1] == "F" {
		// 查询餐饮订单
		if err = dao.FoodOrder.Ctx(ctx).Where(dao.FoodOrder.Columns().OrderSn, in.OrderSn).Scan(&FoodOrder); err != nil {
			return
		}
		if g.IsEmpty(FoodOrder) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if FoodOrder.PayModel != 1 {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if FoodOrder.PayStep == "DEPOSIT" {
			if FoodOrder.DepositExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
				err = gerror.New("订单已过期,无法进行支付")
				return
			}
		} else {
			if FoodOrder.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
				err = gerror.New("订单已过期,无法进行支付")
				return
			}
		}
		OrderAmount = FoodOrder.OrderAmount
		MemberId = gvar.New(FoodOrder.MemberId).Int()
		PayType = "餐饮"
	} else if in.OrderSn[:1] == "S" {
		// 查询 Spa 订单
		if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().OrderSn, in.OrderSn).Scan(&SpaOrder); err != nil {
			return
		}
		if g.IsEmpty(SpaOrder) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if SpaOrder.PayModel != 1 {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if SpaOrder.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
			err = gerror.New("订单已过期,无法进行支付")
			return
		}
		OrderAmount = SpaOrder.OrderAmount
		MemberId = gvar.New(SpaOrder.MemberId).Int()
		PayType = "按摩"
	} else if in.OrderSn[:1] == "B" {
		// 查询储物柜订单
		if err = dao.CabinetOrder.Ctx(ctx).Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&CabinetOrder); err != nil {
			return
		}
		if g.IsEmpty(CabinetOrder) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if (CabinetOrder.PayStep == "BASE" && CabinetOrder.PayModel != 1) || (CabinetOrder.PayStep == "OVERTIME" && CabinetOrder.OvertimePayModel != 1) {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if CabinetOrder.PayStep == "BASE" && CabinetOrder.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
			err = gerror.New("订单已过期,无法进行支付")
			return
		}
		OrderAmount = CabinetOrder.OrderAmount
		MemberId = gvar.New(CabinetOrder.MemberId).Int()
		PayType = "储物柜"
	} else if in.OrderSn[:1] == "T" {
		// 查询一日游订单
		if err = dao.TravelOrder.Ctx(ctx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Scan(&TravelOrder); err != nil {
			return
		}
		if g.IsEmpty(TravelOrder) {
			err = gerror.New("订单错误,无法进行支付")
			return
		}
		if TravelOrder.PayModel != 1 {
			err = gerror.New("不是纯余额支付订单,无法进行支付")
			return
		}
		if TravelOrder.ExpirationTime < gvar.New(gtime.Now().Unix()).Int() {
			err = gerror.New("订单已过期,无法进行支付")
			return
		}
		OrderAmount = TravelOrder.OrderAmount
		MemberId = gvar.New(TravelOrder.MemberId).Int()
		PayType = "一日游"
	} else {
		err = gerror.New("订单错误,无法进行支付")
		return
	}

	// 查询支付订单金额
	if PmsTransactionBalAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
		dao.PmsTransaction.Columns().PayType:   "BAL",
		dao.PmsTransaction.Columns().PayStatus: "WAIT",
	}).Sum(dao.PmsTransaction.Columns().Amount); err != nil {
		return
	}

	if PmsTransactionCouponAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
		dao.PmsTransaction.Columns().PayType:   "COUPON",
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).Sum(dao.PmsTransaction.Columns().Amount); err != nil {
		return
	}

	if in.OrderSn[:1] == "F" || in.OrderSn[:1] == "B" {
		if PmsTransactionThirdAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
			dao.PmsTransaction.Columns().PayStatus: "DONE",
		}).WhereNot(dao.PmsTransaction.Columns().PayType, "BAL").WhereNot(dao.PmsTransaction.Columns().PayType, "COUPON").Sum(dao.PmsTransaction.Columns().Amount); err != nil {
			return
		}
	}

	if in.OrderSn[:1] == "B" {
		if PmsTransactionDoneBalAmount, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
			dao.PmsTransaction.Columns().PayType:   "BAL",
			dao.PmsTransaction.Columns().PayStatus: "DONE",
		}).Sum(dao.PmsTransaction.Columns().Amount); err != nil {
			return
		}
		if PmsTransactionBalAmount > 0 && (PmsTransactionBalAmount+PmsTransactionCouponAmount+PmsTransactionThirdAmount+PmsTransactionDoneBalAmount) != (OrderAmount+float64(CabinetOrder.OvertimeFee)) {
			err = gerror.New("支付金额错误")
			return
		}
	} else {
		if PmsTransactionBalAmount > 0 && (PmsTransactionBalAmount+PmsTransactionCouponAmount+PmsTransactionThirdAmount) != OrderAmount {
			err = gerror.New("支付金额错误")
			return
		}
	}

	if err = service.PayService().OrderBalancePay(ctx, tx, in.OrderSn, MemberId, PayType); err != nil {
		return
	}

	if in.OrderSn[:1] == "H" {
		if _, err = s.HotelCreateOrder(ctx, &input_pay.HotelCreateOrderInp{
			OrderSn:                 in.OrderSn,
			PmsTransactionBalAmount: PmsTransactionBalAmount,
			Booker:                  PmsAppStayInfo.Booker,
		}, tx); err != nil {
			return
		}
	} else if in.OrderSn[:1] == "C" {
		// TODO -出行订单支付完成
		if err = service.PayService().HandleCarOrderMq(ctx, in.OrderSn, g.Log()); err != nil {
			g.Log().Error(ctx, err.Error())
		}
	} else if in.OrderSn[:1] == "S" {
		// TODO -按摩订单支付完成
		if err = service.PayService().HandleSpaOrderMq(ctx, in.OrderSn, g.Log()); err != nil {
			g.Log().Error(ctx, err.Error())
		}
	} else if in.OrderSn[:1] == "B" {
		// TODO -储物柜订单支付完成
		if _, err = s.CabinetCreateOrder(ctx, &input_pay.CabinetCreateOrderInp{
			OrderSn:                 in.OrderSn,
			PmsTransactionBalAmount: PmsTransactionBalAmount,
		}, tx); err != nil {
			return
		}
	} else if in.OrderSn[:1] == "T" {
		// TODO -一日游订单支付完成
		if _, err = s.TravelCreateOrder(ctx, &input_pay.TravelCreateOrderInp{
			OrderSn:                 in.OrderSn,
			PmsTransactionBalAmount: PmsTransactionBalAmount,
		}, tx); err != nil {
			return
		}
	}
	return
}

func (s *sPayService) HotelCreateOrder(ctx context.Context, in *input_pay.HotelCreateOrderInp, tx gdb.TX) (out *input_pay.HotelCreateOrderModel, err error) {
	var (
		cookieBookerParams *airhousePublicApi.CreateStayJSONDataRequest
		BookerResponse     *airhousePublicApi.CreateStayJSONDataResponse
	)
	// 修改住宿订单状态
	if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).Update(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderStatus: "HAVE_PAID",
	}); err != nil {
		return
	}
	// 修改房间订单状态
	if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().OrderSn, in.OrderSn).Update(g.MapStrAny{
		dao.PmsAppReservation.Columns().OrderStatus: "HAVE_PAID",
	}); err != nil {
		return
	}
	// 修改支付订单状态
	if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
		Where(dao.PmsTransaction.Columns().ExpiredTime+" > ?", gtime.Now().Format("Y-m-d H:i:s")).
		Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Update(entity.PmsTransaction{
			PayAmount: in.PmsTransactionBalAmount,
			PayStatus: "DONE",
			PayTime:   gtime.Now(),
		}); err != nil {
		return
	}
	// 获取 AIR_HOST 下单参数
	if cookieBookerParams, err = service.HotelService().CreateStayParams(ctx, in.OrderSn); err != nil {
		return
	}
	// 析构函数 提交事务
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
			if err = service.HotelService().OrderExpiration(ctx, in.OrderSn); err != nil {
				err = gerror.New("订单内部支付失败情况下自动取消订单失败")
				return
			}
		} else {
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
	// 提交订单
	if BookerResponse, err = service.HotelService().CreateStay(ctx, cookieBookerParams, tx); err != nil {
		return
	}
	// 修改住宿订单为支付完成
	if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).Update(g.Map{
		dao.PmsAppStay.Columns().OrderStatus: "HAVE_PAID",
		dao.PmsAppStay.Columns().Uuid:        BookerResponse.Data.ID,
		dao.PmsAppStay.Columns().Booker:      BookerResponse.Data.Booker.ID,
	}); err != nil {
		return
	}
	// 修改入住人ID
	if _, err = dao.PmsGuestProfile.Ctx(ctx).TX(tx).Where(g.Map{
		dao.PmsGuestProfile.Columns().Uid: in.Booker,
	}).Data(g.Map{
		dao.PmsGuestProfile.Columns().Uid: BookerResponse.Data.Booker.ID,
	}).Update(); err != nil {
		return
	}

	// 修改房间订单为支付完成
	if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().OrderSn, in.OrderSn).Update(g.Map{
		dao.PmsAppReservation.Columns().OrderStatus: "HAVE_PAID",
		dao.PmsAppReservation.Columns().MainGuest:   BookerResponse.Data.Booker.ID,
	}); err != nil {
		return
	}

	// 酒店订单支付成功
	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: in.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		err = gerror.New("该订单无需处理")
		return
	}

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

	sendMsg, _ := json.Marshal(g.Map{
		"type":        "AWARD",
		"id":          gvar.New(PmsAppStay.Id).Int(),
		"checkInDate": PmsAppStay.CheckInDate,
		"memberId":    PmsAppStay.MemberId,
	})
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderAward,
		DataByte:     sendMsg,
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
		OperatorId:           PmsAppStay.MemberId,
		OperatorRole:         "MEMBER",
		OrderSn:              PmsAppStay.OrderSn,
	})

	//err = service.HotelService().HotelOrderAward(ctx, tx, gvar.New(PmsAppStay.Id).Int(), PmsAppStay.CheckInDate, PmsAppStay.MemberId)

	return
}

// ThirdPayCompleted 三方支付完成后处理支付流水
func (s *sPayService) ThirdPayCompleted(ctx context.Context, in *input_pay.ThirdPayInp) (err error) {
	var (
		PmsTransaction          *entity.PmsTransaction
		PmsAppStay              *entity.PmsAppStay
		FoodOrder               *entity.FoodOrder
		SpaOrder                *entity.SpaOrder
		CarOrder                *entity.CarOrder
		CabinetOrder            *entity.CabinetOrder
		TravelOrder             *entity.TravelOrder
		PmsAppReservationChange *entity.PmsAppReservationChange
		lockKey                 = "nitifyOrder:" + in.TransNo
		OrderSn                 string
		OrderStatus             string
		MemberId                int
		PayType                 string
	)
	defer func() {
		if err != nil {
			_ = service.WarningWorkWx().SendWarningWorkWx(ctx, `
### 三方支付完成后流水处理 ThirdPayCompleted
@高杨
*  链路ID：`+gctx.CtxId(ctx)+`
* `+err.Error()+`
* 订单号：`+OrderSn+`
* 支付流水号：`+in.TransNo+`
* 用户ID:`+gvar.New(MemberId).String())
		}
	}()
	// 内存锁 防止重复执行该逻辑
	gmlock.Lock(lockKey)
	defer gmlock.Unlock(lockKey)
	defer func() {
		g.Log().Infof(ctx, "支付前订单状态：%s", OrderStatus)
		if err == nil {
			mqstruct := &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNamePlaceOrder,
				DataByte:     []byte(OrderSn),
				Header:       nil,
			}
			g.Log().Info(ctx, mqstruct)
			if err = rabbitmq.SendMqMessage(ctx, mqstruct); err != nil {
				return
			}
		}
	}()
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 查询支付信息
		if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().PaymentRequestId, in.TransNo).Scan(&PmsTransaction); err != nil {
			err = gerror.New("查询支付信息失败")
			return
		}
		if g.IsEmpty(PmsTransaction) {
			err = gerror.New("支付流水不存在")
			return
		}
		defer func() {
			g.Log().Infof(ctx, "支付前订单状态：%s", OrderStatus)
			if OrderStatus == "CANCEL" || OrderStatus == "FAIL" || err != nil {
				// 执行外部支付成功后续失败的情况下退款
				err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
					OrderSn:      OrderSn,
					RefundAmount: PmsTransaction.Amount,
				}, tx)
			}
		}()

		// 查询支付信息
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().PaymentRequestId, in.TransNo).Scan(&PmsTransaction); err != nil {
			err = gerror.New("查询支付信息失败")
			return
		}
		if g.IsEmpty(PmsTransaction) {
			err = gerror.New("支付流水不存在")
			return
		}

		// TODO 支付回调
		if PmsTransaction.OrderType == "CAR" {
			// 查询出行订单
			if err = dao.CarOrder.Ctx(ctx).TX(tx).Where(dao.CarOrder.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&CarOrder); err != nil {
				return
			}
			OrderStatus = CarOrder.OrderStatus
			OrderSn = CarOrder.OrderSn
			MemberId = gvar.New(CarOrder.MemberId).Int()
			PayType = "出行"
		} else if PmsTransaction.OrderType == "SPA" {
			// 查询按摩订单
			if err = dao.SpaOrder.Ctx(ctx).TX(tx).Where(dao.SpaOrder.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&SpaOrder); err != nil {
				return
			}
			OrderStatus = SpaOrder.OrderStatus
			OrderSn = SpaOrder.OrderSn
			MemberId = gvar.New(SpaOrder.MemberId).Int()
			PayType = "按摩"
		} else if PmsTransaction.OrderType == "FOOD" {
			// 查询餐饮订单
			if err = dao.FoodOrder.Ctx(ctx).TX(tx).Where(dao.FoodOrder.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&FoodOrder); err != nil {
				return
			}
			OrderStatus = FoodOrder.OrderStatus
			OrderSn = FoodOrder.OrderSn
			MemberId = gvar.New(FoodOrder.MemberId).Int()
			PayType = "餐饮"
		} else if PmsTransaction.OrderType == "CABINET" {
			// 查询储物柜订单
			if err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&CabinetOrder); err != nil {
				return
			}
			OrderStatus = CabinetOrder.OrderStatus
			OrderSn = CabinetOrder.OrderSn
			MemberId = gvar.New(CabinetOrder.MemberId).Int()
			PayType = "储物柜"
		} else if PmsTransaction.OrderType == "TRAVEL" {
			// 查询储物柜订单
			if err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&TravelOrder); err != nil {
				return
			}
			OrderStatus = TravelOrder.OrderStatus
			OrderSn = TravelOrder.OrderSn
			MemberId = gvar.New(TravelOrder.MemberId).Int()
			PayType = "一日游"
		} else if PmsTransaction.OrderType == "BOOKING" {
			// 查询住宿单信息
			if err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, PmsTransaction.OrderSn).Scan(&PmsAppStay); err != nil {
				err = gerror.New("查询住宿单信息失败")
				return
			}
			OrderStatus = PmsAppStay.OrderStatus
			OrderSn = PmsAppStay.OrderSn
			MemberId = gvar.New(PmsAppStay.MemberId).Int()
			PayType = "住宿"
		} else if PmsTransaction.OrderType == "BOOKING_CHANGE" {
			// 查询住宿变更单
			if err = dao.PmsAppReservationChange.Ctx(ctx).TX(tx).
				Where(dao.PmsAppReservationChange.Columns().ChangeOrderSn, PmsTransaction.ChangeOrderSn).
				Scan(&PmsAppReservationChange); err != nil {
				err = gerror.New("查询住宿单信息失败")
				return
			}
			OrderStatus = PmsAppReservationChange.ChangeStatus
			OrderSn = PmsAppReservationChange.ChangeOrderSn
		} else {
			err = gerror.New("订单类型错误")
			return
		}

		updateData := g.Map{
			dao.PmsTransaction.Columns().PayStatus: "DONE",
			dao.PmsTransaction.Columns().PayTime:   gtime.Now(),
		}
		if PmsTransaction.PayType != "Paypal" {
			updateData[dao.PmsTransaction.Columns().PayAmount] = PmsTransaction.Amount
		}
		// 更新第三方支付信息为支付完成
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().PaymentRequestId, in.TransNo).Update(updateData); err != nil {
			err = gerror.New("第三方支付信息更新失败")
			return
		}
		if PmsTransaction.Scene == "BOOKING" {
			if err = s.PayAmountChange(ctx, tx, PmsTransaction.OrderSn); err != nil {
				return
			}
		}

		if OrderStatus == "WAIT_PAY" || OrderStatus == "ING" {
			if PmsTransaction.OrderType != "BOOKING_CHANGE" {
				if err = service.PayService().OrderBalancePay(ctx, tx, OrderSn, MemberId, PayType); err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (s *sPayService) PayAmountChange(ctx context.Context, tx gdb.TX, OrderSn string) (err error) {
	var (
		PmsTransaction []*entity.PmsTransaction
		AllPayPrice    float64
	)
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   OrderSn,
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).Scan(&PmsTransaction); err != nil {
		return
	}
	if g.IsEmpty(PmsTransaction) {
		return
	}
	// 计算总计支付金额
	for _, Transaction := range PmsTransaction {
		AllPayPrice = decimal.NewFromFloat(AllPayPrice).Add(decimal.NewFromFloat(Transaction.PayAmount)).Round(0).InexactFloat64()
	}
	// 更新到主订单
	_, err = dao.PmsAppStay.Ctx(ctx).TX(tx).
		Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
		Update(g.Map{
			dao.PmsAppStay.Columns().OrderAmount: AllPayPrice,
		})
	return
}

func (s *sPayService) HandleFoodOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error) {
	// TODO 餐饮订单支付完成业务
	var (
		gHttpClient = g.Client()
		response    *gclient.Response
	)
	gHttpClient.SetHeader("Content-Type", "application/json")
	if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/food/pay/payNotify", g.MapStrAny{
		"orderSn": OrderSn,
	}); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())

	// 打印
	_ = service.BasicsPrinter().PrinterFoodOrder(ctx, &input_basics.PrinterCarOrderInp{
		OrderSn: OrderSn,
	})

	if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
		return
	}
	return
}

func (s *sPayService) HandleSpaOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error) {
	// TODO SPA订单支付完成业务
	var (
		gHttpClient = g.Client()
		response    *gclient.Response
	)
	gHttpClient.SetHeader("Content-Type", "application/json")
	if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/spa/pay/payNotify", g.MapStrAny{
		"orderSn": OrderSn,
	}); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())

	// 发送预警短信
	_ = service.BasicsSmsLog().SendPlaceOrderMsg(ctx, &input_basics.SendMsgInp{
		Event:   "spaPlaceOrder",
		OrderSn: OrderSn,
	})

	// 打印
	_ = service.BasicsPrinter().PrinterSpaOrder(ctx, &input_basics.PrinterCarOrderInp{
		OrderSn: OrderSn,
	})

	if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
		return
	}
	return
}

func (s *sPayService) HandleCarOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error) {
	// TODO CAR订单支付完成业务
	var (
		gHttpClient = g.Client()
		response    *gclient.Response
	)
	gHttpClient.SetHeader("Content-Type", "application/json")
	if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/car/pay/payNotify", g.MapStrAny{
		"orderSn": OrderSn,
	}); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())

	// 发送预警短信
	_ = service.BasicsSmsLog().SendPlaceOrderMsg(ctx, &input_basics.SendMsgInp{
		Event:   "carPlaceOrder",
		OrderSn: OrderSn,
	})

	// 打印
	_ = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
		OrderSn: OrderSn,
	})

	if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
		return
	}
	return
}

func (s *sPayService) CabinetCreateOrder(ctx context.Context, in *input_pay.CabinetCreateOrderInp, tx gdb.TX) (out *input_pay.CabinetCreateOrderModel, err error) {
	var CabinetOrder *entity.CabinetOrder
	if err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&CabinetOrder); err != nil {
		err = gerror.New("查询订单信息失败")
		return
	}

	if CabinetOrder.PayStep == "BASE" {
		// 修改订单状态
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Update(g.MapStrAny{
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

		// 请求mch那边下单接口，接口返回成功则继续，失败则回滚
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
			return
		}
		if cabinetResponse.Code != 0 {
			err = gerror.New(cabinetResponse.Msg)
			return
		} else {
			AddressNameJson := &input_cabinet.LanguageJson{
				Zh: cabinetResponse.Data.AddressZh,
				En: cabinetResponse.Data.AddressEn,
				Ja: cabinetResponse.Data.AddressJa,
				Ko: cabinetResponse.Data.AddressKo,
				Tw: cabinetResponse.Data.AddressTw,
			}
			if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
				WherePri(CabinetOrder.Id).Data(g.MapStrAny{
				dao.CabinetOrder.Columns().OutOrderSn:   cabinetResponse.Data.OrderNo,
				dao.CabinetOrder.Columns().BoxId:        cabinetResponse.Data.BoxId,
				dao.CabinetOrder.Columns().BoxNo:        cabinetResponse.Data.BoxNo,
				dao.CabinetOrder.Columns().BoxAlias:     cabinetResponse.Data.BoxAlias,
				dao.CabinetOrder.Columns().Pin:          cabinetResponse.Data.Pin,
				dao.CabinetOrder.Columns().Address:      cabinetResponse.Data.AddressZh,
				dao.CabinetOrder.Columns().AddressJson:  gjson.New(AddressNameJson),
				dao.CabinetOrder.Columns().StartTime:    cabinetResponse.Data.StartTime,
				dao.CabinetOrder.Columns().EndTime:      cabinetResponse.Data.EndTime,
				dao.CabinetOrder.Columns().GraceSeconds: cabinetResponse.Data.GraceSeconds,
				dao.CabinetOrder.Columns().OrderStatus:  "ING",
			}).Update(); err != nil {
				return
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
		// 超时费支付成功
		// 修改订单为支付完成
		OrderAmount := decimal.NewFromFloat(CabinetOrder.OrderAmount).Add(decimal.NewFromFloat(float64(CabinetOrder.OvertimeFee))).Round(2)
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Update(g.Map{
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
			return
		}
		//if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
		//	WherePri(CabinetOrder.Id).Data(g.MapStrAny{
		//	dao.CabinetOrder.Columns().OrderStatus: "GRACE",
		//}).Update(); err != nil {
		//
		//}
		if cabinetResponse.Code == 0 {
			if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
				WherePri(CabinetOrder.Id).Data(g.MapStrAny{
				dao.CabinetOrder.Columns().GraceEndTime: cabinetResponse.Data.GraceEndTime,
				dao.CabinetOrder.Columns().OrderStatus:  "GRACE",
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
			err = gerror.New(cabinetResponse.Msg)
			return
		}
	}
	// 修改支付订单状态
	//if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
	//	Where(dao.PmsTransaction.Columns().ExpiredTime+" > ?", gtime.Now().Format("Y-m-d H:i:s")).
	//	Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
	//	Where(dao.PmsTransaction.Columns().PayType, "BAL").
	//	Update(entity.PmsTransaction{
	//		PayAmount: in.PmsTransactionBalAmount,
	//		PayStatus: "DONE",
	//		PayTime:   gtime.Now(),
	//	}); err != nil {
	//	return
	//}

	return
}

func (s *sPayService) TravelCreateOrder(ctx context.Context, in *input_pay.TravelCreateOrderInp, tx gdb.TX) (out *input_pay.TravelCreateOrderModel, err error) {
	var TravelOrder *entity.TravelOrder
	if err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Scan(&TravelOrder); err != nil {
		err = gerror.New("查询订单信息失败")
		return
	}

	// 修改订单状态
	if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Update(g.MapStrAny{
		dao.TravelOrder.Columns().PayTime:     gtime.Now().Format("Y-m-d H:i:s"),
		dao.TravelOrder.Columns().PayStatus:   "HAVE_PAID",
		dao.TravelOrder.Columns().OrderStatus: "WAIT_VERIFY",
	}); err != nil {
		return
	}

	// 支付成功增加产品销量字段
	if _, err = dao.TravelProduct.Ctx(ctx).TX(tx).
		Where(dao.TravelProduct.Columns().Id, TravelOrder.ProductId).
		Increment(dao.TravelProduct.Columns().SalesNum, TravelOrder.BookingNum); err != nil {
		return
	}

	// 支付成功增加产品Sku销量字段
	if _, err = dao.TravelProductSku.Ctx(ctx).TX(tx).
		Where(dao.TravelProductSku.Columns().Id, TravelOrder.SkuId).
		Increment(dao.TravelProductSku.Columns().SalesNum, TravelOrder.BookingNum); err != nil {
		return
	}

	// 修改支付订单状态
	if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
		Where(dao.PmsTransaction.Columns().ExpiredTime+" > ?", gtime.Now().Format("Y-m-d H:i:s")).
		Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Update(entity.PmsTransaction{
			PayAmount: in.PmsTransactionBalAmount,
			PayStatus: "DONE",
			PayTime:   gtime.Now(),
		}); err != nil {
		return
	}

	// 订单支付成功日志
	if _, err = dao.TravelOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.TravelOrderLog{
		OrderId:     int(TravelOrder.Id),
		OrderStatus: "WAIT_VERIFY",
		ActionWay:   "HAVE_PAID",
		Remark:      "订单支付待核销",
		OperateType: "USER",
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

	// 发送到消息队列
	// systemMessageTitle := map[string]string{
	// 	"zh":    "订单支付成功",
	// 	"en":    "Order payment successful",
	// 	"ja":    "注文の支払いが完了しました",
	// 	"ko":    "주문 결제 성공",
	// 	"zh_CN": "訂單支付成功",
	// }
	// systemMessageContent := map[string]string{
	// 	"zh":    TravelOrder.OrderSn + "订单支付成功",
	// 	"en":    "Order " + TravelOrder.OrderSn + " payment successful",
	// 	"ja":    "注文" + TravelOrder.OrderSn + "の支払いが完了しました",
	// 	"ko":    "주문 " + TravelOrder.OrderSn + " 결제 성공",
	// 	"zh_CN": TravelOrder.OrderSn + "訂單支付成功",
	// }
	// // 查询用户的手机号区号 来判断用户语言
	// phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(TravelOrder.MemberId).Value()
	// var memberLanguage string
	// if phoneArea.String() == "+86" {
	// 	memberLanguage = "zh"
	// } else if phoneArea.String() == "+81" {
	// 	memberLanguage = "ja"
	// } else if phoneArea.String() == "+82" {
	// 	memberLanguage = "ko"
	// } else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
	// 	memberLanguage = "zh_CN"
	// } else {
	// 	memberLanguage = "en"
	// }
	// appPushData := g.MapStrStr{
	// 	"type":   "1",
	// 	"string": TravelOrder.OrderSn,
	// }
	// service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
	// 	SystemMessageTitle:   systemMessageTitle,
	// 	SystemMessageContent: systemMessageContent,
	// 	Scene:                "travel",
	// 	Type:                 "order",
	// 	MemberId:             int(TravelOrder.MemberId),
	// 	Language:             memberLanguage,
	// 	AppPushData:          appPushData,
	// 	AppLink:              "/smart_locker_order_detail",
	// 	WxLink:               fmt.Sprintf("/subpackages/lockers/pages/order-detail?orderSn=%s", TravelOrder.OrderSn),
	// 	EnablePush:           true,
	// 	EnableSms:            false,
	// 	PushTitle:            systemMessageTitle[memberLanguage],
	// 	PushContent:          systemMessageContent[memberLanguage],
	// 	OperatorId:           int(TravelOrder.MemberId),
	// 	OperatorRole:         "MEMBER",
	// 	OrderSn:              TravelOrder.OrderSn,
	// })

	return
}

func (s *sPayService) OrderBalancePay(ctx context.Context, tx gdb.TX, OrderSn string, MemberId int, PayType string) (err error) {
	var (
		PmsBalTransaction    *entity.PmsTransaction
		PmsCouponTransaction *entity.PmsTransaction
	)
	// 查询是否存在未支付的余额支付
	g.Log().Info(ctx, "查询是否存在未支付的余额支付")
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
		Where(dao.PmsTransaction.Columns().PayType, "BAL").
		Scan(&PmsBalTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	// 是否存在未支付的余额支付信息
	if !g.IsEmpty(PmsBalTransaction) {
		// 扣除余额
		if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
			MemberId:      MemberId,
			Scene:         PmsBalTransaction.Scene,
			Type:          "CONSUME",
			ChangeBalance: PmsBalTransaction.Amount * -1,
			OrderSn:       PmsBalTransaction.OrderSn,
			Reason:        fmt.Sprintf("支付%s费用", PayType),
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

	// 查询优惠券支付明细情况
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   OrderSn,
		dao.PmsTransaction.Columns().PayType:   "COUPON",
		dao.PmsTransaction.Columns().PayStatus: "WAIT",
	}).Scan(&PmsCouponTransaction); err != nil {
		return
	}
	if !g.IsEmpty(PmsCouponTransaction) {
		// 核销优惠券
		if _, err = dao.PmsCoupon.Ctx(ctx).TX(tx).Where(dao.PmsCoupon.Columns().Id, PmsCouponTransaction.CouponId).Update(g.MapStrAny{
			dao.PmsCoupon.Columns().State:   2,
			dao.PmsCoupon.Columns().UseTime: gtime.Now(),
		}); err != nil {
			return
		}
	}
	return
}
