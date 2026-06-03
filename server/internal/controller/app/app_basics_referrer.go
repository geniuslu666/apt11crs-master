package app

import (
	"APT/api/app/basics"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/service"
	"APT/utility/convert"
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerBasics) ReferrerInfo(ctx context.Context, req *basics.ReferrerInfoReq) (res *basics.ReferrerInfoRes, err error) {
	var (
		MemberInfo      *model.MemberIdentity
		YYConfig        *model.YYConfig
		Referrer        int
		LastReferrer    int
		InviteCount     int
		TodayInvite     int
		YesterdayInvite int
		WeekInvite      int
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	// 获取配置
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}
	if g.IsEmpty(YYConfig) {
		// 系统异常
		err = gerror.New(gi18n.T(ctx, "system_exception"))
		return
	}
	//if YYConfig.RecommendModel == "FIRST" {
	//	// 查询
	//	LastReferrer = MemberInfo.Id
	//} else {
	//	Referrer = MemberInfo.Id
	//}
	// 按终身推荐制来查询
	Referrer = MemberInfo.Id

	// 查询会员推荐总数量
	res = new(basics.ReferrerInfoRes)
	if InviteCount, err = dao.PmsMember.Ctx(ctx).Where(&entity.PmsMember{
		Referrer:     Referrer,
		LastReferrer: LastReferrer,
	}).OmitEmptyWhere().Count(); err != nil {
		return
	}

	res.InviteCount = gvar.New(InviteCount).Int()

	// 根据推荐模式构建查询条件
	var logQuery *gdb.Model
	logQuery = dao.PmsReferrerLog.Ctx(ctx).Where(dao.PmsReferrerLog.Columns().Referrer, MemberInfo.Id)
	//if YYConfig.RecommendModel == "LAST" {
	//	// LAST模式：直接查询PmsReferrerLog表中Referrer字段
	//	logQuery = dao.PmsReferrerLog.Ctx(ctx).Where(dao.PmsReferrerLog.Columns().Referrer, MemberInfo.Id)
	//} else {
	//	// 非LAST模式：需要先查询当前LastReferrer为该用户的所有会员ID，然后查询这些会员的推荐记录
	//	var memberIds []int64
	//	var memberRecords []struct {
	//		Id int64 `json:"id"`
	//	}
	//	if err = dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().Id).
	//		Where(dao.PmsMember.Columns().LastReferrer, MemberInfo.Id).Scan(&memberRecords); err != nil {
	//		return
	//	}
	//
	//	// 提取ID到切片中
	//	for _, record := range memberRecords {
	//		memberIds = append(memberIds, record.Id)
	//	}
	//	if len(memberIds) == 0 {
	//		// 如果没有被推荐的会员，直接返回空结果
	//		res.TodayInvite = 0
	//		res.YesterdayInvite = 0
	//		res.WeekInvite = 0
	//		res.List = []*struct {
	//			*entity.PmsReferrerLog
	//			MemberInfo *struct {
	//				g.Meta `orm:"table:hg_pms_member"`
	//				*entity.PmsMember
	//			} `json:"memberInfo" orm:"with:id=memberId"`
	//		}{}
	//		res.Count = 0
	//		return
	//	}
	//	logQuery = dao.PmsReferrerLog.Ctx(ctx).WhereIn(dao.PmsReferrerLog.Columns().MemberId, memberIds).Where(dao.PmsReferrerLog.Columns().LastReferrer, MemberInfo.Id)
	//}

	// 今日邀请
	if TodayInvite, err = logQuery.Clone().
		Where("DATE_FORMAT(created_at,'%Y-%m-%d') = ?", gtime.Now().Format("Y-m-d")).
		Count(); err != nil {
		return
	}
	res.TodayInvite = gvar.New(TodayInvite).Int()

	// 昨日邀请
	if YesterdayInvite, err = logQuery.Clone().
		Where("DATE_FORMAT(created_at,'%Y-%m-%d') = ?", gtime.Now().Add(-1*24*time.Hour).Format("Y-m-d")).
		Count(); err != nil {
		return
	}
	res.YesterdayInvite = gvar.New(YesterdayInvite).Int()

	// 本周邀请
	if WeekInvite, err = logQuery.Clone().
		WhereBetween("created_at", gtime.New(convert.WeekStartDate()), gtime.Now()).
		Count(); err != nil {
		return
	}
	res.WeekInvite = gvar.New(WeekInvite).Int()

	// 查询邀请列表
	if err = logQuery.WithAll().Page(req.Page, req.PerPage).ScanAndCount(&res.List, &res.Count, false); err != nil {
		return
	}

	return
}

func (c *ControllerBasics) ReferrerList(ctx context.Context, req *basics.ReferrerListReq) (res *basics.ReferrerListRes, err error) {
	var (
		MemberInfo   *model.MemberIdentity
		YYConfig     *model.YYConfig
		Referrer     int
		LastReferrer int
	)
	MemberInfo = contexts.GetMemberUser(ctx)
	// 获取配置
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}
	if g.IsEmpty(YYConfig) {
		err = gerror.New(gi18n.T(ctx, "system_exception"))
		return
	}
	if YYConfig.RecommendModel == "FIRST" {
		// 查询
		LastReferrer = MemberInfo.Id
	} else {
		Referrer = MemberInfo.Id
	}

	// 查询会员推荐总数量
	res = new(basics.ReferrerListRes)
	// 查询邀请列表
	if err = dao.PmsReferrerLog.Ctx(ctx).
		Where(&entity.PmsReferrerLog{Referrer: Referrer, LastReferrer: LastReferrer}).WithAll().
		OmitEmptyWhere().Page(req.Page, req.PerPage).ScanAndCount(&res.List, &res.Count, false); err != nil {
		return
	}

	return
}
