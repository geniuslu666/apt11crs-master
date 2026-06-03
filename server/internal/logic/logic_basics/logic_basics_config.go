package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/sms"
	"APT/internal/library/storager"
	"APT/internal/library/token"
	"APT/internal/library/toretaApi"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/simple"
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsConfig struct{}

func NewBasicsConfig() *sBasicsConfig {
	return &sBasicsConfig{}
}

func init() {
	service.RegisterBasicsConfig(NewBasicsConfig())
}

func (s *sBasicsConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail：%+v", err)
	}
}

func (s *sBasicsConfig) LoadConfig(ctx context.Context) (err error) {
	upload, err := s.GetUpload(ctx)
	if err != nil {
		return
	}
	storager.SetConfig(upload)

	sm, err := s.GetSms(ctx)
	if err != nil {
		return
	}
	sms.SetConfig(sm)

	tk, err := s.GetLoadToken(ctx)
	if err != nil {
		return
	}
	token.SetConfig(tk)

	return
}

func (s *sBasicsConfig) GetLogin(ctx context.Context) (conf *model.LoginConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "login"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetPay(ctx context.Context) (conf *model.PayConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "pay"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetSms(ctx context.Context) (conf *model.SmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "sms"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetGeeTest(ctx context.Context) (conf *model.GeeTestConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "geetest"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetSpaSmsConfig(ctx context.Context) (conf *model.OrderSmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "spa_sms_config"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetLanguagePackSetting(ctx context.Context) (conf *model.LanguagePackSetting, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "languagePackSetting"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetCarSmsConfig(ctx context.Context) (conf *model.OrderSmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "car_sms_config"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetYYConfig(ctx context.Context) (conf *model.YYConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "yyconfig"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}
func (s *sBasicsConfig) GetWXShareConfig(ctx context.Context) (conf *model.WXShareConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "wxShare"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetMemberRegRewardConfig(ctx context.Context) (conf *model.MemberRegRewardConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "memberregreward"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetInviteNewRewardConfig(ctx context.Context) (conf *model.InviteNewRewardConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "invitenewreward"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetMemberIntentionConfig(ctx context.Context) (conf *model.MemberIntentionConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "member_intention"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetCarPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "carprintersetting"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetSpaPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "spaprintersetting"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetFoodPrinterSettingConfig(ctx context.Context) (conf *model.PrinterSettingConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "foodprintersetting"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetAppDataBoardViewConfig(ctx context.Context) (conf *model.AppDataBoardViewConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "appdataboardview"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetGeo(ctx context.Context) (conf *model.GeoConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "geo"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetUpload(ctx context.Context) (conf *model.UploadConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "upload"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetPmsPrice(ctx context.Context) (conf *model.PmsPriceConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "pms_price"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "smtp"})
	if err != nil {
		return
	}
	if err = gconv.Scan(models.List, &conf); err != nil {
		return
	}

	conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return
}

func (s *sBasicsConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "basic"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error) {
	err = g.Cfg().MustGet(ctx, "tcp").Scan(&conf)
	return
}

func (s *sBasicsConfig) GetApp(ctx context.Context) (conf *model.AppConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "app"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hggen").Scan(&conf)
	return
}

func (s *sBasicsConfig) GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&conf)
	return
}

func (s *sBasicsConfig) GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "system.log").Scan(&conf)
	return
}

func (s *sBasicsConfig) GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "system.serveLog").Scan(&conf)
	return
}

func (s *sBasicsConfig) GetConfigByGroup(ctx context.Context, in *input_basics.GetConfigInp) (res *input_basics.GetConfigModel, err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}

	var models []*entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).Fields("key", "value", "type").Where("group", in.Group).Scan(&models); err != nil {
		err = gerror.Wrapf(err, "获取配置分组[ %v ]失败，请稍后重试！", in.Group)
		return
	}

	res = new(input_basics.GetConfigModel)
	if len(models) > 0 {
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
		}
	}

	res.List = simple.FilterMaskDemo(ctx, res.List)
	return
}

func (s *sBasicsConfig) ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error) {
	if models == nil {
		err = gerror.New("数据不存在")
		return
	}
	return consts.ConvType(models.Value, models.Type), nil
}

func (s *sBasicsConfig) UpdateConfigByGroup(ctx context.Context, in *input_basics.UpdateConfigInp) (err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}
	var (
		mod    = dao.SysConfig.Ctx(ctx)
		models []*entity.SysConfig
	)

	if err = mod.Where("group", in.Group).Scan(&models); err != nil {
		return
	}

	err = dao.SysConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		for k, v := range in.List {
			row := s.getConfigByKey(k, models)
			if row == nil {
				err = gerror.Newf("暂不支持从前台添加变量，请先在数据库表[%v]中配置变量：%v", dao.SysConfig.Table(), k)
				return
			}

			_, err = dao.SysConfig.Ctx(ctx).Where("id", row.Id).Data(g.Map{"value": v, "updated_at": gtime.Now()}).Update()
			if err != nil {
				return
			}
		}
		return s.syncUpdate(ctx, in)
	})

	if err != nil {
		return
	}

	return
}

