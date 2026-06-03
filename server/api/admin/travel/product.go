package travel

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductListReq 获取一日游产品列表
type ProductListReq struct {
	g.Meta `path:"/travel/product/list" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游产品列表"`
	input_travel.TravelProductListInp
}

type ProductListRes struct {
	input_form.PageRes
	List []*input_travel.TravelProductListModel `json:"list" dc:"数据列表"`
}

// ProductViewReq 获取一日游产品详情
type ProductViewReq struct {
	g.Meta `path:"/travel/product/view" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游产品详情"`
	input_travel.TravelProductViewInp
}

type ProductViewRes struct {
	*input_travel.TravelProductViewModel
}

// ProductEditReq 新增/编辑一日游产品
type ProductEditReq struct {
	g.Meta `path:"/travel/product/edit" method:"post" tags:"ADMIN_TRAVEL" summary:"新增/编辑一日游产品"`
	input_travel.TravelProductEditInp
}

type ProductEditRes struct {
	Id int64 `json:"id" dc:"产品ID"`
}

// ProductDeleteReq 删除一日游产品（软删除至回收站）
type ProductDeleteReq struct {
	g.Meta `path:"/travel/product/delete" method:"post" tags:"ADMIN_TRAVEL" summary:"删除一日游产品"`
	input_travel.TravelProductDeleteInp
}

type ProductDeleteRes struct{}

// ProductStatusReq 启用/禁用一日游产品
type ProductStatusReq struct {
	g.Meta `path:"/travel/product/status" method:"post" tags:"ADMIN_TRAVEL" summary:"启用/禁用一日游产品"`
	input_travel.TravelProductStatusInp
}

type ProductStatusRes struct{}

// ProductRecycleListReq 回收站列表
type ProductRecycleListReq struct {
	g.Meta `path:"/travel/product/recycle" method:"get" tags:"ADMIN_TRAVEL" summary:"一日游产品回收站列表"`
	input_travel.TravelProductListInp
}

type ProductRecycleListRes struct {
	input_form.PageRes
	List []*input_travel.TravelProductListModel `json:"list" dc:"数据列表"`
}

// ProductRestoreReq 从回收站恢复产品
type ProductRestoreReq struct {
	g.Meta `path:"/travel/product/restore" method:"post" tags:"ADMIN_TRAVEL" summary:"从回收站恢复一日游产品"`
	input_travel.TravelProductDeleteInp
}

type ProductRestoreRes struct{}
