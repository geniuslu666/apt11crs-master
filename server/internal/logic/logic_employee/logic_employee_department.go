package logic_employee

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_employee"
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sEmployeeDepartment struct{}

func NewEmployeeDepartment() *sEmployeeDepartment {
	return &sEmployeeDepartment{}
}

func init() {
	service.RegisterEmployeeDepartment(NewEmployeeDepartment())
}

// Model 员工部门ORM模型
func (s *sEmployeeDepartment) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.EmployeeDepartment.Ctx(ctx), option...)
}

// List 获取员工部门列表
func (s *sEmployeeDepartment) List(ctx context.Context, in *input_employee.EmployeeDepartmentListInp) (res *input_employee.EmployeeDepartmentListModel, err error) {
	mod := s.Model(ctx)

	// 查询条件
	if in.Name != "" {
		mod = mod.WhereLike(dao.EmployeeDepartment.Columns().Name, "%"+in.Name+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(dao.EmployeeDepartment.Columns().Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.EmployeeDepartment.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	var models []*entity.EmployeeDepartment
	if err = mod.OrderAsc(dao.EmployeeDepartment.Columns().Sort).
		OrderAsc(dao.EmployeeDepartment.Columns().Id).Scan(&models); err != nil {
		return
	}

	var list []*input_employee.EmployeeDepartmentTree
	if len(models) > 0 {
		// 获取所有部门的员工数量统计
		countMap, err := s.getEmployeeCountMap(ctx, models)
		if err != nil {
			// 如果获取员工数量失败，记录错误但不影响主流程
			g.Log().Warning(ctx, "获取部门员工数量失败:", err)
			countMap = make(map[uint64]int)
		}

		for _, v := range models {
			item := &input_employee.EmployeeDepartmentTree{
				EmployeeDepartment: *v,
				Key:                v.Id,
				Title:              v.Name,
				EmployeeCount:      countMap[v.Id], // 设置员工数量
			}
			list = append(list, item)
		}

		// 构建树结构
		list = s.buildTree(list, 0)

		// 可选：将员工数切换为“包含子部门总员工数”
		if in.WithChildrenTotal {
			// 递归聚合：子树员工数 + 本部门直属员工数
			s.aggregateEmployeeCounts(list)
		}
	}

	res = &input_employee.EmployeeDepartmentListModel{
		List: list,
	}
	return
}

// Edit 修改/新增员工部门
func (s *sEmployeeDepartment) Edit(ctx context.Context, in *input_employee.EmployeeDepartmentEditInp) (err error) {
	// 验证唯一性
	if err = s.verifyUnique(ctx, in); err != nil {
		return
	}

	// 处理层级和路径
	if err = s.handleLevelAndPath(ctx, in); err != nil {
		return
	}

	// 使用事务确保数据一致性
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if in.Id > 0 {
			var (
				path          string
				oldDepartment *entity.EmployeeDepartment
			)

			// 获取原部门信息，用于判断是否需要更新子部门
			if err = tx.Model(dao.EmployeeDepartment.Table()).WherePri(in.Id).Scan(&oldDepartment); err != nil {
				return err
			}

			// 更新
			_, err = tx.Model(dao.EmployeeDepartment.Table()).
				Fields(input_employee.EmployeeDepartmentUpdateFields{}).
				WherePri(in.Id).Data(in).Update()
			if err != nil {
				return err
			}

			if !g.IsEmpty(in.Path) {
				path = in.Path + "," + gconv.String(in.Id)
			} else {
				path = gconv.String(in.Id)
			}
			_, err = tx.Model(dao.EmployeeDepartment.Table()).
				WherePri(in.Id).Data(g.Map{
				dao.EmployeeDepartment.Columns().Path: path,
			}).OmitEmptyData().Update()
			if err != nil {
				return err
			}

			// 如果父级部门发生变化，需要递归更新所有子部门的level和path
			if oldDepartment != nil && oldDepartment.ParentId != in.ParentId {
				err = s.updateChildrenLevelAndPathTx(ctx, tx, in.Id, in.Level, path)
				if err != nil {
					return err
				}
			}
		} else {
			var (
				lastInsertId int64
				path         string
			)
			// 新增
			if lastInsertId, err = tx.Model(dao.EmployeeDepartment.Table()).
				Fields(input_employee.EmployeeDepartmentInsertFields{}).
				Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
				return gerror.Wrap(err, "新增部门失败，请稍后重试！")
			}

			if lastInsertId < 1 {
				return gerror.New("新增失败，请稍后重试！")
			}

			if !g.IsEmpty(in.Path) {
				path = in.Path + "," + gconv.String(lastInsertId)
			} else {
				path = gconv.String(lastInsertId)
			}

			// 更新路径
			_, err = tx.Model(dao.EmployeeDepartment.Table()).
				WherePri(lastInsertId).Data(g.Map{
				dao.EmployeeDepartment.Columns().Path: path,
			}).Update()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Delete 删除员工部门
func (s *sEmployeeDepartment) Delete(ctx context.Context, in *input_employee.EmployeeDepartmentDeleteInp) (err error) {
	var models *entity.EmployeeDepartment
	if err = s.Model(ctx).WherePri(in.Id).Scan(&models); err != nil {
		return err
	}

	if models == nil {
		return gerror.New("数据不存在或已删除！")
	}

	// 检查是否有子部门
	childExist, err := s.Model(ctx).Where(dao.EmployeeDepartment.Columns().ParentId, models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !childExist.IsEmpty() {
		return gerror.New("请先删除该部门下的所有子部门！")
	}

	// 检查是否有员工
	employeeExist, err := dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().DepartmentId, models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !employeeExist.IsEmpty() {
		return gerror.New("该部门下还有员工，请先转移员工后再删除！")
	}

	_, err = s.Model(ctx).WherePri(in.Id).Delete()
	return
}

// BatchDelete 批量删除员工部门
func (s *sEmployeeDepartment) BatchDelete(ctx context.Context, in *input_employee.EmployeeDepartmentBatchDeleteInp) (err error) {
	for _, id := range in.Ids {
		if err = s.Delete(ctx, &input_employee.EmployeeDepartmentDeleteInp{Id: id}); err != nil {
			return
		}
	}
	return
}

// View 获取员工部门详情
func (s *sEmployeeDepartment) View(ctx context.Context, in *input_employee.EmployeeDepartmentViewInp) (res *input_employee.EmployeeDepartmentViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		return
	}

	if res == nil {
		return nil, gerror.New("数据不存在")
	}

	// 获取员工数量
	count, err := dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().DepartmentId, res.Id).Count()
	if err == nil {
		res.EmployeeCount = count
	}

	return
}

// Switch 更新员工部门状态
func (s *sEmployeeDepartment) Switch(ctx context.Context, in *input_employee.EmployeeDepartmentSwitchInp) (err error) {
	_, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.EmployeeDepartment.Columns().Status: in.Status,
	}).Update()
	return
}

// TreeOption 获取员工部门树选项
func (s *sEmployeeDepartment) TreeOption(ctx context.Context) (res *input_employee.EmployeeDepartmentOptionModel, err error) {
	var models []*entity.EmployeeDepartment
	if err = s.Model(ctx).
		Where(dao.EmployeeDepartment.Columns().Status, consts.StatusEnabled).
		OrderAsc(dao.EmployeeDepartment.Columns().Sort).
		OrderAsc(dao.EmployeeDepartment.Columns().Id).
		Scan(&models); err != nil {
		return
	}

	var list []*input_employee.EmployeeDepartmentTree
	for _, v := range models {
		item := &input_employee.EmployeeDepartmentTree{
			EmployeeDepartment: *v,
			Key:                v.Id,
			Title:              v.Name,
		}
		list = append(list, item)
	}

	// 构建树结构
	list = s.buildTree(list, 0)

	res = &input_employee.EmployeeDepartmentOptionModel{
		List: list,
	}
	return
}

// MaxSort 获取最大排序
func (s *sEmployeeDepartment) MaxSort(ctx context.Context, in *input_employee.EmployeeDepartmentMaxSortInp) (res *input_employee.EmployeeDepartmentMaxSortModel, err error) {
	if in.Id > 0 {
		if err = s.Model(ctx).Fields(dao.EmployeeDepartment.Columns().Sort).WherePri(in.Id).Scan(&res); err != nil {
			return
		}
	}

	if res == nil {
		res = new(input_employee.EmployeeDepartmentMaxSortModel)
	}

	maxSort, err := s.Model(ctx).Max(dao.EmployeeDepartment.Columns().Sort)
	if err != nil {
		return
	}

	res.Sort = gconv.Int(maxSort) + 1
	return
}

// GenTree 生成关系树
func (s *sEmployeeDepartment) GenTree(ctx context.Context, parentId uint64) (level int, newPath string, err error) {
	if parentId <= 0 {
		level = 1
		newPath = ""
		return
	}

	var parent *entity.EmployeeDepartment
	if err = s.Model(ctx).WherePri(parentId).Scan(&parent); err != nil {
		return
	}

	if parent == nil {
		err = gerror.New("上级部门不存在")
		return
	}

	level = parent.Level + 1
	newPath = parent.Path
	return
}

// 验证唯一性
func (s *sEmployeeDepartment) verifyUnique(ctx context.Context, in *input_employee.EmployeeDepartmentEditInp) (err error) {
	where := g.Map{
		dao.EmployeeDepartment.Columns().Name: in.Name,
	}

	for k, v := range where {
		if v == "" {
			continue
		}
		if err = hgorm.IsUnique(ctx, &dao.EmployeeDepartment, g.Map{k: v}, "部门名称已存在，请换一个", in.Id); err != nil {
			return
		}
	}
	return
}

// 处理层级和路径
func (s *sEmployeeDepartment) handleLevelAndPath(ctx context.Context, in *input_employee.EmployeeDepartmentEditInp) (err error) {
	level, path, err := s.GenTree(ctx, in.ParentId)
	if err != nil {
		return
	}

	in.Level = level
	in.Path = path
	return
}

// 获取员工数量统计 - 高效批量查询，避免N+1问题
func (s *sEmployeeDepartment) getEmployeeCountMap(ctx context.Context, departments []*entity.EmployeeDepartment) (countMap map[uint64]int, err error) {
	countMap = make(map[uint64]int)
	if len(departments) == 0 {
		return
	}

	// 提取部门ID列表
	var departmentIds []uint64
	for _, dept := range departments {
		departmentIds = append(departmentIds, dept.Id)
		// 预初始化为0，确保没有员工的部门也有计数
		countMap[dept.Id] = 0
	}

	// 使用单次SQL查询获取所有部门的员工数量
	var counts []struct {
		DepartmentId int64 `json:"department_id"`
		Count        int   `json:"count"`
	}

	if err = dao.Employee.Ctx(ctx).
		Fields("department_id, COUNT(*) as count").
		WhereIn(dao.Employee.Columns().DepartmentId, departmentIds).
		Group(dao.Employee.Columns().DepartmentId).
		Scan(&counts); err != nil {
		return
	}

	// 将查询结果映射到countMap
	for _, count := range counts {
		countMap[uint64(count.DepartmentId)] = count.Count
	}
	return
}

// 构建树结构
func (s *sEmployeeDepartment) buildTree(items []*input_employee.EmployeeDepartmentTree, parentId uint64) []*input_employee.EmployeeDepartmentTree {
	var tree []*input_employee.EmployeeDepartmentTree
	for _, item := range items {
		if item.ParentId == parentId {
			item.Children = s.buildTree(items, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}

// aggregateEmployeeCounts 递归聚合节点的员工数量（将 EmployeeCount 设置为：本部门直属员工数 + 全部子部门员工数）
func (s *sEmployeeDepartment) aggregateEmployeeCounts(nodes []*input_employee.EmployeeDepartmentTree) int {
	total := 0
	for _, n := range nodes {
		subtotal := n.EmployeeCount
		if len(n.Children) > 0 {
			subtotal += s.aggregateEmployeeCounts(n.Children)
		}
		n.EmployeeCount = subtotal
		total += subtotal
	}
	return total
}

// updateChildrenLevelAndPathTx 递归更新子部门的level和path (事务版本)
func (s *sEmployeeDepartment) updateChildrenLevelAndPathTx(ctx context.Context, tx gdb.TX, parentId uint64, parentLevel int, parentPath string) (err error) {
	var children []*entity.EmployeeDepartment

	// 获取所有直接子部门
	if err = tx.Model(dao.EmployeeDepartment.Table()).
		Where(dao.EmployeeDepartment.Columns().ParentId, parentId).
		Scan(&children); err != nil {
		return
	}

	for _, child := range children {
		// 计算新的level和path
		newLevel := parentLevel + 1
		var newPath string
		if !g.IsEmpty(parentPath) {
			newPath = parentPath + "," + gconv.String(child.Id)
		} else {
			newPath = gconv.String(child.Id)
		}

		// 更新当前子部门的level和path
		_, err = tx.Model(dao.EmployeeDepartment.Table()).
			WherePri(child.Id).
			Data(g.Map{
				dao.EmployeeDepartment.Columns().Level: newLevel,
				dao.EmployeeDepartment.Columns().Path:  newPath,
			}).Update()
		if err != nil {
			return
		}

		// 递归更新子部门的子部门
		if err = s.updateChildrenLevelAndPathTx(ctx, tx, child.Id, newLevel, newPath); err != nil {
			return
		}
	}

	return
}
