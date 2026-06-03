package paycloud

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"sort"
	"strings"
)

type Client struct {
	AppID           string
	StoreNo         string
	WxMiniStoreNo   string
	MerchantNo      string
	PrivateKey      string
	Endpoint        string
	SubAppid        string
	SubWxMiniAppid  string
	RefundUrlWeb    string
	RefundUrlApp    string
	NotifyUrl       string
	RefundNotifyUrl string
}

func NewClient(ctx context.Context, PaypalReturnUrlApp string) (payCloudClient *Client, err error) {
	var (
		PayConfig *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	if !g.IsEmpty(PaypalReturnUrlApp) {
		PayConfig.PayCloudReturnUrlWeb = PaypalReturnUrlApp
		PayConfig.PayCloudReturnUrlApp = PaypalReturnUrlApp
	}
	return &Client{
		AppID:           PayConfig.PayCloudAppID,
		MerchantNo:      PayConfig.PayCloudMerchantNo,
		StoreNo:         PayConfig.PayCloudStoreNo,
		WxMiniStoreNo:   PayConfig.PayCloudWxMiniStoreNo,
		PrivateKey:      PayConfig.PayCloudPrivateKey,
		Endpoint:        PayConfig.PayCloudEndpoint,
		SubAppid:        PayConfig.PayCloudSubAppId,
		SubWxMiniAppid:  PayConfig.PayCloudWxMiniAppid,
		RefundUrlWeb:    PayConfig.PayCloudReturnUrlWeb,
		RefundUrlApp:    PayConfig.PayCloudReturnUrlApp,
		NotifyUrl:       PayConfig.PayCloudNotifyUrl,
		RefundNotifyUrl: PayConfig.PayCloudRefundNotifyUrl,
	}, nil
}

func (c *Client) generateSignature(ctx context.Context, params map[string]string) (string, error) {
	var keys []string
	for k := range params {
		if k != "sign" && k != "timestamp" {
			keys = append(keys, k)
		}
	}
	keys = append(keys, "timestamp")
	sort.Strings(keys)

	var signStrBuilder strings.Builder
	for _, k := range keys {

		if !g.IsEmpty(params[k]) {
			if !g.IsEmpty(signStrBuilder.String()) {
				signStrBuilder.WriteString("&")
			}
			signStrBuilder.WriteString(k + "=" + params[k])
		}

	}
	signStr := signStrBuilder.String()
	g.Log().Info(ctx, signStr)
	return GenerateSign(signStr, c.PrivateKey)
}

// formatPrivateKey 格式化私钥字符串
func formatPrivateKey(key string) string {
	return strings.Join(strings.Split(key, "\n"), "")
}

// GenerateSign 生成签名
func GenerateSign(data, priKey string) (string, error) {
	// 格式化私钥
	privateKeyPEM := fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", formatPrivateKey(priKey))

	// 解析私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("the private key format is incorrect")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// 计算 SHA-256 哈希值
	hashed := sha256.Sum256([]byte(data))

	// 使用私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	// 签名结果进行 base64 编码
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}

func (c *Client) DoRequest(ctx context.Context, paramsReq interface{}) (resp *gclient.Response, err error) {
	var (
		params    = gvar.New(paramsReq).MapStrStr()
		signature string
		jsonData  []byte
		client    *gclient.Client
	)
	if signature, err = c.generateSignature(ctx, params); err != nil {
		g.Log().Path("logs/SDK/PAY_CLOUD").Error(ctx, err)
		return nil, err
	}
	params["sign"] = signature

	if jsonData, err = json.Marshal(params); err != nil {
		g.Log().Path("logs/SDK/PAY_CLOUD").Error(ctx, err)
		return nil, err
	}

	client = g.Client()
	client.SetHeader("Content-Type", "application/json;charset=utf-8")
	if resp, err = client.Post(ctx, c.Endpoint, jsonData); err != nil {
		return nil, err
	}
	g.Log().Path("logs/SDK/PAY_CLOUD").Info(ctx, resp.Raw())
	return resp, nil
}
