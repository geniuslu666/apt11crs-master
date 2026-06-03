package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// CooperateTypeListReq 查询按摩营业类型列表
type CooperateTypeListReq struct {
	g.Meta `path:"/spaCooperateType/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩营业类型列表"`
	input_spa.SpaCooperateTypeListInp
}

type CooperateTypeListRes struct {
	input_form.PageRes
	List []*input_spa.SpaCooperateTypeListModel `json:"list"   dc:"数据列表"`
}

// CooperateTypeExportReq 导出按摩营业类型列表
type CooperateTypeExportReq struct {
	g.Meta `path:"/spaCooperateType/export" method:"get" tags:"ADMIN_SPA" summary:"导出按摩营业类型列表"`
	input_spa.SpaCooperateTypeListInp
}

type CooperateTypeExportRes struct{}

// CooperateTypeViewReq 获取按摩营业类型指定信息
type CooperateTypeViewReq struct {
	g.Meta `path:"/spaCooperateType/view" method:"get" tags:"ADMIN_SPA" summary:"获取按摩营业类型指定信息"`
	input_spa.SpaCooperateTypeViewInp
}

type CooperateTypeViewRes struct {
	*input_spa.SpaCooperateTypeViewModel
}

// CooperateTypeEditReq 修改/新增按摩营业类型
type CooperateTypeEditReq struct {
	g.Meta `path:"/spaCooperateType/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增按摩营业类型"`
	input_spa.SpaCooperateTypeEditInp
}

type CooperateTypeEditRes struct{}

// CooperateTypeDeleteReq 删除按摩营业类型
type CooperateTypeDeleteReq struct {
	g.Meta `path:"/spaCooperateType/delete" method:"post" tags:"ADMIN_SPA" summary:"删除按摩营业类型"`
	input_spa.SpaCooperateTypeDeleteInp
}

type CooperateTypeDeleteRes struct{}

// CooperateTypeStatusReq 更新按摩营业类型状态
type CooperateTypeStatusReq struct {
	g.Meta `path:"/spaCooperateType/status" method:"post" tags:"ADMIN_SPA" summary:"更新按摩营业类型状态"`
	input_spa.SpaCooperateTypeStatusInp
}

type CooperateTypeStatusRes struct{}
