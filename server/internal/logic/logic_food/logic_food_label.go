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

type sFoodLabel struct{}

func NewFoodLabel() *sFoodLabel {
	return &sFoodLabel{}
}

func init() {
	service.RegisterFoodLabel(NewFoodLabel())
}

func (s *sFoodLabel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodLabel.Ctx(ctx), option...)
}

func (s *sFoodLabel) List(ctx context.Context, in *input_food.FoodLabelListInp) (list []*input_food.FoodLabelListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodLabelListModel{})

	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.FoodLabel.Columns().Name, "%"+in.Name+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodLabel.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.Type) {
		mod = mod.Where(dao.FoodLabel.Columns().Type, in.Type)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderAsc(dao.FoodLabel.Columns().Sort).OrderDesc(dao.FoodLabel.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅标签列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅标签列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodLabel) Edit(ctx context.Context, in *input_food.FoodLabelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object         gdb.Record
			Dao            *input_language.LoadLanguage
			LanguageStruct input_language.LanguageModel
		)
		Uuid := guid.S([]byte("name"))

		if in.Id > 0 {
			if Object, err = dao.FoodLabel.Ctx(ctx).Where(dao.FoodLabel.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("CuisineName")]) {
				Uuid = Object["cuisine_name"].String()
			}
			Dao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.FoodLabel.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
				return
			}

			in.Name = Uuid

			if _, err = s.Model(ctx).
				Fields(input_food.FoodLabelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改餐厅标签失败，请稍后重试！")
			}
			return
		}

		in.Name = Uuid
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodLabelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增餐厅标签失败，请稍后重试！")
		}

		Dao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.FoodLabel.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
			return
		}

		return
	})
}

func (s *sFoodLabel) Delete(ctx context.Context, in *input_food.FoodLabelDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除餐厅标签失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodLabel) MaxSort(ctx context.Context, in *input_food.FoodLabelMaxSortInp) (res *input_food.FoodLabelMaxSortModel, err error) {
	if err = dao.FoodLabel.Ctx(ctx).Fields(dao.FoodLabel.Columns().Sort).OrderDesc(dao.FoodLabel.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取餐厅标签最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_food.FoodLabelMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sFoodLabel) View(ctx context.Context, in *input_food.FoodLabelViewInp) (res *input_food.FoodLabelViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取餐厅标签信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodLabel) Status(ctx context.Context, in *input_food.FoodLabelStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodLabel.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新餐厅标签状态失败，请稍后重试！")
		return
	}
	return
}
