package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarOrderLog struct{}

func NewCarOrderLog() *sCarOrderLog {
	return &sCarOrderLog{}
}

func init() {
	service.RegisterCarOrderLog(NewCarOrderLog())
}

func (s *sCarOrderLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarOrderLog.Ctx(ctx), option...)
}

func (s *sCarOrderLog) List(ctx context.Context, in *input_car.CarOrderLogListInp) (list []*input_car.CarOrderLogListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.CarOrderLog.Table(), input_car.CarOrderLogListModel{})
	//mod = mod.Fields(hgorm.JoinFields(ctx, input_car.CarOrderListModel{}, &dao.CarOrder, "carOrder"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsBalanceChangeListModel{}, &dao.AdminMember, "adminMember"))

	//mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.CarOrderLog.Columns().OrderId, "=", dao.CarOrder.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.CarOrderLog.Columns().OperateId, "=", dao.AdminMember.Columns().Id)

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.CarOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.OrderStatus) {
		mod = mod.WhereLike(dao.CarOrderLog.Columns().OrderStatus, in.OrderStatus)
	}

	if !g.IsEmpty(in.ActionWay) {
		mod = mod.WhereLike(dao.CarOrderLog.Columns().ActionWay, in.ActionWay)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CarOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.CarOrderLog.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取接送机订单日志列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			var AdminMemberInfo *entity.AdminMember
			if err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
				return
			}
			v.OperateName = AdminMemberInfo.Username
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			v.OperateName = PmsMemberInfo.FullName
			v.OperatePhone = PmsMemberInfo.Phone
		}

		if v.OperateType == "DRIVER" {
			var CarDriverInfo *entity.CarDriver
			if err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, v.OperateId).Scan(&CarDriverInfo); err != nil {
				return
			}
			v.OperateName = CarDriverInfo.Name
			v.OperatePhone = CarDriverInfo.Phone
		}
	}
	return
}
