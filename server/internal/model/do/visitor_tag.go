// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorTag is the golang structure of table visitor_tag for DAO operations like Where/Data.
type VisitorTag struct {
	g.Meta    `orm:"table:visitor_tag, do:true"`
	Id        interface{} //
	VisitorId interface{} // 访客ID
	TagId     interface{} // 标签ID
	EntId     interface{} // 客服企业ID
	CreatedAt *gtime.Time // 创建时间
	Kefu      interface{} // 客服账户
}
