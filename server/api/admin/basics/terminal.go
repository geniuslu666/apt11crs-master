package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// TerminalListReq 查询终端列表
type TerminalListReq struct {
	g.Meta `path:"/terminal/list" method:"get" tags:"ADMIN" summary:"获取终端列表"`
	input_basics.TerminalListInp
}

type TerminalListRes struct {
	input_form.PageRes
	List []*input_basics.TerminalListModel `json:"list"   dc:"数据列表"`
}

// TerminalViewReq 获取终端指定信息
type TerminalViewReq struct {
	g.Meta `path:"/terminal/view" method:"get" tags:"ADMIN" summary:"获取终端指定信息"`
	input_basics.TerminalViewInp
}

type TerminalViewRes struct {
	*input_basics.TerminalViewModel
}

// TerminalEditReq 修改/新增终端
type TerminalEditReq struct {
	g.Meta `path:"/terminal/edit" method:"post" tags:"ADMIN" summary:"修改/新增终端"`
	input_basics.TerminalEditInp
}

type TerminalEditRes struct{}

// TerminalDeleteReq 删除终端
type TerminalDeleteReq struct {
	g.Meta `path:"/terminal/delete" method:"post" tags:"ADMIN" summary:"删除终端"`
	input_basics.TerminalDeleteInp
}

type TerminalDeleteRes struct{}

// BrandListReq 查询品牌型号列表
type BrandListReq struct {
	g.Meta `path:"/brandModel/list" method:"get" tags:"ADMIN" summary:"获取品牌型号列表"`
	input_basics.BrandListInp
}

type BrandListRes struct {
	input_form.PageRes
	List []*input_basics.BrandListModel `json:"list"   dc:"数据列表"`
}

// BrandViewReq 获取终端指定信息
type BrandViewReq struct {
	g.Meta `path:"/brandModel/view" method:"get" tags:"ADMIN" summary:"获取品牌型号信息"`
	input_basics.BrandViewInp
}

type BrandViewRes struct {
	*input_basics.BrandViewModel
}

// BrandEditReq 修改/新增终端
type BrandEditReq struct {
	g.Meta `path:"/brandModel/edit" method:"post" tags:"ADMIN" summary:"修改品牌型号"`
	input_basics.BrandEditInp
}

type BrandEditRes struct{}

// PrintTestReq 打印测试
type PrintTestReq struct {
	g.Meta `path:"/terminal/printTest" method:"post" tags:"ADMIN" summary:"打印测试"`
	input_basics.PrinterTestInp
}

type PrintTestRes struct{}
