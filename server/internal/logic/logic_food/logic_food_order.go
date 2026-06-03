package logic_food

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/library/toretaApi"
	"APT/internal/library/ws"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_refund"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/encrypt"
	"APT/utility/excel"
	"APT/utility/rabbitmq"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sFoodOrder struct{}

func NewFoodOrder() *sFoodOrder {
	return &sFoodOrder{}
}

func init() {
	service.RegisterFoodOrder(NewFoodOrder())
}

func (s *sFoodOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FoodOrder.Ctx(ctx), option...)
}

func (s *sFoodOrder) List(ctx context.Context, in *input_food.FoodOrderListInp) (list []*input_food.FoodOrderListModel, totalCount int, err error) {
	mod := dao.FoodOrder.Ctx(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.FoodOrder.Table(), input_food.FoodOrderListModel{})
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `goods_deleted`", dao.FoodGoods.Table(), dao.FoodGoods.Columns().DeletedAt))
	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.FoodOrder.Table(), dao.FoodOrder.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))
	mod = mod.LeftJoin(dao.FoodGoods.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.FoodOrder.Table(), dao.FoodOrder.Columns().GoodsId, dao.FoodGoods.Table(), dao.FoodGoods.Columns().Id))

	//mod = mod.WhereNot(dao.FoodOrder.Columns().PayStep, "ALLAMOUNT")

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.FoodOrder.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.OrderType) {
		mod = mod.Where(dao.FoodOrder.Columns().OrderType, in.OrderType)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.FoodOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.RestaurantName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.RestaurantName, "name")
		if err == nil {
			restaurantIds, _ := service.FoodRestaurant().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.FoodOrder.Columns().RestaurantId, restaurantIds)
		}
	}
	if !g.IsEmpty(in.OrderStatus) && in.OrderStatus != "ALL" {
		if in.OrderStatus == "DEPOSIT_WAIT_PAY" {
			// 定金待支付
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().OrderStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().RemainPayStatus, "WAIT_PAY")
		} else if in.OrderStatus == "DEPOSIT_CANCEL" {
			// 定金未支付取消(主动取消 \ 支付超时)
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "CANCEL").Where(dao.FoodOrder.Columns().OrderStatus, "CANCEL").Where(dao.FoodOrder.Columns().RemainPayStatus, "CANCEL")
		} else if in.OrderStatus == "WAIT_CONFIRM" {
			// 定金已支付后台待确认
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "HAVE_PAID").Where(dao.FoodOrder.Columns().OrderStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().RemainPayStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().BookingStatus, "WAIT_CONFIRM")
		} else if in.OrderStatus == "CONFIRM_DISAGREE" {
			// 后台确认失败
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "REFUND").Where(dao.FoodOrder.Columns().OrderStatus, "CANCEL").Where(dao.FoodOrder.Columns().RemainPayStatus, "CANCEL").WhereNot(dao.FoodOrder.Columns().ConfirmRefuseReason, "")
		} else if in.OrderStatus == "REMAIN_WAIT_PAY" {
			// 后台确认成功，尾款待支付
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "HAVE_PAID").Where(dao.FoodOrder.Columns().OrderStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().RemainPayStatus, "WAIT_PAY").Where(dao.FoodOrder.Columns().BookingStatus, "CONFIRMED")
		} else if in.OrderStatus == "REMAIN_CANCEL" {
			// 定金已支付后台待确认主动取消 / 后台确认成功，尾款待支付主动取消 / 尾款支付超时 / 后台操作退款
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "REFUND").Where(dao.FoodOrder.Columns().OrderStatus, "CANCEL").Where(dao.FoodOrder.Columns().RemainPayStatus, "CANCEL").Where(dao.FoodOrder.Columns().ConfirmRefuseReason, "")
		} else if in.OrderStatus == "WAIT_VERIFY" {
			// 定金已支付尾款已支付(待核销)
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().OrderStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().RemainPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().BookingStatus, "CONFIRMED").
				Where(dao.FoodOrder.Columns().VerifyStatus, "WAIT_VERIFY").
				WhereGT(dao.FoodOrder.Columns().BookDatetime, gtime.Now())
		} else if in.OrderStatus == "COMPLETE" {
			// 核销成功（已完成）
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().OrderStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().RemainPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().BookingStatus, "CONFIRMED").
				Where(dao.FoodOrder.Columns().VerifyStatus, "VERIFIED")
		} else if in.OrderStatus == "REFUND" {
			// 退款
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "REFUND").
				Where(dao.FoodOrder.Columns().OrderStatus, "CANCEL").
				Where(dao.FoodOrder.Columns().RemainPayStatus, "REFUND").
				Where(dao.FoodOrder.Columns().VerifyStatus, "WAIT_VERIFY")
		} else if in.OrderStatus == "OVERDUE" {
			// 逾期
			mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().OrderStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().RemainPayStatus, "HAVE_PAID").
				Where(dao.FoodOrder.Columns().BookingStatus, "CONFIRMED").
				Where(dao.FoodOrder.Columns().VerifyStatus, "WAIT_VERIFY").
				WhereLT(dao.FoodOrder.Columns().BookDatetime, gtime.Now())
		}
	}
	if !g.IsEmpty(in.BookingStatus) && in.BookingStatus != "ALL" {
		mod = mod.Where(dao.FoodOrder.Columns().BookingStatus, in.BookingStatus)
	}

	if !g.IsEmpty(in.DepositPayStatus) && in.DepositPayStatus != "ALL" {
		mod = mod.Where(dao.FoodOrder.Columns().DepositPayStatus, in.DepositPayStatus)
	}

	if !g.IsEmpty(in.RemainPayStatus) && in.RemainPayStatus != "ALL" {
		mod = mod.Where(dao.FoodOrder.Columns().RemainPayStatus, in.RemainPayStatus)
	}

	if !g.IsEmpty(in.ToretaPayStatus) && in.ToretaPayStatus != "ALL" {
		mod = mod.Where(dao.FoodOrder.Columns().OrderStatus, in.ToretaPayStatus)
	}

	if !g.IsEmpty(in.GoodsName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.GoodsName, "goods_name")
		if err == nil {
			goodsIds, _ := service.FoodGoods().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.FoodOrder.Columns().GoodsId, goodsIds)
		}
	}

	if len(in.BookDateTime) == 2 {
		mod = mod.WhereBetween(dao.FoodOrder.Columns().BookDatetime, in.BookDateTime[0], in.BookDateTime[1])
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FoodOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	defaultSort := 0
	if !g.IsEmpty(in.BookSort) {
		switch in.BookSort {
		case "ascend":
			mod = mod.Order(dao.FoodOrder.Columns().BookDatetime, "asc")
		case "descend":
			mod = mod.Order(dao.FoodOrder.Columns().BookDatetime, "desc")
		default:
			defaultSort = 1
		}
	}

	if !g.IsEmpty(in.CreateSort) {
		switch in.CreateSort {
		case "ascend":
			mod = mod.Order(dao.FoodOrder.Columns().CreatedAt, "asc")
		case "descend":
			mod = mod.Order(dao.FoodOrder.Columns().CreatedAt, "desc")
		default:
			defaultSort = 1
		}
	}

	if g.IsEmpty(in.BookSort) && g.IsEmpty(in.CreateSort) {
		defaultSort = 1
	}

	if defaultSort == 1 {
		mod = mod.OrderDesc(dao.FoodOrder.Columns().Id)
	}

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		if v.OrderType == "TORETA" || v.OrderType == "CRSALL" {
			if v.OrderStatus == "WAIT_PAY" {
				// 待支付
				v.ActualOrderStatus = "WAIT_PAY"

			} else if v.OrderStatus == "CANCEL" && g.IsEmpty(v.ConfirmRefuseReason) && v.RefundStatus == "WAIT" {
				// 取消（未支付取消）
				v.ActualOrderStatus = "CANCEL"

			} else if v.OrderStatus == "HAVE_PAID" && v.BookingStatus == "WAIT_CONFIRM" {
				// 后台待确认
				v.ActualOrderStatus = "WAIT_CONFIRM"

			} else if v.OrderStatus == "CANCEL" && !g.IsEmpty(v.ConfirmRefuseReason) {
				// 后台确认失败
				v.ActualOrderStatus = "CONFIRM_DISAGREE"

			} else if v.OrderStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "WAIT_VERIFY" && gtime.Now().Before(v.BookDatetime) {
				// 待核销
				v.ActualOrderStatus = "WAIT_VERIFY"

			} else if v.OrderStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "VERIFIED" {
				// 核销成功（已完成）
				v.ActualOrderStatus = "COMPLETE"

			} else if v.OrderStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "WAIT_VERIFY" && gtime.Now().After(v.BookDatetime) {
				// 逾期
				v.ActualOrderStatus = "OVERDUE"

			} else if v.OrderStatus == "CANCEL" && g.IsEmpty(v.ConfirmRefuseReason) && v.RefundStatus != "WAIT" {
				// 退款
				v.ActualOrderStatus = "REFUND"
			}
		} else {
			if v.DepositPayStatus == "WAIT_PAY" && v.OrderStatus == "WAIT_PAY" && v.RemainPayStatus == "WAIT_PAY" {
				//定金待支付
				v.ActualOrderStatus = "DEPOSIT_WAIT_PAY"
			} else if v.DepositPayStatus == "CANCEL" && v.OrderStatus == "CANCEL" && v.RemainPayStatus == "CANCEL" {
				//定金未支付取消(主动取消 | 支付超时)
				v.ActualOrderStatus = "DEPOSIT_CANCEL"
			} else if v.DepositPayStatus == "HAVE_PAID" && v.OrderStatus == "WAIT_PAY" && v.RemainPayStatus == "WAIT_PAY" && v.BookingStatus == "WAIT_CONFIRM" {
				//定金已支付后台待确认
				v.ActualOrderStatus = "WAIT_CONFIRM"
			} else if v.DepositPayStatus == "REFUND" && v.OrderStatus == "CANCEL" && v.RemainPayStatus == "CANCEL" && !g.IsEmpty(v.ConfirmRefuseReason) {
				//后台确认失败
				v.ActualOrderStatus = "CONFIRM_DISAGREE"
			} else if v.DepositPayStatus == "HAVE_PAID" && v.OrderStatus == "WAIT_PAY" && v.RemainPayStatus == "WAIT_PAY" && v.BookingStatus == "CONFIRMED" {
				//后台确认成功，尾款待支付
				v.ActualOrderStatus = "REMAIN_WAIT_PAY"
			} else if v.DepositPayStatus == "REFUND" && v.OrderStatus == "CANCEL" && v.RemainPayStatus == "CANCEL" && g.IsEmpty(v.ConfirmRefuseReason) {
				//定金已支付后台待确认主动取消 / 后台确认成功，尾款待支付主动取消 | 尾款支付超时 / 后台操作退款
				v.ActualOrderStatus = "REMAIN_CANCEL"
			} else if v.DepositPayStatus == "HAVE_PAID" && v.OrderStatus == "HAVE_PAID" && v.RemainPayStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "WAIT_VERIFY" && gtime.Now().Before(v.BookDatetime) {
				//定金已支付尾款已支付(待核销)
				v.ActualOrderStatus = "WAIT_VERIFY"
			} else if v.DepositPayStatus == "HAVE_PAID" && v.OrderStatus == "HAVE_PAID" && v.RemainPayStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "VERIFIED" {
				//核销成功（已完成）
				v.ActualOrderStatus = "COMPLETE"
			} else if v.DepositPayStatus == "REFUND" && v.OrderStatus == "CANCEL" && v.RemainPayStatus == "REFUND" && v.VerifyStatus == "WAIT_VERIFY" {
				//退款
				v.ActualOrderStatus = "REFUND"
			} else if v.DepositPayStatus == "HAVE_PAID" && v.OrderStatus == "HAVE_PAID" && v.RemainPayStatus == "HAVE_PAID" && v.BookingStatus == "CONFIRMED" && v.VerifyStatus == "WAIT_VERIFY" && gtime.Now().After(v.BookDatetime) {
				//逾期
				v.ActualOrderStatus = "OVERDUE"
			}
		}
	}
	return
}

