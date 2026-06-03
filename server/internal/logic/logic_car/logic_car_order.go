package logic_car

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/aladdinApi"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shopspring/decimal"
)

type sCarOrder struct{}

func NewCarOrder() *sCarOrder {
	return &sCarOrder{}
}

func init() {
	service.RegisterCarOrder(NewCarOrder())
}

func (s *sCarOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarOrder.Ctx(ctx), option...)
}

func (s *sCarOrder) List(ctx context.Context, in *input_car.CarOrderListInp) (list []*input_car.CarOrderListModel, totalCount int, err error) {
	mod := dao.CarOrder.Ctx(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.CarOrder.Table(), input_car.CarOrderListModel{})
	mod = mod.Fields(fmt.Sprintf("`%s`.`%s` as `pmsMemberMemberNo`", dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `driver_deleted`", dao.CarDriver.Table(), dao.CarDriver.Columns().DeletedAt))

	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.CarOrder.Table(), dao.CarOrder.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))

	mod = mod.LeftJoin(dao.CarDriver.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.CarOrder.Table(), dao.CarOrder.Columns().DriverId, dao.CarDriver.Table(), dao.CarDriver.Columns().Id))

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.CarOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.CarOrder.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.ServiceType) {
		mod = mod.WhereLike(dao.CarOrder.Columns().ServiceType, in.ServiceType)
	}

	if !g.IsEmpty(in.OrderType) {
		mod = mod.WhereLike(dao.CarOrder.Columns().OrderType, in.OrderType)
	}

	if !g.IsEmpty(in.MemberSearch) {
		mod = mod.Where(mod.Builder().
			WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Id, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.CarOrder.Table(), dao.CarOrder.Columns().BookingName, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.CarOrder.Table(), dao.CarOrder.Columns().BookingMobile, "%"+in.MemberSearch+"%"))
	}

	if !g.IsEmpty(in.DriverName) {
		mod = mod.Where(mod.Builder().
			WherePrefixLike(dao.CarDriver.Table(), dao.CarDriver.Columns().Name, "%"+in.DriverName+"%").
			WhereOrPrefixLike(dao.CarDriver.Table(), dao.CarDriver.Columns().Nickname, "%"+in.DriverName+"%"))
	}

	if !g.IsEmpty(in.OrderStatus) && in.OrderStatus != "ALL" && in.OrderStatus != "ABNORMAL" {
		mod = mod.Where(dao.CarOrder.Columns().OrderStatus, in.OrderStatus)
	}

	if in.OrderStatus == "ABNORMAL" {
		mod = mod.WhereGT(dao.CarOrder.Columns().AbnormalStatus, 1)
	}

	if len(in.BookStartTime) == 2 {
		mod = mod.WhereBetween(dao.CarOrder.Columns().BookStartTime, in.BookStartTime[0], in.BookStartTime[1])
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CarOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	defaultSort := 0
	if !g.IsEmpty(in.BookSort) {
		if in.BookSort == "ascend" {
			mod = mod.Order(dao.CarOrder.Columns().BookStartTime, "asc")
		} else if in.BookSort == "descend" {
			mod = mod.Order(dao.CarOrder.Columns().BookStartTime, "desc")
		} else {
			defaultSort = 1
		}
	}

	if !g.IsEmpty(in.CreateSort) {
		if in.CreateSort == "ascend" {
			mod = mod.Order(dao.CarOrder.Columns().CreatedAt, "asc")
		} else if in.CreateSort == "descend" {
			mod = mod.Order(dao.CarOrder.Columns().CreatedAt, "desc")
		} else {
			defaultSort = 1
		}
	}
	if defaultSort == 1 {
		mod = mod.OrderDesc(dao.CarOrder.Columns().Id)
	}

	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
			return
		}
	}

	// 对于已删除的关联数据，手动加载
	for _, item := range list {
		// 加载已删除的司机
		if item.DriverDeleted && item.DriverId > 0 && item.DriverDetail == nil {
			var driver *entity.CarDriver
			if err = dao.CarDriver.Ctx(ctx).Unscoped().Where(dao.CarDriver.Columns().Id, item.DriverId).Scan(&driver); err == nil && driver != nil {
				item.DriverDetail = &struct {
					gmeta.Meta `orm:"table:hg_car_driver"`
					*entity.CarDriver
				}{CarDriver: driver}
			}
		}
		// 加载已删除的出发地址
		if item.StartAddressId > 0 && item.StartServiceAddress == nil {
			var addr *entity.CarAddress
			if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, item.StartAddressId).Scan(&addr); err == nil && addr != nil {
				item.StartServiceAddress = &struct {
					gmeta.Meta `orm:"table:hg_car_address"`
					*entity.CarAddress
				}{CarAddress: addr}
			}
		}
		// 加载已删除的目的地址
		if item.EndAddressId > 0 && item.EndServiceAddress == nil {
			var addr *entity.CarAddress
			if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, item.EndAddressId).Scan(&addr); err == nil && addr != nil {
				item.EndServiceAddress = &struct {
					gmeta.Meta `orm:"table:hg_car_address"`
					*entity.CarAddress
				}{CarAddress: addr}
			}
		}
		// 加载已删除的车辆
		if item.CarId > 0 && item.CarDetail == nil {
			var car *entity.CarCar
			if err = dao.CarCar.Ctx(ctx).Unscoped().Where(dao.CarCar.Columns().Id, item.CarId).Scan(&car); err == nil && car != nil {
				item.CarDetail = &struct {
					gmeta.Meta `orm:"table:hg_car_car"`
					*entity.CarCar
				}{CarCar: car}
			}
		}
	}

	return
}

