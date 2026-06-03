package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginLogListReq struct {
	g.Meta `path:"/loginLog/list" method:"get" tags:"ADMIN" summary:"登录日志_获取登录日志列表"`
	input_basics.LoginLogListInp
}

type LoginLogListRes struct {
	input_form.PageRes
	List []*input_basics.LoginLogListModel `json:"list"   dc:"数据列表"`
}

type LoginLogExportReq struct {
	g.Meta `path:"/loginLog/export" method:"get" tags:"ADMIN" summary:"登录日志_导出登录日志列表"`
	input_basics.LoginLogListInp
}

type LoginLogExportRes struct{}

type LoginLogDeleteReq struct {
	g.Meta `path:"/loginLog/delete" method:"post" tags:"ADMIN" summary:"登录日志_删除登录日志"`
	input_basics.LoginLogDeleteInp
}

type LoginLogDeleteRes struct{}
