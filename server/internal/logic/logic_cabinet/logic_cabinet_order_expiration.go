package logic_cabinet

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/h5FxPay"
	"APT/internal/model/entity"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderExpiration 订单过期
func (s *sCabinetService) OrderExpiration(ctx context.Context, OrderSn string) (err error) {
	var (
		tx                   gdb.TX
		CabinetOrder         entity.CabinetOrder
		PmsCouponTransaction *entity.PmsTransaction
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.CabinetOrder.Columns().OrderSn:     OrderSn,
		dao.CabinetOrder.Columns().OrderStatus: "WAIT_PAY",
	}).Scan(&CabinetOrder); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(CabinetOrder) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	// 更新当前订单状态-取消
	if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.CabinetOrder.Columns().OrderSn:     OrderSn,
		dao.CabinetOrder.Columns().OrderStatus: "WAIT_PAY",
	}).Data(g.MapStrAny{
		dao.CabinetOrder.Columns().PayStatus:   "CANCEL",
		dao.CabinetOrder.Columns().OrderStatus: "CANCEL",
		dao.CabinetOrder.Columns().CancelTime:  gtime.Now(),
	}).Update(); err != nil {
		return
	}

	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
		Where(dao.PmsTransaction.Columns().PayType, "COUPON").
		Scan(&PmsCouponTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(PmsCouponTransaction) {
		// 释放优惠券
		if _, err = dao.PmsCoupon.Ctx(ctx).TX(tx).
			Where(dao.PmsCoupon.Columns().Id, PmsCouponTransaction.CouponId).
			Update(g.MapStrAny{
				dao.PmsCoupon.Columns().State:   1,
				dao.PmsCoupon.Columns().UseTime: nil,
			}); err != nil {
			return
		}
		// 修改支付订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
			Where(dao.PmsTransaction.Columns().TransactionSn, PmsCouponTransaction.TransactionSn).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Update(entity.PmsTransaction{
				PayAmount: 0,
				PayStatus: "CANCEL",
				PayTime:   gtime.Now(),
			}); err != nil {
			return
		}
	}

	// 更新付费信息状态-取消
	if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   OrderSn,
		dao.PmsTransaction.Columns().PayStatus: "WAIT",
	}).Data(g.MapStrAny{
		dao.PmsTransaction.Columns().PayStatus: "CANCEL",
	}).Update(); err != nil {
		return
	}

	// 订单取消
	if _, err = dao.CabinetOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
		OrderId:     int(CabinetOrder.Id),
		OrderStatus: "CANCEL",
		ActionWay:   "CANCEL",
		Remark:      "订单未支付取消",
		OperateType: "SYSTEM",
	}); err != nil {
		return
	}

	if CabinetOrder.IsFx == "Y" {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
			DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
				OrderNo:      CabinetOrder.OrderSn,
				ChangeStatus: "CANCEL",
			}).MustToJson(),
			Header: nil,
		}); err != nil {
			g.Log().Error(ctx, "发送分销订单消息变更失败", err)
			err = nil
		}
	}

	return
}