func (s *sBasicsConfig) SpaSmsConfigUpdateReq(ctx context.Context, in *input_basics.UpdateSpaSmsConfigInp) (err error) {
	var updateConfigInp input_basics.UpdateConfigInp
	updateConfigInp.Group = in.Group
	updateConfigInp.List = in.List

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 修改按摩短信配置
		err = s.UpdateConfigByGroup(ctx, &updateConfigInp)
		if err != nil {
			return
		}

		// 修改按摩服务商短信配置
		for _, IspItem := range in.IspList {
			if _, err = dao.SpaIsp.Ctx(ctx).TX(tx).Where(dao.SpaIsp.Columns().Id, IspItem.Id).Data(g.MapStrAny{
				dao.SpaIsp.Columns().SendSmsPhone: IspItem.SendSmsPhone,
			}).Update(); err != nil {
				return
			}
		}
		return
	})
}

func (s *sBasicsConfig) getConfigByKey(key string, models []*entity.SysConfig) *entity.SysConfig {
	if len(models) == 0 {
		return nil
	}

	for _, v := range models {
		if key == v.Key {
			return v
		}
	}
	return nil
}

func (s *sBasicsConfig) syncUpdate(ctx context.Context, in *input_basics.UpdateConfigInp) (err error) {
	switch in.Group {
	case "upload":
		upload, err := s.GetUpload(ctx)
		if err == nil {
			storager.SetConfig(upload)
		}
	case "sms":
		sm, err := s.GetSms(ctx)
		if err == nil {
			sms.SetConfig(sm)
		}
	}

	if err != nil {
		err = gerror.Newf("syncUpdate %v conifg fail：%+v", in.Group, err.Error())
	}
	return
}

func (s *sBasicsConfig) ClusterSync(ctx context.Context, message *gredis.Message) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Errorf(ctx, "ClusterSync fail：%+v", err)
	}
}

func (s *sBasicsConfig) UpdateOrderMember(ctx context.Context) (err error) {
	// 更新接送机订单用户ID
	carMemberIds, err := s.getMemberIdsByBatch(ctx, "hg_car_order")
	if err != nil {
		err = gerror.Wrap(err, "获取接送机订单用户ID失败")
		return
	}

	_, err = g.DB().Exec(ctx, `
		UPDATE hg_sys_config 
		SET value = ?
		WHERE `+"`group`"+` = 'wxminAudit' AND `+"`key`"+` = 'carOrderMemberIds'
	`, carMemberIds)
	if err != nil {
		err = gerror.Wrap(err, "更新接送机订单用户ID失败")
		return
	}

	// 更新按摩订单用户ID
	spaMemberIds, err := s.getMemberIdsByBatch(ctx, "hg_spa_order")
	if err != nil {
		err = gerror.Wrap(err, "获取按摩订单用户ID失败")
		return
	}

	_, err = g.DB().Exec(ctx, `
		UPDATE hg_sys_config 
		SET value = ?
		WHERE `+"`group`"+` = 'wxminAudit' AND `+"`key`"+` = 'spaOrderMemberIds'
	`, spaMemberIds)
	if err != nil {
		err = gerror.Wrap(err, "更新按摩订单用户ID失败")
		return
	}

	g.Log().Info(ctx, "订单用户ID更新成功")
	return
}

