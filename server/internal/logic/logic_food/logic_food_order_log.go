package logic_food

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFoodOrderLog struct{}

func NewFoodOrderLog() *sFoodOrderLog {
	return &sFoodOrderLog{}
}

func init() {
	service.RegisterFoodOrderLog(NewFoodOrderLog())
}

func (s *sFoodOrderLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodOrderLog.Ctx(ctx), option...)
}

func (s *sFoodOrderLog) List(ctx context.Context, in *input_food.FoodOrderLogListInp) (list []*input_food.FoodOrderLogListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.FoodOrderLog.Table(), input_food.FoodOrderLogListModel{})
	//mod = mod.Fields(hgorm.JoinFields(ctx, input_food.FoodOrderListModel{}, &dao.FoodOrder, "carOrder"))
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsBalanceChangeListModel{}, &dao.AdminMember, "adminMember"))

	//mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.FoodOrderLog.Columns().OrderId, "=", dao.FoodOrder.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.FoodOrderLog.Columns().OperateId, "=", dao.AdminMember.Columns().Id)

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.FoodOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.ActionWay) {
		mod = mod.WhereLike(dao.FoodOrderLog.Columns().ActionWay, in.ActionWay)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CarOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.FoodOrderLog.Columns().Id)

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
	}
	return
}
