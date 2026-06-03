package input_th

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// ThCouponUpdateFields 修改礼品券字段过滤
type ThCouponUpdateFields struct {
	Logo                     string `json:"logo"           dc:"LOGO"`
	CouponName               string `json:"couponName"    dc:"提货券名称"`
	CouponSubName            string `json:"couponSubName" dc:"提货券副标题"`
	IdentityName             string `json:"identityName"   dc:"券标识名称"`
	CouponNoPrefix           string `json:"couponNoPrefix" dc:"编号前缀"`
	Status                   int    `json:"status"         dc:"发放状态（1立即启用  2暂不启用）"`
	FixedTerm                int    `json:"fixedTerm"      dc:"激活后几天内有效"`
	Desc                     string `json:"desc"           dc:"券说明"`
	UseMode                  string `json:"useMode"        dc:"使用模式"`
	CategoryId               uint   `json:"categoryId"     dc:"分类ID"`
	NeedReservation          int    `json:"needReservation"          dc:"是否需要预约：0-不需要，1-需要"`
	ReservationRestaurantIds string `json:"reservationRestaurantIds" dc:"需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3"`
}

// ThCouponInsertFields 新增礼品券字段过滤
type ThCouponInsertFields struct {
	Logo                     string `json:"logo"           dc:"LOGO"`
	CouponName               string `json:"couponName"    dc:"提货券名称"`
	CouponSubName            string `json:"couponSubName" dc:"提货券副标题"`
	IdentityName             string `json:"identityName"   dc:"券标识名称"`
	CouponNoPrefix           string `json:"couponNoPrefix" dc:"编号前缀"`
	Status                   int    `json:"status"         dc:"发放状态（1立即启用  2暂不启用）"`
	FixedTerm                int    `json:"fixedTerm"      dc:"激活后几天内有效"`
	Desc                     string `json:"desc"           dc:"券说明"`
	UseMode                  string `json:"useMode"        dc:"使用模式"`
	CategoryId               uint   `json:"categoryId"     dc:"分类ID"`
	NeedReservation          int    `json:"needReservation"          dc:"是否需要预约：0-不需要，1-需要"`
	ReservationRestaurantIds string `json:"reservationRestaurantIds" dc:"需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3"`
}

// ThCouponEditInp 修改/新增礼品券
type ThCouponEditInp struct {
	entity.ThCoupon
	NameLanguage    input_language.LanguageModel `json:"nameLanguage"          dc:"多语言礼品券名称"`
	SubNameLanguage input_language.LanguageModel `json:"subNameLanguage"          dc:"多语言礼品券副标题"`
	DescLanguage    input_language.LanguageModel `json:"descLanguage"          dc:"多语言礼品券使用规则"`
	MchList         []*struct {
		MchId int    `json:"mchId"     dc:"商户ID"`
		Name  string `json:"name"     dc:"核销商品名"`
	} `json:"mchList"          dc:"商户列表"`
}

func (in *ThCouponEditInp) Filter(ctx context.Context) (err error) {

	return
}

type ThCouponEditModel struct{}

// ThCouponDeleteInp 删除礼品券
type ThCouponDeleteInp struct {
	Id interface{} `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
}

func (in *ThCouponDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponDeleteModel struct{}

// ThCouponViewInp 获取指定礼品券信息
type ThCouponViewInp struct {
	Id int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
}

func (in *ThCouponViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponViewModel struct {
	entity.ThCoupon
	NameLanguage    []*input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=coupon_name"`
	SubNameLanguage []*input_hotel.LanguageType `json:"subNameLanguage"         dc:"礼品券副标题"   orm:"with:uuid=coupon_sub_name"`
	DescLanguage    []*input_hotel.LanguageType `json:"descLanguage"         dc:"使用说明"   orm:"with:uuid=desc"`
	CategoryInfo    *struct {
		gmeta.Meta `orm:"table:hg_th_coupon_category"`
		Id         int    `json:"id"        dc:""`
		Name       string `json:"name"      dc:"分类名称"`
	} `json:"categoryInfo" orm:"with:id=category_id" dc:"礼品券分类信息"`
	MchList []*struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
		MchInfo *struct {
			gmeta.Meta `orm:"table:hg_th_mch"`
			*entity.ThMch
			CategoryInfo *struct {
				gmeta.Meta `orm:"table:hg_th_mch_category"`
				Id         int    `json:"id"        dc:""`
				Name       string `json:"name"      dc:"分类名称"`
			} `json:"categoryInfo" orm:"with:id=category_id" dc:"商户分类信息"`
		} `json:"mchInfo" orm:"with:id=mch_id" dc:"商户信息"`
	} `json:"mchList" orm:"with:coupon_id=id" dc:"商户列表"`
	RestaurantNames string `json:"restaurantNames"      dc:"需预约的餐饮名称"`
}

