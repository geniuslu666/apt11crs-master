package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarHelp struct{}

func NewCarHelp() *sCarHelp {
	return &sCarHelp{}
}

func init() {
	service.RegisterCarHelp(NewCarHelp())
}

// Model HelpORM模型
func (s *sCarHelp) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarHelp.Ctx(ctx), option...)
}

// List 获取Help列表
func (s *sCarHelp) List(ctx context.Context, in *input_car.CarHelpListInp) (list []*input_car.CarHelpListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarHelpListModel{})

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.CarHelp.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取Help列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增Help
func (s *sCarHelp) Edit(ctx context.Context, in *input_car.CarHelpEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarHelpUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改Help失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarHelpInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增Help失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除Help
func (s *sCarHelp) Delete(ctx context.Context, in *input_car.CarHelpDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除Help失败，请稍后重试！")
		return
	}
	return
}

// View 获取Help指定信息
func (s *sCarHelp) View(ctx context.Context, in *input_car.CarHelpViewInp) (res *input_car.CarHelpViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取Help信息，请稍后重试！")
		return
	}
	return
}

// LanguageList 获取HelpLanguage列表
func (s *sCarHelp) LanguageList(ctx context.Context, in *input_car.CarHelpLanguageListInp) (list []*input_car.CarHelpLanguageListModel, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarHelpLanguageListModel{})

	// 查询数据
	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取HelpLanguage列表失败，请稍后重试！")
		return
	}
	return
}
