package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// SettlementOrderListReq 结算单列表
type SettlementOrderListReq struct {
	g.Meta `path:"/carSettlementOrder/list" method:"get" tags:"ADMIN_CAR" summary:"获取结算订单列表"`
	input_car.CarSettlementOrderListInp
}

type SettlementOrderListRes struct {
	input_form.PageRes
	List []*input_car.CarSettlementOrderListModel `json:"list"   dc:"数据列表"`
}

// SettlementOrderStatReq 结算单概况
type SettlementOrderStatReq struct {
	g.Meta `path:"/carSettlementOrder/stat" method:"get" tags:"ADMIN_CAR" summary:"获取结算单概况"`
	input_car.CarSettlementOrderStatInp
}

type SettlementOrderStatRes struct {
	*input_car.CarSettlementOrderStatModel
}

// SettlementOrderViewReq 结算单详情
type SettlementOrderViewReq struct {
	g.Meta `path:"/carSettlementOrder/view" method:"get" tags:"ADMIN_CAR" summary:"获取结算单详情"`
	input_car.CarSettlementOrderViewInp
}

type SettlementOrderViewRes struct {
	*input_car.CarSettlementOrderViewModel
}

// SettlementOrderVerifyReq 结算单-核账
type SettlementOrderVerifyReq struct {
	g.Meta `path:"/carSettlementOrder/verify" method:"post" tags:"ADMIN_CAR" summary:"结算单-核账"`
	input_car.CarSettlementOrderVerifyInp
}

type SettlementOrderVerifyRes struct {
}
