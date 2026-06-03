package input_hotel

type PreCreateOrderInp struct {
	Puid              string       `json:"puid" dc:"物业ID"`
	CheckinAt         string       `json:"checkin_at" dc:"入住日期"`
	CheckoutAt        string       `json:"checkout_at" dc:"退房日期"`
	RoomItems         []*RoomItems `json:"roomItems" dc:"房间信息"`
	IsLease           int          `json:"isLease" dc:"是否短租 1、否 2、是"`
	IsCheckDaysNotice bool         `json:"isCheckDaysNotice" dc:"是否检查入住天数   false 检查   true 不检查"`
	PricePlanId       int          `json:"pricePlanId" dc:"价格planID"`
}

type PreMoreCreateOrderInp struct {
	Puid              string       `json:"puid" dc:"物业ID"`
	CheckinAt         string       `json:"checkin_at" dc:"入住日期"`
	CheckoutAt        string       `json:"checkout_at" dc:"退房日期"`
	RoomItems         []*RoomItems `json:"roomItems" dc:"房间信息"`
	IsLease           int          `json:"isLease" dc:"是否短租 1、否 2、是"`
	IsCheckDaysNotice bool         `json:"isCheckDaysNotice" dc:"是否检查入住天数   false 检查   true 不检查"`
}

type RoomItems struct {
	RoomId      string  `json:"room_id" dc:"房型ID"`
	RoomNoId    string  `json:"room_no_id" dc:"房间号ID"`
	RoomNum     int     `json:"roomNum" dc:"房间数量"`
	Adult       int     `json:"adult" dc:"成人数量"`
	Child       int     `json:"child" dc:"儿童数量"`
	Infant      int     `json:"infant" dc:"婴儿数量"`
	BookingFee  float64 `json:"booking_fee" dc:"预订费"`
	Remark      string  `json:"remark" dc:"备注"`
	PricePlanId int     `json:"pricePlanId" dc:"价格planID"`
}

type PreCreateOrderModel struct {
	PreOrderSn string `json:"preOrderSn" dc:"预定订单号"`
}
