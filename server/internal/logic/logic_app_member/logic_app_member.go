package logic_app_member

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sAppMember) MemberList(ctx context.Context, in *input_app_member.PmsMemberListInp) (list []*input_app_member.PmsMemberListModel, totalCount int, err error) {
	mod := dao.PmsMember.Ctx(ctx).WithAll()

	mod = mod.Fields(input_app_member.PmsMemberListModel{})

	if in.Id > 0 {
		if in.IsJunior == 1 {
			var YYConfig *model.YYConfig
			// 获取推荐制度
			if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
				return
			}

			// ===邀请注册奖励===
			// 推荐模式
			if YYConfig.RecommendModel == "FIRST" {
				// 最后推荐制
				mod = mod.Where(dao.PmsMember.Columns().LastReferrer, in.Id)
			} else if YYConfig.RecommendModel == "LAST" {
				// 终身推荐制
				mod = mod.Where(dao.PmsMember.Columns().Referrer, in.Id)
			}
		} else {
			mod = mod.Where(dao.PmsMember.Columns().Id, in.Id)
		}
	}

	if in.FullName != "" {
		mod = mod.WhereLike(dao.PmsMember.Columns().FullName, "%"+in.FullName+"%")
	}

	if in.MemberNo != "" {
		mod = mod.WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.MemberNo+"%")
	}

	if in.Phone != "" {
		mod = mod.WhereLike(dao.PmsMember.Columns().Phone, "%"+in.Phone+"%")
	}

	if in.Mail != "" {
		mod = mod.WhereLike(dao.PmsMember.Columns().Mail, "%"+in.Mail+"%")
	}

	if in.Level > 0 {
		mod = mod.Where(dao.PmsMember.Columns().Level, in.Level)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	if len(in.LastLogin) == 2 {
		mod = mod.WhereBetween(dao.PmsMember.Columns().LastLogin, in.LastLogin[0], in.LastLogin[1])
	}

	if !g.IsEmpty(in.MemberIds) {
		mod = mod.WhereIn(dao.PmsMember.Columns().Id, strings.Split(in.MemberIds, ","))
	}

	if !g.IsEmpty(in.Status) {
		mod = mod.Where(dao.PmsMember.Columns().Status, in.Status)
	}

	if !g.IsEmpty(in.RebateMode) {
		mod = mod.Where(dao.PmsMember.Columns().RebateMode, in.RebateMode)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	if !g.IsEmpty(in.Sort) {
		if in.Sort == "ascend" {
			mod = mod.Order(dao.PmsMember.Columns().Balance, "asc")
		} else if in.Sort == "descend" {
			mod = mod.Order(dao.PmsMember.Columns().Balance, "desc")
		} else {
			mod = mod.OrderDesc(dao.PmsMember.Columns().Id)
		}
	} else {
		mod = mod.OrderDesc(dao.PmsMember.Columns().Id)
	}

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取会员信息列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取会员信息列表失败，请稍后重试！")
			return
		}

		var (
			memberList []*input_app_member.PmsMemberListModel
		)

		if !g.IsEmpty(in.MemberIds) {
			memberIdArr := strings.Split(in.MemberIds, ",")
			for _, memberId := range memberIdArr {
				for _, v := range list {
					if v.Id == gvar.New(memberId).Int() {
						memberList = append(memberList, v)
					}
				}
			}
			list = memberList
		}
	}

	return
}

