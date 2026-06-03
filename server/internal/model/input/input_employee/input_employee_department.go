// Package input_employee

package input_employee

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/tree"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeDepartmentMaxSortInp 最大排序
type EmployeeDepartmentMaxSortInp struct {
	Id uint64 `json:"id" dc:"部门ID"`
}

type EmployeeDepartmentMaxSortModel struct {
	Sort int `json:"sort"`
}

// EmployeeDepartmentEditInp 修改/新增员工部门数据
type EmployeeDepartmentEditInp struct {
	entity.EmployeeDepartment
}

func (in *EmployeeDepartmentEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("部门名称不能为空")
		return
	}

	if in.Id > 0 && in.Id == in.ParentId {
		err = gerror.New("上级部门不能是自己")
		return
	}

	return
}

type EmployeeDepartmentEditModel struct{}

// EmployeeDepartmentUpdateFields 修改数据字段过滤
type EmployeeDepartmentUpdateFields struct {
	Id          uint64 `json:"id"          description:"部门ID"`
	Name        string `json:"name"        description:"部门名称"`
	ParentId    uint64 `json:"parentId"    description:"上级部门ID"`
	Level       int    `json:"level"       description:"部门层级"`
	Path        string `json:"path"        description:"部门路径"`
	ManagerId   uint64 `json:"managerId"   description:"部门负责人ID"`
	Description string `json:"description" description:"部门描述"`
	Sort        int    `json:"sort"        description:"排序"`
	Status      int    `json:"status"      description:"状态"`
}

// EmployeeDepartmentInsertFields 新增数据字段过滤
type EmployeeDepartmentInsertFields struct {
	Name        string `json:"name"        description:"部门名称"`
	ParentId    uint64 `json:"parentId"    description:"上级部门ID"`
	Level       int    `json:"level"       description:"部门层级"`
	Path        string `json:"path"        description:"部门路径"`
	ManagerId   uint64 `json:"managerId"   description:"部门负责人ID"`
	Description string `json:"description" description:"部门描述"`
	Sort        int    `json:"sort"        description:"排序"`
	Status      int    `json:"status"      description:"状态"`
}

// EmployeeDepartmentDeleteInp 删除员工部门
type EmployeeDepartmentDeleteInp struct {
	Id interface{} `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
}
type EmployeeDepartmentDeleteModel struct{}

// EmployeeDepartmentBatchDeleteInp 批量删除员工部门
type EmployeeDepartmentBatchDeleteInp struct {
	Ids []int64 `json:"ids" v:"required#部门ID不能为空" dc:"部门ID列表"`
}
type EmployeeDepartmentBatchDeleteModel struct{}

// EmployeeDepartmentViewInp 获取信息
type EmployeeDepartmentViewInp struct {
	Id int64 `json:"id" v:"required#部门ID不能为空" dc:"部门ID"`
}

type EmployeeDepartmentViewModel struct {
	entity.EmployeeDepartment
	EmployeeCount int `json:"employeeCount" dc:"员工数量"`
}

// EmployeeDepartmentListInp 获取列表
type EmployeeDepartmentListInp struct {
	Name              string        `json:"name"             dc:"部门名称"`
	Status            int           `json:"status"           dc:"状态"`
	CreatedAt         []*gtime.Time `json:"createdAt"        dc:"创建时间"`
	WithChildrenTotal bool          `json:"withChildrenTotal" dc:"员工数是否包含子部门（true=包含）"`
}

func (in *EmployeeDepartmentListInp) Filter(ctx context.Context) (err error) {
	return
}

// EmployeeDepartmentTree 员工部门树结构
type EmployeeDepartmentTree struct {
	entity.EmployeeDepartment
	Key           uint64                    `json:"key"           dc:"节点Key"`
	Title         string                    `json:"title"         dc:"节点标题"`
	EmployeeCount int                       `json:"employeeCount" dc:"员工数量"`
	Children      []*EmployeeDepartmentTree `json:"children"      dc:"子节点"`
}

type EmployeeDepartmentListModel struct {
	List []*EmployeeDepartmentTree `json:"list"`
}

// EmployeeDepartmentSwitchInp 更新员工部门状态
type EmployeeDepartmentSwitchInp struct {
	Id     uint64 `json:"id"     v:"required#部门ID不能为空" dc:"部门ID"`
	Status int    `json:"status" v:"required#状态不能为空"   dc:"状态"`
}

func (in *EmployeeDepartmentSwitchInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
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

type EmployeeDepartmentSwitchModel struct{}

type EmployeeDepartmentOptionInp struct {
	input_form.PageReq
}

type EmployeeDepartmentOptionModel struct {
	List []*EmployeeDepartmentTree `json:"list"`
}

// EmployeeDepartmentTreeOption 关系树选项
type EmployeeDepartmentTreeOption struct {
	Id       int64       `json:"id"   dc:"部门ID"`
	ParentId int64       `json:"parentId"  dc:"父部门ID"`
	Name     string      `json:"name" dc:"部门名称"`
	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *EmployeeDepartmentTreeOption) ID() int64 {
	return t.Id
}

// PID 获取父级节点ID
func (t *EmployeeDepartmentTreeOption) PID() int64 {
	return t.ParentId
}

// SetChildren 设置子节点数据
func (t *EmployeeDepartmentTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}
