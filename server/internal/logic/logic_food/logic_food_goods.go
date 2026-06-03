package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sFoodGoods struct{}

func NewFoodGoods() *sFoodGoods {
	return &sFoodGoods{}
}

func init() {
	service.RegisterFoodGoods(NewFoodGoods())
}

func (s *sFoodGoods) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodGoods.Ctx(ctx), option...)
}

func (s *sFoodGoods) List(ctx context.Context, in *input_food.FoodGoodsListInp) (list []*input_food.FoodGoodsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_food.FoodGoodsListModel{})

	if !g.IsEmpty(in.GoodsName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.GoodsName, "goods_name")
		if err == nil {
			goodsIds, _ := service.FoodGoods().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.FoodGoods.Columns().Id, goodsIds)
		}
	}

	if !g.IsEmpty(in.RestaurantId) {
		mod = mod.Where(dao.FoodGoods.Columns().RestaurantId, in.RestaurantId)
	}

	if !g.IsEmpty(in.GoodsState) {
		mod = mod.Where(dao.FoodGoods.Columns().GoodsState, in.GoodsState)
	}

	if !g.IsEmpty(in.PriceRange) {
		if !g.IsEmpty(in.PriceRange[0]) {
			mod = mod.WhereGTE(dao.FoodGoods.Columns().Price, in.PriceRange[0])
		}
		if !g.IsEmpty(in.PriceRange[1]) {
			mod = mod.WhereLTE(dao.FoodGoods.Columns().Price, in.PriceRange[1])
		}
	}

	mod = mod.Page(in.Page, in.PerPage)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.OrderDesc(dao.FoodGoods.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取套餐管理列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodGoods) Export(ctx context.Context, in *input_food.FoodGoodsListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_food.FoodGoodsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出套餐管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("套餐管理")
		exports   []input_food.FoodGoodsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sFoodGoods) Edit(ctx context.Context, in *input_food.FoodGoodsEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object                gdb.Record
			FoodGoodsNameDao      *input_language.LoadLanguage
			NameLanguageStruct    input_language.LanguageModel
			FoodGoodsContentDao   *input_language.LoadLanguage
			ContentLanguageStruct input_language.LanguageModel
			GoodsLabelInfo        []*entity.FoodGoodsLabel
			Restaurant            entity.FoodRestaurant
		)
		NameUuid := guid.S([]byte("name"))
		ContentUuid := guid.S([]byte("content"))

		if in.Id > 0 {

			// 获取餐厅信息
			if err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, in.RestaurantId).Scan(&Restaurant); err != nil {
				err = gerror.Wrap(err, "餐厅信息不存在！")
				return
			}

			if Restaurant.CooperateTypeId != 4 {
				in.ToretaCourseId = ""
			}

			if Object, err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("GoodsName")]) {
				NameUuid = Object["goods_name"].String()
			}
			FoodGoodsNameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.FoodGoods.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("GoodsName"),
			}
			NameLanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, FoodGoodsNameDao); err != nil {
				return
			}

			in.GoodsName = NameUuid

			if !g.IsEmpty(Object[gstr.ToLower("GoodsContent")]) {
				ContentUuid = Object["goods_content"].String()
			}
			FoodGoodsContentDao = &input_language.LoadLanguage{
				Uuid: ContentUuid,
				Tag:  dao.FoodGoods.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("GoodsContent"),
			}
			ContentLanguageStruct = in.DescriptionLanguage
			if err = service.BasicsLanguage().Sync(ctx, ContentLanguageStruct, FoodGoodsContentDao); err != nil {
				return
			}

			in.GoodsContent = ContentUuid

			if _, err = s.Model(ctx).
				Fields(input_food.FoodGoodsUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改套餐管理失败，请稍后重试！")
			}

			if !g.IsEmpty(in.LabelIds) {
				for _, Label := range strings.Split(in.LabelIds, ",") {
					if !g.IsEmpty(Label) {
						LabelId, _ := strconv.ParseInt(Label, 10, 64)
						GoodsLabelInfo = append(GoodsLabelInfo, &entity.FoodGoodsLabel{
							GoodsId: in.Id,
							LabelId: LabelId,
						})
					}
				}
			}

			if !g.IsEmpty(GoodsLabelInfo) {
				if _, err = dao.FoodGoodsLabel.Ctx(ctx).Where(dao.FoodGoodsLabel.Columns().GoodsId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理套餐标签旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.FoodGoodsLabel.Ctx(ctx).OmitEmptyData().Insert(GoodsLabelInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.FoodGoodsLabel.Ctx(ctx).Where(dao.FoodGoodsLabel.Columns().GoodsId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理套餐标签旧数据失败，请稍后重试！")
					return
				}
			}

			return
		}

		var (
			lastInsertId int64
		)
		in.GoodsName = NameUuid
		in.GoodsContent = ContentUuid
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodGoodsInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增套餐管理失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		if !g.IsEmpty(in.LabelIds) {
			for _, Label := range strings.Split(in.LabelIds, ",") {
				if !g.IsEmpty(Label) {
					LabelId, _ := strconv.ParseInt(Label, 10, 64)
					GoodsLabelInfo = append(GoodsLabelInfo, &entity.FoodGoodsLabel{
						GoodsId: lastInsertId,
						LabelId: LabelId,
					})
				}
			}
		}

		if !g.IsEmpty(GoodsLabelInfo) {
			if _, err = dao.FoodGoodsLabel.Ctx(ctx).Where(dao.FoodGoodsLabel.Columns().GoodsId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理套餐标签旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.FoodGoodsLabel.Ctx(ctx).OmitEmptyData().Insert(GoodsLabelInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.FoodGoodsLabel.Ctx(ctx).Where(dao.FoodGoodsLabel.Columns().GoodsId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理套餐标签旧数据失败，请稍后重试！")
				return
			}
		}

		FoodGoodsNameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.FoodGoods.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("GoodsName"),
		}
		NameLanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, FoodGoodsNameDao); err != nil {
			return
		}

		FoodGoodsContentDao = &input_language.LoadLanguage{
			Uuid: ContentUuid,
			Tag:  dao.FoodGoods.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("GoodsContent"),
		}
		ContentLanguageStruct = in.DescriptionLanguage
		if err = service.BasicsLanguage().Sync(ctx, ContentLanguageStruct, FoodGoodsContentDao); err != nil {
			return
		}
		return
	})
}

func (s *sFoodGoods) Delete(ctx context.Context, in *input_food.FoodGoodsDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除套餐管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodGoods) View(ctx context.Context, in *input_food.FoodGoodsViewInp) (res *input_food.FoodGoodsViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取套餐管理信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取套餐管理信息，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodGoods) Status(ctx context.Context, in *input_food.FoodGoodsStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.FoodGoods.Columns().GoodsState: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新套餐状态失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodGoods) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("goods_name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}
