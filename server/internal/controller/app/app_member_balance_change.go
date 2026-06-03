package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/utility/encrypt"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"APT/api/app/member"
)

func (c *ControllerMember) BalanceChange(ctx context.Context, req *member.BalanceChangeReq) (res *member.BalanceChangeRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(member.BalanceChangeRes)
	orm := dao.PmsBalanceChange.Ctx(ctx).Where(dao.PmsBalanceChange.Columns().MemberId, MemberInfo.PmsMember.Id)
	if !g.IsEmpty(req.ChangeMode) {
		switch req.ChangeMode {
		case "in":
			orm = orm.WhereGT(dao.PmsBalanceChange.Columns().ChangePrice, 0)
		case "out":
			orm = orm.WhereLT(dao.PmsBalanceChange.Columns().ChangePrice, 0)
		}
	}
	// 使用Fields方法选择字段，将reason字段作为des字段返回
	if err = orm.Fields("change_price as changePrice, order_sn as orderSn, reason as des, created_at as createdAt").Page(req.PageNum, req.PageSize).OrderDesc(dao.PmsBalanceChange.Columns().Id).ScanAndCount(&res.List, &res.Count, false); err != nil {
		return
	}

	// 定义需要进行多语言处理的消息类型
	/*messageType := []string{
		"注册奖励",
		"下单返积分",
		"退款",
		"推荐人下单返现",
		"邀请人首单奖励",
		"被邀请注册奖励",
		"支付餐饮费用",
		"支付住宿费用",
		"支付储物柜费用",
		"支付订房费用",
		"支付按摩费用",
		"支付出行费用",
	}*/

	for k, v := range res.List {
		if !g.IsEmpty(v.Des) && contexts.GetLanguage(ctx) != consts.Zh {
			res.List[k].Des = gi18n.GetContent(ctx, encrypt.Md5([]byte(v.Des)))
		}
	}
	return
}