func (s *sCarOrder) View(ctx context.Context, in *input_car.CarOrderViewInp) (res *input_car.CarOrderViewModel, err error) {
	if err = dao.CarOrder.Ctx(ctx).Unscoped().Hook(hook2.PmsFindLanguageValueHook).WithAll().
		Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取预订单信息，请稍后重试！")
		return
	}

	if res == nil {
		return
	}

	// 检查订单是否在可退款时间内
	res.CanRefund = false

	// 基础条件：必须已支付
	if res.PayStatus != "HAVE_PAID" {
		res.CanRefund = false
	} else {
		// 已支付状态，根据订单状态进一步判断
		switch res.OrderStatus {
		case "WAIT_CONFIRM":
			// 待确认状态：可以退款
			res.CanRefund = true

		case "WAIT_SERVE":
			// 待服务状态：可以退款
			res.CanRefund = true

		case "SERVING":
			// 服务中状态：可以退款
			res.CanRefund = true

		case "DONE":
			// 已完成状态：需要在配置的天数内
			if res.ActualEndTime != nil {
				// 获取sys_config表中的canRefundDay配置
				config, configErr := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
					Group: "carservicesetting",
				})
				if configErr == nil {
					// 获取允许退款的天数，默认为3天
					canRefundDays := gvar.New(config.List["canRefundDay"]).Int()
					if canRefundDays <= 0 {
						canRefundDays = 3
					}

					// 计算订单完成时间与当前时间的差值
					allowedDays := int64(canRefundDays * 24 * 60 * 60)
					timeDiff := gtime.Now().Timestamp() - res.ActualEndTime.Timestamp()
					res.CanRefund = timeDiff <= allowedDays
				}
			}

		case "CANCEL":
			// 已取消状态：如果已支付则允许退款
			res.CanRefund = true

		case "OVERDUE":
			// 过期状态：不允许退款
			res.CanRefund = false

		default:
			// 其他状态：不允许退款
			res.CanRefund = false
		}
	}

	// 检查会员是否已删除
	if res.MemberId > 0 {
		memberDeletedAt, _ := dao.PmsMember.Ctx(ctx).Unscoped().Fields(dao.PmsMember.Columns().DeletedAt).
			Where(dao.PmsMember.Columns().Id, res.MemberId).Value()
		if !memberDeletedAt.IsNil() && !memberDeletedAt.IsEmpty() {
			res.MemberDeleted = true
			// 手动加载已删除的会员信息
			if res.MemberDetail == nil {
				var member struct {
					Id       int    `json:"id"`
					FullName string `json:"fullName"`
					MemberNo string `json:"memberNo"`
				}
				if err = dao.PmsMember.Ctx(ctx).Unscoped().
					Fields("id, full_name, member_no").
					Where(dao.PmsMember.Columns().Id, res.MemberId).Scan(&member); err == nil {
					res.MemberDetail = &struct {
						gmeta.Meta `orm:"table:hg_pms_member"`
						Id         int    `json:"id"    orm:"id"      dc:"id"`
						FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
						MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
					}{Id: member.Id, FullName: member.FullName, MemberNo: member.MemberNo}
				}
			}
		}
	}

	// 检查司机是否已删除
	if res.DriverId > 0 {
		driverDeletedAt, _ := dao.CarDriver.Ctx(ctx).Unscoped().Fields(dao.CarDriver.Columns().DeletedAt).
			Where(dao.CarDriver.Columns().Id, res.DriverId).Value()
		if !driverDeletedAt.IsNil() && !driverDeletedAt.IsEmpty() {
			res.DriverDeleted = true
			// 手动加载已删除的司机信息
			if res.DriverDetail == nil {
				var driver *entity.CarDriver
				if err = dao.CarDriver.Ctx(ctx).Unscoped().Where(dao.CarDriver.Columns().Id, res.DriverId).Scan(&driver); err == nil && driver != nil {
					res.DriverDetail = &struct {
						gmeta.Meta `orm:"table:hg_car_driver"`
						*entity.CarDriver
					}{CarDriver: driver}
				}
			}
		}
	}

	// 加载已删除的出发地址
	if res.StartAddressId > 0 && res.StartServiceAddress == nil {
		var addr *entity.CarAddress
		if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, res.StartAddressId).Scan(&addr); err == nil && addr != nil {
			res.StartServiceAddress = &struct {
				gmeta.Meta `orm:"table:hg_car_address"`
				*entity.CarAddress
			}{CarAddress: addr}
		}
	}

	// 加载已删除的目的地址
	if res.EndAddressId > 0 && res.EndServiceAddress == nil {
		var addr *entity.CarAddress
		if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, res.EndAddressId).Scan(&addr); err == nil && addr != nil {
			res.EndServiceAddress = &struct {
				gmeta.Meta `orm:"table:hg_car_address"`
				*entity.CarAddress
			}{CarAddress: addr}
		}
	}

	// 加载已删除的车辆
	if res.CarId > 0 && res.CarDetail == nil {
		var car *entity.CarCar
		if err = dao.CarCar.Ctx(ctx).Unscoped().Where(dao.CarCar.Columns().Id, res.CarId).Scan(&car); err == nil && car != nil {
			res.CarDetail = &struct {
				gmeta.Meta `orm:"table:hg_car_car"`
				*entity.CarCar
				CarTypeDetail *struct {
					gmeta.Meta `orm:"table:hg_car_car_type"`
					*entity.CarCarType
				} `json:"carTypeDetail" orm:"with:id=type_id"`
			}{CarCar: car}
		}
	}

	// 加载已删除的服务
	if res.ServiceId > 0 && res.ServiceDetail == nil {
		var svc *entity.CarService
		if err = dao.CarService.Ctx(ctx).Unscoped().Where(dao.CarService.Columns().Id, res.ServiceId).Scan(&svc); err == nil && svc != nil {
			res.ServiceDetail = &struct {
				gmeta.Meta `orm:"table:hg_car_service"`
				*entity.CarService
				CarTypeDetail *struct {
					gmeta.Meta `orm:"table:hg_car_car_type"`
					*entity.CarCarType
				} `json:"carTypeDetail" orm:"with:id=car_type_id"`
			}{CarService: svc}
		}
	}

	// 创建新的切片来存储过滤后的日志
	var filteredLogList []*struct {
		gmeta.Meta  `orm:"table:hg_car_order_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		Images      string      `json:"images"      description:"图集"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	}

	for _, v := range res.LogList {
		// 跳过actionWay是RETURN的记录
		if v.ActionWay == "RETURN" {
			continue
		}

		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			var AdminMemberInfo *entity.AdminMember
			if err = dao.AdminMember.Ctx(ctx).Unscoped().Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
				return
			}
			if AdminMemberInfo != nil {
				v.OperateName = AdminMemberInfo.Username
			}
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			if PmsMemberInfo != nil {
				v.OperateName = PmsMemberInfo.FullName
			}
		}

		if v.OperateType == "DRIVER" {
			var CarDriverInfo *entity.CarDriver
			if err = dao.CarDriver.Ctx(ctx).Unscoped().Where(dao.CarDriver.Columns().Id, v.OperateId).Scan(&CarDriverInfo); err != nil {
				return
			}
			if CarDriverInfo != nil {
				v.OperateName = CarDriverInfo.Name
			}
		}

		// 将处理后的记录添加到过滤后的列表中
		filteredLogList = append(filteredLogList, v)
	}

	// 更新结果中的日志列表
	res.LogList = filteredLogList

	return
}

func (s *sCarOrder) ConfirmAgree(ctx context.Context, in *input_car.CarOrderConfirmAgreeInp) (err error) {

	var models *entity.CarOrder
	var serviceModels *entity.CarService
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.OrderType == "CRS" {
		if err = dao.CarService.Ctx(ctx).Where("id", models.ServiceId).Scan(&serviceModels); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		if g.IsEmpty(serviceModels) {
			err = gerror.New("服务不存在")
			return
		}
		timeout := gtime.New(models.BookStartTime).Add(time.Duration(serviceModels.MaxWaitTime) * time.Minute)
		nowTime := gtime.Now()

		if nowTime.After(timeout) {
			err = gerror.New("服务时间已过")
			return
		}
	} else {
		timeout := gtime.New(models.BookStartTime).EndOfDay()
		nowTime := gtime.Now()

		if nowTime.After(timeout) {
			err = gerror.New("服务时间已过")
			return
		}
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderConfirmAgreeFields{
			OrderStatus: "WAIT_SERVE",
			ConfirmTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   "CONFIRMED",
			Remark:      "订单已确认",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 判断是否自动调度
		var (
			DriverList         []*entity.CarDriver
			WorkingDriverIds   []string
			LastOrderDriverId  []gdb.Value
			CheckOrderList     []*entity.CarOrder
			CanOrderDriverList []*entity.CarDriver
			DriverInfo         *entity.CarDriver
		)
		if models.DispatchType == 2 {
			if err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Status, 1).Order("work_status asc").Scan(&DriverList); err != nil {
				return
			}

			StartTime := gtime.New(models.BookStartTime).Timestamp()
			if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().BookDate, models.BookDate).WhereNot(dao.CarOrder.Columns().Id, in.Id).WhereIn(dao.CarOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Where(dao.CarOrder.Columns().DispatchStatus, "DONE").Scan(&CheckOrderList); err != nil {
				return
			}
			for _, v := range CheckOrderList {
				OrderStartTime := gtime.New(v.BookStartTime).Timestamp()
				if gtime.New(v.BookEndTime).Timestamp() <= StartTime || OrderStartTime >= gtime.New(models.BookEndTime).Timestamp() {

				} else {
					driverIdString := gvar.New(v.DriverId).String()
					WorkingDriverIds = append(WorkingDriverIds, driverIdString)
				}
			}
			if !g.IsEmpty(DriverList) {
				for _, v := range DriverList {
					if v.WorkStatus != "REST" && !gstr.InArray(WorkingDriverIds, strconv.Itoa(v.Id)) {
						CanOrderDriverList = append(CanOrderDriverList, v)
					}
				}
			}

			if !g.IsEmpty(DriverList) && !g.IsEmpty(CanOrderDriverList) {
				if LastOrderDriverId, err = dao.CarOrder.Ctx(ctx).Fields(dao.CarOrder.Columns().DriverId).Where(dao.CarOrder.Columns().DispatchStatus, "DONE").Where(dao.CarOrder.Columns().PayStatus, "HAVE_PAID").Order("dispatch_time desc").Array(); err != nil {
					return
				}

				if len(LastOrderDriverId) > 0 {
					DriverInfo = CanOrderDriverList[0]

					for k, v := range CanOrderDriverList {
						if g.NewVar(LastOrderDriverId).Ints()[0] == v.Id {
							if k != len(CanOrderDriverList)-1 {
								DriverInfo = CanOrderDriverList[k+1]
								break
							}
						}
					}
				} else {
					// 没有指定过就第一个司机
					DriverInfo = CanOrderDriverList[0]
				}

				if _, err = s.Model(ctx).
					WherePri(in.Id).Data(input_car.CarOrderDispatchInp{
					DispatchStatus:     "DONE",
					DispatchTime:       gtime.Now(),
					DispatchOperatorId: 0,
					DispatchDesc:       fmt.Sprintf("系统自动派单给%s", DriverInfo.Nickname),
					DriverId:           int(DriverInfo.Id),
					CarId:              int(DriverInfo.CarId),
				}).OmitEmptyData().Update(); err != nil {
					err = gerror.Wrap(err, "操作失败，请稍后重试！")
					return
				}
				// 订单日志
				if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
					OrderId:     in.Id,
					OrderStatus: "WAIT_SERVE",
					ActionWay:   "DISPATCH",
					Remark:      "预约成功，司机已接单",
					OperateType: "SYSTEM",
				}); err != nil {
					return err
				}

				// 打印
				_ = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
					OrderId: in.Id,
				})

				// 更新司机统计
				if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, DriverInfo.Id).Update(g.MapStrAny{
					dao.CarDriver.Columns().TotalOrderNum:    gdb.Raw("total_order_num+1"),
					dao.CarDriver.Columns().TotalOrderAmount: gdb.Raw(fmt.Sprintf("total_order_amount+%f", models.OrderAmount)),
					dao.CarDriver.Columns().PayOrderNum:      gdb.Raw("pay_order_num+1"),
					dao.CarDriver.Columns().PayOrderAmount:   gdb.Raw(fmt.Sprintf("pay_order_amount+%f", models.OrderAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
					return
				}

				// 写入结算
				var (
					CarSettlement   *entity.CarSettlement
					SettlementRate  float64
					SettlementType  int
					SettlementCycle int
				)
				if err = dao.CarSettlement.Ctx(ctx).WherePri(DriverInfo.SettlementId).Scan(&CarSettlement); err != nil {
					return
				}
				if DriverInfo.SettlementType == 1 {
					// 跟随系统
					if !g.IsEmpty(CarSettlement) {
						SettlementRate = CarSettlement.Rate
					} else {
						SettlementRate = 0
					}
				} else {
					// 自定义
					SettlementRate = DriverInfo.SettlementRate
				}
				if !g.IsEmpty(CarSettlement) {
					SettlementType = CarSettlement.Type
					SettlementCycle = CarSettlement.Cycle
				} else {
					SettlementType = 1
					SettlementCycle = 1
				}
				if _, err = s.Model(ctx).
					WherePri(in.Id).Data(input_car.CarOrderSettlementInp{
					SettlementRate:  SettlementRate,
					SettlementType:  SettlementType,
					SettlementCycle: SettlementCycle,
				}).OmitEmptyData().Update(); err != nil {
					err = gerror.Wrap(err, "操作失败，请稍后重试！")
					return
				}
			}
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已确认",
			"en":    "Order confirmed",
			"ja":    "注文確定",
			"ko":    "주문 확인됨",
			"zh_CN": "訂單已確認",
		}
		systemMessageContent := map[string]string{
			"zh":    models.OrderSn + "订单后台已接单",
			"en":    "Order " + models.OrderSn + " has been received in the system.",
			"ja":    "システムで注文" + models.OrderSn + "が受信されました。",
			"ko":    "시스템에 주문" + models.OrderSn + "이 접수되었습니다.",
			"zh_CN": models.OrderSn + "訂單後台已接單",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		// 发送确认短信(队列)
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderRemind,
			DataByte: gjson.New(g.Map{
				"orderSn": models.OrderSn,
				"event":   "car_order_confirm",
			}).MustToJson(),
			Header: nil,
		})
		return
	})

}

