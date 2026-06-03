package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_hotel"
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

func (s *sHotelService) GuestList(ctx context.Context, in *input_hotel.PmsGuestProfileListInp) (list []*input_hotel.PmsGuestProfileListModel, totalCount int, err error) {
	mod := dao.PmsGuestProfile.Ctx(ctx)

	mod = mod.Fields(input_hotel.PmsGuestProfileListModel{})

	if in.FullName != "" {
		mod = mod.WhereLike(dao.PmsGuestProfile.Columns().FullName, in.FullName)
	}

	if in.Email != "" {
		mod = mod.WhereLike(dao.PmsGuestProfile.Columns().Email, in.Email)
	}

	if in.Phone != "" {
		mod = mod.WhereLike(dao.PmsGuestProfile.Columns().Phone, in.Phone)
	}

	if in.Nationality != "" {
		mod = mod.WhereLike(dao.PmsGuestProfile.Columns().Nationality, in.Nationality)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsGuestProfile.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsGuestProfile.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取客户档案列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) GuestExport(ctx context.Context, in *input_hotel.PmsGuestProfileListInp) (err error) {
	list, _, err := s.GuestList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_hotel.PmsGuestProfileExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出客户档案-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("客户档案")
		exports   []input_hotel.PmsGuestProfileExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sHotelService) GuestEdit(ctx context.Context, in *input_hotel.PmsGuestProfileEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		//if !g.IsEmpty(in.Id) {
		if _, err = dao.PmsGuestProfile.Ctx(ctx).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改客户档案失败，请稍后重试！")
		}
		return
		//}

		//
		//if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		//	Fields(input_hotel.PmsGuestProfileInsertFields{}).
		//	Data(in).OmitEmptyData().Insert(); err != nil {
		//	err = gerror.Wrap(err, "新增客户档案失败，请稍后重试！")
		//}
		//return
	})
}

func (s *sHotelService) GuestDelete(ctx context.Context, in *input_hotel.PmsGuestProfileDeleteInp) (err error) {

	if _, err = dao.PmsGuestProfile.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除客户档案失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) GuestView(ctx context.Context, in *input_hotel.PmsGuestProfileViewInp) (res *input_hotel.PmsGuestProfileViewModel, err error) {
	if err = dao.PmsGuestProfile.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取客户档案信息，请稍后重试！")
		return
	}
	return
}
