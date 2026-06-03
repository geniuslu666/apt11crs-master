package member

import "github.com/gogf/gf/v2/frame/g"

type GetMemberQrCodeReq struct {
	g.Meta `path:"/member/getMemberQrCode" method:"post" tags:"APP_MEMBER" summary:"会员二维码_获取会员二维码"`
}

type GetMemberQrCodeRes struct {
	Code string `json:"code" dc:"AES加密后的动态二维码内容"`
}
