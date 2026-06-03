package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_app_member"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAppMember) MemberSceneList(ctx context.Context, in *input_app_member.PmsMemberSceneListInp) (list []*input_app_member.PmsMemberSceneListModel, totalCount int, err error) {
	mod := dao.PmsMemberScene.Ctx(ctx)

	mod = mod.Fields(input_app_member.PmsMemberSceneListModel{})

	if in.SceneName != "" {
		mod = mod.WhereLike(dao.PmsMemberScene.Columns().SceneName, in.SceneName)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMemberScene.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.PmsMemberScene.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员等级场景列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberSceneEdit(ctx context.Context, in *input_app_member.PmsMemberSceneEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsMemberScene.Ctx(ctx).
				Fields(input_app_member.PmsMemberSceneUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员等级场景失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsMemberScene.Ctx(ctx).
			Fields(input_app_member.PmsMemberSceneInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员等级场景失败，请稍后重试！")
		}
		return
	})
}

func (s *sAppMember) MemberSceneDelete(ctx context.Context, in *input_app_member.PmsMemberSceneDeleteInp) (err error) {

	if _, err = dao.PmsMemberScene.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除会员等级场景失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberSceneView(ctx context.Context, in *input_app_member.PmsMemberSceneViewInp) (res *input_app_member.PmsMemberSceneViewModel, err error) {
	if err = dao.PmsMemberScene.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员等级场景失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberSceneSwitch(ctx context.Context, in *input_app_member.PmsMemberSceneSwitchInp) (err error) {
	var fields = []string{
		dao.PmsMemberScene.Columns().IsGetOpen,
		dao.PmsMemberScene.Columns().IsPayOpen,
	}

	if !validate.InSlice(fields, in.Key) {
		err = gerror.New("开关键名不在白名单")
		return
	}

	if _, err = dao.PmsMemberScene.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		in.Key: in.Value,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新会员等级场景开关失败，请稍后重试！")
		return
	}
	return
}
