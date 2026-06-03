package admin

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"reflect"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
)

func (c *ControllerPms) PropertyLanguage(ctx context.Context, req *pms.PropertyLanguageReq) (res *pms.PropertyLanguageRes, err error) {
	var (
		ormPmsProperty = g.Model(dao.PmsProperty.Table()).Ctx(ctx).Safe()
		Object         gdb.Record
		PmsPropertyDao *input_language.LoadLanguage
		LanguageStruct input_language.LanguageModel
	)
	LanguageStructValue := reflect.ValueOf(req.Data)
	LanguageStructType := reflect.TypeOf(req.Data)
	for i := 0; i < LanguageStructType.NumField(); i++ {
		if LanguageStructType.Field(i).Name == req.Type {
			if req.Id > 0 {
				Uuid := guid.S([]byte(req.Type))
				// 更新数据
				if Object, err = dao.PmsProperty.Ctx(ctx).Where(dao.PmsProperty.Columns().Id, req.Id).One(); err != nil {
					return
				}
				if Object.IsEmpty() {
					err = gerror.New("当前物业不存在")
					return
				}
				Name := LanguageStructType.Field(i).Tag.Get("json")
				// 是否存在languageUuid
				if !g.IsEmpty(Object[gstr.ToLower(Name)]) {
					Uuid = Object[Name].String()
				} else {
					if _, err = ormPmsProperty.Where(dao.PmsProperty.Columns().Id, req.Id).Data(g.MapStrAny{
						gstr.ToLower(req.Type):              Uuid,
						dao.PmsProperty.Columns().UpdatedAt: gtime.Now().Format("Y-m-d H:i:s"),
					}).Update(); err != nil {
						return
					}
				}
				PmsPropertyDao = &input_language.LoadLanguage{
					Uuid: Uuid,
					Tag:  dao.PmsProperty.Table(),
					Type: "table",
					Key:  gstr.CaseSnakeFirstUpper(req.Type),
				}
				if LanguageStructValue.Field(i).IsValid() {
					LanguageStruct = LanguageStructValue.Field(i).Interface().(input_language.LanguageModel)
				} else {
					err = gerror.New("数据错误")
				}
				if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, PmsPropertyDao); err != nil {
					return
				}
			}
		}
	}
	if !g.IsEmpty(req.RegionId) {
		if _, err = g.Model(dao.PmsProperty.Table()).Ctx(ctx).Safe().Where(dao.PmsProperty.Columns().Id, req.Id).Data(g.MapStrAny{
			dao.PmsProperty.Columns().RegionId:  req.RegionId,
			dao.PmsProperty.Columns().UpdatedAt: gtime.Now().Format("Y-m-d H:i:s"),
		}).Update(); err != nil {
			return
		}
	}
	return
}
func (c *ControllerPms) PropertyList(ctx context.Context, req *pms.PropertyListReq) (res *pms.PropertyListRes, err error) {
	list, totalCount, err := service.HotelService().PropertyList(ctx, &req.PmsPropertyListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsPropertyListModel{}
	}

	res = new(pms.PropertyListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) PropertyAll(ctx context.Context, req *pms.PropertyAllReq) (res *pms.PropertyAllRes, err error) {
	list, err := service.HotelService().PropertyAll(ctx, &req.PmsPropertyAllInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsPropertyAllModel{}
	}

	res = new(pms.PropertyAllRes)
	res.List = list
	return
}
func (c *ControllerPms) PropertyExport(ctx context.Context, req *pms.PropertyExportReq) (res *pms.PropertyExportRes, err error) {
	err = service.HotelService().PropertyExport(ctx, &req.PmsPropertyListInp)
	return
}
func (c *ControllerPms) PropertyView(ctx context.Context, req *pms.PropertyViewReq) (res *pms.PropertyViewRes, err error) {
	req.IsLanguage = true
	_, data, err := service.HotelService().PropertyView(ctx, &req.PmsPropertyViewInp)
	if err != nil {
		return
	}

	res = new(pms.PropertyViewRes)
	res.PmsPropertyViewLannguageModel = data
	return
}
func (c *ControllerPms) PropertyEdit(ctx context.Context, req *pms.PropertyEditReq) (res *pms.PropertyEditRes, err error) {
	err = service.HotelService().PropertyEdit(ctx, &req.PmsPropertyEditInp)
	return
}
func (c *ControllerPms) PropertyEditGroup(ctx context.Context, req *pms.PropertyEditGroupReq) (res *pms.PropertyEditGroupRes, err error) {
	err = service.HotelService().PropertyEditGroup(ctx, &req.PmsPropertyEditGroupInp)
	return
}
func (c *ControllerPms) PropertyDelete(ctx context.Context, req *pms.PropertyDeleteReq) (res *pms.PropertyDeleteRes, err error) {
	err = service.HotelService().PropertyDelete(ctx, &req.PmsPropertyDeleteInp)
	return
}
func (c *ControllerPms) PropertySetGalleryCover(ctx context.Context, req *pms.PropertySetGalleryCoverReq) (res *pms.PropertySetGalleryCoverRes, err error) {
	var (
		ormPmsProperty = g.Model(dao.PmsProperty.Table()).Ctx(ctx).Safe()
		result         sql.Result
		updateRow      int64
	)

	if result, err = ormPmsProperty.Where(dao.PmsProperty.Columns().Id, req.Uid).OmitEmptyData().Data(g.MapStrAny{
		dao.PmsProperty.Columns().GalleryCover:  req.GalleryCover,
		dao.PmsProperty.Columns().GalleryImages: req.GalleryImages,
		dao.PmsProperty.Columns().UpdatedAt:     gtime.Now().Format("Y-m-d H:i:s"),
	}).Update(); err != nil {
		return
	}

	if updateRow, err = result.RowsAffected(); err != nil {
		return
	}
	if updateRow != 1 {
		err = gerror.New("更新失败")
		return
	}
	return
}
func (c *ControllerPms) PropertyGetGallery(ctx context.Context, req *pms.PropertyGetGalleryReq) (res *pms.PropertyGetGalleryRes, err error) {
	var (
		ormPmsProperty = g.Model(dao.PmsProperty.Table()).Ctx(ctx).Safe()
	)
	res = new(pms.PropertyGetGalleryRes)
	if err = ormPmsProperty.Where(dao.PmsProperty.Columns().Id, req.Uid).Scan(&res); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	return
}
func (c *ControllerPms) PropertySort(ctx context.Context, req *pms.PropertySortReq) (res *pms.PropertySortRes, err error) {
	err = service.HotelService().PropertySort(ctx, &req.PmsPropertySortInp)
	return
}
func (c *ControllerPms) PropertyRecycle(ctx context.Context, req *pms.PropertyRecycleReq) (res *pms.PropertyRecycleRes, err error) {
	err = service.HotelService().PropertyRecycle(ctx, &req.PmsPropertyRecycleInp)
	return
}
func (c *ControllerPms) PropertySwitchSpaCanOrder(ctx context.Context, req *pms.PropertySwitchSpaCanOrderReq) (res *pms.PropertySwitchSpaCanOrderRes, err error) {
	err = service.HotelService().SwitchSpaCanOrder(ctx, &req.PropertySwitchSpaCanOrderInp)
	return
}
func (c *ControllerPms) PropertyUpdateAccessPass(ctx context.Context, req *pms.PropertyUpdateAccessPassReq) (res *pms.PropertyUpdateAccessPassRes, err error) {
	err = service.HotelService().UpdateAccessPass(ctx, &req.PropertyUpdateAccessPassInp)
	return
}
