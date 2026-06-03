package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/shopspring/decimal"

	"APT/api/app/travel"
)

func (c *ControllerTravel) OrderList(ctx context.Context, req *travel.OrderListReq) (res *travel.OrderListRes, err error) {
	var (
		MemberUser = contexts.GetMemberUser(ctx)
		Orders     []*struct {
			entity.TravelOrder
			ProductInfo *struct {
				gmeta.Meta `orm:"table:hg_travel_product"`
				Id         uint64 `json:"id"              dc:""`
				Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
				SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
			} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
			SkuInfo *struct {
				gmeta.Meta   `orm:"table:hg_travel_product_sku"`
				Id           uint64 `json:"id"              dc:""`
				Name         string `json:"name"           dc:"名称（默认语言；多语言存 hg_pms_language）"`
				MeetingPlace string `json:"meetingPlace"    dc:"集合地点"`
				MeetingTime  string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
				GgLat        string `json:"ggLat"           dc:"谷歌纬度"`
				GgLng        string `json:"ggLng"           dc:"谷歌经度"`
			} `json:"skuInfo" orm:"with:id=sku_id"  dc:"产品Sku"`
		}
	)

	res = new(travel.OrderListRes)

	// 构建查询条件
	mod := dao.TravelOrder.Ctx(ctx).Hook(hook2.PmsFindLanguageValueHook).Where(dao.TravelOrder.Columns().MemberId, MemberUser.Id)

	if !g.IsEmpty(req.OrderStatus) {
		// 根据前端状态映射到数据库状态
		switch req.OrderStatus {
		case "WAIT_PAY":
			mod = mod.Where(dao.TravelOrder.Columns().OrderStatus, "WAIT_PAY")
		case "ORDER_SUCCESS":
			mod = mod.WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"WAIT_VERIFY", "DONE"})
		case "CANCEL":
			mod = mod.WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"CANCEL", "REFUND", "OVERDUE"})
		case "WAIT_START":
			mod = mod.Where(dao.TravelOrder.Columns().OrderStatus, "WAIT_VERIFY")
		case "DONE":
			mod = mod.Where(dao.TravelOrder.Columns().OrderStatus, "DONE")
		case "INVALID":
			mod = mod.WhereIn(dao.TravelOrder.Columns().OrderStatus, []string{"CANCEL", "REFUND", "OVERDUE"})
		}
	}

	// 使用WithAll查询订单列表和产品信息
	if err = mod.WithAll().
		OrderDesc(dao.TravelOrder.Columns().Id).
		Page(req.PageNum, req.PageSize).
		ScanAndCount(&Orders, &res.Count, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, gi18n.T(ctx, "server_exception"))
		return
	}

	// WithAll 不加载软删除记录，对产品/SKU 已删除的情况手动补查
	for _, order := range Orders {
		if order.ProductId > 0 && order.ProductInfo == nil {
			var product *struct {
				gmeta.Meta `orm:"table:hg_travel_product"`
				Id         uint64 `json:"id"              dc:""`
				Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
				SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
			}
			if e := dao.TravelProduct.Ctx(ctx).Unscoped().
				Hook(hook2.PmsFindLanguageValueHook).
				Fields("id,title,sub_title").
				WherePri(order.ProductId).Scan(&product); e == nil && product != nil {
				order.ProductInfo = product
			}
		}
		if order.SkuId > 0 && order.SkuInfo == nil {
			var sku *struct {
				gmeta.Meta   `orm:"table:hg_travel_product_sku"`
				Id           uint64 `json:"id"              dc:""`
				Name         string `json:"name"           dc:"名称（默认语言；多语言存 hg_pms_language）"`
				MeetingPlace string `json:"meetingPlace"    dc:"集合地点"`
				MeetingTime  string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
				GgLat        string `json:"ggLat"           dc:"谷歌纬度"`
				GgLng        string `json:"ggLng"           dc:"谷歌经度"`
			}
			if e := dao.TravelProductSku.Ctx(ctx).Unscoped().
				Hook(hook2.PmsFindLanguageValueHook).
				Fields("id,name,meeting_place,meeting_time,gg_lat,gg_lng").
				WherePri(order.SkuId).Scan(&sku); e == nil && sku != nil {
				order.SkuInfo = sku
			}
		}
	}

	// 构建返回数据
	for _, order := range Orders {
		item := &travel.OrderItem{
			Id:              int(order.Id),
			OrderSn:         order.OrderSn,
			BookDate:        order.BookDate.Format("Y-m-d"),
			BookingNum:      order.BookingNum,
			OrderStatus:     order.OrderStatus,
			OrderAmount:     order.OrderAmount,
			CreateOrderTime: order.CreatedAt.Format("Y-m-d H:i:s"),
			RefundStatus:    order.RefundStatus,
			IsFx:            "N", // 默认设为N
		}

		// 设置产品信息
		if order.ProductInfo != nil {
			item.ProductTitle = order.ProductInfo.Title
		}

		// 设置Sku信息
		if order.SkuInfo != nil {
			item.SkuName = order.SkuInfo.Name
			item.MeetingPlace = order.SkuInfo.MeetingPlace
			item.MeetingTime = order.SkuInfo.MeetingTime
		}

		// 计算倒计时
		if order.OrderStatus == "WAIT_PAY" && order.ExpirationTime > 0 {
			item.Countdown = order.ExpirationTime - gvar.New(gtime.Now().Unix()).Int()
			if item.Countdown < 0 {
				item.Countdown = 0
			}
		}

		// 计算第三方支付金额（订单金额-积分支付金额-优惠券金额）
		item.ActualAmount, _ = decimal.NewFromFloat(order.OrderAmount).Sub(decimal.NewFromFloat(order.BalAmount)).Sub(decimal.NewFromFloat(order.CouponAmount)).Float64()

		// 设置取消状态（根据系统配置）
		config, err := service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
			Group: "travelCancelPolicy",
		})
		item.IsCancel = "N" // 默认不允许取消
		if err == nil && config != nil && config.List != nil {
			allowCancel := gvar.New(config.List["allowCancel"]).Int()
			if allowCancel == 1 {
				item.IsCancel = "Y"
			}
		}

		res.List = append(res.List, item)
	}

	return
}

