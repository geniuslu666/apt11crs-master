package app

import (
	"APT/api/app/hotel"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/convert"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"strings"
)

func (c *ControllerHotel) RoomTypeList(ctx context.Context, req *hotel.RoomTypeListReq) (res *hotel.RoomTypeListRes, err error) {
	var (
		tuids             []string
		PmsAvailabilities []*struct {
			Tuid      string  `json:"tuid"`
			Price     float64 `json:"price"`
			Allotment int     `json:"allotment"`
		}
		PmsPirceConfig *model.PmsPriceConfig
		//BookingPriceModel *input_pay.BookingPriceModel
	)
	g.Log().Debug(ctx, "overdueoverdueoverdue")
	g.Log().Debug(ctx, convert.Diff(gtime.Now().Format("Y-m-d"), req.EndDate, "days"))
	if convert.Diff(gtime.Now().Format("Y-m-d"), req.EndDate, "days") > 179 {
		err = gerror.New(gi18n.T(ctx, "stay_over_179_days_cannot_book"))
		return
	}
	if PmsPirceConfig, err = service.BasicsConfig().GetPmsPrice(ctx); err != nil {
		return
	}
	PricePercent := decimal.NewFromInt(PmsPirceConfig.PricePercent).Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1))
	res = new(hotel.RoomTypeListRes)
	if err = dao.PmsRoomType.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).
		Where(dao.PmsRoomType.Columns().Puid, req.Puid).
		Where(dao.PmsRoomType.Columns().IsShow, 1).
		WhereGTE(dao.PmsRoomType.Columns().Occupancy, req.BookerNumber).
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
					//itemBookingPrice := input_pay.BookingPriceInp{
					//	CheckinAt:  req.StartDate,
					//	CheckoutAt: req.EndDate,
					//	Puid:       req.Puid,
					//	IsPricePlan: true,
					//	RoomInfos: &input_hotel.RoomItems{
					//		RoomId:  v.Uid,
					//		RoomNum: req.RoomNumber,
					//		Adult:   1,
					//		Child:   0,
					//		Infant:  0,
					//	},
					//}
					//if BookingPriceModel, err = service.PayService().BookingPrice(ctx, &itemBookingPrice); err != nil {
					//	return
					//}

					res.List[k].BasePrice = decimal.NewFromFloat(v2.Price).Mul(PricePercent).Round(0).InexactFloat64()
					//res.List[k].PricePlans = BookingPriceModel
					res.List[k].Inventory = v2.Allotment
				}
			}
		}
	}

	return
}
func (c *ControllerHotel) RoomTypeView(ctx context.Context, req *hotel.RoomTypeViewReq) (res *hotel.RoomTypeViewRes, err error) {
	var (
		BedTypeArr []string
		BedTypeStr []string
	)
	res = new(hotel.RoomTypeViewRes)
	if res.PmsRoomTypeViewModel, err = service.HotelService().RoomTypeView(ctx, &input_hotel.PmsRoomTypeViewInp{
		Id:  req.Id,
		Uid: req.Uid,
	}); err != nil {
		return
	}

	for _, item := range res.PmsRoomTypeViewModel.BedType {
		itemStr := ""
		itemNewStr := ""
		switch contexts.GetLanguage(ctx) {
		case "zh":
			itemStr = fmt.Sprintf("%.0f张%.1f床", item.BedNum, item.BedWidth)
			break
		case "en":
			itemStr = fmt.Sprintf("%.0f %.1f-meter bed", item.BedNum, item.BedWidth)
			break
		case "zh_CN":
			itemStr = fmt.Sprintf("%.0f張%.1f公尺的床", item.BedNum, item.BedWidth)
			break
		case "ko":
			itemStr = fmt.Sprintf("%.0f %.1f미터 침대", item.BedNum, item.BedWidth)
			break
		case "ja":
			itemStr = fmt.Sprintf("%.0f %.1fメートルのベッド", item.BedNum, item.BedWidth)
			break
		}
		itemNewStr = fmt.Sprintf("%.1fm%s * %.0f", item.BedWidth, item.BedTypeName, item.BedNum)

		BedTypeArr = append(BedTypeArr, itemStr)
		BedTypeStr = append(BedTypeStr, itemNewStr)
	}
	res.BedTypeDes = strings.Join(BedTypeArr, ",")
	res.BedTypeStr = strings.Join(BedTypeStr, ", ")
	res.BedType = nil
	return
}
