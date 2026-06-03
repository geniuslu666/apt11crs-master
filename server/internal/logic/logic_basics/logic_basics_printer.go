package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/library/xpyun"
	model2 "APT/internal/library/xpyun/model"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	openApi "github.com/Qzm6826/yly-go-sdk"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type sBasicsPrinter struct{}

func NewBasicsPrinter() *sBasicsPrinter {
	return &sBasicsPrinter{}
}

func init() {
	service.RegisterBasicsPrinter(NewBasicsPrinter())
}

// Model 打印机ORM模型
func (s *sBasicsPrinter) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysPrinter.Ctx(ctx), option...)
}

// List 获取打印机列表
func (s *sBasicsPrinter) List(ctx context.Context, in *input_basics.PrinterListInp) (list []*input_basics.PrinterListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_basics.PrinterListModel{})

	// 打印机名称
	if !g.IsEmpty(in.PrinterName) {
		mod = mod.WhereLike(dao.SysPrinter.Columns().PrinterName, "%"+in.PrinterName+"%")
	}

	// 第三方应用ID
	if !g.IsEmpty(in.ClientId) {
		mod = mod.WhereLike(dao.SysPrinter.Columns().ClientId, "%"+in.ClientId+"%")
	}

	// 第三方应用秘钥
	if !g.IsEmpty(in.ClientSecret) {
		mod = mod.WhereLike(dao.SysPrinter.Columns().ClientSecret, "%"+in.ClientSecret+"%")
	}

	// 机器码
	if !g.IsEmpty(in.MachineCode) {
		mod = mod.WhereLike(dao.SysPrinter.Columns().MachineCode, "%"+in.MachineCode+"%")
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SysPrinter.Columns().Sort).OrderDesc(dao.SysPrinter.Columns().Id)
	//mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取打印机列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取打印机列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Edit 修改/新增打印机
func (s *sBasicsPrinter) Edit(ctx context.Context, in *input_basics.PrinterEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PrinterUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改打印机失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PrinterInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增打印机失败，请稍后重试！")
		}

		return
	})
}

// Delete 删除打印机
func (s *sBasicsPrinter) Delete(ctx context.Context, in *input_basics.PrinterDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除打印机失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取打印机最大排序
func (s *sBasicsPrinter) MaxSort(ctx context.Context, in *input_basics.PrinterMaxSortInp) (res *input_basics.PrinterMaxSortModel, err error) {
	if err = dao.SysPrinter.Ctx(ctx).Fields(dao.SysPrinter.Columns().Sort).OrderDesc(dao.SysPrinter.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取打印机最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_basics.PrinterMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取打印机指定信息
func (s *sBasicsPrinter) View(ctx context.Context, in *input_basics.PrinterViewInp) (res *input_basics.PrinterViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取打印机信息，请稍后重试！")
		return
	}

	return
}

// Status 更新司机状态
func (s *sBasicsPrinter) Status(ctx context.Context, in *input_basics.PrinterStatusInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.SysPrinter.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新打印机状态失败，请稍后重试！")
			return
		}

		return
	})
}

