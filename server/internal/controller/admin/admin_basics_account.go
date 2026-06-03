package admin

import (
	"APT/api/admin/basics"
	"APT/internal/consts"
	"APT/internal/library/captcha"
	"APT/internal/library/token"
	"APT/internal/service"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerBasics) AccountLoginLogout(ctx context.Context, _ *basics.AccountLoginLogoutReq) (res *basics.AccountLoginLogoutRes, err error) {
	err = token.Logout(ghttp.RequestFromCtx(ctx))
	return
}
func (c *ControllerBasics) AccountLoginCaptcha(ctx context.Context, _ *basics.AccountLoginCaptchaReq) (res *basics.AccountLoginCaptchaRes, err error) {
	cid, base64 := captcha.Generate(ctx)
	res = &basics.AccountLoginCaptchaRes{Cid: cid, Base64: base64}
	return
}

func (c *ControllerBasics) AccountLogin(ctx context.Context, req *basics.AccountLoginReq) (res *basics.AccountLoginRes, err error) {
	login, err := service.BasicsConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if !req.IsLock && login.CaptchaSwitch == 1 {
		// 校验 验证码
		if !captcha.Verify(req.Cid, req.Code) {
			err = gerror.New("验证码错误")
			return
		}
	}

	model, err := service.BasicsAdminSite().AccountLogin(ctx, &req.AccountLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}
func (c *ControllerBasics) AccountSiteConfig(ctx context.Context, req *basics.AccountSiteConfigReq) (res *basics.AccountSiteConfigRes, err error) {
	request := ghttp.RequestFromCtx(ctx)
	res = &basics.AccountSiteConfigRes{
		Version: consts.VersionApp,
		WsAddr:  getWsAddr(ctx, request),
		Domain:  getDomain(ctx, request),
		Mode:    gmode.Mode(),
	}
	return
}
func (c *ControllerBasics) AccountSiteLoginConfig(ctx context.Context, _ *basics.AccountSiteLoginConfigReq) (res *basics.AccountSiteLoginConfigRes, err error) {
	res = new(basics.AccountSiteLoginConfigRes)
	login, err := service.BasicsConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	res.LoginConfig = login
	return
}
func (c *ControllerBasics) AccountSitePing(_ context.Context, _ *basics.AccountSitePingReq) (_ *basics.AccountSitePingRes, err error) {
	return
}

func getWsAddr(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := ghttp.RequestFromCtx(ctx).GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "ws://" + ip + ":" + gstr.StrEx(request.Host, ":") + "/socket"
	}

	basic, err := service.BasicsConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.WsAddr
}

func getDomain(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := request.GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "http://" + ip + ":" + gstr.StrEx(request.Host, ":")
	}

	basic, err := service.BasicsConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.Domain
}
