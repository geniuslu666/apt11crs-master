package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type SettlementAccountListReq struct {
	g.Meta `path:"/foodSettlementAccount/list" method:"get" tags:"ADMIN_FOOD" summary:"获取结算账户列表"`
	input_food.FoodSettlementAccountListInp
}

type SettlementAccountListRes struct {
	input_form.PageRes
	List []*input_food.FoodSettlementAccountListModel `json:"list"   dc:"数据列表"`
}

type SettlementAccountViewReq struct {
	g.Meta `path:"/foodSettlementAccount/view" method:"get" tags:"ADMIN_FOOD" summary:"获取结算账户指定信息"`
	input_food.FoodSettlementAccountViewInp
}

type SettlementAccountViewRes struct {
	*input_food.FoodSettlementAccountViewModel
}

type SettlementAccountEditReq struct {
	g.Meta `path:"/foodSettlementAccount/edit" method:"post" tags:"ADMIN_FOOD" summary:"修改/新增结算账户"`
	input_food.FoodSettlementAccountEditInp
}

type SettlementAccountEditRes struct{}

type SettlementAccountDeleteReq struct {
	g.Meta `path:"/foodSettlementAccount/delete" method:"post" tags:"ADMIN_FOOD" summary:"删除结算账户"`
	input_food.FoodSettlementAccountDeleteInp
}

type SettlementAccountDeleteRes struct{}

type SettlementAccountStatusReq struct {
	g.Meta `path:"/foodSettlementAccount/status" method:"post" tags:"ADMIN_FOOD" summary:"更新结算账户状态"`
	input_food.FoodSettlementAccountStatusInp
}

type SettlementAccountStatusRes struct{}
