package admin

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
)

func (c *ControllerPms) AppReservationLastOrder(ctx context.Context, req *pms.AppReservationLastOrderReq) (res *pms.AppReservationLastOrderRes, err error) {
	var (
		OrderSnValue *gvar.Var
	)
	if OrderSnValue, err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().MemberId, req.MemberId).Limit(1).OrderDesc(dao.PmsAppStay.Columns().Id).Value(dao.PmsAppStay.Columns().OrderSn); err != nil {
		return
	}
	res = new(pms.AppReservationLastOrderRes)
	res.OrderSn = OrderSnValue.String()
	return
}

func (c *ControllerPms) AppReservationList(ctx context.Context, req *pms.AppReservationListReq) (res *pms.AppReservationListRes, err error) {
	list, totalCount, err := service.HotelService().ReservationList(ctx, &req.PmsAppReservationListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsAppReservationListModel{}
	}

	res = new(pms.AppReservationListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerPms) AppReservationRoomList(ctx context.Context, req *pms.AppReservationRoomListReq) (res *pms.AppReservationRoomListRes, err error) {
	list, totalCount, _, err := service.HotelService().ReservationRoomList(ctx, &req.PmsAppReservationRoomListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsAppReservationRoomListModel{}
	}

	res = new(pms.AppReservationRoomListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerPms) AppReservationRebateList(ctx context.Context, req *pms.AppReservationRebateListReq) (res *pms.AppReservationRebateListRes, err error) {
	list, totalCount, err := service.HotelService().ReservationRebateList(ctx, &req.PmsAppReservationRebateListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsAppReservationRebateListModel{}
	}

	res = new(pms.AppReservationRebateListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerPms) AppReservationExport(ctx context.Context, req *pms.AppReservationExportReq) (res *pms.AppReservationExportRes, err error) {
	err = service.HotelService().ReservationExport(ctx, &req.PmsAppReservationExportListInp)
	return
}

func (c *ControllerPms) AppReservationView(ctx context.Context, req *pms.AppReservationViewReq) (res *pms.AppReservationViewRes, err error) {
	data, err := service.HotelService().ReservationView(ctx, &req.PmsAppReservationViewInp)
	if err != nil {
		return
	}

	res = new(pms.AppReservationViewRes)
	res.PmsAppReservationViewModel = data
	return
}

func (c *ControllerPms) AppReservationEdit(ctx context.Context, req *pms.AppReservationEditReq) (res *pms.AppReservationEditRes, err error) {
	err = service.HotelService().ReservationEdit(ctx, &req.PmsAppReservationEditInp)
	return
}

func (c *ControllerPms) AppReservationDelete(ctx context.Context, req *pms.AppReservationDeleteReq) (res *pms.AppReservationDeleteRes, err error) {
	err = service.HotelService().ReservationDelete(ctx, &req.PmsAppReservationDeleteInp)
	return
}

func (c *ControllerPms) AppReservationStatus(ctx context.Context, req *pms.AppReservationStatusReq) (res *pms.AppReservationStatusRes, err error) {
	err = service.HotelService().ReservationStatus(ctx, &req.PmsAppReservationStatusInp)
	return
}

func (c *ControllerPms) AppReservationCancel(ctx context.Context, req *pms.AppReservationCancelReq) (res *pms.AppReservationCancelRes, err error) {
	var (
		RoomReservation []*entity.PmsAppReservation
	)

	// 查询入住订单信息 进行取消操作
	if err = dao.PmsAppReservation.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppReservation.Columns().OrderSn: req.OrderSn,
		//dao.PmsAppReservation.Columns().Id:      req.Id,
	}).Scan(&RoomReservation); err != nil {
		return
	}
	if g.IsEmpty(RoomReservation) {
		err = gerror.New("不存在入住单")
		return
	}

	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: req.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		err = gerror.New("该订单无需处理")
		return
	}

	if PmsAppStay.OrderStatus == "CANCEL" {
		err = gerror.New("该订单已取消")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, PmsAppStay.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		for _, RoomReservationItem := range RoomReservation {
			if RoomReservationItem.Status != "confirmed" && RoomReservationItem.Status != "overlapped" {
				err = gerror.New("无法取消")
				return
			}
			if _, err = airhousePublicApi.UpdateRoomReservationPost(ctx, RoomReservationItem.Uuid, g.MapStrAny{
				"status":           "cancelled",
				"checkin_date":     RoomReservationItem.CheckinDate,
				"checkout_date":    RoomReservationItem.CheckoutDate,
				"cancellation_fee": RoomReservationItem.BookingFee,
				"booking_fee":      0,
			}); err != nil {
				return
			}
			// 更新入住订单信息
			if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Uuid, RoomReservationItem.Uuid).Data(g.MapStrAny{
				dao.PmsAppReservation.Columns().OrderStatus:     "CANCEL",
				dao.PmsAppReservation.Columns().Status:          "cancelled",
				dao.PmsAppReservation.Columns().CancellationFee: 0,
				dao.PmsAppReservation.Columns().CancelRemake:    req.Reason,
			}).Update(); err != nil {
				return
			}
		}

		// 整单退款
		if err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      req.OrderSn,
			RefundAmount: 0,
		}, tx); err != nil {
			return
		}

		// 失效
		invalidSendMsg, _ := json.Marshal(g.Map{
			"type": "INVALID",
			"id":   gvar.New(PmsAppStay.Id).Int(),
		})
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderAward,
			DataByte:     invalidSendMsg,
			Header:       nil,
		}); err != nil {
			g.Log().Error(ctx, "发送下单奖励失效MQ失败", err)
		}
		return
	}); err != nil {
		return
	}
	// 修改主订单为已取消
	if _, err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, req.OrderSn).Data(g.MapStrAny{
		dao.PmsAppReservation.Columns().OrderStatus: "CANCEL",
	}).Update(); err != nil {
		return
	}

	// 酒店订单取消

	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   "CANCEL",
		Remark:      "订单取消",
		OperateType: "ADMIN",
		OperateId:   int(contexts.GetUserId(ctx)),
	}); err != nil {
		return
	}

	return
}
func (c *ControllerPms) ReservationRoomStatus(ctx context.Context, req *pms.ReservationRoomStatusReq) (res *pms.ReservationRoomStatusRes, err error) {
	var (
		arr            = make(map[string]*g.Var)
		Availabilities []*entity.PmsAvailabilities
	)

	res = new(pms.ReservationRoomStatusRes)
	if res.RoomList, res.Count, err = propertyList(ctx, req); err != nil {
		return
	}
	var rtUids []string
	var roomTypeUids []string
	for _, ru := range res.RoomList {
		rtUids = append(rtUids, ru.Uid)
		roomTypeUids = append(roomTypeUids, ru.RoomType.Uid)
	}
	roomTypeUids = convert.UniqueSlice(roomTypeUids)
	if err = dao.PmsAvailabilities.Ctx(ctx).
		WhereIn(dao.PmsAvailabilities.Columns().Tuid, roomTypeUids).
		WhereBetween(dao.PmsAvailabilities.Columns().Date, req.StartDate, req.EndDate).
		OrderAsc(dao.PmsAvailabilities.Columns().Date).
		Scan(&Availabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	res.Availabilities = make(map[string]map[string][]*entity.PmsAvailabilities)
	for _, v := range Availabilities {
		if g.IsEmpty(res.Availabilities[v.Tuid]) {
			res.Availabilities[v.Tuid] = make(map[string][]*entity.PmsAvailabilities)
		}
		res.Availabilities[v.Tuid][v.Date] = append(res.Availabilities[v.Tuid][v.Date], v)
	}

	var orm = g.Model(dao.PmsAppReservation.Table()).Ctx(ctx).WithAll()
	var where = orm.Builder()
	where = where.WhereOrLTE(dao.PmsAppReservation.Columns().CheckinDate, req.EndDate)
	where = where.WhereOrLTE(dao.PmsAppReservation.Columns().CheckoutDate, req.EndDate)
	where = where.WhereOrGTE(dao.PmsAppReservation.Columns().CheckinDate, req.StartDate)
	where = where.WhereOrGTE(dao.PmsAppReservation.Columns().CheckoutDate, req.StartDate)
	orm = orm.Where(where)
	orm = orm.WhereIn(dao.PmsAppReservation.Columns().RoomUnit, rtUids)
	orm = orm.Where(dao.PmsAppReservation.Columns().Status, "confirmed")
	if !g.IsEmpty(req.PUid) {
		orm = orm.Where(dao.PmsAppReservation.Columns().Puid+" = ?", req.PUid)
	}
	if !g.IsEmpty(req.RoomTypeUid) {
		orm = orm.Where(dao.PmsAppReservation.Columns().RoomType+" = ?", req.RoomTypeUid)
	}
	if !g.IsEmpty(req.RoomUtilUid) {
		orm = orm.Where(dao.PmsAppReservation.Columns().RoomUnit+" = ?", req.RoomUtilUid)
	}
	if err = orm.Scan(&res.Reservation); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	for k, v := range res.Reservation {
		res.Reservation[k].CheckinUnix = gtime.New(v.CheckinDate.String() + " 0:0:0").TimestampMilli()
		res.Reservation[k].CheckOutUnix = gtime.New(v.CheckoutDate.String() + " 23:59:59").Add(-24 * 60 * 60 * time.Second).TimestampMilli()
	}
	res.ShowData = make(map[string]map[string]*gvar.Var)
	for _, RoomList := range res.RoomList {
		arr = make(map[string]*g.Var)
		for _, Reservation := range res.Reservation {
			if Reservation.RoomUnit == RoomList.Uid {
				CheckNum := ((Reservation.CheckOutUnix)-(Reservation.CheckinUnix))/(24*60*60*1000) + 1
				for i := gvar.New(1).Int64(); i <= CheckNum; i++ {
					js := gvar.New(i * -24 * 60 * 60).Duration()
					datedetails := gtime.New(Reservation.CheckoutDate).Add(js * time.Second).Format("Y-m-d")
					days := gvar.New(0).Int64()
					if i == CheckNum {
						days = CheckNum
					}
					if datedetails == req.StartDate {
						days = i
					}
					if gtime.New(datedetails).Before(gtime.New(req.StartDate)) || gtime.New(datedetails).After(gtime.New(req.EndDate)) {
						continue
					}
					arr[datedetails] = gvar.New(g.MapStrAny{
						"orderSn":       Reservation.OrderSn,
						"id":            Reservation.Id,
						"uid":           Reservation.Uuid,
						"checkinStatus": Reservation.CheckinStatus,
						"status":        Reservation.Status,
						"days":          days,
						"selected":      false,
						"reservation":   Reservation,
					})
				}
				if res.ShowData[RoomList.PmsRoomUnit.Uid] == nil {
					res.ShowData[RoomList.PmsRoomUnit.Uid] = make(map[string]*gvar.Var)
				}
				res.ShowData[RoomList.PmsRoomUnit.Uid]["roomNo"] = gvar.New(RoomList.PmsRoomUnit.RoomNo)
				res.ShowData[RoomList.PmsRoomUnit.Uid]["roomtype"] = gvar.New(RoomList.RoomType.Name)
				res.ShowData[RoomList.PmsRoomUnit.Uid]["uid"] = gvar.New(RoomList.PmsRoomUnit.Uid)
				res.ShowData[RoomList.PmsRoomUnit.Uid]["roomtypeuid"] = gvar.New(RoomList.PmsRoomUnit.RtUid)
			}
			ParamsCheckNum := ((gtime.New(req.EndDate).Timestamp())-(gtime.New(req.StartDate).Timestamp()))/(24*60*60) + 1
			for i := gvar.New(0).Int64(); i < ParamsCheckNum; i++ {
				js := gvar.New(i * -24 * 60 * 60).Duration()
				datedetails := gtime.New(req.EndDate).Add(js * time.Second).Format("Y-m-d")
				if gtime.New(datedetails).Before(gtime.New(req.StartDate)) {
					continue
				}
				if arr[datedetails].IsEmpty() {
					arr[datedetails] = g.NewVar(g.MapStrAny{
						"uid":           nil,
						"checkinStatus": nil,
						"status":        nil,
						"days":          1,
						"selected":      false,
						"reservation":   nil,
					})
				}
			}
		}
		if res.ShowData[RoomList.PmsRoomUnit.Uid] == nil {
			res.ShowData[RoomList.PmsRoomUnit.Uid] = make(map[string]*gvar.Var)
		}
		res.ShowData[RoomList.PmsRoomUnit.Uid]["date"] = gvar.New(arr)
	}
	return
}

func propertyList(ctx context.Context, req *pms.ReservationRoomStatusReq) (res []*pms.ReservationRoomUtilListy, total int, err error) {
	var (
		orm           = dao.PmsRoomUnit.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook)
		Tuids         []gdb.Value
		PUID          []string
		TUID          []string
		RoomTypeInfos []*entity.PmsRoomType
		PropertyInfos []*entity.PmsProperty
	)
	if !g.IsEmpty(req.PUid) {
		if Tuids, err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Puid, req.PUid).Array(dao.PmsRoomType.Columns().Uid); err != nil {
			return
		}
		if !g.IsEmpty(Tuids) {
			orm = orm.Where(dao.PmsRoomUnit.Columns().RtUid, Tuids)
		}
	}
	if !g.IsEmpty(req.RoomTypeUid) {
		orm = orm.Where(dao.PmsRoomUnit.Columns().RtUid, req.RoomTypeUid)
	}
	if !g.IsEmpty(req.RoomUtilUid) {
		orm = orm.Where(dao.PmsRoomUnit.Columns().Uid, req.RoomUtilUid)
	}
	orm = orm.OrderAsc(dao.PmsRoomUnit.Columns().RtUid)
	orm = orm.Page(req.Page, req.Limit)
	if err = orm.ScanAndCount(&res, &total, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	for _, v := range res {
		TUID = append(TUID, v.RtUid)
	}
	TUID = convert.UniqueSlice(TUID)
	if err = dao.PmsRoomType.Ctx(ctx).WhereIn(dao.PmsRoomType.Columns().Uid, TUID).Hook(hook.PmsFindLanguageValueHook).Scan(&RoomTypeInfos); err != nil {
		return
	}
	for _, v := range RoomTypeInfos {
		PUID = append(PUID, v.Puid)
	}
	if err = dao.PmsProperty.Ctx(ctx).WhereIn(dao.PmsProperty.Columns().Uid, PUID).Hook(hook.PmsFindLanguageValueHook).Scan(&PropertyInfos); err != nil {
		return
	}
	for k, v := range res {
		for _, vv := range RoomTypeInfos {
			if v.RtUid == vv.Uid {
				res[k].RoomType = &pms.ReservationRoomUtilListRoomType{}
				res[k].RoomType.Name = vv.Name
				res[k].RoomType.Uid = vv.Uid
				res[k].RoomType.Puid = vv.Puid
				for _, vvv := range PropertyInfos {
					if vv.Puid == vvv.Uid {
						res[k].RoomType.Property = &pms.ReservationRoomUtilListProperty{}
						res[k].RoomType.Property.Name = vvv.Name
						res[k].RoomType.Property.Cover = vvv.Cover
						res[k].RoomType.Property.Uid = vvv.Uid
						break
					}
				}
				break
			}
		}
	}
	return
}
func (c *ControllerPms) AppReservationReferrerList(ctx context.Context, req *pms.AppReservationReferrerListReq) (res *pms.AppReservationReferrerListRes, err error) {
	list, totalCount, err := service.HotelService().ReservationReferrerList(ctx, &req.PmsAppReservationReferrerListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsAppReservationReferrerListModel{}
	}

	res = new(pms.AppReservationReferrerListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) AppStayInfo(ctx context.Context, req *pms.AppStayInfoReq) (res *pms.AppStayInfoRes, err error) {
	data, err := service.HotelService().AppStayInfo(ctx, &req.PmsAppReservationViewInp)
	if err != nil {
		return
	}

	res = new(pms.AppStayInfoRes)
	res.PmsAppStayInfoModel = data
	return
}
