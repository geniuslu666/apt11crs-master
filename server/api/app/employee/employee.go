package employee

import (
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeDeleteReq 删除员工
type EmployeeCheckAuthReq struct {
	g.Meta `path:"/employee/checkAuth" method:"post" tags:"APP_EMPLOYEE" summary:"[员工专区]检查是否是员工"`
}

type EmployeeCheckAuthRes struct {
	Auth bool `json:"auth"   dc:"是否是员工"`
}
