// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Oauth is the golang structure for table oauth.
type Oauth struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	UserId    string      `json:"userId"    orm:"user_id"    description:"访客/客服账户"`
	OauthId   string      `json:"oauthId"   orm:"oauth_id"   description:"公众号OPENID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Status    int         `json:"status"    orm:"status"     description:"状态，未启用"`
}
