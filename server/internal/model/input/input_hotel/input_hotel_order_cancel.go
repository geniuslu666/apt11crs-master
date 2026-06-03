package input_hotel

import (
	"APT/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PreRefundIn struct {
	OrderSn string `json:"orderSn" v:"required#order_number_unknown" dc:"订单号"`
}

type PreRefundOut struct {
	PreCancelOrderSn string `json:"preCancelOrderSn" dc:"订单预退款订单号"`
	CancelRate       []*struct {
		Name      string `json:"name"      dc:"规则名"`
		Mode      string `json:"mode"      dc:"规则模式"`
		StartDays int    `json:"startDays" dc:"开始天数"`
		EndDays   int    `json:"endDays"   dc:"结束天数"`
		Rate      int    `json:"rate"      dc:"取消费率"`
		Date      string `json:"date"      dc:"规则解析日期"`
		Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
	} `json:"cancelRate" dc:"取消政策列表"`
	Transaction        []*PreRefundTransaction   `json:"transaction"    dc:"支付流水记录"`
	CancelFee          float64                   `json:"cancelFee"      dc:"取消费用"`
	OutOrderSn         string                    `json:"outOrderSn"     dc:"外部订单号"`
	OrderSn            string                    `json:"orderSn"        dc:"订单号"`
	OrderCreateAt      string                    `json:"orderCreateAt"  dc:"创建订单时间"`
	OrderFee           float64                   `json:"bookerFee"      dc:"订单金额"`
	RefundFee          float64                   `json:"refundFee"      dc:"退款金额"`
	RefundBalance      float64                   `json:"refundBalance"  dc:"退款积分"`
	BalanceAmount      float64                   `json:"balanceAmount"  dc:"积分支付金额"`
	CouponAmount       float64                   `json:"couponAmount"  dc:"优惠券支付金额"`
	ActualAmount       float64                   `json:"actualAmount"  dc:"实付金额"`
	RefundCouponAmount float64                   `json:"refundCouponAmount"  dc:"退款优惠券金额"`
	RefundRecordList   []*RefundRecordDetailItem `json:"refundRecordList" dc:"退款记录列表"`
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

type RefundDetailInp struct {
	OrderSn string `json:"orderSn" v:"required#order_number_miss" dc:"订单号"`
}

type RefundDetailModel struct {
	PreCancelOrderSn string `json:"preCancelOrderSn" dc:"订单预退款订单号"`
	CancelRate       []*struct {
		Mode      string `json:"mode"      dc:"规则模式"`
		StartDays int    `json:"startDays" dc:"开始天数"`
		EndDays   int    `json:"endDays"   dc:"结束天数"`
		Rate      int    `json:"rate"      dc:"取消费率"`
		Date      string `json:"date"      dc:"规则解析日期"`
		Selected  bool   `json:"selected"  dc:"是否对应当前规则"`
	} `json:"cancelRate" dc:"取消政策列表"`
	Transaction       []*entity.PmsTransaction       `json:"transaction"    dc:"支付流水记录"`
	TransactionRefund []*entity.PmsTransactionRefund `json:"transactionRefund"    dc:"支付流水退款记录"`
	CancelFee         float64                        `json:"cancelFee"      dc:"取消费用"`
	OutOrderSn        string                         `json:"outOrderSn"     dc:"外部订单号"`
	OrderSn           string                         `json:"orderSn"        dc:"订单号"`
	OrderCreateAt     string                         `json:"orderCreateAt"  dc:"创建订单时间"`
	OrderFee          float64                        `json:"bookerFee"      dc:"订单金额"`
	BalanceAmount     float64                        `json:"balanceAmount"  dc:"积分支付金额"`
	CouponAmount      float64                        `json:"couponAmount"  dc:"优惠券支付金额"`
	ActualAmount      float64                        `json:"actualAmount"  dc:"实付金额"`
	RefundFee         float64                        `json:"refundFee"      dc:"退款金额"`
	RefundBalance     float64                        `json:"refundBalance"  dc:"退款积分"`
	CancelOrderSn     string                         `json:"cancelOrderSn"  dc:"订单退款订单号"`
	RefundAt          string                         `json:"refundAt"       dc:"退款时间"`
	RefundRecordList  []*RefundRecordDetailItem      `json:"refundRecordList" dc:"退款记录列表"`
}

type PmsAppStayInfo struct {
	*entity.PmsAppStay
	//Transaction []*struct {
	//	g.Meta `orm:"table:hg_pms_transaction"`
	//	*entity.PmsTransaction
	//} `json:"transaction" orm:"with:order_sn=order_sn, where:pay_status='DONE', where:refund_status!='DONE'" dc:"支付流水"`
	TransactionRefund []*struct {
		g.Meta `orm:"table:hg_pms_transaction_refund"`
		*entity.PmsTransactionRefund
	} `json:"transaction_refund" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"支付流水"`
}
