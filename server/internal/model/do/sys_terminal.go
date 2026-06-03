// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminal is the golang structure of table hg_sys_terminal for DAO operations like Where/Data.
type SysTerminal struct {
	g.Meta       `orm:"table:hg_sys_terminal, do:true"`
	Id           interface{} //
	TerminalName interface{} // 终端名称
	TerminalType interface{} // 终端类型
	BrandModel   interface{} // 品牌型号
	Sn           interface{} // 终端编号
	OnlineStatus interface{} // 在线状态 1-在线 2离线
	StoreId      interface{} // 绑定门店
	RestaurantId interface{} // 绑定餐厅
	PrintTimes   interface{} // 打印联数(次数)
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
}
