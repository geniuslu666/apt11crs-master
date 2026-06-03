package aladdinApi

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// BookingBus 巴士预约生单
func (c *AladdinClient) BookingBus(ctx context.Context, params *BookingBusParams) (resp *BookingBusResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", bookingBus, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
