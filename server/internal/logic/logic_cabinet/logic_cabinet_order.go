package logic_cabinet

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/cabinetApi"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpdateOrderStatus 更改订单状态
func (s *sCabinetService) UpdateOrderStatus(ctx context.Context, in *input_cabinet.UpdateOrderStatusInp) (err error) {
	var (
		CabinetOrder         *entity.CabinetOrder
		systemMessageTitle   map[string]string
		systemMessageContent map[string]string
	)
	if err = dao.CabinetOrder.Ctx(ctx).Where(g.Map{
		dao.CabinetOrder.Columns().OrderSn: in.OrderSn,
	}).Scan(&CabinetOrder); err != nil {
		return
	}
	if g.IsEmpty(CabinetOrder) {
		err = gerror.New("不存在订单")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		switch in.EventType {
		case "order_completed":
			// 订单完成
			// 先判断订单状态是否是进行中或是宽限期
			if CabinetOrder.OrderStatus == "DONE" {
				return
			}
			if CabinetOrder.OrderStatus != "ING" && CabinetOrder.OrderStatus != "GRACE" && CabinetOrder.OrderStatus != "OVERTIME" {
				err = gerror.New("订单状态不正确")
				return
			}
			// 变更为已完成
			if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
				WherePri(CabinetOrder.Id).Data(g.MapStrAny{
				dao.CabinetOrder.Columns().OrderStatus: "DONE",
				dao.CabinetOrder.Columns().FinishTime:  gtime.Now(),
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 订单日志
			if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
				OrderId:     int(CabinetOrder.Id),
				OrderStatus: "DONE",
				ActionWay:   "DONE",
				Remark:      "订单已完成",
				OperateType: "SYSTEM",
			}); err != nil {
				return err
			}

			// 发送到消息队列
			systemMessageTitle = map[string]string{
				"zh":    "订单已完成",
				"en":    "Order completed",
				"ja":    "注文完了",
				"ko":    "주문 완료",
				"zh_CN": "訂單已完成",
			}
			systemMessageContent = map[string]string{
				"zh":    CabinetOrder.OrderSn + "订单已完成",
				"en":    "Order " + CabinetOrder.OrderSn + " completed",
				"ja":    "注文番号" + CabinetOrder.OrderSn + " 注文完了",
				"ko":    "주문 " + CabinetOrder.OrderSn + " 완료",
				"zh_CN": CabinetOrder.OrderSn + "訂單已完成",
			}

			// 转发到返利队列
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameRebate,
				DataByte:     gvar.New(CabinetOrder.OrderSn).Bytes(),
				Header:       nil,
			})
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameExp,
				DataByte:     gvar.New(CabinetOrder.OrderSn).Bytes(),
				Header:       nil,
			})

			if CabinetOrder.IsFx == "Y" {
				if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
					ExchangeName: consts.RabbitMQExchangeName,
					QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
					DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
						OrderNo:      CabinetOrder.OrderSn,
						ChangeStatus: "COMPLETE",
					}).MustToJson(),
					Header: nil,
				}); err != nil {
					g.Log().Error(ctx, "发送分销订单消息变更失败", err)
					err = nil
				}
			}

			break
		case "order_timeout":
			// 订单超时
			// 先判断订单状态是否是进行中
			if CabinetOrder.OrderStatus == "OVERTIME" {
				return
			}
			if CabinetOrder.OrderStatus != "ING" {
				err = gerror.New("订单状态不正确")
				return
			}
			// 变更为已超时
			if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
				WherePri(CabinetOrder.Id).Data(g.MapStrAny{
				dao.CabinetOrder.Columns().OrderStatus: "OVERTIME",
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 订单日志
			if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
				OrderId:     int(CabinetOrder.Id),
				OrderStatus: "OVERTIME",
				ActionWay:   "OVERTIME",
				Remark:      "订单已超时",
				OperateType: "SYSTEM",
			}); err != nil {
				return err
			}

			systemMessageTitle = map[string]string{
				"zh":    "订单已超时",
				"en":    "Order timeout",
				"ja":    "注文期限が切れました",
				"ko":    "주문이 만료되었습니다",
				"zh_CN": "訂單已超時",
			}
			systemMessageContent = map[string]string{
				"zh":    CabinetOrder.OrderSn + "订单已超时",
				"en":    "Order " + CabinetOrder.OrderSn + " timeout",
				"ja":    "注文番号" + CabinetOrder.OrderSn + " 注文期限が切れました",
				"ko":    "주문 " + CabinetOrder.OrderSn + " 만료되었습니다",
				"zh_CN": CabinetOrder.OrderSn + "訂單已超時",
			}

			break
		case "order_cancelled":
			// 订单取消
			if CabinetOrder.OrderStatus == "CANCEL" {
				return
			}
			if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
				WherePri(CabinetOrder.Id).Data(g.MapStrAny{
				dao.CabinetOrder.Columns().OrderStatus: "CANCEL",
				dao.CabinetOrder.Columns().CancelTime:  gtime.Now(),
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 订单日志
			if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
				OrderId:     int(CabinetOrder.Id),
				OrderStatus: "CANCEL",
				ActionWay:   "CANCEL",
				Remark:      "订单已取消",
				OperateType: "SYSTEM",
			}); err != nil {
				return err
			}

			// TODO 是否退款待定
			break
		case "order_refunded":
			// 订单退款
			if CabinetOrder.OrderStatus == "CANCEL" {
				return
			}
			if err = s.OrderRefund(ctx, &input_cabinet.OrderRefundInp{
				OrderSn: in.OrderSn,
			}); err != nil {
				return err
			}
			break
		default:
			err = gerror.New("未知事件")
			return
		}

		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(CabinetOrder.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		appPushData := g.MapStrStr{
			"type":   "1",
			"string": CabinetOrder.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "cabinet",
			Type:                 "order",
			MemberId:             int(CabinetOrder.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/smart_locker_order_detail",
			WxLink:               fmt.Sprintf("/subpackages/lockers/pages/order-detail?orderSn=%s", CabinetOrder.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           0,
			OperatorRole:         "SYSTEM",
			OrderSn:              CabinetOrder.OrderSn,
		})

		return
	})
}

