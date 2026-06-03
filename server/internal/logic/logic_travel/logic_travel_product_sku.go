package logic_travel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type sTravelProductSku struct{}

func NewTravelProductSku() *sTravelProductSku {
	return &sTravelProductSku{}
}

func init() {
	service.RegisterTravelProductSku(NewTravelProductSku())
}

func (s *sTravelProductSku) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TravelProductSku.Ctx(ctx), option...)
}

// List 获取车型列表
func (s *sTravelProductSku) List(ctx context.Context, in *input_travel.TravelProductSkuListInp) (list []*input_travel.TravelProductSkuListModel, totalCount int, err error) {
	mod := s.Model(ctx).Fields(input_travel.TravelProductSkuListModel{})

	if in.ProductId > 0 {
		mod = mod.Where(dao.TravelProductSku.Columns().ProductId, in.ProductId)
	}

	if !g.IsEmpty(in.Name) {
		uuIds, uErr := service.BasicsLanguage().GetUuids(ctx, in.Name, "name")
		if uErr == nil {
			skuIds, _ := service.TravelProduct().GetSkuIds(ctx, uuIds)
			mod = mod.WhereIn(dao.TravelProductSku.Columns().Id, skuIds)
		} else {
			mod = mod.WhereLike(dao.TravelProductSku.Columns().Name, "%"+in.Name+"%")
		}
	}

	mod = mod.OrderDesc(dao.TravelProductSku.Columns().Sort).OrderDesc(dao.TravelProductSku.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)
	mod = mod.Page(in.Page, in.PerPage)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取车型列表失败，请稍后重试！")
	}
	return
}

// View 获取车型详情
func (s *sTravelProductSku) View(ctx context.Context, in *input_travel.TravelProductSkuViewInp) (res *input_travel.TravelProductSkuViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取车型信息失败，请稍后重试！")
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取车型信息失败，请稍后重试！")
		}
	}
	return
}

// Edit 新增/编辑车型
func (s *sTravelProductSku) Edit(ctx context.Context, in *input_travel.TravelProductSkuEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		nameUuid := guid.S([]byte("name"))

		if in.Id > 0 {
			// 编辑：获取旧记录的 name UUID
			obj, qErr := dao.TravelProductSku.Ctx(ctx).WherePri(in.Id).One()
			if qErr != nil {
				return gerror.Wrap(qErr, "车型不存在！")
			}
			if !g.IsEmpty(obj[gstr.ToLower("Name")]) {
				nameUuid = obj["name"].String()
			}

			nameDao := &input_language.LoadLanguage{
				Uuid: nameUuid,
				Tag:  dao.TravelProductSku.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("Name"),
			}
			if err = service.BasicsLanguage().Sync(ctx, in.NameLanguage, nameDao); err != nil {
				return
			}

			if _, err = s.Model(ctx).Fields(input_travel.TravelProductSkuUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改车型失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.Name = nameUuid
		lastId, iErr := s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_travel.TravelProductSkuInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId()
		if iErr != nil {
			return gerror.Wrap(iErr, "新增车型失败，请稍后重试！")
		}
		if lastId < 1 {
			return gerror.New("新增失败，请稍后重试！")
		}

		nameDao := &input_language.LoadLanguage{
			Uuid: nameUuid,
			Tag:  dao.TravelProductSku.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("Name"),
		}
		if err = service.BasicsLanguage().Sync(ctx, in.NameLanguage, nameDao); err != nil {
			return
		}
		return
	})
}

// Delete 删除车型
func (s *sTravelProductSku) Delete(ctx context.Context, in *input_travel.TravelProductSkuDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除车型失败，请稍后重试！")
	}
	return
}

// Status 更新车型状态
func (s *sTravelProductSku) Status(ctx context.Context, in *input_travel.TravelProductSkuStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TravelProductSku.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新车型状态失败，请稍后重试！")
	}
	return
}
