// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDouyin is the golang structure for table user_douyin.
type UserDouyin struct {
	Id                 int         `json:"id"                 orm:"id"                   description:""`
	KefuName           string      `json:"kefuName"           orm:"kefu_name"            description:"客服账户"`
	Nickname           string      `json:"nickname"           orm:"nickname"             description:"抖音昵称"`
	Avatar             string      `json:"avatar"             orm:"avatar"               description:"抖音头像"`
	OpenId             string      `json:"openId"             orm:"open_id"              description:"抖音OpenId"`
	UnionId            string      `json:"unionId"            orm:"union_id"             description:"抖音union_id"`
	AccessToken        string      `json:"accessToken"        orm:"access_token"         description:"抖音AccessToken"`
	ExpiresIn          *gtime.Time `json:"expiresIn"          orm:"expires_in"           description:"抖音AccessToken过期时间"`
	RefreshToken       string      `json:"refreshToken"       orm:"refresh_token"        description:"抖音refresh_token"`
	RefreshExpiresIn   *gtime.Time `json:"refreshExpiresIn"   orm:"refresh_expires_in"   description:"抖音refresh_token过期时间"`
	ClientToken        string      `json:"clientToken"        orm:"client_token"         description:"抖音client_token"`
	ClientTokenExpires *gtime.Time `json:"clientTokenExpires" orm:"client_token_expires" description:"抖音client_token过期时间"`
	EntId              string      `json:"entId"              orm:"ent_id"               description:"企业ID"`
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:"创建时间"`
}
