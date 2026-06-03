package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// IspListReq 查询服务商列表
type IspListReq struct {
	g.Meta `path:"/spaIsp/list" method:"get" tags:"ADMIN_SPA" summary:"获取服务商列表"`
	input_spa.SpaIspListInp
}

type IspListRes struct {
	input_form.PageRes
	List []*input_spa.SpaIspListModel `json:"list"   dc:"数据列表"`
}

// IspAllListReq 查询服务商列表
type IspAllListReq struct {
	g.Meta `path:"/spaIsp/all" method:"get" tags:"ADMIN_SPA" summary:"获取服务商列表"`
	input_spa.SpaIspListInp
}

type IspAllListRes struct {
	List []*input_spa.SpaIspAllListModel `json:"list"   dc:"所有服务商-数据列表"`
}

// IspViewReq 获取服务商详情
type IspViewReq struct {
	g.Meta `path:"/spaIsp/view" method:"get" tags:"ADMIN_SPA" summary:"获取服务商详情"`
	input_spa.SpaIspViewInp
}

type IspViewRes struct {
	*input_spa.SpaIspViewModel
}

// IspEditReq 修改/新增服务商
type IspEditReq struct {
	g.Meta `path:"/spaIsp/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增服务商"`
	input_spa.SpaIspEditInp
}

type IspEditRes struct{}

// IspDeleteReq 删除服务商
type IspDeleteReq struct {
	g.Meta `path:"/spaIsp/delete" method:"post" tags:"ADMIN_SPA" summary:"删除服务商"`
	input_spa.SpaIspDeleteInp
}

type IspDeleteRes struct{}

// IspStatusReq 更新服务商状态
type IspStatusReq struct {
	g.Meta `path:"/spaIsp/status" method:"post" tags:"ADMIN_SPA" summary:"更新服务商状态"`
	input_spa.SpaIspStatusInp
}

type IspStatusRes struct{}

// IspWorkStatusReq 更新服务商工作状态
type IspWorkStatusReq struct {
	g.Meta `path:"/spaIsp/workStatus" method:"post" tags:"ADMIN_SPA" summary:"更新服务商工作状态"`
	input_spa.SpaIspWorkStatusInp
}

type IspWorkStatusRes struct{}

type IspBindReq struct {
	g.Meta `path:"/spaIsp/bind" method:"post" tags:"ADMIN_SPA" summary:"服务商绑定用户"`
	input_spa.SpaIspBindInp
}

type IspBindRes struct{}

type IspUnbindReq struct {
	g.Meta `path:"/spaIsp/unbind" method:"post" tags:"ADMIN_SPA" summary:"服务商解绑用户"`
	input_spa.SpaIspUnbindInp
}

type IspUnbindRes struct{}
