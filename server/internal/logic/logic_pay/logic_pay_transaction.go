package logic_pay

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func (s *sPayService) List(ctx context.Context, in *input_pay.PmsTransactionListInp) (list []*input_pay.PmsTransactionListModel, totalCount int, err error) {
	mod := dao.PmsTransaction.Ctx(ctx)

	mod = mod.Fields(input_pay.PmsTransactionListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsTransaction.Columns().Id, in.Id)
	}

	if in.OrderSn != "" {
		mod = mod.WhereLike(dao.PmsTransaction.Columns().OrderSn, in.OrderSn)
	}

	if in.TransactionSn != "" {
		mod = mod.WhereLike(dao.PmsTransaction.Columns().TransactionSn, in.TransactionSn)
	}

	if !g.IsEmpty(in.PayChannel) {
		mod = mod.Where(dao.PmsTransaction.Columns().PayChannel, in.PayChannel)
	}

	if !g.IsEmpty(in.PayType) {
		mod = mod.Where(dao.PmsTransaction.Columns().PayType, in.PayType)
	}

	if !g.IsEmpty(in.PayStatus) {
		mod = mod.Where(dao.PmsTransaction.Columns().PayStatus, in.PayStatus)
	}

	if len(in.PayTime) == 2 {
		mod = mod.WhereBetween(dao.PmsTransaction.Columns().PayTime, in.PayTime[0], in.PayTime[1])
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsTransaction.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.PageReq.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsTransaction.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取支付流水表列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sPayService) Export(ctx context.Context, in *input_pay.PmsTransactionListExportInp) (err error) {
	var (
		list                 []*input_pay.PmsTransactionListModel
		payChannelList       input_basics.DataSelectModel
		payTypeList          input_basics.DataSelectModel
		payStatusList        input_basics.DataSelectModel
		refundStatusList     input_basics.DataSelectModel
		tags                 []string
		PmsTransactionListIn *input_pay.PmsTransactionListInp
	)
	// 无分页
	in.PageReq.Pagination = false
	PmsTransactionListIn = &input_pay.PmsTransactionListInp{
		PageReq:       in.PageReq,
		Id:            in.Id,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		PayChannel:    in.PayChannel,
		PayType:       in.PayType,
		PayStatus:     in.PayStatus,
	}

	if !g.IsEmpty(in.CreatedAt) {
		for _, timeItem := range strings.Split(in.CreatedAt, ",") {
			PmsTransactionListIn.CreatedAt = append(PmsTransactionListIn.CreatedAt, gtime.New(timeItem))
		}
	}

	if !g.IsEmpty(in.PayTime) {
		for _, timeItem := range strings.Split(in.PayTime, ",") {
			PmsTransactionListIn.PayTime = append(PmsTransactionListIn.PayTime, gtime.New(timeItem))
		}
	}

	if list, _, err = s.List(ctx, PmsTransactionListIn); err != nil {
		return
	}

	if payChannelList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_channel"}); err != nil {
		return
	}
	if payTypeList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_type"}); err != nil {
		return
	}
	if payStatusList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_status"}); err != nil {
		return
	}
	if refundStatusList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "refund_status"}); err != nil {
		return
	}

	for k, v := range list {
		for _, v1 := range payChannelList {
			if v.PayChannel == v1.Value {
				list[k].PayChannel = v1.Label
			}
		}
		for _, v2 := range payTypeList {
			if v.PayType == v2.Value {
				list[k].PayType = v2.Label
			}
		}
		for _, v2 := range payStatusList {
			if v.PayStatus == v2.Value {
				list[k].PayStatus = v2.Label
			}
		}
		for _, v2 := range refundStatusList {
			if v.RefundStatus == v2.Value {
				list[k].RefundStatus = v2.Label
			}
		}
	}

	if tags, err = convert.GetEntityDescTags(input_pay.PmsTransactionExportModel{}); err != nil {
		return
	}

	var (
		fileName  = "导出支付流水表-" + gctx.CtxId(ctx)
		sheetName = "支付流水"
		exports   []input_pay.PmsTransactionExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sPayService) View(ctx context.Context, in *input_pay.PmsTransactionViewInp) (res *input_pay.PmsTransactionViewModel, err error) {
	if err = dao.PmsTransaction.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取支付流水表信息，请稍后重试！")
		return
	}
	return
}

func (s *sPayService) GetTransactionSn(ctx context.Context, startDate string, endDate string, payChannel string, payType string) (transactionSnArr []string, err error) {
	mod := dao.PmsTransaction.Ctx(ctx).
		Fields("transaction_sn").
		WhereBetween(dao.PmsTransaction.Columns().PayTime, startDate, endDate).
		Where(dao.PmsTransaction.Columns().PayStatus, "DONE")

	if !g.IsEmpty(payChannel) {
		mod = mod.Where(dao.PmsTransaction.Columns().PayChannel, payChannel)
	}

	if !g.IsEmpty(payType) {
		mod = mod.Where(dao.PmsTransaction.Columns().PayType, payType)
	}

	columns, err := mod.OmitEmptyWhere().Array()
	if err != nil {
		err = gerror.Wrap(err, "获取支付流水号失败！")
		return
	}

	transactionSnArr = g.NewVar(columns).Strings()
	return
}
