package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sHotelService) PropertyList(ctx context.Context, in *input_hotel.PmsPropertyListInp) (list []*input_hotel.PmsPropertyListModel, totalCount int, err error) {
	mod := dao.PmsProperty.Ctx(ctx)
	mod = mod.WithAll()

	mod = mod.Fields(input_hotel.PmsPropertyListModel{})

	if in.ContainDelete {
		mod = mod.Unscoped()
	}

	if in.Uid != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Uid, in.Uid)
	}
	if in.IsFindClose {
		mod = mod.Where(dao.PmsProperty.Columns().Close, 1)
		mod = mod.Where(dao.PmsProperty.Columns().BookingClose, 1)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Name, in.Name)
	}

	if !g.IsEmpty(in.PropertyStatus) {
		if in.PropertyStatus == 1 {
			mod = mod.Where(dao.PmsProperty.Columns().Close, 1)
		} else if in.PropertyStatus == 2 {
			mod = mod.Where(dao.PmsProperty.Columns().Close, 2)
		} else if in.PropertyStatus == 3 {
			mod = mod.Unscoped().WhereNotNull(dao.PmsProperty.Columns().DeletedAt)
		}
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsProperty.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderAsc(dao.PmsProperty.Columns().DeletedAt)
	if !g.IsEmpty(in.Sort) {
		if in.Sort == "ascend" {
			mod = mod.Order(dao.PmsProperty.Columns().Sort, "asc")
		} else {
			mod = mod.Order(dao.PmsProperty.Columns().Sort, "desc")
		}
	} else {
		mod = mod.Order(dao.PmsProperty.Columns().Sort, "desc").OrderDesc(dao.PmsProperty.Columns().Id)
	}

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取物业列表失败，请稍后重试！")
		return
	}

	var (
		puid   []interface{}
		Objeck gdb.Result
	)
	for _, v := range list {
		puid = append(puid, v.Uid)
	}
	if len(puid) > 0 {
		PmsRoomTypeOrm := g.Model(dao.PmsRoomType.Table())
		PmsRoomTypeOrm = PmsRoomTypeOrm.Fields(gdb.Raw("count(1) as roomTypeNum"), dao.PmsRoomType.Columns().Puid)
		PmsRoomTypeOrm = PmsRoomTypeOrm.WhereIn(dao.PmsRoomType.Columns().Puid, puid)
		PmsRoomTypeOrm = PmsRoomTypeOrm.Group(dao.PmsRoomType.Columns().Puid)
		if Objeck, err = PmsRoomTypeOrm.All(); err != nil {
			return
		}
		for _, RoomType := range Objeck {
			for k, v := range list {
				if RoomType.Map()["puid"] == v.Uid {
					list[k].RoomTypeNum = gvar.New(RoomType.Map()["roomTypeNum"]).Int()
				}
			}
		}
		PmsRoomUtilOrm := g.Model(dao.PmsRoomUnit.Table())
		PmsRoomUtilOrm = PmsRoomUtilOrm.InnerJoinOnFields(dao.PmsRoomType.Table(), dao.PmsRoomUnit.Columns().RtUid, "=", dao.PmsRoomType.Columns().Uid)
		PmsRoomUtilOrm = PmsRoomUtilOrm.Fields(gdb.Raw("count(1) as roomUnitNum"))
		PmsRoomUtilOrm = PmsRoomUtilOrm.Fields(dao.PmsRoomType.Columns().Puid)
		PmsRoomUtilOrm = PmsRoomUtilOrm.Group(dao.PmsRoomType.Columns().Puid)
		if Objeck, err = PmsRoomUtilOrm.All(); err != nil {
			return
		}
		for _, RoomUnit := range Objeck {
			for k, v := range list {
				if RoomUnit.Map()["puid"] == v.Uid {
					list[k].RoomUnitNum = gvar.New(RoomUnit.Map()["roomUnitNum"]).Int()
				}
			}
		}
	}

	return
}

func (s *sHotelService) PropertyAll(ctx context.Context, in *input_hotel.PmsPropertyAllInp) (list []*input_hotel.PmsPropertyAllModel, err error) {
	mod := dao.PmsProperty.Ctx(ctx)
	mod = mod.WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.Fields(input_hotel.PmsPropertyAllModel{})

	if in.Uid != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Uid, in.Uid)
	}

	if in.IsFindClose {
		//mod = mod.Where(dao.PmsProperty.Columns().Close, 1)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Name, in.Name)
	}

	if !g.IsEmpty(in.Ids) {
		mod = mod.WhereIn(dao.PmsProperty.Columns().Id, strings.Split(in.Ids, ","))
	}

	mod = mod.OrderDesc(dao.PmsProperty.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取物业列表失败，请稍后重试！")
		return
	}

	return
}

