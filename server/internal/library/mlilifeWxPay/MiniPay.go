package mlilifeWxPay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/guid"
)

type WxappPayRequest struct {
	Pid        string `json:"pid"`          // 必填 - 项目ID
	Sign       string `json:"sign"`         // 必填 - 加密签名
	NonceStr   string `json:"nonce_str"`    // 必填 - 32位随机字符串
	AppId      string `json:"app_id"`       // 必填 - 小程序AppID
	Openid     string `json:"openid"`       // 必填 - 用户OpenID
	Body       string `json:"body"`         // 必填 - 商品描述
	TotalFee   int    `json:"total_fee"`    // 必填 - 支付金额（日元）
	OutTradeNo string `json:"out_trade_no"` // 必填 - 商户订单号
	TimeStart  string `json:"time_start"`   // 可选 - 交易起始时间
	TimeExpire string `json:"time_expire"`  // 可选 - 交易过期时间
	Attach     string `json:"attach"`       // 可选 - 附加数据
	NotifyUrl  string `json:"notify_url"`   // 必填 - 异步通知地址
}

type WxappPayResponse struct {
	S    int         `json:"s"`    // 状态码
	Msg  string      `json:"msg"`  // 响应信息
	Data PaymentData `json:"data"` // 支付参数
}

type PaymentData struct {
	AppId     string `json:"appId"`     // 小程序appId（来自data字段）
	TimeStamp string `json:"timeStamp"` // 时间戳
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"package"`   // 预支付订单ID（格式：prepay_id=xxx）
	SignType  string `json:"signType"`  // 签名类型（MD5）
	PaySign   string `json:"paySign"`   // 支付签名
}

var (
	wxappPayUrl = "https://pay.mlilife.com/v1/api/wxapp_pay"
)

func WxappPay(ctx context.Context, req *WxappPayRequest) (resp *WxappPayResponse, err error) {
	var (
		gClientResponse *gclient.Response
		PayConfig       *model.PayConfig
	)
	// 读取支付配置
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	req.Pid = PayConfig.PayMiniPid
	req.NonceStr = guid.S()
	req.AppId = PayConfig.PayMiniAppID
	req.NotifyUrl = PayConfig.PayMiniNotifyUrl
	if req.Sign, err = GetSign(ctx, req, PayConfig.PayMiniSignKey); err != nil {
		return
	}
	if gClientResponse, err = g.Client().Post(ctx, wxappPayUrl, req); err != nil {
		return
	}
	g.Log().Info(ctx, gClientResponse.Raw())
	if gClientResponse.StatusCode != 200 {
		err = gerror.Newf("微信支付操作失败，状态码：%d", gClientResponse.StatusCode)
		return
	}
	if err = gjson.New(gClientResponse.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return resp, err
}
