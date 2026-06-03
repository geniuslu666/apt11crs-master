package notify

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/api/notify/hook"
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/toretaApi"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
)

func (c *ControllerHook) Toreta(ctx context.Context, req *hook.ToretaReq) (res *hook.ToretaRes, err error) {
	var (
		r              = ghttp.RequestFromCtx(ctx)
		Logger         = g.Log().Path("logs/HOOK/TORETA")
		foodRestaurant *entity.FoodRestaurant
		foodOrder      *entity.FoodOrder
	)

	// 记录接收到的webhook数据
	Logger.Info(ctx, "Toreta webhook received:", gvar.New(r.GetBody()).String())
	Logger.Info(ctx, "Parsed request:", req)

	// 验证必要参数
	// 必要字段校验
	required := []string{"restaurant_key", "resource_key", "resource_action", "resource_type", "updated_at"}
	// 使用 map 便于逐项检查
	bodyMap := g.Map{}
	bodyMap = gconv.Map(req)
	for _, k := range required {
		if gconv.String(bodyMap[k]) == "" {
			Logger.Error(ctx, "缺少必要字段: "+k)
			r.Response.WriteStatus(400)
			r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 400, "detail": "Invalid Request Parameter: " + k + " is invalid."}})
			return
		}
	}

	// 查看餐厅是否存在
	if err = dao.FoodRestaurant.Ctx(ctx).Where("toreta_id", req.RestaurantKey).Scan(&foodRestaurant); err != nil {
		Logger.Error(ctx, "Failed to get restaurant:", err)
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 400, "detail": "Invalid Request Parameter: restaurant_key is invalid."}})
		return
	}

	if g.IsEmpty(foodRestaurant) {
		Logger.Error(ctx, "Restaurant not found")
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 400, "detail": "Invalid Request Parameter: restaurant_key is invalid."}})
		return
	}

	// 查看订单是否存在
	if err = dao.FoodOrder.Ctx(ctx).Where("restaurant_id", foodRestaurant.Id).Where("toreta_reservation_id", req.ResourceKey).Scan(&foodOrder); err != nil {
		Logger.Error(ctx, "Failed to get order:", err)
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 400, "detail": "Invalid Request Parameter: resource_key is invalid."}})
		return
	}

	if g.IsEmpty(foodOrder) {
		Logger.Error(ctx, "Order not found")
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 400, "detail": "Invalid Request Parameter: resource_key is invalid."}})
		return
	}

	// 目前仅支持预约类型的通知
	if req.ResourceType != "Reservation" {
		Logger.Warning(ctx, "Unsupported resource type:", req.ResourceType)
		r.Response.WriteJsonExit(nil)
		return
	}

	// 处理预约通知
	if err = c.handleReservationNotification(ctx, req, Logger); err != nil {
		Logger.Error(ctx, "Failed to handle reservation notification:", err)
		r.Response.WriteStatus(500)
		r.Response.WriteJsonExit(g.Map{"error": g.Map{"status": 500, "detail": "Internal server error"}})
		return
	}

	// 返回成功响应
	r.Response.WriteStatus(200)
	r.Response.WriteJsonExit(g.Map{"status": 200, "detail": "Success"})
	return
}

// handleReservationNotification 处理预约通知
func (c *ControllerHook) handleReservationNotification(ctx context.Context, req *hook.ToretaReq, Logger *glog.Logger) (err error) {
	var (
		toretaConfig *model.ToretaApiConfig
		client       *toretaApi.ToretaClient
		reservation  *toretaApi.ReservationDetailResponse
	)

	// 获取Toreta API配置
	if toretaConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		Logger.Error(ctx, "Failed to get Toreta API config:", err)
		return
	}

	// 创建Toreta客户端
	client = toretaApi.NewClient(ctx, toretaConfig)

	// 根据操作类型处理
	switch req.ResourceAction {
	case "create":
		Logger.Info(ctx, "Processing reservation creation for:", req.ResourceKey)
		// 获取预约详情并创建本地订单
		if reservation, err = client.ReservationDetail(ctx, &toretaApi.ReservationDetailParams{
			ReservationId: req.ResourceKey,
		}); err != nil {
			Logger.Error(ctx, "Failed to get reservation details:", err)
			return
		}
		Logger.Info(ctx, "Reservation details:", reservation)
		// 创建本地订单逻辑
		if err = c.createLocalOrder(ctx, req, reservation, Logger); err != nil {
			Logger.Error(ctx, "Failed to create local order:", err)
			return
		}

	case "update":
		Logger.Info(ctx, "Processing reservation update for:", req.ResourceKey)
		// 获取预约详情并更新本地订单状态
		if reservation, err = client.ReservationDetail(ctx, &toretaApi.ReservationDetailParams{
			ReservationId: req.ResourceKey,
		}); err != nil {
			Logger.Error(ctx, "Failed to get reservation details:", err)
			return
		}
		Logger.Info(ctx, "Updated reservation details:", reservation)
		// 更新本地订单状态逻辑
		if err = c.updateLocalOrder(ctx, req, reservation, Logger); err != nil {
			Logger.Error(ctx, "Failed to update local order:", err)
			return
		}

	case "destroy":
		Logger.Info(ctx, "Processing reservation cancellation for:", req.ResourceKey)
		// 取消本地订单
		if err = c.cancelLocalOrder(ctx, req, Logger); err != nil {
			Logger.Error(ctx, "Failed to cancel local order:", err)
			return
		}

	default:
		Logger.Warning(ctx, "Unknown resource action:", req.ResourceAction)
	}

	return
}