func (s *sCabinetService) OrderRefund(ctx context.Context, in *input_cabinet.OrderRefundInp) (err error) {
	var (
		CabinetOrder *entity.CabinetOrder
	)
	if err = dao.CabinetOrder.Ctx(ctx).Where(g.Map{
		dao.CabinetOrder.Columns().OrderSn: in.OrderSn,
	}).Scan(&CabinetOrder); err != nil {
		return
	}
	if g.IsEmpty(CabinetOrder) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 执行退款
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().OrderSn, CabinetOrder.OrderSn).Scan(&Transaction); err != nil {
			return
		}
		if err = dao.PmsTransactionRefund.Ctx(ctx).Where(dao.PmsTransactionRefund.Columns().OrderSn, CabinetOrder.OrderSn).Scan(&TransactionRefund); err != nil {
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
		if CabinetOrder.OrderAmount > RefundAmount {
			RefundStatus = "PART"
		}
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
			WherePri(CabinetOrder.Id).Data(g.MapStrAny{
			dao.CabinetOrder.Columns().RefundAmount:       RefundAmount,
			dao.CabinetOrder.Columns().RefundBalAmount:    RefundBalance,
			dao.CabinetOrder.Columns().RefundCouponAmount: 0,
			dao.CabinetOrder.Columns().RefundStatus:       RefundStatus,
			dao.CabinetOrder.Columns().RefundTime:         gtime.Now(),
			dao.CabinetOrder.Columns().OrderStatus:        "CANCEL",
			dao.CabinetOrder.Columns().PayStatus:          "REFUND",
			dao.CabinetOrder.Columns().CancelTime:         gtime.Now(),
		}).Update(); err != nil {
			return
		}

		// 订单日志
		if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
			OrderId:     int(CabinetOrder.Id),
			OrderStatus: "CANCEL",
			ActionWay:   "REFUND",
			Remark:      "订单已退款",
			OperateType: "SYSTEM",
		}); err != nil {
			return
		}

		// 全额退款
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      CabinetOrder.OrderSn,
			RefundAmount: RefundAmount,
			OperateType:  "SYSTEM",
		}, tx)
		if err != nil {
			return
		}

		return
	})
}

func (s *sCabinetService) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CabinetOrder.Ctx(ctx), option...)
}

func (s *sCabinetService) AppList(ctx context.Context, in *input_cabinet.OrderAppListInp) (list []*input_cabinet.OrderAppListModel, totalCount int, err error) {
	var (
		CabinetOrder      []*entity.CabinetOrder
		AddressJson       *input_cabinet.LanguageJson
		CabinetNameJson   *input_cabinet.LanguageJson
		MchBranchNameJson *input_cabinet.LanguageJson
	)
	mod := s.Model(ctx)

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.CabinetOrder.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.OrderStatus) {
		if in.OrderStatus == "ING" {
			mod = mod.WhereIn(dao.CabinetOrder.Columns().OrderStatus, g.Slice{"ING", "OVERTIME", "GRACE"})
		} else {
			mod = mod.Where(dao.CabinetOrder.Columns().OrderStatus, in.OrderStatus)
		}
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.CabinetOrder.Columns().CreatedAt)

	if in.Pagination {
		if err = mod.ScanAndCount(&CabinetOrder, &totalCount, false); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_order_list_failed"))
			return
		}
	} else {
		if err = mod.Scan(&CabinetOrder); err != nil {
			err = gerror.Wrap(err, gi18n.T(ctx, "get_order_list_failed"))
			return
		}
	}

	now := gtime.Now()

	var Language = contexts.GetLanguage(ctx)
	if !g.IsEmpty(CabinetOrder) {
		list = make([]*input_cabinet.OrderAppListModel, 0, len(CabinetOrder))
		for _, v := range CabinetOrder {
			item := &input_cabinet.OrderAppListModel{
				Id:           int(v.Id),
				OrderSn:      v.OrderSn,
				Pin:          v.Pin,
				BoxNo:        v.BoxNo,
				BoxAlias:     v.BoxAlias,
				OrderAmount:  v.OrderAmount,
				StartTime:    v.StartTime,
				EndTime:      v.EndTime,
				GraceSeconds: v.GraceSeconds,
				GraceEndTime: v.GraceEndTime,
				OrderStatus:  v.OrderStatus,
				IsFx:         v.IsFx,
			}
			if !g.IsEmpty(v.AddressJson) && !g.IsEmpty(v.CabinetNameJson) && !g.IsEmpty(v.MchBranchNameJson) {
				AddressJson = new(input_cabinet.LanguageJson)
				CabinetNameJson = new(input_cabinet.LanguageJson)
				MchBranchNameJson = new(input_cabinet.LanguageJson)
				if err = json.Unmarshal([]byte(v.AddressJson.String()), &AddressJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				if err = json.Unmarshal([]byte(v.CabinetNameJson.String()), &CabinetNameJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}
				if err = json.Unmarshal([]byte(v.MchBranchNameJson.String()), &MchBranchNameJson); err != nil {
					err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
					return
				}

				switch Language {
				case "zh":
					item.Address = AddressJson.Zh
					item.CabinetName = CabinetNameJson.Zh
					item.MchBranchName = MchBranchNameJson.Zh
				case "zh_CN":
					item.Address = AddressJson.Tw
					item.CabinetName = CabinetNameJson.Tw
					item.MchBranchName = MchBranchNameJson.Tw
				case "en":
					item.Address = AddressJson.En
					item.CabinetName = CabinetNameJson.En
					item.MchBranchName = MchBranchNameJson.En
				case "ja":
					item.Address = AddressJson.Ja
					item.CabinetName = CabinetNameJson.Ja
					item.MchBranchName = MchBranchNameJson.Ja
				case "ko":
					item.Address = AddressJson.Ko
					item.CabinetName = CabinetNameJson.Ko
					item.MchBranchName = MchBranchNameJson.Ko
				}
			} else {
				item.Address = v.Address
				item.CabinetName = v.CabinetName
				item.MchBranchName = v.MchBranchName
			}

			if v.OrderStatus == "GRACE" && !g.IsEmpty(v.GraceEndTime) {
				item.GraceCountdown = int(v.GraceEndTime.Sub(gtime.Now()).Seconds())
			} else {
				item.GraceCountdown = 0
			}
			if v.OrderStatus == "DONE" {
				item.OvertimeAmount = float64(v.OvertimeFee)
				item.UsedTime = int(v.FinishTime.Sub(v.StartTime).Seconds())
			} else if v.OrderStatus == "GRACE" {
				item.OvertimeAmount = float64(v.OvertimeFee)
				item.UsedTime = int(now.Sub(v.StartTime).Seconds())
			} else if v.OrderStatus == "WAIT_PAY" {
				item.OvertimeAmount = 0
				item.UsedTime = 0
			} else if v.OrderStatus == "CANCEL" {
				if !g.IsEmpty(v.EndTime) {
					if v.CancelTime.Before(v.EndTime) {
						item.OvertimeAmount = 0
					} else {
						diffHours := v.CancelTime.Sub(v.EndTime).Hours() // float64 小时数
						hours := math.Ceil(diffHours)                    // 向上取整，不足1小时按1小时

						item.OvertimeAmount = hours * float64(v.BoxTypePrice)
					}
					item.UsedTime = int(v.CancelTime.Sub(v.StartTime).Seconds())
				} else {
					item.OvertimeAmount = 0
					item.UsedTime = 0
				}
			} else {
				if now.Before(v.EndTime) {
					item.OvertimeAmount = 0
				} else {
					diffHours := now.Sub(v.EndTime).Hours() // float64 小时数
					hours := math.Ceil(diffHours)           // 向上取整，不足1小时按1小时

					item.OvertimeAmount = hours * float64(v.BoxTypePrice)
				}
				item.UsedTime = int(now.Sub(v.StartTime).Seconds())
			}

			list = append(list, item)
		}
	}
	return
}

