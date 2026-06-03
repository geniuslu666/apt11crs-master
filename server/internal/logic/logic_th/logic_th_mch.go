package logic_th

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sThMch struct{}

func NewThMch() *sThMch {
	return &sThMch{}
}

func init() {
	service.RegisterThMch(NewThMch())
}

func (s *sThMch) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThMch.Ctx(ctx), option...)
}

func (s *sThMch) List(ctx context.Context, in *input_th.ThMchListInp) (list []*input_th.ThMchListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.ThMch.Table(), input_th.ThMchListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_th.ThMchListModel{}, &dao.ThMchCategory, "thMchCategory"))

	mod = mod.LeftJoinOnFields(dao.ThMchCategory.Table(), dao.ThMch.Columns().CategoryId, "=", dao.ThMchCategory.Columns().Id)

	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.ThMch.Columns().Name, "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.CategoryId) {
		mod = mod.Where(dao.ThMch.Columns().CategoryId, in.CategoryId)
	}

	if !g.IsEmpty(in.ContactInfo) {
		mod = mod.WhereLike(dao.ThMch.Columns().ContactInfo, "%"+in.ContactInfo+"%")
	}

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.ThMch.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.MchIds) {
		mod = mod.WhereIn(dao.ThMch.Columns().Id, in.MchIds)
	}

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.ThMch.Table() + "." + dao.ThMch.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThMch) All(ctx context.Context, in *input_th.ThMchAllInp) (list []*input_th.ThMchAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_th.ThMchAllModel{})

	mod = mod.OrderDesc(dao.ThMch.Columns().Sort)

	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取商户全部列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sThMch) Edit(ctx context.Context, in *input_th.ThMchEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_th.ThMchUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改商户失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_th.ThMchInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增商户失败，请稍后重试！")
		}
		return
	})
}

func (s *sThMch) Delete(ctx context.Context, in *input_th.ThMchDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除商户失败，请稍后重试！")
		return
	}
	return
}

func (s *sThMch) View(ctx context.Context, in *input_th.ThMchViewInp) (res *input_th.ThMchViewModel, err error) {

	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.ThMch.Table(), input_th.ThMchViewModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_th.ThMchViewModel{}, &dao.ThMchCategory, "thMchCategory"))

	mod = mod.LeftJoinOnFields(dao.ThMchCategory.Table(), dao.ThMch.Columns().CategoryId, "=", dao.ThMchCategory.Columns().Id)

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取商户信息，请稍后重试！")
		return
	}
	return
}

func (s *sThMch) Switch(ctx context.Context, in *input_th.ThMchSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThMch.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}

func (s *sThMch) AppView(ctx context.Context, in *input_th.ThMchAppViewInp) (res *input_th.ThMchAppViewModel, err error) {

	var couponMch *entity.ThCouponMch
	if err = s.Model(ctx).Hook(hook.PmsFindLanguageValueHook).WithAll().WherePri(in.MchId).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取商户信息失败，请稍后重试！")
		return
	}

	_ = dao.ThCouponMch.Ctx(ctx).Where(dao.ThCouponMch.Columns().CouponId, in.CouponId).Where(dao.ThCouponMch.Columns().MchId, in.MchId).Scan(&couponMch)

	if !g.IsEmpty(couponMch) {
		res.VerifyGoodsName = couponMch.Name
	}

	return
}