func (s *sAppMember) MemberSelectList(ctx context.Context, in *input_app_member.PmsMemberSelectInp) (list []*input_app_member.PmsMemberListModel, err error) {
	mod := dao.PmsMember.Ctx(ctx)

	mod = mod.Fields(input_app_member.PmsMemberListModel{})

	mod = mod.Where(dao.PmsMember.Columns().RebateMode, "MEMBER")

	if in.Keyword != "" {
		mod = mod.Where(mod.Builder().
			WhereLike(dao.PmsMember.Columns().MemberNo, "%"+in.Keyword+"%").
			WhereOrLike(dao.PmsMember.Columns().FullName, "%"+in.Keyword+"%").
			WhereOrLike(dao.PmsMember.Columns().Phone, "%"+in.Keyword+"%"))
	}

	mod = mod.OrderDesc(dao.PmsMember.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取会员信息列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberExport(ctx context.Context, in *input_app_member.PmsMemberListInp) (err error) {
	var exportList []*input_app_member.PmsMemberExportModel

	in.Pagination = false
	list, _, err := s.MemberList(ctx, in)
	if err != nil {
		return
	}

	var YYConfig *model.YYConfig
	// 获取推荐制度
	if YYConfig, err = service.BasicsConfig().GetYYConfig(ctx); err != nil {
		return
	}

	for _, v := range list {
		var referInfo string

		// 推荐模式
		if YYConfig.RecommendModel == "FIRST" {
			// 最后推荐制
			if !g.IsEmpty(v.LastReferrerDetail) {
				if v.LastReferrerDetail.RebateMode == "CHANNEL" {
					referInfo = "渠道：" + v.LastReferrerDetail.MemberNo
				} else if v.LastReferrerDetail.RebateMode == "STAFF" {
					referInfo = "员工：" + v.LastReferrerDetail.MemberNo
				} else {
					referInfo = "会员：" + v.LastReferrerDetail.MemberNo
				}
			}

		} else if YYConfig.RecommendModel == "LAST" {
			// 终身推荐制
			if !g.IsEmpty(v.ReferrerDetail) {
				if v.ReferrerDetail.RebateMode == "CHANNEL" {
					referInfo = "渠道：" + v.ReferrerDetail.MemberNo
				} else if v.ReferrerDetail.RebateMode == "STAFF" {
					referInfo = "员工：" + v.ReferrerDetail.MemberNo
				} else {
					referInfo = "会员：" + v.ReferrerDetail.MemberNo
				}
			}
		}

		exportList = append(exportList, &input_app_member.PmsMemberExportModel{
			Id:           v.Id,
			MemberNo:     v.MemberNo,
			FullName:     v.FullName,
			ReferInfo:    referInfo,
			Phone:        v.PhoneArea + v.Phone,
			Mail:         v.Mail,
			RegisterTime: v.RegisterTime,
		})
	}

	tags, err := convert.GetEntityDescTags(input_app_member.PmsMemberExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出会员信息-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("会员信息")
		exports   []input_app_member.PmsMemberExportModel
	)

	if err = gconv.Scan(exportList, &exports); err != nil {
		return
	}
	// TODO FIX
	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sAppMember) MemberEdit(ctx context.Context, in *input_app_member.PmsMemberEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = dao.PmsMember.Ctx(ctx).
				Fields(input_app_member.PmsMemberUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员信息失败，请稍后重试！")
			}
			return
		}

		if _, err = dao.PmsMember.Ctx(ctx).
			Fields(input_app_member.PmsMemberInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员信息失败，请稍后重试！")
		}
		return
	})
}

func (s *sAppMember) MemberBaseEdit(ctx context.Context, in *input_app_member.PmsMemberBaseEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		saveData := g.Map{}
		if in.Type == "level" {
			saveData = g.Map{
				dao.PmsMember.Columns().Level: in.Level,
			}
		} else if in.Type == "group" {
			saveData = g.Map{
				dao.PmsMember.Columns().GroupId: in.GroupId,
			}
		} else if in.Type == "firstName" {
			saveData = g.Map{
				dao.PmsMember.Columns().FirstName: in.FirstName,
				dao.PmsMember.Columns().FullName:  fmt.Sprintf("%s %s", in.FirstName, in.LastName),
			}
		} else if in.Type == "lastName" {
			saveData = g.Map{
				dao.PmsMember.Columns().LastName: in.LastName,
				dao.PmsMember.Columns().FullName: fmt.Sprintf("%s %s", in.FirstName, in.LastName),
			}
		} else if in.Type == "phone" {
			saveData = g.Map{
				dao.PmsMember.Columns().PhoneArea: in.PhoneArea,
				dao.PmsMember.Columns().Phone:     in.Phone,
			}
		} else if in.Type == "mail" {
			saveData = g.Map{
				dao.PmsMember.Columns().Mail: in.Mail,
			}
		}

		if in.Id > 0 {
			if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Data(saveData).Update(); err != nil {
				err = gerror.Wrap(err, "修改会员信息失败，可能是存在相同会员信息项！")
			}
			return
		}
		return
	})
}

