package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hggen"
	"APT/internal/model"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"APT/utility/validate"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

type sBasicsGenCodes struct{}

func NewBasicsGenCodes() *sBasicsGenCodes {
	return &sBasicsGenCodes{}
}

func init() {
	service.RegisterBasicsGenCodes(NewBasicsGenCodes())
}

func (s *sBasicsGenCodes) Delete(ctx context.Context, in *input_basics.GenCodesDeleteInp) (err error) {
	_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

func (s *sBasicsGenCodes) Edit(ctx context.Context, in *input_basics.GenCodesEditInp) (res *input_basics.GenCodesEditModel, err error) {
	if in.GenType == 0 {
		err = gerror.New("生成类型不能为空")
		return
	}
	if in.VarName == "" {
		err = gerror.New("实体名称不能为空")
		return
	}

	if in.GenType == consts.GenCodesTypeCurd {
		var temp *model.GenerateAppCrudTemplate
		cfg := fmt.Sprintf("hggen.application.crud.templates.%v", in.GenTemplate)
		if err = g.Cfg().MustGet(ctx, cfg).Scan(&temp); err != nil {
			return
		}

		if temp == nil {
			err = gerror.Newf("选择的模板不存在:%v", cfg)
			return
		}

		if temp.IsAddon && in.AddonName == "" {
			err = gerror.New("插件模板必须选择一个有效的插件")
			return
		}
	}

	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		return &input_basics.GenCodesEditModel{SysGenCodes: in.SysGenCodes}, nil
	}

	if in.Options.IsNil() {
		in.Options = gjson.New(consts.NilJsonToString)
	}

	in.MasterColumns = gjson.New("[]")
	in.Status = consts.GenCodesStatusWait
	in.CreatedAt = gtime.Now()
	id, err := dao.SysGenCodes.Ctx(ctx).Data(in).OmitEmptyData().InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	in.Id = id
	return &input_basics.GenCodesEditModel{SysGenCodes: in.SysGenCodes}, nil
}

func (s *sBasicsGenCodes) Status(ctx context.Context, in *input_basics.GenCodesStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	_, err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

func (s *sBasicsGenCodes) MaxSort(ctx context.Context, in *input_basics.GenCodesMaxSortInp) (res *input_basics.GenCodesMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	if res == nil {
		res = new(input_basics.GenCodesMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sBasicsGenCodes) View(ctx context.Context, in *input_basics.GenCodesViewInp) (res *input_basics.GenCodesViewModel, err error) {
	err = dao.SysGenCodes.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

func (s *sBasicsGenCodes) List(ctx context.Context, in *input_basics.GenCodesListInp) (list []*input_basics.GenCodesListModel, totalCount int, err error) {
	mod := dao.SysGenCodes.Ctx(ctx)

	if in.GenType > 0 {
		mod = mod.Where("gen_type", in.GenType)
	}

	if in.VarName != "" {
		mod = mod.Where("var_name", in.VarName)
	}

	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	typeSelect, err := hggen.GenTypeSelect(ctx)
	if err != nil {
		return
	}

	getTpGroup := func(row *input_basics.GenCodesListModel) string {
		if row == nil {
			return ""
		}
		for _, v := range typeSelect {
			if v.Value == int(row.GenType) {
				for index, template := range v.Templates {
					if index == row.GenTemplate {
						return template.Label
					}
				}
			}
		}

		return ""
	}

	if len(list) > 0 {
		for _, v := range list {
			v.GenTemplateGroup = getTpGroup(v)
		}
	}
	return
}

func (s *sBasicsGenCodes) Selects(ctx context.Context, in *input_basics.GenCodesSelectsInp) (res *input_basics.GenCodesSelectsModel, err error) {
	return hggen.TableSelects(ctx, in)
}

func (s *sBasicsGenCodes) TableSelect(ctx context.Context, in *input_basics.GenCodesTableSelectInp) (res []*input_basics.GenCodesTableSelectModel, err error) {
	var (
		sql           = "SELECT TABLE_NAME as value, TABLE_COMMENT as label FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = '%s'"
		config        = g.DB(in.Name).GetConfig()
		disableTables = g.Cfg().MustGet(ctx, "hggen.disableTables").Strings()
		lists         []*input_basics.GenCodesTableSelectModel
	)

	if err = g.DB(in.Name).Ctx(ctx).Raw(fmt.Sprintf(sql, config.Name)).Scan(&lists); err != nil {
		return
	}

	patternStr := `addon_(\w+)_`
	repStr := ``

	for _, v := range lists {
		if gstr.InArray(disableTables, v.Value) {
			continue
		}

		newValue := v.Value
		if config.Prefix != "" {
			newValue = gstr.SubStrFromEx(v.Value, config.Prefix)
		}
		if newValue == "" {
			err = gerror.Newf("表名[%v]前缀必须和配置中的前缀设置[%v] 保持一致", v.Value, config.Prefix)
			return
		}

		bt, err := gregex.Replace(patternStr, []byte(repStr), []byte(newValue))
		if err != nil {
			err = gerror.Newf("表名[%v] gregex.Replace err:%v", v.Value, err.Error())
			return nil, err
		}

		row := v
		row.DefTableComment = v.Label
		row.DaoName = gstr.CaseCamel(newValue)
		row.DefVarName = gstr.CaseCamel(string(bt))
		row.DefAlias = gstr.CaseCamelLower(newValue)
		row.Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		row.Label = row.Name
		res = append(res, row)
	}
	return
}

func (s *sBasicsGenCodes) ColumnSelect(ctx context.Context, in *input_basics.GenCodesColumnSelectInp) (res []*input_basics.GenCodesColumnSelectModel, err error) {
	var (
		sql    = "select COLUMN_NAME as value,COLUMN_COMMENT as label from information_schema.COLUMNS where TABLE_SCHEMA = '%s' and TABLE_NAME = '%s'"
		config = g.DB(in.Name).GetConfig()
	)

	if err = g.DB(in.Name).Ctx(ctx).Raw(fmt.Sprintf(sql, config.Name, in.Table)).Scan(&res); err != nil {
		return
	}

	if len(res) == 0 {
		err = gerror.Newf("table does not exist:%v", in.Table)
		return
	}

	for k, v := range res {
		res[k].Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		res[k].Label = res[k].Name
	}
	return
}

func (s *sBasicsGenCodes) ColumnList(ctx context.Context, in *input_basics.GenCodesColumnListInp) (res []*input_basics.GenCodesColumnListModel, err error) {
	return hggen.TableColumns(ctx, in)
}

func (s *sBasicsGenCodes) Preview(ctx context.Context, in *input_basics.GenCodesPreviewInp) (res *input_basics.GenCodesPreviewModel, err error) {
	return hggen.Preview(ctx, in)
}

func (s *sBasicsGenCodes) Build(ctx context.Context, in *input_basics.GenCodesBuildInp) (err error) {

	ein := in.SysGenCodes
	if _, err = s.Edit(ctx, &input_basics.GenCodesEditInp{SysGenCodes: ein}); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if err = s.Status(ctx, &input_basics.GenCodesStatusInp{Id: in.Id, Status: consts.GenCodesStatusOk}); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if err = hggen.Build(ctx, in); err != nil {
		_ = s.Status(ctx, &input_basics.GenCodesStatusInp{Id: in.Id, Status: consts.GenCodesStatusFail})
		return err
	}
	return
}
