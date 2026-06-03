// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarHelp is the golang structure of table hg_car_help for DAO operations like Where/Data.
type CarHelp struct {
	g.Meta   `orm:"table:hg_car_help, do:true"`
	Id       interface{} //
	Language interface{} // 语言
	Image    interface{} // 主图
	Content  interface{} // 内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
