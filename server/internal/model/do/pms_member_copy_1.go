// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberCopy1 is the golang structure of table hg_pms_member_copy1 for DAO operations like Where/Data.
type PmsMemberCopy1 struct {
	g.Meta          `orm:"table:hg_pms_member_copy1, do:true"`
	Id              interface{} // 主键
	MemberNo        interface{} // 会员号
	GroupId         interface{} // 会员分组ID
	Avatar          interface{} // 头像
	Sex             interface{} // 1、男 2、女
	FirstName       interface{} // 名
	LastName        interface{} // 姓
	FullName        interface{} // 全名
	Level           interface{} // 等级
	Phone           interface{} // 手机号
	PhoneArea       interface{} // 手机区号
	Mail            interface{} // 邮箱
	Password        interface{} // 密码
	Birthday        *gtime.Time // 生日
	Balance         interface{} // 积分
	Exp             interface{} // 经验
	Referrer        interface{} // 推荐人
	LastReferrer    interface{} // 最后推荐人
	Source          interface{} // 注册来源  IOS,Andriod,H5
	LoginMode       interface{} // password,phone,email,googleOauth
	LastLoginIp     interface{} // 上次登录IP
	LastLogin       *gtime.Time // 上次登录时间
	RegisterIp      interface{} // 注册IP
	RegisterTime    *gtime.Time // 注册时间
	RegisterMdCode  interface{} // 注册设备码
	RegisterMpModel interface{} // 注册设备型号
	Address         interface{} // 地址
	Nationality     interface{} // 国籍
	Replenish       interface{} // 是否完善过用户信息    N 未完善  Y 已完善
	RebateMode      interface{} // MEMBER 会员     STAFF    员工 CHANNEL  渠道
	MemberId        interface{} // 会员ID
	StaffId         interface{} // 员工ID
	ChannelId       interface{} // 渠道ID
	IsNewPrice      interface{} // 是否获取过新人奖  1、获得  0、未获得  2、没有邀新奖励
	IsInviteRewards interface{} //
	Status          interface{} // 状态1、启用 2、禁用
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
