// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_th"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IThCoupon interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThCouponListInp) (list []*input_th.ThCouponListModel, totalCount int, err error)
		All(ctx context.Context, in *input_th.ThCouponListInp) (list []*input_th.ThCouponAllListModel, err error)
		Edit(ctx context.Context, in *input_th.ThCouponEditInp) (err error)
		Delete(ctx context.Context, in *input_th.ThCouponDeleteInp) (err error)
		MaxSort(ctx context.Context, in *input_th.ThCouponMaxSortInp) (res *input_th.ThCouponMaxSortModel, err error)
		View(ctx context.Context, in *input_th.ThCouponViewInp) (res *input_th.ThCouponViewModel, err error)
		AppView(ctx context.Context, in *input_th.ThCouponAppViewInp) (res *input_th.ThCouponAppViewModel, err error)
		Status(ctx context.Context, in *input_th.ThCouponStatusInp) (err error)
		UseStatus(ctx context.Context, in *input_th.ThCouponUseStatusInp) (err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		SendMemberCoupon(ctx context.Context, in *input_th.ThSendMemberCouponInp, source int) (err error)
		SendMemberCouponGetId(ctx context.Context, in *input_th.ThSendMemberCouponInp, source int) (memberCouponId int64, couponNameLanguage []*input_hotel.LanguageType, err error)
		AppReceiveCoupon(ctx context.Context, in *input_th.ThCouponAppReceiveInp) (err error)
	}
	IThCouponCategory interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThCouponCategoryListInp) (list []*input_th.ThCouponCategoryListModel, totalCount int, err error)
		All(ctx context.Context, in *input_th.ThCouponCategoryAllInp) (list []*input_th.ThCouponCategoryAllModel, err error)
		Edit(ctx context.Context, in *input_th.ThCouponCategoryEditInp) (err error)
		Delete(ctx context.Context, in *input_th.ThCouponCategoryDeleteInp) (err error)
		View(ctx context.Context, in *input_th.ThCouponCategoryViewInp) (res *input_th.ThCouponCategoryViewModel, err error)
		Switch(ctx context.Context, in *input_th.ThCouponCategorySwitchInp) (err error)
	}
	IThMch interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThMchListInp) (list []*input_th.ThMchListModel, totalCount int, err error)
		All(ctx context.Context, in *input_th.ThMchAllInp) (list []*input_th.ThMchAllModel, err error)
		Edit(ctx context.Context, in *input_th.ThMchEditInp) (err error)
		Delete(ctx context.Context, in *input_th.ThMchDeleteInp) (err error)
		View(ctx context.Context, in *input_th.ThMchViewInp) (res *input_th.ThMchViewModel, err error)
		Switch(ctx context.Context, in *input_th.ThMchSwitchInp) (err error)
		AppView(ctx context.Context, in *input_th.ThMchAppViewInp) (res *input_th.ThMchAppViewModel, err error)
	}
	IThMchCategory interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThMchCategoryListInp) (list []*input_th.ThMchCategoryListModel, totalCount int, err error)
		All(ctx context.Context, in *input_th.ThMchCategoryAllInp) (list []*input_th.ThMchCategoryAllModel, err error)
		Edit(ctx context.Context, in *input_th.ThMchCategoryEditInp) (err error)
		Delete(ctx context.Context, in *input_th.ThMchCategoryDeleteInp) (err error)
		View(ctx context.Context, in *input_th.ThMchCategoryViewInp) (res *input_th.ThMchCategoryViewModel, err error)
		Switch(ctx context.Context, in *input_th.ThMchCategorySwitchInp) (err error)
	}
	IThMchStore interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThMchStoreListInp) (list []*input_th.ThMchStoreListModel, totalCount int, err error)
		Edit(ctx context.Context, in *input_th.ThMchStoreEditInp) (err error)
		Delete(ctx context.Context, in *input_th.ThMchStoreDeleteInp) (err error)
		View(ctx context.Context, in *input_th.ThMchStoreViewInp) (res *input_th.ThMchStoreViewModel, err error)
		Switch(ctx context.Context, in *input_th.ThMchStoreSwitchInp) (err error)
		GetIds(ctx context.Context, name []string) (ids []int, err error)
		StoreAll(ctx context.Context, in *input_th.ThMchStoreAllInp) (list []*input_th.ThMchStoreAllModel, err error)
	}
	IThMemberCoupon interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *input_th.ThMemberCouponListInp) (list []*input_th.ThMemberCouponListModel, totalCount int, err error)
		RefreshCode(ctx context.Context, in *input_th.ThMemberCouponRefreshCodeInp) (code string, verifyStatus int, err error)
		Verify(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error)
		AppList(ctx context.Context, in *input_th.ThMemberCouponAppListInp) (list []*input_th.ThMemberCouponAppListModel, totalCount int, err error)
		AppView(ctx context.Context, in *input_th.ThMemberCouponAppViewInp) (res *input_th.ThMemberCouponAppViewModel, err error)
		CouponEffect(ctx context.Context, CouponOn string) (err error)
		InvalidMemberCoupon(ctx context.Context, in *input_th.ThInvalidMemberCouponInp) (err error)
		Export(ctx context.Context, in *input_th.ThMemberCouponListInp) (err error)
		View(ctx context.Context, in *input_th.ThMemberCouponViewInp) (res *input_th.ThMemberCouponAdminViewModel, err error)
		Recycle(ctx context.Context, in *input_th.ThMemberCouponRecycleInp) (err error)
		// ManualVerify 会员手动核销
		ManualVerify(ctx context.Context, in *input_th.ThMemberCouponManualVerifyInp) (err error)
	}
)

var (
	localThCoupon         IThCoupon
	localThCouponCategory IThCouponCategory
	localThMch            IThMch
	localThMchCategory    IThMchCategory
	localThMchStore       IThMchStore
	localThMemberCoupon   IThMemberCoupon
)

func ThCoupon() IThCoupon {
	if localThCoupon == nil {
		panic("implement not found for interface IThCoupon, forgot register?")
	}
	return localThCoupon
}

func RegisterThCoupon(i IThCoupon) {
	localThCoupon = i
}

func ThCouponCategory() IThCouponCategory {
	if localThCouponCategory == nil {
		panic("implement not found for interface IThCouponCategory, forgot register?")
	}
	return localThCouponCategory
}

func RegisterThCouponCategory(i IThCouponCategory) {
	localThCouponCategory = i
}

func ThMch() IThMch {
	if localThMch == nil {
		panic("implement not found for interface IThMch, forgot register?")
	}
	return localThMch
}

func RegisterThMch(i IThMch) {
	localThMch = i
}

func ThMchCategory() IThMchCategory {
	if localThMchCategory == nil {
		panic("implement not found for interface IThMchCategory, forgot register?")
	}
	return localThMchCategory
}

func RegisterThMchCategory(i IThMchCategory) {
	localThMchCategory = i
}

func ThMchStore() IThMchStore {
	if localThMchStore == nil {
		panic("implement not found for interface IThMchStore, forgot register?")
	}
	return localThMchStore
}

func RegisterThMchStore(i IThMchStore) {
	localThMchStore = i
}

func ThMemberCoupon() IThMemberCoupon {
	if localThMemberCoupon == nil {
		panic("implement not found for interface IThMemberCoupon, forgot register?")
	}
	return localThMemberCoupon
}

func RegisterThMemberCoupon(i IThMemberCoupon) {
	localThMemberCoupon = i
}