func (s *sAppMember) MemberBalanceEdit(ctx context.Context, in *input_app_member.PmsMemberBalanceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		memberBalance, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().Balance).WherePri(in.Id).Value()
		if (in.Value + memberBalance.Float64()) < 0 {
			err = gerror.Wrap(err, "调整值与当前积分数相加不能小于0！")
			return
		}

		if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Increment(dao.PmsMember.Columns().Balance, in.Value); err != nil {
			err = gerror.Wrap(err, "调整会员积分失败，请稍后重试！")
			return
		}

		if _, err = dao.PmsBalanceChange.Ctx(ctx).OmitEmptyData().Insert(entity.PmsBalanceChange{
			MemberId:    in.Id,
			ChangePrice: in.Value,
			Scene:       "SYSTEM",
			Type:        "SYS",
			OperatorId:  int(contexts.GetUserId(ctx)),
			Des:         in.Des,
			Reason:      in.Reason,
			CreatedAt:   gtime.Now(),
			UpdatedAt:   gtime.Now(),
		}); err != nil {
			return
		}

		// 发送消息
		systemMessageTitle := map[string]string{
			"zh":    "系统消息",
			"en":    "System Message",
			"ja":    "システムメッセージ",
			"ko":    "시스템 메시지",
			"zh_CN": "系統訊息",
		}
		var systemMessageContent map[string]string
		if in.Value > 0 {
			systemMessageContent = map[string]string{
				"zh":    fmt.Sprintf("后台为您新增%d积分", gvar.New(in.Value).Int()),
				"en":    fmt.Sprintf("The system has added %d points to your account.", gvar.New(in.Value).Int()),
				"ja":    fmt.Sprintf("バックエンドで%dポイントを受け取ります。", gvar.New(in.Value).Int()),
				"ko":    fmt.Sprintf("시스템에서 귀하의 계정에 %d포인트를 추가했습니다.", gvar.New(in.Value).Int()),
				"zh_CN": fmt.Sprintf("後台為您新增%d積分", gvar.New(in.Value).Int()),
			}
		} else {
			systemMessageContent = map[string]string{
				"zh":    fmt.Sprintf("后台为您减少%d积分", gvar.New(in.Value).Int()),
				"en":    fmt.Sprintf("The system has deducted %d points from your account.", gvar.New(in.Value).Int()),
				"ja":    fmt.Sprintf("システムによりあなたのアカウントから %d ポイントが差し引かれました。", gvar.New(in.Value).Int()),
				"ko":    fmt.Sprintf("시스템에서 귀하의 계정에서 %d포인트가 차감되었습니다.", gvar.New(in.Value).Int()),
				"zh_CN": fmt.Sprintf("系統為您減少%d積分", gvar.New(in.Value).Int()),
			}
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(in.Id).Value()
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
			"type": "0",
		}

		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "system",
			Type:                 "im",
			MemberId:             in.Id,
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/account/score",
			WxLink:               "/pages/user/points",
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			ShowIndex:            true,
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
		})

		return
	})
}

func (s *sAppMember) MemberExpEdit(ctx context.Context, in *input_app_member.PmsMemberExpEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		memberExp, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().Exp).WherePri(in.Id).Value()
		if (in.Value + memberExp.Float64()) < 0 {
			err = gerror.Wrap(err, "调整值与当前成长值数相加不能小于0！")
			return
		}

		if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Increment(dao.PmsMember.Columns().Exp, in.Value); err != nil {
			err = gerror.Wrap(err, "调整会员成长值失败，请稍后重试！")
			return
		}

		if _, err = dao.PmsExpChange.Ctx(ctx).OmitEmptyData().Insert(entity.PmsExpChange{
			MemberId:   in.Id,
			Exp:        in.Value,
			Scene:      "SYSTEM",
			OperatorId: int(contexts.GetUserId(ctx)),
			Des:        in.Des,
			CreatedAt:  gtime.Now(),
			UpdatedAt:  gtime.Now(),
		}); err != nil {
			return
		}

		var (
			PmsMember      *entity.PmsMember
			PmsMemberLevel *entity.PmsMemberLevel
		)
		dao.PmsMember.Ctx(ctx).WherePri(in.Id).Scan(&PmsMember)
		dao.PmsMemberLevel.Ctx(ctx).
			WhereLTE(dao.PmsMemberLevel.Columns().Exp, PmsMember.Exp).
			OrderDesc(dao.PmsMemberLevel.Columns().Exp).
			Scan(&PmsMemberLevel)
		if !g.IsEmpty(PmsMemberLevel) && PmsMemberLevel.Id != PmsMember.Level {
			dao.PmsMember.Ctx(ctx).WherePri(in.Id).
				Update(g.Map{
					dao.PmsMember.Columns().Level: PmsMemberLevel.Id,
				})
		}
		return
	})
}

