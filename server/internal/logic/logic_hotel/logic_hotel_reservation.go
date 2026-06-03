package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
)

func (s *sHotelService) ReservationList(ctx context.Context, in *input_hotel.PmsAppReservationListInp) (list []*input_hotel.PmsAppReservationListModel, totalCount int, err error) {
	contexts.SetLanguage(ctx, "zh")
	var (
		mod = dao.PmsAppStay.Ctx(ctx).Unscoped().WithAll().Hook(hook.PmsFindLanguageValueHook).OrderDesc(dao.PmsAppReservation.Columns().Id)
	)
	mod = mod.FieldsPrefix(dao.PmsAppStay.Table(), input_hotel.PmsAppReservationListModel{})
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `property_deleted`", dao.PmsProperty.Table(), dao.PmsProperty.Columns().DeletedAt))
	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))
	mod = mod.LeftJoin(dao.PmsProperty.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Puid, dao.PmsProperty.Table(), dao.PmsProperty.Columns().Uid))

	mod = mod.Where(&entity.PmsAppStay{
		Source:       "APP",
		MemberId:     in.MemberId,
		OrderSn:      in.OrderSn,
		OutOrderSn:   in.OutOrderSn,
		OrderStatus:  in.OrderStatus,
		RefundStatus: in.RefundStatus,
	})

	if !g.IsEmpty(in.SnKeyword) {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsAppStay.Columns().OrderSn, "%"+in.SnKeyword+"%").
			WhereOrLike(dao.PmsAppStay.Columns().OutOrderSn, "%"+in.SnKeyword+"%"))
	}

	if !g.IsEmpty(in.MemberNo) {
		mod = mod.WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberNo+"%")
	}

	if len(in.CheckinDate) == 2 {
		mod = mod.WhereBetween(dao.PmsAppStay.Columns().CheckInDate, gtime.New(in.CheckinDate[0]).Format("Y-m-d"), gtime.New(in.CheckinDate[1]).Format("Y-m-d"))
	}
	if len(in.CheckoutDate) == 2 {
		mod = mod.WhereBetween(dao.PmsAppStay.Columns().CheckOutDate, gtime.New(in.CheckoutDate[0]).Format("Y-m-d"), gtime.New(in.CheckoutDate[1]).Format("Y-m-d"))
	}

	if len(in.CreatedAt) > 1 {
		mod = mod.WhereBetween(dao.PmsAppStay.Columns().CreatedAt, gtime.New(in.CreatedAt[0]).Format("Y-m-d H:i:s"), gtime.New(in.CreatedAt[1]).Format("Y-m-d H:i:s"))
	}
	if !g.IsEmpty(in.Puid) {
		mod = mod.Where(dao.PmsAppStay.Columns().Puid, in.Puid)
	}
	if in.PageReq.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}
	if err = mod.OmitEmptyWhere().ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	for k, v := range list {
		if v.TransactionDetail != nil {
			TransactionDetailAmount := 0.0
			TransactionRefundDetailAmount := 0.0
			for _, TransactionDetail := range v.TransactionDetail {
				if TransactionDetail.PayType != "COUPON" && TransactionDetail.PayStatus == "DONE" {
					TransactionDetailAmount += TransactionDetail.PayAmount
				}
				if TransactionDetail.PayType != "COUPON" && TransactionDetail.PayStatus == "DONE" && TransactionDetail.RefundStatus != "DONE" {
					list[k].IsRefund = true
				}
			}
			for _, TransactionRefundDetail := range v.TransactionRefundDetail {
				if TransactionRefundDetail.RefundStatus == "DONE" {
					TransactionRefundDetailAmount += TransactionRefundDetail.RefundAmount
				}
			}
			if TransactionDetailAmount-TransactionRefundDetailAmount > 0 {
				list[k].IsRefund = true
			} else {
				list[k].IsRefund = false
			}
		}
	}
	return
}

