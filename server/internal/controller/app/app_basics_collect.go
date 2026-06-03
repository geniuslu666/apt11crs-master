package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/app/basics"
)

func (c *ControllerBasics) CollectAdd(ctx context.Context, req *basics.CollectAddReq) (res *basics.CollectAddRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	err = service.BasicsCollect().Add(ctx, &input_basics.PmsCollectAddInp{
		MemberId:    MemberInfo.Id,
		CollectType: req.CollectType,
		CollectId:   req.CollectId,
	})
	return
}

func (c *ControllerBasics) CollectList(ctx context.Context, req *basics.CollectListReq) (res *basics.CollectListRes, err error) {
	var (
		CollectIds  []int
		WhereParams = &req.PmsCollectListInp
	)
	WhereParams.MemberId = contexts.GetMemberUser(ctx).Id
	res = new(basics.CollectListRes)
	list, totalCount, err := service.BasicsCollect().List(ctx, WhereParams)
	if err != nil {
		return
	}

	if !g.IsEmpty(list) {
		for _, v := range list {
			CollectIds = append(CollectIds, v.CollectId)
		}
		if err = dao.PmsProperty.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Where(g.MapStrAny{
			dao.PmsProperty.Columns().Id: CollectIds,
		}).Scan(&res.List); err != nil {
			return
		}
		for _, v := range list {
			for kk, vv := range res.List {
				if v.CollectId == vv.Id {
					res.List[kk].CollectId = v.Id
					break
				}
			}
		}
	}

	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerBasics) CollectDelete(ctx context.Context, req *basics.CollectDeleteReq) (res *basics.CollectDeleteRes, err error) {
	err = service.BasicsCollect().Delete(ctx, &req.PmsCollectDeleteInp)
	return
}
