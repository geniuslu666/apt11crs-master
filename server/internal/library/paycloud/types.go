package paycloud

type CommonRequest struct {
	Format     string `json:"format"`             // 公共参数 请求数据格式
	Sign       string `json:"sign"`               // 公共参数 签名
	AppID      string `json:"app_id"`             // 公共参数 应用ID
	SignType   string `json:"sign_type"`          // 公共参数 签名算法RSA
	Version    string `json:"version"`            // 公共参数 版本号
	Timestamp  string `json:"timestamp"`          // 公共参数 时间戳
	Charset    string `json:"charset"`            // 公共参数 请求使用的编码格式
	MerchantNo string `json:"merchant_no"`        // 商户号
	StoreNo    string `json:"store_no,omitempty"` // 门店号
}

type CommonResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
