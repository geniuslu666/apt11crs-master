package app

import (
	"APT/internal/library/cache"
	"APT/internal/library/contexts"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_hotel"
	"APT/internal/service"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/basics"
)

func (c *ControllerBasics) CouponList(ctx context.Context, req *basics.CouponListReq) (res *basics.CouponListRes, err error) {
	var (
		MemberInfo = contexts.GetMemberUser(ctx)
	)
	res = new(basics.CouponListRes)
	if res.List, res.Count, err = service.BasicsCoupon().AppList(ctx, &input_basics.PmsCouponListInp{
		PageReq:  req.PageReq,
		State:    req.State,
		MemberId: MemberInfo.Id,
	}); err != nil {
		return
	}
	return
}

func (c *ControllerBasics) CouponOrderList(ctx context.Context, req *basics.CouponOrderListReq) (res *basics.CouponOrderListRes, err error) {
	var (
		MemberInfo        = contexts.GetMemberUser(ctx)
		PreOrderDetailVar *gvar.Var
		PreOrderDetail    *input_hotel.PreOrderDetailModel
		CabPreOrderDetail *input_cabinet.PreOrderDetailModel
		IsUse             []*input_basics.PmsCouponListModel
		IsNotUse          []*input_basics.PmsCouponListModel
		AtLeast           float64
	)
	// 读取预下单信息
	if PreOrderDetailVar, err = cache.Instance().Get(ctx, "PreOrder_"+req.PreOrderSn); err != nil || PreOrderDetailVar.IsEmpty() {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	if req.Scene == 1 {
		// 住宿
		PreOrderDetail = new(input_hotel.PreOrderDetailModel)
		if err = PreOrderDetailVar.Struct(&PreOrderDetail); err != nil {
			// 订单错误
			err = gerror.New(gi18n.T(ctx, "order_error"))
			return
		}
		AtLeast = PreOrderDetail.PayInfo.AllAmount
	} else if req.Scene == 5 {
		// 储物柜
		CabPreOrderDetail = new(input_cabinet.PreOrderDetailModel)
		if err = PreOrderDetailVar.Struct(&CabPreOrderDetail); err != nil {
			// 订单错误
			err = gerror.New(gi18n.T(ctx, "order_error"))
			return
		}
		AtLeast = CabPreOrderDetail.PayInfo.AllAmount
	}
	res = new(basics.CouponOrderListRes)
	if res.List, res.Count, err = service.BasicsCoupon().List(ctx, &input_basics.PmsCouponListInp{
		PageReq:  req.PageReq,
		State:    1,
		MemberId: MemberInfo.Id,
		Scene:    req.Scene,
		//PropertyId:  PreOrderDetail.Property.Id,
		WantUseTime: gtime.Now().Format("Y-m-d H:i:s"),
		AtLeast:     AtLeast,
	}); err != nil {
		return
	}
	if len(res.List) > 0 {

		// 全部物业语言翻译
		var Language = contexts.GetLanguage(ctx)
		var LanguageMap = map[string]string{
			"zh":    "全部物业可用",
			"zh_CN": "全部物業可用",
			"en":    "All properties",
			"ja":    "全物件",
			"ko":    "모든 숙소에서 사용 가능",
		}

		for _, coupon := range res.List {
			coupon.PropertyNames = LanguageMap[Language]
			if req.Scene == 1 && !g.IsEmpty(coupon.PropertyIds) {
				coupon.PropertyNames = ""
				var propertyArr []*input_hotel.PmsPropertyAllModel
				propertyArr, err = service.HotelService().PropertyAll(ctx, &input_hotel.PmsPropertyAllInp{
					Ids: coupon.PropertyIds,
				})
				if err != nil {
					// 获取物业失败
					err = gerror.Wrap(err, gi18n.T(ctx, "get_property_failed"))
					return
				}
				//if err = gvar.New(propertyArr).Struct(&coupon.Property); err != nil {
				//	return
				//}
				// 获取物业名称以逗号 分隔
				var names []string
				for _, property := range propertyArr {
					names = append(names, property.Name)
				}
				coupon.PropertyNames = gstr.Implode("，", names)
			}

			if !g.IsEmpty(coupon.StartTime) && gtime.Now().Format("y-m-d") < coupon.StartTime.Format("y-m-d") {
				IsNotUse = append(IsNotUse, coupon)
				continue
			}

			if !g.IsEmpty(coupon.EndTime) && gtime.Now().Format("y-m-d") > coupon.EndTime.Format("y-m-d") {
				IsNotUse = append(IsNotUse, coupon)
				continue
			}

			if !g.IsEmpty(coupon.UseTime) {
				IsNotUse = append(IsNotUse, coupon)
				continue
			}
			if coupon.AtLeast > AtLeast {
				IsNotUse = append(IsNotUse, coupon)
				continue
			}
			if req.Scene == 1 && !g.IsEmpty(coupon.PropertyIds) && !gstr.InArray(strings.Split(coupon.PropertyIds, ","), gvar.New(PreOrderDetail.Property.Id).String()) {
				IsNotUse = append(IsNotUse, coupon)
				continue
			}
			coupon.Use = true
			IsUse = append(IsUse, coupon)
		}
		IsUse = append(IsUse, IsNotUse...)
		res.List = IsUse
	}
	return
}
func (c *ControllerBasics) CouponView(ctx context.Context, req *basics.CouponViewReq) (res *basics.CouponViewRes, err error) {
	res = new(basics.CouponViewRes)
	if res.PmsCouponViewModel, err = service.BasicsCoupon().View(ctx, &req.PmsCouponViewInp); err != nil {
		return
	}
	return
}
