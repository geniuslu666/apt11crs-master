package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsFinance struct{}

func NewBasicsFinance() *sBasicsFinance {
	return &sBasicsFinance{}
}

func init() {
	service.RegisterBasicsFinance(NewBasicsFinance())
}

func (s *sBasicsFinance) Stat(ctx context.Context, in *input_basics.FinanceStatInp) (out *input_basics.FinanceStatModel, err error) {
	out = &input_basics.FinanceStatModel{}

	appStayMod := dao.PmsAppStay.Ctx(ctx)
	appStayRefundMod := dao.PmsAppStay.Ctx(ctx)
	staffWithdrawMod := dao.PmsWithdraw.Ctx(ctx).Where(dao.PmsWithdraw.Columns().Type, "CHANNEL").Where(dao.PmsWithdraw.Columns().Transfer, 2)
	channelWithdrawMod := dao.PmsWithdraw.Ctx(ctx).Where(dao.PmsWithdraw.Columns().Type, "STAFF").Where(dao.PmsWithdraw.Columns().Transfer, 2)

	if len(in.SearchTime) == 2 {
		appStayMod = appStayMod.WhereBetween(dao.PmsAppStay.Columns().CreatedAt, in.SearchTime[0], in.SearchTime[1])
		appStayRefundMod = appStayRefundMod.WhereBetween(dao.PmsAppStay.Columns().CreatedAt, in.SearchTime[0], in.SearchTime[1])
		staffWithdrawMod = staffWithdrawMod.WhereBetween(dao.PmsWithdraw.Columns().ApplyAt, in.SearchTime[0], in.SearchTime[1])
		channelWithdrawMod = channelWithdrawMod.WhereBetween(dao.PmsWithdraw.Columns().ApplyAt, in.SearchTime[0], in.SearchTime[1])
	}

	totalIncome, err := appStayMod.Sum(dao.PmsAppStay.Columns().OrderAmount)

	out.ExpendStat.OrderRefund, err = appStayRefundMod.Sum(dao.PmsAppStay.Columns().RefundAmount)

	out.ExpendStat.StaffWithdraw, err = staffWithdrawMod.Sum(dao.PmsWithdraw.Columns().WithdrawAmount)

	out.ExpendStat.ChannelWithdraw, err = channelWithdrawMod.Sum(dao.PmsWithdraw.Columns().WithdrawAmount)

	out.IncomeStat.HomeStayOrderIncome = totalIncome

	out.OverviewStat.TotalExpend = out.ExpendStat.OrderRefund + out.ExpendStat.StaffWithdraw + out.ExpendStat.ChannelWithdraw
	out.OverviewStat.ExpectIncome = totalIncome - out.OverviewStat.TotalExpend
	out.OverviewStat.TotalIncome = totalIncome

	return
}

