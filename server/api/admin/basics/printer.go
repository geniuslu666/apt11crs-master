package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// PrinterListReq 查询打印机列表
type PrinterListReq struct {
	g.Meta `path:"/printer/list" method:"get" tags:"ADMIN" summary:"获取打印机列表"`
	input_basics.PrinterListInp
}

type PrinterListRes struct {
	input_form.PageRes
	List []*input_basics.PrinterListModel `json:"list"   dc:"数据列表"`
}

// PrinterViewReq 获取打印机指定信息
type PrinterViewReq struct {
	g.Meta `path:"/printer/view" method:"get" tags:"ADMIN" summary:"获取打印机指定信息"`
	input_basics.PrinterViewInp
}

type PrinterViewRes struct {
	*input_basics.PrinterViewModel
}

// PrinterEditReq 修改/新增打印机
type PrinterEditReq struct {
	g.Meta `path:"/printer/edit" method:"post" tags:"ADMIN" summary:"修改/新增打印机"`
	input_basics.PrinterEditInp
}

type PrinterEditRes struct{}

// PrinterDeleteReq 删除打印机
type PrinterDeleteReq struct {
	g.Meta `path:"/printer/delete" method:"post" tags:"ADMIN" summary:"删除打印机"`
	input_basics.PrinterDeleteInp
}

type PrinterDeleteRes struct{}

// PrinterMaxSortReq 获取打印机最大排序
type PrinterMaxSortReq struct {
	g.Meta `path:"/printer/maxSort" method:"get" tags:"ADMIN" summary:"获取打印机最大排序"`
	input_basics.PrinterMaxSortInp
}

type PrinterMaxSortRes struct {
	*input_basics.PrinterMaxSortModel
}

// PrinterStatusReq 更新打印机状态
type PrinterStatusReq struct {
	g.Meta `path:"/printer/status" method:"post" tags:"ADMIN" summary:"更新打印机状态"`
	input_basics.PrinterStatusInp
}

type PrinterStatusRes struct{}

// PrinterCarOrderReq 打印接送机订单
type PrinterCarOrderReq struct {
	g.Meta `path:"/printer/printCarOrder" method:"post" tags:"ADMIN" summary:"打印接送机订单"`
	input_basics.PrinterCarOrderInp
}

type PrinterCarOrderRes struct{}

// PrinterFoodOrderReq 打印餐厅订单
type PrinterFoodOrderReq struct {
	g.Meta `path:"/printer/printFoodOrder" method:"post" tags:"ADMIN" summary:"打印餐厅订单"`
	input_basics.PrinterCarOrderInp
}

type PrinterFoodOrderRes struct{}

// PrinterSpaOrderReq 打印按摩订单
type PrinterSpaOrderReq struct {
	g.Meta `path:"/printer/printSpaOrder" method:"post" tags:"ADMIN" summary:"打印按摩订单"`
	input_basics.PrinterCarOrderInp
}

type PrinterSpaOrderRes struct{}
