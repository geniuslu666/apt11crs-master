package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderChangeListReq 查询入住订单变更信息表列表
type OrderChangeListReq struct {
	g.Meta `path:"/pmsAppReservationChange/list" method:"get" tags:"ADMIN_PMS" summary:"获取入住订单变更信息表列表"`
	input_hotel.OrderChangeListInp
}

type OrderChangeListRes struct {
	input_form.PageRes
	List []*input_hotel.OrderChangeListModel `json:"list"   dc:"数据列表"`
}

// OrderChangeExportReq 导出入住订单变更信息表列表
type OrderChangeExportReq struct {
	g.Meta `path:"/pmsAppReservationChange/export" method:"get" tags:"ADMIN_PMS" summary:"导出入住订单变更信息表列表"`
	input_hotel.OrderChangeListInp
}

type OrderChangeExportRes struct{}

// OrderChangeViewReq 获取入住订单变更信息表指定信息
type OrderChangeViewReq struct {
	g.Meta `path:"/pmsAppReservationChange/view" method:"get" tags:"ADMIN_PMS" summary:"获取入住订单变更信息表指定信息"`
	input_hotel.OrderChangeViewInp
}

type OrderChangeViewRes struct {
	*input_hotel.OrderChangeViewModel
}

// OrderChangeEditReq 修改/新增入住订单变更信息表
type OrderChangeEditReq struct {
	g.Meta `path:"/pmsAppReservationChange/edit" method:"post" tags:"ADMIN_PMS" summary:"修改/新增入住订单变更信息表"`
	input_hotel.OrderChangeEditInp
}

type OrderChangeEditRes struct{}

// OrderChangeDeleteReq 删除入住订单变更信息表
type OrderChangeDeleteReq struct {
	g.Meta `path:"/pmsAppReservationChange/delete" method:"post" tags:"ADMIN_PMS" summary:"删除入住订单变更信息表"`
	input_hotel.OrderChangeDeleteInp
}

type OrderChangeDeleteRes struct{}
