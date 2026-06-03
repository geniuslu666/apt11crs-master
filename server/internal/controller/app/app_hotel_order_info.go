package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/convert"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) OrderList(ctx context.Context, req *hotel.OrderListReq) (res *hotel.OrderListRes, err error) {
	var (
		MemberUser   = contexts.GetMemberUser(ctx)
		AppStay      []*entity.PmsAppStay
		propertyInfo *input_hotel.PmsPropertyViewModel
	)
	res = new(hotel.OrderListRes)
	Where := g.MapStrAny{
		dao.PmsAppStay.Columns().MemberId: MemberUser.Id,
	}
	if !g.IsEmpty(req.OrderStatus) {
		Where[dao.PmsAppStay.Columns().OrderStatus] = req.OrderStatus
	}
	if err = dao.PmsAppStay.Ctx(ctx).Where(Where).
		OrderDesc(dao.PmsAppStay.Columns().Id).
		Page(req.PageNum, req.PageSize).
		ScanAndCount(&AppStay, &res.Count, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(AppStay) {
		for _, v := range AppStay {
			if err = dao.PmsProperty.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsProperty.Columns().Uid: v.Puid,
			}).Hook(hook.PmsFindLanguageValueHook).Scan(&propertyInfo); err != nil {
				err = errors.New("获取物业信息失败")
				return
			}
			if g.IsEmpty(propertyInfo) {
				err = errors.New("物业信息不存在")
				return
			}

			// 查询房型订单信息
			var AppReservation []*entity.PmsAppReservation
			if err = dao.PmsAppReservation.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsAppReservation.Columns().OrderSn: v.OrderSn,
			}).Scan(&AppReservation); err != nil {
				return
			}
			if g.IsEmpty(AppReservation) {
				// 订单房间信息不存在
				err = gerror.New(gi18n.T(ctx, "order_room_info_does_not_exist"))
				return
			}

			var IsCancel = "Y"
			if !g.IsEmpty(AppReservation[0].PricePlanInfo) {
				var PricePlanInfo *entity.PmsPricePlan
				if err = json.Unmarshal([]byte(gvar.New(AppReservation[0].PricePlanInfo).String()), &PricePlanInfo); err != nil {
					// 解析路径失败
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				IsCancel = PricePlanInfo.IsCancel
			}

			CheckinDate := gtime.New(v.CheckInDate).Format("Y-m-d")
			CheckoutDate := gtime.New(v.CheckOutDate).Format("Y-m-d")
			Countdown := v.ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
			if Countdown < 0 {
				Countdown = 0
			}
			item := &hotel.OrderItem{
				Id:              v.Id,
				OrderSn:         v.OrderSn,
				OrderStatus:     v.OrderStatus,
				CheckInDate:     CheckinDate,
				CheckOutDate:    CheckoutDate,
				Days:            convert.Diff(CheckinDate, CheckoutDate, "days"),
				PropertyName:    propertyInfo.Name,
				Address:         propertyInfo.Address,
				OrderAmount:     v.OrderAmount,
				RoomNum:         0,
				AdultCount:      0,
				ChildCount:      0,
				CreateOrderTime: gtime.New(v.CreatedAt).Format("Y-m-d H:i:s"),
				Countdown:       Countdown,
				RefundStatus:    v.RefundStatus,
				IsCancel:        IsCancel,
				IsFx:            v.IsFx,
			}
			if g.IsEmpty(item.IsFx) {
				item.IsFx = "N"
			}
			if item.AdultCount, err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().OrderSn, v.OrderSn).Sum(dao.PmsAppReservation.Columns().AdultCount); err != nil {
				return
			}
			if item.ChildCount, err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().OrderSn, v.OrderSn).Sum(dao.PmsAppReservation.Columns().ChildCount); err != nil {
				return
			}
			if item.RoomNum, err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().OrderSn, v.OrderSn).Count(dao.PmsAppReservation.Columns().Id); err != nil {
				return
			}
			//ThirePayAmount := gvar.New(0)
			var PmsTransaction *entity.PmsTransaction
			// 查询第三方支付金额
			if err = dao.PmsTransaction.Ctx(ctx).
				Where(dao.PmsTransaction.Columns().OrderSn, v.OrderSn).
				Where("pay_channel != ?", "SYSTEM").
				Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
				Scan(&PmsTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			if !g.IsEmpty(PmsTransaction) {
				item.ThirePayAmount = PmsTransaction.Amount
			}
			res.List = append(res.List, item)
		}
	}

	return
}
func (c *ControllerHotel) OrderDetail(ctx context.Context, req *hotel.OrderDetailReq) (res *hotel.OrderDetailRes, err error) {
	var (
		PmsPropertyInfo   gdb.Record
		PmsPropertyView   *input_hotel.PmsPropertyViewModel
		PmsRoomType       gdb.Record
		AppStay           *entity.PmsAppStay
		PmsGuestProfile   *entity.PmsGuestProfile
		AppReservation    []*entity.PmsAppReservation
		AppTransaction    []*entity.PmsTransaction
		TransactionRefund []*entity.PmsTransactionRefund
		CancelOrderInfo   *entity.PmsAppCancelOrder
		CheckinTime       string
		CheckoutTime      string
		CancelRate        []*struct {
			Name      string `json:"name"      dc:"规则名"`
			Mode      string `json:"mode"      dc:"规则模式"`
			StartDays int    `json:"startDays" dc:"开始天数"`
			EndDays   int    `json:"endDays"   dc:"结束天数"`
			Rate      int    `json:"rate"      dc:"取消费率"`
			Date      string `json:"date"      dc:"规则解析日期"`
			Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
		}
		CancelTime *gtime.Time
	)
	res = new(hotel.OrderDetailRes)
	res.OrderSn = req.OrderSn
	// 查询订单信息
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: req.OrderSn,
	}).Scan(&AppStay); err != nil {
		return
	}
	if g.IsEmpty(AppStay) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	res.Id = AppStay.Id
	res.Countdown = AppStay.ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
	if res.Countdown < 0 {
		res.Countdown = 0
	}
	res.CreateOrderTime = AppStay.CreatedAt.String()
	CancelTime = AppStay.CancelTime
	if g.IsEmpty(CancelTime) {
		// 查询PmsAppStayLog表 OrderId是id ActionWay是CANCEL 取CreatedAt字段
		cancelLogCreatedAt, _ := dao.PmsAppStayLog.Ctx(ctx).
			Fields(dao.PmsAppStayLog.Columns().CreatedAt).
			Where(dao.PmsAppStayLog.Columns().OrderId, AppStay.Id).
			Where(dao.PmsAppStayLog.Columns().ActionWay, "CANCEL").
			Value()
		if !cancelLogCreatedAt.IsNil() && !cancelLogCreatedAt.IsEmpty() {
			CancelTime = cancelLogCreatedAt.GTime()
		}
	}
	res.CancelTime = CancelTime.String()
	res.IsFx = AppStay.IsFx
	if g.IsEmpty(res.IsFx) {
		res.IsFx = "N"
	}
	if err = dao.PmsGuestProfile.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsGuestProfile.Columns().Uid: AppStay.Booker,
	}).Scan(&PmsGuestProfile); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsGuestProfile) {
		// 订单信息异常
		err = gerror.New(gi18n.T(ctx, "order_info_exception"))
		return
	}
	res.User.FullName = PmsGuestProfile.FullName
	res.User.Phone = PmsGuestProfile.Phone
	res.User.PhoneArea = PmsGuestProfile.AreaNo
	res.User.Mail = PmsGuestProfile.Email
	res.MemberId = AppStay.MemberId
	// 查询物业信息
	if PmsPropertyInfo, err = dao.PmsProperty.Ctx(ctx).Where(dao.PmsProperty.Columns().Uid, AppStay.Puid).One(); err != nil {
		return
	}
	if PmsPropertyInfo.IsEmpty() {
		// 物业不存在
		err = gerror.New(gi18n.T(ctx, "property_does_not_exist"))
	}
	if PmsPropertyView, _, err = service.HotelService().PropertyView(ctx, &input_hotel.PmsPropertyViewInp{
		Id: PmsPropertyInfo["id"].Int(),
	}); err != nil {
		return nil, err
	}
	res.Property.Id = PmsPropertyView.Id
	res.Property.Uid = PmsPropertyView.Uid
	res.Property.Cover = PmsPropertyView.Cover
	res.Property.Name = PmsPropertyView.Name
	res.Property.Address = PmsPropertyView.Address
	res.Property.RequiredBook = PmsPropertyView.RequiredBook
	res.Property.GgLat = PmsPropertyView.GgLat
	res.Property.GgLng = PmsPropertyView.GgLng
	res.OrderDate.StartDate = AppStay.CheckInDate
	res.OrderDate.EndDate = AppStay.CheckOutDate
	res.OrderDate.Days = convert.Diff(res.OrderDate.StartDate, res.OrderDate.EndDate, "days")
	// 查询房型订单信息
	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservation.Columns().OrderSn: req.OrderSn,
	}).Scan(&AppReservation); err != nil {
		return
	}
	if g.IsEmpty(AppReservation) {
		// 订单房间信息不存在
		err = gerror.New(gi18n.T(ctx, "order_room_info_does_not_exist"))
		return
	}
	for _, v := range AppReservation {
		// 查询房型信息
		if PmsRoomType, err = dao.PmsRoomType.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(dao.PmsRoomType.Columns().Uid, v.RoomType).One(); err != nil {
			return
		}
		if PmsRoomType.IsEmpty() {
			// 房型不存在
			err = gerror.New(gi18n.T(ctx, "room_type_does_not_exist"))
			return
		}

		var IsCancel = "Y"
		if !g.IsEmpty(v.PricePlanInfo) {
			var PricePlanInfo *entity.PmsPricePlan
			if err = json.Unmarshal([]byte(gvar.New(v.PricePlanInfo).String()), &PricePlanInfo); err != nil {
				// 解析路径失败
				err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
				return
			}
			IsCancel = PricePlanInfo.IsCancel
		}

		res.CheckinStatus = v.CheckinStatus
		RoomTypeInfo := &input_hotel.OrderRooms{
			Tid:          PmsRoomType["id"].Int(),
			Puid:         v.Puid,
			RoomId:       v.RoomType,
			Cover:        PmsRoomType["cover"].String(),
			Name:         PmsRoomType["name"].String(),
			RoomNum:      1,
			Adult:        v.AdultCount,
			Child:        v.ChildCount,
			Infant:       v.InfantCount,
			BookingFee:   0,
			IsCancel:     IsCancel,
			ChangeAmount: v.ChangeAmount,
			CheckinDate:  v.CheckinDate,
			CheckoutDate: v.CheckoutDate,
			CheckinTime:  v.CheckinTime,
			CheckoutTime: v.CheckoutTime,
		}
		// 查询房型价格
		if err = dao.PmsCharge.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsCharge.Columns().Uid: v.Charges,
		}).Scan(&RoomTypeInfo.Charges); err != nil {
			return
		}
		RoomTypeInfo.BookingFee = v.BookingFee
		res.RatePlans = append(res.RatePlans, RoomTypeInfo)
		CheckinTime = v.CheckinTime
		CheckoutTime = v.CheckoutTime

		res.PayInfo.ChangeAmount += v.ChangeAmount
	}
	res.OrderDate.StartDate = fmt.Sprintf("%s %s", res.OrderDate.StartDate, CheckinTime)
	res.OrderDate.EndDate = fmt.Sprintf("%s %s", res.OrderDate.EndDate, CheckoutTime)
	res.PayInfo.PayModel = AppStay.PayModel
	res.PayInfo.AllAmount = AppStay.OrderAmount
	res.PayInfo.OrderStatus = AppStay.OrderStatus
	res.PayInfo.RefundStatus = AppStay.RefundStatus

	if err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn: req.OrderSn,
	}).Scan(&AppTransaction); err != nil {
		return
	}
	for _, v := range AppTransaction {
		switch v.PayChannel {
		case "SYSTEM":
			if v.PayType == "BAL" {
				//if v.PayStatus == "DONE" {
				res.PayInfo.Balance.BalanceAmount = res.PayInfo.Balance.BalanceAmount + v.Amount
				//}
				res.PayInfo.Balance.BalancePayOrderSn = v.TransactionSn
				res.PayInfo.Balance.PayStatus = v.PayStatus
				res.PayInfo.Balance.RefundAmount = v.RefundAmount
				res.PayInfo.Balance.RefundStatus = v.RefundStatus
			}
			if v.PayType == "COUPON" {
				res.PayInfo.Coupon.CouponAmount = v.Amount
				res.PayInfo.Coupon.CouponPayOrderSn = v.TransactionSn
				res.PayInfo.Coupon.PayStatus = v.PayStatus
				res.PayInfo.Coupon.RefundAmount = v.RefundAmount
				res.PayInfo.Coupon.RefundStatus = v.RefundStatus
			}
		default:
			if v.OrderType == "BOOKING" {
				res.PayInfo.ThirdPay.ThirdAmount = res.PayInfo.ThirdPay.ThirdAmount + v.Amount
			} else {
				if v.PayStatus == "DONE" {
					res.PayInfo.ThirdPay.ThirdAmount = res.PayInfo.ThirdPay.ThirdAmount + v.Amount
				}
			}
			res.PayInfo.ThirdPay.ThirdPayOrderSn = v.TransactionSn
			res.PayInfo.ThirdPay.PayStatus = v.PayStatus
			res.PayInfo.ThirdPay.RefundAmount += v.RefundAmount
			res.PayInfo.ThirdPay.RefundStatus = v.RefundStatus
		}
	}
	//if AppStay.OrderStatus == "WAIT_PAY" {
	//	var WaitThirdTransaction *entity.PmsTransaction
	//	if err = dao.PmsTransaction.Ctx(ctx).
	//		Where(dao.PmsTransaction.Columns().OrderSn, req.OrderSn).
	//		Where("pay_channel != ?", "SYSTEM").
	//		Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").
	//		OrderDesc(dao.PmsTransaction.Columns().Id).
	//		Limit(1).
	//		Scan(&WaitThirdTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
	//		return
	//	}
	//	if !g.IsEmpty(WaitThirdTransaction) {
	//		res.PayInfo.ThirdPay.ThirdAmount = WaitThirdTransaction.Amount
	//		res.PayInfo.ThirdPay.ThirdPayOrderSn = WaitThirdTransaction.TransactionSn
	//		res.PayInfo.ThirdPay.PayStatus = WaitThirdTransaction.PayStatus
	//	}
	//}
	// 查询是否存在退款订单
	if err = dao.PmsAppCancelOrder.Ctx(ctx).Where(dao.PmsAppCancelOrder.Columns().OrderSn, req.OrderSn).Scan(&CancelOrderInfo); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(CancelOrderInfo) {
		if res.RefundDetail, err = service.HotelService().RefundOrderDetail(ctx, &input_hotel.RefundDetailInp{OrderSn: req.OrderSn}); err != nil {
			return
		}
	} else {
		if err = dao.PmsTransactionRefund.Ctx(ctx).Where(dao.PmsTransactionRefund.Columns().OrderSn, req.OrderSn).Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").Scan(&TransactionRefund); err != nil {
			return
		}
		if !g.IsEmpty(TransactionRefund) {
			res.RefundDetail = new(input_hotel.RefundDetailModel)
			for _, v := range TransactionRefund {
				if v.RefundType == "BAL" {
					res.RefundDetail.RefundBalance += v.RefundAmount
				} else {
					res.RefundDetail.RefundFee += v.RefundAmount
				}
			}
			res.RefundDetail.TransactionRefund = TransactionRefund
		}

	}
	if AppStay.OrderStatus == "WAIT_PAY" {
		res.PayAmount = res.PayInfo.AllAmount - res.PayInfo.Balance.BalanceAmount - res.PayInfo.Coupon.CouponAmount
	} else {
		res.PayAmount = res.PayInfo.AllAmount
	}

	NowTime := gtime.Now().Format("Y-m-d")
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	for k, v := range CancelRate {
		switch v.Mode {
		case "after":
			startDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			if startDate <= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = startDate
			break
		case "middle":
			startDay := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if startDay <= NowTime && endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
			break
		case "before":
			endDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				CancelRate[k].Selected = true
			}
			CancelRate[k].Date = endDate
			break
		}
	}
	res.CancelRate = CancelRate

	// 查询退款记录列表
	res.RefundRecordList = getRefundRecordList(ctx, req.OrderSn)

	return
}

