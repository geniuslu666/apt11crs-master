// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBalanceChangeStatInp 获取积分概况信息
type PmsBalanceChangeStatInp struct {
}

type PmsBalanceChangeStatModel struct {
	TotalSendPoints    float64 `json:"totalSendPoints"        dc:"累计发放积分"`
	TotalConsumePoints float64 `json:"totalConsumePoints"        dc:"累计消耗积分"`
	Points             float64 `json:"points"        dc:"累计消耗积分"`
}

// PmsBalanceChangeListInp 获取积分明细列表
type PmsBalanceChangeListInp struct {
	input_form.PageReq
	MemberId   int           `json:"memberId"      dc:"会员ID"`
	Des        string        `json:"des"     dc:"消费描述"`
	Reason     string        `json:"reason"     dc:"原因"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
	MemberKey  string        `json:"memberKey"     dc:"会员信息（会员名/手机号/邮箱）"`
	Direction  int           `json:"direction"      dc:"方向：1-发放， 2-消耗"`
	OrderNo    string        `json:"orderNo"     dc:"订单号"`
	OperatorId int           `json:"operatorId"      dc:"操作员ID"`
}

type PmsBalanceChangeListModel struct {
	entity.PmsBalanceChange
	PmsMemberId         int    `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo   string `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName   string `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone      string `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail       string `json:"pmsMemberMail"           dc:"邮箱"`
	AdminMemberUsername string `json:"adminMemberUsername"           dc:"操作员"`
}
