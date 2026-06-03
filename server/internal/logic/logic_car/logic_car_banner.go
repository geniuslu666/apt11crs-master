package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarBanner struct{}

func NewCarBanner() *sCarBanner {
	return &sCarBanner{}
}

func init() {
	service.RegisterCarBanner(NewCarBanner())
}

// Model BannerORM模型
func (s *sCarBanner) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarBanner.Ctx(ctx), option...)
}

// List 获取Banner列表
func (s *sCarBanner) List(ctx context.Context, in *input_car.CarBannerListInp) (list []*input_car.CarBannerListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarBannerListModel{})

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.CarBanner.Columns().Sort).OrderDesc(dao.CarBanner.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取Banner列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增Banner
func (s *sCarBanner) Edit(ctx context.Context, in *input_car.CarBannerEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarBannerUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改Banner失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarBannerInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增Banner失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除Banner
func (s *sCarBanner) Delete(ctx context.Context, in *input_car.CarBannerDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除Banner失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取Banner最大排序
func (s *sCarBanner) MaxSort(ctx context.Context, in *input_car.CarBannerMaxSortInp) (res *input_car.CarBannerMaxSortModel, err error) {
	if err = dao.CarBanner.Ctx(ctx).Fields(dao.CarBanner.Columns().Sort).OrderDesc(dao.CarBanner.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Banner最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_car.CarBannerMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取Banner指定信息
func (s *sCarBanner) View(ctx context.Context, in *input_car.CarBannerViewInp) (res *input_car.CarBannerViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Banner信息，请稍后重试！")
		return
	}
	return
}

// Status 更新Banner状态
func (s *sCarBanner) Status(ctx context.Context, in *input_car.CarBannerStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarBanner.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新Banner状态失败，请稍后重试！")
		return
	}
	return
}
