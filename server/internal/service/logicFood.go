// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_th"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
)

type (
	IFoodActivity interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodActivityListInp) (list []*input_food.FoodActivityListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_food.FoodActivityListInp) (err error)
		Edit(ctx context.Context, in *input_food.FoodActivityEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodActivityDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodActivityViewInp) (res *input_food.FoodActivityViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodActivityStatusInp) (err error)
	}
	IFoodArea interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodAreaListInp) (res *input_food.FoodAreaListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodAreaEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodAreaDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodAreaViewInp) (res *input_food.FoodAreaViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodAreaStatusInp) (err error)
	}
	IFoodCooperateType interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodCooperateTypeListInp) (list []*input_food.FoodCooperateTypeListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_food.FoodCooperateTypeListInp) (err error)
		Edit(ctx context.Context, in *input_food.FoodCooperateTypeEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodCooperateTypeDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodCooperateTypeViewInp) (res *input_food.FoodCooperateTypeViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodCooperateTypeStatusInp) (err error)
	}
	IFoodCuisine interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodCuisineListInp) (list []*input_food.FoodCuisineListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodCuisineEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodCuisineDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodCuisineViewInp) (res *input_food.FoodCuisineViewModel, err error)
		MaxSort(ctx context.Context, in *input_food.FoodCuisineMaxSortInp) (res *input_food.FoodCuisineMaxSortModel, err error)
		Status(ctx context.Context, in *input_food.FoodCuisineStatusInp) (err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
	}
	IFoodGoods interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodGoodsListInp) (list []*input_food.FoodGoodsListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_food.FoodGoodsListInp) (err error)
		Edit(ctx context.Context, in *input_food.FoodGoodsEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodGoodsDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodGoodsViewInp) (res *input_food.FoodGoodsViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodGoodsStatusInp) (err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
	}
	IFoodLabel interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodLabelListInp) (list []*input_food.FoodLabelListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodLabelEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodLabelDeleteInp) (err error)
		MaxSort(ctx context.Context, in *input_food.FoodLabelMaxSortInp) (res *input_food.FoodLabelMaxSortModel, err error)
		View(ctx context.Context, in *input_food.FoodLabelViewInp) (res *input_food.FoodLabelViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodLabelStatusInp) (err error)
	}
	IFoodMaintenance interface {
		// Model MaintenanceORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Maintenance列表
		List(ctx context.Context, in *input_food.FoodMaintenanceListInp) (list []*input_food.FoodMaintenanceListModel, totalCount int, err error)
		// Edit 修改/新增Maintenance
		Edit(ctx context.Context, in *input_food.FoodMaintenanceEditInp) (err error)
		// Delete 删除Maintenance
		Delete(ctx context.Context, in *input_food.FoodMaintenanceDeleteInp) (err error)
		// View 获取Maintenance指定信息
		View(ctx context.Context, in *input_food.FoodMaintenanceViewInp) (res *input_food.FoodMaintenanceViewModel, err error)
		// LanguageList 获取MaintenanceLanguage列表
		LanguageList(ctx context.Context, in *input_food.FoodMaintenanceLanguageListInp) (list []*input_food.FoodMaintenanceLanguageListModel, err error)
	}
	IFoodOrder interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodOrderListInp) (list []*input_food.FoodOrderListModel, totalCount int, err error)
		View(ctx context.Context, in *input_food.FoodOrderViewInp) (res *input_food.FoodOrderViewModel, err error)
		ConfirmAgree(ctx context.Context, in *input_food.FoodOrderConfirmAgreeInp) (err error)
		// ConfirmDisagree 确认失败，全额退款
		ConfirmDisagree(ctx context.Context, in *input_food.FoodOrderConfirmDisagreeInp) (err error)
		// SettleOrderList 结算订单列表
		SettleOrderList(ctx context.Context, in *input_food.SettleFoodOrderListInp) (list []*input_food.SettleFoodOrderListModel, totalCount int, err error)
		// CancelPay 取消订单
		CancelPay(ctx context.Context, in *input_food.FoodOrderCancelPayInp) (err error)
		RefreshCode(ctx context.Context, in *input_food.FoodsOrderRefreshCodeInp) (code string, verifyStatus int, err error)
		Verify(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error)
		ExportOrder(ctx context.Context, in *input_food.FoodOrderExportInp) (err error)
		StartExport(ctx context.Context, in *input_food.FoodOrderExportInp) (path string, err error)
		ExportList(ctx context.Context, in *input_food.FoodOrderExportListInp) (list []*input_food.FoodOrderExportListModel, totalCount int, err error)
		SyncToretaNotifications(ctx context.Context, Logger *glog.Logger) (err error)
	}
	IFoodOrderLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodOrderLogListInp) (list []*input_food.FoodOrderLogListModel, totalCount int, err error)
	}
	IFoodRestaurant interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodRestaurantListInp) (list []*input_food.FoodRestaurantListModel, totalCount int, err error)
		All(ctx context.Context, in *input_food.FoodRestaurantListInp) (list []*input_food.FoodRestaurantAllListModel, err error)
		Edit(ctx context.Context, in *input_food.FoodRestaurantEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodRestaurantDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodRestaurantViewInp) (res *input_food.FoodRestaurantViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodRestaurantStatusInp) (err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		ResetVerifyCode(ctx context.Context, in *input_food.FoodRestaurantRestVerifyCodeInp) (err error)
		Switch(ctx context.Context, in *input_food.FoodRestaurantSwitchInp) (err error)
		// RestaurantSort 编辑餐厅排序
		RestaurantSort(ctx context.Context, in *input_food.FoodRestaurantSortInp) (err error)
	}
	IFoodRestaurantNotice interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodRestaurantNoticeListInp) (list []*input_food.FoodRestaurantNoticeListModel, totalCount int, err error)
		Export(ctx context.Context, in *input_food.FoodRestaurantNoticeListInp) (err error)
		Edit(ctx context.Context, in *input_food.FoodRestaurantNoticeEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodRestaurantNoticeDeleteInp) (err error)
		MaxSort(ctx context.Context, in *input_food.FoodRestaurantNoticeMaxSortInp) (res *input_food.FoodRestaurantNoticeMaxSortModel, err error)
		View(ctx context.Context, in *input_food.FoodRestaurantNoticeViewInp) (res *input_food.FoodRestaurantNoticeViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodRestaurantNoticeStatusInp) (err error)
	}
	IFoodSeat interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodSeatListInp) (list []*input_food.FoodSeatListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodSeatEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodSeatDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodSeatViewInp) (res *input_food.FoodSeatViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodSeatStatusInp) (err error)
	}
	IFoodSettlement interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodSettlementListInp) (list []*input_food.FoodSettlementListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodSettlementEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodSettlementDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodSettlementViewInp) (res *input_food.FoodSettlementViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodSettlementStatusInp) (err error)
	}
	IFoodSettlementAccount interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_food.FoodSettlementAccountListInp) (list []*input_food.FoodSettlementAccountListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_food.FoodSettlementAccountEditInp) (err error)
		Delete(ctx context.Context, in *input_food.FoodSettlementAccountDeleteInp) (err error)
		View(ctx context.Context, in *input_food.FoodSettlementAccountViewInp) (res *input_food.FoodSettlementAccountViewModel, err error)
		Status(ctx context.Context, in *input_food.FoodSettlementAccountStatusInp) (err error)
	}
	IFoodSettlementOrder interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// DailySettlement 每日结算
		DailySettlement(ctx context.Context) (err error)
		// WeekSettlement 每周结算
		WeekSettlement(ctx context.Context) (err error)
		// MonthSettlement 每月结算
		MonthSettlement(ctx context.Context) (err error)
		// SettlementOrderStat 结算概况
		SettlementOrderStat(ctx context.Context, in *input_food.FoodSettlementOrderStatInp) (res *input_food.FoodSettlementOrderStatModel, err error)
		// List 结算列表
		List(ctx context.Context, in *input_food.FoodSettlementOrderListInp) (list []*input_food.FoodSettlementOrderListModel, totalCount int, err error)
		// View 结算单详情
		View(ctx context.Context, in *input_food.FoodSettlementOrderViewInp) (res *input_food.FoodSettlementOrderViewModel, err error)
		// Verify 核账
		Verify(ctx context.Context, in *input_food.FoodSettlementOrderVerifyInp) (err error)
	}
)

