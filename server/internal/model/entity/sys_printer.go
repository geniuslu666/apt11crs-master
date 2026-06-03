// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPrinter is the golang structure for table sys_printer.
type SysPrinter struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	PrinterType  string      `json:"printerType"  orm:"printer_type"  description:"打印机类型"`
	PrinterName  string      `json:"printerName"  orm:"printer_name"  description:"打印机名称"`
	ClientId     string      `json:"clientId"     orm:"client_id"     description:"易联云第三方应用ID"`
	ClientSecret string      `json:"clientSecret" orm:"client_secret" description:"易联云第三方应用秘钥"`
	MachineCode  string      `json:"machineCode"  orm:"machine_code"  description:"打印机终端号"`
	MachineKey   string      `json:"machineKey"   orm:"machine_key"   description:"终端秘钥"`
	PrintTimes   uint        `json:"printTimes"   orm:"print_times"   description:"打印联数(次数)"`
	Sort         uint        `json:"sort"         orm:"sort"          description:"排序 (数字越大越靠前)"`
	Status       uint        `json:"status"       orm:"status"        description:"状态1、启用 2、禁用"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
