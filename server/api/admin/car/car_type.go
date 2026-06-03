package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// CarTypeListReq 查询车辆车型列表
type CarTypeListReq struct {
	g.Meta `path:"/carCarType/list" method:"get" tags:"ADMIN_CAR" summary:"获取车辆车型列表"`
	input_car.CarCarTypeListInp
}

type CarTypeListRes struct {
	input_form.PageRes
	List []*input_car.CarCarTypeListModel `json:"list"   dc:"数据列表"`
}

// CarTypeViewReq 获取车辆车型指定信息
type CarTypeViewReq struct {
	g.Meta `path:"/carCarType/view" method:"get" tags:"ADMIN_CAR" summary:"获取车辆车型指定信息"`
	input_car.CarCarTypeViewInp
}

type CarTypeViewRes struct {
	*input_car.CarCarTypeViewModel
}

// CarTypeEditReq 修改/新增车辆车型
type CarTypeEditReq struct {
	g.Meta `path:"/carCarType/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增车辆车型"`
	input_car.CarCarTypeEditInp
}

type CarTypeEditRes struct{}

// CarTypeDeleteReq 删除车辆车型
type CarTypeDeleteReq struct {
	g.Meta `path:"/carCarType/delete" method:"post" tags:"ADMIN_CAR" summary:"删除车辆车型"`
	input_car.CarCarTypeDeleteInp
}

type CarTypeDeleteRes struct{}

// CarTypeMaxSortReq 获取车辆车型最大排序
type CarTypeMaxSortReq struct {
	g.Meta `path:"/carCarType/maxSort" method:"get" tags:"ADMIN_CAR" summary:"获取车辆车型最大排序"`
	input_car.CarCarTypeMaxSortInp
}

type CarTypeMaxSortRes struct {
	*input_car.CarCarTypeMaxSortModel
}

// CarTypeStatusReq 更新车辆车型状态
type CarTypeStatusReq struct {
	g.Meta `path:"/carCarType/status" method:"post" tags:"ADMIN_CAR" summary:"更新车辆车型状态"`
	input_car.CarCarTypeStatusInp
}

type CarTypeStatusRes struct{}
