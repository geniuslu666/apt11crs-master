// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDouyin is the golang structure of table user_douyin for DAO operations like Where/Data.
type UserDouyin struct {
	g.Meta             `orm:"table:user_douyin, do:true"`
	Id                 interface{} //
	KefuName           interface{} // 客服账户
	Nickname           interface{} // 抖音昵称
	Avatar             interface{} // 抖音头像
	OpenId             interface{} // 抖音OpenId
	UnionId            interface{} // 抖音union_id
	AccessToken        interface{} // 抖音AccessToken
	ExpiresIn          *gtime.Time // 抖音AccessToken过期时间
	RefreshToken       interface{} // 抖音refresh_token
	RefreshExpiresIn   *gtime.Time // 抖音refresh_token过期时间
	ClientToken        interface{} // 抖音client_token
	ClientTokenExpires *gtime.Time // 抖音client_token过期时间
	EntId              interface{} // 企业ID
	CreatedAt          *gtime.Time // 创建时间
}
