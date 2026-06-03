// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_th"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IAppMember interface {
		MemberList(ctx context.Context, in *input_app_member.PmsMemberListInp) (list []*input_app_member.PmsMemberListModel, totalCount int, err error)
		MemberSelectList(ctx context.Context, in *input_app_member.PmsMemberSelectInp) (list []*input_app_member.PmsMemberListModel, err error)
		MemberExport(ctx context.Context, in *input_app_member.PmsMemberListInp) (err error)
		MemberEdit(ctx context.Context, in *input_app_member.PmsMemberEditInp) (err error)
		MemberBaseEdit(ctx context.Context, in *input_app_member.PmsMemberBaseEditInp) (err error)
		MemberBalanceEdit(ctx context.Context, in *input_app_member.PmsMemberBalanceEditInp) (err error)
		MemberExpEdit(ctx context.Context, in *input_app_member.PmsMemberExpEditInp) (err error)
		MemberDelete(ctx context.Context, in *input_app_member.PmsMemberDeleteInp) (err error)
		MemberView(ctx context.Context, in *input_app_member.PmsMemberViewInp) (res *input_app_member.PmsMemberViewModel, err error)
		MemberAddExp(ctx context.Context, MemberId int, Exp int, Txt string) (err error)
		MemberStat(ctx context.Context, in *input_app_member.PmsMemberStatInp) (out *input_app_member.PmsMemberStatModel, err error)
		// MemberBalanceChange 会员余额变动
		MemberBalanceChange(ctx context.Context, in *input_app_member.MemberBalanceInp, tx gdb.TX) (err error)
		// MemberInfo 查询用户信息
		MemberInfo(ctx context.Context, MemberId int) (MemberInfo *entity.PmsMember, err error)
		// Status 修改会员状态
		Status(ctx context.Context, in *input_app_member.PmsMemberStatusInp) (err error)
		// Cancel 会员注销
		Cancel(ctx context.Context, in *input_app_member.PmsMemberCancelInp) (err error)
		RegisterMemberAward(ctx context.Context, TX gdb.TX, MemberId int) (err error)
		BalanceChangeStat(ctx context.Context, in *input_app_member.PmsBalanceChangeStatInp) (res *input_app_member.PmsBalanceChangeStatModel, err error)
		BalanceChangeList(ctx context.Context, in *input_app_member.PmsBalanceChangeListInp) (list []*input_app_member.PmsBalanceChangeListModel, totalCount int, err error)
		// MemberCancelList 会员注销列表
		MemberCancelList(ctx context.Context, in *input_app_member.PmsMemberCancelListInp) (list []*input_app_member.PmsMemberCancelListModel, totalCount int, err error)
		// MemberCancelAgree 同意注销
		MemberCancelAgree(ctx context.Context, in *input_app_member.PmsMemberCancelAgreeInp) (err error)
		// MemberCancelDisagree 驳回注销
		MemberCancelDisagree(ctx context.Context, in *input_app_member.PmsMemberCancelDisagreeInp) (err error)
		ExpChangeList(ctx context.Context, in *input_app_member.PmsExpChangeListInp) (list []*input_app_member.PmsExpChangeListModel, totalCount int, err error)
		MemberGroupList(ctx context.Context, in *input_app_member.PmsMemberGroupListInp) (list []*input_app_member.PmsMemberGroupListModel, totalCount int, err error)
		MemberGroupExport(ctx context.Context, in *input_app_member.PmsMemberGroupListInp) (err error)
		MemberGroupEdit(ctx context.Context, in *input_app_member.PmsMemberGroupEditInp) (err error)
		MemberGroupDelete(ctx context.Context, in *input_app_member.PmsMemberGroupDeleteInp) (err error)
		MemberGroupView(ctx context.Context, in *input_app_member.PmsMemberGroupViewInp) (res *input_app_member.PmsMemberGroupViewModel, err error)
		MemberGroupAllList(ctx context.Context, in *input_app_member.PmsMemberGroupAllInp) (list []*input_app_member.PmsMemberGroupAllModel, err error)
		MemberIntentionList(ctx context.Context, in *input_app_member.PmsMemberIntentionListInp) (list []*input_app_member.PmsMemberIntentionListModel, totalCount int, err error)
		MemberIntentionEdit(ctx context.Context, in *input_app_member.PmsMemberIntentionEditInp) (err error)
		MemberIntentionExport(ctx context.Context, in *input_app_member.PmsMemberIntentionListInp) (err error)
		MemberIntentionView(ctx context.Context, in *input_app_member.PmsMemberIntentionViewInp) (res *input_app_member.PmsMemberIntentionViewModel, err error)
		MemberLevelList(ctx context.Context, in *input_app_member.PmsMemberLevelListInp) (list []*input_app_member.PmsMemberLevelListModel, totalCount int, err error)
		MemberLevelAllList(ctx context.Context, in *input_app_member.PmsMemberLevelAllInp) (list []*input_app_member.PmsMemberLevelAllModel, err error)
		MemberLevelEdit(ctx context.Context, in *input_app_member.PmsMemberLevelEditInp) (err error)
		MemberLevelDelete(ctx context.Context, in *input_app_member.PmsMemberLevelDeleteInp) (err error)
		MemberLevelView(ctx context.Context, in *input_app_member.PmsMemberLevelViewInp) (res *input_app_member.PmsMemberLevelViewModel, err error)
		MemberLevelComputeLevel(ctx context.Context, MemberId int, tx gdb.TX) (err error)
		MemberLogList(ctx context.Context, in *input_app_member.PmsMemberLogListInp) (list []*input_app_member.PmsMemberLogListModel, totalCount int, err error)
		MemberLogExport(ctx context.Context, in *input_app_member.PmsMemberLogListInp) (err error)
		MemberLogView(ctx context.Context, in *input_app_member.PmsMemberLogViewInp) (res *input_app_member.PmsMemberLogViewModel, err error)
		MemberSceneList(ctx context.Context, in *input_app_member.PmsMemberSceneListInp) (list []*input_app_member.PmsMemberSceneListModel, totalCount int, err error)
		MemberSceneEdit(ctx context.Context, in *input_app_member.PmsMemberSceneEditInp) (err error)
		MemberSceneDelete(ctx context.Context, in *input_app_member.PmsMemberSceneDeleteInp) (err error)
		MemberSceneView(ctx context.Context, in *input_app_member.PmsMemberSceneViewInp) (res *input_app_member.PmsMemberSceneViewModel, err error)
		MemberSceneSwitch(ctx context.Context, in *input_app_member.PmsMemberSceneSwitchInp) (err error)
		// GetMemberQrCode 获取会员二维码
		GetMemberQrCode(ctx context.Context) (code string, err error)
		// VerifyMemberCode 会员码验证（发券+核销）
		VerifyMemberCode(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error)
	}
)

var (
	localAppMember IAppMember
)

func AppMember() IAppMember {
	if localAppMember == nil {
		panic("implement not found for interface IAppMember, forgot register?")
	}
	return localAppMember
}

func RegisterAppMember(i IAppMember) {
	localAppMember = i
}
