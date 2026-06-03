package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/format"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"

	"APT/api/app/basics"
)

func (c *ControllerBasics) HomeDateLoad(ctx context.Context, req *basics.HomeDateLoadReq) (res *basics.HomeDateLoadRes, err error) {
	var (
		pmsBanner     *entity.PmsBanner
		pmsBannerList []*entity.IndexBanner
		MemberInfo    *model.MemberIdentity
		//AlertMessage           []*entity.PmsNotify
		ImmovablesPathJsonData *basics.ImmovablesPathJsonData
	)
	res = new(basics.HomeDateLoadRes)
	res.Banner = make([]*input_basics.PmsBannerListModel, 0)
	//if res.Banner, _, err = service.BasicsBanner().List(ctx, &input_basics.PmsBannerListInp{
	//	PageReq: input_form.PageReq{
	//		Page:       1,
	//		PerPage:    10,
	//		Pagination: false,
	//	},
	//	Model:    "BANNER",
	//	Language: contexts.GetLanguage(ctx),
	//}); err != nil {
	//	return
	//}
	if res.Property, _, err = service.HotelService().PropertyAppIndexList(ctx, &input_hotel.PmsPropertyListInp{
		PageReq: input_form.PageReq{
			Page:       1,
			PerPage:    3,
			Pagination: true,
		},
		IsFindClose: true,
	}); err != nil {
		return
	}
	// 获取导航栏
	// 获取配置信息
	var Config *input_basics.GetConfigModel
	Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "indexnav",
	})
	if err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "config_info_not_found"))
		return
	}

	// 1-开启则显示导航栏  2-关闭则隐藏导航栏
	isOpen := Config.List["isOpen"]
	if isOpen == 1 {
		filter := &input_basics.PmsIndexNavAllInp{
			Status: 1,
		}
		// 如果请求头中的x-channel为miniapp，需要minapp_status字段为1
		if r := ghttp.RequestFromCtx(ctx); r != nil && r.GetHeader("x-channel") == "miniapp" {
			filter.MinappStatus = 1
		}
		if res.NavList, err = service.BasicsIndexNav().ApiNavAll(ctx, filter); err != nil {
			return
		}
	}

	PmsBanner := dao.PmsBanner.Ctx(ctx).
		Where(dao.PmsBanner.Columns().Model, "EVENTS").
		Where(dao.PmsBanner.Columns().Language, contexts.GetLanguage(ctx))

	if ghttp.RequestFromCtx(ctx).GetHeader("X-Version") == "1.1.15" {
		PmsBanner = PmsBanner.WhereLike(dao.PmsBanner.Columns().BannerName, "immovables_%")
	} else {
		PmsBanner = PmsBanner.Where(dao.PmsBanner.Columns().BannerStatus, 1)
	}
	if err = PmsBanner.Scan(&pmsBanner); err != nil {
		return
	}
	if !g.IsEmpty(pmsBanner) {
		res.ImmovablesCover = pmsBanner.BannerImage
		res.ImmovablesPath = pmsBanner.Path
		res.ImmovablesChain = pmsBanner.Chain
		res.ImmovablesAppPath = pmsBanner.Path
		res.ImmovablesWxPath = pmsBanner.Path
		if pmsBanner.Chain == "IN" && pmsBanner.Path != "" {
			if err = json.Unmarshal([]byte(pmsBanner.Path), &ImmovablesPathJsonData); err != nil {
				// 解析路径失败
				err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
				return
			}
			res.ImmovablesAppPath = ImmovablesPathJsonData.App
			res.ImmovablesWxPath = ImmovablesPathJsonData.Weapp
		}
	}

	// 从配置表获取广告位滚动间隔
	var scrollInterval *gvar.Var
	if scrollInterval, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Group, "indexbanner").
		Where(dao.SysConfig.Columns().Key, "scrollInterval").
		Value(dao.SysConfig.Columns().Value); err != nil {
		return
	}
	if scrollInterval == nil {
		res.CarouselScrollInterval = 1
	}
	res.CarouselScrollInterval = gvar.New(scrollInterval).Int()

	mod := dao.IndexBanner.Ctx(ctx)
	// 如果请求头中的x-channel为miniapp，需要minapp_status字段为1
	if r := ghttp.RequestFromCtx(ctx); r != nil && r.GetHeader("x-channel") == "miniapp" {
		mod = mod.Where(dao.IndexBanner.Columns().MinappStatus, 1)
	}
	if err = mod.
		Where(dao.IndexBanner.Columns().Language, contexts.GetLanguage(ctx)).
		Where(dao.IndexBanner.Columns().BannerStatus, 1).
		OrderDesc(dao.IndexBanner.Columns().Sort).
		Scan(&pmsBannerList); err != nil {
		return
	}
	if !g.IsEmpty(pmsBannerList) {
		for _, v := range pmsBannerList {
			var item *basics.CarouselImageListItem
			var CarouselPathJsonData *basics.ImmovablesPathJsonData
			item = &basics.CarouselImageListItem{
				Cover:        v.BannerImage,
				Path:         v.Path,
				Chain:        v.Chain,
				AppPath:      v.Path,
				WxPath:       v.Path,
				LinkOpenType: v.LinkOpenType,
			}
			if v.Chain == "IN" && v.Path != "" {
				if err = json.Unmarshal([]byte(v.Path), &CarouselPathJsonData); err != nil {
					// 解析路径失败
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				item.AppPath = CarouselPathJsonData.App
				item.WxPath = CarouselPathJsonData.Weapp
			}
			res.CarouselImageList = append(res.CarouselImageList, item)
		}
	}

	MemberInfo = contexts.GetMemberUser(ctx)
	//if !g.IsEmpty(MemberInfo) {
	//	// 查询首页弹窗信息
	//	if err = dao.PmsNotify.Ctx(ctx).
	//		Where(dao.PmsNotify.Columns().NotifyType, "GIFT").
	//		Where(dao.PmsNotify.Columns().IsRead, "N").
	//		Where(dao.PmsNotify.Columns().MemberId, MemberInfo.Id).Scan(&AlertMessage); err != nil {
	//		return
	//	}
	//
	//	if len(AlertMessage) > 0 {
	//		for _, v := range AlertMessage {
	//			res.AlertMessage = append(res.AlertMessage, &basics.HomeDateAlertMessage{
	//				Id:            v.Id,
	//				NotifyType:    v.NotifyType,
	//				NotifyContent: v.NotifyData.Get(contexts.GetLanguage(ctx)).String(),
	//			})
	//		}
	//	}
	//
	//}

	// 查询短租模式最小预定区间是多少
	if res.MinDaysNotice, err = dao.PmsProperty.Ctx(ctx).
		Where(dao.PmsProperty.Columns().LeaseClose, 1).
		Min(dao.PmsProperty.Columns().MinDaysNotice); err != nil {
		return
	}
	if res.MinDaysNotice < 7 {
		res.MinDaysNotice = 7
	}

	// 查询行程卡列表
	if !g.IsEmpty(MemberInfo) {
		if res.TripCardList, err = c.getTripCardList(ctx, MemberInfo.Id); err != nil {
			return
		}
	}

	// 未读消息列表列表
	if !g.IsEmpty(MemberInfo) {
		if res.UnReadMessageList, _, err = service.BasicsSystemMessage().AppList(ctx, &input_basics.MessageAppListInp{
			PageReq: input_form.PageReq{
				Pagination: false,
			},
			MemberId:  MemberInfo.Id,
			UnRead:    true,
			ShowIndex: true,
		}); err != nil {
			return
		}
	}

	return
}

