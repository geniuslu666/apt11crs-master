package logic_hotel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/biter777/countries"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/nyaruka/phonenumbers"
)

// SyncAppStay 同步APP订单
func (s *sHotelService) SyncAppStay(ctx context.Context, StayData *airhousePublicApi.StayData) (err error) {
	var (
		PmsAppStayInfoDate *entity.PmsAppStay
		PmsAppStay         *entity.PmsAppStay
		PmsGuestProfile    *entity.PmsGuestProfile
		//OrderAmount        float64
	)

	for _, Reservations := range StayData.RoomReservations {
		if err = s.SyncAppRoomReservations(ctx, &Reservations, StayData.ReservationNumber, StayData.ReservationSourceCode, StayData.ReservationSourceName); err != nil {
			g.Log().Error(ctx, err)
		}
		//OrderAmount += Reservations.BookingFee
	}
	PmsAppStayInfoDate = &entity.PmsAppStay{
		Uuid:       StayData.ID,
		Uid:        StayData.UID,
		Puid:       gvar.New(StayData.Property.ID).String(),
		OutOrderSn: StayData.ReservationNumber,
		Booker:     gvar.New(StayData.Booker.ID).String(),
		//OrderAmount: OrderAmount,
		OrderStatus: "HAVE_PAID",
	}
	// APP永远只要更新数据
	if _, err = dao.PmsAppStay.Ctx(ctx).OmitEmptyData().
		Where(dao.PmsAppStay.Columns().Uid, StayData.UID).
		Update(PmsAppStayInfoDate); err != nil {
		return
	}
	// 查询订单信息
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().Uid, StayData.UID).Scan(&PmsAppStay); err != nil {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		return
	}
	// 更新入住人信息
	if err = gvar.New(StayData.Booker).Scan(&PmsGuestProfile); err != nil {
		return
	}
	if err = s.SyncGuest(ctx, PmsAppStay.Booker, PmsGuestProfile, true); err != nil {
		return
	}
	return
}

