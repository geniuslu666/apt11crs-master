package input_th

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// ThMemberCouponInsertFields 新增会员礼品券券字段过滤
type ThMemberCouponInsertFields struct {
	CouponNo      string      `json:"coupono"              dc:"券号"`
	CouponId      int         `json:"couponId"      dc:"礼品券ID"`
	MemberId      int         `json:"memberId"          dc:"领用人ID"`
	State         int         `json:"state"             dc:"礼品券状态 1已领用（未使用） 2已使用 3已过期 4已关闭"`
	StartTime     *gtime.Time `json:"startTime"         dc:"可使用的开始时间"`
	EndTime       *gtime.Time `json:"endTime"           dc:"有效期结束时间"`
	Source        int         `json:"source"            dc:"来源：1-手动发放 2-自动发放(下单奖励)"`
	CountDown     int         `json:"countDown"            dc:"倒计时"`
	SourceOrderId int         `json:"sourceOrderId"   dc:"来源订单ID"`
	EmployeeId    int         `json:"employeeId"      dc:"员工ID"`
	ActivityId    uint64      `json:"activityId"    dc:"活动ID（员工福利活动）"`
}

// ThMemberCouponListInp 获取已发放礼品券列表
type ThMemberCouponListInp struct {
	input_form.PageReq
	CouponId         int           `json:"couponId"  dc:"礼品券ID"`
	CouponNo         string        `json:"couponNo"      dc:"券号"`
	CouponName       string        `json:"couponName" dc:"券名称"`
	State            int           `json:"state"         dc:"状态 1未使用  2已核销 3已过期"`
	VerifyTime       []*gtime.Time `json:"verifyTime"     dc:"核销时间"`
	VerifyMchId      int           `json:"verifyMchId"   dc:"核销商户ID"`
	VerifyStoreId    int           `json:"verifyStoreId"   dc:"核销门店ID"`
	StoreName        string        `json:"storeName" dc:"门店名称"`
	MemberId         int           `json:"memberId"      dc:"领用人"`
	EmployeeId       int           `json:"employeeId"      dc:"员工ID"`
	ActivityId       uint64        `json:"activityId"    dc:"活动ID（员工福利活动）"`
	IsVerifyTimeDesc bool          `json:"isVerifyTimeDesc" dc:"是否根据核销时间倒序 true-倒序 false-不做排序"`
}

func (in *ThMemberCouponListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMemberCouponListModel struct {
	entity.ThMemberCoupon
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
	ThCoupon *struct {
		gmeta.Meta `orm:"table:hg_th_coupon"`
		*entity.ThCoupon
	} `json:"thCoupon" orm:"with:id=couponId" dc:"礼品券"`
	VerifyMch *struct {
		gmeta.Meta `orm:"table:hg_th_mch"`
		*entity.ThMch
	} `json:"verifyMch" orm:"with:id=verifyMchId" dc:"核销商户信息"`
	VerifyStore *struct {
		gmeta.Meta `orm:"table:hg_th_mch_store"`
		*entity.ThMchStore
	} `json:"verifyStore" orm:"with:id=verifyStoreId" dc:"核销门店信息"`
	ThCouponMch []*struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
	} `json:"thCouponMch" orm:"with:coupon_id=coupon_id" dc:"礼品券适用门店"`
	ThCouponMchName  string `json:"thCouponMchName" dc:"适用门店名称"`
	EmployeeActivity *struct {
		gmeta.Meta `orm:"table:hg_employee_activity"`
		*entity.EmployeeActivity
	} `json:"employeeActivity" orm:"with:id=activity_id" dc:"员工福利活动"`
	EmployeeInfo *struct {
		gmeta.Meta `orm:"table:hg_employee"`
		Id         uint64 `json:"id"           dc:"员工ID"`
		Name       string `json:"name"         dc:"员工姓名"`
		PhoneArea  string `json:"phoneArea"    dc:"电话区号"`
		Phone      string `json:"phone"        dc:"电话号码"`
	} `json:"employeeInfo" orm:"with:id=employee_id" dc:"员工信息"`
	IndexActivity *struct {
		gmeta.Meta `orm:"table:hg_homepage_articles"`
		Id         uint64 `json:"id"              dc:"活动ID"`
		Title      string `json:"title"            dc:"活动标题"`
	} `json:"indexActivity" orm:"with:id=index_activity_id" dc:"首页活动"`
}

