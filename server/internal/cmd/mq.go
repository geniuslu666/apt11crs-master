package cmd

import (
	"APT/internal/queue"
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	MQ = &gcmd.Command{
		Name:  "mq",
		Brief: "run all rabbitmq queue",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			go queue.OrderStay()
			go queue.PlaceOrder()
			go queue.Availabilities()
			go queue.Exp()
			go queue.Rebate()
			go queue.ThCouponEffect()
			queue.OrderExpire()
			return
		},
	}

	MQOrderStay = &gcmd.Command{
		Name:  "MQOrderStay",
		Usage: "start",
		Brief: "队列实现-AirHost同步住房订单信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.OrderStay()
			return
		},
	}

	MQPlaceOrder = &gcmd.Command{
		Name:  "MQPlaceOrder",
		Usage: "start",
		Brief: "队列实现-住宿支付完成之后下单异步操作",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.PlaceOrder()
			return
		},
	}

	MQAvailabilities = &gcmd.Command{
		Name:  "MQAvailabilities",
		Usage: "start",
		Brief: "队列实现-住房库存和房价都同步操作",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.Availabilities()
			return
		},
	}

	MQExp = &gcmd.Command{
		Name:  "MQExp",
		Usage: "start",
		Brief: "队列实现-会员经验值",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.Exp()
			return
		},
	}

	MQRebate = &gcmd.Command{
		Name:  "MQRebate",
		Usage: "start",
		Brief: "队列实现-用户下单推荐人返佣金",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.Rebate()
			return
		},
	}

	MQOrderExpire = &gcmd.Command{
		Name:  "MQOrderExpire",
		Usage: "start",
		Brief: "队列实现-订单取消",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.OrderExpire()
			return
		},
	}

	MQThCouponEffect = &gcmd.Command{
		Name:  "MQThCouponEffect",
		Usage: "start",
		Brief: "队列实现-礼品券生效",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.ThCouponEffect()
			return
		},
	}

	MQOrderAward = &gcmd.Command{
		Name:  "MQOrderAward",
		Usage: "start",
		Brief: "队列实现-下单奖励",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.OrderAward()
			return
		},
	}

	MQOrderRemind = &gcmd.Command{
		Name:  "MQOrderRemind",
		Usage: "start",
		Brief: "队列实现-订单通知",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.OrderRemind()
			return
		},
	}

	MQOrderExport = &gcmd.Command{
		Name:  "MQOrderExport",
		Usage: "start",
		Brief: "队列实现-订单导出",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.OrderExport()
			return
		},
	}

	MQEmployeeActivityEffect = &gcmd.Command{
		Name:  "MQEmployeeActivityEffect",
		Usage: "start",
		Brief: "队列实现-员工活动生效",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.EmployeeActivityEffect()
			return
		},
	}

	MQSystemMessage = &gcmd.Command{
		Name:  "MQSystemMessage",
		Usage: "start",
		Brief: "队列实现-系统消息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.SystemMessage()
			return
		},
	}
	MQFxChangeOrderPush = &gcmd.Command{
		Name:  "MQFxChangeOrderPush",
		Usage: "start",
		Brief: "队列实现-分销订单消息变更",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.FxChangeOrderPush()
			return
		},
	}
)