func (s *sCabinetService) AppView(ctx context.Context, in *input_cabinet.OrderAppViewInp) (res *input_cabinet.OrderAppViewModel, err error) {
	var (
		CabinetOrder      *entity.CabinetOrder
		BoxTypeNameJson   *input_cabinet.LanguageJson
		AddressJson       *input_cabinet.LanguageJson
		CabinetNameJson   *input_cabinet.LanguageJson
		MchBranchNameJson *input_cabinet.LanguageJson
	)

	if err = s.Model(ctx).WithAll().Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&CabinetOrder); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_order_info_failed"))
		return
	}

	if g.IsEmpty(CabinetOrder) {
		err = gerror.New(gi18n.T(ctx, "get_order_info_failed"))
		return
	}

	res = new(input_cabinet.OrderAppViewModel)
	res.Id = int(CabinetOrder.Id)
	res.OrderSn = CabinetOrder.OrderSn
	res.Pin = CabinetOrder.Pin
	res.BoxNo = CabinetOrder.BoxNo
	res.BoxAlias = CabinetOrder.BoxAlias
	res.StartTime = CabinetOrder.StartTime
	res.BuyHours = CabinetOrder.BuyHours
	res.BaseAmount = CabinetOrder.BaseAmount
	res.BoxTypePrice = CabinetOrder.BoxTypePrice
	res.GraceEndTime = CabinetOrder.GraceEndTime
	res.GraceSeconds = CabinetOrder.GraceSeconds
	res.OrderStatus = CabinetOrder.OrderStatus
	res.IsFx = CabinetOrder.IsFx

	now := gtime.Now()
	var Language = contexts.GetLanguage(ctx)

	if !g.IsEmpty(CabinetOrder.BoxTypeNameJson) && !g.IsEmpty(CabinetOrder.AddressJson) && !g.IsEmpty(CabinetOrder.CabinetNameJson) && !g.IsEmpty(CabinetOrder.MchBranchNameJson) {
		BoxTypeNameJson = new(input_cabinet.LanguageJson)
		AddressJson = new(input_cabinet.LanguageJson)
		CabinetNameJson = new(input_cabinet.LanguageJson)
		MchBranchNameJson = new(input_cabinet.LanguageJson)
		if err = json.Unmarshal([]byte(CabinetOrder.BoxTypeNameJson.String()), &BoxTypeNameJson); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
		if err = json.Unmarshal([]byte(CabinetOrder.AddressJson.String()), &AddressJson); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
		if err = json.Unmarshal([]byte(CabinetOrder.CabinetNameJson.String()), &CabinetNameJson); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
		if err = json.Unmarshal([]byte(CabinetOrder.MchBranchNameJson.String()), &MchBranchNameJson); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}

		switch Language {
		case "zh":
			res.BoxTypeName = BoxTypeNameJson.Zh
			res.Address = AddressJson.Zh
			res.CabinetName = CabinetNameJson.Zh
			res.MchBranchName = MchBranchNameJson.Zh
		case "zh_CN":
			res.BoxTypeName = BoxTypeNameJson.Tw
			res.Address = AddressJson.Tw
			res.CabinetName = CabinetNameJson.Tw
			res.MchBranchName = MchBranchNameJson.Tw
		case "en":
			res.BoxTypeName = BoxTypeNameJson.En
			res.Address = AddressJson.En
			res.CabinetName = CabinetNameJson.En
			res.MchBranchName = MchBranchNameJson.En
		case "ja":
			res.BoxTypeName = BoxTypeNameJson.Ja
			res.Address = AddressJson.Ja
			res.CabinetName = CabinetNameJson.Ja
			res.MchBranchName = MchBranchNameJson.Ja
		case "ko":
			res.BoxTypeName = BoxTypeNameJson.Ko
			res.Address = AddressJson.Ko
			res.CabinetName = CabinetNameJson.Ko
			res.MchBranchName = MchBranchNameJson.Ko
		}
	} else {
		res.BoxTypeName = CabinetOrder.BoxTypeName
		res.Address = CabinetOrder.Address
		res.CabinetName = CabinetOrder.CabinetName
		res.MchBranchName = CabinetOrder.MchBranchName
	}

	if CabinetOrder.OrderStatus == "GRACE" && !g.IsEmpty(CabinetOrder.GraceEndTime) {
		res.GraceCountdown = int(CabinetOrder.GraceEndTime.Sub(gtime.Now()).Seconds())
	} else {
		res.GraceCountdown = 0
	}

	if res.OrderStatus == "DONE" {
		res.OvertimeAmount = float64(CabinetOrder.OvertimeFee)
		res.UsedTime = int(CabinetOrder.FinishTime.Sub(CabinetOrder.StartTime).Seconds())
		res.OvertimeTime = CabinetOrder.OvertimeSecs
	} else if res.OrderStatus == "GRACE" {
		res.OvertimeAmount = float64(CabinetOrder.OvertimeFee)
		res.UsedTime = int(now.Sub(CabinetOrder.StartTime).Seconds())
		res.OvertimeTime = CabinetOrder.OvertimeSecs
	} else if res.OrderStatus == "WAIT_PAY" {
		res.OvertimeAmount = 0
		res.UsedTime = 0
		res.OvertimeTime = 0
	} else if res.OrderStatus == "CANCEL" {
		if CabinetOrder.CancelTime.Before(CabinetOrder.EndTime) {
			res.OvertimeAmount = 0
			res.OvertimeTime = 0
		} else {
			diffHours := CabinetOrder.CancelTime.Sub(CabinetOrder.EndTime).Hours() // float64 小时数
			hours := math.Ceil(diffHours)                                          // 向上取整，不足1小时按1小时

			res.OvertimeAmount = hours * float64(CabinetOrder.BoxTypePrice)
			res.OvertimeTime = int(CabinetOrder.CancelTime.Sub(CabinetOrder.EndTime).Seconds())
		}
		res.UsedTime = int(CabinetOrder.CancelTime.Sub(CabinetOrder.StartTime).Seconds())
	} else {
		if now.Before(CabinetOrder.EndTime) {
			res.OvertimeAmount = 0
			res.OvertimeTime = 0
		} else {
			diffHours := now.Sub(CabinetOrder.EndTime).Hours() // float64 小时数
			hours := math.Ceil(diffHours)                      // 向上取整，不足1小时按1小时

			res.OvertimeAmount = hours * float64(CabinetOrder.BoxTypePrice)
			res.OvertimeTime = int(now.Sub(CabinetOrder.EndTime).Seconds())
		}
		res.UsedTime = int(now.Sub(CabinetOrder.StartTime).Seconds())
	}

	var openLog []*input_cabinet.OrderOpenLogModel

	// 请求订单详情接口
	// 请求mch那边下单接口，接口返回成功则继续，失败则回滚
	var (
		cabinetRequest   *cabinetApi.CabinetOrderQueryParams
		cabinetResponse  *cabinetApi.CabinetOrderQueryResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	cabinetRequest = new(cabinetApi.CabinetOrderQueryParams)
	cabinetRequest.OutTradeNo = CabinetOrder.OrderSn
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).OrderQuery(ctx, cabinetRequest); err != nil {
		return
	}
	if cabinetResponse.Code == 0 && !g.IsEmpty(cabinetResponse.Data.OpenLog) {
		if err = json.Unmarshal([]byte(gvar.New(cabinetResponse.Data.OpenLog).String()), &openLog); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
	}
	res.OpenLog = openLog
	res.MchBranchLat = CabinetOrder.MchBranchLat
	res.MchBranchLgt = CabinetOrder.MchBranchLgt

	return
}

