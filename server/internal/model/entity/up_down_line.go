// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UpDownLine is the golang structure for table up_down_line.
type UpDownLine struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	KefuName     string      `json:"kefuName"     orm:"kefu_name"     description:"客服账户"`
	EntId        string      `json:"entId"        orm:"ent_id"        description:"企业ID"`
	OnlineStatus int         `json:"onlineStatus" orm:"online_status" description:"在线状态，1在线，2离线"`
	ClientIp     string      `json:"clientIp"     orm:"client_ip"     description:"ip地址"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}
