package middleware

import (
	"APT/internal/consts"
	"APT/internal/library/contexts"
	"APT/internal/library/token"
	"APT/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

func getModule(path string) (module string) {
	slice := strings.Split(path, "/")
	if len(slice) < 2 {
		module = consts.AppDefault
		return
	}

	if slice[1] == "" {
		module = consts.AppDefault
		return
	}
	return slice[1]
}

func SetLanguage(r *ghttp.Request) {
	// 设置语言
	language := r.GetHeader("x-language")
	if g.IsEmpty(language) {
		language = g.Cfg().MustGet(r.Context(), "system.defaultLanguage").String()
	}
	language = gstr.TrimAll(gstr.ToLower(language), "-")
	language = gstr.TrimAll(gstr.ToLower(language), "_")
	switch language {
	case "tw":
		language = consts.ZhCN
		break
	case "zhtw":
		language = consts.ZhCN
		break
	case "ko":
		language = consts.Ko
		break
	case "kr":
		language = consts.Ko
		break
	case "enus":
		language = consts.En
		break
	case "en":
		language = consts.En
		break
	case "ja":
		language = consts.Ja
		break
	case "jp":
		language = consts.Ja
		break
	case "zhcn":
		language = consts.Zh
		break
	case "zh":
		language = consts.Zh
		break
	case "cn":
		language = consts.Zh
		break
	default:
		language = consts.En
	}
	g.Log().Debugf(r.Context(), language)
	gi18n.SetLanguage(language)
	contexts.SetLanguage(r.Context(), language)
}

// DeliverUserContext 将用户信息传递到上下文中
func DeliverUserContext(r *ghttp.Request) (err error) {
	user, err := token.ParseLoginUser(r)
	if err != nil {
		return
	}

	switch user.App {
	case consts.AppAdmin:
		if err = service.BasicsAdminSite().BindUserContext(r.Context(), user); err != nil {
			return
		}
	default:
		contexts.SetUser(r.Context(), user)
	}
	return
}

// MemberDeliverUserContext 将用户信息传递到上下文中
func MemberDeliverUserContext(r *ghttp.Request) (err error) {
	user, err := token.MemberParseLoginUser(r)
	if err != nil {
		return
	}
	//if (g.IsEmpty(gstr.Trim(user.Phone, "-+=")) || g.IsEmpty(gstr.Trim(user.PhoneArea, "-+="))) && g.IsEmpty(gstr.Trim(user.Mail, "-+=")) {
	//	r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "登录已过期，请重新登录"))
	//	r.SetError(err)
	//	return
	//}
	contexts.SetMemberUser(r.Context(), user)
	return
}

// TerminalDeliverUserContext 将用户信息传递到上下文中
func TerminalDeliverUserContext(r *ghttp.Request) (err error) {
	user, err := token.TerminalParseLoginUser(r)
	if err != nil {
		return
	}
	//if (g.IsEmpty(gstr.Trim(user.Phone, "-+=")) || g.IsEmpty(gstr.Trim(user.PhoneArea, "-+="))) && g.IsEmpty(gstr.Trim(user.Mail, "-+=")) {
	//	r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "登录已过期，请重新登录"))
	//	r.SetError(err)
	//	return
	//}
	contexts.SetTerminalUser(r.Context(), user)
	return
}

// TravelStaffDeliverUserContext 将一日游核销人员信息传递到上下文中
func TravelStaffDeliverUserContext(r *ghttp.Request) (err error) {
	user, err := token.TravelStaffParseLoginUser(r)
	if err != nil {
		return
	}
	//if (g.IsEmpty(gstr.Trim(user.Phone, "-+=")) || g.IsEmpty(gstr.Trim(user.PhoneArea, "-+="))) && g.IsEmpty(gstr.Trim(user.Mail, "-+=")) {
	//	r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "登录已过期，请重新登录"))
	//	r.SetError(err)
	//	return
	//}
	contexts.SetTravelStaffUser(r.Context(), user)
	return
}
