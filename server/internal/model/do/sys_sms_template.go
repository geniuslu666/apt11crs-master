// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsTemplate is the golang structure of table hg_sys_sms_template for DAO operations like Where/Data.
type SysSmsTemplate struct {
	g.Meta          `orm:"table:hg_sys_sms_template, do:true"`
	Id              interface{} // 主键
	Scene           interface{} // 1会员  2住宿  3接送机  4按摩
	Event           interface{} // 别名
	Title           interface{} // 标题
	Content         interface{} // 内容（多语言）
	UmsTemplate     *gjson.Json // 一信通模版json（id：模版id， content：模版内容）
	TencentTemplate *gjson.Json // 腾讯云模版json
	AliyunTemplate  *gjson.Json // 阿里云模版json
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time //
}
