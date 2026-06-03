package admin

import (
	"APT/internal/library/storager"
	"APT/internal/service"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) UploadFile(ctx context.Context, req *basics.UploadFileReq) (res *basics.UploadFileRes, err error) {

	r := g.RequestFromCtx(ctx)
	uploadType := r.Header.Get("uploadType")
	if uploadType != "default" && !validate.InSlice(storager.KindSlice, uploadType) {
		err = gerror.New("上传类型是无效的")
		return
	}

	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}
	res = new(basics.UploadFileRes)
	res.AttachmentListModel, err = service.BasicsUpload().UploadFile(ctx, uploadType, file, req.Kind)
	if err != nil {
		return
	}
	return
}
func (c *ControllerBasics) UploadCheckMultipart(ctx context.Context, req *basics.UploadCheckMultipartReq) (res *basics.UploadCheckMultipartRes, err error) {
	data, err := service.BasicsUpload().CheckMultipart(ctx, &req.CheckMultipartInp)
	if err != nil {
		return nil, err
	}
	res = new(basics.UploadCheckMultipartRes)
	res.CheckMultipartModel = data
	return
}
func (c *ControllerBasics) UploadPart(ctx context.Context, req *basics.UploadPartReq) (res *basics.UploadPartRes, err error) {
	data, err := service.BasicsUpload().UploadPart(ctx, &req.UploadPartInp)
	if err != nil {
		return nil, err
	}
	res = new(basics.UploadPartRes)
	res.UploadPartModel = data
	return
}
