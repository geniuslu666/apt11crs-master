package th

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_th"
	"github.com/gogf/gf/v2/frame/g"
)

type MchListReq struct {
	g.Meta `path:"/thMch/list" method:"get" tags:"ADMIN_TH" summary:"获取商户列表"`
	input_th.ThMchListInp
}

type MchListRes struct {
	input_form.PageRes
	List []*input_th.ThMchListModel `json:"list"   dc:"数据列表"`
}

type MchAllReq struct {
	g.Meta `path:"/thMch/all" method:"get" tags:"ADMIN_TH" summary:"商户全部列表"`
	input_th.ThMchAllInp
}

type MchAllRes struct {
	List []*input_th.ThMchAllModel `json:"list"   dc:"数据列表"`
}

type MchViewReq struct {
	g.Meta `path:"/thMch/view" method:"get" tags:"ADMIN_TH" summary:"获取商户指定信息"`
	input_th.ThMchViewInp
}

type MchViewRes struct {
	*input_th.ThMchViewModel
}

type MchEditReq struct {
	g.Meta `path:"/thMch/edit" method:"post" tags:"ADMIN_TH" summary:"修改/新增商户"`
	input_th.ThMchEditInp
}

type MchEditRes struct{}

type MchDeleteReq struct {
	g.Meta `path:"/thMch/delete" method:"post" tags:"ADMIN_TH" summary:"删除商户"`
	input_th.ThMchDeleteInp
}

type MchDeleteRes struct{}

type MchStatusReq struct {
	g.Meta `path:"/thMch/status" method:"post" tags:"ADMIN_TH" summary:"更新商户状态"`
	input_th.ThMchStatusInp
}

type MchStatusRes struct{}

type MchSwitchReq struct {
	g.Meta `path:"/thMch/switch" method:"post" tags:"ADMIN_TH" summary:"商户_更新状态"`
	input_th.ThMchSwitchInp
}

type MchSwitchRes struct{}
