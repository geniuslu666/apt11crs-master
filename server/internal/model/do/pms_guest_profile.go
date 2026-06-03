// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsGuestProfile is the golang structure of table hg_pms_guest_profile for DAO operations like Where/Data.
type PmsGuestProfile struct {
	g.Meta        `orm:"table:hg_pms_guest_profile, do:true"`
	Id            interface{} // 主键
	Uid           interface{} // 三方系统 ID
	MemberId      interface{} // 会员ID
	FirstName     interface{} // 名
	LastName      interface{} // 姓
	FirstNameKana interface{} // 名的假名
	LastNameKana  interface{} // 姓的假名
	FullName      interface{} // 全名
	Language      interface{} // 语言
	Email         interface{} // 电子邮件
	Phone         interface{} // 电话
	AreaNo        interface{} // 电话国际区号
	Nationality   interface{} // 国籍
	Address       interface{} // 地址
	Password      interface{} // 密码
	Register      interface{} // Y 已注册  N 未注册
	RegisterAt    *gtime.Time // 注册时间
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
