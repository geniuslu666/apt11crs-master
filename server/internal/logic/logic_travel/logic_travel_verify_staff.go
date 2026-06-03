package logic_travel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sTravelVerifyStaff struct{}

func NewTravelVerifyStaff() *sTravelVerifyStaff {
	return &sTravelVerifyStaff{}
}

func init() {
	service.RegisterTravelVerifyStaff(NewTravelVerifyStaff())
}

// Model 核销人员ORM模型
func (s *sTravelVerifyStaff) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TravelVerifyStaff.Ctx(ctx), option...)
}

// List 获取核销人员列表
func (s *sTravelVerifyStaff) List(ctx context.Context, in *input_travel.TravelVerifyStaffListInp) (list []*input_travel.TravelVerifyStaffListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	if !g.IsEmpty(in.Name) {
		mod = mod.WhereLike(dao.TravelVerifyStaff.Columns().Name, "%"+in.Name+"%")
	}
	if in.Status > 0 {
		mod = mod.Where(dao.TravelVerifyStaff.Columns().Status, in.Status)
	}

	mod = mod.OrderDesc(dao.TravelVerifyStaff.Columns().Id)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取核销人员列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取核销人员列表失败，请稍后重试！")
	}
	return
}

// View 获取核销人员详情
func (s *sTravelVerifyStaff) View(ctx context.Context, in *input_travel.TravelVerifyStaffViewInp) (res *input_travel.TravelVerifyStaffViewModel, err error) {
	if err = s.Model(ctx).Where(dao.TravelVerifyStaff.Columns().Id, in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取核销人员详情失败，请稍后重试！")
		return
	}
	if res == nil {
		err = gerror.New("核销人员不存在")
		return
	}

	var scopeList []*entity.TravelVerifyStaffScope
	if err = dao.TravelVerifyStaffScope.Ctx(ctx).
		Where(dao.TravelVerifyStaffScope.Columns().StaffId, res.Id).
		OrderAsc(dao.TravelVerifyStaffScope.Columns().Id).
		Scan(&scopeList); err != nil {
		err = gerror.Wrap(err, "获取核销员权限范围失败")
		return
	}

	if len(scopeList) == 0 {
		//兼容历史数据：未配置范围时按全部可核销处理
		res.IsAllScope = true
		res.ScopeItems = []*input_travel.TravelVerifyStaffScopeItem{}
		return
	}

	res.IsAllScope = false
	scopeItems := make([]*input_travel.TravelVerifyStaffScopeItem, 0, len(scopeList))
	for _, item := range scopeList {
		if item.IsAll == 1 {
			res.IsAllScope = true
			res.ScopeItems = []*input_travel.TravelVerifyStaffScopeItem{}
			return
		}
		scopeItems = append(scopeItems, &input_travel.TravelVerifyStaffScopeItem{
			ProductId: item.ProductId,
			SkuId:     item.SkuId,
		})
	}
	res.ScopeItems = scopeItems
	return
}