func (s *sHotelService) ReservationRoomList(ctx context.Context, in *input_hotel.PmsAppReservationRoomListInp) (list []*input_hotel.PmsAppReservationRoomListModel, totalCount int, orderCount int, err error) {
	var (
		AppStay *entity.PmsAppStay
	)
	mod := dao.PmsAppReservation.Ctx(ctx).Unscoped().WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.FieldsPrefix(dao.PmsAppReservation.Table(), input_hotel.PmsAppReservationRoomListModel{})
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `property_deleted`", dao.PmsProperty.Table(), dao.PmsProperty.Columns().DeletedAt))
	mod = mod.LeftJoin(dao.PmsProperty.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.PmsAppReservation.Table(), dao.PmsAppReservation.Columns().Puid, dao.PmsProperty.Table(), dao.PmsProperty.Columns().Uid))
	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.PmsAppReservation.Columns().OrderSn, "%"+in.OrderSn+"%")
	}
	if !g.IsEmpty(in.OutOrderSn) {
		mod = mod.WhereLike(dao.PmsAppReservation.Columns().OutOrderSn, "%"+in.OutOrderSn+"%")
	}
	//if len(in.CheckinDate) == 2 {
	//	mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckinDate, in.CheckinDate[0], in.CheckinDate[1])
	//}
	//if len(in.CheckoutDate) == 2 {
	//	mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckoutDate, in.CheckoutDate[0], in.CheckoutDate[1])
	//}
	if len(in.CheckinDate) == 2 {
		mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckinDate, gtime.New(in.CheckinDate[0]).Format("Y-m-d"), gtime.New(in.CheckinDate[1]).Format("Y-m-d"))
	}
	if len(in.CheckoutDate) == 2 {
		mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckoutDate, gtime.New(in.CheckoutDate[0]).Format("Y-m-d"), gtime.New(in.CheckoutDate[1]).Format("Y-m-d"))
	}

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsAppReservation.Columns().Status, in.Status)
	}
	if !g.IsEmpty(in.CheckinStatus) {
		mod = mod.Where(dao.PmsAppReservation.Columns().CheckinStatus, in.CheckinStatus)
	}
	if !g.IsEmpty(in.OrderStatus) {
		mod = mod.Where(dao.PmsAppReservation.Columns().OrderStatus, in.OrderStatus)
	}
	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.PmsAppReservation.Columns().MemberId, in.MemberId)
	}
	if !g.IsEmpty(in.Source) {
		mod = mod.Where(dao.PmsAppReservation.Columns().Source, in.Source)
	} else {
		mod = mod.Where(dao.PmsAppReservation.Columns().Source, "APP")
	}

	if !g.IsEmpty(in.PropertyName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.PropertyName, "name")
		if err == nil {
			propertyUids, _ := s.PropertyGetUids(ctx, uuIds)
			mod = mod.WhereIn(dao.PmsAppReservation.Columns().Puid, propertyUids)
		}
	}
	if !g.IsEmpty(in.Puid) {
		mod = mod.Where(dao.PmsAppReservation.Columns().Puid, in.Puid)
	}

	if !g.IsEmpty(in.SelectStatus) {
		if in.SelectStatus == "CONFIRM" {
			mod = mod.Where(dao.PmsAppReservation.Columns().Status, "confirmed")
		} else if in.SelectStatus == "CANCEL" {
			mod = mod.Where(dao.PmsAppReservation.Columns().OrderStatus, "CANCEL")
		}
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsAppReservation.Columns().Id)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取订单列表失败，请稍后重试！")
			return
		}
	} else {
		if in.IsNew == "Y" {
			AppStay = new(entity.PmsAppStay)
			if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().MemberId, in.MemberId).OrderDesc(dao.PmsAppStay.Columns().CreatedAt).Scan(&AppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			if g.IsEmpty(AppStay) {
				return
			}
			mod = mod.Limit(1)
			mod = mod.Where(dao.PmsAppReservation.Columns().OrderSn, AppStay.OrderSn)
		}
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取订单列表失败，请稍后重试！")
			return
		}
	}
	if !g.IsEmpty(list) {
		for k, v := range list {
			if v.TransactionDetail != nil {
				for _, TransactionDetail := range v.TransactionDetail {
					if TransactionDetail.PayStatus == "DONE" && TransactionDetail.RefundStatus != "DONE" {
						list[k].IsRefund = true
					}
				}
			}
		}
	}
	if g.IsEmpty(in.IsNew) {
		orderCount = totalCount
	} else {

		if orderCount, err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsAppStay.Columns().MemberId: in.MemberId,
		}).Count(); err != nil {
			return
		}
	}
	return
}

