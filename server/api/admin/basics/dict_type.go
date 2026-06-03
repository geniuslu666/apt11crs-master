package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type DictTypeTreeReq struct {
	g.Meta `path:"/dictType/tree" tags:"ADMIN" summary:"字典类型_字典类型树列表" method:"get"`
}

type DictTypeTreeRes struct {
	List []*input_basics.DictTypeTree `json:"list"   dc:"数据列表"`
}

type DictTypeEditReq struct {
	g.Meta `path:"/dictType/edit" method:"post" tags:"ADMIN" summary:"字典类型_修改/新增字典类型"`
	input_basics.DictTypeEditInp
}

type DictTypeEditRes struct{}

type DictTypeDeleteReq struct {
	g.Meta `path:"/dictType/delete" method:"post" tags:"ADMIN" summary:"字典类型_删除字典类型"`
	input_basics.DictTypeDeleteInp
}

type DictTypeDeleteRes struct{}
