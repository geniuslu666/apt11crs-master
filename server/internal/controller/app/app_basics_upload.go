package app

import (
	"APT/internal/library/storager"
	"APT/internal/service"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/basics"
)

func (c *ControllerBasics) UploadImage(ctx context.Context, req *basics.UploadImageReq) (res *basics.UploadImageRes, err error) {
	r := g.RequestFromCtx(ctx)
	uploadType := r.Header.Get("uploadType")
	if uploadType != "default" && !validate.InSlice(storager.KindSlice, uploadType) {
		// 上传类型是无效的
		err = gerror.New(gi18n.T(ctx, "the_upload_type_is_invalid"))
		return
	}

	file := r.GetUploadFile("file")
	if file == nil {
		// 没有找到上传的文件
		err = gerror.New(gi18n.T(ctx, "no_upload_file_found"))
		return
	}
	res = new(basics.UploadImageRes)
	res.AttachmentListModel, err = service.BasicsUpload().UploadFile(ctx, uploadType, file, req.Kind)
	return
}
