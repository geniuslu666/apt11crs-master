// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Consumer is the golang structure for table consumer.
type Consumer struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	Company    string      `json:"company"    orm:"company"     description:"公司名"`
	Realname   string      `json:"realname"   orm:"realname"    description:"姓名"`
	Score      string      `json:"score"      orm:"score"       description:"级别"`
	ConsumerSn string      `json:"consumerSn" orm:"consumer_sn" description:"客户编号"`
	EntId      string      `json:"entId"      orm:"ent_id"      description:"企业ID"`
	KefuName   string      `json:"kefuName"   orm:"kefu_name"   description:"kefu名称"`
	Tel        string      `json:"tel"        orm:"tel"         description:"手机"`
	Wechat     string      `json:"wechat"     orm:"wechat"      description:"微信"`
	Qq         string      `json:"qq"         orm:"qq"          description:"qq"`
	Email      string      `json:"email"      orm:"email"       description:"邮箱"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
}
