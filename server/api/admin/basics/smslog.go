package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type SmsLogListReq struct {
	g.Meta `path:"/smsLog/list" method:"get" tags:"ADMIN" summary:"短信_获取短信记录列表"`
	input_basics.SmsLogListInp
}

type SmsLogListRes struct {
	List []*input_basics.SmsLogListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type SmsLogViewReq struct {
	g.Meta `path:"/smsLog/view" method:"get" tags:"ADMIN" summary:"短信_获取指定短信信息"`
	input_basics.SmsLogViewInp
}

type SmsLogViewRes struct {
	*input_basics.SmsLogViewModel
}

type SmsLogDeleteReq struct {
	g.Meta `path:"/smsLog/delete" method:"post" tags:"ADMIN" summary:"短信_删除短信记录"`
	input_basics.SmsLogDeleteInp
}

type SmsLogDeleteRes struct{}
