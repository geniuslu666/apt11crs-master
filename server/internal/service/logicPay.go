// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/model/input/input_pay"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
)

type (
	IPayService interface {
		// SelectThirdPay 选择第三方支付方式
		SelectThirdPay(ctx context.Context, in *input_pay.SelectThirdPayInp) (out *input_pay.SelectThirdPayModel, err error)
		// BalancePay 纯余额支付
		BalancePay(ctx context.Context, in *input_pay.BalancePayInp) (out *input_pay.BalancePayModel, err error)
		HotelCreateOrder(ctx context.Context, in *input_pay.HotelCreateOrderInp, tx gdb.TX) (out *input_pay.HotelCreateOrderModel, err error)
		// ThirdPayCompleted 三方支付完成后处理支付流水
		ThirdPayCompleted(ctx context.Context, in *input_pay.ThirdPayInp) (err error)
		PayAmountChange(ctx context.Context, tx gdb.TX, OrderSn string) (err error)
		HandleFoodOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error)
		HandleSpaOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error)
		HandleCarOrderMq(ctx context.Context, OrderSn string, Logger *glog.Logger) (err error)
		CabinetCreateOrder(ctx context.Context, in *input_pay.CabinetCreateOrderInp, tx gdb.TX) (out *input_pay.CabinetCreateOrderModel, err error)
		TravelCreateOrder(ctx context.Context, in *input_pay.TravelCreateOrderInp, tx gdb.TX) (out *input_pay.TravelCreateOrderModel, err error)
		OrderBalancePay(ctx context.Context, tx gdb.TX, OrderSn string, MemberId int, PayType string) (err error)
		// PayInfo 支付信息计算
		PayInfo(ctx context.Context, in *input_pay.PayInfoInp) (out *input_pay.PayInfoModel, err error)
		// BookingPrice 预定价格计算
		BookingPrice(ctx context.Context, in *input_pay.BookingPriceInp) (out *input_pay.BookingPriceModel, err error)
		List(ctx context.Context, in *input_pay.PmsTransactionListInp) (list []*input_pay.PmsTransactionListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_pay.PmsTransactionListExportInp) (err error)
		View(ctx context.Context, in *input_pay.PmsTransactionViewInp) (res *input_pay.PmsTransactionViewModel, err error)
		GetTransactionSn(ctx context.Context, startDate string, endDate string, payChannel string, payType string) (transactionSnArr []string, err error)
	}
)

var (
	localPayService IPayService
)

func PayService() IPayService {
	if localPayService == nil {
		panic("implement not found for interface IPayService, forgot register?")
	}
	return localPayService
}

func RegisterPayService(i IPayService) {
	localPayService = i
}
