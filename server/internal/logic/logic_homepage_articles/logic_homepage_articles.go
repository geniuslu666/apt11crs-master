package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_homepage_article"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sBasicsHomepageArticles struct{}

func NewBasicsHomepageArticles() *sBasicsHomepageArticles {
	return &sBasicsHomepageArticles{}
}

func init() {
	service.RegisterBasicsHomepageArticles(NewBasicsHomepageArticles())
}

func (s *sBasicsHomepageArticles) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.HomepageArticles.Ctx(ctx), option...)
}

func (s *sBasicsHomepageArticles) List(ctx context.Context, in *input_homepage_article.HomepageArticlesListInp) (list []*input_homepage_article.HomepageArticlesListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_homepage_article.HomepageArticlesListModel{})

	if !g.IsEmpty(in.Title) {
		mod = mod.WhereLike(dao.HomepageArticles.Columns().Title, "%"+in.Title+"%")
	}
	if !g.IsEmpty(in.Keyword) {
		mod = mod.Where("title LIKE %?% OR content LIKE %?% ", in.Keyword)
	}
	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.HomepageArticles.Columns().Status, in.Status)
	}
	if !g.IsEmpty(in.Language) {

		mod = mod.Where(dao.HomepageArticles.Columns().Language, in.Language)
	}
	/*else {

		mod = mod.Where(dao.HomepageArticles.Columns().Language, contexts.GetLanguage(ctx))
	}*/

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.HomepageArticles.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.HomepageArticles.Table() + "." + dao.HomepageArticles.Columns().Sort).OrderDesc(dao.HomepageArticles.Table() + "." + dao.HomepageArticles.Columns().Id)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取首页活动列表失败，请稍后重试！")
			return
		}
	} else {

		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取首页活动列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sBasicsHomepageArticles) Edit(ctx context.Context, in *input_homepage_article.HomepageArticlesEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			ArticleCouponInfo []*entity.HomepageArticleCoupon
			ArticleLinks      []*entity.HomepageArticleLinks
		)
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_homepage_article.HomepageArticlesUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改首页活动失败，请稍后重试！")
			}
			if !g.IsEmpty(in.ThCouponsArr) {
				for _, ThCoupon := range in.ThCouponsArr {
					if !g.IsEmpty(ThCoupon) {
						ThCouponId := ThCoupon.CouponId
						ArticleCouponInfo = append(ArticleCouponInfo, &entity.HomepageArticleCoupon{
							CouponType:        "thCoupon",
							ArticleId:         uint64(in.Id),
							CouponId:          uint64(ThCouponId),
							AvailableQuantity: ThCoupon.AvailableQuantity,
							PerDayAvailable:   ThCoupon.PerDayAvailable,
							LimitDays:         ThCoupon.LimitDays,
							FixedTerm:         ThCoupon.FixedTerm,
						})

						pushData := g.MapStrAny{
							"couponId":        ThCouponId,
							"indexActivityId": in.Id,
						}
						pushDataJson, _ := json.Marshal(pushData)
						appPushData := g.MapStrStr{
							"type":  "2",
							"param": string(pushDataJson),
						}

						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:  "thCoupon",
							LinkId:    ThCouponId,
							ArticleId: uint64(in.Id),
							AppLink:   "/home/homeGiftPage",
							WxLink:    fmt.Sprintf("/pages/gift-coupon/wait-details?id=%d&activityId=%d", ThCouponId, in.Id),
							UrlParam:  gjson.New(appPushData),
						})

					}
				}
			}
			if !g.IsEmpty(in.CouponsArr) {
				for _, Coupon := range in.CouponsArr {
					if !g.IsEmpty(Coupon) {
						CouponId := Coupon.CouponId
						couponLimitDays := Coupon.LimitDays
						if couponLimitDays <= 0 {
							couponLimitDays = 1
						}
						ArticleCouponInfo = append(ArticleCouponInfo, &entity.HomepageArticleCoupon{
							CouponType:        "coupon",
							ArticleId:         uint64(in.Id),
							CouponId:          uint64(CouponId),
							AvailableQuantity: Coupon.AvailableQuantity,
							PerDayAvailable:   Coupon.PerDayAvailable,
							LimitDays:         couponLimitDays,
						})

						pushData := g.MapStrAny{
							"couponId":        CouponId,
							"indexActivityId": in.Id,
						}
						pushDataJson, _ := json.Marshal(pushData)
						appPushData := g.MapStrStr{
							"type":  "2",
							"param": string(pushDataJson),
						}

						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:  "coupon",
							LinkId:    CouponId,
							ArticleId: uint64(in.Id),
							AppLink:   "/home/homeCouponPage",
							WxLink:    fmt.Sprintf("/pages/coupons/wait-details?id=%d&activityId=%d", CouponId, in.Id),
							UrlParam:  gjson.New(appPushData),
						})
					}
				}
			}
			if !g.IsEmpty(ArticleCouponInfo) {
				if _, err = dao.HomepageArticleCoupon.Ctx(ctx).TX(tx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理文章关联券旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.HomepageArticleCoupon.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ArticleCouponInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.HomepageArticleCoupon.Ctx(ctx).TX(tx).Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理文章关联券旧数据失败，请稍后重试！")
					return
				}
			}

			HotelIdsArr := strings.Split(in.HotelIds, ",")
			if !g.IsEmpty(HotelIdsArr) {
				for _, hotelId := range HotelIdsArr {
					if gvar.New(hotelId).Int() > 0 {
						var property *entity.PmsProperty
						// 获取酒店信息
						if err = dao.PmsProperty.Ctx(ctx).Where(dao.PmsProperty.Columns().Id, hotelId).Scan(&property); err != nil {
							err = gerror.Wrap(err, "酒店信息不存在！")
							return
						}

						pushData := g.MapStrAny{
							"id":  gvar.New(hotelId).Int(),
							"uid": property.Uid,
						}
						pushDataJson, _ := json.Marshal(pushData)
						appPushData := g.MapStrStr{
							"type":  "2",
							"param": string(pushDataJson),
						}

						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:  "hotelDetail",
							LinkId:    gvar.New(hotelId).Int(),
							ArticleId: uint64(in.Id),
							AppLink:   "/hotel/hotel-main",
							WxLink:    "/pages/hotels/details?id=" + hotelId + "&puid=" + property.Uid,
							UrlParam:  gjson.New(appPushData),
						})
					}
				}
			}

			RestaurantIdsArr := strings.Split(in.RestaurantIds, ",")
			if !g.IsEmpty(RestaurantIdsArr) {
				for _, restaurantId := range RestaurantIdsArr {
					if gvar.New(restaurantId).Int() > 0 {
						pushData := g.MapStrAny{
							"restaurantId": restaurantId,
							"thCoupon":     0,
						}
						pushDataJson, _ := json.Marshal(pushData)
						appPushData := g.MapStrStr{
							"type":  "2",
							"param": string(pushDataJson),
						}

						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:  "foodDetail",
							LinkId:    gvar.New(restaurantId).Int(),
							ArticleId: uint64(in.Id),
							AppLink:   "/dininRoomDetail",
							WxLink:    "/subpackages/restaurant-booking/pages/restaurant-details?id=" + restaurantId,
							UrlParam:  gjson.New(appPushData),
						})
					}
				}
			}

			SpaIdsArr := strings.Split(in.SpaServiceIds, ",")
			if !g.IsEmpty(SpaIdsArr) {
				for _, spaServiceId := range SpaIdsArr {
					if gvar.New(spaServiceId).Int() > 0 {
						appPushData := g.MapStrStr{
							"type":   "1",
							"string": fmt.Sprintf("%d", gvar.New(spaServiceId).Int()),
						}
						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:  "spaDetail",
							LinkId:    gvar.New(spaServiceId).Int(),
							ArticleId: uint64(in.Id),
							AppLink:   "/spaDetailPage",
							WxLink:    "/subpackages/spa/pages/details?id=" + spaServiceId,
							UrlParam:  gjson.New(appPushData),
						})
					}
				}
			}
			// 判断几个首页推荐是否勾选
			if strings.Contains(in.RecommendLinkType, "foodIndex") {
				appPushData := g.MapStrStr{
					"type": "0",
				}
				ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
					LinkType:  "foodIndex",
					LinkId:    0,
					ArticleId: uint64(in.Id),
					Title:     in.FoodIndexTitle,
					AppLink:   "/diningRoomList",
					WxLink:    "/subpackages/restaurant-booking/pages/home",
					UrlParam:  gjson.New(appPushData),
				})
			}
			if strings.Contains(in.RecommendLinkType, "spaIndex") {
				appPushData := g.MapStrStr{
					"type": "0",
				}
				ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
					LinkType:  "spaIndex",
					LinkId:    0,
					ArticleId: uint64(in.Id),
					Title:     in.SpaIndexTitle,
					AppLink:   "/spaHomePage",
					WxLink:    "/subpackages/spa/pages/home",
					UrlParam:  gjson.New(appPushData),
				})
			}
			if strings.Contains(in.RecommendLinkType, "carIndex") {
				appPushData := g.MapStrStr{
					"type": "0",
				}
				ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
					LinkType:  "carIndex",
					LinkId:    0,
					ArticleId: uint64(in.Id),
					Title:     in.CarIndexTitle,
					AppLink:   "/homePage",
					WxLink:    "/pages/car/index",
					UrlParam:  gjson.New(appPushData),
				})
			}
			// 处理多个外链
			if strings.Contains(in.RecommendLinkType, "outLink") && !g.IsEmpty(in.OutLinks) {
				for _, outLink := range in.OutLinks {
					if !g.IsEmpty(outLink) {
						// 外部是3 内部是4
						appPushData := g.MapStrStr{
							"type": "3",
						}
						if outLink.OpenType == 1 {
							appPushData = g.MapStrStr{
								"type": "4",
							}
						}

						ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
							LinkType:     "outLink",
							LinkId:       0,
							ArticleId:    uint64(in.Id),
							Title:        outLink.Title,
							AppLink:      outLink.Link,
							WxLink:       outLink.Link,
							UrlParam:     gjson.New(appPushData),
							LinkText:     outLink.ButtonTxt,
							LinkOpenType: outLink.OpenType,
						})
					}
				}
			}

			// 添加关联链接
			if _, err = dao.HomepageArticleLinks.Ctx(ctx).TX(tx).Where(dao.HomepageArticleLinks.Columns().ArticleId, in.Id).Delete(); err != nil {
				err = gerror.Wrap(err, "清理文章关联链接旧数据失败，请稍后重试！")
				return
			}
			if !g.IsEmpty(ArticleLinks) {
				if _, err = dao.HomepageArticleLinks.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ArticleLinks); err != nil {
					return err
				}
			}

			return
		}

		// 新增
		// 设置结束时间为设定时间的23:59:59
		in.EndTime = gtime.New(in.EndTime.Format("Y-m-d") + " 23:59:59")
		var lastInsertId int64
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_homepage_article.HomepageArticlesInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增首页活动失败，请稍后重试！")
		}
		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		// 添加关联券
		if !g.IsEmpty(in.ThCouponsArr) {
			for _, ThCoupon := range in.ThCouponsArr {
				if !g.IsEmpty(ThCoupon) {
					ThCouponId := ThCoupon.CouponId
					ArticleCouponInfo = append(ArticleCouponInfo, &entity.HomepageArticleCoupon{
						CouponType:        "thCoupon",
						ArticleId:         uint64(lastInsertId),
						CouponId:          uint64(ThCouponId),
						AvailableQuantity: ThCoupon.AvailableQuantity,
						PerDayAvailable:   ThCoupon.PerDayAvailable,
						LimitDays:         ThCoupon.LimitDays,
						FixedTerm:         ThCoupon.FixedTerm,
					})

					// 关联链接
					pushData := g.MapStrAny{
						"couponId":        ThCouponId,
						"indexActivityId": lastInsertId,
					}
					pushDataJson, _ := json.Marshal(pushData)
					appPushData := g.MapStrStr{
						"type":  "2",
						"param": string(pushDataJson),
					}

					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:  "thCoupon",
						LinkId:    ThCouponId,
						ArticleId: uint64(lastInsertId),
						AppLink:   "/home/homeGiftPage",
						WxLink:    fmt.Sprintf("/pages/gift-coupon/wait-details?id=%d&activityId=%d", ThCouponId, in.Id),
						UrlParam:  gjson.New(appPushData),
					})
				}
			}
		}
		if !g.IsEmpty(in.CouponsArr) {
			for _, Coupon := range in.CouponsArr {
				if !g.IsEmpty(Coupon) {
					CouponId := Coupon.CouponId
					couponLimitDays := Coupon.LimitDays
					if couponLimitDays <= 0 {
						couponLimitDays = 1
					}
					ArticleCouponInfo = append(ArticleCouponInfo, &entity.HomepageArticleCoupon{
						CouponType:        "coupon",
						ArticleId:         uint64(lastInsertId),
						CouponId:          uint64(CouponId),
						AvailableQuantity: Coupon.AvailableQuantity,
						PerDayAvailable:   Coupon.PerDayAvailable,
						LimitDays:         couponLimitDays,
					})

					// 关联链接
					pushData := g.MapStrAny{
						"couponId":        CouponId,
						"indexActivityId": lastInsertId,
					}
					pushDataJson, _ := json.Marshal(pushData)
					appPushData := g.MapStrStr{
						"type":  "2",
						"param": string(pushDataJson),
					}

					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:  "coupon",
						LinkId:    CouponId,
						ArticleId: uint64(lastInsertId),
						AppLink:   "/home/homeCouponPage",
						WxLink:    fmt.Sprintf("/pages/coupons/wait-details?id=%d&activityId=%d", CouponId, lastInsertId),
						UrlParam:  gjson.New(appPushData),
					})
				}
			}
		}
		if !g.IsEmpty(ArticleCouponInfo) {
			// 添加关联券
			if _, err = dao.HomepageArticleCoupon.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ArticleCouponInfo); err != nil {
				return err
			}
		}

		HotelIdsArr := strings.Split(in.HotelIds, ",")
		if !g.IsEmpty(HotelIdsArr) {
			for _, hotelId := range HotelIdsArr {
				if gvar.New(hotelId).Int() > 0 {
					var property *entity.PmsProperty
					// 获取酒店信息
					if err = dao.PmsProperty.Ctx(ctx).Where(dao.PmsProperty.Columns().Id, hotelId).Scan(&property); err != nil {
						err = gerror.Wrap(err, "酒店信息不存在！")
						return
					}

					pushData := g.MapStrAny{
						"id":  gvar.New(hotelId).Int(),
						"uid": property.Uid,
					}
					pushDataJson, _ := json.Marshal(pushData)
					appPushData := g.MapStrStr{
						"type":  "2",
						"param": string(pushDataJson),
					}

					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:  "hotelDetail",
						LinkId:    gvar.New(hotelId).Int(),
						ArticleId: uint64(lastInsertId),
						AppLink:   "/hotel/hotel-main",
						WxLink:    "/pages/hotels/details?id=" + hotelId + "&puid=" + property.Uid,
						UrlParam:  gjson.New(appPushData),
					})
				}
			}
		}

		RestaurantIdsArr := strings.Split(in.RestaurantIds, ",")
		if !g.IsEmpty(RestaurantIdsArr) {
			for _, restaurantId := range RestaurantIdsArr {
				if gvar.New(restaurantId).Int() > 0 {

					pushData := g.MapStrAny{
						"restaurantId": restaurantId,
						"thCoupon":     0,
					}
					pushDataJson, _ := json.Marshal(pushData)
					appPushData := g.MapStrStr{
						"type":  "2",
						"param": string(pushDataJson),
					}

					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:  "foodDetail",
						LinkId:    gvar.New(restaurantId).Int(),
						ArticleId: uint64(lastInsertId),
						AppLink:   "/dininRoomDetail",
						WxLink:    "/subpackages/restaurant-booking/pages/restaurant-details?id=" + restaurantId,
						UrlParam:  gjson.New(appPushData),
					})
				}
			}
		}

		SpaIdsArr := strings.Split(in.SpaServiceIds, ",")
		if !g.IsEmpty(SpaIdsArr) {
			for _, spaServiceId := range SpaIdsArr {
				if gvar.New(spaServiceId).Int() > 0 {

					appPushData := g.MapStrStr{
						"type":   "1",
						"string": fmt.Sprintf("%d", gvar.New(spaServiceId).Int()),
					}
					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:  "spaDetail",
						LinkId:    gvar.New(spaServiceId).Int(),
						ArticleId: uint64(lastInsertId),
						AppLink:   "/spaDetailPage",
						WxLink:    "/subpackages/spa/pages/details?id=" + spaServiceId,
						UrlParam:  gjson.New(appPushData),
					})
				}
			}
		}

		// 判断几个首页推荐是否勾选
		if strings.Contains(in.RecommendLinkType, "foodIndex") {
			appPushData := g.MapStrStr{
				"type": "0",
			}
			ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
				LinkType:  "foodIndex",
				LinkId:    0,
				ArticleId: uint64(lastInsertId),
				Title:     in.FoodIndexTitle,
				AppLink:   "/diningRoomList",
				WxLink:    "/subpackages/restaurant-booking/pages/home",
				UrlParam:  gjson.New(appPushData),
			})
		}
		if strings.Contains(in.RecommendLinkType, "spaIndex") {
			appPushData := g.MapStrStr{
				"type": "0",
			}
			ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
				LinkType:  "spaIndex",
				LinkId:    0,
				ArticleId: uint64(lastInsertId),
				Title:     in.SpaIndexTitle,
				AppLink:   "/spaHomePage",
				WxLink:    "/subpackages/spa/pages/home",
				UrlParam:  gjson.New(appPushData),
			})
		}
		if strings.Contains(in.RecommendLinkType, "carIndex") {
			appPushData := g.MapStrStr{
				"type": "0",
			}
			ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
				LinkType:  "carIndex",
				LinkId:    0,
				ArticleId: uint64(lastInsertId),
				Title:     in.CarIndexTitle,
				AppLink:   "/homePage",
				WxLink:    "/pages/car/index",
				UrlParam:  gjson.New(appPushData),
			})
		}

		// 处理多个外链
		if strings.Contains(in.RecommendLinkType, "outLink") && !g.IsEmpty(in.OutLinks) {
			for _, outLink := range in.OutLinks {
				if !g.IsEmpty(outLink) {
					// 外部是3 内部是4
					appPushData := g.MapStrStr{
						"type": "3",
					}
					if outLink.OpenType == 1 {
						appPushData = g.MapStrStr{
							"type": "4",
						}
					}
					ArticleLinks = append(ArticleLinks, &entity.HomepageArticleLinks{
						LinkType:     "outLink",
						LinkId:       0,
						ArticleId:    uint64(lastInsertId),
						Title:        outLink.Title,
						AppLink:      outLink.Link,
						WxLink:       outLink.Link,
						UrlParam:     gjson.New(appPushData),
						LinkText:     outLink.ButtonTxt,
						LinkOpenType: outLink.OpenType,
					})
				}
			}
		}

		// 添加关联链接
		if !g.IsEmpty(ArticleLinks) {
			if _, err = dao.HomepageArticleLinks.Ctx(ctx).TX(tx).OmitEmptyData().Insert(ArticleLinks); err != nil {
				return err
			}
		}

		return
	})
}

