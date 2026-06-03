package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
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

type sCarCooperateType struct{}

func NewCarCooperateType() *sCarCooperateType {
	return &sCarCooperateType{}
}

func init() {
	service.RegisterCarCooperateType(NewCarCooperateType())
}

// Model 接送机营业类型ORM模型
func (s *sCarCooperateType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarCooperateType.Ctx(ctx), option...)
}

// List 获取接送机营业类型列表
func (s *sCarCooperateType) List(ctx context.Context, in *input_car.CarCooperateTypeListInp) (list []*input_car.CarCooperateTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarCooperateTypeListModel{})

	// 合作类型
	if !g.IsEmpty(in.TypeName) {
		mod = mod.WhereLike(dao.CarCooperateType.Columns().TypeName, "%"+in.TypeName+"%")
	}

	// 查询状态1、启用 2、禁用
	if in.Status > 0 {
		mod = mod.Where(dao.CarCooperateType.Columns().Status, in.Status)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.CarCooperateType.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取接送机营业类型列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出接送机营业类型
func (s *sCarCooperateType) Export(ctx context.Context, in *input_car.CarCooperateTypeListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(input_car.CarCooperateTypeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出接送机营业类型-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("接送机营业类型")
		exports   []input_car.CarCooperateTypeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增接送机营业类型
func (s *sCarCooperateType) Edit(ctx context.Context, in *input_car.CarCooperateTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_car.CarCooperateTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改接送机营业类型失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarCooperateTypeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增接送机营业类型失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除接送机营业类型
func (s *sCarCooperateType) Delete(ctx context.Context, in *input_car.CarCooperateTypeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除接送机营业类型失败，请稍后重试！")
		return
	}
	return
}

// View 获取接送机营业类型指定信息
func (s *sCarCooperateType) View(ctx context.Context, in *input_car.CarCooperateTypeViewInp) (res *input_car.CarCooperateTypeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取接送机营业类型信息，请稍后重试！")
		return
	}
	return
}

// Status 更新接送机营业类型状态
func (s *sCarCooperateType) Status(ctx context.Context, in *input_car.CarCooperateTypeStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarCooperateType.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新接送机营业类型状态失败，请稍后重试！")
		return
	}
	return
}
