package logic_travel

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type sTravelProduct struct{}

func NewTravelProduct() *sTravelProduct {
	return &sTravelProduct{}
}

func init() {
	service.RegisterTravelProduct(NewTravelProduct())
}

// Model 产品ORM模型
func (s *sTravelProduct) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TravelProduct.Ctx(ctx), option...)
}

// List 获取产品列表
func (s *sTravelProduct) List(ctx context.Context, in *input_travel.TravelProductListInp) (list []*input_travel.TravelProductListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	if !g.IsEmpty(in.Title) {
		uuIds, uErr := service.BasicsLanguage().GetUuids(ctx, in.Title, "title")
		if uErr == nil {
			productIds, _ := s.GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.TravelProduct.Columns().Id, productIds)
		} else {
			mod = mod.WhereLike(dao.TravelProduct.Columns().Title, "%"+in.Title+"%")
		}
	}
	if in.Status > 0 {
		mod = mod.Where(dao.TravelProduct.Columns().Status, in.Status)
	}
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TravelProduct.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.OrderDesc(dao.TravelProduct.Columns().Sort).OrderDesc(dao.TravelProduct.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取产品列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取产品列表失败，请稍后重试！")
	}
	return
}

// RecycleList 回收站列表
func (s *sTravelProduct) RecycleList(ctx context.Context, in *input_travel.TravelProductListInp) (list []*input_travel.TravelProductListModel, totalCount int, err error) {
	mod := s.Model(ctx).Unscoped().WhereNotNull(dao.TravelProduct.Columns().DeletedAt)

	if !g.IsEmpty(in.Title) {
		mod = mod.WhereLike(dao.TravelProduct.Columns().Title, "%"+in.Title+"%")
	}
	mod = mod.OrderDesc(dao.TravelProduct.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取回收站列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取回收站列表失败，请稍后重试！")
	}
	return
}

func (s *sTravelProduct) AppList(ctx context.Context, in *input_travel.TravelProductAppListInp) (list []*input_travel.TravelProductAppListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_travel.TravelProductAppListModel{}).WithAll()

	// 只显示在售
	mod = mod.Where(dao.TravelProduct.Columns().Status, 1)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	} else {
		// 不分页限制显示数量 in.PerPage
	}
	mod = mod.OrderDesc(dao.TravelProduct.Columns().Sort).OrderDesc(dao.TravelProduct.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_travel_product_list_failed"))
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_travel_product_list_failed"))
			return
		}
	}

	// 批量查询每个产品的 SKU 最低价并覆盖 Price 字段
	for _, item := range list {
		minPrice, skuErr := dao.TravelProductSku.Ctx(ctx).
			Where(dao.TravelProductSku.Columns().ProductId, item.Id).
			Where(dao.TravelProductSku.Columns().Status, 1).
			Min(dao.TravelProductSku.Columns().Price)
		if skuErr == nil && minPrice > 0 {
			item.Price = minPrice
		}
	}

	return
}

// GetIds 通过多语言 UUID 列表查询产品 ID
func (s *sTravelProduct) GetIds(ctx context.Context, uuIds []string) (ids []int64, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn(dao.TravelProduct.Columns().Title, uuIds).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取产品ID失败！")
		return
	}
	ids = g.NewVar(columns).Int64s()
	return
}

// GetSkuIds 通过多语言 UUID 列表查询产品SKU ID
func (s *sTravelProduct) GetSkuIds(ctx context.Context, uuIds []string) (ids []int64, err error) {
	columns, err := dao.TravelProductSku.Ctx(ctx).
		Fields("id").
		WhereIn(dao.TravelProductSku.Columns().Name, uuIds).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取产品SKU ID失败！")
		return
	}
	ids = g.NewVar(columns).Int64s()
	return
}

// View 获取产品详情
func (s *sTravelProduct) View(ctx context.Context, in *input_travel.TravelProductViewInp) (res *input_travel.TravelProductViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().Where(dao.TravelProduct.Columns().Id, in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取产品详情失败，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().Where(dao.TravelProduct.Columns().Id, in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取产品详情失败，请稍后重试！")
			return
		}
	}
	if res == nil {
		err = gerror.New("产品不存在")
		return
	}
	return
}

