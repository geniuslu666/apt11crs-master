package paycloud

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

type UnifiedOrderRequest struct {
	CommonRequest
	Method          string  `json:"method"`                      // 方法名称
	PayMethodID     string  `json:"pay_method_id,omitempty"`     // 支付方式ID
	SubPayMethodID  string  `json:"sub_pay_method_id,omitempty"` // 子支付方式ID
	MerchantOrderNo string  `json:"merchant_order_no"`           // 商户系统订单号
	PriceCurrency   string  `json:"price_currency"`              // 标价币种
	OrderAmount     float64 `json:"order_amount"`                // 订单金额
	Expires         int     `json:"expires"`                     // 交易有效时间数
	Description     string  `json:"description"`                 // 商品或服务描述
	Attach          string  `json:"attach,omitempty"`            // 附加数据
	ReturnURL       string  `json:"return_url"`                  // 前端重定向用户页面的URL
	NotifyURL       string  `json:"notify_url,omitempty"`        // 支付通知的回调地址
	TerminalType    string  `json:"terminal_type"`               // 商户服务适用的终端类型
	OSType          string  `json:"os_type,omitempty"`           // 手机操作系统类型
	TermIP          string  `json:"term_ip"`                     // 终端的IP地址
	Language        string  `json:"language,omitempty"`          // 收银台界面默认显示语言
	ExtParams       string  `json:"ext_params,omitempty"`        // 扩展参数
}

type UnifiedOrderResponse struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Psn  string           `json:"psn"`
	Sign string           `json:"sign"`
	Data UnifiedOrderData `json:"data"`
}

type UnifiedOrderDataPayOriginalRequest struct {
	Order struct {
		Env struct {
			OsType       string `json:"osType"`
			TerminalType string `json:"terminalType"`
		} `json:"env"`
		Merchant struct {
			MerchantAddress struct {
				City   string `json:"city"`
				Region string `json:"region"`
			} `json:"merchantAddress"`
			MerchantMCC         string `json:"merchantMCC"`
			MerchantName        string `json:"merchantName"`
			ReferenceMerchantId string `json:"referenceMerchantId"`
		} `json:"merchant"`
		OrderAmount struct {
			Currency string `json:"currency"`
			Value    string `json:"value"`
		} `json:"orderAmount"`
		OrderDescription string `json:"orderDescription"`
		ReferenceOrderId string `json:"referenceOrderId"`
	} `json:"order"`
}

type UnifiedOrderPayOriginalResponse struct {
	AcquirerId    string `json:"acquirerId"`
	PaymentAmount struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"paymentAmount"`
	PaymentId  string `json:"paymentId"`
	PaymentUrl string `json:"paymentUrl"`
	PspId      string `json:"pspId"`
	Result     struct {
		ResultCode    string `json:"resultCode"`
		ResultMessage string `json:"resultMessage"`
		ResultStatus  string `json:"resultStatus"`
	} `json:"result"`
}

type UnifiedOrderData struct {
	PayOriginalRequest  UnifiedOrderDataPayOriginalRequest `json:"payOriginalRequest"`
	PayOriginalResponse UnifiedOrderPayOriginalResponse    `json:"payOriginalResponse"`
	PayUrl              string                             `json:"pay_url"`
}

func (c *Client) UnifiedOrder(ctx context.Context, req *UnifiedOrderRequest) (unifiedOrderResponse *UnifiedOrderResponse, err error) {
	var (
		resp   *gclient.Response
		params g.MapStrStr
	)
	// 默认公共参数
	req.Charset = "UTF-8"
	req.Format = "JSON"
	req.AppID = c.AppID
	req.SignType = "RSA2"
	req.Version = "1.0"
	req.Timestamp = gtime.Now().TimestampMilliStr()
	req.MerchantNo = c.MerchantNo
	req.StoreNo = c.StoreNo

	req.Method = "pay.unifiedorder"

	req.PriceCurrency = "JPY"
	req.Expires = 60 * 15
	req.Description = "Apartment 11"
	req.OSType = "ANDROID"
	req.Language = "zh-CN"
	req.ReturnURL = c.RefundUrlApp
	req.NotifyURL = c.NotifyUrl
	//req.TermIP = "123.56.84.630"
	//req.OrderAmount = 1
	//req.MerchantOrderNo = guid.S([]byte("payCloud"))
	params = gvar.New(req).MapStrStr()
	if resp, err = c.DoRequest(ctx, params); err != nil {
		return
	}
	defer resp.Close()
	g.Log().Info(ctx, resp.Raw())
	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to unifiedorder: %s", resp.Status)
		return
	}
	respContent := resp.ReadAllString()
	respContent = gstr.Replace(respContent, `\\\"`, `\\"`)
	respContent = gstr.Replace(respContent, `\\"`, `\"`)
	respContent = gstr.Replace(respContent, `\"`, `"`)
	respContent = gstr.Replace(respContent, `}"`, `}`)
	respContent = gstr.Replace(respContent, `"{`, `{`)
	respContent = gstr.Replace(respContent, `version="1.0"`, `version=\"1.0\"`)
	respContent = gstr.Replace(respContent, `encoding="UTF-8"`, `encoding=\"UTF-8\"`)
	respContent = gstr.Replace(respContent, `standalone="no"`, `standalone=\"no\"`)
	if err = gjson.New(respContent).Scan(&unifiedOrderResponse); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if unifiedOrderResponse.Code != "0" {
		err = gerror.New(fmt.Sprintf("#failed to unifiedorder: %s", unifiedOrderResponse.Msg))
		return
	}

	return
}
