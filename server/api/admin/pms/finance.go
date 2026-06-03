package pms

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type FinanceStatReq struct {
	g.Meta `path:"/finance/stat" method:"get" tags:"ADMIN_PMS" summary:"财务_概况"`
	input_basics.FinanceStatInp
}

type FinanceStatRes struct {
	*input_basics.FinanceStatModel
}

type FinanceListReq struct {
	g.Meta `path:"/finance/list" method:"get" tags:"ADMIN_PMS" summary:"财务_列表"`
	input_basics.FinanceListInp
}

type FinanceListRes struct {
	*input_basics.FinanceListModel
}

type FinanceExportReq struct {
	g.Meta `path:"/finance/export" method:"get" tags:"ADMIN_PMS" summary:"财务_导出"`
	input_basics.FinanceListInp
}

type FinanceExportRes struct{}
