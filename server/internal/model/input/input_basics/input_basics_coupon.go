package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// PmsCouponInsertFields 新增会员优惠券字段过滤
type PmsCouponInsertFields struct {
	Type            string      `json:"type"              dc:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName      string      `json:"couponName"        dc:"优惠券名称"`
	CouponTypeId    int         `json:"couponTypeId"      dc:"优惠券类型ID"`
	MemberId        int         `json:"memberId"          dc:"领用人ID"`
	Scene           int         `json:"scene"             dc:"场景 1-住宿 2-餐饮"`
	PropertyIds     string      `json:"propertyIds"       dc:"物业ID"`
	RestaurantIds   string      `json:"restaurantIds"     dc:"餐厅ID"`
	ServiceIds      string      `json:"serviceIds"        dc:"按摩服务类型"`
	CarServiceTypes string      `json:"carServiceTypes"   dc:"接送机服务类型"`
	AtLeast         float64     `json:"atLeast"           dc:"满多少元使用 0代表无限制"`
	Money           float64     `json:"money"             dc:"发放面额 当type为reward时需要添加"`
	Discount        float64     `json:"discount"          dc:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit   float64     `json:"discountLimit"     dc:"最多折扣金额 当type为discount时可选择性添加"`
	State           int         `json:"state"             dc:"优惠券状态 1已领用（未使用） 2已使用 3已过期 4已关闭"`
	FetchTime       *gtime.Time `json:"fetchTime"         dc:"领取时间"`
	StartTime       *gtime.Time `json:"startTime"         dc:"可使用的开始时间"`
	EndTime         *gtime.Time `json:"endTime"           dc:"有效期结束时间"`
	Source          int         `json:"source"            dc:"来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励"`
	SourceOrderId   int         `json:"sourceOrderId"   dc:"来源订单ID"`
	IndexActivityId int         `json:"indexActivityId"   dc:"首页活动ID"`
}

// PmsCouponListInp 获取已领取优惠券列表
type PmsCouponListInp struct {
	input_form.PageReq
	CouponTypeId int     `json:"couponTypeId"  dc:"优惠券ID"`
	State        int     `json:"state"         dc:"优惠券状态 1已领用（未使用） 2已使用 3已过期 4已关闭"`
	MemberId     int     `json:"memberId"      dc:"会员ID"`
	WantUseTime  string  `json:"wantUseTime"   dc:"想要使用时间"`
	Scene        int     `json:"scene"         dc:"场景 1 住宿 2 餐饮"`
	PropertyId   int     `json:"propertyId"    dc:"物业ID"`
	AtLeast      float64 `json:"atLeast"      dc:"满多少元使用 0代表无限制"`
	MemberKey    string  `json:"memberKey"     dc:"会员信息（会员名/手机号/邮箱）"`
	Source       int     `json:"source"        dc:"来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励"`
}

func (in *PmsCouponListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponListModel struct {
	entity.PmsCoupon
	Use    bool `json:"use" dc:"是否可用  false 不可用  true 可用"`
	Member *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
	PmsCouponType *struct {
		gmeta.Meta `orm:"table:hg_pms_coupon_type"`
		*entity.PmsCouponType
	} `json:"pmsCouponType" orm:"with:id=couponTypeId" dc:"优惠券类型"`
	PropertyNames string `json:"propertyNames"   dc:"允许使用的物业名称"`
}

// PmsCouponStatInp 获取会员优惠券概况信息
type PmsCouponStatInp struct {
}

func (in *PmsCouponStatInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCouponStatModel struct {
	TotalSendNum int `json:"totalSendNum"        dc:"累计发放数量"`
	UsedNum      int `json:"usedNum"        dc:"已使用数量"`
	WaitUseNum   int `json:"waitUseNum"        dc:"待使用数量"`
}

// PmsCouponRecycleInp 回收优惠券
type PmsCouponRecycleInp struct {
	Id int `json:"id" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
}

func (in *PmsCouponRecycleInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("优惠券ID不能为空")
		return
	}
	return
}

type PmsCouponRecycleModel struct{}

// PmsCouponInvalidInp 回收优惠券
type PmsCouponInvalidInp struct {
	SourceOrderId int `json:"sourceOrderId"          dc:"来源订单ID"`
}

type PmsCouponViewInp struct {
	Id int `json:"id" v:"required#coupon_id_cannot_be_empty" dc:"优惠券ID"`
}

type PmsCouponViewModel struct {
	entity.PmsCoupon
	PmsCouponType *struct {
		gmeta.Meta `orm:"table:hg_pms_coupon_type"`
		*entity.PmsCouponType
	} `json:"pmsCouponType" orm:"with:id=couponTypeId" dc:"优惠券类型"`
	Property []*struct {
		Id        int         `json:"id"                  dc:"主键"`
		Uid       string      `json:"uid"                 dc:"物业ID"`
		Name      string      `json:"name"                dc:"物业名称"`
		CreatedAt *gtime.Time `json:"createdAt"           dc:"创建时间"`
		Close     int         `json:"close"               dc:"1、关闭该物业  2、开启该物业"`
	} `json:"property" dc:"物业"`
	PropertyNames     string `json:"propertyNames"   dc:"允许使用的物业名称"`
	RestaurantNames   string `json:"restaurantNames"   dc:"允许使用的餐饮名称"`
	CarServiceTypeStr string `json:"carServiceTypeStr"   dc:"接送机服务类型"`
	SpaServiceTypeStr string `json:"spaServiceTypeStr"   dc:"按摩服务类型"`
}
