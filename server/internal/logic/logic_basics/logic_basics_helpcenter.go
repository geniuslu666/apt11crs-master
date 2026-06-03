package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
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

type sBasicsHelpcenter struct{}

func NewBasicsHelpcenter() *sBasicsHelpcenter {
	return &sBasicsHelpcenter{}
}

func init() {
	service.RegisterBasicsHelpcenter(NewBasicsHelpcenter())
}

func (s *sBasicsHelpcenter) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsHelpcenter.Ctx(ctx), option...)
}

func (s *sBasicsHelpcenter) List(ctx context.Context, in *input_basics.PmsHelpcenterListInp) (list []*input_basics.PmsHelpcenterListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsHelpcenter.Table(), input_basics.PmsHelpcenterListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsHelpcenterListModel{}, &dao.PmsHelpcenterCategory, "pmsHelpcenterCategory"))

	mod = mod.LeftJoinOnFields(dao.PmsHelpcenterCategory.Table(), dao.PmsHelpcenter.Columns().CategoryId, "=", dao.PmsHelpcenterCategory.Columns().Id)

	if !g.IsEmpty(in.Title) {
		mod = mod.WhereLike(dao.PmsHelpcenter.Columns().Title, "%"+in.Title+"%")
	}
	if !g.IsEmpty(in.Keyword) {
		mod = mod.Where("title LIKE %?% OR content LIKE %?% ", in.Keyword)
	}
	if !g.IsEmpty(in.Language) {

		mod = mod.Where(dao.PmsHelpcenter.Columns().Language, in.Language)
	} else {

		mod = mod.Where(dao.PmsHelpcenter.Columns().Language, contexts.GetLanguage(ctx))
	}

	if !g.IsEmpty(in.CategoryId) {
		mod = mod.Where(dao.PmsHelpcenter.Columns().CategoryId, in.CategoryId)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsHelpcenter.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.PmsHelpcenterCategoryName != "" {
		mod = mod.WherePrefixLike(dao.PmsHelpcenterCategory.Table(), dao.PmsHelpcenterCategory.Columns().Name, "%"+in.PmsHelpcenterCategoryName+"%")
	}
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsHelpcenter.Table() + "." + dao.PmsHelpcenter.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取帮助中心列表失败，请稍后重试！")
			return
		}
	} else {

		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取帮助中心列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sBasicsHelpcenter) Edit(ctx context.Context, in *input_basics.PmsHelpcenterEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsHelpcenterUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改帮助中心失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsHelpcenterInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增帮助中心失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsHelpcenter) Delete(ctx context.Context, in *input_basics.PmsHelpcenterDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除帮助中心失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsHelpcenter) View(ctx context.Context, in *input_basics.PmsHelpcenterViewInp) (res *input_basics.PmsHelpcenterViewModel, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsHelpcenter.Table(), input_basics.PmsHelpcenterViewModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsHelpcenterViewModel{}, &dao.PmsHelpcenterCategory, "pmsHelpcenterCategory"))

	mod = mod.LeftJoinOnFields(dao.PmsHelpcenterCategory.Table(), dao.PmsHelpcenter.Columns().CategoryId, "=", dao.PmsHelpcenterCategory.Columns().Id)

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取帮助中心信息，请稍后重试！")
		return
	}
	return
}
