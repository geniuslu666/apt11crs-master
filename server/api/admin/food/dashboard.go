package food

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type DashboardReq struct {
	g.Meta             `path:"/food/dashboard" method:"post" tags:"ADMIN_FOOD" summary:"APP数据看板_展示"`
	Type               string `json:"type"    dc:"类型(base:实时概况  restaurantLine:餐厅排行  goodsLine:套餐排行)"`
	RestaurantLineType string `json:"restaurant_line_type" dc:"餐厅排行类型  num:预定量  amount:预定金额"`
	GoodsLineType      string `json:"goods_line_type" dc:"套餐排行类型  num:预定量  amount:预定金额"`
}

type DashboardRes struct {
	Details struct {
		RestaurantTotalNum         int     `json:"restaurantTotalNum" dc:"餐厅总数"`
		RestaurantTotalOnNum       int     `json:"restaurantTotalOnNum" dc:"营业中餐厅数量"`
		RestaurantTotalOrderNum    int     `json:"restaurantTotalOrderNum" dc:"累计餐厅预约量"`
		RestaurantTotalOrderAmount float64 `json:"restaurantTotalOrderAmount" dc:"累计餐厅预约金额"`
		WaitSettlementOrderNum     int     `json:"waitSettlementOrderNum" dc:"待结算订单量"`
		WaitSettlementOrderAmount  float64 `json:"waitSettlementOrderAmount" dc:"待结算金额"`
		WaitVerifyOrderNum         int     `json:"waitVerifyOrderNum" dc:"待核账订单量"`
		WaitVerifyOrderAmount      float64 `json:"waitVerifyOrderAmount" dc:"待核账金额"`
	} `json:"details" dc:"实时概况"`
	RestaurantLine []*entity.FoodRestaurant `json:"restaurantLine" dc:"实时概况"`
	GoodsLine      []*FoodGoodsModel        `json:"goodsLine" dc:"实时概况"`
}

type FoodGoodsModel struct {
	entity.FoodGoods
	RestaurantDetail *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		Id         int    `json:"id"      description:"餐厅ID"`
		Name       string `json:"name"    description:"餐厅名称"`
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
}