// ConfirmDisagree 确认失败，全额退款
func (s *sCarOrder) ConfirmDisagree(ctx context.Context, in *input_car.CarOrderConfirmDisagreeInp) (err error) {

	var models *entity.CarOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.RefundStatus != "WAIT" {
		err = gerror.New("订单已退款")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 改变订单确认状态
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(input_car.CarOrderConfirmDisagreeFields{
			BookingStatus:       "CANCEL",
			ConfirmRefuseReason: in.ConfirmRefuseReason,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().TX(tx).Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "CANCEL",
			ActionWay:   "DISCONFIRMED",
			Remark:      fmt.Sprintf("订单确认拒绝，原因：%s", in.ConfirmRefuseReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 确认失败、全额退款
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).Scan(&Transaction); err != nil {
			return
		}
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

		for _, v := range Transaction {
			// 可退款金额
			Refundable := v.Amount - v.RefundAmount
			if v.PayType == "BAL" {
				RefundBalance = RefundBalance + Refundable
			} else if v.PayType == "COUPON" {

			} else {
				RefundFee = RefundFee + Refundable
			}
		}
		// 退款金额
		RefundAmount := RefundBalance + RefundFee

		// 修改订单状态
		RefundStatus := "DONE"
		if models.OrderAmount > RefundAmount {
			RefundStatus = "PART"
		}

		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(g.MapStrAny{
			dao.CarOrder.Columns().RefundAmount:       RefundAmount,
			dao.CarOrder.Columns().RefundBalAmount:    RefundBalance,
			dao.CarOrder.Columns().RefundCouponAmount: 0,
			dao.CarOrder.Columns().RefundStatus:       RefundStatus,
			dao.CarOrder.Columns().RefundTime:         gtime.Now(),
			dao.CarOrder.Columns().OrderStatus:        "CANCEL",
			dao.CarOrder.Columns().PayStatus:          "REFUND",
			dao.CarOrder.Columns().CancelTime:         gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 更新服务预定量和预定金额
		if _, err = dao.CarService.Ctx(ctx).TX(tx).Where(dao.CarService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
			dao.CarService.Columns().PayOrderNum:    gdb.Raw("pay_order_num-1"),
			dao.CarService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
			return
		}

		// 更新车辆预定量和预定金额
		if models.CarId > 0 {
			if _, err = dao.CarCar.Ctx(ctx).TX(tx).Where(dao.CarCar.Columns().Id, models.CarId).Update(g.MapStrAny{
				dao.CarCar.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
				dao.CarCar.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新车辆信息失败，请稍后重试！")
				return
			}
		}

		// 更新司机预定量和预定金额
		if models.DriverId > 0 {
			if _, err = dao.CarDriver.Ctx(ctx).TX(tx).Where(dao.CarDriver.Columns().Id, models.DriverId).Update(g.MapStrAny{
				dao.CarDriver.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
				dao.CarDriver.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "CANCEL",
			ActionWay:   "REFUND",
			Remark:      "订单已退款",
			OperateType: "SYSTEM",
		}); err != nil {
			return err
		}

		// 全额退款
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      models.OrderSn,
			RefundAmount: RefundAmount,
			OperateType:  "ADMIN",
			OperateId:    int(contexts.GetUserId(ctx)),
		}, tx)
		if err != nil {
			err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已被拒",
			"en":    "The order has been rejected.",
			"ja":    "注文は拒否されました。",
			"ko":    "주문이 거부되었습니다.",
			"zh_CN": "訂單已被拒",
		}
		systemMessageContent := map[string]string{
			"zh":    "订单" + models.OrderSn + "已被拒，已全额退款",
			"en":    "Order " + models.OrderSn + " has been rejected and fully refunded.",
			"ja":    "注文 " + models.OrderSn + " は拒否され、全額返金されました。",
			"ko":    "주문 " + models.OrderSn + "은 거부되었으며 전액 환불되었습니다.",
			"zh_CN": "訂單" + models.OrderSn + "已被拒，已全額退款",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

func (s *sCarOrder) DriverList(ctx context.Context, in *input_car.CarOrderDriverInp) (list []*input_car.CarOrderDriverModel, totalCount int, err error) {
	var (
		Config           *input_basics.GetConfigModel
		OrderInfo        *entity.CarOrder
		WorkingDriverIds []string
		CheckOrderList   []*entity.CarOrder
	)
	if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().Id, in.Id).Scan(&OrderInfo); err != nil {
		return
	}

	Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "carordersetting",
	})
	if err != nil {
		return
	}
	OrderDispatchModel := Config.List["orderDispatch"]

	mod := dao.CarDriver.Ctx(ctx)
	mod = mod.Fields(input_car.CarOrderDriverModel{})

	if gvar.New(OrderDispatchModel).Int() == 2 {
		StartTime := gtime.New(OrderInfo.BookStartTime).Timestamp()
		if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().BookDate, OrderInfo.BookDate).WhereNot(dao.CarOrder.Columns().Id, in.Id).WhereIn(dao.CarOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Where(dao.CarOrder.Columns().DispatchStatus, "DONE").Scan(&CheckOrderList); err != nil {
			return
		}
		for _, v := range CheckOrderList {
			OrderStartTime := gtime.New(v.BookStartTime).Timestamp()
			if gtime.New(v.BookEndTime).Timestamp() <= StartTime || OrderStartTime >= gtime.New(OrderInfo.BookEndTime).Timestamp() {

			} else {
				driverIdString := gvar.New(v.DriverId).String()
				WorkingDriverIds = append(WorkingDriverIds, driverIdString)
			}
		}
	}
	WorkingDriverIdsStr := strings.Join(WorkingDriverIds, ",")

	if g.IsEmpty(WorkingDriverIdsStr) {
		WorkingDriverIdsStr = "0"
	}
	mod = mod.Fields(fmt.Sprintf(`
	CASE
        WHEN work_status='REST' THEN 1
        WHEN id IN(%s) THEN 2
        ELSE 3
    END AS work_status_enum
`, WorkingDriverIdsStr))
	mod = mod.Fields(fmt.Sprintf(`
	%d AS dispatch_mode
`, gvar.New(OrderDispatchModel).Int()))
	mod = mod.Where(dao.CarDriver.Columns().Status, 1)
	mod = mod.OrderDesc(fmt.Sprintf(`
	CASE
        WHEN work_status='REST' THEN 1
        WHEN id IN(%s) THEN 2
        ELSE 3
    END
`, WorkingDriverIdsStr))
	mod = mod.Page(in.Page, in.PerPage)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取司机列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sCarOrder) Dispatch(ctx context.Context, in *input_car.CarOrderDispatchInp) (err error) {

	var (
		Config         *input_basics.GetConfigModel
		models         *entity.CarOrder
		DriverInfo     *entity.CarDriver
		CheckOrderList []*entity.CarOrder
		CantOrderNum   int
		actionWay      string
		logDesc        string
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" || models.DispatchStatus != "WAIT" {
		err = gerror.New("订单状态不正确")
		return
	}

	if err = dao.CarDriver.Ctx(ctx).WherePri(in.DriverId).Scan(&DriverInfo); err != nil {
		return
	}

	Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "carordersetting",
	})
	if err != nil {
		return
	}
	OrderDispatchModel := Config.List["orderDispatch"]
	if gvar.New(OrderDispatchModel).Int() == 2 {
		// 重新判断下司机

		if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().BookDate, models.BookDate).Where(dao.CarOrder.Columns().DriverId, in.DriverId).WhereNot(dao.CarOrder.Columns().Id, in.Id).WhereIn(dao.CarOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Scan(&CheckOrderList); err != nil {
			return
		}
		for _, v := range CheckOrderList {
			OrderStartTime := gtime.New(v.BookStartTime).Timestamp()
			if gtime.New(v.BookEndTime).Timestamp() <= gtime.New(models.BookStartTime).Timestamp() || OrderStartTime >= gtime.New(models.BookEndTime).Timestamp() {

			} else {
				CantOrderNum++
			}
		}

		if CantOrderNum > 0 {
			err = gerror.New("司机不可派单，请重新选择")
			return
		}
	}

	if models.IsReturn == 1 {
		actionWay = "RE_DISPATCH"
		logDesc = "新司机已接单"
	} else {
		actionWay = "DISPATCH"
		logDesc = "预约成功，司机已接单"
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderDispatchInp{
			IsReturn:           2,
			DispatchStatus:     "DONE",
			DispatchTime:       gtime.Now(),
			DispatchOperatorId: in.DispatchOperatorId,
			DispatchDesc:       in.DispatchDesc,
			DriverId:           in.DriverId,
			CarId:              DriverInfo.CarId,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   actionWay,
			Remark:      logDesc,
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 打印
		_ = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
			OrderId: in.Id,
		})

		// 更新司机统计
		if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, in.DriverId).Update(g.MapStrAny{
			dao.CarDriver.Columns().TotalOrderNum:    gdb.Raw("total_order_num+1"),
			dao.CarDriver.Columns().TotalOrderAmount: gdb.Raw(fmt.Sprintf("total_order_amount+%f", models.OrderAmount)),
			dao.CarDriver.Columns().PayOrderNum:      gdb.Raw("pay_order_num+1"),
			dao.CarDriver.Columns().PayOrderAmount:   gdb.Raw(fmt.Sprintf("pay_order_amount+%f", models.OrderAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
			return
		}

		// 写入结算
		var (
			CarSettlement   *entity.CarSettlement
			SettlementRate  float64
			SettlementType  int
			SettlementCycle int
		)
		if err = dao.CarSettlement.Ctx(ctx).WherePri(DriverInfo.SettlementId).Scan(&CarSettlement); err != nil {
			return
		}
		if DriverInfo.SettlementType == 1 {
			// 跟随系统
			if !g.IsEmpty(CarSettlement) {
				SettlementRate = CarSettlement.Rate
			} else {
				SettlementRate = 0
			}
		} else {
			// 自定义
			SettlementRate = DriverInfo.SettlementRate
		}
		if !g.IsEmpty(CarSettlement) {
			SettlementType = CarSettlement.Type
			SettlementCycle = CarSettlement.Cycle
		} else {
			SettlementType = 1
			SettlementCycle = 1
		}
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderSettlementInp{
			SettlementRate:  SettlementRate,
			SettlementType:  SettlementType,
			SettlementCycle: SettlementCycle,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 发送超时队列
		var TimeDuration int64
		if models.OrderType == "CRS" {
			var CarService *entity.CarService
			if err = dao.CarService.Ctx(ctx).Unscoped().WherePri(models.ServiceId).Scan(&CarService); err != nil {
				return
			}
			// CRS订单：预订开始时间 + 最大等待时间(分钟转秒) - 当前时间
			TimeDuration = models.BookStartTime.Unix() + int64(CarService.MaxWaitTime*60) - gtime.Now().Unix()
		} else {
			// 非CRS订单：预订开始时间当天的23:59:59 - 当前时间
			endOfDay := gtime.New(models.BookStartTime.Format("Y-m-d") + " 23:59:59")
			TimeDuration = endOfDay.Unix() - gtime.Now().Unix()
		}

		// 确保TimeDuration不为负数
		if TimeDuration < 0 {
			TimeDuration = 0
		}
		if TimeDuration > 0 {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeDelayedName,
				QueueName:    consts.RabbitMQQueueNameOrderExpire,
				DataByte:     gvar.New("Y-" + models.OrderSn).Bytes(),
				Header: amqp.Table{
					"x-delay": gvar.New(TimeDuration * 1000).String(),
				},
			})
		} else {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameOrderExpire,
				DataByte:     gvar.New("Y-" + models.OrderSn).Bytes(),
				Header:       nil,
			})
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "司机已接单",
			"en":    "The driver has accepted the order.",
			"ja":    "ドライバーは注文を承諾しました。",
			"ko":    "운전자가 주문을 수락했습니다.",
			"zh_CN": "司機已接單",
		}
		systemMessageContent := map[string]string{
			"zh":    "您的订单" + models.OrderSn + "已派单给司机" + DriverInfo.Nickname,
			"en":    "Your order " + models.OrderSn + " has been dispatched to driver " + DriverInfo.Nickname,
			"ja":    "ご注文番号" + models.OrderSn + "はドライバー" + DriverInfo.Nickname + "に発送されました",
			"ko":    "귀하의 주문 " + models.OrderSn + "이 드라이버 " + DriverInfo.Nickname + "에게 발송되었습니다.",
			"zh_CN": "您的訂單" + models.OrderSn + "已派單給司機" + DriverInfo.Nickname,
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		// 发送确认短信(队列)
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderRemind,
			DataByte: gjson.New(g.Map{
				"orderSn": models.OrderSn,
				"event":   "car_order_dispatch",
			}).MustToJson(),
			Header: nil,
		})

		return
	})

}

