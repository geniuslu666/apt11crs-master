package input_basics

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCouponTypeUpdateFields 修改优惠券字段过滤
type PmsCouponTypeUpdateFields struct {
	CouponName      string  `json:"couponName"                dc:"优惠券名称"`
	Type            string  `json:"type"                      dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	Money           float64 `json:"money"                     dc:"发放面额 当type为reward时需要添加"`
	Discount        float64 `json:"discount"                  dc:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit   float64 `json:"discountLimit"             dc:"最多折扣金额 当type为discount时可选择性添加"`
	AtLeast         float64 `json:"atLeast"                   dc:"满多少元使用 0代表无限制"`
	IsShow          int     `json:"isShow"                    dc:"是否显示"`
	Count           int     `json:"count"                     dc:"发放数量"`
	MaxFetch        int     `json:"maxFetch"                  dc:"每人最大领取个数"`
	ValidityType    int     `json:"validityType"              dc:"过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效"`
	EndUseTime      int     `json:"endUseTime"                dc:"使用结束日期 过期类型0时必填"`
	FixedTerm       int     `json:"fixedTerm"                 dc:"当validity_type为1时需要添加 领取之日起或者次日N天内有效"`
	Scene           int     `json:"scene"                     dc:"场景 1-住宿 2-餐饮"`
	PropertyIds     string  `json:"propertyIds"               dc:"物业ID"`
	RestaurantIds   string  `json:"restaurantIds"             dc:"餐厅ID"`
	ServiceIds      string  `json:"serviceIds"                dc:"服务ID"`
	CarServiceTypes string  `json:"carServiceTypes"   dc:"汽车服务类型"`
	Desc            string  `json:"desc"                      dc:"优惠券使用说明"`
}

// PmsCouponTypeInsertFields 新增优惠券字段过滤
type PmsCouponTypeInsertFields struct {
	CouponName      string  `json:"couponName"                dc:"优惠券名称"`
	Type            string  `json:"type"                      dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	Money           float64 `json:"money"                     dc:"发放面额 当type为reward时需要添加"`
	Discount        float64 `json:"discount"                  dc:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit   float64 `json:"discountLimit"             dc:"最多折扣金额 当type为discount时可选择性添加"`
	AtLeast         float64 `json:"atLeast"                   dc:"满多少元使用 0代表无限制"`
	IsShow          int     `json:"isShow"                    dc:"是否显示"`
	Count           int     `json:"count"                     dc:"发放数量"`
	MaxFetch        int     `json:"maxFetch"                  dc:"每人最大领取个数"`
	ValidityType    int     `json:"validityType"              dc:"过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效"`
	EndUseTime      int     `json:"endUseTime"                dc:"使用结束日期 过期类型0时必填"`
	FixedTerm       int     `json:"fixedTerm"                 dc:"当validity_type为1时需要添加 领取之日起或者次日N天内有效"`
	Scene           int     `json:"scene"                     dc:"场景 1-住宿 2-餐饮"`
	PropertyIds     string  `json:"propertyIds"               dc:"物业ID"`
	RestaurantIds   string  `json:"restaurantIds"             dc:"餐厅ID"`
	ServiceIds      string  `json:"serviceIds"                dc:"服务ID"`
	CarServiceTypes string  `json:"carServiceTypes"   dc:"汽车服务类型"`
	Status          int     `json:"status"                    dc:"状态（1进行中2已结束-1已关闭）"`
	Desc            string  `json:"desc"                      dc:"优惠券使用说明"`
}

// PmsCouponTypeEditInp 修改/新增优惠券
type PmsCouponTypeEditInp struct {
	entity.PmsCouponType
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言优惠券名称"`
	DescLanguage input_language.LanguageModel `json:"descLanguage"          dc:"多语言优惠券使用规则"`
}

