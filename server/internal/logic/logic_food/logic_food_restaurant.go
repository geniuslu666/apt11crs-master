package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/library/toretaApi"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/charset"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

type sFoodRestaurant struct{}

func NewFoodRestaurant() *sFoodRestaurant {
	return &sFoodRestaurant{}
}

func init() {
	service.RegisterFoodRestaurant(NewFoodRestaurant())
}

func (s *sFoodRestaurant) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodRestaurant.Ctx(ctx), option...)
}

func (s *sFoodRestaurant) List(ctx context.Context, in *input_food.FoodRestaurantListInp) (list []*input_food.FoodRestaurantListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_food.FoodRestaurantListModel{})

	if !g.IsEmpty(in.Name) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.Name, "name")
		if err == nil {
			restaurantIds, _ := service.FoodRestaurant().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.FoodRestaurant.Columns().Id, restaurantIds)
		}
	}

	if !g.IsEmpty(in.OpenStatus) {
		mod = mod.Where(dao.FoodRestaurant.Columns().OpenStatus, in.OpenStatus)
	}

	if !g.IsEmpty(in.CanOrder) {
		mod = mod.Where(dao.FoodRestaurant.Columns().CanOrder, in.CanOrder)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	if !g.IsEmpty(in.Sort) {
		if in.Sort == "ascend" {
			mod = mod.Order(dao.FoodRestaurant.Columns().Sort, "asc")
		} else {
			mod = mod.Order(dao.FoodRestaurant.Columns().Sort, "desc")
		}
	} else {
		mod = mod.Order(dao.FoodRestaurant.Columns().Sort, "desc").OrderDesc(dao.FoodRestaurant.Columns().Id)
	}

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取餐厅列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodRestaurant) All(ctx context.Context, in *input_food.FoodRestaurantListInp) (list []*input_food.FoodRestaurantAllListModel, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_food.FoodRestaurantAllListModel{})

	mod = mod.Where(dao.FoodRestaurant.Columns().OpenStatus, "OPENING")

	if !g.IsEmpty(in.RestaurantIds) {
		mod = mod.WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(in.RestaurantIds, ","))
	}

	mod = mod.OrderDesc(dao.FoodRestaurant.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取餐厅列表失败，请稍后重试！")
		return
	}

	return
}

