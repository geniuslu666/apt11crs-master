package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type RoomTypeListReq struct {
	g.Meta `path:"/pmsRoomType/list" method:"get" tags:"ADMIN_PMS" summary:"房型_列表"`
	input_hotel.PmsRoomTypeListInp
}

type RoomTypeListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsRoomTypeListModel `json:"list"   dc:"数据列表"`
}

type RoomTypeViewReq struct {
	g.Meta `path:"/pmsRoomType/view" method:"get" tags:"ADMIN_PMS" summary:"房型_详情"`
	input_hotel.PmsRoomTypeViewInp
}

type RoomTypeViewRes struct {
	*input_hotel.PmsRoomTypeViewModel
}

type RoomTypeEditReq struct {
	g.Meta `path:"/pmsRoomType/edit" method:"post" tags:"ADMIN_PMS" summary:"房型_编辑"`
	input_hotel.PmsRoomTypeEditInp
}

type RoomTypeEditRes struct{}

type RoomTypeDeleteReq struct {
	g.Meta `path:"/pmsRoomType/delete" method:"post" tags:"ADMIN_PMS" summary:"房型_删除"`
	input_hotel.PmsRoomTypeDeleteInp
}

type RoomTypeDeleteRes struct{}

type RoomTypeIsShowReq struct {
	g.Meta `path:"/pmsRoomType/status" method:"post" tags:"ADMIN_PMS" summary:"房型_状态"`
	input_hotel.PmsRoomTypeIsShowInp
}

type RoomTypeIsShowRes struct{}
