package kefu

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/pmsAppReservation/RoomList" method:"get" tags:"NOTIFY_KEFU" summary:"最新订单"`
	input_hotel.PmsAppReservationRoomListInp
}

type OrderListRes struct {
	OrderCount int `json:"orderCount"   dc:"订单总数"`
	input_form.PageRes
	List []*input_hotel.PmsAppReservationRoomListModel `json:"list"   dc:"数据列表"`
}

type AppReservationLanguageTemplateReq struct {
	g.Meta `path:"/appReservation/languageTemplate" method:"post" tags:"NOTIFY_KEFU" summary:"发送订单多语言模板"`
	Id     int64 `json:"id" v:"required#please_select_the_order_id" dc:"ID"`
	Type   int   `json:"type" v:"required#please_select_template_type" dc:"模板类型 1、住宿单 2、支付订单 3、物业咨询 4、餐厅预定 5、按摩服务 6、接送机服务"`
}

type AppReservationLanguageTemplateRes struct {
	Template string `json:"template" dc:"模板"`
}
