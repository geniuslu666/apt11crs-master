// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// New is the golang structure of table new for DAO operations like Where/Data.
type New struct {
	g.Meta    `orm:"table:new, do:true"`
	Id        interface{} //
	Tag       interface{} // 公告标签
	Title     interface{} // 公告标题
	Content   interface{} // 公告内容
	CreatedAt *gtime.Time // 创建时间
	Status    interface{} // 状态，未启用
}