var (
	localFoodActivity          IFoodActivity
	localFoodArea              IFoodArea
	localFoodCooperateType     IFoodCooperateType
	localFoodCuisine           IFoodCuisine
	localFoodGoods             IFoodGoods
	localFoodLabel             IFoodLabel
	localFoodMaintenance       IFoodMaintenance
	localFoodOrder             IFoodOrder
	localFoodOrderLog          IFoodOrderLog
	localFoodRestaurant        IFoodRestaurant
	localFoodRestaurantNotice  IFoodRestaurantNotice
	localFoodSeat              IFoodSeat
	localFoodSettlement        IFoodSettlement
	localFoodSettlementAccount IFoodSettlementAccount
	localFoodSettlementOrder   IFoodSettlementOrder
)

func FoodActivity() IFoodActivity {
	if localFoodActivity == nil {
		panic("implement not found for interface IFoodActivity, forgot register?")
	}
	return localFoodActivity
}

func RegisterFoodActivity(i IFoodActivity) {
	localFoodActivity = i
}

func FoodArea() IFoodArea {
	if localFoodArea == nil {
		panic("implement not found for interface IFoodArea, forgot register?")
	}
	return localFoodArea
}

func RegisterFoodArea(i IFoodArea) {
	localFoodArea = i
}

func FoodCooperateType() IFoodCooperateType {
	if localFoodCooperateType == nil {
		panic("implement not found for interface IFoodCooperateType, forgot register?")
	}
	return localFoodCooperateType
}

