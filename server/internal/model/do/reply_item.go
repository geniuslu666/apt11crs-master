// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyItem is the golang structure of table reply_item for DAO operations like Where/Data.
type ReplyItem struct {
	g.Meta   `orm:"table:reply_item, do:true"`
	Id       interface{} //
	Content  interface{} // 快捷回复内容
	GroupId  interface{} // 快捷回复分组ID
	UserId   interface{} // 客服账户
	ItemName interface{} // 快捷回复标题
	EntId    interface{} // 客服企业ID
	IsTeam   interface{} // 1个人,2团队
}
