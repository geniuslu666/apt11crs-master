package logic_th

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type sThCoupon struct{}

func NewThCoupon() *sThCoupon {
	return &sThCoupon{}
}

func init() {
	service.RegisterThCoupon(NewThCoupon())
}

func (s *sThCoupon) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThCoupon.Ctx(ctx), option...)
}

func (s *sThCoupon) List(ctx context.Context, in *input_th.ThCouponListInp) (list []*input_th.ThCouponListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_th.ThCouponListModel{})

	if !g.IsEmpty(in.CouponName) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.CouponName, "coupon_name")
		if err == nil {
			couponIds, _ := service.ThCoupon().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.ThCoupon.Columns().Id, couponIds)
		}
	}

	if !g.IsEmpty(in.IdentityName) {
		mod = mod.WhereLike(dao.ThCoupon.Columns().IdentityName, "%"+in.IdentityName+"%")
	}

	if !g.IsEmpty(in.CategoryId) {
		mod = mod.Where(dao.ThCoupon.Columns().CategoryId, in.CategoryId)
	}

	if in.Status > 0 {
		mod = mod.Where(dao.ThCoupon.Columns().Status, in.Status)
	}

	if in.UseStatus > 0 {
		mod = mod.Where(dao.ThCoupon.Columns().UseStatus, in.UseStatus)
	}

	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.ThCoupon.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderAsc(dao.ThCoupon.Columns().Sort).OrderDesc(dao.ThCoupon.Columns().Id)

	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取礼品券列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取礼品券列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThCoupon) All(ctx context.Context, in *input_th.ThCouponListInp) (list []*input_th.ThCouponAllListModel, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_th.ThCouponAllListModel{})

	if !g.IsEmpty(in.CouponIds) {
		mod = mod.WhereIn(dao.ThCoupon.Columns().Id, strings.Split(in.CouponIds, ","))
	}

	mod = mod.Where(dao.ThCoupon.Columns().Status, 1)

	mod = mod.OrderDesc(dao.ThCoupon.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取礼品券列表失败，请稍后重试！")
		return
	}

	var (
		thCouponList []*input_th.ThCouponAllListModel
	)

	if !g.IsEmpty(in.CouponIds) {
		couponIdArr := strings.Split(in.CouponIds, ",")
		for _, couponId := range couponIdArr {
			for _, v := range list {
				if v.Id == gvar.New(couponId).Int() {
					thCouponList = append(thCouponList, v)
				}
			}
		}
		list = thCouponList
	}

	return
}

