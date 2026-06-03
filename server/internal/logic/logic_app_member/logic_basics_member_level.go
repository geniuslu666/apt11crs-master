package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_app_member"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAppMember) MemberLevelList(ctx context.Context, in *input_app_member.PmsMemberLevelListInp) (list []*input_app_member.PmsMemberLevelListModel, totalCount int, err error) {
	mod := dao.PmsMemberLevel.Ctx(ctx)

	mod = mod.Fields(input_app_member.PmsMemberLevelListModel{})

	if in.LevelName != "" {
		mod = mod.WhereLike(dao.PmsMemberLevel.Columns().LevelName, "%"+in.LevelName+"%")
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.PmsMemberLevel.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员等级列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		count, _ := dao.PmsMember.Ctx(ctx).Where("level", v.Id).Count()
		v.MemberCount = count
	}
	return
}

func (s *sAppMember) MemberLevelAllList(ctx context.Context, in *input_app_member.PmsMemberLevelAllInp) (list []*input_app_member.PmsMemberLevelAllModel, err error) {
	mod := dao.PmsMemberLevel.Ctx(ctx).WithAll()

	mod = mod.Fields(input_app_member.PmsMemberLevelAllModel{})

	if in.LevelName != "" {
		mod = mod.WhereLike(dao.PmsMemberLevel.Columns().LevelName, in.LevelName)
	}

	mod = mod.OrderDesc(dao.PmsMemberLevel.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取会员等级列表失败，请稍后重试！")
		return
	}

	//for _, v := range list {
	//	count, _ := dao.PmsMember.Ctx(ctx).Where("level", v.Id).Count()
	//	v.MemberCount = count
	//}
	return
}

func (s *sAppMember) MemberLevelEdit(ctx context.Context, in *input_app_member.PmsMemberLevelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsMemberLevel.Ctx(ctx).
				Fields(input_app_member.PmsMemberLevelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员等级失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsMemberLevel.Ctx(ctx).
			Fields(input_app_member.PmsMemberLevelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员等级失败，请稍后重试！")
		}
		return
	})
}

func (s *sAppMember) MemberLevelDelete(ctx context.Context, in *input_app_member.PmsMemberLevelDeleteInp) (err error) {

	if _, err = dao.PmsMemberLevel.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除会员等级失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberLevelView(ctx context.Context, in *input_app_member.PmsMemberLevelViewInp) (res *input_app_member.PmsMemberLevelViewModel, err error) {
	if err = dao.PmsMemberLevel.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员等级信息，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberLevelComputeLevel(ctx context.Context, MemberId int, tx gdb.TX) (err error) {
	var (
		exp      *gvar.Var
		levelId  *gvar.Var
		updateId int64
	)
	if exp, err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberId).TX(tx).Value(dao.PmsMember.Columns().Exp); err != nil {
		return
	}

	if levelId, err = dao.PmsMemberLevel.Ctx(ctx).TX(tx).
		WhereLTE(dao.PmsMemberLevel.Columns().Exp, exp.Int()).
		OrderDesc(dao.PmsMemberLevel.Columns().Exp).
		Value(dao.PmsMemberLevel.Columns().Id); err != nil {
		return
	}

	if updateId, err = dao.PmsMemberLevel.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsMember.Columns().Id: MemberId,
	}).UpdateAndGetAffected(g.MapStrAny{
		dao.PmsMember.Columns().Level: levelId.Int(),
	}); err != nil {
		return
	}

	if g.IsEmpty(updateId) {
		err = gerror.New("用户等级处理失败")
		return
	}

	return
}
