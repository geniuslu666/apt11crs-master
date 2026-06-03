package logic_pay

import (
	"APT/internal/dao"
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
	"math"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// PayInfo 支付信息计算
func (s *sPayService) PayInfo(ctx context.Context, in *input_pay.PayInfoInp) (out *input_pay.PayInfoModel, err error) {
	var (
		MemberInfo  *entity.PmsMember
		MemberLevel *entity.PmsMemberLevel
		MemberScene *entity.PmsMemberScene
		CouponInfo  *entity.PmsCoupon
		YYConfig    *model.YYConfig
		AllScore    float64
	)
	if g.IsEmpty(in.Scene) {
		in.Scene = "HOTEL"
	}
	// 查询当前用户积分余额
	if MemberInfo, err = service.AppMember().MemberInfo(ctx, in.MemberId); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		err = gerror.New("用户信息不存在")
		return
	}

	// 查询会员等级
	if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberLevel) {
		err = gerror.New("用户等级信息不存在")
		return
	}
	// 查询场景使用积分比例
	if in.Scene == "HOTEL" {
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 1).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
	} else if in.Scene == "CABINET" {
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 5).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
	} else if in.Scene == "TRAVEL" {
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 6).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
	}
	if g.IsEmpty(MemberScene) {
		err = gerror.New("场景信息不存在")
		return
	}
	// 获取积分汇率
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}
	if YYConfig.ExchangeRate == 0 || YYConfig.ExchangeRate < 0 {
		YYConfig.ExchangeRate = 1
	}
	out = new(input_pay.PayInfoModel)
	out.MemberBalance = MemberInfo.Balance
	out.AllAmount = in.OrderAmount

	// 查询是否优惠券参与支付
	if !g.IsEmpty(in.CouponId) {
		// 优惠券参与支付
		if err = dao.PmsCoupon.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsCoupon.Columns().Id, in.CouponId).Scan(&CouponInfo); err != nil {
			return
		}

		if g.IsEmpty(CouponInfo) {
			err = gerror.New("优惠券不存在")
			return
		}

		if CouponInfo.State != 1 {
			err = gerror.New("优惠券无法使用")
			return
		}

		if !g.IsEmpty(CouponInfo.UseTime) {
			err = gerror.New("优惠券已使用")
			return
		}

		if CouponInfo.AtLeast > 0 && CouponInfo.AtLeast > in.OrderAmount {
			err = gerror.New(fmt.Sprintf("优惠券需满%vJPY,当前无法使用", CouponInfo.AtLeast))
			return
		}

		if (in.Scene == "HOTEL" && CouponInfo.Scene != 1) || (in.Scene == "FOOD" && CouponInfo.Scene != 2) || (in.Scene == "CABINET" && CouponInfo.Scene != 5) {
			err = gerror.New("优惠券使用场景错误，无法使用")
			return
		}

		if CouponInfo.Type == "reward" { // 满减优惠券

			in.OrderAmount = decimal.NewFromFloat(in.OrderAmount).Sub(decimal.NewFromFloat(CouponInfo.Money)).Round(0).InexactFloat64()
			out.CouponAmount = CouponInfo.Money
			out.CouponName = CouponInfo.CouponName
			out.CouponId = CouponInfo.Id
		} else if CouponInfo.Type == "discount" { // 折扣优惠券
			// 优惠率
			Discount := decimal.NewFromFloat(10 - CouponInfo.Discount).Mul(decimal.NewFromFloat(0.1)).Round(2).InexactFloat64()
			// 优惠金额
			DiscountAmount := decimal.NewFromFloat(in.OrderAmount).Mul(decimal.NewFromFloat(Discount)).Round(0).InexactFloat64()
			in.OrderAmount = decimal.NewFromFloat(in.OrderAmount).Sub(decimal.NewFromFloat(DiscountAmount)).Round(0).InexactFloat64()
			out.CouponAmount = DiscountAmount
			out.CouponName = CouponInfo.CouponName
			out.CouponId = CouponInfo.Id
		}
		if in.OrderAmount < 0 {
			in.OrderAmount = 0
		}
	}

	out.BalanceConfig = new(input_hotel.BalanceConfig)
	out.BalanceConfig.IsPayOpen = MemberScene.IsPayOpen
	// 计算该笔订单能够有多少积分进行抵扣
	// * 计算可以使用的积分总额   =  订单总金额 * 场景抵扣基础比例 * 积分汇率
	AllScore = decimal.NewFromFloat(in.OrderAmount).Mul(decimal.NewFromFloat(MemberScene.PayRate)).Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).Div(decimal.NewFromInt(100)).Round(0).InexactFloat64()
	// * 计算会员可用积分总额 = min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )
	AllScore = math.Min(AllScore, MemberInfo.Balance)
	// * 计算积分的现金价值  min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )  /  积分汇率
	AllScore = decimal.NewFromFloat(AllScore).Div(decimal.NewFromFloat(YYConfig.ExchangeRate)).Round(0).InexactFloat64()

	if MemberScene.IsPayOpen == "N" || in.OrderAmount < MemberScene.LimitMoney {
		AllScore = 0
	}

	out.Score = AllScore
	out.BalanceConfig.ScenePayRate = MemberScene.PayRate
	out.BalanceConfig.ExchangeRate = YYConfig.ExchangeRate
	out.BalanceConfig.Level = MemberLevel.Id
	out.BalanceConfig.LevelName = MemberLevel.LevelName
	if in.OrderAmount == 0 {
		out.PayModel = 1
		out.ThirdAmount = in.OrderAmount
	} else if !in.IsBalance || MemberScene.IsPayOpen == "N" || in.OrderAmount < MemberScene.LimitMoney || AllScore <= 0 {
		// 三方支付
		out.PayModel = 3
		out.ThirdAmount = in.OrderAmount
	} else if in.OrderAmount > AllScore {
		// 组合支付
		out.PayModel = 2
		out.BalanceAmount = AllScore
		out.ThirdAmount = decimal.NewFromFloat(in.OrderAmount).Sub(decimal.NewFromFloat(AllScore)).Round(0).InexactFloat64()
	} else {
		// 纯余额支付
		out.PayModel = 1
		out.BalanceAmount = in.OrderAmount
	}
	return
}

