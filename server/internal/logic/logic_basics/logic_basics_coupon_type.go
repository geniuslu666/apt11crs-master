package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsCouponType struct{}

func NewBasicsCouponType() *sBasicsCouponType {
	return &sBasicsCouponType{}
}

func init() {
	service.RegisterBasicsCouponType(NewBasicsCouponType())
}

func (s *sBasicsCouponType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsCouponType.Ctx(ctx), option...)
}

func (s *sBasicsCouponType) List(ctx context.Context, in *input_basics.PmsCouponTypeListInp) (list []*input_basics.PmsCouponTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx).Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_basics.PmsCouponTypeListModel{})

	if !g.IsEmpty(in.Type) {
		mod = mod.WhereLike(dao.PmsCouponType.Columns().Type, in.Type)
	}

	if !g.IsEmpty(in.CouponName) {
		mod = mod.WhereLike(dao.PmsCouponType.Columns().CouponName, "%"+in.CouponName+"%")
	}

	if in.ValidityType > 0 {
		mod = mod.Where(dao.PmsCouponType.Columns().ValidityType, in.ValidityType)
	}

	if in.Status > 0 {
		mod = mod.Where(dao.PmsCouponType.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.Scene) {
		mod = mod.Where(dao.PmsCouponType.Columns().Scene, in.Scene)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.PmsCouponType.Columns().Sort).OrderDesc(dao.PmsCouponType.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取优惠券列表失败，请稍后重试！")
		return
	}

	// 批量查询已使用优惠券数量，提高效率
	if len(list) > 0 {
		var couponTypeIds []interface{}
		for _, item := range list {
			couponTypeIds = append(couponTypeIds, item.Id)
		}

		// 查询每个优惠券类型的已使用数量
		var usedCountResults []struct {
			CouponTypeId int `json:"coupon_type_id"`
			UsedCount    int `json:"used_count"`
		}

		if err = dao.PmsCoupon.Ctx(ctx).
			Fields("coupon_type_id, COUNT(*) as used_count").
			WhereIn("coupon_type_id", couponTypeIds).
			Where("state = ?", 2). // state=2 表示已使用
			Group("coupon_type_id").
			Scan(&usedCountResults); err != nil {
			err = gerror.Wrap(err, "查询已使用优惠券数量失败")
			return
		}

		// 构建 map 以便快速查找
		usedCountMap := make(map[int]int)
		for _, result := range usedCountResults {
			usedCountMap[result.CouponTypeId] = result.UsedCount
		}

		// 更新列表中的已使用数量
		for _, item := range list {
			if usedCount, exists := usedCountMap[item.Id]; exists {
				item.UsedCount = usedCount
			} else {
				item.UsedCount = 0
			}
		}
	}

	return
}

func (s *sBasicsCouponType) All(ctx context.Context, in *input_basics.PmsCouponTypeListInp) (list []*input_basics.PmsCouponTypeAllListModel, err error) {
	mod := s.Model(ctx).Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_basics.PmsCouponTypeAllListModel{})

	mod = mod.Where(dao.PmsCouponType.Columns().Status, 1)

	if !g.IsEmpty(in.CouponIds) {
		mod = mod.WhereIn(dao.PmsCouponType.Columns().Id, strings.Split(in.CouponIds, ","))
	}

	mod = mod.OrderDesc(dao.PmsCouponType.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取优惠券列表失败，请稍后重试！")
		return
	}

	var (
		couponTypeList []*input_basics.PmsCouponTypeAllListModel
	)

	if !g.IsEmpty(in.CouponIds) {
		couponIdArr := strings.Split(in.CouponIds, ",")
		for _, couponId := range couponIdArr {
			for _, v := range list {
				if v.Id == gvar.New(couponId).Int() {
					couponTypeList = append(couponTypeList, v)
				}
			}
		}
		list = couponTypeList
	}

	return
}

