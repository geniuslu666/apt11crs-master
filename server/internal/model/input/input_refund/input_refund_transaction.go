// Package sysin

package input_refund

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransactionRefundViewInp 获取指定退款流水信息
type PmsTransactionRefundViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsTransactionRefundViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTransactionRefundViewModel struct {
	entity.PmsTransactionRefund
}

// PmsTransactionRefundListInp 获取退款流水列表
type PmsTransactionRefundListInp struct {
	input_form.PageReq
	Id            int           `json:"id"            dc:"主键"`
	OrderSn       string        `json:"orderSn"       dc:"订单号"`
	TransactionSn string        `json:"transactionSn" dc:"支付流水号"`
	RefundChannel string        `json:"refundChannel"    dc:"支付平台"`
	RefundType    string        `json:"refundType"       dc:"支付方式"`
	RefundStatus  string        `json:"refundStatus"     dc:"退款状态"`
	RefundTime    []*gtime.Time `json:"refundTime"    dc:"退款时间"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
}

func (in *PmsTransactionRefundListInp) Filter(ctx context.Context) (err error) {
	return
}

// PmsTransactionRefundListExportInp 获取退款流水导出列表
type PmsTransactionRefundListExportInp struct {
	input_form.PageReq
	Id            int    `json:"id"            dc:"主键"`
	OrderSn       string `json:"orderSn"       dc:"订单号"`
	TransactionSn string `json:"transactionSn" dc:"支付流水号"`
	RefundChannel string `json:"refundChannel"    dc:"支付平台"`
	RefundType    string `json:"refundType"       dc:"支付方式"`
	RefundStatus  string `json:"refundStatus"     dc:"退款状态"`
	RefundTime    string `json:"refundTime"    dc:"退款时间"`
	CreatedAt     string `json:"createdAt"     dc:"创建时间"`
}

func (in *PmsTransactionRefundListExportInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTransactionRefundListModel struct {
	Id            int         `json:"id"            dc:"主键"`
	OrderSn       string      `json:"orderSn"       dc:"订单号"`
	TransactionSn string      `json:"transactionSn" dc:"支付流水号"`
	RefundChannel string      `json:"refundChannel" dc:"支付渠道"`
	RefundType    string      `json:"refundType"    dc:"'支付方式   BAL 余额'"`
	RefundSn      string      `json:"refundSn"      dc:"退款流水号"`
	TransNo       string      `json:"transNo"       dc:"退款交易号"`
	CancelOrderSn string      `json:"cancelOrderSn" dc:"取消订单号"`
	RefundAmount  float64     `json:"refundAmount"  dc:"退款金额"`
	RefundTime    *gtime.Time `json:"refundTime"    dc:"退款时间"`
	RefundStatus  string      `json:"refundStatus"  dc:"退款状态"`
	CancelId      string      `json:"cancelId"      dc:"取消政策ID"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// PmsTransactionRefundExportModel 导出退款流水
type PmsTransactionRefundExportModel struct {
	Id            int         `json:"id"            dc:"主键"`
	OrderSn       string      `json:"orderSn"       dc:"订单号"`
	TransactionSn string      `json:"transactionSn" dc:"支付流水号"`
	RefundChannel string      `json:"refundChannel" dc:"支付渠道"`
	RefundType    string      `json:"refundType"    dc:"支付方式"`
	RefundSn      string      `json:"refundSn"      dc:"退款流水号"`
	TransNo       string      `json:"transNo"       dc:"退款交易号"`
	CancelOrderSn string      `json:"cancelOrderSn" dc:"取消订单号"`
	RefundAmount  float64     `json:"refundAmount"  dc:"退款金额"`
	RefundTime    *gtime.Time `json:"refundTime"    dc:"退款时间"`
	RefundStatus  string      `json:"refundStatus"  dc:"退款状态"`
	CancelId      string      `json:"cancelId"      dc:"取消政策ID"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
