package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) HelpList(ctx context.Context, req *car.HelpListReq) (res *car.HelpListRes, err error) {
	list, totalCount, err := service.CarHelp().List(ctx, &req.CarHelpListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarHelpListModel{}
	}

	res = new(car.HelpListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) HelpView(ctx context.Context, req *car.HelpViewReq) (res *car.HelpViewRes, err error) {
	data, err := service.CarHelp().View(ctx, &req.CarHelpViewInp)
	if err != nil {
		return
	}

	res = new(car.HelpViewRes)
	res.CarHelpViewModel = data
	return
}
func (c *ControllerCar) HelpEdit(ctx context.Context, req *car.HelpEditReq) (res *car.HelpEditRes, err error) {
	err = service.CarHelp().Edit(ctx, &req.CarHelpEditInp)
	return
}
func (c *ControllerCar) HelpDelete(ctx context.Context, req *car.HelpDeleteReq) (res *car.HelpDeleteRes, err error) {
	err = service.CarHelp().Delete(ctx, &req.CarHelpDeleteInp)
	return
}
func (c *ControllerCar) HelpLanguageList(ctx context.Context, req *car.HelpLanguageListReq) (res *car.HelpLanguageListRes, err error) {
	list, err := service.CarHelp().LanguageList(ctx, &req.CarHelpLanguageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarHelpLanguageListModel{}
	}

	res = new(car.HelpLanguageListRes)
	res.List = list
	return
}
