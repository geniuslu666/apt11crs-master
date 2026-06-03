package input_hotel

import (
	"APT/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
)

// PreOrderDetailInp 预订单详情入参
type PreOrderDetailInp struct {
	PreOrderSn string `json:"preOrderSn" dc:"预定订单号"`
}

// PreOrderDetailModel 预订单详情内容
type PreOrderDetailModel struct {
	PricePercent int64  `json:"pricePercent" dc:"价格比例"`
	PreOrderSn   string `json:"preOrderSn"    dc:"预订单号"`
	Remark       string `json:"remark"       dc:"备注"`
	IsFx         string `json:"isFx"       dc:"是否是分销"`
	Property     struct {
		Id           int    `json:"id"               dc:"物业ID"`
		Uid          string `json:"uid"              dc:"AirHost物业ID"`
		Cover        string `json:"cover"            dc:"封面"`
		Name         string `json:"name"             dc:"物业名称"`
		Address      string `json:"address"          dc:"地址描述"`
		RequiredBook string `json:"requiredBook"     dc:"入住必读"`
		GgLat        string `json:"ggLat"            dc:"谷歌纬度"`
		GgLng        string `json:"ggLng"            dc:"谷歌经度"`
	} `json:"Property" dc:"订单物业信息"`
	PreDate struct {
		StartDate string `json:"startDate" dc:"入住日期"`
		StartTime string `json:"startTime" dc:"入住时间"`
		EndDate   string `json:"endDate"   dc:"退房日期"`
		EndTime   string `json:"endTime"   dc:"退房时间"`
		Days      int    `json:"days"      dc:"总天数"`
	} `json:"PreDate" dc:"订单入住时间信息"`
	PreUser struct {
		FullName  string `json:"fullName"      dc:"全名"`
		Phone     string `json:"phone"         dc:"手机号"`
		PhoneArea string `json:"phoneArea"     dc:"手机区号"`
		Mail      string `json:"mail"          dc:"邮箱"`
	} `json:"PreUser" dc:"订单预定人信息"`
	MemberId   int                `json:"memberId"    dc:"会员ID"`
	Rooms      []*OrderRooms      `json:"RatePlans"   dc:"房间信息"`
	PayInfo    *OrderPayInfoModel `json:"PayInfo"     dc:"支付信息"`
	CancelDate []string           `json:"cancelDate"  dc:"免费取消日期提示"`
	CancelRate []*struct {
		Name      string `json:"name"      dc:"规则名"`
		Mode      string `json:"mode"      dc:"规则模式"`
		StartDays int    `json:"startDays" dc:"开始天数"`
		EndDays   int    `json:"endDays"   dc:"结束天数"`
		Rate      int    `json:"rate"      dc:"取消费率"`
		Date      string `json:"date"      dc:"规则解析日期"`
		Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
	} `json:"cancelRate" dc:"取消政策列表"`
	PricePlan    *entity.PmsPricePlan `json:"PricePlan"   dc:"价格plan"`
	ChangeAmount float64              `json:"changeAmount" dc:"对比原价变动金额"`
	TotalAmount  float64              `json:"totalAmount" dc:"房间总价"`
}
type OrderPayInfoModel struct {
	PreOrderSn    string  `json:"preOrderSn"    dc:"预订单号"`
	PayModel      int     `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
	MemberBalance float64 `json:"memberBalance" dc:"用户当前余额"`
	AllAmount     float64 `json:"allAmount"     dc:"订单总金额"`
	Score         float64 `json:"score"         dc:"订单可用积分上限"`
	Balance       struct {
		BalanceAmount     float64 `json:"amount"        dc:"余额支付金额"`
		BalancePayOrderSn string  `json:"payOrderSn"    dc:"余额支付订单号"`
		BalanceConfig     *BalanceConfig
	}
	ThirdPay struct {
		ThirdAmount     float64   `json:"amount"     dc:"外部支付金额"`
		ThirdPayOrderSn string    `json:"payOrderSn" dc:"外部支付订单号"`
		ThirdConfig     *gvar.Var `json:"config"     dc:"外部支付配置"`
	}
	Coupon struct {
		CouponId         int     `json:"couponId"      dc:"优惠券ID"`
		CouponAmount     float64 `json:"couponAmount"  dc:"优惠券抵用金额"`
		CouponName       string  `json:"couponName"    dc:"优惠券名称"`
		CouponPayOrderSn string  `json:"couponPayOrderSn" dc:"优惠券支付订单号"`
	}
}
type OrderRooms struct {
	Tid          int                  `json:"tid"           dc:"房型ID"`
	Puid         string               `json:"puid"          dc:"物业UID"`
	RoomId       string               `json:"room_id"      dc:"房型ID"`
	Cover        string               `json:"cover"        dc:"房型封面"`
	Name         string               `json:"name"         dc:"房型名称"`
	RoomNoId     string               `json:"roomNoId"     dc:"房间号ID"`
	RoomNum      int                  `json:"roomNum"      dc:"房间数量"`
	Adult        int                  `json:"adult"        dc:"成人数量"`
	Child        int                  `json:"child"        dc:"儿童数量"`
	Infant       int                  `json:"infant"       dc:"婴儿数量"`
	BookingFee   float64              `json:"booking_fee"  dc:"预订费"`
	Remark       string               `json:"remark"       dc:"备注"`
	Charges      []*Charges           `json:"charges"      dc:"每日价格明细"`
	ChargesCal   []*ChargesCalItem    `json:"chargesCal"   dc:"每日价格明细(按天数)"`
	IsCancel     string               `json:"isCancel"     dc:"是否可取消 Y-灵活取消 N-不可取消"`
	ChangeAmount float64              `json:"changeAmount" dc:"对比原价变动金额"`
	OrderAmount  float64              `json:"totalAmount"  dc:"房间总价"`
	PricePlan    *entity.PmsPricePlan `json:"PricePlan"   dc:"价格plan"`
	CheckinDate  *gtime.Time          `json:"checkinDate"      dc:"入住日期"`
	CheckoutDate *gtime.Time          `json:"checkoutDate"     dc:"退房日期"`
	CheckinTime  string               `json:"checkinTime"      dc:"入住时间，24小时格式"`
	CheckoutTime string               `json:"checkoutTime"     dc:"退房时间，24小时格式"`
}

type Charges struct {
	ID          interface{} `json:"id,omitempty"  dc:"费用明细ID"`
	UID         interface{} `json:"uid,omitempty" dc:"费用明细UID"`
	Date        string      `json:"date"          dc:"费用日期"`
	FeeType     string      `json:"fee_type"      dc:"费用类型"`
	InitAmount  float64     `json:"initAmount"        dc:"初始费用金额"`
	Amount      float64     `json:"amount"        dc:"费用金额"`
	Description interface{} `json:"description"   dc:"费用描述"`
}
type ChargesCalItem struct {
	Date   string  `json:"date"          dc:"费用日期"`
	Amount float64 `json:"amount"        dc:"费用金额"`
}

type BalanceConfig struct {
	ScenePayRate float64 `json:"scenePayRate"  dc:"支付场景费率"`
	Level        int     `json:"level"         dc:"会员等级"`
	LevelName    string  `json:"levelName"     dc:"会员等级名"`
	ExchangeRate float64 `json:"exchangeRate"  dc:"汇率"`
	IsPayOpen    string  `json:"isPayOpen"     dc:"是否积分抵扣"`
}
