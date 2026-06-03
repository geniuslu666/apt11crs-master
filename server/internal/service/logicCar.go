// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysCarAddress interface {
		// Model 接送机地址ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取接送机地点管理列表
		List(ctx context.Context, in *input_car.CarAddressListInp) (list []*input_car.CarAddressListModel, totalCount int, err error)
		// Export 导出接送机地点管理
		Export(ctx context.Context, in *input_car.CarAddressListInp) (err error)
		// Edit 修改/新增接送机地点管理
		Edit(ctx context.Context, in *input_car.CarAddressEditInp) (err error)
		// Delete 删除接送机地点管理
		Delete(ctx context.Context, in *input_car.CarAddressDeleteInp) (err error)
		// View 获取接送机地点管理指定信息
		View(ctx context.Context, in *input_car.CarAddressViewInp) (res *input_car.CarAddressViewModel, err error)
		// Status 更新接送机地点管理状态
		Status(ctx context.Context, in *input_car.CarAddressStatusInp) (err error)
	}
	ICarAddressType interface {
		// Model 接送机地址类型ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取接送机地址类型列表
		List(ctx context.Context, in *input_car.CarAddressTypeListInp) (list []*input_car.CarAddressTypeListModel, totalCount int, err error)
	}
	ICarBanner interface {
		// Model BannerORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Banner列表
		List(ctx context.Context, in *input_car.CarBannerListInp) (list []*input_car.CarBannerListModel, totalCount int, err error)
		// Edit 修改/新增Banner
		Edit(ctx context.Context, in *input_car.CarBannerEditInp) (err error)
		// Delete 删除Banner
		Delete(ctx context.Context, in *input_car.CarBannerDeleteInp) (err error)
		// MaxSort 获取Banner最大排序
		MaxSort(ctx context.Context, in *input_car.CarBannerMaxSortInp) (res *input_car.CarBannerMaxSortModel, err error)
		// View 获取Banner指定信息
		View(ctx context.Context, in *input_car.CarBannerViewInp) (res *input_car.CarBannerViewModel, err error)
		// Status 更新Banner状态
		Status(ctx context.Context, in *input_car.CarBannerStatusInp) (err error)
	}
	ICarCar interface {
		// Model 车辆ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取车辆列表
		List(ctx context.Context, in *input_car.CarCarListInp) (list []*input_car.CarCarListModel, totalCount int, err error)
		// Edit 修改/新增车辆
		Edit(ctx context.Context, in *input_car.CarCarEditInp) (err error)
		// Delete 删除车辆
		Delete(ctx context.Context, in *input_car.CarCarDeleteInp) (err error)
		// MaxSort 获取车辆最大排序
		MaxSort(ctx context.Context, in *input_car.CarCarMaxSortInp) (res *input_car.CarCarMaxSortModel, err error)
		// View 获取车辆指定信息
		View(ctx context.Context, in *input_car.CarCarViewInp) (res *input_car.CarCarViewModel, err error)
		// Status 更新车辆状态
		Status(ctx context.Context, in *input_car.CarCarStatusInp) (err error)
		// WorkStatus 更新工作状态
		WorkStatus(ctx context.Context, in *input_car.CarCarWorkStatusInp) (err error)
	}
	ICarCarType interface {
		// Model 车辆车型ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取车辆车型列表
		List(ctx context.Context, in *input_car.CarCarTypeListInp) (list []*input_car.CarCarTypeListModel, totalCount int, err error)
		// Edit 修改/新增车辆车型
		Edit(ctx context.Context, in *input_car.CarCarTypeEditInp) (err error)
		// Delete 删除车辆车型
		Delete(ctx context.Context, in *input_car.CarCarTypeDeleteInp) (err error)
		// MaxSort 获取车辆车型最大排序
		MaxSort(ctx context.Context, in *input_car.CarCarTypeMaxSortInp) (res *input_car.CarCarTypeMaxSortModel, err error)
		// View 获取车辆车型指定信息
		View(ctx context.Context, in *input_car.CarCarTypeViewInp) (res *input_car.CarCarTypeViewModel, err error)
		// Status 更新车辆车型状态
		Status(ctx context.Context, in *input_car.CarCarTypeStatusInp) (err error)
	}
	ICarCooperateType interface {
		// Model 接送机营业类型ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取接送机营业类型列表
		List(ctx context.Context, in *input_car.CarCooperateTypeListInp) (list []*input_car.CarCooperateTypeListModel, totalCount int, err error)
		// Export 导出接送机营业类型
		Export(ctx context.Context, in *input_car.CarCooperateTypeListInp) (err error)
		// Edit 修改/新增接送机营业类型
		Edit(ctx context.Context, in *input_car.CarCooperateTypeEditInp) (err error)
		// Delete 删除接送机营业类型
		Delete(ctx context.Context, in *input_car.CarCooperateTypeDeleteInp) (err error)
		// View 获取接送机营业类型指定信息
		View(ctx context.Context, in *input_car.CarCooperateTypeViewInp) (res *input_car.CarCooperateTypeViewModel, err error)
		// Status 更新接送机营业类型状态
		Status(ctx context.Context, in *input_car.CarCooperateTypeStatusInp) (err error)
	}
	ICarDriver interface {
		// Model 司机管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取司机管理列表
		List(ctx context.Context, in *input_car.CarDriverListInp) (list []*input_car.CarDriverListModel, totalCount int, err error)
		// All 获取司机
		All(ctx context.Context, in *input_car.CarDriverListInp) (list []*input_car.CarDriverAllListModel, err error)
		// Edit 修改/新增司机管理
		Edit(ctx context.Context, in *input_car.CarDriverEditInp) (err error)
		// Delete 删除司机管理
		Delete(ctx context.Context, in *input_car.CarDriverDeleteInp) (err error)
		// View 获取司机管理指定信息
		View(ctx context.Context, in *input_car.CarDriverViewInp) (res *input_car.CarDriverViewModel, err error)
		// Status 更新司机状态
		Status(ctx context.Context, in *input_car.CarDriverStatusInp) (err error)
		// WorkStatus 更新司机状态
		WorkStatus(ctx context.Context, in *input_car.CarDriverWorkStatusInp) (err error)
		// GetIds 获取获取司机的id
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		// Bind 绑定用户
		Bind(ctx context.Context, in *input_car.CarDriverBindInp) (err error)
		// Unbind 解绑用户
		Unbind(ctx context.Context, in *input_car.CarDriverUnbindInp) (err error)
		// BindCar 绑定车辆
		BindCar(ctx context.Context, in *input_car.CarDriverBindCarInp) (err error)
	}
	ICarDriverWithdraw interface {
		// Model 司机提现管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取司机提现管理列表
		List(ctx context.Context, in *input_car.DriverWithdrawListInp) (list []*input_car.DriverWithdrawListModel, totalCount int, err error)
		View(ctx context.Context, in *input_car.DriverWithdrawViewInp) (res *input_car.DriverWithdrawViewModel, err error)
		Agree(ctx context.Context, in *input_car.DriverWithdrawAgreeInp) (err error)
		Disagree(ctx context.Context, in *input_car.DriverWithdrawDisagreeInp) (err error)
		Transfer(ctx context.Context, in *input_car.DriverWithdrawTransferInp) (err error)
	}
	ICarHelp interface {
		// Model HelpORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Help列表
		List(ctx context.Context, in *input_car.CarHelpListInp) (list []*input_car.CarHelpListModel, totalCount int, err error)
		// Edit 修改/新增Help
		Edit(ctx context.Context, in *input_car.CarHelpEditInp) (err error)
		// Delete 删除Help
		Delete(ctx context.Context, in *input_car.CarHelpDeleteInp) (err error)
		// View 获取Help指定信息
		View(ctx context.Context, in *input_car.CarHelpViewInp) (res *input_car.CarHelpViewModel, err error)
		// LanguageList 获取HelpLanguage列表
		LanguageList(ctx context.Context, in *input_car.CarHelpLanguageListInp) (list []*input_car.CarHelpLanguageListModel, err error)
	}
	ICarMaintenance interface {
		// Model MaintenanceORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Maintenance列表
		List(ctx context.Context, in *input_car.CarMaintenanceListInp) (list []*input_car.CarMaintenanceListModel, totalCount int, err error)
		// Edit 修改/新增Maintenance
		Edit(ctx context.Context, in *input_car.CarMaintenanceEditInp) (err error)
		// Delete 删除Maintenance
		Delete(ctx context.Context, in *input_car.CarMaintenanceDeleteInp) (err error)
		// View 获取Maintenance指定信息
		View(ctx context.Context, in *input_car.CarMaintenanceViewInp) (res *input_car.CarMaintenanceViewModel, err error)
		// LanguageList 获取MaintenanceLanguage列表
		LanguageList(ctx context.Context, in *input_car.CarMaintenanceLanguageListInp) (list []*input_car.CarMaintenanceLanguageListModel, err error)
	}
	ICarOrder interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_car.CarOrderListInp) (list []*input_car.CarOrderListModel, totalCount int, err error)
		View(ctx context.Context, in *input_car.CarOrderViewInp) (res *input_car.CarOrderViewModel, err error)
		ConfirmAgree(ctx context.Context, in *input_car.CarOrderConfirmAgreeInp) (err error)
		// ConfirmDisagree 确认失败，全额退款
		ConfirmDisagree(ctx context.Context, in *input_car.CarOrderConfirmDisagreeInp) (err error)
		DriverList(ctx context.Context, in *input_car.CarOrderDriverInp) (list []*input_car.CarOrderDriverModel, totalCount int, err error)
		Dispatch(ctx context.Context, in *input_car.CarOrderDispatchInp) (err error)
		TransferOrder(ctx context.Context, in *input_car.CarOrderTransferInp) (err error)
		// SettleOrderList 结算订单列表
		SettleOrderList(ctx context.Context, in *input_car.SettleCarOrderListInp) (list []*input_car.SettleCarOrderListModel, totalCount int, err error)
		// Refund 订单退款
		Refund(ctx context.Context, in *input_car.CarOrderRefundInp) (err error)
		GoOut(ctx context.Context, in *input_car.CarOrderGoOutInp) (err error)
		StartService(ctx context.Context, in *input_car.CarOrderStartServiceInp) (err error)
		EndService(ctx context.Context, in *input_car.CarOrderEndServiceInp) (err error)
		Abnormal(ctx context.Context, in *input_car.CarOrderAbnormalInp) (err error)
		ExportOrder(ctx context.Context, in *input_car.CarOrderExportInp) (err error)
		StartExport(ctx context.Context, in *input_car.CarOrderExportInp) (path string, err error)
		ExportList(ctx context.Context, in *input_car.CarOrderExportListInp) (list []*input_car.CarOrderExportListModel, totalCount int, err error)
	}
	ICarOrderLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_car.CarOrderLogListInp) (list []*input_car.CarOrderLogListModel, totalCount int, err error)
	}
	ICarService interface {
		// Model 服务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务列表
		List(ctx context.Context, in *input_car.CarServiceListInp) (list []*input_car.CarServiceListModel, totalCount int, err error)
		// Edit 修改/新增服务
		Edit(ctx context.Context, in *input_car.CarServiceEditInp) (err error)
		// Delete 删除服务
		Delete(ctx context.Context, in *input_car.CarServiceDeleteInp) (err error)
		// MaxSort 获取服务最大排序
		MaxSort(ctx context.Context, in *input_car.CarServiceMaxSortInp) (res *input_car.CarServiceMaxSortModel, err error)
		// View 获取服务指定信息
		View(ctx context.Context, in *input_car.CarServiceViewInp) (res *input_car.CarServiceViewModel, err error)
		// Status 更新服务状态
		Status(ctx context.Context, in *input_car.CarServiceStatusInp) (err error)
		// ServiceSort 编辑服务排序
		ServiceSort(ctx context.Context, in *input_car.CarServiceSortInp) (err error)
	}
	ICarSettlement interface {
		// Model 结算模式ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取结算模式列表
		List(ctx context.Context, in *input_car.CarSettlementListInp) (list []*input_car.CarSettlementListModel, totalCount int, err error)
		// Edit 修改/新增结算模式
		Edit(ctx context.Context, in *input_car.CarSettlementEditInp) (err error)
		// Delete 删除结算模式
		Delete(ctx context.Context, in *input_car.CarSettlementDeleteInp) (err error)
		// View 获取结算模式指定信息
		View(ctx context.Context, in *input_car.CarSettlementViewInp) (res *input_car.CarSettlementViewModel, err error)
		// Status 更新结算模式状态
		Status(ctx context.Context, in *input_car.CarSettlementStatusInp) (err error)
	}
	ICarSettlementOrder interface {
		// Model 结算模式ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// DailySettlement 每日结算
		DailySettlement(ctx context.Context) (err error)
		// WeekSettlement 每周结算
		WeekSettlement(ctx context.Context) (err error)
		// MonthSettlement 每月结算
		MonthSettlement(ctx context.Context) (err error)
		// SettlementOrderStat 结算概况
		SettlementOrderStat(ctx context.Context, in *input_car.CarSettlementOrderStatInp) (res *input_car.CarSettlementOrderStatModel, err error)
		// List 结算列表
		List(ctx context.Context, in *input_car.CarSettlementOrderListInp) (list []*input_car.CarSettlementOrderListModel, totalCount int, err error)
		// View 结算单详情
		View(ctx context.Context, in *input_car.CarSettlementOrderViewInp) (res *input_car.CarSettlementOrderViewModel, err error)
		// Verify 核账
		Verify(ctx context.Context, in *input_car.CarSettlementOrderVerifyInp) (err error)
	}
)

