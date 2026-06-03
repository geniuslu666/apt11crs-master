package queue

import (
	"APT/internal/consts"
	"APT/internal/dao"
	_ "APT/internal/logic"
	"APT/internal/model/entity"
	"APT/utility/rabbitmq"
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shopspring/decimal"
)

var (
	ExpLogger = g.Log().Path("logs/MQ/" + consts.RabbitMQQueueNameExp)
)

func Exp() {
	var (
		ctx          = gctx.New()
		MQMsg        <-chan amqp.Delivery
		exchangeName = consts.RabbitMQExchangeName
		QueueName    = consts.RabbitMQQueueNameExp
		MQConnStruct *rabbitmq.MQConnection
		err          error
	)
	g.DB().SetLogger(ExpLogger)
	MQConnStruct = &rabbitmq.MQConnection{
		Conn:         rabbitmq.Conn,
		Chan:         nil,
		ExchangeName: exchangeName,
		QueueName:    QueueName,
		RouteKey:     fmt.Sprintf("%s.%s", exchangeName, QueueName),
	}
	if err = MQConnStruct.Channel(); err != nil {
		err = gerror.New("创建队列失败")
		goto ERR
	}
	if err = MQConnStruct.Exchange("topic", true, false); err != nil {
		err = gerror.New("创建交换机失败")
		goto ERR
	}
	if err = MQConnStruct.Queue(true, false, nil); err != nil {
		err = gerror.New("创建通道失败")
		goto ERR
	}
	if err = MQConnStruct.Bind(); err != nil {
		err = gerror.New("交换机绑定队列失败")
		goto ERR
	}
	ExpLogger.Info(ctx, "Exp Queue START SUCCESSFUL")
	if MQMsg, err = MQConnStruct.Consume(guid.S(), false); err != nil {
		ExpLogger.Error(ctx, err.Error())
		return
	}
	for msg := range MQMsg {
		ctx = gctx.New()
		ExpLogger.Info(ctx, "--[start]----------------------------------------------")
		ExpLogger.Info(ctx, msg.Body)
		if !g.IsEmpty(msg.Body) {
			if string(msg.Body)[:1] == "H" {
				if err = HandleHotelExpMQ(ctx, msg); err != nil {
					ExpLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "S" {
				if err = HandleSpaExpMQ(ctx, msg); err != nil {
					ExpLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "C" {
				if err = HandleCarExpMQ(ctx, msg); err != nil {
					ExpLogger.Error(ctx, err.Error())
				}
			}
			if string(msg.Body)[:1] == "F" {
				if err = HandleFoodExpMQ(ctx, msg); err != nil {
					ExpLogger.Error(ctx, err.Error())
				}
			}
		}

		_ = msg.Ack(false)
	}
ERR:
	ExpLogger.Error(ctx, err)
	panic(err)
}

func HandleHotelExpMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body              = msg.Body
		OrderSn           string
		StayInfo          *entity.PmsAppStay
		pmsAppReservation []*entity.PmsAppReservation
		PmsMember         *entity.PmsMember
		PmsMemberLevel    *entity.PmsMemberLevel
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			ExpLogger.Error(ctx, r)
		}
	}()
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		ExpLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.PmsAppStay.Ctx(ctx).TX(tx).
			Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).
			Scan(&StayInfo); err != nil {
			return
		}
		if g.IsEmpty(StayInfo) {
			err = gerror.New("住宿订单数据为空")
			return
		}
		if StayInfo.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单未支付")
			return
		}
		if err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
			Where(dao.PmsAppReservation.Columns().OrderSn, OrderSn).
			Where(dao.PmsAppReservation.Columns().OrderStatus, "HAVE_PAID").
			Where(dao.PmsAppReservation.Columns().Status, "confirmed").
			Where(dao.PmsAppReservation.Columns().CheckinStatus, "checked_out").
			WhereNull(dao.PmsAppReservation.Columns().ExpTime).
			Scan(&pmsAppReservation); err != nil {
			return
		}
		if g.IsEmpty(pmsAppReservation) {
			err = gerror.New("无可结算经验经验值")
			return
		}
		ExpLogger.Info(ctx, "结算经验")
		for _, Reservation := range pmsAppReservation {
			ExpValue := decimal.NewFromFloat(Reservation.BookingFee).Sub(decimal.NewFromFloat(Reservation.CancellationFee)).Round(0).InexactFloat64()
			// 记录经验处理
			if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).
				Where(dao.PmsAppReservation.Columns().Id, Reservation.Id).
				Update(g.Map{
					dao.PmsAppReservation.Columns().ExpTime:  gtime.Now(),
					dao.PmsAppReservation.Columns().ExpValue: ExpValue,
				}); err != nil {
				err = gerror.New("经验记录失败")
				return
			}
			// 插入经验值记录
			if _, err = dao.PmsExpChange.Ctx(ctx).TX(tx).Data(&entity.PmsExpChange{
				Scene:    "HOTEL",
				Exp:      ExpValue,
				OrderSn:  Reservation.OrderSn,
				Des:      "酒店退房结算经验",
				MemberId: Reservation.MemberId,
			}).InsertAndGetId(); err != nil {
				err = gerror.New("经验记录失败")
				return
			}
			// 增加经验值
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
				Where(dao.PmsMember.Columns().Id, Reservation.MemberId).
				Increment(dao.PmsMember.Columns().Exp, ExpValue); err != nil {
				err = gerror.New("经验增加失败")
				return
			}
		}
		// 查询当前会员积分 是否能够升级
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, StayInfo.MemberId).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			err = gerror.New("会员数据为空")
			return
		}
		// 查询当前会员的经验值处理那个会员等级
		if err = dao.PmsMemberLevel.Ctx(ctx).
			WhereLTE(dao.PmsMemberLevel.Columns().Exp, PmsMember.Exp).
			OrderDesc(dao.PmsMemberLevel.Columns().Exp).
			Scan(&PmsMemberLevel); err != nil {
			return
		}
		if g.IsEmpty(PmsMemberLevel) {
			err = gerror.New("会员等级数据为空")
			return
		}
		if PmsMemberLevel.Id != PmsMember.Level {
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
				Where(dao.PmsMember.Columns().Id, StayInfo.MemberId).
				Update(g.Map{
					dao.PmsMember.Columns().Level: PmsMemberLevel.Id,
				}); err != nil {
				err = gerror.New("会员等级更新失败")
				return
			}
		}
		ExpLogger.Info(ctx, "RUN_SUCCESS")
		return err
	}); err != nil {
		ExpLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
func HandleSpaExpMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body           = msg.Body
		OrderSn        string
		SpaOrder       *entity.SpaOrder
		PmsMember      *entity.PmsMember
		PmsMemberLevel *entity.PmsMemberLevel
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		ExpLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.SpaOrder.Ctx(ctx).TX(tx).
			Where(dao.SpaOrder.Columns().OrderSn, OrderSn).
			Scan(&SpaOrder); err != nil {
			return
		}
		if g.IsEmpty(SpaOrder) {
			err = gerror.New("按摩订单数据为空")
			return
		}
		if SpaOrder.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单未支付")
			return
		}
		ExpLogger.Info(ctx, "结算经验")
		ExpValue := decimal.NewFromFloat(SpaOrder.OrderAmount).Sub(decimal.NewFromFloat(SpaOrder.RefundAmount)).Round(0).InexactFloat64()
		// 记录经验处理
		if _, err = dao.SpaOrder.Ctx(ctx).TX(tx).
			Where(dao.SpaOrder.Columns().Id, SpaOrder.Id).
			Update(g.Map{
				dao.SpaOrder.Columns().ExpTime:  gtime.Now(),
				dao.SpaOrder.Columns().ExpValue: ExpValue,
			}); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 插入经验值记录
		if _, err = dao.PmsExpChange.Ctx(ctx).TX(tx).Data(&entity.PmsExpChange{
			Scene:    "SPA",
			Exp:      ExpValue,
			OrderSn:  SpaOrder.OrderSn,
			Des:      "按摩结算经验",
			MemberId: gvar.New(SpaOrder.MemberId).Int(),
		}).InsertAndGetId(); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 增加经验值
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, SpaOrder.MemberId).
			Increment(dao.PmsMember.Columns().Exp, ExpValue); err != nil {
			err = gerror.New("经验增加失败")
			return
		}
		// 查询当前会员积分 是否能够升级
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, SpaOrder.MemberId).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			err = gerror.New("会员数据为空")
			return
		}
		// 查询当前会员的经验值处理那个会员等级
		if err = dao.PmsMemberLevel.Ctx(ctx).
			WhereLTE(dao.PmsMemberLevel.Columns().Exp, PmsMember.Exp).
			OrderDesc(dao.PmsMemberLevel.Columns().Exp).
			Scan(&PmsMemberLevel); err != nil {
			return
		}
		if g.IsEmpty(PmsMemberLevel) {
			err = gerror.New("会员等级数据为空")
			return
		}
		if PmsMemberLevel.Id != PmsMember.Level {
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
				Where(dao.PmsMember.Columns().Id, SpaOrder.MemberId).
				Update(g.Map{
					dao.PmsMember.Columns().Level: PmsMemberLevel.Id,
				}); err != nil {
				err = gerror.New("会员等级更新失败")
				return
			}
		}
		ExpLogger.Info(ctx, "RUN_SUCCESS")
		return err
	}); err != nil {
		ExpLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
func HandleCarExpMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body           = msg.Body
		OrderSn        string
		CarOrder       *entity.CarOrder
		PmsMember      *entity.PmsMember
		PmsMemberLevel *entity.PmsMemberLevel
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		ExpLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.CarOrder.Ctx(ctx).TX(tx).
			Where(dao.CarOrder.Columns().OrderSn, OrderSn).
			Scan(&CarOrder); err != nil {
			return
		}
		if g.IsEmpty(CarOrder) {
			err = gerror.New("出行订单数据为空")
			return
		}
		if CarOrder.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单未支付")
			return
		}
		ExpLogger.Info(ctx, "结算经验")
		ExpValue := decimal.NewFromFloat(CarOrder.OrderAmount).Sub(decimal.NewFromFloat(CarOrder.RefundAmount)).Round(0).InexactFloat64()
		// 记录经验处理
		if _, err = dao.CarOrder.Ctx(ctx).TX(tx).
			Where(dao.CarOrder.Columns().Id, CarOrder.Id).
			Update(g.Map{
				dao.CarOrder.Columns().ExpTime:  gtime.Now(),
				dao.CarOrder.Columns().ExpValue: ExpValue,
			}); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 插入经验值记录
		if _, err = dao.PmsExpChange.Ctx(ctx).TX(tx).Data(&entity.PmsExpChange{
			Scene:    "CAR",
			Exp:      ExpValue,
			OrderSn:  CarOrder.OrderSn,
			Des:      "出行结算经验",
			MemberId: gvar.New(CarOrder.MemberId).Int(),
		}).InsertAndGetId(); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 增加经验值
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, CarOrder.MemberId).
			Increment(dao.PmsMember.Columns().Exp, ExpValue); err != nil {
			err = gerror.New("经验增加失败")
			return
		}
		// 查询当前会员积分 是否能够升级
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, CarOrder.MemberId).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			err = gerror.New("会员数据为空")
			return
		}
		// 查询当前会员的经验值处理那个会员等级
		if err = dao.PmsMemberLevel.Ctx(ctx).
			WhereLTE(dao.PmsMemberLevel.Columns().Exp, PmsMember.Exp).
			OrderDesc(dao.PmsMemberLevel.Columns().Exp).
			Scan(&PmsMemberLevel); err != nil {
			return
		}
		if g.IsEmpty(PmsMemberLevel) {
			err = gerror.New("会员等级数据为空")
			return
		}
		if PmsMemberLevel.Id != PmsMember.Level {
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
				Where(dao.PmsMember.Columns().Id, CarOrder.MemberId).
				Update(g.Map{
					dao.PmsMember.Columns().Level: PmsMemberLevel.Id,
				}); err != nil {
				err = gerror.New("会员等级更新失败")
				return
			}
		}
		ExpLogger.Info(ctx, "RUN_SUCCESS")
		return err
	}); err != nil {
		ExpLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
func HandleFoodExpMQ(ctx context.Context, msg amqp.Delivery) (err error) {
	var (
		body           = msg.Body
		OrderSn        string
		FoodOrder      *entity.FoodOrder
		PmsMember      *entity.PmsMember
		PmsMemberLevel *entity.PmsMemberLevel
	)
	OrderSn = gvar.New(body).String()
	if g.IsEmpty(OrderSn) {
		err = gerror.New("订单号为空")
		return
	}
	// 开启事务处理
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		ExpLogger.Info(ctx, "检测订单信息")
		// 查询 app 订单信息
		if err = dao.FoodOrder.Ctx(ctx).TX(tx).
			Where(dao.FoodOrder.Columns().OrderSn, OrderSn).
			Scan(&FoodOrder); err != nil {
			return
		}
		if g.IsEmpty(FoodOrder) {
			err = gerror.New("餐饮订单数据为空")
			return
		}
		if FoodOrder.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单未支付")
			return
		}
		ExpLogger.Info(ctx, "结算经验")
		ExpValue := decimal.NewFromFloat(FoodOrder.OrderAmount).Sub(decimal.NewFromFloat(FoodOrder.RefundAmount)).Round(0).InexactFloat64()
		// 记录经验处理
		if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).
			Where(dao.FoodOrder.Columns().Id, FoodOrder.Id).
			Update(g.Map{
				dao.FoodOrder.Columns().ExpTime:  gtime.Now(),
				dao.FoodOrder.Columns().ExpValue: ExpValue,
			}); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 插入经验值记录
		if _, err = dao.PmsExpChange.Ctx(ctx).TX(tx).Data(&entity.PmsExpChange{
			Scene:    "FOOD",
			Exp:      ExpValue,
			OrderSn:  FoodOrder.OrderSn,
			Des:      "餐饮结算经验",
			MemberId: gvar.New(FoodOrder.MemberId).Int(),
		}).InsertAndGetId(); err != nil {
			err = gerror.New("经验记录失败")
			return
		}
		// 增加经验值
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, gvar.New(FoodOrder.MemberId).Int()).
			Increment(dao.PmsMember.Columns().Exp, ExpValue); err != nil {
			err = gerror.New("经验增加失败")
			return
		}
		// 查询当前会员积分 是否能够升级
		if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, gvar.New(FoodOrder.MemberId).Int()).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			err = gerror.New("会员数据为空")
			return
		}
		// 查询当前会员的经验值处理那个会员等级
		if err = dao.PmsMemberLevel.Ctx(ctx).
			WhereLTE(dao.PmsMemberLevel.Columns().Exp, PmsMember.Exp).
			OrderDesc(dao.PmsMemberLevel.Columns().Exp).
			Scan(&PmsMemberLevel); err != nil {
			return
		}
		if g.IsEmpty(PmsMemberLevel) {
			err = gerror.New("会员等级数据为空")
			return
		}
		if PmsMemberLevel.Id != PmsMember.Level {
			if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
				Where(dao.PmsMember.Columns().Id, gvar.New(FoodOrder.MemberId).Int()).
				Update(g.Map{
					dao.PmsMember.Columns().Level: PmsMemberLevel.Id,
				}); err != nil {
				err = gerror.New("会员等级更新失败")
				return
			}
		}
		ExpLogger.Info(ctx, "RUN_SUCCESS")
		return err
	}); err != nil {
		ExpLogger.Info(ctx, "RUN_FAIL")
		return err
	}

	return
}
