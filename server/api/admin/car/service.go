package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// ServiceListReq 查询服务列表
type ServiceListReq struct {
	g.Meta `path:"/carService/list" method:"get" tags:"ADMIN_CAR" summary:"获取服务列表"`
	input_car.CarServiceListInp
}

type ServiceListRes struct {
	input_form.PageRes
	List []*input_car.CarServiceListModel `json:"list"   dc:"数据列表"`
}

// ServiceViewReq 获取服务指定信息
type ServiceViewReq struct {
	g.Meta `path:"/carService/view" method:"get" tags:"ADMIN_CAR" summary:"获取服务指定信息"`
	input_car.CarServiceViewInp
}

type ServiceViewRes struct {
	*input_car.CarServiceViewModel
}

// ServiceEditReq 修改/新增服务
type ServiceEditReq struct {
	g.Meta `path:"/carService/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增服务"`
	input_car.CarServiceEditInp
}

type ServiceEditRes struct{}

// ServiceDeleteReq 删除服务
type ServiceDeleteReq struct {
	g.Meta `path:"/carService/delete" method:"post" tags:"ADMIN_CAR" summary:"删除服务"`
	input_car.CarServiceDeleteInp
}

type ServiceDeleteRes struct{}

// ServiceMaxSortReq 获取服务最大排序
type ServiceMaxSortReq struct {
	g.Meta `path:"/carService/maxSort" method:"get" tags:"ADMIN_CAR" summary:"获取服务最大排序"`
	input_car.CarServiceMaxSortInp
}

type ServiceMaxSortRes struct {
	*input_car.CarServiceMaxSortModel
}

// ServiceStatusReq 更新服务状态
type ServiceStatusReq struct {
	g.Meta `path:"/carService/status" method:"post" tags:"ADMIN_CAR" summary:"更新服务状态"`
	input_car.CarServiceStatusInp
}

type ServiceStatusRes struct{}

// ServiceSortReq 餐厅排序
type ServiceSortReq struct {
	g.Meta `path:"/carService/sortUpdate" method:"post" tags:"ADMIN_CAR" summary:"服务_排序"`
	input_car.CarServiceSortInp
}

type ServiceSortRes struct{}