// PayOvertimeInfo 支付超时费页面信息
func (s *sCabinetService) PayOvertimeInfo(ctx context.Context, in *input_cabinet.PayOvertimeInfoInp) (out *input_cabinet.PayOvertimeInfoModel, err error) {
	var (
		CabinetOrder    *entity.CabinetOrder
		IsBalance       bool
		AddressJson     *input_cabinet.LanguageJson
		CabinetNameJson *input_cabinet.LanguageJson
	)
	if err = s.Model(ctx).WithAll().Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&CabinetOrder); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_order_info_failed"))
		return
	}

	if CabinetOrder.OrderStatus != "OVERTIME" {
		err = gerror.New(gi18n.T(ctx, "order_status_incorrect"))
		return
	}

	// 初始化响应变量
	out = new(input_cabinet.PayOvertimeInfoModel)
	out.OrderSn = CabinetOrder.OrderSn
	out.MinHours = CabinetOrder.MinHours
	out.BoxNo = CabinetOrder.BoxNo
	out.BoxAlias = CabinetOrder.BoxAlias
	out.StartTime = CabinetOrder.StartTime
	out.EndTime = CabinetOrder.EndTime
	out.BoxTypePrice = CabinetOrder.BoxTypePrice
	out.MchBranchLat = CabinetOrder.MchBranchLat
	out.MchBranchLgt = CabinetOrder.MchBranchLgt

	var Language = contexts.GetLanguage(ctx)
	if !g.IsEmpty(CabinetOrder.AddressJson) && !g.IsEmpty(CabinetOrder.CabinetNameJson) {
		AddressJson = new(input_cabinet.LanguageJson)
		CabinetNameJson = new(input_cabinet.LanguageJson)
		if err = json.Unmarshal([]byte(CabinetOrder.AddressJson.String()), &AddressJson); err != nil {
			err = gerror.New("解析路径失败")
			return
		}
		if err = json.Unmarshal([]byte(CabinetOrder.CabinetNameJson.String()), &CabinetNameJson); err != nil {
			err = gerror.New("解析路径失败")
			return
		}

		switch Language {
		case "zh":
			out.Address = AddressJson.Zh
			out.CabinetName = CabinetNameJson.Zh
		case "zh_CN":
			out.Address = AddressJson.Tw
			out.CabinetName = CabinetNameJson.Tw
		case "en":
			out.Address = AddressJson.En
			out.CabinetName = CabinetNameJson.En
		case "ja":
			out.Address = AddressJson.Ja
			out.CabinetName = CabinetNameJson.Ja
		case "ko":
			out.Address = AddressJson.Ko
			out.CabinetName = CabinetNameJson.Ko
		}
	} else {
		out.Address = CabinetOrder.Address
		out.CabinetName = CabinetOrder.CabinetName
	}

	if err = json.Unmarshal([]byte(gvar.New(CabinetOrder.BoxTypeJson).String()), &out.BoxTypeInfo); err != nil {
		// 解析路径失败
		err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
		return
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		now := gtime.Now()
		diffHours := now.Sub(CabinetOrder.EndTime).Hours() // float64 小时数
		hours := math.Ceil(diffHours)                      // 向上取整，不足1小时按1小时

		if hours < 0 {
			hours = 0
		}
		OvertimeAmount := hours * float64(CabinetOrder.BoxTypePrice)

		if diffHours < 0 {
			diffHours = 0
		}

		// 将超时的信息写入订单
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
			WherePri(CabinetOrder.Id).Data(g.MapStrAny{
			dao.CabinetOrder.Columns().OvertimeFee:   OvertimeAmount,
			dao.CabinetOrder.Columns().OvertimeHours: hours,
			dao.CabinetOrder.Columns().OvertimeSecs:  int(diffHours * 3600),
			dao.CabinetOrder.Columns().PayStep:       "OVERTIME",
		}).Update(); err != nil {
			return
		}

		out.UsedTime = int(now.Sub(CabinetOrder.StartTime).Seconds())
		out.OvertimeTime = int(now.Sub(CabinetOrder.EndTime).Seconds())

		AllAmount, _ := decimal.NewFromFloat(CabinetOrder.OrderAmount).Add(decimal.NewFromFloat(float64(OvertimeAmount))).Round(2).Float64()
		// 读取第三方支付信息
		var (
			MemberInfo  *entity.PmsMember
			MemberLevel *entity.PmsMemberLevel
			MemberScene *entity.PmsMemberScene
			YYConfig    *model.YYConfig
			AllScore    float64
		)
		// 查询当前用户积分余额
		if MemberInfo, err = service.AppMember().MemberInfo(ctx, int(CabinetOrder.MemberId)); err != nil {
			return
		}
		if g.IsEmpty(MemberInfo) {
			err = gerror.New(gi18n.T(ctx, "user_info_does_not_exist"))
			return
		}
		// 查询会员等级
		if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(MemberLevel) {
			err = gerror.New(gi18n.T(ctx, "user_level_does_not_exist"))
			return
		}
		if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 5).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(MemberScene) {
			err = gerror.New(gi18n.T(ctx, "scene_unknown"))
			return
		}
		// 获取积分汇率
		if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
			return
		}
		if YYConfig.ExchangeRate == 0 || YYConfig.ExchangeRate < 0 {
			YYConfig.ExchangeRate = 1
		}

		out.PayInfo.MemberBalance = MemberInfo.Balance
		out.PayInfo.OvertimeAmount = OvertimeAmount

		// 计算该笔订单能够有多少积分进行抵扣
		// * 计算可以使用的积分总额   =  订单总金额 * 场景抵扣基础比例 * 积分汇率
		AllScore = decimal.NewFromFloat(OvertimeAmount).Mul(decimal.NewFromFloat(MemberScene.PayRate)).Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).Div(decimal.NewFromInt(100)).Round(0).InexactFloat64()
		// * 计算会员可用积分总额 = min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )
		AllScore = math.Min(AllScore, MemberInfo.Balance)
		// * 计算积分的现金价值  min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )  /  积分汇率
		AllScore = decimal.NewFromFloat(AllScore).Div(decimal.NewFromFloat(YYConfig.ExchangeRate)).Round(0).InexactFloat64()

		if MemberScene.IsPayOpen == "N" || AllAmount < MemberScene.LimitMoney {
			AllScore = 0
		}

		if in.IsBalance == 1 {
			IsBalance = true
		}

		if OvertimeAmount == 0 {
			out.PayInfo.PayModel = 1
			out.PayInfo.ThirdPay.ThirdAmount = AllScore
		} else if !IsBalance || MemberScene.IsPayOpen == "N" || AllAmount < MemberScene.LimitMoney || AllScore <= 0 {
			// 三方支付
			out.PayInfo.PayModel = 3
			out.PayInfo.ThirdPay.ThirdAmount = OvertimeAmount
		} else if OvertimeAmount > AllScore {
			// 组合支付
			out.PayInfo.PayModel = 2
			out.PayInfo.Balance.BalanceAmount = AllScore
			out.PayInfo.ThirdPay.ThirdAmount = decimal.NewFromFloat(OvertimeAmount).Sub(decimal.NewFromFloat(AllScore)).Round(0).InexactFloat64()
		} else {
			// 纯余额支付
			out.PayInfo.PayModel = 1
			out.PayInfo.Balance.BalanceAmount = AllScore
		}

		out.PayInfo.Score = AllScore
		out.PayInfo.Balance.BalanceConfig = new(input_cabinet.BalanceConfig)
		out.PayInfo.Balance.BalanceConfig.ScenePayRate = MemberScene.PayRate
		out.PayInfo.Balance.BalanceConfig.ExchangeRate = YYConfig.ExchangeRate
		out.PayInfo.Balance.BalanceConfig.LevelName = MemberLevel.LevelName
		out.PayInfo.Balance.BalanceConfig.Level = MemberLevel.Id

		return
	}); err != nil {
		return
	}
	return
}

