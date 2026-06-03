package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/library/xpyun"
	"APT/internal/library/xpyun/formatter"
	"APT/internal/library/xpyun/model"
	"APT/internal/library/xpyun/util"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsTerminal struct{}

func NewBasicsTerminal() *sBasicsTerminal {
	return &sBasicsTerminal{}
}

func init() {
	service.RegisterBasicsTerminal(NewBasicsTerminal())
}

// Model 终端ORM模型
func (s *sBasicsTerminal) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysTerminal.Ctx(ctx), option...)
}

// List 获取终端列表
func (s *sBasicsTerminal) List(ctx context.Context, in *input_basics.TerminalListInp) (list []*input_basics.TerminalListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_basics.TerminalListModel{})

	if !g.IsEmpty(in.TerminalIds) {
		mod = mod.WhereIn(dao.SysTerminal.Columns().Id, strings.Split(in.TerminalIds, ","))
	} else {
		if in.FromStore == 1 {
			mod = mod.Where(dao.SysTerminal.Columns().Id, -1)
		}
	}

	// 终端名称
	if !g.IsEmpty(in.TerminalName) {
		mod = mod.WhereLike(dao.SysTerminal.Columns().TerminalName, "%"+in.TerminalName+"%")
	}

	// 机器码
	if !g.IsEmpty(in.Sn) {
		mod = mod.WhereLike(dao.SysTerminal.Columns().Sn, "%"+in.Sn+"%")
	}

	// 当前门店ID
	if !g.IsEmpty(in.TerminalStoreId) {
		mod = mod.Where(mod.Builder().
			Where(dao.SysTerminal.Columns().StoreId, in.TerminalStoreId).
			WhereOr(dao.SysTerminal.Columns().StoreId, 0))
		//Where(dao.SysTerminal.Columns().RestaurantId, 0)
	} else {
		if !g.IsEmpty(in.StoreId) {
			if in.StoreId == -1 {
				mod = mod.Where(dao.SysTerminal.Columns().StoreId, 0).Where(dao.SysTerminal.Columns().RestaurantId, 0)
			} else {
				mod = mod.Where(dao.SysTerminal.Columns().StoreId, in.StoreId)
			}
		}
	}

	// 当前餐厅ID
	if !g.IsEmpty(in.TerminalRestaurantId) {
		mod = mod.Where(mod.Builder().
			Where(dao.SysTerminal.Columns().RestaurantId, in.TerminalRestaurantId).
			WhereOr(dao.SysTerminal.Columns().RestaurantId, 0))
		//Where(dao.SysTerminal.Columns().StoreId, 0)
	} else {
		if !g.IsEmpty(in.RestaurantId) {
			if in.RestaurantId == -1 {
				mod = mod.Where(dao.SysTerminal.Columns().RestaurantId, 0).Where(dao.SysTerminal.Columns().StoreId, 0)
			} else {
				mod = mod.Where(dao.SysTerminal.Columns().RestaurantId, in.RestaurantId)
			}
		}
	}

	if !g.IsEmpty(in.TerminalType) {
		mod = mod.Where(dao.SysTerminal.Columns().TerminalType, in.TerminalType)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SysTerminal.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取终端列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取终端列表失败，请稍后重试！")
			return
		}
	}

	// 批量查询打印联数
	if len(list) > 0 {
		// 收集终端ID
		var terminalIds []interface{}
		for _, item := range list {
			terminalIds = append(terminalIds, item.Id)
		}

		// 根据NeedPrintTimesRestaurantId查询餐厅终端打印联数
		if !g.IsEmpty(in.NeedPrintTimesRestaurantId) {
			var restaurantTerminals []struct {
				TerminalId int64 `json:"terminalId"`
				PrintTimes uint  `json:"printTimes"`
			}

			if err = dao.FoodRestaurantTerminal.Ctx(ctx).
				Fields("terminal_id, print_times").
				Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, in.NeedPrintTimesRestaurantId).
				WhereIn(dao.FoodRestaurantTerminal.Columns().TerminalId, terminalIds).
				Scan(&restaurantTerminals); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Warning(ctx, "查询餐厅终端打印联数失败:", err)
			} else {
				// 创建映射表快速匹配
				printTimesMap := make(map[int64]uint)
				for _, rt := range restaurantTerminals {
					printTimesMap[rt.TerminalId] = rt.PrintTimes
				}

				// 更新列表中的打印联数
				for _, item := range list {
					if printTimes, exists := printTimesMap[int64(item.Id)]; exists {
						item.PrintTimes = printTimes
					}
				}
			}
		}

		// 根据NeedPrintTimesStoreId查询门店终端打印联数
		if !g.IsEmpty(in.NeedPrintTimesStoreId) {
			var storeTerminals []struct {
				TerminalId int64 `json:"terminalId"`
				PrintTimes uint  `json:"printTimes"`
			}

			if err = dao.ThStoreTerminal.Ctx(ctx).
				Fields("terminal_id, print_times").
				Where(dao.ThStoreTerminal.Columns().StoreId, in.NeedPrintTimesStoreId).
				WhereIn(dao.ThStoreTerminal.Columns().TerminalId, terminalIds).
				Scan(&storeTerminals); err != nil && !errors.Is(err, sql.ErrNoRows) {
				g.Log().Warning(ctx, "查询门店终端打印联数失败:", err)
			} else {
				// 创建映射表快速匹配
				printTimesMap := make(map[int64]uint)
				for _, st := range storeTerminals {
					printTimesMap[st.TerminalId] = st.PrintTimes
				}

				// 更新列表中的打印联数
				for _, item := range list {
					if printTimes, exists := printTimesMap[int64(item.Id)]; exists {
						item.PrintTimes = printTimes
					}
				}
			}
		}
	}

	return
}

