package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type TranslateTextRes struct {
	Zh   string `json:"zh" dc:"简体中文翻译结果"`
	ZhCN string `json:"zh_CN" dc:"繁体翻译结果"`
	Ko   string `json:"ko" dc:"韩文翻译结果"`
	Ja   string `json:"ja" dc:"日文翻译结果"`
	En   string `json:"en" dc:"英文翻译结果"`
}

func (s *sAppMember) RegisterMemberAward(ctx context.Context, TX gdb.TX, MemberId int) (err error) {
	var (
		MemberInfo            *entity.PmsMember
		RecommendedMemberInfo *entity.PmsMember
		Recommended           int
		MemberRegRewardConfig *model.MemberRegRewardConfig
		InviteNewRewardConfig *model.InviteNewRewardConfig
		YYConfig              *model.YYConfig
		GIFTContent           string
		couponType            []*entity.PmsCouponType
		// WaitGroup             sync.WaitGroup
		// translateNitify       *TranslateTextRes
	)

	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).TX(TX).WherePri(MemberId).Scan(&MemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(MemberInfo) {
		return
	}

	// 新人注册奖励
	if MemberRegRewardConfig, err = service.BasicsConfig().GetMemberRegRewardConfig(ctx); err != nil {
		return
	}
	if MemberRegRewardConfig.IsOpen == 1 {
		// 奖励类型
		RewardTypeArr := strings.Split(MemberRegRewardConfig.RewardType, ",")

		// 开始新人礼奖励
		for _, rewardType := range RewardTypeArr {
			systemMessageTitle := map[string]string{
				"zh":    "注册奖励",
				"en":    "Registration Rewards",
				"ja":    "登録特典",
				"ko":    "등록 보상",
				"zh_CN": "註冊獎勵",
			}
			if rewardType == "balance" && MemberRegRewardConfig.RewardBalance > 0 {
				// 奖励积分
				if err = s.MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
					MemberId:      MemberId,
					Scene:         "SYSTEM",
					Type:          "AWARD",
					ChangeBalance: MemberRegRewardConfig.RewardBalance,
					OrderSn:       "",
					Reason:        "注册奖励",
					MdCode:        MemberInfo.RegisterMdCode,
					MpModel:       MemberInfo.RegisterMpModel,
				}, TX); err != nil {
					continue
				}
				GIFTContent += fmt.Sprintf("%d积分,", gvar.New(MemberRegRewardConfig.RewardBalance).Int())

				// 注册奖励发送消息
				systemMessageContent := map[string]string{
					"zh":    fmt.Sprintf("注册成功，奖励%d积分", gvar.New(MemberRegRewardConfig.RewardBalance).Int()),
					"en":    fmt.Sprintf("Registration successful, %d points awarded.", gvar.New(MemberRegRewardConfig.RewardBalance).Int()),
					"ja":    fmt.Sprintf("登録に成功しました。%dポイント付与されます。", gvar.New(MemberRegRewardConfig.RewardBalance).Int()),
					"ko":    fmt.Sprintf("등록이 성공적으로 완료되어 %d포인트가 지급되었습니다.", gvar.New(MemberRegRewardConfig.RewardBalance).Int()),
					"zh_CN": fmt.Sprintf("註冊成功，獎勵%d積分", gvar.New(MemberRegRewardConfig.RewardBalance).Int()),
				}
				appPushData := g.MapStrStr{
					"type": "0",
				}

				service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
					SystemMessageTitle:   systemMessageTitle,
					SystemMessageContent: systemMessageContent,
					Scene:                "system",
					Type:                 "im",
					MemberId:             MemberId,
					Language:             contexts.GetLanguage(ctx),
					AppPushData:          appPushData,
					AppLink:              "/account/score",
					WxLink:               "/pages/user/points",
					EnablePush:           true,
					EnableSms:            false,
					PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
					PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
					ShowIndex:            true,
					OperatorId:           0,
					OperatorRole:         "SYSTEM",
				})
			}
			if rewardType == "coupon" && !g.IsEmpty(MemberRegRewardConfig.CouponTypeIds) {
				// 奖励会员优惠券 source：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
				CouponTypeIdsArr := strings.Split(MemberRegRewardConfig.CouponTypeIds, ",")
				if err = dao.PmsCouponType.Ctx(ctx).
					Fields("id,max(coupon_name) as coupon_name,count(id) as count").
					Where(dao.PmsCouponType.Columns().Id, MemberRegRewardConfig.CouponTypeIds).
					Group(dao.PmsCouponType.Columns().Id).
					Hook(hook.PmsFindLanguageValueHook).
					Scan(&couponType); err != nil {
					continue
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

				for _, couponTypeId := range CouponTypeIdsArr {
					var memberCouponId int64
					if memberCouponId, couponNameLanguage, err = service.BasicsCouponType().SendMemberCouponGetId(ctx, &input_basics.PmsSendMemberCouponInp{
						MemberId:     MemberId,
						CouponTypeId: gvar.New(couponTypeId).Int(),
					}, 3); err != nil {
						continue
					}

					if err == nil {
						// 注册奖励发送消息
						systemMessageContent := map[string]string{
							"zh":    fmt.Sprintf("注册成功，获得一张优惠券【%s】", getCouponNameByLanguage("zh")),
							"en":    fmt.Sprintf("Registration successful, a coupon received.【%s】", getCouponNameByLanguage("en")),
							"ja":    fmt.Sprintf("登録に成功しました。クーポンを受領しました。【%s】", getCouponNameByLanguage("ja")),
							"ko":    fmt.Sprintf("등록이 성공적으로 완료되어 쿠폰이 지급되었습니다。【%s】", getCouponNameByLanguage("ko")),
							"zh_CN": fmt.Sprintf("註冊成功，獲得一張優惠券【%s】", getCouponNameByLanguage("zh_CN")),
						}
						appPushData := g.MapStrStr{
							"type":   "1",
							"string": fmt.Sprintf("%d", memberCouponId),
						}
						service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
							SystemMessageTitle:   systemMessageTitle,
							SystemMessageContent: systemMessageContent,
							Scene:                "system",
							Type:                 "im",
							MemberId:             MemberId,
							Language:             contexts.GetLanguage(ctx),
							AppPushData:          appPushData,
							AppLink:              "/account/coupon-detail",
							WxLink:               fmt.Sprintf("/pages/coupons/details?id=%d", memberCouponId),
							EnablePush:           true,
							EnableSms:            false,
							PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
							PushContent:          systemMessageContent[contexts.GetLanguage(ctx)],
							ShowIndex:            true,
							OperatorId:           0,
							OperatorRole:         "SYSTEM",
						})
					}
				}
				for _, v := range couponType {
					GIFTContent += fmt.Sprintf("%sx%d,", v.CouponName, v.Count)
				}
			}
			if rewardType == "thcoupon" && !g.IsEmpty(MemberRegRewardConfig.ThCouponIds) {
				//StartTime = gtime.New(CheckInDate).StartOfDay()
				// 奖励会员礼品券  source：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
				ThCouponIdsArr := strings.Split(MemberRegRewardConfig.ThCouponIds, ",")

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
					return "礼品券"
				}

				for _, thCouponId := range ThCouponIdsArr {
					var memberCouponId1 int64
					if memberCouponId1, couponNameLanguage, err = service.ThCoupon().SendMemberCouponGetId(ctx, &input_th.ThSendMemberCouponInp{
						MemberId:  MemberId,
						CouponId:  gvar.New(thCouponId).Int(),
						StartTime: gtime.Now().String(),
					}, 3); err != nil {
					}

					if err == nil {
						// 注册奖励发送消息
						systemMessageContent1 := map[string]string{
							"zh":    fmt.Sprintf("注册成功，获得一张礼品券【%s】", getCouponNameByLanguage("zh")),
							"en":    fmt.Sprintf("Registration successful, a gift voucher received.【%s】", getCouponNameByLanguage("en")),
							"ja":    fmt.Sprintf("登録に成功しました。ギフト券を受領しました。【%s】", getCouponNameByLanguage("ja")),
							"ko":    fmt.Sprintf("등록이 성공적으로 완료되어 상품권 지급되었습니다。【%s】", getCouponNameByLanguage("ko")),
							"zh_CN": fmt.Sprintf("註冊成功，獲得一張禮品券【%s】", getCouponNameByLanguage("zh_CN")),
						}
						appPushData := g.MapStrStr{
							"type":   "1",
							"string": fmt.Sprintf("%d", memberCouponId1),
						}
						service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
							SystemMessageTitle:   systemMessageTitle,
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
							PushTitle:            systemMessageTitle[contexts.GetLanguage(ctx)],
							PushContent:          systemMessageContent1[contexts.GetLanguage(ctx)],
							ShowIndex:            true,
							OperatorId:           0,
							OperatorRole:         "SYSTEM",
						})
					}
				}
			}
		}
		// translateNitify = new(TranslateTextRes)
		// WaitGroup.Add(5)
		// go func() {
		// 	translateNitify.Zh, err = translate.TranslateText(ctx, GIFTContent, "zh")
		// 	WaitGroup.Done()
		// }()
		// go func() {
		// 	translateNitify.ZhCN, err = translate.TranslateText(ctx, GIFTContent, "zh_TW")
		// 	WaitGroup.Done()
		// }()
		// go func() {
		// 	translateNitify.Ko, err = translate.TranslateText(ctx, GIFTContent, "ko")
		// 	WaitGroup.Done()
		// }()
		// go func() {
		// 	translateNitify.Ja, err = translate.TranslateText(ctx, GIFTContent, "ja")
		// 	WaitGroup.Done()
		// }()
		// go func() {
		// 	translateNitify.En, err = translate.TranslateText(ctx, GIFTContent, "en")
		// 	WaitGroup.Done()
		// }()
		// WaitGroup.Wait()
		// PmsLanguage := new(entity.PmsLanguage)
		// if !g.IsEmpty(GIFTContent) {
		// 	if err = dao.PmsLanguage.Ctx(ctx).Where(g.Map{
		// 		dao.PmsLanguage.Columns().Language: contexts.GetLanguage(ctx),
		// 		dao.PmsLanguage.Columns().Key:      "WELCOME",
		// 	}).Scan(&PmsLanguage); err != nil {
		// 		err = gerror.New("获取语言失败，请稍后重试！")
		// 		return
		// 	}
		// 	err = service.BasicsAppNotify().NotifyEdit(ctx, &input_basics.PmsNotifyEditInp{
		// 		PmsNotify: entity.PmsNotify{
		// 			MemberId:      MemberId,
		// 			NotifyTitle:   "SYSTEM",
		// 			NotifyType:    "GIFT", // 新人礼消息
		// 			NotifyContent: PmsLanguage.Content,
		// 			NotifyData:    gjson.New(translateNitify),
		// 		},
		// 	})
		// }
	}
	// 获取推荐制度
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}
	// ===邀请注册奖励===
	// 推荐模式
	if YYConfig.RecommendModel == "FIRST" {
		Recommended = MemberInfo.LastReferrer
	} else if YYConfig.RecommendModel == "LAST" {
		Recommended = MemberInfo.Referrer
	} else {
		return
	}
	// 查询推荐人信息
	if err = dao.PmsMember.Ctx(ctx).TX(TX).WherePri(Recommended).Scan(&RecommendedMemberInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(RecommendedMemberInfo) {
		return
	}

	// 邀请注册配置
	if InviteNewRewardConfig, err = service.BasicsConfig().GetInviteNewRewardConfig(ctx); err != nil {
		return
	}

	// 被邀请人奖励
	if InviteNewRewardConfig.IsInviteRegOpen == 1 {
		// 奖励类型
		RewardTypeArr := strings.Split(InviteNewRewardConfig.RewardType, ",")

		for _, rewardType := range RewardTypeArr {
			if rewardType == "balance" && InviteNewRewardConfig.RewardBalance > 0 {
				// 奖励积分
				if err = s.MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
					MemberId:      MemberId,
					Scene:         "SYSTEM",
					Type:          "AWARD",
					ChangeBalance: InviteNewRewardConfig.RewardBalance,
					OrderSn:       "",
					Reason:        "被邀请注册奖励",
				}, TX); err != nil {
					return
				}

				// 注册奖励发送消息
				systemMessageTitle1 := map[string]string{
					"zh":    "邀请奖励",
					"en":    "Invitation Rewards",
					"ja":    "招待特典",
					"ko":    "초대 보상",
					"zh_CN": "邀請獎勵",
				}
				systemMessageContent1 := map[string]string{
					"zh":    fmt.Sprintf("被邀请注册成功，奖励%d积分", gvar.New(InviteNewRewardConfig.RewardBalance).Int()),
					"en":    fmt.Sprintf("Successful registration by invitation will earn you %d points.", gvar.New(InviteNewRewardConfig.RewardBalance).Int()),
					"ja":    fmt.Sprintf("招待による登録に成功すると、%d ポイントを獲得できます。", gvar.New(InviteNewRewardConfig.RewardBalance).Int()),
					"ko":    fmt.Sprintf("초대를 통해 성공적으로 등록하면 %d포인트를 획득할 수 있습니다.", gvar.New(InviteNewRewardConfig.RewardBalance).Int()),
					"zh_CN": fmt.Sprintf("被邀請註冊成功，獎勵%d積分", gvar.New(InviteNewRewardConfig.RewardBalance).Int()),
				}
				appPushData1 := g.MapStrStr{
					"type": "0",
				}

				service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
					SystemMessageTitle:   systemMessageTitle1,
					SystemMessageContent: systemMessageContent1,
					Scene:                "system",
					Type:                 "im",
					MemberId:             MemberId,
					Language:             contexts.GetLanguage(ctx),
					AppPushData:          appPushData1,
					AppLink:              "/account/score",
					WxLink:               "/pages/user/points",
					EnablePush:           true,
					EnableSms:            false,
					PushTitle:            systemMessageTitle1[contexts.GetLanguage(ctx)],
					PushContent:          systemMessageContent1[contexts.GetLanguage(ctx)],
					ShowIndex:            true,
					OperatorId:           0,
					OperatorRole:         "SYSTEM",
				})

			}
			if rewardType == "coupon" && !g.IsEmpty(InviteNewRewardConfig.CouponTypeIds) {
				// 奖励会员优惠券  source：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
				CouponTypeIdsArr := strings.Split(InviteNewRewardConfig.CouponTypeIds, ",")

				systemMessageTitle1 := map[string]string{
					"zh":    "邀请奖励",
					"en":    "Invitation Rewards",
					"ja":    "招待特典",
					"ko":    "초대 보상",
					"zh_CN": "邀請獎勵",
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

				for _, couponTypeId := range CouponTypeIdsArr {
					var memberCouponId1 int64
					if memberCouponId1, couponNameLanguage, err = service.BasicsCouponType().SendMemberCouponGetId(ctx, &input_basics.PmsSendMemberCouponInp{
						MemberId:     MemberId,
						CouponTypeId: gvar.New(couponTypeId).Int(),
					}, 4); err != nil {
						return
					}

					if err == nil {
						// 注册奖励发送消息
						systemMessageContent1 := map[string]string{
							"zh":    fmt.Sprintf("被邀请成功，获得一张优惠券【%s】", getCouponNameByLanguage("zh")),
							"en":    fmt.Sprintf("Successfully invited and received a coupon【%s】", getCouponNameByLanguage("en")),
							"ja":    fmt.Sprintf("招待されクーポンを受け取りました【%s】", getCouponNameByLanguage("ja")),
							"ko":    fmt.Sprintf("성공적으로 초대를 받고 쿠폰을 받았습니다【%s】", getCouponNameByLanguage("ko")),
							"zh_CN": fmt.Sprintf("被邀請成功，獲得一張優惠券【%s】", getCouponNameByLanguage("zh_CN")),
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
		}
	}

	return
}
