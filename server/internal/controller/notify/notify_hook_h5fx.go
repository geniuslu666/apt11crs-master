package notify

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/h5FxPay"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"net/http"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/notify/hook"
)

func (c *ControllerHook) PaySuccessNotify(ctx context.Context, req *hook.PaySuccessNotifyReq) (res *hook.PaySuccessNotifyRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
	)
	g.Log().Path("logs/HOOK/H5FX_HOOK").Debug(ctx, "h5fx 支付回调")
	g.Log().Path("logs/HOOK/H5FX_HOOK").Debugf(ctx, "%v+f", req)
	defer func() {
		if err != nil {
			r.Response.WriteStatusExit(http.StatusBadRequest, err.Error())
		}
	}()

	// 查询transaction表
	var PmsTransaction *entity.PmsTransaction
	if err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().TransactionSn: req.TradeNo,
	}).Scan(&PmsTransaction); err != nil {
		return
	}
	if g.IsEmpty(PmsTransaction) {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, "交易记录不存在")
		err = gerror.New("交易记录不存在")
		return
	}

	if _, err = dao.PmsTransaction.Ctx(ctx).
		Where(dao.PmsTransaction.Columns().TransactionSn, req.TradeNo).Data(g.Map{
		dao.PmsTransaction.Columns().PayAmount: req.TradeAmount,
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).Update(); err != nil {
		return
	}

	if err = service.PayService().ThirdPayCompleted(ctx, &input_pay.ThirdPayInp{
		//TransNo:   req.Resource.SupplementaryData.RelatedIds.OrderID,
		TransNo: req.TradeNo,
	}); err != nil {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
		err = nil
	}
	if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
		DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
			OrderNo:      PmsTransaction.OrderSn,
			ChangeStatus: "SUCCESS",
		}).MustToJson(),
		Header: nil,
	}); err != nil {
		g.Log().Error(ctx, "发送分销订单消息变更失败", err)
	}
	r.Response.WriteExit("SUCCESS")
	return
}
func (c *ControllerHook) RefundSuccessNotify(ctx context.Context, req *hook.RefundSuccessNotifyReq) (res *hook.RefundSuccessNotifyRes, err error) {

	g.Log().Path("logs/HOOK/H5FX_HOOK").Debug(ctx, "h5fx 退款回调")
	g.Log().Path("logs/HOOK/H5FX_HOOK").Debugf(ctx, "%v+f", req)
	var (
		result               sql.Result
		rows                 int64
		PmsTransactionRefund *entity.PmsTransactionRefund
		tx                   gdb.TX
	)

	// 更新退款状态
	if result, err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().RefundSn:     req.RefundOrderNo,
		dao.PmsTransactionRefund.Columns().RefundStatus: "WAIT",
	}).Data(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
		dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
	}).Update(); err != nil {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
		err = gerror.New("处理退款失败")
		return
	}
	if rows, err = result.RowsAffected(); err != nil {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
		err = gerror.New("处理退款失败")
		return
	}
	if rows != 1 {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
		err = gerror.New("处理退款失败")
		return
	}

	// 查询订单号
	if err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().RefundSn: req.RefundOrderNo,
	}).Scan(&PmsTransactionRefund); err != nil {
		return
	}
	if g.IsEmpty(PmsTransactionRefund) {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, "退款记录不存在")
		err = gerror.New("退款记录不存在")
		return
	}
	tx, err = g.DB().Begin(ctx)
	if err != nil {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if err = service.Refund().RefundStatusChange(ctx, tx, PmsTransactionRefund.OrderSn); err != nil {
		g.Log().Path("logs/HOOK/H5FX_HOOK").Error(ctx, err)
		return
	}

	ghttp.RequestFromCtx(ctx).Response.WriteExit("success")
	return

}
