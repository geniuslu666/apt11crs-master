package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type PricePlanOneReq struct {
	g.Meta `path:"/pricePlan/view" method:"get" tags:"ADMIN_PMS" summary:"PMS价格计划_详情"`
	*input_hotel.FindOneHotelPricePlanInp
}
type PricePlanOneRes struct {
	*input_hotel.FindOneHotelPricePlanRes `response:"arr"`
}

type PricePlanReq struct {
	g.Meta `path:"/pricePlan/list" method:"get" tags:"ADMIN_PMS" summary:"PMS价格计划_列表"`
	*input_hotel.SearchHotelPricePlanInp
}

type PricePlanRes struct {
	input_form.PageRes
	List []*input_hotel.SearchHotelPricePlanRes `json:"list"   dc:"数据列表"`
}

type PricePlanEditReq struct {
	g.Meta `path:"/pricePlan/edit" method:"post" tags:"ADMIN_PMS" summary:"PMS价格计划_新增"`
	*input_hotel.EditHotelPricePlanInp
}

type PricePlanEditRes struct{}

type PricePlanDelReq struct {
	g.Meta `path:"/pricePlan/delete" method:"post" tags:"ADMIN_PMS" summary:"PMS价格计划_删除"`
	*input_hotel.DeleteHotelPricePlanInp
}

type PricePlanDelRes struct{}

type PricePlanStatusReq struct {
	g.Meta `path:"/pricePlan/status" method:"post" tags:"ADMIN_PMS" summary:"更新价格计划状态"`
	input_hotel.PricePlanStatusInp
}

type PricePlanStatusRes struct{}

type PricePlanRecycleReq struct {
	g.Meta `path:"/pricePlan/recycle" method:"post" tags:"ADMIN_PMS" summary:"PMS价格计划_恢复"`
	input_hotel.DeleteHotelPricePlanInp
}

type PricePlanRecycleRes struct{}