// PrinterCarOrder 打印接送机订单
func (s *sBasicsPrinter) PrinterCarOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error) {

	// 获取打印配置
	var PrinterConfig *model.PrinterSettingConfig
	if PrinterConfig, err = service.BasicsConfig().GetCarPrinterSettingConfig(ctx); err != nil {
		return
	}

	// 判断是否开启打印功能
	if PrinterConfig.IsOpen == 2 {
		err = gerror.New("未开启打印功能，请稍后重试！")
		return
	}

	if g.IsEmpty(PrinterConfig.PrinterIds) {
		err = gerror.New("未配置打印机，请稍后重试！")
		return
	}

	// 获取订单信息
	var carOrder *input_car.CarOrderViewModel
	if !g.IsEmpty(in.OrderId) {
		if err = dao.CarOrder.Ctx(ctx).Unscoped().WherePri(in.OrderId).WithAll().Scan(&carOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	} else {
		if err = dao.CarOrder.Ctx(ctx).Unscoped().Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).WithAll().Scan(&carOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	}

	if carOrder == nil {
		err = gerror.New("订单不存在")
		return
	}

	// 手动加载已删除的会员信息
	if carOrder.MemberId > 0 && carOrder.MemberDetail == nil {
		var member struct {
			Id       int    `json:"id"`
			FullName string `json:"fullName"`
			MemberNo string `json:"memberNo"`
		}
		if err = dao.PmsMember.Ctx(ctx).Unscoped().
			Fields("id, full_name, member_no").
			Where(dao.PmsMember.Columns().Id, carOrder.MemberId).Scan(&member); err == nil {
			carOrder.MemberDetail = &struct {
				gmeta.Meta `orm:"table:hg_pms_member"`
				Id         int    `json:"id"    orm:"id"      dc:"id"`
				FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
				MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
			}{Id: member.Id, FullName: member.FullName, MemberNo: member.MemberNo}
		}
	}

	// 手动加载已删除的司机信息
	if carOrder.DriverId > 0 && carOrder.DriverDetail == nil {
		var driver *entity.CarDriver
		if err = dao.CarDriver.Ctx(ctx).Unscoped().Where(dao.CarDriver.Columns().Id, carOrder.DriverId).Scan(&driver); err == nil && driver != nil {
			carOrder.DriverDetail = &struct {
				gmeta.Meta `orm:"table:hg_car_driver"`
				*entity.CarDriver
			}{CarDriver: driver}
		}
	}

	// 手动加载已删除的出发地址
	if carOrder.StartAddressId > 0 && carOrder.StartServiceAddress == nil {
		var addr *entity.CarAddress
		if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, carOrder.StartAddressId).Scan(&addr); err == nil && addr != nil {
			carOrder.StartServiceAddress = &struct {
				gmeta.Meta `orm:"table:hg_car_address"`
				*entity.CarAddress
			}{CarAddress: addr}
		}
	}

	// 手动加载已删除的目的地址
	if carOrder.EndAddressId > 0 && carOrder.EndServiceAddress == nil {
		var addr *entity.CarAddress
		if err = dao.CarAddress.Ctx(ctx).Unscoped().Where(dao.CarAddress.Columns().Id, carOrder.EndAddressId).Scan(&addr); err == nil && addr != nil {
			carOrder.EndServiceAddress = &struct {
				gmeta.Meta `orm:"table:hg_car_address"`
				*entity.CarAddress
			}{CarAddress: addr}
		}
	}

	// 手动加载已删除的车辆信息
	if carOrder.CarId > 0 && carOrder.CarDetail == nil {
		var car *entity.CarCar
		if err = dao.CarCar.Ctx(ctx).Unscoped().Where(dao.CarCar.Columns().Id, carOrder.CarId).Scan(&car); err == nil && car != nil {
			carOrder.CarDetail = &struct {
				gmeta.Meta `orm:"table:hg_car_car"`
				*entity.CarCar
				CarTypeDetail *struct {
					gmeta.Meta `orm:"table:hg_car_car_type"`
					*entity.CarCarType
				} `json:"carTypeDetail" orm:"with:id=type_id"`
			}{CarCar: car}
		}
	}

	// 手动加载已删除的服务信息
	if carOrder.ServiceId > 0 && carOrder.ServiceDetail == nil {
		var svc *entity.CarService
		if err = dao.CarService.Ctx(ctx).Unscoped().Where(dao.CarService.Columns().Id, carOrder.ServiceId).Scan(&svc); err == nil && svc != nil {
			carOrder.ServiceDetail = &struct {
				gmeta.Meta `orm:"table:hg_car_service"`
				*entity.CarService
				CarTypeDetail *struct {
					gmeta.Meta `orm:"table:hg_car_car_type"`
					*entity.CarCarType
				} `json:"carTypeDetail" orm:"with:id=car_type_id"`
			}{CarService: svc}
		}
	}

	// 获取打印机信息，默认一号打印机 todo	mod = mod.
	var printerArr []*entity.SysPrinter
	if err = dao.SysPrinter.Ctx(ctx).
		WhereIn(dao.SysPrinter.Columns().Id, strings.Split(PrinterConfig.PrinterIds, ",")).
		WithAll().
		Scan(&printerArr); err != nil {
		err = gerror.Wrap(err, "获取打印机信息失败，请稍后重试！")
		return
	}

	// 循环打印机, 判断打印机状态，状态为1时打印
	for _, printer := range printerArr {
		if printer.Status != 1 {
			continue
		}

		var content = ""
		var contentInfo *input_basics.GetPrintOrderContentModel
		if printer.PrinterType == "YILINK" {

			// 打印内容-易联云
			contentInfo, err = s.GetYILINKCarOrderContent(ctx, &input_basics.GetYILINKCarOrderContentInp{
				CarOrder:   carOrder,
				PrintTimes: PrinterConfig.PrinterNum,
				IsRefund:   in.IsRefund,
			})
			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content

		} else if printer.PrinterType == "XPYUN503" {
			// 打印内容-芯烨云
			contentInfo, err = s.GetXYUNCarOrderContent(ctx, &input_basics.GetYILINKCarOrderContentInp{
				CarOrder:   carOrder,
				PrintTimes: PrinterConfig.PrinterNum,
				IsRefund:   in.IsRefund,
			})

			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content
		}

		// 打印
		_ = s.PrintOrderContent(ctx, &input_basics.PrintContentInp{
			Content:    content,
			OrderSn:    carOrder.OrderSn,
			PrintTimes: PrinterConfig.PrinterNum,
			Printer:    printer,
		})

	}

	return
}

