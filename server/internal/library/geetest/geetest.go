package geetest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

const defaultValidateURL = "https://gcaptcha4.geetest.com/validate"

var errGeeTestValidateFailed = errors.New("验证失败")

type Config struct {
	Enabled     bool
	CaptchaID   string
	CaptchaKey  string
	ValidateURL string
	Timeout     time.Duration
}

type ValidateParams struct {
	CaptchaOutput string
	LotNumber     string
	PassToken     string
	GenTime       string
}

type ValidateResponse struct {
	Status      string                 `json:"status"`
	Result      string                 `json:"result"`
	Reason      string                 `json:"reason"`
	Code        string                 `json:"code"`
	Msg         string                 `json:"msg"`
	CaptchaArgs map[string]interface{} `json:"captcha_args"`
}

func (c Config) Normalize() Config {
	if strings.TrimSpace(c.ValidateURL) == "" {
		c.ValidateURL = defaultValidateURL
	}
	if c.Timeout <= 0 {
		c.Timeout = 5 * time.Second
	}
	return c
}

func (p ValidateParams) Validate() error {
	if strings.TrimSpace(p.LotNumber) == "" {
		return fmt.Errorf("missing lot_number")
	}
	if strings.TrimSpace(p.CaptchaOutput) == "" {
		return fmt.Errorf("missing captcha_output")
	}
	if strings.TrimSpace(p.PassToken) == "" {
		return fmt.Errorf("missing pass_token")
	}
	if strings.TrimSpace(p.GenTime) == "" {
		return fmt.Errorf("missing gen_time")
	}
	return nil
}

func Validate(ctx context.Context, cfg Config, params ValidateParams) (*ValidateResponse, error) {
	cfg = cfg.Normalize()
	if !cfg.Enabled {
		return &ValidateResponse{Result: "success", Reason: "disabled"}, nil
	}
	if strings.TrimSpace(cfg.CaptchaID) == "" || strings.TrimSpace(cfg.CaptchaKey) == "" {
		return nil, fmt.Errorf("config incomplete")
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("lot_number", params.LotNumber)
	form.Set("captcha_output", params.CaptchaOutput)
	form.Set("pass_token", params.PassToken)
	form.Set("gen_time", params.GenTime)
	form.Set("sign_token", signToken(params.LotNumber, cfg.CaptchaKey))

	validateURL := strings.TrimRight(cfg.ValidateURL, "/")
	validateURL = fmt.Sprintf("%s?captcha_id=%s", validateURL, url.QueryEscape(cfg.CaptchaID))
	startAt := time.Now()
	g.Log().Infof(ctx,
		"validate request url=%s captcha_id=%s lot_number=%s gen_time=%s captcha_output=%s pass_token=%s",
		validateURL,
		maskValue(cfg.CaptchaID, 3, 2),
		maskValue(params.LotNumber, 4, 4),
		params.GenTime,
		maskValue(params.CaptchaOutput, 6, 4),
		maskValue(params.PassToken, 6, 4),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, validateURL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		g.Log().Errorf(ctx, "validate build request failed err=%v", err)
		return nil, errGeeTestValidateFailed
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: cfg.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		g.Log().Errorf(ctx, "validate request failed cost=%s err=%v", time.Since(startAt), err)
		return nil, errGeeTestValidateFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		g.Log().Errorf(ctx, "validate bad status cost=%s status=%d", time.Since(startAt), resp.StatusCode)
		return nil, errGeeTestValidateFailed
	}

	var result ValidateResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		g.Log().Errorf(ctx, "validate decode response failed cost=%s err=%v", time.Since(startAt), err)
		return nil, errGeeTestValidateFailed
	}
	g.Log().Infof(ctx, "validate response cost=%s status=%s result=%s reason=%s code=%s", time.Since(startAt), result.Status, result.Result, result.Reason, result.Code)
	if result.Result != "success" {
		if result.Reason == "" {
			result.Reason = result.Msg
		}
		if result.Reason == "" {
			result.Reason = "validate failed"
		}
		g.Log().Errorf(ctx, "validate business failed cost=%s reason=%s code=%s", time.Since(startAt), result.Reason, result.Code)
		return &result, errGeeTestValidateFailed
	}

	return &result, nil
}

func signToken(lotNumber, captchaKey string) string {
	mac := hmac.New(sha256.New, []byte(captchaKey))
	mac.Write([]byte(lotNumber))
	return hex.EncodeToString(mac.Sum(nil))
}

func maskValue(value string, prefixLen, suffixLen int) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if len(value) <= prefixLen+suffixLen {
		return value
	}
	return value[:prefixLen] + "****" + value[len(value)-suffixLen:]
}
