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

type CreateWxMiniPayOrderParams struct {
	CommonRequest
	Method          string  `json:"method"`                  // 方法名称
	PayMethodID     string  `json:"pay_method_id,omitempty"` // 支付方式ID
	MerchantOrderNo string  `json:"merchant_order_no"`       // 商户系统订单号
	PriceCurrency   string  `json:"price_currency"`          // 标价币种
	OrderAmount     float64 `json:"order_amount"`            // 订单金额
	Expires         int     `json:"expires"`                 // 交易有效时间数
	Description     string  `json:"description"`             // 商品或服务描述
	Attach          string  `json:"attach,omitempty"`        // 附加数据
	SubAppid        string  `json:"sub_appid,omitempty"`     // 子应用ID
	SubOpenid       string  `json:"sub_openid,omitempty"`    // 微信用户ID
	NotifyURL       string  `json:"notify_url,omitempty"`    // 支付通知的回调地址
	ExtParams       string  `json:"ext_params,omitempty"`    // 扩展参数
}

type CreateWxMiniPayOrderResponse struct {
	Code string `json:"code,omitempty"`
	Data struct {
		TransNo   string `json:"trans_no"`
		PayParams struct {
			TimeStamp string `json:"timeStamp"`
			NonceStr  string `json:"nonceStr"`
			Package   string `json:"package"`
			SignType  string `json:"signType"`
			PaySign   string `json:"paySign"`
		} `json:"pay_params"`
		PayOriginalRequest  interface{} `json:"pay_original_request"`
		PayOriginalResponse interface{} `json:"pay_original_response"`
	} `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Psn  string `json:"psn,omitempty"`
	Sign string `json:"sign"`
}

func (c *Client) WxMiniOrder(ctx context.Context, req *CreateWxMiniPayOrderParams) (CreateAppOrder CreateWxMiniPayOrderResponse, err error) {
	var (
		resp               *gclient.Response
		params             g.MapStrStr
		CreateAppOrderJson *gjson.Json
	)
	req.SubAppid = c.SubWxMiniAppid
	req.Charset = "UTF-8"
	req.Format = "JSON"
	req.AppID = c.AppID
	req.SignType = "RSA2"
	req.Version = "1.0"
	req.Timestamp = gtime.Now().TimestampMilliStr()
	req.MerchantNo = c.MerchantNo
	req.StoreNo = c.WxMiniStoreNo
	req.PayMethodID = "WeChatPay"

	req.Method = "pay.miniprogram.order"

	req.PriceCurrency = "JPY"
	req.Expires = 60 * 15
	req.Description = "Apartment 11"
	req.NotifyURL = c.NotifyUrl
	params = gvar.New(req).MapStrStr()
	if resp, err = c.DoRequest(ctx, params); err != nil {
		return
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to query order: %s", resp.Status)
		return
	}
	result := gstr.Replace(resp.ReadAllString(), `\\\"`, `\\"`)
	result = gstr.Replace(result, `\\"`, `\"`)
	result = gstr.Replace(result, `\"`, `"`)
	result = gstr.Replace(result, `"{`, `{`)
	result = gstr.Replace(result, `}"`, `}`)
	result = gstr.Replace(result, `version="1.0"`, `version=\"1.0\"`)
	result = gstr.Replace(result, `encoding="UTF-8"`, `encoding=\"UTF-8\"`)
	result = gstr.Replace(result, `standalone="no"`, `standalone=\"no\"`)
	if CreateAppOrderJson, err = gjson.DecodeToJson(result); err != nil {
		return
	}
	if err = CreateAppOrderJson.Scan(&CreateAppOrder); err != nil {
		return
	}
	if CreateAppOrder.Code != "0" {
		err = gerror.New(CreateAppOrder.Msg)
		return
	}
	return
}
