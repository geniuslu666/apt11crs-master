package aladdinApi

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

func (c *AladdinClient) CancelOrder(ctx context.Context, params *CancelOrderParams) (resp *CancelOrderResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "GET", fmt.Sprintf("%s?orderId=%s", cancelOrder, params.OrderId), nil); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
