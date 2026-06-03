package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodSeat struct{}

func NewFoodSeat() *sFoodSeat {
	return &sFoodSeat{}
}

func init() {
	service.RegisterFoodSeat(NewFoodSeat())
}

func (s *sFoodSeat) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodSeat.Ctx(ctx), option...)
}

func (s *sFoodSeat) List(ctx context.Context, in *input_food.FoodSeatListInp) (list []*input_food.FoodSeatListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodSeatListModel{})

	if in.SeatName != "" {
		mod = mod.WhereLike(dao.FoodSeat.Columns().SeatName, "%"+in.SeatName+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodSeat.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.FoodSeat.Columns().Id)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取订餐-座位表列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取订餐-座位表列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodSeat) Edit(ctx context.Context, in *input_food.FoodSeatEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodSeatUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改订餐-座位表失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodSeatInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增订餐-座位表失败，请稍后重试！")
		}
		return
	})
}

func (s *sFoodSeat) Delete(ctx context.Context, in *input_food.FoodSeatDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除订餐-座位表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSeat) View(ctx context.Context, in *input_food.FoodSeatViewInp) (res *input_food.FoodSeatViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取订餐-座位表信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodSeat) Status(ctx context.Context, in *input_food.FoodSeatStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodSeat.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新订餐-座位表状态失败，请稍后重试！")
		return
	}
	return
}
