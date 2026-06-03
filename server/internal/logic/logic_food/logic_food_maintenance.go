package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodMaintenance struct{}

func NewFoodMaintenance() *sFoodMaintenance {
	return &sFoodMaintenance{}
}

func init() {
	service.RegisterFoodMaintenance(NewFoodMaintenance())
}

// Model MaintenanceORM模型
func (s *sFoodMaintenance) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodMaintenance.Ctx(ctx), option...)
}

// List 获取Maintenance列表
func (s *sFoodMaintenance) List(ctx context.Context, in *input_food.FoodMaintenanceListInp) (list []*input_food.FoodMaintenanceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_food.FoodMaintenanceListModel{})

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.FoodMaintenance.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取Maintenance列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增Maintenance
func (s *sFoodMaintenance) Edit(ctx context.Context, in *input_food.FoodMaintenanceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.IsDefault == 1 {
			// 找出其他默认的设置为否
			if _, err = s.Model(ctx).
				Where(dao.FoodMaintenance.Columns().IsDefault, 1).Data(g.Map{
				dao.FoodMaintenance.Columns().IsDefault: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "修改Maintenance失败，请稍后重试！")
			}
		}

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodMaintenanceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改Maintenance失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodMaintenanceInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增Maintenance失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除Maintenance
func (s *sFoodMaintenance) Delete(ctx context.Context, in *input_food.FoodMaintenanceDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除Maintenance失败，请稍后重试！")
		return
	}
	return
}

// View 获取Maintenance指定信息
func (s *sFoodMaintenance) View(ctx context.Context, in *input_food.FoodMaintenanceViewInp) (res *input_food.FoodMaintenanceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Maintenance信息，请稍后重试！")
		return
	}
	return
}

// LanguageList 获取MaintenanceLanguage列表
func (s *sFoodMaintenance) LanguageList(ctx context.Context, in *input_food.FoodMaintenanceLanguageListInp) (list []*input_food.FoodMaintenanceLanguageListModel, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_food.FoodMaintenanceLanguageListModel{})

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取MaintenanceLanguage列表失败，请稍后重试！")
		return
	}
	return
}