func (s *sThCoupon) Edit(ctx context.Context, in *input_th.ThCouponEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object            gdb.Record
			ThCouponDao       *input_language.LoadLanguage
			ThCouponSubDao    *input_language.LoadLanguage
			ThCouponDescDao   *input_language.LoadLanguage
			LanguageStruct    input_language.LanguageModel
			SubLanguageStruct input_language.LanguageModel
			DescStruct        input_language.LanguageModel
			CouponMchInfo     []*entity.ThCouponMch
		)
		Uuid := guid.S([]byte("coupon_name"))
		SubUuid := guid.S([]byte("coupon_sub_name"))

		DescUuid := guid.S([]byte("desc"))

		if in.Id > 0 {

			if Object, err = dao.ThCoupon.Ctx(ctx).Where(dao.ThCoupon.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object["coupon_name"]) {
				Uuid = Object["coupon_name"].String()
			}

			if !g.IsEmpty(Object["coupon_sub_name"]) {
				SubUuid = Object["coupon_sub_name"].String()
			}

			if !g.IsEmpty(Object["desc"]) {
				DescUuid = Object["desc"].String()
			}

			// 名字
			ThCouponDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.ThCoupon.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("CouponName"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, ThCouponDao); err != nil {
				return
			}

			// 副标题
			ThCouponSubDao = &input_language.LoadLanguage{
				Uuid: SubUuid,
				Tag:  dao.ThCoupon.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("CouponSubName"),
			}
			SubLanguageStruct = in.SubNameLanguage
			if err = service.BasicsLanguage().Sync(ctx, SubLanguageStruct, ThCouponSubDao); err != nil {
				return
			}

			ThCouponDescDao = &input_language.LoadLanguage{
				Uuid: DescUuid,
				Tag:  dao.ThCoupon.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Desc"),
			}
			DescStruct = in.DescLanguage
			if err = service.BasicsLanguage().Sync(ctx, DescStruct, ThCouponDescDao); err != nil {
				return
			}

			in.Desc = DescUuid
			in.CouponName = Uuid
			in.CouponSubName = SubUuid

			if !g.IsEmpty(in.MchList) {
				for _, MchInfo := range in.MchList {
					if !g.IsEmpty(MchInfo) {
						MchId := MchInfo.MchId
						CouponMchInfo = append(CouponMchInfo, &entity.ThCouponMch{
							CouponId: int64(in.Id),
							MchId:    int64(MchId),
							Name:     MchInfo.Name,
						})
					}
				}
			}
			if !g.IsEmpty(CouponMchInfo) {
				if _, err = dao.ThCouponMch.Ctx(ctx).Where(dao.ThCouponMch.Columns().CouponId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理适用门店旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.ThCouponMch.Ctx(ctx).OmitEmptyData().Insert(CouponMchInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.ThCouponMch.Ctx(ctx).Where(dao.ThCouponMch.Columns().CouponId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理适用门店旧数据失败，请稍后重试！")
					return
				}
			}

			if _, err = s.Model(ctx).
				Fields(input_th.ThCouponUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改礼品券失败，请稍后重试！")
			}
			return
		}

		var (
			lastInsertId int64
		)
		in.CouponName = Uuid
		in.CouponSubName = SubUuid
		in.Desc = DescUuid
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_th.ThCouponInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增礼品券失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增礼品券失败，请稍后重试！")
			return
		}

		ThCouponDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.ThCoupon.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("CouponName"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, ThCouponDao); err != nil {
			return
		}

		ThCouponSubDao = &input_language.LoadLanguage{
			Uuid: SubUuid,
			Tag:  dao.ThCoupon.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("CouponSubName"),
		}
		SubLanguageStruct = in.SubNameLanguage
		if err = service.BasicsLanguage().Sync(ctx, SubLanguageStruct, ThCouponSubDao); err != nil {
			return
		}

		ThCouponDescDao = &input_language.LoadLanguage{
			Uuid: DescUuid,
			Tag:  dao.ThCoupon.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Desc"),
		}
		DescStruct = in.DescLanguage
		if err = service.BasicsLanguage().Sync(ctx, DescStruct, ThCouponDescDao); err != nil {
			return
		}

		if !g.IsEmpty(in.MchList) {
			for _, MchInfo := range in.MchList {
				if !g.IsEmpty(MchInfo) {
					MchId := MchInfo.MchId
					CouponMchInfo = append(CouponMchInfo, &entity.ThCouponMch{
						CouponId: lastInsertId,
						MchId:    int64(MchId),
						Name:     MchInfo.Name,
					})
				}
			}
		}

		if !g.IsEmpty(CouponMchInfo) {
			if _, err = dao.ThCouponMch.Ctx(ctx).Where(dao.ThCouponMch.Columns().CouponId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理适用门店旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.ThCouponMch.Ctx(ctx).OmitEmptyData().Insert(CouponMchInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.ThCouponMch.Ctx(ctx).Where(dao.ThCouponMch.Columns().CouponId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理适用门店旧数据失败，请稍后重试！")
				return
			}
		}

		return
	})
}

func (s *sThCoupon) Delete(ctx context.Context, in *input_th.ThCouponDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除礼品券失败，请稍后重试！")
		return
	}
	return
}

func (s *sThCoupon) MaxSort(ctx context.Context, in *input_th.ThCouponMaxSortInp) (res *input_th.ThCouponMaxSortModel, err error) {
	if err = dao.ThCoupon.Ctx(ctx).Fields(dao.ThCoupon.Columns().Sort).OrderDesc(dao.ThCoupon.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取礼品券最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_th.ThCouponMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sThCoupon) View(ctx context.Context, in *input_th.ThCouponViewInp) (res *input_th.ThCouponViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取礼品券信息，请稍后重试！")
		return
	}

	if res.NeedReservation == 1 && !g.IsEmpty(res.ReservationRestaurantIds) {
		res.RestaurantNames = ""
		restaurantArr, _ := dao.FoodRestaurant.Ctx(ctx).Hook(hook2.PmsFindLanguageValueHook).Fields("name").
			WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(res.ReservationRestaurantIds, ",")).Array()

		if !g.IsEmpty(restaurantArr) {
			// 获取餐厅名称以逗号 分隔
			for _, restaurant := range restaurantArr {
				res.RestaurantNames = res.RestaurantNames + g.NewVar(restaurant).String() + "，"
			}
			res.RestaurantNames = gstr.TrimRight(res.RestaurantNames, "，")
		}
	}

	return
}

