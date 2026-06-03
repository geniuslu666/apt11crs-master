package admin

import (
	"APT/internal/consts"
	"APT/internal/library/contexts"
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) RoleList(ctx context.Context, req *basics.RoleListReq) (res *basics.RoleListRes, err error) {
	list, totalCount, err := service.BasicsAdminRole().List(ctx, &req.RoleListInp)
	if err != nil {
		return
	}

	res = new(basics.RoleListRes)
	res.RoleListModel = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) RoleDynamic(ctx context.Context, req *basics.RoleDynamicReq) (res *basics.RoleDynamicRes, err error) {
	return service.BasicsAdminMenu().GetMenuList(ctx, contexts.GetUserId(ctx))
}
func (c *ControllerBasics) RoleUpdatePermissions(ctx context.Context, req *basics.RoleUpdatePermissionsReq) (res *basics.RoleUpdatePermissionsRes, err error) {
	err = service.BasicsAdminRole().UpdatePermissions(ctx, &req.UpdatePermissionsInp)
	return
}
func (c *ControllerBasics) RoleGetPermissions(ctx context.Context, req *basics.RoleGetPermissionsReq) (res *basics.RoleGetPermissionsRes, err error) {
	data, err := service.BasicsAdminRole().GetPermissions(ctx, &req.GetPermissionsInp)
	if err != nil {
		return
	}

	res = new(basics.RoleGetPermissionsRes)
	res.GetPermissionsModel = data
	return
}
func (c *ControllerBasics) RoleEdit(ctx context.Context, req *basics.RoleEditReq) (res *basics.RoleEditRes, err error) {
	err = service.BasicsAdminRole().Edit(ctx, &req.RoleEditInp)
	return
}
func (c *ControllerBasics) RoleDelete(ctx context.Context, req *basics.RoleDeleteReq) (res *basics.RoleDeleteRes, err error) {
	err = service.BasicsAdminRole().Delete(ctx, &req.RoleDeleteInp)
	return
}
func (c *ControllerBasics) RoleDataScopeSelect(_ context.Context, _ *basics.RoleDataScopeSelectReq) (res *basics.RoleDataScopeSelectRes, err error) {
	res = new(basics.RoleDataScopeSelectRes)
	res.List = consts.DataScopeSelect
	return
}
func (c *ControllerBasics) RoleDataScopeEdit(ctx context.Context, req *basics.RoleDataScopeEditReq) (res *basics.RoleDataScopeEditRes, err error) {
	err = service.BasicsAdminRole().DataScopeEdit(ctx, &req.DataScopeEditInp)
	return
}
