package input_basics

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// PrinterUpdateFields 修改打印机字段过滤
type PrinterUpdateFields struct {
	PrinterName  string `json:"printerName"  dc:"打印机名称"`
	PrinterType  string `json:"printerType"  dc:"打印机类型"`
	ClientId     string `json:"clientId"     dc:"易联云第三方应用ID"`
	ClientSecret string `json:"clientSecret" dc:"易联云第三方应用秘钥"`
	MachineCode  string `json:"machineCode"  dc:"打印机终端号"`
	MachineKey   string `json:"machineKey"   dc:"打印机终端秘钥"`
	PrintTimes   int    `json:"printTimes"   dc:"打印联数(次数)"`
	Status       int    `json:"status"       dc:"状态"`
	Sort         int    `json:"sort"         dc:"排序 (数字越大越靠前)"`
}

// PrinterInsertFields 新增打印机字段过滤
type PrinterInsertFields struct {
	PrinterName  string `json:"printerName"  dc:"打印机名称"`
	PrinterType  string `json:"printerType"  dc:"打印机类型"`
	ClientId     string `json:"clientId"     dc:"易联云第三方应用ID"`
	ClientSecret string `json:"clientSecret" dc:"易联云第三方应用秘钥"`
	MachineCode  string `json:"machineCode"  dc:"打印机终端号"`
	MachineKey   string `json:"machineKey"   dc:"打印机终端秘钥"`
	PrintTimes   int    `json:"printTimes"   dc:"打印联数(次数)"`
	Status       int    `json:"status"       dc:"状态"`
	Sort         int    `json:"sort"         dc:"排序 (数字越大越靠前)"`
}

// PrinterEditInp 修改/新增打印机
type PrinterEditInp struct {
	entity.SysPrinter
}

func (in *PrinterEditInp) Filter(ctx context.Context) (err error) {
	if in.PrinterName == "" {
		err = gerror.New("请输入打印机名称")
		return
	}
	if in.ClientId == "" {
		err = gerror.New("请输入第三方应用ID")
		return
	}
	if in.ClientSecret == "" {
		err = gerror.New("请输入第三方应用秘钥")
		return
	}
	if in.MachineCode == "" {
		err = gerror.New("请输入打印机终端号")
		return
	}
	return
}

type PrinterEditModel struct{}

// PrinterDeleteInp 删除打印机
type PrinterDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PrinterDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PrinterDeleteModel struct{}

// PrinterViewInp 获取指定打印机信息
type PrinterViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PrinterViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PrinterViewModel struct {
	entity.SysPrinter
}

// PrinterListInp 获取打印机列表
type PrinterListInp struct {
	input_form.PageReq
	PrinterName  string        `json:"printerName"                     dc:"打印机名称"`
	ClientId     string        `json:"clientId"     dc:"易联云第三方应用ID"`
	ClientSecret string        `json:"clientSecret" dc:"易联云第三方应用秘钥"`
	MachineCode  string        `json:"machineCode"  dc:"打印机终端号"`
	CreatedAt    []*gtime.Time `json:"createdAt"                 dc:"created_at"`
}

func (in *PrinterListInp) Filter(ctx context.Context) (err error) {
	return
}

type PrinterListModel struct {
	Id           int         `json:"id"            dc:"id"`
	PrinterName  string      `json:"printerName"   dc:"打印机名称"`
	PrinterType  string      `json:"printerType"  dc:"打印机类型"`
	ClientId     string      `json:"clientId"      dc:"易联云第三方应用ID"`
	ClientSecret string      `json:"clientSecret"  dc:"易联云第三方应用秘钥"`
	MachineCode  string      `json:"machineCode"   dc:"打印机终端号"`
	MachineKey   string      `json:"machineKey"    dc:"打印机终端秘钥"`
	PrintTimes   uint        `json:"printTimes"    dc:"打印联数(次数)"`
	Sort         int         `json:"sort"          dc:"排序 (数字越大越靠前)"`
	Status       int         `json:"status"        dc:"状态"`
	CreateAt     *gtime.Time `json:"createAt"      dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"      dc:"更新时间"`
}

// PrinterMaxSortInp 获取车辆车型最大排序
type PrinterMaxSortInp struct{}

func (in *PrinterMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PrinterMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// PrinterStatusInp 更新打印机状态
type PrinterStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PrinterStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	return
}

type PrinterStatusModel struct{}

// PrinterCarOrderInp 接送机订单打印
type PrinterCarOrderInp struct {
	OrderId  int    `json:"orderId"  dc:"接送机订单ID"`
	OrderSn  string `json:"orderSn"  dc:"订单号"`
	IsRefund int    `json:"isRefund"  dc:"是否退款 1-打印退款"`
}

type GetYILINKFoodOrderContentInp struct {
	FoodOrder  *input_food.FoodOrderViewModel
	PrintTimes int `json:"printTimes"  dc:"打印联数"`
}

type GetYILINKCarOrderContentInp struct {
	CarOrder   *input_car.CarOrderViewModel
	PrintTimes int `json:"printTimes"  dc:"打印联数"`
	IsRefund   int `json:"isRefund"  dc:"是否退款 1-打印退款"`
}

type GetYILINKSpaOrderContentInp struct {
	SpaOrder   *input_spa.SpaOrderViewModel
	PrintTimes int `json:"printTimes"  dc:"打印联数"`
}

type GetPrintOrderContentModel struct {
	Content string `json:"content"  dc:"打印内容"`
}

type PrintContentInp struct {
	OrderSn    string `json:"orderSn"  dc:"订单号"`
	Content    string `json:"content"  dc:"打印内容"`
	PrintTimes int    `json:"printTimes"  dc:"打印联数"`
	Printer    *entity.SysPrinter
}
