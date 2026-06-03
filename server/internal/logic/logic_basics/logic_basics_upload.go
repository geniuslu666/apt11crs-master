package logic_basics

import (
	"APT/internal/library/storager"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/format"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sBasicsUpload struct{}

func NewBasicsUpload() *sBasicsUpload {
	return &sBasicsUpload{}
}

func init() {
	service.RegisterBasicsUpload(NewBasicsUpload())
}

// UploadFile 上传文件
func (s *sBasicsUpload) UploadFile(ctx context.Context, uploadType string, file *ghttp.UploadFile, kind string) (res *input_basics.AttachmentListModel, err error) {
	attachment, err := storager.DoUpload(ctx, uploadType, file, kind)
	if err != nil {
		return
	}

	attachment.FileUrl = storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	res = &input_basics.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	return
}

// CheckMultipart 检查文件分片
func (s *sBasicsUpload) CheckMultipart(ctx context.Context, in *input_basics.CheckMultipartInp) (res *input_basics.CheckMultipartModel, err error) {
	data, err := storager.CheckMultipart(ctx, in.CheckMultipartParams)
	if err != nil {
		return nil, err
	}
	res = new(input_basics.CheckMultipartModel)
	res.CheckMultipartModel = data
	return
}

// UploadPart 上传分片
func (s *sBasicsUpload) UploadPart(ctx context.Context, in *input_basics.UploadPartInp) (res *input_basics.UploadPartModel, err error) {
	data, err := storager.UploadPart(ctx, in.UploadPartParams)
	if err != nil {
		return nil, err
	}
	res = new(input_basics.UploadPartModel)
	res.UploadPartModel = data
	return
}