// PrinterSpaOrder 打印按摩订单
func (s *sBasicsPrinter) PrinterSpaOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error) {

	// 获取打印配置
	var PrinterConfig *model.PrinterSettingConfig
	if PrinterConfig, err = service.BasicsConfig().GetSpaPrinterSettingConfig(ctx); err != nil {
		return
	}

	// 判断是否开启打印功能
	if PrinterConfig.IsOpen == 2 {
		err = gerror.New("未开启打印功能，请稍后重试！")
		return
	}

	if g.IsEmpty(PrinterConfig.PrinterIds) {
		err = gerror.New("未配置打印机，请稍后重试！")
		return
	}

	// 获取订单信息
	var spaOrder *input_spa.SpaOrderViewModel
	if !g.IsEmpty(in.OrderId) {
		if err = dao.SpaOrder.Ctx(ctx).WherePri(in.OrderId).WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&spaOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	} else {
		if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().OrderSn, in.OrderSn).WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&spaOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	}

	// 获取打印机信息，默认一号打印机 todo	mod = mod.
	var printerArr []*entity.SysPrinter
	if err = dao.SysPrinter.Ctx(ctx).
		WhereIn(dao.SysPrinter.Columns().Id, strings.Split(PrinterConfig.PrinterIds, ",")).
		WithAll().
		Scan(&printerArr); err != nil {
		err = gerror.Wrap(err, "获取打印机信息失败，请稍后重试！")
		return
	}

	// 循环打印机, 判断打印机状态，状态为1时打印
	for _, printer := range printerArr {
		if printer.Status != 1 {
			continue
		}

		var content = ""
		var contentInfo *input_basics.GetPrintOrderContentModel
		if printer.PrinterType == "YILINK" {

			// 打印内容-易联云
			contentInfo, err = s.GetYILINKSpaOrderContent(ctx, &input_basics.GetYILINKSpaOrderContentInp{
				SpaOrder:   spaOrder,
				PrintTimes: PrinterConfig.PrinterNum,
			})
			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content

		} else if printer.PrinterType == "XPYUN503" {
			// 打印内容-芯烨云
			contentInfo, err = s.GetXYUNSpaOrderContent(ctx, &input_basics.GetYILINKSpaOrderContentInp{
				SpaOrder:   spaOrder,
				PrintTimes: PrinterConfig.PrinterNum,
			})

			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content
		}

		// 打印
		_ = s.PrintOrderContent(ctx, &input_basics.PrintContentInp{
			Content:    content,
			OrderSn:    spaOrder.OrderSn,
			PrintTimes: PrinterConfig.PrinterNum,
			Printer:    printer,
		})

	}

	return
}

