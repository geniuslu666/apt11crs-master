package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsTemplateUpdateFields 修改场景模版字段过滤
type SmsTemplateUpdateFields struct {
	Scene           int         `json:"scene"           dc:"1会员  2住宿  3接送机  4按摩"`
	Event           string      `json:"event"           dc:"别名"`
	Title           string      `json:"title"           dc:"标题"`
	Content         string      `json:"content"         dc:"内容（多语言）"`
	UmsTemplate     *gjson.Json `json:"umsTemplate"    dc:"一信通模版json（id：模版id， content：模版内容）"`
	TencentTemplate *gjson.Json `json:"tencentTemplate" dc:"腾讯云模版json"`
	AliyunTemplate  *gjson.Json `json:"aliyunTemplate" dc:"阿里云模版json"`
}

// SmsTemplateInsertFields 新增场景模版字段过滤
type SmsTemplateInsertFields struct {
	Scene           int         `json:"scene"           dc:"1会员  2住宿  3接送机  4按摩"`
	Event           string      `json:"event"           dc:"别名"`
	Title           string      `json:"title"           dc:"标题"`
	Content         string      `json:"content"         dc:"内容（多语言）"`
	UmsTemplate     *gjson.Json `json:"umsTemplate"    dc:"一信通模版json（id：模版id， content：模版内容）"`
	TencentTemplate *gjson.Json `json:"tencentTemplate" dc:"腾讯云模版json"`
	AliyunTemplate  *gjson.Json `json:"aliyunTemplate" dc:"阿里云模版json"`
}

// SmsTemplateEditInp 修改/新增场景模版
type SmsTemplateEditInp struct {
	entity.SysSmsTemplate
	ContentLanguage input_language.LanguageModel `json:"contentLanguage"          dc:"多语言内容"`
}

func (in *SmsTemplateEditInp) Filter(ctx context.Context) (err error) {
	return
}

type SmsTemplateEditModel struct{}

// SmsTemplateDeleteInp 删除场景模版
type SmsTemplateDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SmsTemplateDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SmsTemplateDeleteModel struct{}

// SmsTemplateViewInp 获取指定场景模版信息
type SmsTemplateViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SmsTemplateViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SmsTemplateViewModel struct {
	entity.SysSmsTemplate
	ContentLanguage []*input_hotel.LanguageType `json:"contentLanguage"         dc:"通知内容"   orm:"with:uuid=content"`
}

// SmsTemplateListInp 获取场景模版列表
type SmsTemplateListInp struct {
	input_form.PageReq
	Scene int `json:"scene"           dc:"1会员  2住宿  3接送机  4按摩"`
}

func (in *SmsTemplateListInp) Filter(ctx context.Context) (err error) {
	return
}

type SmsTemplateListModel struct {
	Id        int         `json:"id"                        dc:"id"`
	Scene     int         `json:"scene"           dc:"1会员  2住宿  3接送机  4按摩"`
	Event     string      `json:"event"           dc:"别名"`
	Title     string      `json:"title"           dc:"标题"`
	Content   string      `json:"content"         dc:"内容（多语言）"`
	CreatedAt *gtime.Time `json:"createdAt"                 dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt"                 dc:"updated_at"`
}

// SmsTemplateLogInp 获取场景模版发送日志列表
type SmsTemplateLogInp struct {
	input_form.PageReq
	Type uint `json:"type"           dc:"1短信   2邮件   3推送"`
}

func (in *SmsTemplateLogInp) Filter(ctx context.Context) (err error) {
	return
}

type SmsTemplateLogModel struct {
	Id             int64       `json:"id"             dc:"主键"`
	TemplateId     uint        `json:"templateId"     dc:"系统模版表中主键ID"`
	SendTemplateId string      `json:"sendTemplateId" dc:"发送的模版ID"`
	Type           uint        `json:"type"           dc:"1短信   2邮件   3推送"`
	To             string      `json:"to"             dc:"短信或邮件或推送头"`
	VipId          uint        `json:"vipId"          dc:"会员ID"`
	Content        string      `json:"content"       dc:"发送内容"`
	CreatedAt      *gtime.Time `json:"createdAt"                 dc:"created_at"`
	UpdatedAt      *gtime.Time `json:"updatedAt"                 dc:"updated_at"`
}

// SendTemplateInp 发送 短信
type SendTemplateInp struct {
	OrderSn string `json:"orderSn"     dc:"订单号"` // 必填
	Event   string `json:"event"     dc:"事件"`    // 必填
}

// SendTemplateModel 发送 短信
type SendTemplateModel struct {
	VipId             uint     `json:"vipId"          dc:"会员ID"`
	AreaNo            string   `json:"areaNo"    dc:"区号"`  // 必填
	Mobile            string   `json:"mobile"    dc:"手机号"` // 必填
	TemplateId        string   `json:"templateId"    dc:"模版ID"`
	TemplateContent   string   `json:"templateContent"    dc:"模版内容"`
	TemplateParamsVal []string `json:"templateParams"      dc:"短信模板参数值数组"`
}

type AbroadTemplateModel map[string]AbroadTemplateItemModel

type AbroadTemplateItemModel struct {
	Id    string `json:"id"`
	Param string `json:"param"`
}
