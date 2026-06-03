package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// SettlementListReq 查询结算模式列表
type SettlementListReq struct {
	g.Meta `path:"/spaSettlement/list" method:"get" tags:"ADMIN_SPA" summary:"获取结算模式列表"`
	input_spa.SpaSettlementListInp
}

type SettlementListRes struct {
	input_form.PageRes
	List []*input_spa.SpaSettlementListModel `json:"list"   dc:"数据列表"`
}

// SettlementViewReq 获取结算模式指定信息
type SettlementViewReq struct {
	g.Meta `path:"/spaSettlement/view" method:"get" tags:"ADMIN_SPA" summary:"获取结算模式指定信息"`
	input_spa.SpaSettlementViewInp
}

type SettlementViewRes struct {
	*input_spa.SpaSettlementViewModel
}

// SettlementEditReq 修改/新增结算模式
type SettlementEditReq struct {
	g.Meta `path:"/spaSettlement/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增结算模式"`
	input_spa.SpaSettlementEditInp
}

type SettlementEditRes struct{}

// SettlementDeleteReq 删除结算模式
type SettlementDeleteReq struct {
	g.Meta `path:"/spaSettlement/delete" method:"post" tags:"ADMIN_SPA" summary:"删除结算模式"`
	input_spa.SpaSettlementDeleteInp
}

type SettlementDeleteRes struct{}

// SettlementStatusReq 更新结算模式状态
type SettlementStatusReq struct {
	g.Meta `path:"/spaSettlement/status" method:"post" tags:"ADMIN_SPA" summary:"更新结算模式状态"`
	input_spa.SpaSettlementStatusInp
}

type SettlementStatusRes struct{}
