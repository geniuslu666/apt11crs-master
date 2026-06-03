package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) CooperateTypeList(ctx context.Context, req *car.CooperateTypeListReq) (res *car.CooperateTypeListRes, err error) {
	list, totalCount, err := service.CarCooperateType().List(ctx, &req.CarCooperateTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarCooperateTypeListModel{}
	}

	res = new(car.CooperateTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) CooperateTypeExport(ctx context.Context, req *car.CooperateTypeExportReq) (res *car.CooperateTypeExportRes, err error) {
	err = service.CarCooperateType().Export(ctx, &req.CarCooperateTypeListInp)
	return
}
func (c *ControllerCar) CooperateTypeView(ctx context.Context, req *car.CooperateTypeViewReq) (res *car.CooperateTypeViewRes, err error) {
	data, err := service.CarCooperateType().View(ctx, &req.CarCooperateTypeViewInp)
	if err != nil {
		return
	}

	res = new(car.CooperateTypeViewRes)
	res.CarCooperateTypeViewModel = data
	return
}
func (c *ControllerCar) CooperateTypeEdit(ctx context.Context, req *car.CooperateTypeEditReq) (res *car.CooperateTypeEditRes, err error) {
	err = service.CarCooperateType().Edit(ctx, &req.CarCooperateTypeEditInp)
	return
}
func (c *ControllerCar) CooperateTypeDelete(ctx context.Context, req *car.CooperateTypeDeleteReq) (res *car.CooperateTypeDeleteRes, err error) {
	err = service.CarCooperateType().Delete(ctx, &req.CarCooperateTypeDeleteInp)
	return
}
func (c *ControllerCar) CooperateTypeStatus(ctx context.Context, req *car.CooperateTypeStatusReq) (res *car.CooperateTypeStatusRes, err error) {
	err = service.CarCooperateType().Status(ctx, &req.CarCooperateTypeStatusInp)
	return
}
