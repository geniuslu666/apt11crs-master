// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsTemplateLog is the golang structure for table sys_sms_template_log.
type SysSmsTemplateLog struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键"`
	TemplateId     uint        `json:"templateId"     orm:"template_id"      description:"系统模版表中主键ID"`
	SendTemplateId string      `json:"sendTemplateId" orm:"send_template_id" description:"发送的模版ID"`
	Type           uint        `json:"type"           orm:"type"             description:"1短信   2邮件   3推送"`
	To             string      `json:"to"             orm:"to"               description:"短信或邮件或推送头"`
	VipId          uint        `json:"vipId"          orm:"vip_id"           description:"会员ID"`
	Content        string      `json:"content"        orm:"content"          description:"发送内容"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`
}
