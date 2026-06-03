package pms

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type AppReservationLastOrderReq struct {
	g.Meta   `path:"/pmsAppReservation/lastOrder" method:"get" tags:"ADMIN_PMS" summary:"APP预定_获取最近的订单号"`
	MemberId int64 `v:"required#会员ID不能为空" dc:"会员ID"`
}

type AppReservationLastOrderRes struct {
	OrderSn string `json:"orderSn" dc:"订单编号"`
}

type AppReservationListReq struct {
	g.Meta `path:"/pmsAppReservation/list" method:"get" tags:"ADMIN_PMS" summary:"APP预定_列表"`
	input_hotel.PmsAppReservationListInp
}

type AppReservationListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsAppReservationListModel `json:"list"   dc:"数据列表"`
}

type AppReservationRoomListReq struct {
	g.Meta `path:"/pmsAppReservation/roomList" method:"get" tags:"ADMIN_PMS" summary:"APP预定_房型列表"`
	input_hotel.PmsAppReservationRoomListInp
}

type AppReservationRoomListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsAppReservationRoomListModel `json:"list"   dc:"数据列表"`
}

type AppReservationRebateListReq struct {
	g.Meta `path:"/pmsAppReservation/rebateList" method:"get" tags:"ADMIN_PMS" summary:"APP预定_结算列表"`
	input_hotel.PmsAppReservationRebateListInp
}

type AppReservationRebateListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsAppReservationRebateListModel `json:"list"   dc:"结算数据列表"`
}

type AppReservationExportReq struct {
	g.Meta `path:"/pmsAppReservation/export" method:"get" tags:"ADMIN_PMS" summary:"APP预定_导出"`
	input_hotel.PmsAppReservationExportListInp
}

type AppReservationExportRes struct{}

type AppReservationViewReq struct {
	g.Meta `path:"/pmsAppReservation/view" method:"get" tags:"ADMIN_PMS" summary:"APP预定_详情"`
	input_hotel.PmsAppReservationViewInp
}

type AppReservationViewRes struct {
	*input_hotel.PmsAppReservationViewModel
}

type AppReservationEditReq struct {
	g.Meta `path:"/pmsAppReservation/edit" method:"post" tags:"ADMIN_PMS" summary:"APP预定_编辑"`
	input_hotel.PmsAppReservationEditInp
}

type AppReservationEditRes struct{}

type AppReservationDeleteReq struct {
	g.Meta `path:"/pmsAppReservation/delete" method:"post" tags:"ADMIN_PMS" summary:"APP预定_删除"`
	input_hotel.PmsAppReservationDeleteInp
}

type AppReservationDeleteRes struct{}

type AppReservationStatusReq struct {
	g.Meta `path:"/pmsAppReservation/status" method:"post" tags:"ADMIN_PMS" summary:"APP预定_更新状态"`
	input_hotel.PmsAppReservationStatusInp
}

type AppReservationStatusRes struct{}

type AppReservationCancelReq struct {
	g.Meta  `path:"/pmsAppReservation/cancel" method:"post" tags:"ADMIN_PMS" summary:"APP预定_取消订单"`
	OrderSn string `v:"required#订单编号不能为空" dc:"订单编号"`
	Reason  string `v:"required#取消原因不能为空" dc:"取消原因"`
	Id      int    `v:"required#订单ID不能为空" dc:"订单ID"`
}

type AppReservationCancelRes struct{}

type ReservationRoomStatusReq struct {
	g.Meta `path:"/pms/roomStatus" method:"post" tags:"ADMIN_PMS" summary:"房态_数据"`
	input_hotel.RoomStatus
	Page  int `json:"page" v:"required#请输入页码" dc:"页码"`
	Limit int `json:"limit" v:"required#请输入页长" dc:"页长"`
}

type ReservationRoomUtilListy struct {
	entity.PmsRoomUnit
	RoomType *ReservationRoomUtilListRoomType `json:"room_type" orm:"with:uid=rt_uid"`
}

type ReservationRoomUtilListRoomType struct {
	Uid      string                           `json:"uid"          orm:"uid"           description:"第三方系统的ID"`
	Name     string                           `json:"name"         orm:"name"          description:"房型名称"`
	Puid     string                           `json:"puid"         orm:"puid"          description:"物业ID"`
	Property *ReservationRoomUtilListProperty `json:"property"`
}

type ReservationRoomUtilListProperty struct {
	Uid   string `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
	Name  string `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	Cover string `json:"cover"                orm:"cover"                   description:"封面"`
}

type ReservationRoomStatusRes struct {
	ShowData    map[string]map[string]*gvar.Var
	Reservation []*struct {
		entity.PmsAppReservation
		CheckinUnix        int64
		CheckOutUnix       int64
		GuestProfileDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_guest_profile"`
			entity.PmsGuestProfile
		} `json:"guest_profile_detail" orm:"with:uid=main_guest"`
		RoomTypeDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_type"`
			entity.PmsRoomType
		} `json:"room_type_detail" orm:"with:uid=room_type"`
		RoomUnitDetail *struct {
			gmeta.Meta `orm:"table:hg_pms_room_unit"`
			entity.PmsRoomType
		} `json:"room_unit_detail" orm:"with:uid=room_unit"`
	}
	ReservationDict entity.PmsAppReservation
	RoomList        []*ReservationRoomUtilListy
	Availabilities  map[string]map[string][]*entity.PmsAvailabilities
	Count           int
}

type AppReservationReferrerListReq struct {
	g.Meta `path:"/pmsAppReservation/referrerList" method:"get" tags:"ADMIN_PMS" summary:"APP预定_列表"`
	input_hotel.PmsAppReservationReferrerListInp
}

type AppReservationReferrerListRes struct {
	input_form.PageRes
	List []*input_hotel.PmsAppReservationReferrerListModel `json:"list"   dc:"数据列表"`
}

type AppStayInfoReq struct {
	g.Meta `path:"/pmsAppReservation/singleView" method:"get" tags:"ADMIN_PMS" summary:"APP预定_订单信息"`
	input_hotel.PmsAppReservationViewInp
}

type AppStayInfoRes struct {
	*input_hotel.PmsAppStayInfoModel
}
