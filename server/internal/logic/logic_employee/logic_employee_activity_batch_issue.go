package logic_employee

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_employee"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BatchIssueCoupons 批量发放预约券给符合条件的员工
func (s *sEmployeeActivity) BatchIssueCoupons(ctx context.Context, in *input_employee.EmployeeActivityBatchIssueCouponsInp) (res *input_employee.EmployeeActivityBatchIssueCouponsModel, err error) {
	// 1. 查询活动信息
	var activityInfo *entity.EmployeeActivity
	if err = dao.EmployeeActivity.Ctx(ctx).
		Where(dao.EmployeeActivity.Columns().Id, in.ActivityId).
		Scan(&activityInfo); err != nil {
		err = gerror.Wrap(err, "查询活动失败")
		return
	}
	if g.IsEmpty(activityInfo) {
		err = gerror.New("活动不存在")
		return
	}
	if activityInfo.IsEnabled != 1 {
		err = gerror.New("活动已禁用")
		return
	}
	if activityInfo.Status != 2 {
		err = gerror.New("活动未在进行中")
		return
	}
	// 2. 查询活动中需要预约的礼品券
	type activityCouponWithName struct {
		entity.EmployeeActivityCoupon
		CouponName string `json:"couponName" orm:"coupon_name"`
	}
	var activityCoupons []*activityCouponWithName
	if err = g.Model(dao.EmployeeActivityCoupon.Table()+" eac").Ctx(ctx).
		LeftJoin(dao.ThCoupon.Table()+" tc", "tc.id = eac.coupon_id").
		Where("eac.activity_id", in.ActivityId).
		Where("tc.need_reservation", 1).
		Where("tc.status", 1).
		Fields("eac.*, tc.coupon_name").
		Scan(&activityCoupons); err != nil {
		err = gerror.Wrap(err, "查询活动礼品券失败")
		return
	}
	if len(activityCoupons) == 0 {
		err = gerror.New("该活动没有需要预约的券")
		return
	}

	// 3. 查询符合条件的员工列表
	var employees []*entity.Employee
	switch activityInfo.RestrictionType {
	case 1:
		// 不限制 - 查询所有启用的员工
		if err = dao.Employee.Ctx(ctx).
			Where(dao.Employee.Columns().Status, 1).
			WhereGT(dao.Employee.Columns().MemberId, 0).
			Scan(&employees); err != nil {
			err = gerror.Wrap(err, "查询员工列表失败")
			return
		}
	case 2:
		// 限制部门
		departmentIdsArr, arrErr := dao.EmployeeActivityDepartment.Ctx(ctx).
			Where(dao.EmployeeActivityDepartment.Columns().ActivityId, in.ActivityId).
			Fields(dao.EmployeeActivityDepartment.Columns().DepartmentId).
			Array()
		if arrErr != nil {
			err = gerror.Wrap(arrErr, "查询活动部门失败")
			return
		}
		departmentIds := make([]int, 0, len(departmentIdsArr))
		for _, v := range departmentIdsArr {
			departmentIds = append(departmentIds, v.Int())
		}
		if len(departmentIds) == 0 {
			err = gerror.New("该活动未配置限制部门")
			return
		}
		if err = dao.Employee.Ctx(ctx).
			Where(dao.Employee.Columns().Status, 1).
			WhereGT(dao.Employee.Columns().MemberId, 0).
			WhereIn(dao.Employee.Columns().DepartmentId, departmentIds).
			Scan(&employees); err != nil {
			err = gerror.Wrap(err, "查询员工列表失败")
			return
		}
	case 3:
		// 限制员工
		employeeIdsArr, arrErr := dao.EmployeeActivityEmployee.Ctx(ctx).
			Where(dao.EmployeeActivityEmployee.Columns().ActivityId, in.ActivityId).
			Fields(dao.EmployeeActivityEmployee.Columns().EmployeeId).
			Array()
		if arrErr != nil {
			err = gerror.Wrap(arrErr, "查询活动员工失败")
			return
		}
		employeeIds := make([]int, 0, len(employeeIdsArr))
		for _, v := range employeeIdsArr {
			employeeIds = append(employeeIds, v.Int())
		}
		if len(employeeIds) == 0 {
			err = gerror.New("该活动未配置限制员工")
			return
		}
		if err = dao.Employee.Ctx(ctx).
			Where(dao.Employee.Columns().Status, 1).
			WhereGT(dao.Employee.Columns().MemberId, 0).
			WhereIn(dao.Employee.Columns().Id, employeeIds).
			Scan(&employees); err != nil {
			err = gerror.Wrap(err, "查询员工列表失败")
			return
		}
	}

	if len(employees) == 0 {
		err = gerror.New("没有符合条件的员工")
		return
	}

	// 4. 并发发放券
	res = &input_employee.EmployeeActivityBatchIssueCouponsModel{
		TotalEmployees:  len(employees),
		CouponDetails:   make([]*input_employee.EmployeeActivityBatchIssueCouponDetail, 0),
		FailedEmployees: make([]*input_employee.EmployeeActivityBatchIssueFailedEmployee, 0),
	}

	// 解析券名称（coupon_name 是多语言 UUID，这里用 couponId 做 key 统计）
	couponDetailMap := make(map[uint64]*input_employee.EmployeeActivityBatchIssueCouponDetail)
	for _, ac := range activityCoupons {
		couponDetailMap[ac.CouponId] = &input_employee.EmployeeActivityBatchIssueCouponDetail{
			CouponId:   int(ac.CouponId),
			CouponName: ac.CouponName,
		}
	}

	var (
		issuedCount int64
		failedCount int64
		mu          sync.Mutex
		wg          sync.WaitGroup
		sem         = make(chan struct{}, 50) // 并发限制50
	)

	for _, emp := range employees {
		for _, ac := range activityCoupons {
			wg.Add(1)
			sem <- struct{}{}
			go func(employee *entity.Employee, coupon *activityCouponWithName) {
				defer wg.Done()
				defer func() { <-sem }()

				// 每个 goroutine 使用独立的数据库上下文，避免并发共享连接
				taskCtx := gdb.WithDB(context.Background(), g.DB())

				// 检查是否已领取
				count, countErr := dao.ThMemberCoupon.Ctx(taskCtx).
					Where(dao.ThMemberCoupon.Columns().Source, 3).
					Where(dao.ThMemberCoupon.Columns().EmployeeId, employee.Id).
					Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId).
					Where(dao.ThMemberCoupon.Columns().CouponId, coupon.CouponId).
					Count()
				if countErr != nil {
					atomic.AddInt64(&failedCount, 1)
					mu.Lock()
					res.FailedEmployees = append(res.FailedEmployees, &input_employee.EmployeeActivityBatchIssueFailedEmployee{
						EmployeeId:   int(employee.Id),
						EmployeeName: employee.Name,
						Reason:       fmt.Sprintf("查询已领取记录失败: %v", countErr),
					})
					mu.Unlock()
					return
				}

				// 已领取则跳过
				if coupon.AvailableQuantity > 0 && count >= coupon.AvailableQuantity {
					return
				}

				// 计算需要发放的数量
				needIssue := 1
				if coupon.AvailableQuantity > 0 {
					needIssue = coupon.AvailableQuantity - count
				}

				// 发放券
				startTime := gtime.Now().String()
				if activityInfo.CouponValidity == 1 && activityInfo.StartTime != nil {
					startTime = activityInfo.StartTime.String()
				}

				for i := 0; i < needIssue; i++ {
					_, _, sendErr := service.ThCoupon().SendMemberCouponGetId(taskCtx, &input_th.ThSendMemberCouponInp{
						CouponId:   int(coupon.CouponId),
						MemberId:   int(employee.MemberId),
						StartTime:  startTime,
						ActivityId: activityInfo.Id,
						EmployeeId: employee.Id,
					}, 3) // source=3 员工福利

					if sendErr != nil {
						atomic.AddInt64(&failedCount, 1)
						mu.Lock()
						res.FailedEmployees = append(res.FailedEmployees, &input_employee.EmployeeActivityBatchIssueFailedEmployee{
							EmployeeId:   int(employee.Id),
							EmployeeName: employee.Name,
							Reason:       fmt.Sprintf("发放券[%d]第%d张失败: %v", coupon.CouponId, i+1, sendErr),
						})
						mu.Unlock()
						continue
					}

					atomic.AddInt64(&issuedCount, 1)
					mu.Lock()
					if detail, ok := couponDetailMap[coupon.CouponId]; ok {
						detail.IssuedCount++
					}
					mu.Unlock()
				}
			}(emp, ac)
		}
	}

	wg.Wait()

	// 所有员工都已领取过，直接返回提示
	if issuedCount == 0 && failedCount == 0 {
		err = gerror.New("所有员工已领取，不可重复发放")
		res = nil
		return
	}

	res.IssuedCount = int(issuedCount)
	res.FailedCount = int(failedCount)
	for _, detail := range couponDetailMap {
		res.CouponDetails = append(res.CouponDetails, detail)
	}

	return
}
