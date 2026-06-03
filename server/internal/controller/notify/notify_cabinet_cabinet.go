package notify

import (
	"APT/internal/library/cabinetApi"
	"APT/internal/model"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/cabinet"
)

func (c *ControllerCabinet) CabinetInfo(ctx context.Context, req *cabinet.CabinetInfoReq) (res *cabinet.CabinetInfoRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetInfoResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).Cabinet(ctx, req.CabinetInfoParams); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.CabinetInfoRes)
	res.CabinetInfoResponseItem = cabinetResponse.Data
	return
}
func (c *ControllerCabinet) CreateOrder(ctx context.Context, req *cabinet.CreateOrderReq) (res *cabinet.CreateOrderRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetCreateOrderResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).CreateOrder(ctx, req.CabinetCreateOrderParams); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.CreateOrderRes)
	res.CabinetCreateOrderResponseItem = cabinetResponse.Data
	return
}
func (c *ControllerCabinet) OrderQuery(ctx context.Context, req *cabinet.OrderQueryReq) (res *cabinet.OrderQueryRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetOrderQueryResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).OrderQuery(ctx, req.CabinetOrderQueryParams); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.OrderQueryRes)
	res.CabinetOrderQueryResponseItem = cabinetResponse.Data
	return
}
func (c *ControllerCabinet) PayOvertime(ctx context.Context, req *cabinet.PayOvertimeReq) (res *cabinet.PayOvertimeRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetPayOvertimeResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).PayOvertime(ctx, req.CabinetPayOvertimeParams); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.PayOvertimeRes)
	res.CabinetPayOvertimeResponseItem = cabinetResponse.Data
	return
}
func (c *ControllerCabinet) CabinetList(ctx context.Context, req *cabinet.CabinetListReq) (res *cabinet.CabinetListRes, err error) {
	var (
		cabinetResponse  *cabinetApi.CabinetListResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).CabinetList(ctx); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	res = new(cabinet.CabinetListRes)
	res.List = cabinetResponse.Data
	return
}
func (c *ControllerCabinet) OrderComplete(ctx context.Context, req *cabinet.OrderCompleteReq) (res *cabinet.OrderCompleteRes, err error) {
	var (
		cabinetResponse  *cabinetApi.OrderCompleteResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).OrderComplete(ctx, req.OrderCompleteParams); err != nil {
		return
	}
	if cabinetResponse.Code != 0 {
		err = gerror.New(cabinetResponse.Msg)
		return
	}
	return
}
