package notify

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"html/template"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/kefu"
)

func (c *ControllerKefu) OrderList(ctx context.Context, req *kefu.OrderListReq) (res *kefu.OrderListRes, err error) {
	list, totalCount, orderCount, err := service.HotelService().ReservationRoomList(ctx, &req.PmsAppReservationRoomListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_hotel.PmsAppReservationRoomListModel{}
	}

	res = new(kefu.OrderListRes)
	res.List = list
	res.OrderCount = orderCount
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerKefu) AppReservationLanguageTemplate(ctx context.Context, req *kefu.AppReservationLanguageTemplateReq) (res *kefu.AppReservationLanguageTemplateRes, err error) {
	var (
		appReservationInfo *AppReservationInfo
		appStay            *AppStayInfo
		appSceneInfo       *AppSceneInfo
		Property           *entity.PmsProperty
		foodOrderInfo      *AppFoodInfo
		spaOrderInfo       *AppSpaInfo
		carOrderInfo       *entity.CarOrder
		tmpl               *template.Template
		buffer             bytes.Buffer
	)
	appSceneInfo = new(AppSceneInfo)
	g.Dump(contexts.GetLanguage(ctx))
	tmpl = new(template.Template)
	if req.Type == 1 {
		if err = dao.PmsAppReservation.Ctx(ctx).
			WherePri(req.Id).WithAll().Hook(hook.PmsFindLanguageValueHook).
			Scan(&appReservationInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		if g.IsEmpty(appReservationInfo) {
			err = gerror.New("订单不存在")
			return
		}
		appReservationInfo.PmsAppReservation.CheckinTime = fmt.Sprintf("%s %s", gtime.New(appReservationInfo.CheckinDate).Format("y-m-d"), appReservationInfo.CheckinTime)
		appReservationInfo.PmsAppReservation.CheckoutTime = fmt.Sprintf("%s %s", gtime.New(appReservationInfo.CheckoutDate).Format("y-m-d"), appReservationInfo.CheckoutTime)
		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateReservationZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateReservationZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateReservationEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateReservationJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateReservationKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}

	} else if req.Type == 2 {
		if err = dao.PmsAppStay.Ctx(ctx).
			WherePri(req.Id).WithAll().Hook(hook.PmsFindLanguageValueHook).
			Scan(&appStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(appStay) {
			err = gerror.New("订单不存在")
			return
		}
		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateOrderZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateOrderZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateOrderEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateOrderJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateOrderKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}
	} else if req.Type == 3 {
		if err = dao.PmsProperty.Ctx(ctx).
			WherePri(req.Id).Hook(hook.PmsFindLanguageValueHook).
			Scan(&Property); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(Property) {
			err = gerror.New("物业不存在")
			return
		}
		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplatePropertyZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplatePropertyZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplatePropertyEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplatePropertyJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplatePropertyKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}
	} else if req.Type == 4 {
		// 餐厅订单
		if err = dao.FoodOrder.Ctx(ctx).
			WherePri(req.Id).WithAll().Hook(hook.PmsFindLanguageValueHook).
			Scan(&foodOrderInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		if g.IsEmpty(foodOrderInfo) {
			err = gerror.New("订单不存在")
			return
		}

		SceneMap := g.MapStrStr{
			"zh":    "餐厅预定",
			"zh_CN": "餐廳預定",
			"en":    "Restaurant Reservations",
			"ja":    "レストランの予約",
			"ko":    "레스토랑 예약",
		}

		appSceneInfo.Scene = SceneMap[contexts.GetLanguage(ctx)]
		appSceneInfo.Service = foodOrderInfo.RestaurantDetail.Name
		appSceneInfo.BookingName = foodOrderInfo.BookingName
		appSceneInfo.BookDatetime = foodOrderInfo.BookDatetime
		appSceneInfo.OrderSn = foodOrderInfo.OrderSn
		appSceneInfo.CreatedAt = foodOrderInfo.CreatedAt

		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}
	} else if req.Type == 5 {
		// 按摩订单
		if err = dao.SpaOrder.Ctx(ctx).
			WherePri(req.Id).WithAll().Hook(hook.PmsFindLanguageValueHook).
			Scan(&spaOrderInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		if g.IsEmpty(spaOrderInfo) {
			err = gerror.New("订单不存在")
			return
		}

		SceneMap := g.MapStrStr{
			"zh":    "按摩预约",
			"zh_CN": "按摩預約",
			"en":    "Massage Appointment",
			"ja":    "マッサージの予約",
			"ko":    "마사지 예약",
		}

		ServiceType1Map := g.MapStrStr{
			"zh":    "到店服务",
			"zh_CN": "到店服務",
			"en":    "In-Store",
			"ja":    "来店サービス",
			"ko":    "매장 방문 서비스",
		}

		ServiceType2Map := g.MapStrStr{
			"zh":    "上门服务",
			"zh_CN": "上門服務",
			"en":    "To-Room",
			"ja":    "出張サービス",
			"ko":    "방문 서비스",
		}

		appSceneInfo.Scene = SceneMap[contexts.GetLanguage(ctx)]
		if spaOrderInfo.ServiceType == 1 {
			appSceneInfo.Service = ServiceType1Map[contexts.GetLanguage(ctx)]
		} else {
			appSceneInfo.Service = spaOrderInfo.PropertyDetail.Name + " " + ServiceType2Map[contexts.GetLanguage(ctx)]
		}
		appSceneInfo.BookingName = spaOrderInfo.BookingName
		appSceneInfo.BookDatetime = spaOrderInfo.BookStartTime
		appSceneInfo.OrderSn = spaOrderInfo.OrderSn
		appSceneInfo.CreatedAt = spaOrderInfo.CreatedAt

		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}
	} else if req.Type == 6 {
		// 接送机订单
		if err = dao.CarOrder.Ctx(ctx).
			WherePri(req.Id).WithAll().Hook(hook.PmsFindLanguageValueHook).
			Scan(&carOrderInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		if g.IsEmpty(carOrderInfo) {
			err = gerror.New("订单不存在")
			return
		}

		SceneMap := g.MapStrStr{
			"zh":    "住一接送机/包车",
			"zh_CN": "住一接送機/包車",
			"en":    "Airport Transfer/Charter",
			"ja":    "空港送迎/貸切",
			"ko":    "픽업/전세",
		}

		ServiceType1Map := g.MapStrStr{
			"zh":    "接机",
			"zh_CN": "接機",
			"en":    "Pick-up",
			"ja":    "空港でお迎え",
			"ko":    "공항 픽업",
		}

		ServiceType2Map := g.MapStrStr{
			"zh":    "送机",
			"zh_CN": "送機",
			"en":    "Drop-off",
			"ja":    "空空港に送る",
			"ko":    "공항 드롭오프",
		}

		ServiceType3Map := g.MapStrStr{
			"zh":    "包车",
			"zh_CN": "包車",
			"en":    "Charter",
			"ja":    "貸切",
			"ko":    "전용 차량",
		}

		appSceneInfo.Scene = SceneMap[contexts.GetLanguage(ctx)]
		if carOrderInfo.ServiceType == "PICKUP" {
			appSceneInfo.Service = ServiceType1Map[contexts.GetLanguage(ctx)]
		} else if carOrderInfo.ServiceType == "DELIVERY" {
			appSceneInfo.Service = ServiceType2Map[contexts.GetLanguage(ctx)]
		} else {
			appSceneInfo.Service = ServiceType3Map[contexts.GetLanguage(ctx)]
		}
		appSceneInfo.BookingName = carOrderInfo.BookingName
		appSceneInfo.BookDatetime = carOrderInfo.BookStartTime
		appSceneInfo.OrderSn = carOrderInfo.OrderSn
		appSceneInfo.CreatedAt = carOrderInfo.CreatedAt

		switch contexts.GetLanguage(ctx) {
		case "zh":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZh)
			break
		case "zh_CN":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneZhCN)
			break
		case "en":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneEn)
			break
		case "ja":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneJa)
			break
		case "ko":
			tmpl, err = template.New("bookingInfo").Parse(consts.TemplateSceneKo)
			break
		default:
			err = gerror.New("语言格式错误")
		}
	}
	if err != nil {
		return
	}
	if req.Type == 1 {
		if err = tmpl.Execute(&buffer, appReservationInfo); err != nil {
			return
		}
	} else if req.Type == 2 {
		if err = tmpl.Execute(&buffer, appStay); err != nil {
			return
		}
	} else if req.Type == 3 {
		if err = tmpl.Execute(&buffer, Property); err != nil {
			return
		}
	} else if req.Type == 4 || req.Type == 5 || req.Type == 6 {
		if err = tmpl.Execute(&buffer, appSceneInfo); err != nil {
			return
		}
	}

	res = new(kefu.AppReservationLanguageTemplateRes)
	res.Template = buffer.String()

	return
}

