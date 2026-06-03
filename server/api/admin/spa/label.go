package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"

	"github.com/gogf/gf/v2/frame/g"
)

// LabelListReq 查询服务标签列表
type LabelListReq struct {
	g.Meta `path:"/spaLabel/list" method:"get" tags:"ADMIN_SPA" summary:"获取服务标签列表"`
	input_spa.SpaLabelListInp
}

type LabelListRes struct {
	input_form.PageRes
	List []*input_spa.SpaLabelListModel `json:"list"   dc:"数据列表"`
}

// LabelViewReq 获取服务标签指定信息
type LabelViewReq struct {
	g.Meta `path:"/spaLabel/view" method:"get" tags:"ADMIN_SPA" summary:"获取服务标签指定信息"`
	input_spa.SpaLabelViewInp
}

type LabelViewRes struct {
	*input_spa.SpaLabelViewModel
}

// LabelEditReq 修改/新增服务标签
type LabelEditReq struct {
	g.Meta `path:"/spaLabel/edit" method:"post" tags:"ADMIN_SPA" summary:"修改/新增服务标签"`
	input_spa.SpaLabelEditInp
}

type LabelEditRes struct{}

// LabelDeleteReq 删除服务标签
type LabelDeleteReq struct {
	g.Meta `path:"/spaLabel/delete" method:"post" tags:"ADMIN_SPA" summary:"删除服务标签"`
	input_spa.SpaLabelDeleteInp
}

type LabelDeleteRes struct{}

// LabelMaxSortReq 获取服务标签最大排序
type LabelMaxSortReq struct {
	g.Meta `path:"/spaLabel/maxSort" method:"get" tags:"ADMIN_SPA" summary:"获取服务标签最大排序"`
	input_spa.SpaLabelMaxSortInp
}

type LabelMaxSortRes struct {
	*input_spa.SpaLabelMaxSortModel
}

// LabelStatusReq 更新服务标签状态
type LabelStatusReq struct {
	g.Meta `path:"/spaLabel/status" method:"post" tags:"ADMIN_SPA" summary:"更新服务标签状态"`
	input_spa.SpaLabelStatusInp
}

type LabelStatusRes struct{}
