package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

func (s *sHotelService) CancelRateList(ctx context.Context, in *input_hotel.PmsCancelRateListInp) (list []*input_hotel.PmsCancelRateListModel, totalCount int, err error) {
	mod := dao.PmsCancelRate.Ctx(ctx).WithAll()

	mod = mod.Fields(input_hotel.PmsCancelRateListModel{})

	//mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.PmsCancelRate.Columns().Sort).OrderDesc(dao.PmsCancelRate.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取取消政策列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) CancelRateEdit(ctx context.Context, in *input_hotel.PmsCancelRateEditInp) (err error) {
	var (
		CancelRate        *entity.PmsCancelRate
		CancelRateNameUid string
	)
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			in.Name = ""
			if !g.IsEmpty(in.NameLanguage) {
				// 查询数据
				if err = dao.PmsCancelRate.Ctx(ctx).WherePri(in.Id).Scan(&CancelRate); err != nil {
					return
				}
				if !g.IsEmpty(CancelRate) && !g.IsEmpty(CancelRate.Name) {
					CancelRateNameUid = CancelRate.Name
				} else {
					CancelRateNameUid = guid.S()
				}
				// 更新名字多语言数据
				if err = service.BasicsLanguage().Sync(ctx, in.NameLanguage, &input_language.LoadLanguage{
					Uuid: CancelRateNameUid,
					Tag:  dao.PmsCancelRate.Table(),
					Type: "table",
					Key:  dao.PmsCancelRate.Columns().Name,
				}); err != nil {
					return
				}
				in.Name = CancelRateNameUid
			}
			if _, err = dao.PmsCancelRate.Ctx(ctx).OmitEmptyData().
				Fields(input_hotel.PmsCancelRateUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改取消政策失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsCancelRate.Ctx(ctx).
			Fields(input_hotel.PmsCancelRateInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增取消政策失败，请稍后重试！")
		}
		return
	})
}

func (s *sHotelService) CancelRateDelete(ctx context.Context, in *input_hotel.PmsCancelRateDeleteInp) (err error) {

	if _, err = dao.PmsCancelRate.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除取消政策失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) CancelRateMaxSort(ctx context.Context, in *input_hotel.PmsCancelRateMaxSortInp) (res *input_hotel.PmsCancelRateMaxSortModel, err error) {
	if err = dao.PmsCancelRate.Ctx(ctx).Fields(dao.PmsCancelRate.Columns().Sort).OrderDesc(dao.PmsCancelRate.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取取消政策最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_hotel.PmsCancelRateMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

func (s *sHotelService) CancelRateView(ctx context.Context, in *input_hotel.PmsCancelRateViewInp) (res *input_hotel.PmsCancelRateViewModel, err error) {
	if err = dao.PmsCancelRate.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取取消政策信息，请稍后重试！")
		return
	}
	return
}
