package input_travel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// TravelOrderListInp 获取一日游订单列表
type TravelOrderListInp struct {
	input_form.PageReq
	OrderSn      string        `json:"orderSn"     dc:"预约单号"`
	ProductName  string        `json:"productName" dc:"产品名称"`
	SkuName      string        `json:"skuName" dc:"产品sku名称"`
	MemberSearch string        `json:"memberSearch"   dc:"客户信息"`
	OrderStatus  string        `json:"orderStatus" dc:"订单状态（WAIT_PAY/WAIT_VERIFY/DONE/CANCEL/REFUND，空=全部）"`
	CreatedAt    []*gtime.Time `json:"createdAt"   dc:"创建时间范围 [开始, 结束]"`
	BookDate     []string      `json:"bookDate"    dc:"预约时间范围 [开始, 结束]"`
	VerifyTime   []*gtime.Time `json:"verifyTime"  dc:"核销时间范围 [开始, 结束]"`
	MemberId     uint          `json:"memberId"      dc:"用户ID"`
}

// TravelOrderListModel 订单列表展示模型
type TravelOrderListModel struct {
	entity.TravelOrder
	PmsMemberMemberNo string `json:"pmsMemberMemberNo"        dc:"会员号"`
	MemberDeleted     bool   `json:"memberDeleted"            dc:"会员是否已删除"`
	ProductInfo       *struct {
		gmeta.Meta `orm:"table:hg_travel_product"`
		*entity.TravelProduct
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
	SkuInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_product_sku"`
		Id         uint64 `json:"id"            dc:""`
		ProductId  uint   `json:"productId"     dc:"产品ID"`
		Name       string `json:"name"          dc:"车型名称（默认语言；多语言存 hg_pms_language）"`
	} `json:"skuInfo" orm:"with:id=sku_id"  dc:"产品sku"`
	VerifyStaffInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_verify_staff"`
		*entity.TravelVerifyStaff
	} `json:"verifyStaffInfo" orm:"with:id=verify_staff_id"  dc:"核销人员"`
}

// TravelOrderViewInp 获取订单详情
type TravelOrderViewInp struct {
	Id int64 `json:"id" v:"required#请选择要查看的订单" dc:"订单ID"`
}

// TravelOrderViewModel 订单详情模型
type TravelOrderViewModel struct {
	entity.TravelOrder
	MemberDeleted bool `json:"memberDeleted" dc:"会员是否已删除"`
	MemberDetail  *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	ProductInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_product"`
		*entity.TravelProduct
	} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
	SkuInfo *struct {
		gmeta.Meta    `orm:"table:hg_travel_product_sku"`
		Id            uint64 `json:"id"            dc:""`
		ProductId     uint   `json:"productId"     dc:"产品ID"`
		Name          string `json:"name"          dc:"车型名称（默认语言；多语言存 hg_pms_language）"`
		MeetingPlace  string `json:"meetingPlace"  dc:"集合地点"`
		MeetingTime   string `json:"meetingTime"   dc:"集合时间（HH:MM）"`
		GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
		GgLng         string `json:"ggLng"         dc:"谷歌经度"`
		ContactMobile string `json:"contactMobile" dc:"联系电话"`
	} `json:"skuInfo" orm:"with:id=sku_id"  dc:"产品sku"`
	VerifyStaffInfo *struct {
		gmeta.Meta `orm:"table:hg_travel_verify_staff"`
		*entity.TravelVerifyStaff
	} `json:"verifyStaffInfo" orm:"with:id=verify_staff_id"  dc:"核销人员"`
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
		gmeta.Meta  `orm:"table:hg_travel_order_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	} `json:"logList" orm:"with:order_id=id" dc:"服务流程"`
}

// TravelOrderRefundInp 订单退款
type TravelOrderRefundInp struct {
	Id                int64   `json:"id"           v:"required#请选择要退款的订单" dc:"订单ID"`
	RefundType        int     `json:"refundType"           dc:"退款方式"`
	RefundMoney       float64 `json:"refundMoney"          dc:"退款金额"`
	AdminCancelReason string  `json:"adminCancelReason"    dc:"取消原因"`
}

// TravelOrderRefreshCodeInp 一日游订单-刷新code
type TravelOrderRefreshCodeInp struct {
	OrderSn  string `json:"orderSn"  dc:"订单号"`
	MemberId int    `json:"memberId"      dc:"领用人"`
}

type TravelOrderRefreshCodeModel struct {
	Code        string `json:"code"        dc:"券码"`
	OrderStatus string `json:"orderStatus" dc:"状态"`
}

// TravelOrderApplyRefundDetailInp 一日游订单-申请退款详情
type TravelOrderApplyRefundDetailInp struct {
	OrderSn string `json:"orderSn"  dc:"订单号"`
}

type TravelOrderApplyRefundDetailModel struct {
	OrderSn       string  `json:"orderSn"        dc:"订单号"`
	OrderCreateAt string  `json:"orderCreateAt"  dc:"创建订单时间"`
	OrderAmount   float64 `json:"orderAmount"    dc:"订单金额"`
	BalanceAmount float64 `json:"balanceAmount"  dc:"积分抵扣金额"`
	ActualAmount  float64 `json:"actualAmount"   dc:"实付金额"`
	CancelFee     float64 `json:"cancelFee"      dc:"取消费用"`
	RefundBalance float64 `json:"refundBalance"  dc:"退款积分"`
	RefundAmount  float64 `json:"refundAmount"   dc:"第三方支付的退款金额"`
	CancelFeeRate float64 `json:"cancelFeeRate"  dc:"取消费用比例（单位：百分比）"`
}

type TravelOrderInfo struct {
	*entity.TravelOrder
	//Transaction []*struct {
	//	g.Meta `orm:"table:hg_pms_transaction"`
	//	*entity.PmsTransaction
	//} `json:"transaction" orm:"with:order_sn=order_sn, where:pay_status='DONE', where:refund_status!='DONE'" dc:"支付流水"`
	TransactionRefund []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transaction_refund" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"退款流水"`
}

