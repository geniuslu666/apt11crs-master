// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppconfig is the golang structure of table hg_pms_appconfig for DAO operations like Where/Data.
type PmsAppconfig struct {
	g.Meta    `orm:"table:hg_pms_appconfig, do:true"`
	Id        interface{} //
	Name      interface{} // 配置名称
	Key       interface{} // 配置项
	Value     interface{} // 配置值
	Language  interface{} // 语言
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