func (s *sFoodRestaurant) Edit(ctx context.Context, in *input_food.FoodRestaurantEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			restaurantAccountCount   int
			thAccountCount           int
			Object                   gdb.Record
			FoodRestaurantDao        *input_language.LoadLanguage
			FoodRestaurantContentDao *input_language.LoadLanguage
			LanguageStruct           input_language.LanguageModel
			Area                     entity.FoodArea
			RestaurantCuisineInfo    []*entity.FoodRestaurantCuisine
			RestaurantLabelInfo      []*entity.FoodRestaurantLabel
			RestaurantTerminalInfo   []*entity.FoodRestaurantTerminal
		)
		Uuid := guid.S([]byte("name"))
		ContentUuid := guid.S([]byte("content"))

		if in.AreaId > 0 {
			if err = dao.FoodArea.Ctx(ctx).Where(dao.FoodArea.Columns().Id, in.AreaId).Scan(&Area); err != nil {
				return
			}
			in.AreaPid = int(Area.Pid)
		}

		if in.Id > 0 {
			// 编辑
			if Object, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, in.Id).One(); err != nil {
				return
			}

			// 验证登录名
			if !g.IsEmpty(in.Account) {
				restaurantAccountCount, _ = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Account, in.Account).WhereNot(dao.FoodRestaurant.Columns().Id, in.Id).Count()
				thAccountCount, _ = dao.ThMchStore.Ctx(ctx).Where(dao.ThMchStore.Columns().Account, in.Account).Count()
				if restaurantAccountCount > 0 || thAccountCount > 0 {
					err = gerror.New("账号已存在")
					return
				}
			}

			var (
				RestaurantSeatInfo []*entity.FoodRestaurantSeat
				SeatTotalNum       int
			)
			for _, Seat := range in.Seat {
				if !g.IsEmpty(Seat.SeatId) {
					RestaurantSeatInfo = append(RestaurantSeatInfo, &entity.FoodRestaurantSeat{
						RestaurantId: in.Id,
						SeatId:       int64(Seat.SeatId),
						Num:          Seat.Num,
					})
					SeatTotalNum = SeatTotalNum + Seat.Num
				}
			}
			if !g.IsEmpty(RestaurantSeatInfo) {
				if _, err = dao.FoodRestaurantSeat.Ctx(ctx).Where(dao.FoodRestaurantSeat.Columns().RestaurantId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理餐厅坐席旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.FoodRestaurantSeat.Ctx(ctx).OmitEmptyData().Insert(RestaurantSeatInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.FoodRestaurantSeat.Ctx(ctx).Where(dao.FoodRestaurantSeat.Columns().RestaurantId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理餐厅坐席旧数据失败，请稍后重试！")
					return
				}
			}

			if !g.IsEmpty(in.CuisineIds) {
				for _, Cuisine := range strings.Split(in.CuisineIds, ",") {
					if !g.IsEmpty(Cuisine) {
						CuisineId, _ := strconv.ParseInt(Cuisine, 10, 64)
						RestaurantCuisineInfo = append(RestaurantCuisineInfo, &entity.FoodRestaurantCuisine{
							RestaurantId: in.Id,
							CuisineId:    CuisineId,
						})
					}
				}
			}

			if !g.IsEmpty(in.LabelIds) {
				for _, Label := range strings.Split(in.LabelIds, ",") {
					if !g.IsEmpty(Label) {
						LabelId, _ := strconv.ParseInt(Label, 10, 64)
						RestaurantLabelInfo = append(RestaurantLabelInfo, &entity.FoodRestaurantLabel{
							RestaurantId: in.Id,
							LabelId:      LabelId,
						})
					}
				}
			}

			mod := s.Model(ctx)

			if in.Type == "basic" {
				var isToreta = false
				if in.CooperateTypeId == 4 {
					// 判断toretaId是否已存在, 如果已存在，则不允许添加
					if !g.IsEmpty(in.ToretaId) {
						var count int
						if count, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().ToretaId, in.ToretaId).WhereNot(dao.FoodRestaurant.Columns().Id, in.Id).Count(); err != nil {
							return
						}
						if count > 0 {
							err = gerror.New("该ToretaId已存在")
							return
						}

						isToreta = true
						// 则请求toreta餐厅接口
						var (
							toretaResponse           *toretaApi.RestaurantDetailResponse
							ToretaApiConfig          *model.ToretaApiConfig
							restaurantsDetailRequest *toretaApi.RestaurantDetailParams
						)
						if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
							return
						}

						restaurantsDetailRequest = new(toretaApi.RestaurantDetailParams)
						restaurantsDetailRequest.Key = in.ToretaId
						if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetRestaurantDetail(ctx, restaurantsDetailRequest); err != nil {
							return
						}
						// if !toretaResponse.Success {
						// 	err = gerror.New(toretaResponse.Message)
						// 	return
						// }
						if !g.IsEmpty(toretaResponse) {
							// 预定模式-全款模式
							in.OrderMode = "ALLAMOUNT"
							// toreta订单默认自动确认模式
							in.OrderConfirmType = 2
							// 预约时间 1-每天 2-自定义
							in.OpenTimeType = 1
							// 最大容纳预定数 1-按时段 2-全局
							in.OrderMaxType = 2
							// 时段/全局最大容纳预定数量
							in.TimeDurationMax = toretaResponse.RestaurantSeatsMax
							// 当天预约处理截止时间
							duration := time.Duration(toretaResponse.AcceptLimitHour) * time.Second
							hours := int(duration.Hours())
							minutes := int(duration.Minutes()) % 60
							in.DayTimeLimit = fmt.Sprintf("%02d:%02d", hours, minutes)
							// 提前预约天数
							in.AdvanceOrderDay = toretaResponse.AcceptLimitDay
							// 最长预约天数
							in.MaxOrderDay = toretaResponse.AllowedReservationDays
							// Toreta是否允许取消 1 允许  2 不允许
							in.ToretaCancelEnable = 2
							if toretaResponse.CancelEnabled {
								in.ToretaCancelEnable = 1
							}
							// Toreta允许取消几天前
							in.ToretaCancelLimitDay = toretaResponse.CancelLimitDay
							// Toreta允许取消时间前
							cancelLimitDuration := time.Duration(toretaResponse.CancelLimitHour) * time.Second
							cancelLimitHours := int(cancelLimitDuration.Hours())
							cancelLimitMinutes := int(cancelLimitDuration.Minutes()) % 60
							in.ToretaCancelLimitTime = fmt.Sprintf("%02d:%02d", cancelLimitHours, cancelLimitMinutes)
						} else {
							err = gerror.New("获取toreta餐厅信息失败，请稍后重试！")
							return
						}
					}
				} else {
					in.ToretaId = ""
				}

				if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
					Uuid = Object["name"].String()
				}
				if !g.IsEmpty(Object[gstr.ToLower("Content")]) {
					ContentUuid = Object["content"].String()
				}

				FoodRestaurantDao = &input_language.LoadLanguage{
					Uuid: Uuid,
					Tag:  dao.FoodRestaurant.Table(),
					Type: "table",
					Key:  gstr.CaseSnakeFirstUpper("Name"),
				}
				LanguageStruct = in.NameLanguage
				if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodRestaurantDao); err != nil {
					return
				}

				FoodRestaurantContentDao = &input_language.LoadLanguage{
					Uuid: ContentUuid,
					Tag:  dao.FoodRestaurant.Table(),
					Type: "table",
					Key:  gstr.CaseSnakeFirstUpper("Content"),
				}
				LanguageStruct = in.ContentLanguage
				if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodRestaurantContentDao); err != nil {
					return
				}

				in.Name = Uuid
				in.Content = ContentUuid

				if !g.IsEmpty(RestaurantCuisineInfo) {
					if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).Where(dao.FoodRestaurantCuisine.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅菜系旧数据失败，请稍后重试！")
						return
					}
					if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).OmitEmptyData().Insert(RestaurantCuisineInfo); err != nil {
						return err
					}
				} else {
					if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).Where(dao.FoodRestaurantCuisine.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅菜系旧数据失败，请稍后重试！")
						return
					}
				}

				if !g.IsEmpty(RestaurantLabelInfo) {
					if _, err = dao.FoodRestaurantLabel.Ctx(ctx).Where(dao.FoodRestaurantLabel.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅标签旧数据失败，请稍后重试！")
						return
					}
					if _, err = dao.FoodRestaurantLabel.Ctx(ctx).OmitEmptyData().Insert(RestaurantLabelInfo); err != nil {
						return err
					}
				} else {
					if _, err = dao.FoodRestaurantLabel.Ctx(ctx).Where(dao.FoodRestaurantLabel.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅标签旧数据失败，请稍后重试！")
						return
					}
				}

				if isToreta {
					mod = mod.Fields(input_food.FoodRestaurantBasicUpdateToretaFields{})
				} else {
					mod = mod.Fields(input_food.FoodRestaurantBasicUpdateFields{})
				}
			} else if in.Type == "operation" {
				if Object[gstr.ToLower("cooperate_type_id")].Int() != 4 {
					if !g.IsEmpty(in.OpenTimeType != 1) {
						in.DayTimeLimit = ""
					}
					mod = mod.Fields(input_food.FoodRestaurantOperationUpdateFields{})
				} else {
					// Torate模式 预定模式就是全款模式
					in.OrderMode = "ALLAMOUNT"
					mod = mod.Fields(input_food.FoodRestaurantToretaOperationUpdateFields{})
				}
			} else if in.Type == "other" {

				if Object["time_duration_max"].Int() > 0 && SeatTotalNum > Object["time_duration_max"].Int() {
					err = gerror.New("可预约坐席数大于最大可容纳预定数据量！")
					return
				}

				mod = mod.Fields(input_food.FoodRestaurantOtherUpdateFields{})
			} else if in.Type == "settle" {
				mod = mod.Fields(input_food.FoodRestaurantSettleUpdateFields{})
			} else if in.Type == "terminal" {
				// 更新终端表的restaurantId
				if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().RestaurantId, in.Id).Update(g.MapStrAny{
					dao.SysTerminal.Columns().RestaurantId: 0,
				}); err != nil {
					err = gerror.Wrap(err, "清理终端旧数据失败，请稍后重试！")
					return
				}

				if !g.IsEmpty(in.TerminalList) {
					for _, TerminalInfo := range in.TerminalList {
						if !g.IsEmpty(TerminalInfo) {
							TerminalId := TerminalInfo.Id
							RestaurantTerminalInfo = append(RestaurantTerminalInfo, &entity.FoodRestaurantTerminal{
								RestaurantId: in.Id,
								TerminalId:   int64(TerminalId),
								PrintTimes:   uint(TerminalInfo.PrintTimes),
							})

							if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().Id, TerminalId).Update(g.MapStrAny{
								dao.SysTerminal.Columns().RestaurantId: in.Id,
							}); err != nil {
								err = gerror.Wrap(err, "更新终端信息失败，请稍后重试！")
								return
							}
						}
					}
				}
				if !g.IsEmpty(RestaurantTerminalInfo) {
					if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅终端旧数据失败，请稍后重试！")
						return
					}
					if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).OmitEmptyData().Insert(RestaurantTerminalInfo); err != nil {
						return err
					}
				} else {
					if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, in.Id).Delete(); err != nil {
						err = gerror.Wrap(err, "清理餐厅终端旧数据失败，请稍后重试！")
						return
					}
				}

				mod = mod.Fields(input_food.FoodRestaurantTerminalUpdateFields{})
			} else {
				mod = mod.Fields(input_food.FoodRestaurantUpdateFields{})
			}

			if in.Password != "" {
				if g.IsEmpty(Object[gstr.ToLower("Salt")]) {
					in.Salt = grand.S(6)
				} else {
					in.Salt = Object[gstr.ToLower("Salt")].String()
				}
				// 修改密码，需要获取到密码盐
				in.PasswordHash = gmd5.MustEncryptString(in.Password + in.Salt)
			} else {
				in.Salt = Object[gstr.ToLower("Salt")].String()
				in.PasswordHash = Object["password_hash"].String()
			}

			if _, err = mod.
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改餐厅管理失败，请稍后重试！")
			}

			return
		}

		var (
			lastInsertId int64
			//verifyCode   string
			//isOnly       = false
			//restaurant   entity.FoodRestaurant
		)

		//verifyCode = string(charset.RandomCreateText(6, "0123456789"))

		//for !isOnly {
		//	err = s.Model(ctx).
		//		Where(dao.FoodRestaurant.Columns().VerifyCode, verifyCode).
		//		Where("id <> ?", in.Id).
		//		Scan(&restaurant)
		//	if g.IsEmpty(restaurant) {
		//		isOnly = true
		//	}
		//}
		//in.VerifyCode = verifyCode
		if !g.IsEmpty(in.Account) {
			// 验证登录名
			restaurantAccountCount, _ = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Account, in.Account).Count()
			thAccountCount, _ = dao.ThMchStore.Ctx(ctx).Where(dao.ThMchStore.Columns().Account, in.Account).Count()
			if restaurantAccountCount > 0 || thAccountCount > 0 {
				err = gerror.New("账号已存在")
				return
			}
		}

		in.Name = Uuid
		in.Content = ContentUuid

		in.Salt = grand.S(6)
		in.PasswordHash = gmd5.MustEncryptString(in.Password + in.Salt)

		// 如果合作类型为4-toreta，那么需要请求toreta餐厅信息
		if in.CooperateTypeId == 4 {

			// 判断toretaId是否已存在, 如果已存在，则不允许添加
			if !g.IsEmpty(in.ToretaId) {
				var count int
				if count, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().ToretaId, in.ToretaId).Count(); err != nil {
					return
				}
				if count > 0 {
					err = gerror.New("该ToretaId已存在")
					return
				}
			}

			var (
				toretaResponse           *toretaApi.RestaurantDetailResponse
				ToretaApiConfig          *model.ToretaApiConfig
				restaurantsDetailRequest *toretaApi.RestaurantDetailParams
			)
			if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
				return
			}

			restaurantsDetailRequest = new(toretaApi.RestaurantDetailParams)
			restaurantsDetailRequest.Key = in.ToretaId
			if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetRestaurantDetail(ctx, restaurantsDetailRequest); err != nil {
				return
			}
			// if !toretaResponse.Success {
			// 	err = gerror.New(toretaResponse.Message)
			// 	return
			// }
			if !g.IsEmpty(toretaResponse) {
				// 预定模式-全款模式
				in.OrderMode = "ALLAMOUNT"
				// toreta订单默认自动确认模式
				in.OrderConfirmType = 2
				// 预约时间 1-每天 2-自定义
				in.OpenTimeType = 1
				// 最大容纳预定数 1-按时段 2-全局
				in.OrderMaxType = 2
				// 时段/全局最大容纳预定数量
				in.TimeDurationMax = toretaResponse.RestaurantSeatsMax
				// 当天预约处理截止时间
				duration := time.Duration(toretaResponse.AcceptLimitHour) * time.Second
				hours := int(duration.Hours())
				minutes := int(duration.Minutes()) % 60
				in.DayTimeLimit = fmt.Sprintf("%02d:%02d", hours, minutes)
				// 提前预约天数
				in.AdvanceOrderDay = toretaResponse.AcceptLimitDay
				// 最长预约天数
				in.MaxOrderDay = toretaResponse.AllowedReservationDays
				// Toreta是否允许取消 1 允许  2 不允许
				in.ToretaCancelEnable = 2
				if toretaResponse.CancelEnabled {
					in.ToretaCancelEnable = 1
				}
				// Toreta允许取消几天前
				in.ToretaCancelLimitDay = toretaResponse.CancelLimitDay
				// Toreta允许取消时间前
				cancelLimitDuration := time.Duration(toretaResponse.CancelLimitHour) * time.Second
				cancelLimitHours := int(cancelLimitDuration.Hours())
				cancelLimitMinutes := int(cancelLimitDuration.Minutes()) % 60
				in.ToretaCancelLimitTime = fmt.Sprintf("%02d:%02d", cancelLimitHours, cancelLimitMinutes)
			} else {
				err = gerror.New("获取toreta餐厅信息失败，请稍后重试！")
				return
			}
		}
		if !g.IsEmpty(in.OpenTimeType != 1) {
			in.DayTimeLimit = ""
		}

		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodRestaurantInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增餐厅失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		// 更新终端表的restaurantId
		if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().RestaurantId, in.Id).Update(g.MapStrAny{
			dao.SysTerminal.Columns().RestaurantId: 0,
			dao.SysTerminal.Columns().PrintTimes:   1,
		}); err != nil {
			err = gerror.Wrap(err, "清理终端旧数据失败，请稍后重试！")
			return
		}

		if !g.IsEmpty(in.TerminalList) {
			for _, TerminalInfo := range in.TerminalList {
				if !g.IsEmpty(TerminalInfo) {
					TerminalId := TerminalInfo.Id
					RestaurantTerminalInfo = append(RestaurantTerminalInfo, &entity.FoodRestaurantTerminal{
						RestaurantId: lastInsertId,
						TerminalId:   int64(TerminalId),
						PrintTimes:   uint(TerminalInfo.PrintTimes),
					})

					if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().Id, TerminalId).Update(g.MapStrAny{
						dao.SysTerminal.Columns().RestaurantId: lastInsertId,
					}); err != nil {
						err = gerror.Wrap(err, "更新终端信息失败，请稍后重试！")
						return
					}
				}
			}
		}

		if !g.IsEmpty(RestaurantTerminalInfo) {
			if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅终端旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).OmitEmptyData().Insert(RestaurantTerminalInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.FoodRestaurantTerminal.Ctx(ctx).Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅终端旧数据失败，请稍后重试！")
				return
			}
		}

		if !g.IsEmpty(in.CuisineIds) {
			for _, Cuisine := range strings.Split(in.CuisineIds, ",") {
				if !g.IsEmpty(Cuisine) {
					CuisineId, _ := strconv.ParseInt(Cuisine, 10, 64)
					RestaurantCuisineInfo = append(RestaurantCuisineInfo, &entity.FoodRestaurantCuisine{
						RestaurantId: lastInsertId,
						CuisineId:    CuisineId,
					})
				}
			}
		}

		if !g.IsEmpty(in.LabelIds) {
			for _, Label := range strings.Split(in.LabelIds, ",") {
				if !g.IsEmpty(Label) {
					LabelId, _ := strconv.ParseInt(Label, 10, 64)
					RestaurantLabelInfo = append(RestaurantLabelInfo, &entity.FoodRestaurantLabel{
						RestaurantId: lastInsertId,
						LabelId:      LabelId,
					})
				}
			}
		}

		if !g.IsEmpty(RestaurantCuisineInfo) {
			if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).Where(dao.FoodRestaurantCuisine.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅菜系旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).OmitEmptyData().Insert(RestaurantCuisineInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.FoodRestaurantCuisine.Ctx(ctx).Where(dao.FoodRestaurantCuisine.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅菜系旧数据失败，请稍后重试！")
				return
			}
		}

		if !g.IsEmpty(RestaurantLabelInfo) {
			if _, err = dao.FoodRestaurantLabel.Ctx(ctx).Where(dao.FoodRestaurantLabel.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅标签旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.FoodRestaurantLabel.Ctx(ctx).OmitEmptyData().Insert(RestaurantLabelInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.FoodRestaurantLabel.Ctx(ctx).Where(dao.FoodRestaurantLabel.Columns().RestaurantId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理餐厅标签旧数据失败，请稍后重试！")
				return
			}
		}

		FoodRestaurantDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.FoodRestaurant.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodRestaurantDao); err != nil {
			return
		}

		FoodRestaurantContentDao = &input_language.LoadLanguage{
			Uuid: ContentUuid,
			Tag:  dao.FoodRestaurant.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Content"),
		}
		LanguageStruct = in.ContentLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, FoodRestaurantContentDao); err != nil {
			return
		}

		return
	})
}

