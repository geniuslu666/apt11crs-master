package aladdinApi

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// OfferInfo 查询报价
func (c *AladdinClient) OfferInfo(ctx context.Context, params *OfferInfoParams) (resp *OfferInfoResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", offerInfo, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
