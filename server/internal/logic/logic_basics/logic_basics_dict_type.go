package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/dict"
	"APT/internal/library/hgorm"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsDictType struct{}

func NewBasicsDictType() *sBasicsDictType {
	return &sBasicsDictType{}
}

func init() {
	service.RegisterBasicsDictType(NewBasicsDictType())
}

func (s *sBasicsDictType) Tree(ctx context.Context) (list []*input_basics.DictTypeTree, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	list = s.treeList(0, models)
	return
}

func (s *sBasicsDictType) Delete(ctx context.Context, in *input_basics.DictTypeDeleteInp) (err error) {
	var models *entity.SysDictType
	if err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New("数据不存在或已删除！")
		return
	}

	exist, err := dao.SysDictData.Ctx(ctx).Where("type", models.Type).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		err = gerror.New("请先删除该字典类型下得所有字典数据！")
		return
	}

	pidExist, err := dao.SysDictType.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !pidExist.IsEmpty() {
		err = gerror.New("请先删除该字典类型下得所有子级类型！")
		return
	}

	_, err = dao.SysDictType.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

func (s *sBasicsDictType) Edit(ctx context.Context, in *input_basics.DictTypeEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, &dao.SysDictType, g.Map{dao.SysDictType.Columns().Name: in.Name}, "名称已存在", in.Id); err != nil {
		return
	}

	if in.Id > 0 {
		if _, err = dao.SysDictType.Ctx(ctx).Fields(input_basics.DictTypeUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
		}
		return
	}

	if _, err = dao.SysDictType.Ctx(ctx).Fields(input_basics.DictTypeInsertFields{}).Data(in).OmitEmptyData().Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

func (s *sBasicsDictType) TreeSelect(ctx context.Context, in *input_basics.DictTreeSelectInp) (list []*input_basics.DictTypeTree, err error) {
	var (
		mod    = dao.SysDictType.Ctx(ctx)
		models []*entity.SysDictType
	)

	if err = mod.Order("pid asc,sort asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	list = s.treeList(0, models)
	list = append(list, s.BuiltinSelect()...)
	return
}

func (s *sBasicsDictType) BuiltinSelect() (list []*input_basics.DictTypeTree) {
	top := &input_basics.DictTypeTree{
		SysDictType: entity.SysDictType{
			Id:  dict.BuiltinId,
			Pid: 0,
		},
		Label: "内置字典",
		Value: dict.BuiltinId,
		Key:   dict.BuiltinId,
	}

	enums := &input_basics.DictTypeTree{
		SysDictType: entity.SysDictType{
			Id:  dict.EnumsId,
			Pid: dict.BuiltinId,
		},
		Label: "枚举字典",
		Value: dict.EnumsId,
		Key:   dict.EnumsId,
	}

	for _, v := range dict.GetAllEnums() {
		children := &input_basics.DictTypeTree{
			SysDictType: entity.SysDictType{
				Id:  v.Id,
				Pid: dict.EnumsId,
			},
			Label: v.Label,
			Value: v.Id,
			Key:   v.Id,
		}
		enums.Children = append(enums.Children, children)
	}

	fun := &input_basics.DictTypeTree{
		SysDictType: entity.SysDictType{
			Id:  dict.FuncId,
			Pid: dict.BuiltinId,
		},
		Label: "方法字典",
		Value: dict.FuncId,
		Key:   dict.FuncId,
	}

	for _, v := range dict.GetAllFunc() {
		children := &input_basics.DictTypeTree{
			SysDictType: entity.SysDictType{
				Id:  v.Id,
				Pid: dict.FuncId,
			},
			Label: v.Label,
			Value: v.Id,
			Key:   v.Id,
		}
		fun.Children = append(fun.Children, children)
	}

	top.Children = append(top.Children, enums, fun)
	list = append(list, top)
	return
}

func (s *sBasicsDictType) treeList(pid int64, nodes []*entity.SysDictType) (list []*input_basics.DictTypeTree) {
	list = make([]*input_basics.DictTypeTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(input_basics.DictTypeTree)
			item.SysDictType = *v
			item.Label = v.Name
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
