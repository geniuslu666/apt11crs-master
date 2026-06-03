package logic_terminal

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_terminal"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

type sTerminalTerminal struct{}

func NewTerminalTerminal() *sTerminalTerminal {
	return &sTerminalTerminal{}
}

func init() {
	service.RegisterTerminalTerminal(NewTerminalTerminal())
}

func (s *sTerminalTerminal) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysTerminalVerify.Ctx(ctx), option...)
}

func (s *sTerminalTerminal) List(ctx context.Context, in *input_terminal.VerifyListInp) (list []*input_terminal.VerifyListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.Fields(input_terminal.VerifyListModel{})

	if !g.IsEmpty(in.TerminalId) {
		mod = mod.Where(dao.SysTerminalVerify.Columns().TerminalId, in.TerminalId)
	}
	if !g.IsEmpty(in.MchId) {
		mod = mod.Where(dao.SysTerminalVerify.Columns().MchId, in.MchId)
	}
	if !g.IsEmpty(in.StoreId) {
		mod = mod.Where(dao.SysTerminalVerify.Columns().StoreId, in.StoreId)
	}
	if !g.IsEmpty(in.RestaurantId) {
		mod = mod.Where(dao.SysTerminalVerify.Columns().RestaurantId, in.RestaurantId)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.SysTerminalVerify.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (s *sTerminalTerminal) View(ctx context.Context, in *input_terminal.VerifyLogViewInp) (res *input_terminal.VerifyLogViewModel, err error) {
	var (
		TerminalVerify   *entity.SysTerminalVerify
		StoreInfo        *entity.ThMchStore
		FoodOrderInfo    *input_terminal.VerifyFoodOrderInfo
		MemberCouponInfo *input_terminal.VerifyMemberCouponInfo
	)

	MemberInfo := contexts.GetTerminalUser(ctx)

	if err = s.Model(ctx).Hook(hook.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&TerminalVerify); err != nil {
		err = gerror.Wrap(err, "获取核销记录信息失败，请稍后重试！")
		return
	}

	res = new(input_terminal.VerifyLogViewModel)
	res.Id = in.Id
	res.VerifyType = TerminalVerify.VerifyType
	res.VerifyInfo = new(input_terminal.VerifyInfo)
	res.VerifyInfo.VerifyTime = TerminalVerify.VerifyTime
	res.VerifyInfo.Account = MemberInfo.Account
	if TerminalVerify.VerifyType == "FOOD_ORDER" {
		res.FoodOrderInfo = new(input_terminal.FoodOrderInfo)
		if err = dao.FoodOrder.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(TerminalVerify.FoodOrderId).Scan(&FoodOrderInfo); err != nil {
			err = gerror.Wrap(err, "获取餐厅订单信息失败，请稍后重试！")
			return
		}
		res.VerifyInfo.VerifyStore = FoodOrderInfo.FoodOrderRestaurantInfo.Name
		res.FoodOrderInfo.OrderSn = FoodOrderInfo.OrderSn
		res.FoodOrderInfo.BookingName = FoodOrderInfo.BookingName
		res.FoodOrderInfo.PhoneArea = FoodOrderInfo.PhoneArea
		res.FoodOrderInfo.BookingMobile = FoodOrderInfo.BookingMobile
		res.FoodOrderInfo.GoodsNum = FoodOrderInfo.GoodsNum
		res.FoodOrderInfo.BookDate = FoodOrderInfo.BookDate
		res.FoodOrderInfo.BookTime = FoodOrderInfo.BookTime
		res.FoodOrderInfo.RestaurantName = FoodOrderInfo.FoodOrderRestaurantInfo.Name
		res.FoodOrderInfo.RestaurantAddress = FoodOrderInfo.FoodOrderRestaurantInfo.DetailAddress
		res.FoodOrderInfo.GoodsName = FoodOrderInfo.FoodOrderGoodsInfo.GoodsName
		res.FoodOrderInfo.MemberMessage = FoodOrderInfo.MemberMessage
		res.FoodOrderInfo.MemberMessageJa = FoodOrderInfo.MemberMessageJa
		res.FoodOrderInfo.OrderAmount = FoodOrderInfo.OrderAmount
		res.FoodOrderInfo.CreatedTime = FoodOrderInfo.CreatedAt
	} else if TerminalVerify.VerifyType == "TH_COUPON" {
		res.CouponInfo = new(input_terminal.CouponInfo)
		if err = dao.ThMchStore.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).WherePri(TerminalVerify.StoreId).Scan(&StoreInfo); err != nil {
			err = gerror.Wrap(err, "获取门店信息失败，请稍后重试！")
			return
		}
		res.VerifyInfo.VerifyStore = StoreInfo.StoreName

		if err = dao.ThMemberCoupon.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(TerminalVerify.MemberCouponId).Scan(&MemberCouponInfo); err != nil {
			err = gerror.Wrap(err, "获取券信息失败，请稍后重试！")
			return
		}
		res.CouponInfo.CouponName = MemberCouponInfo.CouponInfo.CouponName
		res.CouponInfo.CouponMchName = TerminalVerify.CouponMchName
		res.CouponInfo.StartTime = MemberCouponInfo.StartTime
		res.CouponInfo.EndTime = MemberCouponInfo.EndTime
		res.CouponInfo.MemberNo = MemberCouponInfo.MemberInfo.MemberNo
		res.CouponInfo.StoreAddress = StoreInfo.DetailAddress
	}

	return
}

func (s *sTerminalTerminal) CodeView(ctx context.Context, in *input_terminal.CodeViewInp) (res *input_terminal.CodeViewModel, err error) {
	var (
		VerifyType       string
		MchInfo          *entity.ThMch
		StoreInfo        *entity.ThMchStore
		CouponMchInfo    *entity.ThCouponMch
		FoodOrderInfo    *input_terminal.VerifyFoodOrderInfo
		MemberCouponInfo *input_terminal.VerifyMemberCouponInfo
	)

	MemberInfo := contexts.GetTerminalUser(ctx)

	ScanCode := strings.Split(in.Code, "|")

	if len(ScanCode) == 2 && ScanCode[1] == "FOOD" {
		// |FOOD区分
		if MemberInfo.Type == "STORE" {
			err = gerror.New("门店不能核销餐厅订单")
			return
		}
		VerifyType = "FOOD_ORDER"
	} else {
		if MemberInfo.Type == "RESTAURANT" {
			err = gerror.New("餐厅不能核销礼品券")
			return
		}
		VerifyType = "TH_COUPON"
	}

	res = new(input_terminal.CodeViewModel)
	res.VerifyType = VerifyType
	if VerifyType == "FOOD_ORDER" {
		if err = dao.FoodOrder.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.FoodOrder.Columns().OrderSn, ScanCode[0]).Scan(&FoodOrderInfo); err != nil {
			err = gerror.Wrap(err, "获取餐厅订单信息失败，请稍后重试！")
			return
		}
		if FoodOrderInfo.RestaurantId != int(MemberInfo.RestaurantId) {
			err = gerror.New("该餐厅订单不属于您的餐厅，不可核销")
			return
		}

		if FoodOrderInfo.OrderStatus != "HAVE_PAID" || FoodOrderInfo.BookingStatus != "CONFIRMED" {
			err = gerror.New("订单状态不正确")
			return
		}

		if gtime.Now().After(FoodOrderInfo.BookDatetime) {
			err = gerror.New("您已逾期，无法核销")
			return
		}

		if FoodOrderInfo.VerifyStatus != "WAIT_VERIFY" {
			err = gerror.New("订单核销状态不正确")
			return
		}

		res.FoodOrderInfo = new(input_terminal.FoodOrderInfo)
		res.FoodOrderInfo.OrderSn = FoodOrderInfo.OrderSn
		res.FoodOrderInfo.BookingName = FoodOrderInfo.BookingName
		res.FoodOrderInfo.PhoneArea = FoodOrderInfo.PhoneArea
		res.FoodOrderInfo.BookingMobile = FoodOrderInfo.BookingMobile
		res.FoodOrderInfo.GoodsNum = FoodOrderInfo.GoodsNum
		res.FoodOrderInfo.BookDate = FoodOrderInfo.BookDate
		res.FoodOrderInfo.BookTime = FoodOrderInfo.BookTime
		res.FoodOrderInfo.RestaurantName = FoodOrderInfo.FoodOrderRestaurantInfo.Name
		res.FoodOrderInfo.RestaurantAddress = FoodOrderInfo.FoodOrderRestaurantInfo.DetailAddress
		res.FoodOrderInfo.GoodsName = FoodOrderInfo.FoodOrderGoodsInfo.GoodsName
		res.FoodOrderInfo.MemberMessage = FoodOrderInfo.MemberMessage
		res.FoodOrderInfo.MemberMessageJa = FoodOrderInfo.MemberMessageJa
		res.FoodOrderInfo.OrderAmount = FoodOrderInfo.OrderAmount
		res.FoodOrderInfo.CreatedTime = FoodOrderInfo.CreatedAt
	} else {
		if err = dao.ThMch.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(MemberInfo.MchId).Scan(&MchInfo); err != nil {
			err = gerror.Wrap(err, "获取商户信息失败，请稍后重试！")
			return
		}
		if err = dao.ThMchStore.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(MemberInfo.StoreId).Scan(&StoreInfo); err != nil {
			err = gerror.Wrap(err, "获取门店信息失败，请稍后重试！")
			return
		}
		if err = dao.ThMemberCoupon.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.ThMemberCoupon.Columns().CouponNo, in.Code).Scan(&MemberCouponInfo); err != nil {
			err = gerror.Wrap(err, "获取券信息失败，请稍后重试！")
			return
		}
		if MemberCouponInfo.CouponInfo.UseStatus != 1 {
			err = gerror.New("该礼品券已停止使用！")
			return
		}
		if MemberCouponInfo.State == 1 {
			err = gerror.New("该礼品券还未生效")
			return
		} else if MemberCouponInfo.State == 3 {
			err = gerror.New("该礼品券已核销")
			return
		} else if MemberCouponInfo.State == 4 {
			err = gerror.New("该礼品券已过期")
			return
		} else if MemberCouponInfo.State == 5 {
			err = gerror.New("该礼品券已失效")
			return
		}
		if err = dao.ThCouponMch.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.ThCouponMch.Columns().CouponId, MemberCouponInfo.CouponId).Where(dao.ThCouponMch.Columns().MchId, MemberInfo.MchId).Scan(&CouponMchInfo); err != nil {
			err = gerror.Wrap(err, "获取券商户关联信息失败，请稍后重试！")
			return
		}
		if g.IsEmpty(CouponMchInfo) {
			err = gerror.New("您的商户不能核销该礼品券")
			return
		}
		res.CouponInfo = new(input_terminal.CodeCouponInfo)
		res.CouponInfo.CouponName = MemberCouponInfo.CouponInfo.CouponName
		res.CouponInfo.CouponMchName = CouponMchInfo.Name
		res.CouponInfo.StartTime = MemberCouponInfo.StartTime
		res.CouponInfo.EndTime = MemberCouponInfo.EndTime
		res.CouponInfo.MemberNo = MemberCouponInfo.MemberInfo.MemberNo
		res.CouponInfo.MchName = MchInfo.Name
		res.CouponInfo.StoreName = StoreInfo.StoreName
		res.CouponInfo.StoreAddress = StoreInfo.DetailAddress
	}

	return
}

