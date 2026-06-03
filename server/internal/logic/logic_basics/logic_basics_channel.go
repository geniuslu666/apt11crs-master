package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsChannel struct{}

func NewBasicsChannel() *sBasicsChannel {
	return &sBasicsChannel{}
}

func init() {
	service.RegisterBasicsChannel(NewBasicsChannel())
}

func (s *sBasicsChannel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsChannel.Ctx(ctx), option...)
}
func (s *sBasicsChannel) List(ctx context.Context, in *input_basics.PmsChannelListInp) (list []*input_basics.PmsChannelListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsChannel.Table(), input_basics.PmsChannelListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsChannelListModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsChannel.Columns().Id, "=", dao.PmsMember.Columns().ChannelId)

	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.PmsChannel.Columns().Name, "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.Phone) {
		mod = mod.WhereLike(dao.PmsChannel.Columns().Phone, "%"+in.Phone+"%")
	}

	if !g.IsEmpty(in.Email) {
		mod = mod.WhereLike(dao.PmsChannel.Columns().Email, "%"+in.Email+"%")
	}

	if len(in.Rate) == 2 {
		mod = mod.WhereBetween(dao.PmsChannel.Columns().Rate, in.Rate[0], in.Rate[1])
	}

	if in.Status > 0 {
		mod = mod.Where(dao.PmsChannel.Columns().Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsChannel.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsChannel.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取渠道管理列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsChannel) Edit(ctx context.Context, in *input_basics.PmsChannelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsChannelUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改渠道管理失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsChannelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增渠道管理失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsChannel) Bind(ctx context.Context, in *input_basics.PmsChannelBindInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		memberCount, err := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, in.PmsMemberId).Where(dao.PmsMember.Columns().RebateMode, "MEMBER").Count()
		if err != nil {
			err = gerror.New("绑定会员失败，未找到用户")
			return
		}
		if memberCount <= 0 {
			err = gerror.New("该用户已被绑定，请重新选择用户")
			return
		}

		if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.PmsMemberId).Data(g.Map{
			dao.PmsMember.Columns().RebateMode: "CHANNEL",
			dao.PmsMember.Columns().ChannelId:  in.Id,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsChannel) Unbind(ctx context.Context, in *input_basics.PmsChannelUnbindInp) (err error) {

	if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().RebateMode, "CHANNEL").Where(dao.PmsMember.Columns().ChannelId, in.Id).Data(g.Map{
		dao.PmsMember.Columns().RebateMode: "MEMBER",
		dao.PmsMember.Columns().ChannelId:  nil,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}

func (s *sBasicsChannel) Delete(ctx context.Context, in *input_basics.PmsChannelDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除渠道管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsChannel) View(ctx context.Context, in *input_basics.PmsChannelViewInp) (res *input_basics.PmsChannelViewModel, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsChannel.Table(), input_basics.PmsChannelViewModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsChannelViewModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsChannel.Columns().Id, "=", dao.PmsMember.Columns().ChannelId)

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取渠道管理信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsChannel) Status(ctx context.Context, in *input_basics.PmsChannelStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsChannel.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新渠道管理状态失败，请稍后重试！")
		return
	}
	return
}
