package app

import (
	"APT/internal/library/cache"
	"APT/internal/model/input/input_app_member"
	"APT/utility/uuid"
	"context"
	"time"

	"APT/api/app/fx"
)

func (c *ControllerFx) Login(ctx context.Context, req *fx.LoginReq) (res *fx.LoginRes, err error) {
	var (
		authCode   = uuid.CreatePayCode("FX")
		memberInfo = &input_app_member.FxMemberInfo{
			AuthId:  req.AuthId,
			Channel: req.Channel,
		}
	)
	res = new(fx.LoginRes)
	if err = cache.Instance().Set(ctx, authCode, memberInfo, time.Second*300); err != nil {
		return
	}
	res.LoginUrl = "https://crsdev-api.yeebok.net/#/?auth_code=" + authCode
	return
}
