// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMember is the golang structure for table pms_member.
type PmsMember struct {
	Id              int         `json:"id"              orm:"id"                description:"主键"`
	MemberNo        string      `json:"memberNo"        orm:"member_no"         description:"会员号"`
	GroupId         int         `json:"groupId"         orm:"group_id"          description:"会员分组ID"`
	Avatar          string      `json:"avatar"          orm:"avatar"            description:"头像"`
	Sex             int         `json:"sex"             orm:"sex"               description:"1、男 2、女"`
	FirstName       string      `json:"firstName"       orm:"first_name"        description:"名"`
	LastName        string      `json:"lastName"        orm:"last_name"         description:"姓"`
	FullName        string      `json:"fullName"        orm:"full_name"         description:"全名"`
	Level           int         `json:"level"           orm:"level"             description:"等级"`
	Phone           string      `json:"phone"           orm:"phone"             description:"手机号"`
	PhoneArea       string      `json:"phoneArea"       orm:"phone_area"        description:"手机区号"`
	Mail            string      `json:"mail"            orm:"mail"              description:"邮箱"`
	Password        string      `json:"password"        orm:"password"          description:"密码"`
	Birthday        *gtime.Time `json:"birthday"        orm:"birthday"          description:"生日"`
	Balance         float64     `json:"balance"         orm:"balance"           description:"积分"`
	Exp             int         `json:"exp"             orm:"exp"               description:"经验"`
	Referrer        int         `json:"referrer"        orm:"referrer"          description:"推荐人"`
	LastReferrer    int         `json:"lastReferrer"    orm:"last_referrer"     description:"最后推荐人"`
	Source          string      `json:"source"          orm:"source"            description:"注册来源  IOS,Andriod,H5"`
	LoginMode       string      `json:"loginMode"       orm:"login_mode"        description:"password,phone,email,googleOauth"`
	LastLoginIp     string      `json:"lastLoginIp"     orm:"last_login_ip"     description:"上次登录IP"`
	LastLogin       *gtime.Time `json:"lastLogin"       orm:"last_login"        description:"上次登录时间"`
	RegisterIp      string      `json:"registerIp"      orm:"register_ip"       description:"注册IP"`
	RegisterTime    *gtime.Time `json:"registerTime"    orm:"register_time"     description:"注册时间"`
	RegisterMdCode  string      `json:"registerMdCode"  orm:"register_md_code"  description:"注册设备码"`
	RegisterMpModel string      `json:"registerMpModel" orm:"register_mp_model" description:"注册设备型号"`
	Address         string      `json:"address"         orm:"address"           description:"地址"`
	Nationality     string      `json:"nationality"     orm:"nationality"       description:"国籍"`
	Replenish       string      `json:"replenish"       orm:"replenish"         description:"是否完善过用户信息    N 未完善  Y 已完善"`
	RebateMode      string      `json:"rebateMode"      orm:"rebate_mode"       description:"MEMBER 会员     STAFF    员工 CHANNEL  渠道"`
	MemberId        int         `json:"memberId"        orm:"member_id"         description:"会员ID"`
	StaffId         int         `json:"staffId"         orm:"staff_id"          description:"员工ID"`
	ChannelId       int         `json:"channelId"       orm:"channel_id"        description:"渠道ID"`
	IsNewPrice      uint        `json:"isNewPrice"      orm:"is_new_price"      description:"是否获取过新人奖  1、获得  0、未获得  2、没有邀新奖励"`
	IsInviteRewards string      `json:"isInviteRewards" orm:"is_Invite_rewards" description:""`
	Status          uint        `json:"status"          orm:"status"            description:"状态1、启用 2、禁用"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}
