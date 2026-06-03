package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type sSysCarAddress struct{}

func NewSysCarAddress() *sSysCarAddress {
	return &sSysCarAddress{}
}

func init() {
	service.RegisterSysCarAddress(NewSysCarAddress())
}

// Model 接送机地址ORM模型
func (s *sSysCarAddress) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarAddress.Ctx(ctx), option...)
}

// List 获取接送机地点管理列表
func (s *sSysCarAddress) List(ctx context.Context, in *input_car.CarAddressListInp) (list []*input_car.CarAddressListModel, totalCount int, err error) {
	mod := dao.CarAddress.Ctx(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarAddressListModel{})

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.CarAddress.Columns().Id, in.Id)
	}

	// 查询地点名称（后台）
	if in.Name != "" {
		mod = mod.WhereLike(dao.CarAddress.Columns().Name, "%"+in.Name+"%")
	}

	// 查询状态1、启用 2、禁用
	if in.Status > 0 {
		mod = mod.Where(dao.CarAddress.Columns().Status, in.Status)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.CarAddress.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取接送机地点列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取接送机地址列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Export 导出接送机地点管理
func (s *sSysCarAddress) Export(ctx context.Context, in *input_car.CarAddressListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(input_car.CarAddressExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = "导出接送机地点管理-" + gctx.CtxId(ctx)
		//sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, input_form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		sheetName = fmt.Sprintf("接送机地点管理")
		exports   []input_car.CarAddressExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增接送机地点管理
func (s *sSysCarAddress) Edit(ctx context.Context, in *input_car.CarAddressEditInp) (err error) {
	// 接送机地点名不能重复
	if err = hgorm.IsUnique(ctx, &dao.CarAddress, g.Map{
		dao.CarAddress.Columns().Name: in.Name,
	}, "地点名称已存在", in.Id); err != nil {
		return
	}

	// 验证'AirportCode'唯一
	if in.TypeId == 1 {
		if err = hgorm.IsUnique(ctx, &dao.CarAddress, g.Map{
			dao.CarAddress.Columns().AirportCode: in.AirportCode,
			dao.CarAddress.Columns().TypeId:      1,
		}, "机场代码已存在", in.Id); err != nil {
			return
		}
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			Object               gdb.Record
			CarAddressSubNameDao *input_language.LoadLanguage
			CarAddressAddressDao *input_language.LoadLanguage
			LanguageStruct       input_language.LanguageModel
			AddressStruct        input_language.LanguageModel
		)
		Uuid := guid.S([]byte("sub_name"))

		AddressUuid := guid.S([]byte("detail_address"))

		// 修改
		if in.Id > 0 {

			if Object, err = dao.CarAddress.Ctx(ctx).Where(dao.CarAddress.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object["sub_name"]) {
				Uuid = Object["sub_name"].String()
				AddressUuid = Object["detail_address"].String()
			}

			CarAddressSubNameDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.CarAddress.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("SubName"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, CarAddressSubNameDao); err != nil {
				return
			}

			CarAddressAddressDao = &input_language.LoadLanguage{
				Uuid: AddressUuid,
				Tag:  dao.PmsCouponType.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("DetailAddress"),
			}
			AddressStruct = in.AddressLanguage
			if err = service.BasicsLanguage().Sync(ctx, AddressStruct, CarAddressAddressDao); err != nil {
				return
			}

			in.SubName = Uuid
			in.DetailAddress = AddressUuid
			if _, err = s.Model(ctx).
				Fields(input_car.CarAddressUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改接送机地点管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		var (
			lastInsertId int64
		)
		in.SubName = Uuid
		in.DetailAddress = AddressUuid
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarAddressInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增接送机地点管理失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增接送机地点管理失败，请稍后重试！")
			return
		}

		CarAddressSubNameDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.CarAddress.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("SubName"),
		}
		LanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, CarAddressSubNameDao); err != nil {
			return
		}

		CarAddressAddressDao = &input_language.LoadLanguage{
			Uuid: AddressUuid,
			Tag:  dao.PmsCouponType.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("DetailAddress"),
		}
		AddressStruct = in.AddressLanguage
		if err = service.BasicsLanguage().Sync(ctx, AddressStruct, CarAddressAddressDao); err != nil {
			return
		}

		return
	})
}

// Delete 删除接送机地点管理
func (s *sSysCarAddress) Delete(ctx context.Context, in *input_car.CarAddressDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除接送机地点管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取接送机地点管理指定信息
func (s *sSysCarAddress) View(ctx context.Context, in *input_car.CarAddressViewInp) (res *input_car.CarAddressViewModel, err error) {
	if err = dao.CarAddress.Ctx(ctx).WherePri(in.Id).WithAll().Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取接送机地点管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新接送机地点管理状态
func (s *sSysCarAddress) Status(ctx context.Context, in *input_car.CarAddressStatusInp) (err error) {
	if _, err = dao.CarAddress.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarAddress.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新接送机地点管理状态失败，请稍后重试！")
		return
	}
	return
}