func (c *ControllerBasics) HomeDateConfig(ctx context.Context, req *basics.HomeDateConfigReq) (res *basics.HomeDateConfigRes, err error) {
	res = new(basics.HomeDateConfigRes)
	if err = dao.PmsAppconfig.Ctx(ctx).Where(dao.PmsAppconfig.Columns().Language, contexts.GetLanguage(ctx)).Scan(&res.ConfigList); err != nil {
		return
	}
	for k, v := range res.ConfigList {
		res.ConfigList[k].Value = format.GetHtmlImgSrc(v.Value)
	}
	return
}
func (c *ControllerBasics) ReadAlertMessage(ctx context.Context, req *basics.ReadAlertMessageReq) (res *basics.ReadAlertMessageRes, err error) {
	if _, err = dao.PmsNotify.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.PmsNotify.Columns().IsRead: "Y",
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) HomeDeviceInfo(ctx context.Context, req *basics.HomeDeviceInfoReq) (res *basics.HomeDeviceInfoRes, err error) {
	var (
		MaxDeviceInfo *entity.PmsScanDevice
		MaxDeviceSn   int
	)
	if err = dao.PmsScanDevice.Ctx(ctx).Fields(dao.PmsScanDevice.Columns().DeviceSn).OrderDesc(dao.PmsScanDevice.Columns().DeviceSn).Scan(&MaxDeviceInfo); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_get_device"))
		return
	}
	if MaxDeviceInfo == nil {
		MaxDeviceSn = 10001
	} else {
		MaxDeviceSn = MaxDeviceInfo.DeviceSn + 1
	}
	if _, err = dao.PmsScanDevice.Ctx(ctx).
		Data(g.Map{
			dao.PmsScanDevice.Columns().DeviceSn: MaxDeviceSn,
		}).OmitEmptyData().Insert(); err != nil {
		err = gerror.Wrap(err, "新增设备码，请稍后重试！")
	}
	res = new(basics.HomeDeviceInfoRes)
	res.DeviceSn = MaxDeviceSn
	return
}
func (c *ControllerBasics) ConfigGet(ctx context.Context, req *basics.ConfigGetReq) (res *basics.ConfigGetRes, err error) {
	res = new(basics.ConfigGetRes)
	res.GetConfigModel, err = service.BasicsConfig().GetConfigByGroup(ctx, &req.GetConfigInp)
	return
}

