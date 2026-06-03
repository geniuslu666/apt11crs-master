package logic_employee

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_employee"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sEmployee struct{}

func NewEmployee() *sEmployee {
	return &sEmployee{}
}

func init() {
	service.RegisterEmployee(NewEmployee())
}

// Model 员工管理ORM模型
func (s *sEmployee) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Employee.Ctx(ctx), option...)
}

// List 获取员工管理列表
func (s *sEmployee) List(ctx context.Context, in *input_employee.EmployeeListInp) (list []*input_employee.EmployeeListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_employee.EmployeeListModel{})

	// 查询名称
	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.Employee.Columns().Name, "%"+in.Name+"%")
	}

	// 查询电话
	if !g.IsEmpty(in.Phone) {
		mod = mod.WhereLike(dao.Employee.Columns().Phone, "%"+in.Phone+"%")
	}

	if !g.IsEmpty(in.Status) && in.Status > 0 {
		mod = mod.Where(dao.Employee.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.DepartmentId) && in.DepartmentId > 0 {
		// 获取当前部门及其所有子部门ID（包括多级嵌套）
		departmentIds := []int{int(in.DepartmentId)} // 包含当前部门ID

		// 通过Path字段查找所有子部门：Path包含当前部门ID的都是子部门
		childDepartmentIds, err := dao.EmployeeDepartment.Ctx(ctx).
			WhereLike(dao.EmployeeDepartment.Columns().Path, "%,"+g.NewVar(in.DepartmentId).String()+",%").
			WhereOrLike(dao.EmployeeDepartment.Columns().Path, g.NewVar(in.DepartmentId).String()+",%").
			WhereOrLike(dao.EmployeeDepartment.Columns().Path, "%,"+g.NewVar(in.DepartmentId).String()).
			Fields(dao.EmployeeDepartment.Columns().Id).Array()

		if err == nil && len(childDepartmentIds) > 0 {
			childIds := g.NewVar(childDepartmentIds).Ints()
			departmentIds = append(departmentIds, childIds...)
		}

		mod = mod.WhereIn(dao.Employee.Columns().DepartmentId, departmentIds)
	}

	// 查询创建时间
	if len(in.CreateAt) == 2 {
		mod = mod.WhereBetween(dao.Employee.Columns().CreatedAt, in.CreateAt[0], in.CreateAt[1])
	}

	if !g.IsEmpty(in.EmployeeIds) {
		mod = mod.WhereIn(dao.Employee.Columns().Id, strings.Split(in.EmployeeIds, ","))
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.Employee.Columns().Id)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取员工列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取员工列表失败，请稍后重试！")
			return
		}
	}

	return
}

// Edit 修改/新增员工
func (s *sEmployee) Edit(ctx context.Context, in *input_employee.EmployeeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_employee.EmployeeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改员工信息失败，请稍后重试！")
			}
			return
		}

		// 新增
		var (
			lastInsertId int64
		)
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_employee.EmployeeInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增员工失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		return
	})
}

// Delete 删除员工管理
func (s *sEmployee) Delete(ctx context.Context, in *input_employee.EmployeeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除员工失败，请稍后重试！")
		return
	}
	return
}

// View 获取员工管理指定信息
func (s *sEmployee) View(ctx context.Context, in *input_employee.EmployeeViewInp) (res *input_employee.EmployeeViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取员工信息失败，请稍后重试！")
		return
	}
	return
}

// Status 更新员工状态
func (s *sEmployee) Status(ctx context.Context, in *input_employee.EmployeeStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.Employee.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新员工状态失败，请稍后重试！")
			return
		}

		return
	})
}

// Bind 绑定用户
func (s *sEmployee) Bind(ctx context.Context, in *input_employee.EmployeeBindInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 判断会员是否已绑定
		memberCount, err := dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, in.MemberId).Count()
		if memberCount > 0 {
			err = gerror.New("该用户已被绑定，请重新选择用户")
			return
		}

		if _, err = dao.Employee.Ctx(ctx).WherePri(in.Id).Data(g.Map{
			dao.Employee.Columns().MemberId: in.MemberId,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		}
		return
	})
}

// Unbind 解绑用户
func (s *sEmployee) Unbind(ctx context.Context, in *input_employee.EmployeeUnbindInp) (err error) {

	if _, err = dao.Employee.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.Employee.Columns().MemberId: 0,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "解绑用户失败，请稍后重试！")
	}

	return
}

// CheckAuth 验证员工
func (s *sEmployee) CheckAuth(ctx context.Context) (auth bool, err error) {
	var (
		MemberInfo *model.MemberIdentity
		Employee   *entity.Employee
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	if err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, MemberInfo.Id).Scan(&Employee); err != nil {
		auth = false
		return
	}
	if g.IsEmpty(Employee) {
		// 员工信息错误
		auth = false
		return
	}

	if Employee.Status != 1 {
		auth = false
		return
	}

	auth = true
	return
}
