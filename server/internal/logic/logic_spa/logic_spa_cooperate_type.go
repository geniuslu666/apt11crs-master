package logic_spa

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSpaCooperateType struct{}

func NewSpaCooperateType() *sSpaCooperateType {
	return &sSpaCooperateType{}
}

func init() {
	service.RegisterSpaCooperateType(NewSpaCooperateType())
}

// Model 按摩营业类型ORM模型
func (s *sSpaCooperateType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaCooperateType.Ctx(ctx), option...)
}

// List 获取按摩营业类型列表
func (s *sSpaCooperateType) List(ctx context.Context, in *input_spa.SpaCooperateTypeListInp) (list []*input_spa.SpaCooperateTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_spa.SpaCooperateTypeListModel{})

	// 合作类型
	if !g.IsEmpty(in.TypeName) {
		mod = mod.WhereLike(dao.SpaCooperateType.Columns().TypeName, "%"+in.TypeName+"%")
	}

	// 查询状态1、启用 2、禁用
	if in.Status > 0 {
		mod = mod.Where(dao.SpaCooperateType.Columns().Status, in.Status)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.SpaCooperateType.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取按摩营业类型列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出按摩营业类型
func (s *sSpaCooperateType) Export(ctx context.Context, in *input_spa.SpaCooperateTypeListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(input_spa.SpaCooperateTypeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出按摩营业类型-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("按摩营业类型")
		exports   []input_spa.SpaCooperateTypeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增按摩营业类型
func (s *sSpaCooperateType) Edit(ctx context.Context, in *input_spa.SpaCooperateTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_spa.SpaCooperateTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改按摩营业类型失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_spa.SpaCooperateTypeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增按摩营业类型失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除按摩营业类型
func (s *sSpaCooperateType) Delete(ctx context.Context, in *input_spa.SpaCooperateTypeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除按摩营业类型失败，请稍后重试！")
		return
	}
	return
}

// View 获取按摩营业类型指定信息
func (s *sSpaCooperateType) View(ctx context.Context, in *input_spa.SpaCooperateTypeViewInp) (res *input_spa.SpaCooperateTypeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取按摩营业类型信息，请稍后重试！")
		return
	}
	return
}

// Status 更新按摩营业类型状态
func (s *sSpaCooperateType) Status(ctx context.Context, in *input_spa.SpaCooperateTypeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SpaCooperateType.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新按摩营业类型状态失败，请稍后重试！")
		return
	}
	return
}