// getTripCardList 获取行程卡列表
func (c *ControllerBasics) getTripCardList(ctx context.Context, memberId int) (tripCardList []*basics.TripCardListItem, err error) {
	type CarOrderItem struct {
		*entity.CarOrder
		StartServiceAddress *struct {
			gmeta.Meta `orm:"table:hg_car_address"`
			Id         int    `json:"id"            dc:""`
			SubName    string `json:"subName"       dc:"地点名称"`
		} `json:"startServiceAddress" orm:"with:id=start_address_id"  dc:"出发地"`
		EndServiceAddress *struct {
			gmeta.Meta `orm:"table:hg_car_address"`
			Id         int    `json:"id"            dc:""`
			SubName    string `json:"subName"       dc:"地点名称"`
		} `json:"endServiceAddress" orm:"with:id=end_address_id"  dc:"目的地"`
		ServiceDetail *struct {
			gmeta.Meta    `orm:"table:hg_car_service"`
			Id            int `json:"id"               dc:"服务ID"`
			CarTypeId     int `json:"carTypeId"        dc:"车型ID"`
			CarTypeDetail *struct {
				gmeta.Meta `orm:"table:hg_car_car_type"`
				Id         int    `json:"id"                  dc:"车型ID"`
				Name       string `json:"name"                dc:"车型名称"`
			} `json:"carTypeDetail" orm:"with:id=car_type_id"`
		} `json:"serviceDetail" orm:"with:id=service_id"`
	}

	type FoodOrderItem struct {
		*entity.FoodOrder
		RestaurantInfo *struct {
			gmeta.Meta `orm:"table:hg_food_restaurant"`
			Id         int    `json:"id"            dc:""`
			Name       string `json:"name"       dc:"餐厅名称"`
		} `json:"restaurantInfo" orm:"with:id=restaurant_id"  dc:"餐厅信息"`
		GoodsInfo *struct {
			gmeta.Meta `orm:"table:hg_food_goods"`
			Id         int    `json:"id"            dc:""`
			GoodsName  string `json:"goodsName"       dc:"套餐名称"`
		} `json:"goodsInfo" orm:"with:id=goods_id"  dc:"套餐信息"`
	}

	type SpaOrderItem struct {
		*entity.SpaOrder
		ServiceInfo *struct {
			gmeta.Meta `orm:"table:hg_spa_service"`
			Id         int    `json:"id"            dc:""`
			Name       string `json:"name"       dc:"服务名称"`
		} `json:"serviceInfo" orm:"with:id=service_id"  dc:"服务信息"`
		GoodsInfo *struct {
			gmeta.Meta `orm:"table:hg_spa_service_goods"`
			Id         int    `json:"id"            dc:""`
			GoodsName  string `json:"goodsName"       dc:"项目名称"`
		} `json:"goodsInfo" orm:"with:id=goods_id"  dc:"项目信息"`
	}

	var (
		hotelOrders []*entity.PmsAppReservation
		foodOrders  []*FoodOrderItem
		carOrders   []*CarOrderItem
		spaOrders   []*SpaOrderItem
	)

	// 查询酒店订单 (APP 订单)
	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().MemberId, memberId).
		Where(dao.PmsAppReservation.Columns().Status, "confirmed").
		Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
		Where(dao.PmsAppReservation.Columns().CheckinStatus, "before_checkin").
		WhereGTE(dao.PmsAppReservation.Columns().CheckinDate, gtime.Now().Format("Y-m-d")).
		Scan(&hotelOrders); err != nil {
		return
	}

	// 查询餐厅订单 - 根据PHP条件转换
	// 构建复杂的OR条件查询
	whereCondition := "(" +
		"(order_type = 'CRS' AND deposit_pay_status = 'HAVE_PAID' AND booking_status = 'CONFIRMED' AND remain_pay_status = 'HAVE_PAID' AND order_status = 'HAVE_PAID' AND verify_status = 'WAIT_VERIFY' AND (UNIX_TIMESTAMP(CONCAT(book_date, ' ', book_time)) > UNIX_TIMESTAMP(NOW())))" +
		" OR " +
		"((order_type = 'TORETA' OR order_type = 'CRSALL') AND order_status = 'HAVE_PAID' AND booking_status = 'CONFIRMED' AND verify_status = 'WAIT_VERIFY' AND (UNIX_TIMESTAMP(CONCAT(book_date, ' ', book_time)) > UNIX_TIMESTAMP(NOW())))" +
		")"

	if err = dao.FoodOrder.Ctx(ctx).WithAll().
		Where(dao.FoodOrder.Columns().MemberId, memberId).
		Where(whereCondition).
		WhereGTE(dao.FoodOrder.Columns().BookDate, gtime.Now().Format("Y-m-d")).
		Hook(hook.PmsFindLanguageValueHook).
		Scan(&foodOrders); err != nil {
		return
	}

	// 查询接送机订单（带关联查询）
	if err = dao.CarOrder.Ctx(ctx).WithAll().
		Where(dao.CarOrder.Columns().MemberId, memberId).
		Where(dao.CarOrder.Columns().OrderStatus, "WAIT_SERVE").
		Where(dao.CarOrder.Columns().DispatchStatus, "DONE").
		WhereGTE(dao.CarOrder.Columns().BookDate, gtime.Now().Format("Y-m-d")).
		Hook(hook.PmsFindLanguageValueHook).
		Scan(&carOrders); err != nil {
		return
	}

	// 查询按摩订单
	if err = dao.SpaOrder.Ctx(ctx).WithAll().
		Where(dao.SpaOrder.Columns().MemberId, memberId).
		Where(dao.SpaOrder.Columns().OrderStatus, "WAIT_SERVE").
		Where(dao.SpaOrder.Columns().DispatchStatus, "DONE").
		WhereGTE(dao.SpaOrder.Columns().BookDate, gtime.Now().Format("Y-m-d")).
		Hook(hook.PmsFindLanguageValueHook).
		Scan(&spaOrders); err != nil {
		return
	}

	// 转换酒店订单数据
	for _, order := range hotelOrders {
		// 计算入住晚数
		nights := 1
		checkinDateStr := ""
		checkoutDateStr := ""
		checkinTimeStr := "15:00"  // 默认入住时间
		checkoutTimeStr := "11:00" // 默认退房时间

		if order.CheckinDate != nil && order.CheckoutDate != nil {
			if !order.CheckinDate.IsZero() && !order.CheckoutDate.IsZero() {
				checkinDateStr = order.CheckinDate.Format("Y-m-d")
				checkoutDateStr = order.CheckoutDate.Format("Y-m-d")
				nights = convert.Diff(checkinDateStr, checkoutDateStr, "days")
			}
		}

		if order.CheckinTime != "" {
			checkinTimeStr = order.CheckinTime
		}
		if order.CheckoutTime != "" {
			checkoutTimeStr = order.CheckoutTime
		}

		// 构建预订时间
		bookingDateTimeStr := ""
		if checkinDateStr != "" {
			bookingDateTimeStr = checkinDateStr + " " + checkinTimeStr
		}

		tripCard := &basics.TripCardListItem{
			OrderId:         int(order.Id),
			OrderSn:         order.OrderSn,
			OrderScene:      "hotel",
			BookingDateTime: bookingDateTimeStr,
			HotelOrderInfo: &basics.HotelOrderInfoItem{
				CheckinDate:  checkinDateStr,
				CheckoutDate: checkoutDateStr,
				CheckinTime:  checkinTimeStr,
				CheckoutTime: checkoutTimeStr,
				Nights:       nights,
			},
		}
		tripCardList = append(tripCardList, tripCard)
	}

	// 转换餐厅订单数据
	for _, order := range foodOrders {
		bookingDateTimeStr := ""
		if order.BookDate != "" && order.BookTime != "" {
			bookingDateTimeStr = order.BookDate + " " + order.BookTime
		}

		tripCard := &basics.TripCardListItem{
			OrderId:         int(order.Id),
			OrderSn:         order.OrderSn,
			OrderScene:      "food",
			BookingDateTime: bookingDateTimeStr,
			FoodOrderInfo: &basics.FoodOrderInfoItem{
				BookDate:       order.BookDate,
				BookTime:       order.BookTime,
				RestaurantName: order.RestaurantInfo.Name,
				GoodsName:      order.GoodsInfo.GoodsName,
			},
		}
		tripCardList = append(tripCardList, tripCard)
	}

	// 转换接送机订单数据
	for _, order := range carOrders {

		carType := ""
		if order.OrderType == "CRS" {
			carType = order.ServiceDetail.CarTypeDetail.Name
		} else {
			carType = "INNN"
		}

		carOrderInfo := &basics.CarOrderInfoItem{
			ServiceType:      order.ServiceType,
			BookDate:         order.BookDate,
			BookTime:         order.BookTime,
			StartAddressName: order.StartServiceAddress.SubName,
			EndAddressName:   order.EndServiceAddress.SubName,
			CarTypeName:      carType,
		}

		// 暂时不查询关联信息，保持简单结构
		// 后续可以根据需要优化为批量查询或使用JOIN

		bookingDateTimeStr := ""
		if order.BookDate != "" && order.BookTime != "" {
			bookingDateTimeStr = order.BookDate + " " + order.BookTime
		}

		tripCard := &basics.TripCardListItem{
			OrderId:         int(order.Id),
			OrderSn:         order.OrderSn,
			OrderScene:      "car",
			BookingDateTime: bookingDateTimeStr,
			CarOrderInfo:    carOrderInfo,
		}
		tripCardList = append(tripCardList, tripCard)
	}

	// 转换按摩订单数据
	for _, order := range spaOrders {
		bookingDateTimeStr := ""
		if order.BookDate != "" && order.BookTime != "" {
			bookingDateTimeStr = order.BookDate + " " + order.BookTime
		}

		tripCard := &basics.TripCardListItem{
			OrderId:         int(order.Id),
			OrderSn:         order.OrderSn,
			OrderScene:      "spa",
			BookingDateTime: bookingDateTimeStr,
			SpaOrderInfo: &basics.SpaOrderInfoItem{
				ServiceType: order.ServiceType,
				BookDate:    order.BookDate,
				BookTime:    order.BookTime,
				ServiceName: order.ServiceInfo.Name,
				GoodsName:   order.GoodsInfo.GoodsName,
			},
		}
		tripCardList = append(tripCardList, tripCard)
	}

	// 按预订时间升序排列所有订单，最早的排在最前面
	for i := 0; i < len(tripCardList)-1; i++ {
		for j := i + 1; j < len(tripCardList); j++ {
			if tripCardList[i].BookingDateTime > tripCardList[j].BookingDateTime {
				tripCardList[i], tripCardList[j] = tripCardList[j], tripCardList[i]
			}
		}
	}

	// 取前20条
	if len(tripCardList) > 20 {
		tripCardList = tripCardList[:20]
	}

	return
}