// PayOvertime 支付超时费
func (s *sCabinetService) PayOvertime(ctx context.Context, in *input_cabinet.PayOvertimeInp) (out *input_cabinet.PayOvertimeModel, err error) {
	var (
		CabinetOrder   *entity.CabinetOrder
		PayConfig      *model.PayConfig
		ExpirationTime int
	)
	if err = s.Model(ctx).WithAll().Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&CabinetOrder); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_order_info_failed"))
		return
	}

	if CabinetOrder.OrderStatus != "OVERTIME" {
		err = gerror.New(gi18n.T(ctx, "order_status_incorrect"))
		return
	}

	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}

	// 初始化响应变量
	out = new(input_cabinet.PayOvertimeModel)
	out.OrderSn = CabinetOrder.OrderSn
	out.CreateOrderTime = gtime.Now().Format("Y-m-d H:i:s")
	out.Countdown = int(PayConfig.HotelStayExp)

	// 读取第三方支付信息
	var (
		MemberInfo  *entity.PmsMember
		MemberLevel *entity.PmsMemberLevel
		MemberScene *entity.PmsMemberScene
		YYConfig    *model.YYConfig
		AllScore    float64
		IsBalance   bool
	)
	// 查询当前用户积分余额
	if MemberInfo, err = service.AppMember().MemberInfo(ctx, int(CabinetOrder.MemberId)); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		err = gerror.New(gi18n.T(ctx, "user_info_does_not_exist"))
		return
	}
	// 查询会员等级
	if err = dao.PmsMemberLevel.Ctx(ctx).Where(dao.PmsMemberLevel.Columns().Id, MemberInfo.Level).Scan(&MemberLevel); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberLevel) {
		err = gerror.New(gi18n.T(ctx, "user_level_does_not_exist"))
		return
	}
	if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 5).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberScene) {
		err = gerror.New(gi18n.T(ctx, "scene_unknown"))
		return
	}
	// 获取积分汇率
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}
	if YYConfig.ExchangeRate == 0 || YYConfig.ExchangeRate < 0 {
		YYConfig.ExchangeRate = 1
	}
	AllAmount, _ := decimal.NewFromFloat(CabinetOrder.OrderAmount).Add(decimal.NewFromFloat(float64(CabinetOrder.OvertimeFee))).Round(2).Float64()
	// 计算该笔订单能够有多少积分进行抵扣
	// * 计算可以使用的积分总额   =  订单总金额 * 场景抵扣基础比例 * 积分汇率
	AllScore = decimal.NewFromFloat(float64(CabinetOrder.OvertimeFee)).Mul(decimal.NewFromFloat(MemberScene.PayRate)).Mul(decimal.NewFromFloat(YYConfig.ExchangeRate)).Div(decimal.NewFromInt(100)).Round(0).InexactFloat64()
	// * 计算会员可用积分总额 = min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )
	AllScore = math.Min(AllScore, MemberInfo.Balance)
	// * 计算积分的现金价值  min(会员的总积分 , 订单总金额 * 场景抵扣基础比例 * 积分汇率 )  /  积分汇率
	AllScore = decimal.NewFromFloat(AllScore).Div(decimal.NewFromFloat(YYConfig.ExchangeRate)).Round(0).InexactFloat64()

	if MemberScene.IsPayOpen == "N" || AllAmount < MemberScene.LimitMoney {
		AllScore = 0
	}

	if in.IsBalance == 1 {
		IsBalance = true
	}

	if CabinetOrder.OvertimeFee == 0 {
		out.PayModel = 1
		out.ThirdPay.ThirdAmount = AllScore
	} else if !IsBalance || MemberScene.IsPayOpen == "N" || AllAmount < MemberScene.LimitMoney || AllScore <= 0 {
		// 三方支付
		out.PayModel = 3
		out.ThirdPay.ThirdAmount = float64(CabinetOrder.OvertimeFee)
	} else if float64(CabinetOrder.OvertimeFee) > AllScore {
		// 组合支付
		out.PayModel = 2
		out.Balance.BalanceAmount = AllScore
		out.ThirdPay.ThirdAmount = decimal.NewFromFloat(float64(CabinetOrder.OvertimeFee)).Sub(decimal.NewFromFloat(AllScore)).Round(0).InexactFloat64()
	} else {
		// 纯余额支付
		out.PayModel = 1
		out.Balance.BalanceAmount = AllScore
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 保存支付方式
		if _, err = dao.CabinetOrder.Ctx(ctx).TX(tx).
			WherePri(CabinetOrder.Id).Data(g.MapStrAny{
			dao.CabinetOrder.Columns().OvertimeBalAmount: out.Balance.BalanceAmount,
			dao.CabinetOrder.Columns().OvertimePayModel:  out.PayModel,
		}).Update(); err != nil {
			return
		}

		// 删除transaction表中未支付的记录，以防重复生成多笔待支付
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, CabinetOrder.OrderSn).Where(dao.PmsTransaction.Columns().OrderType, "CABINET").Where(dao.PmsTransaction.Columns().PayStatus, "WAIT").Delete(); err != nil {
			return
		}

		out.Balance.BalancePayOrderSn = uuid.CreateOrderCode("BL")
		ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()
		// 插入数据库支付信息内容
		if out.Balance.BalanceAmount > 0 {
			// 插入余额支付信息
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:       CabinetOrder.OrderSn,
				OrderType:     "CABINET",
				Scene:         "CABINET",
				TransactionSn: out.Balance.BalancePayOrderSn,
				PayType:       "BAL",
				Amount:        out.Balance.BalanceAmount,
				PayStatus:     "WAIT",
				ExpiredTime:   gtime.New(ExpirationTime),
			}); err != nil {
				return err
			}
		}

		if out.ThirdPay.ThirdAmount > 0 {
			// 插入三方支付信息
			if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsTransaction{
				OrderSn:          CabinetOrder.OrderSn,
				OrderType:        "CABINET",
				Scene:            "CABINET",
				TransactionSn:    out.ThirdPay.ThirdPayOrderSn,
				PaymentRequestId: "",
				PayChannel:       "paycloud",
				PayType:          "",
				PriceCurrency:    "JPY",
				Amount:           out.ThirdPay.ThirdAmount,
				PayStatus:        "WAIT",
				ExpiredTime:      gtime.New(ExpirationTime),
			}); err != nil {
				return err
			}
		}
		return
	}); err != nil {
		return
	}

	return
}