func (s *sBasicsFinance) List(ctx context.Context, in *input_basics.FinanceListInp) (out *input_basics.FinanceListModel, err error) {
	var (
		PayMod     = dao.PmsTransaction.Ctx(ctx).OmitEmptyWhere()
		RefundMod  = dao.PmsTransactionRefund.Ctx(ctx).OmitEmptyWhere()
		PayList    []*entity.PmsTransaction
		RefundList []*entity.PmsTransactionRefund
	)
	if len(in.TimeRange) > 1 {
		PayMod = PayMod.WhereBetween(dao.PmsTransaction.Columns().PayTime, in.TimeRange[0], in.TimeRange[1])
		RefundMod = RefundMod.WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, in.TimeRange[0], in.TimeRange[1])
	} else {
		if !g.IsEmpty(in.StartTime) && !g.IsEmpty(in.EndTime) {
			PayMod = PayMod.WhereBetween(dao.PmsTransaction.Columns().PayTime, gtime.New(in.StartTime), gtime.New(in.EndTime))
			RefundMod = RefundMod.WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, gtime.New(in.StartTime), gtime.New(in.EndTime))
		}
	}
	out = &input_basics.FinanceListModel{}

	if out.PayAmount, err = PayMod.
		Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
		Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
		Where(dao.PmsTransaction.Columns().PayType, in.PayType).
		Where(dao.PmsTransaction.Columns().PayChannel, in.PayChannel).
		Sum(dao.PmsTransaction.Columns().Amount); err != nil {
		return
	}

	if out.RefundAmount, err = RefundMod.
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		Where(dao.PmsTransactionRefund.Columns().OrderSn, in.OrderSn).
		Where(dao.PmsTransactionRefund.Columns().RefundType, in.PayType).
		Where(dao.PmsTransactionRefund.Columns().RefundChannel, in.PayChannel).
		Sum(dao.PmsTransactionRefund.Columns().RefundAmount); err != nil {
		return
	}

	out.TransactionAmount = out.PayAmount - out.RefundAmount

	if in.PageReq.Pagination {
		PayMod = PayMod.Page(in.Page, in.PerPage)
		RefundMod = RefundMod.Page(in.Page, in.PerPage)
	}
	if in.Type == "pay" || g.IsEmpty(in.Type) {
		if err = PayMod.
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).
			Where(dao.PmsTransaction.Columns().PayChannel, in.PayChannel).
			Where(dao.PmsTransaction.Columns().PayType, in.PayType).
			OrderDesc(dao.PmsTransaction.Columns().Id).
			ScanAndCount(&PayList, &out.PayCount, false); err != nil {
			return
		}
		if !g.IsEmpty(PayList) {
			for _, item := range PayList {
				out.PayListItem.List = append(out.PayListItem.List, &input_basics.FinanceListModelItem{
					Id:         item.Id,
					Scene:      item.Scene,
					OrderSn:    item.OrderSn,
					TypeSn:     item.TransactionSn,
					PayChannel: item.PayChannel,
					PayType:    item.PayType,
					Amount:     item.PayAmount,
					PayTime:    item.PayTime.Format("Y-m-d h:i:s"),
				})
			}
		}
	}
	if in.Type == "refund" || g.IsEmpty(in.Type) {
		if err = RefundMod.
			Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
			Where(dao.PmsTransactionRefund.Columns().OrderSn, in.OrderSn).
			Where(dao.PmsTransactionRefund.Columns().RefundChannel, in.PayChannel).
			Where(dao.PmsTransactionRefund.Columns().RefundType, in.PayType).
			OrderDesc(dao.PmsTransactionRefund.Columns().Id).
			ScanAndCount(&RefundList, &out.RefundCount, false); err != nil {
			return
		}
		if !g.IsEmpty(RefundList) {
			for _, item := range RefundList {
				out.RefundListItem.List = append(out.RefundListItem.List, &input_basics.FinanceListModelItem{
					Id:         item.Id,
					Scene:      item.Scene,
					OrderSn:    item.OrderSn,
					TypeSn:     item.RefundSn,
					PayChannel: item.RefundChannel,
					PayType:    item.RefundType,
					Amount:     item.RefundAmount,
					PayTime:    item.RefundTime.Format("Y-m-d h:i:s"),
				})
			}
		}
	}
	return
}

func (s *sBasicsFinance) Export(ctx context.Context, in *input_basics.FinanceListInp) (err error) {
	in.PageReq.Pagination = false
	list, err := s.List(ctx, in)
	if err != nil {
		return
	}

	var (
		tags     []string
		fileName = "导出支付报表-" + gctx.CtxId(ctx)
		//sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount,input_form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		sheetName     = "支付报表"
		payExports    []input_basics.FinancePayExportModel
		refundExports []input_basics.FinanceRefundExportModel
	)

	payChannelList, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_channel"})
	payTypeList, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_type"})

	if in.Type == "pay" {
		tags, err = convert.GetEntityDescTags(input_basics.FinancePayExportModel{})

		if err != nil {
			return
		}

		payList := list.PayListItem.List
		for k, v := range payList {
			if v.Scene == "HOTEL" {
				payList[k].Scene = "住宿"
			} else if v.Scene == "SPA" {
				payList[k].Scene = "按摩"
			} else if v.Scene == "CAR" {
				payList[k].Scene = "接送机"
			} else if v.Scene == "FOOD" {
				payList[k].Scene = "餐厅"
			}
			for _, v1 := range payChannelList {
				if v.PayChannel == v1.Value {
					payList[k].PayChannel = v1.Label
				}
			}
			for _, v2 := range payTypeList {
				if v.PayType == v2.Value {
					payList[k].PayType = v2.Label
				}
			}
		}
		if err = gconv.Scan(payList, &payExports); err != nil {
			return
		}

		err = excel.ExportByStructs(ctx, tags, payExports, fileName, sheetName)
	} else {
		tags, err = convert.GetEntityDescTags(input_basics.FinanceRefundExportModel{})
		fileName = "导出退款报表-" + gctx.CtxId(ctx)
		sheetName = "退款报表"

		if err != nil {
			return
		}

		refundList := list.RefundListItem.List
		for k, v := range refundList {
			if v.Scene == "HOTEL" {
				refundList[k].Scene = "住宿"
			}
			for _, v1 := range payChannelList {
				if v.PayChannel == v1.Value {
					refundList[k].PayChannel = v1.Label
				}
			}
			for _, v2 := range payTypeList {
				if v.PayType == v2.Value {
					refundList[k].PayType = v2.Label
				}
			}
		}
		if err = gconv.Scan(refundList, &refundExports); err != nil {
			return
		}

		err = excel.ExportByStructs(ctx, tags, refundExports, fileName, sheetName)
	}
	return
}
