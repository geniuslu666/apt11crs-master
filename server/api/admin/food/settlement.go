package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type SettlementListReq struct {
	g.Meta `path:"/foodSettlement/list" method:"get" tags:"ADMIN_FOOD" summary:"获取结算模式列表"`
	input_food.FoodSettlementListInp
}

type SettlementListRes struct {
	input_form.PageRes
	List []*input_food.FoodSettlementListModel `json:"list"   dc:"数据列表"`
}

type SettlementViewReq struct {
	g.Meta `path:"/foodSettlement/view" method:"get" tags:"ADMIN_FOOD" summary:"获取结算模式指定信息"`
	input_food.FoodSettlementViewInp
}

type SettlementViewRes struct {
	*input_food.FoodSettlementViewModel
}

type SettlementEditReq struct {
	g.Meta `path:"/foodSettlement/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增结算模式"`
	input_food.FoodSettlementEditInp
}

type SettlementEditRes struct{}

type SettlementDeleteReq struct {
	g.Meta `path:"/foodSettlement/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除结算模式"`
	input_food.FoodSettlementDeleteInp
}

type SettlementDeleteRes struct{}

type SettlementStatusReq struct {
	g.Meta `path:"/foodSettlement/status" method:"post" tags:"ADMIN_FOOD" summary:"更新结算模式状态"`
	input_food.FoodSettlementStatusInp
}

type SettlementStatusRes struct{}