func (s *sTravelProduct) AppView(ctx context.Context, in *input_travel.TravelProductAppViewInp) (res *input_travel.TravelProductAppViewModel, err error) {

	var ProductView *entity.TravelProduct
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&ProductView); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_product_info_failed"))
		return
	}
	if ProductView == nil {
		err = gerror.New(gi18n.T(ctx, "product_not_exist"))
		return
	}

	// 计算有效订单的总人数（待核销或已完成状态的booking_num总和）
	totalBookingNum, err := dao.TravelOrder.Ctx(ctx).
		Where(dao.TravelOrder.Columns().ProductId, in.Id).
		WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "WAIT_VERIFY", "DONE"}).
		Fields(dao.TravelOrder.Columns().BookingNum).
		Sum(dao.TravelOrder.Columns().BookingNum)
	if err != nil {
		// 计算有效订单总人数失败
		err = gerror.Wrap(err, gi18n.T(ctx, "server_exception"))
		return
	}

	// 获取可预约的日期列表
	now := gtime.Now()
	var dateList []string

	// 计算最早可预约日期（至少提前预订天数）
	minBookDate := now.AddDate(0, 0, int(ProductView.AdvanceBookDays))

	// 计算最晚可预约日期（最大可预约天数）
	maxBookDate := now.AddDate(0, 0, int(ProductView.MaxBookDays))

	// 生成日期列表（从最早可预约日期到最晚可预约日期）
	currentDate := minBookDate
	for !currentDate.After(maxBookDate) {
		dateList = append(dateList, currentDate.Format("Y-m-d"))
		currentDate = currentDate.AddDate(0, 0, 1) // 加一天
	}

	// 根据当前语言设置对应的内容
	res = new(input_travel.TravelProductAppViewModel)
	res.Id = ProductView.Id
	res.Title = ProductView.Title
	res.SubTitle = ProductView.SubTitle

	// 根据当前语言选择内容
	currentLang := contexts.GetLanguage(ctx)
	// 映射语言到配置key（参考logic_travel_product.go）
	configLanguageKey := "zh" // 默认中文
	switch currentLang {
	case "zh", "zh_CN":
		res.Content = ProductView.ContentZh
		res.TripPlanning = ProductView.TripPlanningZh
		res.BookingNotes = ProductView.BookingNotesZh
		configLanguageKey = "zh"
	case "en":
		res.Content = ProductView.ContentEn
		res.TripPlanning = ProductView.TripPlanningEn
		res.BookingNotes = ProductView.BookingNotesEn
		configLanguageKey = "en"
	case "ja":
		res.Content = ProductView.ContentJa
		res.TripPlanning = ProductView.TripPlanningJa
		res.BookingNotes = ProductView.BookingNotesJa
		configLanguageKey = "ja"
	case "ko":
		res.Content = ProductView.ContentKo
		res.TripPlanning = ProductView.TripPlanningKo
		res.BookingNotes = ProductView.BookingNotesKo
		configLanguageKey = "ko"
	case "tw", "zh_TW":
		res.Content = ProductView.ContentTw
		res.TripPlanning = ProductView.TripPlanningTw
		res.BookingNotes = ProductView.BookingNotesTw
		configLanguageKey = "zh_CN"
	default:
		// 默认使用中文
		res.Content = ProductView.ContentZh
		res.TripPlanning = ProductView.TripPlanningZh
		res.BookingNotes = ProductView.BookingNotesZh
		configLanguageKey = "zh"
	}

	// 从sys_config表获取预定须知（参考app_travel_order_info.go）
	config, err := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "travelothersetting",
	})
	// 设置预定须知内容
	res.OrderRule = "" // 默认空值
	if err == nil && config != nil && config.List != nil {
		// 获取订单规则
		orderRuleKey := "orderRule_" + configLanguageKey
		if orderRuleValue := config.List[orderRuleKey]; orderRuleValue != nil {
			res.OrderRule = gvar.New(orderRuleValue).String()
		} else {
			// 如果当前语言配置不存在，尝试使用中文版本
			if orderRuleValue := config.List["orderRule_zh"]; orderRuleValue != nil {
				res.OrderRule = gvar.New(orderRuleValue).String()
			}
		}
	}

	// 轮播图
	if err = g.NewVar(ProductView.CarouselImages).Scan(&res.CarouselImages); err != nil {
		res.CarouselImages = []string{}
		err = nil
	}
	res.DailyCapacity = ProductView.DailyCapacity
	// 从 SKU 表查询最低价
	minSkuPrice, skuErr := dao.TravelProductSku.Ctx(ctx).
		Where(dao.TravelProductSku.Columns().ProductId, in.Id).
		Where(dao.TravelProductSku.Columns().Status, 1).
		Min(dao.TravelProductSku.Columns().Price)
	if skuErr == nil && minSkuPrice > 0 {
		res.Price = minSkuPrice
	} else {
		res.Price = ProductView.Price
	}
	res.MeetingPlace = ProductView.MeetingPlace
	res.MeetingTime = ProductView.MeetingTime
	res.GgLat = ProductView.GgLat
	res.GgLng = ProductView.GgLng
	res.MaxBookDays = ProductView.MaxBookDays
	res.AdvanceBookDays = ProductView.AdvanceBookDays
	res.Stock = ProductView.Stock
	res.SalesNum = ProductView.SalesNum
	res.IsStockExhausted = (ProductView.Stock - uint(totalBookingNum)) <= 0
	res.DateList = dateList

	return
}

