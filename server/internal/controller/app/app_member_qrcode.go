package app

import (
	"APT/api/app/member"
	"APT/internal/service"
	"context"
)

// GetMemberQrCode 获取会员二维码
func (c *ControllerMember) GetMemberQrCode(ctx context.Context, req *member.GetMemberQrCodeReq) (res *member.GetMemberQrCodeRes, err error) {

	code, err := service.AppMember().GetMemberQrCode(ctx)
	if err != nil {
		return nil, err
	}

	res = &member.GetMemberQrCodeRes{
		Code: code,
	}

	return res, nil
}
