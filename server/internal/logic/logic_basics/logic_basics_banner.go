package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsBanner struct{}

func NewBasicsBanner() *sBasicsBanner {
	return &sBasicsBanner{}
}

func init() {
	service.RegisterBasicsBanner(NewBasicsBanner())
}

func (s *sBasicsBanner) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PmsBanner.Ctx(ctx), option...)
}

func (s *sBasicsBanner) List(ctx context.Context, in *input_basics.PmsBannerListInp) (list []*input_basics.PmsBannerListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.Fields(input_basics.PmsBannerListModel{})

	if in.Id > 0 {
		mod = mod.Where(dao.PmsBanner.Columns().Id, in.Id)
	}

	if in.Language != "" {
		mod = mod.Where(dao.PmsBanner.Columns().Language, in.Language)
	}

	if in.BannerName != "" {
		mod = mod.WhereLike(dao.PmsBanner.Columns().BannerName, in.BannerName)
	}

	if in.Model != "" {
		mod = mod.Where(dao.PmsBanner.Columns().Model, in.Model)
	}

	if in.Chain != "" {
		mod = mod.Where(dao.PmsBanner.Columns().Chain, in.Chain)
	}

	if in.BannerStatus > 0 {
		mod = mod.Where(dao.PmsBanner.Columns().BannerStatus, in.BannerStatus)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.PmsBanner.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取banner 横幅列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsBanner) Export(ctx context.Context, in *input_basics.PmsBannerListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_basics.PmsBannerExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出BANNER横幅-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("BANNER横幅")
		exports   []input_basics.PmsBannerExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sBasicsBanner) Edit(ctx context.Context, in *input_basics.PmsBannerEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(input_basics.PmsBannerUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改banner 横幅失败，请稍后重试！")
			}
			return
		}

		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_basics.PmsBannerInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增banner 横幅失败，请稍后重试！")
		}
		return
	})
}

func (s *sBasicsBanner) Delete(ctx context.Context, in *input_basics.PmsBannerDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除banner 横幅失败，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsBanner) View(ctx context.Context, in *input_basics.PmsBannerViewInp) (res *input_basics.PmsBannerViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取banner 横幅信息，请稍后重试！")
		return
	}
	return
}

func (s *sBasicsBanner) BannerAppList(ctx context.Context, in *input_basics.PmsBannerListInp) (list []*input_basics.PmsBannerAppListModel, totalCount int, err error) {
	var bannerList []*input_basics.PmsBannerListModel
	if bannerList, _, err = s.List(ctx, in); err != nil {
		return
	}

	// 遍历轮播图获取跳转地址
	for _, v := range bannerList {
		var AppPath string
		var WxPath string
		if v.Chain == "IN" && v.Path != "" {
			var PathJsonData *struct {
				Weapp string `json:"weapp"`
				App   string `json:"app"`
			}
			if err = json.Unmarshal([]byte(v.Path), &PathJsonData); err != nil {
				err = gerror.New("解析路径失败")
				return
			}
			AppPath = PathJsonData.App
			WxPath = PathJsonData.Weapp
		}
		list = append(list, &input_basics.PmsBannerAppListModel{
			BannerImage:  v.BannerImage,
			BannerName:   v.BannerName,
			BannerStatus: v.BannerStatus,
			Chain:        v.Chain,
			Language:     v.Language,
			Model:        v.Model,
			Path:         v.Path,
			AppPath:      AppPath,
			WxPath:       WxPath,
		})
	}
	return
}
