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
)

type OrderRefundSubmitRequest struct {
	CommonRequest
	Method              string `json:"method"`                                                                // 方法名称
	OrigMerchantOrderNo string `json:"orig_merchant_order_no,omitempty" json:"origMerchantOrderNo,omitempty"` // 原商户系统订单号
	MerchantOrderNo     string `json:"merchant_order_no,omitempty" json:"merchantOrderNo,omitempty"`          // 原商户系统订单号
	PriceCurrency       string `json:"price_currency" json:"priceCurrency,omitempty"`                         // 币种
	OrderAmount         string `json:"order_amount" json:"orderAmount,omitempty"`                             // 退款订单金额
	Description         string `json:"description" json:"description,omitempty"`                              // 退款原因
	NotifyUrl           string `json:"notify_url,omitempty"`                                                  // 退款通知的回调地址
}

type OrderRefundSubmitResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Psn  string          `json:"psn"`
	Sign string          `json:"sign"`
	Data RefundOrderData `json:"data,omitempty"`
}

type RefundOrderData struct {
	TransNo     string `json:"trans_no"`
	VatAmount   string `json:"vat_amount"`
	TransFeeC   string `json:"trans_fee_c"`
	TransAmount string `json:"trans_amount"`
	TransStatus int    `json:"trans_status"`
}

func (c *Client) OrderRefundPayCloudSubmit(ctx context.Context, req *OrderRefundSubmitRequest) (orderRefundSubmitResponse *OrderRefundSubmitResponse, err error) {
	var (
		resp   *gclient.Response
		params g.MapStrStr
	)

	req.Charset = "UTF-8"
	req.Format = "JSON"
	req.AppID = c.AppID
	req.SignType = "RSA2"
	req.Version = "1.0"
	req.Timestamp = gtime.Now().TimestampMilliStr()

	req.Method = "order.refund.submit"
	req.MerchantNo = c.MerchantNo
	req.NotifyUrl = c.RefundNotifyUrl
	params = gvar.New(req).MapStrStr()
	if resp, err = c.DoRequest(ctx, params); err != nil {
		return
	}
	defer resp.Close()
	g.Log().Info(ctx, resp.Raw())
	if resp.StatusCode != 200 {
		err = gerror.New(fmt.Sprintf("#failed to refund order: %s", resp.Status))
		return
	}
	if err = gjson.New(resp.ReadAll()).Scan(&orderRefundSubmitResponse); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if orderRefundSubmitResponse.Code != "0" {
		err = gerror.New(fmt.Sprintf("#failed to refund order: %s", orderRefundSubmitResponse.Msg))
		return
	}
	return
}
