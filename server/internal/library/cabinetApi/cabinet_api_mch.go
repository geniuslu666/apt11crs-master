package cabinetApi

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
)

// Cabinet 查询储物柜信息
func (c *CabinetClient) Cabinet(ctx context.Context, params *CabinetInfoParams) (resp *CabinetInfoResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", cabinetInfo, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}

// CreateOrder 下单
func (c *CabinetClient) CreateOrder(ctx context.Context, params *CabinetCreateOrderParams) (resp *CabinetCreateOrderResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", createOrder, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}

// OrderQuery 订单查询
func (c *CabinetClient) OrderQuery(ctx context.Context, params *CabinetOrderQueryParams) (resp *CabinetOrderQueryResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", orderQuery, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}

// PayOvertime 支付超时费
func (c *CabinetClient) PayOvertime(ctx context.Context, params *CabinetPayOvertimeParams) (resp *CabinetPayOvertimeResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", payOvertime, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}

// CabinetList 储物柜列表
func (c *CabinetClient) CabinetList(ctx context.Context) (resp *CabinetListResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", cabinetList, nil); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}

// OrderComplete 订单完成
func (c *CabinetClient) OrderComplete(ctx context.Context, params *OrderCompleteParams) (resp *OrderCompleteResponse, err error) {
	var (
		response *gclient.Response
	)
	if response, err = c.DoRequest(ctx, "POST", orderComplete, params); err != nil {
		return
	}
	if err = gjson.New(response.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
