package admin

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"time"

	"APT/api/admin/pms"
)

func (c *ControllerPms) PriceCalendar(ctx context.Context, req *pms.PriceCalendarReq) (res *pms.PriceCalendarRes, err error) {
	var (
		Availabilities []*entity.PmsAvailabilities
		roomTypeUid    []string
		StartDate      string
		EndDate        string
		PmsPirceConfig *model.PmsPriceConfig
	)
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	StartDate = req.Date
	EndDate = gtime.New(req.Date).Add(time.Duration(30*24) * time.Hour).Format("Y-m-d")
	res = new(pms.PriceCalendarRes)
	// 查询房型
	if err = dao.PmsRoomType.Ctx(ctx).Where(&entity.PmsRoomType{
		Puid: req.Puid,
		Uid:  req.Tuid,
	}).Page(req.Page, req.Limit).WithAll().Hook(hook.PmsFindLanguageValueHook).OmitEmptyWhere().ScanAndCount(&res.RoomTypeList, &res.Count, false); err != nil {
		return
	}
	for _, v := range res.RoomTypeList {
		roomTypeUid = append(roomTypeUid, v.Uid)
	}
	if err = dao.PmsAvailabilities.Ctx(ctx).
		WhereIn(dao.PmsAvailabilities.Columns().Tuid, roomTypeUid).
		WhereBetween(dao.PmsAvailabilities.Columns().Date, StartDate, EndDate).
		OmitEmptyWhere().
		OrderAsc(dao.PmsAvailabilities.Columns().Date).
		Scan(&Availabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	res.Availabilities = make(map[string]map[string][]*entity.PmsAvailabilities)
	for k, v := range Availabilities {
		if g.IsEmpty(res.Availabilities[v.Tuid]) {
			res.Availabilities[v.Tuid] = make(map[string][]*entity.PmsAvailabilities)
		}
		Availabilities[k].BasePrice = v.Price

		Availabilities[k].Price = decimal.NewFromFloat(v.Price).Mul(PricePercent).Round(0).InexactFloat64()
		res.Availabilities[v.Tuid][v.Date] = append(res.Availabilities[v.Tuid][v.Date], v)
	}
	return
}
