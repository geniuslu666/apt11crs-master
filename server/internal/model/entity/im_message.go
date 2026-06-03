// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessage is the golang structure for table im_message.
type ImMessage struct {
	Id           int64       `json:"id"           orm:"id"             description:""`
	Source       string      `json:"source"       orm:"source"         description:"消息发起方"`
	SourceId     int         `json:"sourceId"     orm:"source_id"      description:"身份对应角色ID"`
	SourceName   string      `json:"sourceName"   orm:"source_name"    description:"身份对应角色昵称"`
	SourcePhoto  string      `json:"sourcePhoto"  orm:"source_photo"   description:"身份对应角色照片"`
	OrderSn      string      `json:"orderSn"      orm:"order_sn"       description:"订单号"`
	FromMemberId int64       `json:"fromMemberId" orm:"from_member_id" description:"发起消息方的用户ID"`
	ToMemberId   int64       `json:"toMemberId"   orm:"to_member_id"   description:"接收消息方的用户ID"`
	Content      *gjson.Json `json:"content"      orm:"content"        description:"消息内容"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:"创建消息时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:"更新消息时间"`
}
