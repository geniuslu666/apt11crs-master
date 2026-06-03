package hook

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var PmsFindLanguageValueHook = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		g.Log().Info(ctx, "————————————————————————————————————————————————————————")
		g.Log().Info(ctx, in.Table)
		g.Log().Info(ctx, dao.PmsRoomUnit.Table())
		g.Log().Info(ctx, dao.PmsRoomType.Table())
		result, err = in.Next(ctx)
		if g.IsEmpty(result) {
			return
		}
		if strings.Contains(in.Table, dao.PmsRoomUnit.Table()) {
			return
		}

		//if strings.Contains(in.Table, dao.FoodLabel.Table()) {
		//	return
		//}
		var (
			LanguageFields = []string{
				"name",
				"description",
				"style",
				"tag_list",
				"room_des",
				"surroundings",
				"plan_show_name",
				"subway",
				"bus_station",
				"cancel_policy",
				"address",
				"required_book",
				"coupon_name",
				"coupon_sub_name",
				"desc",
				"goods_name",
				"goods_content",
				"cuisine_name",
				"sub_name",
				"content",
				"detail_address",
				"check_in_guide",
				"store_name",
				"tag",
				"title",
				"sub_title",
			}
			UseLanguageFields = make(map[string][]string)
			LanguageInfo      []*entity.PmsLanguage
			language          = contexts.GetLanguage(ctx)
		)
		g.Log().Debugf(ctx, "--[PmsFindLanguageValueHook]-----------------%s----------------", language)
		if g.IsEmpty(language) {
			language = g.Cfg().MustGet(ctx, "server.defaultLanguage").String()
			g.Log().Debugf(ctx, "--[PmsFindLanguageValueHook]-----------------%s----------------", language)
		}
		g.Log().Debugf(ctx, "--[PmsFindLanguageValueHook]-----------------%s----------------", language)
		for _, v := range LanguageFields {
			if strings.Contains(in.Table, dao.CarAddress.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.CarDriver.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.FoodRestaurant.Table()) && v == "detail_address" {
				continue
			}
			if strings.Contains(in.Table, dao.ThMchStore.Table()) && v == "detail_address" {
				continue
			}
			if strings.Contains(in.Table, dao.SpaTechnician.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.SpaIsp.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.ThMch.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.ThCouponMch.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.SpaStore.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.TravelVerifyStaff.Table()) && v == "name" {
				continue
			}
			if strings.Contains(in.Table, dao.HomepageArticles.Table()) && v == "title" {
				continue
			}

			if strings.Trim(strings.ReplaceAll(in.Table, "`", ""), " ") == dao.Employee.Table() && v == "name" {
				continue
			}
			UseLanguageFields[v] = []string{}
		}

		for _, record := range result {
			for k1 := range UseLanguageFields {
				value, ok := record[k1]
				if ok {
					UseLanguageFields[k1] = append(UseLanguageFields[k1], value.String())
				}
			}
		}
		for k, v := range UseLanguageFields {
			if !g.IsEmpty(v) {
				if err = dao.PmsLanguage.Ctx(ctx).
					Where(dao.PmsLanguage.Columns().Language, language).
					WhereIn(dao.PmsLanguage.Columns().Uuid, v).
					Scan(&LanguageInfo); err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						g.Log().Error(ctx, err)
					}
					err = nil
				}
				for index, record := range result {
					value := record[k]
					result[index][k] = nil
					for _, v := range LanguageInfo {
						if value.String() == v.Uuid {
							result[index][k] = gvar.New(v.Content)
							break
						}
					}
				}
			}
		}
		return
	},
}
