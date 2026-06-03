package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type LogClearReq struct {
	g.Meta `path:"/log/clear" method:"post" tags:"ADMIN" summary:"访问日志_清空日志"`
}

type LogClearRes struct{}

type LogExportReq struct {
	g.Meta `path:"/log/export" method:"get" tags:"ADMIN" summary:"访问日志_导出日志"`
	input_basics.LogListInp
}

type LogExportRes struct{}

type LogListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"ADMIN" summary:"访问日志_获取日志列表"`
	input_basics.LogListInp
}

type LogListRes struct {
	List []*input_basics.LogListModel `json:"list" dc:"数据列表"`
	input_form.PageRes
}

type LogDeleteReq struct {
	g.Meta `path:"/log/delete" method:"post" tags:"ADMIN" summary:"访问日志_删除日志"`
	input_basics.LogDeleteInp
}

type LogDeleteRes struct{}

type LogViewReq struct {
	g.Meta `path:"/log/view" method:"get" tags:"ADMIN" summary:"访问日志_获取指定信息"`
	input_basics.LogViewInp
}

type LogViewRes struct {
	*input_basics.LogViewModel
}
