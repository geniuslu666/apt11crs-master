// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsTemplateLog is the golang structure of table hg_sys_sms_template_log for DAO operations like Where/Data.
type SysSmsTemplateLog struct {
	g.Meta         `orm:"table:hg_sys_sms_template_log, do:true"`
	Id             interface{} // 主键
	TemplateId     interface{} // 系统模版表中主键ID
	SendTemplateId interface{} // 发送的模版ID
	Type           interface{} // 1短信   2邮件   3推送
	To             interface{} // 短信或邮件或推送头
	VipId          interface{} // 会员ID
	Content        interface{} // 发送内容
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time //
}
