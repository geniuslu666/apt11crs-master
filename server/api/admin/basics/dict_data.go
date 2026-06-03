package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type DictDataEditReq struct {
	g.Meta `path:"/dictData/edit" method:"post" tags:"ADMIN" summary:"字典数据_修改/新增字典数据"`
	input_basics.DictDataEditInp
}

type DictDataEditRes struct{}

type DictDataDeleteReq struct {
	g.Meta `path:"/dictData/delete" method:"post" tags:"ADMIN" summary:"字典数据_删除字典数据"`
	input_basics.DictDataDeleteInp
}

type DictDataDeleteRes struct{}

type DictDataListReq struct {
	g.Meta `path:"/dictData/list" method:"get" tags:"ADMIN" summary:"字典数据_获取字典数据列表"`
	input_basics.DictDataListInp
}

type DictDataListRes struct {
	List []*input_basics.DictDataListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type DictDataSelectReq struct {
	g.Meta `path:"/dictData/option/{Type}" method:"get" tags:"ADMIN" summary:"字典数据_获取指定字典选项"`
	input_basics.DataSelectInp
}

type DictDataSelectRes input_basics.DataSelectModel

type DictDataSelectsReq struct {
	g.Meta `path:"/dictData/options" method:"get" tags:"ADMIN" summary:"字典数据_获取多个字典选项"`
	Types  []string `json:"types"`
}

type DictDataSelectsRes map[string]input_basics.DataSelectModel
