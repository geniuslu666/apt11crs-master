package logic_refund

import (
	"APT/internal/dao"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRefundTransaction struct{}

func NewRefundTransaction() *sRefundTransaction {
	return &sRefundTransaction{}
}

func init() {
	service.RegisterRefundTransaction(NewRefundTransaction())
}

func (s *sRefundTransaction) List(ctx context.Context, in *input_refund.PmsTransactionRefundListInp) (list []*input_refund.PmsTransactionRefundListModel, totalCount int, err error) {
	mod := dao.PmsTransactionRefund.Ctx(ctx)

	mod = mod.Fields(input_refund.PmsTransactionRefundListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsTransactionRefund.Columns().Id, in.Id)
	}

	if in.OrderSn != "" {
		mod = mod.WhereLike(dao.PmsTransactionRefund.Columns().OrderSn, in.OrderSn)
	}

	if in.TransactionSn != "" {
		mod = mod.WhereLike(dao.PmsTransactionRefund.Columns().TransactionSn, in.TransactionSn)
	}

	if !g.IsEmpty(in.RefundChannel) {
		mod = mod.Where(dao.PmsTransactionRefund.Columns().RefundChannel, in.RefundChannel)
	}

	if !g.IsEmpty(in.RefundType) {
		mod = mod.Where(dao.PmsTransactionRefund.Columns().RefundType, in.RefundType)
	}

	if !g.IsEmpty(in.RefundStatus) {
		mod = mod.Where(dao.PmsTransactionRefund.Columns().RefundStatus, in.RefundStatus)
	}

	if len(in.RefundTime) == 2 {
		mod = mod.WhereBetween(dao.PmsTransactionRefund.Columns().RefundTime, in.RefundTime[0], in.RefundTime[1])
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsTransactionRefund.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.PageReq.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.PmsTransactionRefund.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取退款流水列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sRefundTransaction) Export(ctx context.Context, in *input_refund.PmsTransactionRefundListExportInp) (err error) {
	var PmsTransactionRefundListIn *input_refund.PmsTransactionRefundListInp

	// 无需分页
	in.PageReq.Pagination = false

	PmsTransactionRefundListIn = &input_refund.PmsTransactionRefundListInp{
		PageReq:       in.PageReq,
		Id:            in.Id,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.RefundChannel,
		RefundType:    in.RefundType,
		RefundStatus:  in.RefundStatus,
	}

	if !g.IsEmpty(in.CreatedAt) {
		for _, timeItem := range strings.Split(in.CreatedAt, ",") {
			PmsTransactionRefundListIn.CreatedAt = append(PmsTransactionRefundListIn.CreatedAt, gtime.New(timeItem))
		}
	}

	if !g.IsEmpty(in.RefundTime) {
		for _, timeItem := range strings.Split(in.RefundTime, ",") {
			PmsTransactionRefundListIn.RefundTime = append(PmsTransactionRefundListIn.RefundTime, gtime.New(timeItem))
		}
	}

	// 获取导出数据
	list, _, err := s.List(ctx, PmsTransactionRefundListIn)
	if err != nil {
		return
	}

	payChannelList, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_channel"})
	payTypeList, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "pay_type"})
	payStatusList, err := service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "refund_status"})

	for k, v := range list {
		for _, v1 := range payChannelList {
			if v.RefundChannel == v1.Value {
				list[k].RefundChannel = v1.Label
			}
		}
		for _, v2 := range payTypeList {
			if v.RefundType == v2.Value {
				list[k].RefundType = v2.Label
			}
		}

		for _, v2 := range payStatusList {
			if v.RefundStatus == v2.Value {
				list[k].RefundStatus = v2.Label
			}
		}
	}

	tags, err := convert.GetEntityDescTags(input_refund.PmsTransactionRefundExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出退款流水-" + gctx.CtxId(ctx)
		sheetName = "退款流水"
		exports   []input_refund.PmsTransactionRefundExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sRefundTransaction) View(ctx context.Context, in *input_refund.PmsTransactionRefundViewInp) (res *input_refund.PmsTransactionRefundViewModel, err error) {
	if err = dao.PmsTransactionRefund.Ctx(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取退款流水信息，请稍后重试！")
		return
	}
	return
}
