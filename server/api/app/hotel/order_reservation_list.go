package hotel

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type ReservationListReq struct {
	g.Meta   `path:"/home/reserveList" method:"post" tags:"APP_HOTEL" summary:"行程_列表"`
	Status   int `json:"status"  dc:"1、即将入住2、过往订单3、失效订单"`
	Page     int `json:"page" v:"required#page_number_miss" dc:"页码"`
	PageSize int `json:"page_size" v:"required#page_length_miss" dc:"页长"`
}

type ReservationListRes struct {
	List []*struct {
		Id                 int     `json:"id"              dc:"主键"`
		Uid                string  `json:"uid"             dc:"三方系统 ID"`
		MemberId           int     `json:"memberId"        dc:"用户ID"`
		Puid               string  `json:"puid"            dc:"物业ID"`
		OrderSn            string  `json:"orderSn"         dc:"系统订单号"`
		OutOrderSn         string  `json:"outOrderSn"      dc:"三方订单号"`
		RoomType           string  `json:"roomType"        dc:"房型信息，参考房型uid"`
		RoomUnit           string  `json:"roomUnit"        dc:"房间单元的id, uid或组合"`
		RatePlanId         string  `json:"ratePlanId"      dc:"费率ID"`
		CheckinDate        string  `json:"checkinDate"     dc:"入住日期"`
		CheckoutDate       string  `json:"checkoutDate"    dc:"退房日期"`
		CheckinTime        string  `json:"checkinTime"     dc:"入住时间，24小时格式"`
		CheckoutTime       string  `json:"checkoutTime"    dc:"退房时间，24小时格式"`
		Status             string  `json:"status"          dc:"预订状态（确认/confirmed、取消/cancelled）"`
		CheckinStatus      string  `json:"checkinStatus"   dc:"入住状态（在入住之前、已入住、已退房）"`
		OrderStatus        string  `json:"orderStatus"     dc:"WAIT_PAY、待支付 CANCEL、支付过期 HAVE_PAID、支付成功"`
		MainGuest          string  `json:"mainGuest"       dc:"住宿人编号"`
		AdultCount         int     `json:"adultCount"      dc:"成人数量"`
		ChildCount         int     `json:"childCount"      dc:"儿童数量"`
		InfantCount        int     `json:"infantCount"     dc:"婴儿数量"`
		BookingFee         float64 `json:"bookingFee"      dc:"预订费"`
		ChannelFee         float64 `json:"channelFee"      dc:"渠道费"`
		CleaningFee        float64 `json:"cleaningFee"     dc:"清洁费"`
		CancellationFee    float64 `json:"cancellationFee" dc:"取消费，仅在取消时适用"`
		Charges            string  `json:"charges"         dc:"费用详情，参考Charge对象"`
		GuestRemarks       string  `json:"guestRemarks"    dc:"备注"`
		Days               int     `json:"days"            dc:"入住时长按天计算"`
		GuestProfileDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_guest_profile"`
			*entity.PmsGuestProfile
		} `json:"guest_profile_detail" orm:"with:uid=mainGuest"`
		RoomTypeDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_type"`
			*entity.PmsRoomType
		} `json:"room_type_detail" orm:"with:uid=room_type"`
		RoomUnitDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_unit"`
			*entity.PmsRoomType
		} `json:"room_unit_detail" orm:"with:uid=room_unit"`
		PropertyDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_property"`
			*entity.PmsProperty
		} `json:"property_detail" orm:"with:uid=puid"`
	}
	Count int `json:"count"`
}

type ReservationDetailReq struct {
	g.Meta `path:"/home/reserveDetail" method:"post" tags:"APP_HOTEL" summary:"行程_详情"`
	Id     int `json:"id" v:"required#order_id_missing" dc:"订单ID"`
}

type ReservationDetailRes struct {
	Id               int                     `json:"id" dc:"订单ID"`
	OrderSn          string                  `json:"orderSn"       dc:"订单号"`
	Days             int                     `json:"days"          dc:"入住时长"`
	CreatedAt        *gtime.Time             `json:"createdAt"     dc:"创建时间"`
	CheckinDate      string                  `json:"checkinDate"   dc:"入住日期"`
	CheckoutDate     string                  `json:"checkoutDate"  dc:"退房日期"`
	CheckoutDateTime string                  `json:"checkoutDateTime"  dc:"退房日期时间"`
	CheckinWeek      string                  `json:"checkinWeek"   dc:"周几入住"`
	CheckoutWeek     string                  `json:"checkoutWeek"  dc:"周几退房"`
	CheckinTime      string                  `json:"checkinTime"   dc:"入住时间，24小时格式"`
	CheckoutTime     string                  `json:"checkoutTime"  dc:"退房时间，24小时格式"`
	RoomTypeName     string                  `json:"roomTypeName"  dc:"房型名称"`
	RoomUnitName     string                  `json:"roomUnitName"  dc:"房间单元名称"`
	RoomNum          int                     `json:"roomNum"       dc:"房间数量"`
	Status           string                  `json:"status"       dc:"订单状态"`
	OrderStatus      string                  `json:"orderStatus"   dc:"订单状态"`
	CheckinStatus    string                  `json:"checkinStatus"   dc:"入住状态（在入住之前、已入住、已退房）"`
	Guest            *entity.PmsGuestProfile `json:"guest"      dc:"入住人"`
	Property         *ReservationDetailProperty
	JpPropertyName   string `json:"jpPropertyName" dc:"日语物业名"`
	JpAddress        string `json:"jpAddress" dc:"日语地址"`
	ChangeDataList   []*struct {
		*entity.PmsAppReservationChange
		Title string `json:"title"`
	}
}

type ReservationDetailProperty struct {
	Name         string `json:"propertyName"  dc:"物业名称"`
	Address      string `json:"address"       dc:"地址"`
	RequiredBook string `json:"requiredBook"  dc:"订房必读"`
	ContactPhone string `json:"contactPhone"  dc:"联系电话"`
	ContactEmail string `json:"contactEmail"  dc:"联系邮箱"`
	GgLat        string `json:"ggLat"         dc:"谷歌纬度"`
	GgLng        string `json:"ggLng"         dc:"谷歌经度"`
}

type FindReservationRoomNoReq struct {
	g.Meta `path:"/home/findReservationRoomNo" method:"post" tags:"APP_HOTEL" summary:"行程_查询房号"`
	Id     int `json:"id" v:"required#order_id_missing" dc:"订单ID"`
}

type FindReservationRoomNoRes struct {
	RoomNo string `json:"roomNo,omitempty" dc:"房号"`
}
