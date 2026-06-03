package logic_th

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sThMchCategory struct{}

func NewThMchCategory() *sThMchCategory {
	return &sThMchCategory{}
}

func init() {
	service.RegisterThMchCategory(NewThMchCategory())
}

func (s *sThMchCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThMchCategory.Ctx(ctx), option...)
}

func (s *sThMchCategory) List(ctx context.Context, in *input_th.ThMchCategoryListInp) (list []*input_th.ThMchCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_th.ThMchCategoryListModel{})
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.ThMchCategory.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户分类列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户分类列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThMchCategory) All(ctx context.Context, in *input_th.ThMchCategoryAllInp) (list []*input_th.ThMchCategoryAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_th.ThMchCategoryAllModel{})

	mod = mod.OrderDesc(dao.ThMchCategory.Columns().Sort)

	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取商户分类全部列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sThMchCategory) Edit(ctx context.Context, in *input_th.ThMchCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_th.ThMchCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改商户分类失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_th.ThMchCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增商户分类失败，请稍后重试！")
		}
		return
	})
}

func (s *sThMchCategory) Delete(ctx context.Context, in *input_th.ThMchCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除商户分类失败，请稍后重试！")
		return
	}
	return
}

func (s *sThMchCategory) View(ctx context.Context, in *input_th.ThMchCategoryViewInp) (res *input_th.ThMchCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取商户分类信息，请稍后重试！")
		return
	}
	return
}

func (s *sThMchCategory) Switch(ctx context.Context, in *input_th.ThMchCategorySwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThMchCategory.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}
