package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_food"
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

type sFoodCuisine struct{}

func NewFoodCuisine() *sFoodCuisine {
	return &sFoodCuisine{}
}

func init() {
	service.RegisterFoodCuisine(NewFoodCuisine())
}

func (s *sFoodCuisine) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodCuisine.Ctx(ctx), option...)
}

func (s *sFoodCuisine) List(ctx context.Context, in *input_food.FoodCuisineListInp) (list []*input_food.FoodCuisineListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_food.FoodCuisineListModel{})

	if !g.IsEmpty(in.CuisineName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.CuisineName, "cuisine_name")
		if err == nil {
			cuisineIds, _ := service.FoodCuisine().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.FoodCuisine.Columns().Id, cuisineIds)
		}
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodCuisine.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderAsc(dao.FoodCuisine.Columns().Sort).OrderDesc(dao.FoodCuisine.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅菜系列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅菜系列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodCuisine) Edit(ctx context.Context, in *input_food.FoodCuisineEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object         gdb.Record
			Dao            *input_language.LoadLanguage
			LanguageStruct input_language.LanguageModel
		)
		Uuid := guid.S([]byte("name"))

		if in.Id > 0 {
			if Object, err = dao.FoodCuisine.Ctx(ctx).Where(dao.FoodCuisine.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("CuisineName")]) {
				Uuid = Object["cuisine_name"].String()
			}
			Dao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.FoodCuisine.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("CuisineName"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
				return
			}

			in.CuisineName = Uuid

			if _, err = s.Model(ctx).
				Fields(input_food.FoodCuisineUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改餐厅菜系失败，请稍后重试！")
			}
			return
		}

		in.CuisineName = Uuid

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodCuisineInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增餐厅菜系失败，请稍后重试！")
		}

		Dao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.FoodCuisine.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("CuisineName"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
			return
		}
		return
	})
}

func (s *sFoodCuisine) Delete(ctx context.Context, in *input_food.FoodCuisineDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除餐厅菜系失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodCuisine) View(ctx context.Context, in *input_food.FoodCuisineViewInp) (res *input_food.FoodCuisineViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取餐厅菜系信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取餐厅菜系信息，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodCuisine) MaxSort(ctx context.Context, in *input_food.FoodCuisineMaxSortInp) (res *input_food.FoodCuisineMaxSortModel, err error) {
	if err = dao.FoodCuisine.Ctx(ctx).Fields(dao.FoodCuisine.Columns().Sort).OrderDesc(dao.FoodCuisine.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取餐厅菜系最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_food.FoodCuisineMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sFoodCuisine) Status(ctx context.Context, in *input_food.FoodCuisineStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodCuisine.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新餐厅菜系状态失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodCuisine) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("cuisine_name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}