// PrinterFoodOrder 打印餐厅订单
func (s *sBasicsPrinter) PrinterFoodOrder(ctx context.Context, in *input_basics.PrinterCarOrderInp) (err error) {

	// 获取打印配置
	var PrinterConfig *model.PrinterSettingConfig
	if PrinterConfig, err = service.BasicsConfig().GetFoodPrinterSettingConfig(ctx); err != nil {
		return
	}

	// 判断是否开启打印功能
	if PrinterConfig.IsOpen == 2 {
		err = gerror.New("未开启打印功能，请稍后重试！")
		return
	}

	if g.IsEmpty(PrinterConfig.PrinterIds) {
		err = gerror.New("未配置打印机，请稍后重试！")
		return
	}

	// 获取订单信息
	var foodOrder *input_food.FoodOrderViewModel
	if !g.IsEmpty(in.OrderId) {
		if err = dao.FoodOrder.Ctx(ctx).WherePri(in.OrderId).WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&foodOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	} else {
		if err = dao.FoodOrder.Ctx(ctx).Where(dao.FoodOrder.Columns().OrderSn, in.OrderSn).WithAll().Hook(hook.PmsFindLanguageValueHook).Scan(&foodOrder); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}
	}

	// 获取打印机信息，默认一号打印机 todo
	var printerArr []*entity.SysPrinter
	if err = dao.SysPrinter.Ctx(ctx).
		WhereIn(dao.SysPrinter.Columns().Id, strings.Split(PrinterConfig.PrinterIds, ",")).
		WithAll().
		Scan(&printerArr); err != nil {
		err = gerror.Wrap(err, "获取打印机信息失败，请稍后重试！")
		return
	}

	// 循环打印机, 判断打印机状态，状态为1时打印
	for _, printer := range printerArr {
		if printer.Status != 1 {
			continue
		}

		var content = ""
		var contentInfo *input_basics.GetPrintOrderContentModel
		if printer.PrinterType == "YILINK" {

			// 打印内容-易联云
			contentInfo, err = s.GetYILINKFoodOrderContent(ctx, &input_basics.GetYILINKFoodOrderContentInp{
				FoodOrder:  foodOrder,
				PrintTimes: PrinterConfig.PrinterNum,
			})
			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content

		} else if printer.PrinterType == "XPYUN503" {
			// 打印内容-芯烨云
			contentInfo, err = s.GetXYUNFoodOrderContent(ctx, &input_basics.GetYILINKFoodOrderContentInp{
				FoodOrder:  foodOrder,
				PrintTimes: PrinterConfig.PrinterNum,
			})

			if err != nil {
				err = gerror.Wrap(err, "获取打印内容失败，请稍后重试！")
				return
			}
			content = contentInfo.Content
		}

		// 打印
		_ = s.PrintOrderContent(ctx, &input_basics.PrintContentInp{
			Content:    content,
			OrderSn:    foodOrder.OrderSn,
			PrintTimes: PrinterConfig.PrinterNum,
			Printer:    printer,
		})
	}

	return
}

func (s *sBasicsPrinter) PrintOrderContent(ctx context.Context, in *input_basics.PrintContentInp) (err error) {

	if in.Printer.PrinterType == "YILINK" {

		// 开始打印
		// 新建一个配置实例
		conf := openApi.NewConfig(in.Printer.ClientId, in.Printer.ClientSecret)

		// 获取 token 并设置
		oauth := openApi.NewAuthClient(conf)
		tokenData := oauth.GetAccessToken()
		conf.SetToken(tokenData)

		// 新建一个 API 实例
		client := openApi.NewClient(conf)

		// 添加一个打印机。未绑定打印机，需先调用此方法
		//_, _ = client.SetPrinter.AddPrinter(printer.MachineCode, printer.MachineKey, printer.PrinterName)

		// 调用服务 API
		_, _ = client.PrintService.TextPrint(in.Printer.MachineCode, in.Content, in.OrderSn, 0)
	} else if in.Printer.PrinterType == "XPYUN503" {

		request := model2.PrintRequest{}
		request.User = in.Printer.ClientId
		request.UserKey = in.Printer.ClientSecret

		//*必填*：打印机编号
		request.Sn = in.Printer.MachineCode

		request.GenerateSign()

		//*必填*：打印内容,不能超过12K
		request.Content = in.Content

		//打印份数，默认为1
		request.Copies = in.PrintTimes

		//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
		request.Voice = 2
		//打印模式：
		//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
		//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
		request.Mode = 1

		// 打印
		result := xpyun.XpYunPrint(ctx, &request)

		//序列化
		reslutJson, _ := json.Marshal(result.Content)
		var msg = fmt.Sprintf("response result: %+v", string(reslutJson))

		g.Log().Info(ctx, "--------打印结果----------")
		g.Log().Info(ctx, msg)
	}
	return
}