// BookingPrice 预定价格计算
func (s *sPayService) BookingPrice(ctx context.Context, in *input_pay.BookingPriceInp) (out *input_pay.BookingPriceModel, err error) {
	var (
		PmsRoomType    *entity.PmsRoomType
		Availabilities []*entity.PmsAvailabilities
		PricePlanList  []*input_pay.BookingPriceItem
		basePrice      float64
		PmsPirceConfig *model.PmsPriceConfig
	)
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	out = new(input_pay.BookingPriceModel)
	defer func() {
		g.Log().Debug(ctx, "BookingPriceModel", out)
	}()
	// 生成预订单详情
	if g.IsEmpty(in.RoomInfos.Adult) {
		in.RoomInfos.Adult = 1
	}
	// 校验库存
	//if err = service.HotelService().CheckHotelInventory(ctx, &input_hotel.CheckHotelInventoryInp{
	//	TUID:      in.RoomInfos.RoomId,
	//	StartDate: in.CheckinAt,
	//	EndDate:   in.CheckoutAt,
	//	Num:       in.RoomInfos.RoomNum,
	//	PUID:      in.Puid,
	//}); err != nil {
	//	return
	//}
	RoomsItems := &input_hotel.OrderRooms{
		RoomId:   in.RoomInfos.RoomId,
		RoomNum:  in.RoomInfos.RoomNum,
		RoomNoId: in.RoomInfos.RoomNoId,
		Adult:    in.RoomInfos.Adult,
		Child:    in.RoomInfos.Child,
		Infant:   in.RoomInfos.Infant,
	}

	// 查询房型信息
	PmsRoomType = nil
	if err = dao.PmsRoomType.Ctx(ctx).
		Hook(hook.PmsFindLanguageValueHook).
		Where(dao.PmsRoomType.Columns().Uid, in.RoomInfos.RoomId).
		Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsRoomType) {
		err = gerror.New("房型不存在")
		return
	}
	RoomsItems.Tid = PmsRoomType.Id
	RoomsItems.Puid = PmsRoomType.Puid
	RoomsItems.Cover = PmsRoomType.Cover
	RoomsItems.Name = PmsRoomType.Name

	Availabilities = nil
	if err = dao.PmsAvailabilities.Ctx(ctx).
		Where(dao.PmsAvailabilities.Columns().Tuid, in.RoomInfos.RoomId).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, in.CheckinAt).
		WhereLT(dao.PmsAvailabilities.Columns().Date, in.CheckoutAt).
		Scan(&Availabilities); err != nil {
		return
	}

	if g.IsEmpty(Availabilities) {
		err = gerror.New("没有可预订的房型")
		return
	}

	PeopleAmount := gconv.Float64(0)
	g.Log().Debugf(ctx, "Availabilities %s", gjson.New(Availabilities).String())
	// 创建费率计划
	for _, Availabilitie := range Availabilities {
		if !g.IsEmpty(Availabilitie.Price) {
			basePrice = decimal.NewFromFloat(Availabilitie.Price).Mul(PricePercent).Round(0).InexactFloat64()
		} else {
			basePrice = PmsRoomType.AdditionalGuestAmounts
		}

		RoomItemNum := 0
		RoomItemNum = in.RoomInfos.Child + in.RoomInfos.Adult
		g.Log().Debugf(ctx, "PmsRoomType.OccupantsForBaseRate %d : Child %d,Adult %d,Infant %d", PmsRoomType.OccupantsForBaseRate, in.RoomInfos.Child, in.RoomInfos.Adult, in.RoomInfos.Infant)
		// 当日应付金额
		ChargeItemAmount := gvar.New(0).Float64()
		if RoomItemNum > PmsRoomType.OccupantsForBaseRate {
			g.Log().Debug(ctx, "人头费")
			OccupantsAmount := decimal.NewFromInt(gvar.New(RoomItemNum - PmsRoomType.OccupantsForBaseRate).Int64())
			g.Log().Debug(ctx, "OccupantsAmount", OccupantsAmount)
			OccupantsAmountRes, ok := OccupantsAmount.Mul(decimal.NewFromFloat(PmsRoomType.AdditionalGuestAmounts)).Float64()
			g.Log().Debug(ctx, "OccupantsAmount", OccupantsAmountRes)
			if !ok {
				err = gerror.New("房型人数金额计算错误")
			}
			RoomsItems.Charges = append(RoomsItems.Charges, &input_hotel.Charges{
				UID:         uuid.GetUuid(),
				Date:        Availabilitie.Date,
				FeeType:     "booking_fee_people",
				InitAmount:  OccupantsAmountRes,
				Amount:      OccupantsAmountRes,
				Description: "超员费",
			})
			RoomsItems.BookingFee += OccupantsAmountRes
			PeopleAmount += OccupantsAmountRes
			ChargeItemAmount += OccupantsAmountRes
		}
		RoomsItems.Charges = append(RoomsItems.Charges, &input_hotel.Charges{
			UID:         uuid.GetUuid(),
			Date:        Availabilitie.Date,
			FeeType:     "booking_fee",
			InitAmount:  basePrice,
			Amount:      basePrice,
			Description: "每日费用",
		})
		RoomsItems.BookingFee += basePrice

		ChargeItemAmount += basePrice

		// 根据每日统计应付金额
		RoomsItems.ChargesCal = append(RoomsItems.ChargesCal, &input_hotel.ChargesCalItem{
			Date:   Availabilitie.Date,
			Amount: ChargeItemAmount,
		})

	}
	g.Log().Debug(ctx, "RoomsItems", RoomsItems)
	out.OrderRooms = RoomsItems
	out.TotalPrice = decimal.NewFromFloat(RoomsItems.BookingFee).Add(decimal.NewFromFloat(out.TotalPrice)).Round(0).InexactFloat64()
	MemberToken := contexts.GetMemberUser(ctx)
	MemberLevel := 0
	MemberGroup := 0
	if !g.IsEmpty(MemberToken) {
		MemberLevel = MemberToken.Level
		MemberGroup = MemberToken.GroupId
	}
	RoomsItems.OrderAmount = out.TotalPrice
	PlanTotalAmount := out.TotalPrice
	// 查询价格plan
	if in.IsPricePlan {
		PmsPricePlanWhere := g.Map{
			dao.PmsPricePlan.Columns().PropertyId:      in.Puid,
			dao.PmsPricePlan.Columns().RoomTypeId:      in.RoomInfos.RoomId,
			dao.PmsPricePlan.Columns().PricePlanStatus: "Y",
			dao.PmsPricePlan.Columns().IsOpenPriceMode: "Y",
			//dao.PmsPricePlan.Columns().Id:in.PricePlanId,
			//dao.PmsPricePlan.Columns().MemberLevelId:   MemberLevel,
			//dao.PmsPricePlan.Columns().MemberGroupId:   MemberGroup,
		}
		if in.PricePlanId != 0 {
			PmsPricePlanWhere[dao.PmsPricePlan.Columns().Id] = in.PricePlanId
		}
		if err = dao.PmsPricePlan.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(PmsPricePlanWhere).Scan(&PricePlanList); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		for _, planPrice := range PricePlanList {
			days := convert.Diff(in.CheckinAt, in.CheckoutAt, "days")
			if g.IsEmpty(MemberLevel) && !g.IsEmpty(planPrice.MemberLevelId) {
				continue
			}
			if g.IsEmpty(MemberGroup) && !g.IsEmpty(planPrice.MemberGroupId) {
				continue
			}
			if !g.IsEmpty(planPrice.MemberLevelId) && !strings.Contains(planPrice.MemberLevelId, gvar.New(MemberLevel).String()) {
				continue
			}
			if !g.IsEmpty(planPrice.MemberGroupId) && !strings.Contains(planPrice.MemberGroupId, gvar.New(MemberGroup).String()) {
				continue
			}

			if g.IsEmpty(planPrice.BookingDays) || planPrice.BookingDays <= days {
				out.PlanPrice = append(out.PlanPrice, planPrice)
			} else {
				continue
			}
		}

		for planIndex, planPrice := range out.PlanPrice {
			if planPrice.IsOpenPriceMode == "Y" {
				for index, Charges := range RoomsItems.Charges {
					if Charges.FeeType == "booking_fee" {
						g.Log().Debug(ctx, "ChargesInitPrice", RoomsItems.Charges[index].InitAmount)
						if planPrice.PriceMode == "+" && planPrice.PriceStandard == "AMOUNT" { // 贵
							// 需要增加的金额
							RoomsItems.Charges[index].Amount = decimal.NewFromFloat(gvar.New(RoomsItems.Charges[index].InitAmount).Float64()).Add(decimal.NewFromFloat(gvar.New(planPrice.PlanValue).Float64())).Round(0).InexactFloat64()
						} else if planPrice.PriceMode == "-" && planPrice.PriceStandard == "AMOUNT" {
							RoomsItems.Charges[index].Amount = decimal.NewFromFloat(gvar.New(RoomsItems.Charges[index].InitAmount).Float64()).Sub(decimal.NewFromFloat(gvar.New(planPrice.PlanValue).Float64())).Round(0).InexactFloat64()

							if RoomsItems.Charges[index].Amount < 0 {
								RoomsItems.Charges[index].Amount = 0
							}
						} else if planPrice.PriceMode == "+" && planPrice.PriceStandard == "PERCENT" {
							PriceStandard := decimal.NewFromFloat(gvar.New(planPrice.PlanValue).Float64()).Div(decimal.NewFromFloat(100))
							PriceStandard = decimal.NewFromInt(1).Add(PriceStandard)

							RoomsItems.Charges[index].Amount = decimal.NewFromFloat(gvar.New(RoomsItems.Charges[index].InitAmount).Float64()).Mul(PriceStandard).Round(0).InexactFloat64()
						} else if planPrice.PriceMode == "-" && planPrice.PriceStandard == "PERCENT" {
							PriceStandard := decimal.NewFromFloat(gvar.New(planPrice.PlanValue).Float64()).Div(decimal.NewFromFloat(100))
							PriceStandard = decimal.NewFromInt(1).Sub(PriceStandard)
							RoomsItems.Charges[index].Amount = decimal.NewFromFloat(gvar.New(RoomsItems.Charges[index].InitAmount).Float64()).Mul(PriceStandard).Round(0).InexactFloat64()

							if RoomsItems.Charges[index].Amount < 0 {
								RoomsItems.Charges[index].Amount = 0
							}
						}
						g.Log().Debug(ctx, "ChargesResultPrice", RoomsItems.Charges[index].Amount)
					}
					out.PlanPrice[planIndex].TotalAmount = decimal.NewFromFloat(out.PlanPrice[planIndex].TotalAmount).Add(decimal.NewFromFloat(RoomsItems.Charges[index].Amount)).Round(0).InexactFloat64()
					g.Log().Debug(ctx, "PlanPrice.TotalAmount", out.PlanPrice[planIndex].TotalAmount)
				}
			}
			if in.PricePlanId > 0 {
				out.PlanPriceTotalAmount = PlanTotalAmount
			} else {
				if out.PlanPriceTotalAmount > PlanTotalAmount || out.PlanPriceTotalAmount == 0 {
					out.PlanPriceTotalAmount = PlanTotalAmount
				}
			}
		}

		// 由于价格计划的原因，价格可能不为基础房费了，每日统计应付金额需要重新计算
		for newCCIndex, NewChargesCal := range RoomsItems.ChargesCal {
			RoomsItems.ChargesCal[newCCIndex].Amount = 0
			for _, NewCharges := range RoomsItems.Charges {
				if NewChargesCal.Date == NewCharges.Date {
					RoomsItems.ChargesCal[newCCIndex].Amount = decimal.NewFromFloat(RoomsItems.ChargesCal[newCCIndex].Amount).Add(decimal.NewFromFloat(NewCharges.Amount)).Round(0).InexactFloat64()
				}
			}
		}

		// 找到 PlanPrice 数组中 TotalAmount 最小的值
		if len(out.PlanPrice) > 0 {
			for i, planPrice := range out.PlanPrice {
				if i == 0 || planPrice.TotalAmount < out.MinTotalPrice {
					out.MinTotalPrice = planPrice.TotalAmount
				}
			}
		}
		RoomsItems.OrderAmount = 0
		for _, Charges := range RoomsItems.Charges {
			RoomsItems.OrderAmount = decimal.NewFromFloat(RoomsItems.OrderAmount).Add(decimal.NewFromFloat(Charges.Amount)).Round(0).InexactFloat64()
		}
		RoomsItems.ChangeAmount = RoomsItems.BookingFee - RoomsItems.OrderAmount
	}
	return
}