func (s *sAppMember) MemberDelete(ctx context.Context, in *input_app_member.PmsMemberDeleteInp) (err error) {

	if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除会员信息失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberView(ctx context.Context, in *input_app_member.PmsMemberViewInp) (res *input_app_member.PmsMemberViewModel, err error) {
	if err = dao.PmsMember.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员信息信息，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberAddExp(ctx context.Context, MemberId int, Exp int, Txt string) (err error) {
	if _, err = dao.PmsMember.Ctx(ctx).Data(g.Map{
		"exp": gdb.Raw(fmt.Sprintf("exp+%d", Exp)),
	}).WherePri(MemberId).Update(); err != nil {
		err = gerror.Wrap(err, "修改会员经验失败，请稍后重试！")
		return
	}
	return
}

func (s *sAppMember) MemberStat(ctx context.Context, in *input_app_member.PmsMemberStatInp) (out *input_app_member.PmsMemberStatModel, err error) {
	out = &input_app_member.PmsMemberStatModel{}
	mod := dao.PmsMember.Ctx(ctx).Safe()

	todayStart := gtime.Now().StartOfDay()
	todayEnd := gtime.Now().EndOfDay()
	today := gtime.Now().Format("m-d")

	lastOneDayStart := gtime.New(todayStart).Add(time.Duration(-24) * time.Hour)
	lastOneDayEnd := gtime.New(todayEnd).Add(time.Duration(-24) * time.Hour)
	lastOneDay := gtime.New(lastOneDayStart).Format("m-d")

	lastTwoDayStart := gtime.New(lastOneDayStart).Add(time.Duration(-24) * time.Hour)
	lastTwoDayEnd := gtime.New(lastOneDayEnd).Add(time.Duration(-24) * time.Hour)
	lastTwoDay := gtime.New(lastTwoDayStart).Format("m-d")

	lastThreeDayStart := gtime.New(lastTwoDayStart).Add(time.Duration(-24) * time.Hour)
	lastThreeDayEnd := gtime.New(lastTwoDayEnd).Add(time.Duration(-24) * time.Hour)
	lastThreeDay := gtime.New(lastThreeDayStart).Format("m-d")

	lastFourDayStart := gtime.New(lastThreeDayStart).Add(time.Duration(-24) * time.Hour)
	lastFourDayEnd := gtime.New(lastThreeDayEnd).Add(time.Duration(-24) * time.Hour)
	lastFourDay := gtime.New(lastFourDayStart).Format("m-d")

	lastFiveDayStart := gtime.New(lastFourDayStart).Add(time.Duration(-24) * time.Hour)
	lastFiveDayEnd := gtime.New(lastFourDayEnd).Add(time.Duration(-24) * time.Hour)
	lastFiveDay := gtime.New(lastFiveDayStart).Format("m-d")

	lastSixDayStart := gtime.New(lastFiveDayStart).Add(time.Duration(-24) * time.Hour)
	lastSixDayEnd := gtime.New(lastFiveDayEnd).Add(time.Duration(-24) * time.Hour)
	lastSixDay := gtime.New(lastSixDayStart).Format("m-d")

	out.MemberBaseStat.MemberTotal, err = mod.Count()
	out.MemberBaseStat.TodayMemberAdd, err = mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, todayStart, todayEnd).Count()

	if in.Type == "memberBaseStat" {

		// 昨日新增会员数
		out.MemberBaseStat.YesterdayMemberAdd, err = mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).Count()

		out.MemberBaseStat.MemberBalanceTotal, err = mod.Sum(dao.PmsMember.Columns().Balance)

		// 今日消耗积分
		out.MemberBaseStat.TodayConsumedBalance, err = dao.PmsBalanceChange.Ctx(ctx).
			WhereBetween(dao.PmsBalanceChange.Columns().CreatedAt, todayStart, todayEnd).
			WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0).
			Sum(dao.PmsBalanceChange.Columns().ChangePrice)

		out.MemberBaseStat.TodayConsumedBalance = math.Abs(out.MemberBaseStat.TodayConsumedBalance)

		/*out.MemberBaseStat.TodayMemberOrderCount, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Where("order_status = ? OR refund_status = ?", "HAVE_PAID", "PART").
		Group(dao.PmsAppStay.Columns().MemberId).
		Count()*/

		// 今日民宿下单数
		out.MemberBaseStat.TodayHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
			WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			Where(dao.PmsAppStay.Columns().Source, "APP").
			Count()

		// 昨日民宿下单数
		out.MemberBaseStat.YesterdayHotelOrderNum, err = dao.PmsAppStay.Ctx(ctx).
			WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).
			Where(dao.PmsAppStay.Columns().OrderStatus, "HAVE_PAID").
			Where(dao.PmsAppStay.Columns().Source, "APP").
			Count()

		/*out.MemberBaseStat.TodayMemberOrderCount, err = dao.PmsAppStay.Ctx(ctx).
		WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).
		Where(dao.PmsAppStay.Columns().Source, "APP").
		Group(dao.PmsAppStay.Columns().MemberId).
		Count()*/

		return
	}

	if in.Type == "memberOrderPie" {

		out.MemberOrderPie.MemberOrder, err = dao.PmsAppStay.Ctx(ctx).Group(dao.PmsAppStay.Columns().MemberId).Count()

		out.MemberOrderPie.MemberOrderPercent = decimal.NewFromInt(gvar.New(out.MemberOrderPie.MemberOrder).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()
		//out.MemberOrderPie.MemberOrderPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(out.MemberOrderPie.MemberOrder)/float64(out.MemberBaseStat.MemberTotal)*100), 64)

		out.MemberOrderPie.MemberNotOrder = out.MemberBaseStat.MemberTotal - out.MemberOrderPie.MemberOrder

		out.MemberOrderPie.MemberNotOrderPercent = decimal.NewFromInt(gvar.New(out.MemberOrderPie.MemberNotOrder).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()

		//out.MemberOrderPie.MemberNotOrderPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(out.MemberOrderPie.MemberNotOrder)/float64(out.MemberBaseStat.MemberTotal)*100), 64)

		return
	}

	if in.Type == "memberCreateTrend" {

		memberCreateTrendLastOneDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).Count()
		memberCreateTrendLastTwoDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastTwoDayStart, lastTwoDayEnd).Count()
		memberCreateTrendLastThreeDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastThreeDayStart, lastThreeDayEnd).Count()
		memberCreateTrendLastFourDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastFourDayStart, lastFourDayEnd).Count()
		memberCreateTrendLastFiveDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastFiveDayStart, lastFiveDayEnd).Count()
		memberCreateTrendLastSixDay, _ := mod.WhereBetween(dao.PmsMember.Columns().CreatedAt, lastSixDayStart, lastSixDayEnd).Count()
		memberCreateTrend := g.Slice{
			g.Map{
				"Num":  memberCreateTrendLastSixDay,
				"Date": lastSixDay,
			},
			g.Map{
				"Num":  memberCreateTrendLastFiveDay,
				"Date": lastFiveDay,
			},
			g.Map{
				"Num":  memberCreateTrendLastFourDay,
				"Date": lastFourDay,
			},
			g.Map{
				"Num":  memberCreateTrendLastThreeDay,
				"Date": lastThreeDay,
			},
			g.Map{
				"Num":  memberCreateTrendLastTwoDay,
				"Date": lastTwoDay,
			},
			g.Map{
				"Num":  memberCreateTrendLastOneDay,
				"Date": lastOneDay,
			},
			g.Map{
				"Num":  out.MemberBaseStat.TodayMemberAdd,
				"Date": today,
			},
		}
		if err = gvar.New(memberCreateTrend).Struct(&out.MemberCreateTrend); err != nil {
			return
		}
		return
	}

	if in.Type == "memberOrderTrend" {
		var (
			memberOrderTrendToday        float64
			memberOrderTrendLastOneDay   float64
			memberOrderTrendLastTwoDay   float64
			memberOrderTrendLastThreeDay float64
			memberOrderTrendLastFourDay  float64
			memberOrderTrendLastFiveDay  float64
			memberOrderTrendLastSixDay   float64
		)

		if memberOrderTrendToday, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, todayStart, todayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastOneDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastOneDayStart, lastOneDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastTwoDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastTwoDayStart, lastTwoDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastThreeDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastThreeDayStart, lastThreeDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastFourDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastFourDayStart, lastFourDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastFiveDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastFiveDayStart, lastFiveDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		if memberOrderTrendLastSixDay, err = dao.PmsAppStay.Ctx(ctx).WhereBetween(dao.PmsAppStay.Columns().CreatedAt, lastSixDayStart, lastSixDayEnd).Sum(dao.PmsAppStay.Columns().OrderAmount); err != nil {
			return
		}
		memberOrderTrend := g.Slice{
			g.Map{
				"Money": memberOrderTrendLastSixDay,
				"Date":  lastSixDay,
			},
			g.Map{
				"Money": memberOrderTrendLastFiveDay,
				"Date":  lastFiveDay,
			},
			g.Map{
				"Money": memberOrderTrendLastFourDay,
				"Date":  lastFourDay,
			},
			g.Map{
				"Money": memberOrderTrendLastThreeDay,
				"Date":  lastThreeDay,
			},
			g.Map{
				"Money": memberOrderTrendLastTwoDay,
				"Date":  lastTwoDay,
			},
			g.Map{
				"Money": memberOrderTrendLastOneDay,
				"Date":  lastOneDay,
			},
			g.Map{
				"Money": memberOrderTrendToday,
				"Date":  today,
			},
		}
		if err = gvar.New(memberOrderTrend).Struct(&out.MemberOrderTrend); err != nil {
			return out, err
		}
		return out, nil
	}

	if in.Type == "memberSource" {

		memberIOSSourceNum, _ := mod.Where(dao.PmsMember.Columns().Source, "IOS").Count()

		memberIOSSourcePercent := decimal.NewFromInt(gvar.New(memberIOSSourceNum).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()
		//memberIOSSourcePercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(memberIOSSourceNum)/float64(out.MemberBaseStat.MemberTotal)*100), 64)

		memberAndroidSourceNum, _ := mod.Where(dao.PmsMember.Columns().Source, "Android").Count()

		memberAndroidSourcePercent := decimal.NewFromInt(gvar.New(memberAndroidSourceNum).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()
		//memberAndroidSourcePercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(memberAndroidSourceNum)/float64(out.MemberBaseStat.MemberTotal)*100), 64)

		memberH5SourceNum, _ := mod.Where(dao.PmsMember.Columns().Source, "H5").Count()

		memberH5SourcePercent := decimal.NewFromInt(gvar.New(memberH5SourceNum).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()

		//memberH5SourcePercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(memberH5SourceNum)/float64(out.MemberBaseStat.MemberTotal)*100), 64)

		memberWxMiniSourceNum, _ := mod.Where(dao.PmsMember.Columns().Source, "WX_MINI").Count()

		memberWxMiniSourcePercent := decimal.NewFromInt(gvar.New(memberWxMiniSourceNum).Int64()).
			Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
			Mul(decimal.NewFromInt(100)).
			Round(2).
			InexactFloat64()

		memberSource := g.Slice{
			g.Map{
				"Source":  "IOS",
				"Num":     memberIOSSourceNum,
				"Percent": memberIOSSourcePercent,
			},
			g.Map{
				"Source":  "Android",
				"Num":     memberAndroidSourceNum,
				"Percent": memberAndroidSourcePercent,
			},
			g.Map{
				"Source":  "H5",
				"Num":     memberH5SourceNum,
				"Percent": memberH5SourcePercent,
			},
			g.Map{
				"Source":  "WX_MINI",
				"Num":     memberWxMiniSourceNum,
				"Percent": memberWxMiniSourcePercent,
			},
		}
		if err = gvar.New(memberSource).Struct(&out.MemberSource); err != nil {
			return
		}
		return
	}

	if in.Type == "memberLevel" {

		var memberLevelList []*input_app_member.PmsMemberLevelAllModel
		_ = dao.PmsMemberLevel.Ctx(ctx).WithAll().Fields(input_app_member.PmsMemberLevelAllModel{}).Scan(&memberLevelList)
		memberLevel := g.Slice{}
		for _, v := range memberLevelList {

			memberLevelItemPercent := decimal.NewFromInt(gvar.New(len(v.MemberDetail)).Int64()).
				Div(decimal.NewFromInt(gvar.New(out.MemberBaseStat.MemberTotal).Int64())).
				Mul(decimal.NewFromInt(100)).
				Round(2).
				InexactFloat64()

			memberLevelItem := g.Map{
				"LevelName": v.LevelName,
				"Num":       len(v.MemberDetail),
				"Percent":   memberLevelItemPercent,
			}
			memberLevel = append(memberLevel, memberLevelItem)
		}
		if err = gvar.New(memberLevel).Struct(&out.MemberLevel); err != nil {
			return
		}
		return
	}

	return
}

// MemberBalanceChange 会员余额变动
func (s *sAppMember) MemberBalanceChange(ctx context.Context, in *input_app_member.MemberBalanceInp, tx gdb.TX) (err error) {
	var (
		MemberInfo *entity.PmsMember
		result     sql.Result
		row        int64
	)
	// 锁住用户余额操作
	if !gmlock.TryLock("BalanceChange:" + gvar.New(in.MemberId).String()) {
		err = gerror.New("用户当前异常无法支付积分")
	}
	// 解锁用户余额操作
	defer gmlock.Unlock("BalanceChange:" + gvar.New(in.MemberId).String())
	// 查询用户信息
	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().Id, in.MemberId).Scan(&MemberInfo); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		err = gerror.New("用户不存在")
		return
	}
	// 验证余额
	if in.ChangeBalance < 0 {
		if math.Abs(in.ChangeBalance) > MemberInfo.Balance && in.ChangeBalance < 0 {
			err = gerror.New("用户积分不足")
			return
		}
	}
	// 更新余额
	if result, err = dao.PmsMember.Ctx(ctx).TX(tx).
		Where(g.MapStrAny{
			dao.PmsMember.Columns().Id:      in.MemberId,
			dao.PmsMember.Columns().Balance: MemberInfo.Balance,
		}).
		Update(g.MapStrAny{
			dao.PmsMember.Columns().Balance: MemberInfo.Balance + in.ChangeBalance,
		}); err != nil {
		return
	}
	// 影响行数
	if row, err = result.RowsAffected(); err != nil {
		return
	}
	// 影响行数
	g.Log().Info(ctx, "MemberBalanceChange", "更新余额影响行数", row)
	if row != 1 {
		err = gerror.New("操作余额失败")
		return
	}
	// 添加余额变动记录
	if _, err = dao.PmsBalanceChange.Ctx(ctx).TX(tx).OmitEmptyData().Insert(entity.PmsBalanceChange{
		MemberId:    in.MemberId,
		ChangePrice: in.ChangeBalance,
		OrderSn:     in.OrderSn,
		Scene:       in.Scene,
		Type:        in.Type,
		Des:         in.Des,
		Reason:      in.Reason,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		MdCode:      in.MdCode,
		MpModel:     in.MpModel,
	}); err != nil {
		return
	}
	return
}

