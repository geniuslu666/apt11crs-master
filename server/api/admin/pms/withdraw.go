package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type WithdrawStaffListReq struct {
	g.Meta `path:"/pmsWithdraw/staffList" method:"get" tags:"ADMIN_PMS" summary:"员工提现申请_列表"`
	input_basics.PmsWithdrawListInp
}

type WithdrawStaffListRes struct {
	input_form.PageRes
	List []*input_basics.PmsWithdrawListModel `json:"list"   dc:"数据列表"`
}

type WithdrawChannelListReq struct {
	g.Meta `path:"/pmsWithdraw/channelList" method:"get" tags:"ADMIN_PMS" summary:"渠道提现申请_列表"`
	input_basics.PmsWithdrawListInp
}

type WithdrawChannelListRes struct {
	input_form.PageRes
	List []*input_basics.PmsWithdrawListModel `json:"list"   dc:"数据列表"`
}

type WithdrawViewReq struct {
	g.Meta `path:"/pmsWithdraw/view" method:"get" tags:"ADMIN_PMS" summary:"提现申请_详情"`
	input_basics.PmsWithdrawViewInp
}

type WithdrawViewRes struct {
	*input_basics.PmsWithdrawViewModel
}

type WithdrawAgreeStaffReq struct {
	g.Meta `path:"/pmsWithdraw/agreeStaff" method:"post" tags:"ADMIN_PMS" summary:"提现申请_同意员工提现"`
	input_basics.PmsWithdrawAgreeInp
}

type WithdrawAgreeStaffRes struct {
}

type WithdrawDisagreeStaffReq struct {
	g.Meta `path:"/pmsWithdraw/disagreeStaff" method:"post" tags:"ADMIN_PMS" summary:"提现申请_拒绝员工提现"`
	input_basics.PmsWithdrawDisagreeInp
}

type WithdrawDisagreeStaffRes struct {
}

type WithdrawAgreeChannelReq struct {
	g.Meta `path:"/pmsWithdraw/agreeChannel" method:"post" tags:"ADMIN_PMS" summary:"提现申请_同意渠道提现"`
	input_basics.PmsWithdrawAgreeInp
}

type WithdrawAgreeChannelRes struct {
}

type WithdrawDisagreeChannelReq struct {
	g.Meta `path:"/pmsWithdraw/disagreeChannel" method:"post" tags:"ADMIN_PMS" summary:"提现申请_拒绝渠道提现"`
	input_basics.PmsWithdrawDisagreeInp
}

type WithdrawDisagreeChannelRes struct {
}

type WithdrawTransferStaffReq struct {
	g.Meta `path:"/pmsWithdraw/transferStaff" method:"post" tags:"ADMIN_PMS" summary:"提现申请_员工提现转账"`
	input_basics.PmsWithdrawTransferInp
}

type WithdrawTransferStaffRes struct {
}

type WithdrawTransferChannelReq struct {
	g.Meta `path:"/pmsWithdraw/transferChannel" method:"post" tags:"ADMIN_PMS" summary:"提现申请_渠道提现转账"`
	input_basics.PmsWithdrawTransferInp
}

type WithdrawTransferChannelRes struct {
}
