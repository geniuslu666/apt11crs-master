// Package car
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.0.0
package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// AddressListReq 查询接送机地点管理列表
type AddressListReq struct {
	g.Meta `path:"/CarAddress/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机地点管理列表"`
	input_car.CarAddressListInp
}

type AddressListRes struct {
	input_form.PageRes
	List []*input_car.CarAddressListModel `json:"list"   dc:"数据列表"`
}

// AddressExportReq 导出接送机地点管理列表
type AddressExportReq struct {
	g.Meta `path:"/CarAddress/export" method:"get" tags:"ADMIN_CAR" summary:"导出接送机地点管理列表"`
	input_car.CarAddressListInp
}

type AddressExportRes struct{}

// AddressViewReq 获取接送机地点详情
type AddressViewReq struct {
	g.Meta `path:"/CarAddress/view" method:"get" tags:"ADMIN_CAR" summary:"获取接送机地点详情"`
	input_car.CarAddressViewInp
}

type AddressViewRes struct {
	*input_car.CarAddressViewModel
}

// AddressEditReq 修改/新增接送机地点
type AddressEditReq struct {
	g.Meta `path:"/CarAddress/edit" method:"post" tags:"ADMIN_CAR" summary:"修改/新增接送机地点"`
	input_car.CarAddressEditInp
}

type AddressEditRes struct{}

// AddressDeleteReq 删除接送机地点
type AddressDeleteReq struct {
	g.Meta `path:"/CarAddress/delete" method:"post" tags:"ADMIN_CAR" summary:"删除接送机地点"`
	input_car.CarAddressDeleteInp
}

type AddressDeleteRes struct{}

// AddressStatusReq CarAddressStatusReq 更新接送机地点状态
type AddressStatusReq struct {
	g.Meta `path:"/CarAddress/status" method:"post" tags:"ADMIN_CAR" summary:"更新接送机地点状态"`
	input_car.CarAddressStatusInp
}

type AddressStatusRes struct{}