func (s *sBasicsCouponType) Edit(ctx context.Context, in *input_basics.PmsCouponTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object               gdb.Record
			PmsCouponTypeDao     *input_language.LoadLanguage
			PmsCouponTypeDescDao *input_language.LoadLanguage
			LanguageStruct       input_language.LanguageModel
			DescStruct           input_language.LanguageModel
		)
		Uuid := guid.S([]byte("coupon_name"))

		DescUuid := guid.S([]byte("desc"))

		if in.ValidityType == 1 {
			// 当日的23点59分59秒
			in.EndUseTime = gtime.New(gtime.New(in.EndUseTime).StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)
		}

		if in.Id > 0 {

			if Object, err = dao.PmsCouponType.Ctx(ctx).Where(dao.PmsCouponType.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object["coupon_name"]) {
				Uuid = Object["coupon_name"].String()
				DescUuid = Object["desc"].String()
			}

			PmsCouponTypeDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.PmsCouponType.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("CouponName"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, PmsCouponTypeDao); err != nil {
				return
			}

			PmsCouponTypeDescDao = &input_language.LoadLanguage{
				Uuid: DescUuid,
				Tag:  dao.PmsCouponType.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Desc"),
			}
			DescStruct = in.DescLanguage
			if err = service.BasicsLanguage().Sync(ctx, DescStruct, PmsCouponTypeDescDao); err != nil {
				return
			}

			in.Desc = DescUuid
			in.CouponName = Uuid
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsCouponTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改优惠券失败，请稍后重试！")
			}
			return
		}

		var (
			lastInsertId int64
		)
		in.CouponName = Uuid
		in.Desc = DescUuid
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsCouponTypeInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增优惠券失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增优惠券失败，请稍后重试！")
			return
		}

		PmsCouponTypeDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.PmsCouponType.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("CouponName"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, PmsCouponTypeDao); err != nil {
			return
		}

		PmsCouponTypeDescDao = &input_language.LoadLanguage{
			Uuid: DescUuid,
			Tag:  dao.PmsCouponType.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Desc"),
		}
		DescStruct = in.DescLanguage
		if err = service.BasicsLanguage().Sync(ctx, DescStruct, PmsCouponTypeDescDao); err != nil {
			return
		}
		return
	})
}

func (s *sBasicsCouponType) Delete(ctx context.Context, in *input_basics.PmsCouponTypeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除优惠券失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsCouponType) MaxSort(ctx context.Context, in *input_basics.PmsCouponTypeMaxSortInp) (res *input_basics.PmsCouponTypeMaxSortModel, err error) {
	if err = dao.PmsCouponType.Ctx(ctx).Fields(dao.PmsCouponType.Columns().Sort).OrderDesc(dao.PmsCouponType.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取优惠券最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_basics.PmsCouponTypeMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sBasicsCouponType) View(ctx context.Context, in *input_basics.PmsCouponTypeViewInp) (res *input_basics.PmsCouponTypeViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取优惠券信息，请稍后重试！")
		return
	}

	if res.Scene == 1 {
		res.PropertyNames = "全部物业可用"
		if !g.IsEmpty(res.PropertyIds) {
			res.PropertyNames = ""
			var propertyArr []*input_hotel.PmsPropertyAllModel
			propertyArr, err = service.HotelService().PropertyAll(ctx, &input_hotel.PmsPropertyAllInp{
				Ids: res.PropertyIds,
			})
			if err != nil {
				err = gerror.Wrap(err, "获取物业失败！")
				return
			}
			if err = gvar.New(propertyArr).Struct(&res.Property); err != nil {
				return
			}
			// 获取物业名称以逗号 分隔
			for _, property := range propertyArr {
				res.PropertyNames = res.PropertyNames + property.Name + "，"
			}
			res.PropertyNames = gstr.TrimRight(res.PropertyNames, "，")
		}
	}

	if res.Scene == 2 {
		res.RestaurantNames = "全部门店可用"
		if !g.IsEmpty(res.RestaurantIds) {
			res.RestaurantNames = ""
			restaurantArr, _ := dao.FoodRestaurant.Ctx(ctx).Hook(hook2.PmsFindLanguageValueHook).Fields("name").
				WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(res.RestaurantIds, ",")).Array()

			if !g.IsEmpty(restaurantArr) {
				// 获取餐厅名称以逗号 分隔
				for _, restaurant := range restaurantArr {
					res.RestaurantNames = res.RestaurantNames + g.NewVar(restaurant).String() + "，"
				}
				res.RestaurantNames = gstr.TrimRight(res.RestaurantNames, "，")
			}
		}
	}

	if res.Scene == 3 {
		spaServiceTypeStr := "到店/上门"
		if !g.IsEmpty(res.ServiceIds) {
			spaServiceTypeStr = ""
			serviceIdsArr := strings.Split(res.ServiceIds, ",")
			for _, serviceType := range serviceIdsArr {
				switch serviceType {
				case "ToStore":
					spaServiceTypeStr = spaServiceTypeStr + "到店/"
					break
				case "ToDoor":
					spaServiceTypeStr = spaServiceTypeStr + "上门/"
					break
				}

			}
			spaServiceTypeStr = gstr.TrimRight(spaServiceTypeStr, "/")
		}
		res.SpaServiceTypeStr = spaServiceTypeStr
	}

	if res.Scene == 4 {
		carServiceTypeStr := "接机/送机/包车"
		if !g.IsEmpty(res.CarServiceTypes) {
			carServiceTypeStr = ""
			serviceTypeArr := strings.Split(res.CarServiceTypes, ",")
			for _, serviceType := range serviceTypeArr {
				switch serviceType {
				case "PICKUP":
					carServiceTypeStr = carServiceTypeStr + "接机/"
					break
				case "DELIVERY":
					carServiceTypeStr = carServiceTypeStr + "送机/"
					break
				case "CHARTERED":
					carServiceTypeStr = carServiceTypeStr + "包车/"
					break
				}
			}
			carServiceTypeStr = gstr.TrimRight(carServiceTypeStr, "/")
		}
		res.CarServiceTypeStr = carServiceTypeStr
	}

	return
}

