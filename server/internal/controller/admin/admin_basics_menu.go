package admin

import (
	"APT/internal/dao"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) MenuEdit(ctx context.Context, req *basics.MenuEditReq) (res *basics.MenuEditRes, err error) {
	err = service.BasicsAdminMenu().Edit(ctx, &req.MenuEditInp)
	return
}
func (c *ControllerBasics) MenuDelete(ctx context.Context, req *basics.MenuDeleteReq) (res *basics.MenuDeleteRes, err error) {
	err = service.BasicsAdminMenu().Delete(ctx, &req.MenuDeleteInp)
	return
}
func (c *ControllerBasics) MenuList(ctx context.Context, req *basics.MenuListReq) (res *basics.MenuListRes, err error) {
	res = new(basics.MenuListRes)
	res.MenuListModel, err = service.BasicsAdminMenu().List(ctx, &req.MenuListInp)
	return
}
func (c *ControllerBasics) Menu(ctx context.Context, req *basics.MenuReq) (res *basics.MenuRes, err error) {
	var (
		orm = g.Model(dao.AdminMenu.Table()).Ctx(ctx)
	)
	res = new(basics.MenuRes)
	orm = orm.Where(dao.AdminMenu.Columns().Pid, req.Pid)
	orm = orm.Where(dao.AdminMenu.Columns().Type, 2)
	orm = orm.Where(dao.AdminMenu.Columns().Hidden, 2)
	if err = orm.Scan(&res.List); err != nil {
		return
	}
	return
}
