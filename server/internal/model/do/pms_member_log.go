// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLog is the golang structure of table hg_pms_member_log for DAO operations like Where/Data.
type PmsMemberLog struct {
	g.Meta    `orm:"table:hg_pms_member_log, do:true"`
	Id        interface{} //
	MemberId  interface{} // 会员ID
	LoginTime *gtime.Time // 登录时间
	LoginType interface{} // 登录方式
	LoginIp   interface{} // 登录IP
	MdCode    interface{} // 登录设备码
	MpModel   interface{} // 登录设备型号
	Token     interface{} // 登录token
	ExpirTime *gtime.Time // 过期时间
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
