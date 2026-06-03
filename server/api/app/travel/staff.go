package travel

import (
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

type StaffConfigReq struct {
	g.Meta `path:"/config" method:"post" tags:"APP_TRAVEL_STAFF" summary:"[一日游核销端]登录页配置"`
}

type StaffConfigRes struct {
	ContactMobile string `json:"contactMobile"          description:"客服联系电话"`
}

type StaffLoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"APP_TRAVEL_STAFF" summary:"[一日游核销端]登录"`
	Username string `json:"username"       v:"required#account_unknown" dc:"账号"`
	Password string `json:"password"   v:"required#password_unknown" dc:"密码"`
}

type StaffLoginRes struct {
	Token    string `json:"token" dc:"用户token令牌身份"`
	Expires  int64  `json:"expires" dc:"用户token令牌有效期"`
	Name     string `json:"name"          description:"姓名"`
	Mobile   string `json:"mobile"        description:"电话"`
	Username string `json:"username"      description:"登录账号"`
}

type StaffLogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"APP_TRAVEL_STAFF" summary:"[一日游核销端]登出"`
}

type StaffLogoutRes struct{}

type VerifyLogReq struct {
	g.Meta   `path:"/verifyLog/list" method:"post" tags:"APP_TRAVEL_STAFF" summary:"核销记录"`
	PageNum  int `p:"pageNum" v:"required#page_number_unknown" dc:"页码"`
	PageSize int `p:"pageSize" v:"required#page_number_unknown" dc:"页数"`
}

type VerifyLogRes struct {
	List  []*input_travel.VerifyListModel `json:"list"   dc:"数据列表"`
	Count int                             `json:"count"   dc:"数据总数"`
}

type VerifyLogViewReq struct {
	g.Meta `path:"/verifyLog/view" method:"post" tags:"APP_TRAVEL_STAFF" summary:"核销记录详情"`
	input_travel.VerifyLogViewInp
}

type VerifyLogViewRes struct {
	*input_travel.VerifyLogViewModel
}

type CodeViewReq struct {
	g.Meta `path:"/code/view" method:"post" tags:"APP_TRAVEL_STAFF" summary:"二维码详情"`
	input_travel.CodeViewInp
}

type CodeViewRes struct {
	*input_travel.CodeViewModel
}

type CodeVerifyReq struct {
	g.Meta `path:"/code/verify" method:"post" tags:"APP_TRAVEL_STAFF" summary:"二维码核销"`
	input_travel.CodeVerifyInp
}

type CodeVerifyRes struct {
	*input_travel.CodeVerifyModel
}
