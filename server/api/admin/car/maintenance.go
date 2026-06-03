package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// MaintenanceListReq 查询Maintenance列表
type MaintenanceListReq struct {
	g.Meta `path:"/carMaintenance/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机维护模式列表"`
	input_car.CarMaintenanceListInp
}

type MaintenanceListRes struct {
	input_form.PageRes
	List []*input_car.CarMaintenanceListModel `json:"list"   dc:"数据列表"`
}

// MaintenanceViewReq 获取Maintenance指定信息
type MaintenanceViewReq struct {
	g.Meta `path:"/carMaintenance/view" method:"get" tags:"ADMIN_CAR" summary:"获取接送机维护模式指定信息"`
	input_car.CarMaintenanceViewInp
}

type MaintenanceViewRes struct {
	*input_car.CarMaintenanceViewModel
}

// MaintenanceEditReq 修改/新增Maintenance
type MaintenanceEditReq struct {
	g.Meta `path:"/carMaintenance/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增接送机维护模式"`
	input_car.CarMaintenanceEditInp
}

type MaintenanceEditRes struct{}

// MaintenanceDeleteReq 删除Maintenance
type MaintenanceDeleteReq struct {
	g.Meta `path:"/carMaintenance/delete" method:"post" tags:"ADMIN_CAR" summary:"删除接送机维护模式"`
	input_car.CarMaintenanceDeleteInp
}

type MaintenanceDeleteRes struct{}

// MaintenanceLanguageListReq 获取Maintenance语言
type MaintenanceLanguageListReq struct {
	g.Meta `path:"/carMaintenance/languageList" method:"get" tags:"ADMIN_CAR" summary:"获取接送机维护模式语言"`
	input_car.CarMaintenanceLanguageListInp
}

type MaintenanceLanguageListRes struct {
	List []*input_car.CarMaintenanceLanguageListModel `json:"list"   dc:"数据列表"`
}
