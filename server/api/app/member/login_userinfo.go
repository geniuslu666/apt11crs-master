package member

import (
	"APT/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type GetUserInfoReq struct {
	g.Meta `path:"/member/info" method:"post" tags:"APP_MEMBER" summary:"用户信息_获取用户信息"`
}

type GetUserInfoRes struct {
	entity.PmsMember
	ShowDataBoard bool `json:"showDataBoard" dc:"是否显示数据看板"`
	LevelInfo     struct {
		gmeta.Meta `orm:"table:hg_pms_member_level"`
		entity.PmsMemberLevel
	} `json:"levelInfo" orm:"with:id=level"`
}

type BindEmailReq struct {
	g.Meta `path:"/member/bindEmail" method:"post" tags:"APP_MEMBER" summary:"用户信息_绑定邮箱"`
	Email  string `v:"required#email_unknown" dc:"邮箱"`
	Code   string `v:"required#verification_code_unknown" dc:"验证码"`
}
type BindEmailRes struct {
}

type BindPhoneReq struct {
	g.Meta `path:"/member/bindPhone" method:"post" tags:"APP_MEMBER" summary:"用户信息_绑定手机号"`
	Phone  string `v:"required#phone_number_unknown" dc:"手机号"`
	AreaNo string `v:"required#phone_international_code_unknown" dc:"手机国际区号"`
	Code   string `v:"required#verification_code_unknown" dc:"验证码"`
}

type BindPhoneRes struct {
}

type StaffInfoReq struct {
	g.Meta  `path:"/member/staffInfo" method:"post" tags:"APP_MEMBER" summary:"用户信息_获取员工信息"`
	StaffId int64 `v:"required#employee_id_unknown" dc:"员工ID"`
}

type StaffInfoRes struct {
	entity.PmsStaff
}

type ChannelInfoReq struct {
	g.Meta    `path:"/member/channelInfo" method:"post" tags:"APP_MEMBER" summary:"用户信息_获取渠道信息"`
	ChannelId int64 `v:"required#channel_id_unknown" dc:"渠道ID"`
}

type ChannelInfoRes struct {
	entity.PmsChannel
}

type UserLevelRuleReq struct {
	g.Meta `path:"/member/levelRule" method:"post" tags:"APP_MEMBER" summary:"用户信息_获取会员权益规则"`
}

type UserLevelRuleRes struct {
	LevelRule string `json:"levelRule" dc:"会员权益规则"`
}

type EditUserInfoReq struct {
	g.Meta         `path:"/member/add_user_information" method:"post" tags:"APP_MEMBER" summary:"[用户]补全"`
	Sex            int    `json:"sex"               v:"in:1,2#gender_error" dc:"性别 1、男 2、女"`
	Avatar         string `json:"avatar"            v:"url#avatar_link_format_error" dc:"头像 url上传之后返回的url"`
	FirstName      string `json:"first_name"        v:"" dc:"名"`
	LastName       string `json:"last_name"         v:"" dc:"姓"`
	FullName       string `json:"full_name"         v:"" dc:"全名"`
	Password       string `json:"password"          v:"password2|eq:RepeatPassword#password2_length_verify|password_inconsistency" dc:"密码"` // 密码长度在6~18之间,要求密码必须包含大小写字母和数字
	RepeatPassword string `json:"repeatPassword"    v:"password2|eq:Password#re_password2_length_verify|password_inconsistency" dc:"重复密码"`
	Phone          string `json:"phone"             v:"" dc:"手机号"`
	AreaNo         string `json:"area_no"           v:"" dc:"手机国际区号"`
	Email          string `json:"email"             v:"email#email_format_error" dc:"邮箱"`
	Birthday       string `json:"birthday"          v:"date#birthday_format_error" dc:"生日"`
	Nationality    string `json:"nationality"       v:"" dc:"国家"`
	Address        string `json:"address"           v:"" dc:"地址"`
}

type EditUserInfoRes struct{}
