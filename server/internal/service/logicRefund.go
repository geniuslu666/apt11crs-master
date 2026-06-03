// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/model/input/input_refund"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IRefund interface {
		// RefundOrder 金额 订单退款
		RefundOrder(ctx context.Context, in *input_refund.RefundAmountInp, tx gdb.TX) (err error)
		// BalanceRefund 余额退款
		BalanceRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
		// PayCloudRefund 云支付退款
		PayCloudRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
		// PaypalRefund paypal退款
		PaypalRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
		MlilifeRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
		// StripeRefund stripe退款
		StripeRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
		RefundStatusChange(ctx context.Context, tx gdb.TX, OrderSn string) (err error)
		// H5FXRefund 分销退款
		H5FXRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error)
	}
	IRefundTransaction interface {
		List(ctx context.Context, in *input_refund.PmsTransactionRefundListInp) (list []*input_refund.PmsTransactionRefundListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_refund.PmsTransactionRefundListExportInp) (err error)
		View(ctx context.Context, in *input_refund.PmsTransactionRefundViewInp) (res *input_refund.PmsTransactionRefundViewModel, err error)
	}
)

var (
	localRefund            IRefund
	localRefundTransaction IRefundTransaction
)

func Refund() IRefund {
	if localRefund == nil {
		panic("implement not found for interface IRefund, forgot register?")
	}
	return localRefund
}

func RegisterRefund(i IRefund) {
	localRefund = i
}

func RefundTransaction() IRefundTransaction {
	if localRefundTransaction == nil {
		panic("implement not found for interface IRefundTransaction, forgot register?")
	}
	return localRefundTransaction
}

func RegisterRefundTransaction(i IRefundTransaction) {
	localRefundTransaction = i
}
