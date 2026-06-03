// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_spa"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISpaBanner interface {
		// Model BannerORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Banner列表
		List(ctx context.Context, in *input_spa.SpaBannerListInp) (list []*input_spa.SpaBannerListModel, totalCount int, err error)
		// Edit 修改/新增Banner
		Edit(ctx context.Context, in *input_spa.SpaBannerEditInp) (err error)
		// Delete 删除Banner
		Delete(ctx context.Context, in *input_spa.SpaBannerDeleteInp) (err error)
		// MaxSort 获取Banner最大排序
		MaxSort(ctx context.Context, in *input_spa.SpaBannerMaxSortInp) (res *input_spa.SpaBannerMaxSortModel, err error)
		// View 获取Banner指定信息
		View(ctx context.Context, in *input_spa.SpaBannerViewInp) (res *input_spa.SpaBannerViewModel, err error)
		// Status 更新Banner状态
		Status(ctx context.Context, in *input_spa.SpaBannerStatusInp) (err error)
	}
	ISpaCooperateType interface {
		// Model 按摩营业类型ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取按摩营业类型列表
		List(ctx context.Context, in *input_spa.SpaCooperateTypeListInp) (list []*input_spa.SpaCooperateTypeListModel, totalCount int, err error)
		// Export 导出按摩营业类型
		Export(ctx context.Context, in *input_spa.SpaCooperateTypeListInp) (err error)
		// Edit 修改/新增按摩营业类型
		Edit(ctx context.Context, in *input_spa.SpaCooperateTypeEditInp) (err error)
		// Delete 删除按摩营业类型
		Delete(ctx context.Context, in *input_spa.SpaCooperateTypeDeleteInp) (err error)
		// View 获取按摩营业类型指定信息
		View(ctx context.Context, in *input_spa.SpaCooperateTypeViewInp) (res *input_spa.SpaCooperateTypeViewModel, err error)
		// Status 更新按摩营业类型状态
		Status(ctx context.Context, in *input_spa.SpaCooperateTypeStatusInp) (err error)
	}
	ISpaIsp interface {
		// Model 服务商管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务商管理列表
		List(ctx context.Context, in *input_spa.SpaIspListInp) (list []*input_spa.SpaIspListModel, totalCount int, err error)
		// All 获取服务商
		All(ctx context.Context, in *input_spa.SpaIspListInp) (list []*input_spa.SpaIspAllListModel, err error)
		// Edit 修改/新增服务商管理
		Edit(ctx context.Context, in *input_spa.SpaIspEditInp) (err error)
		// Delete 删除服务商管理
		Delete(ctx context.Context, in *input_spa.SpaIspDeleteInp) (err error)
		// View 获取服务商管理指定信息
		View(ctx context.Context, in *input_spa.SpaIspViewInp) (res *input_spa.SpaIspViewModel, err error)
		// Status 更新服务商状态
		Status(ctx context.Context, in *input_spa.SpaIspStatusInp) (err error)
		// WorkStatus 更新服务商状态
		WorkStatus(ctx context.Context, in *input_spa.SpaIspWorkStatusInp) (err error)
		// GetIds 获取获取服务商的id
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		// Bind 绑定用户
		Bind(ctx context.Context, in *input_spa.SpaIspBindInp) (err error)
		// Unbind 解绑用户
		Unbind(ctx context.Context, in *input_spa.SpaIspUnbindInp) (err error)
	}
	ISpaLabel interface {
		// Model 服务标签ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务标签列表
		List(ctx context.Context, in *input_spa.SpaLabelListInp) (list []*input_spa.SpaLabelListModel, totalCount int, err error)
		// Edit 修改/新增服务标签
		Edit(ctx context.Context, in *input_spa.SpaLabelEditInp) (err error)
		// Delete 删除服务标签
		Delete(ctx context.Context, in *input_spa.SpaLabelDeleteInp) (err error)
		// MaxSort 获取服务标签最大排序
		MaxSort(ctx context.Context, in *input_spa.SpaLabelMaxSortInp) (res *input_spa.SpaLabelMaxSortModel, err error)
		// View 获取服务标签指定信息
		View(ctx context.Context, in *input_spa.SpaLabelViewInp) (res *input_spa.SpaLabelViewModel, err error)
		// Status 更新服务标签状态
		Status(ctx context.Context, in *input_spa.SpaLabelStatusInp) (err error)
	}
	ISpaMaintenance interface {
		// Model MaintenanceORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取Maintenance列表
		List(ctx context.Context, in *input_spa.SpaMaintenanceListInp) (list []*input_spa.SpaMaintenanceListModel, totalCount int, err error)
		// Edit 修改/新增Maintenance
		Edit(ctx context.Context, in *input_spa.SpaMaintenanceEditInp) (err error)
		// Delete 删除Maintenance
		Delete(ctx context.Context, in *input_spa.SpaMaintenanceDeleteInp) (err error)
		// View 获取Maintenance指定信息
		View(ctx context.Context, in *input_spa.SpaMaintenanceViewInp) (res *input_spa.SpaMaintenanceViewModel, err error)
		// LanguageList 获取MaintenanceLanguage列表
		LanguageList(ctx context.Context, in *input_spa.SpaMaintenanceLanguageListInp) (list []*input_spa.SpaMaintenanceLanguageListModel, err error)
	}
	ISpaOrder interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_spa.SpaOrderListInp) (list []*input_spa.SpaOrderListModel, totalCount int, err error)
		View(ctx context.Context, in *input_spa.SpaOrderViewInp) (res *input_spa.SpaOrderViewModel, err error)
		ConfirmAgree(ctx context.Context, in *input_spa.SpaOrderConfirmAgreeInp) (err error)
		// ConfirmDisagree 确认失败，全额退款
		ConfirmDisagree(ctx context.Context, in *input_spa.SpaOrderConfirmDisagreeInp) (err error)
		TechnicianList(ctx context.Context, in *input_spa.SpaOrderTechnicianInp) (list []*input_spa.SpaOrderTechnicianModel, totalCount int, err error)
		Dispatch(ctx context.Context, in *input_spa.SpaOrderDispatchInp) (err error)
		// TechnicianOrderList 结算订单列表
		TechnicianOrderList(ctx context.Context, in *input_spa.SettleSpaOrderTechnicianListInp) (list []*input_spa.SettleSpaOrderTechnicianListModel, totalCount int, err error)
		GoOut(ctx context.Context, in *input_spa.SpaOrderGoOutInp) (err error)
		Arrive(ctx context.Context, in *input_spa.SpaOrderGoOutInp) (err error)
		StartService(ctx context.Context, in *input_spa.SpaOrderStartServiceInp) (err error)
		EndService(ctx context.Context, in *input_spa.SpaOrderEndServiceInp) (err error)
		Abnormal(ctx context.Context, in *input_spa.SpaOrderAbnormalInp) (err error)
		// CancelPay 取消订单
		CancelPay(ctx context.Context, in *input_spa.SpaOrderCancelPayInp) (err error)
		ExportOrder(ctx context.Context, in *input_spa.SpaOrderExportInp) (err error)
		StartExport(ctx context.Context, in *input_spa.SpaOrderExportInp) (path string, err error)
		ExportList(ctx context.Context, in *input_spa.SpaOrderExportListInp) (list []*input_spa.SpaOrderExportListModel, totalCount int, err error)
	}
	ISpaOrderLog interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_spa.SpaOrderLogListInp) (list []*input_spa.SpaOrderLogListModel, totalCount int, err error)
	}
	ISpaOrderTechnician interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// GetOrderIds 根据技师ID、结算类型、结算周期、开始时间、结束时间获取订单ID
		GetOrderIds(ctx context.Context, technicianId int, ispId int, settlementType int, settlementCycle int, settlementObject string, StartTime string, EndTime string) (ids []int, err error)
		// GetTechnicianIds 根据订单ID获取技师ID
		GetTechnicianIds(ctx context.Context, orderId int) (ids []int, err error)
	}
	ISpaService interface {
		// Model 服务服务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务服务列表
		List(ctx context.Context, in *input_spa.SpaServiceListInp) (list []*input_spa.SpaServiceListModel, totalCount int, err error)
		All(ctx context.Context, in *input_spa.SpaServiceListInp) (list []*input_spa.SpaServiceAllListModel, err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		// Edit 修改/新增服务服务
		Edit(ctx context.Context, in *input_spa.SpaServiceEditInp) (err error)
		// Delete 删除按摩服务
		Delete(ctx context.Context, in *input_spa.SpaServiceDeleteInp) (err error)
		// MaxSort 获取服务服务最大排序
		MaxSort(ctx context.Context, in *input_spa.SpaServiceMaxSortInp) (res *input_spa.SpaServiceMaxSortModel, err error)
		// View 获取服务服务指定信息
		View(ctx context.Context, in *input_spa.SpaServiceViewInp) (res *input_spa.SpaServiceViewModel, err error)
		// Status 更新服务服务状态
		Status(ctx context.Context, in *input_spa.SpaServiceStatusInp) (err error)
		// Recycle 恢复按摩服务
		Recycle(ctx context.Context, in *input_spa.SpaServiceDeleteInp) (err error)
	}
	ISpaSettlement interface {
		// Model 结算模式ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取结算模式列表
		List(ctx context.Context, in *input_spa.SpaSettlementListInp) (list []*input_spa.SpaSettlementListModel, totalCount int, err error)
		// Edit 修改/新增结算模式
		Edit(ctx context.Context, in *input_spa.SpaSettlementEditInp) (err error)
		// Delete 删除结算模式
		Delete(ctx context.Context, in *input_spa.SpaSettlementDeleteInp) (err error)
		// View 获取结算模式指定信息
		View(ctx context.Context, in *input_spa.SpaSettlementViewInp) (res *input_spa.SpaSettlementViewModel, err error)
		// Status 更新结算模式状态
		Status(ctx context.Context, in *input_spa.SpaSettlementStatusInp) (err error)
	}
	ISpaSettlementOrder interface {
		// Model 结算模式ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// DailySettlement 每日结算
		DailySettlement(ctx context.Context) (err error)
		// WeekSettlement 每周结算
		WeekSettlement(ctx context.Context) (err error)
		// MonthSettlement 每月结算
		MonthSettlement(ctx context.Context) (err error)
		// SettlementOrderStat 结算概况
		SettlementOrderStat(ctx context.Context, in *input_spa.SpaSettlementOrderStatInp) (res *input_spa.SpaSettlementOrderStatModel, err error)
		// List 结算列表
		List(ctx context.Context, in *input_spa.SpaSettlementOrderListInp) (list []*input_spa.SpaSettlementOrderListModel, totalCount int, err error)
		// View 结算单详情
		View(ctx context.Context, in *input_spa.SpaSettlementOrderViewInp) (res *input_spa.SpaSettlementOrderViewModel, err error)
		// Verify 核账
		Verify(ctx context.Context, in *input_spa.SpaSettlementOrderVerifyInp) (err error)
	}
	ISpaStore interface {
		// Model 门店管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取门店管理列表
		List(ctx context.Context, in *input_spa.SpaStoreListInp) (list []*input_spa.SpaStoreListModel, totalCount int, err error)
		// Edit 修改/新增门店管理
		Edit(ctx context.Context, in *input_spa.SpaStoreEditInp) (err error)
		// Delete 删除门店管理
		Delete(ctx context.Context, in *input_spa.SpaStoreDeleteInp) (err error)
		// View 获取门店管理指定信息
		View(ctx context.Context, in *input_spa.SpaStoreViewInp) (res *input_spa.SpaStoreViewModel, err error)
		// Latest 获取门店管理最新信息
		Latest(ctx context.Context) (res *input_spa.SpaStoreViewModel, err error)
	}
	ISpaTechnician interface {
		// Model 技师管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取技师管理列表
		List(ctx context.Context, in *input_spa.SpaTechnicianListInp) (list []*input_spa.SpaTechnicianListModel, totalCount int, err error)
		// All 获取技师
		All(ctx context.Context, in *input_spa.SpaTechnicianListInp) (list []*input_spa.SpaTechnicianAllListModel, err error)
		// Edit 修改/新增技师管理
		Edit(ctx context.Context, in *input_spa.SpaTechnicianEditInp) (err error)
		// Delete 删除技师管理
		Delete(ctx context.Context, in *input_spa.SpaTechnicianDeleteInp) (err error)
		// View 获取技师管理指定信息
		View(ctx context.Context, in *input_spa.SpaTechnicianViewInp) (res *input_spa.SpaTechnicianViewModel, err error)
		// Status 更新技师状态
		Status(ctx context.Context, in *input_spa.SpaTechnicianStatusInp) (err error)
		// WorkStatus 更新技师状态
		WorkStatus(ctx context.Context, in *input_spa.SpaTechnicianWorkStatusInp) (err error)
		// GetIds 获取获取技师的id
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		// Bind 绑定用户
		Bind(ctx context.Context, in *input_spa.SpaTechnicianBindInp) (err error)
		// Unbind 解绑用户
		Unbind(ctx context.Context, in *input_spa.SpaTechnicianUnbindInp) (err error)
	}
)

