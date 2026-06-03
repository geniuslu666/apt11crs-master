package logic_th

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_language"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

type sThMchStore struct{}

func NewThMchStore() *sThMchStore {
	return &sThMchStore{}
}

func init() {
	service.RegisterThMchStore(NewThMchStore())
}

func (s *sThMchStore) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThMchStore.Ctx(ctx), option...)
}

func (s *sThMchStore) List(ctx context.Context, in *input_th.ThMchStoreListInp) (list []*input_th.ThMchStoreListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	mod = mod.FieldsPrefix(dao.ThMchStore.Table(), input_th.ThMchStoreListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_th.ThMchStoreListModel{}, &dao.ThMch, "thMch"))

	mod = mod.LeftJoinOnFields(dao.ThMch.Table(), dao.ThMchStore.Columns().MchId, "=", dao.ThMch.Columns().Id)

	if !g.IsEmpty(in.StoreName) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.StoreName, "store_name")
		if err == nil {
			storeIds, _ := service.ThMchStore().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.ThMchStore.Columns().Id, storeIds)
		}
	}

	if !g.IsEmpty(in.MchId) {
		mod = mod.Where(dao.ThMchStore.Columns().MchId, in.MchId)
	}

	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {

		mod = mod.Page(in.Page, in.PerPage)
	}

	mod = mod.OrderDesc(dao.ThMchStore.Table() + "." + dao.ThMchStore.Columns().CreateAt)
	if in.Pagination {

		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户门店列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取商户门店列表失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThMchStore) Edit(ctx context.Context, in *input_th.ThMchStoreEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object             gdb.Record
			StoreTerminalInfo  []*entity.ThStoreTerminal
			ThMchStoreNameDao  *input_language.LoadLanguage
			NameLanguageStruct input_language.LanguageModel
		)
		NameUuid := guid.S([]byte("name"))

		if in.Id > 0 {
			if Object, err = dao.ThMchStore.Ctx(ctx).Where(dao.ThMchStore.Columns().Id, in.Id).One(); err != nil {
				return
			}

			// 更新终端表的storeId
			if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().StoreId, in.Id).Update(g.MapStrAny{
				dao.SysTerminal.Columns().StoreId: 0,
			}); err != nil {
				err = gerror.Wrap(err, "清理终端旧数据失败，请稍后重试！")
				return
			}

			if !g.IsEmpty(in.TerminalList) {
				for _, TerminalInfo := range in.TerminalList {
					if !g.IsEmpty(TerminalInfo) {
						TerminalId := TerminalInfo.Id
						StoreTerminalInfo = append(StoreTerminalInfo, &entity.ThStoreTerminal{
							StoreId:    in.Id,
							TerminalId: int64(TerminalId),
							PrintTimes: uint(TerminalInfo.PrintTimes),
						})

						if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().Id, TerminalId).Update(g.MapStrAny{
							dao.SysTerminal.Columns().StoreId: in.Id,
							//dao.SysTerminal.Columns().PrintTimes: TerminalInfo.PrintTimes,
						}); err != nil {
							err = gerror.Wrap(err, "更新终端信息失败，请稍后重试！")
							return
						}
					}
				}
			}
			if !g.IsEmpty(StoreTerminalInfo) {
				if _, err = dao.ThStoreTerminal.Ctx(ctx).Where(dao.ThStoreTerminal.Columns().StoreId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理门店终端旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.ThStoreTerminal.Ctx(ctx).OmitEmptyData().Insert(StoreTerminalInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.ThStoreTerminal.Ctx(ctx).Where(dao.ThStoreTerminal.Columns().StoreId, in.Id).Delete(); err != nil {
					err = gerror.Wrap(err, "清理门店终端旧数据失败，请稍后重试！")
					return
				}
			}

			if !g.IsEmpty(Object[gstr.ToLower("StoreName")]) {
				NameUuid = Object["store_name"].String()
			}
			ThMchStoreNameDao = &input_language.LoadLanguage{
				Uuid: NameUuid,
				Tag:  dao.ThMchStore.Table(),
				Type: "table",
				Key:  gstr.CaseSnakeFirstUpper("StoreName"),
			}
			NameLanguageStruct = in.NameLanguage
			if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, ThMchStoreNameDao); err != nil {
				return
			}

			in.StoreName = NameUuid

			if in.Password != "" {
				if g.IsEmpty(Object[gstr.ToLower("Salt")]) {
					in.Salt = grand.S(6)
				} else {
					in.Salt = Object[gstr.ToLower("Salt")].String()
				}
				// 修改密码，需要获取到密码盐
				in.PasswordHash = gmd5.MustEncryptString(in.Password + in.Salt)
			} else {
				in.Salt = Object[gstr.ToLower("Salt")].String()
				in.PasswordHash = Object["password_hash"].String()
			}

			if _, err = s.Model(ctx).
				Fields(input_th.ThMchStoreUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改商户门店失败，请稍后重试！")
			}

			// 更新商户门店数量
			if gvar.New(Object["status"]).Int() == 1 {
				if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, Object["mch_id"]).Update(g.MapStrAny{
					dao.ThMch.Columns().StoreOnNum: gdb.Raw("store_on_num-1"),
				}); err != nil {
					err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
					return
				}
				if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, in.MchId).Update(g.MapStrAny{
					dao.ThMch.Columns().StoreOnNum: gdb.Raw("store_on_num+1"),
				}); err != nil {
					err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
					return
				}
			} else {
				if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, Object["mch_id"]).Update(g.MapStrAny{
					dao.ThMch.Columns().StoreOnNum: gdb.Raw("store_off_num-1"),
				}); err != nil {
					err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
					return
				}
				if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, in.MchId).Update(g.MapStrAny{
					dao.ThMch.Columns().StoreOnNum: gdb.Raw("store_off_num+1"),
				}); err != nil {
					err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
					return
				}
			}
			return
		}

		var (
			lastInsertId int64
		)
		in.StoreName = NameUuid

		in.Salt = grand.S(6)
		in.PasswordHash = gmd5.MustEncryptString(in.Password + in.Salt)

		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_th.ThMchStoreInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增套餐管理失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增商户门店失败，请稍后重试！")
			return
		}

		// 更新终端表的storeId和打印联数
		/*if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().StoreId, in.Id).Update(g.MapStrAny{
			dao.SysTerminal.Columns().StoreId:    0,
			dao.SysTerminal.Columns().PrintTimes: 1,
		}); err != nil {
			err = gerror.Wrap(err, "清理终端旧数据失败，请稍后重试！")
			return
		}*/

		if !g.IsEmpty(in.TerminalList) {
			for _, TerminalInfo := range in.TerminalList {
				if !g.IsEmpty(TerminalInfo) {
					TerminalId := TerminalInfo.Id
					StoreTerminalInfo = append(StoreTerminalInfo, &entity.ThStoreTerminal{
						StoreId:    lastInsertId,
						TerminalId: int64(TerminalId),
						PrintTimes: uint(TerminalInfo.PrintTimes),
					})

					if _, err = dao.SysTerminal.Ctx(ctx).Where(dao.SysTerminal.Columns().Id, TerminalId).Update(g.MapStrAny{
						dao.SysTerminal.Columns().StoreId: lastInsertId,
						//dao.SysTerminal.Columns().PrintTimes: TerminalInfo.PrintTimes,
					}); err != nil {
						err = gerror.Wrap(err, "更新终端信息失败，请稍后重试！")
						return
					}
				}
			}
		}

		if !g.IsEmpty(StoreTerminalInfo) {
			if _, err = dao.ThStoreTerminal.Ctx(ctx).Where(dao.ThStoreTerminal.Columns().StoreId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理门店终端旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.ThStoreTerminal.Ctx(ctx).OmitEmptyData().Insert(StoreTerminalInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.ThStoreTerminal.Ctx(ctx).Where(dao.ThStoreTerminal.Columns().StoreId, lastInsertId).Delete(); err != nil {
				err = gerror.Wrap(err, "清理门店终端旧数据失败，请稍后重试！")
				return
			}
		}

		ThMchStoreNameDao = &input_language.LoadLanguage{
			Uuid: NameUuid,
			Tag:  dao.ThMchStore.Table(),
			Type: "table",
			Key:  gstr.CaseSnakeFirstUpper("StoreName"),
		}
		NameLanguageStruct = in.NameLanguage
		if err = service.BasicsLanguage().Sync(ctx, NameLanguageStruct, ThMchStoreNameDao); err != nil {
			return
		}

		// 更新商户门店数量
		if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, in.MchId).Update(g.MapStrAny{
			dao.ThMch.Columns().StoreOnNum: gdb.Raw("store_on_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
			return
		}
		return
	})
}

