package app

import (
	"APT/api/app/basics"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) AdSpaceAfterPayBannerList(ctx context.Context, req *basics.AdSpaceAfterPayBannerListReq) (res *basics.AdSpaceAfterPayBannerListRes, err error) {
	res = new(basics.AdSpaceAfterPayBannerListRes)
	if res.List, _, err = service.BasicsBanner().BannerAppList(ctx, &input_basics.PmsBannerListInp{
		BannerName: "ad_space_after_pay",
		Language:   contexts.GetLanguage(ctx),
		PageReq: input_form.PageReq{
			Pagination: false,
		}}); err != nil {
		return
	}
	return
}
