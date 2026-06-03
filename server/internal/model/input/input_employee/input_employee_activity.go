// Package sysin

package input_employee

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityUpdateFields 修改员工活动表字段过滤
type EmployeeActivityUpdateFields struct {
	Name            string      `json:"name"            dc:"活动名称"`
	Cover           string      `json:"cover"           dc:"活动封面"`
	Description     string      `json:"description"     dc:"活动描述"`
	ValidityType    int         `json:"validityType"    dc:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime       *gtime.Time `json:"startTime"       dc:"开始时间"`
	EndTime         *gtime.Time `json:"endTime"         dc:"结束时间"`
	RestrictionType int         `json:"restrictionType" dc:"限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工"`
	IsEnabled       int         `json:"isEnabled"       dc:"状态：1-正常 2-禁用"`
	Rule            string      `json:"rule"            dc:"活动规则"`
	CouponValidity  int         `json:"couponValidity"  dc:"券有效期  1-跟随活动  2-自身有效期"`
	LimitWeek       string      `json:"limitWeek"       dc:"周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔"`
}

// EmployeeActivityInsertFields 新增员工活动表字段过滤
type EmployeeActivityInsertFields struct {
	Name            string      `json:"name"            dc:"活动名称"`
	Cover           string      `json:"cover"           dc:"活动封面"`
	Description     string      `json:"description"     dc:"活动描述"`
	ValidityType    int         `json:"validityType"    dc:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime       *gtime.Time `json:"startTime"       dc:"开始时间"`
	EndTime         *gtime.Time `json:"endTime"         dc:"结束时间"`
	RestrictionType int         `json:"restrictionType" dc:"限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工"`
	IsEnabled       int         `json:"isEnabled"       dc:"状态：1-正常 2-禁用"`
	Status          int         `json:"status"          dc:"状态：1-未开始 2-进行中 3-已结束"`
	Rule            string      `json:"rule"            dc:"活动规则"`
	CouponValidity  int         `json:"couponValidity"  dc:"券有效期  1-跟随活动  2-自身有效期"`
	LimitWeek       string      `json:"limitWeek"       dc:"周末领取限制 6-周六不允许,7-周日不允许，多个用逗号分隔"`
}

// EmployeeActivityEditInp 修改/新增员工活动表
type EmployeeActivityEditInp struct {
	entity.EmployeeActivity
	NameLanguage        input_language.LanguageModel `json:"nameLanguage"          dc:"多语言名称"`
	DescriptionLanguage input_language.LanguageModel `json:"descriptionLanguage"   dc:"多语言描述"`
	ThCouponsArr        []*struct {
		CouponId          int `json:"couponId"          dc:"提货券ID"`
		AvailableQuantity int `json:"availableQuantity" dc:"领取数量"`
		PerDayAvailable   int `json:"perDayAvailable"   dc:"每人每天可领取数量（0表示无限制）"`
		PerDayVerify      int `json:"perDayVerify"      dc:"每人每天可核销数量（0表示无限制）"`
		LimitDays         int `json:"limitDays"         dc:"限制天数"`
	} `json:"thCouponsArr"          dc:"提货券列表"`
	DepartmentIds []int `json:"departmentIds" dc:"员工部门IDs"`
	EmployeeIds   []int `json:"employeeIds" dc:"员工IDs"`
}

