package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/model/input/input_app_member"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sAppMember) MemberLogList(ctx context.Context, in *input_app_member.PmsMemberLogListInp) (list []*input_app_member.PmsMemberLogListModel, totalCount int, err error) {
	mod := dao.PmsMemberLog.Ctx(ctx)

	mod = mod.FieldsPrefix(dao.PmsMemberLog.Table(), input_app_member.PmsMemberLogListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsMemberLogListModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.InnerJoinOnFields(dao.PmsMember.Table(), dao.PmsMemberLog.Columns().MemberId, "=", dao.PmsMember.Columns().Id)

	if in.Id > 0 {
		mod = mod.Where(dao.PmsMemberLog.Columns().Id, in.Id)
	}

	if len(in.LoginTime) == 2 {
		mod = mod.WhereBetween(dao.PmsMemberLog.Columns().LoginTime, in.LoginTime[0], in.LoginTime[1])
	}

	if in.LoginType != "" {
		mod = mod.Where(dao.PmsMemberLog.Columns().LoginType, in.LoginType)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMemberLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.PmsMemberMemberNo != "" {
		mod = mod.WherePrefix(dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo, in.PmsMemberMemberNo)
	}

	if in.PmsMemberFullName != "" {
		mod = mod.WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().FullName, in.PmsMemberFullName)
	}

	if in.PmsMemberPhone != "" {
		mod = mod.WherePrefix(dao.PmsMember.Table(), dao.PmsMember.Columns().Phone, in.PmsMemberPhone)
	}

	if in.PmsMemberPhoneArea != "" {
		mod = mod.WherePrefix(dao.PmsMember.Table(), dao.PmsMember.Columns().PhoneArea, in.PmsMemberPhoneArea)
	}

	if in.PmsMemberMail != "" {
		mod = mod.WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Mail, in.PmsMemberMail)
	}

	if in.PmsMemberSource != "" {
		mod = mod.WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Source, in.PmsMemberSource)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsMemberLog.Table() + "." + dao.PmsMemberLog.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员登录日志列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberLogExport(ctx context.Context, in *input_app_member.PmsMemberLogListInp) (err error) {
	list, _, err := s.MemberLogList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_app_member.PmsMemberLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出会员登录日志-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("会员登录日志")
		exports   []input_app_member.PmsMemberLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sAppMember) MemberLogView(ctx context.Context, in *input_app_member.PmsMemberLogViewInp) (res *input_app_member.PmsMemberLogViewModel, err error) {
	if err = dao.PmsMemberLog.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员登录日志信息，请稍后重试！")
		return
	}
	return
}
