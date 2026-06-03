package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsAttachmentKind struct{}

func NewBasicsAttachmentKind() *sBasicsAttachmentKind {
	return &sBasicsAttachmentKind{}
}

func init() {
	service.RegisterBasicsAttachmentKind(NewBasicsAttachmentKind())
}

func (s *sBasicsAttachmentKind) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysAttachmentKind.Ctx(ctx), option...)
}

func (s *sBasicsAttachmentKind) List(ctx context.Context, in *input_basics.SysAttachmentKindListInp) (list []*input_basics.SysAttachmentKindListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.SysAttachmentKindListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.SysAttachmentKind.Columns().Id, in.Id)
	}

	if in.Label != "" {
		mod = mod.WhereLike(dao.SysAttachmentKind.Columns().Label, in.Label)
	}

	if in.Key != "" {
		mod = mod.WhereLike(dao.SysAttachmentKind.Columns().Key, in.Key)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.SysAttachmentKind.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取附件分类列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsAttachmentKind) Edit(ctx context.Context, in *input_basics.SysAttachmentKindEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.SysAttachmentKindUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改附件分类失败，请稍后重试！")
			}
			return
		}
		Sn := guid.S()
		if g.IsEmpty(in.Key) {
			in.Key = Sn
		}
		if g.IsEmpty(in.Value) {
			in.Value = Sn
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.SysAttachmentKindInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增附件分类失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsAttachmentKind) Delete(ctx context.Context, in *input_basics.SysAttachmentKindDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除附件分类失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsAttachmentKind) View(ctx context.Context, in *input_basics.SysAttachmentKindViewInp) (res *input_basics.SysAttachmentKindViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取附件分类信息，请稍后重试！")
		return
	}
	return
}