func (s *sBasicsCouponType) AppView(ctx context.Context, in *input_basics.PmsCouponTypeAppViewInp) (res *input_basics.PmsCouponTypeAppViewModel, err error) {
	if err = s.Model(ctx).Hook(hook.PmsFindLanguageValueHook).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取优惠券信息失败，请稍后重试！")
		return
	}

	// 当前语言
	var Language = contexts.GetLanguage(ctx)

	if res.Scene == 1 {
		// 全部物业语言翻译
		var LanguageMap = map[string]string{
			"zh":    "全部物业可用",
			"zh_CN": "全部物業可用",
			"en":    "All properties",
			"ja":    "全物件",
			"ko":    "모든 숙소에서 사용 가능",
		}
		res.PropertyNames = LanguageMap[Language]
		if !g.IsEmpty(res.PropertyIds) {
			res.PropertyNames = ""
			var propertyArr []*input_hotel.PmsPropertyAllModel
			propertyArr, err = service.HotelService().PropertyAll(ctx, &input_hotel.PmsPropertyAllInp{
				Ids: res.PropertyIds,
			})
			if err != nil {
				err = gerror.Wrap(err, "获取物业失败！")
				return
			}
			// 获取物业名称以逗号 分隔
			for _, property := range propertyArr {
				res.PropertyNames = res.PropertyNames + property.Name + "，"
			}
			res.PropertyNames = gstr.TrimRight(res.PropertyNames, "，")
		}
	}

	if res.Scene == 2 {
		// 全部门店语言翻译
		var LanguageMap = map[string]string{
			"zh":    "全部门店可用",
			"zh_CN": "全部門店可用",
			"en":    "All stores",
			"ja":    "全店舗対応",
			"ko":    "모든 매장에서 사용 가능",
		}
		res.RestaurantNames = LanguageMap[Language]
		if !g.IsEmpty(res.RestaurantIds) {
			res.RestaurantNames = ""
			restaurantArr, _ := dao.FoodRestaurant.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).Fields("name").
				WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(res.RestaurantIds, ",")).Array()

			if !g.IsEmpty(restaurantArr) {
				// 获取餐厅名称以逗号 分隔
				for _, restaurant := range restaurantArr {
					res.RestaurantNames = res.RestaurantNames + g.NewVar(restaurant).String() + "，"
				}
				res.RestaurantNames = gstr.TrimRight(res.RestaurantNames, "，")
			}
		}
	}

	if res.Scene == 3 {
		var ToStoreLangMap = map[string]string{
			"zh":    "到店服务",
			"zh_CN": "到店服務",
			"en":    "In-Store",
			"ja":    "来店サービス",
			"ko":    "매장 방문 서비스",
		}
		var ToDoorLangMap = map[string]string{
			"zh":    "上门服务",
			"zh_CN": "上門服務",
			"en":    "To-Room",
			"ja":    "出張サービス",
			"ko":    "방문 서비스",
		}

		spaServiceTypeStr := ToStoreLangMap[Language] + "/" + ToDoorLangMap[Language]
		if !g.IsEmpty(res.ServiceIds) {
			spaServiceTypeStr = ""
			serviceIdsArr := strings.Split(res.ServiceIds, ",")
			for _, serviceType := range serviceIdsArr {
				switch serviceType {
				case "ToStore":
					spaServiceTypeStr = spaServiceTypeStr + ToStoreLangMap[Language] + "/"
					break
				case "ToDoor":
					spaServiceTypeStr = spaServiceTypeStr + ToDoorLangMap[Language] + "/"
					break
				}

			}
			spaServiceTypeStr = gstr.TrimRight(spaServiceTypeStr, "/")
		}
		res.SpaServiceTypeStr = spaServiceTypeStr
	}

	if res.Scene == 4 {

		var PickupLangMap = map[string]string{
			"zh":    "接机",
			"zh_CN": "接機",
			"en":    "Pick-up",
			"ja":    "空港でお迎え",
			"ko":    "공항 픽업",
		}
		var DeliveryLangMap = map[string]string{
			"zh":    "送机",
			"zh_CN": "送機",
			"en":    "Drop-off",
			"ja":    "空空港に送る",
			"ko":    "공항 드롭오프",
		}
		var CharteredLangMap = map[string]string{
			"zh":    "包车",
			"zh_CN": "包車",
			"en":    "Charter",
			"ja":    "貸切",
			"ko":    "전용 차량",
		}
		carServiceTypeStr := PickupLangMap[Language] + "/" + DeliveryLangMap[Language] + "/" + CharteredLangMap[Language]
		if !g.IsEmpty(res.CarServiceTypes) {
			carServiceTypeStr = ""
			serviceTypeArr := strings.Split(res.CarServiceTypes, ",")
			for _, serviceType := range serviceTypeArr {
				switch serviceType {
				case "PICKUP":
					carServiceTypeStr = carServiceTypeStr + PickupLangMap[Language] + "/"
					break
				case "DELIVERY":
					carServiceTypeStr = carServiceTypeStr + DeliveryLangMap[Language] + "/"
					break
				case "CHARTERED":
					carServiceTypeStr = carServiceTypeStr + CharteredLangMap[Language] + "/"
					break
				}
			}
			carServiceTypeStr = gstr.TrimRight(carServiceTypeStr, "/")
		}
		res.CarServiceTypeStr = carServiceTypeStr
	}

	var memberCouponStatus int
	switch res.Status {
	case -1:
		memberCouponStatus = 4 // 已关闭
	case 2:
		memberCouponStatus = 5 // 已结束
	default:
		// 判断是否领完
		if res.Count > 0 && res.LeadCount >= res.Count {
			// 已领完
			memberCouponStatus = 3
		} else {
			memberCouponStatus = 1
		}
	}

	if !g.IsEmpty(in.IndexActivityId) {
		var (
			MemberInfo          *model.MemberIdentity
			IndexActivity       *entity.HomepageArticles
			IndexActivityCoupon *entity.HomepageArticleCoupon
		)
		MemberInfo = contexts.GetMemberUser(ctx)
		if err = dao.HomepageArticles.Ctx(ctx).Where(dao.HomepageArticles.Columns().Id, in.IndexActivityId).Scan(&IndexActivity); err != nil {
			err = gerror.Wrap(err, "获取活动信息失败，请稍后重试！")
			return
		}
		if g.IsEmpty(IndexActivity) {
			err = gerror.New("活动信息错误，请稍后重试！")
			return
		}

		res.ActivityId = int(IndexActivity.Id)
		// 活动状态只能根据当前时间和活动时间做对比
		now := gtime.Now()
		if now.Before(IndexActivity.StartTime) {
			res.ActivityStatus = 1 // 未开始
		} else if now.After(IndexActivity.EndTime) {
			res.ActivityStatus = 3 // 已结束
		} else {
			res.ActivityStatus = 2 // 进行中
		}

		if err = dao.HomepageArticleCoupon.Ctx(ctx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).Where(dao.HomepageArticleCoupon.Columns().CouponType, "coupon").Where(dao.HomepageArticleCoupon.Columns().CouponId, in.Id).Scan(&IndexActivityCoupon); err != nil {
			err = gerror.Wrap(err, "获取活动礼品券信息失败，请稍后重试！")
			return
		}
		if g.IsEmpty(IndexActivityCoupon) {
			err = gerror.New("获取活动礼品券信息失败，请稍后重试！")
			return
		}

		// 已领完状态要先判断是否领取过
		// 先判断总库存是否已领完
		if res.Count > 0 && res.LeadCount >= res.Count {
			// 总库存已领完，需要判断用户是否已领取过
			var memberHaveReceiveCount int
			memberHaveReceiveCount, err = dao.PmsCoupon.Ctx(ctx).
				Where(dao.PmsCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.PmsCoupon.Columns().CouponTypeId, in.Id).
				Where(dao.PmsCoupon.Columns().IndexActivityId, in.IndexActivityId).
				Where(dao.PmsCoupon.Columns().Source, 5).
				Count()
			if err != nil {
				err = gerror.Wrap(err, "获取活动信息失败，请稍后重试！")
				return
			}

			if memberHaveReceiveCount > 0 {
				memberCouponStatus = 2 // 用户已领过，但总库存已领完
			} else {
				memberCouponStatus = 3 // 用户还没领过，但总库存已领完
			}
		} else {
			// 总库存未领完，判断用户个人领取状态
			var memberHaveReceiveCount int
			memberHaveReceiveCount, err = dao.PmsCoupon.Ctx(ctx).
				Where(dao.PmsCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.PmsCoupon.Columns().CouponTypeId, in.Id).
				Where(dao.PmsCoupon.Columns().IndexActivityId, in.IndexActivityId).
				Where(dao.PmsCoupon.Columns().Source, 5).
				Count()
			if err != nil {
				err = gerror.Wrap(err, "获取活动信息失败，请稍后重试！")
				return
			}

			if IndexActivityCoupon.AvailableQuantity > 0 && memberHaveReceiveCount >= IndexActivityCoupon.AvailableQuantity {
				memberCouponStatus = 2
			}
		}

		// 今日是否领取
		if memberCouponStatus == 1 && IndexActivityCoupon.PerDayAvailable > 0 {
			var todayReceiveCount int
			todayReceiveCount, err = dao.PmsCoupon.Ctx(ctx).
				Where(dao.PmsCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.PmsCoupon.Columns().CouponTypeId, in.Id).
				Where(dao.PmsCoupon.Columns().IndexActivityId, in.IndexActivityId).
				Where(dao.PmsCoupon.Columns().Source, 5).
				WhereGTE(dao.PmsCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
				WhereLTE(dao.PmsCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 23:59:59").
				Count()
			if err != nil {
				err = gerror.Wrap(err, "获取今日已领取数量失败，请稍后重试！")
				return
			}
			if todayReceiveCount >= IndexActivityCoupon.PerDayAvailable {
				memberCouponStatus = 6 // 今日已领
			}
		}
	}

	res.ReceiveStatus = memberCouponStatus

	return
}

func (s *sBasicsCouponType) Status(ctx context.Context, in *input_basics.PmsCouponTypeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsCouponType.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新优惠券状态失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsCouponType) SendMemberCoupon(ctx context.Context, in *input_basics.PmsSendMemberCouponInp, source int) (err error) {

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.CouponTypeId > 0 {

			var couponType *entity.PmsCouponType
			if err = s.Model(ctx).Where(dao.PmsCouponType.Columns().Id, in.CouponTypeId).Scan(&couponType); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(couponType) {
				err = gerror.New("优惠券不存在")
				return
			}

			if couponType.Status != 1 {
				err = gerror.New("该优惠券已关闭，无法发放")
				return
			}

			if couponType.Count <= couponType.LeadCount && couponType.Count > 0 {
				err = gerror.New("该优惠券已达领取上限，无法发放")
				return
			}

			var endTime *gtime.Time

			if couponType.ValidityType == 1 {
				endTime = couponType.EndUseTime
			} else if couponType.ValidityType == 2 {
				todayStart := gtime.Now()
				endTime = gtime.New(todayStart).Add(time.Duration(24*couponType.FixedTerm) * time.Hour)
				endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)
			}

			if _, err = g.Model(dao.PmsCoupon.Table()).Ctx(ctx).Safe().
				Fields(input_basics.PmsCouponInsertFields{}).
				Data(g.MapStrAny{
					dao.PmsCoupon.Columns().Type:            couponType.Type,
					dao.PmsCoupon.Columns().CouponName:      couponType.CouponName,
					dao.PmsCoupon.Columns().CouponTypeId:    couponType.Id,
					dao.PmsCoupon.Columns().MemberId:        in.MemberId,
					dao.PmsCoupon.Columns().Scene:           couponType.Scene,
					dao.PmsCoupon.Columns().PropertyIds:     couponType.PropertyIds,
					dao.PmsCoupon.Columns().RestaurantIds:   couponType.RestaurantIds,
					dao.PmsCoupon.Columns().ServiceIds:      couponType.ServiceIds,
					dao.PmsCoupon.Columns().CarServiceTypes: couponType.CarServiceTypes,
					dao.PmsCoupon.Columns().AtLeast:         couponType.AtLeast,

					dao.PmsCoupon.Columns().Money:         couponType.Money,
					dao.PmsCoupon.Columns().Discount:      couponType.Discount,
					dao.PmsCoupon.Columns().DiscountLimit: couponType.DiscountLimit,
					dao.PmsCoupon.Columns().State:         1,
					dao.PmsCoupon.Columns().FetchTime:     gtime.Now().Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().StartTime:     gtime.Now().Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().EndTime:       endTime,
					dao.PmsCoupon.Columns().Source:        source,
					dao.PmsCoupon.Columns().SourceOrderId: in.SourceOrderId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "发放优惠券失败，请稍后重试！")
			}

			if _, err = s.Model(ctx).WherePri(in.CouponTypeId).Increment(dao.PmsCouponType.Columns().LeadCount, 1); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
			return
		}

		err = gerror.Wrap(err, "优惠券不存在，请稍后重试！")
		return
	})
}

