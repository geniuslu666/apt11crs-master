// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarMaintenance is the golang structure of table hg_car_maintenance for DAO operations like Where/Data.
type CarMaintenance struct {
	g.Meta    `orm:"table:hg_car_maintenance, do:true"`
	Id        interface{} //
	Type      interface{} // 类型 1-全部 2-接送机 3-包机
	Language  interface{} // 语言
	Content   interface{} // 内容
	IsDefault interface{} // 是否默认  1是 2否
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
}
