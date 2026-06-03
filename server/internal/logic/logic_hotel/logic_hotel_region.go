package logic_hotel

import (
	"APT/internal/dao"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_hotel"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sHotelService) RegionList(ctx context.Context, in *input_hotel.PropertyRegionListInp) (list []*input_hotel.PropertyRegionListModel, totalCount int, err error) {
	mod := dao.PmsPropertyRegion.Ctx(ctx).Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_hotel.PropertyRegionListModel{})

	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.PmsPropertyRegion.Columns().Status, in.Status)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsPropertyRegion.Columns().Id)
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取地区列表失败，请稍后重试！")
			return
		}
	} else {

		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取地区列表失败，请稍后重试！")
			return
		}
	}
	return
}
