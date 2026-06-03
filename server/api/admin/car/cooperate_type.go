package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type CooperateTypeListReq struct {
	g.Meta `path:"/carCooperateType/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机合作类型列表"`
	input_car.CarCooperateTypeListInp
}

type CooperateTypeListRes struct {
	input_form.PageRes
	List []*input_car.CarCooperateTypeListModel `json:"list"   dc:"数据列表"`
}

type CooperateTypeExportReq struct {
	g.Meta `path:"/carCooperateType/export" method:"get" tags:"ADMIN_CAR" summary:"导出接送机合作类型列表"`
	input_car.CarCooperateTypeListInp
}

type CooperateTypeExportRes struct{}

type CooperateTypeViewReq struct {
	g.Meta `path:"/carCooperateType/view" method:"get" tags:"ADMIN_CAR" summary:"获取接送机合作类型指定信息"`
	input_car.CarCooperateTypeViewInp
}

type CooperateTypeViewRes struct {
	*input_car.CarCooperateTypeViewModel
}

type CooperateTypeEditReq struct {
	g.Meta `path:"/carCooperateType/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增接送机合作类型"`
	input_car.CarCooperateTypeEditInp
}

type CooperateTypeEditRes struct{}

type CooperateTypeDeleteReq struct {
	g.Meta `path:"/carCooperateType/delete" method:"post" tags:"ADMIN_CAR" summary:"删除接送机合作类型"`
	input_car.CarCooperateTypeDeleteInp
}

type CooperateTypeDeleteRes struct{}

type CooperateTypeStatusReq struct {
	g.Meta `path:"/carCooperateType/status" method:"post" tags:"ADMIN_CAR" summary:"更新接送机合作类型状态"`
	input_car.CarCooperateTypeStatusInp
}

type CooperateTypeStatusRes struct{}