func (s *sTerminalTerminal) CodeVerify(ctx context.Context, in *input_terminal.CodeVerifyInp) (res *input_terminal.CodeVerifyModel, err error) {
	var (
		VerifyType       string
		MchInfo          *entity.ThMch
		StoreInfo        *entity.ThMchStore
		CouponMchInfo    *entity.ThCouponMch
		FoodOrderInfo    *input_terminal.VerifyFoodOrderInfo
		MemberCouponInfo *input_terminal.VerifyMemberCouponInfo
	)

	MemberInfo := contexts.GetTerminalUser(ctx)

	ScanCode := strings.Split(in.Code, "|")

	if len(ScanCode) == 2 && ScanCode[1] == "FOOD" {
		// |FOOD区分
		if MemberInfo.Type == "STORE" {
			err = gerror.New("门店不能核销餐厅订单")
			return
		}
		VerifyType = "FOOD_ORDER"
	} else {
		if MemberInfo.Type == "RESTAURANT" {
			err = gerror.New("餐厅不能核销礼品券")
			return
		}
		VerifyType = "TH_COUPON"
	}

	res = new(input_terminal.CodeVerifyModel)
	res.VerifyType = VerifyType
	if VerifyType == "FOOD_ORDER" {
		if err = dao.FoodOrder.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.FoodOrder.Columns().OrderSn, ScanCode[0]).Scan(&FoodOrderInfo); err != nil {
			err = gerror.Wrap(err, "获取餐厅订单信息失败，请稍后重试！")
			return
		}
		if FoodOrderInfo.RestaurantId != int(MemberInfo.RestaurantId) {
			err = gerror.New("该餐厅订单不属于您的餐厅，不可核销")
			return
		}

		if FoodOrderInfo.OrderStatus != "HAVE_PAID" || FoodOrderInfo.BookingStatus != "CONFIRMED" {
			err = gerror.New("订单状态不正确")
			return
		}

		if gtime.Now().After(FoodOrderInfo.BookDatetime) {
			err = gerror.New("您已逾期，无法核销")
			return
		}

		if FoodOrderInfo.VerifyStatus != "WAIT_VERIFY" {
			err = gerror.New("订单核销状态不正确")
			return
		}

		if err = s.FoodOrderVerify(ctx, &input_terminal.FoodOrderVerifyInp{
			OrderId:    FoodOrderInfo.Id,
			TerminalId: int(MemberInfo.TerminalId),
		}); err != nil {
			return
		}

		res.FoodOrderInfo = new(input_terminal.FoodOrderInfo)
		res.FoodOrderInfo.OrderSn = FoodOrderInfo.OrderSn
		res.FoodOrderInfo.BookingName = FoodOrderInfo.BookingName
		res.FoodOrderInfo.PhoneArea = FoodOrderInfo.PhoneArea
		res.FoodOrderInfo.BookingMobile = FoodOrderInfo.BookingMobile
		res.FoodOrderInfo.GoodsNum = FoodOrderInfo.GoodsNum
		res.FoodOrderInfo.BookDate = FoodOrderInfo.BookDate
		res.FoodOrderInfo.BookTime = FoodOrderInfo.BookTime
		res.FoodOrderInfo.RestaurantName = FoodOrderInfo.FoodOrderRestaurantInfo.Name
		res.FoodOrderInfo.RestaurantAddress = FoodOrderInfo.FoodOrderRestaurantInfo.DetailAddress
		res.FoodOrderInfo.GoodsName = FoodOrderInfo.FoodOrderGoodsInfo.GoodsName
		res.FoodOrderInfo.MemberMessage = FoodOrderInfo.MemberMessage
		res.FoodOrderInfo.MemberMessageJa = FoodOrderInfo.MemberMessageJa
		res.FoodOrderInfo.OrderAmount = FoodOrderInfo.OrderAmount
		res.FoodOrderInfo.CreatedTime = FoodOrderInfo.CreatedAt
	} else {
		if err = dao.ThMch.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(MemberInfo.MchId).Scan(&MchInfo); err != nil {
			err = gerror.Wrap(err, "获取商户信息失败，请稍后重试！")
			return
		}
		if err = dao.ThMchStore.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(MemberInfo.StoreId).Scan(&StoreInfo); err != nil {
			err = gerror.Wrap(err, "获取门店信息失败，请稍后重试！")
			return
		}
		if err = dao.ThMemberCoupon.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.ThMemberCoupon.Columns().CouponNo, in.Code).Scan(&MemberCouponInfo); err != nil {
			err = gerror.Wrap(err, "获取券信息失败，请稍后重试！")
			return
		}
		if MemberCouponInfo.CouponInfo.UseStatus != 1 {
			err = gerror.New("该礼品券已停止使用！")
			return
		}
		if MemberCouponInfo.State == 1 {
			err = gerror.New("该礼品券还未生效")
			return
		} else if MemberCouponInfo.State == 3 {
			err = gerror.New("该礼品券已核销")
			return
		} else if MemberCouponInfo.State == 4 {
			err = gerror.New("该礼品券已过期")
			return
		} else if MemberCouponInfo.State == 5 {
			err = gerror.New("该礼品券已失效")
			return
		}
		if err = dao.ThCouponMch.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Where(dao.ThCouponMch.Columns().CouponId, MemberCouponInfo.CouponId).Where(dao.ThCouponMch.Columns().MchId, MemberInfo.MchId).Scan(&CouponMchInfo); err != nil {
			err = gerror.Wrap(err, "获取券商户关联信息失败，请稍后重试！")
			return
		}
		if g.IsEmpty(CouponMchInfo) {
			err = gerror.New("您的商户不能核销该礼品券")
			return
		}

		if err = s.MemberCouponVerify(ctx, &input_terminal.MemberCouponVerifyInp{
			MemberCouponId: MemberCouponInfo.Id,
			TerminalId:     int(MemberInfo.TerminalId),
			CouponMchName:  CouponMchInfo.Name,
		}); err != nil {
			return
		}

		res.CouponInfo = new(input_terminal.CodeCouponInfo)
		res.CouponInfo.CouponName = MemberCouponInfo.CouponInfo.CouponName
		res.CouponInfo.CouponMchName = CouponMchInfo.Name
		res.CouponInfo.StartTime = MemberCouponInfo.StartTime
		res.CouponInfo.EndTime = MemberCouponInfo.EndTime
		res.CouponInfo.MemberNo = MemberCouponInfo.MemberInfo.MemberNo
		res.CouponInfo.MchName = MchInfo.Name
		res.CouponInfo.StoreName = StoreInfo.StoreName
		res.CouponInfo.StoreAddress = StoreInfo.DetailAddress
	}

	res.VerifyTime = gtime.Now()

	return
}

