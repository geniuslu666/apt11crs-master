// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WechatLogin is the golang structure for table wechat_login.
type WechatLogin struct {
	Id         int         `json:"id"         orm:"id"           description:""`
	KefuName   string      `json:"kefuName"   orm:"kefu_name"    description:"客服账户"`
	OpenId     string      `json:"openId"     orm:"open_id"      description:"微信公众号openid"`
	TempKefuId string      `json:"tempKefuId" orm:"temp_kefu_id" description:"临时客服ID"`
	Status     string      `json:"status"     orm:"status"       description:"当前状态"`
	EntId      string      `json:"entId"      orm:"ent_id"       description:"企业ID"`
	LoginIp    string      `json:"loginIp"    orm:"login_ip"     description:"登录IP"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`
}
