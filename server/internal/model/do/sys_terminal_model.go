// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminalModel is the golang structure of table hg_sys_terminal_model for DAO operations like Where/Data.
type SysTerminalModel struct {
	g.Meta       `orm:"table:hg_sys_terminal_model, do:true"`
	Id           interface{} //
	BrandModel   interface{} // 品牌型号
	ClientId     interface{} // 开发者ID
	ClientSecret interface{} // 开发者秘钥
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
	DeletedAt    *gtime.Time //
}
