package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/model/input/input_app_member"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sAppMember) BalanceChangeStat(ctx context.Context, in *input_app_member.PmsBalanceChangeStatInp) (res *input_app_member.PmsBalanceChangeStatModel, err error) {
	res = &input_app_member.PmsBalanceChangeStatModel{}
	mod := dao.PmsBalanceChange.Ctx(ctx)

	res.TotalSendPoints, err = mod.WhereGT(dao.PmsBalanceChange.Columns().ChangePrice, 0).Sum(dao.PmsBalanceChange.Columns().ChangePrice)

	consumePoints, err := mod.WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0).Sum(dao.PmsBalanceChange.Columns().ChangePrice)
	res.TotalConsumePoints = math.Abs(consumePoints)

	res.Points, err = dao.PmsMember.Ctx(ctx).Sum(dao.PmsMember.Columns().Balance)

	fmt.Println("-------------********************--------------------")

	return
}

func (s *sAppMember) BalanceChangeList(ctx context.Context, in *input_app_member.PmsBalanceChangeListInp) (list []*input_app_member.PmsBalanceChangeListModel, totalCount int, err error) {
	mod := dao.PmsBalanceChange.Ctx(ctx)

	mod = mod.FieldsPrefix(dao.PmsBalanceChange.Table(), input_app_member.PmsBalanceChangeListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsBalanceChangeListModel{}, &dao.PmsMember, "pmsMember"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsBalanceChangeListModel{}, &dao.AdminMember, "adminMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsBalanceChange.Columns().MemberId, "=", dao.PmsMember.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.PmsBalanceChange.Columns().OperatorId, "=", dao.AdminMember.Columns().Id)

	if in.MemberId > 0 {
		mod = mod.Where(dao.PmsBalanceChange.Columns().MemberId, in.MemberId)
	}

	if in.OperatorId > 0 {
		mod = mod.Where(dao.PmsBalanceChange.Columns().OperatorId, in.OperatorId)
	}

	if in.Reason != "" {
		mod = mod.WhereLike(dao.PmsBalanceChange.Columns().Reason, "%"+in.Reason+"%")
	}

	if in.Des != "" {
		mod = mod.WhereLike(dao.PmsBalanceChange.Columns().Des, "%"+in.Des+"%")
	}

	if in.OrderNo != "" {
		mod = mod.WhereLike(dao.PmsBalanceChange.Columns().OrderSn, "%"+in.OrderNo+"%")
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.MemberKey != "" {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().FullName, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Phone, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Mail, "%"+in.MemberKey+"%"))
	}

	if in.Direction == 1 {
		mod = mod.WhereGT(dao.PmsBalanceChange.Columns().ChangePrice, 0)
	}

	if in.Direction == 2 {
		mod = mod.WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsBalanceChange.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取积分明细列表失败，请稍后重试！")
		return
	}
	return
}