// ThMemberCouponVerifyInp 会员礼品券核销
type ThMemberCouponVerifyInp struct {
	Sign      string `json:"sign"      dc:"签名，sn+timestamp+val+key 拼接得到的字符串，再使用 SHA-1 算法加密得到"`
	Sn        string `json:"sn"        dc:"打印机编号"`
	Timestamp int    `json:"timestamp" dc:"时间戳"`
	Val       string `json:"val"       dc:"扫描枪回调内容"`
}

// ThMemberCouponVerifyFields 会员礼品券核销字段过滤
type ThMemberCouponVerifyFields struct {
	State         int         `json:"state"          dc:"状态 1待生效 2未使用 3已核销 4已过期  5已失效"`
	VerifyTime    *gtime.Time `json:"verifyTime"     dc:"核销时间"`
	VerifyMchId   int         `json:"verifyMchId"    dc:"核销商户ID"`
	VerifyStoreId int         `json:"verifyStoreId"  dc:"核销门店ID"`
}

type ThMemberCouponViewModel struct {
	entity.ThMemberCoupon
	Member *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
	ThCoupon *struct {
		gmeta.Meta `orm:"table:hg_th_coupon"`
		*entity.ThCoupon
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=coupon_name, where:language='ja'"`
	} `json:"thCoupon" orm:"with:id=couponId" dc:"礼品券"`
	VerifyMch *struct {
		gmeta.Meta `orm:"table:hg_th_mch"`
		*entity.ThMch
	} `json:"verifyMch" orm:"with:id=verifyMchId" dc:"核销商户信息"`
	VerifyStore *struct {
		gmeta.Meta `orm:"table:hg_th_mch_store"`
		*entity.ThMchStore
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=store_name, where:language='ja'"`
	} `json:"verifyStore" orm:"with:id=verifyStoreId" dc:"核销门店信息"`
	ThCouponMch *struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
	} `json:"thCouponMch" orm:"with:coupon_id=coupon_id, with:mch_id=verify_mch_id" dc:"礼品券适用门店"`
}

// ThMemberCouponAppListInp 获取已发放礼品券列表
type ThMemberCouponAppListInp struct {
	input_form.PageReq
	State    int `json:"state" dc:"状态 0 全部 |  1 待生效 | 2 待使用 | 3 已失效"`
	MemberId int `json:"memberId"      dc:"领用人"`
	Source   int `json:"source"        dc:"来源：1-手动发放 2-自动发放(下单奖励) 3-员工福利领取"`
}

