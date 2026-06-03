// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLog is the golang structure for table pms_member_log.
type PmsMemberLog struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	MemberId  int         `json:"memberId"  orm:"member_id"  description:"会员ID"`
	LoginTime *gtime.Time `json:"loginTime" orm:"login_time" description:"登录时间"`
	LoginType string      `json:"loginType" orm:"login_type" description:"登录方式"`
	LoginIp   string      `json:"loginIp"   orm:"login_ip"   description:"登录IP"`
	MdCode    string      `json:"mdCode"    orm:"md_code"    description:"登录设备码"`
	MpModel   string      `json:"mpModel"   orm:"mp_model"   description:"登录设备型号"`
	Token     string      `json:"token"     orm:"token"      description:"登录token"`
	ExpirTime *gtime.Time `json:"expirTime" orm:"expir_time" description:"过期时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
