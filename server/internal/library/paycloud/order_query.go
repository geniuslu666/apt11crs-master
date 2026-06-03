package paycloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gclient"
)

type OrderQueryRequest struct {
	MerchantNo      string `json:"merchant_no"`
	Charset         string `json:"charset"`
	Method          string `json:"method"`
	Format          string `json:"format"`
	Sign            string `json:"sign"`
	AppID           string `json:"app_id"`
	SignType        string `json:"sign_type"`
	Version         string `json:"version"`
	MerchantOrderNo string `json:"merchant_order_no"`
	Timestamp       string `json:"timestamp"`
}

type OrderQueryResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	OrderNo string `json:"order_no"`
	Status  string `json:"status"`
}

func (c *Client) OrderQuery(ctx context.Context, req OrderQueryRequest) (orderQueryResp OrderQueryResponse, err error) {
	var (
		resp   *gclient.Response
		params = gvar.New(req).MapStrStr()
	)

	if resp, err = c.DoRequest(ctx, params); err != nil {
		return
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to query order: %s", resp.Status)
		return
	}

	if err = json.Unmarshal(resp.ReadAll(), &orderQueryResp); err != nil {
		return
	}
	if orderQueryResp.Code != "0" {
		err = gerror.New(orderQueryResp.Message)
		return
	}
	return
}