func (s *sBasicsHomepageArticles) Delete(ctx context.Context, in *input_homepage_article.HomepageArticlesDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除首页活动失败，请稍后重试！")
		return
	}
	return
}

// Status 更新活动状态
func (s *sBasicsHomepageArticles) Status(ctx context.Context, in *input_homepage_article.HomepageArticlesStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.HomepageArticles.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新首页活动状态失败，请稍后重试！")
			return
		}

		return
	})
}

func (s *sBasicsHomepageArticles) View(ctx context.Context, in *input_homepage_article.HomepageArticlesViewInp) (res *input_homepage_article.HomepageArticlesViewModel, err error) {
	mod := s.Model(ctx).WithAll()

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取首页活动信息失败，请稍后重试！")
		return
	}

	// 查询外链数据
	var outLinks []*entity.HomepageArticleLinks
	if err = dao.HomepageArticleLinks.Ctx(ctx).
		Where(dao.HomepageArticleLinks.Columns().ArticleId, in.Id).
		Where(dao.HomepageArticleLinks.Columns().LinkType, "outLink").
		Scan(&outLinks); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取外链信息失败，请稍后重试！")
		return
	}
	// 重置err，避免sql.ErrNoRows影响后续流程
	err = nil

	// 组装外链数据
	if len(outLinks) > 0 {
		res.OutLinks = make([]*struct {
			Title     string `json:"title"      dc:"外链标题"`
			Link      string `json:"link"       dc:"外链链接"`
			ButtonTxt string `json:"buttonTxt"  dc:"按钮文字"`
			OpenType  int    `json:"openType"   dc:"跳转方式 1-内部webview 2-外部浏览器"`
		}, 0, len(outLinks))

		for _, link := range outLinks {
			res.OutLinks = append(res.OutLinks, &struct {
				Title     string `json:"title"      dc:"外链标题"`
				Link      string `json:"link"       dc:"外链链接"`
				ButtonTxt string `json:"buttonTxt"  dc:"按钮文字"`
				OpenType  int    `json:"openType"   dc:"跳转方式 1-内部webview 2-外部浏览器"`
			}{
				Title:     link.Title,
				Link:      link.AppLink,
				ButtonTxt: link.LinkText,
				OpenType:  link.LinkOpenType,
			})
		}
	}

	return
}