// getMemberIdsByBatch 分批获取用户ID并拼接
func (s *sBasicsConfig) getMemberIdsByBatch(ctx context.Context, tableName string) (memberIds string, err error) {
	const batchSize = 10000 // 每批处理10000条记录
	var offset int
	var allMemberIds []int
	memberIdMap := make(map[int]bool) // 用于去重

	for {
		// 分批查询用户ID
		sql := fmt.Sprintf(`
			SELECT DISTINCT member_id 
			FROM %s 
			WHERE member_id IS NOT NULL AND member_id > 0 
			ORDER BY member_id ASC 
			LIMIT %d OFFSET %d
		`, tableName, batchSize, offset)

		result, err := g.DB().Query(ctx, sql)
		if err != nil {
			return "", gerror.Wrapf(err, "查询表 %s 失败", tableName)
		}

		// 转换为整数数组
		var batchMemberIds []int
		for _, record := range result {
			memberId := record["member_id"].Int()
			if memberId > 0 {
				// 去重处理
				if !memberIdMap[memberId] {
					memberIdMap[memberId] = true
					batchMemberIds = append(batchMemberIds, memberId)
				}
			}
		}

		// 如果本批次没有数据，说明已经处理完所有数据
		if len(batchMemberIds) == 0 {
			break
		}

		// 添加到总列表
		allMemberIds = append(allMemberIds, batchMemberIds...)

		// 如果本批次数据少于批次大小，说明已经是最后一批
		if len(batchMemberIds) < batchSize {
			break
		}

		offset += batchSize

		g.Log().Infof(ctx, "已处理表 %s 的 %d 条记录", tableName, offset)
	}

	// 排序并转换为字符串
	sort.Ints(allMemberIds)

	// 转换为字符串数组
	memberIdStrs := make([]string, len(allMemberIds))
	for i, id := range allMemberIds {
		memberIdStrs[i] = fmt.Sprintf("%d", id)
	}

	// 用逗号连接
	memberIds = strings.Join(memberIdStrs, ",")

	g.Log().Infof(ctx, "表 %s 共获取到 %d 个唯一用户ID", tableName, len(allMemberIds))

	return memberIds, nil
}

func (s *sBasicsConfig) GetCabinetApi(ctx context.Context) (conf *model.CabinetApiConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "cabinetApi"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetToretaApi(ctx context.Context) (conf *model.ToretaApiConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "toretaApi"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetXpyunApi(ctx context.Context) (conf *model.XpyunConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "xpyun"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) RefreshToretaAccessToken(ctx context.Context) (err error) {
	var (
		toretaConfig *model.ToretaApiConfig
		toretaClient *toretaApi.ToretaClient
		refreshResp  *toretaApi.RefreshOATokenResponse
	)

	// 获取当前 TORETA 配置
	if toretaConfig, err = s.GetToretaApi(ctx); err != nil {
		return gerror.Wrap(err, "获取TORETA配置失败")
	}

	// 检查是否有 RefreshToken
	if toretaConfig.RefreshToken == "" {
		return gerror.New("RefreshToken为空，无法刷新AccessToken")
	}

	// 创建 TORETA 客户端
	toretaClient = toretaApi.NewClient(ctx, toretaConfig)

	// 调用刷新 Token API
	if refreshResp, err = toretaClient.RefreshOAToken(ctx, &toretaApi.RefreshOATokenParams{
		RefreshToken: toretaConfig.RefreshToken,
	}); err != nil {
		return gerror.Wrap(err, "调用TORETA刷新Token API失败")
	}

	// 更新配置
	// 计算具体的过期时间戳：当前时间 + 有效期秒数
	expiresAt := gtime.Now().Add(time.Duration(refreshResp.ExpiresIn) * time.Second).Unix()
	updateData := map[string]interface{}{
		"accessToken": refreshResp.AccessToken,
		"expiresAt":   expiresAt,
	}

	if err = s.UpdateConfigByGroup(ctx, &input_basics.UpdateConfigInp{
		Group: "toretaApi",
		List:  updateData,
	}); err != nil {
		return gerror.Wrap(err, "更新TORETA配置失败")
	}

	g.Log().Infof(ctx, "TORETA AccessToken刷新成功，新Token过期时间: %d秒", refreshResp.ExpiresIn)
	return
}

func (s *sBasicsConfig) GetSystemMessageConfig(ctx context.Context) (conf *model.SystemMessageConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "systemMessage"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

func (s *sBasicsConfig) GetHotelSettingConfig(ctx context.Context) (conf *model.HotelSettingConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &input_basics.GetConfigInp{Group: "hotelSetting"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}