func (s *sBasicsPrinter) GetYILINKCarOrderContent(ctx context.Context, in *input_basics.GetYILINKCarOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "住一自营"
	var pickUpSign = "否"
	if in.CarOrder.OrderType == "INNN" {
		orderType = "INNN"
	}
	if in.CarOrder.ServiceType == "PICKUP" && in.CarOrder.PickUpSign == "Y" {
		pickUpSign = "是"
	}

	var content = ""
	content = content +
		"<MN>" + gvar.New(in.PrintTimes).String() + "</MN>"
	if (in.CarOrder.OrderStatus == "CANCEL" && in.CarOrder.PayStatus == "REFUND") || in.IsRefund == 1 {
		content = content +
			"<FS><center>" + "-------订单已取消-------" + "</center></FS>" + "\\r\\n"
	}
	content = content +
		"<FS><center>" + orderType + "</center></FS>" + "\\r\\n" +
		//"<FH>" +
		"订单编号：" + in.CarOrder.OrderSn + "\\r\\n" + "\\r" +
		"出发地：" + in.CarOrder.StartServiceAddress.Name + "\\r\\n" +
		"到达地：" + in.CarOrder.EndServiceAddress.Name + "\\r\\n" +
		"用车时间：" + gvar.New(in.CarOrder.BookStartTime).String() + "\\r\\n" +
		"\\r\\n" +
		"订单金额：" + gvar.New(in.CarOrder.OrderAmount).String() + "JPY\\r\\n" +
		"支付时间：" + gvar.New(in.CarOrder.PayTime).String() + "\\r\\n" +
		"\\r\\n" +
		"预订人信息：\\r\\n" +
		"预订人：" + in.CarOrder.BookingName + "\\r\\n" +
		"电话：" + in.CarOrder.PhoneArea + in.CarOrder.BookingMobile + "\\r\\n" +
		"\\r\\n" +
		"人数：\\r\\n" +
		"成人：" + gvar.New(in.CarOrder.AdultNum).String() + "\\r\\n" +
		"儿童：" + gvar.New(in.CarOrder.ChildNum).String() + "\\r\\n" +
		"\\r\\n" +
		"附加服务：\\r\\n" +
		"举牌服务：" + pickUpSign + "\\r\\n" +
		"婴儿座椅：" + gvar.New(in.CarOrder.ChildSeatAddNum).String() + "\\r\\n" +
		"\\r\\n" +
		"紧急联系人：\\r\\n" +
		"联系人：" + in.CarOrder.EmergencyName + "\\r\\n" +
		"电话：" + in.CarOrder.EmergencyPhoneArea + in.CarOrder.EmergencyMobile + "\\r\\n" +
		"备注：" + in.CarOrder.MemberMessage + "\\r\\n" +
		//"</FH>" +
		"\\r\\n"
	if in.CarOrder.DispatchStatus != "WAIT" && in.CarOrder.PayStatus == "HAVE_PAID" {
		content = content +
			"<FS>司机：" + "\\r\\n" +
			"司机姓名：" + in.CarOrder.DriverDetail.Name + "\\r\\n" +
			"司机手机：" + in.CarOrder.DriverDetail.PhoneArea + in.CarOrder.DriverDetail.Phone + "</FS>\\r\\n"
	}
	if (in.CarOrder.OrderStatus == "CANCEL" && in.CarOrder.PayStatus == "REFUND") || in.IsRefund == 1 {
		content = content +
			"<FS><center>" + "-------订单已取消-------" + "</center></FS>" + "\\r\\n"
	}
	content = content + "\\r\\n" + "\\r\\n"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = content

	return
}

