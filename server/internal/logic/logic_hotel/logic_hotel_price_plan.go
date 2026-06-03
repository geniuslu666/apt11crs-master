package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"strings"
)

func (s *sHotelService) HotelPricePlanList(ctx context.Context, ipt *input_hotel.SearchHotelPricePlanInp) (list []*input_hotel.SearchHotelPricePlanRes, count int, err error) {
	mod := dao.PmsPricePlan.Ctx(ctx).WithAll()

	mod = mod.Fields(input_hotel.SearchHotelPricePlanRes{})

	if ipt.PlanName != "" {
		mod = mod.WhereLike(dao.PmsPricePlan.Columns().PlanName, "%"+ipt.PlanName+"%")
	}

	if !g.IsEmpty(ipt.RoomTypeId) {
		mod = mod.Where(dao.PmsPricePlan.Columns().RoomTypeId, ipt.RoomTypeId)
	}

	if !g.IsEmpty(ipt.PropertyId) {
		mod = mod.Where(dao.PmsPricePlan.Columns().PropertyId, ipt.PropertyId)
	}

	if !g.IsEmpty(ipt.PricePlanStatus) {
		mod = mod.Where(dao.PmsPricePlan.Columns().PricePlanStatus, ipt.PricePlanStatus)
	}

	if !g.IsEmpty(ipt.Type) {
		mod = mod.Unscoped().WhereNotNull(dao.PmsPricePlan.Columns().DeletedAt)
	}

	mod = mod.Hook(hook.PmsFindLanguageValueHook).Page(ipt.Page, ipt.PerPage)

	mod = mod.OrderDesc(dao.PmsRoomUnit.Columns().Id)

	if err = mod.ScanAndCount(&list, &count, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}

	for k, v := range list {
		if !g.IsEmpty(v.MemberLevelId) {
			if err = dao.PmsMemberLevel.Ctx(ctx).WhereIn(dao.PmsMemberLevel.Columns().Id, strings.Split(v.MemberLevelId, ",")).Scan(&list[k].MemberLevelDetail); err != nil {
				err = gerror.Wrap(err, "获取会员等级失败，请稍后重试！")
				return
			}
		}
		if !g.IsEmpty(v.MemberGroupId) {
			if err = dao.PmsMemberGroup.Ctx(ctx).WhereIn(dao.PmsMemberGroup.Columns().Id, strings.Split(v.MemberGroupId, ",")).Scan(&list[k].MemberGroupDetail); err != nil {
				err = gerror.Wrap(err, "获取会员等级失败，请稍后重试！")
				return
			}
		}
	}

	return
}

func (s *sHotelService) HotelPricePlanEdit(ctx context.Context, ipt *input_hotel.EditHotelPricePlanInp) (err error) {
	var (
		PricePlan *entity.PmsPricePlan
	)
	ipt.PlanShowName = ""
	if !g.IsEmpty(ipt.PlanShowNameLanguage) {
		// 查询数据
		if err = dao.PmsPricePlan.Ctx(ctx).WherePri(ipt.Id).Scan(&PricePlan); err != nil {
			return
		}
		if !g.IsEmpty(PricePlan) && !g.IsEmpty(PricePlan.PlanShowName) {
			ipt.PlanShowName = PricePlan.PlanShowName
		} else {
			ipt.PlanShowName = guid.S()
		}
		// 更新名字多语言数据
		if err = service.BasicsLanguage().Sync(ctx, ipt.PlanShowNameLanguage, &input_language.LoadLanguage{
			Uuid: ipt.PlanShowName,
			Tag:  dao.PmsPricePlan.Table(),
			Type: "table",
			Key:  dao.PmsPricePlan.Columns().PlanShowName,
		}); err != nil {
			return
		}
	}
	if ipt.Id > 0 {
		if _, err = dao.PmsPricePlan.Ctx(ctx).
			WherePri(ipt.Id).Data(ipt).Update(); err != nil {
			err = gerror.Wrap(err, "修改价格plan失败，请稍后重试！")
		}
		return
	} else {
		if _, err = dao.PmsPricePlan.Ctx(ctx).
			Data(ipt).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增价格plan失败，请稍后重试！")
		}
	}
	return
}

func (s *sHotelService) HotelPricePlanDelete(ctx context.Context, ipt *input_hotel.DeleteHotelPricePlanInp) (err error) {
	if _, err = dao.PmsPricePlan.Ctx(ctx).
		WherePri(ipt.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除价格plan失败，请稍后重试！")
	}
	return
}

func (s *sHotelService) HotelPricePlanRecycle(ctx context.Context, in *input_hotel.DeleteHotelPricePlanInp) (err error) {
	if _, err = dao.PmsPricePlan.Ctx(ctx).Unscoped().WherePri(in.Id).Data(g.Map{
		dao.PmsPricePlan.Columns().DeletedAt: nil,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "恢复价格计划失败，请稍后重试！")
		return
	}
	return
}

func (s *sHotelService) HotelPricePlanGet(ctx context.Context, ipt *input_hotel.FindOneHotelPricePlanInp) (res *input_hotel.FindOneHotelPricePlanRes, err error) {
	if err = dao.PmsPricePlan.Ctx(ctx).
		WithAll().WherePri(ipt.Id).Scan(&res); err != nil {
		return
	}
	if !g.IsEmpty(res.RoomTypeId) {
		if err = dao.PmsRoomType.Ctx(ctx).Where(dao.PmsRoomType.Columns().Uid, res.RoomTypeId).Hook(hook.PmsFindLanguageValueHook).Scan(&res.RoomTypeDetail); err != nil {
			err = gerror.Wrap(err, "获取房型失败，请稍后重试！")
			return
		}
	}
	if !g.IsEmpty(res.MemberLevelId) {
		if err = dao.PmsMemberLevel.Ctx(ctx).WhereIn(dao.PmsMemberLevel.Columns().Id, strings.Split(res.MemberLevelId, ",")).Scan(&res.MemberLevelDetail); err != nil {
			err = gerror.Wrap(err, "获取会员等级失败，请稍后重试！")
			return
		}
	}
	if !g.IsEmpty(res.MemberGroupId) {
		if err = dao.PmsMemberGroup.Ctx(ctx).WhereIn(dao.PmsMemberGroup.Columns().Id, strings.Split(res.MemberGroupId, ",")).Scan(&res.MemberGroupDetail); err != nil {
			err = gerror.Wrap(err, "获取会员组失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sHotelService) HotelPricePlanStatus(ctx context.Context, in *input_hotel.PricePlanStatusInp) (err error) {
	if _, err = dao.PmsPricePlan.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsPricePlan.Columns().PricePlanStatus: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
		return
	}
	return
}