func (s *sBasicsHomepageArticles) AppView(ctx context.Context, in *input_homepage_article.HomepageArticlesAppViewInp) (res *input_homepage_article.HomepageArticlesAppViewModel, err error) {
	mod := s.Model(ctx)

	var AppUrlParamJson *input_basics.AppUrlParamModel

	// 先查询到临时结构体，包含时间字段
	var detail struct {
		Id                    int         `json:"id"`
		Title                 string      `json:"title"`
		Content               string      `json:"content"`
		Views                 int         `json:"views"`
		ThumbNum              int         `json:"thumbNum"`
		StartTime             *gtime.Time `json:"startTime"`
		EndTime               *gtime.Time `json:"endTime"`
		CouponLimit           int         `json:"couponLimit"`
		Status                int         `json:"status"`
		CouponLinkTitle       string      `json:"couponLinkTitle"`
		CouponLinkSubTitle    string      `json:"couponLinkSubTitle"`
		RecommendLinkTitle    string      `json:"recommendLinkTitle"`
		RecommendLinkSubTitle string      `json:"recommendLinkSubTitle"`
		LimitWeek             string      `json:"limitWeek"`
	}

	if err = mod.WherePri(in.Id).Scan(&detail); err != nil {
		err = gerror.Wrap(err, "获取首页活动信息失败，请稍后重试！")
		return
	}

	if detail.Status != 1 {
		err = gerror.New("活动已关闭！")
		return
	}

	// 增加浏览量
	if _, err = s.Model(ctx).WherePri(in.Id).Increment(dao.HomepageArticles.Columns().Views, 1); err != nil {
		err = gerror.Wrap(err, "更新浏览量失败")
		return
	}

	// 构建返回结果
	res = new(input_homepage_article.HomepageArticlesAppViewModel)
	res.Id = detail.Id
	res.Title = detail.Title
	res.Content = detail.Content
	res.Views = detail.Views + 1 // 加上刚才增加的浏览量
	res.ThumbNum = detail.ThumbNum
	res.CouponLimit = detail.CouponLimit
	res.CouponLinkTitle = detail.CouponLinkTitle
	res.CouponLinkSubTitle = detail.CouponLinkSubTitle
	res.RecommendLinkTitle = detail.RecommendLinkTitle
	res.RecommendLinkSubTitle = detail.RecommendLinkSubTitle

	// 活动状态用时间判断
	now := gtime.Now()
	if now.Before(detail.StartTime) {
		res.ActivityStatus = 1 // 未开始
	} else if now.After(detail.EndTime) {
		res.ActivityStatus = 3 // 已结束
	} else {
		res.ActivityStatus = 2 // 进行中
	}

	// 格式化时间字段
	if detail.StartTime != nil {
		res.StartTime = detail.StartTime.Format("Y-m-d")
	}
	if detail.EndTime != nil {
		res.EndTime = detail.EndTime.Format("Y-m-d")
	}

	// 判断当前用户是否已点赞
	memberInfo := contexts.GetMemberUser(ctx)
	if memberInfo != nil {
		var thumbRecord *entity.HomepageArticleThumb
		if err = dao.HomepageArticleThumb.Ctx(ctx).
			Where(dao.HomepageArticleThumb.Columns().MemberId, memberInfo.Id).
			Where(dao.HomepageArticleThumb.Columns().ArticleId, in.Id).
			Scan(&thumbRecord); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "查询点赞记录失败")
			return
		}
		res.HasThumb = thumbRecord != nil
	}

	// 查询文章关联的链接
	var links []*entity.HomepageArticleLinks
	if err = dao.HomepageArticleLinks.Ctx(ctx).Where(dao.HomepageArticleLinks.Columns().ArticleId, in.Id).Scan(&links); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取文章链接失败")
		return
	}

	// 收集所有优惠券ID用于批量查询用户领取记录
	var couponIds []int64
	var couponLinkTypes []string
	for _, v := range links {
		if v.LinkType == "coupon" || v.LinkType == "thCoupon" {
			couponIds = append(couponIds, int64(v.LinkId))
			couponLinkTypes = append(couponLinkTypes, v.LinkType)
		}
	}

	// 批量查询用户已领取的优惠券数量
	type couponReceivedMapItem struct {
		memberReceived      int
		memberTodayReceived int
		totalNum            int
		totalReceived       int
		availableQuantity   int
		perDayAvailable     int
		limitDays           int
	}
	// 提前查询文章关联的优惠券配置信息，以便查询用户领取数量时使用 limitDays
	limitDaysMap := make(map[string]int) // key: linkType-couponId, value: limitDays
	if len(couponIds) > 0 {
		var articleCouponsEarly []*entity.HomepageArticleCoupon
		if err = dao.HomepageArticleCoupon.Ctx(ctx).
			Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.Id).
			Scan(&articleCouponsEarly); err == nil {
			for _, c := range articleCouponsEarly {
				limitDaysMap[c.CouponType+"-"+gconv.String(c.CouponId)] = c.LimitDays
			}
		}
		err = nil
	}

	// 查询当前活动的编号，找出所有相同编号的活动ID（防止用户切换同一活动的不同语言版本后重复领券）
	var sameNoActivityIds []int64
	var currentArticleNo string
	_ = dao.HomepageArticles.Ctx(ctx).
		Fields(dao.HomepageArticles.Columns().No).
		WherePri(in.Id).
		Scan(&currentArticleNo)
	if currentArticleNo != "" {
		var sameNoRows []struct{ Id int64 }
		_ = dao.HomepageArticles.Ctx(ctx).
			Fields(dao.HomepageArticles.Columns().Id).
			Where(dao.HomepageArticles.Columns().No, currentArticleNo).
			Scan(&sameNoRows)
		for _, row := range sameNoRows {
			sameNoActivityIds = append(sameNoActivityIds, row.Id)
		}
	}
	if len(sameNoActivityIds) == 0 {
		sameNoActivityIds = []int64{int64(in.Id)}
	}

	memberInfo = contexts.GetMemberUser(ctx)
	couponReceivedMap := make(map[string]*couponReceivedMapItem) // key: linkType-linkId, value: 已领取数量
	if memberInfo != nil && len(couponIds) > 0 {
		// 根据不同的优惠券类型查询不同的表
		for i, couponId := range couponIds {
			linkType := couponLinkTypes[i]
			var receivedCount int
			var todayReceivedCount int

			if linkType == "coupon" {
				// 查询 hg_pms_coupon 表 - 总领取数量
				count, err := dao.PmsCoupon.Ctx(ctx).
					Where(dao.PmsCoupon.Columns().MemberId, memberInfo.Id).
					Where(dao.PmsCoupon.Columns().CouponTypeId, couponId).
					Where(dao.PmsCoupon.Columns().Source, 5).
					WhereIn(dao.PmsCoupon.Columns().IndexActivityId, sameNoActivityIds).
					Count()
				if err != nil {
					g.Log().Errorf(ctx, "查询用户已领取优惠券数量失败: %v", err)
					receivedCount = 0
				} else {
					receivedCount = int(count)
				}

				// 查询当前周期内领取数量（以第一次领取日期为起点，每N天为一个周期）
				couponLimitDays := limitDaysMap["coupon-"+gconv.String(couponId)]
				if couponLimitDays > 0 {
					var firstReceiveDateRow *struct{ D string }
					_ = dao.PmsCoupon.Ctx(ctx).
						Fields("DATE("+dao.PmsCoupon.Columns().CreateAt+") as d").
						Where(dao.PmsCoupon.Columns().MemberId, memberInfo.Id).
						Where(dao.PmsCoupon.Columns().CouponTypeId, couponId).
						Where(dao.PmsCoupon.Columns().Source, 5).
						WhereIn(dao.PmsCoupon.Columns().IndexActivityId, sameNoActivityIds).
						OrderAsc(dao.PmsCoupon.Columns().CreateAt).
						Limit(1).
						Scan(&firstReceiveDateRow)
					var firstReceiveDate string
					if firstReceiveDateRow != nil {
						firstReceiveDate = firstReceiveDateRow.D
					}
					if firstReceiveDate != "" {
						firstDate, _ := gtime.StrToTime(firstReceiveDate)
						nowDate, _ := gtime.StrToTime(gtime.Now().Format("Y-m-d"))
						daysDiff := int(nowDate.Sub(firstDate).Hours() / 24)
						periodIndex := daysDiff / couponLimitDays
						periodStartDate := firstDate.AddDate(0, 0, periodIndex*couponLimitDays).Format("Y-m-d")
						periodStart := periodStartDate + " 00:00:00"
						periodEnd := firstDate.AddDate(0, 0, (periodIndex+1)*couponLimitDays).Format("Y-m-d") + " 00:00:00"
						todayCount, err := dao.PmsCoupon.Ctx(ctx).
							Where(dao.PmsCoupon.Columns().MemberId, memberInfo.Id).
							Where(dao.PmsCoupon.Columns().CouponTypeId, couponId).
							Where(dao.PmsCoupon.Columns().Source, 5).
							WhereIn(dao.PmsCoupon.Columns().IndexActivityId, sameNoActivityIds).
							WhereGTE(dao.PmsCoupon.Columns().CreateAt, periodStart).
							WhereLT(dao.PmsCoupon.Columns().CreateAt, periodEnd).
							Count()
						if err != nil {
							g.Log().Errorf(ctx, "查询用户周期内已领取优惠券数量失败: %v", err)
						} else {
							todayReceivedCount = int(todayCount)
						}
					}
				}
			} else if linkType == "thCoupon" {
				// 查询 hg_th_member_coupon 表（礼品券也用同一个表） - 总领取数量
				count, err := dao.ThMemberCoupon.Ctx(ctx).
					Where(dao.ThMemberCoupon.Columns().MemberId, memberInfo.Id).
					Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
					Where(dao.ThMemberCoupon.Columns().Source, 4).
					WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
					Count()
				if err != nil {
					g.Log().Errorf(ctx, "查询用户已领取礼品券数量失败: %v", err)
					receivedCount = 0
				} else {
					receivedCount = int(count)
				}

				// 查询当前周期内领取数量（以第一次领取日期为起点，每N天为一个周期）
				thCouponLimitDays := limitDaysMap["thCoupon-"+gconv.String(couponId)]
				if thCouponLimitDays > 0 {
					var thFirstReceiveDateRow *struct{ D string }
					_ = dao.ThMemberCoupon.Ctx(ctx).
						Fields("DATE("+dao.ThMemberCoupon.Columns().CreateAt+") as d").
						Where(dao.ThMemberCoupon.Columns().MemberId, memberInfo.Id).
						Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
						Where(dao.ThMemberCoupon.Columns().Source, 4).
						WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
						OrderAsc(dao.ThMemberCoupon.Columns().CreateAt).
						Limit(1).
						Scan(&thFirstReceiveDateRow)
					var thFirstReceiveDate string
					if thFirstReceiveDateRow != nil {
						thFirstReceiveDate = thFirstReceiveDateRow.D
					}
					if thFirstReceiveDate != "" {
						thFirstDate, _ := gtime.StrToTime(thFirstReceiveDate)
						thNowDate, _ := gtime.StrToTime(gtime.Now().Format("Y-m-d"))
						thDaysDiff := int(thNowDate.Sub(thFirstDate).Hours() / 24)
						thPeriodIndex := thDaysDiff / thCouponLimitDays
						thPeriodStartDate := thFirstDate.AddDate(0, 0, thPeriodIndex*thCouponLimitDays).Format("Y-m-d")
						thPeriodStart := thPeriodStartDate + " 00:00:00"
						thPeriodEnd := thFirstDate.AddDate(0, 0, (thPeriodIndex+1)*thCouponLimitDays).Format("Y-m-d") + " 00:00:00"
						todayCount, err := dao.ThMemberCoupon.Ctx(ctx).
							Where(dao.ThMemberCoupon.Columns().MemberId, memberInfo.Id).
							Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
							Where(dao.ThMemberCoupon.Columns().Source, 4).
							WhereIn(dao.ThMemberCoupon.Columns().IndexActivityId, sameNoActivityIds).
							WhereGTE(dao.ThMemberCoupon.Columns().CreateAt, thPeriodStart).
							WhereLT(dao.ThMemberCoupon.Columns().CreateAt, thPeriodEnd).
							Count()
						if err != nil {
							g.Log().Errorf(ctx, "查询用户周期内已领取礼品券数量失败: %v", err)
						} else {
							todayReceivedCount = int(todayCount)
						}
					}
				}
			}

			key := gconv.String(linkType) + "-" + gconv.String(couponId)
			couponReceivedMap[key] = &couponReceivedMapItem{
				memberReceived:      receivedCount,
				memberTodayReceived: todayReceivedCount,
			}
		}
	}

	// 查询文章关联的优惠券配置信息（包含AvailableQuantity）
	if len(couponIds) > 0 {
		var articleCoupons []*entity.HomepageArticleCoupon
		if err = dao.HomepageArticleCoupon.Ctx(ctx).
			Where(dao.HomepageArticleCoupon.Columns().ArticleId, in.Id).
			Scan(&articleCoupons); err != nil && !errors.Is(err, sql.ErrNoRows) {
			g.Log().Errorf(ctx, "查询文章关联优惠券配置失败: %v", err)
		} else {
			// 将查询结果映射到couponReceivedMap中
			for _, coupon := range articleCoupons {
				key := coupon.CouponType + "-" + gconv.String(coupon.CouponId)
				if item, exists := couponReceivedMap[key]; exists {
					item.availableQuantity = coupon.AvailableQuantity
					item.perDayAvailable = coupon.PerDayAvailable
					item.limitDays = coupon.LimitDays
				} else {
					// 如果map中不存在该key，创建新的item
					couponReceivedMap[key] = &couponReceivedMapItem{
						memberReceived:      0,
						memberTodayReceived: 0,
						totalNum:            0,
						totalReceived:       0,
						availableQuantity:   coupon.AvailableQuantity,
						perDayAvailable:     coupon.PerDayAvailable,
						limitDays:           coupon.LimitDays,
					}
				}
			}
		}
	}

	// 处理链接列表
	for _, v := range links {
		// 根据linkType查询对应表的名称字段
		title := v.Title // 默认使用原标题
		subTitle := ""   // 副标题
		atLeast := 0.0   // 满多少元使用
		switch v.LinkType {
		case "coupon":
			// 查询coupon_type表的coupon_name字段
			var couponType *entity.PmsCouponType
			if err = dao.PmsCouponType.Ctx(ctx).WherePri(v.LinkId).Hook(hook.PmsFindLanguageValueHook).Scan(&couponType); err == nil && couponType != nil {
				title = couponType.CouponName
				atLeast = couponType.AtLeast
			}
			key := "coupon-" + gconv.String(v.LinkId)
			item, ok := couponReceivedMap[key]
			if !ok || item == nil {
				item = &couponReceivedMapItem{}
				couponReceivedMap[key] = item
			}
			if couponType != nil {
				item.totalNum = couponType.Count
				item.totalReceived = couponType.LeadCount
			}
		case "thCoupon":
			// 查询th_coupon表的coupon_name字段
			var thCoupon *entity.ThCoupon
			if err = dao.ThCoupon.Ctx(ctx).WherePri(v.LinkId).Hook(hook.PmsFindLanguageValueHook).Scan(&thCoupon); err == nil && thCoupon != nil {
				title = thCoupon.CouponName
				subTitle = thCoupon.CouponSubName
			}
			key := "thCoupon-" + gconv.String(v.LinkId)
			item, ok := couponReceivedMap[key]
			if !ok || item == nil {
				item = &couponReceivedMapItem{}
				couponReceivedMap[key] = item
			}
			if thCoupon != nil {
				item.totalNum = 0
				item.totalReceived = thCoupon.Count
			}
		case "hotelDetail":
			// 查询pms_property表的name字段
			var property *entity.PmsProperty
			if err = dao.PmsProperty.Ctx(ctx).WherePri(v.LinkId).Hook(hook.PmsFindLanguageValueHook).Scan(&property); err == nil && property != nil {
				title = property.Name
			}
		case "foodDetail":
			// 查询food_restaurant表的name字段
			var restaurant *entity.FoodRestaurant
			if err = dao.FoodRestaurant.Ctx(ctx).WherePri(v.LinkId).Hook(hook.PmsFindLanguageValueHook).Scan(&restaurant); err == nil && restaurant != nil {
				title = restaurant.Name
			}
		case "spaDetail":
			// 查询spa_service表的name字段
			var spaService *entity.SpaService
			if err = dao.SpaService.Ctx(ctx).WherePri(v.LinkId).Hook(hook.PmsFindLanguageValueHook).Scan(&spaService); err == nil && spaService != nil {
				title = spaService.Name
			}
		}
		// 重置err，避免查询失败影响主流程
		err = nil

		linkListItem := &input_homepage_article.AppLinkListItem{
			Title:     title,
			LinkScene: v.LinkType,
			AppLink:   v.AppLink,
			WxLink:    v.WxLink,
			LinkText:  v.LinkText,
		}

		if err = json.Unmarshal([]byte(v.UrlParam.String()), &AppUrlParamJson); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
		linkListItem.AppLinkType = gconv.Int(AppUrlParamJson.Type)
		var pushUrlParam g.MapStrStr
		switch linkListItem.AppLinkType {
		case 1:
			pushUrlParam = g.MapStrStr{
				"string": AppUrlParamJson.String,
			}
		case 2:
			pushUrlParam = g.MapStrStr{
				"param": AppUrlParamJson.Param,
			}
		}
		if !g.IsEmpty(pushUrlParam) {
			pushUrlParamJson, _ := json.Marshal(pushUrlParam)
			linkListItem.UrlParam = string(pushUrlParamJson)
		} else {
			linkListItem.UrlParam = ""
		}

		if v.LinkType == "coupon" || v.LinkType == "thCoupon" {
			// 获取用户已领取数量
			key := gconv.String(v.LinkType) + "-" + gconv.String(v.LinkId)

			// 确保 map 中有对应的项
			item, ok := couponReceivedMap[key]
			if !ok || item == nil {
				item = &couponReceivedMapItem{}
				couponReceivedMap[key] = item
			}

			// 判断优惠券状态
			var status int

			// 1. 先判断总库存是否已领完
			if item.totalNum > 0 && item.totalReceived >= item.totalNum {
				if item.memberReceived > 0 {
					// 用户已领过，但总库存已领完
					status = 2
				} else {
					// 用户还没领过，但总库存已领完
					status = 3
				}
			} else if item.availableQuantity > 0 && item.memberReceived >= item.availableQuantity {
				// 2. 判断用户是否已达到个人总领取限制
				status = 2
			} else {
				// 3. 判断N天内领取限制（如果设置了limitDays和perDayAvailable）
				if item.limitDays > 0 && item.perDayAvailable > 0 && item.memberTodayReceived >= item.perDayAvailable {
					// N天内已达到领取限制
					status = 4
				} else {
					// 可领取
					status = 1
				}
			}

			// 判断礼品券领取限制，是否周六周日不可领取， status=5表示今日不可领取
			if v.LinkType == "thCoupon" {
				// 获取当前日期的星期几（0=周日, 1=周一, ..., 6=周六）
				currentWeekday := gtime.Now().Weekday()
				// 检查活动是否有周末限制
				if detail.LimitWeek != "" {
					limits := gstr.Split(detail.LimitWeek, ",")
					// 如果是周六（6）且限制中包含6，或者周日（0）且限制中包含7
					if (currentWeekday == 6 && gstr.InArray(limits, "6")) || (currentWeekday == 0 && gstr.InArray(limits, "7")) {
						status = 5 // 今日不可领取
					}
				}
			}

			// 创建CouponAppLinkListItem，包含额外的subTitle和atLeast字段
			couponLinkItem := &input_homepage_article.CouponAppLinkListItem{
				CouponId:           v.LinkId,
				SubTitle:           subTitle,
				AtLeast:            atLeast,
				Status:             status,
				TotalNum:           item.totalNum,
				TotalReceive:       item.totalReceived,
				MemberReceive:      item.memberReceived,
				MemberTodayReceive: item.memberTodayReceived,
				AvailableQuantity:  item.availableQuantity,
				PerDayAvailable:    item.perDayAvailable,
				LimitDays:          item.limitDays,
				AppLinkListItem:    linkListItem,
			}
			res.CouponLinkList = append(res.CouponLinkList, couponLinkItem)
		} else {
			res.RecommendLinkList = append(res.RecommendLinkList, linkListItem)
		}
	}

	return
}