func RegisterFoodCooperateType(i IFoodCooperateType) {
	localFoodCooperateType = i
}

func FoodCuisine() IFoodCuisine {
	if localFoodCuisine == nil {
		panic("implement not found for interface IFoodCuisine, forgot register?")
	}
	return localFoodCuisine
}

func RegisterFoodCuisine(i IFoodCuisine) {
	localFoodCuisine = i
}

func FoodGoods() IFoodGoods {
	if localFoodGoods == nil {
		panic("implement not found for interface IFoodGoods, forgot register?")
	}
	return localFoodGoods
}

func RegisterFoodGoods(i IFoodGoods) {
	localFoodGoods = i
}

func FoodLabel() IFoodLabel {
	if localFoodLabel == nil {
		panic("implement not found for interface IFoodLabel, forgot register?")
	}
	return localFoodLabel
}

func RegisterFoodLabel(i IFoodLabel) {
	localFoodLabel = i
}

func FoodMaintenance() IFoodMaintenance {
	if localFoodMaintenance == nil {
		panic("implement not found for interface IFoodMaintenance, forgot register?")
	}
	return localFoodMaintenance
}

func RegisterFoodMaintenance(i IFoodMaintenance) {
	localFoodMaintenance = i
}

func FoodOrder() IFoodOrder {
	if localFoodOrder == nil {
		panic("implement not found for interface IFoodOrder, forgot register?")
	}
	return localFoodOrder
}

func RegisterFoodOrder(i IFoodOrder) {
	localFoodOrder = i
}

func FoodOrderLog() IFoodOrderLog {
	if localFoodOrderLog == nil {
		panic("implement not found for interface IFoodOrderLog, forgot register?")
	}
	return localFoodOrderLog
}

func RegisterFoodOrderLog(i IFoodOrderLog) {
	localFoodOrderLog = i
}

func FoodRestaurant() IFoodRestaurant {
	if localFoodRestaurant == nil {
		panic("implement not found for interface IFoodRestaurant, forgot register?")
	}
	return localFoodRestaurant
}

func RegisterFoodRestaurant(i IFoodRestaurant) {
	localFoodRestaurant = i
}

func FoodRestaurantNotice() IFoodRestaurantNotice {
	if localFoodRestaurantNotice == nil {
		panic("implement not found for interface IFoodRestaurantNotice, forgot register?")
	}
	return localFoodRestaurantNotice
}

func RegisterFoodRestaurantNotice(i IFoodRestaurantNotice) {
	localFoodRestaurantNotice = i
}

func FoodSeat() IFoodSeat {
	if localFoodSeat == nil {
		panic("implement not found for interface IFoodSeat, forgot register?")
	}
	return localFoodSeat
}

func RegisterFoodSeat(i IFoodSeat) {
	localFoodSeat = i
}

func FoodSettlement() IFoodSettlement {
	if localFoodSettlement == nil {
		panic("implement not found for interface IFoodSettlement, forgot register?")
	}
	return localFoodSettlement
}

func RegisterFoodSettlement(i IFoodSettlement) {
	localFoodSettlement = i
}

func FoodSettlementAccount() IFoodSettlementAccount {
	if localFoodSettlementAccount == nil {
		panic("implement not found for interface IFoodSettlementAccount, forgot register?")
	}
	return localFoodSettlementAccount
}

func RegisterFoodSettlementAccount(i IFoodSettlementAccount) {
	localFoodSettlementAccount = i
}

func FoodSettlementOrder() IFoodSettlementOrder {
	if localFoodSettlementOrder == nil {
		panic("implement not found for interface IFoodSettlementOrder, forgot register?")
	}
	return localFoodSettlementOrder
}

func RegisterFoodSettlementOrder(i IFoodSettlementOrder) {
	localFoodSettlementOrder = i
}
