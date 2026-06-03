// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorBlack is the golang structure of table visitor_black for DAO operations like Where/Data.
type VisitorBlack struct {
	g.Meta    `orm:"table:visitor_black, do:true"`
	Id        interface{} //
	VisitorId interface{} // 访客ID
	Name      interface{} // 访客名称
	CreatedAt *gtime.Time // 创建时间
	EntId     interface{} // 客服企业ID
	KefuName  interface{} // 客服账户
}
