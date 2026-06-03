// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_pay"
	"APT/internal/model/input/input_travel"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ITravelOrder interface {
		// Model 订单ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取订单列表
		List(ctx context.Context, in *input_travel.TravelOrderListInp) (list []*input_travel.TravelOrderListModel, totalCount int, err error)
		// View 获取订单详情
		View(ctx context.Context, in *input_travel.TravelOrderViewInp) (res *input_travel.TravelOrderViewModel, err error)
		// Refund 订单退款
		Refund(ctx context.Context, in *input_travel.TravelOrderRefundInp) (err error)
		// OrderExpiration 订单过期
		OrderExpiration(ctx context.Context, OrderSn string) (err error)
		// OrderOverdue 订单过期
		OrderOverdue(ctx context.Context, OrderSn string) (err error)
		RefreshCode(ctx context.Context, in *input_travel.TravelOrderRefreshCodeInp) (res *input_travel.TravelOrderRefreshCodeModel, err error)
	}
	ITravelOrderCreateService interface {
		// CreateOrder 创建订单
		CreateOrder(ctx context.Context, in *input_travel.CreateOrderInp) (out *input_travel.CreateOrderModel, err error)
	}
	ITravelOrderCreatePreService interface {
		// PreOrder 预下单
		PreOrder(ctx context.Context, in *input_travel.PreCreateOrderInp) (out *input_travel.PreCreateOrderModel, err error)
		// PreOrderDetail 预下单详情
		PreOrderDetail(ctx context.Context, in *input_travel.PreOrderDetailInp) (out *input_travel.PreOrderDetailModel, err error)
		// PrePayInfo 预下单重新计算支付金额支付模式
		PrePayInfo(ctx context.Context, in *input_pay.PrePayInfoInp) (out *input_pay.PrePayInfoModel, err error)
	}
	ITravelOrderRefundService interface {
		// ApplyRefundDetail 预退款订单详情
		ApplyRefundDetail(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (out *input_travel.TravelOrderApplyRefundDetailModel, err error)
		// RefundOrder 退款
		RefundOrder(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (err error)
		// RefundOrderDetail 退款订单详情
		RefundOrderDetail(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (out *input_travel.RefundDetailModel, err error)
	}
	ITravelProduct interface {
		// Model 产品ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取产品列表
		List(ctx context.Context, in *input_travel.TravelProductListInp) (list []*input_travel.TravelProductListModel, totalCount int, err error)
		// RecycleList 回收站列表
		RecycleList(ctx context.Context, in *input_travel.TravelProductListInp) (list []*input_travel.TravelProductListModel, totalCount int, err error)
		AppList(ctx context.Context, in *input_travel.TravelProductAppListInp) (list []*input_travel.TravelProductAppListModel, totalCount int, err error)
		// GetIds 通过多语言 UUID 列表查询产品 ID
		GetIds(ctx context.Context, uuIds []string) (ids []int64, err error)
		// GetSkuIds 通过多语言 UUID 列表查询产品SKU ID
		GetSkuIds(ctx context.Context, uuIds []string) (ids []int64, err error)
		// View 获取产品详情
		View(ctx context.Context, in *input_travel.TravelProductViewInp) (res *input_travel.TravelProductViewModel, err error)
		AppView(ctx context.Context, in *input_travel.TravelProductAppViewInp) (res *input_travel.TravelProductAppViewModel, err error)
		// Edit 新增/编辑产品
		Edit(ctx context.Context, in *input_travel.TravelProductEditInp) (err error)
		// Delete 软删除产品（移入回收站）
		Delete(ctx context.Context, in *input_travel.TravelProductDeleteInp) (err error)
		// Restore 从回收站恢复产品
		Restore(ctx context.Context, in *input_travel.TravelProductDeleteInp) (err error)
		// Status 更新产品状态
		Status(ctx context.Context, in *input_travel.TravelProductStatusInp) (err error)
		// SkuStock 查询SKU库存
		SkuStock(ctx context.Context, in *input_travel.TravelProductSkuStockInp) (list []*input_travel.TravelProductSkuStockModel, err error)
	}
	ITravelProductSku interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取车型列表
		List(ctx context.Context, in *input_travel.TravelProductSkuListInp) (list []*input_travel.TravelProductSkuListModel, totalCount int, err error)
		// View 获取车型详情
		View(ctx context.Context, in *input_travel.TravelProductSkuViewInp) (res *input_travel.TravelProductSkuViewModel, err error)
		// Edit 新增/编辑车型
		Edit(ctx context.Context, in *input_travel.TravelProductSkuEditInp) (err error)
		// Delete 删除车型
		Delete(ctx context.Context, in *input_travel.TravelProductSkuDeleteInp) (err error)
		// Status 更新车型状态
		Status(ctx context.Context, in *input_travel.TravelProductSkuStatusInp) (err error)
	}
	ITravelVerifyRecord interface {
		// Model 核销记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取核销记录列表
		List(ctx context.Context, in *input_travel.TravelVerifyRecordListInp) (list []*input_travel.TravelVerifyRecordListModel, totalCount int, err error)
		// VerifyList 获取核销端核销记录列表
		VerifyList(ctx context.Context, in *input_travel.VerifyListInp) (list []*input_travel.VerifyListModel, totalCount int, err error)
		// VerifyView 获取核销记录信息
		VerifyView(ctx context.Context, in *input_travel.VerifyLogViewInp) (res *input_travel.VerifyLogViewModel, err error)
		CodeView(ctx context.Context, in *input_travel.CodeViewInp) (res *input_travel.CodeViewModel, err error)
		CodeVerify(ctx context.Context, in *input_travel.CodeVerifyInp) (res *input_travel.CodeVerifyModel, err error)
	}
	ITravelVerifyStaff interface {
		// Model 核销人员ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取核销人员列表
		List(ctx context.Context, in *input_travel.TravelVerifyStaffListInp) (list []*input_travel.TravelVerifyStaffListModel, totalCount int, err error)
		// View 获取核销人员详情
		View(ctx context.Context, in *input_travel.TravelVerifyStaffViewInp) (res *input_travel.TravelVerifyStaffViewModel, err error)
		// Edit 新增/编辑核销人员
		Edit(ctx context.Context, in *input_travel.TravelVerifyStaffEditInp) (err error)
		// Delete 删除核销人员
		Delete(ctx context.Context, in *input_travel.TravelVerifyStaffDeleteInp) (err error)
		// ScopeOptions 获取核销权限范围选项（产品与SKU）
		ScopeOptions(ctx context.Context) (list []*input_travel.TravelVerifyStaffScopeOptionModel, err error)
		// Status 更新核销人员状态
		Status(ctx context.Context, in *input_travel.TravelVerifyStaffStatusInp) (err error)
	}
)

var (
	localTravelOrder                 ITravelOrder
	localTravelOrderCreateService    ITravelOrderCreateService
	localTravelOrderCreatePreService ITravelOrderCreatePreService
	localTravelOrderRefundService    ITravelOrderRefundService
	localTravelProduct               ITravelProduct
	localTravelProductSku            ITravelProductSku
	localTravelVerifyRecord          ITravelVerifyRecord
	localTravelVerifyStaff           ITravelVerifyStaff
)

func TravelOrder() ITravelOrder {
	if localTravelOrder == nil {
		panic("implement not found for interface ITravelOrder, forgot register?")
	}
	return localTravelOrder
}

func RegisterTravelOrder(i ITravelOrder) {
	localTravelOrder = i
}

func TravelOrderCreateService() ITravelOrderCreateService {
	if localTravelOrderCreateService == nil {
		panic("implement not found for interface ITravelOrderCreateService, forgot register?")
	}
	return localTravelOrderCreateService
}

func RegisterTravelOrderCreateService(i ITravelOrderCreateService) {
	localTravelOrderCreateService = i
}

func TravelOrderCreatePreService() ITravelOrderCreatePreService {
	if localTravelOrderCreatePreService == nil {
		panic("implement not found for interface ITravelOrderCreatePreService, forgot register?")
	}
	return localTravelOrderCreatePreService
}

func RegisterTravelOrderCreatePreService(i ITravelOrderCreatePreService) {
	localTravelOrderCreatePreService = i
}

func TravelOrderRefundService() ITravelOrderRefundService {
	if localTravelOrderRefundService == nil {
		panic("implement not found for interface ITravelOrderRefundService, forgot register?")
	}
	return localTravelOrderRefundService
}

func RegisterTravelOrderRefundService(i ITravelOrderRefundService) {
	localTravelOrderRefundService = i
}

func TravelProduct() ITravelProduct {
	if localTravelProduct == nil {
		panic("implement not found for interface ITravelProduct, forgot register?")
	}
	return localTravelProduct
}

func RegisterTravelProduct(i ITravelProduct) {
	localTravelProduct = i
}

func TravelProductSku() ITravelProductSku {
	if localTravelProductSku == nil {
		panic("implement not found for interface ITravelProductSku, forgot register?")
	}
	return localTravelProductSku
}

func RegisterTravelProductSku(i ITravelProductSku) {
	localTravelProductSku = i
}

func TravelVerifyRecord() ITravelVerifyRecord {
	if localTravelVerifyRecord == nil {
		panic("implement not found for interface ITravelVerifyRecord, forgot register?")
	}
	return localTravelVerifyRecord
}

func RegisterTravelVerifyRecord(i ITravelVerifyRecord) {
	localTravelVerifyRecord = i
}

func TravelVerifyStaff() ITravelVerifyStaff {
	if localTravelVerifyStaff == nil {
		panic("implement not found for interface ITravelVerifyStaff, forgot register?")
	}
	return localTravelVerifyStaff
}

func RegisterTravelVerifyStaff(i ITravelVerifyStaff) {
	localTravelVerifyStaff = i
}