var (
	localSpaBanner          ISpaBanner
	localSpaCooperateType   ISpaCooperateType
	localSpaIsp             ISpaIsp
	localSpaLabel           ISpaLabel
	localSpaMaintenance     ISpaMaintenance
	localSpaOrder           ISpaOrder
	localSpaOrderLog        ISpaOrderLog
	localSpaOrderTechnician ISpaOrderTechnician
	localSpaService         ISpaService
	localSpaSettlement      ISpaSettlement
	localSpaSettlementOrder ISpaSettlementOrder
	localSpaStore           ISpaStore
	localSpaTechnician      ISpaTechnician
)

func SpaBanner() ISpaBanner {
	if localSpaBanner == nil {
		panic("implement not found for interface ISpaBanner, forgot register?")
	}
	return localSpaBanner
}

func RegisterSpaBanner(i ISpaBanner) {
	localSpaBanner = i
}

func SpaCooperateType() ISpaCooperateType {
	if localSpaCooperateType == nil {
		panic("implement not found for interface ISpaCooperateType, forgot register?")
	}
	return localSpaCooperateType
}

func RegisterSpaCooperateType(i ISpaCooperateType) {
	localSpaCooperateType = i
}

func SpaIsp() ISpaIsp {
	if localSpaIsp == nil {
		panic("implement not found for interface ISpaIsp, forgot register?")
	}
	return localSpaIsp
}

