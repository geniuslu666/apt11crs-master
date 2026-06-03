// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberAuth is the golang structure for table pms_member_auth.
type PmsMemberAuth struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	AuthId    string      `json:"authId"    orm:"auth_id"    description:"第三方用户ID标识"`
	WxUnionid string      `json:"wxUnionid" orm:"wx_unionid" description:"微信unionid"`
	Channel   string      `json:"channel"   orm:"channel"    description:"WX、APPLE"`
	MemberId  int         `json:"memberId"  orm:"member_id"  description:"关联的会员ID"`
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`
	Phone     string      `json:"phone"     orm:"phone"      description:"手机号"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
	IsFxAuth  string      `json:"isFxAuth"  orm:"is_fx_auth" description:"是否是分销授权用户"`
}
