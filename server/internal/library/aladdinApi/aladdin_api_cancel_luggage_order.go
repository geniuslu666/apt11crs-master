package aladdinApi

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// CancelLuggageOrder 取消行李订单
func (c *AladdinClient) CancelLuggageOrder(ctx context.Context, params *CancelLuggageOrderParams) (resp *CancelLuggageOrderResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "GET", fmt.Sprintf("%s?luggageOrderId=%s", cancelLuggageOrder, params.LuggageOrderId), params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
