// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IpAuth is the golang structure for table ip_auth.
type IpAuth struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	Content    string      `json:"content"    orm:"content"     description:"备注"`
	IpAddress  string      `json:"ipAddress"  orm:"ip_address"  description:"IP地址"`
	ExpireTime string      `json:"expireTime" orm:"expire_time" description:"过期时间，未启用"`
	Phone      string      `json:"phone"      orm:"phone"       description:"客服账户手机号"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	Status     int         `json:"status"     orm:"status"      description:"开启状态，1正常，2关闭"`
	Code       string      `json:"code"       orm:"code"        description:"授权码"`
}