func (s *sFoodRestaurant) Delete(ctx context.Context, in *input_food.FoodRestaurantDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除餐厅管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodRestaurant) View(ctx context.Context, in *input_food.FoodRestaurantViewInp) (res *input_food.FoodRestaurantViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取餐厅管理信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取餐厅管理信息，请稍后重试！")
			return
		}
	}
	return
}

func (s *sFoodRestaurant) Status(ctx context.Context, in *input_food.FoodRestaurantStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.FoodRestaurant.Columns().OpenStatus: in.Status,
			dao.FoodRestaurant.Columns().Desc:       in.Desc,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新餐厅状态失败，请稍后重试！")
			return
		}

		if in.Status == "OPENING" || in.Status == "CLOSE" {
			if _, err = dao.FoodRestaurantOperateLog.Ctx(ctx).
				Data(g.Map{
					dao.FoodRestaurantOperateLog.Columns().RestaurantId:    in.Id,
					dao.FoodRestaurantOperateLog.Columns().OperateType:     in.Status,
					dao.FoodRestaurantOperateLog.Columns().OperateMemberId: contexts.GetUserId(ctx),
					dao.FoodRestaurantOperateLog.Columns().Remark:          in.Desc,
				}).OmitEmptyData().Insert(); err != nil {
				err = gerror.Wrap(err, "餐厅日志表写入失败，请稍后重试！")
			}
		}

		return
	})
}

