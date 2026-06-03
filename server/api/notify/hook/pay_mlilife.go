package hook

import (
	"APT/internal/library/mlilifeWxPay"
	"github.com/gogf/gf/v2/frame/g"
)

type MlilifePayReq struct {
	g.Meta `path:"/pay/mlilife/return" method:"post" tags:"NOTIFY_HOOK" summary:"mlilife_支付回调RETURN"`
	*mlilifeWxPay.NotifyRes
}

type MlilifePayRes struct {
}
