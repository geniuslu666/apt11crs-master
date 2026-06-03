package aladdinApi

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// ConformLuggageOrder 确认行李订单
func (c *AladdinClient) ConformLuggageOrder(ctx context.Context, params *ConformLuggageOrderParams) (resp *ConformLuggageOrderResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "GET", fmt.Sprintf("%s?luggageOrderId=%s", conformLuggageOrder, params.LuggageOrderId), params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
