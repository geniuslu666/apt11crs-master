package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/model/input/input_app_member"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sAppMember) ExpChangeList(ctx context.Context, in *input_app_member.PmsExpChangeListInp) (list []*input_app_member.PmsExpChangeListModel, totalCount int, err error) {
	mod := dao.PmsExpChange.Ctx(ctx)

	mod = mod.FieldsPrefix(dao.PmsExpChange.Table(), input_app_member.PmsExpChangeListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsExpChangeListModel{}, &dao.PmsMember, "pmsMember"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsExpChangeListModel{}, &dao.AdminMember, "adminMember"))

	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.PmsExpChange.Columns().MemberId, "=", dao.PmsMember.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.PmsExpChange.Columns().OperatorId, "=", dao.AdminMember.Columns().Id)

	if in.MemberId > 0 {
		mod = mod.Where(dao.PmsExpChange.Columns().MemberId, in.MemberId)
	}

	if in.OperatorId > 0 {
		mod = mod.Where(dao.PmsExpChange.Columns().OperatorId, in.OperatorId)
	}

	if in.Des != "" {
		mod = mod.WhereLike(dao.PmsExpChange.Columns().Des, "%"+in.Des+"%")
	}

	if in.OrderNo != "" {
		mod = mod.WhereLike(dao.PmsExpChange.Columns().OrderSn, "%"+in.OrderNo+"%")
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsExpChange.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.MemberKey != "" {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().FullName, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Phone, "%"+in.MemberKey+"%").
			WhereOrLike(dao.PmsMember.Columns().Mail, "%"+in.MemberKey+"%"))
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsExpChange.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取成长值明细列表失败，请稍后重试！")
		return
	}
	return
}
