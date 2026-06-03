package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) DictDataEdit(ctx context.Context, req *basics.DictDataEditReq) (res *basics.DictDataEditRes, err error) {
	err = service.BasicsDictData().Edit(ctx, &req.DictDataEditInp)
	return
}
func (c *ControllerBasics) DictDataDelete(ctx context.Context, req *basics.DictDataDeleteReq) (res *basics.DictDataDeleteRes, err error) {
	err = service.BasicsDictData().Delete(ctx, &req.DictDataDeleteInp)
	return
}
func (c *ControllerBasics) DictDataList(ctx context.Context, req *basics.DictDataListReq) (res *basics.DictDataListRes, err error) {
	list, totalCount, err := service.BasicsDictData().List(ctx, &req.DictDataListInp)
	if err != nil {
		return
	}

	res = new(basics.DictDataListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) DictDataSelect(ctx context.Context, req *basics.DictDataSelectReq) (res *basics.DictDataSelectRes, err error) {
	var (
		list input_basics.DataSelectModel
	)
	if list, err = service.BasicsDictData().Select(ctx, &req.DataSelectInp); err != nil {
		return
	}
	result := basics.DictDataSelectRes(list)
	res = &result
	return
}
func (c *ControllerBasics) DictDataSelects(ctx context.Context, req *basics.DictDataSelectsReq) (res *basics.DictDataSelectsRes, err error) {
	var (
		result = make(basics.DictDataSelectsRes)
	)
	res = new(basics.DictDataSelectsRes)
	for _, v := range req.Types {
		option, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: v})
		if err != nil {
			return nil, err
		}
		result[v] = option
	}
	res = &result
	return
}