// Edit 修改/新增终端
func (s *sBasicsTerminal) Edit(ctx context.Context, in *input_basics.TerminalEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.TerminalUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改终端失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.TerminalInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增终端失败，请稍后重试！")
		}

		return
	})
}

// Delete 删除终端
func (s *sBasicsTerminal) Delete(ctx context.Context, in *input_basics.TerminalDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除终端失败，请稍后重试！")
		return
	}
	return
}

// View 获取终端指定信息
func (s *sBasicsTerminal) View(ctx context.Context, in *input_basics.TerminalViewInp) (res *input_basics.TerminalViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取终端信息，请稍后重试！")
		return
	}

	return
}

// BrandList 获取品牌型号列表
func (s *sBasicsTerminal) BrandList(ctx context.Context, in *input_basics.BrandListInp) (list []*input_basics.BrandListModel, totalCount int, err error) {
	mod := dao.SysTerminalModel.Ctx(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_basics.BrandListModel{})

	// 终端名称
	if !g.IsEmpty(in.BrandModel) {
		mod = mod.WhereLike(dao.SysTerminalModel.Columns().BrandModel, "%"+in.BrandModel+"%")
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.SysTerminalModel.Columns().Id)
	//mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取品牌型号列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取品牌型号列表失败，请稍后重试！")
			return
		}
	}
	return
}

// BrandEdit 修改/新增品牌型号
func (s *sBasicsTerminal) BrandEdit(ctx context.Context, in *input_basics.BrandEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = dao.SysTerminalModel.Ctx(ctx).
				Fields(input_basics.BrandUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改品牌型号失败，请稍后重试！")
			}
			return
		}
		err = gerror.New("修改失败，请稍后重试！")
		return
	})
}

// BrandView 获取终端指定信息
func (s *sBasicsTerminal) BrandView(ctx context.Context, in *input_basics.BrandViewInp) (res *input_basics.BrandViewModel, err error) {
	if err = dao.SysTerminalModel.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取品牌型号失败，请稍后重试！")
		return
	}

	return
}

