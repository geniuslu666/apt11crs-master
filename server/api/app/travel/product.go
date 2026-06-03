package travel

import (
	"APT/internal/model/input/input_travel"

	"github.com/gogf/gf/v2/frame/g"
)

type ProductListReq struct {
	g.Meta   `path:"/travel/productList" method:"post" tags:"APP_TRAVEL" summary:"[一日游]产品列表"`
	PageNum  int `p:"pageNum" v:"required#page_number_unknown" dc:"页码"`
	PageSize int `p:"pageSize" v:"required#page_number_unknown" dc:"页数"`
}

type ProductListRes struct {
	List  []*input_travel.TravelProductAppListModel `json:"list" dc:"数据列表"`
	Count int                                       `json:"count"   dc:"数据总数"`
}

type ProductViewReq struct {
	g.Meta `path:"/travel/productView" method:"post" tags:"APP_TRAVEL" summary:"[一日游]产品详情"`
	input_travel.TravelProductAppViewInp
}

type ProductViewRes struct {
	*input_travel.TravelProductAppViewModel
}

type ProductSkuStockReq struct {
	g.Meta `path:"/travel/productSkuStock" method:"post" tags:"APP_TRAVEL" summary:"[一日游]SKU库存查询"`
	input_travel.TravelProductSkuStockInp
}

type ProductSkuStockRes struct {
	List []*input_travel.TravelProductSkuStockModel `json:"list" dc:"SKU库存列表"`
}