func (s *sThCoupon) AppView(ctx context.Context, in *input_th.ThCouponAppViewInp) (res *input_th.ThCouponAppViewModel, err error) {
	if err = s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取礼品券信息，请稍后重试！")
		return
	}

	// 提前查询首页活动信息（用于周末限制判断）
	var IndexActivity *entity.HomepageArticles
	if !g.IsEmpty(in.IndexActivityId) {
		if err = dao.HomepageArticles.Ctx(ctx).WherePri(in.IndexActivityId).Scan(&IndexActivity); err != nil {
			err = gerror.Wrap(err, "获取活动信息失败，请稍后重试！")
			return
		}
	}

	if res.NeedReservation == 1 && !g.IsEmpty(res.ReservationRestaurantIds) {
		if err = dao.FoodRestaurant.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Fields("id,name").
			WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(res.ReservationRestaurantIds, ",")).Scan(&res.RestaurantList); err != nil {
			err = gerror.Wrap(err, "获取餐厅列表信息失败，请稍后重试！")
			return
		}
	}

	if !g.IsEmpty(in.ActivityId) {
		var (
			MemberInfo             *model.MemberIdentity
			EmployeeActivity       *entity.EmployeeActivity
			EmployeeActivityCoupon *entity.EmployeeActivityCoupon
		)
		MemberInfo = contexts.GetMemberUser(ctx)
		if err = dao.EmployeeActivity.Ctx(ctx).Where(dao.EmployeeActivity.Columns().Id, in.ActivityId).Scan(&EmployeeActivity); err != nil {
			err = gerror.Wrap(err, "获取活动信息，请稍后重试！")
			return
		}
		if g.IsEmpty(EmployeeActivity) {
			err = gerror.New("活动信息错误，请稍后重试！")
			return
		}

		res.ActivityId = int(EmployeeActivity.Id)
		res.ActivityStatus = EmployeeActivity.Status

		if err = dao.EmployeeActivityCoupon.Ctx(ctx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.ActivityId).Where(dao.EmployeeActivityCoupon.Columns().CouponId, in.Id).Scan(&EmployeeActivityCoupon); err != nil {
			err = gerror.Wrap(err, "获取活动信息，请稍后重试！")
			return
		}
		if g.IsEmpty(EmployeeActivityCoupon) {
			err = gerror.New("活动信息错误，请稍后重试！")
			return
		}
		res.AvailableQuantity = EmployeeActivityCoupon.AvailableQuantity
		res.PerDayAvailable = EmployeeActivityCoupon.PerDayAvailable
		res.LimitDays = EmployeeActivityCoupon.LimitDays

		res.HaveReceived, err = dao.ThMemberCoupon.Ctx(ctx).
			Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
			Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
			Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
			Where(dao.ThMemberCoupon.Columns().Source, 3).
			Count()
		if err != nil {
			err = gerror.Wrap(err, "获取活动信息，请稍后重试！")
			return
		}

		// 当前周期内已领取数量（以第一次领取日期为起点，每N天为一个周期）
		if EmployeeActivityCoupon.LimitDays > 0 {
			var firstReceiveDateRow *struct{ D string }
			_ = dao.ThMemberCoupon.Ctx(ctx).
				Fields("DATE("+dao.ThMemberCoupon.Columns().CreateAt+") as d").
				Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
				Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
				Where(dao.ThMemberCoupon.Columns().Source, 3).
				OrderAsc(dao.ThMemberCoupon.Columns().CreateAt).
				Limit(1).
				Scan(&firstReceiveDateRow)
			var firstReceiveDate string
			if firstReceiveDateRow != nil {
				firstReceiveDate = firstReceiveDateRow.D
			}
			if firstReceiveDate != "" {
				firstDate, _ := gtime.StrToTime(firstReceiveDate)
				nowDate, _ := gtime.StrToTime(gtime.Now().Format("Y-m-d"))
				daysDiff := int(nowDate.Sub(firstDate).Hours() / 24)
				periodIndex := daysDiff / EmployeeActivityCoupon.LimitDays
				periodStart := firstDate.AddDate(0, 0, periodIndex*EmployeeActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
				periodEnd := firstDate.AddDate(0, 0, (periodIndex+1)*EmployeeActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
				res.TodayReceive, err = dao.ThMemberCoupon.Ctx(ctx).
					Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
					Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
					Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
					Where(dao.ThMemberCoupon.Columns().Source, 3).
					WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, periodStart).
					WhereLT(dao.ThMemberCoupon.Columns().CreateAt, periodEnd).
					Count()
				if err != nil {
					err = gerror.Wrap(err, "获取周期内已领取数量失败，请稍后重试！")
					return
				}
			}
		}
	}

	if !g.IsEmpty(in.IndexActivityId) {
		var (
			MemberInfo          *model.MemberIdentity
			IndexActivityCoupon *entity.HomepageArticleCoupon
		)
		MemberInfo = contexts.GetMemberUser(ctx)
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

		// 查找所有相同编号的活动ID（防止用户切换同一活动的不同语言版本后重复领券）
		var sameNoActivityIds []int64
		if IndexActivity.No != "" {
			var sameNoRows []struct{ Id int64 }
			_ = dao.HomepageArticles.Ctx(ctx).
				Fields(dao.HomepageArticles.Columns().Id).
				Where(dao.HomepageArticles.Columns().No, IndexActivity.No).
				Scan(&sameNoRows)
			for _, row := range sameNoRows {
				sameNoActivityIds = append(sameNoActivityIds, row.Id)
			}
		}
		if len(sameNoActivityIds) == 0 {
			sameNoActivityIds = []int64{int64(in.IndexActivityId)}
		}

		if err = dao.HomepageArticleCoupon.Ctx(ctx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).Where(dao.HomepageArticleCoupon.Columns().CouponType, "thCoupon").Where(dao.HomepageArticleCoupon.Columns().CouponId, in.Id).Scan(&IndexActivityCoupon); err != nil {
			err = gerror.Wrap(err, "获取活动礼品券信息失败，请稍后重试！")
			return
		}
		if g.IsEmpty(IndexActivityCoupon) {
			err = gerror.New("获取活动礼品券信息失败，请稍后重试！")
			return
		}
		res.AvailableQuantity = IndexActivityCoupon.AvailableQuantity
		res.PerDayAvailable = IndexActivityCoupon.PerDayAvailable
		res.LimitDays = IndexActivityCoupon.LimitDays

		res.HaveReceived, err = dao.ThMemberCoupon.Ctx(ctx).
			Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
			Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
			WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
			Where(dao.ThMemberCoupon.Columns().Source, 4).
			Count()
		if err != nil {
			err = gerror.Wrap(err, "获取活动信息失败，请稍后重试！")
			return
		}

		// 当前周期内已领取数量（以第一次领取日期为起点，每N天为一个周期）
		if IndexActivityCoupon.LimitDays > 0 {
			var firstReceiveDateRow *struct{ D string }
			err = dao.ThMemberCoupon.Ctx(ctx).
				Fields("DATE("+dao.ThMemberCoupon.Columns().CreateAt+") as d").
				Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
				WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
				Where(dao.ThMemberCoupon.Columns().Source, 4).
				OrderAsc(dao.ThMemberCoupon.Columns().CreateAt).
				Limit(1).
				Scan(&firstReceiveDateRow)
			var firstReceiveDate string
			if firstReceiveDateRow != nil {
				firstReceiveDate = firstReceiveDateRow.D
			}
			if firstReceiveDate != "" {
				firstDate, _ := gtime.StrToTime(firstReceiveDate)
				nowDate, _ := gtime.StrToTime(gtime.Now().Format("Y-m-d"))
				daysDiff := int(nowDate.Sub(firstDate).Hours() / 24)
				periodIndex := daysDiff / IndexActivityCoupon.LimitDays
				periodStart := firstDate.AddDate(0, 0, periodIndex*IndexActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
				periodEnd := firstDate.AddDate(0, 0, (periodIndex+1)*IndexActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
				res.TodayReceive, err = dao.ThMemberCoupon.Ctx(ctx).
					Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
					Where(dao.ThMemberCoupon.Columns().CouponId, in.Id).
					WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
					Where(dao.ThMemberCoupon.Columns().Source, 4).
					WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, periodStart).
					WhereLT(dao.ThMemberCoupon.Columns().CreateAt, periodEnd).
					Count()
				if err != nil {
					err = gerror.Wrap(err, "获取周期内已领取数量失败，请稍后重试！")
					return
				}
			}
		}
	}

	// 判断礼品券状态
	var status int
	// 1. 先判断用户是否已达到个人总领取限制
	if res.AvailableQuantity > 0 && res.HaveReceived >= res.AvailableQuantity {
		status = 2 // 已领取
	} else {
		// 2. 判断N天周期内限制（如果设置了limitDays和perDayAvailable）
		if res.LimitDays > 0 && res.PerDayAvailable > 0 && res.TodayReceive >= res.PerDayAvailable {
			status = 4 // 周期内已领满
		} else {
			// 3. 判断周末限制（仅对首页活动中的礼品券生效）
			if IndexActivity != nil && IndexActivity.LimitWeek != "" {
				// 获取当前日期的星期几（0=周日, 1=周一, ..., 6=周六）
				currentWeekday := gtime.Now().Weekday()
				limits := gstr.Split(IndexActivity.LimitWeek, ",")
				// 如果是周六（6）且限制中包含6，或者周日（0）且限制中包含7
				if (currentWeekday == 6 && gstr.InArray(limits, "6")) || (currentWeekday == 0 && gstr.InArray(limits, "7")) {
					status = 5 // 今日不可领取
				} else {
					status = 1 // 可领取
				}
			} else {
				status = 1 // 可领取
			}
		}
	}
	res.Status = status

	return
}

func (s *sThCoupon) Status(ctx context.Context, in *input_th.ThCouponStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.ThCoupon.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新礼品券发行状态失败，请稍后重试！")
		return
	}
	return
}

