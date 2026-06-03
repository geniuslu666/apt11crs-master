package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type BlacklistListReq struct {
	g.Meta `path:"/blacklist/list" method:"get" tags:"ADMIN" summary:"黑名单_获取黑名单列表"`
	input_basics.BlacklistListInp
}

type BlacklistListRes struct {
	List []*input_basics.BlacklistListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type BlacklistViewReq struct {
	g.Meta `path:"/blacklist/view" method:"get" tags:"ADMIN" summary:"黑名单_获取指定信息"`
	input_basics.BlacklistViewInp
}

type BlacklistViewRes struct {
	*input_basics.BlacklistViewModel
}

type BlacklistEditReq struct {
	g.Meta `path:"/blacklist/edit" method:"post" tags:"ADMIN" summary:"黑名单_修改/新增黑名单"`
	input_basics.BlacklistEditInp
}

type BlacklistEditRes struct{}

type BlacklistDeleteReq struct {
	g.Meta `path:"/blacklist/delete" method:"post" tags:"ADMIN" summary:"黑名单_删除黑名单"`
	input_basics.BlacklistDeleteInp
}

type BlacklistDeleteRes struct{}

type BlacklistStatusReq struct {
	g.Meta `path:"/blacklist/status" method:"post" tags:"ADMIN" summary:"黑名单_更新黑名单状态"`
	input_basics.BlacklistStatusInp
}

type BlacklistStatusRes struct{}
