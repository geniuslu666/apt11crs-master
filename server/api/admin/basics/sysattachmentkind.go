package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

type SysAttachmentKindListReq struct {
	g.Meta `path:"/sysAttachmentKind/list" method:"get" tags:"ADMIN" summary:"获取附件分类列表"`
	input_basics.SysAttachmentKindListInp
}

type SysAttachmentKindListRes struct {
	input_form.PageRes
	List []*input_basics.SysAttachmentKindListModel `json:"list"   dc:"数据列表"`
}

type SysAttachmentKindViewReq struct {
	g.Meta `path:"/sysAttachmentKind/view" method:"get" tags:"ADMIN" summary:"获取附件分类指定信息"`
	input_basics.SysAttachmentKindViewInp
}

type SysAttachmentKindViewRes struct {
	*input_basics.SysAttachmentKindViewModel
}

type SysAttachmentKindEditReq struct {
	g.Meta `path:"/sysAttachmentKind/edit" method:"post" tags:"ADMIN" summary:"修改/新增附件分类"`
	input_basics.SysAttachmentKindEditInp
}

type SysAttachmentKindEditRes struct{}

type SysAttachmentKindDeleteReq struct {
	g.Meta `path:"/sysAttachmentKind/delete" method:"post" tags:"ADMIN" summary:"删除附件分类"`
	input_basics.SysAttachmentKindDeleteInp
}

type SysAttachmentKindDeleteRes struct{}
