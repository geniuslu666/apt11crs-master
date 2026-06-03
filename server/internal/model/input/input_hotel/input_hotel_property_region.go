// Package sysin

package input_hotel

import (
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PropertyRegionListInp 获取地区列表
type PropertyRegionListInp struct {
	input_form.PageReq
	Status int `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *PropertyRegionListInp) Filter(ctx context.Context) (err error) {
	return
}

type PropertyRegionListModel struct {
	Id       int         `json:"id"       dc:"id"`
	Name     string      `json:"name"     dc:"地区名称"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}
