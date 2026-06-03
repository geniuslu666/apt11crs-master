package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarSettlement struct{}

func NewCarSettlement() *sCarSettlement {
	return &sCarSettlement{}
}

func init() {
	service.RegisterCarSettlement(NewCarSettlement())
}

// Model 结算模式ORM模型
func (s *sCarSettlement) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarSettlement.Ctx(ctx), option...)
}

// List 获取结算模式列表
func (s *sCarSettlement) List(ctx context.Context, in *input_car.CarSettlementListInp) (list []*input_car.CarSettlementListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarSettlementListModel{})

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.CarSettlement.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算模式列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增结算模式
func (s *sCarSettlement) Edit(ctx context.Context, in *input_car.CarSettlementEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarSettlementUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改结算模式失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarSettlementInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增结算模式失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除结算模式
func (s *sCarSettlement) Delete(ctx context.Context, in *input_car.CarSettlementDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除结算模式失败，请稍后重试！")
		return
	}
	return
}

// View 获取结算模式指定信息
func (s *sCarSettlement) View(ctx context.Context, in *input_car.CarSettlementViewInp) (res *input_car.CarSettlementViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算模式信息，请稍后重试！")
		return
	}
	return
}

// Status 更新结算模式状态
func (s *sCarSettlement) Status(ctx context.Context, in *input_car.CarSettlementStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarSettlement.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新结算模式状态失败，请稍后重试！")
		return
	}
	return
}
