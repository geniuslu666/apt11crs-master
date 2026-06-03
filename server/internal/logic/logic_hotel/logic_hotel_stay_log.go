package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sHotelService) AppStayLogModel(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsAppStayLog.Ctx(ctx), option...)
}

func (s *sHotelService) AppStayLogList(ctx context.Context, in *input_hotel.AppStayLogListInp) (list []*input_hotel.AppStayLogListModel, totalCount int, err error) {
	mod := s.AppStayLogModel(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.PmsAppStayLog.Table(), input_hotel.AppStayLogListModel{})
	//mod = mod.Fields(hgorm.JoinFields(ctx, input_hotel.PmsAppReservationListModel{}, &dao.PmsAppStay, "appStay"))

	mod = mod.LeftJoinOnFields(dao.PmsAppStay.Table(), dao.PmsAppStayLog.Columns().OrderId, "=", dao.PmsAppStay.Columns().Id)

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.PmsAppStay.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.ActionWay) {
		mod = mod.WhereLike(dao.PmsAppStayLog.Columns().ActionWay, in.ActionWay)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsAppStayLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsAppStayLog.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取酒店订单日志列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			if v.OperateId > 0 {
				var AdminMemberInfo *entity.AdminMember
				if err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
					return
				}
				v.OperateName = AdminMemberInfo.Username
			} else {
				v.OperateName = "后台"
			}
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			v.OperateName = PmsMemberInfo.FullName
			v.OperatePhone = PmsMemberInfo.Phone
		}

		if v.OperateType == "TECHNICIAN" {
			var SpaTechnicianInfo *entity.SpaTechnician
			if err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, v.OperateId).Scan(&SpaTechnicianInfo); err != nil {
				return
			}
			v.OperateName = SpaTechnicianInfo.Name
			v.OperatePhone = SpaTechnicianInfo.Phone
		}
	}
	return
}