func (s *sFoodOrder) View(ctx context.Context, in *input_food.FoodOrderViewInp) (res *input_food.FoodOrderViewModel, err error) {
	if err = dao.FoodOrder.Ctx(ctx).Unscoped().WithAll().Where(dao.FoodOrder.Columns().OrderSn, in.OrderSn).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
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
				res.MemberDetail = &struct {
					gmeta.Meta `orm:"table:hg_pms_member"`
					Id         int    `json:"id"    orm:"id"      dc:"id"`
					FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
					MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
				}{
					Id:       memberInfo.Id,
					FullName: memberInfo.FullName,
					MemberNo: memberInfo.MemberNo,
				}
			}
		}
	}

	// Check goods deleted
	if res.GoodsId > 0 {
		goodsDeletedAt, _ := dao.FoodGoods.Ctx(ctx).Unscoped().Fields(dao.FoodGoods.Columns().DeletedAt).
			Where(dao.FoodGoods.Columns().Id, res.GoodsId).Value()
		if !goodsDeletedAt.IsNil() && !goodsDeletedAt.IsEmpty() {
			res.GoodsDeleted = true
			// Manually load deleted goods info
			var goodsInfo *entity.FoodGoods
			if err = dao.FoodGoods.Ctx(ctx).Unscoped().Where(dao.FoodGoods.Columns().Id, res.GoodsId).Scan(&goodsInfo); err == nil && goodsInfo != nil {
				res.GoodsDetail = &struct {
					gmeta.Meta `orm:"table:hg_food_goods"`
					*entity.FoodGoods
				}{
					FoodGoods: goodsInfo,
				}
			}
		}
	}

	if res.OrderType == "TORETA" || res.OrderType == "CRSALL" {
		if res.OrderStatus == "WAIT_PAY" {
			// 待支付
			res.ActualOrderStatus = "WAIT_PAY"

		} else if res.OrderStatus == "CANCEL" && g.IsEmpty(res.ConfirmRefuseReason) && res.RefundStatus == "WAIT" {
			// 取消（未支付取消）
			res.ActualOrderStatus = "CANCEL"

		} else if res.OrderStatus == "HAVE_PAID" && res.BookingStatus == "WAIT_CONFIRM" {
			// 后台待确认
			res.ActualOrderStatus = "WAIT_CONFIRM"

		} else if res.OrderStatus == "CANCEL" && !g.IsEmpty(res.ConfirmRefuseReason) {
			// 后台确认失败
			res.ActualOrderStatus = "CONFIRM_DISAGREE"

		} else if res.OrderStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "WAIT_VERIFY" && gtime.Now().Before(res.BookDatetime) {
			// 待核销
			res.ActualOrderStatus = "WAIT_VERIFY"

		} else if res.OrderStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "VERIFIED" {
			// 核销成功（已完成）
			res.ActualOrderStatus = "COMPLETE"

		} else if res.OrderStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "WAIT_VERIFY" && gtime.Now().After(res.BookDatetime) {
			// 逾期
			res.ActualOrderStatus = "OVERDUE"

		} else if res.OrderStatus == "CANCEL" && g.IsEmpty(res.ConfirmRefuseReason) && res.RefundStatus != "WAIT" {
			// 退款
			res.ActualOrderStatus = "REFUND"
		}
	} else {
		if res.DepositPayStatus == "WAIT_PAY" && res.OrderStatus == "WAIT_PAY" && res.RemainPayStatus == "WAIT_PAY" {
			//定金待支付
			res.ActualOrderStatus = "DEPOSIT_WAIT_PAY"
		} else if res.DepositPayStatus == "CANCEL" && res.OrderStatus == "CANCEL" && res.RemainPayStatus == "CANCEL" {
			//定金未支付取消(主动取消 | 支付超时)
			res.ActualOrderStatus = "DEPOSIT_CANCEL"
		} else if res.DepositPayStatus == "HAVE_PAID" && res.OrderStatus == "WAIT_PAY" && res.RemainPayStatus == "WAIT_PAY" && res.BookingStatus == "WAIT_CONFIRM" {
			//定金已支付后台待确认
			res.ActualOrderStatus = "WAIT_CONFIRM"
		} else if res.DepositPayStatus == "REFUND" && res.OrderStatus == "CANCEL" && res.RemainPayStatus == "CANCEL" && !g.IsEmpty(res.ConfirmRefuseReason) {
			//后台确认失败
			res.ActualOrderStatus = "CONFIRM_DISAGREE"
		} else if res.DepositPayStatus == "HAVE_PAID" && res.OrderStatus == "WAIT_PAY" && res.RemainPayStatus == "WAIT_PAY" && res.BookingStatus == "CONFIRMED" {
			//后台确认成功，尾款待支付
			res.ActualOrderStatus = "REMAIN_WAIT_PAY"
		} else if res.DepositPayStatus == "REFUND" && res.OrderStatus == "CANCEL" && res.RemainPayStatus == "CANCEL" && g.IsEmpty(res.ConfirmRefuseReason) {
			//定金已支付后台待确认主动取消 / 后台确认成功，尾款待支付主动取消 | 尾款支付超时 / 后台操作退款
			res.ActualOrderStatus = "REMAIN_CANCEL"
		} else if res.DepositPayStatus == "HAVE_PAID" && res.OrderStatus == "HAVE_PAID" && res.RemainPayStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "WAIT_VERIFY" && gtime.Now().Before(res.BookDatetime) {
			//定金已支付尾款已支付(待核销)
			res.ActualOrderStatus = "WAIT_VERIFY"
		} else if res.DepositPayStatus == "HAVE_PAID" && res.OrderStatus == "HAVE_PAID" && res.RemainPayStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "VERIFIED" {
			//核销成功（已完成）
			res.ActualOrderStatus = "COMPLETE"
		} else if res.DepositPayStatus == "REFUND" && res.OrderStatus == "CANCEL" && res.RemainPayStatus == "REFUND" && res.VerifyStatus == "WAIT_VERIFY" {
			//退款
			res.ActualOrderStatus = "REFUND"
		} else if res.DepositPayStatus == "HAVE_PAID" && res.OrderStatus == "HAVE_PAID" && res.RemainPayStatus == "HAVE_PAID" && res.BookingStatus == "CONFIRMED" && res.VerifyStatus == "WAIT_VERIFY" && gtime.Now().After(res.BookDatetime) {
			//逾期
			res.ActualOrderStatus = "OVERDUE"
		}
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

	return
}