func (in *EmployeeActivityEditInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityEditModel struct{}

// EmployeeActivityDeleteInp 删除员工活动表
type EmployeeActivityDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *EmployeeActivityDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityDeleteModel struct{}

// EmployeeActivityViewInp 获取指定员工活动表信息
type EmployeeActivityViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *EmployeeActivityViewInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityViewModel struct {
	entity.EmployeeActivity
	NameLanguage        []*input_hotel.LanguageType `json:"nameLanguage"         dc:"名称"   orm:"with:uuid=name"`
	DescriptionLanguage []*input_hotel.LanguageType `json:"descriptionLanguage"         dc:"描述"   orm:"with:uuid=description"`
	CouponList          []*struct {
		gmeta.Meta        `orm:"table:hg_employee_activity_coupon"`
		CouponId          int64 `json:"couponId"  dc:"提货券ID"`
		ActivityId        int   `json:"activityId"        dc:"活动ID"`
		AvailableQuantity int   `json:"availableQuantity" dc:"可领取数量（>0）"`
		PerDayAvailable   int   `json:"perDayAvailable"   dc:"每人每天可领取数量（0表示无限制）"`
		PerDayVerify      int   `json:"perDayVerify"   dc:"每人每天可核销数量（0表示无限制）"`
	} `json:"couponList" orm:"with:activity_id=id" dc:"提货券列表"`
	EmployeeDepartmentList []*struct {
		gmeta.Meta       `orm:"table:hg_employee_activity_department"`
		DepartmentId     int64 `json:"departmentId"  dc:"部门ID"`
		ActivityId       int   `json:"activityId"        dc:"活动ID"`
		DepartmentDetail *struct {
			gmeta.Meta `orm:"table:hg_employee_department"`
			Id         int    `json:"id"              dc:"ID"`
			Name       string `json:"name"            dc:"名称"`
		} `json:"departmentDetail" orm:"with:id=department_id"`
	} `json:"employeeDepartmentList" orm:"with:activity_id=id" dc:"员工部门列表"`
	EmployeeList []*struct {
		gmeta.Meta `orm:"table:hg_employee_activity_employee"`
		EmployeeId int64 `json:"employeeId"  dc:"员工ID"`
		ActivityId int   `json:"activityId"        dc:"活动ID"`
	} `json:"employeeList" orm:"with:activity_id=id" dc:"员工列表"`
}

// EmployeeActivityListInp 获取员工活动表列表
type EmployeeActivityListInp struct {
	input_form.PageReq
	Name   string `json:"name" dc:"活动名称"`
	Status int    `json:"status" dc:"状态"`
}

func (in *EmployeeActivityListInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityListModel struct {
	Id           int64       `json:"id"       dc:"id"`
	Name         string      `json:"name"            dc:"活动名称"`
	Cover        string      `json:"cover"           dc:"活动封面"`
	Description  string      `json:"description"     dc:"活动描述"`
	ValidityType int         `json:"validityType"    dc:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime    *gtime.Time `json:"startTime"       dc:"开始时间"`
	EndTime      *gtime.Time `json:"endTime"         dc:"结束时间"`
	StartDate    string      `json:"startDate"       dc:"开始日期"`
	EndDate      string      `json:"endDate"         dc:"结束日期"`
	Status       int         `json:"status"          dc:"状态：1-未开始 2-进行中 3-已结束"`
	IsEnabled    int         `json:"isEnabled"       dc:"状态：1-正常 2-禁用"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CouponList   []*struct {
		gmeta.Meta   `orm:"table:hg_employee_activity_coupon"`
		CouponId     int64 `json:"couponId"  orm:"coupon_id"     description:"提货券ID"`
		ActivityId   int   `json:"activityId"        orm:"activity_id"           description:"活动ID"`
		CouponDetail *struct {
			gmeta.Meta `orm:"table:hg_th_coupon"`
			Id         int    `json:"id"              dc:"ID"`
			Name       string `json:"name"            dc:"名称"`
		} `json:"couponDetail" orm:"with:id=coupon_id"`
	} `json:"couponList" orm:"with:activity_id=id" dc:"提货券列表"`
}

// EmployeeActivityStatusInp 更新员工活动表状态
type EmployeeActivityStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *EmployeeActivityStatusInp) Filter(ctx context.Context) (err error) {
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

type EmployeeActivityStatusModel struct{}

// EmployeeActivityAppListInp 获取员工活动表列表
type EmployeeActivityAppListInp struct {
	input_form.PageReq
	Status int `json:"status" dc:"状态 1-未开始 2-进行中 3-已结束"`
}

func (in *EmployeeActivityAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityAppListModel struct {
	Id           int64       `json:"id"       dc:"id"`
	Name         string      `json:"name"            dc:"活动名称"`
	Cover        string      `json:"cover"           dc:"活动封面"`
	Description  string      `json:"description"     dc:"活动描述"`
	ValidityType int         `json:"validityType"    dc:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime    *gtime.Time `json:"startTime"       dc:"开始时间"`
	EndTime      *gtime.Time `json:"endTime"         dc:"结束时间"`
	Status       int         `json:"status"          dc:"状态：1-未开始 2-进行中 3-已结束"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// EmployeeActivityAppViewInp 获取指定员工活动表信息
type EmployeeActivityAppViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *EmployeeActivityAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type EmployeeActivityAppViewModel struct {
	Id           int64       `json:"id"       dc:"id"`
	Name         string      `json:"name"            dc:"活动名称"`
	Cover        string      `json:"cover"           dc:"活动封面"`
	Description  string      `json:"description"     dc:"活动描述"`
	ValidityType int         `json:"validityType"    dc:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime    *gtime.Time `json:"startTime"       dc:"开始时间"`
	EndTime      *gtime.Time `json:"endTime"         dc:"结束时间"`
	Status       int         `json:"status"          dc:"状态：1-未开始 2-进行中 3-已结束"`
	Rule         string      `json:"rule"           dc:"活动规则"`
	CouponList   []*struct {
		gmeta.Meta         `orm:"table:hg_employee_activity_coupon"`
		CouponId           int64 `json:"couponId"  dc:"提货券ID"`
		ActivityId         int   `json:"activityId"        dc:"活动ID"`
		HaveReceived       int   `json:"haveReceived"        dc:"本人已领取数量"`
		MemberTodayReceive int   `json:"memberTodayReceive" dc:"会员今日已领取数量"`
		AvailableQuantity  int   `json:"availableQuantity" dc:"每人可领取总数量（0表示无限制）"`
		PerDayAvailable    int   `json:"perDayAvailable" dc:"每人每天可领取总数量（0是不限制）"`
		ReceiveStatus      int   `json:"receiveStatus"        dc:"状态 1-可领取 2-已领取 3-已领完 4-今日已领"`
		CouponDetail       *struct {
			gmeta.Meta    `orm:"table:hg_th_coupon"`
			Id            int    `json:"id"              dc:"ID"`
			CouponName    string `json:"couponName"     dc:"券名称"`
			CouponSubName string `json:"couponSubName"  dc:"券副标题"`
			Desc          string `json:"desc"           dc:"券说明"`
			Logo          string `json:"logo"          dc:"LOGO"`
		} `json:"couponDetail" orm:"with:id=coupon_id" dc:"提货券详情"`
	} `json:"couponList" orm:"with:activity_id=id" dc:"提货券列表"`
}

// EmployeeActivityReceiveCouponInp 领取礼品券
type EmployeeActivityReceiveCouponInp struct {
	ActivityId int `json:"activityId"          dc:"活动ID"`
	CouponId   int `json:"couponId"          dc:"礼品券ID"`
}

type EmployeeActivityReceiveCouponModel struct{}

// EmployeeActivityBatchIssueCouponsInp 批量发放预约券
type EmployeeActivityBatchIssueCouponsInp struct {
	ActivityId int `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
}

func (in *EmployeeActivityBatchIssueCouponsInp) Filter(ctx context.Context) (err error) {
	if in.ActivityId <= 0 {
		err = gerror.New("活动ID不能为空")
	}
	return
}

// EmployeeActivityBatchIssueCouponsModel 批量发放结果
type EmployeeActivityBatchIssueCouponsModel struct {
	TotalEmployees  int                                         `json:"totalEmployees" dc:"总员工数"`
	IssuedCount     int                                         `json:"issuedCount"    dc:"成功发放数量"`
	FailedCount     int                                         `json:"failedCount"    dc:"失败数量"`
	CouponDetails   []*EmployeeActivityBatchIssueCouponDetail   `json:"couponDetails"  dc:"券发放详情"`
	FailedEmployees []*EmployeeActivityBatchIssueFailedEmployee `json:"failedEmployees" dc:"失败员工列表"`
}

type EmployeeActivityBatchIssueCouponDetail struct {
	CouponId    int    `json:"couponId"    dc:"券ID"`
	CouponName  string `json:"couponName"  dc:"券名称"`
	IssuedCount int    `json:"issuedCount" dc:"发放数量"`
}

type EmployeeActivityBatchIssueFailedEmployee struct {
	EmployeeId   int    `json:"employeeId"   dc:"员工ID"`
	EmployeeName string `json:"employeeName" dc:"员工姓名"`
	Reason       string `json:"reason"       dc:"失败原因"`
}

// EmployeeActivityCouponRecordListInp 券领取记录列表
type EmployeeActivityCouponRecordListInp struct {
	input_form.PageReq
	ActivityId int    `json:"activityId" v:"required#活动ID不能为空" dc:"活动ID"`
	IssueType  int    `json:"issueType"  dc:"发放类型：0-全部 1-自动领取 2-批量发放"`
	State      int    `json:"state"      dc:"券状态：0-全部 1-待生效 2-未使用 3-已核销 4-已过期 5-已失效 6-已回收"`
	Keyword    string `json:"keyword"    dc:"搜索关键词（员工姓名/券号）"`
}

func (in *EmployeeActivityCouponRecordListInp) Filter(ctx context.Context) (err error) {
	if in.ActivityId <= 0 {
		err = gerror.New("活动ID不能为空")
	}
	return
}

type EmployeeActivityCouponRecordListModel struct {
	Id              int         `json:"id"              dc:"记录ID"`
	CouponNo        string      `json:"couponNo"        dc:"券号"`
	CouponId        int         `json:"couponId"        dc:"券ID"`
	CouponName      string      `json:"couponName"      dc:"券名称"`
	EmployeeId      int         `json:"employeeId"      dc:"员工ID"`
	EmployeeName    string      `json:"employeeName"    dc:"员工姓名"`
	MemberId        int         `json:"memberId"        dc:"会员ID"`
	State           int         `json:"state"           dc:"状态 1待生效 2未使用 3已核销 4已过期 5已失效 6已回收"`
	Source          int         `json:"source"          dc:"来源"`
	NeedReservation int         `json:"needReservation" dc:"是否需要预约：0-不需要 1-需要"`
	IssueType       int         `json:"issueType"       dc:"发放类型：1-自动领取 2-批量发放"`
	StartTime       *gtime.Time `json:"startTime"       dc:"有效期开始时间"`
	EndTime         *gtime.Time `json:"endTime"         dc:"有效期结束时间"`
	VerifyTime      *gtime.Time `json:"verifyTime"      dc:"核销时间"`
	VerifyMchId     int         `json:"verifyMchId"     dc:"核销商户ID"`
	StoreName       string      `json:"storeName"       dc:"核销门店名称"`
	CreateAt        *gtime.Time `json:"createAt"        dc:"领取时间"`
}