func (s *sTerminalTerminal) FoodOrderVerify(ctx context.Context, in *input_terminal.FoodOrderVerifyInp) (err error) {

	var foodOrderInfo *input_food.FoodOrderVerifyViewModel

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = dao.FoodOrder.Ctx(ctx).WithAll().WherePri(in.OrderId).Scan(&foodOrderInfo); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}

		// 记录分佣 并 核销
		cost := foodOrderInfo.RestaurantDetail.SettlementDetail.Cost
		settlementAmount := foodOrderInfo.OrderAmount
		if !g.IsEmpty(cost) {
			costArr := strings.Split(cost, ",")
			constIsCoupon := false
			constIsBal := false
			for _, costItem := range costArr {
				if gvar.New(costItem).Int() == 1 {
					constIsCoupon = true
				}
				if gvar.New(costItem).Int() == 2 {
					constIsBal = true
				}
			}
			if !constIsCoupon {
				settlementAmount = settlementAmount - foodOrderInfo.CouponAmount
			}
			if !constIsBal {
				settlementAmount = settlementAmount - foodOrderInfo.BalAmount
			}
		}

		settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(foodOrderInfo.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()

		if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).
			WherePri(foodOrderInfo.Id).Update(g.MapStrAny{
			dao.FoodOrder.Columns().SettlementAmount: settlementAmount,
			dao.FoodOrder.Columns().VerifyStatus:     "VERIFIED",
			dao.FoodOrder.Columns().VerifyTime:       gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "核销失败，请稍后重试！")
			return
		}

		// 写入终端核销日志
		if _, err = dao.SysTerminalVerify.Ctx(ctx).Insert(&entity.SysTerminalVerify{
			TerminalId:     in.TerminalId,
			VerifyType:     "FOOD_ORDER",
			MchId:          0,
			RestaurantId:   foodOrderInfo.RestaurantId,
			FoodOrderId:    gvar.New(foodOrderInfo.Id).Int(),
			VerifyMemberId: gvar.New(foodOrderInfo.MemberId).Int(),
			VerifyTime:     gtime.Now(),
		}); err != nil {
			return
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     gvar.New(foodOrderInfo.Id).Int(),
			ActionWay:   "VERIFIED",
			Remark:      "订单已核销",
			OperateType: "SYSTEM",
			//OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 转发到返利队列 订单计算佣金/计算经验需要扔队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(foodOrderInfo.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(foodOrderInfo.OrderSn).Bytes(),
			Header:       nil,
		})

		// 发送分销订单变更队列
		if foodOrderInfo.IsFx == "Y" {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
				DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
					OrderNo:      foodOrderInfo.OrderSn,
					ChangeStatus: "COMPLETE",
				}).MustToJson(),
				Header: nil,
			})
		}

		return
	})

	return
}

