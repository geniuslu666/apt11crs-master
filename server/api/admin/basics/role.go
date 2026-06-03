package basics

import (
	"APT/internal/consts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"ADMIN" summary:"角色_获取角色列表"`
	input_basics.RoleListInp
}

type RoleListRes struct {
	*input_basics.RoleListModel
	input_form.PageRes
}

type RoleDynamicReq struct {
	g.Meta `path:"/role/dynamic" method:"get" tags:"ADMIN" summary:"角色_获取动态路由" description:"获取登录用户动态路由"`
}

type RoleDynamicRes struct {
	List []*input_basics.MenuRoute `json:"list"   description:"数据列表"`
}

type RoleUpdatePermissionsReq struct {
	g.Meta `path:"/role/updatePermissions" method:"post" tags:"ADMIN" summary:"角色_编辑角色菜单权限"`
	input_basics.UpdatePermissionsInp
}

type RoleUpdatePermissionsRes struct{}

type RoleGetPermissionsReq struct {
	g.Meta `path:"/role/getPermissions" method:"get" tags:"ADMIN" summary:"角色_获取指定角色权限"`
	input_basics.GetPermissionsInp
}

type RoleGetPermissionsRes struct {
	*input_basics.GetPermissionsModel
}

type RoleEditReq struct {
	g.Meta `path:"/role/edit" method:"post" tags:"ADMIN" summary:"角色_修改/新增角色"`
	input_basics.RoleEditInp
}

type RoleEditRes struct{}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"ADMIN" summary:"角色_删除角色"`
	input_basics.RoleDeleteInp
}

type RoleDeleteRes struct{}

type RoleDataScopeSelectReq struct {
	g.Meta `path:"/role/dataScope/select" method:"get" tags:"ADMIN" summary:"角色_获取数据权限选项"`
}

type RoleDataScopeSelectRes struct {
	List []consts.GroupScopeSelect `json:"list" dc:"数据选项"`
}

type RoleDataScopeEditReq struct {
	g.Meta `path:"/role/dataScope/edit" method:"post" tags:"ADMIN" summary:"角色_修改指定角色的数据权限"`
	input_basics.DataScopeEditInp
}

type RoleDataScopeEditRes struct{}