func (c *ControllerTravel) OrderDetail(ctx context.Context, req *travel.OrderDetailReq) (res *travel.OrderDetailRes, err error) {
	var (
		Order *struct {
			*entity.TravelOrder
			ProductInfo *struct {
				gmeta.Meta `orm:"table:hg_travel_product"`
				Id         uint64 `json:"id"              dc:""`
				Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
				SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
			} `json:"productInfo" orm:"with:id=product_id"  dc:"产品"`
			SkuInfo *struct {
				gmeta.Meta       `orm:"table:hg_travel_product_sku"`
				Id               uint64 `json:"id"              dc:""`
				Name             string `json:"name"           dc:"名称（默认语言；多语言存 hg_pms_language）"`
				ContactMobile    string `json:"contactMobile"   dc:"联系电话"`
				MeetingPlace     string `json:"meetingPlace"    dc:"集合地点"`
				MeetingTime      string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
				GgLat            string `json:"ggLat"           dc:"谷歌纬度"`
				GgLng            string `json:"ggLng"           dc:"谷歌经度"`
				AllowCancel      int    `json:"allowCancel"      dc:"是否允许取消"`
				FreeCancelHours  int    `json:"freeCancelHours"  dc:"几小时前免费"`
				CancelFeePercent int    `json:"cancelFeePercent" dc:"取消费率%"`
			} `json:"skuInfo" orm:"with:id=sku_id"  dc:"产品Sku"`
		}
	)

	res = new(travel.OrderDetailRes)
	res.OrderSn = req.OrderSn

	// 使用WithAll查询订单信息
	if err = dao.TravelOrder.Ctx(ctx).Hook(hook2.PmsFindLanguageValueHook).WithAll().
		Where(dao.TravelOrder.Columns().OrderSn, req.OrderSn).
		Where(dao.TravelOrder.Columns().MemberId, uint64(contexts.GetMemberUser(ctx).Id)).
		Scan(&Order); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "server_exception"))
		return
	}

	if g.IsEmpty(Order) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	// WithAll 不加载软删除记录，对产品/SKU 已删除的情况手动补查
	if Order.ProductId > 0 && Order.ProductInfo == nil {
		var product *struct {
			gmeta.Meta `orm:"table:hg_travel_product"`
			Id         uint64 `json:"id"              dc:""`
			Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
			SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
		}
		if e := dao.TravelProduct.Ctx(ctx).Unscoped().
			Hook(hook2.PmsFindLanguageValueHook).
			Fields("id,title,sub_title").
			WherePri(Order.ProductId).Scan(&product); e == nil && product != nil {
			Order.ProductInfo = product
		}
	}
	if Order.SkuId > 0 && Order.SkuInfo == nil {
		var sku *struct {
			gmeta.Meta       `orm:"table:hg_travel_product_sku"`
			Id               uint64 `json:"id"              dc:""`
			Name             string `json:"name"           dc:"名称（默认语言；多语言存 hg_pms_language）"`
			ContactMobile    string `json:"contactMobile"   dc:"联系电话"`
			MeetingPlace     string `json:"meetingPlace"    dc:"集合地点"`
			MeetingTime      string `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
			GgLat            string `json:"ggLat"           dc:"谷歌纬度"`
			GgLng            string `json:"ggLng"           dc:"谷歌经度"`
			AllowCancel      int    `json:"allowCancel"      dc:"是否允许取消"`
			FreeCancelHours  int    `json:"freeCancelHours"  dc:"几小时前免费"`
			CancelFeePercent int    `json:"cancelFeePercent" dc:"取消费率%"`
		}
		if e := dao.TravelProductSku.Ctx(ctx).Unscoped().
			Hook(hook2.PmsFindLanguageValueHook).
			Fields("id,name,contact_mobile,meeting_place,meeting_time,gg_lat,gg_lng,allow_cancel,free_cancel_hours,cancel_fee_percent").
			WherePri(Order.SkuId).Scan(&sku); e == nil && sku != nil {
			Order.SkuInfo = sku
		}
	}

	// 设置基本信息
	res.Id = int(Order.Id)
	res.OrderSn = Order.OrderSn
	res.OrderStatus = Order.OrderStatus
	res.CreateOrderTime = Order.CreatedAt.String()
	res.IsFx = Order.IsFx
	res.MemberId = int(Order.MemberId)
	res.OrderAmount = Order.OrderAmount

	// 设置取消权限及取消政策（根据SKU配置）
	if Order.SkuInfo != nil {
		allowCancel := Order.SkuInfo.AllowCancel
		res.CanCancel = allowCancel == 1

		if allowCancel == 2 {
			res.CancelFreeDate = ""
			res.CancelPolicy = gi18n.T(ctx, "cancel_policy_no_cancel")
		} else if allowCancel == 1 {
			freeCancelHours := Order.SkuInfo.FreeCancelHours
			cancelFeePercent := Order.SkuInfo.CancelFeePercent
			meetingTimeStr := Order.SkuInfo.MeetingTime

			// 以预定日期 + SKU 集合时间作为服务开始时间，再减去免费取消小时数
			// 默认服务开始时间为预定日期 00:00
			serviceStartStr := Order.BookDate.Format("Y-m-d") + " 00:00:00"
			if meetingTimeStr != "" {
				serviceStartStr = Order.BookDate.Format("Y-m-d") + " " + meetingTimeStr + ":00"
			}
			serviceStart, parseErr := gtime.StrToTime(serviceStartStr)
			var cancelFreeDateStr string
			if parseErr == nil {
				cancelFreeTime := serviceStart.Add(-time.Duration(freeCancelHours) * time.Hour)
				cancelFreeDateStr = cancelFreeTime.Format("Y-m-d H:i")
			} else {
				cancelFreeDateStr = Order.BookDate.Format("Y-m-d")
			}
			res.CancelFreeDate = cancelFreeDateStr
			res.CancelPolicy = gi18n.Tf(ctx, "cancel_policy_free_before", gvar.New(freeCancelHours).String(), cancelFeePercent)
		}
	}

	// 设置时间信息
	res.CancelTime = Order.CancelTime.String()
	res.PayTime = Order.PayTime.String()
	res.VerifyTime = Order.VerifyTime.String()
	res.RefundTime = Order.RefundTime.String()
	res.PayModel = Order.PayModel

	// 设置产品信息
	if Order.ProductInfo != nil {
		res.ProductInfo.Id = Order.ProductInfo.Id
		res.ProductInfo.Title = Order.ProductInfo.Title
		res.ProductInfo.SubTitle = Order.ProductInfo.SubTitle
	}

	// 设置Sku信息
	if Order.SkuInfo != nil {
		res.SkuInfo.Id = Order.SkuInfo.Id
		res.SkuInfo.Name = Order.SkuInfo.Name
		res.SkuInfo.ContactMobile = Order.SkuInfo.ContactMobile
		res.SkuInfo.MeetingPlace = Order.SkuInfo.MeetingPlace
		res.SkuInfo.MeetingTime = Order.SkuInfo.MeetingTime
		res.SkuInfo.GgLat = Order.SkuInfo.GgLat
		res.SkuInfo.GgLng = Order.SkuInfo.GgLng

		res.ProductInfo.ContactMobile = Order.SkuInfo.ContactMobile
		res.ProductInfo.MeetingPlace = Order.SkuInfo.MeetingPlace
		res.ProductInfo.MeetingTime = Order.SkuInfo.MeetingTime
		res.ProductInfo.GgLat = Order.SkuInfo.GgLat
		res.ProductInfo.GgLng = Order.SkuInfo.GgLng
	}

	// 设置预订信息
	res.BookInfo.BookDate = Order.BookDate.Format("Y-m-d")
	res.BookInfo.BookNum = Order.BookingNum
	res.BookInfo.BookingName = Order.BookingName
	res.BookInfo.BookingMobile = Order.BookingMobile
	res.BookInfo.PhoneArea = Order.PhoneArea
	res.BookInfo.BookingEmail = Order.BookingEmail

	// 直接使用订单字段
	res.CancelFee = Order.CancelFee
	res.BalAmount = Order.BalAmount
	res.CouponAmount = Order.CouponAmount
	// 计算第三方支付金额（订单金额-积分支付金额-优惠券金额）
	res.ActualAmount, _ = decimal.NewFromFloat(Order.OrderAmount).Sub(decimal.NewFromFloat(Order.BalAmount)).Sub(decimal.NewFromFloat(Order.CouponAmount)).Float64()
	res.RefundBalAmount = Order.RefundBalAmount
	res.RefundAmount = Order.RefundAmount
	res.RefundActualAmount, _ = decimal.NewFromFloat(Order.RefundAmount).Sub(decimal.NewFromFloat(Order.RefundBalAmount)).Float64()

	// 查询退款记录
	var refundLogs []*entity.OrderRefundLog
	if err = dao.OrderRefundLog.Ctx(ctx).
		Where(dao.OrderRefundLog.Columns().OrderSn, Order.OrderSn).
		Where(dao.OrderRefundLog.Columns().RefundStatus, "DONE").
		OrderDesc(dao.OrderRefundLog.Columns().Id).
		Scan(&refundLogs); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	// 处理退款记录
	for _, refund := range refundLogs {

		res.RefundRecordList = append(res.RefundRecordList, &travel.RefundRecordItem{
			RefundType:   refund.RefundType,
			RefundAmount: refund.RefundAmount,
			ApplyTime:    refund.CreatedAt.Format("Y-m-d H:i:s"),
		})
	}

	// 计算支付金额
	if Order.OrderStatus == "WAIT_PAY" {
		res.PayAmount = res.OrderAmount - res.BalAmount - res.CouponAmount
	} else {
		res.PayAmount = res.OrderAmount
	}

	// 计算倒计时
	if Order.OrderStatus == "WAIT_PAY" && Order.ExpirationTime > 0 {
		res.Countdown = Order.ExpirationTime - int(gtime.Now().Unix())
		if res.Countdown < 0 {
			res.Countdown = 0
		}
	}

	return
}
func (c *ControllerTravel) MemberVerifyCodeRefresh(ctx context.Context, req *travel.MemberVerifyCodeRefreshReq) (res *travel.MemberVerifyCodeRefreshRes, err error) {
	res = new(travel.MemberVerifyCodeRefreshRes)
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res.TravelOrderRefreshCodeModel, err = service.TravelOrder().RefreshCode(ctx, &input_travel.TravelOrderRefreshCodeInp{
		OrderSn:  req.OrderSn,
		MemberId: MemberInfo.Id,
	})
	if err != nil {
		return
	}
	return
}
