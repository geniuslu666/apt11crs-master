package miniGateway

import (
	"APT/internal/library/contexts"
	_ "APT/internal/logic"
	"APT/internal/model/model_gateway"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	AesKey     = []byte("202CB962AC59075B964B07152D234B70")
	AesIv      = []byte("233FA6B19FDE8CAF")
	Domain     = "appapi.apt11.net"
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 20,
			IdleConnTimeout:     90 * time.Second,
		},
		Timeout: 30 * time.Second,
	}
	//SignKey = "D9840773233FA6B19FDE8CAF765402F5"
)

// AppGateway 总入口
func AppGateway() {
	//var (
	//	ctx = context.Background()
	//)
	//global.Init(ctx)
	//tk, err := service.BasicsConfig().GetLoadToken(ctx)
	//if err != nil {
	//	return
	//}
	//token.SetConfig(tk)
	s := g.Server()
	g.Log().Path("logs/APP_GATEWAY").Info(context.TODO(), "start logs")
	s.BindHandler("GET:/api/ws", gatewayHandler)
	// 注册全局中间件
	s.BindHandler("GET:/*", gatewayHandler)
	s.BindHandler("POST:/*", gatewayHandler)

	s.SetPort(8008)
	s.Run()
}

// websocket 函数处理
func wsHandler(greq *ghttp.Request) {
	uri := fmt.Sprintf("ws://%s/api/ws", Domain)
	targetURL, err := url.Parse(uri) // 替换为你的目标服务器地址和路径
	if err != nil {
		log.Fatalf("invalid target URL: %v", err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = targetURL.Scheme
			r.URL.Host = targetURL.Host
			r.Host = r.URL.Host
			r.Header = greq.Header.Clone() // 克隆原始 Header（避免覆盖）
			// 关键：保留 WebSocket 升级所需的头（可能被默认 Director 清除）
			r.Header.Set("Upgrade", "websocket")
			r.Header.Set("Connection", "Upgrade")
			// 路径处理：假设客户端访问的路径如`/proxy/chat`，需要去掉路径前缀以匹配后端
			//r.URL.Path = "/" + r.URL.Path[len("/proxy"):] // 剔除路径中的 `/proxy` 部分
		},
	}

	proxy.ServeHTTP(greq.Response.ResponseWriter, greq.Request)
}

// 网关处理函数
func gatewayHandler(greq *ghttp.Request) {
	uri := fmt.Sprintf("http://%s", Domain)
	var (
		targetPmsURL  = uri
		targetTpURL   = uri
		decryptedBody []byte
		RequestBody   []byte
		err           error
		req           *http.Request
		resp          *http.Response
		params        map[string]*g.Var
		body          []byte
		buffParams    *bytes.Buffer
		jumpUrl       = []string{
			"/api/upload/file",
		}
		ctx = greq.Context()
	)
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, greq.URL.Path)
	if greq.URL.Path == "/api/ws" {
		g.Log().Path("logs/APP_GATEWAY").Info(ctx, "websocket proxy")
		wsHandler(greq)
		return
	}
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, "--[REQUEST]----------------------------------------------------------------")
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, greq.Request.URL.Path)
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, greq.Request.Method)
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, greq.GetBodyString())
	if g.IsEmpty(greq.GetBodyString()) {
		buffParams = bytes.NewBuffer(greq.GetBody())
	} else if !containsString(jumpUrl, greq.URL.Path) {
		if RequestBody, err = io.ReadAll(greq.Body); err != nil {
			err = errors.New("parsing fail")
			goto ERR
		}
		// 重新设置请求体
		greq.Body = io.NopCloser(bytes.NewReader(RequestBody))
		// 调用 ReloadParam 以确保请求体的参数被重新加载
		greq.ReloadParam()

		contexts.SetData(ctx, "request.body", RequestBody)
		buffParams = bytes.NewBuffer(decryptedBody)
	} else {
		buffParams = bytes.NewBuffer(greq.GetBody())
	}

	if strings.HasPrefix(greq.URL.Path, "/api") {
		// 创建新的请求
		if req, err = http.NewRequest(greq.Method, targetPmsURL+greq.URL.Path, buffParams); err != nil {
			goto ERR
		}
		// 复制请求头
		req.Header = greq.Header
		req.Header.Set("X-Forwarded-For", greq.GetClientIp())
	} else if strings.HasPrefix(greq.URL.Path, "/terminal") {
		// 创建新的请求
		if req, err = http.NewRequest(greq.Method, targetPmsURL+greq.URL.Path, buffParams); err != nil {
			goto ERR
		}
		// 复制请求头
		req.Header = greq.Header
		req.Header.Set("X-Forwarded-For", greq.GetClientIp())
	} else {
		// 创建新的请求
		if req, err = http.NewRequest(greq.Method, targetTpURL+greq.URL.Path, buffParams); err != nil {
			goto ERR
		}
		// 复制请求头
		g.Log().Info(ctx, contexts.GetMemberUser(ctx))
		req.Header = greq.Header
		req.Header.Set("MemberInfo", gjson.New(contexts.GetMemberUser(ctx)).String())
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Forwarded-For", greq.GetClientIp())
	}
	if !g.IsEmpty(req.Header.Get("Accept-Encoding")) {
		req.Header.Set("Accept-Encoding", "identity")
	}
	req.Header.Set("x-language", contexts.GetLanguage(ctx))
	for k, v := range req.Header {
		g.Log().Path("logs/APP_GATEWAY").Info(ctx, " -H ", k, v[0])
	}
	// 复制请求头
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, greq.URL.Path, greq.Method, greq.Header, params)
	// 发送请求到目标服务器
	if resp, err = httpClient.Do(req); err != nil {
		goto ERR
	}
	defer func() { _ = resp.Body.Close() }()
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, "--[RESPONSE]----------------------------------------------------------------")
	// 读取目标服务器响应
	if body, err = io.ReadAll(resp.Body); err != nil {
		goto ERR
	}
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, resp.StatusCode)
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, resp.Header)
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, string(body))

	greq.Response.Header().Set("Content-Type", "application/json")
	greq.Response.Write(body)
	return