func (s *sThMchStore) Delete(ctx context.Context, in *input_th.ThMchStoreDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除商户门店失败，请稍后重试！")
		return
	}
	return
}

func (s *sThMchStore) View(ctx context.Context, in *input_th.ThMchStoreViewInp) (res *input_th.ThMchStoreViewModel, err error) {

	mod := s.Model(ctx).WithAll()

	mod = mod.FieldsPrefix(dao.ThMchStore.Table(), input_th.ThMchStoreViewModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, input_th.ThMchStoreViewModel{}, &dao.ThMch, "thMch"))

	mod = mod.LeftJoinOnFields(dao.ThMch.Table(), dao.ThMchStore.Columns().MchId, "=", dao.ThMch.Columns().Id)

	if !in.IsLanguage {
		if err = mod.WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取商户门店信息，请稍后重试！")
			return
		}
	} else {
		if err = mod.WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取商户门店信息，请稍后重试！")
			return
		}
	}
	return
}

func (s *sThMchStore) Switch(ctx context.Context, in *input_th.ThMchStoreSwitchInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			Object *entity.ThMchStore
		)

		if err = dao.ThMchStore.Ctx(ctx).WherePri(in.Id).Scan(&Object); err != nil {
			err = gerror.Wrap(err, "获取商户门店信息，请稍后重试！")
			return
		}

		if in.Status == 1 {
			// 启用
			// 更新商户门店数量
			if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, Object.MchId).Update(g.MapStrAny{
				dao.ThMch.Columns().StoreOffNum: gdb.Raw("store_off_num-1"),
				dao.ThMch.Columns().StoreOnNum:  gdb.Raw("store_on_num+1"),
			}); err != nil {
				err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
				return
			}
		} else {
			// 停用
			// 更新商户门店数量
			if _, err = dao.ThMch.Ctx(ctx).Where(dao.ThMch.Columns().Id, Object.MchId).Update(g.MapStrAny{
				dao.ThMch.Columns().StoreOffNum: gdb.Raw("store_off_num+1"),
				dao.ThMch.Columns().StoreOnNum:  gdb.Raw("store_on_num-1"),
			}); err != nil {
				err = gerror.Wrap(err, "更新商户信息失败，请稍后重试！")
				return
			}
		}

		if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
			dao.ThMchStore.Columns().Status: in.Status,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "更新状态失败，请稍后重试！")
			return
		}

		return
	})
}

