package cabinet

import (
	"APT/internal/library/cabinetApi"

	"github.com/gogf/gf/v2/frame/g"
)

type CabinetInfoReq struct {
	g.Meta `path:"/cabinetNotify/cabinetInfo" method:"post" tags:"CABINET_HOOK" summary:"查询储物柜"`
	*cabinetApi.CabinetInfoParams
}

type CabinetInfoRes struct {
	cabinetApi.CabinetInfoResponseItem
}

type CreateOrderReq struct {
	g.Meta `path:"/cabinetNotify/createOrder" method:"post" tags:"CABINET_HOOK" summary:"下单"`
	*cabinetApi.CabinetCreateOrderParams
}

type CreateOrderRes struct {
	cabinetApi.CabinetCreateOrderResponseItem
}

type OrderQueryReq struct {
	g.Meta `path:"/cabinetNotify/orderQuery" method:"post" tags:"CABINET_HOOK" summary:"查询订单"`
	*cabinetApi.CabinetOrderQueryParams
}

type OrderQueryRes struct {
	cabinetApi.CabinetOrderQueryResponseItem
}

type PayOvertimeReq struct {
	g.Meta `path:"/cabinetNotify/payOvertime" method:"post" tags:"CABINET_HOOK" summary:"支付超时费"`
	*cabinetApi.CabinetPayOvertimeParams
}

type PayOvertimeRes struct {
	cabinetApi.CabinetPayOvertimeResponseItem
}

type CabinetListReq struct {
	g.Meta `path:"/cabinetNotify/cabinetList" method:"post" tags:"CABINET_HOOK" summary:"储物柜列表"`
}

type CabinetListRes struct {
	List []*cabinetApi.CabinetListResponseItem `json:"list"   dc:"数据列表"`
}

type OrderCompleteReq struct {
	g.Meta `path:"/cabinetNotify/orderComplete" method:"post" tags:"CABINET_HOOK" summary:"订单完成"`
	*cabinetApi.OrderCompleteParams
}

type OrderCompleteRes struct {
}
