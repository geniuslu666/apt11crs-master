package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarMaintenance struct{}

func NewCarMaintenance() *sCarMaintenance {
	return &sCarMaintenance{}
}

func init() {
	service.RegisterCarMaintenance(NewCarMaintenance())
}

// Model MaintenanceORM模型
func (s *sCarMaintenance) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarMaintenance.Ctx(ctx), option...)
}

// List 获取Maintenance列表
func (s *sCarMaintenance) List(ctx context.Context, in *input_car.CarMaintenanceListInp) (list []*input_car.CarMaintenanceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarMaintenanceListModel{})

	if in.Type > 0 {
		mod = mod.Where(dao.CarMaintenance.Columns().Type, in.Type)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.CarMaintenance.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取Maintenance列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增Maintenance
func (s *sCarMaintenance) Edit(ctx context.Context, in *input_car.CarMaintenanceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.IsDefault == 1 {
			// 找出其他默认的设置为否
			if _, err = s.Model(ctx).
				Where(dao.CarMaintenance.Columns().Type, in.Type).
				Where(dao.CarMaintenance.Columns().IsDefault, 1).Data(g.Map{
				dao.CarMaintenance.Columns().IsDefault: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "修改Maintenance失败，请稍后重试！")
				return
			}
		}

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarMaintenanceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改Maintenance失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarMaintenanceInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增Maintenance失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除Maintenance
func (s *sCarMaintenance) Delete(ctx context.Context, in *input_car.CarMaintenanceDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除Maintenance失败，请稍后重试！")
		return
	}
	return
}

// View 获取Maintenance指定信息
func (s *sCarMaintenance) View(ctx context.Context, in *input_car.CarMaintenanceViewInp) (res *input_car.CarMaintenanceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Maintenance信息，请稍后重试！")
		return
	}
	return
}

// LanguageList 获取MaintenanceLanguage列表
func (s *sCarMaintenance) LanguageList(ctx context.Context, in *input_car.CarMaintenanceLanguageListInp) (list []*input_car.CarMaintenanceLanguageListModel, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarMaintenanceLanguageListModel{})

	if in.Type > 0 {
		mod = mod.Where(dao.CarMaintenance.Columns().Type, in.Type)
	}

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取MaintenanceLanguage列表失败，请稍后重试！")
		return
	}
	return
}
