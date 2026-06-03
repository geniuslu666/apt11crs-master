package member

import (
	"APT/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberIntentionEditReq struct {
	g.Meta    `path:"/member/add_intention" method:"post" tags:"APP_MEMBER" summary:"意向_修改/新增"`
	Name      string `json:"name"      orm:"name"       description:"称呼"`
	Sex       int    `json:"sex"       orm:"sex"        description:"1、男 2、女"`
	Country   string `json:"country"   orm:"country"    description:"国家/地区"`
	Phone     string `json:"phone"     orm:"phone"      description:"手机号"`
	PhoneArea string `json:"phoneArea" orm:"phone_area" description:"手机区号"`
	Language  string `json:"language"  orm:"language"   description:"喜好语言（简体中文：zh  日语：ja  韩语：ko  英语：en  繁体中文：zh_CN）"`
	Mail      string `json:"mail"      orm:"mail"       description:"邮箱"`
	Remark    string `json:"remark"    orm:"remark"     description:"备注"`
	WechatNo  string `json:"wechatNo"    orm:"wechat_no"     description:"微信号"`
	LineId    string `json:"lineId"    orm:"line_id"     description:"LINE ID"`
}

type MemberIntentionEditRes struct{}

type GetMemberIntentionConfigReq struct {
	g.Meta `path:"/member/intention_config" method:"post" tags:"APP_MEMBER" summary:"意向_会员意向基础配置" `
}

type GetMemberIntentionConfigRes struct {
	*model.MemberIntentionConfig
}
