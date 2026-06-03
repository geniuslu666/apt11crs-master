package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// DriverListReq 查询司机管理列表
type DriverListReq struct {
	g.Meta `path:"/carDriver/list" method:"get" tags:"ADMIN_CAR" summary:"获取司机管理列表"`
	input_car.CarDriverListInp
}

type DriverListRes struct {
	input_form.PageRes
	List []*input_car.CarDriverListModel `json:"list"   dc:"数据列表"`
}

// DriverAllListReq 查询司机列表
type DriverAllListReq struct {
	g.Meta `path:"/carDriver/all" method:"get" tags:"ADMIN_CAR" summary:"获取司机列表"`
	input_car.CarDriverListInp
}

type DriverAllListRes struct {
	List []*input_car.CarDriverAllListModel `json:"list"   dc:"所有司机-数据列表"`
}

// DriverViewReq 获取司机详情
type DriverViewReq struct {
	g.Meta `path:"/carDriver/view" method:"get" tags:"ADMIN_CAR" summary:"获取司机详情"`
	input_car.CarDriverViewInp
}

type DriverViewRes struct {
	*input_car.CarDriverViewModel
}

// DriverEditReq 修改/新增司机
type DriverEditReq struct {
	g.Meta `path:"/carDriver/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增司机"`
	input_car.CarDriverEditInp
}

type DriverEditRes struct{}

// DriverDeleteReq 删除司机
type DriverDeleteReq struct {
	g.Meta `path:"/carDriver/delete" method:"post" tags:"ADMIN_CAR" summary:"删除司机"`
	input_car.CarDriverDeleteInp
}

type DriverDeleteRes struct{}

// DriverStatusReq 更新司机状态
type DriverStatusReq struct {
	g.Meta `path:"/carDriver/status" method:"post" tags:"ADMIN_CAR" summary:"更新司机状态"`
	input_car.CarDriverStatusInp
}

type DriverStatusRes struct{}

// DriverWorkStatusReq 更新工作状态
type DriverWorkStatusReq struct {
	g.Meta `path:"/carDriver/workStatus" method:"post" tags:"ADMIN_CAR" summary:"更新司机工作状态"`
	input_car.CarDriverWorkStatusInp
}

type DriverWorkStatusRes struct{}

type DriverBindReq struct {
	g.Meta `path:"/carDriver/bind" method:"post" tags:"ADMIN_CAR" summary:"司机绑定用户"`
	input_car.CarDriverBindInp
}

type DriverBindRes struct{}

type DriverUnbindReq struct {
	g.Meta `path:"/carDriver/unbind" method:"post" tags:"ADMIN_CAR" summary:"司机解绑用户"`
	input_car.CarDriverUnbindInp
}

type DriverUnbindRes struct{}

type DriverBindCarReq struct {
	g.Meta `path:"/carDriver/bindCar" method:"post" tags:"ADMIN_CAR" summary:"司机绑定车辆"`
	input_car.CarDriverBindCarInp
}

type DriverBindCarRes struct{}