func (s *sThCoupon) UseStatus(ctx context.Context, in *input_th.ThCouponUseStatusInp) (err error) {
	if in.Status == 2 {
		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThCoupon.Columns().Status:    in.Status,
			dao.ThCoupon.Columns().UseStatus: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新礼品券使用状态失败，请稍后重试！")
			return
		}
	} else {
		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThCoupon.Columns().UseStatus: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新礼品券使用状态失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sThCoupon) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("coupon_name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}

func (s *sThCoupon) SendMemberCoupon(ctx context.Context, in *input_th.ThSendMemberCouponInp, source int) (err error) {
	var (
		couponInfo *entity.ThCoupon
		CouponNo   string
		TimeDiff   int
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.CouponId > 0 {

			if err = s.Model(ctx).Where(dao.ThCoupon.Columns().Id, in.CouponId).Scan(&couponInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(couponInfo) {
				err = gerror.New("礼品券不存在")
				return
			}

			if couponInfo.Status != 1 {
				err = gerror.New("该礼品券已停止发行")
				return
			}

			var endTime *gtime.Time

			todayStart := in.StartTime

			//if gtime.Now().After(gtime.New(todayStart)) {
			//	err = gerror.New("启用时间不能小于当前时间")
			//	return
			//}

			CouponNo = uuid.CreateOrderCode(couponInfo.CouponNoPrefix)

			TimeDiff = int(gtime.New(todayStart).Sub(gtime.Now()).Seconds())

			endTime = gtime.New(todayStart).Add(time.Duration(24*couponInfo.FixedTerm) * time.Hour)
			endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)

			if _, err = g.Model(dao.ThMemberCoupon.Table()).Ctx(ctx).Safe().
				Fields(input_th.ThMemberCouponInsertFields{}).
				Data(g.MapStrAny{
					dao.ThMemberCoupon.Columns().CouponNo:      CouponNo,
					dao.ThMemberCoupon.Columns().CouponId:      couponInfo.Id,
					dao.ThMemberCoupon.Columns().MemberId:      in.MemberId,
					dao.ThMemberCoupon.Columns().State:         1,
					dao.ThMemberCoupon.Columns().StartTime:     gtime.New(todayStart).Format("Y-m-d H:i:s"),
					dao.ThMemberCoupon.Columns().EndTime:       endTime,
					dao.ThMemberCoupon.Columns().Source:        source,
					dao.ThMemberCoupon.Columns().CountDown:     TimeDiff,
					dao.ThMemberCoupon.Columns().SourceOrderId: in.SourceOrderId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "发放礼品券失败，请稍后重试！")
				return
			}

			if _, err = s.Model(ctx).WherePri(in.CouponId).Increment(dao.ThCoupon.Columns().Count, 1); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			return
		}

		err = gerror.Wrap(err, "礼品券不存在，请稍后重试！")
		return
	}); err != nil {
		return
	}

	// 投递自动生效礼品券队列
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeDelayedName,
		QueueName:    consts.RabbitMQQueueNameThCouponEffect,
		DataByte:     gvar.New(CouponNo).Bytes(),
		Header: amqp.Table{
			"x-delay": gvar.New(TimeDiff * 1000).String(),
		},
	}); err != nil {
		g.Log().Error(ctx, "发送自动生效礼品券MQ失败", err)
	}

	return
}

