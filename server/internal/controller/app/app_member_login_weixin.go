package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/cache"
	"APT/internal/library/token"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"context"
	"encoding/json"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"APT/api/app/member"
)

func (c *ControllerMember) OauthLogin(ctx context.Context, req *member.OauthLoginReq) (res *member.OauthLoginRes, err error) {
	var (
		ghttpclient        = g.Client()
		openUri            = "https://api.weixin.qq.com/sns/oauth2/access_token"
		wxMiniUri          = "https://api.weixin.qq.com/sns/jscode2session"
		WeixinAuthResponse *member.WeixinAuthResponse
		PmsMemberAuth      *entity.PmsMemberAuth
		PmsMember          *entity.PmsMember
		MemberTokenInfo    *model.MemberIdentity
		AuthId             string
		WxUnionid          string
		Channel            string
		r                  *gclient.Response
		tx                 gdb.TX
	)
	res = new(member.OauthLoginRes)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if !g.IsEmpty(req.Code) {
		ghttpclient.SetHeader("Accept", "*/*")
		if r, err = ghttpclient.Get(ctx, openUri, g.MapStrAny{
			"appid":      g.Cfg().MustGet(ctx, "WeixinOpenplatform.appid"),
			"secret":     g.Cfg().MustGet(ctx, "WeixinOpenplatform.appsecret"),
			"code":       req.Code,
			"grant_type": "authorization_code",
		}); err != nil {
			return
		}

		defer r.Close()
		g.Log().Info(ctx, r.Raw())
		if err = json.Unmarshal(r.ReadAll(), &WeixinAuthResponse); err != nil {
			return
		}
		if g.IsEmpty(WeixinAuthResponse) {
			// 微信登录失败
			err = gerror.New(gi18n.T(ctx, "wechat_login_failed"))
			return
		}
		AuthId = WeixinAuthResponse.Openid
		WxUnionid = WeixinAuthResponse.Unionid
		if g.IsEmpty(AuthId) {
			// 微信登录失败
			err = gerror.New(gi18n.T(ctx, "wechat_login_failed"))
			return
		}
		Channel = "WX"
	} else if !g.IsEmpty(req.MiniCode) {
		ghttpclient.SetHeader("Accept", "*/*")
		if r, err = ghttpclient.Get(ctx, wxMiniUri, g.MapStrAny{
			"appid":      g.Cfg().MustGet(ctx, "WeixinMiniprogram.appid"),
			"secret":     g.Cfg().MustGet(ctx, "WeixinMiniprogram.appsecret"),
			"js_code":    req.MiniCode,
			"grant_type": "authorization_code",
		}); err != nil {
			return
		}

		defer r.Close()
		g.Log().Info(ctx, r.Raw())
		if err = json.Unmarshal(r.ReadAll(), &WeixinAuthResponse); err != nil {
			return
		}
		if g.IsEmpty(WeixinAuthResponse) {
			err = gerror.New(gi18n.T(ctx, "wechat_login_failed"))
			return
		}
		AuthId = WeixinAuthResponse.Openid
		if g.IsEmpty(AuthId) {
			err = gerror.New(gi18n.T(ctx, "wechat_login_failed"))
			return
		}
		WxUnionid = WeixinAuthResponse.Unionid
		Channel = "WX_MINI"
	} else {
		AuthId = req.AppleId
		Channel = "APPLE"
	}

	// 查询是否存在这个openid
	if err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsMemberAuth.Columns().AuthId: AuthId,
	}).Scan(&PmsMemberAuth); err != nil {
		return
	}
	if g.IsEmpty(PmsMemberAuth) {
		// 插入数据
		if _, err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsMemberAuth{
			AuthId:    AuthId,
			WxUnionid: WxUnionid,
			Channel:   Channel,
		}); err != nil {
			return
		}
		res.IsBind = false
		res.AuthId = AuthId
	} else if !g.IsEmpty(PmsMemberAuth.MemberId) {
		// 登录流程 查询用户信息
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Id: PmsMemberAuth.MemberId,
		}).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			if _, err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.PmsMemberAuth.Columns().AuthId: AuthId,
			}).Delete(); err != nil {
				return
			}
			// 如果该用户绑定的会员ID 查询不到   让用户继续绑定手机号
			res.IsBind = false
			res.AuthId = AuthId
			return
		}
		// 更新登录时间和IP

		PmsMemberInfoUpdate := g.MapStrAny{
			dao.PmsMember.Columns().LastLogin:   gtime.Now(),
			dao.PmsMember.Columns().LoginMode:   Channel,
			dao.PmsMember.Columns().LastLoginIp: ghttp.RequestFromCtx(ctx).GetClientIp(),
		}

		if !g.IsEmpty(req.Referrer) && req.Referrer != PmsMember.Id {

			PmsReferrerLogUpdate := g.MapStrAny{
				dao.PmsReferrerLog.Columns().MemberId:     PmsMember.Id,
				dao.PmsReferrerLog.Columns().LastReferrer: req.Referrer,
			}

			if g.IsEmpty(PmsMember.Referrer) {
				PmsMemberInfoUpdate[dao.PmsMember.Columns().Referrer] = req.Referrer
				PmsReferrerLogUpdate[dao.PmsReferrerLog.Columns().Referrer] = req.Referrer
			}
			PmsMemberInfoUpdate[dao.PmsMember.Columns().LastReferrer] = req.Referrer

			// 插入记录
			if _, err = dao.PmsReferrerLog.Ctx(ctx).Data(PmsReferrerLogUpdate).InsertAndGetId(); err != nil {
				return
			}
		}

		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, PmsMember.Id).
			OmitEmptyData().Update(PmsMemberInfoUpdate); err != nil {
			return
		}
		if err = gvar.New(PmsMember).Struct(&MemberTokenInfo); err != nil {
			return
		}

		MemberTokenInfo.LoginAt = gtime.Now()
		MemberTokenInfo.App = consts.AppMember
		if res.Token, res.Expires, err = token.MemberLogin(ctx, MemberTokenInfo); err != nil {
			return
		}

		// 记录登录记录
		if _, err = dao.PmsMemberLog.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsMemberLog{
			MemberId:  PmsMemberAuth.MemberId,
			LoginTime: gtime.Now(),
			LoginType: Channel,
			LoginIp:   ghttp.RequestFromCtx(ctx).GetClientIp(),
			Token:     res.Token,
			ExpirTime: gtime.New(res.Expires + gtime.Now().Unix()),
			MdCode:    req.MdCode,
		}).Insert(); err != nil {
			return
		}
		res.IsBind = true
		res.AuthId = AuthId
	} else {
		res.IsBind = false
		res.AuthId = AuthId
	}
	return
}
func (c *ControllerMember) PhoneNumber(ctx context.Context, req *member.PhoneNumberReq) (res *member.PhoneNumberRes, err error) {
	res = new(member.PhoneNumberRes)
	MiniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  g.Cfg().MustGet(ctx, "WeixinMiniprogram.appid").String(),     // 小程序appid
		Secret: g.Cfg().MustGet(ctx, "WeixinMiniprogram.appsecret").String(), // 小程序app secret
	})
	if err != nil {
		// 初始化微信小程序失败
		err = gerror.New(gi18n.T(ctx, "failed_to_init_wechat_mini_program"))
		return
	}
	response, err := MiniProgramApp.PhoneNumber.GetUserPhoneNumber(ctx, req.MiniCode)
	g.Log().Info(ctx, response)
	if g.IsEmpty(response) {
		// 获取手机号失败
		err = gerror.New(gi18n.T(ctx, "failed_to_obtain_phone_number"))
		return
	}
	if response.Code != "0" && response.Code != "" {
		err = gerror.New(gi18n.T(ctx, "failed_to_obtain_phone_number"))
		return
	}
	if err != nil {
		err = gerror.New(gi18n.T(ctx, "failed_to_obtain_phone_number"))
		return
	}
	if response.ErrCode != 0 {
		err = gerror.New(response.ErrMsg)
		return
	}
	res.PhoneNumber = response.PhoneInfo.PhoneNumber
	res.PurePhoneNumber = response.PhoneInfo.PurePhoneNumber
	res.CountryCode = response.PhoneInfo.CountryCode
	return
}
func (c *ControllerMember) FxAuthLogin(ctx context.Context, req *member.FxAuthLoginReq) (res *member.FxAuthLoginRes, err error) {
	var (
		PmsMemberAuth   *entity.PmsMemberAuth
		PmsMember       *entity.PmsMember
		MemberTokenInfo *model.MemberIdentity
		tx              gdb.TX
		FxLoginInfo     *input_app_member.FxMemberInfo
		gvarValue       *gvar.Var
	)
	res = new(member.FxAuthLoginRes)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if gvarValue, err = cache.Instance().Get(ctx, req.Code); err != nil {
		return
	}
	if err = gvarValue.Struct(&FxLoginInfo); err != nil {
		return
	}
	if g.IsEmpty(FxLoginInfo) {
		// 微信登录失败
		err = gerror.New(gi18n.T(ctx, "wechat_login_failed"))
		return
	}

	// 查询是否存在这个openid
	if err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsMemberAuth.Columns().AuthId: FxLoginInfo.AuthId,
	}).Scan(&PmsMemberAuth); err != nil {
		return
	}
	g.Log().Info(ctx, FxLoginInfo)
	res.AuthId = FxLoginInfo.AuthId
	if g.IsEmpty(PmsMemberAuth) {
		// 插入数据
		if _, err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.PmsMemberAuth{
			AuthId:  FxLoginInfo.AuthId,
			Channel: FxLoginInfo.Channel,
		}); err != nil {
			return
		}
		res.IsBind = false
	} else if !g.IsEmpty(PmsMemberAuth.MemberId) {
		// 登录流程 查询用户信息
		if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(g.MapStrAny{
			dao.PmsMember.Columns().Id: PmsMemberAuth.MemberId,
		}).Scan(&PmsMember); err != nil {
			return
		}
		if g.IsEmpty(PmsMember) {
			if _, err = dao.PmsMemberAuth.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.PmsMemberAuth.Columns().AuthId: FxLoginInfo.AuthId,
			}).Delete(); err != nil {
				return
			}
			// 如果该用户绑定的会员ID 查询不到   让用户继续绑定手机号
			res.IsBind = false
			return
		}
		// 更新登录时间和IP
		if _, err = dao.PmsMember.Ctx(ctx).TX(tx).
			Where(dao.PmsMember.Columns().Id, PmsMember.Id).
			OmitEmptyData().Update(g.MapStrAny{
			dao.PmsMember.Columns().LastLogin:    gtime.Now(),
			dao.PmsMember.Columns().LoginMode:    FxLoginInfo.Channel,
			dao.PmsMember.Columns().LastLoginIp:  ghttp.RequestFromCtx(ctx).GetClientIp(),
			dao.PmsMember.Columns().LastReferrer: req.Referrer,
		}); err != nil {
			return
		}
		if err = gvar.New(PmsMember).Struct(&MemberTokenInfo); err != nil {
			return
		}

		MemberTokenInfo.LoginAt = gtime.Now()
		MemberTokenInfo.App = consts.AppMember
		MemberTokenInfo.Channel = FxLoginInfo.Channel
		MemberTokenInfo.PmsMemberAuth = PmsMemberAuth
		MemberTokenInfo.IsFx = true
		if res.Token, res.Expires, err = token.MemberLogin(ctx, MemberTokenInfo); err != nil {
			return
		}

		// 记录登录记录
		if _, err = dao.PmsMemberLog.Ctx(ctx).TX(tx).OmitEmptyData().Data(&entity.PmsMemberLog{
			MemberId:  PmsMemberAuth.MemberId,
			LoginTime: gtime.Now(),
			LoginType: FxLoginInfo.Channel,
			LoginIp:   ghttp.RequestFromCtx(ctx).GetClientIp(),
			Token:     res.Token,
			ExpirTime: gtime.New(res.Expires + gtime.Now().Unix()),
			MdCode:    req.MdCode,
		}).Insert(); err != nil {
			return
		}
		res.IsBind = true
	} else {
		res.IsBind = false
	}
	return
}