// SyncAppRoomReservations 同步APP入住单
func (s *sHotelService) SyncAppRoomReservations(ctx context.Context, RoomReservation *airhousePublicApi.RoomReservations, OrderSn interface{}, SourceCode interface{}, SourceName interface{}) (err error) {
	var (
		HandlePmsAppReservationData *entity.PmsAppReservation
		PmsAppReservation           *entity.PmsAppReservation
		//RetrieveAFolioResponse      *airhousePublicApi.RetrieveAFolioJSONDataResponse
		PmsGuestProfile *entity.PmsGuestProfile
		//PmsCharges                  []*entity.PmsCharge
		DataByte []byte
	)
	HandlePmsAppReservationData = &entity.PmsAppReservation{
		Uuid:          gvar.New(RoomReservation.ID).String(),
		Uid:           gvar.New(RoomReservation.UID).String(),
		Puid:          gvar.New(RoomReservation.Property.ID).String(),
		RoomType:      gvar.New(RoomReservation.RoomType.ID).String(),
		RoomUnit:      gvar.New(RoomReservation.RoomUnit.ID).String(),
		RatePlanId:    gvar.New(RoomReservation.Folio.ID).String(),
		CheckinDate:   gtime.New(RoomReservation.CheckinDate),
		CheckoutDate:  gtime.New(RoomReservation.CheckoutDate),
		CheckinTime:   RoomReservation.CheckinTime,
		CheckoutTime:  RoomReservation.CheckoutTime,
		Status:        RoomReservation.Status,
		CheckinStatus: RoomReservation.CheckinStatus,
		AdultCount:    RoomReservation.AdultCount,
		ChildCount:    RoomReservation.ChildCount,
		InfantCount:   RoomReservation.InfantCount,
		GuestRemarks:  RoomReservation.GuestRemarks,
	}
	if !g.IsEmpty(OrderSn) {
		HandlePmsAppReservationData.OutOrderSn = gvar.New(OrderSn).String()
	}
	if !g.IsEmpty(SourceCode) {
		HandlePmsAppReservationData.SourceCode = gvar.New(SourceCode).String()
	}
	if !g.IsEmpty(SourceName) {
		HandlePmsAppReservationData.SourceName = gvar.New(SourceName).String()
	}

	var (
		PmsAppReservationInfo *entity.PmsAppReservation
		OldCheckinStatus      string
		OldStatus             string
	)
	// APP入住订单更新
	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().Uuid, RoomReservation.ID).
		Scan(&PmsAppReservationInfo); err != nil {
		return
	}

	OldCheckinStatus = PmsAppReservationInfo.CheckinStatus
	OldStatus = PmsAppReservationInfo.Status

	// APP入住订单更新
	if _, err = dao.PmsAppReservation.Ctx(ctx).OmitEmptyData().
		Where(dao.PmsAppReservation.Columns().Uuid, RoomReservation.ID).
		Data(HandlePmsAppReservationData).UpdateAndGetAffected(); err != nil {
		// 更新入住订单信息失败
		err = gerror.New(gi18n.T(ctx, "update_checkin_info_failed"))
		return
	}
	// 查询入住单信息
	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().Uuid, RoomReservation.ID).
		Scan(&PmsAppReservation); err != nil {
		return
	}
	if g.IsEmpty(PmsAppReservation) {
		return
	}
	// 更新入住人信息
	if err = gvar.New(RoomReservation.MainGuest).Scan(&PmsGuestProfile); err != nil {
		return
	}
	if err = s.SyncGuest(ctx, PmsAppReservation.MainGuest, PmsGuestProfile, true); err != nil {
		return
	}

	// 回写房间信息
	// 检测日期发生变化移除之前的数据进行重新写入房态
	if _, err = dao.PmsRoomStatus.Ctx(ctx).OmitEmptyWhere().Where(&entity.PmsRoomStatus{
		Puid:      gvar.New(RoomReservation.Property.ID).String(),
		Tuid:      gvar.New(RoomReservation.RoomType.ID).String(),
		Ruid:      gvar.New(RoomReservation.RoomUnit.ID).String(),
		ReserveId: gvar.New(RoomReservation.ID).String(),
	}).Data(dao.PmsRoomStatus.Columns().ReserveId, nil).Update(); err != nil {
		return
	}
	if _, err = dao.PmsRoomStatus.Ctx(ctx).OmitEmptyWhere().
		Where(&entity.PmsRoomStatus{
			Puid: gvar.New(RoomReservation.Property.ID).String(),
			Tuid: gvar.New(RoomReservation.RoomType.ID).String(),
			Ruid: gvar.New(RoomReservation.RoomUnit.ID).String(),
		}).
		WhereGTE(dao.PmsRoomStatus.Columns().Date, RoomReservation.CheckinDate).
		WhereLT(dao.PmsRoomStatus.Columns().Date, RoomReservation.CheckoutDate).
		Update(g.MapStrAny{
			dao.PmsRoomStatus.Columns().ReserveId: RoomReservation.ID,
		}); err != nil {
		return
	}

	//if !g.IsEmpty(RoomReservation.Folio.ID) && SourceCode != "manually_created" {
	//	// 请求获取账单明细
	//	if RetrieveAFolioResponse, err = airhousePublicApi.GetRetrieveAFolio(ctx, gvar.New(RoomReservation.Folio.ID).String()); err != nil {
	//		return
	//	}
	//	for _, Charges := range RetrieveAFolioResponse.Data.Charges {
	//		PmsCharges = append(PmsCharges, &entity.PmsCharge{
	//			Uid:         gvar.New(RoomReservation.Folio.ID).String(),
	//			Date:        gtime.New(Charges.Date),
	//			Name:        consts.Charges[Charges.FeeType],
	//			FeeType:     Charges.FeeType,
	//			Amount:      gvar.New(Charges.Amount).Float64(),
	//			Description: gvar.New(Charges.Description).String(),
	//		})
	//	}
	//	if err = s.SyncCharge(ctx, PmsCharges); err != nil {
	//		g.Log().Error(ctx, err)
	//	}
	//}

	// 同步日志
	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: PmsAppReservation.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		// 该订单无需处理
		err = gerror.New(gi18n.T(ctx, "order_need_not_handle"))
		return
	}
	var ActionWay string
	if RoomReservation.CheckinStatus == "checked_in" && OldCheckinStatus != "checked_in" {
		ActionWay = "ORDER_CHECKED_IN"
	} else if RoomReservation.CheckinStatus == "checked_out" && OldCheckinStatus != "checked_out" {
		ActionWay = "ORDER_CHECKED_OUT"
	} else if RoomReservation.Status == "blocked" && OldStatus != "blocked" {
		ActionWay = "ORDER_BLOCKED"
	} else if RoomReservation.Status == "user_cancelled" && OldStatus != "user_cancelled" {
		ActionWay = "ORDER_USER_CANCELLED"
	} else if RoomReservation.Status == "pending" && OldStatus != "pending" {
		ActionWay = "ORDER_PENDING"
	} else if RoomReservation.Status == "overlapped" && OldStatus != "overlapped" {
		ActionWay = "ORDER_OVERLAPPED"
	} else if RoomReservation.Status == "cancelled" && OldStatus != "cancelled" {
		ActionWay = "ORDER_CANCELLED"
	} else if RoomReservation.Status == "confirmed" && OldStatus != "confirmed" {
		ActionWay = "ORDER_CONFIRMED"
	}
	if !g.IsEmpty(ActionWay) {
		if DataByte, err = json.Marshal(RoomReservation); err != nil {
			return
		}
		if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
			OrderId:            PmsAppStay.Id,
			ReservationOrderId: PmsAppReservation.Id,
			ActionWay:          ActionWay,
			Remark:             string(DataByte),
			OperateType:        "SYSTEM",
		}); err != nil {
			return
		}
	}

	if ActionWay == "ORDER_CONFIRMED" {
		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "房间已确认",
			"en":    "Room confirmed",
			"ja":    "部屋が確定しました",
			"ko":    "객실이 확인되었습니다",
			"zh_CN": "房間已確認",
		}
		systemMessageContent := map[string]string{
			"zh":    "房间已确认",
			"en":    "Room confirmed",
			"ja":    "部屋が確定しました",
			"ko":    "객실이 확인되었습니다",
			"zh_CN": "房間已確認",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(PmsAppStay.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": gvar.New(PmsAppReservation.Id).String(),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "hotel",
			Type:                 "order",
			MemberId:             int(PmsAppStay.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/trip/detail",
			WxLink:               fmt.Sprintf("/subpackages/trips/details?id=%d", PmsAppReservation.Id),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           0,
			OperatorRole:         "SYSTEM",
			OrderSn:              PmsAppReservation.OrderSn,
		})
	}

	return
}