// Edit 新增/编辑产品
func (s *sTravelProduct) Edit(ctx context.Context, in *input_travel.TravelProductEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			TitleDao               *input_language.LoadLanguage
			TitleLanguageStruct    input_language.LanguageModel
			SubTitleDao            *input_language.LoadLanguage
			SubTitleLanguageStruct input_language.LanguageModel
		)
		titleUuid := guid.S([]byte("title"))
		subTitleUuid := guid.S([]byte("sub_title"))

		// 编辑
		if in.Id > 0 {
			var object gdb.Record
			if object, err = dao.TravelProduct.Ctx(ctx).Where(dao.TravelProduct.Columns().Id, in.Id).One(); err != nil {
				return
			}

			// 读取已有的 UUID（多语言存的是 uuid，不是原文）
			if !g.IsEmpty(object[gstr.ToLower("Title")]) {
				titleUuid = object["title"].String()
			}
			if !g.IsEmpty(object["sub_title"]) {
				subTitleUuid = object["sub_title"].String()
			}

			TitleDao = &input_language.LoadLanguage{
				Uuid: titleUuid,
				Tag:  dao.TravelProduct.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Title"),
			}
			TitleLanguageStruct = in.TitleLanguage
			if err = service.BasicsLanguage().Sync(ctx, TitleLanguageStruct, TitleDao); err != nil {
				return
			}

			SubTitleDao = &input_language.LoadLanguage{
				Uuid: subTitleUuid,
				Tag:  dao.TravelProduct.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("SubTitle"),
			}
			SubTitleLanguageStruct = in.SubTitleLanguage
			if err = service.BasicsLanguage().Sync(ctx, SubTitleLanguageStruct, SubTitleDao); err != nil {
				return
			}

			in.Title = titleUuid
			in.SubTitle = subTitleUuid

			if _, err = s.Model(ctx).
				Fields(input_travel.TravelProductUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "编辑产品失败，请稍后重试！")
				return
			}

			// 处理SKU数据
			if err = s.handleSkus(ctx, int(in.Id), in.SkuList, tx); err != nil {
				err = gerror.Wrap(err, "处理产品SKU失败，请稍后重试！")
				return
			}

			return
		}

		// 新增：先同步多语言，再插入主记录
		TitleDao = &input_language.LoadLanguage{
			Uuid: titleUuid,
			Tag:  dao.TravelProduct.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Title"),
		}
		TitleLanguageStruct = in.TitleLanguage
		if err = service.BasicsLanguage().Sync(ctx, TitleLanguageStruct, TitleDao); err != nil {
			return
		}

		SubTitleDao = &input_language.LoadLanguage{
			Uuid: subTitleUuid,
			Tag:  dao.TravelProduct.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("SubTitle"),
		}
		SubTitleLanguageStruct = in.SubTitleLanguage
		if err = service.BasicsLanguage().Sync(ctx, SubTitleLanguageStruct, SubTitleDao); err != nil {
			return
		}

		in.Title = titleUuid
		in.SubTitle = subTitleUuid

		lastInsertId, err := s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_travel.TravelProductInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, "新增产品失败，请稍后重试！")
			return
		}

		// 处理SKU数据
		productId := lastInsertId
		if err = s.handleSkus(ctx, int(productId), in.SkuList, tx); err != nil {
			err = gerror.Wrap(err, "处理产品SKU失败，请稍后重试！")
			return
		}

		return
	})
}