func (in *ThMemberCouponAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMemberCouponAppListModel struct {
	Id        int         `json:"id"            dc:"领取记录ID"`
	CouponId  int         `json:"couponId"      dc:"券id"`
	StartTime *gtime.Time `json:"startTime"    dc:"有效期开始时间"`
	EndTime   *gtime.Time `json:"endTime"       dc:"有效期结束时间"`
	State     int         `json:"state"         dc:"状态 1待生效 2未使用 3已核销 4已过期  5已失效"`
	CreateAt  *gtime.Time `json:"createAt"      dc:"领取时间"`
	ThCoupon  *struct {
		gmeta.Meta    `orm:"table:hg_th_coupon"`
		Id            int    `json:"id"             dc:"券ID"`
		CouponName    string `json:"couponName"     dc:"券名称"`
		CouponSubName string `json:"couponSubName"  dc:"券副标题"`
		Logo          string `json:"logo"          dc:"LOGO"`
	} `json:"thCoupon" orm:"with:id=couponId" dc:"礼品券"`
}

type ThMemberCouponAppViewInp struct {
	Id int `json:"id" v:"required#The_collection_record_ID_cannot_be_empty" dc:"领取记录ID"`
}

func (in *ThMemberCouponAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMemberCouponAppViewModel struct {
	Id         int         `json:"id"           dc:"领取记录ID"`
	CouponNo   string      `json:"couponNo"      dc:"券号"`
	Code       string      `json:"code"          dc:"加密后券码（以此券码生成二维码）"`
	MemberId   int         `json:"memberId"      dc:"会员ID"`
	CouponId   int         `json:"couponId"      dc:"券id"`
	State      int         `json:"state"         dc:"状态 1待生效 2未使用 3已核销 4已过期  5已失效"`
	StartTime  *gtime.Time `json:"startTime"     dc:"有效期开始时间"`
	EndTime    *gtime.Time `json:"endTime"       dc:"有效期结束时间"`
	VerifyTime *gtime.Time `json:"verifyTime"    dc:"核销时间"`
	ThCoupon   *struct {
		gmeta.Meta               `orm:"table:hg_th_coupon"`
		Id                       int    `json:"id"             dc:"券ID"`
		CouponName               string `json:"couponName"     dc:"券名称"`
		CouponSubName            string `json:"couponSubName"  dc:"券副标题"`
		Desc                     string `json:"desc"           dc:"券说明"`
		NeedReservation          int    `json:"needReservation"          dc:"是否需要预约：0-不需要，1-需要"`
		ReservationRestaurantIds string `json:"reservationRestaurantIds" dc:"需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3"`
		RestaurantList           []*struct {
			Id   int    `json:"id"                  dc:"主键"`
			Name string `json:"name"                dc:"餐厅名称"`
		} `json:"restaurantList" dc:"需要预约的餐厅"`
	} `json:"thCoupon" orm:"with:id=couponId" dc:"礼品券"`
	MchList []*struct {
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
	} `json:"mchList" orm:"with:coupon_id=coupon_id" dc:"商户列表"`
}

type ThMchAppViewInp struct {
	CouponId int `json:"couponId" v:"required#券ID不能为空" dc:"券ID"`
	MchId    int `json:"mchId" v:"required#商户ID不能为空" dc:"商户ID"`
}

func (in *ThMchAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchAppViewModel struct {
	Id              int    `json:"id"         dc:"商户ID"`
	Name            string `json:"name"        dc:"名称"`
	Logo            string `json:"logo"        dc:"LOGO"`
	StoreOnNum      int    `json:"storeOnNum"  dc:"门店数量"`
	VerifyGoodsName string `json:"verifyGoodsName"        dc:"可兑换商品"`
	StoreList       []*struct {
		gmeta.Meta    `orm:"table:hg_th_mch_store"`
		Id            int64  `json:"id"            dc:"门店ID"`
		MchId         uint64 `json:"mchId"         dc:"商户ID"`
		StoreName     string `json:"storeName"     dc:"门店名称"`
		Phone         string `json:"phone"         dc:"电话"`
		DetailAddress string `json:"detailAddress" dc:"详细地址"`
		GgLat         string `json:"ggLat"         dc:"谷歌纬度"`
		GgLng         string `json:"ggLng"         dc:"谷歌经度"`
	} `json:"storeList" orm:"with:mch_id=id" dc:"门店列表"`
}

// ThInvalidMemberCouponInp 失效会员礼品券
type ThInvalidMemberCouponInp struct {
	Id            int `json:"id"           dc:"主键ID"`
	SourceOrderId int `json:"sourceOrderId"          dc:"来源订单ID"`
}

// ThMemberCouponExportModel 导出会员礼品券
type ThMemberCouponExportModel struct {
	Id                   int         `json:"id"               dc:"主键"`
	CouponNo             string      `json:"couponNo"         dc:"クーポンコード"`  // 券号
	CouponName           string      `json:"couponName"       dc:"クーポン券"`    // 券名称
	IdentityName         string      `json:"identityName"     dc:"券识别名称"`    // 券识别名称
	MemberNo             string      `json:"memberNo"         dc:"ID"`       // 会员编号
	MemberName           string      `json:"memberName"       dc:"名前"`       // 会员名称
	State                string      `json:"state"            dc:"使用状況"`     // 核销状态
	VerifyTime           *gtime.Time `json:"verifyTime"       dc:"使用時間"`     // 核销时间
	VerifyMchName        string      `json:"verifyMchName"    dc:"店舗名"`      // 核销商户名称
	VerifyStoreName      string      `json:"verifyStoreName"  dc:"使用店舗"`     // 核销门店名称
	ThCouponMchName      string      `json:"thCouponMchName"  dc:"使用クーポン"`   // 核销商品名
	StartTime            *gtime.Time `json:"startTime"        dc:"クーポン付与日時"` // 有效期开始时间
	EndTime              *gtime.Time `json:"endTime"          dc:"クーポン有効期限"` // 有效期结束时间
	CreatedAt            *gtime.Time `json:"createdAt"        dc:"クーポン受取日時"` // 创建时间
	EmployeeActivityName string      `json:"employeeActivityName" dc:"従業員アクティビティ"`
	EmployeeName         string      `json:"employeeName" dc:"従業員名"`
	IndexActivityTitle   string      `json:"indexActivityTitle" dc:"ホームイベント"`
}

// ThMemberCouponViewInp 获取指定礼品券信息
type ThMemberCouponViewInp struct {
	Id int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
}

func (in *ThMemberCouponViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMemberCouponAdminViewModel struct {
	entity.ThMemberCoupon
	Member *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
	ThCoupon *struct {
		gmeta.Meta `orm:"table:hg_th_coupon"`
		*entity.ThCoupon
		NameLanguage    []*input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=coupon_name"`
		SubNameLanguage []*input_hotel.LanguageType `json:"subNameLanguage"         dc:"礼品券副标题"   orm:"with:uuid=coupon_sub_name"`
		DescLanguage    []*input_hotel.LanguageType `json:"descLanguage"         dc:"使用说明"   orm:"with:uuid=desc"`
		MchList         []*struct {
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
	} `json:"thCoupon" orm:"with:id=couponId" dc:"礼品券"`
	VerifyMch *struct {
		gmeta.Meta `orm:"table:hg_th_mch"`
		*entity.ThMch
	} `json:"verifyMch" orm:"with:id=verifyMchId" dc:"核销商户信息"`
	VerifyStore *struct {
		gmeta.Meta `orm:"table:hg_th_mch_store"`
		*entity.ThMchStore
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"         dc:"礼品券名称"   orm:"with:uuid=store_name, where:language='ja'"`
	} `json:"verifyStore" orm:"with:id=verifyStoreId" dc:"核销门店信息"`
	ThCouponMch *struct {
		gmeta.Meta `orm:"table:hg_th_coupon_mch"`
		*entity.ThCouponMch
	} `json:"thCouponMch" orm:"with:coupon_id=coupon_id, with:mch_id=verify_mch_id" dc:"礼品券适用门店"`
	EmployeeActivity *struct {
		gmeta.Meta   `orm:"table:hg_employee_activity"`
		Id           uint64                    `json:"id"              dc:"活动ID"`
		Name         string                    `json:"name"            dc:"活动名称（多语言）"`
		NameLanguage *input_hotel.LanguageType `json:"nameLanguage"    dc:"活动名称"   orm:"with:uuid=name, where:language='ja'"`
	} `json:"employeeActivity" orm:"with:id=activity_id" dc:"员工福利活动"`
	IndexActivity *struct {
		gmeta.Meta `orm:"table:hg_homepage_articles"`
		Id         uint64 `json:"id"              dc:"活动ID"`
		Title      string `json:"title"            dc:"活动标题"`
	} `json:"indexActivity" orm:"with:id=index_activity_id" dc:"首页活动"`
	EmployeeInfo *struct {
		gmeta.Meta `orm:"table:hg_employee"`
		Id         uint64 `json:"id"              dc:"员工ID"`
		Name       string `json:"name"            dc:"员工姓名"`
	} `json:"employeeInfo" orm:"with:id=employee_id" dc:"员工信息"`
}

// ThMemberCouponRecycleInp 回收礼品券
type ThMemberCouponRecycleInp struct {
	Id int `json:"id" v:"required#礼品券ID不能为空" dc:"礼品券ID"`
}

func (in *ThMemberCouponRecycleInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("礼品券ID不能为空")
		return
	}
	return
}

type ThMemberCouponRecycleModel struct{}

// ThMemberCouponRefreshCodeInp 会员礼品券-刷新code
type ThMemberCouponRefreshCodeInp struct {
	Id       int `json:"id" v:"required#The_collection_record_ID_cannot_be_empty" dc:"领取记录ID"`
	MemberId int `json:"memberId"      dc:"领用人"`
}

func (in *ThMemberCouponRefreshCodeInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMemberCouponRefreshCodeModel struct {
	Code string `json:"code"        dc:"券码"`
}

// ThMemberCouponManualVerifyInp 会员礼品券手动核销
type ThMemberCouponManualVerifyInp struct {
	Id         int         `json:"id"        dc:"领取记录ID"`
	StoreId    int         `json:"storeId"   dc:"门店ID"`
	VerifyTime *gtime.Time `json:"verifyTime" dc:"核销时间"`
}
