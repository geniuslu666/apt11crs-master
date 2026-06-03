package h5FxPay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
)

type ChangeOrderPushRequest struct {
	OrderNo      string `json:"orderNo" dc:"分销中心业务订单号，标识一笔业务订单，跳转至收银台时传入"`
	ChangeStatus string `json:"changeStatus" dc:"变更状态"`
}

type ChangeOrderPushResponse struct {
}

func ChangeOrderPush(ctx context.Context, params *ChangeOrderPushRequest, Logger *glog.Logger) (res *ChangeOrderPushResponse, err error) {
	var (
		gHttpClient   = g.Client()
		gHttpResponse *gclient.Response
		PayConfig     *model.PayConfig
	)
	defer func() {
		if r := recover(); r != nil {
			g.Log().Error(ctx, r)
		}
	}()
	if Logger == nil {
		Logger = g.Log().Path("logs/SDK/H5_FX")
	}
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	gHttpClient.SetHeader("Content-Type", "application/json")
	if gHttpResponse, err = gHttpClient.Post(ctx, PayConfig.H5FxDomain+"/external/apt11/state", params); err != nil {
		return
	}
	Logger.Info(ctx, gHttpResponse.Raw())
	defer gHttpResponse.Close()
	if gHttpResponse.StatusCode != 200 {
		err = gerror.New("失败")
		return
	}
	return
}
