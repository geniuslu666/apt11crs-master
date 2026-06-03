package logic_employee

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_employee"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	hook2 "APT/internal/library/hgorm/hook"

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

type sEmployeeActivity struct{}

func NewEmployeeActivity() *sEmployeeActivity {
	return &sEmployeeActivity{}
}

func init() {
	service.RegisterEmployeeActivity(NewEmployeeActivity())
}

func (s *sEmployeeActivity) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.EmployeeActivity.Ctx(ctx), option...)
}

func (s *sEmployeeActivity) List(ctx context.Context, in *input_employee.EmployeeActivityListInp) (list []*input_employee.EmployeeActivityListModel, totalCount int, err error) {

	mod := s.Model(ctx)

	mod = mod.Fields(input_employee.EmployeeActivityListModel{}).WithAll()

	if !g.IsEmpty(in.Name) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.Name, "name")
		if err == nil {
			activityIds, _ := service.EmployeeActivity().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.EmployeeActivity.Columns().Id, activityIds)
		}
	}

	if in.Status > 0 {
		mod = mod.Where(dao.EmployeeActivity.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.EmployeeActivity.Columns().Id)

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取活动表列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取活动表列表失败，请稍后重试！")
			return
		}
	}

	// 处理日期字段，只显示日期不显示时间
	for _, v := range list {
		if v.StartTime != nil {
			v.StartDate = v.StartTime.Format("Y-m-d")
		}
		if v.EndTime != nil {
			v.EndDate = v.EndTime.Format("Y-m-d")
		}
	}

	return
}

