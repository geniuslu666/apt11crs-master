package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/utility/convert"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) ReservationList(ctx context.Context, req *hotel.ReservationListReq) (res *hotel.ReservationListRes, err error) {
	var (
		MemberInfo *model.MemberIdentity
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	res = new(hotel.ReservationListRes)
	mdl := dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().MemberId, MemberInfo.Id)
	switch req.Status {
	case 1:
		mdl = mdl.
			Where(dao.PmsAppReservation.Columns().Status, g.Slice{"confirmed"}).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, g.Slice{"checked_in", "before_checkin"}).
			Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID")
		break
	case 2:
		mdl = mdl.
			Where(dao.PmsAppReservation.Columns().Status, g.Slice{"confirmed"}).
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_out").
			Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID")
		break
	case 3:
		mdl = mdl.
			Where(dao.PmsAppReservation.Columns().Status, g.Slice{"cancelled", "user_cancelled"}).
			Where(dao.PmsAppReservation.Columns().OrderStatus, "CANCEL")
		break
	default:
		// default: 组合 1、2、3 三种状态，用 OR 进行筛选
		colStatus := dao.PmsAppReservation.Columns().Status
		colCheckinStatus := dao.PmsAppReservation.Columns().CheckinStatus
		colOrderStatus := dao.PmsAppReservation.Columns().OrderStatus
		confirmed := g.Slice{"confirmed"}
		checkedInOut := g.Slice{"checked_in", "checked_out"}
		cancelled := g.Slice{"cancelled", "user_cancelled"}
		// memberId AND ( (status IN confirmed AND checkin_status = 'before_checkin' AND order_status = 'HAVE_PAID')
		//              OR (status IN confirmed AND checkin_status IN ('checked_in','checked_out') AND order_status = 'HAVE_PAID')
		//              OR (status IN ('cancelled','user_cancelled') AND order_status = 'CANCEL') )
		cond := fmt.Sprintf(
			"((%s IN (?) AND %s = ? AND %s = ?) OR (%s IN (?) AND %s IN (?) AND %s = ?) OR (%s IN (?) AND %s = ?))",
			colStatus, colCheckinStatus, colOrderStatus,
			colStatus, colCheckinStatus, colOrderStatus,
			colStatus, colOrderStatus,
		)
		mdl = mdl.Where(
			cond,
			confirmed, "before_checkin", "HAVE_PAID",
			confirmed, checkedInOut, "HAVE_PAID",
			cancelled, "CANCEL",
		)
	}
	if err = mdl.WithAll().
		Page(req.Page, req.PageSize).
		OrderDesc(dao.PmsAppReservation.Columns().CreatedAt).
		Hook(hook.PmsFindLanguageValueHook).
		ScanAndCount(&res.List, &res.Count, false); err != nil {
		return
	}
	if !g.IsEmpty(res.List) {
		for k, v := range res.List {
			res.List[k].Days = convert.Diff(v.CheckinDate, v.CheckoutDate, "days")
		}
	}
	return
}
func (c *ControllerHotel) ReservationDetail(ctx context.Context, req *hotel.ReservationDetailReq) (res *hotel.ReservationDetailRes, err error) {
	var (
		appReservation *entity.PmsAppReservation
		pmsProperty    *entity.PmsProperty
		pmsPropertys   []*entity.PmsLanguage
		pmsRoomType    *entity.PmsRoomType
		//pmsRoomUnit    *entity.PmsRoomUnit
		weekday = make(map[string][]string)
	)
	weekday["zh"] = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	weekday["zh_CN"] = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	weekday["ko"] = []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}
	weekday["ja"] = []string{"日曜", "月曜", "火曜", "水曜", "木曜", "金曜", "土曜"}
	weekday["en"] = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	res = new(hotel.ReservationDetailRes)
	appReservation = new(entity.PmsAppReservation)
	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservation.Columns().Id: req.Id,
	}).Scan(&appReservation); err != nil {
		return
	}
	if g.IsEmpty(appReservation) {
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}
	res.Id = appReservation.Id
	res.RoomNum = 1
	res.Days = convert.Diff(appReservation.CheckinDate.Format("Y-m-d"), appReservation.CheckoutDate.Format("Y-m-d"), "days")
	res.CheckinDate = appReservation.CheckinDate.Format("m-d")
	res.CheckoutDate = appReservation.CheckoutDate.Format("m-d")
	res.CheckoutDateTime = appReservation.CheckoutDate.Format("Y-m-d") + " " + appReservation.CheckoutTime
	res.CheckinTime = appReservation.CheckinTime
	res.CheckoutTime = appReservation.CheckoutTime
	InWeek := gvar.New(gtime.New(appReservation.CheckinDate).Format("w")).Int()
	OutWeek := gvar.New(gtime.New(appReservation.CheckoutDate).Format("w")).Int()
	res.CheckinWeek = weekday[contexts.GetLanguage(ctx)][InWeek]
	res.CheckoutWeek = weekday[contexts.GetLanguage(ctx)][OutWeek]
	res.OrderSn = appReservation.OrderSn
	res.CreatedAt = appReservation.CreatedAt
	res.Status = appReservation.Status
	res.OrderStatus = appReservation.OrderStatus
	res.CheckinStatus = appReservation.CheckinStatus
	// 查询物业信息
	if err = dao.PmsProperty.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsProperty.Columns().Uid: appReservation.Puid,
	}).Hook(hook.PmsFindLanguageValueHook).Scan(&pmsProperty); err != nil {
		return
	}
	if g.IsEmpty(pmsProperty) {
		err = errors.New("物业信息不存在")
		return
	}
	res.Property = new(hotel.ReservationDetailProperty)
	res.Property.Name = pmsProperty.Name
	res.Property.Address = pmsProperty.Address
	res.Property.RequiredBook = pmsProperty.RequiredBook
	res.Property.ContactPhone = pmsProperty.Phone
	res.Property.ContactEmail = pmsProperty.ContactEmail
	res.Property.GgLng = pmsProperty.GgLng
	res.Property.GgLat = pmsProperty.GgLat

	// 查询入住人信息
	if err = dao.PmsGuestProfile.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsGuestProfile.Columns().Uid: appReservation.MainGuest,
	}).Scan(&res.Guest); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	// 查询房型信息
	if err = dao.PmsRoomType.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(g.MapStrAny{
		dao.PmsRoomType.Columns().Uid: appReservation.RoomType,
	}).Scan(&pmsRoomType); err != nil {
		return
	}
	if g.IsEmpty(pmsRoomType) {
		err = errors.New("房型信息不存在")
		return
	}
	res.RoomTypeName = pmsRoomType.Name
	// 查询房间信息
	//if !g.IsEmpty(appReservation.RoomUnit) {
	//	if err = dao.PmsRoomUnit.Ctx(ctx).Where(g.MapStrAny{
	//		dao.PmsRoomUnit.Columns().Uid: appReservation.RoomUnit,
	//	}).Scan(&pmsRoomUnit); err != nil {
	//		return
	//	}
	//	res.RoomUnitName = pmsRoomUnit.RoomNo
	//}
	// 查询物业信息
	if err = dao.PmsProperty.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsProperty.Columns().Uid: appReservation.Puid,
	}).Scan(&pmsProperty); err != nil {
		return
	}

	// 查询日语地址和日语物业名
	if err = dao.PmsLanguage.Ctx(ctx).
		Where(dao.PmsLanguage.Columns().Uuid, g.Slice{pmsProperty.Name, pmsProperty.Address}).
		Where(dao.PmsLanguage.Columns().Language, "ja").
		Scan(&pmsPropertys); err != nil {
		return
	}
	for _, v := range pmsPropertys {
		if v.Uuid == pmsProperty.Name {
			res.JpPropertyName = v.Content
		}
		if v.Uuid == pmsProperty.Address {
			res.JpAddress = v.Content
		}
	}

	// 查询变更单
	if err = dao.PmsAppReservationChange.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservationChange.Columns().OrderId:      appReservation.Id,
		dao.PmsAppReservationChange.Columns().ChangeStatus: "DONE",
	}).Scan(&res.ChangeDataList); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(res.ChangeDataList) {
		for k, v := range res.ChangeDataList {
			switch v.ChangeType {
			case "GUEST":
				res.ChangeDataList[k].Title = consts.BookingChangeTypeGuest[contexts.GetLanguage(ctx)]
				break
			case "PEOPLE":
				res.ChangeDataList[k].Title = consts.BookingChangeTypePeople[contexts.GetLanguage(ctx)]
				break
			case "DATE":
				res.ChangeDataList[k].Title = consts.BookingChangeTypeDate[contexts.GetLanguage(ctx)]
				break
			default:
				res.ChangeDataList[k].Title = ""
			}
		}
	}
	return
}
func (c *ControllerHotel) FindReservationRoomNo(ctx context.Context, req *hotel.FindReservationRoomNoReq) (res *hotel.FindReservationRoomNoRes, err error) {
	var (
		appReservation *entity.PmsAppReservation
		RoomNo         *gvar.Var
	)
	res = new(hotel.FindReservationRoomNoRes)
	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservation.Columns().Id: req.Id,
	}).Scan(&appReservation); err != nil {
		return
	}

	if g.IsEmpty(appReservation) {
		// 不存在入住单
		err = gerror.New(gi18n.T(ctx, "no_checkin_form"))
		return
	}
	// 判断日期是否是当天并且时间是否到达了下午13点
	if gtime.Now().After(gtime.New(fmt.Sprintf("%s 13:00:00", appReservation.CheckinDate.Format("Y-m-d")))) {
		if RoomNo, err = dao.PmsRoomUnit.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsRoomUnit.Columns().Uid: appReservation.RoomUnit,
		}).Value(dao.PmsRoomUnit.Columns().RoomNo); err != nil {
			return
		}
		if RoomNo.IsEmpty() {
			// 没有房间号
			err = gerror.New(gi18n.T(ctx, "no_room_number"))
			return
		}
		res.RoomNo = RoomNo.String()
	}
	return
}