var (
	localSysCarAddress      ISysCarAddress
	localCarAddressType     ICarAddressType
	localCarBanner          ICarBanner
	localCarCar             ICarCar
	localCarCarType         ICarCarType
	localCarCooperateType   ICarCooperateType
	localCarDriver          ICarDriver
	localCarDriverWithdraw  ICarDriverWithdraw
	localCarHelp            ICarHelp
	localCarMaintenance     ICarMaintenance
	localCarOrder           ICarOrder
	localCarOrderLog        ICarOrderLog
	localCarService         ICarService
	localCarSettlement      ICarSettlement
	localCarSettlementOrder ICarSettlementOrder
)

func SysCarAddress() ISysCarAddress {
	if localSysCarAddress == nil {
		panic("implement not found for interface ISysCarAddress, forgot register?")
	}
	return localSysCarAddress
}

func RegisterSysCarAddress(i ISysCarAddress) {
	localSysCarAddress = i
}

func CarAddressType() ICarAddressType {
	if localCarAddressType == nil {
		panic("implement not found for interface ICarAddressType, forgot register?")
	}
	return localCarAddressType
}

func RegisterCarAddressType(i ICarAddressType) {
	localCarAddressType = i
}

func CarBanner() ICarBanner {
	if localCarBanner == nil {
		panic("implement not found for interface ICarBanner, forgot register?")
	}
	return localCarBanner
}

