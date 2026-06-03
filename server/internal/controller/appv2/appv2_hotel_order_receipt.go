package appv2

import (
	"APT/internal/dao"
	"APT/internal/library/hgrds/lock"
	"APT/internal/library/storager"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/api/appv2/hotel"
)

func (c *ControllerHotel) OrderReceipt(ctx context.Context, req *hotel.OrderReceiptReq) (res *hotel.OrderReceiptRes, err error) {
	var (
		Receipt      *entity.Receipt
		PmsAppStay   *entity.PmsAppStay
		PmsProperty  *entity.PmsProperty
		PmsLanguage  *entity.PmsLanguage
		EmsParams    *input_basics.SendEmsInp
		FilePath     string
		UploadConfig *model.UploadConfig
	)
	l := lock.Mutex(fmt.Sprintf("Receipt_%s", req.OrderSn))
	if err = l.TryLock(ctx); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "try_again_later"))
		return
	}
	defer func() {
		if LockErr := l.Unlock(ctx); LockErr != nil {
			g.Log().Errorf(ctx, "解锁失败：%s", LockErr.Error())
		}
	}()
	if err = dao.Receipt.Ctx(ctx).Where(g.MapStrAny{
		dao.Receipt.Columns().OrderSn: req.OrderSn,
	}).Scan(&Receipt); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(Receipt) {

		if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsAppStay.Columns().OrderSn: req.OrderSn,
		}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(PmsAppStay) {
			// 未找到该订单信息
			err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
			return
		}
		if PmsAppStay.OrderStatus != "HAVE_PAID" {
			// 该订单未支付无法申领
			err = gerror.New(gi18n.T(ctx, "order_has_not_paid_and_cannot_be_claimed"))
			return
		}
		// 查询物业信息
		if err = dao.PmsProperty.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsProperty.Columns().Uid: PmsAppStay.Puid,
		}).Scan(&PmsProperty); err != nil && errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(PmsProperty) {
			// 未找到该物业信息
			err = gerror.New(gi18n.T(ctx, "property_info_not_found"))
			return
		}
		// 查询物业名
		if err = dao.PmsLanguage.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsLanguage.Columns().Uuid:     PmsProperty.Name,
			dao.PmsLanguage.Columns().Language: "ja",
		}).Scan(&PmsLanguage); err != nil && errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(PmsLanguage) {
			PmsLanguage.Content = "住一"
		}
		if FilePath, err = service.BasicsPDF().OrderReceipt(ctx, &input_basics.PdfLssIpt{
			UserName:     req.UserName,
			OrderSn:      req.OrderSn,
			OrderAmount:  gconv.Int(PmsAppStay.OrderAmount),
			StartTime:    PmsAppStay.CheckInDate,
			EndTime:      PmsAppStay.CheckOutDate,
			CreateTime:   PmsAppStay.CreatedAt.Format("Y.m.d"),
			PropertyName: PmsLanguage.Content,
		}); err != nil {
			return
		}
		// 上传文件
		if UploadConfig, err = service.BasicsConfig().GetUpload(ctx); err != nil {
			return
		}
		g.Log().Debug(ctx, "上传文件：", FilePath)
		// 上传到驱动
		if FilePath, err = storager.New(UploadConfig.Drive).UploadFile(ctx, FilePath); err != nil {
			return
		}
		g.Log().Debug(ctx, "云存储：", FilePath)
		FilePath = storager.LastUrl(ctx, FilePath, UploadConfig.Drive)
		g.Log().Debug(ctx, "云存储：", FilePath)
		// 插入订单发送记录表
		if _, err = dao.Receipt.Ctx(ctx).Data(g.MapStrAny{
			dao.Receipt.Columns().UserName:    req.UserName,
			dao.Receipt.Columns().OrderSn:     req.OrderSn,
			dao.Receipt.Columns().ReceiptPath: FilePath,
		}).Insert(); err != nil {
			return
		}
	} else {
		FilePath = Receipt.ReceiptPath
	}

	EmsParams = new(input_basics.SendEmsInp)
	EmsParams.Content = `
<p>お客様へ</p>
<p style="text-indent:2em">ご注文の領収書を添付いたしましたので、ご確認ください。</p>
<p style="text-indent:2em">Dear Customer, please find attached the receipt for your order. Kindly check it.</p>
<p style="text-indent:2em">親愛的顧客您好，已隨信附上您的訂單收據，敬請查收。</p>
<a href="` + FilePath + `">領収書/电子收据下载Download</a>`
	EmsParams.Event = "text"
	EmsParams.Email = req.Mail
	if err = service.BasicsEmsLog().Send(ctx, EmsParams); err != nil {
		return
	}
	return
}

func (c *ControllerHotel) OrderReceiptView(ctx context.Context, req *hotel.OrderReceiptViewReq) (res *hotel.OrderReceiptViewRes, err error) {

	if err = dao.Receipt.Ctx(ctx).Where(g.MapStrAny{
		dao.Receipt.Columns().OrderSn: req.OrderSn,
	}).Scan(&res); err != nil && errors.Is(err, sql.ErrNoRows) {
		// 获取发票信息失败，请稍后重试
		err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_get_invoice_info"))
		return
	}

	return
}
