// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_pay"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IHotelService interface {
		// CheckHotelInventory 校验酒店库存
		CheckHotelInventory(ctx context.Context, in *input_hotel.CheckHotelInventoryInp) (err error)
		CancelRateList(ctx context.Context, in *input_hotel.PmsCancelRateListInp) (list []*input_hotel.PmsCancelRateListModel, totalCount int, err error)
		CancelRateEdit(ctx context.Context, in *input_hotel.PmsCancelRateEditInp) (err error)
		CancelRateDelete(ctx context.Context, in *input_hotel.PmsCancelRateDeleteInp) (err error)
		CancelRateMaxSort(ctx context.Context, in *input_hotel.PmsCancelRateMaxSortInp) (res *input_hotel.PmsCancelRateMaxSortModel, err error)
		CancelRateView(ctx context.Context, in *input_hotel.PmsCancelRateViewInp) (res *input_hotel.PmsCancelRateViewModel, err error)
		GuestList(ctx context.Context, in *input_hotel.PmsGuestProfileListInp) (list []*input_hotel.PmsGuestProfileListModel, totalCount int, err error)
		GuestExport(ctx context.Context, in *input_hotel.PmsGuestProfileListInp) (err error)
		GuestEdit(ctx context.Context, in *input_hotel.PmsGuestProfileEditInp) (err error)
		GuestDelete(ctx context.Context, in *input_hotel.PmsGuestProfileDeleteInp) (err error)
		GuestView(ctx context.Context, in *input_hotel.PmsGuestProfileViewInp) (res *input_hotel.PmsGuestProfileViewModel, err error)
		HotelOrderAward(ctx context.Context, OrderId int, CheckInDate string, MemberId int) (err error)
		HotelOrderAwardInvalid(ctx context.Context, OrderId int) (err error)
		// OrderChangeCheckIsExist 查询变更单是否存在
		OrderChangeCheckIsExist(ctx context.Context, OrderId int, ChangeType string) (err error)
		// OrderChangeList 获取入住订单变更信息表列表
		OrderChangeList(ctx context.Context, in *input_hotel.OrderChangeListInp) (list []*input_hotel.OrderChangeListModel, totalCount int, err error)
		// OrderChangeExport 导出入住订单变更信息表
		OrderChangeExport(ctx context.Context, in *input_hotel.OrderChangeListInp) (err error)
		// OrderChangeEdit 修改/新增入住订单变更信息表
		OrderChangeEdit(ctx context.Context, in *input_hotel.OrderChangeEditInp) (err error)
		// OrderChangeDelete 删除入住订单变更信息表
		OrderChangeDelete(ctx context.Context, in *input_hotel.OrderChangeDeleteInp) (err error)
		// OrderChangeView 获取入住订单变更信息表指定信息
		OrderChangeView(ctx context.Context, in *input_hotel.OrderChangeViewInp) (res *input_hotel.OrderChangeViewModel, err error)
		// OrderChangeGuestEdit 订单变更用户信息
		OrderChangeGuestEdit(ctx context.Context, in *input_hotel.OrderChangeGuestReq) (err error)
		// CreateHotelOrder 创建酒店订单
		CreateHotelOrder(ctx context.Context, in *input_hotel.CreateOrderInp) (out *input_hotel.CreateOrderModel, err error)
		// CreateStayParams 创建预订参数
		CreateStayParams(ctx context.Context, OrderSn string) (out *airhousePublicApi.CreateStayJSONDataRequest, err error)
		// CreateStay 提交预订
		CreateStay(ctx context.Context, in *airhousePublicApi.CreateStayJSONDataRequest, tx gdb.TX) (BookerResponse *airhousePublicApi.CreateStayJSONDataResponse, err error)
		// PreOrder 预下单
		PreOrder(ctx context.Context, in *input_hotel.PreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error)
		// PrePricePlanOrder 价格计划预下单
		PrePricePlanOrder(ctx context.Context, in *input_hotel.PreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error)
		// PreMorePricePlanOrder 多房间价格计划预下单
		PreMorePricePlanOrder(ctx context.Context, in *input_hotel.PreMoreCreateOrderInp) (out *input_hotel.PreCreateOrderModel, err error)
		// PreOrderDetail 预下单详情
		PreOrderDetail(ctx context.Context, in *input_hotel.PreOrderDetailInp) (out *input_hotel.PreOrderDetailModel, err error)
		// PrePayInfo 预下单重新计算支付金额支付模式
		PrePayInfo(ctx context.Context, in *input_pay.PrePayInfoInp) (out *input_pay.PrePayInfoModel, err error)
		// OrderExpiration 订单过期
		OrderExpiration(ctx context.Context, OrderSn string) (err error)
		// PreRefundOrderDetail 预退款订单详情
		PreRefundOrderDetail(ctx context.Context, in *input_hotel.PreRefundIn) (out *input_hotel.PreRefundOut, err error)
		// RefundOrderDetail 退款订单详情
		RefundOrderDetail(ctx context.Context, in *input_hotel.RefundDetailInp) (out *input_hotel.RefundDetailModel, err error)
		// SyncAppStay 同步APP订单
		SyncAppStay(ctx context.Context, StayData *airhousePublicApi.StayData) (err error)
		// SyncAppRoomReservations 同步APP入住单
		SyncAppRoomReservations(ctx context.Context, RoomReservation *airhousePublicApi.RoomReservations, OrderSn interface{}, SourceCode interface{}, SourceName interface{}) (err error)
		// SyncOtaStay 同步OTA订单
		SyncOtaStay(ctx context.Context, StayData *airhousePublicApi.StayData) (err error)
		// SyncOtaRoomReservations 同步OTA订单
		SyncOtaRoomReservations(ctx context.Context, RoomReservation *airhousePublicApi.RoomReservations, OrderSn interface{}, SourceCode interface{}, SourceName interface{}) (err error)
		// SyncOtaReservation 同步OTA订单
		SyncOtaReservation(ctx context.Context, RoomReservation *entity.PmsAppReservation) (err error)
		// SyncGuest 同步入住人信息
		SyncGuest(ctx context.Context, Uid string, PmsGuest *entity.PmsGuestProfile, IsApp bool) (err error)
		// SyncCharge 同步账单
		SyncCharge(ctx context.Context, PmsCharge []*entity.PmsCharge) (err error)
		HotelPricePlanList(ctx context.Context, ipt *input_hotel.SearchHotelPricePlanInp) (list []*input_hotel.SearchHotelPricePlanRes, count int, err error)
		HotelPricePlanEdit(ctx context.Context, ipt *input_hotel.EditHotelPricePlanInp) (err error)
		HotelPricePlanDelete(ctx context.Context, ipt *input_hotel.DeleteHotelPricePlanInp) (err error)
		HotelPricePlanRecycle(ctx context.Context, in *input_hotel.DeleteHotelPricePlanInp) (err error)
		HotelPricePlanGet(ctx context.Context, ipt *input_hotel.FindOneHotelPricePlanInp) (res *input_hotel.FindOneHotelPricePlanRes, err error)
		HotelPricePlanStatus(ctx context.Context, in *input_hotel.PricePlanStatusInp) (err error)
		PropertyList(ctx context.Context, in *input_hotel.PmsPropertyListInp) (list []*input_hotel.PmsPropertyListModel, totalCount int, err error)
		PropertyAll(ctx context.Context, in *input_hotel.PmsPropertyAllInp) (list []*input_hotel.PmsPropertyAllModel, err error)
		PropertyAppIndexList(ctx context.Context, in *input_hotel.PmsPropertyListInp) (list []*input_hotel.PmsPropertyAppIndexModel, totalCount int, err error)
		PropertyExport(ctx context.Context, in *input_hotel.PmsPropertyListInp) (err error)
		PropertyEdit(ctx context.Context, in *input_hotel.PmsPropertyEditInp) (err error)
		PropertyEditGroup(ctx context.Context, in *input_hotel.PmsPropertyEditGroupInp) (err error)
		PropertyDelete(ctx context.Context, in *input_hotel.PmsPropertyDeleteInp) (err error)
		PropertyRecycle(ctx context.Context, in *input_hotel.PmsPropertyRecycleInp) (err error)
		PropertyView(ctx context.Context, in *input_hotel.PmsPropertyViewInp) (res *input_hotel.PmsPropertyViewModel, resLanguage *input_hotel.PmsPropertyViewLannguageModel, err error)
		PropertyGetUids(ctx context.Context, name []string) (uids []string, err error)
		// PropertySort 编辑物业排序
		PropertySort(ctx context.Context, in *input_hotel.PmsPropertySortInp) (err error)
		SwitchSpaCanOrder(ctx context.Context, in *input_hotel.PropertySwitchSpaCanOrderInp) (err error)
		UpdateAccessPass(ctx context.Context, in *input_hotel.PropertyUpdateAccessPassInp) (err error)
		RegionList(ctx context.Context, in *input_hotel.PropertyRegionListInp) (list []*input_hotel.PropertyRegionListModel, totalCount int, err error)
		ReservationList(ctx context.Context, in *input_hotel.PmsAppReservationListInp) (list []*input_hotel.PmsAppReservationListModel, totalCount int, err error)
		ReservationRoomList(ctx context.Context, in *input_hotel.PmsAppReservationRoomListInp) (list []*input_hotel.PmsAppReservationRoomListModel, totalCount int, orderCount int, err error)
		ReservationRebateList(ctx context.Context, in *input_hotel.PmsAppReservationRebateListInp) (list []*input_hotel.PmsAppReservationRebateListModel, totalCount int, err error)
		ReservationExport(ctx context.Context, in *input_hotel.PmsAppReservationExportListInp) (err error)
		ReservationEdit(ctx context.Context, in *input_hotel.PmsAppReservationEditInp) (err error)
		ReservationDelete(ctx context.Context, in *input_hotel.PmsAppReservationDeleteInp) (err error)
		ReservationView(ctx context.Context, in *input_hotel.PmsAppReservationViewInp) (res *input_hotel.PmsAppReservationViewModel, err error)
		ReservationStatus(ctx context.Context, in *input_hotel.PmsAppReservationStatusInp) (err error)
		ReservationReferrerList(ctx context.Context, in *input_hotel.PmsAppReservationReferrerListInp) (list []*input_hotel.PmsAppReservationReferrerListModel, totalCount int, err error)
		AppStayInfo(ctx context.Context, in *input_hotel.PmsAppReservationViewInp) (res *input_hotel.PmsAppStayInfoModel, err error)
		RoomTypeList(ctx context.Context, in *input_hotel.PmsRoomTypeListInp) (list []*input_hotel.PmsRoomTypeListModel, totalCount int, err error)
		RoomTypeEdit(ctx context.Context, in *input_hotel.PmsRoomTypeEditInp) (err error)
		RoomTypeDelete(ctx context.Context, in *input_hotel.PmsRoomTypeDeleteInp) (err error)
		RoomTypeView(ctx context.Context, in *input_hotel.PmsRoomTypeViewInp) (res *input_hotel.PmsRoomTypeViewModel, err error)
		RoomTypeGetUidsByPuid(ctx context.Context, puid string) (uids []string, err error)
		RoomTypeStatus(ctx context.Context, in *input_hotel.PmsRoomTypeIsShowInp) (err error)
		RoomUnitList(ctx context.Context, in *input_hotel.PmsRoomUnitListInp) (list []*input_hotel.PmsRoomUnitListModel, totalCount int, err error)
		RoomUnitExport(ctx context.Context, in *input_hotel.PmsRoomUnitListInp) (err error)
		RoomUnitEdit(ctx context.Context, in *input_hotel.PmsRoomUnitEditInp) (err error)
		RoomUnitDelete(ctx context.Context, in *input_hotel.PmsRoomUnitDeleteInp) (err error)
		RoomUnitView(ctx context.Context, in *input_hotel.PmsRoomUnitViewInp) (res *input_hotel.PmsRoomUnitViewModel, err error)
		AppStayLogModel(ctx context.Context, option ...*handler.Option) *gdb.Model
		AppStayLogList(ctx context.Context, in *input_hotel.AppStayLogListInp) (list []*input_hotel.AppStayLogListModel, totalCount int, err error)
	}
)

var (
	localHotelService IHotelService
)

func HotelService() IHotelService {
	if localHotelService == nil {
		panic("implement not found for interface IHotelService, forgot register?")
	}
	return localHotelService
}

func RegisterHotelService(i IHotelService) {
	localHotelService = i
}
