// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCooperateType is the golang structure of table hg_car_cooperate_type for DAO operations like Where/Data.
type CarCooperateType struct {
	g.Meta    `orm:"table:hg_car_cooperate_type, do:true"`
	Id        interface{} //
	TypeName  interface{} //
	Status    interface{} // 状态1、启用 2、禁用
	IsThird   interface{} // 是否第三方 1-是 2-否
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
