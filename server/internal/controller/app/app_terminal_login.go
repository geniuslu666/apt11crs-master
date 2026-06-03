package app

import (
	"APT/internal/dao"
	"APT/internal/library/token"
	"APT/internal/model"
	"APT/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/terminal"
)

func (c *ControllerTerminal) TerminalLogin(ctx context.Context, req *terminal.TerminalLoginReq) (res *terminal.TerminalLoginRes, err error) {
	var (
		TerminalInfo      *entity.SysTerminal
		RestaurantInfo    *entity.FoodRestaurant
		StoreInfo         *entity.ThMchStore
		LoginType         string
		LoginMchId        int64
		LoginStoreId      int64
		LoginRestaurantId int64
	)

	res = new(terminal.TerminalLoginRes)
	if err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().Sn, req.TerminalSn).Scan(&TerminalInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(TerminalInfo) {
		// 设备号不存在
		err = gerror.New(gi18n.T(ctx, "device_number_does_not_exist"))
		return
	}

	if err = dao.FoodRestaurant.Ctx(ctx).Where(dao.FoodRestaurant.Columns().Account, req.Account).Scan(&RestaurantInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if err = dao.ThMchStore.Ctx(ctx).Where(dao.ThMchStore.Columns().Account, req.Account).Scan(&StoreInfo); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	if g.IsEmpty(RestaurantInfo) && g.IsEmpty(StoreInfo) {
		err = gerror.New(gi18n.T(ctx, "account_does_not_exist"))
		return
	}

	if !g.IsEmpty(RestaurantInfo) && RestaurantInfo.Id != int64(TerminalInfo.RestaurantId) {
		// 账号不存在
		err = gerror.New(gi18n.T(ctx, "account_does_not_exist"))
		return
	}

	if !g.IsEmpty(StoreInfo) && StoreInfo.Id != int64(TerminalInfo.StoreId) {
		err = gerror.New(gi18n.T(ctx, "account_does_not_exist"))
		return
	}

	if !g.IsEmpty(RestaurantInfo) {
		if RestaurantInfo.PasswordHash != gmd5.MustEncryptString(req.Password+RestaurantInfo.Salt) {
			// 密码不正确
			err = gerror.New(gi18n.T(ctx, "password_incorrect"))
			return
		}
		LoginType = "RESTAURANT"
		LoginRestaurantId = RestaurantInfo.Id
	} else {
		if StoreInfo.PasswordHash != gmd5.MustEncryptString(req.Password+StoreInfo.Salt) {
			err = gerror.New(gi18n.T(ctx, "password_incorrect"))
			return
		}
		LoginType = "STORE"
		LoginMchId = int64(StoreInfo.MchId)
		LoginStoreId = StoreInfo.Id
	}

	user := &model.TerminalIdentity{
		Type:         LoginType,
		Account:      req.Account,
		TerminalId:   int64(TerminalInfo.Id),
		MchId:        LoginMchId,
		StoreId:      LoginStoreId,
		RestaurantId: LoginRestaurantId,
		LoginAt:      gtime.Now(),
	}

	if res.Token, res.Expires, err = token.TerminalLogin(ctx, user); err != nil {
		return
	}
	return
}
func (c *ControllerTerminal) TerminalLogout(ctx context.Context, req *terminal.TerminalLogoutReq) (res *terminal.TerminalLogoutRes, err error) {
	err = token.TerminalLogout(ghttp.RequestFromCtx(ctx))
	return
}
