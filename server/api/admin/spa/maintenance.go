package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// MaintenanceListReq 查询Maintenance列表
type MaintenanceListReq struct {
	g.Meta `path:"/spaMaintenance/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩维护模式列表"`
	input_spa.SpaMaintenanceListInp
}

type MaintenanceListRes struct {
	input_form.PageRes
	List []*input_spa.SpaMaintenanceListModel `json:"list"   dc:"数据列表"`
}

// MaintenanceViewReq 获取Maintenance指定信息
type MaintenanceViewReq struct {
	g.Meta `path:"/spaMaintenance/view" method:"get" tags:"ADMIN_SPA" summary:"获取按摩维护模式指定信息"`
	input_spa.SpaMaintenanceViewInp
}

type MaintenanceViewRes struct {
	*input_spa.SpaMaintenanceViewModel
}

// MaintenanceEditReq 修改/新增Maintenance
type MaintenanceEditReq struct {
	g.Meta `path:"/spaMaintenance/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增按摩维护模式"`
	input_spa.SpaMaintenanceEditInp
}

type MaintenanceEditRes struct{}

// MaintenanceDeleteReq 删除Maintenance
type MaintenanceDeleteReq struct {
	g.Meta `path:"/spaMaintenance/delete" method:"post" tags:"ADMIN_SPA" summary:"删除按摩维护模式"`
	input_spa.SpaMaintenanceDeleteInp
}

type MaintenanceDeleteRes struct{}

// MaintenanceLanguageListReq 获取Maintenance语言
type MaintenanceLanguageListReq struct {
	g.Meta `path:"/spaMaintenance/languageList" method:"get" tags:"ADMIN_SPA" summary:"获取按摩维护模式语言"`
	input_spa.SpaMaintenanceLanguageListInp
}

type MaintenanceLanguageListRes struct {
	List []*input_spa.SpaMaintenanceLanguageListModel `json:"list"   dc:"数据列表"`
}