func (s *sBasicsHomepageArticles) AppList(ctx context.Context, in *input_homepage_article.HomepageArticlesAppListInp) (list []*input_homepage_article.HomepageArticlesAppListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 先查询到临时结构体，包含时间字段
	type tempListModel struct {
		Id           int         `json:"id"`
		Language     string      `json:"language"`
		ListPic      string      `json:"listPic"`
		Title        string      `json:"title"`
		Author       string      `json:"author"`
		Views        int         `json:"views"`
		ThumbNum     int         `json:"thumbNum"`
		StartTime    *gtime.Time `json:"startTime"`
		EndTime      *gtime.Time `json:"endTime"`
		Sort         int         `json:"sort"`
		Chain        string      `json:"chain"             dc:"活动链接方式 IN-内部链接 OUT-外部链接"`
		Path         string      `json:"path"              dc:"活动外链链接"`
		LinkOpenType string      `json:"linkOpenType"      dc:"外链打开方式 1-内部webview 2-外部浏览器"`
	}

	var tempList []*tempListModel

	mod = mod.Fields(tempListModel{}).WithAll()

	mod = mod.Where(dao.HomepageArticles.Columns().Language, contexts.GetLanguage(ctx))
	mod = mod.WhereLTE(dao.HomepageArticles.Columns().StartTime, gtime.Now().Format("Y-m-d H:i:s"))
	mod = mod.WhereGTE(dao.HomepageArticles.Columns().EndTime, gtime.Now().Format("Y-m-d H:i:s"))
	mod = mod.Where(dao.HomepageArticles.Columns().Status, 1)

	// 如果请求头中的x-channel为miniapp，需要minapp_status字段为1
	if r := ghttp.RequestFromCtx(ctx); r != nil && r.GetHeader("x-channel") == "miniapp" {
		mod = mod.Where(dao.HomepageArticles.Columns().MinappStatus, 1)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.HomepageArticles.Columns().Sort).OrderDesc(dao.HomepageArticles.Columns().Id)

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&tempList, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_activity_list_failed"))
			return
		}
	} else {
		if err = mod.Scan(&tempList); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_activity_list_failed"))
			return
		}
	}

	// 转换为最终结果格式
	for _, temp := range tempList {
		item := &input_homepage_article.HomepageArticlesAppListModel{
			Id:           temp.Id,
			Language:     temp.Language,
			ListPic:      temp.ListPic,
			Title:        temp.Title,
			Author:       temp.Author,
			Views:        temp.Views,
			ThumbNum:     temp.ThumbNum,
			Sort:         temp.Sort,
			Chain:        temp.Chain,
			Path:         temp.Path,
			LinkOpenType: temp.LinkOpenType,
		}

		// 格式化时间字段
		if temp.StartTime != nil {
			item.StartTime = temp.StartTime.Format("Y-m-d")
		}
		if temp.EndTime != nil {
			item.EndTime = temp.EndTime.Format("Y-m-d")
		}

		list = append(list, item)
	}

	// 判断是否已点赞
	memberInfo := contexts.GetMemberUser(ctx)
	if memberInfo != nil && len(list) > 0 {
		// 收集所有文章ID
		var articleIds []interface{}
		for _, article := range list {
			articleIds = append(articleIds, article.Id)
		}

		// 批量查询当前用户对这些文章的点赞记录
		var thumbRecords []*entity.HomepageArticleThumb
		if err = dao.HomepageArticleThumb.Ctx(ctx).
			Where(dao.HomepageArticleThumb.Columns().MemberId, memberInfo.Id).
			WhereIn(dao.HomepageArticleThumb.Columns().ArticleId, articleIds).
			Scan(&thumbRecords); err != nil {
			err = gerror.Wrap(err, "查询点赞记录失败")
			return
		}

		// 构建点赞记录映射
		thumbMap := make(map[int]bool)
		for _, thumb := range thumbRecords {
			thumbMap[int(thumb.ArticleId)] = true
		}

		// 设置每篇文章的点赞状态
		for _, article := range list {
			article.HasThumb = thumbMap[article.Id]
		}
	}

	return
}

