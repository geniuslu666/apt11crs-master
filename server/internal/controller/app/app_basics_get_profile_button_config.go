package app

import (
	"APT/internal/dao"
	"APT/internal/model"
	"context"

	"APT/api/app/basics"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
)

func (c *ControllerBasics) GetProfileButtonConfig(ctx context.Context, req *basics.GetProfileButtonConfigReq) (res *basics.GetProfileButtonConfigRes, err error) {
	var (
		dataVar *gvar.Var
	)
	res = new(basics.GetProfileButtonConfigRes)

	// 从配置表获取个人中心按钮配置
	if dataVar, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Group, "profileButtons").
		Where(dao.SysConfig.Columns().Key, "profileButtonData").
		Value(dao.SysConfig.Columns().Value); err != nil {
		return
	}
	if dataVar == nil || dataVar.IsEmpty() {
		return
	}

	// 解析 JSON 字符串
	var config model.ProfileButtonInfo
	if err = gjson.New(dataVar.String()).Scan(&config); err != nil {
		return
	}

	res.Data = &basics.ProfileButtonInfo{
		Qrcode:    config.Qrcode,
		Coupon:    config.Coupon,
		ThCoupon:  config.ThCoupon,
		DataBoard: config.DataBoard,
		Employee:  config.Employee,
		Workbench: config.Workbench,
		Fx:        config.Fx,
		Help:      config.Help,
		Contact:   config.Contact,
		Agreement: config.Agreement,
		Privacy:   config.Privacy,
	}

	return
}
