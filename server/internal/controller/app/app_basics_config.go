package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"APT/api/app/basics"
)

func (c *ControllerBasics) GetAppConfig(ctx context.Context, req *basics.GetAppConfigReq) (res *basics.GetAppConfigRes, err error) {
	var (
		r              = ghttp.RequestFromCtx(ctx)
		XVersion       = r.GetHeader("X-Version")
		XChannel       = r.GetHeader("x-channel")
		XLaguage       = contexts.GetLanguage(ctx)
		AppVersionInfo []*model.AppVersionInfo
	)
	if g.IsEmpty(XChannel) {
		XChannel = "history"
	}

	res = new(basics.GetAppConfigRes)
	if res.AppConfig, err = service.BasicsConfig().GetApp(ctx); err != nil {
		return
	}
	if err = gjson.New(res.AppConfig.AppVersion).Scan(&AppVersionInfo); err != nil {
		return
	}

	for _, v := range AppVersionInfo {
		if v.Name == XChannel {
			res.AppConfig.DownloadUrl = v.Info.AppDownload
			res.AppConfig.AppOpen = v.Info.AppOpen
			res.AppConfig.Appforce = v.Info.AppForcedUpload
			res.AppConfig.PayModelList = v.Info.AppPayModel
			res.AppConfig.AppApply = v.Info.AppAudit
			res.AppConfig.AppVersionContent = v.Info.AppVersionContent
			res.AppConfig.AppVersion = v.Info.AppVersion
			switch XLaguage {
			case consts.Zh:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyZh
				break
			case consts.ZhCN:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyZhCN
				break
			case consts.En:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyEn
				break
			case consts.Ja:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyJa
				break
			case consts.Ko:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyKo
				break
			default:
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyZh
				break
			}
			if g.IsEmpty(res.AppConfig.AppPrivacy) {
				res.AppConfig.AppPrivacy = v.Info.AppPrivacyZh
			}
		}
	}
	v1, _ := semver.Make(res.AppConfig.AppVersion)
	v2, _ := semver.Make(XVersion)
	if !g.IsEmpty(XVersion) {
		if v2.LT(v1) {
			res.AppConfig.IsUpdate = true
		}
	}
	return
}

func (c *ControllerBasics) GetPayModeConfig(ctx context.Context, req *basics.GetPayModeConfigReq) (res *basics.GetPayModeConfigRes, err error) {
	var (
		payModeVar *gvar.Var
	)
	if payModeVar, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, "paymode").
		Value(dao.SysConfig.Columns().Value); err != nil {
	}
	if payModeVar == nil {
		return
	}
	res = new(basics.GetPayModeConfigRes)
	res.PayModeList = strings.Split(payModeVar.String(), ",")
	return
}
func (c *ControllerBasics) GetAppLanguageConfig(ctx context.Context, req *basics.GetAppLanguageConfigReq) (res *basics.GetAppLanguageConfigRes, err error) {
	var (
		LanguagePackSetting *model.LanguagePackSetting
	)
	if LanguagePackSetting, err = service.BasicsConfig().GetLanguagePackSetting(ctx); err != nil {
		return
	}
	res = new(basics.GetAppLanguageConfigRes)
	res.LanguageVersion = LanguagePackSetting.LanguagePackVersion
	res.LanguageVersionUrl = LanguagePackSetting.LanguagePackFilePath
	return
}

func (c *ControllerBasics) GetContactConfig(ctx context.Context, req *basics.GetContactConfigReq) (res *basics.GetContactConfigRes, err error) {
	var (
		contactDataVar *gvar.Var
	)
	res = new(basics.GetContactConfigRes)

	// 从配置表获取联系我们配置
	if contactDataVar, err = dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Group, "contactCenter").
		Where(dao.SysConfig.Columns().Key, "contactData").
		Value(dao.SysConfig.Columns().Value); err != nil {
		return
	}
	if contactDataVar == nil || contactDataVar.IsEmpty() {
		return
	}

	// 解析 JSON 字符串
	var contactConfig model.ContactCenterConfig
	if err = gjson.New(contactDataVar.String()).Scan(&contactConfig); err != nil {
		return
	}

	if contactConfig.Japan != nil {
		res.Japan = &basics.ContactInfo{
			Phone:    contactConfig.Japan.Phone,
			Email:    contactConfig.Japan.Email,
			WorkTime: contactConfig.Japan.WorkTime,
		}
	}
	if contactConfig.China != nil {
		res.China = &basics.ContactInfo{
			Phone:    contactConfig.China.Phone,
			Email:    contactConfig.China.Email,
			WorkTime: contactConfig.China.WorkTime,
		}
	}

	return
}