func (s *sBasicsPrinter) GetXYUNCarOrderContent(ctx context.Context, in *input_basics.GetYILINKCarOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "住一自营"
	var pickUpSign = "否"
	if in.CarOrder.OrderType == "INNN" {
		orderType = "INNN"
	}
	if in.CarOrder.ServiceType == "PICKUP" && in.CarOrder.PickUpSign == "Y" {
		pickUpSign = "是"
	}

	// 打印内容
	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"

	if (in.CarOrder.OrderStatus == "CANCEL" && in.CarOrder.PayStatus == "REFUND") || in.IsRefund == 1 {
		printContent = printContent + "<CB>" + "---订单已取消---" + "<BR>"
		printContent = printContent + "<BR>"
	}

	printContent = printContent + "<CB>" + orderType + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.New(in.CarOrder.PayTime).Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>订单编号:" + in.CarOrder.OrderSn + "<BR>"
	printContent = printContent + "出发地：" + in.CarOrder.StartServiceAddress.Name + "<BR>"
	printContent = printContent + "到达地：" + in.CarOrder.EndServiceAddress.Name + "<BR>"
	printContent = printContent + "用车时间：" + gvar.New(in.CarOrder.BookStartTime).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "订单金额：" + gvar.New(in.CarOrder.OrderAmount).String() + "JPY<BR>"
	printContent = printContent + "支付时间：" + gvar.New(in.CarOrder.PayTime).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "预订人信息：" + "<BR>"
	printContent = printContent + "预订人：" + in.CarOrder.BookingName + "<BR>"
	printContent = printContent + "电话：" + in.CarOrder.PhoneArea + in.CarOrder.BookingMobile + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "人数：" + "<BR>"
	printContent = printContent + "成人：" + gvar.New(in.CarOrder.AdultNum).String() + "<BR>"
	printContent = printContent + "儿童：" + gvar.New(in.CarOrder.ChildNum).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "附加服务：" + "<BR>"
	printContent = printContent + "举牌服务：" + pickUpSign + "<BR>"
	printContent = printContent + "婴儿座椅：" + gvar.New(in.CarOrder.ChildSeatAddNum).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "紧急联系人：" + "<BR>"
	printContent = printContent + "联系人：" + in.CarOrder.EmergencyName + "<BR>"
	printContent = printContent + "电话：" + in.CarOrder.EmergencyPhoneArea + in.CarOrder.EmergencyMobile + "<BR>"
	printContent = printContent + "备注：" + in.CarOrder.MemberMessage + "<BR>"

	if in.CarOrder.DispatchStatus != "WAIT" && in.CarOrder.PayStatus == "HAVE_PAID" {
		printContent = printContent + "<BR>"
		printContent = printContent + "司机：" + "<BR>"
		printContent = printContent + "司机姓名：" + in.CarOrder.DriverDetail.Name + "<BR>"
		printContent = printContent + "司机手机：" + in.CarOrder.DriverDetail.PhoneArea + in.CarOrder.DriverDetail.Phone + "<BR>"
	}
	printContent = printContent + "<BR>"
	if (in.CarOrder.OrderStatus == "CANCEL" && in.CarOrder.PayStatus == "REFUND") || in.IsRefund == 1 {
		printContent = printContent + "<CB>" + "---订单已取消---" + "<BR>"
		printContent = printContent + "<BR>"
	}
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR><BR>"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = printContent

	return
}