ERR:
	greq.SetError(err)
	return
}

// Response 网关响应中间件统一处理
func Response(r *ghttp.Request) {
	r.Middleware.Next()
	var (
		ghttpcode     model_gateway.GatewayCode
		err           error
		encryptedBody []byte
		ctx           = r.Context()
		body          []byte
		resJsonData   interface{}
	)
	if r.GetError() != nil {
		err = r.GetError()
		goto ERR
	}
	body = r.Response.Buffer()
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, body)
	// 加密响应
	if encryptedBody, err = Encrypt(body); err != nil {
		goto ERR
	}
	ghttpcode.ResCode = gcode.CodeOK.Code()
	ghttpcode.ResMessage = gcode.CodeOK.Message()
	ghttpcode.ResData = string(encryptedBody)
	g.Log().Info(ctx, body)
	if !g.IsEmpty(body) {
		if err = json.Unmarshal(body, &resJsonData); err != nil {
			goto ERR
		}
	}
	ghttpcode.ResJsonData = resJsonData
	contexts.SetGatewayCode(ctx, &ghttpcode)
	contexts.SetDataMap(r.Context(), g.Map{
		"request.takeUpTime": gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds(),
		// ...
	})
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, ghttpcode)
	r.Response.ClearBuffer()
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteJsonExit(ghttpcode)
	return
ERR:
	g.Log().Path("logs/APP_GATEWAY").Error(ctx, err)
	ghttpcode.ResCode = gcode.CodeBusinessValidationFailed.Code()
	ghttpcode.ResMessage = err.Error()
	g.Log().Path("logs/APP_GATEWAY").Info(ctx, ghttpcode)
	r.Response.ClearBuffer()
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteJsonExit(ghttpcode)
}

// HexDecode 16进制解码
func HexDecode(s string) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst, []byte(s))        //进制转换, src->dst
	if err != nil {
		g.Dump(err)
		err = errors.New("解析失败")
		return nil, err
	}
	return dst[:n], err //返回0:n的数据.
}

// HexEncode 字符串转为16进制
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s))) //申请一个切片, 指明大小. 必须使用hex.EncodedLen
	n := hex.Encode(dst, []byte(s))             //字节流转化成16进制
	return dst[:n]
}

// Decrypt 解密
func Decrypt(cipherText string) (bodyText []byte, err error) {
	var (
		BodyHex []byte
	)
	// 解密
	if BodyHex, err = HexDecode(cipherText); err != nil {
		return
	}
	if bodyText, err = gaes.DecryptCBC(BodyHex, AesKey, AesIv); err != nil {
		return
	}
	return
}

// Encrypt 加密
func Encrypt(plainText []byte) (bodyText []byte, err error) {
	var (
		BodyHex []byte
	)
	if BodyHex, err = gaes.EncryptCBC(plainText, AesKey, AesIv); err != nil {
		return
	}
	bodyText = HexEncode(string(BodyHex))
	return
}

// 判断字符串是否在slice中
func containsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
