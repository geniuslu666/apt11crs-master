// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberCancelUpdateFields 修改会员信息字段过滤
type PmsMemberCancelUpdateFields struct {
	MemberId     int    `json:"memberId"     dc:"用户ID"`
	MemberNo     string `json:"memberNo"     dc:"会员号"`
	Phone        string `json:"phone"        dc:"手机号"`
	PhoneArea    string `json:"phoneArea"    dc:"手机号区号"`
	Mail         string `json:"mail"         dc:"邮箱"`
	CancelReason string `json:"cancelReason" dc:"注销原因"`
	OperatorId   int    `json:"operatorId"   dc:"操作人ID"`
	CreateAt     string `json:"createAt"     dc:"创建时间"`
	UpdateAt     string `json:"updateAt"     dc:"更新时间"`
}

// PmsMemberCancelInsertFields 新增会员信息字段过滤
type PmsMemberCancelInsertFields struct {
	MemberId     int    `json:"memberId"     dc:"用户ID"`
	MemberNo     string `json:"memberNo"     dc:"会员号"`
	Phone        string `json:"phone"        dc:"手机号"`
	PhoneArea    string `json:"phoneArea"    dc:"手机号区号"`
	Mail         string `json:"mail"         dc:"邮箱"`
	CancelReason string `json:"cancelReason" dc:"注销原因"`
	OperatorId   int    `json:"operatorId"   dc:"操作人ID"`
	CreateAt     string `json:"createAt"     dc:"创建时间"`
	UpdateAt     string `json:"updateAt"     dc:"更新时间"`
}

// PmsMemberCancelListInp 获取会员注销列表
type PmsMemberCancelListInp struct {
	input_form.PageReq
	Id          int           `json:"id"          dc:"主键"`
	MemberNo    string        `json:"memberNo"    dc:"会员号"`
	Phone       string        `json:"phone"       dc:"手机号"`
	Mail        string        `json:"mail"        dc:"邮箱"`
	AuditStatus string        `json:"auditStatus" dc:"审核状态"`
	CreatedAt   []*gtime.Time `json:"createdAt" dc:"创建时间"`
	AuditTime   []*gtime.Time `json:"auditTime" dc:"审核时间"`
}

func (in *PmsMemberCancelListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberCancelListModel struct {
	entity.PmsMemberCancel
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"           dc:"id"`
		MemberNo   string `json:"memberNo"     dc:"会员号"`
		Phone      string `json:"phone"        dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    dc:"手机号区号"`
		Mail       string `json:"mail"         dc:"邮箱"`
		FullName   string `json:"fullName"     dc:"全名"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
	AdminMemberUsername string `json:"adminMemberUsername"           dc:"操作员"`
}

// PmsMemberCancelAgreeFields 同意注销字段过滤
type PmsMemberCancelAgreeFields struct {
	AuditStatus int         `json:"auditStatus"   dc:"申请状态"`
	OperatorId  int         `json:"operatorId"    dc:"操作员ID"`
	AuditReason string      `json:"auditReason"   dc:"审核备注"`
	AuditTime   *gtime.Time `json:"auditTime"     dc:"审核时间"`
}

// PmsMemberCancelAgreeInp 同意注销
type PmsMemberCancelAgreeInp struct {
	Id          interface{} `json:"id" v:"required#id不能为空" dc:"id"`
	AuditReason string      `json:"auditReason"           dc:"审核通过/拒绝原因"`
}

func (in *PmsMemberCancelAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawAgreeModel struct{}

// PmsMemberCancelDisagreeFields 拒绝注销字段过滤
type PmsMemberCancelDisagreeFields struct {
	AuditStatus int         `json:"auditStatus"   dc:"申请状态"`
	OperatorId  int         `json:"operatorId"    dc:"操作员ID"`
	AuditReason string      `json:"auditReason"   dc:"审核备注"`
	AuditTime   *gtime.Time `json:"auditTime"     dc:"审核时间"`
}

// PmsMemberCancelDisagreeInp 拒绝注销
type PmsMemberCancelDisagreeInp struct {
	Id          interface{} `json:"id" v:"required#id不能为空" dc:"id"`
	AuditReason string      `json:"auditReason"           dc:"审核通过/拒绝原因"`
}

func (in *PmsMemberCancelDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberCancelDisagreeModel struct{}
