package logic_employee

import (
	"APT/internal/dao"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_employee"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponRecordList 获取活动券领取记录列表
func (s *sEmployeeActivity) CouponRecordList(ctx context.Context, in *input_employee.EmployeeActivityCouponRecordListInp) (list []*input_employee.EmployeeActivityCouponRecordListModel, totalCount int, err error) {
	if err = in.Filter(ctx); err != nil {
		return
	}

	mod := g.Model(dao.ThMemberCoupon.Table()+" mc").Ctx(ctx).
		LeftJoin(dao.ThCoupon.Table()+" tc", "tc.id = mc.coupon_id").
		LeftJoin(dao.Employee.Table()+" emp", "emp.id = mc.employee_id").
		LeftJoin(dao.ThMchStore.Table()+" store", "store.id = mc.verify_store_id").
		Where("mc.source", 3).
		Where("mc.activity_id", in.ActivityId).
		Fields(
			"mc.id, mc.coupon_no, mc.coupon_id, mc.member_id, mc.employee_id",
			"mc.state, mc.source, mc.start_time, mc.end_time",
			"mc.verify_time, mc.verify_mch_id, mc.create_at",
			"tc.coupon_name, tc.need_reservation",
			"emp.name as employee_name",
			"store.store_name",
		)

	// 发放类型筛选：1-自动领取(need_reservation=0) 2-批量发放(need_reservation=1)
	if in.IssueType == 1 {
		mod = mod.Where("tc.need_reservation", 0)
	} else if in.IssueType == 2 {
		mod = mod.Where("tc.need_reservation", 1)
	}

	// 券状态筛选
	if in.State > 0 {
		mod = mod.Where("mc.state", in.State)
	}

	// 关键词搜索
	if in.Keyword != "" {
		mod = mod.Where(
			g.Model().Builder().
				WhereLike("emp.name", "%"+in.Keyword+"%").
				WhereOrLike("mc.coupon_no", "%"+in.Keyword+"%"),
		)
	}

	mod = mod.Page(in.Page, in.PerPage).OrderDesc("mc.create_at")

	// Hook 自动将 coupon_name、store_name 的多语言UUID替换为对应语言文本
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "查询券领取记录列表失败")
		return
	}

	// 填充 issueType
	for _, item := range list {
		if item.NeedReservation == 1 {
			item.IssueType = 2 // 批量发放
		} else {
			item.IssueType = 1 // 自动领取
		}
	}

	return
}