func (s *sCabinetService) List(ctx context.Context, in *input_cabinet.OrderListInp) (list []*input_cabinet.OrderListModel, totalCount int, err error) {
	mod := dao.CabinetOrder.Ctx(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.CabinetOrder.Table(), input_cabinet.OrderListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_cabinet.OrderListModel{}, &dao.PmsMember, "pmsMember"))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))

	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.CabinetOrder.Table(), dao.CabinetOrder.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.CabinetOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.MchBranchName) {
		mod = mod.WhereLike(dao.CabinetOrder.Columns().MchBranchName, "%"+in.MchBranchName+"%")
	}

	if !g.IsEmpty(in.CabinetName) {
		mod = mod.WhereLike(dao.CabinetOrder.Columns().CabinetName, "%"+in.CabinetName+"%")
	}

	if !g.IsEmpty(in.OrderStatus) {
		mod = mod.Where(dao.CabinetOrder.Columns().OrderStatus, in.OrderStatus)
	}

	if !g.IsEmpty(in.BoxNo) {
		mod = mod.WhereLike(dao.CabinetOrder.Columns().BoxNo, "%"+in.BoxNo+"%")
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.CabinetOrder.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.MemberSearch) {
		mod = mod.Where(mod.Builder().
			WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Id, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Phone, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Mail, "%"+in.MemberSearch+"%"))
	}

	if !g.IsEmpty(in.OrderStatus) && in.OrderStatus != "ALL" && in.OrderStatus != "ABNORMAL" {
		mod = mod.Where(dao.CabinetOrder.Columns().OrderStatus, in.OrderStatus)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.CabinetOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	if !g.IsEmpty(in.CreateSort) {
		if in.CreateSort == "ascend" {
			mod = mod.Order(dao.CabinetOrder.Columns().CreatedAt, "asc")
		} else if in.CreateSort == "descend" {
			mod = mod.Order(dao.CabinetOrder.Columns().CreatedAt, "desc")
		} else {
			mod = mod.OrderDesc(dao.CabinetOrder.Columns().Id)
		}
	} else {
		mod = mod.OrderDesc(dao.CabinetOrder.Columns().Id)
	}

	if in.Pagination {
		// 分页
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
			return
		}
	} else {
		// 不分页
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
			return
		}
	}
	for _, v := range list {
		// 计算格口使用时长
		// 如果订单状态为ING、OVERTIME、GRACE，则计算时长为当前时间减去订单开始时间
		if v.OrderStatus == "ING" || v.OrderStatus == "OVERTIME" || v.OrderStatus == "GRACE" {
			v.UsedTime = int(gtime.Now().Sub(v.StartTime).Seconds())
		} else if v.OrderStatus == "WAIT_PAY" || v.OrderStatus == "CANCEL" || v.OrderStatus == "HAVE_PAID" {
			v.UsedTime = 0
		} else if v.OrderStatus == "DONE" {
			v.UsedTime = int(v.FinishTime.Sub(v.StartTime).Seconds())
		}
	}

	return
}