func RegisterCarBanner(i ICarBanner) {
	localCarBanner = i
}

func CarCar() ICarCar {
	if localCarCar == nil {
		panic("implement not found for interface ICarCar, forgot register?")
	}
	return localCarCar
}

func RegisterCarCar(i ICarCar) {
	localCarCar = i
}

func CarCarType() ICarCarType {
	if localCarCarType == nil {
		panic("implement not found for interface ICarCarType, forgot register?")
	}
	return localCarCarType
}

func RegisterCarCarType(i ICarCarType) {
	localCarCarType = i
}

func CarCooperateType() ICarCooperateType {
	if localCarCooperateType == nil {
		panic("implement not found for interface ICarCooperateType, forgot register?")
	}
	return localCarCooperateType
}

func RegisterCarCooperateType(i ICarCooperateType) {
	localCarCooperateType = i
}

func CarDriver() ICarDriver {
	if localCarDriver == nil {
		panic("implement not found for interface ICarDriver, forgot register?")
	}
	return localCarDriver
}

func RegisterCarDriver(i ICarDriver) {
	localCarDriver = i
}

func CarDriverWithdraw() ICarDriverWithdraw {
	if localCarDriverWithdraw == nil {
		panic("implement not found for interface ICarDriverWithdraw, forgot register?")
	}
	return localCarDriverWithdraw
}

