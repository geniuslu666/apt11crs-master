package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaStore struct{}

func NewSpaStore() *sSpaStore {
	return &sSpaStore{}
}

func init() {
	service.RegisterSpaStore(NewSpaStore())
}

// Model 门店管理ORM模型
func (s *sSpaStore) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaStore.Ctx(ctx), option...)
}

// List 获取门店管理列表
func (s *sSpaStore) List(ctx context.Context, in *input_spa.SpaStoreListInp) (list []*input_spa.SpaStoreListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_spa.SpaStoreListModel{})

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.SpaStore.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.SpaStore.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取门店管理列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增门店管理
func (s *sSpaStore) Edit(ctx context.Context, in *input_spa.SpaStoreEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Area entity.FoodArea
		)
		if in.AreaId > 0 {
			if err = dao.FoodArea.Ctx(ctx).Where(dao.FoodArea.Columns().Id, in.AreaId).Scan(&Area); err != nil {
				return
			}
			in.AreaPid = int(Area.Pid)
		}

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_spa.SpaStoreUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改门店管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaStoreInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增门店管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除门店管理
func (s *sSpaStore) Delete(ctx context.Context, in *input_spa.SpaStoreDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除门店管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取门店管理指定信息
func (s *sSpaStore) View(ctx context.Context, in *input_spa.SpaStoreViewInp) (res *input_spa.SpaStoreViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取门店管理信息，请稍后重试！")
		return
	}
	return
}

// Latest 获取门店管理最新信息
func (s *sSpaStore) Latest(ctx context.Context) (res *input_spa.SpaStoreViewModel, err error) {
	if err = s.Model(ctx).OrderDesc(dao.SpaStore.Columns().Id).Limit(1).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取门店信息失败，请稍后重试！")
		return
	}
	return
}