func (s *sHotelService) PropertyAppIndexList(ctx context.Context, in *input_hotel.PmsPropertyListInp) (list []*input_hotel.PmsPropertyAppIndexModel, totalCount int, err error) {
	mod := dao.PmsProperty.Ctx(ctx)
	mod = mod.WithAll()

	mod = mod.Fields(input_hotel.PmsPropertyAppIndexModel{})

	if in.ContainDelete {
		mod = mod.Unscoped()
	}

	if in.Uid != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Uid, in.Uid)
	}
	if in.IsFindClose {
		mod = mod.Where(dao.PmsProperty.Columns().Close, 1)
		mod = mod.Where(dao.PmsProperty.Columns().BookingClose, 1)
	}

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsProperty.Columns().Name, in.Name)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsProperty.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.Order(dao.PmsProperty.Columns().Sort, "desc").OrderDesc(dao.PmsProperty.Columns().Id)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
			err = gerror.Wrap(err, "获取物业列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil {
			err = gerror.Wrap(err, "获取物业列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sHotelService) PropertyExport(ctx context.Context, in *input_hotel.PmsPropertyListInp) (err error) {
	list, _, err := s.PropertyList(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_hotel.PmsPropertyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出物业-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("物业列表")
		exports   []input_hotel.PmsPropertyExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sHotelService) PropertyEdit(ctx context.Context, in *input_hotel.PmsPropertyEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			var (
				PmsProperty *entity.PmsProperty
			)
			if err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Scan(&PmsProperty); err != nil && errors.Is(err, sql.ErrNoRows) {
				return
			}

			if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).OmitEmpty().Data(in.PmsProperty).Update(); err != nil {
				err = gerror.Wrap(err, "修改物业失败，请稍后重试！")
			}

			if !g.IsEmpty(in.ShowType) && in.ShowType == 1 {
				if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Data(g.Map{
					dao.PmsProperty.Columns().GroupIds: nil,
				}).Update(); err != nil {
					err = gerror.Wrap(err, "修改物业失败，请稍后重试！")
				}
			}

			if !g.IsEmpty(in.CheckinAt) {
				if _, err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Puid, PmsProperty.Uid).
					Update(g.Map{
						dao.PmsRoomType.Columns().CheckinAt: in.CheckinAt,
					}); err != nil {
					return
				}
			}
			if !g.IsEmpty(in.CheckoutAt) {
				if _, err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Puid, PmsProperty.Uid).
					Update(g.Map{
						dao.PmsRoomType.Columns().CheckoutAt: in.CheckoutAt,
					}); err != nil {
					return
				}
			}
			return
		} else {

			if _, err = dao.PmsProperty.Ctx(ctx).Data(in.PmsProperty).OmitEmptyData().Insert(); err != nil {
				err = gerror.Wrap(err, "新增物业失败，请稍后重试！")
			}
		}
		return
	})
}

func (s *sHotelService) PropertyEditGroup(ctx context.Context, in *input_hotel.PmsPropertyEditGroupInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			var (
				PmsProperty *entity.PmsProperty
				GroupIdsStr string
			)
			err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Scan(&PmsProperty)
			if err != nil {
				err = gerror.Wrap(err, "修改物业失败，请稍后重试！")
				return
			}

			if !g.IsEmpty(in.GroupIds) {
				GroupIdsStr = strings.Join(gvar.New(in.GroupIds).Strings(), ",")
			} else {
				GroupIdsStr = ""
			}
			if _, err = dao.PmsProperty.Ctx(ctx).Fields(entity.PmsProperty{}).WherePri(in.Id).Data(g.Map{
				dao.PmsProperty.Columns().GroupIds: GroupIdsStr,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "修改物业失败，请稍后重试！")
			}
			return
		} else {
			err = gerror.Wrap(err, "修改物业失败，请稍后重试！")
		}
		return
	})
}

func (s *sHotelService) PropertyDelete(ctx context.Context, in *input_hotel.PmsPropertyDeleteInp) (err error) {
	if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除物业失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) PropertyRecycle(ctx context.Context, in *input_hotel.PmsPropertyRecycleInp) (err error) {
	if _, err = dao.PmsProperty.Ctx(ctx).Unscoped().WherePri(in.Id).Data(g.Map{
		dao.PmsProperty.Columns().DeletedAt: nil,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "回收物业失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) PropertyView(ctx context.Context, in *input_hotel.PmsPropertyViewInp) (res *input_hotel.PmsPropertyViewModel, resLanguage *input_hotel.PmsPropertyViewLannguageModel, err error) {
	res = new(input_hotel.PmsPropertyViewModel)
	resLanguage = new(input_hotel.PmsPropertyViewLannguageModel)
	if !in.IsLanguage {
		if err = dao.PmsProperty.Ctx(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook).OmitEmptyWhere().Where(&entity.PmsProperty{
			Id:  in.Id,
			Uid: in.Uid,
		}).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取物业信息，请稍后重试！")
			return
		}
	} else {
		if err = dao.PmsProperty.Ctx(ctx).
			OmitEmpty().
			Where(&entity.PmsProperty{
				Id:  in.Id,
				Uid: in.Uid,
			}).
			WithAll().
			Scan(&resLanguage); err != nil {
			err = gerror.Wrap(err, "获取物业信息，请稍后重试！")
			return
		}
	}

	return
}

func (s *sHotelService) PropertyGetUids(ctx context.Context, name []string) (uids []string, err error) {
	columns, err := dao.PmsProperty.Ctx(ctx).
		Fields("uid").
		WhereIn("name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取uid失败！")
		return
	}

	uids = g.NewVar(columns).Strings()
	return
}

// PropertySort 编辑物业排序
func (s *sHotelService) PropertySort(ctx context.Context, in *input_hotel.PmsPropertySortInp) (err error) {
	if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Update(g.Map{
		dao.PmsProperty.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改物业排序，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) SwitchSpaCanOrder(ctx context.Context, in *input_hotel.PropertySwitchSpaCanOrderInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.PmsProperty.Columns().SpaCanOrder: in.SpaCanOrder,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新预定状态失败，请稍后重试！")
			return
		}

		return
	})
}

func (s *sHotelService) UpdateAccessPass(ctx context.Context, in *input_hotel.PropertyUpdateAccessPassInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = dao.PmsProperty.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.PmsProperty.Columns().AccessPass: in.AccessPass,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新门禁密码失败，请稍后重试！")
			return
		}

		return
	})
}
