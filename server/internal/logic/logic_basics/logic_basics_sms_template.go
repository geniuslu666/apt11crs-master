package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/library/sms"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_language"
	"APT/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sBasicsSmsTemplate struct{}

func NewBasicsSmsTemplate() *sBasicsSmsTemplate {
	return &sBasicsSmsTemplate{}
}

func init() {
	service.RegisterBasicsSmsTemplate(NewBasicsSmsTemplate())
}

func (s *sBasicsSmsTemplate) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysSmsTemplate.Ctx(ctx), option...)
}

func (s *sBasicsSmsTemplate) List(ctx context.Context, in *input_basics.SmsTemplateListInp) (list []*input_basics.SmsTemplateListModel, totalCount int, err error) {
	mod := s.Model(ctx).Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_basics.SmsTemplateListModel{})

	if !g.IsEmpty(in.Scene) {
		mod = mod.Where(dao.SysSmsTemplate.Columns().Scene, in.Scene)
	}

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.SysSmsTemplate.Columns().Id)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取通知模版列表失败，请稍后重试！")
			return
		}
	} else {

		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取通知模版列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sBasicsSmsTemplate) Edit(ctx context.Context, in *input_basics.SmsTemplateEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object             gdb.Record
			SmsTemplateContDao *input_language.LoadLanguage
			LanguageStruct     input_language.LanguageModel
		)
		Uuid := guid.S([]byte("content"))

		if in.Id > 0 {

			if Object, err = dao.SysSmsTemplate.Ctx(ctx).Where(dao.SysSmsTemplate.Columns().Id, in.Id).One(); err != nil {
				return
			}
			if !g.IsEmpty(Object["content"]) {
				Uuid = Object["content"].String()
			}

			SmsTemplateContDao = &input_language.LoadLanguage{
				Uuid: Uuid,
				Tag:  dao.SysSmsTemplate.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Content"),
			}
			LanguageStruct = in.ContentLanguage
			if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, SmsTemplateContDao); err != nil {
				return
			}

			in.Content = Uuid

			if _, err = s.Model(ctx).
				Fields(input_basics.SmsTemplateUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改通知模版失败，请稍后重试！")
			}
			return
		}

		in.Content = Uuid

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.SmsTemplateInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增通知模版失败，请稍后重试！")
		}

		SmsTemplateContDao = &input_language.LoadLanguage{
			Uuid: Uuid,
			Tag:  dao.SysSmsTemplate.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Content"),
		}
		LanguageStruct = in.ContentLanguage
		if err = service.BasicsLanguage().Sync(ctx, LanguageStruct, SmsTemplateContDao); err != nil {
			return
		}
		return
	})
}

