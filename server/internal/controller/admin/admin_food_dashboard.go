package admin

import (
	"APT/internal/dao"
	hook2 "APT/internal/library/hgorm/hook"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/admin/food"
)

func (c *ControllerFood) Dashboard(ctx context.Context, req *food.DashboardReq) (res *food.DashboardRes, err error) {
	res = new(food.DashboardRes)
	if req.Type == "base" {
		//餐厅总数
		if res.Details.RestaurantTotalNum, err = dao.FoodRestaurant.Ctx(ctx).Count(); err != nil {
			return
		}
		// 营业中餐厅数量
		if res.Details.RestaurantTotalOnNum, err = dao.FoodRestaurant.Ctx(ctx).
			Where(dao.FoodRestaurant.Columns().OpenStatus, "OPENING").
			Count(); err != nil {
			return
		}
		// 累计餐厅预约量
		var (
			RestaurantTotalOrderNum float64
		)
		if RestaurantTotalOrderNum, err = dao.FoodRestaurant.Ctx(ctx).Sum(dao.FoodRestaurant.Columns().TotalOrderNum); err != nil {
			return
		}
		res.Details.RestaurantTotalOrderNum = int(RestaurantTotalOrderNum)
		// 累计餐厅预约金额
		if res.Details.RestaurantTotalOrderAmount, err = dao.FoodRestaurant.Ctx(ctx).Sum(dao.FoodRestaurant.Columns().TotalOrderAmount); err != nil {
			return
		}
		// 待结算订单量
		if res.Details.WaitSettlementOrderNum, err = dao.FoodOrder.Ctx(ctx).
			Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
			Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
			Count(); err != nil {
			return
		}
		// 待结算金额
		if res.Details.WaitSettlementOrderAmount, err = dao.FoodOrder.Ctx(ctx).
			Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED").
			Where(dao.FoodOrder.Columns().SettlementStatus, "WAIT").
			Sum(dao.FoodOrder.Columns().SettlementAmount); err != nil {
			return
		}
		// 待核账订单量
		var (
			SettlementOrderIds []gdb.Value
		)
		if SettlementOrderIds, err = dao.FoodSettlementOrder.Ctx(ctx).
			Where(dao.FoodSettlementOrder.Columns().VerifyStatus, "WAIT_VERIFY").
			Array(dao.FoodSettlementOrder.Columns().Id); err != nil {
			return
		}
		if res.Details.WaitVerifyOrderNum, err = dao.FoodOrder.Ctx(ctx).
			WhereIn(dao.FoodOrder.Columns().SettlementOrderId, SettlementOrderIds).
			Count(); err != nil {
			return
		}
		// 待核账金额
		if res.Details.WaitVerifyOrderAmount, err = dao.FoodSettlementOrder.Ctx(ctx).
			Where(dao.FoodSettlementOrder.Columns().VerifyStatus, "WAIT_VERIFY").
			Sum(dao.FoodSettlementOrder.Columns().SettlementAmount); err != nil {
			return
		}
	} else if req.Type == "restaurantLine" {
		if req.RestaurantLineType == "num" {
			if err = dao.FoodRestaurant.Ctx(ctx).WhereIn(dao.FoodRestaurant.Columns().OpenStatus, g.Slice{"OPENING", "REST"}).Page(1, 5).OrderDesc(dao.FoodRestaurant.Columns().TotalOrderNum).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.RestaurantLine); err != nil {
				return
			}
		} else {
			if err = dao.FoodRestaurant.Ctx(ctx).WhereIn(dao.FoodRestaurant.Columns().OpenStatus, g.Slice{"OPENING", "REST"}).Page(1, 5).OrderDesc(dao.FoodRestaurant.Columns().TotalOrderAmount).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.RestaurantLine); err != nil {
				return
			}
		}
	} else if req.Type == "goodsLine" {
		if req.GoodsLineType == "num" {
			if err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().GoodsState, 1).Page(1, 5).WithAll().OrderDesc(dao.FoodGoods.Columns().TotalOrderNum).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.GoodsLine); err != nil {
				return
			}
		} else {
			if err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().GoodsState, 1).Page(1, 5).WithAll().OrderDesc(dao.FoodGoods.Columns().TotalOrderAmount).Hook(hook2.PmsFindLanguageValueHook).Scan(&res.GoodsLine); err != nil {
				return
			}
		}
	}
	return
}
