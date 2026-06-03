package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type SmsTemplateListReq struct {
	g.Meta `path:"/smsTemplate/list" method:"get" tags:"ADMIN_PMS" summary:"场景通知模版_列表"`
	input_basics.SmsTemplateListInp
}

type SmsTemplateListRes struct {
	input_form.PageRes
	List []*input_basics.SmsTemplateListModel `json:"list"   dc:"数据列表"`
}

type SmsTemplateViewReq struct {
	g.Meta `path:"/smsTemplate/view" method:"get" tags:"ADMIN_PMS" summary:"场景通知模版_详情"`
	input_basics.SmsTemplateViewInp
}

type SmsTemplateViewRes struct {
	*input_basics.SmsTemplateViewModel
}

type SmsTemplateEditReq struct {
	g.Meta `path:"/smsTemplate/edit" method:"post" tags:"ADMIN_PMS" summary:"场景通知模版_修改/新增"`
	input_basics.SmsTemplateEditInp
}

type SmsTemplateEditRes struct{}

type SmsTemplateDeleteReq struct {
	g.Meta `path:"/smsTemplate/delete" method:"post" tags:"ADMIN_PMS" summary:"场景通知模版_删除"`
	input_basics.SmsTemplateDeleteInp
}

type SmsTemplateDeleteRes struct{}

type SmsTemplateLogReq struct {
	g.Meta `path:"/smsTemplate/log" method:"get" tags:"ADMIN_PMS" summary:"场景通知模版_发送记录"`
	input_basics.SmsTemplateLogInp
}

type SmsTemplateLogRes struct {
	input_form.PageRes
	List []*input_basics.SmsTemplateLogModel `json:"list"   dc:"数据列表"`
}
