package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarDriver struct{}

func NewCarDriver() *sCarDriver {
	return &sCarDriver{}
}

func init() {
	service.RegisterCarDriver(NewCarDriver())
}

// Model 司机管理ORM模型
func (s *sCarDriver) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarDriver.Ctx(ctx), option...)
}

// List 获取司机管理列表
func (s *sCarDriver) List(ctx context.Context, in *input_car.CarDriverListInp) (list []*input_car.CarDriverListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarDriverListModel{})

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.WhereLike(dao.CarDriver.Columns().Name, "%"+in.Keywords+"%").WhereOrLike(dao.CarDriver.Columns().Nickname, "%"+in.Keywords+"%").WhereOrLike(dao.CarDriver.Columns().Phone, "%"+in.Keywords+"%")
	}

	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.CarDriver.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.CarDriver.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询创建时间
	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.CarDriver.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	if !g.IsEmpty(in.OrderType) {
		if in.OrderType == "choose" {
			// 排序
			mod = mod.OrderDesc(dao.CarDriver.Columns().WorkStatus).OrderDesc(dao.CarDriver.Columns().Id)
		} else {
			// 排序
			mod = mod.OrderDesc(dao.CarDriver.Columns().Id)
		}
	} else {
		// 排序
		mod = mod.OrderDesc(dao.CarDriver.Columns().Id)
	}

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取司机列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取司机列表失败，请稍后重试！")
			return
		}
	}
	for k, v := range list {
		if v.WorkStatus == "REST" {
			// 休息
			list[k].WorkStatusEnum = 1
		} else if v.WorkStatus == "WORKING" {
			workingCount, _ := dao.CarOrder.Ctx(ctx).
				Where(dao.CarOrder.Columns().DriverId, v.Id).
				Where(dao.CarOrder.Columns().OrderStatus, "SERVING").Count()
			if workingCount > 0 {
				// 服务中
				list[k].WorkStatusEnum = 2
			} else {
				// 待服务
				list[k].WorkStatusEnum = 3
			}
		}
	}

	return
}

// All 获取司机
func (s *sCarDriver) All(ctx context.Context, in *input_car.CarDriverListInp) (list []*input_car.CarDriverAllListModel, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarDriverAllListModel{})

	// 正常状态
	mod = mod.Where(dao.CarDriver.Columns().Status, 1)

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.CarDriver.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.Where(dao.CarDriver.Columns().Name, in.Keywords).WhereOr(dao.CarDriver.Columns().Phone, in.Keywords)
	}

	// 排序
	mod = mod.OrderDesc(dao.CarDriver.Columns().Id)

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取司机列表失败，请稍后重试！")
		return
	}

	return
}

// Edit 修改/新增司机管理
func (s *sCarDriver) Edit(ctx context.Context, in *input_car.CarDriverEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			mod := s.Model(ctx)

			if in.Type == "basic" {

				mod = mod.Fields(input_car.CarDriverBasicUpdateFields{})
			} else if in.Type == "qualification" {
				mod = mod.Fields(input_car.CarDriverQualificationUpdateFields{})
			} else if in.Type == "settle" {
				mod = mod.Fields(input_car.CarDriverSettleUpdateFields{})
			} else {
				mod = mod.Fields(input_car.CarDriverUpdateFields{})
			}

			if _, err = mod.
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改司机管理失败，请稍后重试！")
			}

			return
		}

		// 新增
		var (
			lastInsertId int64
		)
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarDriverInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增司机失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		return
	})
}

// Delete 删除司机管理
func (s *sCarDriver) Delete(ctx context.Context, in *input_car.CarDriverDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除司机管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取司机管理指定信息
func (s *sCarDriver) View(ctx context.Context, in *input_car.CarDriverViewInp) (res *input_car.CarDriverViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取司机管理信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取司机管理信息，请稍后重试！")
			return
		}
	}
	return
}

// Status 更新司机状态
func (s *sCarDriver) Status(ctx context.Context, in *input_car.CarDriverStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.CarDriver.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新司机状态失败，请稍后重试！")
			return
		}

		return
	})
}

// WorkStatus 更新司机状态
func (s *sCarDriver) WorkStatus(ctx context.Context, in *input_car.CarDriverWorkStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.CarDriver.Columns().WorkStatus: in.WorkStatus,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新司机工作状态失败，请稍后重试！")
			return
		}

		return
	})
}

// GetIds 获取获取司机的id
func (s *sCarDriver) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}

// Bind 绑定用户
func (s *sCarDriver) Bind(ctx context.Context, in *input_car.CarDriverBindInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 判断会员是否已绑定
		memberCount, err := dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().MemberId, in.MemberId).Count()
		if memberCount > 0 {
			err = gerror.New("该用户已被绑定，请重新选择用户")
			return
		}

		if _, err = dao.CarDriver.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.CarDriver.Columns().MemberId: in.MemberId,
			dao.CarDriver.Columns().IsLeader: in.IsLeader,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

// Unbind 解绑用户
func (s *sCarDriver) Unbind(ctx context.Context, in *input_car.CarDriverUnbindInp) (err error) {

	if _, err = dao.CarDriver.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarDriver.Columns().MemberId: 0,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}

// BindCar 绑定车辆
func (s *sCarDriver) BindCar(ctx context.Context, in *input_car.CarDriverBindCarInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = dao.CarDriver.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.CarDriver.Columns().CarId: in.CarId,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定车辆失败，请稍后重试！")
		}
		return
	})
}
