package input_basics

import (
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceStatInp 获取财务概况
type FinanceStatInp struct {
	SearchTime []*gtime.Time `json:"searchTime" dc:"时间搜索"`
}

type FinanceStatModel struct {
	OverviewStat struct {
		ExpectIncome float64 `json:"expectIncome"           dc:"预计收入"`
		TotalIncome  float64 `json:"totalIncome"           dc:"收入总额"`
		TotalExpend  float64 `json:"totalExpend"      dc:"支出总额"`
	} `json:"overviewStat" dc:"财务概况"`
	IncomeStat struct {
		HomeStayOrderIncome float64 `json:"HomeStayOrderIncome"           dc:"民宿订单收入"`
		FoodOrderIncome     float64 `json:"foodOrderIncome"           dc:"餐饮收入"`
		CarOrderIncome      float64 `json:"carOrderIncome"           dc:"车队订单收入"`
		MassageOrderIncome  float64 `json:"massageOrderIncome"           dc:"按摩订单收入"`
	} `json:"incomeStat" dc:"收入概况"`
	ExpendStat struct {
		OrderRefund     float64 `json:"orderRefund"           description:"订单退款"`
		StaffWithdraw   float64 `json:"staffWithdraw"    description:"员工提现"`
		ChannelWithdraw float64 `json:"channelWithdraw"    description:"渠道提现"`
	} `json:"expendStat" dc:"支出概况"`
}

type FinanceListInp struct {
	input_form.PageReq
	PayChannel string        `json:"payChannel" v:"in:SYSTEM,PAYCLOUD,PAYPAL,STRIPE" dc:"支付渠道  SYSTEM PAYCLOUD PAYPAL STRIPE"`
	PayType    string        `json:"payType" v:"in:CreditLinkPay,WeChatMiniPay,COUPON,BAL,WeChatPay,Alipay+,Paypal,PaypalCard,StripeCard" dc:"支付类型 CreditLinkPay,WeChatMiniPay,COUPON,BAL,WeChatPay,Alipay+,Paypal,PaypalCard,StripeCard"`
	TimeRange  []*gtime.Time `json:"timeRange" dc:"时间区间"`
	StartTime  string        `json:"startTime" v:"date#日期格式错误" dc:"开始时间"`
	EndTime    string        `json:"endTime" v:"date#日期格式错误" dc:"结束时间"`
	Type       string        `json:"type" v:"in:pay,refund" dc:"pay:支付 refund:退款"`
	OrderSn    string        `json:"orderSn" dc:"系统订单编号"`
}

type FinanceListModel struct {
	PayListItem       FinanceListItem `json:"pay_list_item" dc:"支付列表"`
	PayCount          int             `json:"pay_count" dc:"支付总数"`
	RefundListItem    FinanceListItem `json:"refund_list_item" dc:"退款列表"`
	RefundCount       int             `json:"refund_count" dc:"退款总数"`
	Count             int             `json:"total" dc:"总数"`
	TransactionAmount float64         `json:"transactionAmount" dc:"总交易量"`
	PayAmount         float64         `json:"payAmount" dc:"总支付金额"`
	RefundAmount      float64         `json:"refundAmount" dc:"总退款金额"`
}

type FinanceListItem struct {
	input_form.PageRes
	List []*FinanceListModelItem `json:"list" dc:"列表"`
}

type FinanceListModelItem struct {
	Id         int     `json:"id" dc:"ID"`
	Scene      string  `json:"Scene" dc:"场景"`
	OrderSn    string  `json:"orderSn" dc:"订单编号"`
	TypeSn     string  `json:"TypeSn" dc:"支付/退款单号"`
	Amount     float64 `json:"amount" dc:"金额"`
	PayChannel string  `json:"payChannel" dc:"支付渠道"`
	PayType    string  `json:"payType" dc:"支付类型"`
	PayTime    string  `json:"payTime" dc:"支付时间"`
}

type FinancePayExportModel struct {
	Id         int     `json:"id" dc:"ID"`
	Scene      string  `json:"Scene" dc:"场景"`
	OrderSn    string  `json:"orderSn" dc:"订单编号"`
	TypeSn     string  `json:"TypeSn" dc:"支付单号"`
	PayChannel string  `json:"payChannel" dc:"支付平台"`
	PayType    string  `json:"payType" dc:"支付类型"`
	Amount     float64 `json:"amount" dc:"金额"`
	PayTime    string  `json:"payTime" dc:"支付时间"`
}

type FinanceRefundExportModel struct {
	Id         int     `json:"id" dc:"ID"`
	Scene      string  `json:"Scene" dc:"场景"`
	OrderSn    string  `json:"orderSn" dc:"订单编号"`
	TypeSn     string  `json:"TypeSn" dc:"退款单号"`
	PayChannel string  `json:"payChannel" dc:"支付平台"`
	PayType    string  `json:"payType" dc:"支付类型"`
	Amount     float64 `json:"amount" dc:"金额"`
	PayTime    string  `json:"payTime" dc:"退款时间"`
}
