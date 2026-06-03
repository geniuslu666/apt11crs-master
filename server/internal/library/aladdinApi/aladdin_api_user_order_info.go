package aladdinApi

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

func (c *AladdinClient) UserOrderInfo(ctx context.Context, params *UserOrderInfoParams) (resp *UserOrderInfoResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "GET", fmt.Sprintf("%s?orderId=%s", userOrderInfo, params.OrderId), params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
