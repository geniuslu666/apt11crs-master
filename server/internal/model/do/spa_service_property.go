// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServiceProperty is the golang structure of table hg_spa_service_property for DAO operations like Where/Data.
type SpaServiceProperty struct {
	g.Meta     `orm:"table:hg_spa_service_property, do:true"`
	ServiceId  interface{} // 服务ID
	PropertyId interface{} // 物业ID
}
