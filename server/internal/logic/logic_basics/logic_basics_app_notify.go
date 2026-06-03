package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsAppNotify struct{}

func NewBasicsAppNotify() *sBasicsAppNotify {
	return &sBasicsAppNotify{}
}

func init() {
	service.RegisterBasicsAppNotify(NewBasicsAppNotify())
}

func (s *sBasicsAppNotify) NotifyList(ctx context.Context, in *input_basics.PmsNotifyListInp) (list []*input_basics.PmsNotifyListModel, totalCount int, err error) {
	mod := dao.PmsNotify.Ctx(ctx)

	mod = mod.Fields(input_basics.PmsNotifyListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsNotify.Columns().Id, in.Id)
	}

	if in.MemberId > 0 {
		mod = mod.Where(dao.PmsNotify.Columns().MemberId, in.MemberId)
	}

	if in.NotifyTitle != "" {
		mod = mod.WhereLike(dao.PmsNotify.Columns().NotifyTitle, in.NotifyTitle)
	}

	if in.NotifyType != "" {
		mod = mod.WhereLike(dao.PmsNotify.Columns().NotifyType, in.NotifyType)
	}

	if in.IsRead != "" {
		mod = mod.WhereLike(dao.PmsNotify.Columns().IsRead, in.IsRead)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsNotify.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取站内信列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsAppNotify) NotifyEdit(ctx context.Context, in *input_basics.PmsNotifyEditInp) (err error) {
	var (
		MemberInfo  *entity.PmsMember
		PmsLanguage *entity.PmsLanguage
	)
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		switch in.NotifyTitle {
		case "SYSTEM":
			if err = dao.PmsLanguage.Ctx(ctx).Where(g.Map{
				dao.PmsLanguage.Columns().Language: contexts.GetLanguage(ctx),
				dao.PmsLanguage.Columns().Key:      in.NotifyTitle,
			}).Scan(&PmsLanguage); err != nil {
				err = gerror.New("获取语言失败，请稍后重试！")
				return
			}
			in.NotifyTitle = PmsLanguage.Content
			break
		}
		if in.Id > 0 {
			if _, err = dao.PmsNotify.Ctx(ctx).TX(tx).
				Fields(input_basics.PmsNotifyUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改站内信失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsNotify.Ctx(ctx).TX(tx).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增站内信失败，请稍后重试！")
		}
		if !g.IsEmpty(in.MemberId) {
			_ = dao.PmsMember.Ctx(ctx).WherePri(in.MemberId).Scan(&MemberInfo)
			if !g.IsEmpty(MemberInfo.MemberNo) {
				// _ = MobilePush.SendPush(ctx, "welcome to live one member", "welcome to live one member", g.SliceStr{MemberInfo.MemberNo}, g.MapStrStr{
				// 	"path": "/account/message",
				// })
			}
		}

		return
	})
}

func (s *sBasicsAppNotify) NotifyDelete(ctx context.Context, in *input_basics.PmsNotifyDeleteInp) (err error) {

	if _, err = dao.PmsNotify.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除站内信失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsAppNotify) NotifyView(ctx context.Context, in *input_basics.PmsNotifyViewInp) (res *input_basics.PmsNotifyViewModel, err error) {
	if err = dao.PmsNotify.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取站内信信息，请稍后重试！")
		return
	}
	return
}
