package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) StoreList(ctx context.Context, req *spa.StoreListReq) (res *spa.StoreListRes, err error) {
	list, totalCount, err := service.SpaStore().List(ctx, &req.SpaStoreListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaStoreListModel{}
	}

	res = new(spa.StoreListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) StoreView(ctx context.Context, req *spa.StoreViewReq) (res *spa.StoreViewRes, err error) {
	data, err := service.SpaStore().View(ctx, &req.SpaStoreViewInp)
	if err != nil {
		return
	}

	res = new(spa.StoreViewRes)
	res.SpaStoreViewModel = data
	return
}
func (c *ControllerSpa) StoreEdit(ctx context.Context, req *spa.StoreEditReq) (res *spa.StoreEditRes, err error) {
	err = service.SpaStore().Edit(ctx, &req.SpaStoreEditInp)
	return
}
func (c *ControllerSpa) StoreDelete(ctx context.Context, req *spa.StoreDeleteReq) (res *spa.StoreDeleteRes, err error) {
	err = service.SpaStore().Delete(ctx, &req.SpaStoreDeleteInp)
	return
}
func (c *ControllerSpa) StoreLatest(ctx context.Context, req *spa.StoreLatestReq) (res *spa.StoreLatestRes, err error) {
	data, err := service.SpaStore().Latest(ctx)
	if err != nil {
		return
	}

	res = new(spa.StoreLatestRes)
	res.SpaStoreViewModel = data
	return
}