func (s *sThCoupon) SendMemberCouponGetId(ctx context.Context, in *input_th.ThSendMemberCouponInp, source int) (memberCouponId int64, couponNameLanguage []*input_hotel.LanguageType, err error) {
	var (
		ActivityInfo *entity.EmployeeActivity
		couponInfo   *struct {
			entity.ThCoupon
			NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=coupon_name"`
		}
		CouponNo string
		TimeDiff int
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.ActivityId > 0 {
			if err = dao.EmployeeActivity.Ctx(ctx).Where(dao.EmployeeActivity.Columns().Id, in.ActivityId).Scan(&ActivityInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(ActivityInfo) {
				err = gerror.New(gi18n.T(ctx, "activity_not_exist"))
				return
			}
		}

		if in.CouponId > 0 {

			if err = s.Model(ctx).WithAll().Where(dao.ThCoupon.Columns().Id, in.CouponId).Scan(&couponInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(couponInfo) {
				err = gerror.New("礼品券不存在")
				return
			}

			if couponInfo.Status != 1 {
				err = gerror.New("该礼品券已停止发行")
				return
			}

			couponNameLanguage = couponInfo.NameLanguage

			var endTime *gtime.Time

			todayStart := in.StartTime
			if in.ActivityId > 0 {
				if ActivityInfo.CouponValidity == 1 {
					// 跟随活动
					todayStart = gtime.New(ActivityInfo.StartTime).String()
				}
			}

			//if gtime.Now().After(gtime.New(todayStart)) {
			//	err = gerror.New("启用时间不能小于当前时间")
			//	return
			//}

			CouponNo = uuid.CreateOrderCode(couponInfo.CouponNoPrefix)

			TimeDiff = int(gtime.New(todayStart).Sub(gtime.Now()).Seconds())

			endTime = gtime.New(todayStart).Add(time.Duration(24*couponInfo.FixedTerm) * time.Hour)
			endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)

			if in.ActivityId > 0 {
				if ActivityInfo.CouponValidity == 1 {
					// 跟随活动
					endTime = ActivityInfo.EndTime
				}
			}

			var CouponStatus int
			if TimeDiff > 0 {
				CouponStatus = 1
			} else {
				CouponStatus = 2
			}
			if memberCouponId, err = g.Model(dao.ThMemberCoupon.Table()).Ctx(ctx).Safe().
				Fields(input_th.ThMemberCouponInsertFields{}).
				Data(g.MapStrAny{
					dao.ThMemberCoupon.Columns().CouponNo:      CouponNo,
					dao.ThMemberCoupon.Columns().CouponId:      couponInfo.Id,
					dao.ThMemberCoupon.Columns().MemberId:      in.MemberId,
					dao.ThMemberCoupon.Columns().State:         CouponStatus,
					dao.ThMemberCoupon.Columns().StartTime:     gtime.New(todayStart).Format("Y-m-d H:i:s"),
					dao.ThMemberCoupon.Columns().EndTime:       endTime,
					dao.ThMemberCoupon.Columns().Source:        source,
					dao.ThMemberCoupon.Columns().CountDown:     TimeDiff,
					dao.ThMemberCoupon.Columns().SourceOrderId: in.SourceOrderId,
					dao.ThMemberCoupon.Columns().ActivityId:    in.ActivityId,
					dao.ThMemberCoupon.Columns().EmployeeId:    in.EmployeeId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, "发放礼品券失败，请稍后重试！")
				return
			}

			if _, err = s.Model(ctx).WherePri(in.CouponId).Increment(dao.ThCoupon.Columns().Count, 1); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			if in.ActivityId > 0 {
				if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.ActivityId).Where(dao.EmployeeActivityCoupon.Columns().CouponId, in.CouponId).Increment(dao.EmployeeActivityCoupon.Columns().TotalReceived, 1); err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
					return
				}
			}

			return
		}

		err = gerror.Wrap(err, "礼品券不存在，请稍后重试！")
		return
	}); err != nil {
		return
	}

	if TimeDiff > 0 {
		// 投递自动生效礼品券队列
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeDelayedName,
			QueueName:    consts.RabbitMQQueueNameThCouponEffect,
			DataByte:     gvar.New(CouponNo).Bytes(),
			Header: amqp.Table{
				"x-delay": gvar.New(TimeDiff * 1000).String(),
			},
		}); err != nil {
			g.Log().Error(ctx, "发送自动生效礼品券MQ失败", err)
		}
	}

	return
}

