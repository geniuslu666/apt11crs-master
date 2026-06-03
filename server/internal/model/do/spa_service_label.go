// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServiceLabel is the golang structure of table hg_spa_service_label for DAO operations like Where/Data.
type SpaServiceLabel struct {
	g.Meta    `orm:"table:hg_spa_service_label, do:true"`
	ServiceId interface{} // 服务ID
	LabelId   interface{} // 标签ID
}
