// Package sysin

package input_car

import (
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// CarAddressTypeListInp 获取订餐合作类型列表
type CarAddressTypeListInp struct {
	input_form.PageReq
	TypeName string `json:"typeName"     dc:"地址类型"`
	Status   int    `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *CarAddressTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarAddressTypeListModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"地址类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	Sort     int         `json:"sort"     dc:"排序 越大越靠前"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}