// ThCouponListInp 获取礼品券列表
type ThCouponListInp struct {
	input_form.PageReq
	CouponName   string        `json:"couponName"   dc:"礼品券名称"`
	IdentityName string        `json:"identityName"   dc:"券标识名称"`
	Status       int           `json:"status"         dc:"发放状态（1立即启用  2暂不启用）"`
	UseStatus    int           `json:"useStatus"     dc:"使用状态（1开始使用  2停止使用）"`
	CategoryId   uint          `json:"categoryId"     dc:"分类ID"`
	CreateAt     []*gtime.Time `json:"createAt"     dc:"创建时间"`
	CouponIds    string        `json:"couponIds"   dc:"券Ids"`
}

func (in *ThCouponListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponListModel struct {
	Id           int    `json:"id"             dc:"礼品券ID"`
	Logo         string `json:"logo"           dc:"LOGO"`
	CouponName   string `json:"couponName"     dc:"礼品券名称"`
	IdentityName string `json:"identityName"   dc:"券标识名称"`
	CategoryId   uint   `json:"categoryId"     dc:"分类ID"`
	FixedTerm    int    `json:"fixedTerm"      dc:"激活后几天内有效"`
	Count        int    `json:"count"          dc:"发放数量"`
	UsedCount    int    `json:"usedCount"      dc:"已使用数量"`
	Status       int    `json:"status"         dc:"发放状态（1立即启用  2暂不启用）"`
	UseStatus    int    `json:"useStatus"      dc:"使用状态（1开始使用  2停止使用）"`
	MchList      []*struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
		MchInfo *struct {
			gmeta.Meta `orm:"table:hg_th_mch"`
			*entity.ThMch
			CategoryInfo *struct {
				gmeta.Meta `orm:"table:hg_th_mch_category"`
				Id         int    `json:"id"        dc:""`
				Name       string `json:"name"      dc:"分类名称"`
			} `json:"categoryInfo" orm:"with:id=category_id" dc:"商户分类信息"`
		} `json:"mchInfo" orm:"with:id=mch_id" dc:"商户信息"`
	} `json:"mchList" orm:"with:coupon_id=id" dc:"商户列表"`
	CreateAt *gtime.Time `json:"createAt"                  dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt"                  dc:"修改时间"`
}

type ThCouponAllListModel struct {
	Id         int    `json:"id"                        dc:"礼品券ID"`
	CouponName string `json:"couponName"                dc:"礼品券名称"`
	CategoryId uint   `json:"categoryId"     dc:"分类ID"`
	FixedTerm  int    `json:"fixedTerm"      dc:"激活后几天内有效"`
	Count      int    `json:"count"          dc:"发放数量"`
	UsedCount  int    `json:"usedCount"      dc:"已使用数量"`
	Status     int    `json:"status"         dc:"发放状态（1立即启用  2暂不启用）"`
	UseStatus  int    `json:"useStatus"     dc:"使用状态（1开始使用  2停止使用）"`
	MchList    []*struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
		MchInfo *struct {
			gmeta.Meta `orm:"table:hg_th_mch"`
			*entity.ThMch
			CategoryInfo *struct {
				gmeta.Meta `orm:"table:hg_th_mch_category"`
				Id         int    `json:"id"        dc:""`
				Name       string `json:"name"      dc:"分类名称"`
			} `json:"categoryInfo" orm:"with:id=category_id" dc:"商户分类信息"`
		} `json:"mchInfo" orm:"with:id=mch_id" dc:"商户信息"`
	} `json:"mchList" orm:"with:coupon_id=id" dc:"商户列表"`
	CreateAt *gtime.Time `json:"createAt"                  dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt"                  dc:"修改时间"`
}

