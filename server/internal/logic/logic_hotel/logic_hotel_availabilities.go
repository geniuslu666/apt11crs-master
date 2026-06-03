package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

// CheckHotelInventory 校验酒店库存
func (s *sHotelService) CheckHotelInventory(ctx context.Context, in *input_hotel.CheckHotelInventoryInp) (err error) {
	var (
		allotment *gvar.Var
	)
	// 校验库存
	if allotment, err = dao.PmsAvailabilities.Ctx(ctx).OmitEmptyWhere().
		Where(&entity.PmsAvailabilities{
			Puid: gvar.New(in.PUID).String(),
			Tuid: gvar.New(in.TUID).String(),
		}).
		WhereGTE(dao.PmsAvailabilities.Columns().Date, in.StartDate).
		WhereLT(dao.PmsAvailabilities.Columns().Date, in.EndDate).
		Value("MIN(allotment) as allotment"); err != nil {
		return
	}
	if allotment.Int() < in.Num {
		// 库存不足
		err = gerror.New(gi18n.T(ctx, "insufficient_inventory"))
		return
	}
	return
}