func (in *PmsCouponTypeEditInp) Filter(ctx context.Context) (err error) {
	// 验证优惠券类型 reward-满减 discount-折扣 random-随机
	if err := g.Validator().Rules("required").Data(in.Type).Messages("优惠券类型 reward-满减 discount-折扣 random-随机不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证优惠券名称
	if err := g.Validator().Rules("required").Data(in.CouponName).Messages("优惠券名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证名称备注
	//if err := g.Validator().Rules("required").Data(in.CouponNameRemark).Messages("名称备注不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证发放数量
	if err := g.Validator().Rules("required").Data(in.Count).Messages("发放数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证已领取数量
	if err := g.Validator().Rules("required").Data(in.LeadCount).Messages("已领取数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证已使用数量
	if err := g.Validator().Rules("required").Data(in.UsedCount).Messages("已使用数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证使用门槛0-无门槛 1-有门槛
	//if err := g.Validator().Rules("required").Data(in.IsLimit).Messages("使用门槛0-无门槛 1-有门槛不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证满多少元使用 0代表无限制
	if err := g.Validator().Rules("required").Data(in.AtLeast).Messages("满多少元使用 0代表无限制不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证发放面额 当type为reward时需要添加
	if err := g.Validator().Rules("required").Data(in.Money).Messages("发放面额 当type为reward时需要添加不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证1 =< 折扣 <= 9.9 当type为discount时需要添加
	if err := g.Validator().Rules("required").Data(in.Discount).Messages("1 =< 折扣 <= 9.9 当type为discount时需要添加不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证最多折扣金额 当type为discount时可选择性添加
	if err := g.Validator().Rules("required").Data(in.DiscountLimit).Messages("最多折扣金额 当type为discount时可选择性添加不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证最低金额 当type为random时需要添加
	//if err := g.Validator().Rules("required").Data(in.MinMoney).Messages("最低金额 当type为random时需要添加不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证最大金额 当type为random时需要添加
	//if err := g.Validator().Rules("required").Data(in.MaxMoney).Messages("最大金额 当type为random时需要添加不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效
	if err := g.Validator().Rules("required").Data(in.ValidityType).Messages("过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证使用开始日期 过期类型0时必填
	if err := g.Validator().Rules("required").Data(in.StartUseTime).Messages("使用开始日期 过期类型0时必填不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证使用结束日期 过期类型0时必填
	if err := g.Validator().Rules("required").Data(in.EndUseTime).Messages("使用结束日期 过期类型0时必填不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证当validity_type为1时需要添加 领取之日起或者次日N天内有效
	if err := g.Validator().Rules("required").Data(in.FixedTerm).Messages("当validity_type为1时需要添加 领取之日起或者次日N天内有效不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否无限制0-否 1是
	//if err := g.Validator().Rules("required").Data(in.IsLimitless).Messages("是否无限制0-否 1是不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证每人最大领取个数
	if err := g.Validator().Rules("required").Data(in.MaxFetch).Messages("每人最大领取个数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否开启过期提醒0-不开启 1-开启
	//if err := g.Validator().Rules("required").Data(in.IsExpireNotice).Messages("是否开启过期提醒0-不开启 1-开启不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证过期前N天提醒
	//if err := g.Validator().Rules("required").Data(in.ExpireNoticeFixedTerm).Messages("过期前N天提醒不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证是否显示
	if err := g.Validator().Rules("required").Data(in.IsShow).Messages("是否显示不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证住宿订单的优惠总金额
	if err := g.Validator().Rules("required").Data(in.DiscountAppStayOrderMoney).Messages("住宿订单的优惠总金额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证住宿订单用券总成交额
	if err := g.Validator().Rules("required").Data(in.AppStayOrderMoney).Messages("住宿订单用券总成交额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否禁止发放0-否 1-是
	//if err := g.Validator().Rules("required").Data(in.IsForbidden).Messages("是否禁止发放0-否 1-是不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	// 验证状态（1进行中2已结束-1已关闭）
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态（1进行中2已结束-1已关闭）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsCouponTypeEditModel struct{}

// PmsCouponTypeDeleteInp 删除优惠券
type PmsCouponTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
}

func (in *PmsCouponTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponTypeDeleteModel struct{}

// PmsCouponTypeViewInp 获取指定优惠券信息
type PmsCouponTypeViewInp struct {
	Id int `json:"id" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
}

func (in *PmsCouponTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponTypeViewModel struct {
	entity.PmsCouponType
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"优惠券名称"   orm:"with:uuid=coupon_name"`
	DescLanguage []*input_hotel.LanguageType `json:"descLanguage"         dc:"使用说明"   orm:"with:uuid=desc"`
	Property     []*struct {
		Id        int         `json:"id"                  dc:"主键"`
		Uid       string      `json:"uid"                 dc:"物业ID"`
		Name      string      `json:"name"                dc:"物业名称"`
		CreatedAt *gtime.Time `json:"createdAt"           dc:"创建时间"`
		Close     int         `json:"close"               dc:"1、关闭该物业  2、开启该物业"`
	} `json:"property" dc:"物业"`
	Restaurant []*struct {
		Id        int         `json:"id"                  dc:"主键"`
		Name      string      `json:"name"                dc:"餐厅名称"`
		CreatedAt *gtime.Time `json:"createdAt"           dc:"创建时间"`
	} `json:"restaurant" dc:"餐厅"`
	Service []*struct {
		Id        int         `json:"id"                  dc:"主键"`
		Name      string      `json:"name"                dc:"服务名称"`
		CreatedAt *gtime.Time `json:"createdAt"           dc:"创建时间"`
	} `json:"service" dc:"服务"`
	PropertyNames     string `json:"propertyNames"        dc:"允许使用的物业名称"`
	RestaurantNames   string `json:"restaurantNames"      dc:"允许使用的餐饮名称"`
	CarServiceTypeStr string `json:"carServiceTypeStr"    dc:"接送机服务类型"`
	SpaServiceTypeStr string `json:"spaServiceTypeStr"    dc:"按摩服务类型"`
}

// PmsCouponTypeListInp 获取优惠券列表
type PmsCouponTypeListInp struct {
	input_form.PageReq
	Type         string `json:"type"         dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName   string `json:"couponName"   dc:"优惠券名称"`
	ValidityType int    `json:"validityType" dc:"过期类型1-固定时间范围过期 2-领取之日固定日期后过期 3长期有效"`
	Status       int    `json:"status"       dc:"状态（1进行中2已结束-1已关闭）"`
	Scene        int    `json:"scene"       dc:"场景（1-酒店民宿）"`
	CouponIds    string `json:"couponIds"   dc:"优惠券Ids"`
}

func (in *PmsCouponTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponTypeListModel struct {
	Id                        int         `json:"id"                        dc:"优惠券ID"`
	Type                      string      `json:"type"                      dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName                string      `json:"couponName"                dc:"优惠券名称"`
	Count                     int         `json:"count"                     dc:"发放数量"`
	LeadCount                 int         `json:"leadCount"                 dc:"已领取数量"`
	UsedCount                 int         `json:"usedCount"                 dc:"已使用数量"`
	AtLeast                   float64     `json:"atLeast"                   dc:"满多少元使用 0代表无限制"`
	Money                     float64     `json:"money"                     dc:"发放面额 当type为reward时需要添加"`
	Discount                  float64     `json:"discount"                  dc:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit             float64     `json:"discountLimit"             dc:"最多折扣金额 当type为discount时可选择性添加"`
	ValidityType              int         `json:"validityType"              dc:"过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效"`
	StartUseTime              int         `json:"startUseTime"              dc:"使用开始日期 过期类型0时必填"`
	EndUseTime                string      `json:"endUseTime"                dc:"使用结束日期 过期类型0时必填"`
	FixedTerm                 int         `json:"fixedTerm"                 dc:"当validity_type为1时需要添加 领取之日起或者次日N天内有效"`
	Sort                      int         `json:"sort"                      dc:"排序"`
	MaxFetch                  int         `json:"maxFetch"                  dc:"每人最大领取个数"`
	IsShow                    int         `json:"isShow"                    dc:"是否显示"`
	DiscountAppStayOrderMoney float64     `json:"discountAppStayOrderMoney" dc:"住宿订单的优惠总金额"`
	AppStayOrderMoney         float64     `json:"appStayOrderMoney"         dc:"住宿订单用券总成交额"`
	IsForbidden               int         `json:"isForbidden"               dc:"是否禁止发放0-否 1-是"`
	Status                    int         `json:"status"                    dc:"状态（1进行中2已结束-1已关闭）"`
	Scene                     int         `json:"scene"                     dc:"场景（1-酒店民宿）"`
	CreateAt                  *gtime.Time `json:"createAt"                  dc:"创建时间"`
	UpdateAt                  *gtime.Time `json:"updateAt"                  dc:"修改时间"`
}

type PmsCouponTypeAllListModel struct {
	Id           int     `json:"id"                        dc:"优惠券ID"`
	Type         string  `json:"type"                      dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName   string  `json:"couponName"                dc:"优惠券名称"`
	AtLeast      float64 `json:"atLeast"                   dc:"满多少元使用 0代表无限制"`
	Money        float64 `json:"money"                     dc:"发放面额 当type为reward时需要添加"`
	Discount     float64 `json:"discount"                  dc:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	Scene        int     `json:"scene"                     dc:"场景（1-酒店民宿）"`
	ValidityType int     `json:"validityType"              dc:"过期类型0-古固定时间范围过期 1-领取之日固定日期后过期 2长期有效"`
	StartUseTime int     `json:"startUseTime"              dc:"使用开始日期 过期类型0时必填"`
	EndUseTime   string  `json:"endUseTime"                dc:"使用结束日期 过期类型0时必填"`
	FixedTerm    int     `json:"fixedTerm"                 dc:"当validity_type为1时需要添加 领取之日起或者次日N天内有效"`
}

// PmsCouponTypeMaxSortInp 获取优惠券最大排序
type PmsCouponTypeMaxSortInp struct{}

func (in *PmsCouponTypeMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponTypeMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// PmsCouponTypeStatusInp 更新优惠券状态
type PmsCouponTypeStatusInp struct {
	Id     int `json:"id" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsCouponTypeStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("优惠券ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type PmsCouponTypeStatusModel struct{}

// PmsSendMemberCouponInp 发放会员优惠券
type PmsSendMemberCouponInp struct {
	CouponTypeId  int `json:"couponTypeId"          dc:"优惠券ID"`
	MemberId      int `json:"memberId"          dc:"会员ID"`
	SourceOrderId int `json:"sourceOrderId"          dc:"来源订单ID"`
}

type PmsSendMemberCouponModel struct{}

// PmsCouponTypeAppViewInp 获取指定优惠券信息
type PmsCouponTypeAppViewInp struct {
	Id              int `json:"id" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
	IndexActivityId int `json:"indexActivityId" dc:"首页活动ID"`
}

func (in *PmsCouponTypeAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponTypeAppViewModel struct {
	entity.PmsCouponType
	ActivityId        int    `json:"activityId"     dc:"活动ID"`
	ActivityStatus    int    `json:"activityStatus"     dc:"活动状态 1-未开始  2-进行中  3-已结束"`
	ReceiveStatus     int    `json:"receiveStatus" dc:"领取状态 1-可领取  2-已领取  3-已领完  4-已关闭  5-已结束  6-今日已领"`
	PropertyNames     string `json:"propertyNames"        dc:"允许使用的物业名称"`
	RestaurantNames   string `json:"restaurantNames"      dc:"允许使用的餐饮名称"`
	CarServiceTypeStr string `json:"carServiceTypeStr"    dc:"接送机服务类型"`
	SpaServiceTypeStr string `json:"spaServiceTypeStr"    dc:"按摩服务类型"`
}

// PmsCouponTypeAppReceiveInp 领取优惠券
type PmsCouponTypeAppReceiveInp struct {
	CouponTypeId    int `json:"couponTypeId"          dc:"优惠券ID"`
	IndexActivityId int `json:"indexActivityId" dc:"首页活动ID"`
}

type PmsCouponTypeAppReceiveModel struct{}