func (s *sEmployeeActivity) GetIds(ctx context.Context, name []string) (ids []int, err error) {
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

func (s *sEmployeeActivity) Edit(ctx context.Context, in *input_employee.EmployeeActivityEditInp) (err error) {
	if in.RestrictionType == 2 && g.IsEmpty(in.DepartmentIds) {
		err = gerror.New("请选择部门！")
		return
	}
	if in.RestrictionType == 3 && g.IsEmpty(in.EmployeeIds) {
		err = gerror.New("请选择员工！")
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object                         gdb.Record
			EmployeeActivityNameDao        *input_language.LoadLanguage
			EmployeeActivityDescriptionDao *input_language.LoadLanguage
			LanguageStruct                 input_language.LanguageModel
			ActivityCouponInfo             []*entity.EmployeeActivityCoupon
			ActivityDepartmentInfo         []*entity.EmployeeActivityDepartment
			ActivityEmployeeInfo           []*entity.EmployeeActivityEmployee
		)
		NameUuid := guid.S([]byte("name"))
		DescriptionUuid := guid.S([]byte("description"))

		if in.Id > 0 {
			if Object, err = dao.EmployeeActivity.Ctx(ctx).Where(dao.EmployeeActivity.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
				NameUuid = Object["name"].String()
			}
			if !g.IsEmpty(Object[gstr.ToLower("Description")]) {
				DescriptionUuid = Object["description"].String()
			}

			EmployeeActivityNameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.EmployeeActivity.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, EmployeeActivityNameDao); err != nil {
				return
			}

			EmployeeActivityDescriptionDao = &input_language.LoadLanguage{
				Uuid: DescriptionUuid,
				Tag:  dao.EmployeeActivity.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Description"),
			}
			LanguageStruct = in.DescriptionLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, EmployeeActivityDescriptionDao); err != nil {
				return
			}

			in.Name = NameUuid
			in.Description = DescriptionUuid

			if _, err = s.Model(ctx).TX(tx).
				Fields(input_employee.EmployeeActivityUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改活动失败，请稍后重试！")
			}

			if !g.IsEmpty(in.ThCouponsArr) {
				for _, ThCoupon := range in.ThCouponsArr {
					if !g.IsEmpty(ThCoupon) {
						ThCouponId := ThCoupon.CouponId
						thLimitDays := ThCoupon.LimitDays
						if thLimitDays <= 0 {
							thLimitDays = 1
						}
						ActivityCouponInfo = append(ActivityCouponInfo, &entity.EmployeeActivityCoupon{
							ActivityId:        uint64(in.Id),
							CouponId:          uint64(ThCouponId),
							AvailableQuantity: ThCoupon.AvailableQuantity,
							PerDayAvailable:   ThCoupon.PerDayAvailable,
							PerDayVerify:      ThCoupon.PerDayVerify,
							LimitDays:         thLimitDays,
						})
					}
				}

				if !g.IsEmpty(ActivityCouponInfo) {
					if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理活动提货券旧数据失败，请稍后重试！")
						return
					}
					if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityCouponInfo); err != nil {
						return err
					}
				} else {
					if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理活动提货券旧数据失败，请稍后重试！")
						return
					}
				}
			}

			if in.RestrictionType == 2 && !g.IsEmpty(in.DepartmentIds) {
				for _, Department := range in.DepartmentIds {
					if !g.IsEmpty(Department) {
						ActivityDepartmentInfo = append(ActivityDepartmentInfo, &entity.EmployeeActivityDepartment{
							ActivityId:   uint64(in.Id),
							DepartmentId: uint64(Department),
						})
					}
				}

				if _, err = dao.EmployeeActivityDepartment.Ctx(ctx).Where(dao.EmployeeActivityDepartment.Columns().ActivityId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理活动部门旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.EmployeeActivityEmployee.Ctx(ctx).Where(dao.EmployeeActivityEmployee.Columns().ActivityId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理活动员工旧数据失败，请稍后重试！")
					return
				}

				if !g.IsEmpty(ActivityDepartmentInfo) {
					if _, err = dao.EmployeeActivityDepartment.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityDepartmentInfo); err != nil {
						return err
					}
				}
			}

			if in.RestrictionType == 3 && !g.IsEmpty(in.EmployeeIds) {
				for _, Employee := range in.EmployeeIds {
					if !g.IsEmpty(Employee) {
						ActivityEmployeeInfo = append(ActivityEmployeeInfo, &entity.EmployeeActivityEmployee{
							ActivityId: uint64(in.Id),
							EmployeeId: uint64(Employee),
						})
					}
				}

				if _, err = dao.EmployeeActivityDepartment.Ctx(ctx).TX(tx).Where(dao.EmployeeActivityDepartment.Columns().ActivityId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理活动部门旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.EmployeeActivityEmployee.Ctx(ctx).TX(tx).Where(dao.EmployeeActivityEmployee.Columns().ActivityId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理活动员工旧数据失败，请稍后重试！")
					return
				}

				if !g.IsEmpty(ActivityEmployeeInfo) {
					if _, err = dao.EmployeeActivityEmployee.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityEmployeeInfo); err != nil {
						return err
					}
				}
			}

			return
		}

		var (
			lastInsertId int64
		)

		in.Name = NameUuid
		in.Description = DescriptionUuid

		// 处理开始和结束时间
		if in.ValidityType == 1 {
			// 检查开始和结束时间是否都有传入
			if in.StartTime == nil {
				err = gerror.New("请选择活动开始时间！")
				return
			}
			if in.EndTime == nil {
				err = gerror.New("请选择活动结束时间！")
				return
			}

			in.EndTime = gtime.New(in.EndTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)

			// 检查开始时间是否小于结束时间
			if in.StartTime.After(in.EndTime) {
				err = gerror.New("活动开始时间不能大于结束时间！")
				return
			}

			// 根据当前时间与开始结束时间的关系设置状态
			now := gtime.Now()
			if now.Before(in.StartTime) {
				// 当前时间小于开始时间，状态为未开始
				in.Status = 1
			} else if now.After(in.EndTime) {
				// 当前时间大于结束时间，状态为已结束
				in.Status = 3
			} else {
				// 当前时间在开始和结束时间之间，状态为进行中
				in.Status = 2
			}
		} else {
			in.Status = 2
		}

		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).TX(tx).
			Fields(input_employee.EmployeeActivityInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增活动失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		EmployeeActivityNameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.EmployeeActivity.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, EmployeeActivityNameDao); err != nil {
			return
		}

		EmployeeActivityDescriptionDao = &input_language.LoadLanguage{
			Uuid: DescriptionUuid,
			Tag:  dao.EmployeeActivity.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Description"),
		}
		LanguageStruct = in.DescriptionLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, EmployeeActivityDescriptionDao); err != nil {
			return
		}

		if !g.IsEmpty(in.ThCouponsArr) {
			for _, ThCoupon := range in.ThCouponsArr {
				if !g.IsEmpty(ThCoupon) {
					ThCouponId := ThCoupon.CouponId
					thLimitDays := ThCoupon.LimitDays
					if thLimitDays <= 0 {
						thLimitDays = 1
					}
					ActivityCouponInfo = append(ActivityCouponInfo, &entity.EmployeeActivityCoupon{
						ActivityId:        uint64(lastInsertId),
						CouponId:          uint64(ThCouponId),
						AvailableQuantity: ThCoupon.AvailableQuantity,
						PerDayAvailable:   ThCoupon.PerDayAvailable,
						PerDayVerify:      ThCoupon.PerDayVerify,
						LimitDays:         thLimitDays,
					})
				}
			}
		}
		if !g.IsEmpty(ActivityCouponInfo) {
			if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityCouponInfo); err != nil {
				return err
			}
		}

		if in.RestrictionType == 2 && !g.IsEmpty(in.DepartmentIds) {
			for _, EmployeeDepartment := range in.DepartmentIds {
				if !g.IsEmpty(EmployeeDepartment) {
					ActivityDepartmentInfo = append(ActivityDepartmentInfo, &entity.EmployeeActivityDepartment{
						ActivityId:   uint64(lastInsertId),
						DepartmentId: uint64(EmployeeDepartment),
					})
				}
			}
		}
		if !g.IsEmpty(ActivityDepartmentInfo) {
			if _, err = dao.EmployeeActivityDepartment.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityDepartmentInfo); err != nil {
				return err
			}
		}

		if in.RestrictionType == 3 && !g.IsEmpty(in.EmployeeIds) {
			for _, Employee := range in.EmployeeIds {
				if !g.IsEmpty(Employee) {
					ActivityEmployeeInfo = append(ActivityEmployeeInfo, &entity.EmployeeActivityEmployee{
						ActivityId: uint64(lastInsertId),
						EmployeeId: uint64(Employee),
					})
				}
			}
		}
		if !g.IsEmpty(ActivityEmployeeInfo) {
			if _, err = dao.EmployeeActivityEmployee.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ActivityEmployeeInfo); err != nil {
				return err
			}
		}

		// 未开始的活动进入队列
		if in.Status == 1 {
			var TimeDiff int
			TimeDiff = int(gtime.New(in.StartTime).Sub(gtime.Now()).Seconds())
			if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeDelayedName,
				QueueName:    consts.RabbitMQQueueNameEmployeeActivityEffect,
				DataByte:     gvar.New(lastInsertId).Bytes(),
				Header: amqp.Table{
					"x-delay": gvar.New(TimeDiff * 1000).String(),
				},
			}); err != nil {
				g.Log().Error(ctx, "发送自动生效活动MQ失败", err)
			}
		}

		return
	})
}

