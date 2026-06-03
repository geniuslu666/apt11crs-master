// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPrinter is the golang structure of table hg_sys_printer for DAO operations like Where/Data.
type SysPrinter struct {
	g.Meta       `orm:"table:hg_sys_printer, do:true"`
	Id           interface{} //
	PrinterType  interface{} // 打印机类型
	PrinterName  interface{} // 打印机名称
	ClientId     interface{} // 易联云第三方应用ID
	ClientSecret interface{} // 易联云第三方应用秘钥
	MachineCode  interface{} // 打印机终端号
	MachineKey   interface{} // 终端秘钥
	PrintTimes   interface{} // 打印联数(次数)
	Sort         interface{} // 排序 (数字越大越靠前)
	Status       interface{} // 状态1、启用 2、禁用
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
}
