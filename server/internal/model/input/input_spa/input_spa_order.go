package input_spa

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SpaOrderViewInp 获取指定餐厅套餐预订单信息
type SpaOrderViewInp struct {
	Id      int64  `json:"id" dc:"id"`
	OrderSn string `json:"orderSn" dc:"系统订单号"`
}

func (in *SpaOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderViewModel struct {
	entity.SpaOrder
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	MemberDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	StoreDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_store"`
		*entity.SpaStore
	} `json:"storeDetail" orm:"with:id=store_id"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:id=property_id"`
	ServiceDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_service"`
		*entity.SpaService
	} `json:"serviceDetail" orm:"with:id=service_id"`
	GoodsDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_service_goods"`
		*entity.SpaServiceGoods
	} `json:"goodsDetail" orm:"with:id=goods_id"`
	TechnicianList []*struct {
		gmeta.Meta `orm:"table:hg_spa_order_technician"`
		*entity.SpaOrderTechnician
		TechnicianDetail *struct {
			gmeta.Meta `orm:"table:hg_spa_technician"`
			Id         int    `json:"id"                   dc:""`
			Name       string `json:"name"                 dc:"技师真实姓名"`
			Nickname   string `json:"nickname"              dc:"技师昵称"`
			Phone      string `json:"phone"                 dc:"手机号"`
			PhoneArea  string `json:"phoneArea"             dc:"手机区号"`
		} `json:"technicianDetail" orm:"with:id=technician_id"`
	} `json:"technicianList" orm:"with:order_id=id"`
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
		gmeta.Meta  `orm:"table:hg_spa_order_log"`
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

// SpaOrderListInp 获取餐厅套餐预订单列表
type SpaOrderListInp struct {
	input_form.PageReq
	MemberId       uint   `json:"memberId"                dc:"用户ID"`
	OrderSn        string `json:"orderSn" dc:"订单编号"`
	TechnicianName string `json:"technicianName" dc:"技师名称"`
	MemberSearch   string `json:"memberSearch"   dc:"客户信息"`
	OrderStatus    string `json:"orderStatus" dc:"订单状态"`
}

func (in *SpaOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderListModel struct {
	entity.SpaOrder
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	SpaIspDetail  *struct {
		gmeta.Meta `orm:"table:hg_spa_isp"`
		Id         int    `json:"id"      orm:"id"    dc:""`
		Name       string `json:"name"    orm:"name"  dc:"服务商名称"`
	} `json:"spaIspDetail" orm:"with:id=isp_id" dc:"服务商信息"`
	PmsMemberMemberNo string `json:"pmsMemberMemberNo"        dc:"会员号"`
	StoreDetail       *struct {
		gmeta.Meta `orm:"table:hg_spa_store"`
		*entity.SpaStore
	} `json:"storeDetail" orm:"with:id=store_id"`
	PropertyDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_property"`
		*entity.PmsProperty
	} `json:"propertyDetail" orm:"with:id=property_id"`
	ServiceDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_service"`
		*entity.SpaService
	} `json:"serviceDetail" orm:"with:id=service_id"`
	GoodsDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_service_goods"`
		*entity.SpaServiceGoods
	} `json:"goodsDetail" orm:"with:id=goods_id"`
	TechnicianList []*struct {
		gmeta.Meta `orm:"table:hg_spa_order_technician"`
		*entity.SpaOrderTechnician
		TechnicianDetail *struct {
			gmeta.Meta `orm:"table:hg_spa_technician"`
			Id         int    `json:"id"                   dc:""`
			Name       string `json:"name"                 dc:"技师真实姓名"`
			Nickname   string `json:"nickname"              dc:"技师昵称"`
			Phone      string `json:"phone"                 dc:"手机号"`
			PhoneArea  string `json:"phoneArea"             dc:"手机区号"`
		} `json:"technicianDetail" orm:"with:id=technician_id"`
	} `json:"technicianList" orm:"with:order_id=id"`
	TransactionDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction"`
		*entity.PmsTransaction
	} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
	TransactionRefundDetail []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
}

// SpaOrderConfirmDisagreeFields 订单确认拒绝字段过滤
type SpaOrderConfirmDisagreeFields struct {
	BookingStatus       string `json:"bookingStatus"   dc:"确认状态"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"      dc:"确认拒绝原因"`
}

// SpaOrderConfirmAgreeFields 订单同意字段过滤
type SpaOrderConfirmAgreeFields struct {
	OrderStatus string      `json:"orderStatus"   dc:"确认状态"`
	ConfirmTime *gtime.Time `json:"confirmTime"        dc:"确认时间"`
}

// SpaOrderConfirmAgreeInp 订单确认同意
type SpaOrderConfirmAgreeInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaOrderConfirmAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderConfirmAgreeModel struct{}

