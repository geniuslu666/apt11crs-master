// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberIntention is the golang structure of table hg_pms_member_intention for DAO operations like Where/Data.
type PmsMemberIntention struct {
	g.Meta    `orm:"table:hg_pms_member_intention, do:true"`
	Id        interface{} // 主键ID
	MemberId  interface{} // 意向用户ID
	Name      interface{} //
	Sex       interface{} // 1、男 2、女
	Country   interface{} //
	Phone     interface{} //
	PhoneArea interface{} //
	Language  interface{} //
	Mail      interface{} //
	Remark    interface{} // 备注
	WechatNo  interface{} // 微信号
	LineId    interface{} // line_id
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
