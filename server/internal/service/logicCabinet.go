// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_cabinet"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ICabinetService interface {
		// UpdateOrderStatus 更改订单状态
		UpdateOrderStatus(ctx context.Context, in *input_cabinet.UpdateOrderStatusInp) (err error)
		OrderRefund(ctx context.Context, in *input_cabinet.OrderRefundInp) (err error)
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		AppList(ctx context.Context, in *input_cabinet.OrderAppListInp) (list []*input_cabinet.OrderAppListModel, totalCount int, err error)
		AppView(ctx context.Context, in *input_cabinet.OrderAppViewInp) (res *input_cabinet.OrderAppViewModel, err error)
		// PayOvertimeInfo 支付超时费页面信息
		PayOvertimeInfo(ctx context.Context, in *input_cabinet.PayOvertimeInfoInp) (out *input_cabinet.PayOvertimeInfoModel, err error)
		// PayOvertime 支付超时费
		PayOvertime(ctx context.Context, in *input_cabinet.PayOvertimeInp) (out *input_cabinet.PayOvertimeModel, err error)
		List(ctx context.Context, in *input_cabinet.OrderListInp) (list []*input_cabinet.OrderListModel, totalCount int, err error)
		OrderView(ctx context.Context, in *input_cabinet.OrderViewInp) (res *input_cabinet.OrderViewModel, err error)
		ExportOrder(ctx context.Context, in *input_cabinet.OrderExportInp) (err error)
		StartExport(ctx context.Context, in *input_cabinet.OrderExportInp) (path string, err error)
		ExportList(ctx context.Context, in *input_cabinet.OrderExportListInp) (list []*input_cabinet.OrderExportListModel, totalCount int, err error)
		CompleteOrder(ctx context.Context, in *input_cabinet.OrderCompleteInp) (err error)
		// Refund 订单退款
		Refund(ctx context.Context, in *input_cabinet.CabinetOrderRefundInp) (err error)
		// CreateOrder 创建订单
		CreateOrder(ctx context.Context, in *input_cabinet.CreateOrderInp) (out *input_cabinet.CreateOrderModel, err error)
		// PreOrder 预下单
		PreOrder(ctx context.Context, in *input_cabinet.PreCreateOrderInp) (out *input_cabinet.PreCreateOrderModel, err error)
		// PreOrderDetail 预下单详情
		PreOrderDetail(ctx context.Context, in *input_cabinet.PreOrderDetailInp) (out *input_cabinet.PreOrderDetailModel, err error)
		// PrePayInfo 预下单重新计算支付金额支付模式
		PrePayInfo(ctx context.Context, in *input_cabinet.PrePayInfoInp) (out *input_cabinet.OrderPayInfoModel, err error)
		// OrderExpiration 订单过期
		OrderExpiration(ctx context.Context, OrderSn string) (err error)
	}
)

var (
	localCabinetService ICabinetService
)

func CabinetService() ICabinetService {
	if localCabinetService == nil {
		panic("implement not found for interface ICabinetService, forgot register?")
	}
	return localCabinetService
}

func RegisterCabinetService(i ICabinetService) {
	localCabinetService = i
}
