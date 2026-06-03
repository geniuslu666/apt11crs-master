package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sFoodRestaurantNotice struct{}

func NewFoodRestaurantNotice() *sFoodRestaurantNotice {
	return &sFoodRestaurantNotice{}
}

func init() {
	service.RegisterFoodRestaurantNotice(NewFoodRestaurantNotice())
}

func (s *sFoodRestaurantNotice) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodRestaurantNotice.Ctx(ctx), option...)
}

func (s *sFoodRestaurantNotice) List(ctx context.Context, in *input_food.FoodRestaurantNoticeListInp) (list []*input_food.FoodRestaurantNoticeListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.Fields(input_food.FoodRestaurantNoticeListModel{})

	if in.RestaurantId > 0 {
		mod = mod.Where(dao.FoodRestaurantNotice.Columns().RestaurantId, in.RestaurantId)
	}

	if in.Title != "" {
		mod = mod.WhereLike(dao.FoodRestaurantNotice.Columns().Title, "%"+in.Title+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodRestaurantNotice.Columns().Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FoodRestaurantNotice.Columns().CreateAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.FoodRestaurantNotice.Table() + "." + dao.FoodRestaurantNotice.Columns().Sort).OrderDesc(dao.FoodRestaurantNotice.Table() + "." + dao.FoodRestaurantNotice.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取通知管理列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodRestaurantNotice) Export(ctx context.Context, in *input_food.FoodRestaurantNoticeListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_food.FoodRestaurantNoticeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出通知管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("通知管理")
		exports   []input_food.FoodRestaurantNoticeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sFoodRestaurantNotice) Edit(ctx context.Context, in *input_food.FoodRestaurantNoticeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodRestaurantNoticeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改通知管理失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodRestaurantNoticeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增通知管理失败，请稍后重试！")
		}
		return
	})
}

func (s *sFoodRestaurantNotice) Delete(ctx context.Context, in *input_food.FoodRestaurantNoticeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除通知管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodRestaurantNotice) MaxSort(ctx context.Context, in *input_food.FoodRestaurantNoticeMaxSortInp) (res *input_food.FoodRestaurantNoticeMaxSortModel, err error) {
	if err = dao.FoodRestaurantNotice.Ctx(ctx).Fields(dao.FoodRestaurantNotice.Columns().Sort).OrderDesc(dao.FoodRestaurantNotice.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取通知管理最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_food.FoodRestaurantNoticeMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sFoodRestaurantNotice) View(ctx context.Context, in *input_food.FoodRestaurantNoticeViewInp) (res *input_food.FoodRestaurantNoticeViewModel, err error) {
	if err = s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取通知管理信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodRestaurantNotice) Status(ctx context.Context, in *input_food.FoodRestaurantNoticeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodRestaurantNotice.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新通知管理状态失败，请稍后重试！")
		return
	}
	return
}
