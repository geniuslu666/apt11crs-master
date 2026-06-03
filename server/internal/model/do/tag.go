// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta    `orm:"table:tag, do:true"`
	Id        interface{} //
	Name      interface{} // 标签名称
	CreatedAt *gtime.Time // 创建时间
	Kefu      interface{} // 客服账户
	EntId     interface{} // 客服企业ID
}
