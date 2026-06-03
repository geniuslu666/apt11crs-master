package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodSettlement struct{}

func NewFoodSettlement() *sFoodSettlement {
	return &sFoodSettlement{}
}

func init() {
	service.RegisterFoodSettlement(NewFoodSettlement())
}

func (s *sFoodSettlement) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodSettlement.Ctx(ctx), option...)
}

func (s *sFoodSettlement) List(ctx context.Context, in *input_food.FoodSettlementListInp) (list []*input_food.FoodSettlementListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodSettlementListModel{})

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.FoodSettlement.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算模式列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlement) Edit(ctx context.Context, in *input_food.FoodSettlementEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodSettlementUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改结算模式失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodSettlementInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增结算模式失败，请稍后重试！")
		}
		return
	})
}

func (s *sFoodSettlement) Delete(ctx context.Context, in *input_food.FoodSettlementDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除结算模式失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlement) View(ctx context.Context, in *input_food.FoodSettlementViewInp) (res *input_food.FoodSettlementViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算模式信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlement) Status(ctx context.Context, in *input_food.FoodSettlementStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodSettlement.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新结算模式状态失败，请稍后重试！")
		return
	}
	return
}
