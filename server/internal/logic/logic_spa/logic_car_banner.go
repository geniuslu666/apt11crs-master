package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaBanner struct{}

func NewSpaBanner() *sSpaBanner {
	return &sSpaBanner{}
}

func init() {
	service.RegisterSpaBanner(NewSpaBanner())
}

// Model BannerORM模型
func (s *sSpaBanner) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaBanner.Ctx(ctx), option...)
}

// List 获取Banner列表
func (s *sSpaBanner) List(ctx context.Context, in *input_spa.SpaBannerListInp) (list []*input_spa.SpaBannerListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_spa.SpaBannerListModel{})

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.SpaBanner.Columns().Sort).OrderDesc(dao.SpaBanner.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取Banner列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增Banner
func (s *sSpaBanner) Edit(ctx context.Context, in *input_spa.SpaBannerEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_spa.SpaBannerUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改Banner失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaBannerInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增Banner失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除Banner
func (s *sSpaBanner) Delete(ctx context.Context, in *input_spa.SpaBannerDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除Banner失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取Banner最大排序
func (s *sSpaBanner) MaxSort(ctx context.Context, in *input_spa.SpaBannerMaxSortInp) (res *input_spa.SpaBannerMaxSortModel, err error) {
	if err = dao.SpaBanner.Ctx(ctx).Fields(dao.SpaBanner.Columns().Sort).OrderDesc(dao.SpaBanner.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Banner最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_spa.SpaBannerMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取Banner指定信息
func (s *sSpaBanner) View(ctx context.Context, in *input_spa.SpaBannerViewInp) (res *input_spa.SpaBannerViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Banner信息，请稍后重试！")
		return
	}
	return
}

// Status 更新Banner状态
func (s *sSpaBanner) Status(ctx context.Context, in *input_spa.SpaBannerStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaBanner.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新Banner状态失败，请稍后重试！")
		return
	}
	return
}