func RegisterSpaIsp(i ISpaIsp) {
	localSpaIsp = i
}

func SpaLabel() ISpaLabel {
	if localSpaLabel == nil {
		panic("implement not found for interface ISpaLabel, forgot register?")
	}
	return localSpaLabel
}

func RegisterSpaLabel(i ISpaLabel) {
	localSpaLabel = i
}

func SpaMaintenance() ISpaMaintenance {
	if localSpaMaintenance == nil {
		panic("implement not found for interface ISpaMaintenance, forgot register?")
	}
	return localSpaMaintenance
}

func RegisterSpaMaintenance(i ISpaMaintenance) {
	localSpaMaintenance = i
}

func SpaOrder() ISpaOrder {
	if localSpaOrder == nil {
		panic("implement not found for interface ISpaOrder, forgot register?")
	}
	return localSpaOrder
}

func RegisterSpaOrder(i ISpaOrder) {
	localSpaOrder = i
}

func SpaOrderLog() ISpaOrderLog {
	if localSpaOrderLog == nil {
		panic("implement not found for interface ISpaOrderLog, forgot register?")
	}
	return localSpaOrderLog
}

func RegisterSpaOrderLog(i ISpaOrderLog) {
	localSpaOrderLog = i
}

func SpaOrderTechnician() ISpaOrderTechnician {
	if localSpaOrderTechnician == nil {
		panic("implement not found for interface ISpaOrderTechnician, forgot register?")
	}
	return localSpaOrderTechnician
}

