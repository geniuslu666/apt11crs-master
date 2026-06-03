// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Rate is the golang structure of table rate for DAO operations like Where/Data.
type Rate struct {
	g.Meta       `orm:"table:rate, do:true"`
	Id           interface{} //
	KefuName     interface{} // 客服账户
	KefuNickname interface{} // 客服昵称
	VisitorId    interface{} // 访客id
	Content      interface{} // 评价内容
	EntId        interface{} // 企业ID
	Score        interface{} // 评价分数
	CreatedAt    *gtime.Time // 创建时间
}