func (s *sCarOrder) TransferOrder(ctx context.Context, in *input_car.CarOrderTransferInp) (err error) {

	var (
		models         *entity.CarOrder
		OldDriverInfo  *entity.CarDriver
		DriverInfo     *entity.CarDriver
		CheckOrderList []*entity.CarOrder
		CantOrderNum   int
	)

	// 订单信息
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	// 判断订单状态是否是待服务并且是已派单状态
	if models.OrderStatus != "WAIT_SERVE" || models.DispatchStatus != "DONE" {
		err = gerror.New("订单状态不正确")
		return
	}

	// 旧司机信息
	if err = dao.CarDriver.Ctx(ctx).Unscoped().WherePri(models.DriverId).Scan(&OldDriverInfo); err != nil {
		return
	}

	// 新司机信息
	if err = dao.CarDriver.Ctx(ctx).WherePri(in.DriverId).Scan(&DriverInfo); err != nil {
		return
	}

	// 重新判断下司机
	if err = dao.CarOrder.Ctx(ctx).Where(dao.CarOrder.Columns().BookDate, models.BookDate).Where(dao.CarOrder.Columns().DriverId, in.DriverId).WhereNot(dao.CarOrder.Columns().Id, in.Id).WhereIn(dao.CarOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Scan(&CheckOrderList); err != nil {
		return
	}
	for _, v := range CheckOrderList {
		OrderStartTime := gtime.New(v.BookStartTime).Timestamp()
		if gtime.New(v.BookEndTime).Timestamp() <= gtime.New(models.BookStartTime).Timestamp() || OrderStartTime >= gtime.New(models.BookEndTime).Timestamp() {

		} else {
			CantOrderNum++
		}
	}

	if CantOrderNum > 0 {
		err = gerror.New("司机不可派单，请重新选择")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 先退单
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(g.Map{
			dao.CarOrder.Columns().DispatchStatus:  "WAIT",
			dao.CarOrder.Columns().DispatchTime:    nil,
			dao.CarOrder.Columns().DispatchDesc:    "",
			dao.CarOrder.Columns().IsReturn:        1,
			dao.CarOrder.Columns().DriverId:        0,
			dao.CarOrder.Columns().CarId:           0,
			dao.CarOrder.Columns().ReturnDriverId:  OldDriverInfo.Id,
			dao.CarOrder.Columns().ReturnCarId:     OldDriverInfo.CarId,
			dao.CarOrder.Columns().SettlementRate:  0,
			dao.CarOrder.Columns().SettlementType:  1,
			dao.CarOrder.Columns().SettlementCycle: 1,
		}).Update(); err != nil {
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			ActionWay:   "RETURN",
			OrderStatus: "WAIT_SERVE",
			Remark:      "司机已退单",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 更新旧司机统计
		if _, err = dao.CarDriver.Ctx(ctx).Unscoped().WherePri(OldDriverInfo.Id).Data(g.Map{
			dao.CarDriver.Columns().TotalOrderNum:    gdb.Raw("total_order_num-1"),
			dao.CarDriver.Columns().TotalOrderAmount: gdb.Raw(fmt.Sprintf("total_order_amount-%f", models.OrderAmount)),
			dao.CarDriver.Columns().PayOrderNum:      gdb.Raw("pay_order_num-1"),
			dao.CarDriver.Columns().PayOrderAmount:   gdb.Raw(fmt.Sprintf("pay_order_amount-%f", models.OrderAmount)),
		}).Update(); err != nil {
			return err
		}

		// 重新派单
		var OldCarInfo *entity.CarCar
		var CarInfo *entity.CarCar
		if OldDriverInfo.CarId > 0 {
			if err = dao.CarCar.Ctx(ctx).WherePri(OldDriverInfo.CarId).Scan(&OldCarInfo); err != nil {
				return
			}
		}

		if DriverInfo.Status != 1 || DriverInfo.WorkStatus == "REST" {
			err = gerror.New("司机状态不正确，请稍后重试！")
			return
		}

		if DriverInfo.CarId > 0 {
			if err = dao.CarCar.Ctx(ctx).WherePri(DriverInfo.CarId).Scan(&CarInfo); err != nil {
				return
			}
		}

		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(g.Map{
			dao.CarOrder.Columns().IsReturn:           2,
			dao.CarOrder.Columns().DispatchStatus:     "DONE",
			dao.CarOrder.Columns().DispatchTime:       gtime.Now(),
			dao.CarOrder.Columns().DispatchOperatorId: int(contexts.GetUserId(ctx)),
			dao.CarOrder.Columns().DispatchDesc:       in.DispatchDesc,
			dao.CarOrder.Columns().DriverId:           in.DriverId,
			dao.CarOrder.Columns().CarId:              DriverInfo.CarId,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			ActionWay:   "TRANSFER",
			OrderStatus: "WAIT_SERVE",
			Remark:      "新司机已接单",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 打印
		_ = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
			OrderId: in.Id,
		})

		// 更新司机统计
		if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().Id, in.DriverId).Update(g.MapStrAny{
			dao.CarDriver.Columns().TotalOrderNum:    gdb.Raw("total_order_num+1"),
			dao.CarDriver.Columns().TotalOrderAmount: gdb.Raw(fmt.Sprintf("total_order_amount+%f", models.OrderAmount)),
			dao.CarDriver.Columns().PayOrderNum:      gdb.Raw("pay_order_num+1"),
			dao.CarDriver.Columns().PayOrderAmount:   gdb.Raw(fmt.Sprintf("pay_order_amount+%f", models.OrderAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
			return
		}

		// 写入结算
		var (
			CarSettlement   *entity.CarSettlement
			SettlementRate  float64
			SettlementType  int
			SettlementCycle int
		)
		if err = dao.CarSettlement.Ctx(ctx).WherePri(DriverInfo.SettlementId).Scan(&CarSettlement); err != nil {
			return
		}
		if DriverInfo.SettlementType == 1 {
			// 跟随系统
			if !g.IsEmpty(CarSettlement) {
				if CarSettlement.Type == 1 {
					SettlementRate = 0
				} else {
					SettlementRate = CarSettlement.Rate
				}
			} else {
				SettlementRate = 0
			}
		} else {
			// 自定义
			SettlementRate = DriverInfo.SettlementRate
		}
		if !g.IsEmpty(CarSettlement) {
			SettlementType = CarSettlement.Type
			SettlementCycle = CarSettlement.Cycle
		} else {
			SettlementType = 1
			SettlementCycle = 1
		}
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderSettlementInp{
			SettlementRate:  SettlementRate,
			SettlementType:  SettlementType,
			SettlementCycle: SettlementCycle,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单转单",
			"en":    "Order transfer",
			"ja":    "注文転送",
			"ko":    "주문 전송",
			"zh_CN": "訂單轉單",
		}
		systemMessageContent := map[string]string{
			"zh":    "订单已派单给新司机" + DriverInfo.Nickname,
			"en":    "The order has been assigned to the new driver " + DriverInfo.Nickname,
			"ja":    "注文は新しいドライバー" + DriverInfo.Nickname + "に割り当てられました",
			"ko":    "새로운 운전자에게 주문이 할당되었습니다. " + DriverInfo.Nickname,
			"zh_CN": "訂單已派單給新司機" + DriverInfo.Nickname,
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		//// 发送确认短信(队列)
		//_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		//	ExchangeName: consts.RabbitMQExchangeName,
		//	QueueName:    consts.RabbitMQQueueNameOrderRemind,
		//	DataByte: gjson.New(g.Map{
		//		"orderSn": models.OrderSn,
		//		"event":   "car_order_dispatch",
		//	}).MustToJson(),
		//	Header: nil,
		//})

		return
	})

}

// SettleOrderList 结算订单列表
func (s *sCarOrder) SettleOrderList(ctx context.Context, in *input_car.SettleCarOrderListInp) (list []*input_car.SettleCarOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_car.SettleCarOrderListModel{})

	if !g.IsEmpty(in.SettlementOrderId) {
		mod = mod.Where(dao.CarOrder.Columns().SettlementOrderId, in.SettlementOrderId)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.CarOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CarOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.CarOrder.Columns().Id)
	//mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算订单列表失败，请稍后重试！")
		return
	}
	return
}

// Refund 订单退款
func (s *sCarOrder) Refund(ctx context.Context, in *input_car.CarOrderRefundInp) (err error) {

	var models *entity.CarOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.PayStatus != "HAVE_PAID" {
		err = gerror.New("订单支付状态不正确")
		return
	}

	// 如果订单已完成，需要判断是否在配置的天数内
	if models.OrderStatus == "DONE" {
		if models.ActualEndTime == nil {
			err = gerror.New("订单完成时间异常，无法退款")
			return
		}

		// 获取sys_config表中的canRefundDay配置
		config, configErr := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
			Group: "carservicesetting",
		})
		if configErr != nil {
			err = gerror.Wrap(configErr, "获取退款配置失败")
			return
		}

		// 获取允许退款的天数，默认为3天
		canRefundDays := gvar.New(config.List["canRefundDay"]).Int()
		if canRefundDays <= 0 {
			canRefundDays = 3
		}

		// 计算订单完成时间与当前时间的差值
		allowedDays := int64(canRefundDays * 24 * 60 * 60) // 配置天数的秒数
		timeDiff := gtime.Now().Timestamp() - models.ActualEndTime.Timestamp()
		if timeDiff > allowedDays {
			err = gerror.Newf("订单已完成超过%d天，无法退款", canRefundDays)
			return
		}
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			CancelFee         float64 // 退款手续费
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		CancelFee = models.OrderAmount - models.CouponAmount - models.RefundAmount - in.RefundMoney // 订单金额100， 优惠券10， 余额支付30，三方支付 60， 申请退款5， fee 85
		CancelFeeCalc := CancelFee
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().PayStatus, "DONE").Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).OrderDesc(`
			CASE pay_type
				WHEN 'StripeCard' THEN 1
				WHEN 'PaypalCard' THEN 2
				WHEN 'Paypal' THEN 3
				WHEN 'WeChatPay' THEN 4
				WHEN 'Alipay+' THEN 5
				WHEN 'WeChatMiniPay' THEN 6
				WHEN 'BAL' THEN 7
				ELSE 8
			END
		`).Scan(&Transaction); err != nil {
			return
		}
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

		for _, v := range Transaction {
			// 可退款金额
			Refundable := v.Amount - v.RefundAmount
			if v.PayType == "BAL" {
				if g.IsEmpty(CancelFeeCalc) {
					RefundBalance += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundBalance += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundBalance += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			} else if v.PayType == "COUPON" {

			} else {
				if g.IsEmpty(CancelFeeCalc) {
					RefundFee += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundFee += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundFee += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			}
		}
		// 退款金额
		RefundAmount := RefundBalance + RefundFee

		OrderTotalRefundAmount := RefundAmount + models.RefundAmount
		OrderTotalRefundBalance := RefundBalance + models.RefundBalAmount

		AdminOrderTotalRefundAmount := RefundAmount + models.AdminRefundAmount
		AdminOrderTotalRefundBalance := RefundBalance + models.AdminRefundBalAmount

		// 修改订单状态
		RefundStatus := "DONE"
		if models.OrderAmount > OrderTotalRefundAmount {
			RefundStatus = "PART"
		}
		if in.RefundType == 1 {
			// 仅退款
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				dao.CarOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.CarOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.CarOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.CarOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.CarOrder.Columns().RefundCouponAmount:   0,
				dao.CarOrder.Columns().RefundStatus:         RefundStatus,
				dao.CarOrder.Columns().RefundTime:           gtime.Now(),
				dao.CarOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
			// 更新服务预定量和预定金额
			if _, err = dao.CarService.Ctx(ctx).TX(tx).Where(dao.CarService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
				dao.CarService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
				return
			}

			// 更新车辆预定量和预定金额
			if models.CarId > 0 {
				if _, err = dao.CarCar.Ctx(ctx).TX(tx).Where(dao.CarCar.Columns().Id, models.CarId).Update(g.MapStrAny{
					dao.CarCar.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新车辆信息失败，请稍后重试！")
					return
				}
			}

			// 更新司机预定量和预定金额
			if models.DriverId > 0 {
				if _, err = dao.CarDriver.Ctx(ctx).TX(tx).Where(dao.CarDriver.Columns().Id, models.DriverId).Update(g.MapStrAny{
					dao.CarDriver.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
					return
				}
			}
		} else {
			// 退款并取消
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				dao.CarOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.CarOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.CarOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.CarOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.CarOrder.Columns().RefundCouponAmount:   0,
				dao.CarOrder.Columns().RefundStatus:         RefundStatus,
				dao.CarOrder.Columns().RefundTime:           gtime.Now(),
				dao.CarOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
				dao.CarOrder.Columns().OrderStatus:          "CANCEL",
				dao.CarOrder.Columns().PayStatus:            "REFUND",
				dao.CarOrder.Columns().CancelTime:           gtime.Now(),
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
			// 更新服务预定量和预定金额
			if _, err = dao.CarService.Ctx(ctx).TX(tx).Where(dao.CarService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
				dao.CarService.Columns().PayOrderNum:    gdb.Raw("pay_order_num-1"),
				dao.CarService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
				return
			}

			// 更新车辆预定量和预定金额
			if models.CarId > 0 {
				if _, err = dao.CarCar.Ctx(ctx).TX(tx).Where(dao.CarCar.Columns().Id, models.CarId).Update(g.MapStrAny{
					dao.CarCar.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
					dao.CarCar.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新车辆信息失败，请稍后重试！")
					return
				}
			}

			// 更新司机预定量和预定金额
			if models.DriverId > 0 {
				if _, err = dao.CarDriver.Ctx(ctx).TX(tx).Where(dao.CarDriver.Columns().Id, models.DriverId).Update(g.MapStrAny{
					dao.CarDriver.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
					dao.CarDriver.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
				}); err != nil {
					err = gerror.Wrap(err, "更新司机信息失败，请稍后重试！")
					return
				}
			}

			// 如果是异常单，则需要异常完成处理
			if models.AbnormalStatus == 2 {
				if _, err = s.Model(ctx).
					WherePri(in.Id).Data(g.Map{
					dao.CarOrder.Columns().AbnormalStatus: 3,
					dao.CarOrder.Columns().AbnormalTime:   gtime.Now(),
					dao.CarOrder.Columns().AbnormalReason: "后台取消订单",
				}).OmitEmptyData().Update(); err != nil {
					err = gerror.Wrap(err, "操作失败，请稍后重试！")
					return
				}
			}
		}
		if _, err = dao.CarOrder.Ctx(ctx).TX(tx).WherePri(in.Id).Update(g.MapStrAny{
			dao.CarOrder.Columns().AdminCancelNum: gdb.Raw("admin_cancel_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新失败，请稍后重试！")
			return
		}

		if RefundAmount > 0 {
			// 退款
			err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
				OrderSn:      models.OrderSn,
				RefundAmount: RefundAmount,
				Remark:       in.AdminCancelReason,
				OperateType:  "ADMIN",
				OperateId:    int(contexts.GetUserId(ctx)),
			}, tx)
			if err != nil {
				err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     int(in.Id),
			OrderStatus: models.OrderStatus,
			ActionWay:   "ADMIN_REFUND",
			Remark:      fmt.Sprintf("后台退款，原因：%s", in.AdminCancelReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
			Images:      in.Images,
		}); err != nil {
			return err
		}

		if in.RefundType == 2 {
			// 订单日志
			if _, err = dao.CarOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CarOrderLog{
				OrderId:     int(in.Id),
				OrderStatus: "CANCEL",
				ActionWay:   "CANCEL",
				Remark:      "后台订单取消",
				OperateType: "ADMIN",
				OperateId:   int(contexts.GetUserId(ctx)),
				Images:      in.Images,
			}); err != nil {
				return err
			}

			if models.OrderType == "INNN" && !g.IsEmpty(models.InnnOrderId) {
				if _, err = aladdinApi.NewClient(ctx).CancelOrder(ctx, &aladdinApi.CancelOrderParams{OrderId: models.InnnOrderId}); err != nil {
					return
				}
			}

			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单已取消",
				"en":    "Order has been canceled",
				"ja":    "注文はキャンセルされました",
				"ko":    "주문이 취소되었습니다",
				"zh_CN": "訂單已取消",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and " + gvar.New(RefundBalance).String() + " points have been refunded",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消",
			// 		"en":    "Order has been canceled",
			// 		"ja":    "注文はキャンセルされました",
			// 		"ko":    "주문이 취소되었습니다",
			// 		"zh_CN": "訂單已取消",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消，并成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Order has been canceled, and a refund of " + gvar.New(RefundAmount).String() + "JPY has been issued",
					"ja":    "注文はキャンセルされ、" + gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    "주문이 취소되었으며 " + gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消",
					"en":    "Order has been canceled",
					"ja":    "注文はキャンセルされました",
					"ko":    "주문이 취소되었습니다",
					"zh_CN": "訂單已取消",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
			var memberLanguage string
			if phoneArea.String() == "+86" {
				memberLanguage = "zh"
			} else if phoneArea.String() == "+81" {
				memberLanguage = "ja"
			} else if phoneArea.String() == "+82" {
				memberLanguage = "ko"
			} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
				memberLanguage = "zh_CN"
			} else {
				memberLanguage = "en"
			}
			appPushData := g.MapStrStr{
				"type":   "1",
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "car",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/transferDetailScene",
				WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		} else {
			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单成功退款",
				"en":    "Refund Successful",
				"ja":    "返金が正常に完了しました",
				"ko":    "환불이 성공적으로 완료되었습니다",
				"zh_CN": "訂單成功退款",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款",
			// 		"en":    "Refund Successful",
			// 		"ja":    "返金が正常に完了しました",
			// 		"ko":    "환불이 성공적으로 완료되었습니다",
			// 		"zh_CN": "訂單成功退款",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Refund of " + gvar.New(RefundAmount).String() + "JPY completed",
					"ja":    gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單成功退款" + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款",
					"en":    "Refund Successful",
					"ja":    "返金が正常に完了しました",
					"ko":    "환불이 성공적으로 완료되었습니다",
					"zh_CN": "訂單成功退款",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
			var memberLanguage string
			if phoneArea.String() == "+86" {
				memberLanguage = "zh"
			} else if phoneArea.String() == "+81" {
				memberLanguage = "ja"
			} else if phoneArea.String() == "+82" {
				memberLanguage = "ko"
			} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
				memberLanguage = "zh_CN"
			} else {
				memberLanguage = "en"
			}
			appPushData := g.MapStrStr{
				"type":   "1",
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "car",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/transferDetailScene",
				WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		}

		return
	})

}

func (s *sCarOrder) GoOut(ctx context.Context, in *input_car.CarOrderGoOutInp) (err error) {

	var models *entity.CarOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.DispatchStatus != "DONE" {
		err = gerror.New("调度状态不正确")
		return
	}

	if !g.IsEmpty(models.DriverGoTime) {
		err = gerror.New("司机已出发，请勿重复点击")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderGoOutFields{
			DriverGoTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "SERVING",
			ActionWay:   "OUT",
			Remark:      "司机已出发",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 司机状态改为工作中
		if _, err = dao.CarDriver.Ctx(ctx).
			WherePri(models.DriverId).
			Data(g.MapStrAny{
				dao.CarDriver.Columns().WorkStatus: "WORKING",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改司机工作状态失败，请稍后重试！")
			return
		}

		// 车辆状态改为工作中
		if models.CarId > 0 {
			if _, err = dao.CarCar.Ctx(ctx).
				WherePri(models.CarId).
				Data(g.MapStrAny{
					dao.CarCar.Columns().WorkStatus: "WORKING",
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改车辆工作状态失败，请稍后重试！")
				return
			}
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "司机已出发",
			"en":    "The driver has departed.",
			"ja":    "運転手は出発しました。",
			"ko":    "운전자가 떠났습니다.",
			"zh_CN": "司機已出發",
		}
		systemMessageContent := map[string]string{
			"zh":    "您的订单" + models.OrderSn + "司机已出发，正在赶往您的出发地",
			"en":    "Your order number " + models.OrderSn + " has been dispatched by driver and is en route to your departure point.",
			"ja":    "ご注文番号 " + models.OrderSn + " はドライバーによって発送され、出発地に向けて輸送中です。",
			"ko":    "주문 번호 " + models.OrderSn + "이 운전자에 의해 배송되었으며 출발 지점으로 이동 중입니다.",
			"zh_CN": "您的訂單" + models.OrderSn + "司機已出發，正在趕往您的出發地",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		// 发送确认短信(队列)
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderRemind,
			DataByte: gjson.New(g.Map{
				"orderSn": models.OrderSn,
				"event":   "car_order_out",
			}).MustToJson(),
			Header: nil,
		})

		return
	})

}

func (s *sCarOrder) StartService(ctx context.Context, in *input_car.CarOrderStartServiceInp) (err error) {

	var models *entity.CarOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.DispatchStatus != "DONE" {
		err = gerror.New("调度状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderStartServiceFields{
			OrderStatus:     "SERVING",
			ActualStartTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "SERVING",
			ActionWay:   "SERVING",
			Remark:      "司机已到达出发地",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		//// 司机状态改为工作中
		//if _, err = dao.CarDriver.Ctx(ctx).
		//	WherePri(models.DriverId).
		//	Data(g.MapStrAny{
		//		dao.CarDriver.Columns().WorkStatus: "WORKING",
		//	}).Update(); err != nil {
		//	err = gerror.Wrap(err, "修改司机工作状态失败，请稍后重试！")
		//	return
		//}
		//
		//// 车辆状态改为工作中
		//if models.CarId > 0 {
		//	if _, err = dao.CarCar.Ctx(ctx).
		//		WherePri(models.CarId).
		//		Data(g.MapStrAny{
		//			dao.CarCar.Columns().WorkStatus: "WORKING",
		//		}).Update(); err != nil {
		//		err = gerror.Wrap(err, "修改车辆工作状态失败，请稍后重试！")
		//		return
		//	}
		//}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "司机达到接送点",
			"en":    "The driver arrived at the pick-up point",
			"ja":    "運転手がピックアップ地点に到着した",
			"ko":    "운전자가 픽업 지점에 도착했습니다.",
			"zh_CN": "司機達到接送點",
		}
		systemMessageContent := map[string]string{
			"zh":    "司机已经到达约定接送地点，如您未找到司机，请在行程内联系司机或客服。",
			"en":    "The driver has arrived at the agreed pick-up location. If you cannot find the driver, please contact the driver or customer service within the trip details.",
			"ja":    "ドライバーは予定の乗車場所に到着しました。ドライバーが見つからない場合は、乗車情報に記載されているドライバーまたはカスタマーサービスまでご連絡ください。",
			"ko":    "운전기사가 약속된 픽업 장소에 도착했습니다. 운전기사를 찾을 수 없는 경우, 여정 세부 정보에 기재된 운전기사 또는 고객센터로 문의해 주세요.",
			"zh_CN": "司機已到達約定接送地點，如您未找到司機，請在行程內聯絡司機或客服。",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		// 发送确认短信(队列)
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderRemind,
			DataByte: gjson.New(g.Map{
				"orderSn": models.OrderSn,
				"event":   "car_order_driver_arrive",
			}).MustToJson(),
			Header: nil,
		})

		return
	})

}