// getRefundRecordList 获取退款记录列表
// 逻辑：
// 1. 查询 transaction_refund 表中的退款记录
// 2. 查询 app_cancel_order 表中的取消记录
// 3. 如果 cancel_order 中有记录但 refund 表中没有 operate_type 为 USER 的记录，说明有一笔退款为0的记录
// 4. 合并结果并按申请时间升序排序
func getRefundRecordList(ctx context.Context, orderSn string) []*hotel.RefundRecordItem {
	var (
		refundRecords []*entity.PmsTransactionRefund
		cancelOrder   *entity.PmsAppCancelOrder
		result        []*hotel.RefundRecordItem
		hasUserRefund bool
	)

	// 1. 查询 transaction_refund 表中的退款记录（已完成的退款）
	_ = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().OrderSn, orderSn).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		OrderAsc(dao.PmsTransactionRefund.Columns().CreatedAt).
		Scan(&refundRecords)

	// 转换退款记录
	for _, v := range refundRecords {
		refundType := "AMOUNT"
		if v.RefundType == "BAL" {
			refundType = "BAL"
		}
		result = append(result, &hotel.RefundRecordItem{
			RefundType:   refundType,
			RefundAmount: v.RefundAmount,
			ApplyTime:    v.CreatedAt.String(),
		})
		// 检查是否有用户发起的退款
		if v.OperateType == "USER" {
			hasUserRefund = true
		}
	}

	// 2. 查询 app_cancel_order 表中的取消记录
	_ = dao.PmsAppCancelOrder.Ctx(ctx).
		Where(dao.PmsAppCancelOrder.Columns().OrderSn, orderSn).
		Scan(&cancelOrder)

	// 3. 如果有取消记录但没有用户发起的退款记录，说明是退款金额为0的情况
	if !g.IsEmpty(cancelOrder) && !hasUserRefund {
		result = append(result, &hotel.RefundRecordItem{
			RefundType:   "AMOUNT",
			RefundAmount: 0,
			ApplyTime:    cancelOrder.CreatedAt.String(),
		})
	}

	// 4. 按申请时间升序排序
	if len(result) > 1 {
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].ApplyTime > result[j].ApplyTime {
					result[i], result[j] = result[j], result[i]
				}
			}
		}
	}

	return result
}
