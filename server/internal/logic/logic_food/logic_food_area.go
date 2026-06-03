package logic_food

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"APT/utility/tree"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodArea struct{}

func NewFoodArea() *sFoodArea {
	return &sFoodArea{}
}

func init() {
	service.RegisterFoodArea(NewFoodArea())
}

func (s *sFoodArea) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodArea.Ctx(ctx), option...)
}

func (s *sFoodArea) List(ctx context.Context, in *input_food.FoodAreaListInp) (res *input_food.FoodAreaListModel, totalCount int, err error) {
	var (
		models []*entity.FoodArea
		pid    int64 = 0
	)
	mod := s.Model(ctx)

	if !g.IsEmpty(in.Level) {
		mod = mod.Where(dao.FoodArea.Columns().Level, in.Level)
	}

	if in.Status > 0 {
		mod = mod.Where(dao.FoodArea.Columns().AreaStatus, in.Status)
	}

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	if err = mod.ScanAndCount(&models, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取区域管理列表失败，请稍后重试！")
		return
	}
	res = new(input_food.FoodAreaListModel)
	res.List = s.treeList(pid, models)

	return
}

func (s *sFoodArea) treeList(pid int64, nodes []*entity.FoodArea) (list []*input_food.AreaTree) {
	list = make([]*input_food.AreaTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(input_food.AreaTree)
			item.FoodArea = *v

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}

func (s *sFoodArea) Edit(ctx context.Context, in *input_food.FoodAreaEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if err = hgorm.IsUnique(ctx, &dao.FoodArea, g.Map{dao.FoodArea.Columns().AreaName: in.AreaName}, "名称已存在", in.Id); err != nil {
			return
		}

		if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.FoodArea, in.Pid); err != nil {
			return
		}

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_food.FoodAreaUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改区域管理失败，请稍后重试！")
			}

			return updateRoleChildrenTree(ctx, in.Id, in.Level, in.Tree)
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_food.FoodAreaInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增区域管理失败，请稍后重试！")
		}
		return
	})
}

func updateRoleChildrenTree(ctx context.Context, _id int64, _level int, _tree string) (err error) {
	var list []*entity.FoodArea
	if err = dao.FoodArea.Ctx(ctx).Where("pid", _id).Scan(&list); err != nil {
		return
	}
	for _, child := range list {
		child.Level = _level + 1
		child.Tree = tree.GenLabel(_tree, child.Pid)

		if _, err = dao.FoodArea.Ctx(ctx).Where("id", child.Id).Data("level", child.Level, "tree", child.Tree).Update(); err != nil {
			return
		}

		if err = updateRoleChildrenTree(ctx, child.Id, child.Level, child.Tree); err != nil {
			return
		}
	}
	return
}

func (s *sFoodArea) Delete(ctx context.Context, in *input_food.FoodAreaDeleteInp) (err error) {
	has, err := s.Model(ctx).Where("pid", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if !has.IsEmpty() {
		return gerror.New("请先删除该区域下得所有子级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除区域管理失败，请稍后重试！")
		return
	}
	return
}

func (s *sFoodArea) View(ctx context.Context, in *input_food.FoodAreaViewInp) (res *input_food.FoodAreaViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取区域管理信息，请稍后重试！")
		return
	}
	return
}

func (s *sFoodArea) Status(ctx context.Context, in *input_food.FoodAreaStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var list []*entity.FoodArea
		if err = dao.FoodArea.Ctx(ctx).Where(dao.FoodArea.Columns().Pid, in.Id).Scan(&list); err != nil {
			return
		}
		for _, child := range list {
			if _, err = dao.FoodArea.Ctx(ctx).WherePri(child.Id).Data(g.Map{
				dao.FoodArea.Columns().AreaStatus: in.Status,
			}).Update(); err != nil {
				return
			}
		}

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.FoodArea.Columns().AreaStatus: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新区域状态失败，请稍后重试！")
			return
		}

		return
	})
}
