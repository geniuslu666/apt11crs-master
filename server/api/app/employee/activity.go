package employee

import (
	"APT/internal/model/input/input_employee"

	"github.com/gogf/gf/v2/frame/g"
)

type EmployeeActivityListReq struct {
	g.Meta `path:"/employee/activityList" method:"post" tags:"APP_EMPLOYEE" summary:"[员工活动]列表"`
	input_employee.EmployeeActivityAppListInp
}

type EmployeeActivityListRes struct {
	List  []*input_employee.EmployeeActivityAppListModel `json:"list"   dc:"数据列表"`
	Count int                                            `json:"count"   dc:"数据总数"`
}

type EmployeeActivityViewReq struct {
	g.Meta `path:"/employee/activityDetail" method:"post" tags:"APP_EMPLOYEE" summary:"[员工活动]详情"`
	input_employee.EmployeeActivityAppViewInp
}

type EmployeeActivityViewRes struct {
	*input_employee.EmployeeActivityAppViewModel
}

type EmployeeActivityReceiveCouponReq struct {
	g.Meta `path:"/employee/activityReceiveCoupon" method:"post" tags:"APP_EMPLOYEE" summary:"[员工活动]领取礼品券"`
	input_employee.EmployeeActivityReceiveCouponInp
}

type EmployeeActivityReceiveCouponRes struct{}