// PrinterTest 打印测试
func (s *sBasicsTerminal) PrinterTest(ctx context.Context, in *input_basics.PrinterTestInp) (err error) {

	// 终端信息
	var terminalModel *input_basics.TerminalViewModel
	if err = dao.SysTerminal.Ctx(ctx).
		WhereIn(dao.SysTerminal.Columns().Id, in.TerminalId).
		WithAll().
		Scan(&terminalModel); err != nil {
		err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
		return
	}

	printContent := ""

	printContent = printContent + "<C>" + "<B>芯烨云小票</B>" + "<BR></C>"
	printContent = printContent + "<BR>"

	printContent = printContent + "菜名" + util.StrRepeat(" ", 16) + "数量" + util.StrRepeat(" ", 2) + "单价" + util.StrRepeat(" ", 2) + "<BR>"
	printContent = printContent + util.StrRepeat("-", 32) + "<BR>"
	printContent = printContent + formatter.FormatPrintOrderItem("可乐鸡翅", 2, 9.99)
	printContent = printContent + formatter.FormatPrintOrderItem("水煮鱼特辣", 1, 108.0)
	printContent = printContent + formatter.FormatPrintOrderItem("豪华版超级无敌龙虾炒饭", 1, 99.9)
	printContent = printContent + formatter.FormatPrintOrderItem("炭烤鳕鱼", 5, 19.99)
	printContent = printContent + util.StrRepeat("-", 32) + "<BR>"
	printContent = printContent + "<R>" + "合计：" + "327.83" + "元" + "<BR></R>"

	printContent = printContent + "<BR>"
	printContent = printContent + "<L>" + "客户地址：" + "珠海市香洲区xx路xx号" + "<BR>" + "客户电话：" + "1363*****88" + "<BR>" + "下单时间：" + "2020-9-9 15:07:57" + "<BR>" + "备注：" + "少放辣 不吃香菜" + "<BR>"

	printContent = printContent + "<C>" + "<QRCODE s=6 e=L l=center>https://www.xpyun.net</QRCODE>" + "</C>"

	request := model.PrintRequest{}
	request.User = terminalModel.BrandInfo.ClientId
	request.UserKey = terminalModel.BrandInfo.ClientSecret

	//*必填*：打印机编号
	request.Sn = terminalModel.Sn

	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = printContent

	//打印份数，默认为1
	request.Copies = 1

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
	request.Voice = 2
	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1
	//支付方式：
	//取值范围41~55：
	//支付宝 41、微信 42、云支付 43、银联刷卡 44、银联支付 45、会员卡消费 46、会员卡充值 47、翼支付 48、成功收款 49、嘉联支付 50、壹钱包 51、京东支付 52、快钱支付 53、威支付 54、享钱支付 55
	//仅用于支持金额播报的芯烨云打印机。
	request.PayType = 41

	//支付与否：
	//取值范围59~61：
	//退款 59 到账 60 消费 61。
	//仅用于支持金额播报的芯烨云打印机。
	request.PayMode = 60

	//支付金额：
	//最多允许保留2位小数。
	//仅用于支持金额播报的芯烨云打印机。
	request.Money = 20.15

	// 支持来单播放tts语音文本，目前不能播英语，若英语单词，将会按照单字母方式逐个播报英文字母，需要设备支持
	// request.Tts = "芯烨云来单了，请及时处理"

	result := xpyun.XpYunPrint(ctx, &request)
	//序列化
	reslutJson, _ := json.Marshal(result.Content)
	var msg = fmt.Sprintf("response result: %+v", string(reslutJson))

	fmt.Println("--------QAQ-----打印结果---------")
	fmt.Println(msg)

	if result.Content.Code != 0 {
		err = errors.New("打印失败，请稍后重试！")
		return
	}
	return
}

// Printer 打印方法
func (s *sBasicsTerminal) Printer(ctx context.Context, in *input_basics.PrinterInp) (err error) {

	Logger := g.Log().Path("logs/Verify")

	// 终端信息
	var terminalModel *input_basics.TerminalViewModel
	if err = dao.SysTerminal.Ctx(ctx).
		WhereIn(dao.SysTerminal.Columns().Sn, in.Sn).
		WithAll().
		Scan(&terminalModel); err != nil {
		err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
		return
	}

	request := model.PrintRequest{}
	request.User = terminalModel.BrandInfo.ClientId
	request.UserKey = terminalModel.BrandInfo.ClientSecret

	//*必填*：打印机编号
	request.Sn = terminalModel.Sn

	request.GenerateSign()

	//*必填*：打印内容,不能超过12K
	request.Content = in.PrintContent

	// 打印份数，默认为1（此处手动循环打印）
	request.Copies = 1

	//声音播放模式，0 为取消订单模式，1 为静音模式，2 为来单播放模式，3为有用户申请退单了。默认为 2 来单播放模式
	request.Voice = 2
	//打印模式：
	//值为 0 或不指定则会检查打印机是否在线，如果不在线 则不生成打印订单，直接返回设备不在线状态码；如果在线则生成打印订单，并返回打印订单号。
	//值为 1不检查打印机是否在线，直接生成打印订单，并返回打印订单号。如果打印机不在线，订单将缓存在打印队列中，打印机正常在线时会自动打印。
	request.Mode = 1

	// 支持来单播放tts语音文本，目前不能播英语，若英语单词，将会按照单字母方式逐个播报英文字母，需要设备支持
	request.Tts = "核销成功"

	printTimes := in.PrintTimes
	if printTimes <= 0 {
		printTimes = 1
	}

	for i := 1; i <= printTimes; i++ {
		Logger.Info(ctx, "--------开始打印----------"+fmt.Sprintf("%d回目の印刷", i))
		request.Content = fmt.Sprintf("%s<C>%d回目の印刷</C><BR><BR>", in.PrintContent, i)
		// 打印
		result := xpyun.XpYunPrint(ctx, &request)

		// 序列化
		reslutJson, _ := json.Marshal(result.Content)
		var msg = fmt.Sprintf("response result: %+v", string(reslutJson))

		Logger.Info(ctx, "--------打印结果----------")
		Logger.Info(ctx, msg)

		/*if result.Content.Code != 0 {
			err = errors.New("打印失败，请稍后重试！")
			return
		}*/
	}

	return
}
