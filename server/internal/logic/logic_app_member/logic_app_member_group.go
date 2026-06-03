package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_app_member"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sAppMember) MemberGroupList(ctx context.Context, in *input_app_member.PmsMemberGroupListInp) (list []*input_app_member.PmsMemberGroupListModel, totalCount int, err error) {
	mod := dao.PmsMemberGroup.Ctx(ctx)

	mod = mod.Fields(input_app_member.PmsMemberGroupListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsMemberGroup.Columns().Id, in.Id)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMemberGroup.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsMemberGroup.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员分组列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberGroupExport(ctx context.Context, in *input_app_member.PmsMemberGroupListInp) (err error) {
	list, _, err := s.MemberGroupList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_app_member.PmsMemberGroupExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出会员分组-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("会员分组")
		exports   []input_app_member.PmsMemberGroupExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sAppMember) MemberGroupEdit(ctx context.Context, in *input_app_member.PmsMemberGroupEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsMemberGroup.Ctx(ctx).
				Fields(input_app_member.PmsMemberGroupUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员分组失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsMemberGroup.Ctx(ctx).
			Fields(input_app_member.PmsMemberGroupInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员分组失败，请稍后重试！")
		}
		return
	})
}

func (s *sAppMember) MemberGroupDelete(ctx context.Context, in *input_app_member.PmsMemberGroupDeleteInp) (err error) {

	if _, err = dao.PmsMemberGroup.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除会员分组失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberGroupView(ctx context.Context, in *input_app_member.PmsMemberGroupViewInp) (res *input_app_member.PmsMemberGroupViewModel, err error) {
	if err = dao.PmsMemberGroup.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员分组信息，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberGroupAllList(ctx context.Context, in *input_app_member.PmsMemberGroupAllInp) (list []*input_app_member.PmsMemberGroupAllModel, err error) {
	mod := dao.PmsMemberGroup.Ctx(ctx).WithAll()

	mod = mod.Fields(input_app_member.PmsMemberGroupAllModel{})

	if in.MemberGroup != "" {
		mod = mod.WhereLike(dao.PmsMemberGroup.Columns().MemberGroup, in.MemberGroup)
	}

	mod = mod.OrderDesc(dao.PmsMemberGroup.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取会员分组列表失败，请稍后重试！")
		return
	}
	return
}
