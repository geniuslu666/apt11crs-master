package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// TechnicianListReq 查询技师列表
type TechnicianListReq struct {
	g.Meta `path:"/spaTechnician/list" method:"get" tags:"ADMIN_SPA" summary:"获取技师列表"`
	input_spa.SpaTechnicianListInp
}

type TechnicianListRes struct {
	input_form.PageRes
	List []*input_spa.SpaTechnicianListModel `json:"list"   dc:"数据列表"`
}

// TechnicianAllListReq 查询技师列表
type TechnicianAllListReq struct {
	g.Meta `path:"/spaTechnician/all" method:"get" tags:"ADMIN_SPA" summary:"获取技师列表"`
	input_spa.SpaTechnicianListInp
}

type TechnicianAllListRes struct {
	List []*input_spa.SpaTechnicianAllListModel `json:"list"   dc:"所有技师-数据列表"`
}

// TechnicianViewReq 获取技师详情
type TechnicianViewReq struct {
	g.Meta `path:"/spaTechnician/view" method:"get" tags:"ADMIN_SPA" summary:"获取技师详情"`
	input_spa.SpaTechnicianViewInp
}

type TechnicianViewRes struct {
	*input_spa.SpaTechnicianViewModel
}

// TechnicianEditReq 修改/新增技师
type TechnicianEditReq struct {
	g.Meta `path:"/spaTechnician/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增技师"`
	input_spa.SpaTechnicianEditInp
}

type TechnicianEditRes struct{}

// TechnicianDeleteReq 删除技师
type TechnicianDeleteReq struct {
	g.Meta `path:"/spaTechnician/delete" method:"post" tags:"ADMIN_SPA" summary:"删除技师"`
	input_spa.SpaTechnicianDeleteInp
}

type TechnicianDeleteRes struct{}

// TechnicianStatusReq 更新技师状态
type TechnicianStatusReq struct {
	g.Meta `path:"/spaTechnician/status" method:"post" tags:"ADMIN_SPA" summary:"更新技师状态"`
	input_spa.SpaTechnicianStatusInp
}

type TechnicianStatusRes struct{}

// TechnicianWorkStatusReq 更新技师工作状态
type TechnicianWorkStatusReq struct {
	g.Meta `path:"/spaTechnician/workStatus" method:"post" tags:"ADMIN_SPA" summary:"更新技师工作状态"`
	input_spa.SpaTechnicianWorkStatusInp
}

type TechnicianWorkStatusRes struct{}

type TechnicianBindReq struct {
	g.Meta `path:"/spaTechnician/bind" method:"post" tags:"ADMIN_SPA" summary:"技师绑定用户"`
	input_spa.SpaTechnicianBindInp
}

type TechnicianBindRes struct{}

type TechnicianUnbindReq struct {
	g.Meta `path:"/spaTechnician/unbind" method:"post" tags:"ADMIN_SPA" summary:"技师解绑用户"`
	input_spa.SpaTechnicianUnbindInp
}

type TechnicianUnbindRes struct{}
