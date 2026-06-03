package basics

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ChannelFxCenterReq struct {
	g.Meta    `path:"/channel/fxCenter" method:"post" tags:"APP_BASICS" summary:"分销_员工分销中心信息"`
	ChannelId int64 `v:"required#channel_id_unknown" dc:"渠道ID"`
}

type ChannelFxCenterRes struct {
	Brokerage         float64 `json:"brokerage" dc:"可提现佣金"`
	AllBrokerage      float64 `json:"allBrokerage" dc:"总佣金"`
	SettlingBrokerage float64 `json:"settlingBrokerage" dc:"结算中佣金"`
	SettlingWithdraw  float64 `json:"settlingWithdraw" dc:"提现中"`
	Withdraw          float64 `json:"withdraw" dc:"已提现"`
	BrokerageList     []*struct {
		Id           int64       `json:"id"             orm:"id"              description:"ID"`
		OrderAmount  float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
		Referrer     int         `json:"referrer"       orm:"referrer"        description:"推荐人"`
		RebateRate   float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
		RebateStatus string      `json:"rebateStatus"   orm:"rebate_status"   description:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
		RebateAmount float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
		OrderSn      string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
		RebateTime   *gtime.Time `json:"rebateTime"     orm:"rebate_time"     description:"分佣结算时间"`
		CreatedAt    *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	}
}