func (s *sEmployeeActivity) Delete(ctx context.Context, in *input_employee.EmployeeActivityDeleteInp) (err error) {

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.EmployeeActivityDepartment.Ctx(ctx).Where(dao.EmployeeActivityDepartment.Columns().ActivityId, in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "清理活动部门旧数据失败，请稍后重试！")
			return
		}
		if _, err = dao.EmployeeActivityEmployee.Ctx(ctx).Where(dao.EmployeeActivityEmployee.Columns().ActivityId, in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "清理活动员工旧数据失败，请稍后重试！")
			return
		}

		if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, "删除活动失败，请稍后重试！")
			return
		}
		return
	})
}

func (s *sEmployeeActivity) View(ctx context.Context, in *input_employee.EmployeeActivityViewInp) (res *input_employee.EmployeeActivityViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取活动表信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取活动表信息，请稍后重试！")
			return
		}
	}
	return
}

// Status 更新员工活动状态
func (s *sEmployeeActivity) Status(ctx context.Context, in *input_employee.EmployeeActivityStatusInp) (err error) {
	_, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.EmployeeActivity.Columns().IsEnabled: in.Status,
	}).Update()
	return
}

func (s *sEmployeeActivity) ActivityEffect(ctx context.Context, ActivityId uint64) (err error) {
	var (
		tx               gdb.TX
		EmployeeActivity entity.EmployeeActivity
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if err = dao.EmployeeActivity.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.EmployeeActivity.Columns().Id: ActivityId,
	}).Scan(&EmployeeActivity); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(EmployeeActivity) {
		err = gerror.New("该活动无需处理")
		return
	}

	if EmployeeActivity.Status != 1 {
		err = gerror.New("该活动无需处理")
		return
	}

	// 更新当前活动为进行中
	if _, err = dao.EmployeeActivity.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.EmployeeActivity.Columns().Id: ActivityId,
	}).Data(g.MapStrAny{
		dao.EmployeeActivity.Columns().Status: 2,
	}).Update(); err != nil {
		return
	}

	return
}

