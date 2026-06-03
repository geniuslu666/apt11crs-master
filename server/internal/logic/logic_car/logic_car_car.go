package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarCar struct{}

func NewCarCar() *sCarCar {
	return &sCarCar{}
}

func init() {
	service.RegisterCarCar(NewCarCar())
}

// Model 车辆ORM模型
func (s *sCarCar) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarCar.Ctx(ctx), option...)
}

// List 获取车辆列表
func (s *sCarCar) List(ctx context.Context, in *input_car.CarCarListInp) (list []*input_car.CarCarListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarCarListModel{})

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.CarCar.Columns().Sort).OrderDesc(dao.CarCar.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取车辆列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取车辆列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Edit 修改/新增车辆
func (s *sCarCar) Edit(ctx context.Context, in *input_car.CarCarEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarCarUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改车辆失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarCarInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增车辆失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除车辆
func (s *sCarCar) Delete(ctx context.Context, in *input_car.CarCarDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除车辆失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取车辆最大排序
func (s *sCarCar) MaxSort(ctx context.Context, in *input_car.CarCarMaxSortInp) (res *input_car.CarCarMaxSortModel, err error) {
	if err = dao.CarCar.Ctx(ctx).Fields(dao.CarCar.Columns().Sort).OrderDesc(dao.CarCar.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取车辆最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_car.CarCarMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取车辆指定信息
func (s *sCarCar) View(ctx context.Context, in *input_car.CarCarViewInp) (res *input_car.CarCarViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取车辆信息，请稍后重试！")
		return
	}

	return
}

// Status 更新车辆状态
func (s *sCarCar) Status(ctx context.Context, in *input_car.CarCarStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarCar.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新车辆状态失败，请稍后重试！")
		return
	}
	return
}

// WorkStatus 更新工作状态
func (s *sCarCar) WorkStatus(ctx context.Context, in *input_car.CarCarWorkStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.CarCar.Columns().WorkStatus: in.WorkStatus,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新车辆工作状态失败，请稍后重试！")
			return
		}

		return
	})
}
