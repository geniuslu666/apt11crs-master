package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CarSettlementOrderInsertFields 新增接送机结算单字段过滤
type CarSettlementOrderInsertFields struct {
	OrderSn          string      `json:"orderSn"          dc:"结算单编号"`
	DriverId         int         `json:"driverId"         dc:"司机ID"`
	OrderAmount      float64     `json:"orderAmount"      dc:"订单总额"`
	SettlementAmount float64     `json:"settlementAmount" dc:"结算总额"`
	Status           string      `json:"status"           dc:"结算状态"`
	StartTime        *gtime.Time `json:"startTime"        dc:"账期开始时间"`
	EndTime          *gtime.Time `json:"endTime"          dc:"账期结束时间"`
}

// CarSettlementOrderStatInp 获取财务概况信息
type CarSettlementOrderStatInp struct {
}

type CarSettlementOrderStatModel struct {
	TotalOrderAmount          float64 `json:"totalOrderAmount"            dc:"订单总额(已支付)"`
	WaitSettlementAmount      float64 `json:"waitSettlementAmount"        dc:"待结算(已支付已核销但还未结算的金额)"`
	DoneSettlementAmount      float64 `json:"doneSettlementAmount"        dc:"已结算金额"`
	HadVerifySettlementAmount float64 `json:"hadVerifySettlementAmount"   dc:"已核账(已核销且完成手动对账的金额)"`
}

// CarSettlementOrderListInp 获取结算单列表
type CarSettlementOrderListInp struct {
	input_form.PageReq
	OrderSn      string        `json:"orderSn"      dc:"结算单编号"`
	DriverId     int           `json:"driverId"     dc:"司机ID"`
	VerifyStatus string        `json:"verifyStatus" dc:"核账状态"`
	VerifyTime   []*gtime.Time `json:"verifyTime"   dc:"核账时间"`
}

func (in *CarSettlementOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementOrderListModel struct {
	Id               int64       `json:"id"               orm:"id"                description:""`
	OrderSn          string      `json:"orderSn"          orm:"order_sn"          description:"结算单编号"`
	DriverId         int         `json:"driverId"         orm:"driver_id"         description:"司机ID"`
	OrderAmount      float64     `json:"orderAmount"      orm:"order_amount"      description:"订单总额"`
	SettlementAmount float64     `json:"settlementAmount" orm:"settlement_amount" description:"结算总额"`
	Status           string      `json:"status"           orm:"status"            description:"结算状态"`
	StartTime        *gtime.Time `json:"startTime"        orm:"start_time"        description:"账期开始时间"`
	EndTime          *gtime.Time `json:"endTime"          orm:"end_time"          description:"账期结束时间"`
	VerifyStatus     string      `json:"verifyStatus"     orm:"verify_status"     description:"核账状态"`
	VerifyTime       *gtime.Time `json:"verifyTime"       orm:"verify_time"       description:"核帐时间"`
	VerifyImg        string      `json:"verifyImg"        orm:"verify_img"        description:"核账凭证"`
	VerifyDesc       string      `json:"verifyDesc"       orm:"verify_desc"       description:"核账说明"`
	VerifyOperateId  string      `json:"verifyOperateId"  orm:"verify_operate_id" description:"核账操作人ID"`
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"         description:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"         description:"更新时间"`
	DriverDetail     *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
	OperateDetail *struct {
		gmeta.Meta `orm:"table:hg_admin_member"`
		Id         int64  `json:"id"           dc:"管理员ID"`
		DeptId     int64  `json:"deptId"       dc:"部门ID"`
		RealName   string `json:"realName"     dc:"真实姓名"`
		Username   string `json:"username"     dc:"帐号"`
	} `json:"operateDetail" orm:"with:id=verify_operate_id"`
}

// CarSettlementOrderViewInp 结算单详情
type CarSettlementOrderViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarSettlementOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementOrderViewModel struct {
	entity.CarSettlementOrder
	DriverDetail *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
}

// CarSettlementOrderVerifyFields 结算单核账字段过滤
type CarSettlementOrderVerifyFields struct {
	VerifyStatus    string      `json:"verifyStatus"   dc:"核账状态"`
	VerifyImg       string      `json:"verifyImg"      dc:"核账凭证"`
	VerifyDesc      string      `json:"verifyDesc"     dc:"核账说明"`
	VerifyOperateId int64       `json:"verifyOperateId"  dc:"操作人id"`
	VerifyTime      *gtime.Time `json:"verifyTime"     dc:"核帐时间"`
}

// CarSettlementOrderVerifyInp 结算单-核账
type CarSettlementOrderVerifyInp struct {
	Id         int64  `json:"id" v:"required#id不能为空" dc:"id"`
	VerifyImg  string `json:"verifyImg"           dc:"核账凭证"`
	VerifyDesc string `json:"verifyDesc"          dc:"核账说明"`
}

func (in *CarSettlementOrderVerifyInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementOrderVerifyModel struct{}