func (s *sBasicsHomepageArticles) Thumb(ctx context.Context, in *input_homepage_article.HomepageArticlesThumbInp) (err error) {
	var (
		lastInsertId int64
		MemberInfo   *model.MemberIdentity
		thumb        *entity.HomepageArticleThumb
	)
	MemberInfo = contexts.GetMemberUser(ctx)

	if err = dao.HomepageArticleThumb.Ctx(ctx).
		Where(dao.HomepageArticleThumb.Columns().ArticleId, in.ArticleId).
		Where(dao.HomepageArticleThumb.Columns().MemberId, MemberInfo.Id).Scan(&thumb); err != nil {
		// 点赞失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "thumb_failed"))
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 取消点赞
		if in.Type == 2 {
			if g.IsEmpty(thumb) {
				// 未点赞
				err = gerror.New(gi18n.T(ctx, "have_not_thumb_cannot_cancel_thumb"))
				return
			}

			// 取消点赞
			if _, err = dao.HomepageArticleThumb.Ctx(ctx).Where(dao.HomepageArticleThumb.Columns().ArticleId, in.ArticleId).Where(dao.HomepageArticleThumb.Columns().MemberId, MemberInfo.Id).Delete(); err != nil {
				// 取消点赞失败，请稍后重试！
				err = gerror.Wrap(err, gi18n.T(ctx, "cancel_thumb_failed"))
			}
			// 同步减少活动点赞数
			if _, err = dao.HomepageArticles.Ctx(ctx).Where(dao.HomepageArticles.Columns().Id, in.ArticleId).Update(g.MapStrAny{
				dao.HomepageArticles.Columns().ThumbNum: gdb.Raw("thumb_num-1"),
			}); err != nil {
				err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_update_likes"))
				return
			}
			return
		}

		// 已点赞
		if !g.IsEmpty(thumb) {
			err = gerror.New(gi18n.T(ctx, "already_thumb_cannot_thumb_again"))
			return
		}

		// 添加点赞
		if lastInsertId, err = dao.HomepageArticleThumb.Ctx(ctx).Data(&entity.HomepageArticleThumb{
			MemberId:  uint64(MemberInfo.Id),
			ArticleId: uint64(in.ArticleId),
		}).InsertAndGetId(); err != nil {
			// 点赞失败，请稍后重试！
			err = gerror.Wrap(err, gi18n.T(ctx, "thumb_failed"))
			return
		}
		if lastInsertId < 1 {
			// 点赞失败，请稍后重试！
			err = gerror.New(gi18n.T(ctx, "thumb_failed"))
			return
		}

		// 更新活动点赞数
		if _, err = dao.HomepageArticles.Ctx(ctx).Where(dao.HomepageArticles.Columns().Id, in.ArticleId).Update(g.MapStrAny{
			dao.HomepageArticles.Columns().ThumbNum: gdb.Raw("thumb_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_update_likes"))
			return
		}

		return
	})

}
