package admin

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"APT/api/admin/pms"
)

func (c *ControllerPms) CouponTypeList(ctx context.Context, req *pms.CouponTypeListReq) (res *pms.CouponTypeListRes, err error) {
	list, totalCount, err := service.BasicsCouponType().List(ctx, &req.PmsCouponTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsCouponTypeListModel{}
	}

	res = new(pms.CouponTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) CouponTypeAllList(ctx context.Context, req *pms.CouponTypeAllListReq) (res *pms.CouponTypeAllListRes, err error) {
	list, err := service.BasicsCouponType().All(ctx, &req.PmsCouponTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsCouponTypeAllListModel{}
	}

	res = new(pms.CouponTypeAllListRes)
	res.List = list
	return
}
func (c *ControllerPms) CouponTypeView(ctx context.Context, req *pms.CouponTypeViewReq) (res *pms.CouponTypeViewRes, err error) {
	data, err := service.BasicsCouponType().View(ctx, &req.PmsCouponTypeViewInp)
	if err != nil {
		return
	}

	res = new(pms.CouponTypeViewRes)
	res.PmsCouponTypeViewModel = data
	return
}
func (c *ControllerPms) CouponTypeEdit(ctx context.Context, req *pms.CouponTypeEditReq) (res *pms.CouponTypeEditRes, err error) {
	err = service.BasicsCouponType().Edit(ctx, &req.PmsCouponTypeEditInp)
	return
}
func (c *ControllerPms) CouponTypeDelete(ctx context.Context, req *pms.CouponTypeDeleteReq) (res *pms.CouponTypeDeleteRes, err error) {
	err = service.BasicsCouponType().Delete(ctx, &req.PmsCouponTypeDeleteInp)
	return
}
func (c *ControllerPms) CouponTypeMaxSort(ctx context.Context, req *pms.CouponTypeMaxSortReq) (res *pms.CouponTypeMaxSortRes, err error) {
	data, err := service.BasicsCouponType().MaxSort(ctx, &req.PmsCouponTypeMaxSortInp)
	if err != nil {
		return
	}

	res = new(pms.CouponTypeMaxSortRes)
	res.PmsCouponTypeMaxSortModel = data
	return
}
func (c *ControllerPms) CouponTypeStatus(ctx context.Context, req *pms.CouponTypeStatusReq) (res *pms.CouponTypeStatusRes, err error) {
	err = service.BasicsCouponType().Status(ctx, &req.PmsCouponTypeStatusInp)
	return
}
func (c *ControllerPms) SendMemberCoupon(ctx context.Context, req *pms.SendMemberCouponReq) (res *pms.SendMemberCouponRes, err error) {
	// source 来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励
	memberCouponId, couponNameLanguage, err := service.BasicsCouponType().SendMemberCouponGetId(ctx, &req.PmsSendMemberCouponInp, 2)
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
		return "优惠券"
	}

	// 注册奖励发送消息
	systemMessageTitle := map[string]string{
		"zh":    "系统消息",
		"en":    "System Message",
		"ja":    "システムメッセージ",
		"ko":    "시스템 메시지",
		"zh_CN": "系統訊息",
	}
	systemMessageContent := map[string]string{
		"zh":    fmt.Sprintf("系统赠送您一张优惠券【%s】", getCouponNameByLanguage("zh")),
		"en":    fmt.Sprintf("The system will give you a coupon【%s】.", getCouponNameByLanguage("en")),
		"ja":    fmt.Sprintf("システムからクーポン【%s】が提供されます。", getCouponNameByLanguage("ja")),
		"ko":    fmt.Sprintf("시스템에서 쿠폰【%s】을 제공합니다.", getCouponNameByLanguage("ko")),
		"zh_CN": fmt.Sprintf("系統贈送您一張優惠券【%s】", getCouponNameByLanguage("zh_CN")),
	}
	appPushData := g.MapStrStr{
		"type":   "1",
		"string": fmt.Sprintf("%d", memberCouponId),
	}
	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(req.MemberId).Value()
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
	service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle:   systemMessageTitle,
		SystemMessageContent: systemMessageContent,
		Scene:                "system",
		Type:                 "im",
		MemberId:             req.MemberId,
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/account/coupon-detail",
		WxLink:               fmt.Sprintf("/pages/coupons/details?id=%d", memberCouponId),
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
