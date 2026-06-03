package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

func (s *sHotelService) RoomTypeList(ctx context.Context, in *input_hotel.PmsRoomTypeListInp) (list []*input_hotel.PmsRoomTypeListModel, totalCount int, err error) {
	mod := dao.PmsRoomType.Ctx(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.PmsRoomType.Table(), input_hotel.PmsRoomTypeListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_hotel.PmsRoomTypeListModel{}, &dao.PmsProperty, "pmsProperty"))

	mod = mod.InnerJoinOnFields(dao.PmsProperty.Table(), dao.PmsRoomType.Columns().Puid, "=", dao.PmsProperty.Columns().Uid)

	if in.Name != "" {
		mod = mod.WhereLike(dao.PmsRoomType.Columns().Name, in.Name)
	}

	if !g.IsEmpty(in.Puid) {
		mod = mod.WherePrefix(dao.PmsProperty.Table(), dao.PmsProperty.Columns().Uid, in.Puid)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsRoomType.Table() + "." + dao.PmsRoomType.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
			err = gerror.Wrap(err, "获取房型列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil {
			err = gerror.Wrap(err, "获取房型列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sHotelService) RoomTypeEdit(ctx context.Context, in *input_hotel.PmsRoomTypeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {

			var (
				BedTypeInfo []*entity.PmsBedType
			)

			for _, BedType := range in.BedType {
				if !g.IsEmpty(BedType.BedTypeName) {
					BedTypeInfo = append(BedTypeInfo, &entity.PmsBedType{
						RoomTypeId:  in.Id,
						BedTypeName: BedType.BedTypeName,
						BedWidth:    BedType.BedWidth,
						BedNum:      BedType.BedNum,
					})
				}

			}
			if !g.IsEmpty(BedTypeInfo) {

				if _, err = dao.PmsBedType.Ctx(ctx).Where(dao.PmsBedType.Columns().RoomTypeId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理房型床型旧数据失败，请稍后重试！")
					return
				}

				if _, err = dao.PmsBedType.Ctx(ctx).OmitEmptyData().Insert(BedTypeInfo); err != nil {
					return err
				}
			}

			var (
				Object         gdb.Record
				PmsRoomTypeDao *input_language.LoadLanguage
				LanguageStruct input_language.LanguageModel
			)
			Uuid := guid.S([]byte("name"))
			if Object, err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Id, in.Id).One(); err != nil {
				return
			}

			if !g.IsEmpty(Object[gstr.ToLower("Name")]) {
				Uuid = Object["name"].String()
			}
			PmsRoomTypeDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.PmsRoomType.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			LanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, PmsRoomTypeDao); err != nil {
				return
			}

			in.Name = Uuid
			if _, err = dao.PmsRoomType.Ctx(ctx).
				Fields(input_hotel.PmsRoomTypeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改房型失败，请稍后重试！")
			}
			return
		}

		var (
			lastInsertId int64
		)
		if lastInsertId, err = dao.PmsRoomType.Ctx(ctx).
			Fields(input_hotel.PmsRoomTypeInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增房型失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "收藏失败，请稍后重试！")
			return
		}
		var (
			BedTypeInfo []*entity.PmsBedType
		)

		for _, BedType := range in.BedType {
			BedTypeInfo = append(BedTypeInfo, &entity.PmsBedType{
				RoomTypeId:  int(lastInsertId),
				BedTypeName: BedType.BedTypeName,
				BedWidth:    BedType.BedWidth,
				BedNum:      BedType.BedNum,
			})
		}
		if !g.IsEmpty(BedTypeInfo) {
			if _, err = dao.PmsBedType.Ctx(ctx).OmitEmptyData().Insert(BedTypeInfo); err != nil {
				return err
			}
		}

		return
	})
}

func (s *sHotelService) RoomTypeDelete(ctx context.Context, in *input_hotel.PmsRoomTypeDeleteInp) (err error) {

	if _, err = dao.PmsRoomType.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除房型失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) RoomTypeView(ctx context.Context, in *input_hotel.PmsRoomTypeViewInp) (res *input_hotel.PmsRoomTypeViewModel, err error) {
	if err = dao.PmsRoomType.Ctx(ctx).WithAll().Where(g.MapStrAny{
		dao.PmsRoomType.Columns().Id:  in.Id,
		dao.PmsRoomType.Columns().Uid: in.Uid,
	}).OmitEmptyWhere().Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取房型信息，请稍后重试！")
		return
	}

	return
}

func (s *sHotelService) RoomTypeGetUidsByPuid(ctx context.Context, puid string) (uids []string, err error) {
	columns, err := dao.PmsRoomType.Ctx(ctx).
		Fields("uid").
		Where("puid", puid).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取uid失败！")
		return
	}

	uids = g.NewVar(columns).Strings()
	return
}

func (s *sHotelService) RoomTypeStatus(ctx context.Context, in *input_hotel.PmsRoomTypeIsShowInp) (err error) {
	if _, err = dao.PmsRoomType.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsRoomType.Columns().IsShow: in.IsShow,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "状态更新失败，请稍后重试！")
		return
	}
	return
}