func (s *sBasicsCouponType) SendMemberCouponGetId(ctx context.Context, in *input_basics.PmsSendMemberCouponInp, source int) (memberCouponId int64, couponNameLanguage []*input_hotel.LanguageType, err error) {

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.CouponTypeId > 0 {

			var couponType *struct {
				entity.PmsCouponType
				NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"优惠券名称"   orm:"with:uuid=coupon_name"`
			}
			if err = s.Model(ctx).WithAll().Where(dao.PmsCouponType.Columns().Id, in.CouponTypeId).Scan(&couponType); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(couponType) {
				err = gerror.New("优惠券不存在")
				return
			}

			if couponType.Status != 1 {
				err = gerror.New("该优惠券已关闭，无法发放")
				return
			}

			if couponType.Count <= couponType.LeadCount && couponType.Count > 0 {
				err = gerror.New("该优惠券已达领取上限，无法发放")
				return
			}

			couponNameLanguage = couponType.NameLanguage

			var endTime *gtime.Time

			if couponType.ValidityType == 1 {
				endTime = couponType.EndUseTime
			} else if couponType.ValidityType == 2 {
				todayStart := gtime.Now()
				endTime = gtime.New(todayStart).Add(time.Duration(24*couponType.FixedTerm) * time.Hour)
				endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)
			}

			if memberCouponId, err = g.Model(dao.PmsCoupon.Table()).Ctx(ctx).Safe().
				Fields(input_basics.PmsCouponInsertFields{}).
				Data(g.MapStrAny{
					dao.PmsCoupon.Columns().Type:            couponType.Type,
					dao.PmsCoupon.Columns().CouponName:      couponType.CouponName,
					dao.PmsCoupon.Columns().CouponTypeId:    couponType.Id,
					dao.PmsCoupon.Columns().MemberId:        in.MemberId,
					dao.PmsCoupon.Columns().Scene:           couponType.Scene,
					dao.PmsCoupon.Columns().PropertyIds:     couponType.PropertyIds,
					dao.PmsCoupon.Columns().RestaurantIds:   couponType.RestaurantIds,
					dao.PmsCoupon.Columns().ServiceIds:      couponType.ServiceIds,
					dao.PmsCoupon.Columns().CarServiceTypes: couponType.CarServiceTypes,
					dao.PmsCoupon.Columns().AtLeast:         couponType.AtLeast,

					dao.PmsCoupon.Columns().Money:         couponType.Money,
					dao.PmsCoupon.Columns().Discount:      couponType.Discount,
					dao.PmsCoupon.Columns().DiscountLimit: couponType.DiscountLimit,
					dao.PmsCoupon.Columns().State:         1,
					dao.PmsCoupon.Columns().FetchTime:     gtime.Now().Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().StartTime:     gtime.Now().Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().EndTime:       endTime,
					dao.PmsCoupon.Columns().Source:        source,
					dao.PmsCoupon.Columns().SourceOrderId: in.SourceOrderId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "发放优惠券失败，请稍后重试！")
				return
			}

			if _, err = s.Model(ctx).WherePri(in.CouponTypeId).Increment(dao.PmsCouponType.Columns().LeadCount, 1); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
			return
		}

		err = gerror.New("优惠券不存在，请稍后重试！")
		return
	})
	return
}

func (s *sBasicsCouponType) AppReceiveCouponType(ctx context.Context, in *input_basics.PmsCouponTypeAppReceiveInp) (err error) {
	var (
		IndexActivity  *entity.HomepageArticles
		CouponTypeInfo *entity.PmsCouponType
		MemberInfo     *model.MemberIdentity
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		MemberInfo = contexts.GetMemberUser(ctx)

		if in.IndexActivityId > 0 {
			if err = dao.HomepageArticles.Ctx(ctx).Where(dao.EmployeeActivity.Columns().Id, in.IndexActivityId).Scan(&IndexActivity); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(IndexActivity) {
				err = gerror.New(gi18n.T(ctx, "activity_not_exist"))
				return
			}

			if IndexActivity.Status != 1 {
				err = gerror.New(gi18n.T(ctx, "activity_disabled"))
				return
			}

			// 判断当前日期是否在活动开始时间和结束时间内
			currentTime := gtime.Now()
			if IndexActivity.StartTime != nil && currentTime.Before(IndexActivity.StartTime) {
				err = gerror.New(gi18n.T(ctx, "activity_not_started"))
				return
			}
			if IndexActivity.EndTime != nil && currentTime.After(IndexActivity.EndTime) {
				err = gerror.New(gi18n.T(ctx, "activity_expired"))
				return
			}
		}

		if in.CouponTypeId > 0 {

			if in.IndexActivityId > 0 {
				// 判断活动是否绑定这个券
				var couponCount int
				couponCount, err = dao.HomepageArticleCoupon.Ctx(ctx).
					Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).
					Where(dao.HomepageArticleCoupon.Columns().CouponType, "coupon").
					Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponTypeId).
					Count()
				if err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
					return
				}
				if couponCount == 0 {
					err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_failed"))
					return
				}

				if gtime.Now().After(gtime.New(IndexActivity.EndTime)) {
					err = gerror.New(gi18n.T(ctx, "activity_has_ended"))
					return
				}

				// 判断已经领取的数量
				var receiveCouponCount int
				receiveCouponCount, err = dao.PmsCoupon.Ctx(ctx).
					Where(dao.PmsCoupon.Columns().MemberId, MemberInfo.Id).
					Where(dao.PmsCoupon.Columns().Source, 5).
					Where(dao.PmsCoupon.Columns().IndexActivityId, in.IndexActivityId).
					Where(dao.PmsCoupon.Columns().CouponTypeId, in.CouponTypeId).
					Count()
				if err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
					return
				}

				var ActivityCoupon *entity.HomepageArticleCoupon
				if err = dao.HomepageArticleCoupon.Ctx(ctx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).Where(dao.HomepageArticleCoupon.Columns().CouponType, "coupon").Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponTypeId).Scan(&ActivityCoupon); err != nil {
					return
				}
				if ActivityCoupon.AvailableQuantity > 0 && receiveCouponCount >= ActivityCoupon.AvailableQuantity {
					err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_limit"))
					return
				}

				// 判断今日领取数量
				if ActivityCoupon.PerDayAvailable > 0 {
					var todayReceiveCount int
					todayReceiveCount, err = dao.PmsCoupon.Ctx(ctx).
						Where(dao.PmsCoupon.Columns().MemberId, MemberInfo.Id).
						Where(dao.PmsCoupon.Columns().Source, 5).
						Where(dao.PmsCoupon.Columns().IndexActivityId, in.IndexActivityId).
						Where(dao.PmsCoupon.Columns().CouponTypeId, in.CouponTypeId).
						WhereGTE(dao.PmsCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
						WhereLTE(dao.PmsCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 23:59:59").
						Count()
					if err != nil {
						err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
						return
					}
					if ActivityCoupon.PerDayAvailable > 0 && todayReceiveCount >= ActivityCoupon.PerDayAvailable {
						err = gerror.New(gi18n.T(ctx, "validate_coupon_daily_limit"))
						return
					}
				}

			}

			if err = dao.PmsCouponType.Ctx(ctx).Where(dao.PmsCouponType.Columns().Id, in.CouponTypeId).Scan(&CouponTypeInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(CouponTypeInfo) {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_type_activity_failed"))
				return
			}

			if CouponTypeInfo.Status != 1 {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_type_activity_stop"))
				return
			}

			if CouponTypeInfo.Count <= CouponTypeInfo.LeadCount && CouponTypeInfo.Count > 0 {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_type_activity_limit"))
				return
			}

			var endTime *gtime.Time
			var memberCouponSource int
			todayStart := gtime.Now()
			if in.IndexActivityId > 0 {
				memberCouponSource = 5
			} else {
				memberCouponSource = 1
			}
			if CouponTypeInfo.ValidityType == 1 {
				endTime = CouponTypeInfo.EndUseTime
			} else if CouponTypeInfo.ValidityType == 2 {
				endTime = gtime.New(todayStart).Add(time.Duration(24*CouponTypeInfo.FixedTerm) * time.Hour)
				endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)
			}

			if _, err = g.Model(dao.PmsCoupon.Table()).Ctx(ctx).Safe().
				Fields(input_basics.PmsCouponInsertFields{}).
				Data(g.MapStrAny{
					dao.PmsCoupon.Columns().Type:            CouponTypeInfo.Type,
					dao.PmsCoupon.Columns().CouponName:      CouponTypeInfo.CouponName,
					dao.PmsCoupon.Columns().CouponTypeId:    CouponTypeInfo.Id,
					dao.PmsCoupon.Columns().MemberId:        MemberInfo.Id,
					dao.PmsCoupon.Columns().Scene:           CouponTypeInfo.Scene,
					dao.PmsCoupon.Columns().PropertyIds:     CouponTypeInfo.PropertyIds,
					dao.PmsCoupon.Columns().RestaurantIds:   CouponTypeInfo.RestaurantIds,
					dao.PmsCoupon.Columns().ServiceIds:      CouponTypeInfo.ServiceIds,
					dao.PmsCoupon.Columns().CarServiceTypes: CouponTypeInfo.CarServiceTypes,
					dao.PmsCoupon.Columns().AtLeast:         CouponTypeInfo.AtLeast,

					dao.PmsCoupon.Columns().Money:           CouponTypeInfo.Money,
					dao.PmsCoupon.Columns().Discount:        CouponTypeInfo.Discount,
					dao.PmsCoupon.Columns().DiscountLimit:   CouponTypeInfo.DiscountLimit,
					dao.PmsCoupon.Columns().State:           1,
					dao.PmsCoupon.Columns().FetchTime:       gtime.Now().Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().StartTime:       gtime.New(todayStart).Format("Y-m-d H:i:s"),
					dao.PmsCoupon.Columns().EndTime:         endTime,
					dao.PmsCoupon.Columns().Source:          memberCouponSource,
					dao.PmsCoupon.Columns().IndexActivityId: in.IndexActivityId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "receive_coupon_type_failed"))
				return
			}
			if _, err = s.Model(ctx).WherePri(in.CouponTypeId).Increment(dao.PmsCouponType.Columns().LeadCount, 1); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
				return
			}

			if in.IndexActivityId > 0 {
				if _, err = dao.HomepageArticleCoupon.Ctx(ctx).
					Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).
					Where(dao.HomepageArticleCoupon.Columns().CouponType, "coupon").
					Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponTypeId).
					Increment(dao.HomepageArticleCoupon.Columns().TotalReceived, 1); err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
					return
				}
			}

			return
		}

		err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_type_activity_failed"))
		return
	}); err != nil {
		return
	}

	return
}