func (s *sHotelService) ReservationRebateList(ctx context.Context, in *input_hotel.PmsAppReservationRebateListInp) (list []*input_hotel.PmsAppReservationRebateListModel, totalCount int, err error) {
	sql, err := gdb.ToSQL(ctx, func(ctx context.Context) error {
		AppReservationMod := dao.PmsAppReservation.Ctx(ctx)
		AppReservationMod = AppReservationMod.Fields(dao.PmsAppReservation.Columns().OrderSn)
		AppReservationMod = AppReservationMod.Fields(gdb.Raw("(SUM( adult_count ) + SUM( child_count )) as guestNum"))
		AppReservationMod = AppReservationMod.Fields(gdb.Raw("COUNT(" + dao.PmsAppReservation.Columns().Id + ") as roomNum"))
		AppReservationMod = AppReservationMod.Where(dao.PmsAppReservation.Columns().Status, in.Status)
		AppReservationMod = AppReservationMod.Where(dao.PmsAppReservation.Columns().CheckinStatus, in.CheckinStatus)
		AppReservationMod = AppReservationMod.Where(dao.PmsAppReservation.Columns().OrderSn, in.OrderSn)
		_, err = AppReservationMod.OmitEmptyWhere().Group(dao.PmsAppReservation.Columns().OrderSn).All()
		return err
	})

	sql = fmt.Sprintf("(%s) as %s", sql, dao.PmsAppReservation.Table())
	g.Log().Info(ctx, sql)

	mod := dao.PmsAppStay.Ctx(ctx).WithAll().
		FieldsPrefix(dao.PmsAppReservation.Table(), "*").
		FieldsPrefix(dao.PmsAppStay.Table(),
			dao.PmsAppStay.Columns().Id,
			dao.PmsAppStay.Columns().Puid,
			dao.PmsAppStay.Columns().OrderSn,
			dao.PmsAppStay.Columns().OutOrderSn,
			dao.PmsAppStay.Columns().CreatedAt,
			dao.PmsAppStay.Columns().OrderAmount,
			dao.PmsAppStay.Columns().PayModel,
			dao.PmsAppStay.Columns().OrderStatus,
			dao.PmsAppStay.Columns().Booker,
			dao.PmsAppStay.Columns().ExpirationTime,
			dao.PmsAppStay.Columns().RebateRate,
			dao.PmsAppStay.Columns().RebateAmount,
			dao.PmsAppStay.Columns().RebateStatus).
		WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().OrderSn, in.OrderSn).
		WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().OutOrderSn, in.OutOrderSn).
		WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().OrderStatus, in.OrderStatus).
		WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().MemberId, in.MemberId)

	if in.FxType == "channel" {

		memberIds, err := dao.PmsMember.Ctx(ctx).
			Fields(dao.PmsMember.Columns().Id).
			Where(dao.PmsMember.Columns().RebateMode, "CHANNEL").
			Where(dao.PmsMember.Columns().ChannelId, in.FxId).Array()
		if err != nil || g.IsEmpty(memberIds) {
			mod = mod.WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Id, -1)
		}
		mod = mod.WherePrefixIn(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Referrer, memberIds)
	}
	if in.FxType == "staff" {

		memberIds, err := dao.PmsMember.Ctx(ctx).
			Fields(dao.PmsMember.Columns().Id).
			Where(dao.PmsMember.Columns().RebateMode, "STAFF").
			Where(dao.PmsMember.Columns().StaffId, in.FxId).Array()
		if err != nil || g.IsEmpty(memberIds) {
			mod = mod.WherePrefix(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Id, -1)
		}
		mod = mod.WherePrefixIn(dao.PmsAppStay.Table(), dao.PmsAppStay.Columns().Referrer, memberIds)
	}

	if err = mod.
		Page(in.Page, in.PerPage).OmitEmptyWhere().OrderDesc(dao.PmsAppReservation.Columns().Id).
		InnerJoin(sql, "hg_pms_app_stay.order_sn = hg_pms_app_reservation.order_sn").Hook(hook.PmsFindLanguageValueHook).ScanAndCount(&list, &totalCount, false); err != nil {
		return
	}
	for k, v := range list {
		list[k].TotalRefundAmount = 0
		if v.TransactionDetail != nil {
			for _, TransactionDetail := range v.TransactionDetail {
				if TransactionDetail.PayStatus == "DONE" && TransactionDetail.RefundStatus != "DONE" {
					list[k].IsRefund = true
				}
			}
		}
		if v.TransactionRefundDetail != nil {
			for _, TransactionRefundDetail := range v.TransactionRefundDetail {
				list[k].TotalRefundAmount += TransactionRefundDetail.RefundAmount
			}
		}
	}
	return
}

