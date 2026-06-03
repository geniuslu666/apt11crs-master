// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	KefuId    string      `json:"kefuId"    orm:"kefu_id"    description:"客服账户"`
	VisitorId string      `json:"visitorId" orm:"visitor_id" description:"访客ID"`
	Content   string      `json:"content"   orm:"content"    description:"消息内容"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
	MesType   string      `json:"mesType"   orm:"mes_type"   description:"消息类型：kefu为客服发送，visitor为访客发送"`
	Status    string      `json:"status"    orm:"status"     description:"已读状态：read为已读，unread为未读"`
	EntId     uint        `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
}
