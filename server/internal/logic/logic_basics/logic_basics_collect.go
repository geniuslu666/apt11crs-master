package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type sBasicsCollect struct{}

func NewBasicsCollect() *sBasicsCollect {
	return &sBasicsCollect{}
}

func init() {
	service.RegisterBasicsCollect(NewBasicsCollect())
}

func (s *sBasicsCollect) Add(ctx context.Context, in *input_basics.PmsCollectAddInp) (err error) {
	var (
		lastInsertId int64
	)
	if lastInsertId, err = dao.PmsCollect.Ctx(ctx).Data(&entity.PmsCollect{
		MemberId:    in.MemberId,
		CollectType: in.CollectType,
		CollectId:   in.CollectId,
	}).InsertAndGetId(); err != nil {
		// 收藏失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "collect_failed"))
		return
	}
	if lastInsertId < 1 {
		// 收藏失败，请稍后重试！
		err = gerror.Wrap(err, gi18n.T(ctx, "collect_failed"))
		return
	}
	return
}

func (s *sBasicsCollect) List(ctx context.Context, in *input_basics.PmsCollectListInp) (list []*input_basics.PmsCollectListModel, totalCount int, err error) {
	mod := dao.PmsCollect.Ctx(ctx)

	mod = mod.Fields(input_basics.PmsCollectListModel{})

	if !g.IsEmpty(in.CollectType) {
		mod = mod.Where(dao.PmsCollect.Columns().CollectType, in.CollectType)
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.PmsCollect.Columns().MemberId, in.MemberId)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsCollect.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		// 获取收藏夹列表失败，请稍后重试
		err = gerror.Wrap(err, gi18n.T(ctx, "get_collect_list_failed"))
		return
	}
	return
}

func (s *sBasicsCollect) Delete(ctx context.Context, in *input_basics.PmsCollectDeleteInp) (err error) {

	if _, err = dao.PmsCollect.Ctx(ctx).WherePri(in.Id).Delete(); err != nil {
		// 删除收藏夹失败，请稍后重试
		err = gerror.Wrap(err, gi18n.T(ctx, "delete_collect_failed"))
		return
	}
	return
}
