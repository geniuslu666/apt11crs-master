package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaService struct{}

func NewSpaService() *sSpaService {
	return &sSpaService{}
}

func init() {
	service.RegisterSpaService(NewSpaService())
}

// Model 服务服务ORM模型
func (s *sSpaService) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaService.Ctx(ctx), option...)
}

// List 获取服务服务列表
func (s *sSpaService) List(ctx context.Context, in *input_spa.SpaServiceListInp) (list []*input_spa.SpaServiceListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaServiceListModel{})

	// 查询服务名称
	if !g.IsEmpty(in.Name) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.Name, "name")
		if err == nil {
			serviceIds, _ := service.SpaService().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.SpaService.Columns().Id, serviceIds)
		}
	}

	// 查询状态
	if !g.IsEmpty(in.ServiceStatus) && in.ServiceStatus > 0 && in.ServiceStatus < 3 {
		mod = mod.Where(dao.SpaService.Columns().ServiceState, in.ServiceStatus)
	}

	// 回收站
	if !g.IsEmpty(in.ServiceStatus) && in.ServiceStatus == 3 {
		mod = mod.Unscoped().WhereNotNull(dao.SpaService.Columns().DeletedAt)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaService.Columns().Sort).OrderDesc(dao.SpaService.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sSpaService) All(ctx context.Context, in *input_spa.SpaServiceListInp) (list []*input_spa.SpaServiceAllListModel, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook)

	// 字段过滤
	mod = mod.Fields(input_spa.SpaServiceAllListModel{})

	// 查询服务名称
	if !g.IsEmpty(in.Name) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.Name, "name")
		if err == nil {
			serviceIds, _ := service.SpaService().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.SpaService.Columns().Id, serviceIds)
		}
	}

	if !g.IsEmpty(in.ServiceIds) {
		mod = mod.WhereIn(dao.SpaService.Columns().Id, strings.Split(in.ServiceIds, ","))
	}

	// 查询状态
	if in.ServiceStatus > 0 {
		mod = mod.Where(dao.SpaService.Columns().ServiceState, in.ServiceStatus)
	}

	mod = mod.OrderDesc(dao.FoodRestaurant.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取服务列表失败，请稍后重试！")
		return
	}

	return
}

func (s *sSpaService) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}

