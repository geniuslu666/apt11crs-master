// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsPropertyRegion is the golang structure of table hg_pms_property_region for DAO operations like Where/Data.
type PmsPropertyRegion struct {
	g.Meta    `orm:"table:hg_pms_property_region, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 区域名称（多语言）
	Status    interface{} // 状态1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