// MemberInfo 查询用户信息
func (s *sAppMember) MemberInfo(ctx context.Context, MemberId int) (MemberInfo *entity.PmsMember, err error) {
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, MemberId).Scan(&MemberInfo); err != nil {
		return
	}
	if g.IsEmpty(MemberInfo) {
		err = gerror.New("用户不存在")
		return
	}
	return
}

// Status 修改会员状态
func (s *sAppMember) Status(ctx context.Context, in *input_app_member.PmsMemberStatusInp) (err error) {
	if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		dao.PmsMember.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新会员状态失败，请稍后重试！")
		return
	}
	return
}

// Cancel 会员注销
func (s *sAppMember) Cancel(ctx context.Context, in *input_app_member.PmsMemberCancelInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		var (
			PmsMember       *entity.PmsMember
			PmsMemberCancel *entity.PmsMemberCancel
			InsertData      *entity.PmsMemberCancel
		)

		// 获取用户信息
		if err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Scan(&PmsMember); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return
		}

		// 根据会员ID找到会员注销记录
		if err = dao.PmsMemberCancel.Ctx(ctx).
			Where(g.Map{
				dao.PmsMemberCancel.Columns().MemberId: PmsMember.Id,
			}).
			WhereIn(dao.PmsMemberCancel.Columns().AuditStatus, []int{1, 2}).OrderDesc(dao.PmsMemberCancel.Columns().Id).
			Scan(&PmsMemberCancel); err != nil {
			return
		}

		if !g.IsEmpty(PmsMemberCancel) {
			err = gerror.New("您已申请注销会员，请勿重复申请！")
			return
		}

		// 获取会员注销设置
		var Config *input_basics.GetConfigModel
		Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
			Group: "membercancelsetting",
		})
		if err != nil {
			err = gerror.Wrap(err, "未找到注销设置！")
			return
		}

		IsEnableCancel := Config.List["isEnableCancel"]
		IsCancelAudit := Config.List["isCancelAudit"]
		if IsEnableCancel != 1 {
			err = gerror.New("不可注销会员账号！")
			return err
		}

		if IsCancelAudit == 1 {
			InsertData = &entity.PmsMemberCancel{
				MemberId:     gvar.New(in.Id).Uint(),
				MemberNo:     PmsMember.MemberNo,
				Phone:        PmsMember.Phone,
				PhoneArea:    PmsMember.PhoneArea,
				Mail:         PmsMember.Mail,
				CancelReason: in.CancelReason,
			}
		} else {
			// 无需审核
			InsertData = &entity.PmsMemberCancel{
				MemberId:     gvar.New(in.Id).Uint(),
				MemberNo:     PmsMember.MemberNo,
				Phone:        PmsMember.Phone,
				PhoneArea:    PmsMember.PhoneArea,
				Mail:         PmsMember.Mail,
				CancelReason: in.CancelReason,
				AuditStatus:  2,
				AuditTime:    gtime.Now(),
				AuditReason:  "自动审核成功",
				OperatorId:   int(contexts.GetUserId(ctx)),
			}

			// 若会员绑定了渠道或员工，则解除渠道或员工的绑定关系
			if PmsMember.RebateMode == "CHANNEL" {
				// 解绑渠道
				if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Data(g.Map{
					dao.PmsMember.Columns().RebateMode: "MEMBER",
					dao.PmsMember.Columns().ChannelId:  nil,
				}).Update(); err != nil {
					err = gerror.Wrap(err, "解绑渠道失败，请稍后重试！")
				}

			}
			if PmsMember.RebateMode == "STAFF" {
				// 解绑员工
				if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Data(g.Map{
					dao.PmsMember.Columns().RebateMode: "MEMBER",
					dao.PmsMember.Columns().StaffId:    nil,
				}).Update(); err != nil {
					err = gerror.Wrap(err, "解绑员工失败，请稍后重试！")
				}
			}

			// 查询是否绑定按摩服务商、按摩技师、接送机司机、员工活动中的员工，进行解绑
			// 解绑按摩服务商
			if _, err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().MemberId, in.Id).Data(g.Map{
				dao.SpaIsp.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑按摩服务商失败，请稍后重试！")
				return
			}
			// 解绑按摩技师
			if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().MemberId, in.Id).Data(g.Map{
				dao.SpaTechnician.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑按摩技师失败，请稍后重试！")
				return
			}
			// 解绑接送机司机
			if _, err = dao.CarDriver.Ctx(ctx).Where(dao.CarDriver.Columns().MemberId, in.Id).Data(g.Map{
				dao.CarDriver.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑接送机司机失败，请稍后重试！")
				return
			}
			// 解绑员工活动中的员工
			if _, err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, in.Id).Data(g.Map{
				dao.Employee.Columns().MemberId: 0,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "解绑员工活动员工失败，请稍后重试！")
				return
			}

			// 删除会员信息
			if _, err = dao.PmsMember.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
				err = gerror.Wrap(err, "删除会员失败，请稍后重试！")
				return
			}
		}
		// 写入会员注销记录
		if _, err = dao.PmsMemberCancel.Ctx(ctx).Data(InsertData).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增会员注销记录失败，请稍后重试！")
			return
		}
		return
	})

}
