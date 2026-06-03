package member

import "github.com/gogf/gf/v2/frame/g"

type CancelReq struct {
	g.Meta       `path:"/member/cancel" method:"post" tags:"APP_MEMBER" summary:"申请注销会员"`
	CancelReason string `json:"cancelReason" dc:"注销原因"`
}

type CancelRes struct{}
