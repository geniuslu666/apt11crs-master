package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

type sBasicsCoupon struct{}

func NewBasicsCoupon() *sBasicsCoupon {
	return &sBasicsCoupon{}
}

func init() {
	service.RegisterBasicsCoupon(NewBasicsCoupon())
}

func (s *sBasicsCoupon) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsCoupon.Ctx(ctx), option...)
}

func (s *sBasicsCoupon) View(ctx context.Context, in *input_basics.PmsCouponViewInp) (res *input_basics.PmsCouponViewModel, err error) {
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

	return
}

func (s *sBasicsCoupon) List(ctx context.Context, in *input_basics.PmsCouponListInp) (list []*input_basics.PmsCouponListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.FieldsPrefix(dao.PmsCoupon.Table(), input_basics.PmsCouponListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsCouponListModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsCoupon.Columns().MemberId, "=", dao.PmsMember.Columns().Id)

	if in.CouponTypeId > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().CouponTypeId, in.CouponTypeId)
	}

	if in.State > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().State, in.State)
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.PmsCoupon.Columns().MemberId, in.MemberId)
		mod = mod.OrderAsc(dao.PmsCoupon.Columns().State)
	}

	if !g.IsEmpty(in.WantUseTime) {
		//	mod = mod.Where("start_time IS NULL OR start_time <= ?", in.WantUseTime)
		//	mod = mod.Where("end_time IS NULL OR end_time >= ?", in.WantUseTime)
		mod = mod.WhereNull(dao.PmsCoupon.Columns().UseTime)
	}

	//if !g.IsEmpty(in.PropertyId) {
	//	mod = mod.Where("FIND_IN_SET(?, `property_ids`) > 0", in.PropertyId)
	//}

	if !g.IsEmpty(in.Scene) {
		mod = mod.Where(dao.PmsCoupon.Columns().Scene, in.Scene)
	}

	//if !g.IsEmpty(in.AtLeast) {
	//	mod = mod.WhereLTE(dao.PmsCoupon.Columns().AtLeast, in.AtLeast)
	//}

	if in.MemberKey != "" {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().FullName, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Phone, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Mail, "%"+in.MemberKey+"%"))
	}

	if in.Source > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().Source, in.Source)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsCoupon.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		//err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsCoupon) AppList(ctx context.Context, in *input_basics.PmsCouponListInp) (list []*input_basics.PmsCouponListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.FieldsPrefix(dao.PmsCoupon.Table(), input_basics.PmsCouponListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsCouponListModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsCoupon.Columns().MemberId, "=", dao.PmsMember.Columns().Id)

	if in.CouponTypeId > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().CouponTypeId, in.CouponTypeId)
	}

	if in.State > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().State, in.State)
	} else {
		mod = mod.WhereNot(dao.PmsCoupon.Columns().State, 5)
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.PmsCoupon.Columns().MemberId, in.MemberId)
		mod = mod.OrderAsc(dao.PmsCoupon.Columns().State)
	}

	if !g.IsEmpty(in.WantUseTime) {
		//	mod = mod.Where("start_time IS NULL OR start_time <= ?", in.WantUseTime)
		//	mod = mod.Where("end_time IS NULL OR end_time >= ?", in.WantUseTime)
		mod = mod.WhereNull(dao.PmsCoupon.Columns().UseTime)
	}

	//if !g.IsEmpty(in.PropertyId) {
	//	mod = mod.Where("FIND_IN_SET(?, `property_ids`) > 0", in.PropertyId)
	//}

	if !g.IsEmpty(in.Scene) {
		mod = mod.Where(dao.PmsCoupon.Columns().Scene, in.Scene)
	}

	//if !g.IsEmpty(in.AtLeast) {
	//	mod = mod.WhereLTE(dao.PmsCoupon.Columns().AtLeast, in.AtLeast)
	//}

	if in.MemberKey != "" {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().FullName, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Phone, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Mail, "%"+in.MemberKey+"%"))
	}

	if in.Source > 0 {
		mod = mod.Where(dao.PmsCoupon.Columns().Source, in.Source)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsCoupon.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		//err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsCoupon) Stat(ctx context.Context, in *input_basics.PmsCouponStatInp) (res *input_basics.PmsCouponStatModel, err error) {
	res = &input_basics.PmsCouponStatModel{}
	mod := s.Model(ctx)

	totalSendNum, err := mod.Safe().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.TotalSendNum = gvar.New(totalSendNum).Int()

	usedCount, err := mod.Safe().Where(dao.PmsCoupon.Columns().State, 2).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.UsedNum = gvar.New(usedCount).Int()

	waitUseNum, err := mod.Safe().Where(dao.PmsCoupon.Columns().State, 1).Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.WaitUseNum = gvar.New(waitUseNum).Int()

	return
}

func (s *sBasicsCoupon) Recycle(ctx context.Context, in *input_basics.PmsCouponRecycleInp) (err error) {
	// 查询优惠券
	var coupon *entity.PmsCoupon
	if err = s.Model(ctx).WherePri(in.Id).Scan(&coupon); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsCoupon.Columns().State:        5,
		dao.PmsCoupon.Columns().OperatorId:   int(contexts.GetUserId(ctx)),
		dao.PmsCoupon.Columns().RecoveryTime: gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "优惠券回收失败，请稍后重试！")
		return
	}

	// 回收后已发放数量要减1
	if _, err = dao.PmsCouponType.Ctx(ctx).Where(dao.PmsCouponType.Columns().Id, coupon.CouponTypeId).Decrement(dao.PmsCouponType.Columns().LeadCount, 1); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsCoupon) InvalidCoupon(ctx context.Context, in *input_basics.PmsCouponInvalidInp) (err error) {
	if _, err = s.Model(ctx).Where(dao.PmsCoupon.Columns().SourceOrderId, in.SourceOrderId).Data(g.Map{
		dao.PmsCoupon.Columns().State:        5,
		dao.PmsCoupon.Columns().RecoveryTime: gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "优惠券回收失败，请稍后重试！")
		return
	}

	return
}
