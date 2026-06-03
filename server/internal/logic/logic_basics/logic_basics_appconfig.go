package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsPmsAppconfig struct{}

func NewBasicsPmsAppconfig() *sBasicsPmsAppconfig {
	return &sBasicsPmsAppconfig{}
}

func init() {
	service.RegisterBasicsPmsAppconfig(NewBasicsPmsAppconfig())
}

func (s *sBasicsPmsAppconfig) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsAppconfig.Ctx(ctx), option...)
}

func (s *sBasicsPmsAppconfig) List(ctx context.Context, in *input_basics.PmsAppconfigListInp) (list []*input_basics.PmsAppconfigListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsAppconfigListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsAppconfig.Columns().Id, in.Id)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsAppconfig.Columns().Name, in.Name)
	}

	if in.Key != "" {
		mod = mod.WhereLike(dao.PmsAppconfig.Columns().Key, in.Key)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsAppconfig.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsAppconfig.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取APP配置列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsPmsAppconfig) Edit(ctx context.Context, in *input_basics.PmsAppconfigEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsAppconfigUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改APP配置失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsAppconfigInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增APP配置失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsPmsAppconfig) Delete(ctx context.Context, in *input_basics.PmsAppconfigDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除APP配置失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsPmsAppconfig) View(ctx context.Context, in *input_basics.PmsAppconfigViewInp) (res *input_basics.PmsAppconfigViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取APP配置信息，请稍后重试！")
		return
	}
	return
}
