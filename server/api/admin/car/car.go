package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// CarListReq 查询车辆列表
type CarListReq struct {
	g.Meta `path:"/carCar/list" method:"get" tags:"ADMIN_CAR" summary:"获取车辆列表"`
	input_car.CarCarListInp
}

type CarListRes struct {
	input_form.PageRes
	List []*input_car.CarCarListModel `json:"list"   dc:"数据列表"`
}

// CarViewReq 获取车辆指定信息
type CarViewReq struct {
	g.Meta `path:"/carCar/view" method:"get" tags:"ADMIN_CAR" summary:"获取车辆指定信息"`
	input_car.CarCarViewInp
}

type CarViewRes struct {
	*input_car.CarCarViewModel
}

// CarEditReq 修改/新增车辆
type CarEditReq struct {
	g.Meta `path:"/carCar/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增车辆"`
	input_car.CarCarEditInp
}

type CarEditRes struct{}

// CarDeleteReq 删除车辆
type CarDeleteReq struct {
	g.Meta `path:"/carCar/delete" method:"post" tags:"ADMIN_CAR" summary:"删除车辆"`
	input_car.CarCarDeleteInp
}

type CarDeleteRes struct{}

// CarMaxSortReq 获取车辆最大排序
type CarMaxSortReq struct {
	g.Meta `path:"/carCar/maxSort" method:"get" tags:"ADMIN_CAR" summary:"获取车辆最大排序"`
	input_car.CarCarMaxSortInp
}

type CarMaxSortRes struct {
	*input_car.CarCarMaxSortModel
}

// CarStatusReq 更新车辆状态
type CarStatusReq struct {
	g.Meta `path:"/carCar/status" method:"post" tags:"ADMIN_CAR" summary:"更新车辆状态"`
	input_car.CarCarStatusInp
}

type CarStatusRes struct{}

// WorkStatusReq 更新工作状态
type WorkStatusReq struct {
	g.Meta `path:"/carCar/workStatus" method:"post" tags:"ADMIN_CAR" summary:"更新车辆工作状态"`
	input_car.CarCarWorkStatusInp
}

type WorkStatusRes struct{}
