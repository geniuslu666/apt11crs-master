package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBannerUpdateFields 修改banner 横幅字段过滤
type PmsBannerUpdateFields struct {
	Language     string `json:"language"     dc:"语言"`
	BannerName   string `json:"bannerName"   dc:"名称"`
	BannerImage  string `json:"bannerImage"  dc:"轮播图"`
	Model        string `json:"model"        dc:"模块"`
	Chain        string `json:"chain"        dc:"内外联"`
	Path         string `json:"path"         dc:"链接内容"`
	BannerStatus int    `json:"bannerStatus" dc:"状态"`
}

// PmsBannerInsertFields 新增banner 横幅字段过滤
type PmsBannerInsertFields struct {
	Language     string `json:"language"     dc:"语言"`
	BannerName   string `json:"bannerName"   dc:"名称"`
	BannerImage  string `json:"bannerImage"  dc:"轮播图"`
	Model        string `json:"model"        dc:"模块"`
	Chain        string `json:"chain"        dc:"内外联"`
	Path         string `json:"path"         dc:"链接内容"`
	BannerStatus int    `json:"bannerStatus" dc:"状态"`
}

// PmsBannerEditInp 修改/新增banner 横幅
type PmsBannerEditInp struct {
	entity.PmsBanner
}

func (in *PmsBannerEditInp) Filter(ctx context.Context) (err error) {
	// 验证语言
	if err := g.Validator().Rules("required").Data(in.Language).Messages("语言不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:en,zh,ja,ko,zh_CN").Data(in.Language).Messages("语言值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证名称
	if err := g.Validator().Rules("required").Data(in.BannerName).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证轮播图
	if err := g.Validator().Rules("required").Data(in.BannerImage).Messages("轮播图不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证模块
	if err := g.Validator().Rules("required").Data(in.Model).Messages("模块不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:BANNER,EVENTS").Data(in.Model).Messages("模块值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证内外联
	if err := g.Validator().Rules("required").Data(in.Chain).Messages("内外联不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:OUT,IN").Data(in.Chain).Messages("内外联值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.BannerStatus).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.BannerStatus).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsBannerEditModel struct{}

// PmsBannerDeleteInp 删除banner 横幅
type PmsBannerDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsBannerDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsBannerDeleteModel struct{}

// PmsBannerViewInp 获取指定banner 横幅信息
type PmsBannerViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsBannerViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsBannerViewModel struct {
	entity.PmsBanner
}

// PmsBannerListInp 获取banner 横幅列表
type PmsBannerListInp struct {
	input_form.PageReq
	Id           int    `json:"id"           dc:"id"`
	Language     string `json:"language"     dc:"语言"`
	BannerName   string `json:"bannerName"   dc:"名称"`
	Model        string `json:"model"        dc:"模块"`
	Chain        string `json:"chain"        dc:"内外联"`
	BannerStatus int    `json:"bannerStatus" dc:"状态"`
}

func (in *PmsBannerListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsBannerListModel struct {
	Id           int         `json:"id"           dc:"id"`
	Language     string      `json:"language"     dc:"语言"`
	BannerName   string      `json:"bannerName"   dc:"名称"`
	BannerImage  string      `json:"bannerImage"  dc:"轮播图"`
	Model        string      `json:"model"        dc:"模块"`
	Chain        string      `json:"chain"        dc:"内外联"`
	Path         string      `json:"path"         dc:"链接内容"`
	BannerStatus int         `json:"bannerStatus" dc:"状态"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}

type PmsBannerAppListModel struct {
	Id           int         `json:"id"           dc:"id"`
	Language     string      `json:"language"     dc:"语言"`
	BannerName   string      `json:"bannerName"   dc:"名称"`
	BannerImage  string      `json:"bannerImage"  dc:"轮播图"`
	Model        string      `json:"model"        dc:"模块"`
	Chain        string      `json:"chain"        dc:"内外联"`
	Path         string      `json:"path"         dc:"链接内容"`
	WxPath       string      `json:"wxPath"       dc:"微信链接内容"`
	AppPath      string      `json:"appPath"      dc:"app链接内容"`
	BannerStatus int         `json:"bannerStatus" dc:"状态"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}

// PmsBannerExportModel 导出banner 横幅
type PmsBannerExportModel struct {
	Id           int         `json:"id"           dc:"id"`
	Language     string      `json:"language"     dc:"语言"`
	BannerName   string      `json:"bannerName"   dc:"名称"`
	BannerImage  string      `json:"bannerImage"  dc:"轮播图"`
	Model        string      `json:"model"        dc:"模块"`
	Chain        string      `json:"chain"        dc:"内外联"`
	Path         string      `json:"path"         dc:"链接内容"`
	BannerStatus int         `json:"bannerStatus" dc:"状态"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}
