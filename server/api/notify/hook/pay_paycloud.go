package hook

import (
	"APT/internal/model/input/input_pay"
	"github.com/gogf/gf/v2/frame/g"
)

type DoPayCloudReq struct {
	g.Meta `path:"/paycloud/callback" method:"post" tags:"NOTIFY_HOOK" summary:"Paycloud_支付"`
	input_pay.ThirdPayInp
}

type DoPayCloudRes struct {
	*input_pay.ThirdPayModel
}

type RefundCloudReq struct {
	g.Meta            `path:"/payCloud/refund" method:"post" tags:"NOTIFY_HOOK" summary:"Paycloud_退款"`
	TransEndTime      string `json:"trans_end_time"`
	Charset           string `json:"charset"`
	StoreNo           string `json:"store_no"`
	PayScenario       string `json:"pay_scenario"`
	Sign              string `json:"sign"`
	TransFeeC         string `json:"trans_fee_c"`
	MerchantOrderNo   string `json:"merchant_order_no"`
	DiscountBpc       string `json:"discount_bpc"`
	VatAmount         string `json:"vat_amount"`
	CashbackAmount    string `json:"cashback_amount"`
	OrderAmount       string `json:"order_amount"`
	AppId             string `json:"app_id"`
	SignType          string `json:"sign_type"`
	TransStatus       int    `json:"trans_status"`
	PriceCurrency     string `json:"price_currency"`
	TransType         int    `json:"trans_type"`
	Timestamp         string `json:"timestamp"`
	TransNo           string `json:"trans_no"`
	MerchantNo        string `json:"merchant_no"`
	Method            string `json:"method"`
	PayUserAccountId  string `json:"pay_user_account_id"`
	Format            string `json:"format"`
	PayMethodId       string `json:"pay_method_id"`
	TransAmount       string `json:"trans_amount"`
	HttpRequestId     string `json:"http_request_id"`
	Version           string `json:"version"`
	PayChannelTransNo string `json:"pay_channel_trans_no"`
	PaidAmount        string `json:"paid_amount"`
	DiscountBmopc     string `json:"discount_bmopc"`
}

type RefundCloudRes struct{}