func (s *sFoodRestaurant) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}

func (s *sFoodRestaurant) ResetVerifyCode(ctx context.Context, in *input_food.FoodRestaurantRestVerifyCodeInp) (err error) {
	var (
		verifyCode string
		isOnly     = false
		restaurant entity.FoodRestaurant
	)

	verifyCode = string(charset.RandomCreateText(6, "0123456789"))

	for !isOnly {
		err = s.Model(ctx).
			Where(dao.FoodRestaurant.Columns().VerifyCode, verifyCode).
			Where("id <> ?", in.Id).
			Scan(&restaurant)
		if g.IsEmpty(restaurant) {
			isOnly = true
		}
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{"verify_code": verifyCode}).Update(); err != nil {
		err = gerror.Wrap(err, "重置核销码失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodRestaurant) Switch(ctx context.Context, in *input_food.FoodRestaurantSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.FoodRestaurant.Columns().CanOrder: in.CanOrder,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新预定状态失败，请稍后重试！")
			return
		}

		return
	})
}

// RestaurantSort 编辑餐厅排序
func (s *sFoodRestaurant) RestaurantSort(ctx context.Context, in *input_food.FoodRestaurantSortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Update(g.Map{
		dao.FoodRestaurant.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改餐厅排序，请稍后重试！")
		return
	}
	return
}
