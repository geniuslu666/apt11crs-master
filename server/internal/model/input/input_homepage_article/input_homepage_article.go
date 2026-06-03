package input_homepage_article

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// HomepageArticlesUpdateFields 修改首页活动字段过滤
type HomepageArticlesUpdateFields struct {
	Language              string      `json:"language"   dc:"语言"`
	No                    string      `json:"no"          dc:"活动编号"`
	ListPic               string      `json:"listPic"    dc:"列表图"`
	Title                 string      `json:"title"      dc:"标题"`
	Author                string      `json:"author"     dc:"作者"`
	Content               string      `json:"content"    dc:"内容"`
	RecommendLinkType     string      `json:"recommendLinkType" dc:"推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页"`
	CouponLimit           int         `json:"couponLimit"       dc:"每人每张券可领取数量"`
	HotelIds              string      `json:"hotelIds"          dc:"民宿ids"`
	RestaurantIds         string      `json:"restaurantIds"     dc:"餐厅ids"`
	SpaServiceIds         string      `json:"spaServiceIds"     dc:"按摩ids"`
	StartTime             *gtime.Time `json:"startTime"         dc:"开始时间"`
	EndTime               *gtime.Time `json:"endTime"           dc:"结束时间"`
	Status                uint        `json:"status"    dc:"1、启用 2、禁用"`
	Sort                  int         `json:"sort"       dc:"排序(越大越靠前)"`
	FoodIndexTitle        string      `json:"foodIndexTitle" dc:"餐厅首页标题"`
	SpaIndexTitle         string      `json:"spaIndexTitle" dc:"按摩首页标题"`
	CarIndexTitle         string      `json:"carIndexTitle" dc:"接送机首页标题"`
	CouponLinkTitle       string      `json:"couponLinkTitle" dc:"券链接区域标题"`
	CouponLinkSubTitle    string      `json:"couponLinkSubTitle" dc:"券链接区域副标题"`
	RecommendLinkTitle    string      `json:"recommendLinkTitle" dc:"推荐链接区域标题"`
	RecommendLinkSubTitle string      `json:"recommendLinkSubTitle" dc:"推荐链接区域标题"`
	ButtonTxt             string      `json:"buttonTxt"             dc:"外链按钮文字"`
	OutLinkTitle          string      `json:"outLinkTitle"          dc:"外链推荐标题"`
	OutLinkContent        string      `json:"outLinkContent"        dc:"外链推荐链接"`
	MinappStatus          int         `json:"minappStatus"          dc:"是否排除小程序显示 1-不排除 2-排除"`
	Chain                 string      `json:"chain" dc:"活动链接方式 IN-内部链接 OUT-外部链接"`
	Path                  string      `json:"path" dc:"活动外链链接"`
	LinkOpenType          string      `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
	LimitWeek             string      `json:"limitWeek"             dc:"周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔"`
}

// HomepageArticlesInsertFields 新增首页活动字段过滤
type HomepageArticlesInsertFields struct {
	Language              string      `json:"language"   dc:"语言"`
	No                    string      `json:"no"          dc:"活动编号"`
	ListPic               string      `json:"listPic"    dc:"列表图"`
	Title                 string      `json:"title"      dc:"标题"`
	Author                string      `json:"author"     dc:"作者"`
	Content               string      `json:"content"    dc:"内容"`
	RecommendLinkType     string      `json:"recommendLinkType" dc:"推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页"`
	CouponLimit           int         `json:"couponLimit"       dc:"每人每张券可领取数量"`
	HotelIds              string      `json:"hotelIds"          dc:"民宿ids"`
	RestaurantIds         string      `json:"restaurantIds"     dc:"餐厅ids"`
	SpaServiceIds         string      `json:"spaServiceIds"     dc:"按摩ids"`
	StartTime             *gtime.Time `json:"startTime"         dc:"开始时间"`
	EndTime               *gtime.Time `json:"endTime"           dc:"结束时间"`
	Status                uint        `json:"status"     dc:"1、启用 2、禁用"`
	Sort                  int         `json:"sort"       dc:"排序(越大越靠前)"`
	FoodIndexTitle        string      `json:"foodIndexTitle" dc:"餐厅首页标题"`
	SpaIndexTitle         string      `json:"spaIndexTitle" dc:"按摩首页标题"`
	CarIndexTitle         string      `json:"carIndexTitle" dc:"接送机首页标题"`
	CouponLinkTitle       string      `json:"couponLinkTitle" dc:"券链接区域标题"`
	CouponLinkSubTitle    string      `json:"couponLinkSubTitle" dc:"券链接区域副标题"`
	RecommendLinkTitle    string      `json:"recommendLinkTitle" dc:"推荐链接区域标题"`
	RecommendLinkSubTitle string      `json:"recommendLinkSubTitle" dc:"推荐链接区域标题"`
	ButtonTxt             string      `json:"buttonTxt"             dc:"外链按钮文字"`
	OutLinkTitle          string      `json:"outLinkTitle"          dc:"外链推荐标题"`
	OutLinkContent        string      `json:"outLinkContent"        dc:"外链推荐链接"`
	MinappStatus          int         `json:"minappStatus"          dc:"是否排除小程序显示 1-不排除 2-排除"`
	Chain                 string      `json:"chain" dc:"活动链接方式 IN-内部链接 OUT-外部链接"`
	Path                  string      `json:"path" dc:"活动外链链接"`
	LinkOpenType          string      `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
	LimitWeek             string      `json:"limitWeek"             dc:"周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔"`
}

