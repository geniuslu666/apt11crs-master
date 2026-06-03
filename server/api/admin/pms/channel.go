package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type ChannelListReq struct {
	g.Meta `path:"/pmsChannel/list" method:"get" tags:"ADMIN_PMS" summary:"渠道_管理列表"`
	input_basics.PmsChannelListInp
}

type ChannelListRes struct {
	input_form.PageRes
	List []*input_basics.PmsChannelListModel `json:"list"   dc:"数据列表"`
}

type ChannelViewReq struct {
	g.Meta `path:"/pmsChannel/view" method:"get" tags:"ADMIN_PMS" summary:"渠道_详情"`
	input_basics.PmsChannelViewInp
}

type ChannelViewRes struct {
	*input_basics.PmsChannelViewModel
}

type ChannelEditReq struct {
	g.Meta `path:"/pmsChannel/edit" method:"post" tags:"ADMIN_PMS" summary:"渠道_修改/新增"`
	input_basics.PmsChannelEditInp
}

type ChannelEditRes struct{}

type ChannelDeleteReq struct {
	g.Meta `path:"/pmsChannel/delete" method:"post" tags:"ADMIN_PMS" summary:"渠道_删除"`
	input_basics.PmsChannelDeleteInp
}

type ChannelDeleteRes struct{}

type ChannelStatusReq struct {
	g.Meta `path:"/pmsChannel/status" method:"post" tags:"ADMIN_PMS" summary:"渠道_更新渠道状态"`
	input_basics.PmsChannelStatusInp
}

type ChannelStatusRes struct{}

type ChannelBindReq struct {
	g.Meta `path:"/pmsChannel/bind" method:"post" tags:"ADMIN_PMS" summary:"渠道_绑定会员"`
	input_basics.PmsChannelBindInp
}

type ChannelBindRes struct{}

type ChannelUnbindReq struct {
	g.Meta `path:"/pmsChannel/unbind" method:"post" tags:"ADMIN_PMS" summary:"渠道_解绑会员"`
	input_basics.PmsChannelUnbindInp
}

type ChannelUnbindRes struct{}