func (s *sCarOrder) EndService(ctx context.Context, in *input_car.CarOrderEndServiceInp) (err error) {

	var (
		models         *entity.CarOrder
		DriverInfo     *entity.CarDriver
		SettlementInfo *entity.CarSettlement
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "SERVING" {
		err = gerror.New("订单状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderEndServiceFields{
			OrderStatus:   "DONE",
			ActualEndTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "DONE",
			ActionWay:   "DONE",
			Remark:      "成功到达目的地",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 司机状态改为可约
		if _, err = dao.CarDriver.Ctx(ctx).
			WherePri(models.DriverId).
			Data(g.MapStrAny{
				dao.CarDriver.Columns().WorkStatus: "CANORDER",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改司机工作状态失败，请稍后重试！")
			return
		}

		// 车辆状态改为可约
		if models.CarId > 0 {
			if _, err = dao.CarCar.Ctx(ctx).
				WherePri(models.CarId).
				Data(g.MapStrAny{
					dao.CarCar.Columns().WorkStatus: "CANORDER",
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改车辆工作状态失败，请稍后重试！")
				return
			}
		}

		// 如果是异常单，则需要异常完成处理
		if models.AbnormalStatus == 2 {
			if _, err = s.Model(ctx).
				WherePri(in.Id).Data(g.Map{
				dao.CarOrder.Columns().AbnormalStatus: 3,
				dao.CarOrder.Columns().AbnormalTime:   gtime.Now(),
				dao.CarOrder.Columns().AbnormalReason: "后台已结束服务",
			}).OmitEmptyData().Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
		}

		// 记录分佣
		var cost string
		_ = dao.CarDriver.Ctx(ctx).WherePri(models.DriverId).Scan(&DriverInfo)
		if !g.IsEmpty(DriverInfo) {
			_ = dao.CarSettlement.Ctx(ctx).WherePri(DriverInfo.SettlementId).Scan(&SettlementInfo)
			if !g.IsEmpty(SettlementInfo) {
				cost = SettlementInfo.Cost
			}
		}
		settlementAmount := models.OrderAmount - models.NightAmount // 初始结算金额扣除深夜费
		if !g.IsEmpty(cost) {
			costArr := strings.Split(cost, ",")
			constIsCoupon := false
			constIsBal := false
			for _, costItem := range costArr {
				if gvar.New(costItem).Int() == 1 {
					constIsCoupon = true
				}
				if gvar.New(costItem).Int() == 2 {
					constIsBal = true
				}
			}
			if !constIsCoupon {
				settlementAmount = settlementAmount - models.CouponAmount
			}
			if !constIsBal {
				settlementAmount = settlementAmount - models.BalAmount
			}
		}
		// 结算金额 = 计算后的结算金额 + 深夜费
		settlementAmount = settlementAmount + models.NightAmount
		settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(models.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()
		if _, err = dao.CarOrder.Ctx(ctx).TX(tx).Data(g.Map{
			dao.CarOrder.Columns().SettlementAmount: settlementAmount,
		}).WherePri(in.Id).Update(); err != nil {
			return
		}

		// 转发到返利队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "服务结束",
			"en":    "Service ended",
			"ja":    "サービス終了",
			"ko":    "서비스 종료",
			"zh_CN": "服務結束",
		}
		systemMessageContent := map[string]string{
			"zh":    "司机已到达目的地，本次服务已结束。",
			"en":    "The driver has arrived at the destination; this service has ended.",
			"ja":    "ドライバーは目的地に到着しました。サービスは終了しました。",
			"ko":    "운전자가 목적지에 도착했습니다. 이 서비스는 종료되었습니다.",
			"zh_CN": "司機已到達目的地，本次服務已結束。",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "car",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/transferDetailScene",
			WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		if models.IsFx == "Y" {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
				DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
					OrderNo:      models.OrderSn,
					ChangeStatus: "COMPLETE",
				}).MustToJson(),
				Header: nil,
			})
		}

		return
	})

}

