package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaIsp struct{}

func NewSpaIsp() *sSpaIsp {
	return &sSpaIsp{}
}

func init() {
	service.RegisterSpaIsp(NewSpaIsp())
}

// Model 服务商管理ORM模型
func (s *sSpaIsp) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaIsp.Ctx(ctx), option...)
}

// List 获取服务商管理列表
func (s *sSpaIsp) List(ctx context.Context, in *input_spa.SpaIspListInp) (list []*input_spa.SpaIspListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaIspListModel{})

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.WhereLike(dao.SpaIsp.Columns().Name, "%"+in.Keywords+"%")
	}

	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.SpaIsp.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.SpaIsp.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询创建时间
	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.SpaIsp.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaIsp.Columns().Id)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务商列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务商列表失败，请稍后重试！")
			return
		}
	}
	return
}

// All 获取服务商
func (s *sSpaIsp) All(ctx context.Context, in *input_spa.SpaIspListInp) (list []*input_spa.SpaIspAllListModel, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaIspAllListModel{})

	// 正常状态
	mod = mod.Where(dao.SpaIsp.Columns().Status, 1)

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.SpaIsp.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.Where(dao.SpaIsp.Columns().Name, in.Keywords)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaIsp.Columns().Id)

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取服务商列表失败，请稍后重试！")
		return
	}

	return
}

// Edit 修改/新增服务商管理
func (s *sSpaIsp) Edit(ctx context.Context, in *input_spa.SpaIspEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			mod := s.Model(ctx)

			mod = mod.Fields(input_spa.SpaIspUpdateFields{})

			if _, err = mod.
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改服务商管理失败，请稍后重试！")
			}

			return
		}

		// 新增
		var (
			lastInsertId int64
		)
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaIspInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增服务商失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		return
	})
}

// Delete 删除服务商管理
func (s *sSpaIsp) Delete(ctx context.Context, in *input_spa.SpaIspDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除服务商管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取服务商管理指定信息
func (s *sSpaIsp) View(ctx context.Context, in *input_spa.SpaIspViewInp) (res *input_spa.SpaIspViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取服务商管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新服务商状态
func (s *sSpaIsp) Status(ctx context.Context, in *input_spa.SpaIspStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaIsp.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新服务商状态失败，请稍后重试！")
			return
		}

		return
	})
}

// WorkStatus 更新服务商状态
func (s *sSpaIsp) WorkStatus(ctx context.Context, in *input_spa.SpaIspWorkStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaIsp.Columns().WorkStatus: in.WorkStatus,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新服务商工作状态失败，请稍后重试！")
			return
		}

		return
	})
}

// GetIds 获取获取服务商的id
func (s *sSpaIsp) GetIds(ctx context.Context, name []string) (ids []int, err error) {
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
func (s *sSpaIsp) Bind(ctx context.Context, in *input_spa.SpaIspBindInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 判断会员是否已绑定技师
		memberCount, err := dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().MemberId, in.MemberId).Count()
		if memberCount > 0 {
			err = gerror.New("该用户已被绑定技师，请重新选择用户")
			return
		}

		// 判断会员是否已绑定服务商
		ispCount, err := dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().MemberId, in.MemberId).Count()
		if ispCount > 0 {
			err = gerror.New("该用户已被绑定，请重新选择用户")
			return
		}

		if _, err = dao.SpaIsp.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaIsp.Columns().MemberId: in.MemberId,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

// Unbind 解绑用户
func (s *sSpaIsp) Unbind(ctx context.Context, in *input_spa.SpaIspUnbindInp) (err error) {

	if _, err = dao.SpaIsp.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaIsp.Columns().MemberId: 0,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}
