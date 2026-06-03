package hook

import "github.com/gogf/gf/v2/frame/g"

type StripeWebhookReq struct {
	g.Meta     `path:"/stripe/webhook" method:"post" tags:"NOTIFY_HOOK" summary:"Stripe_支付回调HOOK"`
	ID         string `json:"id"`
	Object     string `json:"object"`
	APIVersion string `json:"api_version"`
	Created    int    `json:"created"`
	Data       struct {
		Object struct {
			ID                  string      `json:"id"`
			Object              string      `json:"object"`
			AfterExpiration     interface{} `json:"after_expiration"`
			AllowPromotionCodes interface{} `json:"allow_promotion_codes"`
			AmountSubtotal      int         `json:"amount_subtotal"`
			AmountTotal         int         `json:"amount_total"`
			AutomaticTax        struct {
				Enabled   bool `json:"enabled"`
				Liability struct {
					Type string `json:"type"`
				} `json:"liability"`
				Status string `json:"status"`
			} `json:"automatic_tax"`
			BillingAddressCollection interface{}   `json:"billing_address_collection"`
			CancelURL                string        `json:"cancel_url"`
			ClientReferenceID        string        `json:"client_reference_id"`
			ClientSecret             interface{}   `json:"client_secret"`
			Consent                  interface{}   `json:"consent"`
			ConsentCollection        interface{}   `json:"consent_collection"`
			Created                  int           `json:"created"`
			Currency                 string        `json:"currency"`
			CurrencyConversion       interface{}   `json:"currency_conversion"`
			CustomFields             []interface{} `json:"custom_fields"`
			CustomText               struct {
				AfterSubmit              interface{} `json:"after_submit"`
				ShippingAddress          interface{} `json:"shipping_address"`
				Submit                   interface{} `json:"submit"`
				TermsOfServiceAcceptance interface{} `json:"terms_of_service_acceptance"`
			} `json:"custom_text"`
			Customer         interface{} `json:"customer"`
			CustomerCreation string      `json:"customer_creation"`
			CustomerDetails  struct {
				Address struct {
					City       interface{} `json:"city"`
					Country    string      `json:"country"`
					Line1      interface{} `json:"line1"`
					Line2      interface{} `json:"line2"`
					PostalCode interface{} `json:"postal_code"`
					State      interface{} `json:"state"`
				} `json:"address"`
				Email     string        `json:"email"`
				Name      string        `json:"name"`
				Phone     interface{}   `json:"phone"`
				TaxExempt string        `json:"tax_exempt"`
				TaxIds    []interface{} `json:"tax_ids"`
			} `json:"customer_details"`
			CustomerEmail   interface{} `json:"customer_email"`
			ExpiresAt       int         `json:"expires_at"`
			Invoice         interface{} `json:"invoice"`
			InvoiceCreation struct {
				Enabled     bool `json:"enabled"`
				InvoiceData struct {
					AccountTaxIds interface{} `json:"account_tax_ids"`
					CustomFields  interface{} `json:"custom_fields"`
					Description   interface{} `json:"description"`
					Footer        interface{} `json:"footer"`
					Issuer        interface{} `json:"issuer"`
					Metadata      struct {
					} `json:"metadata"`
					RenderingOptions interface{} `json:"rendering_options"`
				} `json:"invoice_data"`
			} `json:"invoice_creation"`
			Livemode bool        `json:"livemode"`
			Locale   interface{} `json:"locale"`
			Metadata struct {
			} `json:"metadata"`
			Mode                              string      `json:"mode"`
			PaymentIntent                     string      `json:"payment_intent"`
			PaymentLink                       interface{} `json:"payment_link"`
			PaymentMethodCollection           string      `json:"payment_method_collection"`
			PaymentMethodConfigurationDetails struct {
				ID     string      `json:"id"`
				Parent interface{} `json:"parent"`
			} `json:"payment_method_configuration_details"`
			PaymentMethodOptions struct {
				Card struct {
					RequestThreeDSecure string `json:"request_three_d_secure"`
				} `json:"card"`
			} `json:"payment_method_options"`
			PaymentMethodTypes    []string `json:"payment_method_types"`
			PaymentStatus         string   `json:"payment_status"`
			PhoneNumberCollection struct {
				Enabled bool `json:"enabled"`
			} `json:"phone_number_collection"`
			RecoveredFrom             interface{}   `json:"recovered_from"`
			SavedPaymentMethodOptions interface{}   `json:"saved_payment_method_options"`
			SetupIntent               interface{}   `json:"setup_intent"`
			ShippingAddressCollection interface{}   `json:"shipping_address_collection"`
			ShippingCost              interface{}   `json:"shipping_cost"`
			ShippingDetails           interface{}   `json:"shipping_details"`
			ShippingOptions           []interface{} `json:"shipping_options"`
			Status                    string        `json:"status"`
			SubmitType                interface{}   `json:"submit_type"`
			Subscription              interface{}   `json:"subscription"`
			SuccessURL                string        `json:"success_url"`
			TotalDetails              struct {
				AmountDiscount int `json:"amount_discount"`
				AmountShipping int `json:"amount_shipping"`
				AmountTax      int `json:"amount_tax"`
			} `json:"total_details"`
			UIMode string      `json:"ui_mode"`
			URL    interface{} `json:"url"`
		} `json:"object"`
	} `json:"data"`
	Livemode        bool `json:"livemode"`
	PendingWebhooks int  `json:"pending_webhooks"`
	Request         struct {
		ID             interface{} `json:"id"`
		IdempotencyKey interface{} `json:"idempotency_key"`
	} `json:"request"`
	Type string `json:"type"`
}
type StripeWebhookRes struct{}
type StripeSuccessReq struct {
	g.Meta `path:"/stripe/success" method:"get" tags:"NOTIFY_HOOK" summary:"Stripe_支付成功"`
}
type StripeSuccessRes struct{}
type StripeCancelReq struct {
	g.Meta `path:"/stripe/cancel" method:"get" tags:"NOTIFY_HOOK" summary:"Stripe_支付取消"`
}
type StripeCancelRes struct{}
