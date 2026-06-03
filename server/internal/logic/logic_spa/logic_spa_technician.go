package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaTechnician struct{}

func NewSpaTechnician() *sSpaTechnician {
	return &sSpaTechnician{}
}

func init() {
	service.RegisterSpaTechnician(NewSpaTechnician())
}

// Model 技师管理ORM模型
func (s *sSpaTechnician) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaTechnician.Ctx(ctx), option...)
}

// List 获取技师管理列表
func (s *sSpaTechnician) List(ctx context.Context, in *input_spa.SpaTechnicianListInp) (list []*input_spa.SpaTechnicianListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaTechnicianListModel{})

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.WhereLike(dao.SpaTechnician.Columns().Name, "%"+in.Keywords+"%").WhereOrLike(dao.SpaTechnician.Columns().Phone, "%"+in.Keywords+"%")
	}

	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.SpaTechnician.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.SpaTechnician.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询创建时间
	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.SpaTechnician.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaTechnician.Columns().Id)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取技师列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取技师列表失败，请稍后重试！")
			return
		}
	}
	return
}

// All 获取技师
func (s *sSpaTechnician) All(ctx context.Context, in *input_spa.SpaTechnicianListInp) (list []*input_spa.SpaTechnicianAllListModel, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaTechnicianAllListModel{})

	// 正常状态
	mod = mod.Where(dao.SpaTechnician.Columns().Status, 1)

	if !g.IsEmpty(in.WorkStatus) {
		mod = mod.Where(dao.SpaTechnician.Columns().WorkStatus, in.WorkStatus)
	}

	// 查询名称
	if !g.IsEmpty(in.Keywords) {
		mod = mod.Where(dao.SpaTechnician.Columns().Name, in.Keywords).WhereOr(dao.SpaTechnician.Columns().Phone, in.Keywords)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaTechnician.Columns().Id)

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取技师列表失败，请稍后重试！")
		return
	}

	return
}

// Edit 修改/新增技师管理
func (s *sSpaTechnician) Edit(ctx context.Context, in *input_spa.SpaTechnicianEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			mod := s.Model(ctx)

			if in.Type == "basic" {

				mod = mod.Fields(input_spa.SpaTechnicianBasicUpdateFields{})
			} else if in.Type == "qualification" {
				mod = mod.Fields(input_spa.SpaTechnicianQualificationUpdateFields{})
			} else if in.Type == "settle" {
				mod = mod.Fields(input_spa.SpaTechnicianSettleUpdateFields{})
			} else if in.Type == "order" {
				mod = mod.Fields(input_spa.SpaTechnicianOrderUpdateFields{})
			} else {
				mod = mod.Fields(input_spa.SpaTechnicianUpdateFields{})
			}

			if _, err = mod.
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改技师管理失败，请稍后重试！")
			}

			return
		}

		// 新增
		var (
			lastInsertId int64
		)
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaTechnicianInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增技师失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		return
	})
}

// Delete 删除技师管理
func (s *sSpaTechnician) Delete(ctx context.Context, in *input_spa.SpaTechnicianDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除技师管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取技师管理指定信息
func (s *sSpaTechnician) View(ctx context.Context, in *input_spa.SpaTechnicianViewInp) (res *input_spa.SpaTechnicianViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取技师管理信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取技师管理信息，请稍后重试！")
			return
		}
	}
	return
}

// Status 更新技师状态
func (s *sSpaTechnician) Status(ctx context.Context, in *input_spa.SpaTechnicianStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaTechnician.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新技师状态失败，请稍后重试！")
			return
		}

		return
	})
}

// WorkStatus 更新技师状态
func (s *sSpaTechnician) WorkStatus(ctx context.Context, in *input_spa.SpaTechnicianWorkStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaTechnician.Columns().WorkStatus: in.WorkStatus,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新技师工作状态失败，请稍后重试！")
			return
		}

		return
	})
}

// GetIds 获取获取技师的id
func (s *sSpaTechnician) GetIds(ctx context.Context, name []string) (ids []int, err error) {
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
func (s *sSpaTechnician) Bind(ctx context.Context, in *input_spa.SpaTechnicianBindInp) (err error) {
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
			err = gerror.New("该用户已被绑定服务商，请重新选择用户")
			return
		}

		if _, err = dao.SpaTechnician.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.SpaTechnician.Columns().MemberId: in.MemberId,
			dao.SpaTechnician.Columns().IsLeader: in.IsLeader,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

// Unbind 解绑用户
func (s *sSpaTechnician) Unbind(ctx context.Context, in *input_spa.SpaTechnicianUnbindInp) (err error) {

	if _, err = dao.SpaTechnician.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaTechnician.Columns().MemberId: 0,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}
