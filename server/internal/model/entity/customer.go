// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure for table customer.
type Customer struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	Name         string      `json:"name"         orm:"name"          description:"会员名称"`
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"会员头像"`
	KefuName     string      `json:"kefuName"     orm:"kefu_name"     description:"客服账户"`
	Tel          string      `json:"tel"          orm:"tel"           description:"会员手机号"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	Openid       string      `json:"openid"       orm:"openid"        description:"会员ID"`
	AcountOpenid string      `json:"acountOpenid" orm:"acount_openid" description:"公众号会员ID"`
	Extra        string      `json:"extra"        orm:"extra"         description:"会员扩展信息"`
	EntId        uint        `json:"entId"        orm:"ent_id"        description:"对接的企业ID"`
	Score        uint        `json:"score"        orm:"score"         description:"会员积分"`
}