func (s *sThMchStore) GetIds(ctx context.Context, name []string) (ids []int, err error) {
	columns, err := s.Model(ctx).
		Fields("id").
		WhereIn("store_name", name).Array()
	if err != nil {
		err = gerror.Wrap(err, "获取id失败！")
		return
	}

	ids = g.NewVar(columns).Ints()
	return
}

func (s *sThMchStore) StoreAll(ctx context.Context, in *input_th.ThMchStoreAllInp) (list []*input_th.ThMchStoreAllModel, err error) {
	mod := dao.ThMchStore.Ctx(ctx)
	mod = mod.WithAll().Hook(hook2.PmsFindLanguageValueHook)

	mod = mod.Fields(input_th.ThMchStoreAllModel{})

	// 搜索状态为启用的门店
	mod = mod.Where(dao.ThMchStore.Columns().Status, 1)

	if !g.IsEmpty(in.CouponId) {
		// 根据券ID找到关联的商户，然后查出对应的门店
		mchIds, _ := dao.ThCouponMch.Ctx(ctx).Fields(dao.ThCouponMch.Columns().MchId).Where(dao.ThCouponMch.Columns().CouponId, in.CouponId).Array()
		mod = mod.WhereIn(dao.ThMchStore.Columns().MchId, mchIds)
	}

	mod = mod.OrderDesc(dao.ThMchStore.Columns().Id)

	if err = mod.Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取门店列表失败，请稍后重试！")
		return
	}
	return
}
