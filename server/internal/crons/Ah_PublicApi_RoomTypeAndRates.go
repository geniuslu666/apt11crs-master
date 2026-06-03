package crons

import (
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func AhPublicApiRoomTypeAndRates(ctx context.Context) (err error) {
	var (
		ListRoomTypesAndRates *airhousePublicApi.ListRoomTypesAndRatesJSONDataResponse
		PmsRoomType           []*entity.PmsRoomType
		UpdateRow             int64
		Logger                = g.Log().Path("logs/CRON/AhPublicApiRoomTypeAndRates")
		IsUpateRoomRatePlanId string
		IsSyncRatePlanId      string
	)
	Logger.Info(ctx, "--[执行开始]----------------------------------------")
	g.DB().SetLogger(Logger)
	fmt.Println("是否需要更新费率计划数据 Y/N")
	fmt.Scan(&IsSyncRatePlanId)
	fmt.Println("是否需要更新现有房型费率计划 Y/N")
	fmt.Scan(&IsUpateRoomRatePlanId)
	if err = dao.PmsRoomType.Ctx(ctx).Scan(&PmsRoomType); err != nil && !errors.Is(err, sql.ErrNoRows) {
		Logger.Error(ctx, "同步房型和费率计划数据失败", err)
		return
	}

	for key, v := range PmsRoomType {
		// 查询房型和费率计划
		if ListRoomTypesAndRates, err = airhousePublicApi.GetListRoomTypesAndRates(ctx, &airhousePublicApi.ListRoomTypesAndRatesRequest{
			PropertyId: v.Puid,
			RoomTypeId: v.Uid,
		}); err != nil {
			Logger.Error(ctx, "同步房型和费率计划数据失败 物业 ID:"+v.Puid, err)
			continue
		}
		for _, ListRoomTypesAndRatesData := range ListRoomTypesAndRates.Data {
			RatePlanId := ""
			RatePlanIndex := 0
			for RatePlansIndex, RatePlans := range ListRoomTypesAndRatesData.RatePlans {
				if IsSyncRatePlanId == "Y" {

					// 同步本地费率计划ID库
					PmsRoomRatePlan := new(entity.PmsRoomRatePlan)
					if err = dao.PmsRoomRatePlan.Ctx(ctx).Where(g.Map{
						dao.PmsRoomRatePlan.Columns().RatePlanId: RatePlans.ID,
					}).Scan(&PmsRoomRatePlan); err != nil && !errors.Is(err, sql.ErrNoRows) {
						return
					}

					if g.IsEmpty(PmsRoomRatePlan.Id) {
						// 插入计划
						if _, err = dao.PmsRoomRatePlan.Ctx(ctx).Insert(g.MapStrAny{
							dao.PmsRoomRatePlan.Columns().Tuid:       ListRoomTypesAndRatesData.RoomType.ID,
							dao.PmsRoomRatePlan.Columns().TName:      ListRoomTypesAndRatesData.RoomTypeName,
							dao.PmsRoomRatePlan.Columns().RateName:   RatePlans.Name,
							dao.PmsRoomRatePlan.Columns().RatePlanId: RatePlans.ID,
						}); err != nil {
							return
						}
					} else {
						if _, err = dao.PmsRoomRatePlan.Ctx(ctx).Where(g.Map{
							dao.PmsRoomRatePlan.Columns().RatePlanId: RatePlans.ID,
						}).Update(g.MapStrAny{
							dao.PmsRoomRatePlan.Columns().Tuid:       ListRoomTypesAndRatesData.RoomType.ID,
							dao.PmsRoomRatePlan.Columns().TName:      ListRoomTypesAndRatesData.RoomTypeName,
							dao.PmsRoomRatePlan.Columns().RateName:   RatePlans.Name,
							dao.PmsRoomRatePlan.Columns().RatePlanId: RatePlans.ID,
						}); err != nil {
							return
						}
					}
				}
				if gstr.HasPrefix(gstr.Trim(RatePlans.Name), "Standard") {
					RatePlanId = gvar.New(RatePlans.ID).String()
					RatePlanIndex = RatePlansIndex
				}
			}
			if IsUpateRoomRatePlanId == "Y" {
				if RatePlanId == v.RatePlanId {
					Logger.Error(ctx, "同步房型和费率 RatePlansId 一致无需同步")
					continue
				}
				// 更新房型信息
				if UpdateRow, err = dao.PmsRoomType.Ctx(ctx).Where(g.MapStrAny{
					dao.PmsRoomType.Columns().Uid: ListRoomTypesAndRatesData.RoomType.ID,
				}).OmitEmptyData().Data(&entity.PmsRoomType{
					BasePrice:              gvar.New(ListRoomTypesAndRatesData.RoomTypeMinBaseRate).Float64(),
					RatePlanId:             RatePlanId,
					CleaningFee:            0,
					AdditionalGuestAmounts: gvar.New(ListRoomTypesAndRatesData.RatePlans[RatePlanIndex].PerDayPricing.AdditionalGuestAmounts[0].Amount).Float64(),
				}).UpdateAndGetAffected(); err != nil {
					Logger.Error(ctx, "同步房型和费率计划数据失败 物业 ID:"+v.Uid, err)
					return
				}
				if UpdateRow != 1 {
					Logger.Error(ctx, "同步房型和费率计划数据失败 物业 ID:"+v.Uid, err)
					continue
				}
			}
		}
		Logger.Infof(ctx, "[INDEX]:%d", key)
	}
	Logger.Info(ctx, "执行成功")
	return
}
