// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Qrcode is the golang structure of table qrcode for DAO operations like Where/Data.
type Qrcode struct {
	g.Meta    `orm:"table:qrcode, do:true"`
	Id        interface{} //
	EntId     interface{} // 客服企业ID
	KefuName  interface{} // 客服账户
	Uuid      interface{} // 唯一ID
	Url       interface{} // 跳转的URL
	CreatedAt *gtime.Time // 创建时间
}