func (s *sCarOrder) Abnormal(ctx context.Context, in *input_car.CarOrderAbnormalInp) (err error) {

	var (
		models         *entity.CarOrder
		DriverInfo     *entity.CarDriver
		SettlementInfo *entity.CarSettlement
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	if models.AbnormalStatus != 2 {
		err = gerror.New("订单异常处理状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderAbnormalInp{
			AbnormalStatus:     3,
			AbnormalTime:       gtime.Now(),
			AbnormalOperatorId: in.AbnormalOperatorId,
			AbnormalReason:     in.AbnormalReason,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 如果当前订单状态为已完成DONE,则不进行后续操作
		if models.OrderStatus == "DONE" {
			return
		}

		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_car.CarOrderEndServiceFields{
			OrderStatus:   "DONE",
			ActualEndTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.CarOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CarOrderLog{
			OrderId:     in.Id,
			OrderStatus: "DONE",
			ActionWay:   "DONE",
			Remark:      "成功到达目的地",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 司机状态改为可约
		if models.DriverId > 0 {
			if _, err = dao.CarDriver.Ctx(ctx).
				WherePri(models.DriverId).
				Data(g.MapStrAny{
					dao.CarDriver.Columns().WorkStatus: "CANORDER",
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改司机工作状态失败，请稍后重试！")
				return
			}
		}

		// 车辆状态改为可约
		if models.CarId > 0 {
			if _, err = dao.CarCar.Ctx(ctx).
				WherePri(models.CarId).
				Data(g.MapStrAny{
					dao.CarCar.Columns().WorkStatus: "CANORDER",
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改车辆工作状态失败，请稍后重试！")
				return
			}
		}

		if models.DriverId > 0 {
			// 记录分佣
			_ = dao.CarDriver.Ctx(ctx).WherePri(models.DriverId).Scan(&DriverInfo)
			if !g.IsEmpty(DriverInfo) {
				_ = dao.CarSettlement.Ctx(ctx).WherePri(DriverInfo.SettlementId).Scan(&SettlementInfo)
				var cost string
				if !g.IsEmpty(SettlementInfo) {
					cost = SettlementInfo.Cost
				}
				settlementAmount := models.OrderAmount
				if !g.IsEmpty(cost) {
					costArr := strings.Split(cost, ",")
					constIsCoupon := false
					constIsBal := false
					for _, costItem := range costArr {
						if gvar.New(costItem).Int() == 1 {
							constIsCoupon = true
						}
						if gvar.New(costItem).Int() == 2 {
							constIsBal = true
						}
					}
					if !constIsCoupon {
						settlementAmount = settlementAmount - models.CouponAmount
					}
					if !constIsBal {
						settlementAmount = settlementAmount - models.BalAmount
					}
				}
				settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(models.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()
				if _, err = dao.CarOrder.Ctx(ctx).TX(tx).Data(g.Map{
					dao.CarOrder.Columns().SettlementAmount: settlementAmount,
				}).WherePri(in.Id).Update(); err != nil {
					return
				}
			}

		}

		// 转发到返利队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})

		return
	})

}

func (s *sCarOrder) ExportOrder(ctx context.Context, in *input_car.CarOrderExportInp) (err error) {
	var (
		lastInsertId int64
	)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		resultJson, _ := json.Marshal(in)

		if lastInsertId, err = dao.OrderExport.Ctx(ctx).
			Data(entity.OrderExport{
				Scene:     4,
				Condition: string(resultJson),
			}).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		return
	})

	if err != nil {
		return
	}

	// 导出(队列)
	_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderExport,
		DataByte:     gvar.New(lastInsertId).Bytes(),
		Header:       nil,
	})

	return
}

func (s *sCarOrder) StartExport(ctx context.Context, in *input_car.CarOrderExportInp) (path string, err error) {
	var (
		ChangeList  []*input_car.CarOrderExportModel
		OrderListIn *input_car.CarOrderListInp
	)

	OrderListIn = &input_car.CarOrderListInp{
		PageReq: input_form.PageReq{
			Pagination: false,
		},
		MemberId:      in.MemberId,
		OrderSn:       in.OrderSn,
		ServiceType:   in.ServiceType,
		OrderType:     in.OrderType,
		MemberSearch:  in.MemberSearch,
		OrderStatus:   in.OrderStatus,
		DriverName:    in.DriverName,
		BookStartTime: in.BookStartTime,
		CreatedAt:     in.CreatedAt,
	}
	//if !g.IsEmpty(in.BookStartTime) {
	//	for _, timeItem := range strings.Split(in.BookStartTime, ",") {
	//		fmt.Println(timeItem)
	//		tsMilli := gconv.Int64(strings.TrimSpace(timeItem))
	//		fmt.Println(tsMilli)
	//		OrderListIn.BookStartTime = append(OrderListIn.BookStartTime, gtime.New(time.UnixMilli(tsMilli)))
	//	}
	//}
	//if !g.IsEmpty(in.CreatedAt) {
	//	for _, timeItem := range strings.Split(in.CreatedAt, ",") {
	//		tsMilli := gconv.Int64(strings.TrimSpace(timeItem))
	//		OrderListIn.CreatedAt = append(OrderListIn.BookStartTime, gtime.New(time.UnixMilli(tsMilli)))
	//	}
	//}
	list, _, err := s.List(ctx, OrderListIn)
	if err != nil {
		return
	}
	for _, item := range list {
		itemData := &input_car.CarOrderExportModel{
			OrderSn:           item.OrderSn,
			MemberNo:          item.PmsMemberMemberNo,
			BookingName:       item.BookingName,
			BookStartTime:     item.BookDate + " " + item.BookTime,
			OrderTime:         item.CreatedAt.Format("Y-m-d H:i:s"),
			TotalAmount:       item.OrderAmount,
			RefundTotalAmount: item.RefundAmount,
			RefundTime:        item.RefundTime.Format("Y-m-d H:i:s"),
		}
		if item.ServiceType == "PICKUP" {
			itemData.ServiceType = "接机"
		} else if item.ServiceType == "DELIVERY" {
			itemData.ServiceType = "送机"
		} else {
			itemData.ServiceType = "包车"
		}
		for _, TransactionDetail := range item.TransactionDetail {
			if TransactionDetail.PayStatus == "DONE" {
				switch TransactionDetail.PayType {
				case "BAL":
					itemData.PointsPayment += TransactionDetail.PayAmount
					break
				case "COUPON":
					itemData.CouponPayment += TransactionDetail.PayAmount
					break
				case "WeChatPay":
					itemData.PaycloudWechatPay += TransactionDetail.PayAmount
					break
				case "Alipay+":
					itemData.PaycloudAlipayPay += TransactionDetail.PayAmount
					break
				case "Paypal":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "PaypalCard":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "StripeCard":
					itemData.StripeCreditPay += TransactionDetail.PayAmount
					break
				case "WeChatMiniPay":
					itemData.MlilifeWeChatMiniPay += TransactionDetail.PayAmount
					break
				}
			}
		}

		for _, TransactionRefundDetail := range item.TransactionRefundDetail {
			switch TransactionRefundDetail.RefundType {
			case "BAL":
				itemData.PointsRefund += TransactionRefundDetail.RefundAmount
				break
			case "COUPON":
				itemData.CouponRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatPay":
				itemData.PaycloudWechatRefund += TransactionRefundDetail.RefundAmount
				break
			case "Alipay+":
				itemData.PaycloudAlipayRefund += TransactionRefundDetail.RefundAmount
				break
			case "Paypal":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "PaypalCard":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "StripeCard":
				itemData.StripeCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatMiniPay":
				itemData.MlilifeWeChatMiniRefund += TransactionRefundDetail.RefundAmount
				break
			}
		}
		ChangeList = append(ChangeList, itemData)
	}
	tags, err := convert.GetEntityDescTags(input_car.CarOrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出接送机订单-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("接送机订单")
		exports   []input_car.CarOrderExportModel
	)

	if err = gconv.Scan(ChangeList, &exports); err != nil {
		return
	}

	path, err = excel.ExportByStructsFile(ctx, tags, exports, fileName, sheetName)

	return
}

func (s *sCarOrder) ExportList(ctx context.Context, in *input_car.CarOrderExportListInp) (list []*input_car.CarOrderExportListModel, totalCount int, err error) {
	mod := dao.OrderExport.Ctx(ctx)

	mod = mod.Fields(input_car.CarOrderExportListModel{})

	mod = mod.Where(dao.OrderExport.Columns().Scene, 4)

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.OrderExport.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		v.Path = g.Cfg().MustGet(ctx, "localUploadDomain").String() + "/" + v.Path
	}

	return
}