// HomepageArticlesEditInp 修改/新增首页活动
type HomepageArticlesEditInp struct {
	entity.HomepageArticles
	CouponsArr []*struct {
		CouponId          int `json:"couponId"          dc:"提货券ID"`
		AvailableQuantity int `json:"availableQuantity" dc:"领取数量"`
		PerDayAvailable   int `json:"perDayAvailable"   dc:"N天内可领取数量"`
		LimitDays         int `json:"limitDays"         dc:"限制天数"`
	} `json:"couponsArr"          dc:"优惠券列表"`
	ThCouponsArr []*struct {
		CouponId          int `json:"couponId"          dc:"提货券ID"`
		AvailableQuantity int `json:"availableQuantity" dc:"领取数量"`
		PerDayAvailable   int `json:"perDayAvailable"   dc:"N天内可领取数量"`
		LimitDays         int `json:"limitDays"         dc:"限制天数"`
		FixedTerm         int `json:"fixedTerm"         dc:"领取后几天内有效（0表示领取当日23:59:59）"`
	} `json:"thCouponsArr"          dc:"提货券列表"`
	OutLinks []*struct {
		Title     string `json:"title"      dc:"外链标题"`
		Link      string `json:"link"       dc:"外链链接"`
		ButtonTxt string `json:"buttonTxt"  dc:"按钮文字"`
		OpenType  int    `json:"openType"   dc:"跳转方式 1-内部webview 2-外部浏览器"`
	} `json:"outLinks"          dc:"外链列表"`
}

func (in *HomepageArticlesEditInp) Filter(ctx context.Context) (err error) {
	if in.Language == "" {
		err = gerror.New("请选择语言")
		return
	}
	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}
	if in.Content == "" {
		err = gerror.New("内容不能为空")
		return
	}
	return
}

type HomepageArticlesEditModel struct{}

// HomepageArticlesDeleteInp 删除首页活动
type HomepageArticlesDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *HomepageArticlesDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesDeleteModel struct{}

// HomepageArticlesStatusInp 更新首页活动状态
type HomepageArticlesStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *HomepageArticlesStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
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

type HomepageArticlesStatusModel struct{}

// HomepageArticlesViewInp 获取指定首页活动信息
type HomepageArticlesViewInp struct {
	Id int `json:"id" v:"required#id_cannot_be_empty" dc:"id"`
}

