// Package storager
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package storager

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"google.golang.org/api/option"
)

// GoogleDrive 七牛云对象存储驱动
type GoogleDrive struct {
}

func (d *GoogleDrive) UploadFile(ctx context.Context, filePath string) (fullPath string, err error) {
	return
}

// Upload 上传到Google云对象存储
func (d *GoogleDrive) Upload(ctx context.Context, file *ghttp.UploadFile) (fullPath string, err error) {
	if g.IsEmpty(config.UploadGoogleCredentials) || g.IsEmpty(config.UploadGoogleBucketName) || g.IsEmpty(config.UploadGoogleEndpoint) || g.IsEmpty(config.UploadGooglePath) {
		err = gerror.New("Google云存储驱动必须配置服务凭证和存储桶名称!")
		return
	}
	googlePath := GenFullPath(config.UploadGooglePath, gfile.Ext(file.Filename))
	// 初始化客户端，使用服务账号密钥文件
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.UploadGoogleCredentials)))
	if err != nil {
		g.Log().Errorf(ctx, "无法创建存储客户端: %v", err)
		return
	}
	defer client.Close()

	// 打开本地文件
	f, err := file.Open()
	if err != nil {
		g.Log().Errorf(ctx, "无法打开文件: %v", err)
		return
	}
	defer f.Close()

	// 获取存储桶句柄
	bucket := client.Bucket(config.UploadGoogleBucketName)

	// 创建目标对象的句柄
	obj := bucket.Object(googlePath)

	// 创建写入器
	wc := obj.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		g.Log().Errorf(ctx, "无法写入到 GCS: %v", err)
		return
	}

	// 关闭写入器
	if err = wc.Close(); err != nil {
		g.Log().Errorf(ctx, "无法关闭写入器: %v", err)
		return
	}
	fullPath = googlePath
	return
}

// CreateMultipart 创建分片事件
func (d *GoogleDrive) CreateMultipart(ctx context.Context, in *CheckMultipartParams) (res *MultipartProgress, err error) {
	err = gerror.New("当前驱动暂不支持分片上传！")
	return
}

// UploadPart 上传分片
func (d *GoogleDrive) UploadPart(ctx context.Context, in *UploadPartParams) (res *UploadPartModel, err error) {
	err = gerror.New("当前驱动暂不支持分片上传！")
	return
}

func (d *GoogleDrive) UploadFileName(ctx context.Context, filePath string, fileName string) (fullPath string, err error) {
	return
}