func (s *sHotelService) ReservationExport(ctx context.Context, in *input_hotel.PmsAppReservationExportListInp) (err error) {
	var (
		ChangeList     []*input_hotel.PmsAppReservationExportModel
		AppReservation []*entity.PmsAppReservation
		OutOrderSn     []string

		ReservationListIn *input_hotel.PmsAppReservationListInp
	)
	in.PageReq.Pagination = false

	ReservationListIn = &input_hotel.PmsAppReservationListInp{
		PageReq:      in.PageReq,
		OrderSn:      in.OrderSn,
		OutOrderSn:   in.OutOrderSn,
		RefundStatus: in.RefundStatus,
		OrderStatus:  in.OrderStatus,
		MemberId:     in.MemberId,
		MemberNo:     in.MemberNo,
		Puid:         in.Puid,
		SnKeyword:    in.SnKeyword,
	}
	if !g.IsEmpty(in.CreatedAt) {
		for _, timeItem := range strings.Split(in.CreatedAt, ",") {
			ReservationListIn.CreatedAt = append(ReservationListIn.CreatedAt, gtime.New(timeItem).Format("Y-m-d H:i:s"))
		}
	}
	if !g.IsEmpty(in.CheckinDate) {
		for _, timeItem := range strings.Split(in.CheckinDate, ",") {
			ReservationListIn.CheckinDate = append(ReservationListIn.CheckinDate, gtime.New(timeItem).Format("Y-m-d"))
		}
	}
	if !g.IsEmpty(in.CheckoutDate) {
		for _, timeItem := range strings.Split(in.CheckoutDate, ",") {
			ReservationListIn.CheckoutDate = append(ReservationListIn.CheckoutDate, gtime.New(timeItem).Format("Y-m-d"))
		}
	}

	list, _, err := s.ReservationList(ctx, ReservationListIn)
	if err != nil {
		return
	}
	for _, item := range list {
		FullName := ""
		if !g.IsEmpty(item.GuestProfileDetail) {
			FullName = item.GuestProfileDetail.FullName
		}
		OutOrderSn = append(OutOrderSn, item.OutOrderSn)
		itemData := &input_hotel.PmsAppReservationExportModel{
			PropertyName:      item.PropertyDetail.Name,
			GuestName:         FullName,
			SystemOrderID:     item.OrderSn,
			ChannelOrderID:    item.OutOrderSn,
			PaymentStatus:     item.OrderStatus,
			CheckInDate:       item.CheckInDate,
			CheckOutDate:      item.CheckOutDate,
			NumberOfGuests:    0,
			NumberOfNights:    convert.Diff(item.CheckInDate, item.CheckOutDate, "days"),
			OrderTime:         item.CreatedAt.Format("Y-m-d H:i:s"),
			TotalAmount:       item.OrderAmount,
			RefundTotalAmount: item.RefundAmount,
			RefundTime:        item.RefundTime.Format("Y-m-d H:i:s"),
		}
		if !g.IsEmpty(item.MemberDetail) {
			itemData.MemberNo = item.MemberDetail.MemberNo
		}
		for _, TransactionDetail := range item.TransactionDetail {
			if TransactionDetail.PayStatus == "DONE" {
				switch TransactionDetail.PayType {
				case "BAL":
					itemData.PointsPayment += TransactionDetail.PayAmount
					break
				case "COUPON":
					itemData.CouponPayment += TransactionDetail.PayAmount
					break
				case "WeChatPay":
					itemData.PaycloudWechatPay += TransactionDetail.PayAmount
					break
				case "Alipay+":
					itemData.PaycloudAlipayPay += TransactionDetail.PayAmount
					break
				case "Paypal":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "PaypalCard":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "StripeCard":
					itemData.StripeCreditPay += TransactionDetail.PayAmount
					break
				case "WeChatMiniPay":
					itemData.MlilifeWeChatMiniPay += TransactionDetail.PayAmount
					break
				}
			}
		}

		for _, TransactionRefundDetail := range item.TransactionRefundDetail {
			switch TransactionRefundDetail.RefundType {
			case "BAL":
				itemData.PointsRefund += TransactionRefundDetail.RefundAmount
				break
			case "COUPON":
				itemData.CouponRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatPay":
				itemData.PaycloudWechatRefund += TransactionRefundDetail.RefundAmount
				break
			case "Alipay+":
				itemData.PaycloudAlipayRefund += TransactionRefundDetail.RefundAmount
				break
			case "Paypal":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "PaypalCard":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "StripeCard":
				itemData.StripeCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatMiniPay":
				itemData.MlilifeWeChatMiniRefund += TransactionRefundDetail.RefundAmount
				break
			}
		}
		ChangeList = append(ChangeList, itemData)
	}

	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().OutOrderSn, OutOrderSn).
		Group(dao.PmsAppReservation.Columns().OutOrderSn).
		Fields("out_order_sn,SUM(infant_count) as infant_count,SUM(child_count) as child_count,SUM(adult_count) as adult_count").
		Scan(&AppReservation); err != nil {
		return
	}
	for _, item := range ChangeList {
		for _, item2 := range AppReservation {
			if item.ChannelOrderID == item2.OutOrderSn {
				item.NumberOfGuests = item2.AdultCount + item2.ChildCount + item2.InfantCount
			}
		}
	}

	tags, err := convert.GetEntityDescTags(input_hotel.PmsAppReservationExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出系统入住订单-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("系统入住订单")
		exports   []input_hotel.PmsAppReservationExportModel
	)

	if err = gconv.Scan(ChangeList, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sHotelService) ReservationEdit(ctx context.Context, in *input_hotel.PmsAppReservationEditInp) (err error) {

	if err = hgorm.IsUnique(ctx, &dao.PmsAppReservation, g.Map{dao.PmsAppReservation.Columns().OrderSn: in.OrderSn}, "系统订单号已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if !g.IsEmpty(in.Id) {
			if _, err = dao.PmsAppReservation.Ctx(ctx).
				Fields(input_hotel.PmsAppReservationUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改系统入住订单失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsAppReservation.Ctx(ctx).
			Fields(input_hotel.PmsAppReservationInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增系统入住订单失败，请稍后重试！")
		}
		return
	})
}

