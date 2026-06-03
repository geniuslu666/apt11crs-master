// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsExpChangeListInp 获取积分明细列表
type PmsExpChangeListInp struct {
	input_form.PageReq
	MemberId   int           `json:"memberId"      dc:"会员ID"`
	Des        string        `json:"des"     dc:"消费描述"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
	MemberKey  string        `json:"memberKey"     dc:"会员信息（会员名/手机号/邮箱）"`
	OrderNo    string        `json:"orderNo"     dc:"订单号"`
	OperatorId int           `json:"operatorId"      dc:"操作员ID"`
}

func (in *PmsExpChangeListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsExpChangeListModel struct {
	entity.PmsExpChange
	PmsMemberId         int    `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo   string `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName   string `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone      string `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail       string `json:"pmsMemberMail"           dc:"邮箱"`
	AdminMemberUsername string `json:"adminMemberUsername"           dc:"操作员"`
}
