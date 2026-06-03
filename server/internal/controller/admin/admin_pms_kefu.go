package admin

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/pms"
)

func (c *ControllerPms) OssKefuLogin(ctx context.Context, req *pms.OssKefuLoginReq) (res *pms.OssKefuLoginRes, err error) {
	var (
		UserInfo   *model.Identity
		account    string
		resp       *gclient.Response
		token      *pms.OssKefuToken
		AccountObj gdb.Record
	)
	UserInfo = contexts.GetUser(ctx)
	if g.IsEmpty(UserInfo) {
		err = gerror.New("没有客服权限")
		return
	}
	account = UserInfo.KefuAccount
	if g.IsEmpty(account) {
		err = gerror.New("没有客服权限")
		return
	}
	if AccountObj, err = dao.AdminMember.Ctx(ctx).Where(g.MapStrAny{
		dao.AdminMember.Columns().KefuAccount: account,
	}).One(); err != nil {
		err = gerror.New("没有客服权限")
		return
	}
	if AccountObj.IsEmpty() {
		err = gerror.New("没有客服权限")
		return
	}
	httpclient := g.Client()
	if resp, err = httpclient.Get(ctx, g.Cfg().MustGet(ctx, "kefu.kefuApiHost").String()+"/other/internalToken", g.MapStrAny{
		"account": account,
	}); err != nil {
		return
	}

	defer resp.Close()
	g.Log().Info(ctx, resp.Raw())
	if err = json.Unmarshal(resp.ReadAll(), &token); err != nil {
		return
	}
	if token.Code != 20000 {
		err = gerror.New("进入客服功能失败")
		return
	}
	res = new(pms.OssKefuLoginRes)
	res.IframeUrl = g.Cfg().MustGet(ctx, "kefu.kefuUrlHost").String() + "/sso?token=" + token.Result.Token + "&redirect=" + req.Redirect
	return
}
