// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorAttr is the golang structure of table visitor_attr for DAO operations like Where/Data.
type VisitorAttr struct {
	g.Meta        `orm:"table:visitor_attr, do:true"`
	Id            interface{} //
	VisitorId     interface{} // 访客ID
	RealName      interface{} // 访客真实姓名
	Tel           interface{} // 访客手机号
	Email         interface{} // 访客邮箱
	Qq            interface{} // 访客QQ
	Wechat        interface{} // 访客微信
	Comment       interface{} // 访客备注
	CreatedAt     *gtime.Time // 创建时间
	EntId         interface{} // 对接企业ID
	MaxMessageNum interface{} // 访客最大消息数
}
