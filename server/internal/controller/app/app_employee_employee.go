package app

import (
	"context"

	"APT/api/app/employee"
	"APT/internal/service"
)

func (c *ControllerEmployee) EmployeeCheckAuth(ctx context.Context, req *employee.EmployeeCheckAuthReq) (res *employee.EmployeeCheckAuthRes, err error) {
	res = new(employee.EmployeeCheckAuthRes)
	if res.Auth, err = service.Employee().CheckAuth(ctx); err != nil {
		return
	}
	return
}