// SyncOtaStay 同步OTA订单
func (s *sHotelService) SyncOtaStay(ctx context.Context, StayData *airhousePublicApi.StayData) (err error) {
	var (
		PmsAppStayInfo  *entity.PmsAppStay
		PmsGuestProfile *entity.PmsGuestProfile
		OrderAmount     float64
	)
	for _, Reservations := range StayData.RoomReservations {
		if err = s.SyncOtaRoomReservations(ctx, &Reservations, StayData.ReservationNumber, StayData.ReservationSourceCode, StayData.ReservationSourceName); err != nil {
			g.Log().Error(ctx, err)
		}
		OrderAmount += Reservations.BookingFee
	}

	PmsAppStayInfoDate := &entity.PmsAppStay{
		Uuid:        StayData.ID,
		Uid:         StayData.UID,
		Puid:        gvar.New(StayData.Property.ID).String(),
		OutOrderSn:  StayData.ReservationNumber,
		Booker:      gvar.New(StayData.Booker.ID).String(),
		OrderAmount: OrderAmount,
		OrderStatus: "HAVE_PAID",
	}
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().Uuid, StayData.ID).Scan(&PmsAppStayInfo); err != nil {
		return
	}
	if g.IsEmpty(PmsAppStayInfo) {
		if _, err = dao.PmsAppStay.Ctx(ctx).OmitEmptyData().InsertAndGetId(PmsAppStayInfoDate); err != nil {
			return
		}
	} else {
		if _, err = dao.PmsAppStay.Ctx(ctx).OmitEmptyData().
			Where(dao.PmsAppStay.Columns().Uuid, StayData.ID).
			UpdateAndGetAffected(PmsAppStayInfoDate); err != nil {
			return
		}
	}

	if err = gvar.New(StayData.Booker).Scan(&PmsGuestProfile); err != nil {
		return
	}

	if err = s.SyncGuest(ctx, gvar.New(StayData.Booker.ID).String(), PmsGuestProfile, false); err != nil {
		return
	}

	return
}

