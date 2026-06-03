package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type UploadFileReq struct {
	g.Meta `path:"/upload/file" tags:"ADMIN" summary:"附件_上传附件" method:"post"`
	Kind   string `json:"kind" dc:"类型"`
}

type UploadFileRes struct {
	*input_basics.AttachmentListModel
}

type UploadCheckMultipartReq struct {
	g.Meta `path:"/upload/checkMultipart" tags:"ADMIN" summary:"附件_检查文件分片" method:"post"`
	input_basics.CheckMultipartInp
}

type UploadCheckMultipartRes struct {
	*input_basics.CheckMultipartModel
}

type UploadPartReq struct {
	g.Meta `path:"/upload/uploadPart" tags:"ADMIN" summary:"附件_分片上传" method:"post"`
	input_basics.UploadPartInp
}

type UploadPartRes struct {
	*input_basics.UploadPartModel
}
