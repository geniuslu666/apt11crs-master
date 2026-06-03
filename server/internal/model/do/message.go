// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta    `orm:"table:message, do:true"`
	Id        interface{} //
	KefuId    interface{} // 客服账户
	VisitorId interface{} // 访客ID
	Content   interface{} // 消息内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	MesType   interface{} // 消息类型：kefu为客服发送，visitor为访客发送
	Status    interface{} // 已读状态：read为已读，unread为未读
	EntId     interface{} // 客服企业ID
}
