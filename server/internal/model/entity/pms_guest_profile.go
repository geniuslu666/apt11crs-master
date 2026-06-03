// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsGuestProfile is the golang structure for table pms_guest_profile.
type PmsGuestProfile struct {
	Id            int         `json:"id"            orm:"id"              description:"主键"`
	Uid           string      `json:"uid"           orm:"uid"             description:"三方系统 ID"`
	MemberId      int         `json:"memberId"      orm:"member_id"       description:"会员ID"`
	FirstName     string      `json:"firstName"     orm:"first_name"      description:"名"`
	LastName      string      `json:"lastName"      orm:"last_name"       description:"姓"`
	FirstNameKana string      `json:"firstNameKana" orm:"first_name_kana" description:"名的假名"`
	LastNameKana  string      `json:"lastNameKana"  orm:"last_name_kana"  description:"姓的假名"`
	FullName      string      `json:"fullName"      orm:"full_name"       description:"全名"`
	Language      string      `json:"language"      orm:"language"        description:"语言"`
	Email         string      `json:"email"         orm:"email"           description:"电子邮件"`
	Phone         string      `json:"phone"         orm:"phone"           description:"电话"`
	AreaNo        string      `json:"areaNo"        orm:"area_no"         description:"电话国际区号"`
	Nationality   string      `json:"nationality"   orm:"nationality"     description:"国籍"`
	Address       string      `json:"address"       orm:"address"         description:"地址"`
	Password      string      `json:"password"      orm:"password"        description:"密码"`
	Register      string      `json:"register"      orm:"register"        description:"Y 已注册  N 未注册"`
	RegisterAt    *gtime.Time `json:"registerAt"    orm:"register_at"     description:"注册时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:""`
}