func (s *sBasicsSmsTemplate) Delete(ctx context.Context, in *input_basics.SmsTemplateDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除通知模版失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsSmsTemplate) View(ctx context.Context, in *input_basics.SmsTemplateViewInp) (res *input_basics.SmsTemplateViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取通知模版信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsSmsTemplate) Log(ctx context.Context, in *input_basics.SmsTemplateLogInp) (list []*input_basics.SmsTemplateLogModel, totalCount int, err error) {
	mod := dao.SysSmsTemplateLog.Ctx(ctx)

	mod = mod.Fields(input_basics.SmsTemplateLogModel{})

	if !g.IsEmpty(in.Type) {
		mod = mod.Where(dao.SysSmsTemplateLog.Columns().Type, in.Type)
	}

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.SysSmsTemplateLog.Columns().Id)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取发送记录失败，请稍后重试！")
			return
		}
	} else {

		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取发送记录失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sBasicsSmsTemplate) SendTemplate(ctx context.Context, in *input_basics.SendTemplateInp) (err error) {
	var (
		language          string
		paramsArr         []string
		SmsTemplate       *input_basics.SmsTemplateViewModel
		SendTemplateModel *input_basics.SendTemplateModel
		out               *input_basics.SendMsgInp
		UmsTemplate       *struct {
			Id      string `json:"id"`
			Content string `json:"content"`
			Param   string `json:"param"`
		}
		AbroadTemplate *input_basics.AbroadTemplateModel
	)
	out = new(input_basics.SendMsgInp)
	SendTemplateModel = new(input_basics.SendTemplateModel)
	if g.IsEmpty(in.Event) {
		err = gerror.New("事件不能为空")
		return
	}

	if g.IsEmpty(in.OrderSn) {
		err = gerror.New("事件不能为空")
		return
	}

	//out.Mobile = in.Mobile
	//out.Event = in.Event
	//out.TemplateParams = in.TemplateParams

	g.Log().Info(ctx, gjson.New(in))

	if err = s.Model(ctx).WithAll().Where(dao.SysSmsTemplate.Columns().Event, in.Event).Scan(&SmsTemplate); err != nil {
		err = gerror.Wrap(err, "获取通知模版信息失败")
		return
	}

	if in.OrderSn[:1] == "C" {
		if SendTemplateModel, err = s.getCarTemplate(ctx, in); err != nil {
			err = gerror.Wrap(err, "获取通知模版信息失败")
			return
		}
	} else if in.OrderSn[:1] == "S" {
		if SendTemplateModel, err = s.getSpaTemplate(ctx, in); err != nil {
			err = gerror.Wrap(err, "获取通知模版信息失败")
			return
		}
	} else if in.OrderSn[:1] == "H" {
		if SendTemplateModel, err = s.getHotelTemplate(ctx, in); err != nil {
			err = gerror.Wrap(err, "获取通知模版信息失败")
			return
		}
		if in.Event == "hotel_order_pay" && SendTemplateModel.AreaNo != "+86" {
			err = gerror.New("只发送国内短信")
			return
		}
	}

	out.Mobile = SendTemplateModel.Mobile

	config, err := service.BasicsConfig().GetSms(ctx)
	if err != nil {
		return
	}

	if SendTemplateModel.AreaNo == "+86" {
		config.SmsDrive = config.NoticeSmsDriveCN
	} else {
		config.SmsDrive = config.NoticeSmsDriveAbroad
	}

	if config.SmsDrive == "ums" {
		if err = json.Unmarshal([]byte(gvar.New(SmsTemplate.UmsTemplate).String()), &UmsTemplate); err != nil {
			err = gerror.New("解析路径失败")
			return
		}
		out.Content = UmsTemplate.Content

		if !g.IsEmpty(UmsTemplate.Param) {
			paramsArr = strings.Split(UmsTemplate.Param, ",")
			for index, para := range paramsArr {
				out.Content = gstr.Replace(out.Content, "{"+para+"}", SendTemplateModel.TemplateParamsVal[index])
			}
		}
	} else if config.SmsDrive == "tencent" {
		if err = json.Unmarshal([]byte(gvar.New(SmsTemplate.TencentTemplate).String()), &AbroadTemplate); err != nil {
			err = gerror.New("解析路径失败")
			return
		}
	} else if config.SmsDrive == "aliyun" {
		if err = json.Unmarshal([]byte(gvar.New(SmsTemplate.AliyunTemplate).String()), &AbroadTemplate); err != nil {
			err = gerror.New("解析路径失败")
			return
		}
	}

	switch SendTemplateModel.AreaNo {
	case "+86":
		language = "zh"
		out.Template = UmsTemplate.Id
		break
	case "+886", "+852", "+853":
		language = "zh_CN"
		// tw
		out.Template = (*AbroadTemplate)["tw"].Id
		paramsArr = strings.Split((*AbroadTemplate)["tw"].Param, ",")
		break
	case "+82":
		language = "ko"
		// ko
		out.Template = (*AbroadTemplate)["ko"].Id
		paramsArr = strings.Split((*AbroadTemplate)["ko"].Param, ",")
		break
	case "+81":
		language = "ja"
		// ja
		out.Template = (*AbroadTemplate)["ja"].Id
		paramsArr = strings.Split((*AbroadTemplate)["ja"].Param, ",")
		break
	default:
		language = "en"
		// en
		out.Template = (*AbroadTemplate)["en"].Id
		paramsArr = strings.Split((*AbroadTemplate)["en"].Param, ",")
		break
	}

	if language != "zh" {
		for _, v := range SmsTemplate.ContentLanguage {
			if v.Language == language {
				out.Content = v.Content
			}
		}
	}

	if !g.IsEmpty(AbroadTemplate) {
		TemplateParams := make(map[string]string)
		for i := 0; i < len(SendTemplateModel.TemplateParamsVal) && i < len(paramsArr); i++ {
			TemplateParams[paramsArr[i]] = SendTemplateModel.TemplateParamsVal[i]
		}
		out.TemplateParams = TemplateParams
	}

	g.Log().Info(ctx, config.SmsDrive)
	g.Log().Info(ctx, gjson.New(out))

	if err = sms.New(config.SmsDrive).SendMsg(ctx, out); err != nil {
		return
	}

	var data = new(entity.SysSmsTemplateLog)
	data.TemplateId = uint(SmsTemplate.Id)
	data.SendTemplateId = out.Template
	data.Type = 1
	data.To = out.Mobile
	data.VipId = SendTemplateModel.VipId
	data.Content = out.Content
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysSmsTemplateLog.Ctx(ctx).Data(data).OmitEmptyData().Insert()
	return
}