// handleSkus 处理产品SKU数据
func (s *sTravelProduct) handleSkus(ctx context.Context, productId int, skuList []*input_travel.TravelProductSkuInp, tx gdb.TX) error {
	if skuList == nil || len(skuList) == 0 {
		return nil
	}

	// 获取现有的SKU列表
	var existingSkus []*entity.TravelProductSku
	if err := dao.TravelProductSku.Ctx(ctx).Where(dao.TravelProductSku.Columns().ProductId, productId).Scan(&existingSkus); err != nil {
		return gerror.Wrap(err, "查询现有SKU失败")
	}

	// 创建现有SKU的map
	existingSkuMap := make(map[int]*entity.TravelProductSku)
	for _, sku := range existingSkus {
		existingSkuMap[int(sku.Id)] = sku
	}

	// 处理传入的SKU列表
	var processedSkuIds []int
	for _, sku := range skuList {
		if sku.Id > 0 {
			// 更新现有SKU
			if existingSku, exists := existingSkuMap[int(sku.Id)]; exists {
				updateData := g.Map{
					dao.TravelProductSku.Columns().Price:         sku.Price,
					dao.TravelProductSku.Columns().DailyCapacity: sku.DailyCapacity,
					dao.TravelProductSku.Columns().Price:         sku.Price,
					dao.TravelProductSku.Columns().Status:        sku.Status,
					dao.TravelProductSku.Columns().Sort:          sku.Sort,
				}

				var NameLanguageStruct input_language.LanguageModel

				// 处理多语言名称 - 使用UUID方式
				if !g.IsEmpty(sku.NameLanguage) {
					// 读取已有的UUID
					if !g.IsEmpty(existingSku.Name) {
						// 使用现有UUID
						nameUuid := existingSku.Name
						NameDao := &input_language.LoadLanguage{
							Uuid: nameUuid,
							Tag:  dao.TravelProductSku.Table(),
							Type: "table",
							Key:  gstr.CaseSnakeFirstUpper("Name"),
						}
						NameLanguageStruct = sku.NameLanguage
						if err := service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
							return gerror.Wrap(err, "同步SKU多语言失败")
						}
					} else {
						// 生成新UUID
						nameUuid := guid.S([]byte("sku_name"))
						NameDao := &input_language.LoadLanguage{
							Uuid: nameUuid,
							Tag:  dao.TravelProductSku.Table(),
							Type: "table",
							Key:  gstr.CaseSnakeFirstUpper("Name"),
						}
						NameLanguageStruct := sku.NameLanguage
						if err := service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
							return gerror.Wrap(err, "同步SKU多语言失败")
						}
						// 更新SKU记录的UUID
						updateData[dao.TravelProductSku.Columns().Name] = nameUuid
					}
				}

				if _, err := dao.TravelProductSku.Ctx(ctx).TX(tx).
					Where(dao.TravelProductSku.Columns().Id, sku.Id).
					Data(updateData).Update(); err != nil {
					return gerror.Wrap(err, "更新SKU失败")
				}
				processedSkuIds = append(processedSkuIds, int(sku.Id))
			}
		} else {
			// 新增SKU
			insertData := g.Map{
				dao.TravelProductSku.Columns().ProductId:     productId,
				dao.TravelProductSku.Columns().DailyCapacity: sku.DailyCapacity,
				dao.TravelProductSku.Columns().Price:         sku.Price,
				dao.TravelProductSku.Columns().Status:        sku.Status,
				dao.TravelProductSku.Columns().Sort:          sku.Sort,
				dao.TravelProductSku.Columns().SalesNum:      0,
			}

			// 处理多语言名称 - 使用UUID方式
			if !g.IsEmpty(sku.NameLanguage) {
				// 生成新UUID
				nameUuid := guid.S([]byte("sku_name"))
				NameDao := &input_language.LoadLanguage{
					Uuid: nameUuid,
					Tag:  dao.TravelProductSku.Table(),
					Type: "table",
					Key:  gstr.CaseSnakeFirstUpper("Name"),
				}
				NameLanguageStruct := sku.NameLanguage
				if err := service.BasicsLanguage().Sync(ctx, NameLanguageStruct, NameDao); err != nil {
					return gerror.Wrap(err, "同步SKU多语言失败")
				}
				// 设置SKU记录的UUID
				insertData[dao.TravelProductSku.Columns().Name] = nameUuid
			} else {
				// 如果没有多语言数据，设置默认名称
				insertData[dao.TravelProductSku.Columns().Name] = ""
			}

			if _, err := dao.TravelProductSku.Ctx(ctx).TX(tx).
				Data(insertData).Insert(); err != nil {
				return gerror.Wrap(err, "新增SKU失败")
			}
		}
	}

	// 删除不在传入列表中的现有SKU
	for _, existingSku := range existingSkus {
		found := false
		for _, processedId := range processedSkuIds {
			if int(existingSku.Id) == processedId {
				found = true
				break
			}
		}
		if !found {
			if _, err := dao.TravelProductSku.Ctx(ctx).TX(tx).
				Where(dao.TravelProductSku.Columns().Id, existingSku.Id).
				Delete(); err != nil {
				return gerror.Wrap(err, "删除SKU失败")
			}
		}
	}

	return nil
}

