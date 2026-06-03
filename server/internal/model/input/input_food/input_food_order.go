// Package sysin

package input_food

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// FoodOrderViewInp 获取指定餐厅套餐预订单信息
type FoodOrderViewInp struct {
	Id      int64  `json:"id" dc:"id"`
	OrderSn string `json:"orderSn" dc:"系统订单号"`
}

func (in *FoodOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderViewModel struct {
	entity.FoodOrder
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	GoodsDeleted  bool `json:"goodsDeleted" dc:"套餐是否已删除"`
	MemberDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	ActualOrderStatus string `json:"actualOrderStatus"      dc:"订单状态"`
	RestaurantDetail  *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		*entity.FoodRestaurant
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
	GoodsDetail *struct {
		gmeta.Meta `orm:"table:hg_food_goods"`
		*entity.FoodGoods
	} `json:"goodsDetail" orm:"with:id=goods_id"`
	TransactionDetail []*struct {
		gmeta.Meta       `orm:"table:hg_pms_transaction"`
		Id               int         `json:"id"               orm:"id"                 description:"主键"`
		OrderSn          string      `json:"orderSn"          orm:"order_sn"           description:"订单号"`
		TransactionSn    string      `json:"transactionSn"    orm:"transaction_sn"     description:"支付流水号"`
		PaymentRequestId string      `json:"paymentRequestId" orm:"payment_request_id" description:"第三方支付流水号"`
		PayChannel       string      `json:"payChannel"       orm:"pay_channel"        description:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
		PayType          string      `json:"payType"          orm:"pay_type"           description:"支付方式   BAL 余额"`
		Amount           float64     `json:"amount"           orm:"amount"             description:"总金额"`
		PayParams        string      `json:"payParams"        orm:"pay_params"         description:"支付参数"`
		PriceCurrency    string      `json:"priceCurrency"    orm:"price_currency"     description:"币种"`
		PayAmount        float64     `json:"payAmount"        orm:"pay_amount"         description:"支付金额"`
		PayStatus        string      `json:"payStatus"        orm:"pay_status"         description:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
		PayTime          *gtime.Time `json:"payTime"          orm:"pay_time"           description:"支付时间"`
		ExpiredTime      *gtime.Time `json:"expiredTime"      orm:"expired_time"       description:"过期时间"`
		RefundAmount     float64     `json:"refundAmount"     orm:"refund_amount"      description:"退款金额"`
		RefundStatus     string      `json:"refundStatus"     orm:"refund_status"      description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
		CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
		UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
		DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		gmeta.Meta    `orm:"table:hg_pms_transaction_refund"`
		Id            int         `json:"id"            orm:"id"             description:"主键"`
		OrderSn       string      `json:"orderSn"       orm:"order_sn"       description:"订单号"`
		TransactionSn string      `json:"transactionSn" orm:"transaction_sn" description:"支付流水号"`
		RefundType    string      `json:"refundType"    orm:"refund_type"    description:"'支付方式   BAL 余额'"`
		RefundSn      string      `json:"refundSn"      orm:"refund_sn"      description:"退款流水号"`
		TransNo       string      `json:"transNo"       orm:"trans_no"       description:"退款交易号"`
		RefundAmount  float64     `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
		RefundTime    *gtime.Time `json:"refundTime"    orm:"refund_time"    description:"退款时间"`
		RefundStatus  string      `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
		CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
		UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"退款明细"`
	LogList []*struct {
		gmeta.Meta  `orm:"table:hg_food_order_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		Images      string      `json:"images"      description:"图集"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	} `json:"logList" orm:"with:order_id=id" dc:"服务流程"`
}

// FoodOrderListInp 获取餐厅套餐预订单列表
type FoodOrderListInp struct {
	input_form.PageReq
	MemberId         uint          `json:"memberId"                dc:"用户ID"`
	OrderType        string        `json:"orderType" dc:"订单类型"`
	OrderSn          string        `json:"orderSn" dc:"预约单号"`
	RestaurantName   string        `json:"restaurantName" dc:"餐厅名称"`
	OrderStatus      string        `json:"orderStatus"   dc:"订单付款状态"`
	BookingStatus    string        `json:"bookingStatus" dc:"订单预定状态"`
	DepositPayStatus string        `json:"depositPayStatus"   dc:"定金状态"`
	RemainPayStatus  string        `json:"remainPayStatus"   dc:"尾款状态"`
	ToretaPayStatus  string        `json:"toretaPayStatus"   dc:"toreta订单支付状态"`
	GoodsName        string        `json:"goodsName" dc:"套餐名称"`
	BookDateTime     []*gtime.Time `json:"bookDateTime" dc:"预定时间"`
	CreatedAt        []*gtime.Time `json:"createdAt" dc:"创建时间"`
	BookSort         string        `json:"bookSort"      dc:"预定时间排序"`
	CreateSort       string        `json:"createSort"    dc:"下单时间排序"`
}

func (in *FoodOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderListModel struct {
	entity.FoodOrder
	ActualOrderStatus string `json:"actualOrderStatus"      dc:"订单状态"`
	MemberDeleted     bool   `json:"memberDeleted" dc:"会员是否已删除"`
	GoodsDeleted      bool   `json:"goodsDeleted" dc:"套餐是否已删除"`
	RestaurantDetail  *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		*entity.FoodRestaurant
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
	GoodsDetail *struct {
		gmeta.Meta `orm:"table:hg_food_goods"`
		*entity.FoodGoods
	} `json:"goodsDetail" orm:"with:id=goods_id"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		*entity.PmsMember
	} `json:"memberDetail" orm:"with:id=member_id"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		*entity.PmsTransaction
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

// FoodOrderConfirmDisagreeFields 订单确认拒绝字段过滤
type FoodOrderConfirmDisagreeFields struct {
	BookingStatus       string `json:"bookingStatus"   dc:"确认状态"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"      dc:"确认拒绝原因"`
}

// FoodOrderConfirmAgreeFields 订单同意字段过滤
type FoodOrderConfirmAgreeFields struct {
	BookingStatus   string      `json:"bookingStatus"   dc:"确认状态"`
	BookingTime     *gtime.Time `json:"bookingTime"        dc:"预定时间"`
	BookDate        string      `json:"bookDate"                dc:"预定日期"`
	BookTime        string      `json:"bookTime"                dc:"预定时间"`
	BookDatetime    *gtime.Time `json:"bookDatetime"            dc:"预定日期时间"`
	OldBookDate     string      `json:"oldBookDate"             dc:"原预定日期"`
	OldBookTime     string      `json:"oldBookTime"             dc:"原预定时间"`
	OldBookDatetime *gtime.Time `json:"oldBookDatetime"         dc:"原预定日期时间"`
	ExpirationTime  int         `json:"expirationTime"         dc:"订单过期时间"`
}

// FoodOrderConfirmAgreeInp 订单确认同意
type FoodOrderConfirmAgreeInp struct {
	Id       int64  `json:"id" v:"required#id不能为空" dc:"id"`
	BookDate string `json:"bookDate"                dc:"预定日期"`
	BookTime string `json:"bookTime"                dc:"预定时间"`
}

func (in *FoodOrderConfirmAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderConfirmAgreeModel struct{}

// FoodOrderConfirmDisagreeInp 订单确认拒绝
type FoodOrderConfirmDisagreeInp struct {
	Id                    int64  `json:"id" v:"required#id不能为空" dc:"id"`
	BookingDisagreeReason string `json:"bookingDisagreeReason"           dc:"确认拒绝原因"`
}

func (in *FoodOrderConfirmDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderConfirmDisagreeModel struct{}

// SettleFoodOrderListInp 获取餐厅结算订单列表
type SettleFoodOrderListInp struct {
	input_form.PageReq
	OrderSn           string        `json:"orderSn" dc:"预约单号"`
	CreatedAt         []*gtime.Time `json:"createdAt" dc:"创建时间"`
	SettlementOrderId uint          `json:"settlementOrderId"      dc:"结算单ID"`
}

func (in *SettleFoodOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type SettleFoodOrderListModel struct {
	Id                int64       `json:"id"                     dc:"id"`
	OrderSn           string      `json:"orderSn"                dc:"订单编号"`
	OutOrderSn        string      `json:"outOrderSn"             dc:"三方订单号"`
	RestaurantId      int         `json:"restaurantId"           dc:"餐厅ID"`
	OrderAmount       float64     `json:"orderAmount"            dc:"订单金额"`
	CouponAmount      float64     `json:"couponAmount"           dc:"优惠券抵扣金额"`
	BalAmount         float64     `json:"balAmount"              dc:"积分抵扣金额"`
	VerifyTime        *gtime.Time `json:"verifyTime"             dc:"核销时间"`
	SettlementRate    float64     `json:"settlementRate"         dc:"结算比例"`
	SettlementStatus  string      `json:"settlementStatus"       dc:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount  float64     `json:"settlementAmount"       dc:"结算金额"`
	SettlementTime    *gtime.Time `json:"settlementTime"         dc:"结算时间"`
	SettlementOrderId uint        `json:"settlementOrderId"      dc:"结算单ID"`
}

type FoodOrderSettleInfoModel struct {
	RestaurantId          int     `json:"restaurantId"           dc:"餐厅ID"`
	TotalOrderNum         int     `json:"totalOrderNum"           dc:"总订单数量"`
	TotalOrderAmount      float64 `json:"totalOrderAmount"       dc:"总订单金额"`
	TotalSettlementAmount float64 `json:"totalSettlementAmount"  dc:"总结算金额"`
}

// FoodOrderCancelPayInp 订单取消
type FoodOrderCancelPayInp struct {
	Id                int64   `json:"id" v:"required#id不能为空" dc:"id"`
	RefundType        int     `json:"refundType"           dc:"退款方式"`
	RefundMoney       float64 `json:"refundMoney"           dc:"退款金额"`
	AdminCancelReason string  `json:"adminCancelReason"           dc:"取消原因"`
}

func (in *FoodOrderCancelPayInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderCancelPayModel struct{}

// FoodOrderVerifyViewModel 核销-订单详情
type FoodOrderVerifyViewModel struct {
	entity.FoodOrder
	Member *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FirstName  string `json:"firstName"    orm:"first_name"     dc:"姓"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
	RestaurantDetail *struct {
		gmeta.Meta `orm:"table:hg_food_restaurant"`
		*entity.FoodRestaurant
		NameLanguage     *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=name, where:language='ja'"`
		SettlementDetail *struct {
			gmeta.Meta `orm:"table:hg_food_settlement"`
			*entity.FoodSettlement
		} `json:"settlementDetail" orm:"with:id=settlement_id"`
	} `json:"restaurantDetail" orm:"with:id=restaurant_id"`
	GoodsDetail *struct {
		gmeta.Meta `orm:"table:hg_food_goods"`
		*entity.FoodGoods
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=goods_name, where:language='ja'"`
	} `json:"goodsDetail" orm:"with:id=goods_id"`
}

// FoodOrderExportInp 获取订单列表
type FoodOrderExportInp struct {
	MemberId         uint          `json:"memberId"                dc:"用户ID"`
	OrderType        string        `json:"orderType" dc:"订单类型"`
	OrderSn          string        `json:"orderSn" dc:"预约单号"`
	RestaurantName   string        `json:"restaurantName" dc:"餐厅名称"`
	OrderStatus      string        `json:"orderStatus"   dc:"订单付款状态"`
	BookingStatus    string        `json:"bookingStatus" dc:"订单预定状态"`
	DepositPayStatus string        `json:"depositPayStatus"   dc:"定金状态"`
	RemainPayStatus  string        `json:"remainPayStatus"   dc:"尾款状态"`
	ToretaPayStatus  string        `json:"toretaPayStatus"   dc:"toreta订单支付状态"`
	GoodsName        string        `json:"goodsName" dc:"套餐名称"`
	CreatedAt        []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FoodOrderExportInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderExportModel struct {
	OrderSn                 string  `json:"orderSn" dc:"系统单号"`
	MemberNo                string  `json:"memberNo" dc:"下单会员ID"`
	BookingName             string  `json:"bookingName" dc:"预定人姓名"`
	BookStartTime           string  `json:"bookStartTime" dc:"预定时间"`
	OrderTime               string  `json:"orderTime" dc:"下单时间"`
	TotalAmount             float64 `json:"totalAmount" dc:"订单总金额"`
	RefundTotalAmount       float64 `json:"refundTotalAmount" dc:"退款总金额"`
	NowPayAmount            float64 `json:"nowPayAmount" dc:"实付金额"`
	PointsPayment           float64 `json:"pointsPayment" dc:"积分支付"`
	CouponPayment           float64 `json:"couponPayment" dc:"优惠券支付"`
	PaycloudWechatPay       float64 `json:"paycloudWechatPay" dc:"PaycloudWechat支付"`
	PaycloudAlipayPay       float64 `json:"paycloudAlipayPay" dc:"PaycloudAlipay+支付"`
	PaycloudCreditPay       float64 `json:"paycloudCreditPay" dc:"Paycloud信用卡支付"`
	PaypelPay               float64 `json:"paypelPay" dc:"Paypel支付"`
	PaypelCreditPay         float64 `json:"paypelCreditPay" dc:"Paypel信用卡支付"`
	StripeCreditPay         float64 `json:"stripeCreditPay" dc:"Stripe信用卡支付"`
	MlilifeWeChatMiniPay    float64 `json:"mlilifeWeChatMiniPay" dc:"WeChatMini支付"`
	PointsRefund            float64 `json:"pointsRefund" dc:"积分退款"`
	CouponRefund            float64 `json:"couponRefund" dc:"优惠券退款"`
	PaycloudWechatRefund    float64 `json:"paycloudWechatRefund" dc:"PaycloudWechat退款"`
	PaycloudAlipayRefund    float64 `json:"paycloudAlipayRefund" dc:"PaycloudAlipay+退款"`
	PaycloudCreditRefund    float64 `json:"paycloudCreditRefund" dc:"Paycloud信用卡退款"`
	PaypelCreditRefund      float64 `json:"paypelCreditRefund" dc:"Paypel信用卡退款"`
	PaypelRefund            float64 `json:"paypelRefund" dc:"Paypel信用卡退款"`
	StripeCreditRefund      float64 `json:"stripeCreditRefund" dc:"Stripe信用卡退款"`
	MlilifeWeChatMiniRefund float64 `json:"mlilifeWeChatMiniRefund" dc:"WeChatMini退款"`
	RefundTime              string  `json:"refundTime" dc:"退款/取消时间"`
}

// FoodOrderExportListInp 获取列表
type FoodOrderExportListInp struct {
	input_form.PageReq
}

func (in *FoodOrderExportListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodOrderExportListModel struct {
	entity.OrderExport
}

// FoodsOrderRefreshCodeInp 餐厅订单-刷新code
type FoodsOrderRefreshCodeInp struct {
	Id       int `json:"id" v:"required#order_id_cannot_be_empty" dc:"订单ID"`
	MemberId int `json:"memberId"      dc:"领用人"`
}

func (in *FoodsOrderRefreshCodeInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodsOrderRefreshCodeModel struct {
	Code string `json:"code"        dc:"券码"`
}