func (s *sHotelService) ReservationDelete(ctx context.Context, in *input_hotel.PmsAppReservationDeleteInp) (err error) {

	if _, err = dao.PmsAppReservation.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除系统入住订单失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) ReservationView(ctx context.Context, in *input_hotel.PmsAppReservationViewInp) (res *input_hotel.PmsAppReservationViewModel, err error) {
	orm := dao.PmsAppStay.Ctx(ctx).Unscoped()
	if !g.IsEmpty(in.Id) {
		orm = orm.WherePri(in.Id)
	} else if !g.IsEmpty(in.OrderSn) {
		orm = orm.Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn)
	} else if !g.IsEmpty(in.OutOrderSn) {
		orm = orm.Where(dao.PmsAppStay.Columns().OutOrderSn, in.OutOrderSn)
	} else {
		err = gerror.Wrap(err, "参数错误，Id和OrderSn不能同时为空")
		return
	}
	if err = orm.WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取系统入住订单信息，请稍后重试！")
		return
	}

	// Check member deleted
	if res.MemberId > 0 {
		memberDeletedAt, _ := dao.PmsMember.Ctx(ctx).Unscoped().Fields(dao.PmsMember.Columns().DeletedAt).
			Where(dao.PmsMember.Columns().Id, res.MemberId).Value()
		if !memberDeletedAt.IsNil() && !memberDeletedAt.IsEmpty() {
			res.MemberDeleted = true
			// Manually load deleted member info
			var memberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, res.MemberId).Scan(&memberInfo); err == nil && memberInfo != nil {
				res.MemberDetail = &struct {
					gmeta.Meta `orm:"table:hg_pms_member"`
					Id         int    `json:"id"              orm:"id"                description:"主键"`
					MemberNo   string `json:"memberNo"        orm:"member_no"         description:"会员号"`
					FirstName  string `json:"firstName"       orm:"first_name"        description:"名"`
					LastName   string `json:"lastName"        orm:"last_name"         description:"姓"`
					FullName   string `json:"fullName"        orm:"full_name"         description:"全名"`
					Phone      string `json:"phone"           orm:"phone"             description:"手机号"`
					PhoneArea  string `json:"phoneArea"       orm:"phone_area"        description:"手机区号"`
					Mail       string `json:"mail"            orm:"mail"              description:"邮箱"`
				}{
					Id:        memberInfo.Id,
					MemberNo:  memberInfo.MemberNo,
					FirstName: memberInfo.FirstName,
					LastName:  memberInfo.LastName,
					FullName:  memberInfo.FullName,
					Phone:     memberInfo.Phone,
					PhoneArea: memberInfo.PhoneArea,
					Mail:      memberInfo.Mail,
				}
			}
		}
	}

	// Check referrer deleted
	if res.Referrer > 0 && res.ReferrerDetail == nil {
		referrerDeletedAt, _ := dao.PmsMember.Ctx(ctx).Unscoped().Fields(dao.PmsMember.Columns().DeletedAt).
			Where(dao.PmsMember.Columns().Id, res.Referrer).Value()
		if !referrerDeletedAt.IsNil() && !referrerDeletedAt.IsEmpty() {
			res.ReferrerDeleted = true
			// Manually load deleted referrer info
			var referrerInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, res.Referrer).Scan(&referrerInfo); err == nil && referrerInfo != nil {
				res.ReferrerDetail = &struct {
					gmeta.Meta `orm:"table:hg_pms_member"`
					*entity.PmsMember
					Staff *struct {
						gmeta.Meta `orm:"table:hg_pms_staff"`
						*entity.PmsStaff
					} `json:"staff" orm:"with:id=staff_id"`
					Channel *struct {
						gmeta.Meta `orm:"table:hg_pms_channel"`
						*entity.PmsChannel
					} `json:"channel" orm:"with:id=channel_id"`
				}{
					PmsMember: referrerInfo,
				}
			}
		}
	}

	// Check property deleted
	if !g.IsEmpty(res.Puid) {
		propertyDeletedAt, _ := dao.PmsProperty.Ctx(ctx).Unscoped().Fields(dao.PmsProperty.Columns().DeletedAt).
			Where(dao.PmsProperty.Columns().Uid, res.Puid).Value()
		if !propertyDeletedAt.IsNil() && !propertyDeletedAt.IsEmpty() {
			res.PropertyDeleted = true
			// Manually load deleted property info
			var propertyInfo *entity.PmsProperty
			if err = dao.PmsProperty.Ctx(ctx).Unscoped().Where(dao.PmsProperty.Columns().Uid, res.Puid).Scan(&propertyInfo); err == nil && propertyInfo != nil {
				res.PropertyDetail = &struct {
					gmeta.Meta `orm:"table:hg_pms_property"`
					*entity.PmsProperty
				}{
					PmsProperty: propertyInfo,
				}
			}
		}
	}

	for _, v := range res.LogList {
		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			if v.OperateId > 0 {
				var AdminMemberInfo *entity.AdminMember
				if err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
					return
				}
				if AdminMemberInfo != nil {
					v.OperateName = AdminMemberInfo.Username
				}
			} else {
				v.OperateName = "后台"
			}
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			if PmsMemberInfo != nil {
				v.OperateName = PmsMemberInfo.FullName
			}
		}
	}

	return
}

