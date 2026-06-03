package logic_car

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sCarDriverWithdraw struct{}

func NewCarDriverWithdraw() *sCarDriverWithdraw {
	return &sCarDriverWithdraw{}
}

func init() {
	service.RegisterCarDriverWithdraw(NewCarDriverWithdraw())
}

// Model 司机提现管理ORM模型
func (s *sCarDriverWithdraw) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarDriverWithdraw.Ctx(ctx), option...)
}

// List 获取司机提现管理列表
func (s *sCarDriverWithdraw) List(ctx context.Context, in *input_car.DriverWithdrawListInp) (list []*input_car.DriverWithdrawListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.FieldsPrefix(dao.CarDriverWithdraw.Table(), input_car.DriverWithdrawListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_car.DriverWithdrawListModel{}, &dao.CarDriver, "carDriver"))

	mod = mod.LeftJoinOnFields(dao.CarDriver.Table(), dao.CarDriverWithdraw.Columns().DriverId, "=", dao.CarDriver.Columns().Id)

	// 查询名称
	if !g.IsEmpty(in.Type) {
		mod = mod.Where(dao.CarDriverWithdraw.Columns().Type, in.Type)
	}

	if !g.IsEmpty(in.WithdrawSn) {
		mod = mod.WhereLike(dao.CarDriverWithdraw.Columns().WithdrawSn, "%"+in.WithdrawSn+"%")
	}

	if !g.IsEmpty(in.DriverId) {
		mod = mod.Where(dao.CarDriverWithdraw.Columns().DriverId, in.DriverId)
	}

	if !g.IsEmpty(in.WithdrawStatus) {
		mod = mod.Where(dao.CarDriverWithdraw.Columns().WithdrawStatus, in.WithdrawStatus)
	}

	if !g.IsEmpty(in.DriverName) {
		mod = mod.WherePrefixLike(dao.CarDriver.Table(), dao.CarDriver.Columns().Name, "%"+in.DriverName+"%")
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CarDriverWithdraw.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.CarDriverWithdraw.Table() + "." + dao.CarDriverWithdraw.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取提现申请表列表失败，请稍后重试！")
		return
	}

	return
}

func (s *sCarDriverWithdraw) View(ctx context.Context, in *input_car.DriverWithdrawViewInp) (res *input_car.DriverWithdrawViewModel, err error) {
	mod := s.Model(ctx).WithAll()

	if err = mod.WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取提现申请表信息，请稍后重试！")
		return
	}
	return
}

func (s *sCarDriverWithdraw) Agree(ctx context.Context, in *input_car.DriverWithdrawAgreeInp) (err error) {

	var models *entity.CarDriverWithdraw
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
		WherePri(in.Id).Data(input_car.DriverWithdrawAgreeFields{
		WithdrawStatus: "SUCCESS",
		Transfer:       1,
		ApplyAt:        gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
	}
	return

}

func (s *sCarDriverWithdraw) Disagree(ctx context.Context, in *input_car.DriverWithdrawDisagreeInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 获取提现申请表信息
		var models *entity.CarDriverWithdraw
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
			WherePri(in.Id).Data(input_car.DriverWithdrawDisagreeFields{
			WithdrawStatus: "FAIL",
			ApplyRemark:    in.ApplyRemark,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
		}

		// 减少司机提现中金额，增加司机可提现金额
		if _, err = dao.CarDriver.Ctx(ctx).TX(tx).WherePri(models.DriverId).Update(g.MapStrAny{
			dao.CarDriver.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
			dao.CarDriver.Columns().Balance:              gdb.Raw(fmt.Sprintf("balance+%f", models.WithdrawAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}
		return
	})

}

func (s *sCarDriverWithdraw) Transfer(ctx context.Context, in *input_car.DriverWithdrawTransferInp) (err error) {

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

	var models *entity.CarDriverWithdraw
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

	if _, err = dao.CarDriver.Ctx(ctx).TX(tx).WherePri(models.DriverId).Update(g.MapStrAny{
		dao.CarDriver.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance-%f", models.WithdrawAmount)),
		dao.CarDriver.Columns().WithdrawBalance:      gdb.Raw(fmt.Sprintf("withdraw_balance+%f", models.WithdrawAmount)),
	}); err != nil {
		err = gerror.Wrap(err, "提现转账失败，请稍后重试！")
		return
	}

	// 写入司机积分变更日志
	if _, err = dao.CarDriverBalanceChange.Ctx(ctx).TX(tx).OmitEmptyData().Insert(entity.CarDriverBalanceChange{
		DriverId:        gvar.New(models.DriverId).Int(),
		Type:            "WITHDRAW",
		ChangePrice:     models.WithdrawAmount,
		WithdrawOrderId: models.Id,
		Des:             "提现",
		CreatedAt:       gtime.Now(),
		UpdatedAt:       gtime.Now(),
	}); err != nil {
		err = gerror.Wrap(err, "写入司机积分变更日志失败，请稍后重试！")
		return
	}

	return

}
