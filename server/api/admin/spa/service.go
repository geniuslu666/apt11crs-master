package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// ServiceListReq 查询服务管理列表
type ServiceListReq struct {
	g.Meta `path:"/spaService/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩服务列表"`
	input_spa.SpaServiceListInp
}

type ServiceListRes struct {
	input_form.PageRes
	List []*input_spa.SpaServiceListModel `json:"list"   dc:"数据列表"`
}

type ServiceAllListReq struct {
	g.Meta `path:"/spaService/all" method:"get" tags:"ADMIN_SPA" summary:"获取按摩服务列表-ALL"`
	input_spa.SpaServiceListInp
}

type ServiceAllListRes struct {
	List []*input_spa.SpaServiceAllListModel `json:"list"   dc:"所有按摩服务-数据列表"`
}

// ServiceViewReq 获取按摩服务详情
type ServiceViewReq struct {
	g.Meta `path:"/spaService/view" method:"get" tags:"ADMIN_SPA" summary:"获取按摩服务详情"`
	input_spa.SpaServiceViewInp
}

type ServiceViewRes struct {
	*input_spa.SpaServiceViewModel
}

// ServiceEditReq 修改/新增服务管理
type ServiceEditReq struct {
	g.Meta `path:"/spaService/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增按摩服务"`
	input_spa.SpaServiceEditInp
}

type ServiceEditRes struct{}

// ServiceDeleteReq 删除服务管理
type ServiceDeleteReq struct {
	g.Meta `path:"/spaService/delete" method:"post" tags:"ADMIN_SPA" summary:"删除按摩服务"`
	input_spa.SpaServiceDeleteInp
}

type ServiceDeleteRes struct{}

// ServiceMaxSortReq 获取服务最大排序
type ServiceMaxSortReq struct {
	g.Meta `path:"/spaService/maxSort" method:"get" tags:"ADMIN_SPA" summary:"获取按摩服务最大排序"`
	input_spa.SpaServiceMaxSortInp
}

type ServiceMaxSortRes struct {
	*input_spa.SpaServiceMaxSortModel
}

// ServiceStatusReq 更新订餐-座位表状态
type ServiceStatusReq struct {
	g.Meta `path:"/spaService/status" method:"post" tags:"ADMIN_SPA" summary:"更新按摩服务状态"`
	input_spa.SpaServiceStatusInp
}

type ServiceStatusRes struct{}

// ServiceRecycleReq 按摩服务_恢复
type ServiceRecycleReq struct {
	g.Meta `path:"/spaService/recycle" method:"post" tags:"ADMIN_SPA" summary:"按摩服务_恢复"`
	input_spa.SpaServiceDeleteInp
}

type ServiceRecycleRes struct{}
