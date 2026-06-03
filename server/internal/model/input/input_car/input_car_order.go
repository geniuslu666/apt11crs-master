package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CarOrderViewInp 获取指定餐厅套餐预订单信息
type CarOrderViewInp struct {
	Id      int64  `json:"id" dc:"id"`
	OrderSn string `json:"orderSn" dc:"系统订单号"`
}

func (in *CarOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderViewModel struct {
	entity.CarOrder
	CanRefund     bool `json:"canRefund" dc:"是否可退款"`
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	DriverDeleted bool `json:"driverDeleted" dc:"司机是否已删除"`
	MemberDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	StartServiceAddress *struct {
		gmeta.Meta `orm:"table:hg_car_address"`
		*entity.CarAddress
	} `json:"startServiceAddress" orm:"with:id=start_address_id"  dc:"出发地"`
	EndServiceAddress *struct {
		gmeta.Meta `orm:"table:hg_car_address"`
		*entity.CarAddress
	} `json:"endServiceAddress" orm:"with:id=end_address_id"  dc:"目的地"`
	ServiceDetail *struct {
		gmeta.Meta `orm:"table:hg_car_service"`
		*entity.CarService
		CarTypeDetail *struct {
			gmeta.Meta `orm:"table:hg_car_car_type"`
			*entity.CarCarType
		} `json:"carTypeDetail" orm:"with:id=car_type_id"`
	} `json:"serviceDetail" orm:"with:id=service_id"`
	DriverDetail *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
	CarDetail *struct {
		gmeta.Meta `orm:"table:hg_car_car"`
		*entity.CarCar
		CarTypeDetail *struct {
			gmeta.Meta `orm:"table:hg_car_car_type"`
			*entity.CarCarType
		} `json:"carTypeDetail" orm:"with:id=type_id"`
	} `json:"carDetail" orm:"with:id=car_id"`
	AdminMember *struct {
		gmeta.Meta `orm:"table:hg_admin_member"`
		Username   string `json:"username"           dc:"帐号"`
	} `json:"adminMember" orm:"with:id=dispatch_operator_id"`
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
		gmeta.Meta  `orm:"table:hg_car_order_log"`
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

// CarOrderListInp 获取餐厅套餐预订单列表
type CarOrderListInp struct {
	input_form.PageReq
	MemberId      uint          `json:"memberId"                dc:"用户ID"`
	OrderSn       string        `json:"orderSn" dc:"订单编号"`
	ServiceType   string        `json:"serviceType"          dc:"服务类型"`
	OrderType     string        `json:"orderType" dc:"订单类型"`
	MemberSearch  string        `json:"memberSearch"   dc:"客户信息"`
	OrderStatus   string        `json:"orderStatus" dc:"订单状态"`
	BookStartTime []*gtime.Time `json:"bookStartTime" dc:"预定用车时间"`
	DriverName    string        `json:"driverName" dc:"司机"`
	CreatedAt     []*gtime.Time `json:"createdAt" dc:"创建时间"`
	BookSort      string        `json:"bookSort"      dc:"用车时间排序"`
	CreateSort    string        `json:"createSort"    dc:"下单时间排序"`
}

func (in *CarOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderListModel struct {
	entity.CarOrder
	PmsMemberMemberNo   string `json:"pmsMemberMemberNo"        dc:"会员号"`
	MemberDeleted       bool   `json:"memberDeleted"            dc:"会员是否已删除"`
	DriverDeleted       bool   `json:"driverDeleted"            dc:"司机是否已删除"`
	StartServiceAddress *struct {
		gmeta.Meta `orm:"table:hg_car_address"`
		*entity.CarAddress
	} `json:"startServiceAddress" orm:"with:id=start_address_id"  dc:"出发地"`
	EndServiceAddress *struct {
		gmeta.Meta `orm:"table:hg_car_address"`
		*entity.CarAddress
	} `json:"endServiceAddress" orm:"with:id=end_address_id"  dc:"目的地"`

	DriverDetail *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
	CarDetail *struct {
		gmeta.Meta `orm:"table:hg_car_car"`
		*entity.CarCar
	} `json:"carDetail" orm:"with:id=car_id"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		*entity.PmsTransaction
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

// CarOrderExportInp 获取订单列表
type CarOrderExportInp struct {
	MemberId      uint          `json:"memberId"                dc:"用户ID"`
	OrderSn       string        `json:"orderSn" dc:"订单编号"`
	ServiceType   string        `json:"serviceType"          dc:"服务类型"`
	OrderType     string        `json:"orderType" dc:"订单类型"`
	MemberSearch  string        `json:"memberSearch"   dc:"客户信息"`
	OrderStatus   string        `json:"orderStatus" dc:"订单状态"`
	BookStartTime []*gtime.Time `json:"bookStartTime" dc:"预定用车时间"`
	DriverName    string        `json:"driverName" dc:"司机"`
	CreatedAt     []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *CarOrderExportInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderExportModel struct {
	OrderSn                 string  `json:"orderSn" dc:"系统单号"`
	ServiceType             string  `json:"serviceType" dc:"服务类型"`
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

// CarOrderExportListInp 获取列表
type CarOrderExportListInp struct {
	input_form.PageReq
}

func (in *CarOrderExportListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderExportListModel struct {
	entity.OrderExport
}

// CarOrderConfirmDisagreeFields 订单确认拒绝字段过滤
type CarOrderConfirmDisagreeFields struct {
	BookingStatus       string `json:"bookingStatus"   dc:"确认状态"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"      dc:"确认拒绝原因"`
}

// CarOrderConfirmAgreeFields 订单同意字段过滤
type CarOrderConfirmAgreeFields struct {
	OrderStatus string      `json:"orderStatus"   dc:"确认状态"`
	ConfirmTime *gtime.Time `json:"confirmTime"        dc:"确认时间"`
}

// CarOrderConfirmAgreeInp 订单确认同意
type CarOrderConfirmAgreeInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarOrderConfirmAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderConfirmAgreeModel struct{}

// CarOrderConfirmDisagreeInp 订单确认拒绝
type CarOrderConfirmDisagreeInp struct {
	Id                  int    `json:"id" v:"required#id不能为空" dc:"id"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"           dc:"确认拒绝原因"`
}

func (in *CarOrderConfirmDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderConfirmDisagreeModel struct{}

// CarOrderDriverInp 获取指定司机信息
type CarOrderDriverInp struct {
	input_form.PageReq
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarOrderDriverInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderDriverModel struct {
	Id             int64  `json:"id"                      dc:"id"`
	Nickname       string `json:"nickname"              dc:"昵称"`
	Name           string `json:"name"                dc:"真实姓名"`
	Phone          string `json:"phone"         dc:"手机号"`
	PhoneArea      string `json:"phoneArea"                    dc:"区号"`
	WorkStatus     string `json:"workStatus"       dc:"工作状态"`
	WorkStatusEnum int    `json:"workStatusEnum"       dc:"工作状态"`
	DispatchMode   int    `json:"dispatchMode"       dc:"派单限制"`
}

// CarOrderDispatchInp 订单调度
type CarOrderDispatchInp struct {
	Id                 int         `json:"id" v:"required#id不能为空" dc:"id"`
	IsReturn           int         `json:"isReturn"               dc:"是否退单中 1是  2否"`
	DispatchStatus     string      `json:"dispatchStatus"      dc:"订单调度状态"`
	DispatchDesc       string      `json:"dispatchDesc"       dc:"订单调度备注"`
	DispatchOperatorId uint        `json:"dispatchOperatorId"  dc:"调度操作人ID"`
	DispatchTime       *gtime.Time `json:"dispatchTime"        dc:"订单调度时间"`
	DriverId           int         `json:"driverId"       dc:"司机ID"`
	CarId              int         `json:"carId"       dc:"汽车ID"`
}

func (in *CarOrderDispatchInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderDispatchModel struct{}

// CarOrderTransferInp 订单转单
type CarOrderTransferInp struct {
	Id                 int         `json:"id" v:"required#id不能为空" dc:"id"`
	DispatchDesc       string      `json:"dispatchDesc"       dc:"订单调度备注"`
	DispatchOperatorId uint        `json:"dispatchOperatorId"  dc:"调度操作人ID"`
	DispatchTime       *gtime.Time `json:"dispatchTime"        dc:"订单调度时间"`
	DriverId           int         `json:"driverId"       dc:"司机ID"`
}

func (in *CarOrderTransferInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderTransferModel struct{}

// SettleCarOrderListInp 获取餐厅结算订单列表
type SettleCarOrderListInp struct {
	input_form.PageReq
	OrderSn           string        `json:"orderSn" dc:"预约单号"`
	CreatedAt         []*gtime.Time `json:"createdAt" dc:"创建时间"`
	SettlementOrderId uint          `json:"settlementOrderId"      dc:"结算单ID"`
}

func (in *SettleCarOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type SettleCarOrderListModel struct {
	Id                int64       `json:"id"                     dc:"id"`
	OrderSn           string      `json:"orderSn"                dc:"订单编号"`
	OutOrderSn        string      `json:"outOrderSn"             dc:"三方订单号"`
	DriverId          int         `json:"driverId"               dc:"司机ID"`
	OrderAmount       float64     `json:"orderAmount"            dc:"订单金额"`
	CouponAmount      float64     `json:"couponAmount"           dc:"优惠券抵扣金额"`
	BalAmount         float64     `json:"balAmount"              dc:"积分抵扣金额"`
	ActualEndTime     *gtime.Time `json:"actualEndTime"          dc:"实际结束时间"`
	SettlementRate    float64     `json:"settlementRate"         dc:"结算比例"`
	SettlementStatus  string      `json:"settlementStatus"       dc:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount  float64     `json:"settlementAmount"       dc:"结算金额"`
	SettlementTime    *gtime.Time `json:"settlementTime"         dc:"结算时间"`
	SettlementOrderId uint        `json:"settlementOrderId"      dc:"结算单ID"`
	AbnormalStatus    int         `json:"abnormalStatus"         dc:"订单异常处理状态"`
	AbnormalReason    string      `json:"abnormalReason"         dc:"订单异常处理备注"`
}

type CarOrderSettleInfoModel struct {
	DriverId              int     `json:"driverId"               dc:"司机ID"`
	TotalOrderNum         int     `json:"totalOrderNum"           dc:"总订单数量"`
	TotalOrderAmount      float64 `json:"totalOrderAmount"       dc:"总订单金额"`
	TotalSettlementAmount float64 `json:"totalSettlementAmount"  dc:"总结算金额"`
}

// CarOrderSettlementInp 订单结算
type CarOrderSettlementInp struct {
	Id                int         `json:"id" v:"required#id不能为空" dc:"id"`
	SettlementRate    float64     `json:"settlementRate"         dc:"结算比例"`
	SettlementStatus  string      `json:"settlementStatus"       dc:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount  float64     `json:"settlementAmount"       dc:"结算金额"`
	SettlementTime    *gtime.Time `json:"settlementTime"         dc:"结算时间"`
	SettlementOrderId uint        `json:"settlementOrderId"      dc:"结算单ID"`
	SettlementType    int         `json:"settlementType"         dc:"结算方式  1无需结算 2按周期自动结算  3手动申请结算"`
	SettlementCycle   int         `json:"settlementCycle"        dc:"结算周期  1每日结算  2每周结算  3每月结算"`
}

// CarOrderRefundInp 订单退款
type CarOrderRefundInp struct {
	Id                int64   `json:"id" v:"required#id不能为空" dc:"id"`
	RefundType        int     `json:"refundType"            dc:"退款方式"`
	RefundMoney       float64 `json:"refundMoney"           dc:"退款金额"`
	AdminCancelReason string  `json:"adminCancelReason"     dc:"取消原因"`
	Images            string  `json:"images"                dc:"图片"`
}

func (in *CarOrderRefundInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderRefundModel struct{}

// CarOrderGoOutInp 订单司机出发
type CarOrderGoOutInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarOrderGoOutInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderGoOutModel struct{}

// CarOrderGoOutFields 订单司机出发字段过滤
type CarOrderGoOutFields struct {
	DriverGoTime *gtime.Time `json:"driverGoTime"           description:"司机出发时间"`
}

// CarOrderStartServiceInp 订单开始服务
type CarOrderStartServiceInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarOrderStartServiceInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderStartServiceModel struct{}

// CarOrderStartServiceFields 订单开始服务字段过滤
type CarOrderStartServiceFields struct {
	OrderStatus     string      `json:"orderStatus"         dc:"订单状态"`
	ActualStartTime *gtime.Time `json:"actualStartTime"     dc:"实际开始日期时间"`
}

// CarOrderEndServiceInp 订单结束服务
type CarOrderEndServiceInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarOrderEndServiceInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderEndServiceModel struct{}

// CarOrderEndServiceFields 订单结束服务字段过滤
type CarOrderEndServiceFields struct {
	OrderStatus   string      `json:"orderStatus"         dc:"订单状态"`
	ActualEndTime *gtime.Time `json:"actualEndTime"     dc:"实际结束日期时间"`
}

// CarOrderAbnormalInp 订单异常处理
type CarOrderAbnormalInp struct {
	Id                 int         `json:"id" v:"required#id不能为空" dc:"id"`
	AbnormalStatus     int         `json:"abnormalStatus"      dc:"订单异常处理状态"`
	AbnormalReason     string      `json:"abnormalReason"       dc:"订单异常处理备注"`
	AbnormalOperatorId uint        `json:"abnormalOperatorId"  dc:"异常处理操作人ID"`
	AbnormalTime       *gtime.Time `json:"abnormalTime"        dc:"订单异常处理时间"`
}

func (in *CarOrderAbnormalInp) Filter(ctx context.Context) (err error) {
	return
}

type CarOrderAbnormalModel struct{}
