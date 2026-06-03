package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type sBasicsIndexNav struct{}

func NewBasicsIndexNav() *sBasicsIndexNav {
	return &sBasicsIndexNav{}
}

func init() {
	service.RegisterBasicsIndexNav(NewBasicsIndexNav())
}

func (s *sBasicsIndexNav) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsIndexNav.Ctx(ctx), option...)
}

func (s *sBasicsIndexNav) List(ctx context.Context, in *input_basics.PmsIndexNavListInp) (list []*input_basics.PmsIndexNavListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsIndexNavListModel{})
	mod = mod.Hook(hook.PmsFindLanguageValueHook)
	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsIndexNav.Columns().Sort)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sBasicsIndexNav) ApiNavAll(ctx context.Context, in *input_basics.PmsIndexNavAllInp) (list []*input_basics.PmsIndexNavApiListModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsIndexNavApiListModel{})

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsIndexNav.Columns().Status, in.Status)
	}
	if !g.IsEmpty(in.MinappStatus) {
		mod = mod.Where(dao.PmsIndexNav.Columns().MinappStatus, in.MinappStatus)
	}
	mod = mod.OrderDesc(dao.PmsIndexNav.Columns().Sort)
	mod = mod.Hook(hook.PmsFindLanguageValueHook)
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsIndexNav) All(ctx context.Context, in *input_basics.PmsIndexNavAllInp) (list []*input_basics.PmsIndexNavAllModel, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsIndexNavAllModel{})

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsIndexNav.Columns().Status, in.Status)
	}
	mod = mod.OrderDesc(dao.PmsIndexNav.Columns().Sort)
	mod = mod.Hook(hook.PmsFindLanguageValueHook)
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取导航列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsIndexNav) Edit(ctx context.Context, in *input_basics.PmsIndexNavEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object             gdb.Record
			IndexNavNameDao    *input_language.LoadLanguage
			NameLanguageStruct input_language.LanguageModel

			IndexNavTagDao    *input_language.LoadLanguage
			TagLanguageStruct input_language.LanguageModel
		)
		NameUuid := guid.S([]byte("name"))
		TagUuid := guid.S([]byte("tag"))

		if in.Id > 0 {
			if Object, err = dao.PmsIndexNav.Ctx(ctx).Where(dao.PmsIndexNav.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
				NameUuid = Object["name"].String()
			}
			IndexNavNameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.PmsIndexNav.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			NameLanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, IndexNavNameDao); err != nil {
				return
			}

			in.Name = NameUuid

			if !g.IsEmpty(Object[gstr.ToLower("Tag")]) {
				TagUuid = Object["tag"].String()
			}
			IndexNavTagDao = &input_language.LoadLanguage{
				Uuid: TagUuid,
				Tag:  dao.PmsIndexNav.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Tag"),
			}
			TagLanguageStruct = in.TagLanguage
			if err = service.BasicsLanguage().Sync(ctx, TagLanguageStruct, IndexNavTagDao); err != nil {
				return
			}

			in.Tag = TagUuid

			if _, err = s.Model(ctx).
				Fields(input_basics.PmsIndexNavUpdateFields{}).
				WherePri(in.Id).Data(in).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "修改失败，请稍后重试！")
			}
			return
		}

		var (
			lastInsertId int64
		)
		in.Name = NameUuid
		in.Tag = TagUuid

		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsIndexNavInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		IndexNavNameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.PmsIndexNav.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		NameLanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, IndexNavNameDao); err != nil {
			return
		}

		IndexNavTagDao = &input_language.LoadLanguage{
			Uuid: TagUuid,
			Tag:  dao.PmsIndexNav.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Tag"),
		}
		TagLanguageStruct = in.TagLanguage
		if err = service.BasicsLanguage().Sync(ctx, TagLanguageStruct, IndexNavTagDao); err != nil {
			return
		}
		return
	})
}

func (s *sBasicsIndexNav) Delete(ctx context.Context, in *input_basics.PmsIndexNavDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsIndexNav) View(ctx context.Context, in *input_basics.PmsIndexNavViewInp) (res *input_basics.PmsIndexNavViewModel, err error) {

	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取信息失败，请稍后重试！")
		return
	}

	return
}

func (s *sBasicsIndexNav) Switch(ctx context.Context, in *input_basics.PmsIndexNavSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.PmsIndexNav.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}

// MinappStatus 更新小程序显示状态
func (s *sBasicsIndexNav) MinappStatus(ctx context.Context, in *input_basics.PmsIndexNavSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.PmsIndexNav.Columns().MinappStatus: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}

// NavSort 编辑排序
func (s *sBasicsIndexNav) NavSort(ctx context.Context, in *input_basics.PmsIndexNavSortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Update(g.Map{
		dao.PmsIndexNav.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改排序失败，请稍后重试！")
		return
	}
	return
}
