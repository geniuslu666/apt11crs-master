package paypay

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DoPayPayCallbackReq struct {
	g.Meta `path:"/paypay/callback" method:"post" tags:"NOTIFY_HOOK" summary:"paypay_支付通知"`
}

type DoPayPayCallbackRes struct {
}