type AppReservationInfo struct {
	entity.PmsAppReservation
	GuestProfileDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_guest_profile"`
		*entity.PmsGuestProfile
	} `json:"guestProfileDetail" orm:"with:uid=mainGuest"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:uid=puid"`
	RoomTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_room_type"`
		*entity.PmsRoomType
	} `json:"roomTypeDetail" orm:"with:uid=room_type"  dc:"房型信息"`
}

type AppStayInfo struct {
	entity.PmsAppStay
	GuestProfileDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_guest_profile"`
		*entity.PmsGuestProfile
	} `json:"guestProfileDetail" orm:"with:uid=booker"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:uid=puid"`
}

type AppFoodInfo struct {
	entity.FoodOrder
	RestaurantDetail *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		*entity.FoodRestaurant
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=name, where:language='ja'"`
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
}

type AppSpaInfo struct {
	entity.SpaOrder
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:id=property_id"`
}

type AppSceneInfo struct {
	Scene        string      `json:"scene"           dc:"场景"`
	Service      string      `json:"service"           dc:"服务"`
	BookingName  string      `json:"bookingName"           dc:"预订人"`
	BookDatetime *gtime.Time `json:"bookDatetime"           dc:"服务时间"`
	OrderSn      string      `json:"orderSn"           dc:"订单编号"`
	CreatedAt    *gtime.Time `json:"createdAt"           dc:"下单时间"`
}