// createLocalOrder 创建本地订单（当Toreta有新预约创建时）
func (c *ControllerHook) createLocalOrder(ctx context.Context, req *hook.ToretaReq, reservation *toretaApi.ReservationDetailResponse, Logger *glog.Logger) (err error) {
	// 注意：通常情况下，Toreta的预约是通过我们的系统创建的，所以create事件可能不需要特殊处理
	// 这里主要是记录日志，实际的订单创建逻辑在预约接口中已经处理
	Logger.Info(ctx, "Toreta reservation created - this might be from our own system")
	Logger.Info(ctx, "Reservation ID:", req.ResourceKey)
	Logger.Info(ctx, "Restaurant Key:", req.RestaurantKey)

	// 如果需要，可以在这里添加订单状态同步逻辑
	// 比如确保本地订单状态与Toreta状态一致

	// Toreta状态映射：
	// 0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置
	var (
		toretaStatus int
	)

	// 解析Toreta状态
	if reservation.Status != "" {
		toretaStatus = gconv.Int(reservation.Status)
	}

	// 更新数据库中的订单状态
	updateData := g.Map{
		"toreta_reservation_status": toretaStatus,
	}

	// 如果有结束时间，也更新
	if reservation.EndAt > 0 {
		endTime := gtime.NewFromTimeStamp(int64(reservation.EndAt))
		updateData["toreta_reservation_endtime"] = endTime
	}

	// 根据Toreta预约ID更新订单
	if _, err = dao.FoodOrder.Ctx(ctx).
		Where("toreta_reservation_id", req.ResourceKey).
		Data(updateData).
		Update(); err != nil {
		Logger.Error(ctx, "Failed to create order in database:", err)
		return
	}

	Logger.Info(ctx, "Successfully create local order status")
	return
}

