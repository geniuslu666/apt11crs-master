package basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type AttachmentListReq struct {
	g.Meta `path:"/attachment/list" method:"get" tags:"ADMIN" summary:"管理/附件_获取附件列表"`
	input_basics.AttachmentListInp
}

type AttachmentListRes struct {
	List []*input_basics.AttachmentListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type AttachmentViewReq struct {
	g.Meta `path:"/attachment/view" method:"get" tags:"ADMIN" summary:"附件_获取指定附件信息"`
	input_basics.AttachmentViewInp
}

type AttachmentViewRes struct {
	*input_basics.AttachmentViewModel
}

type AttachmentDeleteReq struct {
	g.Meta `path:"/attachment/delete" method:"post" tags:"ADMIN" summary:"附件_删除附件"`
	input_basics.AttachmentDeleteInp
}

type AttachmentDeleteRes struct{}

type AttachmentClearKindReq struct {
	g.Meta `path:"/attachment/clearKind" method:"post" tags:"ADMIN" summary:"附件_清空上传类型"`
	input_basics.AttachmentClearKindInp
}

type AttachmentClearKindRes struct{}

type AttachmentChooserOptionReq struct {
	g.Meta `path:"/attachment/chooserOption" method:"get" tags:"ADMIN" summary:"附件_获取选择器选项"`
}

type AttachmentChooserOptionRes struct {
	Drive input_basics.DataSelectModel `json:"drive" dc:"驱动"`
	Kind  []entity.SysAttachmentKind   `json:"kind"  dc:"上传类型"`
}

type AttachmentKindSelect struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Tag   string `json:"listClass"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
}
