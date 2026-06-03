package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TranslateTextRes struct {
	Zh   string `json:"zh" dc:"简体中文翻译结果"`
	ZhCN string `json:"zh_CN" dc:"繁体翻译结果"`
	Ko   string `json:"ko" dc:"韩文翻译结果"`
	Ja   string `json:"ja" dc:"日文翻译结果"`
	En   string `json:"en" dc:"英文翻译结果"`
}

func (s *sHotelService) HotelOrderAward(ctx context.Context, OrderId int, CheckInDate string, MemberId int) (err error) {
	var (
		MemberInfo  *entity.PmsMember
		MemberScene *entity.PmsMemberScene
		StartTime   *gtime.Time
	)

	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).WherePri(MemberId).Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberInfo) {
		err = gerror.New("会员已注销")
		return
	}

	if err = dao.PmsMemberScene.Ctx(ctx).Where(dao.PmsMemberScene.Columns().Id, 1).Scan(&MemberScene); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if !g.IsEmpty(MemberScene) {
		if MemberScene.IsOpenReward == 1 {
			systemMessageTitle1 := map[string]string{
				"zh":    "下单奖励",
				"en":    "Order Rewards",
				"ja":    "注文特典",
				"ko":    "주문 보상",
				"zh_CN": "下單獎勵",
			}

			// 根据语言获取优惠券名称
			var couponNameLanguage []*input_hotel.LanguageType
			getCouponNameByLanguage := func(language string) string {
				for _, lang := range couponNameLanguage {
					if lang.Language == language {
						return lang.Content
					}
				}
				// 如果没找到对应语言，返回第一个或默认值
				if len(couponNameLanguage) > 0 {
					return couponNameLanguage[0].Content
				}
				return "优惠券"
			}

			// 奖励类型
			RewardTypeArr := strings.Split(MemberScene.RewardType, ",")
			for _, rewardType := range RewardTypeArr {
				if rewardType == "coupon" && !g.IsEmpty(MemberScene.RewardCouponTypeIds) {
					// 奖励会员优惠券  source：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
					CouponTypeIdsArr := strings.Split(MemberScene.RewardCouponTypeIds, ",")

					for _, couponTypeId := range CouponTypeIdsArr {
						var memberCouponId1 int64
						if memberCouponId1, couponNameLanguage, err = service.BasicsCouponType().SendMemberCouponGetId(ctx, &input_basics.PmsSendMemberCouponInp{
							MemberId:      MemberId,
							CouponTypeId:  gvar.New(couponTypeId).Int(),
							SourceOrderId: OrderId,
						}, 2); err != nil {
						}

						if err == nil {
							// 下单奖励发送消息
							systemMessageContent1 := map[string]string{
								"zh":    fmt.Sprintf("下单成功，获得一张优惠券【%s】", getCouponNameByLanguage("zh")),
								"en":    fmt.Sprintf("Order placed successfully. You have received a coupon 【%s】.", getCouponNameByLanguage("en")),
								"ja":    fmt.Sprintf("ご注文は正常に完了しました。クーポン【%s】を受け取りました。", getCouponNameByLanguage("ja")),
								"ko":    fmt.Sprintf("주문이 성공적으로 완료되었습니다. 쿠폰 【%s】을 받으셨습니다.", getCouponNameByLanguage("ko")),
								"zh_CN": fmt.Sprintf("下單成功，獲得一張優惠券【%s】", getCouponNameByLanguage("zh_CN")),
							}
							appPushData := g.MapStrStr{
								"type":   "1",
								"string": fmt.Sprintf("%d", memberCouponId1),
							}
							service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
								SystemMessageTitle:   systemMessageTitle1,
								SystemMessageContent: systemMessageContent1,
								Scene:                "system",
								Type:                 "im",
								MemberId:             MemberId,
								Language:             contexts.GetLanguage(ctx),
								AppPushData:          appPushData,
								AppLink:              "/account/coupon-detail",
								WxLink:               fmt.Sprintf("/pages/coupons/details?id=%d", memberCouponId1),
								EnablePush:           true,
								EnableSms:            false,
								PushTitle:            systemMessageTitle1[contexts.GetLanguage(ctx)],
								PushContent:          systemMessageContent1[contexts.GetLanguage(ctx)],
								ShowIndex:            true,
								OperatorId:           0,
								OperatorRole:         "SYSTEM",
							})
						}

					}
				}
				if rewardType == "thcoupon" && !g.IsEmpty(MemberScene.RewardThCouponIds) {
					StartTime = gtime.New(CheckInDate).StartOfDay()
					// 奖励会员优惠券  source：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
					ThCouponIdsArr := strings.Split(MemberScene.RewardThCouponIds, ",")
					for _, thCouponId := range ThCouponIdsArr {
						var memberCouponId1 int64
						if memberCouponId1, couponNameLanguage, err = service.ThCoupon().SendMemberCouponGetId(ctx, &input_th.ThSendMemberCouponInp{
							MemberId:      MemberId,
							CouponId:      gvar.New(thCouponId).Int(),
							StartTime:     StartTime.String(),
							SourceOrderId: OrderId,
						}, 2); err != nil {
						}

						if err == nil {
							// 下单奖励发送消息
							systemMessageContent1 := map[string]string{
								"zh":    fmt.Sprintf("下单成功，获得一张礼品券【%s】", getCouponNameByLanguage("zh")),
								"en":    fmt.Sprintf("Order placed successfully. You have received a gift voucher 【%s】.", getCouponNameByLanguage("en")),
								"ja":    fmt.Sprintf("ご注文が正常に完了しました。ギフト券【%s】を受け取りました。", getCouponNameByLanguage("ja")),
								"ko":    fmt.Sprintf("주문이 성공적으로 완료되었습니다. 상품권 【%s】을(를) 받으셨습니다.", getCouponNameByLanguage("ko")),
								"zh_CN": fmt.Sprintf("下單成功，獲得一張禮品券【%s】", getCouponNameByLanguage("zh_CN")),
							}
							appPushData := g.MapStrStr{
								"type":   "1",
								"string": fmt.Sprintf("%d", memberCouponId1),
							}
							service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
								SystemMessageTitle:   systemMessageTitle1,
								SystemMessageContent: systemMessageContent1,
								Scene:                "system",
								Type:                 "im",
								MemberId:             MemberId,
								Language:             contexts.GetLanguage(ctx),
								AppPushData:          appPushData,
								AppLink:              "/giftDetailPage",
								WxLink:               fmt.Sprintf("/pages/gift-coupon/details?id=%d", memberCouponId1),
								EnablePush:           true,
								EnableSms:            false,
								PushTitle:            systemMessageTitle1[contexts.GetLanguage(ctx)],
								PushContent:          systemMessageContent1[contexts.GetLanguage(ctx)],
								ShowIndex:            true,
								OperatorId:           0,
								OperatorRole:         "SYSTEM",
							})
						}
					}
				}
			}
		}
	}

	return
}

func (s *sHotelService) HotelOrderAwardInvalid(ctx context.Context, OrderId int) (err error) {

	if err = service.BasicsCoupon().InvalidCoupon(ctx, &input_basics.PmsCouponInvalidInp{
		SourceOrderId: OrderId,
	}); err != nil {
		return
	}

	if err = service.ThMemberCoupon().InvalidMemberCoupon(ctx, &input_th.ThInvalidMemberCouponInp{
		SourceOrderId: OrderId,
	}); err != nil {
		return
	}

	return
}
