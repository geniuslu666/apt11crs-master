package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type MemberListReq struct {
	g.Meta `path:"/pmsMember/list" method:"get" tags:"ADMIN_PMS" summary:"会员信息_列表"`
	input_app_member.PmsMemberListInp
}

type MemberListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberListModel `json:"list"   dc:"数据列表"`
}

type MemberSelectReq struct {
	g.Meta `path:"/pmsMember/select" method:"get" tags:"ADMIN_PMS" summary:"会员信息_选择列表"`
	input_app_member.PmsMemberSelectInp
}

type MemberSelectRes struct {
	List []*input_app_member.PmsMemberListModel `json:"list"   dc:"数据列表"`
}

type MemberExportReq struct {
	g.Meta `path:"/pmsMember/export" method:"get" tags:"ADMIN_PMS" summary:"会员信息_导出"`
	input_app_member.PmsMemberListInp
}

type MemberExportRes struct{}

type MemberViewReq struct {
	g.Meta `path:"/pmsMember/view" method:"get" tags:"ADMIN_PMS" summary:"会员信息_详情"`
	input_app_member.PmsMemberViewInp
}

type MemberViewRes struct {
	*input_app_member.PmsMemberViewModel
}

type MemberEditReq struct {
	g.Meta `path:"/pmsMember/edit" method:"post" tags:"ADMIN_PMS" summary:"会员信息_编辑"`
	input_app_member.PmsMemberEditInp
}

type MemberEditRes struct{}

type MemberBaseEditReq struct {
	g.Meta `path:"/pmsMember/baseEdit" method:"post" tags:"ADMIN_PMS" summary:"会员信息_编辑"`
	input_app_member.PmsMemberBaseEditInp
}

type MemberBaseEditRes struct{}

type MemberBalanceEditReq struct {
	g.Meta `path:"/pmsMember/balanceEdit" method:"post" tags:"ADMIN_PMS" summary:"会员信息_调整会员积分"`
	input_app_member.PmsMemberBalanceEditInp
}

type MemberBalanceEditRes struct{}

type MemberExpEditReq struct {
	g.Meta `path:"/pmsMember/expEdit" method:"post" tags:"ADMIN_PMS" summary:"会员信息_调整会员成长值"`
	input_app_member.PmsMemberExpEditInp
}

type MemberExpEditRes struct{}

type MemberDeleteReq struct {
	g.Meta `path:"/pmsMember/delete" method:"post" tags:"ADMIN_PMS" summary:"会员信息_删除"`
	input_app_member.PmsMemberDeleteInp
}

type MemberDeleteRes struct{}

type MemberSendMsgReq struct {
	g.Meta   `path:"/pmsMember/sendMsg" method:"post" tags:"ADMIN_PMS" summary:"会员信息_发送通知"`
	MemberId int    `json:"member_id" v:"required#请填入会员ID" dc:"会员ID"`
	Title    string `json:"title" dc:"标题"`
	Content  string `json:"content" v:"required#请填入内容" dc:"内容"`
}

type MemberSendMsgRes struct {
}

type MemberSendEmailReq struct {
	g.Meta  `path:"/pmsMember/sendEmail" method:"post" tags:"ADMIN_PMS" summary:"会员信息_发送邮件"`
	Email   string `json:"email" v:"required#请填入邮箱" dc:"邮箱"`
	Content string `json:"content" v:"required#请填入发送内容" dc:"发送内容"`
}
type MemberSendEmailRes struct{}

type MemberSendNotifyReq struct {
	g.Meta   `path:"/pmsMember/sendNotify" method:"post" tags:"ADMIN_PMS" summary:"会员信息_发送通知"`
	MemberId int    `json:"member_id" v:"required#请填入会员ID" dc:"会员ID"`
	Type     string `json:"type" dc:"类型-通知类型 SYSTEM、系统消息    BOOKING、预定消息"`
	Title    string `json:"title" dc:"标题"`
	Content  string `json:"content" v:"required#请填入内容" dc:"内容"`
}

type MemberSendNotifyRes struct {
}

type MemberSendSmsReq struct {
	g.Meta  `path:"/pmsMember/sendSms" method:"post" tags:"ADMIN_PMS" summary:"会员信息_发送短信"`
	Phone   string `json:"phone" v:"required#请填入手机号" dc:"手机号"`
	Content string `json:"content" v:"required#请填入发送内容" dc:"发送内容"`
	AreaNo  string `json:"area_no" v:"required#请填入区号" dc:"区号"`
}

type MemberSendSmsRes struct {
}

type MemberStatReq struct {
	g.Meta `path:"/pmsMember/stat" method:"get" tags:"ADMIN_PMS" summary:"会员信息_会员概况信息"`
	input_app_member.PmsMemberStatInp
}

type MemberStatRes struct {
	*input_app_member.PmsMemberStatModel
}

type MemberStatusReq struct {
	g.Meta `path:"/pmsMember/status" method:"post" tags:"ADMIN_PMS" summary:"会员信息_更新会员状态"`
	input_app_member.PmsMemberStatusInp
}

type MemberStatusRes struct{}

// MemberCancelReq 注销
type MemberCancelReq struct {
	g.Meta `path:"/pmsMember/cancel" method:"post" tags:"ADMIN_PMS" summary:"会员信息_注销"`
	input_app_member.PmsMemberCancelInp
}

type MemberCancelRes struct {
}

type MemberSendRegisterCodeReq struct {
	g.Meta `path:"/pmsMember/sendRegisterCode" method:"post" tags:"APP_MEMBER" summary:"会员信息_发送验证码"`
	Type   int    `json:"type" v:"required#请选择验证码发送类型" dc:"验证码发送类型-1短信 2邮箱"`
	Phone  string `json:"phone"    dc:"手机号"`
	AreaNo string `json:"areaNo"  dc:"手机区号-例如国内+86 传入+86"`
	EMail  string `json:"email"    dc:"邮箱"`
	// 仅短信发送时需要传入 GeeTest GT4 校验结果
	CaptchaOutput string `json:"captcha_output" dc:"极验输出信息"`
	LotNumber     string `json:"lot_number"     dc:"极验验证流水号"`
	PassToken     string `json:"pass_token"     dc:"极验验证通过标识"`
	GenTime       string `json:"gen_time"       dc:"极验验证通过时间戳"`
}

type MemberSendRegisterCodeRes struct{}