// ThCouponMaxSortInp 获取礼品券最大排序
type ThCouponMaxSortInp struct{}

func (in *ThCouponMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// ThCouponStatusInp 更新礼品券状态
type ThCouponStatusInp struct {
	Id     int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *ThCouponStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("礼品券ID不能为空")
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

type ThCouponStatusModel struct{}

// ThCouponUseStatusInp 更新礼品券适用状态
type ThCouponUseStatusInp struct {
	Id     int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *ThCouponUseStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("礼品券ID不能为空")
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

type ThCouponUseStatusModel struct{}

// ThSendMemberCouponInp 发放会员礼品券
type ThSendMemberCouponInp struct {
	CouponId      int    `json:"couponId"          dc:"礼品券ID"`
	MemberId      int    `json:"memberId"          dc:"会员ID"`
	StartTime     string `json:"startTime"            dc:"开始时间"`
	SourceOrderId int    `json:"sourceOrderId"          dc:"来源订单ID"`
	ActivityId    uint64 `json:"activityId"        dc:"活动ID（员工福利活动）"`
	EmployeeId    uint64 `json:"employeeId"        dc:"员工ID（员工福利发放）"`
}

type ThSendMemberCouponModel struct{}

// ThCouponAppViewInp 获取指定礼品券App信息
type ThCouponAppViewInp struct {
	Id              int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
	ActivityId      int `json:"activityId" dc:"员工活动ID"`
	IndexActivityId int `json:"indexActivityId" dc:"首页活动ID"`
}

func (in *ThCouponAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponAppViewModel struct {
	Id                int    `json:"id"             dc:"券ID"`
	CouponName        string `json:"couponName"     dc:"券名称"`
	CouponSubName     string `json:"couponSubName"  dc:"券副标题"`
	Desc              string `json:"desc"           dc:"券说明"`
	Logo              string `json:"logo"          dc:"LOGO"`
	ActivityId        int    `json:"activityId"     dc:"活动ID"`
	ActivityStatus    int    `json:"activityStatus"     dc:"活动状态 1-未开始  2-进行中  3-已结束"`
	HaveReceived      int    `json:"haveReceived"        dc:"本人已领取总数量"`
	AvailableQuantity int    `json:"availableQuantity" dc:"每人可领取总数量（0表示无限制）"`
	TodayReceive      int    `json:"todayReceive" dc:"本人N天内已领取数量"`
	PerDayAvailable   int    `json:"perDayAvailable" dc:"N天内可领取总数量（0是不限制）"`
	LimitDays         int    `json:"limitDays" dc:"限制天数"`
	Status            int    `json:"status"        dc:"状态 1-可领取 2-已领取 3-已领完 4-周期内已领满  5-今日不可领"`
	MchList           []*struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		CouponId   int64  `json:"couponId" dc:"券ID"`
		MchId      int64  `json:"mchId"   dc:"商户ID"`
		Name       string `json:"name"     dc:"核销商品名(可兑商品)"`
		MchInfo    *struct {
			gmeta.Meta `orm:"table:hg_th_mch"`
			Id         int    `json:"id"             dc:"门店ID"`
			Name       string `json:"name"        dc:"名称"`
			Logo       string `json:"logo"        dc:"LOGO"`
			StoreOnNum int    `json:"storeOnNum"  dc:"门店数量"`
		} `json:"mchInfo" orm:"with:id=mch_id" dc:"商户信息"`
	} `json:"mchList" orm:"with:coupon_id=id" dc:"商户列表"`
	NeedReservation          int    `json:"needReservation"          dc:"是否需要预约：0-不需要，1-需要"`
	ReservationRestaurantIds string `json:"reservationRestaurantIds" dc:"需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3"`
	RestaurantList           []*struct {
		Id   int    `json:"id"                  dc:"主键"`
		Name string `json:"name"                dc:"餐厅名称"`
	} `json:"restaurantList" dc:"需要预约的餐厅"`
}

// ThCouponAppReceiveInp 领取礼品券
type ThCouponAppReceiveInp struct {
	CouponId        int `json:"couponId"          dc:"礼品券ID"`
	IndexActivityId int `json:"indexActivityId" dc:"首页活动ID"`
}

type ThCouponAppReceiveModel struct{}