// Edit 修改/新增服务服务
func (s *sSpaService) Edit(ctx context.Context, in *input_spa.SpaServiceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object                gdb.Record
			NameDao               *input_language.LoadLanguage
			NameLanguageStruct    input_language.LanguageModel
			SubNameDao            *input_language.LoadLanguage
			SubNameLanguageStruct input_language.LanguageModel
			ContentDao            *input_language.LoadLanguage
			ContentLanguageStruct input_language.LanguageModel
			ServiceLabelInfo      []*entity.SpaServiceLabel
			ServicePropertyInfo   []*entity.SpaServiceProperty
			ServiceGoodsInfo      []*entity.SpaServiceGoods
		)
		NameUuid := guid.S([]byte("name"))
		SubNameUuid := guid.S([]byte("sub_name"))
		ContentUuid := guid.S([]byte("content"))

		// 修改
		if in.Id > 0 {
			if Object, err = dao.SpaService.Ctx(ctx).Where(dao.SpaService.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(in.LabelIds) {
				for _, Label := range strings.Split(in.LabelIds, ",") {
					if !g.IsEmpty(Label) {
						LabelId, _ := strconv.ParseInt(Label, 10, 64)
						ServiceLabelInfo = append(ServiceLabelInfo, &entity.SpaServiceLabel{
							ServiceId: in.Id,
							LabelId:   LabelId,
						})
					}
				}
			}
			if !g.IsEmpty(ServiceLabelInfo) {
				if _, err = dao.SpaServiceLabel.Ctx(ctx).Where(dao.SpaServiceLabel.Columns().ServiceId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务标签旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.SpaServiceLabel.Ctx(ctx).OmitEmptyData().Insert(ServiceLabelInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.SpaServiceLabel.Ctx(ctx).Where(dao.SpaServiceLabel.Columns().ServiceId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务标签旧数据失败，请稍后重试！")
					return
				}
			}
			if !g.IsEmpty(in.PropertyIds) {
				for _, Property := range strings.Split(in.PropertyIds, ",") {
					if !g.IsEmpty(Property) {
						PropertyId, _ := strconv.ParseInt(Property, 10, 64)
						ServicePropertyInfo = append(ServicePropertyInfo, &entity.SpaServiceProperty{
							ServiceId:  in.Id,
							PropertyId: PropertyId,
						})
					}
				}
			}
			if !g.IsEmpty(ServicePropertyInfo) {
				if _, err = dao.SpaServiceProperty.Ctx(ctx).Where(dao.SpaServiceProperty.Columns().ServiceId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务物业旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.SpaServiceProperty.Ctx(ctx).OmitEmptyData().Insert(ServicePropertyInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.SpaServiceProperty.Ctx(ctx).Where(dao.SpaServiceProperty.Columns().ServiceId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务物业旧数据失败，请稍后重试！")
					return
				}
			}

			// 名称多语言
			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
				NameUuid = Object["name"].String()
			}
			NameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.SpaService.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			NameLanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
				return
			}

			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("SubName")]) {
				SubNameUuid = Object["sub_name"].String()
			}
			SubNameDao = &input_language.LoadLanguage{
				Uuid: SubNameUuid,
				Tag:  dao.SpaService.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("SubName"),
			}
			SubNameLanguageStruct = in.SubNameLanguage
			if err = service.BasicsLanguage().Sync(ctx, SubNameLanguageStruct, SubNameDao); err != nil {
				return
			}

			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("Content")]) {
				ContentUuid = Object["content"].String()
			}
			ContentDao = &input_language.LoadLanguage{
				Uuid: ContentUuid,
				Tag:  dao.SpaService.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Content"),
			}
			ContentLanguageStruct = in.ContentLanguage
			if err = service.BasicsLanguage().Sync(ctx, ContentLanguageStruct, ContentDao); err != nil {
				return
			}

			// 新增价格列表
			// 先删除所有价格
			if _, err = dao.SpaServiceGoods.Ctx(ctx).Where(dao.SpaServiceGoods.Columns().ServiceId, in.Id).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务价格旧数据失败，请稍后重试！")
				return
			}
			if !g.IsEmpty(in.PriceList) {
				for _, PriceItem := range in.PriceList {
					var (
						PriceItemNameDao            *input_language.LoadLanguage
						PriceItemNameLanguageStruct input_language.LanguageModel
					)
					PriceItemNameUuid := guid.S([]byte("name"))
					if PriceItem.Id > 0 {
						// 更新
						if !g.IsEmpty(PriceItem.GoodsName) {
							PriceItemNameUuid = PriceItem.GoodsName
						}
						PriceItemNameDao = &input_language.LoadLanguage{
							Uuid: PriceItemNameUuid,
							Tag:  dao.SpaServiceGoods.Table(),
							Type: "table",
							Key:  gstr.CaseSnakeFirstUpper("Name"),
						}
						PriceItemNameLanguageStruct = PriceItem.NameLanguage
						if err = service.BasicsLanguage().Sync(ctx, PriceItemNameLanguageStruct, PriceItemNameDao); err != nil {
							return
						}
						if _, err = dao.SpaServiceGoods.Ctx(ctx).WherePri(PriceItem.Id).Data(g.Map{
							dao.SpaServiceGoods.Columns().GoodsName:   PriceItemNameUuid,
							dao.SpaServiceGoods.Columns().Image:       PriceItem.Image,
							dao.SpaServiceGoods.Columns().DetailImage: PriceItem.DetailImage,
							dao.SpaServiceGoods.Columns().Price:       PriceItem.Price,
							dao.SpaServiceGoods.Columns().Duration:    PriceItem.Duration,
							dao.SpaServiceGoods.Columns().Status:      PriceItem.Status,
							dao.SpaServiceGoods.Columns().DeletedAt:   nil,
						}).Unscoped().Update(); err != nil {
							return err
						}
					} else {
						// 新增
						PriceItemNameDao = &input_language.LoadLanguage{
							Uuid: PriceItemNameUuid,
							Tag:  dao.SpaServiceGoods.Table(),
							Type: "table",
							Key:  gstr.CaseSnakeFirstUpper("Name"),
						}
						PriceItemNameLanguageStruct = PriceItem.NameLanguage
						if err = service.BasicsLanguage().Sync(ctx, PriceItemNameLanguageStruct, PriceItemNameDao); err != nil {
							return
						}
						if _, err = dao.SpaServiceGoods.Ctx(ctx).OmitEmptyData().Insert(g.Map{
							dao.SpaServiceGoods.Columns().ServiceId:   uint(in.Id),
							dao.SpaServiceGoods.Columns().GoodsName:   PriceItemNameUuid,
							dao.SpaServiceGoods.Columns().Image:       PriceItem.Image,
							dao.SpaServiceGoods.Columns().DetailImage: PriceItem.DetailImage,
							dao.SpaServiceGoods.Columns().Price:       PriceItem.Price,
							dao.SpaServiceGoods.Columns().Duration:    PriceItem.Duration,
							dao.SpaServiceGoods.Columns().Status:      PriceItem.Status,
						}); err != nil {
							return err
						}
					}
				}
			}

			// 修改
			in.Name = NameUuid
			in.SubName = SubNameUuid
			in.Content = ContentUuid

			if _, err = s.Model(ctx).
				Fields(input_spa.SpaServiceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改服务服务失败，请稍后重试！")
			}
			return
		}
		in.Name = NameUuid
		in.SubName = SubNameUuid
		in.Content = ContentUuid

		var (
			lastInsertId int64
		)

		// 新增
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaServiceInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增服务失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		if !g.IsEmpty(in.LabelIds) {
			for _, Label := range strings.Split(in.LabelIds, ",") {
				if !g.IsEmpty(Label) {
					LabelId, _ := strconv.ParseInt(Label, 10, 64)
					ServiceLabelInfo = append(ServiceLabelInfo, &entity.SpaServiceLabel{
						ServiceId: lastInsertId,
						LabelId:   LabelId,
					})
				}
			}
		}
		if !g.IsEmpty(ServiceLabelInfo) {
			if _, err = dao.SpaServiceLabel.Ctx(ctx).Where(dao.SpaServiceLabel.Columns().ServiceId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务标签旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.SpaServiceLabel.Ctx(ctx).OmitEmptyData().Insert(ServiceLabelInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.SpaServiceLabel.Ctx(ctx).Where(dao.SpaServiceLabel.Columns().ServiceId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务标签旧数据失败，请稍后重试！")
				return
			}
		}
		if !g.IsEmpty(in.PropertyIds) {
			for _, Property := range strings.Split(in.PropertyIds, ",") {
				if !g.IsEmpty(Property) {
					PropertyId, _ := strconv.ParseInt(Property, 10, 64)
					ServicePropertyInfo = append(ServicePropertyInfo, &entity.SpaServiceProperty{
						ServiceId:  lastInsertId,
						PropertyId: PropertyId,
					})
				}
			}
		}
		if !g.IsEmpty(ServicePropertyInfo) {
			if _, err = dao.SpaServiceProperty.Ctx(ctx).Where(dao.SpaServiceProperty.Columns().ServiceId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务物业旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.SpaServiceProperty.Ctx(ctx).OmitEmptyData().Insert(ServicePropertyInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.SpaServiceProperty.Ctx(ctx).Where(dao.SpaServiceProperty.Columns().ServiceId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务物业旧数据失败，请稍后重试！")
				return
			}
		}

		// 多语言
		NameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.SpaService.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		NameLanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
			return
		}

		SubNameDao = &input_language.LoadLanguage{
			Uuid: SubNameUuid,
			Tag:  dao.SpaService.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("SubName"),
		}
		SubNameLanguageStruct = in.SubNameLanguage
		if err = service.BasicsLanguage().Sync(ctx, SubNameLanguageStruct, SubNameDao); err != nil {
			return
		}

		ContentDao = &input_language.LoadLanguage{
			Uuid: ContentUuid,
			Tag:  dao.SpaService.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Content"),
		}
		ContentLanguageStruct = in.ContentLanguage
		if err = service.BasicsLanguage().Sync(ctx, ContentLanguageStruct, ContentDao); err != nil {
			return
		}

		// 新增价格列表
		if !g.IsEmpty(in.PriceList) {
			for _, PriceItem := range in.PriceList {
				PriceItemNameUuid := guid.S([]byte("name"))
				var (
					PriceItemNameDao            *input_language.LoadLanguage
					PriceItemNameLanguageStruct input_language.LanguageModel
				)
				PriceItemNameDao = &input_language.LoadLanguage{
					Uuid: PriceItemNameUuid,
					Tag:  dao.SpaServiceGoods.Table(),
					Type: "table",
					Key:  gstr.CaseSnakeFirstUpper("Name"),
				}
				PriceItemNameLanguageStruct = PriceItem.NameLanguage
				if err = service.BasicsLanguage().Sync(ctx, PriceItemNameLanguageStruct, PriceItemNameDao); err != nil {
					return
				}
				ServiceGoodsInfo = append(ServiceGoodsInfo, &entity.SpaServiceGoods{
					ServiceId:   uint(lastInsertId),
					GoodsName:   PriceItemNameUuid,
					Image:       PriceItem.Image,
					DetailImage: PriceItem.DetailImage,
					Price:       PriceItem.Price,
					Duration:    PriceItem.Duration,
					Status:      PriceItem.Status,
				})
			}
		}
		if !g.IsEmpty(ServiceGoodsInfo) {
			if _, err = dao.SpaServiceGoods.Ctx(ctx).OmitEmptyData().Insert(ServiceGoodsInfo); err != nil {
				return err
			}
		}

		return
	})
}

