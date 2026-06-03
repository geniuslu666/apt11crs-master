package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type ServeLogListReq struct {
	g.Meta `path:"/serveLog/list" method:"get" tags:"ADMIN" summary:"服务日志_获取服务日志列表"`
	input_basics.ServeLogListInp
}

type ServeLogListRes struct {
	input_form.PageRes
	List []*input_basics.ServeLogListModel `json:"list"   dc:"数据列表"`
}

type ServeLogExportReq struct {
	g.Meta `path:"/serveLog/export" method:"get" tags:"ADMIN" summary:"服务日志_导出服务日志列表"`
	input_basics.ServeLogListInp
}

type ServeLogExportRes struct{}

type ServeLogViewReq struct {
	g.Meta `path:"/serveLog/view" method:"get" tags:"ADMIN" summary:"服务日志_获取服务日志指定信息"`
	input_basics.ServeLogViewInp
}

type ServeLogViewRes struct {
	*input_basics.ServeLogViewModel
}

type ServeLogDeleteReq struct {
	g.Meta `path:"/serveLog/delete" method:"post" tags:"ADMIN" summary:"服务日志_删除服务日志"`
	input_basics.ServeLogDeleteInp
}

type ServeLogDeleteRes struct{}
