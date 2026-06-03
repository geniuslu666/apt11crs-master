package hook

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type PaypelPayReturnReq struct {
	g.Meta  `path:"/pay/paypel/return" method:"get" tags:"NOTIFY_HOOK" summary:"Paypel_支付回调RETURN"`
	Token   string `json:"token" v:"required#token_cannot_be_empty" dc:"token"`
	PayerID string `json:"PayerID" v:"required#payerId_cannot_be_empty" dc:"PayerID"`
}

type PaypelPayReturnRes struct {
}

type PaypelPayCancelReq struct {
	g.Meta `path:"/pay/paypel/cancel" method:"get" tags:"NOTIFY_HOOK" summary:"Paypel_支付回调CALCEL"`
}

type PaypelPayCancelRes struct {
}

type PaypelPayHookReq struct {
	g.Meta          `path:"/pay/paypel/hook" method:"post" tags:"NOTIFY_HOOK" summary:"Paypel_支付回调CALCEL"`
	ID              string    `json:"id"`
	EventVersion    string    `json:"event_version"`
	CreateTime      time.Time `json:"create_time"`
	ResourceType    string    `json:"resource_type"`
	ResourceVersion string    `json:"resource_version"`
	EventType       string    `json:"event_type"`
	Summary         string    `json:"summary"`
	Resource        struct {
		Payee struct {
			EmailAddress string `json:"email_address"`
			MerchantID   string `json:"merchant_id"`
		} `json:"payee"`
		Amount struct {
			Value        string `json:"value"`
			CurrencyCode string `json:"currency_code"`
		} `json:"amount"`
		SellerProtection struct {
			DisputeCategories []string `json:"dispute_categories"`
			Status            string   `json:"status"`
		} `json:"seller_protection"`
		SupplementaryData struct {
			RelatedIds struct {
				OrderID string `json:"order_id"`
			} `json:"related_ids"`
		} `json:"supplementary_data"`
		UpdateTime                time.Time `json:"update_time"`
		CreateTime                time.Time `json:"create_time"`
		FinalCapture              bool      `json:"final_capture"`
		SellerReceivableBreakdown struct {
			PaypalFee struct {
				Value        string `json:"value"`
				CurrencyCode string `json:"currency_code"`
			} `json:"paypal_fee"`
			GrossAmount struct {
				Value        string `json:"value"`
				CurrencyCode string `json:"currency_code"`
			} `json:"gross_amount"`
			NetAmount struct {
				Value        string `json:"value"`
				CurrencyCode string `json:"currency_code"`
			} `json:"net_amount"`
		} `json:"seller_receivable_breakdown"`
		Links []struct {
			Method string `json:"method"`
			Rel    string `json:"rel"`
			Href   string `json:"href"`
		} `json:"links"`
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"resource"`
	Links []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}

type PaypelPayHookRes struct {
}
