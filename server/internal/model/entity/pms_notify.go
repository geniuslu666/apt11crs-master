// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsNotify is the golang structure for table pms_notify.
type PmsNotify struct {
	Id            int         `json:"id"            orm:"id"             description:""`
	MemberId      int         `json:"memberId"      orm:"member_id"      description:"会员ID"`
	NotifyTitle   string      `json:"notifyTitle"   orm:"notify_title"   description:"通知标题"`
	NotifyType    string      `json:"notifyType"    orm:"notify_type"    description:"通知类型 SYSTEM、系统消息    BOOKING、预定消息"`
	NotifyContent string      `json:"notifyContent" orm:"notify_content" description:"通知内容"`
	NotifyData    *gjson.Json `json:"notifyData"    orm:"notify_data"    description:"通知数据"`
	IsRead        string      `json:"isRead"        orm:"is_read"        description:"Y 已读  N 未读"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`
}
