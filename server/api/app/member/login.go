package member

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta    `path:"/member/login" method:"post" tags:"APP_MEMBER" summary:"授权_登录"`
	LoginType string `json:"login_type" v:"required|in:password,phone,email,googleOauth,yahooOauth#login_method_unknown|login_method_not_exist" dc:"必填 登录方式password（手机号或者邮箱作为用户名使用密码登录）,phone（手机号验证登录）,email（邮箱验证登录）,googleOauth（谷歌联登）"`
	Source    string `json:"source"     v:"in:Android,IOS,H5,WX_MINI#login_source_unknown|login_source_not_exist" dc:"必填 来源AndroidApp（安卓APP登录）,IosApp（IosApp登录）,H5（网页登录）Android,IOS,H5,WX_MINI"`
	Phone     string `json:"phone"      v:"required-if:login_type,phone|max-length:20#phone_number_unknown|phone_number_length_limit" dc:"手机号"`
	AreaNo    string `json:"area_no"    v:"required-if:login_type,phone|max-length:10#phone_international_code_unknown|phone_international_code_length_limit" dc:"区号"`
	Email     string `json:"email"      v:"required-if:login_type,email,googleOauth,yahooOauth|max-length:255#email_unknown|email_too_long"`
	Code      string `json:"code"       v:"required-unless:login_type,password,login_type,login_type,login_type,googleOauth,login_type,yahooOauth|length:6,6#verification_code_unknown|verification_code_too_long"`
	Password  string `json:"password"   v:"required-if:login_type,password|max-length:255#password_unknown|password_too_long" dc:"密码"`
	Referrer  int    `json:"referrer"   dc:"推荐人ID"`
	MdCode    string `json:"mdCode"     dc:"登录设备码"`
	MpModel   string `json:"mpModel"    dc:"登录设备型号"`
}

type LoginRes struct {
	Token     string `json:"token" dc:"用户token令牌身份"`
	Expires   int64  `json:"expires" dc:"用户token令牌有效期"`
	Replenish bool   `json:"isNewMember" dc:"是否新用户 true、未补充   false、已补充"`
}

type AuthIdBindLoginReq struct {
	g.Meta   `path:"/member/bindLogin" method:"post" tags:"APP_MEMBER" summary:"授权_登录绑定"`
	BindType string `json:"bindType" v:"required|in:PHONE,phone,EMAIL,email,WX_MINI,wx_mini#bindTypeUnknown|bindTypeFormatError" dc:"绑定类型phone,email,WX_MINI"`
	AuthId   string `json:"authId"   v:"required#auth_id_identifier_unknown" dc:"授权ID标识"`
	Email    string `json:"email"    v:"required-unless:bindType,PHONE,bindType,phone,bindType,WX_MINI|email#email_unknown|email_format_error" dc:"邮箱"`
	Code     string `json:"code"     v:"required#verification_code_unknown" dc:"验证码"`
	Phone    string `json:"phone"    v:"required-unless:bindType,EMAIL,bindType,email|max-length:20#phone_number_unknown|phone_number_length_limit" dc:"手机号"`
	AreaNo   string `json:"area_no"  v:"required-unless:bindType,EMAIL,bindType,email|max-length:10#phone_international_code_unknown|phone_international_code_length_limit" dc:"区号 +86"`
	Source   string `json:"source"   v:"required|in:Android,IOS,H5,WX_MINI,H5_FX#source_unknown|source_format_error" dc:"来源"`
	Referrer int    `json:"referrer" dc:"推荐人ID"`
	MdCode   string `json:"mdCode"     dc:"登录设备码"`
	MpModel  string `json:"mpModel"    dc:"登录设备型号"`
}

type AuthIdBindLoginRes struct {
	Token     string `json:"token" dc:"用户token令牌身份"`
	Expires   int64  `json:"expires" dc:"用户token令牌有效期"`
	Replenish bool   `json:"isNewMember" dc:"是否新用户 true、未补充   false、已补充"`
}
