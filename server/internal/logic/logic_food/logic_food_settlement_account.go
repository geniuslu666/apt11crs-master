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

type sFoodSettlementAccount struct{}

func NewFoodSettlementAccount() *sFoodSettlementAccount {
	return &sFoodSettlementAccount{}
}

func init() {
	service.RegisterFoodSettlementAccount(NewFoodSettlementAccount())
}

func (s *sFoodSettlementAccount) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodSettlementAccount.Ctx(ctx), option...)
}

func (s *sFoodSettlementAccount) List(ctx context.Context, in *input_food.FoodSettlementAccountListInp) (list []*input_food.FoodSettlementAccountListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodSettlementAccountListModel{})

	mod = mod.Where(dao.FoodSettlementAccount.Columns().RestaurantId, in.RestaurantId)

	if in.Status > 0 {
		mod = mod.Where(dao.FoodSettlementAccount.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.FoodSettlementAccount.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算账户列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlementAccount) Edit(ctx context.Context, in *input_food.FoodSettlementAccountEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodSettlementAccountUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改结算账户失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodSettlementAccountInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增结算账户失败，请稍后重试！")
		}
		return
	})
}

func (s *sFoodSettlementAccount) Delete(ctx context.Context, in *input_food.FoodSettlementAccountDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除结算账户失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlementAccount) View(ctx context.Context, in *input_food.FoodSettlementAccountViewInp) (res *input_food.FoodSettlementAccountViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取结算账户信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSettlementAccount) Status(ctx context.Context, in *input_food.FoodSettlementAccountStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodSettlementAccount.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新结算账户状态失败，请稍后重试！")
		return
	}
	return
}