func (s *sFoodOrder) ConfirmAgree(ctx context.Context, in *input_food.FoodOrderConfirmAgreeInp) (err error) {

	var (
		models          *entity.FoodOrder
		oldBookDate     string
		oldBookTime     string
		oldBookDateTime *gtime.Time
		newBookDatetime *gtime.Time
		ExpirationTime  int
		PayConfig       *model.PayConfig
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderType != "CRS" {
		// 全款模式或toreta模式
		if models.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单支付状态不正确")
			return
		}

	} else {
		if models.DepositPayStatus != "HAVE_PAID" {
			err = gerror.New("定金支付状态不正确")
			return
		}
	}

	if models.BookingStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单确认状态不正确")
		return
	}

	if in.BookDate != models.BookDate || in.BookTime != models.BookTime {
		oldBookDate = models.BookDate
		oldBookTime = models.BookTime
		oldBookDateTime = gtime.NewFromStr(fmt.Sprintf("%s %s", oldBookDate, oldBookTime))
	}

	newBookDatetime = gtime.NewFromStr(fmt.Sprintf("%s %s", in.BookDate, in.BookTime))

	// 读取支付配置
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	ExpirationTime = gvar.New(gtime.Now().Unix() + PayConfig.HotelStayExp).Int()

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_food.FoodOrderConfirmAgreeFields{
			BookingStatus:   "CONFIRMED",
			BookingTime:     gtime.Now(),
			BookDate:        in.BookDate,
			BookTime:        in.BookTime,
			BookDatetime:    newBookDatetime,
			OldBookDate:     oldBookDate,
			OldBookTime:     oldBookTime,
			OldBookDatetime: oldBookDateTime,
			ExpirationTime:  ExpirationTime,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     int(in.Id),
			ActionWay:   "CONFIRMED",
			Remark:      "订单已确认",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 更新餐厅预定量和预定金额
		if _, err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
			dao.FoodRestaurant.Columns().WaitConfirmOrderNum: gdb.Raw("wait_confirm_order_num-1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
			return
		}

		// 投递自动过期订单队列 如果尾款在规定时间内未支付 必须取消订单
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeDelayedName,
			QueueName:    consts.RabbitMQQueueNameOrderExpire,
			DataByte:     gvar.New("R-" + models.OrderSn).Bytes(),
		}); err != nil {
			g.Log().Error(ctx, "发送过期自动取消订单MQ失败", err)
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已确认",
			"en":    "Order confirmed",
			"ja":    "注文確定",
			"ko":    "주문 확인됨",
			"zh_CN": "訂單已確認",
		}
		systemMessageContent := map[string]string{
			"zh":    models.OrderSn + "订单后台已接单",
			"en":    "Order " + models.OrderSn + " has been received in the system.",
			"ja":    "システムで注文" + models.OrderSn + "が受信されました。",
			"ko":    "시스템에 주문" + models.OrderSn + "이 접수되었습니다.",
			"zh_CN": models.OrderSn + "訂單後台已接單",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
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
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "food",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/DiningOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/restaurant-booking/pages/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

// ConfirmDisagree 确认失败，全额退款
func (s *sFoodOrder) ConfirmDisagree(ctx context.Context, in *input_food.FoodOrderConfirmDisagreeInp) (err error) {

	var models *entity.FoodOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderType != "CRS" {
		// 全款模式或toreta模式
		if models.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单支付状态不正确")
			return
		}

	} else {
		if models.DepositPayStatus != "HAVE_PAID" {
			err = gerror.New("定金支付状态不正确")
			return
		}
	}

	if models.BookingStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单确认状态不正确")
		return
	}

	var restaurantModel *entity.FoodRestaurant
	if err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Scan(&restaurantModel); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if restaurantModel == nil {
		err = gerror.New("餐厅信息不存在或已被删除")
		return
	}

	//if models.RefundStatus != "WAIT" {
	//	err = gerror.New("订单已退款")
	//	return
	//}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 改变订单确认状态
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(input_food.FoodOrderConfirmDisagreeFields{
			BookingStatus:       "CANCEL",
			ConfirmRefuseReason: in.BookingDisagreeReason,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     int(in.Id),
			ActionWay:   "DISCONFIRMED",
			Remark:      fmt.Sprintf("订单被拒，原因：%s", in.BookingDisagreeReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 确认失败、全额退款
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
		var DepositPayStatus string
		var DepositRefundTime *gtime.Time
		if models.OrderType != "CRS" {
			// 全款模式或toreta模式
			DepositPayStatus = "CANCEL"
			DepositRefundTime = nil
		} else {
			DepositPayStatus = "REFUND"
			DepositRefundTime = gtime.Now()
		}
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(g.MapStrAny{
			dao.FoodOrder.Columns().RefundAmount:       RefundAmount,
			dao.FoodOrder.Columns().RefundBalAmount:    RefundBalance,
			dao.FoodOrder.Columns().RefundCouponAmount: 0,
			dao.FoodOrder.Columns().RefundStatus:       RefundStatus,
			dao.FoodOrder.Columns().RefundTime:         gtime.Now(),
			dao.FoodOrder.Columns().OrderStatus:        "CANCEL",
			dao.FoodOrder.Columns().DepositPayStatus:   DepositPayStatus,
			dao.FoodOrder.Columns().DepositRefundTime:  DepositRefundTime,
			dao.FoodOrder.Columns().RemainPayStatus:    "CANCEL",
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if models.OrderType != "CRS" {
			// 全款模式或toreta模式
			// 更新餐厅预定量和预定金额
			if _, err = dao.FoodRestaurant.Ctx(ctx).TX(tx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().PayOrderNum:         gdb.Raw("pay_order_num-1"),
				dao.FoodRestaurant.Columns().PayOrderAmount:      gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
				dao.FoodRestaurant.Columns().WaitConfirmOrderNum: gdb.Raw("wait_confirm_order_num-1"),
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
		} else {
			// 更新餐厅预定量和预定金额
			if _, err = dao.FoodRestaurant.Ctx(ctx).TX(tx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().WaitConfirmOrderNum: gdb.Raw("wait_confirm_order_num-1"),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     int(in.Id),
			ActionWay:   "REFUND",
			Remark:      "订单已退款",
			OperateType: "SYSTEM",
		}); err != nil {
			return err
		}

		// 全额退款
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      models.OrderSn,
			RefundAmount: RefundAmount,
			OperateType:  "ADMIN",
			OperateId:    int(contexts.GetUserId(ctx)),
		}, tx)
		if err != nil {
			err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
			return
		}

		// 如果是toreta订单，需要走toreta取消接口
		if models.OrderType == "TORETA" {
			var (
				ToretaApiConfig *model.ToretaApiConfig
			)
			if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
				return
			}
			if _, err = toretaApi.NewClient(ctx, ToretaApiConfig).CancelReservation(ctx, &toretaApi.CancelReservationParams{
				ReservationId: models.ToretaReservationId,
				RestaurantId:  restaurantModel.ToretaId,
			}); err != nil {
				return
			}
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已被拒",
			"en":    "The order has been rejected.",
			"ja":    "注文は拒否されました。",
			"ko":    "주문이 거부되었습니다.",
			"zh_CN": "訂單已被拒",
		}
		systemMessageContent := map[string]string{
			"zh":    "订单" + models.OrderSn + "已被拒，已全额退款",
			"en":    "Order " + models.OrderSn + " has been rejected and fully refunded.",
			"ja":    "注文 " + models.OrderSn + " は拒否され、全額返金されました。",
			"ko":    "주문 " + models.OrderSn + "은 거부되었으며 전액 환불되었습니다.",
			"zh_CN": "訂單" + models.OrderSn + "已被拒，已全額退款",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
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
			"string": models.OrderSn,
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "food",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/DiningOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/restaurant-booking/pages/order-detail?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

// SettleOrderList 结算订单列表
func (s *sFoodOrder) SettleOrderList(ctx context.Context, in *input_food.SettleFoodOrderListInp) (list []*input_food.SettleFoodOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	mod = mod.Fields(input_food.SettleFoodOrderListModel{})

	if !g.IsEmpty(in.SettlementOrderId) {
		mod = mod.Where(dao.FoodOrder.Columns().SettlementOrderId, in.SettlementOrderId)
	}

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.FoodOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FoodOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.FoodOrder.Columns().Id)
	//mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算订单列表失败，请稍后重试！")
		return
	}
	return
}

// CancelPay 取消订单
func (s *sFoodOrder) CancelPay(ctx context.Context, in *input_food.FoodOrderCancelPayInp) (err error) {

	var models *entity.FoodOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderType == "CRS" {
		if models.DepositPayStatus != "HAVE_PAID" {
			err = gerror.New("定金支付状态不正确")
			return
		}

		if models.RemainPayStatus != "HAVE_PAID" {
			err = gerror.New("尾款支付状态不正确")
			return
		}
	}

	if models.OrderType == "TORETA" || models.OrderType == "CRSALL" {
		if models.OrderStatus != "HAVE_PAID" {
			err = gerror.New("订单状态不正确")
			return
		}
	}

	if models.VerifyStatus != "WAIT_VERIFY" {
		err = gerror.New("订单已核销")
		return
	}

	if models.AdminCancelNum > 0 {
		err = gerror.New("后台不可多次退款")
		return
	}

	var restaurantModel *entity.FoodRestaurant
	if err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Scan(&restaurantModel); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if restaurantModel == nil {
		err = gerror.New("餐厅信息不存在或已被删除")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 确认失败、全额退款
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			CancelFee         float64 // 退款手续费
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		CancelFee = models.OrderAmount - models.CouponAmount - in.RefundMoney
		CancelFeeCalc := CancelFee
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().PayStatus, "DONE").Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).OrderDesc(`
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
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

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
		// 退款金额
		RefundAmount := RefundBalance + RefundFee

		OrderTotalRefundAmount := RefundAmount + models.RefundAmount
		OrderTotalRefundBalance := RefundBalance + models.RefundBalAmount

		AdminOrderTotalRefundAmount := RefundAmount + models.AdminRefundAmount
		AdminOrderTotalRefundBalance := RefundBalance + models.AdminRefundBalAmount

		// 修改订单状态
		RefundStatus := "DONE"
		if models.OrderAmount > OrderTotalRefundAmount {
			RefundStatus = "PART"
		}
		if in.RefundType == 1 {
			// 仅退款
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				//dao.FoodOrder.Columns().RefundFee:          CancelFee,
				dao.FoodOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.FoodOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.FoodOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.FoodOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.FoodOrder.Columns().RefundCouponAmount:   0,
				dao.FoodOrder.Columns().RefundStatus:         RefundStatus,
				dao.FoodOrder.Columns().RefundTime:           gtime.Now(),
				dao.FoodOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 更新餐厅预定量和预定金额
			if _, err = dao.FoodRestaurant.Ctx(ctx).TX(tx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
				dao.FoodRestaurant.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
				return
			}

			// 更新套餐预定量和预定金额
			if _, err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
				dao.FoodGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新套餐信息失败，请稍后重试！")
				return
			}
		} else {
			// 退款并取消
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				dao.FoodOrder.Columns().RefundFee:          CancelFee,
				dao.FoodOrder.Columns().RefundAmount:       OrderTotalRefundAmount,
				dao.FoodOrder.Columns().RefundBalAmount:    OrderTotalRefundBalance,
				dao.FoodOrder.Columns().RefundCouponAmount: 0,
				dao.FoodOrder.Columns().RefundStatus:       RefundStatus,
				dao.FoodOrder.Columns().RefundTime:         gtime.Now(),
				dao.FoodOrder.Columns().OrderStatus:        "CANCEL",
				dao.FoodOrder.Columns().BookingStatus:      "CANCEL",
				dao.FoodOrder.Columns().DepositPayStatus:   "REFUND",
				dao.FoodOrder.Columns().RemainPayStatus:    "REFUND",
				dao.FoodOrder.Columns().DepositRefundTime:  gtime.Now(),
				dao.FoodOrder.Columns().RemainRefundTime:   gtime.Now(),
				dao.FoodOrder.Columns().RemainCancelSource: "ADMIN",
				dao.FoodOrder.Columns().RemainCancelReason: "后台取消订单",
				dao.FoodOrder.Columns().AdminCancelReason:  in.AdminCancelReason,
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

			if models.BookingStatus == "WAIT_CONFIRM" {
				if _, err = dao.FoodRestaurant.Ctx(ctx).TX(tx).Where(dao.FoodRestaurant.Columns().Id, models.RestaurantId).Update(g.MapStrAny{
					dao.FoodRestaurant.Columns().WaitConfirmOrderNum: gdb.Raw("wait_confirm_order_num-1"),
				}); err != nil {
					err = gerror.Wrap(err, "更新餐厅信息失败，请稍后重试！")
					return
				}
			}

			// 更新套餐预定量和预定金额
			if _, err = dao.FoodGoods.Ctx(ctx).Where(dao.FoodGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
				dao.FoodGoods.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", models.GoodsNum)),
				dao.FoodGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新套餐信息失败，请稍后重试！")
				return
			}
		}

		if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).WherePri(in.Id).Update(g.MapStrAny{
			dao.FoodOrder.Columns().AdminCancelNum: gdb.Raw("admin_cancel_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新失败，请稍后重试！")
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
		if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     int(in.Id),
			ActionWay:   "ADMIN_REFUND",
			Remark:      fmt.Sprintf("后台退款，原因：%s", in.AdminCancelReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		if in.RefundType == 2 {
			// 订单日志
			if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
				OrderId:     int(in.Id),
				ActionWay:   "CANCEL",
				Remark:      "后台订单取消",
				OperateType: "ADMIN",
				OperateId:   int(contexts.GetUserId(ctx)),
			}); err != nil {
				return err
			}
		}

		if in.RefundType == 2 {
			// 如果是toreta订单，需要走toreta取消接口
			if models.OrderType == "TORETA" {
				var (
					ToretaApiConfig *model.ToretaApiConfig
				)
				if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
					return
				}
				if _, err = toretaApi.NewClient(ctx, ToretaApiConfig).CancelReservation(ctx, &toretaApi.CancelReservationParams{
					ReservationId: models.ToretaReservationId,
					RestaurantId:  restaurantModel.ToretaId,
				}); err != nil {
					return
				}
			}

			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单已取消",
				"en":    "Order has been canceled",
				"ja":    "注文はキャンセルされました",
				"ko":    "주문이 취소되었습니다",
				"zh_CN": "訂單已取消",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and " + gvar.New(RefundBalance).String() + " points have been refunded",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消",
			// 		"en":    "Order has been canceled",
			// 		"ja":    "注文はキャンセルされました",
			// 		"ko":    "주문이 취소되었습니다",
			// 		"zh_CN": "訂單已取消",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消，并成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Order has been canceled, and a refund of " + gvar.New(RefundAmount).String() + "JPY has been issued",
					"ja":    "注文はキャンセルされ、" + gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    "주문이 취소되었으며 " + gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消",
					"en":    "Order has been canceled",
					"ja":    "注文はキャンセルされました",
					"ko":    "주문이 취소되었습니다",
					"zh_CN": "訂單已取消",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
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
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "food",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/DiningOrderDetailPage",
				WxLink:               fmt.Sprintf("/subpackages/restaurant-booking/pages/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		} else {
			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单成功退款",
				"en":    "Refund Successful",
				"ja":    "返金が正常に完了しました",
				"ko":    "환불이 성공적으로 완료되었습니다",
				"zh_CN": "訂單成功退款",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款",
			// 		"en":    "Refund Successful",
			// 		"ja":    "返金が正常に完了しました",
			// 		"ko":    "환불이 성공적으로 완료되었습니다",
			// 		"zh_CN": "訂單成功退款",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Refund of " + gvar.New(RefundAmount).String() + "JPY completed",
					"ja":    gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單成功退款" + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款",
					"en":    "Refund Successful",
					"ja":    "返金が正常に完了しました",
					"ko":    "환불이 성공적으로 완료되었습니다",
					"zh_CN": "訂單成功退款",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
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
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "food",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/DiningOrderDetailPage",
				WxLink:               fmt.Sprintf("/subpackages/restaurant-booking/pages/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		}

		return
	})

}

func (s *sFoodOrder) RefreshCode(ctx context.Context, in *input_food.FoodsOrderRefreshCodeInp) (code string, verifyStatus int, err error) {

	var foodOrderInfo *struct {
		Id           int
		OrderSn      string
		MemberId     int
		VerifyStatus string
	}
	if err = s.Model(ctx).WherePri(in.Id).Scan(&foodOrderInfo); err != nil {
		err = gerror.Wrap(err, "获取餐厅订单信息失败，请稍后重试！")
		return
	}
	if foodOrderInfo.MemberId != in.MemberId {
		err = gerror.New("会员信息不匹配")
		return
	}

	verifyStatus = 1
	if foodOrderInfo.VerifyStatus == "VERIFIED" {
		verifyStatus = 2
	}

	// 生成券码，采用aes加密，将couponNo|会员ID|时间戳进行加密
	timestamp := gtime.Now().Unix()
	plainText := fmt.Sprintf("%s|%d|%d", foodOrderInfo.OrderSn, foodOrderInfo.MemberId, timestamp)

	// 使用AES加密生成动态券码
	code = encrypt.MustAesECBEncryptToString(plainText, string(consts.RequestEncryptKey))

	return
}

func (s *sFoodOrder) Verify(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error) {
	Logger := g.Log().Path("logs/Verify")
	Logger.Info(ctx, "--------餐厅订单核销进入----------")
	Logger.Info(ctx, gjson.New(in))

	if g.IsEmpty(in.Sn) {
		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
	}

	if g.IsEmpty(in.Val) {
		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
	}

	var terminalModel *input_basics.TerminalViewModel
	var foodOrderInfo *input_food.FoodOrderVerifyViewModel
	var terminalRestaurantModel *entity.FoodRestaurantTerminal

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 终端信息
		if err = dao.SysTerminal.Ctx(ctx).TX(tx).
			Where(dao.SysTerminal.Columns().Sn, in.Sn).
			WithAll().
			Scan(&terminalModel); err != nil {
			err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
			return
		}

		// 验签
		//sign := util2.Sign(in.Sn + gvar.New(in.Timestamp).String() + in.Val + terminalModel.BrandInfo.ClientSecret)
		//
		//if sign != in.Sign {
		//	err = gerror.New("签名不正确")
		//	return
		//}

		if terminalModel.RestaurantId <= 0 {
			err = gerror.New("不可核销餐厅订单")
			return
		}

		// 解密动态券码并验证时间戳
		var orderSn string
		var memberId int
		var timestamp int64
		var isDynamicCode bool = false

		valArr := strings.Split(in.Val, "|")
		encCode := valArr[0]
		Logger.Info(ctx, "--------未解密字符串----------")
		Logger.Info(ctx, encCode)

		// 尝试解密动态券码
		// 首先尝试base64解码
		encryptedData, base64Err := base64.StdEncoding.DecodeString(encCode)

		if base64Err != nil {
			// 如果base64解码失败，可能是旧的券码格式，直接使用原券号查询
			orderSn = encCode
			Logger.Info(ctx, "--------base64Err----------")
			Logger.Info(ctx, "base64转义失败")
		} else {
			// base64解码成功，尝试AES解密
			decryptedText, decryptErr := encrypt.AesECBDecrypt(encryptedData, consts.RequestEncryptKey)
			Logger.Info(ctx, "--------核销解码----------")
			Logger.Info(ctx, decryptedText)
			if decryptErr != nil {
				// 如果解密失败，可能是旧的券码格式，直接使用原券号查询
				orderSn = encCode
				Logger.Info(ctx, "base64解码成功，AES解密失败")
			} else {
				// 解密成功，解析格式：orderSn|memberId|timestamp
				parts := strings.Split(string(decryptedText), "|")
				if len(parts) != 3 {
					err = gerror.New("动态券码格式不正确！")
					return
				}

				orderSn = parts[0]
				memberId, err = strconv.Atoi(parts[1])
				if err != nil {
					err = gerror.New("动态券码会员ID格式不正确！")
					return
				}

				timestamp, err = strconv.ParseInt(parts[2], 10, 64)
				if err != nil {
					err = gerror.New("动态券码时间戳格式不正确！")
					return
				}

				// 验证时间戳是否在30秒内
				currentTime := gtime.Now().Unix()
				if currentTime-timestamp > 3000 {
					err = gerror.New("动态券码已过期，请重新生成！")
					return
				}

				isDynamicCode = true
			}
		}

		Logger.Info(ctx, "--------OrderSn----------")
		Logger.Info(ctx, orderSn)

		// 订单信息
		if err = dao.FoodOrder.Ctx(ctx).WithAll().Where(dao.FoodOrder.Columns().OrderSn, orderSn).Scan(&foodOrderInfo); err != nil {
			err = gerror.Wrap(err, "获取订单信息失败，请稍后重试！")
			return
		}

		if g.IsEmpty(foodOrderInfo) {
			err = gerror.New("订单不存在")
			return
		}

		// 判断会员是否已注销
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, foodOrderInfo.MemberId).Count()
		if memberCount == 0 {
			err = gerror.New("会员已注销，无法操作此订单")
			return
		}

		// 如果使用的是动态券码，还需要验证会员ID是否匹配
		if isDynamicCode && foodOrderInfo.MemberId != uint(memberId) {
			err = gerror.New("动态券码会员信息不匹配！")
			return
		}

		if foodOrderInfo.RestaurantId != terminalModel.RestaurantId {
			err = gerror.New("该餐厅订单【" + gvar.New(foodOrderInfo.RestaurantId).String() + "】不属于您的餐厅【" + gvar.New(terminalModel.RestaurantId).String() + "】，不可核销")
			return
		}

		// 判断订单状态
		if foodOrderInfo.OrderStatus != "HAVE_PAID" || foodOrderInfo.BookingStatus != "CONFIRMED" {
			err = gerror.New("订单状态不正确")
			return
		}

		if gtime.Now().After(foodOrderInfo.BookDatetime) {
			err = gerror.New("您已逾期，无法核销")
			return
		}

		if foodOrderInfo.VerifyStatus != "WAIT_VERIFY" {
			err = gerror.New("订单核销状态不正确")
			return
		}

		// 获取餐厅终端关联表信息，获取打印次数
		if err = dao.FoodRestaurantTerminal.Ctx(ctx).TX(tx).
			Where(dao.FoodRestaurantTerminal.Columns().TerminalId, terminalModel.Id).
			Where(dao.FoodRestaurantTerminal.Columns().RestaurantId, terminalModel.RestaurantId).
			Scan(&terminalRestaurantModel); err != nil {
			err = gerror.Wrap(err, "获取餐厅终端关联信息失败，请稍后重试！")
			return
		}

		// 记录分佣 并 核销
		cost := foodOrderInfo.RestaurantDetail.SettlementDetail.Cost
		settlementAmount := foodOrderInfo.OrderAmount
		if !g.IsEmpty(cost) {
			costArr := strings.Split(cost, ",")
			constIsCoupon := false
			constIsBal := false
			for _, costItem := range costArr {
				if gvar.New(costItem).Int() == 1 {
					constIsCoupon = true
				}
				if gvar.New(costItem).Int() == 2 {
					constIsBal = true
				}
			}
			if !constIsCoupon {
				settlementAmount = settlementAmount - foodOrderInfo.CouponAmount
			}
			if !constIsBal {
				settlementAmount = settlementAmount - foodOrderInfo.BalAmount
			}
		}

		settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(foodOrderInfo.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()

		if _, err = dao.FoodOrder.Ctx(ctx).TX(tx).
			WherePri(foodOrderInfo.Id).Update(g.MapStrAny{
			dao.FoodOrder.Columns().SettlementAmount: settlementAmount,
			dao.FoodOrder.Columns().VerifyStatus:     "VERIFIED",
			dao.FoodOrder.Columns().VerifyTime:       gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "核销失败，请稍后重试！")
			return
		}

		// 写入终端核销日志
		if _, err = dao.SysTerminalVerify.Ctx(ctx).Insert(&entity.SysTerminalVerify{
			TerminalId:     terminalModel.Id,
			VerifyType:     "FOOD_ORDER",
			MchId:          0,
			RestaurantId:   foodOrderInfo.RestaurantId,
			FoodOrderId:    gvar.New(foodOrderInfo.Id).Int(),
			VerifyMemberId: gvar.New(foodOrderInfo.MemberId).Int(),
			VerifyTime:     gtime.Now(),
		}); err != nil {
			return
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     gvar.New(foodOrderInfo.Id).Int(),
			ActionWay:   "VERIFIED",
			Remark:      "订单已核销",
			OperateType: "SYSTEM",
			//OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 转发到返利队列 订单计算佣金/计算经验需要扔队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(foodOrderInfo.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(foodOrderInfo.OrderSn).Bytes(),
			Header:       nil,
		})

		// 发送分销订单变更队列
		if foodOrderInfo.IsFx == "Y" {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
				DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
					OrderNo:      foodOrderInfo.OrderSn,
					ChangeStatus: "COMPLETE",
				}).MustToJson(),
				Header: nil,
			})
		}

		return
	})

	if err != nil {
		Logger.Info(ctx, "--------核销失败----------")
		Logger.Info(ctx, err)

		// 发送消息
		_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
			MemberId: gvar.New(foodOrderInfo.MemberId).Int(),
			Code:     -1,
			Event:    "VERIFY",
			Message:  "核销失败",
		})

		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
		Logger.Info(ctx, "--------核销失败111----------")
		return
	}

	Logger.Info(ctx, "--------核销成功----------")

	printTimes := 1
	if !g.IsEmpty(terminalRestaurantModel) {
		printTimes = int(terminalRestaurantModel.PrintTimes)
	}

	// 打印
	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"
	printContent = printContent + "<CB>引換券<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.Now().Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><BOLD>ご予約者様名：" + foodOrderInfo.Member.FirstName + " 様</BOLD><BR>"
	printContent = printContent + "<L><BOLD>人数：" + gvar.New(foodOrderInfo.BookingCount).String() + "人</BOLD><BR>"
	printContent = printContent + "<L><BOLD>予約日：" + gvar.New(foodOrderInfo.BookDate).String() + "</BOLD><BR>"
	printContent = printContent + "<L><BOLD>予約時間：" + foodOrderInfo.BookTime + "</BOLD><BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><BOLD>" + foodOrderInfo.GoodsDetail.NameLanguage.Content + " *" + gvar.New(foodOrderInfo.BookingCount).String() + "</BOLD><BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>--------------------------------"
	printContent = printContent + "<BR>"
	if !g.IsEmpty(foodOrderInfo.MemberMessage) {
		printContent = printContent + "<L><HB>コメント：" + foodOrderInfo.MemberMessageJa + "<BR>"
		printContent = printContent + "<BR>"
		printContent = printContent + "<L>备注：" + foodOrderInfo.MemberMessage + "<BR>"
		printContent = printContent + "<BR>"
	}
	printContent = printContent + "<L><N>注文番号:" + foodOrderInfo.OrderSn + "<BR>"
	printContent = printContent + "顧客情報：" + foodOrderInfo.BookingName + "<BR>"
	printContent = printContent + "連絡先：" + foodOrderInfo.PhoneArea + "-" + foodOrderInfo.BookingMobile + "<BR>"
	printContent = printContent + "使用時間：" + gtime.Now().Format("Y-m-d H:i:s") + "<BR>"
	printContent = printContent + "場所：" + foodOrderInfo.RestaurantDetail.NameLanguage.Content + "<BR>"
	printContent = printContent + "<N>〒" + foodOrderInfo.RestaurantDetail.DetailAddress + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR><BR>"

	Logger.Info(ctx, "--------打印内容----------")
	Logger.Info(ctx, printContent)
	if printTimes > 0 {
		Logger.Info(ctx, "--------开始打印----------")
		err = service.BasicsTerminal().Printer(ctx, &input_basics.PrinterInp{
			Sn:           terminalModel.Sn,
			PrintContent: printContent,
			PrintTimes:   printTimes,
		})
		if err != nil {
			return
		}
	}

	// 发送消息
	_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
		MemberId: gvar.New(foodOrderInfo.MemberId).Int(),
		Code:     200,
		Event:    "VERIFY",
		Message:  "核销成功",
	})

	ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"data": "OK",
	})

	return
}

func (s *sFoodOrder) ExportOrder(ctx context.Context, in *input_food.FoodOrderExportInp) (err error) {
	var (
		lastInsertId int64
	)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		resultJson, _ := json.Marshal(in)

		if lastInsertId, err = dao.OrderExport.Ctx(ctx).
			Data(entity.OrderExport{
				Scene:     2,
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

func (s *sFoodOrder) StartExport(ctx context.Context, in *input_food.FoodOrderExportInp) (path string, err error) {
	var (
		ChangeList  []*input_food.FoodOrderExportModel
		OrderListIn *input_food.FoodOrderListInp
	)

	OrderListIn = &input_food.FoodOrderListInp{
		PageReq: input_form.PageReq{
			Pagination: false,
		},
		MemberId:         in.MemberId,
		OrderType:        in.OrderType,
		OrderSn:          in.OrderSn,
		RestaurantName:   in.RestaurantName,
		BookingStatus:    in.BookingStatus,
		DepositPayStatus: in.DepositPayStatus,
		OrderStatus:      in.OrderStatus,
		RemainPayStatus:  in.RemainPayStatus,
		ToretaPayStatus:  in.ToretaPayStatus,
		GoodsName:        in.GoodsName,
		CreatedAt:        in.CreatedAt,
	}
	//if !g.IsEmpty(in.CreatedAt) {
	//	for _, timeItem := range strings.Split(in.CreatedAt, ",") {
	//		OrderListIn.CreatedAt = append(OrderListIn.CreatedAt, gtime.New(timeItem).Format("Y-m-d H:i:s"))
	//	}
	//}
	list, _, err := s.List(ctx, OrderListIn)
	if err != nil {
		return
	}
	for _, item := range list {
		itemData := &input_food.FoodOrderExportModel{
			OrderSn:           item.OrderSn,
			MemberNo:          item.MemberDetail.MemberNo,
			BookingName:       item.BookingName,
			BookStartTime:     item.BookDate + " " + item.BookTime,
			OrderTime:         item.CreatedAt.Format("Y-m-d H:i:s"),
			TotalAmount:       item.OrderAmount,
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
	tags, err := convert.GetEntityDescTags(input_food.FoodOrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出餐厅订单-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("餐厅订单")
		exports   []input_food.FoodOrderExportModel
	)

	if err = gconv.Scan(ChangeList, &exports); err != nil {
		return
	}

	path, err = excel.ExportByStructsFile(ctx, tags, exports, fileName, sheetName)

	return
}

func (s *sFoodOrder) ExportList(ctx context.Context, in *input_food.FoodOrderExportListInp) (list []*input_food.FoodOrderExportListModel, totalCount int, err error) {
	mod := dao.OrderExport.Ctx(ctx)

	mod = mod.Fields(input_food.FoodOrderExportListModel{})

	mod = mod.Where(dao.OrderExport.Columns().Scene, 2)

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

func (s *sFoodOrder) SyncToretaNotifications(ctx context.Context, Logger *glog.Logger) (err error) {
	var (
		ToretaApiConfig *model.ToretaApiConfig
		toretaResponse  *toretaApi.NotificationsResponse
		startTime       int64
		endTime         int64
		failedCount     int
	)

	// 先处理TORETA通知失败日志表中未处理和处理失败的数据
	_ = s.handleFailedNotifications(ctx, Logger)

	// 计算上一天的00:00:00和23:59:59
	yesterday := gtime.Now().AddDate(0, 0, -1)
	startTime = gtime.New(yesterday.Format("Y-m-d") + " 00:00:00").Timestamp()
	endTime = gtime.New(yesterday.Format("Y-m-d") + " 23:59:59").Timestamp()

	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).Notifications(ctx, &toretaApi.NotificationsParams{
		StartAt: int(startTime),
		EndAt:   int(endTime),
	}); err != nil {
		return
	}

	for _, notification := range *toretaResponse {
		if notification.StatusCode != 200 {
			// 目前仅支持预约类型的通知
			if notification.ResourceType != "Reservation" {
				Logger.Warning(ctx, "Unsupported resource type:", notification.ResourceType)
				continue
			}
			failedCount++
			// 处理预约通知
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				if err = s.handleReservationNotification(ctx, tx, &notification, Logger); err != nil {
					Logger.Error(ctx, "Failed to handle reservation notification:", err)
					// 将记录写入TORETA通知失败日志表中
					if _, err = dao.ToretaNotificationFailedLog.Ctx(ctx).Insert(g.Map{
						dao.ToretaNotificationFailedLog.Columns().ResourceKey:    notification.ResourceKey,
						dao.ToretaNotificationFailedLog.Columns().ResourceType:   notification.ResourceType,
						dao.ToretaNotificationFailedLog.Columns().ResourceAction: notification.ResourceAction,
						dao.ToretaNotificationFailedLog.Columns().RestaurantKey:  notification.RestaurantKey,
						dao.ToretaNotificationFailedLog.Columns().StatusCode:     notification.StatusCode,
					}); err != nil {
						Logger.Error(ctx, "Failed to insert toreta notification failed log:", gjson.New(notification).String())
						return err
					}
					return nil
				}

				return nil
			})

			if err != nil {
				Logger.Error(ctx, "Failed to process notification in transaction:", notification.ResourceKey, err)
				continue
			}
		}
	}

	Logger.Info(ctx, "同步TORETA通知失败次数：", failedCount)
	return
}

func (s *sFoodOrder) handleFailedNotifications(ctx context.Context, Logger *glog.Logger) (err error) {
	var (
		failedNotifications []*entity.ToretaNotificationFailedLog
	)

	// 获取TORETA通知失败日志表中statuscode为500的数据
	if err = dao.ToretaNotificationFailedLog.Ctx(ctx).Where(dao.ToretaNotificationFailedLog.Columns().StatusCode, 500).WhereLTE(dao.ToretaNotificationFailedLog.Columns().RetryCount, 3).Scan(&failedNotifications); err != nil {
		return
	}

	for _, notification := range failedNotifications {
		toretaNotification := &toretaApi.NotificationItem{
			ResourceKey:    notification.ResourceKey,
			ResourceType:   notification.ResourceType,
			ResourceAction: notification.ResourceAction,
			RestaurantKey:  notification.RestaurantKey,
			StatusCode:     notification.StatusCode,
		}

		// 使用事务保证原子性
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 更新重试次数
			if _, err := dao.ToretaNotificationFailedLog.Ctx(ctx).TX(tx).WherePri(notification.Id).Update(g.Map{
				dao.ToretaNotificationFailedLog.Columns().RetryCount: notification.RetryCount + 1,
			}); err != nil {
				return err
			}

			// 处理预约通知（简化版，只更新订单状态，不涉及复杂的退款逻辑）
			if err := s.handleReservationNotification(ctx, tx, toretaNotification, Logger); err != nil {
				Logger.Error(ctx, "Failed to handle reservation notification database:"+notification.ResourceKey, err)
				return err
			}

			// 将statusCode变为200
			if _, err := dao.ToretaNotificationFailedLog.Ctx(ctx).TX(tx).WherePri(notification.Id).Update(g.Map{
				dao.ToretaNotificationFailedLog.Columns().StatusCode: 200,
			}); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			Logger.Error(ctx, "Failed to process notification in transaction:", notification.ResourceKey, err)
			continue
		}
	}

	return
}

func (s *sFoodOrder) handleReservationNotification(ctx context.Context, tx gdb.TX, notification *toretaApi.NotificationItem, Logger *glog.Logger) (err error) {
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
	switch notification.ResourceAction {
	case "create":
		Logger.Info(ctx, "Processing reservation creation for:", notification.ResourceKey)
		// 获取预约详情并创建本地订单
		if reservation, err = client.ReservationDetail(ctx, &toretaApi.ReservationDetailParams{
			ReservationId: notification.ResourceKey,
		}); err != nil {
			Logger.Error(ctx, "Failed to get reservation details:", err)
			return
		}
		Logger.Info(ctx, "Reservation details:", reservation)
		// 创建本地订单逻辑
		if err = s.createLocalOrder(ctx, tx, notification, reservation, Logger); err != nil {
			Logger.Error(ctx, "Failed to create local order:", err)
			return
		}
	case "update":
		Logger.Info(ctx, "Processing reservation update for:", notification.ResourceKey)
		// 获取预约详情并更新本地订单状态
		if reservation, err = client.ReservationDetail(ctx, &toretaApi.ReservationDetailParams{
			ReservationId: notification.ResourceKey,
		}); err != nil {
			Logger.Error(ctx, "Failed to get reservation details:", err)
			return
		}
		Logger.Info(ctx, "Updated reservation details:", reservation)
		// 更新本地订单状态逻辑
		if err = s.updateLocalOrder(ctx, tx, notification, reservation, Logger); err != nil {
			Logger.Error(ctx, "Failed to update local order:", err)
			return
		}

	case "destroy":
		Logger.Info(ctx, "Processing reservation cancellation for:", notification.ResourceKey)
		// 取消本地订单
		if err = s.cancelLocalOrder(ctx, tx, notification, Logger); err != nil {
			Logger.Error(ctx, "Failed to cancel local order:", err)
			return
		}

	default:
		Logger.Warning(ctx, "Unknown resource action:", notification.ResourceAction)
	}
	return
}

func (s *sFoodOrder) createLocalOrder(ctx context.Context, tx gdb.TX, notification *toretaApi.NotificationItem, reservation *toretaApi.ReservationDetailResponse, Logger *glog.Logger) (err error) {
	// 注意：通常情况下，Toreta的预约是通过我们的系统创建的，所以create事件可能不需要特殊处理
	// 这里主要是记录日志，实际的订单创建逻辑在预约接口中已经处理
	Logger.Info(ctx, "Toreta reservation created - this might be from our own system")
	Logger.Info(ctx, "Reservation ID:", notification.ResourceKey)
	Logger.Info(ctx, "Restaurant Key:", notification.RestaurantKey)

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
		TX(tx).
		Where("toreta_reservation_id", notification.ResourceKey).
		Data(updateData).
		Update(); err != nil {
		Logger.Error(ctx, "Failed to create order in database:", err)
		return
	}

	Logger.Info(ctx, "Successfully create local order status")
	return
}

func (s *sFoodOrder) updateLocalOrder(ctx context.Context, tx gdb.TX, notification *toretaApi.NotificationItem, reservation *toretaApi.ReservationDetailResponse, Logger *glog.Logger) (err error) {
	Logger.Info(ctx, "Updating local order for Toreta reservation:", notification.ResourceKey)

	// 根据Toreta预约状态更新本地订单
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
		"booking_count":             reservation.Seats,
		"goods_num":                 reservation.Seats,
	}

	// 如果有开始时间，也更新
	if reservation.StartAt > 0 {
		startTime := gtime.NewFromTimeStamp(int64(reservation.StartAt))
		updateData["book_date"] = startTime.Format("Y-m-d")
		updateData["book_time"] = startTime.Format("H:i")
		updateData["book_datetime"] = startTime
	}

	// 如果有结束时间，也更新
	if reservation.EndAt > 0 {
		endTime := gtime.NewFromTimeStamp(int64(reservation.EndAt))
		updateData["toreta_reservation_endtime"] = endTime
	}

	// 根据Toreta预约ID更新订单
	if _, err = dao.FoodOrder.Ctx(ctx).
		TX(tx).
		Where("toreta_reservation_id", notification.ResourceKey).
		Data(updateData).
		Update(); err != nil {
		Logger.Error(ctx, "Failed to update order in database:", err)
		return
	}

	// 如果状态码返回的是2：预约取消，则订单进行退款取消处理
	if toretaStatus == 2 {
		var models *entity.FoodOrder
		if err = dao.FoodOrder.Ctx(ctx).Where("toreta_reservation_id", notification.ResourceKey).Scan(&models); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		// 如果订单已经退款 则不做任何处理
		if models.OrderStatus == "CANCEL" {
			return
		}
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
			dao.FoodOrder.Columns().RefundAmount:       RefundAmount,
			dao.FoodOrder.Columns().RefundBalAmount:    RefundBalance,
			dao.FoodOrder.Columns().RefundCouponAmount: 0,
			dao.FoodOrder.Columns().RefundStatus:       RefundStatus,
			dao.FoodOrder.Columns().RefundTime:         gtime.Now(),
			dao.FoodOrder.Columns().OrderStatus:        "CANCEL",
			dao.FoodOrder.Columns().DepositPayStatus:   "REFUND",
			dao.FoodOrder.Columns().DepositRefundTime:  gtime.Now(),
			dao.FoodOrder.Columns().RemainPayStatus:    "CANCEL",
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.FoodOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.FoodOrderLog{
			OrderId:     int(models.Id),
			ActionWay:   "CANCEL",
			Remark:      "订单取消",
			OperateType: "SYSTEM",
		}); err != nil {
			return err
		}

		// 全额退款
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      models.OrderSn,
			RefundAmount: RefundAmount,
		}, tx)
		if err != nil {
			err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
			return
		}

		return
	}

	Logger.Info(ctx, "Successfully updated local order status")
	return
}

func (s *sFoodOrder) cancelLocalOrder(ctx context.Context, tx gdb.TX, notification *toretaApi.NotificationItem, Logger *glog.Logger) (err error) {
	Logger.Info(ctx, "Cancelling local order for Toreta reservation:", notification.ResourceKey)
	return
}