func (s *sHotelService) ReservationStatus(ctx context.Context, in *input_hotel.PmsAppReservationStatusInp) (err error) {
	if _, err = dao.PmsAppReservation.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsAppReservation.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新系统入住订单状态失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) ReservationReferrerList(ctx context.Context, in *input_hotel.PmsAppReservationReferrerListInp) (list []*input_hotel.PmsAppReservationReferrerListModel, totalCount int, err error) {

	var (
		mod = dao.PmsAppStay.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).OrderDesc(dao.PmsAppReservation.Columns().Id)
	)
	mod = mod.Where(&entity.PmsAppStay{
		Source:       "APP",
		MemberId:     in.MemberId,
		OrderSn:      in.OrderSn,
		OutOrderSn:   in.OutOrderSn,
		OrderStatus:  in.OrderStatus,
		RefundStatus: in.RefundStatus,
	}).WhereNotNull(dao.PmsAppStay.Columns().Referrer)

	if !g.IsEmpty(in.CheckinDate) {
		mod = mod.Where(dao.PmsAppStay.Columns().CheckInDate, gtime.New(in.CheckinDate).Format("Y-m-d"))
	}
	if !g.IsEmpty(in.CheckoutDate) {
		mod = mod.Where(dao.PmsAppStay.Columns().CheckOutDate, gtime.New(in.CheckoutDate).Format("Y-m-d"))
	}
	if len(in.CreatedAt) > 1 {
		mod = mod.WhereBetween(dao.PmsAppStay.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}
	if !g.IsEmpty(in.Puid) {
		mod = mod.Where(dao.PmsAppStay.Columns().Puid, in.Puid)
	}
	if in.PageReq.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}
	if err = mod.OmitEmptyWhere().ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	return
}

func (s *sHotelService) AppStayInfo(ctx context.Context, in *input_hotel.PmsAppReservationViewInp) (res *input_hotel.PmsAppStayInfoModel, err error) {
	orm := dao.PmsAppStay.Ctx(ctx)
	if !g.IsEmpty(in.Id) {
		orm = orm.WherePri(in.Id)
	} else if !g.IsEmpty(in.OrderSn) {
		orm = orm.Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn)
	} else if !g.IsEmpty(in.OutOrderSn) {
		orm = orm.Where(dao.PmsAppStay.Columns().OutOrderSn, in.OutOrderSn)
	} else {
		err = gerror.Wrap(err, "参数错误，Id和OrderSn不能同时为空")
		return
	}
	if err = orm.WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取系统入住订单信息，请稍后重试！")
		return
	}
	return
}
