package member

import (
	"github.com/gogf/gf/v2/frame/g"
)

type OauthLoginReq struct {
	g.Meta   `path:"/auth/oauthLogin" method:"post" tags:"APP_MEMBER" summary:"授权_微信授权登录"`
	Code     string `json:"Code" v:"required-without-all:AppleId,MiniCode#wechat_auth_code_cannot_be_empty" dc:"微信授权码"`
	MiniCode string `json:"MiniCode" v:"required-without-all:AppleId,Code#mini_code_cannot_be_empty" dc:"微信小程序授权码"`
	AppleId  string `json:"AppleId" v:"required-without-all:Code,MiniCode#appleId_cannot_be_empty" dc:"苹果授权码"`
	Referrer int    `json:"referrer" dc:"推荐人ID"`
	MdCode   string `json:"mdCode"   dc:"登录设备码"`
}

type OauthLoginRes struct {
	Token   string `json:"token" dc:"token"`
	Expires int64  `json:"expires" dc:"过期时长"`
	IsBind  bool   `json:"isBind" dc:"true:已绑定，false:未绑定"`
	AuthId  string `json:"authId" dc:"授权ID标识"`
}

type WeixinAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type PhoneNumberReq struct {
	g.Meta   `path:"/auth/getPhoneNumber" method:"post" tags:"APP_MEMBER" summary:"获取用户手机号"`
	MiniCode string `json:"MiniCode" dc:"微信小程序授权码"`
}

type PhoneNumberRes struct {
	PhoneNumber     string `json:"phoneNumber" dc:"用户绑定的手机号（国外手机号会有区号）"`
	PurePhoneNumber string `json:"purePhoneNumber" dc:"没有区号的手机号"`
	CountryCode     string `json:"countryCode" dc:"区号"`
}

type FxAuthLoginReq struct {
	g.Meta   `path:"/fx/auth/login" method:"post" tags:"APP_MEMBER" summary:"授权_登录"`
	Code     string `json:"Code" v:"required-without-all:AppleId,MiniCode#wechat_auth_code_cannot_be_empty" dc:"分销授权码"`
	Referrer int    `json:"referrer" dc:"推荐人ID"`
	MdCode   string `json:"mdCode"   dc:"登录设备码"`
}

type FxAuthLoginRes struct {
	Token   string `json:"token" dc:"token"`
	Expires int64  `json:"expires" dc:"过期时长"`
	IsBind  bool   `json:"isBind" dc:"true:已绑定，false:未绑定"`
	AuthId  string `json:"authId" dc:"授权ID标识"`
}
