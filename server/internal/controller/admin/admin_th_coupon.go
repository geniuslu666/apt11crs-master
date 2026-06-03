package admin

import (
	"APT/api/admin/th"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerTh) CouponList(ctx context.Context, req *th.CouponListReq) (res *th.CouponListRes, err error) {
	list, totalCount, err := service.ThCoupon().List(ctx, &req.ThCouponListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThCouponListModel{}
	}

	res = new(th.CouponListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerTh) CouponAllList(ctx context.Context, req *th.CouponAllListReq) (res *th.CouponAllListRes, err error) {
	list, err := service.ThCoupon().All(ctx, &req.ThCouponListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_th.ThCouponAllListModel{}
	}

	res = new(th.CouponAllListRes)
	res.List = list
	return
}
func (c *ControllerTh) CouponView(ctx context.Context, req *th.CouponViewReq) (res *th.CouponViewRes, err error) {
	data, err := service.ThCoupon().View(ctx, &req.ThCouponViewInp)
	if err != nil {
		return
	}

	res = new(th.CouponViewRes)
	res.ThCouponViewModel = data
	return
}
func (c *ControllerTh) CouponEdit(ctx context.Context, req *th.CouponEditReq) (res *th.CouponEditRes, err error) {
	err = service.ThCoupon().Edit(ctx, &req.ThCouponEditInp)
	return
}
func (c *ControllerTh) CouponDelete(ctx context.Context, req *th.CouponDeleteReq) (res *th.CouponDeleteRes, err error) {
	err = service.ThCoupon().Delete(ctx, &req.ThCouponDeleteInp)
	return
}
func (c *ControllerTh) CouponMaxSort(ctx context.Context, req *th.CouponMaxSortReq) (res *th.CouponMaxSortRes, err error) {
	data, err := service.ThCoupon().MaxSort(ctx, &req.ThCouponMaxSortInp)
	if err != nil {
		return
	}

	res = new(th.CouponMaxSortRes)
	res.ThCouponMaxSortModel = data
	return
}
func (c *ControllerTh) CouponStatus(ctx context.Context, req *th.CouponStatusReq) (res *th.CouponStatusRes, err error) {
	err = service.ThCoupon().Status(ctx, &req.ThCouponStatusInp)
	return
}
func (c *ControllerTh) CouponUseStatus(ctx context.Context, req *th.CouponUseStatusReq) (res *th.CouponUseStatusRes, err error) {
	err = service.ThCoupon().UseStatus(ctx, &req.ThCouponUseStatusInp)
	return
}
func (c *ControllerTh) SendMemberCoupon(ctx context.Context, req *th.SendMemberCouponReq) (res *th.SendMemberCouponRes, err error) {
	// source 来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
	memberCouponId, couponNameLanguage, err := service.ThCoupon().SendMemberCouponGetId(ctx, &req.ThSendMemberCouponInp, 1)
	if err != nil {
		return
	}

	// 根据语言获取优惠券名称
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

	// 发送消息
	systemMessageTitle := map[string]string{
		"zh":    "系统消息",
		"en":    "System Message",
		"ja":    "システムメッセージ",
		"ko":    "시스템 메시지",
		"zh_CN": "系統訊息",
	}
	systemMessageContent := map[string]string{
		"zh":    fmt.Sprintf("系统赠送您一张礼品券【%s】", getCouponNameByLanguage("zh")),
		"en":    fmt.Sprintf("The system will give you a gift certificate【%s】.", getCouponNameByLanguage("en")),
		"ja":    fmt.Sprintf("システムからクーポン【%s】が提供されます。", getCouponNameByLanguage("ja")),
		"ko":    fmt.Sprintf("시스템에서 쿠폰【%s】을 제공합니다.", getCouponNameByLanguage("ko")),
		"zh_CN": fmt.Sprintf("系統贈送您一張禮品券【%s】", getCouponNameByLanguage("zh_CN")),
	}
	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(req.ThSendMemberCouponInp.MemberId).Value()
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
		"string": fmt.Sprintf("%d", memberCouponId),
	}
	service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle:   systemMessageTitle,
		SystemMessageContent: systemMessageContent,
		Scene:                "system",
		Type:                 "im",
		MemberId:             req.ThSendMemberCouponInp.MemberId,
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/giftDetailPage",
		WxLink:               fmt.Sprintf("/pages/gift-coupon/details?id=%d", memberCouponId),
		EnablePush:           true,
		EnableSms:            false,
		PushTitle:            systemMessageTitle[memberLanguage],
		PushContent:          systemMessageContent[memberLanguage],
		ShowIndex:            true,
		OperatorId:           int(contexts.GetUserId(ctx)),
		OperatorRole:         "ADMIN",
	})
	return
}
