package notify

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/hook"
)

func (c *ControllerHook) DoPayCloud(ctx context.Context, req *hook.DoPayCloudReq) (res *hook.DoPayCloudRes, err error) {
	g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Debugf(ctx, "paycloud 支付回调")
	g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Debugf(ctx, gjson.New(req).String())
	if _, err = dao.PmsTransaction.Ctx(ctx).
		Where(dao.PmsTransaction.Columns().TransactionSn, req.TransNo).Data(g.Map{
		dao.PmsTransaction.Columns().PayAmount: req.PaidAmount,
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).Update(); err != nil {
		return
	}
	if err = service.PayService().ThirdPayCompleted(ctx, &req.ThirdPayInp); err != nil {
		g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Error(ctx, err)
		err = nil
	}
	ghttp.RequestFromCtx(ctx).Response.WriteExit("success")
	return
}
func (c *ControllerHook) RefundCloud(ctx context.Context, req *hook.RefundCloudReq) (res *hook.RefundCloudRes, err error) {
	g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Debug(ctx, "paycloud 退款回调")
	g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Debugf(ctx, "%v+f", req)
	var (
		result               sql.Result
		rows                 int64
		PmsTransactionRefund *entity.PmsTransactionRefund
		tx                   gdb.TX
	)
	if req.TransStatus == 2 {
		// 更新退款状态
		if result, err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().RefundSn:     req.MerchantOrderNo,
			dao.PmsTransactionRefund.Columns().RefundStatus: "WAIT",
		}).Data(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
			dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
		}).Update(); err != nil {
			g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Error(ctx, err)
			err = gerror.New("处理退款失败")
			return
		}
		if rows, err = result.RowsAffected(); err != nil {
			g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Error(ctx, err)
			err = gerror.New("处理退款失败")
			return
		}
		if rows != 1 {
			g.Log().Path("logs/HOOK/PAYCLOUD_HOOK").Error(ctx, err)
			err = gerror.New("处理退款失败")
			return
		}

		// 查询订单号
		if err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().RefundSn: req.MerchantOrderNo,
		}).Scan(&PmsTransactionRefund); err != nil {
			return
		}
		tx, err = g.DB().Begin(ctx)
		if err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
		}
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			} else {
				_ = tx.Commit()
			}
		}()
		if err = service.Refund().RefundStatusChange(ctx, tx, PmsTransactionRefund.OrderSn); err != nil {
			return
		}

		ghttp.RequestFromCtx(ctx).Response.WriteExit("success")
	}
	return
}
