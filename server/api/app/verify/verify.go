package verify

import (
	"APT/internal/model/input/input_th"

	"github.com/gogf/gf/v2/frame/g"
)

// MemberVerifyReq 核销
type MemberVerifyReq struct {
	g.Meta `path:"/memberVerify/verify" method:"get,post" tags:"APP_BASICS" summary:"会员_核销"`
	input_th.ThMemberCouponVerifyInp
}

type MemberVerifyRes struct{}

type MemberVerifyCodeRefreshReq struct {
	g.Meta `path:"/memberVerify/codeRefresh" method:"get,post" tags:"APP_BASICS" summary:"会员_核销码刷新"`
	Scene  string `json:"scene"        dc:"场景 food-餐饮 thCoupon-礼品券"`
	Id     int    `json:"id"           dc:"ID 餐厅订单ID/礼品券领取记录ID"`
}

type MemberVerifyCodeRefreshRes struct {
	Code         string `json:"code"        dc:"券码"`
	VerifyStatus int    `json:"verifyStatus"        dc:"核销状态 礼品券状态 1待生效 2未使用 3已核销 4已过期  5已失效  6已回收； 餐厅：1-未核销,2-已核销"`
}
