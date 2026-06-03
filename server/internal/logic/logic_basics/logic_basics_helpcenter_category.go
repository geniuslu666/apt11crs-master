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

type sBasicsHelpcenterCategory struct{}

func NewBasicsHelpcenterCategory() *sBasicsHelpcenterCategory {
	return &sBasicsHelpcenterCategory{}
}

func init() {
	service.RegisterBasicsHelpcenterCategory(NewBasicsHelpcenterCategory())
}

func (s *sBasicsHelpcenterCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsHelpcenterCategory.Ctx(ctx), option...)
}

func (s *sBasicsHelpcenterCategory) List(ctx context.Context, in *input_basics.PmsHelpcenterCategoryListInp) (list []*input_basics.PmsHelpcenterCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsHelpcenterCategoryListModel{})
	if !g.IsEmpty(in.Language) {
		mod = mod.Where(dao.PmsHelpcenterCategory.Columns().Language, in.Language)
	}
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsHelpcenterCategory.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取帮助中心分类列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取帮助中心分类列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sBasicsHelpcenterCategory) All(ctx context.Context, in *input_basics.PmsHelpcenterCategoryAllInp) (list []*input_basics.PmsHelpcenterCategoryAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsHelpcenterCategoryAllModel{})

	if in.Language != "" {
		mod = mod.Where(dao.PmsHelpcenterCategory.Columns().Language, in.Language)
	}

	mod = mod.OrderDesc(dao.PmsHelpcenterCategory.Columns().Id)

	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取帮助中心分类全部列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsHelpcenterCategory) Edit(ctx context.Context, in *input_basics.PmsHelpcenterCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsHelpcenterCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改帮助中心分类失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsHelpcenterCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增帮助中心分类失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsHelpcenterCategory) Delete(ctx context.Context, in *input_basics.PmsHelpcenterCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除帮助中心分类失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsHelpcenterCategory) View(ctx context.Context, in *input_basics.PmsHelpcenterCategoryViewInp) (res *input_basics.PmsHelpcenterCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取帮助中心分类信息，请稍后重试！")
		return
	}
	return
}
