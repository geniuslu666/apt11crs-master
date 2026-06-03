// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBrokerage is the golang structure for table pms_brokerage.
type PmsBrokerage struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Identity  string      `json:"identity"  orm:"identity"   description:"类型    CHANNEL   渠道   STAFF  员工"`
	Scene     string      `json:"scene"     orm:"scene"      description:"场景值   HOTEL 酒店  SYSTEM 系统"`
	Type      string      `json:"type"      orm:"type"       description:"REBATE 返佣  WITHDRAW  提现  SYS 系统调整"`
	Balance   float64     `json:"balance"   orm:"balance"    description:"变动金额"`
	StaffId   int         `json:"staffId"   orm:"staff_id"   description:"员工ID"`
	ChannelId int         `json:"channelId" orm:"channel_id" description:"渠道ID"`
	MemberId  int         `json:"memberId"  orm:"member_id"  description:"会员ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
