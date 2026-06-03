package app

import (
	"context"

	"APT/api/app/basics"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
)

func (c *ControllerBasics) TestNavList(ctx context.Context, req *basics.TestNavListReq) (res *basics.TestNavListRes, err error) {
	res = new(basics.TestNavListRes)
	if res.List, err = service.BasicsTestNav().ApiNavAll(ctx, &input_basics.PmsTestNavAllInp{
		Status: 1,
	}); err != nil {
		return
	}
	return
}
