package travel

import (
	"APT/internal/model/input/input_travel"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta      `path:"/travel/orderList" method:"post" tags:"APP_TRAVEL" summary:"[一日游]_订单列表"`
	PageNum     int    `p:"pageNum" v:"required#page_number_unknown" dc:"页码"`
	PageSize    int    `p:"pageSize" v:"required#page_number_unknown" dc:"页数"`
	OrderStatus string `p:"orderStatus" v:"in:WAIT_PAY,ORDER_SUCCESS,CANCEL,WAIT_START,DONE,INVALID#order_status_format_error" dc:"订单状态【WAIT_PAY：待支付 ORDER_SUCCESS：预定成功  CANCEL：已取消  WAIT_START：即将开始  DONE：过往订单  INVALID：已失效    全部的话传空值或不传此字段】"`
}

type OrderListRes struct {
	List  []*OrderItem `json:"list" dc:"数据列表"`
	Count int          `json:"count"`
}

type OrderItem struct {
	Id              int     `json:"id" dc:"订单ID"`
	ProductTitle    string  `json:"productTitle" dc:"产品名称"`
	SkuName         string  `json:"skuName" dc:"车型名称"`
	MeetingPlace    string  `json:"meetingPlace" dc:"集合地点"`
	BookDate        string  `json:"bookDate" dc:"预约日期"`
	MeetingTime     string  `json:"meetingTime" dc:"集合时间（格式：HH:MM）"`
	BookingNum      uint    `json:"bookingNum"            dc:"预约人数"`
	ExpirationTime  int     `json:"expiration_time" dc:"订单过期时间"`
	OrderStatus     string  `json:"orderStatus" dc:"订单状态_WAIT_PAY、待支付 WAIT_VERIFY、已支付待核销 DONE、已完成 CANCEL、已取消 REFUND、已退款 OVERDUE、已逾期"`
	OrderAmount     float64 `json:"order_amount" dc:"订单金额"`
	OrderSn         string  `json:"orderSn" dc:"订单号"`
	CreateOrderTime string  `json:"createTime" dc:"创建时间"`
	Countdown       int     `json:"countdown" dc:"订单支付倒计时秒"`
	RefundStatus    string  `json:"refundStatus" dc:"退款状态"`
	ActualAmount    float64 `json:"actualAmount" dc:"第三方支付金额"`
	IsCancel        string  `json:"isCancel" dc:"是否可以取消 Y-灵活取消 N-不可取消"`
	IsFx            string  `json:"isFx" dc:"是否是分销订单   Y   是    N   不是"`
}

type OrderDetailReq struct {
	g.Meta  `path:"/travel/orderDetail" method:"post" tags:"APP_TRAVEL" summary:"[一日游]_订单详情"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type OrderDetailRes struct {
	Id              int    `json:"id" dc:"订单id"`
	OrderSn         string `json:"orderSn" dc:"订单号"`
	OrderStatus     string `json:"orderStatus" dc:"订单状态_WAIT_PAY、待支付 WAIT_VERIFY、已支付待核销 DONE、已完成 CANCEL、已取消 REFUND、已退款 OVERDUE、已逾期"`
	CreateOrderTime string `json:"createOrderTime" dc:"创建订单时间"`
	CancelTime      string `json:"cancelTime" dc:"订单未支付取消时间"`
	PayTime         string `json:"payTime" dc:"订单支付时间"`
	Countdown       int    `json:"countdown" dc:"订单支付倒计时秒"`
	IsFx            string `json:"isFx" dc:"是否是分销订单   Y   是    N   否"`
	ProductInfo     struct {
		Id            uint64 `json:"id"              dc:""`
		Title         string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
		SubTitle      string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
		ContactMobile string `json:"contactMobile"   dc:"联系电话"`
		MeetingPlace  string `json:"meetingPlace"    dc:"集合地点"`
		MeetingTime   string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
		GgLat         string `json:"ggLat"           dc:"谷歌纬度"`
		GgLng         string `json:"ggLng"           dc:"谷歌经度"`
	} `json:"productInfo"   dc:"产品"`
	SkuInfo struct {
		Id            uint64 `json:"id"              dc:""`
		Name          string `json:"name"           dc:"名称（默认语言；多语言存 hg_pms_language）"`
		ContactMobile string `json:"contactMobile" dc:"联系电话"`
		MeetingPlace  string `json:"meetingPlace"  dc:"集合地点"`
		MeetingTime   string `json:"meetingTime"   dc:"集合时间（HH:MM）"`
		GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
		GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	} `json:"skuInfo"   dc:"车型"`
	BookInfo struct {
		BookDate      string `json:"bookDate"       dc:"预定日期"`
		BookNum       uint   `json:"bookNum"     dc:"预定人数"`
		BookingName   string `json:"bookingName"    dc:"预订人姓名"`
		BookingMobile string `json:"bookingMobile"  dc:"手机号"`
		PhoneArea     string `json:"phoneArea"      dc:"手机区号"`
		BookingEmail  string `json:"bookingEmail"   dc:"邮箱"`
	} `json:"bookInfo"  dc:"预约信息"`
	MemberId           int                 `json:"memberId" dc:"会员ID"`
	PayModel           int                 `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
	OrderAmount        float64             `json:"orderAmount" dc:"订单金额"`
	CouponAmount       float64             `json:"couponAmount" dc:"优惠券支付金额"`
	BalAmount          float64             `json:"balAmount" dc:"余额支付金额"`
	ActualAmount       float64             `json:"actualAmount" dc:"第三方支付金额"`
	RefundTime         string              `json:"returnTime" dc:"退款时间"`
	CancelFee          float64             `json:"cancelFee" dc:"取消手续费"`
	RefundAmount       float64             `json:"refundAmount" dc:"退款金额(积分退款+实际退款)"`
	RefundActualAmount float64             `json:"refundActualAmount" dc:"实际退款金额"`
	RefundBalAmount    float64             `json:"refundBalAmount" dc:"退款余额"`
	PayAmount          float64             `json:"payAmount" dc:"支付金额"`
	VerifyTime         string              `json:"verifyTime" dc:"核销时间"`
	CanCancel          bool                `json:"canCancel" dc:"是否可以取消"`
	RefundRecordList   []*RefundRecordItem `json:"refundRecordList" dc:"退款记录列表"`
	CancelFreeDate     string              `json:"cancelFreeDate"       dc:"免费取消日期"`
	CancelPolicy       string              `json:"cancelPolicy"       dc:"取消政策"`
}

type RefundRecordItem struct {
	RefundType   string  `json:"refundType"   dc:"退款方式 BAL-积分退款 AMOUNT-金额退款"`
	RefundAmount float64 `json:"refundAmount" dc:"退款金额"`
	ApplyTime    string  `json:"applyTime"    dc:"申请时间"`
}

type MemberVerifyCodeRefreshReq struct {
	g.Meta  `path:"/travel/orderCodeRefresh" method:"get,post" tags:"APP_TRAVEL" summary:"[一日游]_核销码刷新"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type MemberVerifyCodeRefreshRes struct {
	*input_travel.TravelOrderRefreshCodeModel
}
