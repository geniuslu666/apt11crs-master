// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package crons

import (
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

var (
	IsUpdateRoomUnit = "N"
)

// AhPublicApiProperties 执行任务
func AhPublicApiProperties(ctx context.Context, Uid string) (err error) {
	var (
		ListProperties     airhousePublicApi.ListPropertiesJSONDataResponse
		RetrieveProperties airhousePublicApi.RetrievePropertiesJSONDataResponse
		Logger             = g.Log().Path("logs/CRON/AhPublicApiProperties")
	)
	Logger.Info(ctx, "--[执行开始]----------------------------------------")
	Logger.Info(ctx, "是否需要更新房型间数  Y/N")
	fmt.Scan(&IsUpdateRoomUnit)
	g.DB().SetLogger(Logger)
	if !g.IsEmpty(Uid) {
		if RetrieveProperties, err = airhousePublicApi.GetRetrieveProperties(ctx, Uid); err != nil {
			return
		}
		if err = SyncProperty(ctx, RetrieveProperties.Data); err != nil {
			Logger.Error(ctx, err)
		}
	} else {
		if ListProperties, err = airhousePublicApi.GetListProperties(ctx, 100, nil); err != nil {
			return
		}
		for _, v := range ListProperties.Data {
			if err = SyncProperty(ctx, v); err != nil {
				Logger.Error(ctx, err)
			}
		}
	}
	Logger.Info(ctx, "执行成功")
	return err
}

func SyncProperty(ctx context.Context, HandleData airhousePublicApi.PropertiesData) (err error) {
	var (
		PmsPropertyInfo    *entity.PmsProperty
		PmsRoomTypeInfo    *entity.PmsRoomType
		PmsRoomUnitInfo    *entity.PmsRoomUnit
		updateResult       sql.Result
		insertResult       sql.Result
		lastInsertId       int64
		count              int64
		Logger             = g.Log().Path("logs/CRON/AhPublicApiProperties")
		PropertyHandleData g.MapStrAny
	)
	Logger.Info(ctx, "同步物业信息")
	PmsPropertyInfo = nil
	// 查询是否存在该物业
	if err = dao.PmsProperty.Ctx(ctx).Where(dao.PmsProperty.Columns().Uid, HandleData.ID).Scan(&PmsPropertyInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if g.IsEmpty(PmsPropertyInfo) {
		PropertyNameUid := guid.S([]byte(dao.PmsProperty.Columns().Name))
		PropertyHandleData = g.MapStrAny{
			dao.PmsProperty.Columns().Uid:           HandleData.ID,
			dao.PmsProperty.Columns().Name:          PropertyNameUid,
			dao.PmsProperty.Columns().Currency:      HandleData.Currency,      // 默认货币
			dao.PmsProperty.Columns().Language:      HandleData.Language,      // 默认语言
			dao.PmsProperty.Columns().TimeZone:      HandleData.TimeZone,      // 默认时区
			dao.PmsProperty.Columns().MaxDaysNotice: HandleData.BookingWindow, // 最大可预定
		}
		// 插入
		if insertResult, err = dao.PmsProperty.Ctx(ctx).Data().Insert(PropertyHandleData); err != nil {
			return err
		}
		if lastInsertId, err = insertResult.LastInsertId(); err != nil {
			return err
		}
		Logger.Info(ctx, lastInsertId)
		// 查询物业多语言数据
		NameLanguageModel := input_language.LanguageModel{
			EN: "",
			ZH: "",
			JA: HandleData.Name,
			KO: "",
		}
		// 更新名字多语言数据
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageModel, &input_language.LoadLanguage{
			Uuid: PropertyNameUid,
			Tag:  dao.PmsProperty.Table(),
			Type: "table",
			Key:  dao.PmsProperty.Columns().Name,
		}); err != nil {
			return
		}
	} else {
		PropertyHandleData = g.MapStrAny{
			dao.PmsProperty.Columns().Uid:           HandleData.ID,
			dao.PmsProperty.Columns().Currency:      HandleData.Currency,      // 默认货币
			dao.PmsProperty.Columns().Language:      HandleData.Language,      // 默认语言
			dao.PmsProperty.Columns().TimeZone:      HandleData.TimeZone,      // 默认时区
			dao.PmsProperty.Columns().MaxDaysNotice: HandleData.BookingWindow, // 最大可预定
		}
		// 更新
		if updateResult, err = g.Model(dao.PmsProperty.Table()).
			Ctx(ctx).
			Where(dao.PmsProperty.Columns().Uid, HandleData.ID).
			Data(PropertyHandleData).
			Update(); err != nil {
			return err
		}
		if count, err = updateResult.RowsAffected(); err != nil {
			return err
		}
		Logger.Info(ctx, count)
	}

	for _, vv := range HandleData.RoomTypes {
		Logger.Info(ctx, "同步房型")
		PmsRoomTypeInfo = nil
		// 查询房型
		if err = dao.PmsRoomType.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsRoomType.Columns().Uid: vv.ID,
		}).Scan(&PmsRoomTypeInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		RoomTypeNameUid := guid.S([]byte(dao.PmsRoomType.Columns().Name))
		if g.IsEmpty(PmsRoomTypeInfo) {
			// 插入
			if insertResult, err = g.Model(dao.PmsRoomType.Table()).Ctx(ctx).Data(g.MapStrAny{
				dao.PmsRoomType.Columns().Uid:       vv.ID,
				dao.PmsRoomType.Columns().Puid:      HandleData.ID,
				dao.PmsRoomType.Columns().CoverList: "[]",
				dao.PmsRoomType.Columns().Name:      RoomTypeNameUid,
				dao.PmsRoomType.Columns().RoomNum:   len(vv.RoomUnits),
			}).Insert(); err != nil {
				return err
			}
			if lastInsertId, err = insertResult.LastInsertId(); err != nil {
				return err
			}
			Logger.Info(ctx, lastInsertId)
			// 查询物业多语言数据
			NameLanguageModel := input_language.LanguageModel{
				EN: vv.Name,
				ZH: "",
				JA: "",
				KO: "",
			}
			// 更新名字多语言数据
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageModel, &input_language.LoadLanguage{
				Uuid: RoomTypeNameUid,
				Tag:  dao.PmsRoomType.Table(),
				Type: "table",
				Key:  dao.PmsRoomType.Columns().Name,
			}); err != nil {
				return
			}
		}
		if IsUpdateRoomUnit == "Y" {
			if insertResult, err = g.Model(dao.PmsRoomType.Table()).Ctx(ctx).Where(g.Map{
				dao.PmsRoomType.Columns().Uid: vv.ID,
			}).Data(g.MapStrAny{
				dao.PmsRoomType.Columns().RoomNum: len(vv.RoomUnits),
			}).Update(); err != nil {
				return err
			}
		}
		for _, vvv := range vv.RoomUnits {
			Logger.Info(ctx, "同步房间")
			PmsRoomUnitInfo = nil
			if err = dao.PmsRoomUnit.Ctx(ctx).Where(g.MapStrAny{
				dao.PmsRoomUnit.Columns().Uid:    vvv.ID,
				dao.PmsRoomUnit.Columns().RoomNo: vvv.RoomNo,
			}).Scan(&PmsRoomUnitInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return err
			}
			if g.IsEmpty(PmsRoomUnitInfo) {
				// 插入
				if insertResult, err = g.Model(dao.PmsRoomUnit.Table()).Ctx(ctx).Data(g.MapStrAny{
					dao.PmsRoomUnit.Columns().Uid:    vvv.ID,
					dao.PmsRoomUnit.Columns().RtUid:  vv.ID,
					dao.PmsRoomUnit.Columns().RoomNo: vvv.RoomNo,
				}).Insert(); err != nil {
					return err
				}
				if lastInsertId, err = insertResult.LastInsertId(); err != nil {
					return err
				}
				Logger.Info(ctx, lastInsertId)
			} else {
				// 更新
				if updateResult, err = g.Model(dao.PmsRoomUnit.Table()).Ctx(ctx).Where(g.MapStrAny{
					dao.PmsRoomUnit.Columns().Uid:   vv.ID,
					dao.PmsRoomUnit.Columns().RtUid: vv.ID,
				}).Data(g.MapStrAny{
					dao.PmsRoomUnit.Columns().RoomNo: vvv.RoomNo,
				}).Update(); err != nil {
					return err
				}
				if count, err = updateResult.RowsAffected(); err != nil {
					return err
				}
				Logger.Info(ctx, count)
			}
		}
	}
	return
}