func (s *sThCoupon) AppReceiveCoupon(ctx context.Context, in *input_th.ThCouponAppReceiveInp) (err error) {
	var (
		IndexActivity  *entity.HomepageArticles
		CouponInfo     *entity.ThCoupon
		MemberInfo     *model.MemberIdentity
		CouponNo       string
		TimeDiff       int
		ActivityCoupon *entity.HomepageArticleCoupon
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

		if in.CouponId > 0 {

			if in.IndexActivityId > 0 {
				// 判断活动是否绑定这个券
				var couponCount int
				couponCount, err = dao.HomepageArticleCoupon.Ctx(ctx).
					Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).
					Where(dao.HomepageArticleCoupon.Columns().CouponType, "thCoupon").
					Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponId).
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

				// 查找所有相同编号的活动ID（防止用户切换同一活动的不同语言版本后重复领券）
				var sameNoActivityIds []int64
				if IndexActivity.No != "" {
					var sameNoRows []struct{ Id int64 }
					_ = dao.HomepageArticles.Ctx(ctx).
						Fields(dao.HomepageArticles.Columns().Id).
						Where(dao.HomepageArticles.Columns().No, IndexActivity.No).
						Scan(&sameNoRows)
					for _, row := range sameNoRows {
						sameNoActivityIds = append(sameNoActivityIds, row.Id)
					}
				}
				if len(sameNoActivityIds) == 0 {
					sameNoActivityIds = []int64{int64(in.IndexActivityId)}
				}

				// 判断已经领取的数量
				var receiveCouponCount int
				receiveCouponCount, err = dao.ThMemberCoupon.Ctx(ctx).
					Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
					Where(dao.ThMemberCoupon.Columns().Source, 4).
					WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
					Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId).
					Count()
				if err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
					return
				}

				if err = dao.HomepageArticleCoupon.Ctx(ctx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).Where(dao.HomepageArticleCoupon.Columns().CouponType, "thCoupon").Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponId).Scan(&ActivityCoupon); err != nil {
					return
				}
				if ActivityCoupon.AvailableQuantity > 0 && receiveCouponCount >= ActivityCoupon.AvailableQuantity {
					err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_limit"))
					return
				}

				// 判断周期内领取数量
				if ActivityCoupon.LimitDays > 0 && ActivityCoupon.PerDayAvailable > 0 {
					var firstReceiveDateRow *struct{ D string }
					_ = dao.ThMemberCoupon.Ctx(ctx).
						Fields("DATE("+dao.ThMemberCoupon.Columns().CreateAt+") as d").
						Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
						Where(dao.ThMemberCoupon.Columns().Source, 4).
						WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
						Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId).
						OrderAsc(dao.ThMemberCoupon.Columns().CreateAt).
						Limit(1).
						Scan(&firstReceiveDateRow)
					var firstReceiveDate string
					if firstReceiveDateRow != nil {
						firstReceiveDate = firstReceiveDateRow.D
					}
					if firstReceiveDate != "" {
						firstDate, _ := gtime.StrToTime(firstReceiveDate)
						nowDate, _ := gtime.StrToTime(gtime.Now().Format("Y-m-d"))
						daysDiff := int(nowDate.Sub(firstDate).Hours() / 24)
						periodIndex := daysDiff / ActivityCoupon.LimitDays
						periodStart := firstDate.AddDate(0, 0, periodIndex*ActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
						periodEnd := firstDate.AddDate(0, 0, (periodIndex+1)*ActivityCoupon.LimitDays).Format("Y-m-d") + " 00:00:00"
						var periodReceiveCount int
						periodReceiveCount, err = dao.ThMemberCoupon.Ctx(ctx).
							Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
							Where(dao.ThMemberCoupon.Columns().Source, 4).
							WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
							Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId).
							WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, periodStart).
							WhereLT(dao.ThMemberCoupon.Columns().CreateAt, periodEnd).
							Count()
						if err != nil {
							err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
							return
						}
						if periodReceiveCount >= ActivityCoupon.PerDayAvailable {
							err = gerror.New(gi18n.T(ctx, "validate_coupon_period_limit"))
							return
						}
					}
				}

				// 判断首页活动是否限制周六日领取
				if IndexActivity.LimitWeek != "" {
					// 获取当前日期的星期几（0=周日, 1=周一, ..., 6=周六）
					currentWeekday := gtime.Now().Weekday()
					limits := gstr.Split(IndexActivity.LimitWeek, ",")
					// 如果是周六（6）且限制中包含6，或者周日（0）且限制中包含7
					if (currentWeekday == 6 && gstr.InArray(limits, "6")) || (currentWeekday == 0 && gstr.InArray(limits, "7")) {
						err = gerror.New(gi18n.T(ctx, "today_not_available"))
						return
					}
				}
			}

			if err = dao.ThCoupon.Ctx(ctx).Where(dao.ThCoupon.Columns().Id, in.CouponId).Scan(&CouponInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}

			if g.IsEmpty(CouponInfo) {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_failed"))
				return
			}

			if CouponInfo.Status != 1 {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_stop"))
				return
			}

			var State int

			todayStart := gtime.Now()

			CouponNo = uuid.CreateOrderCode(CouponInfo.CouponNoPrefix)

			TimeDiff = int(gtime.New(todayStart).Sub(gtime.Now()).Seconds())

			if TimeDiff <= 0 {
				State = 2
				TimeDiff = 0
			} else {
				State = 1
			}

			var endTime *gtime.Time
			var memberCouponSource int
			if in.IndexActivityId > 0 {
				memberCouponSource = 4
			} else {
				memberCouponSource = 5
			}

			// 优先使用首页活动的FixedTerm，如果没有则使用礼品券本身的FixedTerm
			var fixedTerm int
			if ActivityCoupon != nil && ActivityCoupon.FixedTerm >= 0 {
				fixedTerm = ActivityCoupon.FixedTerm
			} else {
				fixedTerm = CouponInfo.FixedTerm
			}

			// 计算有效期：0表示领取当日23:59:59，1表示次日23:59:59，以此类推
			endTime = gtime.New(todayStart).Add(time.Duration(24*fixedTerm) * time.Hour)
			endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)

			if _, err = g.Model(dao.ThMemberCoupon.Table()).Ctx(ctx).Safe().
				Data(g.MapStrAny{
					dao.ThMemberCoupon.Columns().CouponNo:        CouponNo,
					dao.ThMemberCoupon.Columns().CouponId:        CouponInfo.Id,
					dao.ThMemberCoupon.Columns().MemberId:        MemberInfo.Id,
					dao.ThMemberCoupon.Columns().Source:          memberCouponSource,
					dao.ThMemberCoupon.Columns().State:           State,
					dao.ThMemberCoupon.Columns().StartTime:       gtime.New(todayStart).Format("Y-m-d H:i:s"),
					dao.ThMemberCoupon.Columns().EndTime:         endTime,
					dao.ThMemberCoupon.Columns().CountDown:       TimeDiff,
					dao.ThMemberCoupon.Columns().IndexActivityId: in.IndexActivityId,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "receive_coupon_failed"))
				return
			}

			if _, err = dao.ThCoupon.Ctx(ctx).WherePri(in.CouponId).Increment(dao.ThCoupon.Columns().Count, 1); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
				return
			}

			if in.IndexActivityId > 0 {
				if _, err = dao.HomepageArticleCoupon.Ctx(ctx).
					Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.IndexActivityId).
					Where(dao.HomepageArticleCoupon.Columns().CouponType, "thCoupon").
					Where(dao.HomepageArticleCoupon.Columns().CouponId, in.CouponId).
					Increment(dao.HomepageArticleCoupon.Columns().TotalReceived, 1); err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
					return
				}
			}

			return
		}

		err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_failed"))
		return
	}); err != nil {
		return
	}

	// 投递自动生效礼品券队列
	if TimeDiff > 0 {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeDelayedName,
			QueueName:    consts.RabbitMQQueueNameThCouponEffect,
			DataByte:     gvar.New(CouponNo).Bytes(),
			Header: amqp.Table{
				"x-delay": gvar.New(TimeDiff * 1000).String(),
			},
		}); err != nil {
			g.Log().Error(ctx, "发送自动生效礼品券MQ失败", err)
		}
	}

	return
}