func (s *sBasicsPrinter) GetYILINKSpaOrderContent(ctx context.Context, in *input_basics.GetYILINKSpaOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "按摩"
	var serviceType = "到店"
	if in.SpaOrder.ServiceType != 1 {
		serviceType = "上门"
	}

	var content = ""
	content = content +
		"<MN>" + gvar.New(in.PrintTimes).String() + "</MN>" +
		"<FS><center>" + orderType + "</center></FS>" + "\\r\\n" +
		//"<FH>" +
		"订单编号：" + in.SpaOrder.OrderSn + "\\r\\n" +
		"订单金额：" + gvar.New(in.SpaOrder.OrderAmount).String() + "JPY\\r\\n" +
		"支付时间：" + gvar.New(in.SpaOrder.PayTime).String() + "\\r\\n" +
		"\\r" +
		"服务方式：" + serviceType + "\\r\\n"
	if in.SpaOrder.ServiceType == 2 {
		content = content +
			"服务地址：" + in.SpaOrder.PropertyDetail.Name + "\\r\\n" +
			"房间号：" + in.SpaOrder.RoomNo + "\\r\\n"
		if !g.IsEmpty(in.SpaOrder.PropertyDetail.AccessPass) {
			content = content +
				"门禁密码：" + in.SpaOrder.PropertyDetail.AccessPass + "\\r\\n"
		} else {
			content = content +
				"门禁密码：" + "无" + "\\r\\n"
		}
	}
	content = content +
		"预约时间：" + gvar.New(in.SpaOrder.BookStartTime).String() + "\\r\\n" +
		"\\r\\n" +
		"预约项目：\\r\\n" +
		"\\t" + in.SpaOrder.ServiceDetail.Name + "\\r\\n" +
		"\\t" + in.SpaOrder.GoodsDetail.GoodsName + " *" + gvar.New(in.SpaOrder.GoodsNum).String() + "次\\r\\n" +
		"\\r\\n" +
		"预订人信息：\\r\\n" +
		"预订人：" + in.SpaOrder.BookingName + "\\r\\n" +
		"电话：" + in.SpaOrder.PhoneArea + in.SpaOrder.BookingMobile + "\\r\\n" +
		"\\r\\n" +
		"备注：" + in.SpaOrder.MemberMessage + "\\r\\n" +
		//"</FH>" +
		"\\r\\n"
	if in.SpaOrder.DispatchStatus == "DONE" && in.SpaOrder.PayStatus == "HAVE_PAID" {
		content = content + "<FS>技师：" + "\\r\\n"
		for _, technician := range in.SpaOrder.TechnicianList {
			content = content +
				"技师姓名：" + technician.TechnicianDetail.Name + "\\r\\n" +
				"技师手机：" + technician.TechnicianDetail.PhoneArea + technician.TechnicianDetail.Phone
		}
		content = content + "</FS>\\r\\n"
	}

	content = content + "\\r\\n"
	content = content + "<center>---扫码导航---</center>\\r\\n"
	content = content + "\\r\\n"
	content = content + "<QR>" + "https://www.google.com/maps?q=" + in.SpaOrder.PropertyDetail.GgLat + "," + in.SpaOrder.PropertyDetail.GgLng + "</QR>"

	content = content + "\\r\\n" + "\\r\\n"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = content

	return
}

func (s *sBasicsPrinter) GetXYUNSpaOrderContent(ctx context.Context, in *input_basics.GetYILINKSpaOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "按摩"
	var serviceType = "到店"
	if in.SpaOrder.ServiceType != 1 {
		serviceType = "上门"
	}

	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"
	printContent = printContent + "<CB>" + orderType + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.New(in.SpaOrder.PayTime).Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>订单编号:" + in.SpaOrder.OrderSn + "<BR>"
	printContent = printContent + "订单金额：" + gvar.New(in.SpaOrder.OrderAmount).String() + "JPY<BR>"
	printContent = printContent + "支付时间：" + gvar.New(in.SpaOrder.PayTime).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "服务方式：" + serviceType + "<BR>"
	if in.SpaOrder.ServiceType == 2 {
		printContent = printContent + "服务地址：" + in.SpaOrder.PropertyDetail.Name + "<BR>"
		printContent = printContent + "房间号：" + in.SpaOrder.RoomNo + "<BR>"
		if !g.IsEmpty(in.SpaOrder.PropertyDetail.AccessPass) {
			printContent = printContent + "门禁密码：" + in.SpaOrder.PropertyDetail.AccessPass + "<BR>"
		} else {
			printContent = printContent + "门禁密码：无<BR>"
		}
	}
	printContent = printContent + "预约时间：" + gvar.New(in.SpaOrder.BookStartTime).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "预约项目：" + "<BR>"
	printContent = printContent + in.SpaOrder.ServiceDetail.Name + "<BR>"
	printContent = printContent + in.SpaOrder.GoodsDetail.GoodsName + " *" + gvar.New(in.SpaOrder.GoodsNum).String() + "次<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "预订人信息：" + "<BR>"
	printContent = printContent + "预订人：" + in.SpaOrder.BookingName + "<BR>"
	printContent = printContent + "电话：" + in.SpaOrder.PhoneArea + in.SpaOrder.BookingMobile + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "备注：" + in.SpaOrder.MemberMessage + "<BR>"

	if in.SpaOrder.DispatchStatus == "DONE" && in.SpaOrder.PayStatus == "HAVE_PAID" {
		printContent = printContent + "<BR>"
		printContent = printContent + "技师：" + "<BR>"
		for _, technician := range in.SpaOrder.TechnicianList {
			printContent = printContent +
				"技师姓名：" + technician.TechnicianDetail.Name + "<BR>" +
				"技师手机：" + technician.TechnicianDetail.PhoneArea + technician.TechnicianDetail.Phone + "<BR>"
		}
	}
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---扫码导航---<BR>"
	printContent = printContent + " <QRCODE s=6 e=L l=center>" + "https://www.google.com/maps?q=" + in.SpaOrder.PropertyDetail.GgLat + "," + in.SpaOrder.PropertyDetail.GgLng + "</QRCODE>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR><BR>"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = printContent

	return
}

