package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type UploadImageReq struct {
	g.Meta `path:"/upload/file" method:"post" tags:"APP_BASICS" summary:"[上传]附件"`
	Kind   string `json:"kind" dc:"类型"`
}

type UploadImageRes struct {
	*input_basics.AttachmentListModel
}
