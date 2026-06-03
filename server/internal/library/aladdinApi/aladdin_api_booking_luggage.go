package aladdinApi

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// BookingLuggage 行李生单
func (c *AladdinClient) BookingLuggage(ctx context.Context, params *BookingLuggageParams) (resp *BookingLuggageResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", bookingLuggage, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
