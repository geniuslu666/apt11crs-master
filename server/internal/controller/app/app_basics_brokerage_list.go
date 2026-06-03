package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/shopspring/decimal"

	"APT/api/app/basics"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/utility/uuid"
)

func (c *ControllerBasics) BrokerageList(ctx context.Context, req *basics.BrokerageListReq) (res *basics.BrokerageListRes, err error) {
	var (
		MemberInfo  *entity.PmsMember
		PaidAppStay []struct {
			Id                int64       `json:"id"             orm:"id"              description:"ID"`
			OrderAmount       float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
			Referrer          int         `json:"referrer"       orm:"referrer"        description:"推荐人"`
			RebateRate        float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
			RebateStatus      string      `json:"rebateStatus"   orm:"rebate_status"   description:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
			RebateAmount      float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
			OrderSn           string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
			RebateTime        *gtime.Time `json:"rebateTime"     orm:"rebate_time"     description:"分佣结算时间"`
			CreatedAt         *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
			TransactionDetail []struct {
				gmeta.Meta `orm:"table:hg_pms_transaction"`
				OrderSn    string  `json:"orderSn"          orm:"order_sn"           description:"订单号"`
				PayAmount  float64 `json:"payAmount"  orm:"pay_amount"  description:"支付金额"`
				PayType    string  `json:"payType"    orm:"pay_type"    description:"支付类型"`
				PayStatus  string  `json:"payStatus"  orm:"pay_status"  description:"支付状态"`
			} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
			TransactionRefundDetail []struct {
				gmeta.Meta   `orm:"table:hg_pms_transaction_refund"`
				OrderSn      string  `json:"orderSn"       orm:"order_sn"       description:"订单号"`
				RefundType   string  `json:"refundType"    orm:"refund_type"    description:"退款类型"`
				RefundAmount float64 `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
				RefundStatus string  `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
			} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
		}
	)
	if err = dao.PmsMember.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsMember.Columns().StaffId:   req.StaffId,
		dao.PmsMember.Columns().ChannelId: req.ChannelId,
	}).OmitEmptyWhere().Scan(&MemberInfo); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		// 会员信息不存在
		err = gerror.New(gi18n.T(ctx, "member_info_does_not_exist"))
		return
	}
	res = new(basics.BrokerageListRes)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().Referrer:     MemberInfo.Id,
		dao.PmsAppStay.Columns().RebateStatus: req.RebateStatus,
		dao.PmsAppStay.Columns().OrderStatus:  "HAVE_PAID",
	}).OmitEmptyWhere().Page(req.PageNum, req.PageSize).WithAll().OrderDesc(dao.PmsAppStay.Columns().Id).ScanAndCount(&PaidAppStay, &res.Count, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		// 列表信息异常
		err = gerror.New(gi18n.T(ctx, "list_info_exception"))
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
				Round(0). // 四舍五入到整数
				InexactFloat64()
			PaidAppStay[k].RebateAmount += rebateAmount
		}

		res.List = append(res.List, &struct {
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
			Id:           v.Id,
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
	return
}
func (c *ControllerBasics) BrokerageDetail(ctx context.Context, req *basics.BrokerageDetailReq) (res *basics.BrokerageDetailRes, err error) {
	var (
		orderDetail struct {
			Id                int64       `json:"id"             orm:"id"              description:"ID"`
			OrderAmount       float64     `json:"orderAmount"    orm:"order_amount"    description:"订单金额"`
			Referrer          int         `json:"referrer"       orm:"referrer"        description:"推荐人"`
			RebateRate        float64     `json:"rebateRate"     orm:"rebate_rate"     description:"分佣比例"`
			RebateStatus      string      `json:"rebateStatus"   orm:"rebate_status"   description:"结算状态 WAIT:待结算 SUCCESS 结算成功  FAIL 结算失败"`
			RebateAmount      float64     `json:"rebateAmount"   orm:"rebate_amount"   description:"分佣结算金额"`
			OrderSn           string      `json:"orderSn"        orm:"order_sn"        description:"订单号"`
			OrderStatus       string      `json:"orderStatus"    orm:"order_status"    description:"订单状态"`
			RebateTime        *gtime.Time `json:"rebateTime"     orm:"rebate_time"     description:"分佣结算时间"`
			CreatedAt         *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
			TransactionDetail []struct {
				gmeta.Meta `orm:"table:hg_pms_transaction"`
				OrderSn    string  `json:"orderSn"          orm:"order_sn"           description:"订单号"`
				PayAmount  float64 `json:"payAmount"  orm:"pay_amount"  description:"支付金额"`
				PayType    string  `json:"payType"    orm:"pay_type"    description:"支付类型"`
				PayStatus  string  `json:"payStatus"  orm:"pay_status"  description:"支付状态"`
			} `json:"transactionDetail" orm:"with:order_sn=order_sn" dc:"支付明细"`
			TransactionRefundDetail []struct {
				gmeta.Meta   `orm:"table:hg_pms_transaction_refund"`
				OrderSn      string  `json:"orderSn"       orm:"order_sn"       description:"订单号"`
				RefundType   string  `json:"refundType"    orm:"refund_type"    description:"退款类型"`
				RefundAmount float64 `json:"refundAmount"  orm:"refund_amount"  description:"退款金额"`
				RefundStatus string  `json:"refundStatus"  orm:"refund_status"  description:"退款状态"`
			} `json:"transactionRefundDetail" orm:"with:order_sn=order_sn" dc:"退款明细"`
		}
	)

	// 查询订单详情，包含交易明细
	if err = dao.PmsAppStay.Ctx(ctx).Where("id", req.ID).WithAll().Scan(&orderDetail); err != nil {
		return
	}

	// 判断订单orderStatus是否为HAVE_PAID
	if orderDetail.OrderStatus != "HAVE_PAID" {
		err = gerror.New("订单状态不正确，无法查看佣金详情")
		return
	}

	res = new(basics.BrokerageDetailRes)
	res.Id = orderDetail.Id
	res.OrderAmount = orderDetail.OrderAmount
	res.Referrer = orderDetail.Referrer
	res.RebateRate = orderDetail.RebateRate
	res.RebateStatus = orderDetail.RebateStatus
	res.OrderSn = orderDetail.OrderSn
	res.RebateTime = orderDetail.RebateTime
	res.CreatedAt = orderDetail.CreatedAt

	// 如果订单的RebateStatus是WAIT，则需要手动计算RebateAmount
	if orderDetail.RebateStatus == "WAIT" {
		var balPayAmount float64
		var thirdRefundAmount float64
		var couponPayAmount float64

		// 计算积分支付金额
		if orderDetail.TransactionDetail != nil {
			for _, t := range orderDetail.TransactionDetail {
				if t.PayType == "BAL" && t.PayStatus == "DONE" {
					balPayAmount += t.PayAmount
				}
				if t.PayType == "COUPON" && t.PayStatus == "DONE" {
					couponPayAmount += t.PayAmount
				}
			}
		}

		// 计算第三方退款金额（排除积分退款）
		if orderDetail.TransactionRefundDetail != nil {
			for _, t := range orderDetail.TransactionRefundDetail {
				if t.RefundType != "BAL" && t.RefundStatus == "DONE" {
					thirdRefundAmount += t.RefundAmount
				}
			}
		}

		// 使用decimal进行精确计算，避免浮点数精度问题
		netAmount := decimal.NewFromFloat(orderDetail.OrderAmount).
			Sub(decimal.NewFromFloat(balPayAmount)).
			Sub(decimal.NewFromFloat(thirdRefundAmount)).
			Sub(decimal.NewFromFloat(couponPayAmount))
		rebateAmount := netAmount.
			Mul(decimal.NewFromFloat(orderDetail.RebateRate)).
			Div(decimal.NewFromFloat(100)).
			Round(0). // 四舍五入到整数
			InexactFloat64()

		res.RebateAmount = rebateAmount
	} else {
		// 如果已经结算，使用数据库中的金额
		res.RebateAmount = orderDetail.RebateAmount
	}

	return
}
func (c *ControllerBasics) BrokerageWithdrawBase(ctx context.Context, req *basics.BrokerageWithdrawBaseReq) (res *basics.BrokerageWithdrawBaseRes, err error) {
	var (
		Channel    *entity.PmsChannel
		Staff      *entity.PmsStaff
		MemberInfo *model.MemberIdentity
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	res = new(basics.BrokerageWithdrawBaseRes)
	if req.ChannelId > 0 && MemberInfo.RebateMode == "CHANNEL" {
		if err = dao.PmsChannel.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsChannel.Columns().Id: req.ChannelId,
		}).OmitEmptyWhere().Scan(&Channel); err != nil {
			return
		}
		res.MinWithdrawalAmount = Channel.MinWithdrawalAmount
		res.Balance = Channel.Balance
		res.ServiceCharge = Channel.ServiceCharge
		res.AfterDay = Channel.AfterDay
	} else if req.StaffId > 0 && MemberInfo.RebateMode == "STAFF" {
		if err = dao.PmsStaff.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsStaff.Columns().Id: req.StaffId,
		}).OmitEmptyWhere().Scan(&Staff); err != nil {
			return
		}
		res.MinWithdrawalAmount = Staff.MinWithdrawalAmount
		res.Balance = Staff.Balance
		res.ServiceCharge = Staff.ServiceCharge
		res.AfterDay = Staff.AfterDay
	} else {
		// 参数错误
		err = gerror.New(gi18n.T(ctx, "parameter_error"))
	}
	return
}
func (c *ControllerBasics) BrokerageWithdraw(ctx context.Context, req *basics.BrokerageWithdrawReq) (res *basics.BrokerageWithdrawRes, err error) {
	var (
		Staff       *entity.PmsStaff
		Channel     *entity.PmsChannel
		ChannelType string
		insertId    int64
		rate        float64
		MemberInfo  *model.MemberIdentity
	)
	res = new(basics.BrokerageWithdrawRes)
	MemberInfo = contexts.GetMemberUser(ctx)
	if req.ChannelId > 0 && MemberInfo.RebateMode == "CHANNEL" {
		ChannelType = "CHANNEL"
		if err = dao.PmsChannel.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsChannel.Columns().Id: req.ChannelId,
		}).OmitEmptyWhere().Scan(&Channel); err != nil {
			return
		}

		if Channel.Balance < req.Amount {
			// 申请提现金额不足
			err = gerror.New(gi18n.T(ctx, "insufficient_apply_withdraw_amount"))
			return
		}

		res.AfterDay = Channel.AfterDay
		rate = Channel.ServiceCharge
	} else if req.StaffId > 0 && MemberInfo.RebateMode == "STAFF" {
		ChannelType = "STAFF"
		if err = dao.PmsStaff.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsStaff.Columns().Id: req.StaffId,
		}).OmitEmptyWhere().Scan(&Staff); err != nil {
			return
		}

		if Staff.Balance < req.Amount {
			// 申请提现金额不足
			err = gerror.New(gi18n.T(ctx, "insufficient_apply_withdraw_amount"))
			return
		}

		res.AfterDay = Staff.AfterDay
		rate = Staff.ServiceCharge
	} else {
		// 参数错误
		err = gerror.New(gi18n.T(ctx, "parameter_error"))
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 减少渠道、员工余额，增加提现中金额
		if ChannelType == "CHANNEL" {
			if _, err = dao.PmsChannel.Ctx(ctx).TX(tx).OmitEmptyData().Where("id", req.ChannelId).Data(g.MapStrAny{
				dao.PmsChannel.Columns().Balance:              gdb.Raw(fmt.Sprintf("balance-%f", req.Amount)),
				dao.PmsChannel.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance+%f", req.Amount)),
			}).Update(); err != nil {
				return
			}
		} else if ChannelType == "STAFF" {
			if _, err = dao.PmsStaff.Ctx(ctx).TX(tx).OmitEmptyData().Where("id", req.StaffId).Data(g.MapStrAny{
				dao.PmsStaff.Columns().Balance:              gdb.Raw(fmt.Sprintf("balance-%f", req.Amount)),
				dao.PmsStaff.Columns().ApplyWithdrawBalance: gdb.Raw(fmt.Sprintf("apply_withdraw_balance+%f", req.Amount)),
			}).Update(); err != nil {
				return
			}
		}

		// 写入提现记录
		if insertId, err = dao.PmsWithdraw.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsWithdraw{
			Type:           ChannelType,
			StaffId:        req.StaffId,
			ChannelId:      req.ChannelId,
			WithdrawAmount: req.Amount,
			WithdrawSn:     uuid.CreateOrderCode("W"),
			ServiceCharge:  rate,
			ArrivalAmount:  req.Amount - (req.Amount * (rate / 100)),
		}).InsertAndGetId(); err != nil {
			return
		}
		if g.IsEmpty(insertId) {
			// 提现申请失败
			err = gerror.New(gi18n.T(ctx, "withdraw_apply_failed"))
			return
		}
		res.Amount = req.Amount
		res.Balance = req.Amount - (req.Amount * (rate / 100))
		res.ServiceCharge = rate
		res.CreateTime = gtime.Now()
		return
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) BrokerageWithdrawList(ctx context.Context, req *basics.BrokerageWithdrawListReq) (res *basics.BrokerageWithdrawListRes, err error) {
	res = new(basics.BrokerageWithdrawListRes)

	mod := dao.PmsWithdraw.Ctx(ctx).OrderDesc("id")

	if !g.IsEmpty(req.StaffId) {
		mod = mod.Where(dao.PmsWithdraw.Columns().StaffId, req.StaffId)
	}
	if !g.IsEmpty(req.ChannelId) {
		mod = mod.Where(dao.PmsWithdraw.Columns().ChannelId, req.ChannelId)
	}
	if !g.IsEmpty(req.WithdrawStatus) {
		mod = mod.Where(dao.PmsWithdraw.Columns().WithdrawStatus, req.WithdrawStatus)
	}
	if !g.IsEmpty(req.PageNum) {
		mod = mod.Page(req.PageNum, req.PageSize)
	}
	if err = mod.OmitEmptyWhere().ScanAndCount(&res.List, &res.Count, false); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) BrokerageWithdrawDetail(ctx context.Context, req *basics.BrokerageWithdrawDetailReq) (res *basics.BrokerageWithdrawDetailRes, err error) {
	res = new(basics.BrokerageWithdrawDetailRes)
	if err = dao.PmsWithdraw.Ctx(ctx).Where("id", req.ID).Scan(&res); err != nil {
		return
	}

	switch res.WithdrawStatus {
	case "WAIT":
		res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
			Des:    "提交申请待审核",
			Time:   res.CreatedAt,
			Remark: "",
		})
	case "SUCCESS":
		res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
			Des:    "提交申请待审核",
			Time:   res.CreatedAt,
			Remark: "",
		})
		res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
			Des:    "提现申请已通过",
			Time:   res.ApplyAt,
			Remark: res.ApplyRemark,
		})
		if res.Transfer == 2 {
			res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
				Des:    "提现打款已处理",
				Time:   res.UpdatedAt,
				Remark: "",
			})
		}
	case "FAIL":
		res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
			Des:    "提交申请待审核",
			Time:   res.CreatedAt,
			Remark: "",
		})
		res.Schedule = append(res.Schedule, &basics.BrokerageWithdrawSchedule{
			Des:    "提交失败申请被驳回",
			Time:   res.ApplyAt,
			Remark: res.ApplyRemark,
		})
	}
	return
}