// Delete 软删除产品（移入回收站）
func (s *sTravelProduct) Delete(ctx context.Context, in *input_travel.TravelProductDeleteInp) (err error) {
	_, err = s.Model(ctx).WhereIn(dao.TravelProduct.Columns().Id, in.Id).
		Delete()
	if err != nil {
		err = gerror.Wrap(err, "删除产品失败，请稍后重试！")
	}
	return
}

// Restore 从回收站恢复产品
func (s *sTravelProduct) Restore(ctx context.Context, in *input_travel.TravelProductDeleteInp) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn(dao.TravelProduct.Columns().Id, in.Id).
		Data(g.Map{dao.TravelProduct.Columns().DeletedAt: nil}).Update()
	if err != nil {
		err = gerror.Wrap(err, "恢复产品失败，请稍后重试！")
	}
	return
}

// Status 更新产品状态
func (s *sTravelProduct) Status(ctx context.Context, in *input_travel.TravelProductStatusInp) (err error) {
	_, err = s.Model(ctx).Where(dao.TravelProduct.Columns().Id, in.Id).
		Data(g.Map{dao.TravelProduct.Columns().Status: in.Status}).Update()
	if err != nil {
		err = gerror.Wrap(err, "更新产品状态失败，请稍后重试！")
	}
	return
}

// SkuStock 查询SKU库存
func (s *sTravelProduct) SkuStock(ctx context.Context, in *input_travel.TravelProductSkuStockInp) (list []*input_travel.TravelProductSkuStockModel, err error) {
	// 查询该产品的所有SKU
	var skuList []*entity.TravelProductSku
	if err = dao.TravelProductSku.Ctx(ctx).
		Where(dao.TravelProductSku.Columns().ProductId, in.ProductId).
		Where(dao.TravelProductSku.Columns().Status, 1).
		OrderDesc(dao.TravelProductSku.Columns().Sort).
		OrderAsc(dao.TravelProductSku.Columns().Id).
		Hook(hook2.PmsFindLanguageValueHook).
		Scan(&skuList); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_sku_list_failed"))
		return
	}

	if len(skuList) == 0 {
		return []*input_travel.TravelProductSkuStockModel{}, nil
	}

	// 解析预约日期
	bookDate, err := gtime.StrToTime(in.BookDate)
	if err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "invalid_date_format"))
		return
	}

	// 遍历每个SKU，计算已预订人数和剩余座位数
	list = make([]*input_travel.TravelProductSkuStockModel, 0, len(skuList))
	for _, sku := range skuList {
		// 查询该SKU在指定日期的有效订单总人数
		// 有效订单状态：WAIT_PAY（待支付）、WAIT_VERIFY（待核销）、DONE（已完成）
		bookedNum, err := dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().SkuId, sku.Id).
			Where(dao.TravelOrder.Columns().BookDate, bookDate.Format("Y-m-d")).
			WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_PAY", "WAIT_VERIFY", "DONE"}).
			Sum(dao.TravelOrder.Columns().BookingNum)
		if err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "calculate_booked_num_failed"))
			return nil, err
		}

		// 计算剩余座位数
		bookedNumUint := uint(bookedNum)
		var remainingSeats uint
		if sku.DailyCapacity > bookedNumUint {
			remainingSeats = sku.DailyCapacity - bookedNumUint
		} else {
			remainingSeats = 0
		}

		list = append(list, &input_travel.TravelProductSkuStockModel{
			Id:             sku.Id,
			Name:           sku.Name,
			Price:          sku.Price,
			DailyCapacity:  sku.DailyCapacity,
			BookedNum:      bookedNumUint,
			RemainingSeats: remainingSeats,
		})
	}

	return
}