func (s *sEmployeeActivity) AppList(ctx context.Context, in *input_employee.EmployeeActivityAppListInp) (list []*input_employee.EmployeeActivityAppListModel, totalCount int, err error) {
	var (
		MemberInfo *model.MemberIdentity
		Employee   *entity.Employee
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	if err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, MemberInfo.Id).Scan(&Employee); err != nil {
		return
	}
	if g.IsEmpty(Employee) {
		// 员工信息错误
		err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
		return
	}

	if Employee.Status != 1 {
		err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
		return
	}

	mod := s.Model(ctx)

	mod = mod.Fields(input_employee.EmployeeActivityAppListModel{}).WithAll()

	mod = mod.Where(dao.EmployeeActivity.Columns().IsEnabled, 1)

	// 根据限制类型过滤活动 - 使用原生SQL
	// 1-不做任何限制 2-限制指定部门 3-限制指定员工

	// 构建权限查询的原生SQL
	restrictionSQL := fmt.Sprintf(`(
		%s = 1 OR 
		(%s = 2 AND %s IN (
			SELECT activity_id FROM %s WHERE department_id = ?
		)) OR
		(%s = 3 AND %s IN (
			SELECT activity_id FROM %s WHERE employee_id = ?
		))
	)`,
		dao.EmployeeActivity.Columns().RestrictionType,
		dao.EmployeeActivity.Columns().RestrictionType,
		dao.EmployeeActivity.Columns().Id,
		dao.EmployeeActivityDepartment.Table(),
		dao.EmployeeActivity.Columns().RestrictionType,
		dao.EmployeeActivity.Columns().Id,
		dao.EmployeeActivityEmployee.Table(),
	)

	mod = mod.Where(restrictionSQL, Employee.DepartmentId, Employee.Id)

	if in.Status > 0 {
		mod = mod.Where(dao.EmployeeActivity.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.EmployeeActivity.Columns().Id)

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_activity_list_failed"))
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_activity_list_failed"))
			return
		}
	}

	return
}

