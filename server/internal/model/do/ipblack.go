// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Ipblack is the golang structure of table ipblack for DAO operations like Where/Data.
type Ipblack struct {
	g.Meta   `orm:"table:ipblack, do:true"`
	Id       interface{} //
	Ip       interface{} // IP地址
	Name     interface{} // 名称
	CreateAt *gtime.Time // 创建时间
	KefuId   interface{} // 操作的客服账户
	EntId    interface{} // 客服企业ID
}
