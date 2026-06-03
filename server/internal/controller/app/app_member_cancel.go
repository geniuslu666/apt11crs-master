package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/app/member"
)

func (c *ControllerMember) Cancel(ctx context.Context, req *member.CancelReq) (res *member.CancelRes, err error) {
	var (
		PmsMember       *entity.PmsMember
		MemberUser      = contexts.GetMemberUser(ctx)
		PmsMemberCancel *entity.PmsMemberCancel
		InsertData      *entity.PmsMemberCancel
		Config          *input_basics.GetConfigModel
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 获取用户信息
		if err = dao.PmsMember.Ctx(ctx).WherePri(MemberUser.Id).Scan(&PmsMember); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		// 根据会员ID找到会员注销记录
		if err = dao.PmsMemberCancel.Ctx(ctx).TX(tx).
			Where(g.Map{
				dao.PmsMemberCancel.Columns().MemberId: MemberUser.Id,
			}).
			WhereIn(dao.PmsMemberCancel.Columns().AuditStatus, []int{1, 2}).OrderDesc(dao.PmsMemberCancel.Columns().Id).
			Scan(&PmsMemberCancel); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		if !g.IsEmpty(PmsMemberCancel) {
			// 您已申请注销会员，请勿重复申请
			err = gerror.New(gi18n.T(ctx, "has_applied_to_cancel_membership"))
			return
		}

		// 获取会员注销设置
		if Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
			Group: "membercancelsetting",
		}); err != nil {
			// 未找到注销设置
			err = gerror.Wrap(err, gi18n.T(ctx, "not_found_unregister_settings"))
			return
		}

		IsEnableCancel := Config.List["isEnableCancel"]
		IsCancelAudit := Config.List["isCancelAudit"]
		if IsEnableCancel != 1 {
			// 不可注销会员账号
			err = gerror.New(gi18n.T(ctx, "cannot_cancel_member_account"))
			return
		}

		if IsCancelAudit == 1 {
			InsertData = &entity.PmsMemberCancel{
				MemberId:     gvar.New(MemberUser.Id).Uint(),
				MemberNo:     MemberUser.MemberNo,
				Phone:        MemberUser.Phone,
				PhoneArea:    MemberUser.PhoneArea,
				Mail:         MemberUser.Mail,
				CancelReason: req.CancelReason,
			}
		} else {
			// 无需审核
			InsertData = &entity.PmsMemberCancel{
				MemberId:     gvar.New(MemberUser.Id).Uint(),
				MemberNo:     MemberUser.MemberNo,
				Phone:        MemberUser.Phone,
				PhoneArea:    MemberUser.PhoneArea,
				Mail:         MemberUser.Mail,
				CancelReason: req.CancelReason,
				AuditStatus:  2,
				AuditTime:    gtime.Now(),
				AuditReason:  "用户申请注销自动审核成功",
				OperatorId:   0,
			}

			// 若会员绑定了渠道或员工，则解除渠道或员工的绑定关系
			if PmsMember.RebateMode == "CHANNEL" {
				// 解绑渠道
				if _, err = dao.PmsMember.Ctx(ctx).WherePri(MemberUser.Id).Data(g.Map{
					dao.PmsMember.Columns().RebateMode: "MEMBER",
					dao.PmsMember.Columns().ChannelId:  nil,
				}).Update(); err != nil {
					err = gerror.Wrap(err, "解绑渠道失败，请稍后重试！")
				}

			}
			if PmsMember.RebateMode == "STAFF" {
				// 解绑员工
				if _, err = dao.PmsMember.Ctx(ctx).WherePri(MemberUser.Id).Data(g.Map{
					dao.PmsMember.Columns().RebateMode: "MEMBER",
					dao.PmsMember.Columns().StaffId:    nil,
				}).Update(); err != nil {
					err = gerror.Wrap(err, "解绑员工失败，请稍后重试！")
				}
			}

			// 查询是否绑定按摩服务商、按摩技师、接送机司机、员工活动中的员工，进行解绑
			// 解绑按摩服务商
			if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().MemberId, MemberUser.Id).Data(g.Map{
				dao.SpaIsp.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑按摩服务商失败，请稍后重试！")
				return
			}
			// 解绑按摩技师
			if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().MemberId, MemberUser.Id).Data(g.Map{
				dao.SpaTechnician.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑按摩技师失败，请稍后重试！")
				return
			}
			// 解绑接送机司机
			if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().MemberId, MemberUser.Id).Data(g.Map{
				dao.CarDriver.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑接送机司机失败，请稍后重试！")
				return
			}
			// 解绑员工活动中的员工
			if _, err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, MemberUser.Id).Data(g.Map{
				dao.Employee.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑员工活动员工失败，请稍后重试！")
				return
			}

			// 删除会员信息
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).WherePri(MemberUser.Id).Delete(); err != nil {
				// 注销会员失败，请稍后重试
				err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_cancel_membership"))
				return
			}
		}

		// 写入会员注销记录
		if _, err = dao.PmsMemberCancel.Ctx(ctx).TX(tx).Data(InsertData).OmitEmptyData().Insert(); err != nil {
			// 申请注销失败，请稍后重试
			err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_apply_for_cancellation"))
			return
		}
		return
	}); err != nil {
		return
	}
	return
}