// Edit 新增/编辑核销人员
func (s *sTravelVerifyStaff) Edit(ctx context.Context, in *input_travel.TravelVerifyStaffEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 检查账号唯一（新增和编辑都需要）
		var count int
		if in.Id > 0 {
			// 编辑时排除自己的账号
			count, err = dao.TravelVerifyStaff.Ctx(ctx).TX(tx).
				Where(dao.TravelVerifyStaff.Columns().Username, in.Username).
				WhereNot(dao.TravelVerifyStaff.Columns().Id, in.Id).
				Count()
		} else {
			// 新增时直接检查
			count, err = dao.TravelVerifyStaff.Ctx(ctx).TX(tx).
				Where(dao.TravelVerifyStaff.Columns().Username, in.Username).
				Count()
		}
		if err != nil {
			return gerror.Wrap(err, "检查账号失败")
		}
		if count > 0 {
			return gerror.New("登录账号已存在")
		}

		staffID := uint64(in.Id)

		// 修改
		if in.Id > 0 {
			// 获取原记录信息
			var oldStaff *entity.TravelVerifyStaff
			if err = dao.TravelVerifyStaff.Ctx(ctx).TX(tx).WherePri(in.Id).Scan(&oldStaff); err != nil {
				return gerror.Wrap(err, "获取核销人员信息失败")
			}
			if oldStaff == nil {
				return gerror.New("核销人员不存在")
			}

			// 构建更新数据
			editData := in
			if !g.IsEmpty(in.Password) {
				// 修改密码，生成新的盐值
				editData.Salt = grand.S(6)
				editData.PasswordHash = gmd5.MustEncryptString(in.Password + editData.Salt)
			} else {
				// 不修改密码，保持原盐值和密码哈希
				editData.Salt = oldStaff.Salt
				editData.PasswordHash = oldStaff.PasswordHash
			}
			if _, err = dao.TravelVerifyStaff.Ctx(ctx).TX(tx).
				Fields(input_travel.TravelVerifyStaffUpdateFields{}).
				WherePri(in.Id).Data(editData).OmitEmptyData().Update(); err != nil {
				return gerror.Wrap(err, "编辑核销人员失败，请稍后重试！")
			}

			hasScopePayload := in.IsAllScope || len(in.ScopeItems) > 0
			if !hasScopePayload {
				return nil
			}
			if err = s.replaceStaffScopes(ctx, tx, staffID, in); err != nil {
				return err
			}
			return nil
		}

		// 新增
		if g.IsEmpty(in.Password) {
			return gerror.New("新增核销人员时密码不能为空")
		}
		//生成盐值和加密密码
		insertData := in
		insertData.Salt = grand.S(6)
		insertData.PasswordHash = gmd5.MustEncryptString(in.Password + insertData.Salt)
		insertResult, insertErr := dao.TravelVerifyStaff.Ctx(ctx).TX(tx).
			Fields(input_travel.TravelVerifyStaffInsertFields{}).
			Data(insertData).OmitEmptyData().Insert()
		if insertErr != nil {
			return gerror.Wrap(insertErr, "新增核销人员失败，请稍后重试！")
		}
		lastInsertId, lastInsertErr := insertResult.LastInsertId()
		if lastInsertErr != nil {
			return gerror.Wrap(lastInsertErr, "获取核销人员ID失败")
		}
		staffID = uint64(lastInsertId)

		if err = s.replaceStaffScopes(ctx, tx, staffID, in); err != nil {
			return err
		}
		return nil
	})
}

// Delete 删除核销人员
func (s *sTravelVerifyStaff) Delete(ctx context.Context, in *input_travel.TravelVerifyStaffDeleteInp) (err error) {
	_, err = s.Model(ctx).WhereIn(dao.TravelVerifyStaff.Columns().Id, in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, "删除核销人员失败，请稍后重试！")
	}
	return
}

// Status 更新核销人员状态
func (s *sTravelVerifyStaff) Status(ctx context.Context, in *input_travel.TravelVerifyStaffStatusInp) (err error) {
	_, err = s.Model(ctx).Where(dao.TravelVerifyStaff.Columns().Id, in.Id).
		Data(g.Map{dao.TravelVerifyStaff.Columns().Status: in.Status}).Update()
	if err != nil {
		err = gerror.Wrap(err, "更新核销人员状态失败，请稍后重试！")
	}
	return
}

// ScopeOptions 获取核销权限范围选项（产品与SKU）
func (s *sTravelVerifyStaff) ScopeOptions(ctx context.Context) (list []*input_travel.TravelVerifyStaffScopeOptionModel, err error) {
	var products []*entity.TravelProduct
	if err = dao.TravelProduct.Ctx(ctx).
		Fields(dao.TravelProduct.Columns().Id, dao.TravelProduct.Columns().Title).
		Where(dao.TravelProduct.Columns().Status, 1).
		OrderDesc(dao.TravelProduct.Columns().Sort).
		OrderDesc(dao.TravelProduct.Columns().Id).
		Hook(hook2.PmsFindLanguageValueHook).
		Scan(&products); err != nil {
		err = gerror.Wrap(err, "获取产品选项失败")
		return
	}
	if len(products) == 0 {
		list = []*input_travel.TravelVerifyStaffScopeOptionModel{}
		return
	}

	productIds := make([]uint64, 0, len(products))
	for _, product := range products {
		productIds = append(productIds, product.Id)
	}

	var skuRows []*entity.TravelProductSku
	if err = dao.TravelProductSku.Ctx(ctx).
		Fields(dao.TravelProductSku.Columns().Id, dao.TravelProductSku.Columns().ProductId, dao.TravelProductSku.Columns().Name).
		WhereIn(dao.TravelProductSku.Columns().ProductId, productIds).
		Where(dao.TravelProductSku.Columns().Status, 1).
		OrderDesc(dao.TravelProductSku.Columns().Sort).
		OrderDesc(dao.TravelProductSku.Columns().Id).
		Hook(hook2.PmsFindLanguageValueHook).
		Scan(&skuRows); err != nil {
		err = gerror.Wrap(err, "获取SKU选项失败")
		return
	}

	skuMap := make(map[uint64][]*input_travel.TravelVerifyStaffScopeSkuOptionModel)
	for _, sku := range skuRows {
		pid := uint64(sku.ProductId)
		skuMap[pid] = append(skuMap[pid], &input_travel.TravelVerifyStaffScopeSkuOptionModel{
			Id:        sku.Id,
			ProductId: pid,
			Name:      sku.Name,
		})
	}

	list = make([]*input_travel.TravelVerifyStaffScopeOptionModel, 0, len(products))
	for _, product := range products {
		list = append(list, &input_travel.TravelVerifyStaffScopeOptionModel{
			Id:      product.Id,
			Title:   product.Title,
			SkuList: skuMap[product.Id],
		})
	}
	return
}

