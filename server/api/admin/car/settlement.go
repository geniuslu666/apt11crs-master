package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// SettlementListReq 查询结算模式列表
type SettlementListReq struct {
	g.Meta `path:"/carSettlement/list" method:"get" tags:"ADMIN_CAR" summary:"获取结算模式列表"`
	input_car.CarSettlementListInp
}

type SettlementListRes struct {
	input_form.PageRes
	List []*input_car.CarSettlementListModel `json:"list"   dc:"数据列表"`
}

// SettlementViewReq 获取结算模式指定信息
type SettlementViewReq struct {
	g.Meta `path:"/carSettlement/view" method:"get" tags:"ADMIN_CAR" summary:"获取结算模式指定信息"`
	input_car.CarSettlementViewInp
}

type SettlementViewRes struct {
	*input_car.CarSettlementViewModel
}

// SettlementEditReq 修改/新增结算模式
type SettlementEditReq struct {
	g.Meta `path:"/carSettlement/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增结算模式"`
	input_car.CarSettlementEditInp
}

type SettlementEditRes struct{}

// SettlementDeleteReq 删除结算模式
type SettlementDeleteReq struct {
	g.Meta `path:"/carSettlement/delete" method:"post" tags:"ADMIN_CAR" summary:"删除结算模式"`
	input_car.CarSettlementDeleteInp
}

type SettlementDeleteRes struct{}

// SettlementStatusReq 更新结算模式状态
type SettlementStatusReq struct {
	g.Meta `path:"/carSettlement/status" method:"post" tags:"ADMIN_CAR" summary:"更新结算模式状态"`
	input_car.CarSettlementStatusInp
}

type SettlementStatusRes struct{}