// Delete 删除按摩服务
func (s *sSpaService) Delete(ctx context.Context, in *input_spa.SpaServiceDeleteInp) (err error) {

	//  先判断是否有订单关联了服务
	var useCount int
	if useCount, err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().ServiceId, in.Id).Count(); err != nil {
		err = gerror.Wrap(err, "判断服务关联订单失败，请稍后重试！")
		return
	}

	if useCount > 0 {
		err = gerror.New("该服务有订单关联，无法删除！")
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除服务服务失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取服务服务最大排序
func (s *sSpaService) MaxSort(ctx context.Context, in *input_spa.SpaServiceMaxSortInp) (res *input_spa.SpaServiceMaxSortModel, err error) {
	if err = dao.SpaService.Ctx(ctx).Fields(dao.SpaService.Columns().Sort).OrderDesc(dao.SpaService.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取服务服务最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_spa.SpaServiceMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取服务服务指定信息
func (s *sSpaService) View(ctx context.Context, in *input_spa.SpaServiceViewInp) (res *input_spa.SpaServiceViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务服务信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务服务信息，请稍后重试！")
			return
		}
	}

	return
}

// Status 更新服务服务状态
func (s *sSpaService) Status(ctx context.Context, in *input_spa.SpaServiceStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaService.Columns().ServiceState: in.ServiceStatus,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新服务服务状态失败，请稍后重试！")
		return
	}
	return
}

// Recycle 恢复按摩服务
func (s *sSpaService) Recycle(ctx context.Context, in *input_spa.SpaServiceDeleteInp) (err error) {
	if _, err = dao.SpaService.Ctx(ctx).Unscoped().WherePri(in.Id).Data(g.Map{
		dao.SpaService.Columns().DeletedAt: nil,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "恢复按摩服务失败，请稍后重试！")
		return
	}
	return
}
