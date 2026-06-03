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

type sThCouponCategory struct{}

func NewThCouponCategory() *sThCouponCategory {
	return &sThCouponCategory{}
}

func init() {
	service.RegisterThCouponCategory(NewThCouponCategory())
}

func (s *sThCouponCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThCouponCategory.Ctx(ctx), option...)
}

func (s *sThCouponCategory) List(ctx context.Context, in *input_th.ThCouponCategoryListInp) (list []*input_th.ThCouponCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_th.ThCouponCategoryListModel{})
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.ThCouponCategory.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取礼品券分类列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取礼品券分类列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThCouponCategory) All(ctx context.Context, in *input_th.ThCouponCategoryAllInp) (list []*input_th.ThCouponCategoryAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_th.ThCouponCategoryAllModel{})

	mod = mod.OrderDesc(dao.ThCouponCategory.Columns().Sort)

	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取礼品券分类全部列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sThCouponCategory) Edit(ctx context.Context, in *input_th.ThCouponCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_th.ThCouponCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改礼品券分类失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_th.ThCouponCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增礼品券分类失败，请稍后重试！")
		}
		return
	})
}

func (s *sThCouponCategory) Delete(ctx context.Context, in *input_th.ThCouponCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除礼品券分类失败，请稍后重试！")
		return
	}
	return
}

func (s *sThCouponCategory) View(ctx context.Context, in *input_th.ThCouponCategoryViewInp) (res *input_th.ThCouponCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取礼品券分类信息，请稍后重试！")
		return
	}
	return
}

func (s *sThCouponCategory) Switch(ctx context.Context, in *input_th.ThCouponCategorySwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThCouponCategory.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}
