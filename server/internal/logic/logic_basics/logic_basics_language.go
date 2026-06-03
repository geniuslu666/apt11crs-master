package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"reflect"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsLanguage struct{}

func NewBasicsLanguage() *sBasicsLanguage {
	return &sBasicsLanguage{}
}

func init() {
	service.RegisterBasicsLanguage(NewBasicsLanguage())
}

func (s *sBasicsLanguage) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsLanguage.Ctx(ctx), option...)
}

func (s *sBasicsLanguage) List(ctx context.Context, in *input_language.PmsLanguageListInp) (list []*input_language.PmsLanguageListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_language.PmsLanguageListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsLanguage.Columns().Id, in.Id)
	}

	if in.Uuid != "" {
		mod = mod.Where(dao.PmsLanguage.Columns().Uuid, in.Uuid)
	}

	if in.Tag != "" {
		mod = mod.Where(dao.PmsLanguage.Columns().Tag, in.Tag)
	}

	if in.Type != "" {
		mod = mod.Where(dao.PmsLanguage.Columns().Type, in.Type)
	}

	if in.Language != "" {
		mod = mod.Where(dao.PmsLanguage.Columns().Language, in.Language)
	}

	if in.Content != "" {
		mod = mod.WhereLike(dao.PmsLanguage.Columns().Content, "%"+in.Content+"%")
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsLanguage.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取语言字典列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsLanguage) Export(ctx context.Context, in *input_language.PmsLanguageListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_language.PmsLanguageExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出语言字典-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("语言字典")
		exports   []input_language.PmsLanguageExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sBasicsLanguage) Edit(ctx context.Context, in *input_language.PmsLanguageEditInp) (err error) {

	if err = hgorm.IsUnique(ctx, &dao.PmsLanguage, g.Map{dao.PmsLanguage.Columns().Uuid: in.Uuid}, "标签ID已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_language.PmsLanguageUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改语言字典失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_language.PmsLanguageInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增语言字典失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsLanguage) Delete(ctx context.Context, in *input_language.PmsLanguageDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除语言字典失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsLanguage) View(ctx context.Context, in *input_language.PmsLanguageViewInp) (res *input_language.PmsLanguageViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取语言字典信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsLanguage) Sync(ctx context.Context, LanguageStruct input_language.LanguageModel, In *input_language.LoadLanguage) (err error) {
	var (
		Object       gdb.Result
		LanguageInfo *entity.PmsLanguage
	)
	if Object, err = s.Model(ctx).Where(g.Map{
		dao.PmsLanguage.Columns().Uuid: In.Uuid,
		dao.PmsLanguage.Columns().Tag:  In.Tag,
		dao.PmsLanguage.Columns().Type: In.Type,
		dao.PmsLanguage.Columns().Key:  In.Key,
	}).All(); err != nil {
		return
	}
	LanguageStructValue := reflect.ValueOf(LanguageStruct)
	LanguageStructType := reflect.TypeOf(LanguageStruct)
	if !g.IsEmpty(Object) {
		g.Log().Info(ctx, LanguageStructType.NumField())
		for i := 0; i < LanguageStructType.NumField(); i++ {
			LanguageInfo = nil
			Name := LanguageStructType.Field(i).Tag.Get("json")
			if err = s.Model(ctx).Where(g.Map{
				dao.PmsLanguage.Columns().Uuid:     In.Uuid,
				dao.PmsLanguage.Columns().Tag:      In.Tag,
				dao.PmsLanguage.Columns().Type:     In.Type,
				dao.PmsLanguage.Columns().Key:      In.Key,
				dao.PmsLanguage.Columns().Language: Name,
			}).Scan(&LanguageInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
				return
			}
			UpdateData := g.MapStrAny{
				//dao.PmsLanguage.Columns().UpdateAt: gtime.Now().Unix(),
				dao.PmsLanguage.Columns().Content:  LanguageStructValue.Field(i).String(),
				dao.PmsLanguage.Columns().Uuid:     In.Uuid,
				dao.PmsLanguage.Columns().Tag:      In.Tag,
				dao.PmsLanguage.Columns().Type:     In.Type,
				dao.PmsLanguage.Columns().Key:      In.Key,
				dao.PmsLanguage.Columns().Language: Name,
			}
			if !g.IsEmpty(LanguageInfo) {
				UpdateDataWhere := g.MapStrAny{
					dao.PmsLanguage.Columns().Uuid:     In.Uuid,
					dao.PmsLanguage.Columns().Language: gstr.ToLower(Name),
				}
				if _, err = s.Model(ctx).Where(UpdateDataWhere).Data(UpdateData).Update(); err != nil {
					return err
				}
			} else {
				if _, err = s.Model(ctx).OmitEmptyData().Data(UpdateData).Insert(); err != nil {
					return err
				}
			}
		}
	} else {
		var InsertDataArrStruct []*entity.PmsLanguage
		for i := 0; i < LanguageStructType.NumField(); i++ {
			Name := LanguageStructType.Field(i).Tag.Get("json")
			InsertStruct := &entity.PmsLanguage{}
			InsertStruct.Uuid = In.Uuid
			InsertStruct.Tag = In.Tag
			InsertStruct.Type = In.Type
			InsertStruct.Key = In.Key
			InsertStruct.Language = Name
			if LanguageStructValue.Field(i).IsValid() {
				InsertStruct.Content = LanguageStructValue.Field(i).String()
			} else {
				InsertStruct.Content = ""
			}
			g.Dump(LanguageStructValue.Field(i).String())
			InsertDataArrStruct = append(InsertDataArrStruct, InsertStruct)
		}
		if _, err = s.Model(ctx).OmitEmptyData().Data(InsertDataArrStruct).Insert(); err != nil {
			return err
		}
	}
	return
}

func (s *sBasicsLanguage) GetUuids(ctx context.Context, content string, key string) (uuids []string, err error) {
	columns, err := s.Model(ctx).
		Fields("uuid").
		Where("key", key).
		WhereLike("content", "%"+content+"%").Array()
	if err != nil {
		err = gerror.Wrap(err, "获取uuid失败！")
		return
	}

	uuids = g.NewVar(columns).Strings()
	return
}