// SyncOtaRoomReservations 同步OTA订单
func (s *sHotelService) SyncOtaRoomReservations(ctx context.Context, RoomReservation *airhousePublicApi.RoomReservations, OrderSn interface{}, SourceCode interface{}, SourceName interface{}) (err error) {
	var (
		RetrieveAFolioResponse      *airhousePublicApi.RetrieveAFolioJSONDataResponse
		InsertDataPmsAppReservation *entity.PmsAppReservation
		PmsGuestProfile             *entity.PmsGuestProfile
		PmsCharges                  []*entity.PmsCharge
	)
	InsertDataPmsAppReservation = &entity.PmsAppReservation{
		Uuid:            gvar.New(RoomReservation.ID).String(),
		Uid:             gvar.New(RoomReservation.UID).String(),
		Puid:            gvar.New(RoomReservation.Property.ID).String(),
		RoomType:        gvar.New(RoomReservation.RoomType.ID).String(),
		RoomUnit:        gvar.New(RoomReservation.RoomUnit.ID).String(),
		RatePlanId:      gvar.New(RoomReservation.Folio.ID).String(),
		CheckinDate:     gtime.New(RoomReservation.CheckinDate),
		CheckoutDate:    gtime.New(RoomReservation.CheckoutDate),
		MainGuest:       gvar.New(RoomReservation.MainGuest.ID).String(),
		CheckinTime:     RoomReservation.CheckinTime,
		CheckoutTime:    RoomReservation.CheckoutTime,
		Status:          RoomReservation.Status,
		CheckinStatus:   RoomReservation.CheckinStatus,
		AdultCount:      RoomReservation.AdultCount,
		ChildCount:      RoomReservation.ChildCount,
		InfantCount:     RoomReservation.InfantCount,
		BookingFee:      gvar.New(RoomReservation.BookingFee).Float64(),
		ChannelFee:      gvar.New(RoomReservation.ChannelFee).Float64(),
		CleaningFee:     gvar.New(RoomReservation.CleaningFee).Float64(),
		CancellationFee: gvar.New(RoomReservation.CancellationFee).Float64(),
		Charges:         gvar.New(RoomReservation.Folio.ID).String(),
		GuestRemarks:    RoomReservation.GuestRemarks,
	}
	if !g.IsEmpty(OrderSn) {
		InsertDataPmsAppReservation.OutOrderSn = gvar.New(OrderSn).String()
	}
	if !g.IsEmpty(SourceCode) {
		InsertDataPmsAppReservation.SourceCode = gvar.New(SourceCode).String()
	}
	if !g.IsEmpty(SourceName) {
		InsertDataPmsAppReservation.SourceName = gvar.New(SourceName).String()
	}
	if err = s.SyncOtaReservation(ctx, InsertDataPmsAppReservation); err != nil {
		g.Log().Error(ctx, err)
	}
	if err = gvar.New(RoomReservation.MainGuest).Scan(&PmsGuestProfile); err != nil {
		return
	}
	if err = s.SyncGuest(ctx, gvar.New(RoomReservation.MainGuest.ID).String(), PmsGuestProfile, false); err != nil {
		g.Log().Error(ctx, err)
	}
	if !g.IsEmpty(RoomReservation.Folio.ID) {
		// 请求获取账单明细
		if RetrieveAFolioResponse, err = airhousePublicApi.GetRetrieveAFolio(ctx, gvar.New(RoomReservation.Folio.ID).String()); err != nil {
			return
		}
		for _, Charges := range RetrieveAFolioResponse.Data.Charges {
			PmsCharges = append(PmsCharges, &entity.PmsCharge{
				Uid:         gvar.New(RoomReservation.Folio.ID).String(),
				Date:        gtime.New(Charges.Date),
				Name:        consts.Charges[Charges.FeeType],
				FeeType:     Charges.FeeType,
				Amount:      gvar.New(Charges.Amount).Float64(),
				Description: gvar.New(Charges.Description).String(),
			})
		}
		if err = s.SyncCharge(ctx, PmsCharges); err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// SyncOtaReservation 同步OTA订单
func (s *sHotelService) SyncOtaReservation(ctx context.Context, RoomReservation *entity.PmsAppReservation) (err error) {
	var (
		RoomReservationInfo *entity.PmsAppReservation
	)
	// 查询是否存在入住管理
	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().Uuid, RoomReservation.Uuid).
		Scan(&RoomReservationInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(ctx, err)
		return
	}
	if g.IsEmpty(RoomReservationInfo) {
		if _, err = dao.PmsAppReservation.Ctx(ctx).
			OmitEmptyData().Data(RoomReservation).Insert(); err != nil {
			return
		}
	} else {
		if _, err = dao.PmsAppReservation.Ctx(ctx).
			Where(dao.PmsAppReservation.Columns().Uuid, RoomReservation.Uuid).
			OmitEmptyData().Data(RoomReservation).
			Update(); err != nil {
			return
		}

		// 回写房间信息
		// 检测日期发生变化移除之前的数据进行重新写入房态
		if _, err = dao.PmsRoomStatus.Ctx(ctx).OmitEmptyWhere().Where(&entity.PmsRoomStatus{
			Puid:      RoomReservation.Puid,
			Tuid:      RoomReservation.RoomType,
			Ruid:      RoomReservation.RoomUnit,
			ReserveId: RoomReservation.Uuid,
		}).Data(dao.PmsRoomStatus.Columns().ReserveId, nil).Update(); err != nil {
			return
		}
		if _, err = dao.PmsRoomStatus.Ctx(ctx).OmitEmptyWhere().
			Where(&entity.PmsRoomStatus{
				Puid: RoomReservation.Puid,
				Tuid: RoomReservation.RoomType,
				Ruid: RoomReservation.RoomUnit,
			}).
			WhereGTE(dao.PmsRoomStatus.Columns().Date, RoomReservation.CheckinDate.Format("Y-m-d")).
			WhereLT(dao.PmsRoomStatus.Columns().Date, RoomReservation.CheckoutDate.Format("Y-m-d")).
			Update(g.MapStrAny{
				dao.PmsRoomStatus.Columns().ReserveId: RoomReservation.Uuid,
			}); err != nil {
			return
		}
	}
	return
}

// SyncGuest 同步入住人信息
func (s *sHotelService) SyncGuest(ctx context.Context, Uid string, PmsGuest *entity.PmsGuestProfile, IsApp bool) (err error) {
	var (
		PhoneNum        *phonenumbers.PhoneNumber
		Phone           string
		AreaNo          string
		Country         countries.CountryCode
		PmsGuestProfile *entity.PmsGuestProfile
	)
	// 手机号解析
	if !g.IsEmpty(PmsGuest.Phone) {
		if gvar.New(PmsGuest.Phone).Bytes()[0] != '+' {
			PmsGuest.Phone = fmt.Sprintf("+%v", PmsGuest.Phone)
		}
		PhoneNum, _ = phonenumbers.Parse(gvar.New(PmsGuest.Phone).String(), "")
		if !g.IsEmpty(PhoneNum.GetCountryCode()) {
			if !g.IsEmpty(PhoneNum.GetNationalNumber()) {
				Phone = fmt.Sprintf("%d", PhoneNum.GetNationalNumber())
			}
			if !g.IsEmpty(PhoneNum.GetCountryCode()) {
				AreaNo = fmt.Sprintf("+%d", PhoneNum.GetCountryCode())
			}
			PhoneNum.GetCountryCodeSource()
		}
	}
	if !g.IsEmpty(PmsGuest.Nationality) {
		Country = countries.ByName(PmsGuest.Nationality)
		if Country != countries.Unknown {
			PmsGuest.Nationality = Country.Alpha2()
		} else {
			PmsGuest.Nationality = "UNKNOW"
		}
	}
	HandlePmsGuestProfileData := &entity.PmsGuestProfile{
		FirstName:     gvar.New(PmsGuest.FirstName).String(),
		LastName:      gvar.New(PmsGuest.LastName).String(),
		FirstNameKana: gvar.New(PmsGuest.FirstNameKana).String(),
		LastNameKana:  gvar.New(PmsGuest.LastNameKana).String(),
		FullName:      gvar.New(PmsGuest.FullName).String(),
		Language:      gvar.New(PmsGuest.Language).String(),
		Email:         gvar.New(PmsGuest.Email).String(),
		Phone:         Phone,
		AreaNo:        AreaNo,
		Nationality:   gvar.New(PmsGuest.Nationality).String(),
		Address:       gvar.New(PmsGuest.Address).String(),
	}
	if IsApp {
		// 更新入住人信息
		if _, err = dao.PmsGuestProfile.Ctx(ctx).OmitEmptyData().
			Where(dao.PmsGuestProfile.Columns().Uid, Uid).
			Data(HandlePmsGuestProfileData).Update(); err != nil {
			return
		}
	} else {
		// 查询是否存在该入住人信息
		if err = dao.PmsGuestProfile.Ctx(ctx).
			Where(dao.PmsGuestProfile.Columns().Uid, Uid).
			Scan(&PmsGuestProfile); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if g.IsEmpty(PmsGuestProfile) {
			// 插入住人信息
			HandlePmsGuestProfileData.Uid = Uid
			if _, err = dao.PmsGuestProfile.Ctx(ctx).OmitEmptyData().
				Where(dao.PmsGuestProfile.Columns().Uid, Uid).
				Data(HandlePmsGuestProfileData).Insert(); err != nil {
				return
			}
		} else {
			// 更新入住人信息
			if _, err = dao.PmsGuestProfile.Ctx(ctx).OmitEmptyData().
				Where(dao.PmsGuestProfile.Columns().Uid, Uid).
				Data(HandlePmsGuestProfileData).Update(); err != nil {
				return
			}
		}

	}
	return
}

// SyncCharge 同步账单
func (s *sHotelService) SyncCharge(ctx context.Context, PmsCharge []*entity.PmsCharge) (err error) {
	var (
		ChargeInfo *entity.PmsCharge
	)
	for _, Charge := range PmsCharge {
		// 查询是否存在入住管理
		if err = dao.PmsCharge.Ctx(ctx).OmitEmptyData().Where(dao.PmsCharge.Columns().Uid, Charge.Uid).Scan(&ChargeInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			g.Log().Error(ctx, err)
			return
		}
		if g.IsEmpty(ChargeInfo) {
			if _, err = dao.PmsCharge.Ctx(ctx).OmitEmptyData().Data(Charge).InsertAndGetId(); err != nil {
				return
			}
		} else {
			if _, err = dao.PmsCharge.Ctx(ctx).OmitEmptyData().
				Where(dao.PmsCharge.Columns().Uid, Charge.Uid).
				Data(Charge).UpdateAndGetAffected(); err != nil {
				return
			}
		}
	}
	return
}
