package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type EmsLogListReq struct {
	g.Meta `path:"/emsLog/list" method:"get" tags:"ADMIN" summary:"邮件记录_获取邮件记录列表"`
	input_basics.EmsLogListInp
}

type EmsLogListRes struct {
	List []*input_basics.EmsLogListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type EmsLogViewReq struct {
	g.Meta `path:"/emsLog/view" method:"get" tags:"ADMIN" summary:"邮件记录_获取指定信息"`
	input_basics.EmsLogViewInp
}

type EmsLogViewRes struct {
	*input_basics.EmsLogViewModel
}

type EmsLogEditReq struct {
	g.Meta `path:"/emsLog/edit" method:"post" tags:"ADMIN" summary:"邮件记录_修改/新增邮件记录"`
	input_basics.EmsLogEditInp
}

type EmsLogEditRes struct{}

type EmsLogDeleteReq struct {
	g.Meta `path:"/emsLog/delete" method:"post" tags:"ADMIN" summary:"邮件记录_删除邮件记录"`
	input_basics.EmsLogDeleteInp
}

type EmsLogDeleteRes struct{}

type EmsLogStatusReq struct {
	g.Meta `path:"/emsLog/status" method:"post" tags:"ADMIN" summary:"邮件记录_更新邮件记录状态"`
	input_basics.EmsLogStatusInp
}

type EmsLogStatusRes struct{}
