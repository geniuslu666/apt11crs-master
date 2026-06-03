package travel

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

type ProductSkuListReq struct {
	g.Meta `path:"/travel/productSku/list" method:"get" tags:"ADMIN_TRAVEL" summary:"获取车型列表"`
	input_travel.TravelProductSkuListInp
}

type ProductSkuListRes struct {
	input_form.PageRes
	List []*input_travel.TravelProductSkuListModel `json:"list" dc:"数据列表"`
}

type ProductSkuViewReq struct {
	g.Meta `path:"/travel/productSku/view" method:"get" tags:"ADMIN_TRAVEL" summary:"获取车型指定信息"`
	input_travel.TravelProductSkuViewInp
}

type ProductSkuViewRes struct {
	*input_travel.TravelProductSkuViewModel
}

type ProductSkuEditReq struct {
	g.Meta `path:"/travel/productSku/edit" method:"post" tags:"ADMIN_TRAVEL" summary:"修改/新增车型"`
	input_travel.TravelProductSkuEditInp
}

type ProductSkuEditRes struct{}

type ProductSkuDeleteReq struct {
	g.Meta `path:"/travel/productSku/delete" method:"post" tags:"ADMIN_TRAVEL" summary:"删除车型"`
	input_travel.TravelProductSkuDeleteInp
}

type ProductSkuDeleteRes struct{}

type ProductSkuStatusReq struct {
	g.Meta `path:"/travel/productSku/status" method:"post" tags:"ADMIN_TRAVEL" summary:"更新车型状态"`
	input_travel.TravelProductSkuStatusInp
}

type ProductSkuStatusRes struct{}
