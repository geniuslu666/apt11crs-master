package admin

import (
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) PostEdit(ctx context.Context, req *basics.PostEditReq) (res *basics.PostEditRes, err error) {
	err = service.BasicsAdminPost().Edit(ctx, &req.PostEditInp)
	return
}
func (c *ControllerBasics) PostDelete(ctx context.Context, req *basics.PostDeleteReq) (res *basics.PostDeleteRes, err error) {
	err = service.BasicsAdminPost().Delete(ctx, &req.PostDeleteInp)
	return
}
func (c *ControllerBasics) PostMaxSort(ctx context.Context, req *basics.PostMaxSortReq) (res *basics.PostMaxSortRes, err error) {
	res = new(basics.PostMaxSortRes)
	res.PostMaxSortModel, err = service.BasicsAdminPost().MaxSort(ctx, &req.PostMaxSortInp)
	return
}
func (c *ControllerBasics) PostList(ctx context.Context, req *basics.PostListReq) (res *basics.PostListRes, err error) {
	list, totalCount, err := service.BasicsAdminPost().List(ctx, &req.PostListInp)
	if err != nil {
		return
	}

	res = new(basics.PostListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) PostView(ctx context.Context, req *basics.PostViewReq) (res *basics.PostViewRes, err error) {
	data, err := service.BasicsAdminPost().View(ctx, &req.PostViewInp)
	if err != nil {
		return
	}

	res = new(basics.PostViewRes)
	res.PostViewModel = data
	return
}
func (c *ControllerBasics) PostStatus(ctx context.Context, req *basics.PostStatusReq) (res *basics.PostStatusRes, err error) {
	err = service.BasicsAdminPost().Status(ctx, &req.PostStatusInp)
	return
}
