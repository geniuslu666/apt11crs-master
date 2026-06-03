package notify

import (
	"APT/internal/model/input/input_cabinet"
	"APT/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/api/notify/hook"
)

func (c *ControllerHook) Cabinet(ctx context.Context, req *hook.CabinetReq) (res *hook.CabinetRes, err error) {
	var (
		r      = ghttp.RequestFromCtx(ctx)
		Logger = g.Log().Path("logs/HOOK/CABINET")
	)
	Logger.Info(ctx, gvar.New(r.GetBody()).String())
	Logger.Info(ctx, req)

	// 验证必要参数
	// 必要字段校验
	required := []string{"appid", "orderId", "outTradeNo", "eventType", "timestamp", "sign"}
	// 使用 map 便于逐项检查
	bodyMap := g.Map{}
	bodyMap = gconv.Map(req)
	for _, k := range required {
		if gconv.String(bodyMap[k]) == "" {
			Logger.Error(ctx, "缺少必要字段: "+k)
			r.Response.WriteStatus(400)
			r.Response.WriteJsonExit(g.Map{"code": 400, "msg": "缺少必要字段: " + k})
			return
		}
	}

	// 验证时间戳
	now := gtime.Now().Timestamp() // 当前时间戳（秒）
	diff := now - int64(req.Timestamp)
	if diff > 5*60 {
		Logger.Error(ctx, "时间戳过期")
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"code": 400, "msg": "时间戳过期"})
		return
	}

	// 验证签名
	// 从配置表获取储物柜API配置
	cabinetConfig, err := service.BasicsConfig().GetCabinetApi(ctx)
	if err != nil {
		Logger.Error(ctx, "获取储物柜API配置失败: "+err.Error())
		r.Response.WriteStatus(500)
		r.Response.WriteJsonExit(g.Map{"code": 500, "msg": "获取配置失败"})
		return
	}
	rawSignStr := fmt.Sprintf("%s%s%d%s", req.Appid, cabinetConfig.CabinetAppkey, req.Timestamp, cabinetConfig.CabinetApiSecret)
	Logger.Info(ctx, rawSignStr)
	sign := gmd5.MustEncryptString(rawSignStr)
	Logger.Info(ctx, req.Sign)
	Logger.Info(ctx, sign)
	if sign != req.Sign {
		Logger.Error(ctx, "签名错误")
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"code": 400, "msg": "签名错误"})
		return
	}

	// 处理订单
	if err = service.CabinetService().UpdateOrderStatus(ctx, &input_cabinet.UpdateOrderStatusInp{
		OrderSn:   req.OutTradeNo,
		EventType: req.EventType,
	}); err != nil {
		Logger.Error(ctx, err)
		r.Response.WriteStatus(400)
		r.Response.WriteJsonExit(g.Map{"code": 400, "msg": "订单处理失败"})
		return
	}

	r.Response.WriteJsonExit(g.Map{"code": 0, "msg": "success"})
	return
}