func RegisterSpaOrderTechnician(i ISpaOrderTechnician) {
	localSpaOrderTechnician = i
}

func SpaService() ISpaService {
	if localSpaService == nil {
		panic("implement not found for interface ISpaService, forgot register?")
	}
	return localSpaService
}

func RegisterSpaService(i ISpaService) {
	localSpaService = i
}

func SpaSettlement() ISpaSettlement {
	if localSpaSettlement == nil {
		panic("implement not found for interface ISpaSettlement, forgot register?")
	}
	return localSpaSettlement
}

func RegisterSpaSettlement(i ISpaSettlement) {
	localSpaSettlement = i
}

func SpaSettlementOrder() ISpaSettlementOrder {
	if localSpaSettlementOrder == nil {
		panic("implement not found for interface ISpaSettlementOrder, forgot register?")
	}
	return localSpaSettlementOrder
}

func RegisterSpaSettlementOrder(i ISpaSettlementOrder) {
	localSpaSettlementOrder = i
}

func SpaStore() ISpaStore {
	if localSpaStore == nil {
		panic("implement not found for interface ISpaStore, forgot register?")
	}
	return localSpaStore
}

func RegisterSpaStore(i ISpaStore) {
	localSpaStore = i
}

func SpaTechnician() ISpaTechnician {
	if localSpaTechnician == nil {
		panic("implement not found for interface ISpaTechnician, forgot register?")
	}
	return localSpaTechnician
}

func RegisterSpaTechnician(i ISpaTechnician) {
	localSpaTechnician = i
}