func (in *HomepageArticlesViewInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesViewModel struct {
	entity.HomepageArticles
	CouponList []*struct {
		gmeta.Meta        `orm:"table:hg_homepage_article_coupon"`
		ArticleId         int    `json:"articleId"         dc:"活动ID"`
		CouponId          int64  `json:"couponId"          dc:"提货券ID"`
		CouponType        string `json:"couponType"        dc:"优惠券类型 coupon-优惠券 thCoupon-提货券"`
		AvailableQuantity int    `json:"availableQuantity" dc:"可领取数量"`
		PerDayAvailable   int    `json:"perDayAvailable"   dc:"N天内可领取数量（0表示无限制）"`
		LimitDays         int    `json:"limitDays"         dc:"限制天数"`
	} `json:"couponList" orm:"with:article_id=id" dc:"券列表"`
	OutLinks []*struct {
		Title     string `json:"title"      dc:"外链标题"`
		Link      string `json:"link"       dc:"外链链接"`
		ButtonTxt string `json:"buttonTxt"  dc:"按钮文字"`
		OpenType  int    `json:"openType"   dc:"跳转方式 1-内部webview 2-外部浏览器"`
	} `json:"outLinks" dc:"外链列表"`
}

// HomepageArticlesListInp 获取首页活动列表
type HomepageArticlesListInp struct {
	input_form.PageReq
	Title     string        `json:"title"              dc:"标题"`
	Language  string        `json:"language"           dc:"语言"`
	Keyword   string        `json:"keyword"            dc:"关键字"`
	Status    int           `json:"status"             dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt"          dc:"created_at"`
}

func (in *HomepageArticlesListInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesListModel struct {
	Id        int         `json:"id"                dc:"id"`
	Language  string      `json:"language"          dc:"语言"`
	No        string      `json:"no"                dc:"活动编号"`
	ListPic   string      `json:"listPic"           dc:"列表图"`
	Title     string      `json:"title"             dc:"标题"`
	Author    string      `json:"author"            dc:"作者"`
	Views     int         `json:"views"             dc:"浏览量"`
	ThumbNum  int         `json:"thumbNum"          dc:"点赞数量"`
	Status    uint        `json:"status"            dc:"1、启用 2、禁用"`
	Sort      int         `json:"sort"              dc:"排序(越大越靠前)"`
	CreatedAt *gtime.Time `json:"createdAt"         dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt"         dc:"修改时间"`
}

// HomepageArticlesAppListInp 获取首页活动列表
type HomepageArticlesAppListInp struct {
	input_form.PageReq
}

func (in *HomepageArticlesAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesAppListModel struct {
	Id           int    `json:"id"                dc:"id"`
	Language     string `json:"language"          dc:"语言"`
	ListPic      string `json:"listPic"           dc:"列表图"`
	Title        string `json:"title"             dc:"标题"`
	Author       string `json:"author"            dc:"作者"`
	Views        int    `json:"views"             dc:"浏览量"`
	ThumbNum     int    `json:"thumbNum"          dc:"点赞数量"`
	HasThumb     bool   `json:"hasThumb"          dc:"是否点赞"`
	StartTime    string `json:"startTime"         dc:"开始时间"`
	EndTime      string `json:"endTime"           dc:"结束时间"`
	Sort         int    `json:"sort"              dc:"排序(越大越靠前)"`
	Chain        string `json:"chain"             dc:"活动链接方式 IN-内部链接 OUT-外部链接"`
	Path         string `json:"path"              dc:"活动外链链接"`
	LinkOpenType string `json:"linkOpenType"      dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// HomepageArticlesAppViewInp 获取指定活动信息
type HomepageArticlesAppViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *HomepageArticlesAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesAppViewModel struct {
	Id                    int                      `json:"id"                dc:""`
	Title                 string                   `json:"title"             dc:"标题"`
	Content               string                   `json:"content"           dc:"内容"`
	Views                 int                      `json:"views"             dc:"浏览量"`
	ThumbNum              int                      `json:"thumbNum"          dc:"点赞数量"`
	HasThumb              bool                     `json:"hasThumb"          dc:"是否点赞"`
	StartTime             string                   `json:"startTime"         dc:"开始时间"`
	EndTime               string                   `json:"endTime"           dc:"结束时间"`
	ActivityStatus        int                      `json:"activityStatus"     dc:"活动状态 1-未开始  2-进行中  3-已结束"`
	CouponLimit           int                      `json:"couponLimit"       dc:"每人每张券可领取数量(已废弃)"`
	CouponLinkTitle       string                   `json:"couponLinkTitle"   dc:"券链接区域标题"`
	CouponLinkSubTitle    string                   `json:"couponLinkSubTitle" dc:"券链接区域副标题"`
	RecommendLinkTitle    string                   `json:"recommendLinkTitle" dc:"推荐链接区域标题"`
	RecommendLinkSubTitle string                   `json:"recommendLinkSubTitle" dc:"推荐链接区域副标题"`
	CouponLinkList        []*CouponAppLinkListItem `json:"couponLinkList" dc:"券链接列表"`
	RecommendLinkList     []*AppLinkListItem       `json:"recommendLinkList" dc:"推荐链接列表"`
}

type CouponAppLinkListItem struct {
	CouponId           int     `json:"couponId"      dc:"券ID"`
	SubTitle           string  `json:"subTitle" dc:"副标题（礼品券才有值，优惠券为空）"`
	AtLeast            float64 `json:"atLeast" dc:"满多少元使用 0代表无限制（优惠券才有值，礼品券为空）"`
	Status             int     `json:"status"        dc:"状态 1-可领取 2-已领取 3-已领完 4-今日已领 5-今日不可领"`
	TotalNum           int     `json:"totalNum"      dc:"券总数量"`
	TotalReceive       int     `json:"totalReceive"  dc:"总领取数量"`
	MemberReceive      int     `json:"memberReceive" dc:"会员已领取数量"`
	MemberTodayReceive int     `json:"memberTodayReceive" dc:"会员N天内已领取数量"`
	AvailableQuantity  int     `json:"availableQuantity" dc:"每人可领取总数量"`
	PerDayAvailable    int     `json:"perDayAvailable" dc:"N天内可领取总数量（0是不限制）"`
	LimitDays          int     `json:"limitDays"       dc:"限制天数"`
	*AppLinkListItem
}

type AppLinkListItem struct {
	Title       string `json:"title" dc:"标题"`
	LinkScene   string `json:"linkScene" dc:"链接场景 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页"`
	WxLink      string `json:"wxLink" dc:"微信跳转链接"`
	AppLinkType int    `json:"appLinkType" dc:"app跳转类型 【0-不带参数跳转|1-带string参数跳转|2-带param参数跳转|3-跳转网页|4-无需跳转】"`
	AppLink     string `json:"appLink" dc:"app跳转链接"`
	UrlParam    string `json:"urlParam" dc:"app跳转参数"`
	LinkText    string `json:"linkText"  dc:"链接显示的文字"`
}

// HomepageArticlesThumbInp 首页活动-点赞
type HomepageArticlesThumbInp struct {
	ArticleId int `json:"articleId" v:"required#id_unknown" dc:"活动ID"`
	Type      int `json:"type" v:"required" d:"1" dc:"点赞类型 1-点赞，2-取消点赞"`
}

func (in *HomepageArticlesThumbInp) Filter(ctx context.Context) (err error) {
	return
}

type HomepageArticlesThumbModel struct{}
