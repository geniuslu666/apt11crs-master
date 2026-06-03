package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSpaLabel struct{}

func NewSpaLabel() *sSpaLabel {
	return &sSpaLabel{}
}

func init() {
	service.RegisterSpaLabel(NewSpaLabel())
}

// Model 服务标签ORM模型
func (s *sSpaLabel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaLabel.Ctx(ctx), option...)
}

// List 获取服务标签列表
func (s *sSpaLabel) List(ctx context.Context, in *input_spa.SpaLabelListInp) (list []*input_spa.SpaLabelListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_spa.SpaLabelListModel{})

	// 查询标签名称
	if !g.IsEmpty(in.LabelName) {
		mod = mod.WhereLike(dao.SpaLabel.Columns().LabelName, "%"+in.LabelName+"%")
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SpaLabel.Columns().Status, in.Status)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SpaLabel.Columns().Sort).OrderDesc(dao.SpaLabel.Columns().Id)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务标签列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务标签列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Edit 修改/新增服务标签
func (s *sSpaLabel) Edit(ctx context.Context, in *input_spa.SpaLabelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object         gdb.Record
			Dao            *input_language.LoadLanguage
			LanguageStruct input_language.LanguageModel
		)
		Uuid := guid.S([]byte("content"))

		// 修改
		if in.Id > 0 {
			if Object, err = dao.SpaLabel.Ctx(ctx).Where(dao.SpaLabel.Columns().Id, in.Id).One(); err != nil {
				return
			}
			// 名称多语言
			// 是否存在languageUuid
			if !g.IsEmpty(Object[gstr.ToLower("LabelContent")]) {
				Uuid = Object["label_content"].String()
			}
			Dao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.SpaLabel.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("LabelContent"),
			}
			LanguageStruct = in.ContentLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
				return
			}
			// 修改
			in.LabelContent = Uuid

			if _, err = s.Model(ctx).
				Fields(input_spa.SpaLabelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改服务标签失败，请稍后重试！")
			}
			return
		}
		in.LabelContent = Uuid

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaLabelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增服务标签失败，请稍后重试！")
		}

		// 多语言
		Dao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.SpaLabel.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("LabelContent"),
		}
		LanguageStruct = in.ContentLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, Dao); err != nil {
			return
		}
		return
	})
}

// Delete 删除服务标签
func (s *sSpaLabel) Delete(ctx context.Context, in *input_spa.SpaLabelDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除服务标签失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取服务标签最大排序
func (s *sSpaLabel) MaxSort(ctx context.Context, in *input_spa.SpaLabelMaxSortInp) (res *input_spa.SpaLabelMaxSortModel, err error) {
	if err = dao.SpaLabel.Ctx(ctx).Fields(dao.SpaLabel.Columns().Sort).OrderDesc(dao.SpaLabel.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取服务标签最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_spa.SpaLabelMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取服务标签指定信息
func (s *sSpaLabel) View(ctx context.Context, in *input_spa.SpaLabelViewInp) (res *input_spa.SpaLabelViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务标签信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务标签信息，请稍后重试！")
			return
		}
	}

	return
}

// Status 更新服务标签状态
func (s *sSpaLabel) Status(ctx context.Context, in *input_spa.SpaLabelStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaLabel.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新服务标签状态失败，请稍后重试！")
		return
	}
	return
}
