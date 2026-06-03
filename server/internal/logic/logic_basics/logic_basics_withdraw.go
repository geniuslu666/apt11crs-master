package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sBasicsWithdraw struct{}

func NewBasicsWithdraw() *sBasicsWithdraw {
	return &sBasicsWithdraw{}
}

func init() {
	service.RegisterBasicsWithdraw(NewBasicsWithdraw())
}

func (s *sBasicsWithdraw) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsWithdraw.Ctx(ctx), option...)
}

func (s *sBasicsWithdraw) List(ctx context.Context, in *input_basics.PmsWithdrawListInp) (list []*input_basics.PmsWithdrawListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsWithdraw.Table(), input_basics.PmsWithdrawListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsWithdrawListModel{}, &dao.PmsStaff, "pmsStaff"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsWithdrawListModel{}, &dao.PmsChannel, "pmsChannel"))

	mod = mod.LeftJoinOnFields(dao.PmsStaff.Table(), dao.PmsWithdraw.Columns().StaffId, "=", dao.PmsStaff.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.PmsChannel.Table(), dao.PmsWithdraw.Columns().ChannelId, "=", dao.PmsChannel.Columns().Id)

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsWithdraw.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.Type != "" {
		mod = mod.Where(dao.PmsWithdraw.Columns().Type, in.Type)
	}

	if in.PmsStaffName != "" {
		mod = mod.WherePrefixLike(dao.PmsStaff.Table(), dao.PmsStaff.Columns().Name, "%"+in.PmsStaffName+"%")
	}

	if in.PmsChannelName != "" {
		mod = mod.WherePrefixLike(dao.PmsChannel.Table(), dao.PmsChannel.Columns().Name, "%"+in.PmsChannelName+"%")
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsWithdraw.Table() + "." + dao.PmsWithdraw.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取提现申请表列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsWithdraw) View(ctx context.Context, in *input_basics.PmsWithdrawViewInp) (res *input_basics.PmsWithdrawViewModel, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.PmsWithdraw.Table(), input_basics.PmsWithdrawListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsWithdrawListModel{}, &dao.PmsStaff, "pmsStaff"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_basics.PmsWithdrawListModel{}, &dao.PmsChannel, "pmsChannel"))

	mod = mod.LeftJoinOnFields(dao.PmsStaff.Table(), dao.PmsWithdraw.Columns().StaffId, "=", dao.PmsStaff.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.PmsChannel.Table(), dao.PmsWithdraw.Columns().ChannelId, "=", dao.PmsChannel.Columns().Id)

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取提现申请表信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsWithdraw) Agree(ctx context.Context, in *input_basics.PmsWithdrawAgreeInp) (err error) {

	var models *entity.PmsWithdraw
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("提现信息不存在或已被删除")
		return
	}

	if models.WithdrawStatus != "WAIT" {
		err = gerror.New("提现状态不正确")
		return
	}

	if _, err = s.Model(ctx).
		WherePri(in.Id).Data(input_basics.PmsWithdrawAgreeFields{
		WithdrawStatus: "SUCCESS",
		Transfer:       1,
		ApplyAt:        gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
	}
	return

}

func (s *sBasicsWithdraw) Disagree(ctx context.Context, in *input_basics.PmsWithdrawDisagreeInp) (err error) {

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var models *entity.PmsWithdraw
		if err = s.Model(ctx).TX(tx).Where("id", in.Id).Scan(&models); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if models == nil {
			err = gerror.New("提现信息不存在或已被删除")
			return
		}

		// 更新提现状态为失败，并记录失败原因
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(input_basics.PmsWithdrawDisagreeFields{
			WithdrawStatus: "FAIL",
			ApplyRemark:    in.ApplyRemark,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
		}

		// 渠道：减少提现中余额，增加可提现余额
		if models.Type == "CHANNEL" {
			if _, err = dao.PmsChannel.Ctx(ctx).TX(tx).WherePri(models.ChannelId).Update(g.MapStrAny{
				dao.PmsChannel.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
				dao.PmsChannel.Columns().Balance:              gdb.Raw(fmt.Sprintf("balance+%f", models.WithdrawAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
		}

		// 员工：减少提现中余额，增加可提现余额
		if models.Type == "STAFF" {
			if _, err = dao.PmsStaff.Ctx(ctx).TX(tx).WherePri(models.StaffId).Update(g.MapStrAny{
				dao.PmsStaff.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
				dao.PmsStaff.Columns().Balance:              gdb.Raw(fmt.Sprintf("balance+%f", models.WithdrawAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
		}

		return
	})

}

func (s *sBasicsWithdraw) Transfer(ctx context.Context, in *input_basics.PmsWithdrawTransferInp) (err error) {

	var (
		tx gdb.TX
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var models *entity.PmsWithdraw
	if err = s.Model(ctx).TX(tx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("提现信息不存在或已被删除")
		return
	}

	if models.WithdrawStatus != "SUCCESS" {
		err = gerror.New("提现状态不正确")
		return
	}

	if models.Transfer != 1 {
		err = gerror.New("转账状态不正确")
		return
	}

	if _, err = s.Model(ctx).TX(tx).
		WherePri(in.Id).Data(input_basics.PmsWithdrawTransferFields{
		Transfer: 2,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
	}

	// 减少提现中余额，增加已提现金额
	if models.Type == "CHANNEL" {
		if _, err = dao.PmsChannel.Ctx(ctx).TX(tx).WherePri(models.ChannelId).Update(g.MapStrAny{
			dao.PmsChannel.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
			dao.PmsChannel.Columns().WithdrawBalance:      gdb.Raw(fmt.Sprintf("withdraw_balance+%f", models.WithdrawAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "提现转账失败，请稍后重试！")
			return
		}
	}

	// 减少提现中余额，增加已提现金额
	if models.Type == "STAFF" {
		if _, err = dao.PmsStaff.Ctx(ctx).TX(tx).WherePri(models.StaffId).Update(g.MapStrAny{
			dao.PmsStaff.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
			dao.PmsStaff.Columns().WithdrawBalance:      gdb.Raw(fmt.Sprintf("withdraw_balance+%f", models.WithdrawAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "提现转账失败，请稍后重试！")
			return
		}
	}

	return

}