func (s *sCabinetService) OrderView(ctx context.Context, in *input_cabinet.OrderViewInp) (res *input_cabinet.OrderViewModel, err error) {
	if err = dao.CabinetOrder.Ctx(ctx).Unscoped().WithAll().Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取预订单信息，请稍后重试！")
		return
	}

	// Check member deleted
	if res.MemberId > 0 {
		memberDeletedAt, _ := dao.PmsMember.Ctx(ctx).Unscoped().Fields(dao.PmsMember.Columns().DeletedAt).
			Where(dao.PmsMember.Columns().Id, res.MemberId).Value()
		if !memberDeletedAt.IsNil() && !memberDeletedAt.IsEmpty() {
			res.MemberDeleted = true
			// Manually load deleted member info
			var memberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, res.MemberId).Scan(&memberInfo); err == nil && memberInfo != nil {
				res.MemberInfo = &struct {
					gmeta.Meta `orm:"table:hg_pms_member"`
					Id         int    `json:"id"             dc:""`
					MemberNo   string `json:"memberNo"       dc:"会员号"`
					FullName   string `json:"fullName"       dc:"会员姓名"`
					Phone      string `json:"phone"          dc:"手机号"`
					PhoneArea  string `json:"phoneArea"      dc:"区号"`
					Mail       string `json:"mail"           dc:"邮箱"`
				}{
					Id:        memberInfo.Id,
					MemberNo:  memberInfo.MemberNo,
					FullName:  memberInfo.FullName,
					Phone:     memberInfo.Phone,
					PhoneArea: memberInfo.PhoneArea,
					Mail:      memberInfo.Mail,
				}
			}
		}
	}

	if res.OrderStatus == "ING" || res.OrderStatus == "OVERTIME" || res.OrderStatus == "GRACE" {
		res.UsedTime = int(gtime.Now().Sub(res.StartTime).Seconds())
	} else if res.OrderStatus == "WAIT_PAY" || res.OrderStatus == "CANCEL" || res.OrderStatus == "HAVE_PAID" {
		res.UsedTime = 0
	} else if res.OrderStatus == "DONE" {
		res.UsedTime = int(res.FinishTime.Sub(res.StartTime).Seconds())
	}

	for _, v := range res.LogList {
		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			var AdminMemberInfo *entity.AdminMember
			if err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
				return
			}
			if AdminMemberInfo != nil {
				v.OperateName = AdminMemberInfo.Username
			}
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			if PmsMemberInfo != nil {
				v.OperateName = PmsMemberInfo.FullName
			}
		}
	}

	var openLog []*input_cabinet.OrderOpenLogModel

	// 请求订单详情接口
	// 请求mch那边下单接口，接口返回成功则继续，失败则回滚
	var (
		cabinetRequest   *cabinetApi.CabinetOrderQueryParams
		cabinetResponse  *cabinetApi.CabinetOrderQueryResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	cabinetRequest = new(cabinetApi.CabinetOrderQueryParams)
	cabinetRequest.OutTradeNo = res.OrderSn
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).OrderQuery(ctx, cabinetRequest); err != nil {
		return
	}
	if cabinetResponse.Code == 0 && !g.IsEmpty(cabinetResponse.Data.OpenLog) {
		if err = json.Unmarshal([]byte(gvar.New(cabinetResponse.Data.OpenLog).String()), &openLog); err != nil {
			err = gerror.New(gi18n.T(ctx, "failed_to_parse_path"))
			return
		}
	}
	res.OpenLogList = openLog

	return
}

func (s *sCabinetService) ExportOrder(ctx context.Context, in *input_cabinet.OrderExportInp) (err error) {
	var (
		lastInsertId int64
	)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		resultJson, _ := json.Marshal(in)

		if lastInsertId, err = dao.OrderExport.Ctx(ctx).
			Data(entity.OrderExport{
				Scene:     5,
				Condition: string(resultJson),
			}).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		return
	})

	if err != nil {
		return
	}

	// 导出(队列)
	_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderExport,
		DataByte:     gvar.New(lastInsertId).Bytes(),
		Header:       nil,
	})

	return
}

func (s *sCabinetService) StartExport(ctx context.Context, in *input_cabinet.OrderExportInp) (path string, err error) {
	var (
		ChangeList      []*input_cabinet.OrderExportModel
		OrderListIn     *input_cabinet.OrderListInp
		OrderStatusList input_basics.DataSelectModel
		PayStatusList   input_basics.DataSelectModel
	)

	OrderListIn = &input_cabinet.OrderListInp{
		PageReq: input_form.PageReq{
			Pagination: false,
		},
		OrderSn:       in.OrderSn,
		MchBranchName: in.MchBranchName,
		OrderStatus:   in.OrderStatus,
		MemberSearch:  in.MemberSearch,
		BoxNo:         in.BoxNo,
		CreatedAt:     in.CreatedAt,
	}

	list, _, err := s.List(ctx, OrderListIn)
	if err != nil {
		return
	}

	if OrderStatusList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "cabinet_order_status"}); err != nil {
		return
	}
	if PayStatusList, err = service.BasicsDictData().Select(ctx, &input_basics.DataSelectInp{Type: "spa_order_pay_status"}); err != nil {
		return
	}

	for _, item := range list {

		for _, v2 := range PayStatusList {
			if item.PayStatus == v2.Value {
				item.PayStatus = v2.Label
			}
			if item.OvertimePayStatus == v2.Value {
				item.OvertimePayStatus = v2.Label
			}
		}

		for _, v2 := range OrderStatusList {
			if item.OrderStatus == v2.Value {
				item.OrderStatus = v2.Label
			}
		}

		itemData := &input_cabinet.OrderExportModel{
			OrderSn:           item.OrderSn,
			OutOrderSn:        item.OutOrderSn,
			MemberNo:          item.MemberInfo.MemberNo,
			MemberFullName:    item.MemberInfo.FullName,
			CabinetName:       item.CabinetName,
			CityName:          item.CityName,
			MchBranchName:     item.MchBranchName,
			Address:           item.Address,
			BoxTypeName:       item.BoxTypeName,
			BoxTypePrice:      item.BoxTypePrice,
			BoxNo:             item.BoxNo,
			BoxAlias:          item.BoxAlias,
			Pin:               item.Pin,
			OrderAmount:       item.OrderAmount,
			BaseAmount:        item.BaseAmount,
			PayTime:           item.PayTime,
			PayStatus:         item.PayStatus,
			OvertimeSecs:      item.OvertimeSecs,
			OvertimeHours:     item.OvertimeHours,
			OvertimeFee:       item.OvertimeFee,
			OvertimePayTime:   item.OvertimePayTime,
			OvertimePayStatus: item.OvertimePayStatus,
			OrderStatus:       item.OrderStatus,
			OrderTime:         item.CreatedAt.Format("Y-m-d H:i:s"),
			StartTime:         item.StartTime,
			EndTime:           item.EndTime,
			FinishTime:        item.FinishTime,
			RefundTotalAmount: item.RefundAmount,
			RefundTime:        item.RefundTime.Format("Y-m-d H:i:s"),
		}

		for _, TransactionDetail := range item.TransactionDetail {
			if TransactionDetail.PayStatus == "DONE" {
				switch TransactionDetail.PayType {
				case "BAL":
					itemData.PointsPayment += TransactionDetail.PayAmount
					break
				case "COUPON":
					itemData.CouponPayment += TransactionDetail.PayAmount
					break
				case "WeChatPay":
					itemData.PaycloudWechatPay += TransactionDetail.PayAmount
					break
				case "Alipay+":
					itemData.PaycloudAlipayPay += TransactionDetail.PayAmount
					break
				case "Paypal":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "PaypalCard":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "StripeCard":
					itemData.StripeCreditPay += TransactionDetail.PayAmount
					break
				case "WeChatMiniPay":
					itemData.MlilifeWeChatMiniPay += TransactionDetail.PayAmount
					break
				}
			}
		}

		for _, TransactionRefundDetail := range item.TransactionRefundDetail {
			switch TransactionRefundDetail.RefundType {
			case "BAL":
				itemData.PointsRefund += TransactionRefundDetail.RefundAmount
				break
			case "COUPON":
				itemData.CouponRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatPay":
				itemData.PaycloudWechatRefund += TransactionRefundDetail.RefundAmount
				break
			case "Alipay+":
				itemData.PaycloudAlipayRefund += TransactionRefundDetail.RefundAmount
				break
			case "Paypal":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "PaypalCard":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "StripeCard":
				itemData.StripeCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatMiniPay":
				itemData.MlilifeWeChatMiniRefund += TransactionRefundDetail.RefundAmount
				break
			}
		}
		ChangeList = append(ChangeList, itemData)
	}
	tags, err := convert.GetEntityDescTags(input_cabinet.OrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出储物柜订单-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("储物柜订单")
		exports   []input_cabinet.OrderExportModel
	)

	if err = gconv.Scan(ChangeList, &exports); err != nil {
		return
	}

	path, err = excel.ExportByStructsFile(ctx, tags, exports, fileName, sheetName)

	return
}

