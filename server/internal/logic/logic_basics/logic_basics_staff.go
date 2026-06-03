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

type sBasicsStaff struct{}

func NewBasicsStaff() *sBasicsStaff {
	return &sBasicsStaff{}
}

func init() {
	service.RegisterBasicsStaff(NewBasicsStaff())
}

func (s *sBasicsStaff) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsStaff.Ctx(ctx), option...)
}

func (s *sBasicsStaff) List(ctx context.Context, in *input_basics.PmsStaffListInp) (list []*input_basics.PmsStaffListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsStaff.Table(), input_basics.PmsStaffListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsStaffListModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsStaff.Columns().Id, "=", dao.PmsMember.Columns().StaffId)

	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.PmsStaff.Columns().Name, "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.Department) {
		mod = mod.WhereLike(dao.PmsStaff.Columns().Department, "%"+in.Department+"%")
	}

	if !g.IsEmpty(in.Phone) {
		mod = mod.WhereLike(dao.PmsStaff.Columns().Phone, "%"+in.Phone+"%")
	}

	if !g.IsEmpty(in.Email) {
		mod = mod.WhereLike(dao.PmsStaff.Columns().Email, "%"+in.Email+"%")
	}

	if len(in.Rate) == 2 {
		mod = mod.WhereBetween(dao.PmsStaff.Columns().Rate, in.Rate[0], in.Rate[1])
	}

	if in.Status > 0 {
		mod = mod.Where(dao.PmsStaff.Columns().Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsStaff.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsStaff.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取员工管理列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsStaff) Edit(ctx context.Context, in *input_basics.PmsStaffEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsStaffUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改员工管理失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsStaffInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增员工管理失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsStaff) Bind(ctx context.Context, in *input_basics.PmsStaffBindInp) (err error) {
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
			dao.PmsMember.Columns().RebateMode: "STAFF",
			dao.PmsMember.Columns().StaffId:    in.Id,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsStaff) Unbind(ctx context.Context, in *input_basics.PmsStaffUnbindInp) (err error) {

	if _, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().RebateMode, "STAFF").Where(dao.PmsMember.Columns().StaffId, in.Id).Data(g.Map{
		dao.PmsMember.Columns().RebateMode: "MEMBER",
		dao.PmsMember.Columns().StaffId:    nil,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}

func (s *sBasicsStaff) Delete(ctx context.Context, in *input_basics.PmsStaffDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除员工管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsStaff) View(ctx context.Context, in *input_basics.PmsStaffViewInp) (res *input_basics.PmsStaffViewModel, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsStaff.Table(), input_basics.PmsStaffViewModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsStaffViewModel{}, &dao.PmsMember, "pmsMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsStaff.Columns().Id, "=", dao.PmsMember.Columns().StaffId)

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取员工管理信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsStaff) Status(ctx context.Context, in *input_basics.PmsStaffStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsStaff.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新员工管理状态失败，请稍后重试！")
		return
	}
	return
}
