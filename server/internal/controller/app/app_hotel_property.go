package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"sort"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) PropertySearch(ctx context.Context, req *hotel.PropertySearchReq) (res *hotel.PropertySearchRes, err error) {
	var (
		PmsAvailabilities []*struct {
			Puid      string  `json:"puid"      orm:"puid"       description:"物业UID"`
			Allotment int     `json:"allotment" orm:"allotment"  description:"库存"`
			Price     float64 `json:"price"     orm:"price"      description:"价格"`
		}
		orm            = dao.PmsProperty.Ctx(ctx)
		puids          []string
		MemberInfo     = contexts.GetMemberUser(ctx)
		WhereOr        *gdb.WhereBuilder
		PmsRoomType    []*entity.PmsRoomType
		PmsPirceConfig *model.PmsPriceConfig
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
	if g.IsEmpty(req.IsLease) {
		req.IsLease = 1
	}
	res = new(hotel.PropertySearchRes)
	WhereOr = orm.Builder()
	WhereOr = WhereOr.WhereNull(dao.PmsProperty.Columns().GroupIds)
	WhereOr = WhereOr.WhereOr(dao.PmsProperty.Columns().GroupIds, "")
	//orm = orm.WhereNull(dao.PmsProperty.Columns().GroupIds)
	//orm = orm.WhereOr(dao.PmsProperty.Columns().GroupIds, "")
	if !g.IsEmpty(MemberInfo) && !g.IsEmpty(MemberInfo.GroupId) {
		WhereOr = WhereOr.WhereOr("FIND_IN_SET(?, group_ids) > 0", MemberInfo.GroupId)
	}
	orm = orm.Where(dao.PmsProperty.Columns().Close, 1)
	if req.IsLease == 1 {
		orm = orm.Where(dao.PmsProperty.Columns().BookingClose, 1)
	} else if req.IsLease == 2 {
		orm = orm.Where(dao.PmsProperty.Columns().LeaseClose, 1)
	} else {
		err = errors.New("模式错误")
	}
	if !g.IsEmpty(req.RegionId) {
		orm = orm.Where(dao.PmsProperty.Columns().RegionId, req.RegionId)
	}
	orm = orm.OrderDesc(dao.PmsProperty.Columns().Sort)
	if err = orm.Where(WhereOr).WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&res.List); err != nil {
		return
	}
	for _, v := range res.List {
		puids = append(puids, v.Uid)
	}
	// 批量查询物业在时间段中的最低价格
	if err = dao.PmsAvailabilities.Ctx(ctx).
		Fields("puid", "tuid", "min(price) as price", "min(allotment) as allotment").
		WhereIn(dao.PmsAvailabilities.Columns().Puid, puids).
		WhereGT(dao.PmsAvailabilities.Columns().Price, 0).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, req.StartDate).
		WhereLT(dao.PmsAvailabilities.Columns().Date, req.EndDate).
		Group(dao.PmsAvailabilities.Columns().Puid, dao.PmsAvailabilities.Columns().Tuid).Scan(&PmsAvailabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(PmsAvailabilities) {
		for k, v := range res.List {
			for _, v2 := range PmsAvailabilities {
				if v.Uid == v2.Puid {
					res.List[k].Price = decimal.NewFromFloat(v2.Price).Mul(PricePercent).Round(0).InexactFloat64()
					if v2.Allotment > res.List[k].Inventory {
						res.List[k].Inventory = v2.Allotment
						if res.List[k].Inventory <= 0 {
							res.List[k].IsDisable = true
						} else {
							res.List[k].IsDisable = false
						}
					}
				}
			}
		}
	}

	// 查询所有物业的可容纳人数排序
	if err = dao.PmsRoomType.Ctx(ctx).
		OrderDesc("max(occupancy)").
		Group(dao.PmsRoomType.Columns().Puid).
		Fields(dao.PmsRoomType.Columns().Puid, "max(occupancy) as occupancy").Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	days := convert.Diff(req.StartDate, req.EndDate, "days")
	for k, v := range res.List {
		for _, v2 := range PmsRoomType {
			if v.Uid == v2.Puid {
				res.List[k].Occupancy = v2.Occupancy
				if req.BookerNumber > v2.Occupancy {
					res.List[k].IsDisable = true
				} else if req.RoomNumber > 0 && req.BookerNumber > v2.Occupancy {
					res.List[k].IsDisable = true
				} else if req.AdultNumber+req.ChildNumber > 0 && req.AdultNumber+req.ChildNumber > v2.Occupancy {
					res.List[k].IsDisable = true
				}
			}
		}
		if req.IsLease == 2 && v.MinDaysNotice > 0 && days < v.MinDaysNotice {
			res.List[k].IsDisable = true
		}
	}

	List := res.List
	res.List = []*hotel.PropertyOnlineSearchItem{}
	for _, item := range List {
		if !item.IsDisable {
			res.List = append(res.List, item)
		}
	}

	sort.Slice(res.List, func(i, j int) bool {
		return res.List[i].Occupancy > res.List[j].Occupancy
	})
	sort.Slice(res.List, func(i, j int) bool {
		return res.List[i].Inventory > res.List[j].Inventory
	})

	// 如果是短租的情况下查询是否满足短租最低要求  不满足禁用

	return
}
func (c *ControllerHotel) PropertyView(ctx context.Context, req *hotel.PropertyViewReq) (res *hotel.PropertyViewRes, err error) {
	var (
		CollectId  *gvar.Var
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(hotel.PropertyViewRes)
	if res.PmsPropertyViewModel, _, err = service.HotelService().PropertyView(ctx, &req.PmsPropertyViewInp); err != nil {
		return
	}
	where := g.MapStrAny{
		dao.PmsCollect.Columns().CollectId:   res.PmsPropertyViewModel.Id,
		dao.PmsCollect.Columns().CollectType: "HOST",
	}
	if !g.IsEmpty(MemberInfo) {
		where[dao.PmsCollect.Columns().MemberId] = MemberInfo.PmsMember.Id
	}
	if CollectId, err = dao.PmsCollect.Ctx(ctx).Where(where).OmitEmptyWhere().Value(dao.PmsCollect.Columns().Id); err != nil {
		return
	}
	if CollectId.IsEmpty() {
		res.CollectId = 0
	} else {
		res.CollectId = CollectId.Int()
	}

	res.JaRate = 0.05
	// 获取配置信息
	var Config *input_basics.GetConfigModel
	if Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "rmbExchangeRate",
	}); err != nil {
		return
	}
	if g.IsEmpty(Config) {
		return
	}
	val, ok := Config.List["jaRate"].(float64)
	if !ok {
		return
	}
	res.JaRate = val

	return
}