func (s *sTerminalTerminal) MemberCouponVerify(ctx context.Context, in *input_terminal.MemberCouponVerifyInp) (err error) {

	var thMemberCouponInfo *input_th.ThMemberCouponViewModel

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		MemberInfo := contexts.GetTerminalUser(ctx)

		if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).WithAll().
			WherePri(in.MemberCouponId).
			Scan(&thMemberCouponInfo); err != nil {
			err = gerror.New("礼品券信息不存在！")
			return
		}

		// 如果是员工活动领取的券，限制同个活动同个券一天只能核销一张
		if thMemberCouponInfo.Source == 3 && thMemberCouponInfo.ActivityId > 0 {
			var ActivityCoupon *entity.EmployeeActivityCoupon
			if err = dao.EmployeeActivityCoupon.Ctx(ctx).
				Where(dao.EmployeeActivityCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.EmployeeActivityCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Scan(&ActivityCoupon); err != nil {
				return
			}

			var todayVerifiedCount int
			todayStart := gtime.Now().Format("Y-m-d") + " 00:00:00"
			todayEnd := gtime.Now().Format("Y-m-d") + " 23:59:59"

			if todayVerifiedCount, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
				Where(dao.ThMemberCoupon.Columns().Source, 3).
				Where(dao.ThMemberCoupon.Columns().MemberId, thMemberCouponInfo.MemberId).
				Where(dao.ThMemberCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.ThMemberCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Where(dao.ThMemberCoupon.Columns().State, 3).
				WhereBetween(dao.ThMemberCoupon.Columns().VerifyTime, todayStart, todayEnd).
				Count(); err != nil {
				err = gerror.Wrap(err, "查询今日核销记录失败！")
				return
			}

			if todayVerifiedCount >= ActivityCoupon.PerDayVerify {
				err = gerror.New("今日不可再进行核销！")
				return
			}
		}

		// 核销礼品券
		if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			WherePri(in.MemberCouponId).Data(input_th.ThMemberCouponVerifyFields{
			State:         3,
			VerifyTime:    gtime.Now(),
			VerifyMchId:   int(MemberInfo.MchId),
			VerifyStoreId: int(MemberInfo.StoreId),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 更新门店核销数
		if _, err = dao.ThMchStore.Ctx(ctx).Data(g.Map{
			dao.ThMchStore.Columns().VerifyNum: gdb.Raw(fmt.Sprintf("verify_num+%d", 1)),
		}).WherePri(MemberInfo.StoreId).Update(); err != nil {
			err = gerror.Wrap(err, "修改门店核销数，请稍后重试！")
			return
		}

		// 写入终端核销日志
		if _, err = dao.SysTerminalVerify.Ctx(ctx).OmitEmptyData().Insert(&entity.SysTerminalVerify{
			TerminalId:     in.TerminalId,
			VerifyType:     "TH_COUPON",
			MchId:          gvar.New(MemberInfo.MchId).Int(),
			StoreId:        int(MemberInfo.StoreId),
			MemberCouponId: thMemberCouponInfo.Id,
			CouponMchName:  in.CouponMchName,
			VerifyMemberId: gvar.New(thMemberCouponInfo.MemberId).Int(),
			VerifyTime:     gtime.Now(),
		}); err != nil {
			return err
		}

		return
	})

	return
}
