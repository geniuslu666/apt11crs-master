// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_employee"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IEmployee interface {
		// Model 员工管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取员工管理列表
		List(ctx context.Context, in *input_employee.EmployeeListInp) (list []*input_employee.EmployeeListModel, totalCount int, err error)
		// Edit 修改/新增员工
		Edit(ctx context.Context, in *input_employee.EmployeeEditInp) (err error)
		// Delete 删除员工管理
		Delete(ctx context.Context, in *input_employee.EmployeeDeleteInp) (err error)
		// View 获取员工管理指定信息
		View(ctx context.Context, in *input_employee.EmployeeViewInp) (res *input_employee.EmployeeViewModel, err error)
		// Status 更新员工状态
		Status(ctx context.Context, in *input_employee.EmployeeStatusInp) (err error)
		// Bind 绑定用户
		Bind(ctx context.Context, in *input_employee.EmployeeBindInp) (err error)
		// Unbind 解绑用户
		Unbind(ctx context.Context, in *input_employee.EmployeeUnbindInp) (err error)
		// CheckAuth 验证员工
		CheckAuth(ctx context.Context) (auth bool, err error)
	}
	IEmployeeActivity interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_employee.EmployeeActivityListInp) (list []*input_employee.EmployeeActivityListModel, totalCount int, err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		Edit(ctx context.Context, in *input_employee.EmployeeActivityEditInp) (err error)
		Delete(ctx context.Context, in *input_employee.EmployeeActivityDeleteInp) (err error)
		View(ctx context.Context, in *input_employee.EmployeeActivityViewInp) (res *input_employee.EmployeeActivityViewModel, err error)
		// Status 更新员工活动状态
		Status(ctx context.Context, in *input_employee.EmployeeActivityStatusInp) (err error)
		ActivityEffect(ctx context.Context, ActivityId uint64) (err error)
		AppList(ctx context.Context, in *input_employee.EmployeeActivityAppListInp) (list []*input_employee.EmployeeActivityAppListModel, totalCount int, err error)
		AppView(ctx context.Context, in *input_employee.EmployeeActivityAppViewInp) (res *input_employee.EmployeeActivityAppViewModel, err error)
		ReceiveActivityCoupon(ctx context.Context, in *input_employee.EmployeeActivityReceiveCouponInp) (err error)
		BatchIssueCoupons(ctx context.Context, in *input_employee.EmployeeActivityBatchIssueCouponsInp) (res *input_employee.EmployeeActivityBatchIssueCouponsModel, err error)
		CouponRecordList(ctx context.Context, in *input_employee.EmployeeActivityCouponRecordListInp) (list []*input_employee.EmployeeActivityCouponRecordListModel, totalCount int, err error)
	}
	IEmployeeDepartment interface {
		// Model 员工部门ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取员工部门列表
		List(ctx context.Context, in *input_employee.EmployeeDepartmentListInp) (res *input_employee.EmployeeDepartmentListModel, err error)
		// Edit 修改/新增员工部门
		Edit(ctx context.Context, in *input_employee.EmployeeDepartmentEditInp) (err error)
		// Delete 删除员工部门
		Delete(ctx context.Context, in *input_employee.EmployeeDepartmentDeleteInp) (err error)
		// BatchDelete 批量删除员工部门
		BatchDelete(ctx context.Context, in *input_employee.EmployeeDepartmentBatchDeleteInp) (err error)
		// View 获取员工部门详情
		View(ctx context.Context, in *input_employee.EmployeeDepartmentViewInp) (res *input_employee.EmployeeDepartmentViewModel, err error)
		// Switch 更新员工部门状态
		Switch(ctx context.Context, in *input_employee.EmployeeDepartmentSwitchInp) (err error)
		// TreeOption 获取员工部门树选项
		TreeOption(ctx context.Context) (res *input_employee.EmployeeDepartmentOptionModel, err error)
		// MaxSort 获取最大排序
		MaxSort(ctx context.Context, in *input_employee.EmployeeDepartmentMaxSortInp) (res *input_employee.EmployeeDepartmentMaxSortModel, err error)
		// GenTree 生成关系树
		GenTree(ctx context.Context, parentId uint64) (level int, newPath string, err error)
	}
)

var (
	localEmployee           IEmployee
	localEmployeeActivity   IEmployeeActivity
	localEmployeeDepartment IEmployeeDepartment
)

func Employee() IEmployee {
	if localEmployee == nil {
		panic("implement not found for interface IEmployee, forgot register?")
	}
	return localEmployee
}

func RegisterEmployee(i IEmployee) {
	localEmployee = i
}

func EmployeeActivity() IEmployeeActivity {
	if localEmployeeActivity == nil {
		panic("implement not found for interface IEmployeeActivity, forgot register?")
	}
	return localEmployeeActivity
}

func RegisterEmployeeActivity(i IEmployeeActivity) {
	localEmployeeActivity = i
}

func EmployeeDepartment() IEmployeeDepartment {
	if localEmployeeDepartment == nil {
		panic("implement not found for interface IEmployeeDepartment, forgot register?")
	}
	return localEmployeeDepartment
}

func RegisterEmployeeDepartment(i IEmployeeDepartment) {
	localEmployeeDepartment = i
}
