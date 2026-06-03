package admin

import (
	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
)

// CooperateTypeList 查看按摩营业类型列表
func (c *ControllerSpa) CooperateTypeList(ctx context.Context, req *spa.CooperateTypeListReq) (res *spa.CooperateTypeListRes, err error) {
	list, totalCount, err := service.SpaCooperateType().List(ctx, &req.SpaCooperateTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaCooperateTypeListModel{}
	}

	res = new(spa.CooperateTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// CooperateTypeExport 导出按摩营业类型列表
func (c *ControllerSpa) CooperateTypeExport(ctx context.Context, req *spa.CooperateTypeExportReq) (res *spa.CooperateTypeExportRes, err error) {
	err = service.SpaCooperateType().Export(ctx, &req.SpaCooperateTypeListInp)
	return
}

// CooperateTypeEdit 更新按摩营业类型
func (c *ControllerSpa) CooperateTypeEdit(ctx context.Context, req *spa.CooperateTypeEditReq) (res *spa.CooperateTypeEditRes, err error) {
	err = service.SpaCooperateType().Edit(ctx, &req.SpaCooperateTypeEditInp)
	return
}

// CooperateTypeView 获取指定按摩营业类型信息
func (c *ControllerSpa) CooperateTypeView(ctx context.Context, req *spa.CooperateTypeViewReq) (res *spa.CooperateTypeViewRes, err error) {
	data, err := service.SpaCooperateType().View(ctx, &req.SpaCooperateTypeViewInp)
	if err != nil {
		return
	}

	res = new(spa.CooperateTypeViewRes)
	res.SpaCooperateTypeViewModel = data
	return
}

// CooperateTypeDelete 删除按摩营业类型
func (c *ControllerSpa) CooperateTypeDelete(ctx context.Context, req *spa.CooperateTypeDeleteReq) (res *spa.CooperateTypeDeleteRes, err error) {
	err = service.SpaCooperateType().Delete(ctx, &req.SpaCooperateTypeDeleteInp)
	return
}

// CooperateTypeStatus 更新按摩营业类型状态
func (c *ControllerSpa) CooperateTypeStatus(ctx context.Context, req *spa.CooperateTypeStatusReq) (res *spa.CooperateTypeStatusRes, err error) {
	err = service.SpaCooperateType().Status(ctx, &req.SpaCooperateTypeStatusInp)
	return
}
