package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"strconv"
	"strings"
)

type sFoodActivity struct{}

func NewFoodActivity() *sFoodActivity {
	return &sFoodActivity{}
}

func init() {
	service.RegisterFoodActivity(NewFoodActivity())
}

func (s *sFoodActivity) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodActivity.Ctx(ctx), option...)
}

func (s *sFoodActivity) List(ctx context.Context, in *input_food.FoodActivityListInp) (list []*input_food.FoodActivityListModel, totalCount int, err error) {

	mod := s.Model(ctx)
	if !in.IsLanguage {
		mod = mod.Hook(hook2.PmsFindLanguageValueHook)
	} else {
		mod = mod.WithAll()
	}

	mod = mod.Fields(input_food.FoodActivityListModel{})

	if len(in.Date) == 2 {
		mod = mod.WhereBetween(dao.FoodActivity.Columns().Date, in.Date[0], in.Date[1])
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodActivity.Columns().Status, in.Status)
	}

	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.FoodActivity.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.FoodActivity.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取订餐活动表列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodActivity) Export(ctx context.Context, in *input_food.FoodActivityListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_food.FoodActivityExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出订餐活动表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("订餐活动表")
		exports   []input_food.FoodActivityExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sFoodActivity) Edit(ctx context.Context, in *input_food.FoodActivityEditInp) (err error) {
	//活动日期不能重复
	if err = hgorm.IsUnique(ctx, &dao.FoodActivity, g.Map{
		dao.FoodActivity.Columns().Date: in.Date,
	}, "该日期已存在其他活动，同一日期只能存在一场活动", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object                 gdb.Record
			FoodActivityDao        *input_language.LoadLanguage
			LanguageStruct         input_language.LanguageModel
			ActivityRestaurantInfo []*entity.FoodActivityRestaurant
		)
		Uuid := guid.S([]byte("name"))

		if in.Id > 0 {

			if Object, err = dao.FoodActivity.Ctx(ctx).Where(dao.FoodActivity.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object["name"]) {
				Uuid = Object["name"].String()
			}

			FoodActivityDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.FoodActivity.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodActivityDao); err != nil {
				return
			}

			in.Name = Uuid
			if _, err = s.Model(ctx).
				Fields(input_food.FoodActivityUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改订餐活动失败，请稍后重试！")
			}

			if !g.IsEmpty(in.RestaurantIds) {
				for _, Restaurant := range strings.Split(in.RestaurantIds, ",") {
					if !g.IsEmpty(Restaurant) {
						RestaurantId, _ := strconv.ParseInt(Restaurant, 10, 64)
						ActivityRestaurantInfo = append(ActivityRestaurantInfo, &entity.FoodActivityRestaurant{
							ActivityId:   in.Id,
							RestaurantId: RestaurantId,
						})
					}
				}

				if !g.IsEmpty(ActivityRestaurantInfo) {
					if _, err = dao.FoodActivityRestaurant.Ctx(ctx).Where(dao.FoodActivityRestaurant.Columns().ActivityId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理活动餐厅旧数据失败，请稍后重试！")
						return
					}
					if _, err = dao.FoodActivityRestaurant.Ctx(ctx).OmitEmptyData().Insert(ActivityRestaurantInfo); err != nil {
						return err
					}
				} else {
					if _, err = dao.FoodActivityRestaurant.Ctx(ctx).Where(dao.FoodActivityRestaurant.Columns().ActivityId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理活动餐厅旧数据失败，请稍后重试！")
						return
					}
				}
			}

			return
		}

		var (
			lastInsertId int64
		)
		in.Name = Uuid
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodActivityInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增订餐活动失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		if !g.IsEmpty(in.RestaurantIds) {
			for _, Restaurant := range strings.Split(in.RestaurantIds, ",") {
				if !g.IsEmpty(Restaurant) {
					RestaurantId, _ := strconv.ParseInt(Restaurant, 10, 64)
					ActivityRestaurantInfo = append(ActivityRestaurantInfo, &entity.FoodActivityRestaurant{
						ActivityId:   lastInsertId,
						RestaurantId: RestaurantId,
					})
				}
			}
		}
		if !g.IsEmpty(ActivityRestaurantInfo) {
			if _, err = dao.FoodActivityRestaurant.Ctx(ctx).OmitEmptyData().Insert(ActivityRestaurantInfo); err != nil {
				return err
			}
		}

		FoodActivityDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.FoodActivity.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodActivityDao); err != nil {
			return
		}

		return
	})
}

func (s *sFoodActivity) Delete(ctx context.Context, in *input_food.FoodActivityDeleteInp) (err error) {

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		/*if _, err = dao.FoodActivityRestaurant.Ctx(ctx).Where(dao.FoodActivityRestaurant.Columns().ActivityId, in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "清理活动餐厅旧数据失败，请稍后重试！")
			return
		}*/

		if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "删除订餐活动失败，请稍后重试！")
			return
		}
		return
	})
}

func (s *sFoodActivity) View(ctx context.Context, in *input_food.FoodActivityViewInp) (res *input_food.FoodActivityViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取订餐活动表信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodActivity) Status(ctx context.Context, in *input_food.FoodActivityStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodActivity.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新订餐活动表状态失败，请稍后重试！")
		return
	}
	return
}
