package logic_app_member

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberCancelList 会员注销列表
func (s *sAppMember) MemberCancelList(ctx context.Context, in *input_app_member.PmsMemberCancelListInp) (list []*input_app_member.PmsMemberCancelListModel, totalCount int, err error) {
	mod := dao.PmsMemberCancel.Ctx(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.PmsMemberCancel.Table(), input_app_member.PmsMemberCancelListModel{})
	//mod = mod.Fields(input_app_member.PmsMemberCancelListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_app_member.PmsBalanceChangeListModel{}, &dao.AdminMember, "adminMember"))

	mod = mod.LeftJoinOnFields(dao.AdminMember.Table(), dao.PmsMemberCancel.Columns().OperatorId, "=", dao.AdminMember.Columns().Id)

	if in.Id > 0 {
		mod = mod.Where(dao.PmsMemberCancel.Columns().Id, in.Id)
	}

	if in.MemberNo != "" {
		mod = mod.WhereLike(dao.PmsMemberCancel.Columns().MemberNo, "%"+in.MemberNo+"%")
	}

	if in.Phone != "" {
		mod = mod.WhereLike(dao.PmsMemberCancel.Columns().Phone, "%"+in.Phone+"%")
	}

	if in.Mail != "" {
		mod = mod.WhereLike(dao.PmsMemberCancel.Columns().Mail, "%"+in.Mail+"%")
	}

	if !g.IsEmpty(in.AuditStatus) && in.AuditStatus != "ALL" {
		mod = mod.Where(dao.PmsMemberCancel.Columns().AuditStatus, in.AuditStatus)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsMemberCancel.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取会员列表注销失败，请稍后重试！")
		return
	}
	return
}

// MemberCancelAgree 同意注销
func (s *sAppMember) MemberCancelAgree(ctx context.Context, in *input_app_member.PmsMemberCancelAgreeInp) (err error) {

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var models *entity.PmsMemberCancel
		if err = dao.PmsMemberCancel.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if models == nil {
			err = gerror.New("注销信息不存在或已被删除")
			return
		}

		if models.AuditStatus != 1 {
			err = gerror.New("注销申请状态不正确")
			return
		}

		// 若会员绑定了渠道或员工，则解除渠道或员工的绑定关系
		if _, err = dao.PmsMember.Ctx(ctx).WherePri(models.MemberId).Data(g.Map{
			dao.PmsMember.Columns().RebateMode: "MEMBER",
			dao.PmsMember.Columns().ChannelId:  nil,
			dao.PmsMember.Columns().StaffId:    nil,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "解绑渠道/员工失败，请稍后重试！")
		}

		// 查询是否绑定按摩服务商、按摩技师、接送机司机、员工活动中的员工，进行解绑
		// 解绑按摩服务商
		if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().MemberId, in.Id).Data(g.Map{
			dao.SpaIsp.Columns().MemberId: 0,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "解绑按摩服务商失败，请稍后重试！")
			return
		}
		// 解绑按摩技师
		if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().MemberId, in.Id).Data(g.Map{
			dao.SpaTechnician.Columns().MemberId: 0,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "解绑按摩技师失败，请稍后重试！")
			return
		}
		// 解绑接送机司机
		if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().MemberId, in.Id).Data(g.Map{
			dao.CarDriver.Columns().MemberId: 0,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "解绑接送机司机失败，请稍后重试！")
			return
		}
		// 解绑员工活动中的员工
		if _, err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, in.Id).Data(g.Map{
			dao.Employee.Columns().MemberId: 0,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "解绑员工活动员工失败，请稍后重试！")
			return
		}

		// 删除会员
		if _, err = dao.PmsMember.Ctx(ctx).WherePri(models.MemberId).Delete(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 更新注销申请
		if _, err = dao.PmsMemberCancel.Ctx(ctx).
			WherePri(in.Id).Data(input_app_member.PmsMemberCancelAgreeFields{
			AuditStatus: 2,
			AuditReason: in.AuditReason,
			OperatorId:  int(contexts.GetUserId(ctx)),
			AuditTime:   gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		return
	})

}

// MemberCancelDisagree 驳回注销
func (s *sAppMember) MemberCancelDisagree(ctx context.Context, in *input_app_member.PmsMemberCancelDisagreeInp) (err error) {

	if _, err = dao.PmsMemberCancel.Ctx(ctx).
		WherePri(in.Id).Data(input_app_member.PmsMemberCancelAgreeFields{
		AuditStatus: 3,
		AuditReason: in.AuditReason,
		OperatorId:  int(contexts.GetUserId(ctx)),
		AuditTime:   gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
	}
	return

}