type PreRefundTransaction struct {
	TransactionSn string  `json:"transactionSn"       dc:"支付流水号"`
	PayType       string  `json:"payType"             dc:"支付方式   BAL 余额"`
	Amount        float64 `json:"amount"              dc:"总金额"`
	RefundAmount  float64 `json:"refundAmount"        dc:"退款金额"`
	PriceCurrency string  `json:"priceCurrency"       dc:"币种"`
	Refundable    float64 `json:"refundable"          dc:"可退金额"`
	ActualRefund  float64 `json:"actualRefund"        dc:"实际退款金额"`
	PayStatus     string  `json:"payStatus"           dc:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
}

type RefundDetailModel struct {
	OrderSn            string                    `json:"orderSn"        dc:"订单号"`
	OrderCreateAt      string                    `json:"orderCreateAt"  dc:"创建订单时间"`
	OrderAmount        float64                   `json:"orderAmount"      dc:"订单金额"`
	BalanceAmount      float64                   `json:"balanceAmount"  dc:"积分支付金额"`
	CouponAmount       float64                   `json:"couponAmount"  dc:"优惠券支付金额"`
	ActualAmount       float64                   `json:"actualAmount"  dc:"实付金额"`
	CancelFee          float64                   `json:"cancelFee"      dc:"取消费用"`
	RefundAmount       float64                   `json:"refundAmount" dc:"退款金额(积分退款+实际退款)"`
	RefundActualAmount float64                   `json:"refundActualAmount" dc:"实际退款金额"`
	RefundBalance      float64                   `json:"refundBalance"  dc:"退款积分"`
	RefundTime         string                    `json:"refundTime"       dc:"退款时间"`
	RefundRecordList   []*RefundRecordDetailItem `json:"refundRecordList" dc:"退款记录列表"`
	CancelPolicy       string                    `json:"cancelPolicy"      dc:"取消政策"`
}

type RefundRecordDetailItem struct {
	RefundType   string  `json:"refundType"   dc:"退款方式 BAL-积分退款 AMOUNT-金额退款"`
	RefundAmount float64 `json:"refundAmount" dc:"退款金额"`
	ApplyTime    string  `json:"applyTime"    dc:"申请时间"`
	OperateType  string  `json:"operateType"  dc:"退款类型 USER-用户主动申请 ADMIN-后台退款"`
	RefundTime   string  `json:"refundTime"   dc:"到账时间"`
	RefundStatus string  `json:"refundStatus" dc:"退款状态 DONE-已退款 WAIT-退款中"`
	Remark       string  `json:"remark"       dc:"退款说明"`
}
