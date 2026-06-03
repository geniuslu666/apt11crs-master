package aladdinApi

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// ConformOrder 巴士预约订单确认
func (c *AladdinClient) ConformOrder(ctx context.Context, params *ConformOrderParams) (resp *ConformOrderResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", conformOrder, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
