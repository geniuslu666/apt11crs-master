package basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuEditReq struct {
	g.Meta `path:"/menu/edit" method:"post" tags:"ADMIN" summary:"菜单_修改/新增菜单"`
	input_basics.MenuEditInp
}

type MenuEditRes struct{}

type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"ADMIN" summary:"菜单_删除菜单"`
	input_basics.MenuDeleteInp
}

type MenuDeleteRes struct{}

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"ADMIN" summary:"菜单_获取菜单列表"`
	input_basics.MenuListInp
}

type MenuListRes struct {
	*input_basics.MenuListModel
	input_form.PageRes
}

type MenuReq struct {
	g.Meta `path:"/pms/menu" method:"post" tags:"ADMIN" summary:"PMS菜单_列表"`
	Pid    int `json:"pid" dc:"父级菜单id"`
}
type MenuRes struct {
	List []entity.AdminMenu `json:"list" dc:"菜单列表"`
}
