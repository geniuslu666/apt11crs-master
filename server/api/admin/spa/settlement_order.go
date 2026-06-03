package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"github.com/gogf/gf/v2/frame/g"
)

// SettlementOrderListReq 结算单列表
type SettlementOrderListReq struct {
	g.Meta `path:"/spaSettlementOrder/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩结算单列表"`
	input_spa.SpaSettlementOrderListInp
}

type SettlementOrderListRes struct {
	input_form.PageRes
	List []*input_spa.SpaSettlementOrderListModel `json:"list"   dc:"数据列表"`
}

type SettlementOrderStatReq struct {
	g.Meta `path:"/spaSettlementOrder/stat" method:"get" tags:"ADMIN_SPA" summary:"获取按摩结算单概况"`
	input_spa.SpaSettlementOrderStatInp
}

type SettlementOrderStatRes struct {
	*input_spa.SpaSettlementOrderStatModel
}

type SettlementOrderViewReq struct {
	g.Meta `path:"/spaSettlementOrder/view" method:"get" tags:"ADMIN_SPA" summary:"获取按摩结算单详情"`
	input_spa.SpaSettlementOrderViewInp
}

type SettlementOrderViewRes struct {
	*input_spa.SpaSettlementOrderViewModel
}

type SettlementOrderVerifyReq struct {
	g.Meta `path:"/spaSettlementOrder/verify" method:"post" tags:"ADMIN_SPA" summary:"按摩结算单-核账"`
	input_spa.SpaSettlementOrderVerifyInp
}

type SettlementOrderVerifyRes struct {
}
