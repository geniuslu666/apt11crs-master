// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberIntention is the golang structure for table pms_member_intention.
type PmsMemberIntention struct {
	Id        int         `json:"id"        orm:"id"         description:"主键ID"`
	MemberId  int         `json:"memberId"  orm:"member_id"  description:"意向用户ID"`
	Name      string      `json:"name"      orm:"name"       description:""`
	Sex       int         `json:"sex"       orm:"sex"        description:"1、男 2、女"`
	Country   string      `json:"country"   orm:"country"    description:""`
	Phone     string      `json:"phone"     orm:"phone"      description:""`
	PhoneArea string      `json:"phoneArea" orm:"phone_area" description:""`
	Language  string      `json:"language"  orm:"language"   description:""`
	Mail      string      `json:"mail"      orm:"mail"       description:""`
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`
	WechatNo  string      `json:"wechatNo"  orm:"wechat_no"  description:"微信号"`
	LineId    string      `json:"lineId"    orm:"line_id"    description:"line_id"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