// updateLocalOrder 更新本地订单状态（当Toreta预约状态发生变化时）
func (c *ControllerHook) updateLocalOrder(ctx context.Context, req *hook.ToretaReq, reservation *toretaApi.ReservationDetailResponse, Logger *glog.Logger) (err error) {
	Logger.Info(ctx, "Updating local order for Toreta reservation:", req.ResourceKey)

	// 根据Toreta预约状态更新本地订单
	// Toreta状态映射：
	// 0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置

	var (
		foodOrder    *entity.FoodOrder
		toretaStatus int
	)

	// 获取订单
	if err = dao.FoodOrder.Ctx(ctx).Where("toreta_reservation_id", req.ResourceKey).Scan(&foodOrder); err != nil {
		Logger.Error(ctx, "Failed to get order in database:", err)
		return
	}

	// 解析Toreta状态
	if reservation.Status != "" {
		toretaStatus = gconv.Int(reservation.Status)
	}

	// 更新数据库中的订单状态
	updateData := g.Map{
		"toreta_reservation_status": toretaStatus,
		"booking_count":             reservation.Seats,
		"goods_num":                 reservation.Seats,
	}

	// 判断人数是否有变化
	if foodOrder.BookingCount != reservation.Seats {
		updateData["old_booking_count"] = foodOrder.BookingCount
	}

	// 如果有开始时间，也更新
	if reservation.StartAt > 0 {
		startTime := gtime.NewFromTimeStamp(int64(reservation.StartAt))
		updateData["book_date"] = startTime.Format("Y-m-d")
		updateData["book_time"] = startTime.Format("H:i")
		updateData["book_datetime"] = startTime

		// 判断预定日期时间是否变化
		if foodOrder.BookDate != startTime.Format("Y-m-d") || foodOrder.BookTime != startTime.Format("H:i") {
			updateData["old_book_date"] = foodOrder.BookDate
			updateData["old_book_time"] = foodOrder.BookTime
			updateData["old_book_datetime"] = foodOrder.BookDatetime
		}
	}

	// 如果有结束时间，也更新
	if reservation.EndAt > 0 {
		endTime := gtime.NewFromTimeStamp(int64(reservation.EndAt))
		updateData["toreta_reservation_endtime"] = endTime
	}

	// 根据Toreta预约ID更新订单
	if _, err = dao.FoodOrder.Ctx(ctx).
		Where("toreta_reservation_id", req.ResourceKey).
		Data(updateData).
		Update(); err != nil {
		Logger.Error(ctx, "Failed to update order in database:", err)
		return
	}

	// 如果状态码返回的是2：预约取消，则订单进行退款取消处理
	if toretaStatus == 2 {
		var models *entity.FoodOrder
		if err = dao.FoodOrder.Ctx(ctx).Where("toreta_reservation_id", req.ResourceKey).Scan(&models); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			// 计算退款金额
			var (
				Transaction       []*entity.PmsTransaction
				TransactionRefund []*entity.PmsTransactionRefund
				RefundBalance     float64 // 可退款积分
				RefundFee         float64 // 第三方支付
			)
			if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).Scan(&Transaction); err != nil {
				return
			}
			if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
				return
			}

			for _, v := range Transaction {
				// 可退款金额
				Refundable := v.Amount - v.RefundAmount
				if v.PayType == "BAL" {
					RefundBalance = RefundBalance + Refundable
				} else if v.PayType == "COUPON" {

				} else {
					RefundFee = RefundFee + Refundable
				}
			}
			// 退款金额
			RefundAmount := RefundBalance + RefundFee

			// 修改订单状态
			RefundStatus := "DONE"
			//if models.OrderAmount > RefundAmount {
			//	RefundStatus = "PART"
			//}
			if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).
				WherePri(models.Id).Data(g.MapStrAny{
				dao.FoodOrder.Columns().RefundAmount:        RefundAmount,
				dao.FoodOrder.Columns().RefundBalAmount:     RefundBalance,
				dao.FoodOrder.Columns().RefundCouponAmount:  0,
				dao.FoodOrder.Columns().RefundStatus:        RefundStatus,
				dao.FoodOrder.Columns().RefundTime:          gtime.Now(),
				dao.FoodOrder.Columns().OrderStatus:         "CANCEL",
				dao.FoodOrder.Columns().BookingStatus:       "CANCEL",
				dao.FoodOrder.Columns().DepositPayStatus:    "CANCEL",
				dao.FoodOrder.Columns().RemainPayStatus:     "CANCEL",
				dao.FoodOrder.Columns().DepositCancelReason: "Toreta预约取消",
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 更新餐厅预定量和预定金额
			if _, err = dao.FoodRestaurant.Ctx(ctx).TX(tx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().PayOrderNum:    gdb.Raw("pay_order_num-1"),
				dao.FoodRestaurant.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}

			// 更新套餐预定量和预定金额
			if _, err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
				dao.FoodGoods.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", models.GoodsNum)),
				dao.FoodGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新套餐信息失败，请稍后重试！")
				return
			}

			// 订单日志
			if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
				OrderId:     int(models.Id),
				ActionWay:   "REFUND",
				Remark:      "订单取消",
				OperateType: "SYSTEM",
			}); err != nil {
				return err
			}

			if RefundAmount > 0 {
				// 全额退款
				err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
					OrderSn:      models.OrderSn,
					RefundAmount: RefundAmount,
				}, tx)
				if err != nil {
					err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
					return
				}
			}

			return
		})
		if err != nil {
			return
		}
	}

	Logger.Info(ctx, "Successfully updated local order status")
	return
}

// cancelLocalOrder 取消本地订单（当Toreta预约被删除时）
func (c *ControllerHook) cancelLocalOrder(ctx context.Context, req *hook.ToretaReq, Logger *glog.Logger) (err error) {
	Logger.Info(ctx, "Cancelling local order for Toreta reservation:", req.ResourceKey)
	return
}
