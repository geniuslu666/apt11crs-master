// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsTemplate is the golang structure for table sys_sms_template.
type SysSmsTemplate struct {
	Id              int64       `json:"id"              orm:"id"               description:"主键"`
	Scene           int         `json:"scene"           orm:"scene"            description:"1会员  2住宿  3接送机  4按摩"`
	Event           string      `json:"event"           orm:"event"            description:"别名"`
	Title           string      `json:"title"           orm:"title"            description:"标题"`
	Content         string      `json:"content"         orm:"content"          description:"内容（多语言）"`
	UmsTemplate     *gjson.Json `json:"umsTemplate"     orm:"ums_template"     description:"一信通模版json（id：模版id， content：模版内容）"`
	TencentTemplate *gjson.Json `json:"tencentTemplate" orm:"tencent_template" description:"腾讯云模版json"`
	AliyunTemplate  *gjson.Json `json:"aliyunTemplate"  orm:"aliyun_template"  description:"阿里云模版json"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:""`
}
