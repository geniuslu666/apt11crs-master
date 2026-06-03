package hotel

import (
	"APT/internal/model/input/input_hotel"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderChangeOptionsReq struct {
	g.Meta  `path:"/changeOrder/options" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_操作列表"`
	OrderId string `json:"orderId" v:"required#order_id_unkonwn" dc:"订单ID"`
}

type OrderChangeOptionsRes struct {
	OrderSn          string `json:"orderSn"          dc:"系统订单号"`
	IsChangeGuest    string `json:"isChangeGuest"    dc:"是否变更入住人信息   Y 是  N 否"`
	IsChangeGuestId  int    `json:"isChangeGuestId"  dc:"入住人信息变更ID"`
	IsChangePeople   string `json:"isChangePeople"   dc:"是否变更入住人数   Y 是  N 否"`
	IsChangePeopleId int    `json:"isChangePeopleId" dc:"入住人数信息变更ID"`
	IsChangeDate     string `json:"isChangeDate"     dc:"是否变更入住日期   Y 是  N 否"`
	IsChangeDateId   int    `json:"isChangeDateId"   dc:"入住日期变更ID"`
	MainGuest        string `json:"mainGuest"        dc:"住宿人编号"`
	MainGuestDetail  *struct {
		g.Meta   `orm:"table:hg_pms_guest_profile"`
		FullName string `json:"fullName"      dc:"全名"`
		Email    string `json:"email"         dc:"电子邮件"`
		Phone    string `json:"phone"         dc:"电话"`
		AreaNo   string `json:"areaNo"        dc:"电话国际区号"`
		Uid      string `json:"uid"           dc:"三方系统 ID"`
	} `json:"mainGuestDetail" orm:"with:uid=mainGuest"`
	CheckinDate      string `json:"checkinDate"      dc:"入住日期"`
	CheckoutDate     string `json:"checkoutDate"     dc:"退房日期"`
	AdultCount       int    `json:"adultCount"       dc:"成人数量"`
	ChildCount       int    `json:"childCount"       dc:"儿童数量"`
	InfantCount      int    `json:"infantCount"      dc:"婴儿数量"`
	CheckinStatus    string `json:"checkinStatus"    dc:"入住状态  before_checkin  在入住之前[显示取消订单，三种操作菜单根据状态显示操作]  checked_in  已入住[显示续住，三种操作都不可变更禁用]  checked_out  已退房[任何按钮都不能操作]"`
	IsNextChange     bool   `json:"isNextChange"     dc:"是否可以继续变更 false   预定人、日期、人数  不可变更 true 可以"`
	IsNextDateChange bool   `json:"isNextDateChange" dc:"是否可以变更日期 false 不可变更日期 true 可以"`
	IsCancelOrder    bool   `json:"isCancelOrder"    dc:"是否可以取消订单 false 不可取消订单 true 可以"`
	IsStayOn         bool   `json:"isStayOn"         dc:"是否可以续住 false 不可续住 true 可以"`
	OrderStatus      string `json:"orderStatus"      dc:"入住日期"`
}

type OrderChangeGuestEditReq struct {
	g.Meta `path:"/changeOrder/guestEdit" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_修改预定人信息"`
	input_hotel.OrderChangeGuestReq
}

type OrderChangeGuestEditRes struct{}

type OrderChangeGuestInfoReq struct {
	g.Meta  `path:"/changeOrder/findGuestInfo" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_查询预定人信息"`
	OrderId int `json:"orderId" v:"required#order_id_unkonwn" dc:"订单ID"`
}

type OrderChangeGuestInfoRes struct {
	AirhostOrderUuid string `json:"airhostOrderUuid" v:"required#airhost_order_id_unkonwn" dc:"airhost订单ID"`
	OrderId          int    `json:"orderId"    dc:"订单ID"`
	FirstName        string `json:"firstName"  dc:"名"`
	LastName         string `json:"lastName"   dc:"姓"`
	Email            string `json:"email"      dc:"电子邮件"`
	Phone            string `json:"phone"      dc:"电话"`
	AreaNo           string `json:"areaNo"     dc:"电话国际区号"`
}

type OrderChangeBookingPeopleInfoReq struct {
	g.Meta  `path:"/changeOrder/bookingPeopleInfo" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_变更入住人信息"`
	OrderId string `json:"orderId" v:"required#order_id_unkonwn" dc:"订单ID"`
}

type OrderChangeBookingPeopleInfoRes struct {
	Adult     int `json:"adult"     dc:"成人数量"`
	Child     int `json:"child"     dc:"儿童数量"`
	Infant    int `json:"infant"    dc:"婴儿数量"`
	Occupancy int `json:"occupancy" dc:"占用"`
}

type OrderChangeBookingPeopleReq struct {
	g.Meta  `path:"/changeOrder/bookingPeople" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_变更入住人数"`
	OrderId int `json:"orderId" v:"required#order_id_unkonwn" dc:"订单ID"`
	Adult   int `json:"adult" dc:"成人数量"`
	Child   int `json:"child" dc:"儿童数量"`
	Infant  int `json:"infant" dc:"婴儿数量"`
}

type OrderChangeBookingPeopleRes struct {
	ChangeOrderSn   string  `json:"changeOrderSn" dc:"变更订单号"`
	TransactionSn   string  `json:"transactionSn" dc:"支付订单号"`
	ChangePrice     float64 `json:"changePrice" dc:"变更价格"`
	OldOrderPrice   float64 `json:"oldOrderPrice" dc:"原订单价格"`
	NewOrderPrice   float64 `json:"orderPrice" dc:"订单价格"`
	CreateOrderTime string  `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int     `json:"countdown"       dc:"订单支付倒计时秒"`
}

type OrderChangeBookingPeopleSubmitReq struct {
	g.Meta        `path:"/changeOrder/bookingPeopleSubmit" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_变更入住人数提交"`
	ChangeOrderSn string `json:"changeOrderSn" v:"required#change_order_number_unknown" dc:"变更订单号"`
}

type OrderChangeBookingPeopleSubmitRes struct {
	ChangeOrderSn   string  `json:"changeOrderSn" dc:"变更订单号"`
	TransactionSn   string  `json:"transactionSn" dc:"支付订单号"`
	ChangePrice     float64 `json:"changePrice" dc:"变更价格"`
	CreateOrderTime string  `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int     `json:"countdown"       dc:"订单支付倒计时秒"`
	IsFx            string  `json:"isFx" dc:"是否分销订单   Y   是    N   否"`
}

type OrderChangeDatePreInfoReq struct {
	g.Meta       `path:"/changeOrder/datePreInfo" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_日期预变更详情"`
	OrderId      int    `json:"orderId" v:"required#order_id_missing" dc:"订单ID"`
	CheckInDate  string `json:"checkInDate" v:"required#check_in_date_missing" dc:"入住日期"`
	CheckOutDate string `json:"checkOutDate" v:"required#departure_date_unknown" dc:"离店日期"`
}

type OrderChangeDatePreInfoRes struct {
	OldCheckInDate  string   `json:"oldCheckInDate"  dc:"原入住日期"`
	OldCheckOutDate string   `json:"oldCheckOutDate" dc:"原离店日期"`
	NewCheckInDate  string   `json:"newCheckInDate"  dc:"新入住日期"`
	NewCheckOutDate string   `json:"newCheckOutDate" dc:"新离店日期"`
	RoomName        string   `json:"roomName"        dc:"房型名称"`
	RoomNumber      int      `json:"roomNumber"      dc:"房间数量"`
	CancelDate      []string `json:"cancelDate"      dc:"免费取消日期提示"`
	OldOrderPrice   float64  `json:"oldOrderPrice"   dc:"原订单价格"`
	NewOrderPrice   float64  `json:"newOrderPrice"   dc:"新订单价格"`
	DiffOrderPrice  float64  `json:"diffOrderPrice"  dc:"订单价格差"`
	TransactionSn   string   `json:"transactionSn"   dc:"支付订单号"`
	CreateOrderTime string   `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int      `json:"countdown"       dc:"订单支付倒计时秒"`
	ChangeOrderSn   string   `json:"changeOrderSn"   dc:"变更订单号"`
	IsCancel        string   `json:"isCancel" dc:"是否可以取消 Y-灵活取消 N-不可取消"`
	CancelRate      []*struct {
		Name      string `json:"name"      dc:"规则名"`
		Mode      string `json:"mode"      dc:"规则模式"`
		StartDays int    `json:"startDays" dc:"开始天数"`
		EndDays   int    `json:"endDays"   dc:"结束天数"`
		Rate      int    `json:"rate"      dc:"取消费率"`
		Date      string `json:"date"      dc:"规则解析日期"`
		Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
	} `json:"cancelRate" dc:"取消政策列表"`
}

type OrderChangeDatePreInfoSubmitReq struct {
	g.Meta        `path:"/changeOrder/datePreInfoSubmit" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_日期预变更提交"`
	ChangeOrderSn string `json:"changeOrderSn" v:"required#change_order_number_unknown" dc:"变更订单号"`
}

type OrderChangeDatePreInfoSubmitRes struct {
	DiffOrderPrice  float64 `json:"diffOrderPrice"  dc:"订单价格差"`
	TransactionSn   string  `json:"transactionSn"   dc:"支付订单号"`
	CreateOrderTime string  `json:"createOrderTime" dc:"创建订单时间"`
	Countdown       int     `json:"countdown"       dc:"订单支付倒计时秒"`
	ChangeOrderSn   string  `json:"changeOrderSn"   dc:"变更订单号"`
}

type OrderChangeStayOnInfoReq struct {
	g.Meta   `path:"/changeOrder/stayOn" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_续住"`
	OrderId  string `json:"orderId" v:"required#order_id_unknown"  dc:"订单ID"`
	NightNum int    `json:"nightNum" v:"required#how_many_more_nights_of_stay_unknown"  dc:"续住几晚"`
}

type OrderChangeStayOnInfoRes struct {
	CheckOutDate  string  `json:"checkOutDate" dc:"离店日期"`
	Cover         string  `json:"cover"     dc:"封面"`
	Name          string  `json:"name"      dc:"房型封面"`
	BasePrice     float64 `json:"basePrice" dc:"预计每晚价格"`
	Inventory     int     `json:"inventory" dc:"库存量"`
	PreOrderSn    string  `json:"preOrderSn" dc:"预下单订单号"`
	MaxStayNoDays int     `json:"maxStayNoDays" dc:"最大可入住天数"`
}

type OrderChangeInfoDetailReq struct {
	g.Meta        `path:"/changeOrder/infoDetail" method:"post" tags:"APP_HOTEL" summary:"酒店订单变更_变更详情"`
	ChangeOrderSn string `json:"changeOrderSn" v:"required#change_order_number_missing" dc:"变更订单号"`
}

type OrderChangeInfoDetailRes struct {
	Id              int         `json:"id"              orm:"id"                description:"主键"`
	ChangeOrderSn   string      `json:"changeOrderSn"   orm:"change_order_sn"   description:"变更订单号"`
	OrderId         int         `json:"orderId"         orm:"order_id"          description:"订单ID"`
	OrderSn         string      `json:"orderSn"         orm:"order_sn"          description:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      orm:"out_order_sn"      description:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      orm:"change_type"       description:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  string      `json:"oldCheckinDate"  orm:"old_checkin_date"  description:"入住日期"`
	OldCheckoutDate string      `json:"oldCheckoutDate" orm:"old_checkout_date" description:"退房日期"`
	NewCheckinDate  string      `json:"newCheckinDate"  orm:"new_checkin_date"  description:"入住日期"`
	NewCheckoutDate string      `json:"newCheckoutDate" orm:"new_checkout_date" description:"退房日期"`
	OldMainGuest    *gjson.Json `json:"oldMainGuest"    orm:"old_main_guest"    description:"住宿人编号"`
	NewMainGuest    *gjson.Json `json:"newMainGuest"    orm:"new_main_guest"    description:"住宿人编号"`
	OldAdultCount   int         `json:"oldAdultCount"   orm:"old_adult_count"   description:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   orm:"new_adult_count"   description:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   orm:"old_child_count"   description:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   orm:"new_child_count"   description:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  orm:"old_infant_count"  description:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  orm:"new_infant_count"  description:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    orm:"change_status"     description:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    orm:"change_amount"     description:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      orm:"submit_date"       description:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        orm:"done_date"         description:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   orm:"old_order_price"   description:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   orm:"new_order_price"   description:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  orm:"expiration_time"   description:"订单过期时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""`
	Countdown       int         `json:"countdown"       dc:"订单支付倒计时秒"`
	IsFx            string      `json:"isFx"            dc:"是否分销订单   Y   是    N   否"`
}