func (s *sBasicsPrinter) GetYILINKFoodOrderContent(ctx context.Context, in *input_basics.GetYILINKFoodOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "餐厅"
	var content = ""
	content = content +
		"<MN>" + gvar.New(in.PrintTimes).String() + "</MN>" +
		"<FS><center>" + orderType + "</center></FS>" + "\\r\\n" +
		//"<FH>" +
		"订单编号：" + in.FoodOrder.OrderSn + "\\r\\n" +
		"订单金额：" + gvar.New(in.FoodOrder.OrderAmount).String() + "JPY\\r\\n"

	if in.FoodOrder.OrderType == "CRS" {
		content = content +
			"支付定金：" + gvar.New(in.FoodOrder.DepositAmount).String() + "\\r\\n" +
			"定金支付：" + gvar.New(in.FoodOrder.DepositPayTime).String() + "\\r\\n"
		if in.FoodOrder.PayStep == "REMAIN" {
			content = content +
				"尾款支付：" + gvar.New(in.FoodOrder.PayTime).String() + "\\r\\n"
		}
	} else {
		content = content +
			"支付时间：" + gvar.New(in.FoodOrder.PayTime).String() + "\\r\\n"
	}
	content = content +
		"\\r" +
		"餐厅：" + in.FoodOrder.RestaurantDetail.Name + "\\r\\n" +
		"套餐：" + in.FoodOrder.GoodsDetail.GoodsName + "\\r\\n" +
		"\\r\\n" +
		"预订人信息：\\r\\n" +
		"预订人：" + in.FoodOrder.BookingName + "\\r\\n" +
		"电话：" + in.FoodOrder.PhoneArea + in.FoodOrder.BookingMobile + "\\r\\n" +
		"人数：" + gvar.New(in.FoodOrder.BookingCount).String() + "\\r\\n" +
		"\\r\\n" +
		"备注：" + in.FoodOrder.MemberMessage + "\\r\\n" +
		//"</FH>" +
		"\\r\\n" + "\\r\\n"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = content

	return
}

func (s *sBasicsPrinter) GetXYUNFoodOrderContent(ctx context.Context, in *input_basics.GetYILINKFoodOrderContentInp) (res *input_basics.GetPrintOrderContentModel, err error) {

	// 打印内容
	var orderType = "餐厅"

	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"
	printContent = printContent + "<CB>" + orderType + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.New(in.FoodOrder.DepositPayTime).Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>订单编号:" + in.FoodOrder.OrderSn + "<BR>"
	printContent = printContent + "订单金额：" + gvar.New(in.FoodOrder.OrderAmount).String() + "JPY<BR>"

	if in.FoodOrder.OrderType == "CRS" {
		printContent = printContent + "支付定金：" + gvar.New(in.FoodOrder.DepositAmount).String() + "<BR>"
		printContent = printContent + "定金支付：" + gvar.New(in.FoodOrder.DepositPayTime).String() + "<BR>"
		if in.FoodOrder.PayStep == "REMAIN" {
			printContent = printContent +
				"尾款支付：" + gvar.New(in.FoodOrder.PayTime).String() + "<BR>"
		}
	} else {
		printContent = printContent + "支付时间：" + gvar.New(in.FoodOrder.PayTime).String() + "<BR>"
	}

	printContent = printContent + "<BR>"
	printContent = printContent + "餐厅：" + in.FoodOrder.RestaurantDetail.Name + "<BR>"
	printContent = printContent + "套餐：" + in.FoodOrder.GoodsDetail.GoodsName + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "预订人信息：" + "<BR>"
	printContent = printContent + "预订人：" + in.FoodOrder.BookingName + "<BR>"
	printContent = printContent + "电话：" + in.FoodOrder.PhoneArea + in.FoodOrder.BookingMobile + "<BR>"
	printContent = printContent + "人数：" + gvar.New(in.FoodOrder.BookingCount).String() + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "备注：" + in.FoodOrder.MemberMessage + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR><BR>"

	res = &input_basics.GetPrintOrderContentModel{}
	res.Content = printContent

	return
}
