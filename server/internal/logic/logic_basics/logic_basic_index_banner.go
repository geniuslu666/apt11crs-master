package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsIndexBanner struct{}

func NewBasicsIndexBanner() *sBasicsIndexBanner {
	return &sBasicsIndexBanner{}
}

func init() {
	service.RegisterBasicsIndexBanner(NewBasicsIndexBanner())
}

func (s *sBasicsIndexBanner) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.IndexBanner.Ctx(ctx), option...)
}

func (s *sBasicsIndexBanner) List(ctx context.Context, in *input_basics.IndexBannerListInp) (list []*input_basics.IndexBannerListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.IndexBannerListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.IndexBanner.Columns().Id, in.Id)
	}

	if in.Language != "" {
		mod = mod.Where(dao.IndexBanner.Columns().Language, in.Language)
	}

	if in.Model != "" {
		mod = mod.Where(dao.IndexBanner.Columns().Model, in.Model)
	}

	if in.Chain != "" {
		mod = mod.Where(dao.IndexBanner.Columns().Chain, in.Chain)
	}

	if in.BannerStatus > 0 {
		mod = mod.Where(dao.IndexBanner.Columns().BannerStatus, in.BannerStatus)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsIndexNav.Columns().Sort).OrderDesc(dao.IndexBanner.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取banner列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsIndexBanner) Edit(ctx context.Context, in *input_basics.IndexBannerEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.IndexBannerUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改banner失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.IndexBannerInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增banner失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsIndexBanner) Delete(ctx context.Context, in *input_basics.IndexBannerDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除banner失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取Banner最大排序
func (s *sBasicsIndexBanner) MaxSort(ctx context.Context, in *input_basics.IndexBannerMaxSortInp) (res *input_basics.IndexBannerMaxSortModel, err error) {
	if err = dao.IndexBanner.Ctx(ctx).Fields(dao.IndexBanner.Columns().Sort).OrderDesc(dao.IndexBanner.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Banner最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_basics.IndexBannerMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sBasicsIndexBanner) View(ctx context.Context, in *input_basics.IndexBannerViewInp) (res *input_basics.IndexBannerViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取banner信息失败，请稍后重试！")
		return
	}
	return
}

// Status 更新Banner状态
func (s *sBasicsIndexBanner) Status(ctx context.Context, in *input_basics.IndexBannerStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.IndexBanner.Columns().BannerStatus: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新Banner状态失败，请稍后重试！")
		return
	}
	return
}

// BannerSort 编辑排序
func (s *sBasicsIndexBanner) BannerSort(ctx context.Context, in *input_basics.IndexBannerSortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Update(g.Map{
		dao.IndexBanner.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改排序失败，请稍后重试！")
		return
	}
	return
}
