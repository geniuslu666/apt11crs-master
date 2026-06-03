package admin

import (
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) LanguageList(ctx context.Context, req *pms.LanguageListReq) (res *pms.LanguageListRes, err error) {
	list, totalCount, err := service.BasicsLanguage().List(ctx, &req.PmsLanguageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_language.PmsLanguageListModel{}
	}

	res = new(pms.LanguageListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) LanguageExport(ctx context.Context, req *pms.LanguageExportReq) (res *pms.LanguageExportRes, err error) {
	err = service.BasicsLanguage().Export(ctx, &req.PmsLanguageListInp)
	return
}
func (c *ControllerPms) LanguageView(ctx context.Context, req *pms.LanguageViewReq) (res *pms.LanguageViewRes, err error) {
	data, err := service.BasicsLanguage().View(ctx, &req.PmsLanguageViewInp)
	if err != nil {
		return
	}

	res = new(pms.LanguageViewRes)
	res.PmsLanguageViewModel = data
	return
}
func (c *ControllerPms) LanguageEdit(ctx context.Context, req *pms.LanguageEditReq) (res *pms.LanguageEditRes, err error) {
	err = service.BasicsLanguage().Edit(ctx, &req.PmsLanguageEditInp)
	return
}
func (c *ControllerPms) LanguageDelete(ctx context.Context, req *pms.LanguageDeleteReq) (res *pms.LanguageDeleteRes, err error) {
	err = service.BasicsLanguage().Delete(ctx, &req.PmsLanguageDeleteInp)
	return
}
