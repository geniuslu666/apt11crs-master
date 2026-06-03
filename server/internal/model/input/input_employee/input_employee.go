package input_employee

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// EmployeeUpdateFields 修改员工管理字段过滤
type EmployeeUpdateFields struct {
	Name         string `json:"name"              dc:"真实姓名"`
	Phone        string `json:"phone"             dc:"手机号"`
	PhoneArea    string `json:"phoneArea"         dc:"区号"`
	MemberId     int    `json:"memberId"          dc:"会员ID"`
	DepartmentId int    `json:"departmentId"      dc:"员工部门ID"`
	EmployeeNo   string `json:"employeeNo"        dc:"员工编号"`
	Status       int    `json:"status"            dc:"状态"`
	Remark       string `json:"remark"            dc:"备注"`
}

// EmployeeInsertFields 新增员工管理字段过滤
type EmployeeInsertFields struct {
	Name         string `json:"name"              dc:"真实姓名"`
	Phone        string `json:"phone"             dc:"手机号"`
	PhoneArea    string `json:"phoneArea"         dc:"区号"`
	MemberId     int    `json:"memberId"          dc:"会员ID"`
	DepartmentId int    `json:"departmentId"      dc:"员工部门ID"`
	EmployeeNo   string `json:"employeeNo"        dc:"员工编号"`
	Status       int    `json:"status"            dc:"状态"`
	Remark       string `json:"remark"            dc:"备注"`
}

// EmployeeEditInp 修改/新增员工
type EmployeeEditInp struct {
	entity.Employee
}

func (in *EmployeeEditInp) Filter(ctx context.Context) (err error) {

	return
}

type EmployeeEditModel struct{}

// EmployeeDeleteInp 删除员工管理
type EmployeeDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *EmployeeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeDeleteModel struct{}

// EmployeeViewInp 获取指定员工管理信息
type EmployeeViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *EmployeeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeViewModel struct {
	entity.Employee
	DepartmentDetail *struct {
		gmeta.Meta `orm:"table:hg_employee_department"`
		*entity.EmployeeDepartment
	} `json:"departmentDetail" orm:"with:id=department_id"  dc:"部门信息"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"     orm:"member_no"     dc:"会员号"`
		Mail       string `json:"mail"         orm:"mail"          dc:"邮箱"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
}

// EmployeeListInp 获取员工管理列表
type EmployeeListInp struct {
	input_form.PageReq
	Name         string        `json:"name"      dc:"姓名"`
	Phone        string        `json:"phone"     dc:"手机号"`
	DepartmentId int           `json:"departmentId"  dc:"员工部门ID"`
	Status       int           `json:"status"    dc:"状态"`
	CreateAt     []*gtime.Time `json:"createdAt" dc:"加入时间"`
	EmployeeIds  string        `json:"employeeIds"   dc:"员工Ids"`
}

func (in *EmployeeListInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeListModel struct {
	Id               int64       `json:"id"                      dc:"id"`
	Name             string      `json:"name"              dc:"真实姓名"`
	Phone            string      `json:"phone"             dc:"手机号"`
	PhoneArea        string      `json:"phoneArea"         dc:"区号"`
	MemberId         int         `json:"memberId"          dc:"会员ID"`
	DepartmentId     int         `json:"departmentId"      dc:"员工部门ID"`
	EmployeeNo       string      `json:"employeeNo"        dc:"员工编号"`
	Status           int         `json:"status"            dc:"状态"`
	Remark           string      `json:"remark"            dc:"备注"`
	CreatedAt        *gtime.Time `json:"createdAt"                dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"                dc:"更新时间"`
	DepartmentDetail *struct {
		gmeta.Meta `orm:"table:hg_employee_department"`
		*entity.EmployeeDepartment
	} `json:"departmentDetail" orm:"with:id=department_id"  dc:"部门信息"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"     orm:"member_no"     dc:"会员号"`
		Mail       string `json:"mail"         orm:"mail"          dc:"邮箱"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
}

// EmployeeStatusInp 更新员工状态
type EmployeeStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *EmployeeStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}

	return
}

type EmployeeStatusModel struct{}

// EmployeeBindInp 绑定用户
type EmployeeBindInp struct {
	Id       int `json:"id" dc:"ID"`
	MemberId int `json:"memberId" dc:"会员ID"`
}

func (in *EmployeeBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.MemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type EmployeeBindModel struct{}

// EmployeeUnbindInp 解绑会员
type EmployeeUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *EmployeeUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeUnbindModel struct{}
