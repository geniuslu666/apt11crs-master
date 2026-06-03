// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyGroup is the golang structure of table reply_group for DAO operations like Where/Data.
type ReplyGroup struct {
	g.Meta    `orm:"table:reply_group, do:true"`
	Id        interface{} //
	GroupName interface{} // 组名
	UserId    interface{} // 客服账户
	EntId     interface{} // 客服企业ID
	IsTeam    interface{} // 1个人,2团队
}
