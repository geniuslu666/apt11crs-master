package hotel

import (
	"APT/internal/model/input/input_hotel"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta      `path:"/home/OrderList" method:"post" tags:"APP_HOTEL" summary:"酒店订单_订单列表"`
	PageNum     int    `p:"pageNum" v:"required#page_number_unknown" dc:"页码"`
	PageSize    int    `p:"pageSize" v:"required#page_number_unknown" dc:"页数"`
	OrderStatus string `p:"orderStatus" v:"in:WAIT_PAY,HAVE_PAID,CANCEL#order_status_format_error" dc:"订单状态【WAIT_PAY、待支付 HAVE_PAID、已支付 CANCEL、已取消】"`
}

type OrderListRes struct {
	List  []*OrderItem
	Count int `json:"count"`
}

type OrderItem struct {
	Id              int     `json:"id" dc:"订单ID"`
	PropertyName    string  `json:"propertyName" dc:"物业名称"`
	Address         string  `jsonn:"address" dc:"地址"`
	ExpirationTime  int     `json:"expiration_time" dc:"订单过期时间"`
	OrderStatus     string  `json:"orderStatus" dc:"订单状态___WAIT_PAY、待支付 HAVE_PAID、已支付 CANCEL、已取消"`
	OrderAmount     float64 `json:"order_amount" dc:"订单金额"`
	OrderSn         string  `json:"orderSn" dc:"订单号"`
	RoomNum         int     `json:"roomNum" dc:"房间数量"`
	CheckInDate     string  `json:"checkInDate" dc:"入住时间"`
	CheckOutDate    string  `json:"checkOutDate" dc:"退房时间"`
	Days            int     `json:"days" dc:"入住市场按天计算"`
	AdultCount      float64 `json:"adult_count" dc:"成人数量"`
	ChildCount      float64 `json:"child_count" dc:"儿童数量"`
	CreateOrderTime string  `json:"createTime" dc:"创建时间"`
	Countdown       int     `json:"countdown" dc:"订单支付倒计时秒"`
	RefundStatus    string  `json:"refundStatus" dc:"退款状态"`
	ThirePayAmount  float64 `json:"thirePayAmount" dc:"第三方支付金额"`
	IsCancel        string  `json:"isCancel" dc:"是否可以取消 Y-灵活取消 N-不可取消"`
	IsFx            string  `json:"isFx" dc:"是否是分销订单   Y   是    N   不是"`
}

type OrderDetailReq struct {
	g.Meta  `path:"/home/OrderDetail" method:"post" tags:"APP_HOTEL" summary:"酒店订单_订单详情"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type OrderDetailRes struct {
	Id              int    `json:"id" dc:"订单id"`
	OrderSn         string `json:"orderSn" dc:"订单号"`
	CreateOrderTime string `json:"createOrderTime" dc:"创建订单时间"`
	CancelTime      string `json:"cancelTime" dc:"订单未支付取消时间"`
	Countdown       int    `json:"countdown" dc:"订单支付倒计时秒"`
	CheckinStatus   string `json:"checkinStatus" dc:"入住状态 "`
	IsFx            string `json:"isFx" dc:"是否是分销订单   Y   是    N   否"`
	Property        struct {
		Id           int    `json:"id"                   dc:"主键"`
		Uid          string `json:"uid"                  dc:"在API合作伙伴系统中的物业ID"`
		Cover        string `json:"cover"            dc:"封面"`
		Name         string `json:"name"             dc:"物业名称"`
		Address      string `json:"address"          dc:"地址描述"`
		RequiredBook string `json:"requiredBook"     dc:"订房必读"`
		GgLat        string `json:"ggLat"            dc:"谷歌纬度"`
		GgLng        string `json:"ggLng"            dc:"谷歌经度"`
	}
	OrderDate struct {
		StartDate string `json:"startDate" dc:"入住日期"`
		EndDate   string `json:"endDate"   dc:"退房日期"`
		Days      int    `json:"days"      dc:"总天数"`
	}
	User struct {
		FullName  string `json:"fullName"      dc:"全名"`
		Phone     string `json:"phone"         dc:"手机号"`
		PhoneArea string `json:"phoneArea"     dc:"手机区号"`
		Mail      string `json:"mail"          dc:"邮箱"`
	}
	MemberId  int `json:"memberId" dc:"会员ID"`
	RatePlans []*input_hotel.OrderRooms
	PayInfo   struct {
		PayModel     int     `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
		AllAmount    float64 `json:"allAmount"     dc:"订单总金额"`
		ChangeAmount float64 `json:"changeAmount"     dc:"变更金额"`
		OrderStatus  string  `json:"orderStatus" dc:"订单支付状态 WAIT_PAY-待支付  HAVE_PAID-已支付  CANCEL-取消支付"`
		RefundStatus string  `json:"refundStatus" dc:"退款状态 WAIT 等待退款、DONE 完成退款、PART 部分退款"`
		Balance      struct {
			BalanceAmount     float64 `json:"amount" dc:"余额支付金额"`
			BalancePayOrderSn string  `json:"payOrderSn" dc:"余额支付订单号"`
			PayStatus         string  `json:"payStatus" dc:"支付状态 WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
			RefundStatus      string  `json:"refundStatus" dc:"退款状态 WAIT 等待退款、DONE 完成退款、CANCEL 取消退款"`
			RefundAmount      float64 `json:"refundAmount" dc:"退款金额"`
		}
		ThirdPay struct {
			ThirdAmount     float64 `json:"amount" dc:"外部支付金额"`
			ThirdPayOrderSn string  `json:"payOrderSn" dc:"外部支付订单号"`
			PayStatus       string  `json:"payStatus" dc:"支付状态 WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
			RefundStatus    string  `json:"refundStatus" dc:"退款状态 WAIT 等待退款、DONE 完成退款、CANCEL 取消退款"`
			RefundAmount    float64 `json:"refundAmount" dc:"退款金额"`
		}
		Coupon struct {
			CouponAmount     float64 `json:"amount" dc:"优惠券支付金额"`
			CouponPayOrderSn string  `json:"payOrderSn" dc:"优惠券支付订单号"`
			PayStatus        string  `json:"payStatus" dc:"支付状态 WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
			RefundStatus     string  `json:"refundStatus" dc:"退款状态 WAIT 等待退款、DONE 完成退款、CANCEL 取消退款"`
			RefundAmount     float64 `json:"refundAmount" dc:"退款金额"`
		}
	}
	PayAmount    float64                        `json:"payAmount" dc:"支付金额"`
	RefundDetail *input_hotel.RefundDetailModel `json:"refundDetail" dc:"订单退款详情"`
	CancelRate   []*struct {
		Name      string `json:"name"      dc:"规则名"`
		Mode      string `json:"mode"      dc:"规则模式"`
		StartDays int    `json:"startDays" dc:"开始天数"`
		EndDays   int    `json:"endDays"   dc:"结束天数"`
		Rate      int    `json:"rate"      dc:"取消费率"`
		Date      string `json:"date"      dc:"规则解析日期"`
		Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
	}
	RefundRecordList []*RefundRecordItem `json:"refundRecordList" dc:"退款记录列表"`
}

type RefundRecordItem struct {
	RefundType   string  `json:"refundType"   dc:"退款方式 BAL-积分退款 AMOUNT-金额退款"`
	RefundAmount float64 `json:"refundAmount" dc:"退款金额"`
	ApplyTime    string  `json:"applyTime"    dc:"申请时间"`
}
