package inner

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/inner/order"
)

func (c *ControllerOrder) OrderList(ctx context.Context, req *order.OrderListReq) (res *order.OrderListRes, err error) {
	var (
		Order        []*entity.PmsAppReservation
		Puid         []string
		PuidInfo     []*entity.PmsProperty
		RoomType     []string
		RoomTypeInfo []*entity.PmsRoomType
		mod          = dao.PmsAppReservation.Ctx(ctx)
	)
	contexts.SetLanguage(ctx, "zh")
	res = new(order.OrderListRes)
	mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckinDate, req.StartTime, req.EndTime)
	mod = mod.WhereBetween(dao.PmsAppReservation.Columns().CheckoutDate, req.StartTime, req.EndTime)
	mod = mod.Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID")
	if req.NextId > 0 {
		mod = mod.WhereGT(dao.PmsAppReservation.Columns().Id, req.NextId)
	}
	mod = mod.Limit(req.Limit)
	if err = mod.Scan(&Order); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	for _, v := range Order {
		Puid = append(Puid, v.Puid)
		RoomType = append(RoomType, v.RoomType)
		res.List = append(res.List, &order.OrderListItem{
			Puid:         v.Puid,
			RoomType:     v.RoomType,
			SourceCode:   v.SourceCode,
			SourceName:   v.SourceName,
			OrderSn:      v.OrderSn,
			OrderIndex:   v.OrderIndex,
			OutOrderSn:   v.OutOrderSn,
			BookingFee:   v.BookingFee,
			CheckinDate:  v.CheckinDate,
			CheckoutDate: v.CheckoutDate,
			CreateTime:   v.CreatedAt,
		})
	}
	if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).WhereIn(dao.PmsProperty.Columns().Uid, Puid).Scan(&PuidInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	g.Log().Debug(ctx, "PuidInfo", PuidInfo)
	if err = dao.PmsRoomType.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).WhereIn(dao.PmsRoomType.Columns().Uid, RoomType).Scan(&RoomTypeInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	g.Log().Debug(ctx, "RoomTypeInfo", RoomTypeInfo)
	for i, v := range res.List {
		for _, v1 := range PuidInfo {
			if v.Puid == v1.Uid {
				res.List[i].PuidName = v1.Name
			}
		}
		for _, v2 := range RoomTypeInfo {
			if v.RoomType == v2.Uid {
				res.List[i].RoomTypeName = v2.Name
			}
		}
	}
	if len(res.List) < req.Limit {
		res.NextId = -1
	} else {
		res.NextId = Order[len(Order)-1].Id
	}
	return
}
