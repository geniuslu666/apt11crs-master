package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sFoodCooperateType struct{}

func NewFoodCooperateType() *sFoodCooperateType {
	return &sFoodCooperateType{}
}

func init() {
	service.RegisterFoodCooperateType(NewFoodCooperateType())
}

func (s *sFoodCooperateType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodCooperateType.Ctx(ctx), option...)
}

func (s *sFoodCooperateType) List(ctx context.Context, in *input_food.FoodCooperateTypeListInp) (list []*input_food.FoodCooperateTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodCooperateTypeListModel{})

	if !g.IsEmpty(in.TypeName) {
		mod = mod.WhereLike(dao.FoodCooperateType.Columns().TypeName, "%"+in.TypeName+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodCooperateType.Columns().Status, in.Status)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.FoodCooperateType.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取订餐合作类型列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodCooperateType) Export(ctx context.Context, in *input_food.FoodCooperateTypeListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_food.FoodCooperateTypeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出订餐合作类型-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("订餐合作类型")
		exports   []input_food.FoodCooperateTypeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sFoodCooperateType) Edit(ctx context.Context, in *input_food.FoodCooperateTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodCooperateTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改订餐合作类型失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodCooperateTypeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增订餐合作类型失败，请稍后重试！")
		}
		return
	})
}

func (s *sFoodCooperateType) Delete(ctx context.Context, in *input_food.FoodCooperateTypeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除订餐合作类型失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodCooperateType) View(ctx context.Context, in *input_food.FoodCooperateTypeViewInp) (res *input_food.FoodCooperateTypeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取订餐合作类型信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodCooperateType) Status(ctx context.Context, in *input_food.FoodCooperateTypeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodCooperateType.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新订餐合作类型状态失败，请稍后重试！")
		return
	}
	return
}