func (s *sEmployeeActivity) AppView(ctx context.Context, in *input_employee.EmployeeActivityAppViewInp) (res *input_employee.EmployeeActivityAppViewModel, err error) {
	var (
		MemberInfo *model.MemberIdentity
		Employee   *entity.Employee
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	if err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, MemberInfo.Id).Scan(&Employee); err != nil {
		return
	}
	if g.IsEmpty(Employee) {
		// 员工信息错误
		err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
		return
	}

	if Employee.Status != 1 {
		err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
		return
	}

	if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_activity_info_failed"))
		return
	}

	// 查询本人已领取数量
	if len(res.CouponList) > 0 {
		// 收集所有优惠券ID
		var couponIds []int64
		for _, coupon := range res.CouponList {
			couponIds = append(couponIds, coupon.CouponId)
		}

		// 一次性查询所有优惠券的已领取数量
		var receivedCoupons []struct {
			CouponId int64 `json:"coupon_id"`
			Count    int   `json:"count"`
		}

		err := dao.ThMemberCoupon.Ctx(ctx).
			Fields("coupon_id, COUNT(*) as count").
			Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
			WhereIn(dao.ThMemberCoupon.Columns().CouponId, couponIds).
			Where(dao.ThMemberCoupon.Columns().ActivityId, in.Id).
			Where(dao.ThMemberCoupon.Columns().Source, 3).
			Group("coupon_id").
			Scan(&receivedCoupons)

		if err != nil {
			g.Log().Error(ctx, "查询已领取数量失败:", err)
		}

		// 构建优惠券ID到已领取数量的映射
		receivedMap := make(map[int64]int)
		for _, item := range receivedCoupons {
			receivedMap[item.CouponId] = item.Count
		}

		// 今日已领取数量
		var todayReceivedCoupons []struct {
			CouponId int64 `json:"coupon_id"`
			Count    int   `json:"count"`
		}

		err = dao.ThMemberCoupon.Ctx(ctx).
			Fields("coupon_id, COUNT(*) as count").
			Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
			WhereIn(dao.ThMemberCoupon.Columns().CouponId, couponIds).
			Where(dao.ThMemberCoupon.Columns().ActivityId, in.Id).
			Where(dao.ThMemberCoupon.Columns().Source, 3).
			WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
			WhereLTE(dao.ThMemberCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 23:59:59").
			Group("coupon_id").
			Scan(&todayReceivedCoupons)

		if err != nil {
			g.Log().Error(ctx, "查询今日已领取数量失败:", err)
		}

		// 构建优惠券ID到今日已领取数量的映射
		todayReceivedMap := make(map[int64]int)
		for _, item := range todayReceivedCoupons {
			todayReceivedMap[item.CouponId] = item.Count
		}

		// 设置每个优惠券的今日已领取数量和状态
		for _, coupon := range res.CouponList {
			coupon.HaveReceived = receivedMap[coupon.CouponId]            // 如果没有记录，默认为0
			coupon.MemberTodayReceive = todayReceivedMap[coupon.CouponId] // 如果没有记录，默认为0

			// 判断优惠券状态
			var status int
			// 1. 先判断用户是否已达到个人总领取限制
			if coupon.AvailableQuantity > 0 && coupon.HaveReceived >= coupon.AvailableQuantity {
				status = 2 // 已领取
			} else {
				// 2. 判断每日限制（如果设置了每日限制）
				if coupon.PerDayAvailable > 0 && coupon.MemberTodayReceive >= coupon.PerDayAvailable {
					status = 4 // 今日已领
				} else {
					status = 1 // 可领取
				}
			}
			coupon.ReceiveStatus = status
		}
	}
	return
}

