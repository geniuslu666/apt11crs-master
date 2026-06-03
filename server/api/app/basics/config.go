package basics

import (
	"APT/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAppConfigReq struct {
	g.Meta `path:"/getAppConfig" method:"post" tags:"APP_BASICS" summary:"[配置]应用配置" `
}

type GetAppConfigRes struct {
	*model.AppConfig
}

type GetAppLanguageConfigReq struct {
	g.Meta `path:"/getAppLanguageConfig" method:"post" tags:"APP_BASICS" summary:"[配置]语言配置" `
}

type GetAppLanguageConfigRes struct {
	LanguageVersion    string `json:"languageVersion" response:"string" dc:"语言版本"`
	LanguageVersionUrl string `json:"languageVersionUrl" response:"string" dc:"语言版本数据URL"`
}

type GetPayModeConfigReq struct {
	g.Meta `path:"/getPayModeConfig" method:"post" tags:"APP_BASICS" summary:"[配置]支付配置" `
}

type GetPayModeConfigRes struct {
	PayModeList []string `json:"payModeList" response:"arr" dc:"支付方式列表"`
}

// GetContactConfigReq 获取联系我们配置请求
type GetContactConfigReq struct {
	g.Meta `path:"/getContactConfig" method:"post" tags:"APP_BASICS" summary:"[配置]联系我们配置" `
}

// GetContactConfigRes 获取联系我们配置响应
type GetContactConfigRes struct {
	Japan *ContactInfo `json:"japan" dc:"日本客服中心"`
	China *ContactInfo `json:"china" dc:"中国客服中心"`
}

// ContactInfo 客服信息
type ContactInfo struct {
	Phone    string `json:"phone"    dc:"电话"`
	Email    string `json:"email"    dc:"邮箱"`
	WorkTime string `json:"workTime" dc:"工作时间"`
}

// GetProfileButtonConfigReq 获取个人中心按钮配置请求
type GetProfileButtonConfigReq struct {
	g.Meta `path:"/getProfileButtonConfig" method:"post" tags:"APP_BASICS" summary:"[配置]个人中心按钮配置" `
}

// GetProfileButtonConfigRes 获取个人中心按钮配置响应
type GetProfileButtonConfigRes struct {
	Data *ProfileButtonInfo `json:"data" dc:"数据"`
}

// ProfileButtonInfo 按钮信息
type ProfileButtonInfo struct {
	Qrcode    int `json:"qrcode"    dc:"会员码开关（1-开启  0-关闭）"`
	Coupon    int `json:"coupon"    dc:"优惠券开关（1-开启  0-关闭）"`
	ThCoupon  int `json:"thCoupon"    dc:"礼品券开关（1-开启  0-关闭）"`
	DataBoard int `json:"dataBoard"    dc:"数据看板开关（1-开启  0-关闭）"`
	Employee  int `json:"employee"    dc:"员工专区开关（1-开启  0-关闭）"`
	Workbench int `json:"workbench"    dc:"工作台开关（1-开启  0-关闭）"`
	Fx        int `json:"fx"    dc:"渠道分销开关（1-开启  0-关闭）"`
	Help      int `json:"help"    dc:"帮助中心开关（1-开启  0-关闭）"`
	Contact   int `json:"contact"    dc:"联系我们开关（1-开启  0-关闭）"`
	Agreement int `json:"agreement"    dc:"用户协议开关（1-开启  0-关闭）"`
	Privacy   int `json:"privacy"    dc:"隐私政策开关（1-开启  0-关闭）"`
}
