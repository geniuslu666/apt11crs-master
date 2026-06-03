package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/storager"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/format"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

type sBasicsAttachment struct{}

func NewBasicsAttachment() *sBasicsAttachment {
	return &sBasicsAttachment{}
}

func init() {
	service.RegisterBasicsAttachment(NewBasicsAttachment())
}

func (s *sBasicsAttachment) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysAttachment.Ctx(ctx), option...)
}

func (s *sBasicsAttachment) Delete(ctx context.Context, in *input_basics.AttachmentDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除附件失败，请稍后重试！")
	}
	return
}

func (s *sBasicsAttachment) View(ctx context.Context, in *input_basics.AttachmentViewInp) (res *input_basics.AttachmentViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取附件信息失败，请稍后重试！")
	}
	return
}

func (s *sBasicsAttachment) List(ctx context.Context, in *input_basics.AttachmentListInp) (list []*input_basics.AttachmentListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if in.Drive != "" {
		mod = mod.Where(dao.SysAttachment.Columns().Drive, in.Drive)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.SysAttachment.Columns().Name, "%"+in.Name+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.SysAttachment.Columns().Status, in.Status)
	}

	if len(in.UpdatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysAttachment.Columns().UpdatedAt, in.UpdatedAt[0], in.UpdatedAt[1])
	}

	if in.Kind != "" {
		mod = mod.Where(dao.SysAttachment.Columns().Kind, in.Kind)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取附件数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).OrderDesc(dao.SysAttachment.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取附件列表失败！")
		return
	}

	for _, v := range list {
		v.SizeFormat = format.FileSize(v.Size)
		v.FileUrl = storager.LastUrl(ctx, v.FileUrl, v.Drive)

		// 获取URL路径部分
		urlParts := strings.Split(v.FileUrl, "/")
		fileName := urlParts[len(urlParts)-1]

		// 处理文件名
		fileNameWithoutExt := gfile.Name(fileName)
		newFileName := fileNameWithoutExt + "_MID.jpg"

		// 构建新URL
		urlParts[len(urlParts)-1] = newFileName
		v.MidFileUrl = strings.Join(urlParts, "/")

	}
	return
}

func (s *sBasicsAttachment) ClearKind(ctx context.Context, in *input_basics.AttachmentClearKindInp) (err error) {
	memberId := contexts.GetUserId(ctx)
	if _, err = s.Model(ctx).Where(dao.SysAttachment.Columns().MemberId, memberId).Where(dao.SysAttachment.Columns().Kind, in.Kind).Delete(); err != nil {
		err = gerror.Wrap(err, "删除附件上传类型失败，请稍后重试！")
	}
	return
}