// SpaOrderConfirmDisagreeInp 订单确认拒绝
type SpaOrderConfirmDisagreeInp struct {
	Id                  int    `json:"id" v:"required#id不能为空" dc:"id"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"           dc:"确认拒绝原因"`
}

func (in *SpaOrderConfirmDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderConfirmDisagreeModel struct{}

// SpaOrderTechnicianInp 获取指定餐厅套餐预订单信息
type SpaOrderTechnicianInp struct {
	input_form.PageReq
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaOrderTechnicianInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderTechnicianModel struct {
	Id             int64  `json:"id"                      dc:"id"`
	Nickname       string `json:"nickname"              dc:"昵称"`
	Name           string `json:"name"                dc:"真实姓名"`
	Phone          string `json:"phone"         dc:"手机号"`
	PhoneArea      string `json:"phoneArea"                    dc:"区号"`
	WorkStatus     string `json:"workStatus"             dc:"工作状态"`
	WorkStatusEnum int    `json:"workStatusEnum"       dc:"工作状态"`
}

// SpaOrderDispatchInp 订单调度
type SpaOrderDispatchInp struct {
	Id                 int         `json:"id" v:"required#id不能为空" dc:"id"`
	DispatchStatus     string      `json:"dispatchStatus"      dc:"订单调度状态"`
	DispatchDesc       string      `json:"dispatchDesc"       dc:"订单调度备注"`
	DispatchOperatorId uint        `json:"dispatchOperatorId"  dc:"调度操作人ID"`
	DispatchTime       *gtime.Time `json:"dispatchTime"        dc:"订单调度时间"`
	TechnicianIds      string      `json:"technicianIds"       dc:"技师ID ,分隔"`
}

func (in *SpaOrderDispatchInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderDispatchModel struct{}

type SpaOrderDispatchTechnicianModel struct {
	Id             int     `json:"id"               dc:"Id"`
	SettlementType int     `json:"settlementType"   dc:"服务分成类型 1跟随系统  2自定义"`
	SettlementId   int     `json:"settlementId"     dc:"结算模式ID"`
	SettlementRate float64 `json:"settlementRate"   dc:"服务分成%"`
	SettlementInfo *struct {
		gmeta.Meta `orm:"table:hg_spa_settlement"`
		Id         int64   `json:"id"        dc:"Id"`
		Type       int     `json:"type"      dc:"结算方式  1无需结算 2按周期自动结算"`
		Rate       float64 `json:"rate"      dc:"结算比例"`
		Cycle      int     `json:"cycle"    dc:"结算周期  1每日结算  2每周结算  3每月结算"`
		Cost       string  `json:"cost"            dc:"结算成本控制"`
	} `json:"settlementInfo" orm:"with:id=settlement_id" dc:"结算模式"`
}

type SpaOrderDispatchIspModel struct {
	Id               int     `json:"id"               dc:"Id"`
	SettlementObject string  `json:"settlementObject"      dc:"结算对象"`
	SettlementType   int     `json:"settlementType"   dc:"服务分成类型 1跟随系统  2自定义"`
	SettlementId     int     `json:"settlementId"     dc:"结算模式ID"`
	SettlementRate   float64 `json:"settlementRate"   dc:"服务分成%"`
	SettlementInfo   *struct {
		gmeta.Meta `orm:"table:hg_spa_settlement"`
		Id         int64   `json:"id"        dc:"Id"`
		Type       int     `json:"type"      dc:"结算方式  1无需结算 2按周期自动结算"`
		Rate       float64 `json:"rate"      dc:"结算比例"`
		Cycle      int     `json:"cycle"    dc:"结算周期  1每日结算  2每周结算  3每月结算"`
		Cost       string  `json:"cost"            dc:"结算成本控制"`
	} `json:"settlementInfo" orm:"with:id=settlement_id" dc:"结算模式"`
}

type SpaOrderSettleInfoModel struct {
	IspId                 int     `json:"ispId"           dc:"服务商ID"`
	TechnicianId          int     `json:"technicianId"           dc:"技师ID"`
	TotalOrderNum         int     `json:"totalOrderNum"          dc:"总订单数量"`
	TotalOrderAmount      float64 `json:"totalOrderAmount"       dc:"总订单金额"`
	TotalSettlementAmount float64 `json:"totalSettlementAmount"  dc:"总结算金额"`
}

// SettleSpaOrderTechnicianListInp 获按摩结算技师订单列表
type SettleSpaOrderTechnicianListInp struct {
	input_form.PageReq
	OrderSn           string `json:"orderSn" dc:"订单号"`
	SettlementOrderId uint   `json:"settlementOrderId"      dc:"结算单ID"`
}

func (in *SettleSpaOrderTechnicianListInp) Filter(ctx context.Context) (err error) {
	return
}

type SettleSpaOrderTechnicianListModel struct {
	OrderId           int64       `json:"id"                   dc:"id"`
	TechnicianId      int64       `json:"technicianId"         dc:"技师ID"`
	SettlementRate    float64     `json:"settlementRate"       dc:"结算比例"`
	SettlementStatus  string      `json:"settlementStatus"     dc:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount  float64     `json:"settlementAmount"     dc:"结算金额"`
	SettlementTime    *gtime.Time `json:"settlementTime"       dc:"结算时间"`
	SettlementOrderId uint        `json:"settlementOrderId"    dc:"结算单ID"`
	SettlementType    int         `json:"settlementType"       dc:"结算方式  1无需结算 2按周期自动结算  3手动申请结算"`
	SettlementCycle   int         `json:"settlementCycle"      dc:"结算周期  1每日结算  2每周结算  3每月结算"`
	OrderAmount       float64     `json:"orderAmount"          dc:"单笔订单金额"`
	ActualEndTime     *gtime.Time `json:"actualEndTime"        dc:"实际结束日期时间"`
	OrderDetail       *struct {
		gmeta.Meta  `orm:"table:hg_spa_order"`
		Id          int64   `json:"id"              dc:"Id"`
		OrderSn     string  `json:"orderSn"         dc:"订单编号"`
		OrderAmount float64 `json:"orderAmount"     dc:"订单金额"`
	} `json:"orderDetail" orm:"with:id=order_id" dc:"订单详情"`
}

// SpaOrderGoOutInp 订单出发
type SpaOrderGoOutInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaOrderGoOutInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderGoOutModel struct{}

// SpaOrderGoOutFields 订单出发字段过滤
type SpaOrderGoOutFields struct {
	TechnicianGoTime *gtime.Time `json:"technicianGoTime"    dc:"技师上门出发时间"`
}

// SpaOrderArriveFields 订单客人到店字段过滤
type SpaOrderArriveFields struct {
	MemberArriveTime *gtime.Time `json:"memberArriveTime"    dc:"客人到店时间"`
}

// SpaOrderStartServiceInp 订单开始服务
type SpaOrderStartServiceInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaOrderStartServiceInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderStartServiceModel struct{}

// SpaOrderStartServiceFields 订单开始服务字段过滤
type SpaOrderStartServiceFields struct {
	OrderStatus     string      `json:"orderStatus"         dc:"订单状态"`
	ActualStartTime *gtime.Time `json:"actualStartTime"     dc:"实际开始日期时间"`
}

// SpaOrderEndServiceInp 订单结束服务
type SpaOrderEndServiceInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaOrderEndServiceInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderEndServiceModel struct{}

// SpaOrderEndServiceFields 订单结束服务字段过滤
type SpaOrderEndServiceFields struct {
	OrderStatus   string      `json:"orderStatus"         dc:"订单状态"`
	ActualEndTime *gtime.Time `json:"actualEndTime"     dc:"实际结束日期时间"`
}

// SpaOrderCancelPayInp 订单取消
type SpaOrderCancelPayInp struct {
	Id                int64   `json:"id" v:"required#id不能为空" dc:"id"`
	RefundType        int     `json:"refundType"           dc:"退款方式"`
	RefundMoney       float64 `json:"refundMoney"           dc:"退款金额"`
	AdminCancelReason string  `json:"adminCancelReason"           dc:"取消原因"`
}

func (in *SpaOrderCancelPayInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderCancelPayModel struct{}

// SpaOrderExportInp 获取订单列表
type SpaOrderExportInp struct {
	MemberId       uint   `json:"memberId"                dc:"用户ID"`
	OrderSn        string `json:"orderSn" dc:"订单编号"`
	TechnicianName string `json:"technicianName" dc:"技师名称"`
	MemberSearch   string `json:"memberSearch"   dc:"客户信息"`
	OrderStatus    string `json:"orderStatus" dc:"订单状态"`
}

func (in *SpaOrderExportInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderExportModel struct {
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

// SpaOrderExportListInp 获取列表
type SpaOrderExportListInp struct {
	input_form.PageReq
}

func (in *SpaOrderExportListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderExportListModel struct {
	entity.OrderExport
}

// SpaOrderAbnormalInp 订单异常处理
type SpaOrderAbnormalInp struct {
	Id                 int         `json:"id" v:"required#id不能为空" dc:"id"`
	AbnormalStatus     int         `json:"abnormalStatus"      dc:"订单异常处理状态"`
	AbnormalReason     string      `json:"abnormalReason"       dc:"订单异常处理备注"`
	AbnormalOperatorId uint        `json:"abnormalOperatorId"  dc:"异常处理操作人ID"`
	AbnormalTime       *gtime.Time `json:"abnormalTime"        dc:"订单异常处理时间"`
}

func (in *SpaOrderAbnormalInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaOrderAbnormalModel struct{}
