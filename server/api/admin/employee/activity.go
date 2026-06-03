package employee

import (
	"APT/internal/model/input/input_employee"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type ActivityListReq struct {
	g.Meta `path:"/employeeActivity/list" method:"get" tags:"ADMIN_EMPLOYEE" summary:"员工活动_列表"`
	input_employee.EmployeeActivityListInp
}

type ActivityListRes struct {
	input_form.PageRes
	List []*input_employee.EmployeeActivityListModel `json:"list"   dc:"数据列表"`
}

type ActivityViewReq struct {
	g.Meta `path:"/employeeActivity/view" method:"get" tags:"ADMIN_EMPLOYEE" summary:"员工活动_详情"`
	input_employee.EmployeeActivityViewInp
}

type ActivityViewRes struct {
	*input_employee.EmployeeActivityViewModel
}

type ActivityEditReq struct {
	g.Meta `path:"/employeeActivity/edit" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工活动_修改/新增"`
	input_employee.EmployeeActivityEditInp
}

type ActivityEditRes struct{}

type ActivityDeleteReq struct {
	g.Meta `path:"/employeeActivity/delete" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工活动_删除"`
	input_employee.EmployeeActivityDeleteInp
}

type ActivityDeleteRes struct{}

type ActivityStatusReq struct {
	g.Meta `path:"/employeeActivity/status" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工活动_更新"`
	input_employee.EmployeeActivityStatusInp
}

type ActivityStatusRes struct{}

type ActivityBatchIssueCouponsReq struct {
	g.Meta `path:"/employeeActivity/batchIssueCoupons" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工活动_批量发放预约券"`
	input_employee.EmployeeActivityBatchIssueCouponsInp
}

type ActivityBatchIssueCouponsRes struct {
	*input_employee.EmployeeActivityBatchIssueCouponsModel
}

type ActivityCouponRecordListReq struct {
	g.Meta `path:"/employeeActivity/couponRecordList" method:"get" tags:"ADMIN_EMPLOYEE" summary:"员工活动_券领取记录列表"`
	input_employee.EmployeeActivityCouponRecordListInp
}

type ActivityCouponRecordListRes struct {
	input_form.PageRes
	List []*input_employee.EmployeeActivityCouponRecordListModel `json:"list" dc:"数据列表"`
}