func (s *sBasicsSmsTemplate) getCarTemplate(ctx context.Context, in *input_basics.SendTemplateInp) (out *input_basics.SendTemplateModel, err error) {
	var (
		language string
		sceneMap map[string]string
		CarOrder *entity.CarOrder
	)

	out = new(input_basics.SendTemplateModel)

	// 查询 app 订单信息
	if err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).
		Scan(&CarOrder); err != nil {
		return
	}
	if g.IsEmpty(CarOrder) {
		err = gerror.New("出行订单数据为空")
		return
	}
	out.VipId = CarOrder.MemberId
	out.AreaNo = CarOrder.PhoneArea
	out.Mobile = CarOrder.PhoneArea + CarOrder.BookingMobile

	switch CarOrder.PhoneArea {
	case "+86":
		language = "zh"
		break
	case "+886", "+852", "+853":
		// tw
		language = "tw"
		break
	case "+82":
		// ko
		language = "ko"
		break
	case "+81":
		// ja
		language = "ja"
		break
	default:
		// en
		language = "en"
		break
	}

	switch in.Event {
	case "car_order_confirm":
		if CarOrder.ServiceType == "PICKUP" {
			sceneMap = map[string]string{
				"zh": "接机",
				"en": "Airport Pickup",
				"ja": "空港お迎え",
				"ko": "공항 픽업",
				"tw": "接機",
			}
		} else if CarOrder.ServiceType == "DELIVERY" {
			sceneMap = map[string]string{
				"zh": "送机",
				"en": "Drop-off",
				"ja": "空港お送り",
				"ko": "공항 배웅",
				"tw": "送機",
			}
		} else {
			sceneMap = map[string]string{
				"zh": "包车",
				"en": "Charter",
				"ja": "貸切",
				"ko": "전용 차량",
				"tw": "包車",
			}
		}
		out.TemplateParamsVal = []string{sceneMap[language]}
		break
	case "car_order_dispatch":
		out.TemplateParamsVal = []string{CarOrder.BookDate}
		break
	default:
		break
	}

	return
}

func (s *sBasicsSmsTemplate) getSpaTemplate(ctx context.Context, in *input_basics.SendTemplateInp) (out *input_basics.SendTemplateModel, err error) {
	var (
		SpaOrder *entity.SpaOrder
	)

	out = new(input_basics.SendTemplateModel)

	// 查询 app 订单信息
	if err = dao.CarOrder.Ctx(ctx).
		Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).
		Scan(&SpaOrder); err != nil {
		return
	}
	if g.IsEmpty(SpaOrder) {
		err = gerror.New("按摩订单数据为空")
		return
	}
	out.VipId = SpaOrder.MemberId
	out.AreaNo = SpaOrder.PhoneArea
	out.Mobile = SpaOrder.PhoneArea + SpaOrder.BookingMobile

	return
}

func (s *sBasicsSmsTemplate) getHotelTemplate(ctx context.Context, in *input_basics.SendTemplateInp) (out *input_basics.SendTemplateModel, err error) {
	var (
		AppStay   *entity.PmsAppStay
		GuestInfo *entity.PmsGuestProfile
	)

	out = new(input_basics.SendTemplateModel)

	// 查询 app 订单信息
	if err = dao.PmsAppStay.Ctx(ctx).
		Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).
		Scan(&AppStay); err != nil {
		return
	}
	if g.IsEmpty(AppStay) {
		err = gerror.New("住宿订单数据为空")
		return
	}
	if err = dao.PmsGuestProfile.Ctx(ctx).
		Where(dao.PmsGuestProfile.Columns().Uid, AppStay.Booker).
		Scan(&GuestInfo); err != nil {
		return
	}
	if g.IsEmpty(GuestInfo) {
		err = gerror.New("入住客人信息数据为空")
		return
	}
	out.VipId = uint(AppStay.MemberId)
	out.AreaNo = GuestInfo.AreaNo
	out.Mobile = GuestInfo.AreaNo + GuestInfo.Phone

	return
}
