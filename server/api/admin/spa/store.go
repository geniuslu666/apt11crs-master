package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// StoreListReq 查询门店列表
type StoreListReq struct {
	g.Meta `path:"/spaStore/list" method:"get" tags:"ADMIN_SPA" summary:"获取门店列表"`
	input_spa.SpaStoreListInp
}

type StoreListRes struct {
	input_form.PageRes
	List []*input_spa.SpaStoreListModel `json:"list"   dc:"数据列表"`
}

// StoreViewReq 获取门店指定信息
type StoreViewReq struct {
	g.Meta `path:"/spaStore/view" method:"get" tags:"ADMIN_SPA" summary:"获取门店管指定信息"`
	input_spa.SpaStoreViewInp
}

type StoreViewRes struct {
	*input_spa.SpaStoreViewModel
}

// StoreLatestReq 获取门店最新信息
type StoreLatestReq struct {
	g.Meta `path:"/spaStore/latest" method:"get" tags:"ADMIN_SPA" summary:"获取门店最新信息"`
}

type StoreLatestRes struct {
	*input_spa.SpaStoreViewModel
}

// StoreEditReq 修改/新增门店
type StoreEditReq struct {
	g.Meta `path:"/spaStore/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增门店"`
	input_spa.SpaStoreEditInp
}

type StoreEditRes struct{}

// StoreDeleteReq 删除门店
type StoreDeleteReq struct {
	g.Meta `path:"/spaStore/delete" method:"post" tags:"ADMIN_SPA" summary:"删除门店"`
	input_spa.SpaStoreDeleteInp
}

type StoreDeleteRes struct{}