func RegisterCarDriverWithdraw(i ICarDriverWithdraw) {
	localCarDriverWithdraw = i
}

func CarHelp() ICarHelp {
	if localCarHelp == nil {
		panic("implement not found for interface ICarHelp, forgot register?")
	}
	return localCarHelp
}

func RegisterCarHelp(i ICarHelp) {
	localCarHelp = i
}

func CarMaintenance() ICarMaintenance {
	if localCarMaintenance == nil {
		panic("implement not found for interface ICarMaintenance, forgot register?")
	}
	return localCarMaintenance
}

func RegisterCarMaintenance(i ICarMaintenance) {
	localCarMaintenance = i
}

func CarOrder() ICarOrder {
	if localCarOrder == nil {
		panic("implement not found for interface ICarOrder, forgot register?")
	}
	return localCarOrder
}

func RegisterCarOrder(i ICarOrder) {
	localCarOrder = i
}

func CarOrderLog() ICarOrderLog {
	if localCarOrderLog == nil {
		panic("implement not found for interface ICarOrderLog, forgot register?")
	}
	return localCarOrderLog
}

func RegisterCarOrderLog(i ICarOrderLog) {
	localCarOrderLog = i
}

func CarService() ICarService {
	if localCarService == nil {
		panic("implement not found for interface ICarService, forgot register?")
	}
	return localCarService
}

func RegisterCarService(i ICarService) {
	localCarService = i
}

func CarSettlement() ICarSettlement {
	if localCarSettlement == nil {
		panic("implement not found for interface ICarSettlement, forgot register?")
	}
	return localCarSettlement
}

func RegisterCarSettlement(i ICarSettlement) {
	localCarSettlement = i
}

func CarSettlementOrder() ICarSettlementOrder {
	if localCarSettlementOrder == nil {
		panic("implement not found for interface ICarSettlementOrder, forgot register?")
	}
	return localCarSettlementOrder
}

func RegisterCarSettlementOrder(i ICarSettlementOrder) {
	localCarSettlementOrder = i
}
