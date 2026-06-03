// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorAttr is the golang structure for table visitor_attr.
type VisitorAttr struct {
	Id            int         `json:"id"            orm:"id"              description:""`
	VisitorId     string      `json:"visitorId"     orm:"visitor_id"      description:"访客ID"`
	RealName      string      `json:"realName"      orm:"real_name"       description:"访客真实姓名"`
	Tel           string      `json:"tel"           orm:"tel"             description:"访客手机号"`
	Email         string      `json:"email"         orm:"email"           description:"访客邮箱"`
	Qq            string      `json:"qq"            orm:"qq"              description:"访客QQ"`
	Wechat        string      `json:"wechat"        orm:"wechat"          description:"访客微信"`
	Comment       string      `json:"comment"       orm:"comment"         description:"访客备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	EntId         uint        `json:"entId"         orm:"ent_id"          description:"对接企业ID"`
	MaxMessageNum string      `json:"maxMessageNum" orm:"max_message_num" description:"访客最大消息数"`
}
