package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarCarType struct{}

func NewCarCarType() *sCarCarType {
	return &sCarCarType{}
}

func init() {
	service.RegisterCarCarType(NewCarCarType())
}

// Model 车辆车型ORM模型
func (s *sCarCarType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarCarType.Ctx(ctx), option...)
}

// List 获取车辆车型列表
func (s *sCarCarType) List(ctx context.Context, in *input_car.CarCarTypeListInp) (list []*input_car.CarCarTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarCarTypeListModel{})

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.CarCarType.Columns().Sort).OrderDesc(dao.CarCarType.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取车辆车型列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取车辆车型列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Edit 修改/新增车辆车型
func (s *sCarCarType) Edit(ctx context.Context, in *input_car.CarCarTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object             gdb.Record
			NameDao            *input_language.LoadLanguage
			NameLanguageStruct input_language.LanguageModel
			DescDao            *input_language.LoadLanguage
			DescLanguageStruct input_language.LanguageModel
		)
		NameUuid := guid.S([]byte("name"))
		DescUuid := guid.S([]byte("desc"))

		// 修改
		if in.Id > 0 {
			if Object, err = dao.CarCarType.Ctx(ctx).Where(dao.CarCarType.Columns().Id, in.Id).One(); err != nil {
				return
			}
			// 名称多语言
			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
				NameUuid = Object["name"].String()
			}
			NameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.SpaService.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			NameLanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
				return
			}

			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("Desc")]) {
				DescUuid = Object["desc"].String()
			}
			DescDao = &input_language.LoadLanguage{
				Uuid: DescUuid,
				Tag:  dao.SpaService.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Desc"),
			}
			DescLanguageStruct = in.DescLanguage
			if err = service.BasicsLanguage().Sync(ctx, DescLanguageStruct, DescDao); err != nil {
				return
			}
			// 修改
			in.Name = NameUuid
			in.Desc = DescUuid

			if _, err = s.Model(ctx).
				Fields(input_car.CarCarTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改车辆车型失败，请稍后重试！")
			}
			return
		}
		in.Name = NameUuid
		in.Desc = DescUuid

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarCarTypeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增车辆车型失败，请稍后重试！")
		}

		// 多语言
		NameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.SpaService.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		NameLanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
			return
		}

		DescDao = &input_language.LoadLanguage{
			Uuid: DescUuid,
			Tag:  dao.SpaService.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Desc"),
		}
		DescLanguageStruct = in.DescLanguage
		if err = service.BasicsLanguage().Sync(ctx, DescLanguageStruct, DescDao); err != nil {
			return
		}
		return
	})
}

// Delete 删除车辆车型
func (s *sCarCarType) Delete(ctx context.Context, in *input_car.CarCarTypeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除车辆车型失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取车辆车型最大排序
func (s *sCarCarType) MaxSort(ctx context.Context, in *input_car.CarCarTypeMaxSortInp) (res *input_car.CarCarTypeMaxSortModel, err error) {
	if err = dao.CarCarType.Ctx(ctx).Fields(dao.CarCarType.Columns().Sort).OrderDesc(dao.CarCarType.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取车辆车型最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_car.CarCarTypeMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取车辆车型指定信息
func (s *sCarCarType) View(ctx context.Context, in *input_car.CarCarTypeViewInp) (res *input_car.CarCarTypeViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取车辆车型信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取车辆车型信息，请稍后重试！")
			return
		}
	}

	return
}

// Status 更新车辆车型状态
func (s *sCarCarType) Status(ctx context.Context, in *input_car.CarCarTypeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarCarType.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新车辆车型状态失败，请稍后重试！")
		return
	}
	return
}
