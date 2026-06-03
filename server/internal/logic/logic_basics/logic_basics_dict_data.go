package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/dict"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsDictData struct{}

func NewBasicsDictData() *sBasicsDictData {
	return &sBasicsDictData{}
}

func init() {
	service.RegisterBasicsDictData(NewBasicsDictData())
}

func (s *sBasicsDictData) Delete(ctx context.Context, in *input_basics.DictDataDeleteInp) error {
	_, err := dao.SysDictData.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

func (s *sBasicsDictData) Edit(ctx context.Context, in *input_basics.DictDataEditInp) (err error) {
	if in.Id > 0 {
		_, err = dao.SysDictData.Ctx(ctx).Fields(input_basics.DictDataUpdateFields{}).WherePri(in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		return nil
	}

	in.Type, err = s.GetType(ctx, in.TypeID)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if in.Type == "" {
		return gerror.Wrap(err, "类型选择无效，请检查")
	}

	_, err = dao.SysDictData.Ctx(ctx).Fields(input_basics.DictDataInsertFields{}).Data(in).OmitEmptyData().Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

func (s *sBasicsDictData) List(ctx context.Context, in *input_basics.DictDataListInp) (list []*input_basics.DictDataListModel, totalCount int, err error) {
	mod := dao.SysDictData.Ctx(ctx)
	if in.TypeID > 0 {
		types, err := s.GetTypes(ctx, in.TypeID)
		if err != nil {
			return list, totalCount, err
		}
		mod = mod.WhereIn("type", types)
	}

	if in.Type != "" {
		mod = mod.Where("type", in.Type)
	}

	if in.Value != "" {
		mod = mod.Where("value", in.Value)
	}

	if in.Label != "" {
		mod = mod.WhereLike("label", "%"+in.Label+"%")
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		v.TypeID, _ = s.GetId(ctx, v.Type)
	}
	return list, totalCount, err
}

func (s *sBasicsDictData) GetId(ctx context.Context, t string) (id int64, err error) {
	m := dao.SysDictType.Ctx(ctx).Fields("id").Where("type", t).Where("status", consts.StatusEnabled)
	val, err := m.Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return 0, err
	}
	return val.Int64(), nil
}

func (s *sBasicsDictData) GetType(ctx context.Context, id int64) (types string, err error) {
	m := dao.SysDictType.Ctx(ctx).Fields("type").Where("id", id).Where("status", consts.StatusEnabled)
	val, err := m.Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return types, err
	}
	return val.String(), nil
}

func (s *sBasicsDictData) GetTypes(ctx context.Context, id int64) (types []string, err error) {
	columns, err := dao.SysDictType.Ctx(ctx).Fields("type").
		Where("id", id).WhereOr("pid", id).Where("status", consts.StatusEnabled).
		Array()
	types = g.NewVar(columns).Strings()
	return
}

func (s *sBasicsDictData) Select(ctx context.Context, in *input_basics.DataSelectInp) (list input_basics.DataSelectModel, err error) {
	options, err := dict.GetOptions(ctx, in.Type)
	if err == nil {
		return options, nil
	}
	if !errors.Is(err, dict.NotExistKeyError) {
		return nil, err
	}

	mod := dao.SysDictData.Ctx(ctx).Where("type", in.Type)
	if err = mod.Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	for k, v := range list {
		list[k].Value = consts.ConvType(v.Value, v.ValueType)
		list[k].Key = list[k].Value
	}

	if len(list) == 0 {
		list = make(input_basics.DataSelectModel, 0)
	}
	return
}
