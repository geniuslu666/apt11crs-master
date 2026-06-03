package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"APT/api/app/basics"
)

func (c *ControllerBasics) ShareInfo(ctx context.Context, req *basics.ShareInfoReq) (res *basics.ShareInfoRes, err error) {
	var (
		MemberInfo    = contexts.GetMemberUser(ctx)
		PmsProperty   entity.PmsProperty
		WXShareConfig *model.WXShareConfig
	)
	if WXShareConfig, err = service.BasicsConfig().GetWXShareConfig(ctx); err != nil {
		return
	}
	res = new(basics.ShareInfoRes)
	switch contexts.GetLanguage(ctx) {
	case consts.Zh:
		res.Title = WXShareConfig.TitleZh
		res.Content = WXShareConfig.ContentZh
		break
	case consts.En:
		res.Title = WXShareConfig.TitleEn
		res.Content = WXShareConfig.ContentEn
		break
	case consts.Ko:
		res.Title = WXShareConfig.TitleKo
		res.Content = WXShareConfig.ContentKo
		break
	case consts.Ja:
		res.Title = WXShareConfig.TitleJa
		res.Content = WXShareConfig.ContentJa
		break
	case consts.ZhCN:
		res.Title = WXShareConfig.TitleZhCn
		res.Content = WXShareConfig.ContentZhCn
		break
	default:
		// 语言格式错误
		err = gerror.New(gi18n.T(ctx, "language_format_error"))
		return
	}
	res.Icon = WXShareConfig.Icon
	res.Url = WXShareConfig.ShareUrl
	res.ShareDomain = WXShareConfig.ShareDomain
	if req.ShareType == "member" {
		res.Url = fmt.Sprintf("%s?referrer=%d", res.Url, MemberInfo.Id)
	} else if req.ShareType == "host" {
		if err = dao.PmsProperty.Ctx(ctx).
			Where(dao.PmsProperty.Columns().Id, req.Pid).
			Hook(hook.PmsFindLanguageValueHook).
			Scan(&PmsProperty); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(PmsProperty) {
			g.Log().Info(ctx, "物业不存在", err)
			return
		}
		res.Url = fmt.Sprintf("%s?referrer=%d&Property=%s", res.Url, MemberInfo.Id, PmsProperty.Uid)
	}
	return
}

func (c *ControllerBasics) ShareInvitePageInfo(ctx context.Context, req *basics.ShareInvitePageInfoReq) (res *basics.ShareInvitePageInfoRes, err error) {
	var (
		MemberInfo   *model.MemberIdentity
		ChannelInfo  *entity.PmsChannel
		StaffInfo    *entity.PmsStaff
		YYConfig     *model.YYConfig
		Referrer     int
		LastReferrer int
		InviteCount  int
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	res = new(basics.ShareInvitePageInfoRes)

	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}

	if MemberInfo.RebateMode == "CHANNEL" {
		if err = dao.PmsChannel.Ctx(ctx).
			Where(dao.PmsChannel.Columns().Id, MemberInfo.ChannelId).
			Scan(&ChannelInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(ChannelInfo) {
			g.Log().Info(ctx, "渠道不存在", err)
			return
		}
		res.AwardCommission = ChannelInfo.Rate
		res.GetCommission = ChannelInfo.AllBalance
		res.WithdrawPrice = ChannelInfo.Balance
	} else if MemberInfo.RebateMode == "STAFF" {
		if err = dao.PmsStaff.Ctx(ctx).
			Where(dao.PmsStaff.Columns().Id, MemberInfo.StaffId).
			Scan(&StaffInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}
		if g.IsEmpty(StaffInfo) {
			g.Log().Info(ctx, "员工不存在", err)
			return
		}
		res.AwardCommission = StaffInfo.Rate
		res.GetCommission = StaffInfo.AllBalance
		res.WithdrawPrice = StaffInfo.Balance
	} else {
		// 会员
		res.AwardCommission = YYConfig.MemberBrokerageRate
		res.GetCommission = 0
		res.WithdrawPrice = 0
	}

	// 获取配置
	//if YYConfig.RecommendModel == "FIRST" {
	//	// 查询
	//	LastReferrer = MemberInfo.Id
	//} else {
	//	Referrer = MemberInfo.Id
	//}
	// 按终身推荐制来查询
	Referrer = MemberInfo.Id

	// 查询会员推荐总数量
	InviteCount, err = dao.PmsMember.Ctx(ctx).Where(&entity.PmsMember{
		Referrer:     Referrer,
		LastReferrer: LastReferrer,
	}).OmitEmptyWhere().Count()

	if err == nil {
		res.SuccessMember = gvar.New(InviteCount).Int()
	}
	res.AwardScore = 0

	// 新人注册奖励
	var MemberRegRewardConfig *model.MemberRegRewardConfig
	MemberRegRewardConfig, err = service.BasicsConfig().GetMemberRegRewardConfig(ctx)
	if err == nil {
		if MemberRegRewardConfig.IsOpen == 1 {
			// 奖励类型
			RewardTypeArr := strings.Split(MemberRegRewardConfig.RewardType, ",")
			for _, rewardType := range RewardTypeArr {
				if rewardType == "balance" && MemberRegRewardConfig.RewardBalance > 0 {
					res.AwardScore = MemberRegRewardConfig.RewardBalance + res.AwardScore
				}
			}
		}
	}

	var InviteNewRewardConfig *model.InviteNewRewardConfig
	InviteNewRewardConfig, err = service.BasicsConfig().GetInviteNewRewardConfig(ctx)
	if err == nil {
		// 被邀请人奖励
		if InviteNewRewardConfig.IsInviteRegOpen == 1 {
			if InviteNewRewardConfig.RewardType == "balance" && InviteNewRewardConfig.RewardBalance > 0 {
				res.AwardScore = InviteNewRewardConfig.RewardBalance + res.AwardScore
			}
		}
	}

	return
}