func (s *sEmployeeActivity) ReceiveActivityCoupon(ctx context.Context, in *input_employee.EmployeeActivityReceiveCouponInp) (err error) {
	var (
		ActivityInfo *entity.EmployeeActivity
		CouponInfo   *entity.ThCoupon
		MemberInfo   *model.MemberIdentity
		Employee     *entity.Employee
		CouponNo     string
		TimeDiff     int
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		MemberInfo = contexts.GetMemberUser(ctx)
		if err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, MemberInfo.Id).Scan(&Employee); err != nil {
			return
		}
		if g.IsEmpty(Employee) {
			// 员工信息错误
			err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
			return
		}

		if Employee.Status != 1 {
			err = gerror.New(gi18n.T(ctx, "non_employee_cannot_view"))
			return
		}

		if err = s.Model(ctx).Where(dao.EmployeeActivity.Columns().Id, in.ActivityId).Scan(&ActivityInfo); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if g.IsEmpty(ActivityInfo) {
			err = gerror.New(gi18n.T(ctx, "activity_not_exist"))
			return
		}

		// 验证活动是否可行
		if ActivityInfo.Status == 1 {
			err = gerror.New(gi18n.T(ctx, "activity_not_start"))
			return
		}

		if ActivityInfo.Status == 3 {
			err = gerror.New(gi18n.T(ctx, "activity_has_ended"))
			return
		}

		if ActivityInfo.IsEnabled != 1 {
			err = gerror.New(gi18n.T(ctx, "activity_disabled"))
			return
		}

		if ActivityInfo.RestrictionType == 2 {
			// 限制部门 - 检查员工的部门是否在允许的部门列表中
			var departmentCount int
			departmentCount, err = dao.EmployeeActivityDepartment.Ctx(ctx).
				Where(dao.EmployeeActivityDepartment.Columns().ActivityId, in.ActivityId).
				Where(dao.EmployeeActivityDepartment.Columns().DepartmentId, Employee.DepartmentId).
				Count()
			if err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "validate_department_permission_failed"))
				return
			}
			if departmentCount == 0 {
				err = gerror.New(gi18n.T(ctx, "validate_department_activity_failed"))
				return
			}
		}

		if ActivityInfo.RestrictionType == 3 {
			// 限制员工 - 检查当前员工是否在允许的员工列表中
			var employeeCount int
			employeeCount, err = dao.EmployeeActivityEmployee.Ctx(ctx).
				Where(dao.EmployeeActivityEmployee.Columns().ActivityId, in.ActivityId).
				Where(dao.EmployeeActivityEmployee.Columns().EmployeeId, Employee.Id).
				Count()
			if err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "validate_employee_permission_failed"))
				return
			}
			if employeeCount == 0 {
				err = gerror.New(gi18n.T(ctx, "validate_employee_activity_failed"))
				return
			}
		}

		if in.CouponId > 0 {
			// 判断活动是否绑定这个券
			var couponCount int
			couponCount, err = dao.EmployeeActivityCoupon.Ctx(ctx).
				Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.ActivityId).
				Where(dao.EmployeeActivityCoupon.Columns().CouponId, in.CouponId).
				Count()
			if err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
				return
			}
			if couponCount == 0 {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_failed"))
				return
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

			if gtime.Now().After(gtime.New(ActivityInfo.EndTime)) {
				err = gerror.New(gi18n.T(ctx, "activity_has_ended"))
				return
			}

			// 判断已经领取的数量
			var receiveCouponCount int
			receiveCouponCount, err = dao.ThMemberCoupon.Ctx(ctx).
				Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
				Where(dao.ThMemberCoupon.Columns().Source, 3).
				Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
				Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId).
				Count()
			if err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
				return
			}
			var ActivityCoupon *entity.EmployeeActivityCoupon
			if err = dao.EmployeeActivityCoupon.Ctx(ctx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.ActivityId).Where(dao.EmployeeActivityCoupon.Columns().CouponId, in.CouponId).Scan(&ActivityCoupon); err != nil {
				return
			}
			if ActivityCoupon.AvailableQuantity > 0 && receiveCouponCount >= ActivityCoupon.AvailableQuantity {
				err = gerror.New(gi18n.T(ctx, "validate_coupon_activity_limit"))
				return
			}

			// 判断今日领取数量
			if ActivityCoupon.PerDayAvailable > 0 {
				var todayReceiveCount int
				todayReceiveCount, err = dao.ThMemberCoupon.Ctx(ctx).
					Where(dao.ThMemberCoupon.Columns().MemberId, MemberInfo.Id).
					Where(dao.ThMemberCoupon.Columns().Source, 3).
					Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
					Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId).
					WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
					WhereLTE(dao.ThMemberCoupon.Columns().CreateAt, gtime.Now().Format("Y-m-d")+" 23:59:59").
					Count()
				if err != nil {
					err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_permission_failed"))
					return
				}

				if todayReceiveCount >= ActivityCoupon.PerDayAvailable {
					err = gerror.New(gi18n.T(ctx, "validate_coupon_daily_limit"))
					return
				}
			}

			// 判断员工活动是否限制当天领取（1-7 分别对应周一到周日）
			if ActivityInfo.LimitWeek != "" {
				// Go Weekday: 0=周日,1=周一,...,6=周六 → 转换为 LimitWeek 约定: 周日=7, 其余不变
				goWeekday := int(gtime.Now().Weekday())
				limitWeekDay := goWeekday
				if goWeekday == 0 {
					limitWeekDay = 7
				}
				limits := gstr.Split(ActivityInfo.LimitWeek, ",")
				if gstr.InArray(limits, gvar.New(limitWeekDay).String()) {
					err = gerror.New(gi18n.T(ctx, "today_not_available"))
					return
				}
			}

			var endTime *gtime.Time
			var State int

			todayStart := gtime.Now()
			if ActivityInfo.CouponValidity == 1 {
				// 跟随活动
				todayStart = ActivityInfo.StartTime
			}

			CouponNo = uuid.CreateOrderCode(CouponInfo.CouponNoPrefix)

			TimeDiff = int(gtime.New(todayStart).Sub(gtime.Now()).Seconds())

			if TimeDiff <= 0 {
				State = 2
				TimeDiff = 0
			} else {
				State = 1
			}

			if ActivityInfo.CouponValidity == 1 {
				// 跟随活动
				endTime = ActivityInfo.EndTime
			} else {
				endTime = gtime.New(todayStart).Add(time.Duration(24*CouponInfo.FixedTerm) * time.Hour)
				endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23) * time.Hour).Add(time.Duration(59) * time.Minute).Add(time.Duration(59) * time.Second)
			}

			if _, err = g.Model(dao.ThMemberCoupon.Table()).Ctx(ctx).Safe().
				Data(g.MapStrAny{
					dao.ThMemberCoupon.Columns().CouponNo:   CouponNo,
					dao.ThMemberCoupon.Columns().CouponId:   CouponInfo.Id,
					dao.ThMemberCoupon.Columns().MemberId:   MemberInfo.Id,
					dao.ThMemberCoupon.Columns().Source:     3,
					dao.ThMemberCoupon.Columns().State:      State,
					dao.ThMemberCoupon.Columns().StartTime:  gtime.New(todayStart).Format("Y-m-d H:i:s"),
					dao.ThMemberCoupon.Columns().EndTime:    endTime,
					dao.ThMemberCoupon.Columns().CountDown:  TimeDiff,
					dao.ThMemberCoupon.Columns().ActivityId: in.ActivityId,
					dao.ThMemberCoupon.Columns().EmployeeId: Employee.Id,
				}).OmitEmptyData().InsertAndGetId(); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "receive_coupon_failed"))
				return
			}

			if _, err = dao.ThCoupon.Ctx(ctx).WherePri(in.CouponId).Increment(dao.ThCoupon.Columns().Count, 1); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
				return
			}

			if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).Where(dao.EmployeeActivityCoupon.Columns().ActivityId, in.ActivityId).Where(dao.EmployeeActivityCoupon.Columns().CouponId, in.CouponId).Increment(dao.EmployeeActivityCoupon.Columns().TotalReceived, 1); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "operation_failed"))
				return
			}
			return
		}

		err = gerror.Wrap(err, gi18n.T(ctx, "validate_coupon_activity_failed"))
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
