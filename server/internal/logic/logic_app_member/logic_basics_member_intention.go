package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_app_member"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sAppMember) MemberIntentionList(ctx context.Context, in *input_app_member.PmsMemberIntentionListInp) (list []*input_app_member.PmsMemberIntentionListModel, totalCount int, err error) {
	mod := dao.PmsMemberIntention.Ctx(ctx).WithAll()

	mod = mod.Fields(input_app_member.PmsMemberIntentionListModel{})

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsMemberIntention.Columns().Name, in.Name)
	}

	if in.Sex > 0 {
		mod = mod.Where(dao.PmsMemberIntention.Columns().Sex, in.Sex)
	}

	if in.Phone != "" {
		mod = mod.WhereLike(dao.PmsMemberIntention.Columns().Phone, in.Phone)
	}

	if in.Mail != "" {
		mod = mod.WhereLike(dao.PmsMemberIntention.Columns().Mail, in.Mail)
	}

	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMemberIntention.Columns().CreateAt, in.CreateAt[0], in.CreateAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsMemberIntention.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员意向表列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberIntentionEdit(ctx context.Context, in *input_app_member.PmsMemberIntentionEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsMemberIntention.Ctx(ctx).
				Fields(input_app_member.PmsMemberIntentionUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员意向表失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsMemberIntention.Ctx(ctx).
			Fields(input_app_member.PmsMemberIntentionInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员意向表失败，请稍后重试！")
		}
		return
	})
}

func (s *sAppMember) MemberIntentionExport(ctx context.Context, in *input_app_member.PmsMemberIntentionListInp) (err error) {
	list, _, err := s.MemberIntentionList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_app_member.PmsMemberIntentionExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出会员意向表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("会员意向表")
		exports   []input_app_member.PmsMemberIntentionExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sAppMember) MemberIntentionView(ctx context.Context, in *input_app_member.PmsMemberIntentionViewInp) (res *input_app_member.PmsMemberIntentionViewModel, err error) {
	if err = dao.PmsMemberIntention.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员意向表信息，请稍后重试！")
		return
	}
	return
}
