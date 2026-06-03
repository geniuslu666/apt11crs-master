package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsProvinces struct{}

func NewBasicsProvinces() *sBasicsProvinces {
	return &sBasicsProvinces{}
}

func init() {
	service.RegisterBasicsProvinces(NewBasicsProvinces())
}

func (s *sBasicsProvinces) Tree(ctx context.Context) (list []*input_basics.ProvincesTree, err error) {
	var models []*entity.SysProvinces
	if err = dao.SysProvinces.Ctx(ctx).Order("pid asc,id asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取省市区关系树选项列表失败！")
		return
	}

	list = s.treeList(0, models)
	return
}

func (s *sBasicsProvinces) Delete(ctx context.Context, in *input_basics.ProvincesDeleteInp) (err error) {
	var models *entity.SysProvinces
	if err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取省市区数据失败！")
		return
	}

	if models == nil {
		err = gerror.New("数据不存在或已删除！")
		return
	}

	has, err := dao.SysProvinces.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, "删除省市区数据时获取上级数据失败！")
		return
	}

	if !has.IsEmpty() {
		err = gerror.New("请先删除该地区下得所有子级！")
		return
	}

	if _, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除省市区数据失败！")
		return
	}
	return
}

func (s *sBasicsProvinces) Edit(ctx context.Context, in *input_basics.ProvincesEditInp) (err error) {

	in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.SysProvinces, in.Pid)
	if err != nil {
		return
	}

	models, err := s.View(ctx, &input_basics.ProvincesViewInp{Id: in.Id})
	if err != nil {
		return
	}

	if models != nil {
		if _, err = dao.SysProvinces.Ctx(ctx).Fields(input_basics.ProvincesUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改省市区数据失败！")
		}
		return
	}

	if _, err = dao.SysProvinces.Ctx(ctx).Fields(input_basics.ProvincesInsertFields{}).Data(in).OmitEmptyData().Insert(); err != nil {
		err = gerror.Wrap(err, "新增省市区数据失败！")
	}
	return
}

func (s *sBasicsProvinces) Status(ctx context.Context, in *input_basics.ProvincesStatusInp) (err error) {
	if _, err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update(); err != nil {
		err = gerror.Wrap(err, "更新省市区状态失败！")
	}
	return
}

func (s *sBasicsProvinces) MaxSort(ctx context.Context, in *input_basics.ProvincesMaxSortInp) (res *input_basics.ProvincesMaxSortModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Fields(dao.SysProvinces.Columns().Sort).OrderDesc(dao.SysProvinces.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取省市区最大排序失败！")
		return
	}

	if res == nil {
		res = new(input_basics.ProvincesMaxSortModel)
	}
	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sBasicsProvinces) View(ctx context.Context, in *input_basics.ProvincesViewInp) (res *input_basics.ProvincesViewModel, err error) {
	if err = dao.SysProvinces.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取省市区信息失败！")
	}
	return
}

func (s *sBasicsProvinces) List(ctx context.Context, in *input_basics.ProvincesListInp) (list []*input_basics.ProvincesListModel, totalCount int, err error) {
	mod := dao.SysProvinces.Ctx(ctx)

	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取省市区数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取省市区列表失败！")
	}
	return
}

func (s *sBasicsProvinces) ChildrenList(ctx context.Context, in *input_basics.ProvincesChildrenListInp) (list []*input_basics.ProvincesChildrenListModel, totalCount int, err error) {
	mod := dao.SysProvinces.Ctx(ctx)

	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	if in.Pid > 0 {
		mod = mod.Where("pid", in.Pid)
	}

	if in.Id > 0 {
		mod = mod.Where("id", in.Id)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取省市区下级数据行失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取省市区下级列表失败！")
	}
	return
}

func (s *sBasicsProvinces) UniqueId(ctx context.Context, in *input_basics.ProvincesUniqueIdInp) (res *input_basics.ProvincesUniqueIdModel, err error) {
	res = new(input_basics.ProvincesUniqueIdModel)
	res.IsUnique = true
	if in.NewId == 0 {
		return
	}

	if err = hgorm.IsUnique(ctx, &dao.SysProvinces, g.Map{dao.SysProvinces.Columns().Id: in.NewId}, "", in.OldId); err != nil {
		res.IsUnique = false
		return
	}
	return
}

func (s *sBasicsProvinces) Select(ctx context.Context, in *input_basics.ProvincesSelectInp) (res *input_basics.ProvincesSelectModel, err error) {
	res = new(input_basics.ProvincesSelectModel)
	mod := dao.SysProvinces.Ctx(ctx).Fields("id as value, title as label, level").Where("pid", in.Value)

	if err = mod.Order("sort asc,id asc").Scan(&res.List); err != nil {
		err = gerror.Wrap(err, "获取省市区选项失败！")
		return
	}

	for _, v := range res.List {
		if in.DataType == "p" {
			v.IsLeaf = true
			continue
		}

		if in.DataType == "pc" && v.Level >= 2 {
			v.IsLeaf = true
			continue
		}

		if in.DataType == "pca" && v.Level >= 3 {
			v.IsLeaf = true
			continue
		}
	}
	return
}

func (s *sBasicsProvinces) treeList(pid int64, nodes []*entity.SysProvinces) (list []*input_basics.ProvincesTree) {
	list = make([]*input_basics.ProvincesTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(input_basics.ProvincesTree)
			item.SysProvinces = *v
			item.Label = v.Title
			item.Value = v.Id
			item.Key = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
