// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorExt is the golang structure of table visitor_ext for DAO operations like Where/Data.
type VisitorExt struct {
	g.Meta    `orm:"table:visitor_ext, do:true"`
	Id        interface{} //
	VisitorId interface{} // 访客ID
	EntId     interface{} // 对接的企业ID
	Ua        interface{} // 访客浏览器UserAgent
	Title     interface{} // 页面标题
	Url       interface{} // 页面地址
	Refer     interface{} // 页面来源
	ReferUrl  interface{} // 页面来源地址
	ClientIp  interface{} // 访客IP
	City      interface{} // 访客城市
	Language  interface{} // 浏览器语言
	CreatedAt *gtime.Time // 创建时间
}
