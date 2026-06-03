package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsTestNav struct{}

func NewBasicsTestNav() *sBasicsTestNav {
	return &sBasicsTestNav{}
}

func init() {
	service.RegisterBasicsTestNav(NewBasicsTestNav())
}

func (s *sBasicsTestNav) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsTestNav.Ctx(ctx), option...)
}

func (s *sBasicsTestNav) List(ctx context.Context, in *input_basics.PmsTestNavListInp) (list []*input_basics.PmsTestNavListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsTestNavListModel{})
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsTestNav.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sBasicsTestNav) ApiNavAll(ctx context.Context, in *input_basics.PmsTestNavAllInp) (list []*input_basics.PmsTestNavApiListModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsTestNavApiListModel{})

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsTestNav.Columns().Status, in.Status)
	}
	mod = mod.OrderDesc(dao.PmsTestNav.Columns().Sort)
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsTestNav) All(ctx context.Context, in *input_basics.PmsTestNavAllInp) (list []*input_basics.PmsTestNavAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsTestNavAllModel{})

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsTestNav.Columns().Status, in.Status)
	}
	mod = mod.OrderDesc(dao.PmsTestNav.Columns().Sort)
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsTestNav) Edit(ctx context.Context, in *input_basics.PmsTestNavEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsTestNavUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsTestNavInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsTestNav) Delete(ctx context.Context, in *input_basics.PmsTestNavDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsTestNav) View(ctx context.Context, in *input_basics.PmsTestNavViewInp) (res *input_basics.PmsTestNavViewModel, err error) {

	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取信息失败，请稍后重试！")
		return
	}

	return
}

func (s *sBasicsTestNav) Switch(ctx context.Context, in *input_basics.PmsTestNavSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.PmsTestNav.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}

// NavSort 编辑排序
func (s *sBasicsTestNav) NavSort(ctx context.Context, in *input_basics.PmsTestNavSortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Update(g.Map{
		dao.PmsTestNav.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改排序失败，请稍后重试！")
		return
	}
	return
}
