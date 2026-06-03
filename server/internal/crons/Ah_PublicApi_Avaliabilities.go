package crons

import (
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/model/entity"
	"APT/utility/convert"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"strings"
	"time"
)

type PmsRoomTypes struct {
	*entity.PmsRoomType
	Properties struct {
		gmeta.Meta    `orm:"table:hg_pms_property"`
		Uid           string `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
		MaxDaysNotice int    `json:"maxDaysNotice"        orm:"max_days_notice"         description:"最大预定区间"`
	} `json:"properties" orm:"with:uid=puid"`
}

func AhPublicApiAvailabilities(ctx context.Context, puid string, tuid string) (err error) {
	var (
		StartDate string
		EndDate   string

		roomTypeInfos        []*PmsRoomTypes
		response             *airhousePublicApi.AvailabilitiesJSONDataResponse
		RoomUnit             []*entity.PmsRoomUnit
		RoomUnitStatus       []*entity.PmsRoomStatus
		Availabilities       []*entity.PmsAvailabilities
		InsertRoomUnitStatus []*entity.PmsRoomStatus
		InsertRoomTypeStatus []*entity.PmsAvailabilities
		RoomUnitStatusMap    map[string]*entity.PmsRoomStatus
		RoomTypeStatusMap    map[string]*entity.PmsAvailabilities
		InsertId             int64
		Dates                []*convert.DateRange
		Logger               = g.Log().Path("logs/CRON/AhPublicApiAvailabilities")
	)
	Logger.Info(ctx, "--[执行开始]----------------------------------------")
	PmsRoomType := g.MapStrAny{}
	if !g.IsEmpty(puid) {
		PmsRoomType[dao.PmsRoomType.Columns().Puid] = strings.Split(puid, ",")
	}
	if !g.IsEmpty(tuid) {
		PmsRoomType[dao.PmsRoomType.Columns().Uid] = tuid
	}
	g.DB().SetLogger(Logger)
	if err = dao.PmsRoomType.Ctx(ctx).WithAll().OmitEmptyWhere().Where(PmsRoomType).Scan(&roomTypeInfos); err != nil {
		Logger.Error(ctx, err)
		return
	}
	for _, v := range roomTypeInfos {
		Logger.Infof(ctx, "%s  -   %s", v.Puid, v.Uid)
	}
	if g.IsEmpty(roomTypeInfos) {
		Logger.Error(ctx, "没有找到房型信息")
		return
	}
	for index, roomTypeInfo := range roomTypeInfos {
		StartDate = gtime.Now().Format("Y-m-d")
		//EndDate = gtime.New(StartDate).Add(time.Duration(roomTypeInfo.Properties.MaxDaysNotice*24) * time.Hour).Format("Y-m-d")
		EndDate = gtime.New(StartDate).Add(time.Duration(180*24) * time.Hour).Format("Y-m-d")
		Logger.Infof(ctx, "%d 物业 %s 房型 %s 开始日期  ：  %s   结束日期 :  %s\n", index, roomTypeInfo.Puid, roomTypeInfo.Uid, StartDate, EndDate)

		// 初始化库存信息
		Availabilities = nil
		if err = dao.PmsAvailabilities.Ctx(ctx).
			WhereBetween(dao.PmsAvailabilities.Columns().Date, StartDate, EndDate).
			Where(dao.PmsAvailabilities.Columns().Tuid, roomTypeInfo.Uid).
			Scan(&Availabilities); err != nil && !errors.Is(err, sql.ErrNoRows) {
			Logger.Error(ctx, err)
			continue
		}
		RoomTypeStatusMap = make(map[string]*entity.PmsAvailabilities)
		for _, Availabilitie := range Availabilities {
			RoomTypeStatusMap[Availabilitie.Date] = Availabilitie
		}

		InsertRoomTypeStatus = nil
		for i := 0; i < convert.Diff(StartDate, EndDate, "days"); i++ {
			Date := gtime.New(StartDate).Add(time.Duration(i*24) * time.Hour).Format("Y-m-d")
			if _, exists := RoomTypeStatusMap[Date]; !exists {
				InsertRoomTypeStatus = append(InsertRoomTypeStatus, &entity.PmsAvailabilities{
					Puid: roomTypeInfo.Puid,
					Tuid: roomTypeInfo.Uid,
					Date: Date,
				})
			}
		}
		if !g.IsEmpty(InsertRoomTypeStatus) {
			if InsertId, err = dao.PmsAvailabilities.Ctx(ctx).Batch(100).OmitEmpty().InsertAndGetId(InsertRoomTypeStatus); err != nil {
				Logger.Error(ctx, err)
				continue
			}
			if g.IsEmpty(InsertId) {
				Logger.Error(ctx, err)
				continue
			}
		}

		if err = dao.PmsRoomUnit.Ctx(ctx).Where(dao.PmsRoomUnit.Columns().RtUid, roomTypeInfo.Uid).Scan(&RoomUnit); err != nil {
			Logger.Error(ctx, err)
			continue
		}
		if g.IsEmpty(RoomUnit) {
			Logger.Error(ctx, "房间单元为空")
			continue
		}

		if Dates, err = convert.GenerateDateRanges(StartDate, EndDate, 25); err != nil {
			Logger.Error(ctx, err)
			return
		}

		for _, DateItem := range Dates {
			ToStartDate := DateItem.StartDate.Format("Y-m-d")
			ToEndDate := DateItem.EndDate.Format("Y-m-d")
			Logger.Infof(ctx, "-- %d 物业 %s 房型 %s 开始日期  ：  %s   结束日期 :  %s\n", index, roomTypeInfo.Puid, roomTypeInfo.Uid, ToStartDate, ToEndDate)
			if response, err = airhousePublicApi.GetAvailabilities(ctx, &airhousePublicApi.AvailabilitiesJSONDataRequest{
				RoomTypeId: roomTypeInfo.Uid,
				StartDate:  ToStartDate,
				EndDate:    ToEndDate,
			}); err != nil {
				Logger.Error(ctx, err)
				break
			}
			// 插入不存在的房态信息
			for _, roomUnit := range RoomUnit {
				RoomUnitStatus = nil
				if err = dao.PmsRoomStatus.Ctx(ctx).
					WhereBetween(dao.PmsRoomStatus.Columns().Date, StartDate, EndDate).
					Where(dao.PmsRoomStatus.Columns().Puid, roomTypeInfo.Puid).
					Where(dao.PmsRoomStatus.Columns().Tuid, roomTypeInfo.Uid).
					Where(dao.PmsRoomStatus.Columns().Ruid, roomUnit.Uid).
					Scan(&RoomUnitStatus); err != nil && !errors.Is(err, sql.ErrNoRows) {
					Logger.Error(ctx, err)
					continue
				}
				Logger.Infof(ctx, "房态数据：%d", len(RoomUnitStatus))
				RoomUnitStatusMap = make(map[string]*entity.PmsRoomStatus)
				for _, roomUnitItem := range RoomUnitStatus {
					RoomUnitStatusMap[roomUnitItem.Date.Format("Y-m-d")] = roomUnitItem
				}
				Logger.Infof(ctx, "房态MAP数据：%d", len(RoomUnitStatusMap))
				InsertRoomUnitStatus = nil
				for i := 0; i < convert.Diff(ToStartDate, ToEndDate, "days"); i++ {
					Date := gtime.New(ToStartDate).Add(time.Duration(i*24) * time.Hour).Format("Y-m-d")
					Logger.Infof(ctx, "--  [%d] -- 物业 %s 房型 %s 开始日期  ：  %s   结束日期 :  %s  检测日期 %s \n", index, roomTypeInfo.Puid, roomTypeInfo.Uid, ToStartDate, ToEndDate, Date)
					if g.IsEmpty(RoomUnitStatusMap[Date]) {
						InsertRoomUnitStatus = append(InsertRoomUnitStatus, &entity.PmsRoomStatus{
							Puid:   roomTypeInfo.Puid,
							Tuid:   roomTypeInfo.Uid,
							Ruid:   roomUnit.Uid,
							RoomNo: roomUnit.RoomNo,
							Date:   gtime.New(Date),
							Status: "WAIT",
						})
					}
				}
				if !g.IsEmpty(InsertRoomUnitStatus) {
					if InsertId, err = dao.PmsRoomStatus.Ctx(ctx).Batch(100).OmitEmptyData().InsertAndGetId(InsertRoomUnitStatus); err != nil {
						Logger.Error(ctx, err)
						continue
					}
					if g.IsEmpty(InsertId) {
						Logger.Error(ctx, err)
						continue
					}
				}
			}

			// 批量更新房型库存信息和价格
			for _, AvailabilitiesItem := range response.Data.Inventories {
				// 批量更新库存
				if _, err = dao.PmsAvailabilities.Ctx(ctx).
					WhereBetween(dao.PmsAvailabilities.Columns().Date, AvailabilitiesItem.StartDate, AvailabilitiesItem.EndDate).
					Where(dao.PmsAvailabilities.Columns().Tuid, roomTypeInfo.Uid).
					Data(g.MapStrAny{
						dao.PmsAvailabilities.Columns().Allotment: AvailabilitiesItem.InventoryCount,
					}).UpdateAndGetAffected(); err != nil {
					Logger.Error(ctx, err)
					continue
				}
			}
			for _, RatePlan := range response.Data.Rates {
				if roomTypeInfo.RatePlanId == RatePlan.RatePlan.ID {
					// 批量更新价格
					if _, err = dao.PmsAvailabilities.Ctx(ctx).
						WhereBetween(dao.PmsAvailabilities.Columns().Date, RatePlan.StartDate, RatePlan.EndDate).
						Where(dao.PmsAvailabilities.Columns().Tuid, roomTypeInfo.Uid).
						Data(&entity.PmsAvailabilities{
							Price:     gvar.New(RatePlan.PerDayBaseRate).Float64(),
							UpdatedAt: gtime.Now().Add(time.Duration(10) * time.Second),
						}).OmitEmptyData().UpdateAndGetAffected(); err != nil {
						Logger.Error(ctx, err)
						continue
					}
				}

				if RatePlan.Closed == "true" && roomTypeInfo.RatePlanId == RatePlan.RatePlan.ID {
					if _, err = dao.PmsAvailabilities.Ctx(ctx).
						WhereBetween(dao.PmsAvailabilities.Columns().Date, RatePlan.StartDate, RatePlan.EndDate).
						Where(dao.PmsAvailabilities.Columns().Tuid, roomTypeInfo.Uid).
						Data(dao.PmsAvailabilities.Columns().Allotment, 0).UpdateAndGetAffected(); err != nil {
						Logger.Error(ctx, err)
						continue
					}
				}
			}
			// 先批量解锁所有房间
			//if _, err = dao.PmsRoomStatus.Ctx(ctx).
			//	Where(dao.PmsRoomStatus.Columns().Tuid, roomTypeInfo.Uid).
			//	Where(dao.PmsRoomStatus.Columns().Status, "LOCK").
			//	OmitEmptyData().
			//	UpdateAndGetAffected(&entity.PmsRoomStatus{
			//		Status: "WAIT",
			//	}); err != nil {
			//	Logger.Error(ctx, err)
			//	continue
			//}
			// 批量更新房态信息
			//for _, Blocks := range response.Data.Blocks {
			//	if Affected, err = dao.PmsRoomStatus.Ctx(ctx).
			//		WhereBetween(dao.PmsRoomStatus.Columns().Date, Blocks.StartDate, Blocks.EndDate).
			//		Where(dao.PmsRoomStatus.Columns().Ruid, Blocks.RoomUnit.ID).
			//		OmitEmptyData().
			//		Data(&entity.PmsRoomStatus{
			//			Status: "LOCK",
			//		}).UpdateAndGetAffected(); err != nil {
			//		Logger.Error(ctx, err)
			//		continue
			//	}
			//	if g.IsEmpty(Affected) {
			//		Logger.Error(ctx, "更新失败")
			//		continue
			//	}
			//}
		}
	}
	Logger.Info(ctx, "执行成功")
	return
}
