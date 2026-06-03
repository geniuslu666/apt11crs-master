package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// SettlementOrderListReq 结算单列表
type SettlementOrderListReq struct {
	g.Meta `path:"/foodSettlementOrder/list" method:"get" tags:"ADMIN_FOOD" summary:"获取结算账户列表"`
	input_food.FoodSettlementOrderListInp
}

type SettlementOrderListRes struct {
	input_form.PageRes
	List []*input_food.FoodSettlementOrderListModel `json:"list"   dc:"数据列表"`
}

type SettlementOrderStatReq struct {
	g.Meta `path:"/foodSettlementOrder/stat" method:"get" tags:"ADMIN_FOOD" summary:"获取结算单概况"`
	input_food.FoodSettlementOrderStatInp
}

type SettlementOrderStatRes struct {
	*input_food.FoodSettlementOrderStatModel
}

type SettlementOrderViewReq struct {
	g.Meta `path:"/foodSettlementOrder/view" method:"get" tags:"ADMIN_FOOD" summary:"获取结算账户指定信息"`
	input_food.FoodSettlementOrderViewInp
}

type SettlementOrderViewRes struct {
	*input_food.FoodSettlementOrderViewModel
}

type SettlementOrderVerifyReq struct {
	g.Meta `path:"/foodSettlementOrder/verify" method:"post" tags:"ADMIN_FOOD" summary:"结算单核账"`
	input_food.FoodSettlementOrderVerifyInp
}

type SettlementOrderVerifyRes struct {
}
