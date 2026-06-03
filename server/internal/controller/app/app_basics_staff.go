package app

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/basics"
)

func (c *ControllerBasics) StaffFxCenter(ctx context.Context, req *basics.StaffFxCenterReq) (res *basics.StaffFxCenterRes, err error) {
	var (
		PmsAppStay  []*entity.PmsAppStay
		PaidAppStay []*struct {
			*entity.PmsAppStay
			TransactionDetail []*struct {
				gmeta.Meta `orm:"table:hg_pms_transaction"`
				Id         int     `json:"id"               orm:"id"                 description:"主键"`
				OrderSn    string  `json:"orderSn"          orm:"order_sn"           description:"订单号"`
				PayType    string  `json:"payType"          orm:"pay_type"           description:"支付方式   BAL 余额"`
				Amount     float64 `json:"amount"           orm:"amount"             description:"总金额"`
				PayAmount  float64 `json:"payAmount"        orm:"pay_amount"         description:"支付金额"`
				PayStatus  string  `json:"payStatus"        orm:"pay_status"         description:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
			} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
			TransactionRefundDetail []*struct {
				gmeta.Meta   `orm:"table:hg_pms_transaction_refund"`
				Id           int     `json:"id"            orm:"id"             description:"主键"`
				OrderSn      string  `json:"orderSn"       orm:"order_sn"       description:"订单号"`
				RefundType   string  `json:"refundType"    orm:"refund_type"    description:"'支付方式   BAL 余额'"`
				RefundAmount float64 `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
				RefundStatus string  `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
			} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
		}
		MemberInfo *entity.PmsMember
		Staff      *entity.PmsStaff
	)
	res = new(basics.StaffFxCenterRes)
	if err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsMember.Columns().StaffId:    req.StaffId,
		dao.PmsMember.Columns().RebateMode: "STAFF",
	}).Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberInfo) {
		// 员工信息不存在
		err = gerror.New(gi18n.T(ctx, "employee_info_does_not_exist"))
		return
	}
	if err = dao.PmsStaff.Ctx(ctx).Where(dao.PmsStaff.Columns().Id, req.StaffId).Scan(&Staff); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(Staff) {
		// 员工信息不存在
		err = gerror.New(gi18n.T(ctx, "employee_info_does_not_exist"))
		return
	}
	res.AllBrokerage = Staff.AllBalance
	res.Brokerage = Staff.Balance
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().Referrer, MemberInfo.Id).Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").Page(1, 4).WithAll().Scan(&PaidAppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 计算单数据异常
		err = gerror.New(gi18n.T(ctx, "abnormal_cal_of_single_data"))
		return
	}
	var balPayAmount float64
	var thirdRefundAmount float64
	var couponPayAmount float64
	for k, v := range PaidAppStay {
		if v.RebateStatus == "WAIT" {
			// 需要排除掉积分、优惠券支付的部分
			balPayAmount = 0
			thirdRefundAmount = 0
			couponPayAmount = 0
			if v.TransactionDetail != nil {
				for _, t := range v.TransactionDetail {
					if t.PayType == "BAL" && t.PayStatus == "DONE" {
						balPayAmount += t.PayAmount
					}
					if t.PayType == "COUPON" && t.PayStatus == "DONE" {
						couponPayAmount += t.PayAmount
					}
				}
			}
			if v.TransactionRefundDetail != nil {
				for _, t := range v.TransactionRefundDetail {
					if t.RefundType != "BAL" && t.RefundStatus == "DONE" {
						thirdRefundAmount += t.RefundAmount
					}
				}
			}
			// 使用decimal进行精确计算，避免浮点数精度问题
			netAmount := decimal.NewFromFloat(v.OrderAmount).
				Sub(decimal.NewFromFloat(balPayAmount)).
				Sub(decimal.NewFromFloat(thirdRefundAmount)).
				Sub(decimal.NewFromFloat(couponPayAmount))
			rebateAmount := netAmount.
				Mul(decimal.NewFromFloat(v.RebateRate)).
				Div(decimal.NewFromFloat(100)).
				Round(0) // 保留2位小数
			PaidAppStay[k].RebateAmount += rebateAmount.InexactFloat64()
		}
		res.BrokerageList = append(res.BrokerageList, &struct {
			Id           int64       `json:"id"             orm:"id"              description:"ID"`
			OrderAmount  float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
			Referrer     int         `json:"referrer"       orm:"referrer"        description:"推荐人"`
			RebateRate   float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
			RebateStatus string      `json:"rebateStatus"   orm:"rebate_status"   description:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
			RebateAmount float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
			OrderSn      string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
			RebateTime   *gtime.Time `json:"rebateTime"     orm:"rebate_time"     description:"分佣结算时间"`
			CreatedAt    *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
		}{
			Id:           int64(v.Id),
			OrderAmount:  v.OrderAmount,
			Referrer:     v.Referrer,
			RebateRate:   v.RebateRate,
			RebateStatus: v.RebateStatus,
			RebateAmount: PaidAppStay[k].RebateAmount,
			OrderSn:      v.OrderSn,
			RebateTime:   v.RebateTime,
			CreatedAt:    v.CreatedAt,
		})
	}
	// 结算中
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().Referrer:     MemberInfo.Id,
		dao.PmsAppStay.Columns().RebateStatus: "WAIT",
	}).Scan(&PmsAppStay); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 计算单数据异常
		err = gerror.New(gi18n.T(ctx, "abnormal_cal_of_single_data"))
		return
	}
	for _, v := range PmsAppStay {
		if v.RebateStatus == "WAIT" {
			res.SettlingBrokerage += v.OrderAmount * (v.RebateRate / 100)
		} else if v.RebateStatus == "SUCCESS" {
			//res.SettlingWithdraw += v.RebateAmount
		}
	}

	// 提现中金额
	res.SettlingWithdraw = Staff.ApplyWithdrawBalance
	// 已提现金额
	res.Withdraw = Staff.WithdrawBalance

	/*if res.SettlingWithdraw, err = dao.PmsWithdraw.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsWithdraw.Columns().StaffId:        req.StaffId,
		dao.PmsWithdraw.Columns().WithdrawStatus: "WAIT",
	}).Sum(dao.PmsWithdraw.Columns().WithdrawAmount); err != nil {
		err = gerror.New("计算单数据异常")
		return
	}*/

	return
}