func (s *sCabinetService) ExportList(ctx context.Context, in *input_cabinet.OrderExportListInp) (list []*input_cabinet.OrderExportListModel, totalCount int, err error) {
	mod := dao.OrderExport.Ctx(ctx)

	mod = mod.Fields(input_cabinet.OrderExportListModel{})

	mod = mod.Where(dao.OrderExport.Columns().Scene, 5)

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.OrderExport.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		v.Path = g.Cfg().MustGet(ctx, "localUploadDomain").String() + "/" + v.Path
	}

	return
}

func (s *sCabinetService) CompleteOrder(ctx context.Context, in *input_cabinet.OrderCompleteInp) (err error) {

	var models *entity.CabinetOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	if models.OrderStatus != "ING" && models.OrderStatus != "OVERTIME" && models.OrderStatus != "GRACE" {
		err = gerror.New("订单状态不正确")
		return
	}

	// 将操作人的信息写入订单
	if _, err = dao.CabinetOrder.Ctx(ctx).
		WherePri(models.Id).Data(g.MapStrAny{
		dao.CabinetOrder.Columns().IsAdminComplete:         1,
		dao.CabinetOrder.Columns().AdminCompleteOperatorId: int(contexts.GetUserId(ctx)),
	}).Update(); err != nil {
		return
	}

	// 请求mch订单完成
	var (
		cabinetRequest   *cabinetApi.OrderCompleteParams
		cabinetResponse  *cabinetApi.OrderCompleteResponse
		CabinetApiConfig *model.CabinetApiConfig
	)
	if CabinetApiConfig, err = service.BasicsConfig().GetCabinetApi(ctx); err != nil {
		return
	}
	cabinetRequest = new(cabinetApi.OrderCompleteParams)
	cabinetRequest.OutTradeNo = models.OrderSn
	if cabinetResponse, err = cabinetApi.NewClient(ctx, CabinetApiConfig).OrderComplete(ctx, cabinetRequest); err != nil {
	}

	if cabinetResponse.Code == 0 {

	} else {
		err = gerror.New(cabinetResponse.Msg)
		return
	}

	return

}

// Refund 订单退款
func (s *sCabinetService) Refund(ctx context.Context, in *input_cabinet.CabinetOrderRefundInp) (err error) {

	var models *entity.CabinetOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	if models.PayStatus != "HAVE_PAID" {
		err = gerror.New("订单支付状态不正确")
		return
	}

	if models.OrderStatus == "CANCEL" || models.OrderStatus == "WAIT_PAY" {
		err = gerror.New("订单状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			CancelFee         float64 // 手续费
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		// 手续费 = 订单金额 - 订单已退款金额 - 优惠券减免金额 - 退款金额
		CancelFee = models.OrderAmount - models.RefundAmount - models.CouponAmount - in.RefundMoney
		CancelFeeCalc := CancelFee

		// 支付流水
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).OrderDesc(`
			CASE pay_type
				WHEN 'StripeCard' THEN 1
				WHEN 'PaypalCard' THEN 2
				WHEN 'Paypal' THEN 3
				WHEN 'WeChatPay' THEN 4
				WHEN 'Alipay+' THEN 5
				WHEN 'WeChatMiniPay' THEN 6
				WHEN 'BAL' THEN 7
				ELSE 8
			END
		`).Scan(&Transaction); err != nil {
			return
		}

		// 退款流水
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

		// 遍历支付流水 2000 800 积分支付 500 三方支付 1500
		// 手续费 = 订单金额 - 优惠金额 - 退款金额 = 2000-0-800 = 1200
		for _, v := range Transaction {
			// 可退款金额
			Refundable := v.Amount - v.RefundAmount
			if v.PayType == "BAL" {
				if g.IsEmpty(CancelFeeCalc) {
					RefundBalance += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundBalance += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundBalance += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			} else if v.PayType == "COUPON" {

			} else {
				if g.IsEmpty(CancelFeeCalc) {
					RefundFee += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundFee += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundFee += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			}
		}

		// 退款金额 = 可退款积分 + 第三方支付
		RefundAmount := RefundBalance + RefundFee

		// 订单退款金额
		OrderTotalRefundAmount := RefundAmount + models.RefundAmount
		// 订单退款积分
		OrderTotalRefundBalance := RefundBalance + models.RefundBalAmount

		// 修改订单退款状态
		RefundStatus := "DONE"
		if models.OrderAmount > OrderTotalRefundAmount {
			RefundStatus = "PART"
		}
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(g.MapStrAny{
			dao.CabinetOrder.Columns().RefundAmount:       OrderTotalRefundAmount,
			dao.CabinetOrder.Columns().RefundBalAmount:    OrderTotalRefundBalance,
			dao.CabinetOrder.Columns().RefundCouponAmount: 0,
			dao.CabinetOrder.Columns().RefundStatus:       RefundStatus,
			dao.CabinetOrder.Columns().RefundTime:         gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if RefundAmount > 0 {
			// 退款
			err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
				OrderSn:      models.OrderSn,
				RefundAmount: RefundAmount,
				Remark:       in.AdminCancelReason,
				OperateType:  "ADMIN",
				OperateId:    int(contexts.GetUserId(ctx)),
			}, tx)
			if err != nil {
				err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.CabinetOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
			OrderId:     int(in.Id),
			OrderStatus: models.OrderStatus,
			ActionWay:   "ADMIN_REFUND",
			Remark:      fmt.Sprintf("后台退款，原因：%s", in.AdminCancelReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		return
	})

}
