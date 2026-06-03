package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type RoomUnitListReq struct {
	g.Meta `path:"/pmsRoomUnit/list" method:"get" tags:"ADMIN_PMS" summary:"房间_列表"`
	input_hotel.PmsRoomUnitListInp
}

type RoomUnitListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsRoomUnitListModel `json:"list"   dc:"数据列表"`
}

type RoomUnitExportReq struct {
	g.Meta `path:"/pmsRoomUnit/export" method:"get" tags:"ADMIN_PMS" summary:"房间_导出"`
	input_hotel.PmsRoomUnitListInp
}

type RoomUnitExportRes struct{}

type RoomUnitViewReq struct {
	g.Meta `path:"/pmsRoomUnit/view" method:"get" tags:"ADMIN_PMS" summary:"房间_详情"`
	input_hotel.PmsRoomUnitViewInp
}

type RoomUnitViewRes struct {
	*input_hotel.PmsRoomUnitViewModel
}

type RoomUnitEditReq struct {
	g.Meta `path:"/pmsRoomUnit/edit" method:"post" tags:"ADMIN_PMS" summary:"房间_编辑"`
	input_hotel.PmsRoomUnitEditInp
}

type RoomUnitEditRes struct{}

type RoomUnitDeleteReq struct {
	g.Meta `path:"/pmsRoomUnit/delete" method:"post" tags:"ADMIN_PMS" summary:"房间_删除"`
	input_hotel.PmsRoomUnitDeleteInp
}

type RoomUnitDeleteRes struct{}
