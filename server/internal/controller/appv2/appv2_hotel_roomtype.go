package appv2

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"

	"APT/api/appv2/hotel"
)

func (c *ControllerHotel) RoomTypeList(ctx context.Context, req *hotel.RoomTypeListReq) (res *hotel.RoomTypeListRes, err error) {
	var (
		tuids             []string
		PmsAvailabilities []*struct {
			Tuid      string  `json:"tuid"`
			Price     float64 `json:"price"`
			Allotment int     `json:"allotment"`
		}
		BookingPriceModel *input_pay.BookingPriceModel
		CancelRate        []*entity.PmsCancelRate
		PmsPirceConfig    *model.PmsPriceConfig
	)
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	if g.IsEmpty(PmsPirceConfig) {
		err = errors.New("价格配置不存在")
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	res = new(hotel.RoomTypeListRes)
	if err = dao.PmsRoomType.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).
		Where(dao.PmsRoomType.Columns().Puid, req.Puid).
		Where(dao.PmsRoomType.Columns().IsShow, 1).
		Scan(&res.List); err != nil {
		return
	}

	if !g.IsEmpty(&res.List) {
		for _, v := range res.List {
			tuids = append(tuids, v.Uid)
		}
		// 批量查询物业在时间段中的最低价格
		if err = dao.PmsAvailabilities.Ctx(ctx).
			Fields(dao.PmsAvailabilities.Columns().Tuid, "min(price) as price", "min(allotment) as allotment").
			WhereIn(dao.PmsAvailabilities.Columns().Tuid, tuids).
			WhereGT(dao.PmsAvailabilities.Columns().Price, 0).
			WhereGTE(dao.PmsAvailabilities.Columns().Date, req.StartDate).
			WhereLT(dao.PmsAvailabilities.Columns().Date, req.EndDate).
			Group(dao.PmsAvailabilities.Columns().Tuid).Scan(&PmsAvailabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		for k, v := range res.List {
			for _, v2 := range PmsAvailabilities {
				if v.Uid == v2.Tuid {
					itemBookingPrice := input_pay.BookingPriceInp{
						CheckinAt:   req.StartDate,
						CheckoutAt:  req.EndDate,
						Puid:        req.Puid,
						IsPricePlan: true,
						RoomInfos: &input_hotel.RoomItems{
							RoomId:  v.Uid,
							RoomNum: req.RoomNumber,
							Adult:   1,
							Child:   0,
							Infant:  0,
						},
					}
					if BookingPriceModel, err = service.PayService().BookingPrice(ctx, &itemBookingPrice); err != nil {
						return
					}
					// 计算不对 BasePrice作废
					res.List[k].BasePrice = decimal.NewFromFloat(BookingPriceModel.PlanPriceTotalAmount).Mul(PricePercent).Round(0).InexactFloat64()
					//res.List[k].BasePrice = BookingPriceModel.PlanPriceTotalAmount
					if g.IsEmpty(BookingPriceModel.PlanPrice) {
						PricePlan := &entity.PmsPricePlan{
							PropertyId:      req.Puid,
							RoomTypeId:      v.Uid,
							BookingDays:     0,
							IsCancel:        "Y",
							IsOpenPriceMode: "N",
						}
						BookingPriceInfo := &input_pay.BookingPriceItem{
							PmsPricePlan:         PricePlan,
							PlanShowNameLanguage: nil,
							TotalAmount:          BookingPriceModel.TotalPrice,
						}

						BookingPriceModel.PlanPrice = append(BookingPriceModel.PlanPrice, BookingPriceInfo)
						BookingPriceModel.MinTotalPrice = BookingPriceInfo.TotalAmount
					}
					res.List[k].PricePlans = BookingPriceModel
					res.List[k].Inventory = v2.Allotment
					if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
						return
					}
					for _, RateItem := range CancelRate {
						NowTime := gtime.Now().Format("Y-m-d")
						switch RateItem.Mode {
						case "after":
							startDate := gtime.New(req.StartDate).Add(time.Duration(-RateItem.StartDays*24) * time.Hour).Format("Y-m-d")
							//res.CancelDate = append(res.CancelDate, fmt.Sprintf("%s%s", startDate, v.Name))
							if len(res.List[k].CancelDate) >= 1 {
								res.List[k].CancelDate = append(res.List[k].CancelDate, fmt.Sprintf("%s%s", startDate, RateItem.Name))
							} else {
								res.List[k].CancelDate = append(res.List[k].CancelDate, fmt.Sprintf("%s", RateItem.Name))
							}
							break
						case "middle":
							startDay := gtime.New(req.StartDate).Add(time.Duration(-RateItem.StartDays*24) * time.Hour).Format("Y-m-d")
							endDate := gtime.New(req.StartDate).Add(time.Duration(-RateItem.EndDays*24) * time.Hour).Format("Y-m-d")
							if NowTime <= endDate {
								res.List[k].CancelDate = append(res.List[k].CancelDate, fmt.Sprintf("%s%s", fmt.Sprintf("%s ~ %s", startDay, endDate), RateItem.Name))
							}
							break
						case "before":
							endDate := gtime.New(req.StartDate).Add(time.Duration(-RateItem.EndDays*24) * time.Hour).Format("Y-m-d")
							if endDate >= NowTime {
								res.List[k].CancelDate = append(res.List[k].CancelDate, fmt.Sprintf("%s%s", endDate, RateItem.Name))
							}
							break
						}
						if len(res.List[k].CancelDate) > 0 {
							break
						}
					}
				}
			}
		}
	}

	return
}