func (s *sTravelVerifyStaff) replaceStaffScopes(ctx context.Context, tx gdb.TX, staffID uint64, in *input_travel.TravelVerifyStaffEditInp) (err error) {
	if staffID == 0 {
		return gerror.New("核销人员ID无效")
	}

	if _, err = dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).Unscoped().
		Where(dao.TravelVerifyStaffScope.Columns().StaffId, staffID).
		Delete(); err != nil {
		return gerror.Wrap(err, "清理核销员权限范围失败")
	}

	if in.IsAllScope {
		_, err = dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).Data(g.Map{
			dao.TravelVerifyStaffScope.Columns().StaffId:   staffID,
			dao.TravelVerifyStaffScope.Columns().ProductId: 0,
			dao.TravelVerifyStaffScope.Columns().SkuId:     0,
			dao.TravelVerifyStaffScope.Columns().IsAll:     1,
		}).Insert()
		if err != nil {
			return gerror.Wrap(err, "保存全部核销权限失败")
		}
		return nil
	}

	scopeItems, err := s.normalizeAndValidateScopeItems(ctx, tx, in.ScopeItems)
	if err != nil {
		return err
	}
	if len(scopeItems) == 0 {
		return gerror.New("请至少配置一条核销权限范围")
	}

	insertData := make([]g.Map, 0, len(scopeItems))
	for _, item := range scopeItems {
		insertData = append(insertData, g.Map{
			dao.TravelVerifyStaffScope.Columns().StaffId:   staffID,
			dao.TravelVerifyStaffScope.Columns().ProductId: item.ProductId,
			dao.TravelVerifyStaffScope.Columns().SkuId:     item.SkuId,
			dao.TravelVerifyStaffScope.Columns().IsAll:     0,
		})
	}
	if _, err = dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).Data(insertData).Insert(); err != nil {
		return gerror.Wrap(err, "保存核销权限范围失败")
	}
	return nil
}

func (s *sTravelVerifyStaff) normalizeAndValidateScopeItems(ctx context.Context, tx gdb.TX, items []*input_travel.TravelVerifyStaffScopeItem) (res []*input_travel.TravelVerifyStaffScopeItem, err error) {
	if len(items) == 0 {
		return []*input_travel.TravelVerifyStaffScopeItem{}, nil
	}

	res = make([]*input_travel.TravelVerifyStaffScopeItem, 0, len(items))
	uniq := make(map[string]struct{})

	for _, item := range items {
		if item == nil {
			continue
		}
		if item.ProductId == 0 {
			return nil, gerror.New("授权范围中的产品ID不能为空")
		}

		if item.SkuId > 0 {
			skuCount, skuErr := dao.TravelProductSku.Ctx(ctx).TX(tx).
				Where(dao.TravelProductSku.Columns().Id, item.SkuId).
				Where(dao.TravelProductSku.Columns().ProductId, item.ProductId).
				Count()
			if skuErr != nil {
				return nil, gerror.Wrap(skuErr, "校验SKU失败")
			}
			if skuCount == 0 {
				return nil, gerror.Newf("SKU(%d)不属于产品(%d)", item.SkuId, item.ProductId)
			}
		} else {
			productCount, productErr := dao.TravelProduct.Ctx(ctx).TX(tx).
				Where(dao.TravelProduct.Columns().Id, item.ProductId).
				Count()
			if productErr != nil {
				return nil, gerror.Wrap(productErr, "校验产品失败")
			}
			if productCount == 0 {
				return nil, gerror.Newf("产品(%d)不存在", item.ProductId)
			}
		}

		key := fmt.Sprintf("%d_%d", item.ProductId, item.SkuId)
		if _, ok := uniq[key]; ok {
			continue
		}
		uniq[key] = struct{}{}
		res = append(res, &input_travel.TravelVerifyStaffScopeItem{
			ProductId: item.ProductId,
			SkuId:     item.SkuId,
		})
	}
	return res, nil
}
