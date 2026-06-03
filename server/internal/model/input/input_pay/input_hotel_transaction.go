// Package sysin

package input_pay

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransactionViewInp 获取指定支付流水表信息
type PmsTransactionViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsTransactionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTransactionViewModel struct {
	entity.PmsTransaction
}

// PmsTransactionListInp 获取支付流水表列表
type PmsTransactionListInp struct {
	input_form.PageReq
	Id            int           `json:"id"            dc:"主键"`
	OrderSn       string        `json:"orderSn"       dc:"订单号"`
	TransactionSn string        `json:"transactionSn" dc:"支付流水号"`
	PayChannel    string        `json:"payChannel"    dc:"支付平台"`
	PayType       string        `json:"payType"       dc:"支付方式"`
	PayStatus     string        `json:"payStatus"       dc:"支付状态"`
	PayTime       []*gtime.Time `json:"payTime"       dc:"支付时间"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
}

func (in *PmsTransactionListInp) Filter(ctx context.Context) (err error) {
	return
}

// PmsTransactionListExportInp 获取支付流水表导出列表
type PmsTransactionListExportInp struct {
	input_form.PageReq
	Id            int    `json:"id"            dc:"主键"`
	OrderSn       string `json:"orderSn"       dc:"订单号"`
	TransactionSn string `json:"transactionSn" dc:"支付流水号"`
	PayChannel    string `json:"payChannel"    dc:"支付平台"`
	PayType       string `json:"payType"       dc:"支付方式"`
	PayStatus     string `json:"payStatus"     dc:"支付状态"`
	PayTime       string `json:"payTime"       dc:"支付时间"`
	CreatedAt     string `json:"createdAt"     dc:"创建时间"`
}

func (in *PmsTransactionListExportInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTransactionListModel struct {
	Id               int         `json:"id"               dc:"主键"`
	OrderSn          string      `json:"orderSn"          dc:"订单号"`
	TransactionSn    string      `json:"transactionSn"    dc:"支付流水号"`
	PaymentRequestId string      `json:"paymentRequestId" dc:"第三方支付流水号"`
	PayChannel       string      `json:"payChannel"       dc:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
	PayType          string      `json:"payType"          dc:"支付方式   BAL 余额"`
	Amount           float64     `json:"amount"           dc:"总金额"`
	PriceCurrency    string      `json:"priceCurrency"    dc:"币种"`
	PayAmount        float64     `json:"payAmount"        dc:"支付金额"`
	PayStatus        string      `json:"payStatus"        dc:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
	PayTime          *gtime.Time `json:"payTime"          dc:"支付时间"`
	ExpiredTime      *gtime.Time `json:"expiredTime"      dc:"过期时间"`
	RefundAmount     float64     `json:"refundAmount"     dc:"退款金额"`
	RefundStatus     string      `json:"refundStatus"     dc:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	ScenePayRate     float64     `json:"scenePayRate"     dc:"场景积分抵扣比例"`
	ExchangeRate     float64     `json:"exchangeRate"     dc:"积分汇率"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"更新时间"`
}

// PmsTransactionExportModel 导出支付流水表
type PmsTransactionExportModel struct {
	Id               int         `json:"id"               dc:"主键"`
	OrderSn          string      `json:"orderSn"          dc:"订单号"`
	TransactionSn    string      `json:"transactionSn"    dc:"支付流水号"`
	PaymentRequestId string      `json:"paymentRequestId" dc:"第三方支付流水号"`
	PayChannel       string      `json:"payChannel"       dc:"支付渠道"` // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
	PayType          string      `json:"payType"          dc:"支付方式"`
	Amount           float64     `json:"amount"           dc:"总金额"`
	PriceCurrency    string      `json:"priceCurrency"    dc:"币种"`
	PayAmount        float64     `json:"payAmount"        dc:"支付金额"`
	PayStatus        string      `json:"payStatus"        dc:"支付状态"` // WAIT 等待支付、DONE 完成支付、CANCEL 取消支付
	PayTime          *gtime.Time `json:"payTime"          dc:"支付时间"`
	ExpiredTime      *gtime.Time `json:"expiredTime"      dc:"过期时间"`
	RefundAmount     float64     `json:"refundAmount"     dc:"退款金额"`
	RefundStatus     string      `json:"refundStatus"     dc:"退款状态"` //  WAIT 未退款   PART  部分退款 DONE 全部退款
	ScenePayRate     float64     `json:"scenePayRate"     dc:"场景积分抵扣比例"`
	ExchangeRate     float64     `json:"exchangeRate"     dc:"积分汇率"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"更新时间"`
}
