package basics

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type BrokerageListReq struct {
	g.Meta       `path:"/brokerage/list" method:"post" tags:"APP_BASICS" summary:"佣金_获取佣金列表"`
	StaffId      int64  `json:"StaffId"     dc:"员工ID"`
	PageNum      int    `json:"PageNum"      dc:"页码"`
	PageSize     int    `json:"PageSize"    dc:"每页数量"`
	ChannelId    int64  `json:"ChannelId"    dc:"渠道ID"`
	RebateStatus string `json:"rebateStatus"   dc:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
}

type BrokerageListRes struct {
	List []*struct {
		Id           int64       `json:"id"             orm:"id"              description:"ID"`
		OrderAmount  float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
		Referrer     int         `json:"referrer"       orm:"referrer"        description:"推荐人"`
		RebateRate   float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
		RebateStatus string      `json:"rebateStatus"   orm:"rebate_status"   description:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
		RebateAmount float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
		OrderSn      string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
		RebateTime   *gtime.Time `json:"rebateTime"     orm:"rebate_time"     description:"分佣结算时间"`
		CreatedAt    *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	} `json:"list" dc:"佣金列表"`
	Count int `json:"count"`
}

type BrokerageDetailReq struct {
	g.Meta `path:"/brokerage/detail" method:"post" tags:"APP_BASICS" summary:"佣金_获取佣金详情"`
	ID     int64 `json:"id"  v:"required#id_unknown"   dc:"ID"`
}

type BrokerageDetailRes struct {
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

type BrokerageWithdrawBaseReq struct {
	g.Meta    `path:"/brokerage/withdraw" method:"post" tags:"APP_BASICS" summary:"佣金_提现基础信息"`
	StaffId   int64 `dc:"员工ID"`
	ChannelId int64 `dc:"渠道ID"`
}

type BrokerageWithdrawBaseRes struct {
	Balance             float64 `json:"balance" dc:"余额"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount" dc:"最小提现金额"`
	ServiceCharge       float64 `json:"serviceCharge" dc:"手续费"`
	AfterDay            int     `json:"afterDay" dc:"提现后多少天可提现"`
}

type BrokerageWithdrawReq struct {
	g.Meta    `path:"/brokerage/withdrawApply" method:"post" tags:"APP_BASICS" summary:"佣金_提现"`
	StaffId   int     `dc:"员工ID"`
	ChannelId int     `dc:"渠道ID"`
	Amount    float64 `json:"amount" dc:"提现金额"`
}

type BrokerageWithdrawRes struct {
	Amount        float64     `json:"amount" dc:"提现金额"`
	ServiceCharge float64     `json:"serviceCharge" dc:"手续费"`
	Balance       float64     `json:"balance" dc:"余额"`
	CreateTime    *gtime.Time `json:"createTime" dc:"创建时间"`
	AfterDay      int         `json:"afterDay" dc:"提现后多少天可提现"`
}

type BrokerageWithdrawListReq struct {
	g.Meta         `path:"/brokerage/withdrawList" method:"post" tags:"APP_BASICS" summary:"佣金_提现列表"`
	StaffId        int64  `dc:"员工ID"`
	PageNum        int    `json:"PageNum"      dc:"页码"`
	PageSize       int    `json:"PageSize"    dc:"每页数量"`
	ChannelId      int64  ` dc:"渠道ID"`
	WithdrawStatus string `json:"withdrawStatus" v:"in:WAIT,SUCCESS,FAIL#withdraw_status_incorrect" dc:"提现状态 不填此参数则查询全部，WAIT 提现中 SUCCESS 提现成功 FAIL 提现失败 "`
}

type BrokerageWithdrawListRes struct {
	List  []*entity.PmsWithdraw `json:"list" dc:"提现列表"`
	Count int                   `json:"count"`
}

type BrokerageWithdrawDetailReq struct {
	g.Meta `path:"/brokerage/withdrawDetail" method:"post" tags:"APP_BASICS" summary:"佣金_提现详情"`
	ID     int64 `json:"id"  v:"required#id_unknown"   dc:"ID"`
}

type BrokerageWithdrawDetailRes struct {
	*entity.PmsWithdraw
	Schedule []*BrokerageWithdrawSchedule
}

type BrokerageWithdrawSchedule struct {
	Des    string      `json:"des"`
	Time   *gtime.Time `json:"time"`
	Remark string      `json:"remark"`
}
