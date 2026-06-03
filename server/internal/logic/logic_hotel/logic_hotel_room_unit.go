package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
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

func (s *sHotelService) RoomUnitList(ctx context.Context, in *input_hotel.PmsRoomUnitListInp) (list []*input_hotel.PmsRoomUnitListModel, totalCount int, err error) {
	var (
		Objeck []gdb.Value
	)
	mod := dao.PmsRoomUnit.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.Fields(input_hotel.PmsRoomUnitListModel{})

	if in.RoomNo != "" {
		mod = mod.WhereLike(dao.PmsRoomUnit.Columns().RoomNo, in.RoomNo)
	}
	if !g.IsEmpty(in.Puid) {
		if Objeck, err = g.Model(dao.PmsRoomType.Table()).Ctx(ctx).Where(dao.PmsRoomType.Columns().Puid, in.Puid).Fields(dao.PmsRoomType.Columns().Uid).Array(); err != nil {
			return
		}
		if !g.IsEmpty(Objeck) {
			mod = mod.Where(dao.PmsRoomUnit.Columns().RtUid, Objeck)
		}
	}
	if !g.IsEmpty(in.RtUid) {
		mod = mod.Where(dao.PmsRoomUnit.Columns().RtUid, in.RtUid)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsRoomUnit.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsRoomUnit.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取房间列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) RoomUnitExport(ctx context.Context, in *input_hotel.PmsRoomUnitListInp) (err error) {
	list, _, err := s.RoomUnitList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_hotel.PmsRoomUnitExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出房间-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("房间列表")
		exports   []input_hotel.PmsRoomUnitExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sHotelService) RoomUnitEdit(ctx context.Context, in *input_hotel.PmsRoomUnitEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsRoomUnit.Ctx(ctx).
				Fields(input_hotel.PmsRoomUnitUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改房间失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsRoomUnit.Ctx(ctx).
			Fields(input_hotel.PmsRoomUnitInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增房间失败，请稍后重试！")
		}
		return
	})
}

func (s *sHotelService) RoomUnitDelete(ctx context.Context, in *input_hotel.PmsRoomUnitDeleteInp) (err error) {

	if _, err = dao.PmsRoomUnit.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除房间失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) RoomUnitView(ctx context.Context, in *input_hotel.PmsRoomUnitViewInp) (res *input_hotel.PmsRoomUnitViewModel, err error) {
	if err = dao.PmsRoomUnit.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取房间信息，请稍后重试！")
		return
	}
	return
}
