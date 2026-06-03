// Package sysin

package input_food

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// FoodRestaurantUpdateFields 修改餐厅管理字段过滤
type FoodRestaurantUpdateFields struct {
	Name                    string  `json:"name"                    dc:"名称"`
	CuisineIds              string  `json:"cuisineIds"              dc:"菜系（多选）"`
	LabelIds                string  `json:"labelIds"                dc:"标签（多选）"`
	CooperateTypeId         int     `json:"cooperateTypeId"         dc:"合作类型ID"`
	Logo                    string  `json:"logo"                    dc:"LOGO"`
	Images                  string  `json:"images"                  dc:"图集"`
	Content                 string  `json:"content"                 dc:"简介"`
	Phone                   string  `json:"phone"                   dc:"联系电话"`
	OpenTime                string  `json:"openTime"                dc:"营业时间"`
	AreaPid                 int     `json:"areaPid"                 dc:"地区省级ID"`
	AreaId                  int     `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress           string  `json:"detailAddress"                 dc:"详细地址"`
	GgLat                   string  `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng                   string  `json:"ggLng"                   dc:"谷歌经度"`
	Lat                     string  `json:"lat"                     dc:"纬度"`
	Lng                     string  `json:"lng"                     dc:"经度"`
	OpenStatus              string  `json:"openStatus"              dc:"营业状态"`
	OrderTimeType           int     `json:"orderTimeType"           dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek           string  `json:"orderTimeWeek"           dc:"预约日期自定义周数"`
	RestTimeWeek            string  `json:"restTimeWeek"            dc:"定休日周数"`
	RestTimeDate            string  `json:"restTimeDate"            dc:"定休日固定日期"`
	OrderConfirmType        int     `json:"orderConfirmType"        dc:"预约确定模式 1手动确认 2自动确认"`
	OrderConfirmBeforeDays  int     `json:"orderConfirmBeforeDays"  dc:"多少天内手动确认"`
	OrderConfirmTime        int     `json:"orderConfirmTime"        dc:"预约确认自定义分钟数"`
	OpenTimeType            int     `json:"openTimeType"            dc:"1 按时段  2 按时间点"`
	AmOpenTime              string  `json:"amOpenTime"              dc:"早市开始时间"`
	AmCloseTime             string  `json:"amCloseTime"             dc:"早市结束时间"`
	PmOpenTime              string  `json:"pmOpenTime"              dc:"晚市开始时间"`
	PmCloseTime             string  `json:"pmCloseTime"             dc:"晚市结束时间"`
	TimeDuration            string  `json:"timeDuration"            dc:"时间间隔"`
	TimePoints              string  `json:"timePoints"              dc:"时间点"`
	OrderMaxType            int     `json:"orderMaxType"            dc:"最大容纳数  1按时段  2按时间点"`
	TimeDurationMax         int     `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	AdvanceOrderDay         int     `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	MaxOrderDay             int     `json:"maxOrderDay"             dc:"最长预约天数"`
	CancelPolicyOpen        int     `json:"cancelPolicyOpen"        dc:"是否允许取消  1允许  2不允许"`
	BeforeConfirmCancelRate float64 `json:"beforeConfirmCancelRate" dc:"订单确认前取消费用为订单的%，0则免费"`
	AfterConfirmCancelDay   int     `json:"afterConfirmCancelDay"   dc:"订单确认后距离到店多少天"`
	AfterConfirmCancelRate1 float64 `json:"afterConfirmCancelRate1" dc:"订单确认后距离到店天数前取消费用为订单的%，0则免费"`
	AfterConfirmCancelRate2 float64 `json:"afterConfirmCancelRate2" dc:"订单确认后距离到店天数后取消费用为订单的%，0则免费"`
	MaxSeat                 int     `json:"maxSeat"                 dc:"最大席位数"`
	CanSmoking              int     `json:"canSmoking"              dc:"是否允许抽烟  1允许  2不允许"`
	SettlementType          int     `json:"settlementType"          dc:"门店抽成类型 1跟随系统  2自定义"`
	SettlementId            int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementRate          float64 `json:"settlementRate"          dc:"门店结算比例"`
	DepositRate             float64 `json:"depositRate"             dc:"定金比例"`
	Account                 string  `json:"account"       dc:"账号"`
	Password                string  `json:"password"     dc:"密码"`
}

// FoodRestaurantBasicUpdateFields 修改餐厅基础信息字段过滤
type FoodRestaurantBasicUpdateFields struct {
	ToretaId        string `json:"toretaId"                dc:"Toreta的餐厅ID"`
	Name            string `json:"name"                    dc:"名称"`
	CuisineIds      string `json:"cuisineIds"              dc:"菜系（多选）"`
	LabelIds        string `json:"labelIds"                dc:"标签（多选）"`
	CooperateTypeId int    `json:"cooperateTypeId"         dc:"合作类型ID"`
	Logo            string `json:"logo"                    dc:"LOGO"`
	Images          string `json:"images"                  dc:"图集"`
	Content         string `json:"content"                 dc:"简介"`
	Phone           string `json:"phone"                   dc:"联系电话"`
	OpenTime        string `json:"openTime"                dc:"营业时间"`
	AreaPid         int    `json:"areaPid"                 dc:"地区省级ID"`
	AreaId          int    `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress   string `json:"detailAddress"                 dc:"详细地址"`
	GgLat           string `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng           string `json:"ggLng"                   dc:"谷歌经度"`
	Lat             string `json:"lat"                     dc:"纬度"`
	Lng             string `json:"lng"                     dc:"经度"`
}

// FoodRestaurantBasicUpdateToretaFields 修改餐厅基础信息字段过滤-合作类型变更为toreta，且toretaid与之前不一致
type FoodRestaurantBasicUpdateToretaFields struct {
	ToretaId              string `json:"toretaId"                dc:"Toreta的餐厅ID"`
	Name                  string `json:"name"                    dc:"名称"`
	CuisineIds            string `json:"cuisineIds"              dc:"菜系（多选）"`
	LabelIds              string `json:"labelIds"                dc:"标签（多选）"`
	CooperateTypeId       int    `json:"cooperateTypeId"         dc:"合作类型ID"`
	Logo                  string `json:"logo"                    dc:"LOGO"`
	Images                string `json:"images"                  dc:"图集"`
	Content               string `json:"content"                 dc:"简介"`
	Phone                 string `json:"phone"                   dc:"联系电话"`
	OpenTime              string `json:"openTime"                dc:"营业时间"`
	AreaPid               int    `json:"areaPid"                 dc:"地区省级ID"`
	AreaId                int    `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress         string `json:"detailAddress"                 dc:"详细地址"`
	GgLat                 string `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng                 string `json:"ggLng"                   dc:"谷歌经度"`
	Lat                   string `json:"lat"                     dc:"纬度"`
	Lng                   string `json:"lng"                     dc:"经度"`
	OrderMode             string `json:"orderMode"               dc:"预定模式 ALLAMOUNT-全款模式 DEPOSIT-定金模式"`
	OrderConfirmType      int    `json:"orderConfirmType"        dc:"预约确定模式 1手动确认 2自动确认"`
	OpenTimeType          int    `json:"openTimeType"            dc:"1 按时段  2 按时间点"`
	OrderMaxType          int    `json:"orderMaxType"            dc:"最大容纳数  1按时段  2按时间点"`
	TimeDurationMax       int    `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	DayTimeLimit          string `json:"dayTimeLimit"            dc:"当天预约处理截止时间"`
	AdvanceOrderDay       int    `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	MaxOrderDay           int    `json:"maxOrderDay"             dc:"最长预约天数"`
	ToretaCancelEnable    int    `json:"toretaCancelEnable"      dc:"Toreta是否允许取消 1 允许  2 不允许"`
	ToretaCancelLimitDay  int    `json:"toretaCancelLimitDay"    dc:"Toreta允许取消几天前"`
	ToretaCancelLimitTime string `json:"toretaCancelLimitTime"   dc:"Toreta允许取消时间前"`
}

// FoodRestaurantOperationUpdateFields 修改餐厅运营信息字段过滤
type FoodRestaurantOperationUpdateFields struct {
	OrderTimeType           int     `json:"orderTimeType"           dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek           string  `json:"orderTimeWeek"           dc:"预约日期自定义周数"`
	RestTimeWeek            string  `json:"restTimeWeek"            dc:"定休日周数"`
	RestTimeDate            string  `json:"restTimeDate"            dc:"定休日固定日期"`
	OrderConfirmType        int     `json:"orderConfirmType"        dc:"预约确定模式 1手动确认 2自动确认"`
	OrderConfirmBeforeDays  int     `json:"orderConfirmBeforeDays"  dc:"多少天内手动确认"`
	OrderConfirmTime        int     `json:"orderConfirmTime"        dc:"预约确认自定义分钟数"`
	OpenTimeType            int     `json:"openTimeType"            dc:"1 按时段  2 按时间点"`
	AmOpenTime              string  `json:"amOpenTime"              dc:"早市开始时间"`
	AmCloseTime             string  `json:"amCloseTime"             dc:"早市结束时间"`
	PmOpenTime              string  `json:"pmOpenTime"              dc:"晚市开始时间"`
	PmCloseTime             string  `json:"pmCloseTime"             dc:"晚市结束时间"`
	TimeDuration            string  `json:"timeDuration"            dc:"时间间隔"`
	TimePoints              string  `json:"timePoints"              dc:"时间点"`
	OrderMaxType            int     `json:"orderMaxType"            dc:"最大容纳数  1按时段  2按时间点"`
	TimeDurationMax         int     `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	DayTimeLimit            string  `json:"dayTimeLimit"            dc:"当天预约处理截止时间"`
	AdvanceOrderDay         int     `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	MaxOrderDay             int     `json:"maxOrderDay"             dc:"最长预约天数"`
	CancelPolicyOpen        int     `json:"cancelPolicyOpen"        dc:"是否允许取消  1允许  2不允许"`
	BeforeConfirmCancelRate float64 `json:"beforeConfirmCancelRate" dc:"订单确认前取消费用为订单的%，0则免费"`
	AfterConfirmCancelDay   int     `json:"afterConfirmCancelDay"   dc:"订单确认后距离到店多少天"`
	AfterConfirmCancelRate1 float64 `json:"afterConfirmCancelRate1" dc:"订单确认后距离到店天数前取消费用为订单的%，0则免费"`
	AfterConfirmCancelRate2 float64 `json:"afterConfirmCancelRate2" dc:"订单确认后距离到店天数后取消费用为订单的%，0则免费"`
	DepositRate             float64 `json:"depositRate"             dc:"定金比例"`
	OrderMode               string  `json:"orderMode"               dc:"预定模式 ALLAMOUNT-全款模式 DEPOSIT-定金模式"`
	MaxDatetimeOrderOpen    uint    `json:"maxDatetimeOrderOpen"    dc:"日期限制 1 开启  2关闭"`
	MaxDatetimeOrderDate    string  `json:"maxDatetimeOrderDate"    dc:"限制的日期 逗号分隔"`
}

// FoodRestaurantToretaOperationUpdateFields 修改餐厅Toreta运营信息字段过滤
type FoodRestaurantToretaOperationUpdateFields struct {
	OrderTimeType         int    `json:"orderTimeType"           dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek         string `json:"orderTimeWeek"           dc:"预约日期自定义周数"`
	TimeDurationMax       int    `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	DayTimeLimit          string `json:"dayTimeLimit"            dc:"当天预约处理截止时间"`
	AdvanceOrderDay       int    `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	MaxOrderDay           int    `json:"maxOrderDay"             dc:"最长预约天数"`
	ToretaCancelEnable    int    `json:"toretaCancelEnable"      dc:"Toreta是否允许取消 1 允许  2 不允许"`
	ToretaCancelLimitDay  int    `json:"toretaCancelLimitDay"    dc:"Toreta允许取消几天前"`
	ToretaCancelLimitTime string `json:"toretaCancelLimitTime"   dc:"Toreta允许取消时间前"`
	OrderMode             string `json:"orderMode"               dc:"预定模式 ALLAMOUNT-全款模式 DEPOSIT-定金模式"`
	MaxDatetimeOrderOpen  uint   `json:"maxDatetimeOrderOpen"    dc:"日期限制 1 开启  2关闭"`
	MaxDatetimeOrderDate  string `json:"maxDatetimeOrderDate"    dc:"限制的日期 逗号分隔"`
}

// FoodRestaurantOtherUpdateFields 修改餐厅其他信息字段过滤
type FoodRestaurantOtherUpdateFields struct {
	MaxSeat    int `json:"maxSeat"                 dc:"最大席位数"`
	CanSmoking int `json:"canSmoking"              dc:"是否允许抽烟  1允许  2不允许"`
}

// FoodRestaurantSettleUpdateFields 修改餐厅结算信息字段过滤
type FoodRestaurantSettleUpdateFields struct {
	SettlementId   int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType int     `json:"settlementType"          dc:"门店抽成类型 1跟随系统  2自定义"`
	SettlementRate float64 `json:"settlementRate"          dc:"门店结算比例"`
}

// FoodRestaurantTerminalUpdateFields 修改餐厅终端信息字段过滤
type FoodRestaurantTerminalUpdateFields struct {
	Account      string `json:"account"       dc:"账号"`
	Password     string `json:"password"     dc:"密码"`
	PasswordHash string `json:"passwordHash"            dc:"密码"`
	Salt         string `json:"salt"                    dc:"密码盐"`
}

// FoodRestaurantInsertFields 新增餐厅管理字段过滤
type FoodRestaurantInsertFields struct {
	ToretaId                string      `json:"toretaId"                dc:"Toreta的餐厅ID"`
	OrderMode               string      `json:"orderMode"               dc:"预定模式 ALLAMOUNT-全款模式 DEPOSIT-定金模式"`
	Name                    string      `json:"name"                    dc:"名称"`
	CuisineIds              string      `json:"cuisineIds"              dc:"菜系（多选）"`
	LabelIds                string      `json:"labelIds"                dc:"标签（多选）"`
	CooperateTypeId         int         `json:"cooperateTypeId"         dc:"合作类型ID"`
	Logo                    string      `json:"logo"                    dc:"LOGO"`
	Images                  string      `json:"images"                  dc:"图集"`
	Content                 string      `json:"content"                 dc:"简介"`
	Phone                   string      `json:"phone"                   dc:"联系电话"`
	OpenTime                string      `json:"openTime"                dc:"营业时间"`
	AreaPid                 int         `json:"areaPid"                 dc:"地区省级ID"`
	AreaId                  int         `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress           string      `json:"detailAddress"                 dc:"详细地址"`
	GgLat                   string      `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng                   string      `json:"ggLng"                   dc:"谷歌经度"`
	Lat                     string      `json:"lat"                     dc:"纬度"`
	Lng                     string      `json:"lng"                     dc:"经度"`
	OpenStatus              string      `json:"openStatus"              dc:"营业状态"`
	OrderTimeType           int         `json:"orderTimeType"           dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek           string      `json:"orderTimeWeek"           dc:"预约日期自定义周数"`
	RestTimeWeek            string      `json:"restTimeWeek"            dc:"定休日周数"`
	RestTimeDate            string      `json:"restTimeDate"            dc:"定休日固定日期"`
	OrderConfirmType        int         `json:"orderConfirmType"        dc:"预约确定模式 1手动确认 2自动确认"`
	OrderConfirmTime        int         `json:"orderConfirmTime"        dc:"预约确认自定义分钟数"`
	OpenTimeType            int         `json:"openTimeType"            dc:"1 按时段  2 按时间点"`
	AmOpenTime              *gtime.Time `json:"amOpenTime"              dc:"早市开始时间"`
	AmCloseTime             *gtime.Time `json:"amCloseTime"             dc:"早市结束时间"`
	PmOpenTime              *gtime.Time `json:"pmOpenTime"              dc:"晚市开始时间"`
	PmCloseTime             *gtime.Time `json:"pmCloseTime"             dc:"晚市结束时间"`
	TimeDuration            string      `json:"timeDuration"            dc:"时间间隔"`
	TimePoints              string      `json:"timePoints"              dc:"时间点"`
	OrderMaxType            int         `json:"orderMaxType"            dc:"最大容纳数  1按时段  2按时间点"`
	TimeDurationMax         int         `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	AdvanceOrderDay         int         `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	DayTimeLimit            string      `json:"dayTimeLimit"            dc:"当天预约处理截止时间"`
	MaxOrderDay             int         `json:"maxOrderDay"             dc:"最长预约天数"`
	CancelPolicyOpen        int         `json:"cancelPolicyOpen"        dc:"是否允许取消  1允许  2不允许"`
	BeforeConfirmCancelRate float64     `json:"beforeConfirmCancelRate" dc:"订单确认前取消费用为订单的%，0则免费"`
	AfterConfirmCancelDay   int         `json:"afterConfirmCancelDay"   dc:"订单确认后距离到店多少天"`
	AfterConfirmCancelRate1 float64     `json:"afterConfirmCancelRate1" dc:"订单确认后距离到店天数前取消费用为订单的%，0则免费"`
	AfterConfirmCancelRate2 float64     `json:"afterConfirmCancelRate2" dc:"订单确认后距离到店天数后取消费用为订单的%，0则免费"`
	MaxSeat                 int         `json:"maxSeat"                 dc:"最大席位数"`
	CanSmoking              int         `json:"canSmoking"              dc:"是否允许抽烟  1允许  2不允许"`
	SettlementType          int         `json:"settlementType"          dc:"门店抽成类型 1跟随系统  2自定义"`
	SettlementId            int         `json:"settlementId"            dc:"结算模式ID"`
	SettlementRate          float64     `json:"settlementRate"          dc:"门店结算比例"`
	VerifyCode              string      `json:"verifyCode"            dc:"核销码"`
	DepositRate             float64     `json:"depositRate"             dc:"定金比例"`
	Account                 string      `json:"account"       dc:"账号"`
	Password                string      `json:"password"     dc:"密码"`
	PasswordHash            string      `json:"passwordHash"            dc:"密码"`
	Salt                    string      `json:"salt"                    dc:"密码盐"`
	ToretaCancelEnable      int         `json:"toretaCancelEnable"      dc:"Toreta是否允许取消 1 允许  2 不允许"`
	ToretaCancelLimitDay    int         `json:"toretaCancelLimitDay"    dc:"Toreta允许取消几天前"`
	ToretaCancelLimitTime   string      `json:"toretaCancelLimitTime"   dc:"Toreta允许取消时间前"`
	MaxDatetimeOrderOpen    uint        `json:"maxDatetimeOrderOpen"    dc:"日期限制 1 开启  2关闭"`
	MaxDatetimeOrderDate    string      `json:"maxDatetimeOrderDate"    dc:"限制的日期 逗号分隔"`
}

// FoodRestaurantEditInp 修改/新增餐厅管理
type FoodRestaurantEditInp struct {
	entity.FoodRestaurant
	Password        string                       `json:"password"     dc:"密码"`
	NameLanguage    input_language.LanguageModel `json:"nameLanguage"          dc:"多语言餐厅名称"`
	ContentLanguage input_language.LanguageModel `json:"contentLanguage"          dc:"多语言简介名称"`
	Type            string                       `json:"type"     dc:"修改类型：basic、operation、other、settle"`
	Seat            []*struct {
		SeatId int `json:"seatId"     dc:"座位ID"`
		Num    int `json:"num"     dc:"同时段可预约数量"`
	} `json:"seatList"          dc:"座位"`
	TerminalList []*struct {
		Id         int `json:"id"     dc:"终端ID"`
		PrintTimes int `json:"printTimes"     dc:"打印次数"`
	} `json:"terminalList"          dc:"终端列表"`
}

func (in *FoodRestaurantEditInp) Filter(ctx context.Context) (err error) {
	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证菜系（多选）
	if err := g.Validator().Rules("required").Data(in.CuisineIds).Messages("菜系（多选）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证标签（多选）
	if err := g.Validator().Rules("required").Data(in.LabelIds).Messages("标签（多选）不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证合作类型ID
	if err := g.Validator().Rules("required").Data(in.CooperateTypeId).Messages("合作类型ID 不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证LOGO
	if err := g.Validator().Rules("required").Data(in.Logo).Messages("LOGO不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证联系电话
	if err := g.Validator().Rules("required").Data(in.Phone).Messages("联系电话不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证营业时间
	if err := g.Validator().Rules("required").Data(in.OpenTime).Messages("营业时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证详细地址
	if err := g.Validator().Rules("required").Data(in.DetailAddress).Messages("详细地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证谷歌纬度
	if err := g.Validator().Rules("required").Data(in.GgLat).Messages("谷歌纬度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证谷歌经度
	if err := g.Validator().Rules("required").Data(in.GgLng).Messages("谷歌经度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证纬度
	if err := g.Validator().Rules("required").Data(in.Lat).Messages("纬度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证经度
	if err := g.Validator().Rules("required").Data(in.Lng).Messages("经度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证营业状态
	if err := g.Validator().Rules("required").Data(in.OpenStatus).Messages("营业状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:OPENING,REST,CLOSE").Data(in.OpenStatus).Messages("营业状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证预约日期  1每天 2自定义
	if err := g.Validator().Rules("required").Data(in.OrderTimeType).Messages("预约日期  1每天 2自定义不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证预约日期自定义周数
	if err := g.Validator().Rules("required").Data(in.OrderTimeWeek).Messages("预约日期自定义周数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证定休日周数
	if err := g.Validator().Rules("required").Data(in.RestTimeWeek).Messages("定休日周数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证定休日固定日期
	if err := g.Validator().Rules("required").Data(in.RestTimeDate).Messages("定休日固定日期不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证预约确定模式 1手动确认 2自动确认
	if err := g.Validator().Rules("required").Data(in.OrderConfirmType).Messages("预约确定模式 1手动确认 2自动确认不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证预约确认自定义小时数
	if err := g.Validator().Rules("required").Data(in.OrderConfirmTime).Messages("预约确认自定义小时数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证时间间隔
	if err := g.Validator().Rules("required").Data(in.TimeDuration).Messages("时间间隔不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证时段最大容纳预定数量  0不限制
	if err := g.Validator().Rules("required").Data(in.TimeDurationMax).Messages("时段最大容纳预定数量  0不限制不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证至少提前预约天数  0当日可约
	if err := g.Validator().Rules("required").Data(in.AdvanceOrderDay).Messages("至少提前预约天数  0当日可约不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证最长预约天数
	if err := g.Validator().Rules("required").Data(in.MaxOrderDay).Messages("最长预约天数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否允许取消  1允许  2不允许
	if err := g.Validator().Rules("required").Data(in.CancelPolicyOpen).Messages("是否允许取消  1允许  2不允许不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证订单确认前取消费用为订单的%，0则免费
	if err := g.Validator().Rules("required").Data(in.BeforeConfirmCancelRate).Messages("订单确认前取消费用为订单的%，0则免费不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证订单确认后距离到店多少天
	if err := g.Validator().Rules("required").Data(in.AfterConfirmCancelDay).Messages("订单确认后距离到店多少天不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证订单确认后距离到店天数前取消费用为订单的%，0则免费
	if err := g.Validator().Rules("required").Data(in.AfterConfirmCancelRate1).Messages("订单确认后距离到店天数前取消费用为订单的%，0则免费不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证订单确认后距离到店天数后取消费用为订单的%，0则免费
	if err := g.Validator().Rules("required").Data(in.AfterConfirmCancelRate2).Messages("订单确认后距离到店天数后取消费用为订单的%，0则免费不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证最大席位数
	if err := g.Validator().Rules("required").Data(in.MaxSeat).Messages("最大席位数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证是否允许抽烟  1允许  2不允许
	if err := g.Validator().Rules("required").Data(in.CanSmoking).Messages("是否允许抽烟  1允许  2不允许不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证门店抽成类型 1跟随系统  2自定义
	if err := g.Validator().Rules("required").Data(in.SettlementType).Messages("门店抽成类型 1跟随系统  2自定义不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证结算模式ID
	if err := g.Validator().Rules("required").Data(in.SettlementId).Messages("结算模式ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证门店结算比例
	if err := g.Validator().Rules("required").Data(in.SettlementRate).Messages("门店结算比例不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodRestaurantEditModel struct{}

// FoodRestaurantDeleteInp 删除餐厅管理
type FoodRestaurantDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodRestaurantDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantDeleteModel struct{}

// FoodRestaurantViewInp 获取指定餐厅管理信息
type FoodRestaurantViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *FoodRestaurantViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantViewModel struct {
	entity.FoodRestaurant
	NameLanguage    []*input_hotel.LanguageType `json:"nameLanguage"         dc:"餐厅名称"   orm:"with:uuid=name"`
	ContentLanguage []*input_hotel.LanguageType `json:"contentLanguage"         dc:"简介"   orm:"with:uuid=content"`
	Seat            []*struct {
		gmeta.Meta   `orm:"table:hg_food_restaurant_seat"`
		RestaurantId int64 `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		SeatId       int   `json:"seatId"        orm:"seat_id"           description:"座位ID"`
		Num          int   `json:"num"        orm:"num"               description:"同时段可预约数量"`
		SeatDetail   *struct {
			gmeta.Meta `orm:"table:hg_food_seat"`
			*entity.FoodSeat
		} `json:"seatDetail" orm:"with:id=seat_id"  dc:"坐席信息"`
	} `json:"seat" orm:"with:restaurant_id=id" dc:"坐席列表"`
	TerminalList []*struct {
		gmeta.Meta `orm:"table:hg_food_restaurant_terminal"`
		*entity.FoodRestaurantTerminal
		TerminalInfo *struct {
			gmeta.Meta   `orm:"table:hg_sys_terminal"`
			Id           int    `json:"id"           dc:""`
			TerminalName string `json:"terminalName" dc:"终端名称"`
			TerminalType string `json:"terminalType" dc:"终端类型"`
			BrandModel   string `json:"brandModel"   dc:"品牌型号"`
			Sn           string `json:"sn"           dc:"终端编号"`
			PrintTimes   uint   `json:"printTimes" dc:"打印联数(次数)"`
		} `json:"terminalInfo" orm:"with:id=terminal_id" dc:"终端信息"`
	} `json:"terminalList" orm:"with:restaurant_id=id" dc:"终端列表"`
}

// FoodRestaurantListInp 获取餐厅管理列表
type FoodRestaurantListInp struct {
	input_form.PageReq
	Name          string `json:"name" dc:"名称"`
	OpenStatus    string `json:"openStatus"             dc:"营业状态"`
	CanOrder      int    `json:"canOrder"               dc:"是否开放预定  1开放  2关闭"`
	RestaurantIds string `json:"restaurantIds"   dc:"餐厅Ids"`
	Sort          string `json:"sort"      dc:"排序"`
}

func (in *FoodRestaurantListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantListModel struct {
	Id                      int64       `json:"id"                      dc:"id"`
	Name                    string      `json:"name"                    dc:"名称"`
	CuisineIds              string      `json:"cuisineIds"              dc:"菜系（多选）"`
	LabelIds                string      `json:"labelIds"                dc:"标签（多选）"`
	CooperateTypeId         int         `json:"cooperateTypeId"         dc:"合作类型ID"`
	Logo                    string      `json:"logo"                    dc:"LOGO"`
	Phone                   string      `json:"phone"                   dc:"联系电话"`
	OpenTime                string      `json:"openTime"                dc:"营业时间"`
	AreaPid                 int         `json:"areaPid"                 dc:"地区省级ID"`
	AreaId                  int         `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress           string      `json:"detailAddress"           dc:"详细地址"`
	GgLat                   string      `json:"ggLat"                   dc:"谷歌纬度"`
	GgLng                   string      `json:"ggLng"                   dc:"谷歌经度"`
	Lat                     string      `json:"lat"                     dc:"纬度"`
	Lng                     string      `json:"lng"                     dc:"经度"`
	OpenStatus              string      `json:"openStatus"              dc:"营业状态"`
	OrderTimeType           int         `json:"orderTimeType"           dc:"预约日期  1每天 2自定义"`
	OrderTimeWeek           string      `json:"orderTimeWeek"           dc:"预约日期自定义周数"`
	RestTimeWeek            string      `json:"restTimeWeek"            dc:"定休日周数"`
	OrderConfirmType        int         `json:"orderConfirmType"        dc:"预约确定模式 1手动确认 2自动确认"`
	OrderConfirmHour        int         `json:"orderConfirmHour"        dc:"预约确认自定义小时数"`
	AmOpenTime              *gtime.Time `json:"amOpenTime"              dc:"早市开始时间"`
	AmCloseTime             *gtime.Time `json:"amCloseTime"             dc:"早市结束时间"`
	PmOpenTime              *gtime.Time `json:"pmOpenTime"              dc:"晚市开始时间"`
	PmCloseTime             *gtime.Time `json:"pmCloseTime"             dc:"晚市结束时间"`
	TimeDuration            string      `json:"timeDuration"            dc:"时间间隔"`
	TimeDurationMax         int         `json:"timeDurationMax"         dc:"时段最大容纳预定数量  0不限制"`
	AdvanceOrderDay         int         `json:"advanceOrderDay"         dc:"至少提前预约天数  0当日可约"`
	MaxOrderDay             int         `json:"maxOrderDay"             dc:"最长预约天数"`
	CancelPolicyOpen        int         `json:"cancelPolicyOpen"        dc:"是否允许取消  1允许  2不允许"`
	BeforeConfirmCancelRate float64     `json:"beforeConfirmCancelRate" dc:"订单确认前取消费用为订单的%，0则免费"`
	AfterConfirmCancelDay   int         `json:"afterConfirmCancelDay"   dc:"订单确认后距离到店多少天"`
	AfterConfirmCancelRate1 float64     `json:"afterConfirmCancelRate1" dc:"订单确认后距离到店天数前取消费用为订单的%，0则免费"`
	AfterConfirmCancelRate2 float64     `json:"afterConfirmCancelRate2" dc:"订单确认后距离到店天数后取消费用为订单的%，0则免费"`
	MaxSeat                 int         `json:"maxSeat"                 dc:"最大席位数"`
	CanSmoking              int         `json:"canSmoking"              dc:"是否允许抽烟  1允许  2不允许"`
	SettlementType          int         `json:"settlementType"          dc:"门店抽成类型 1跟随系统  2自定义"`
	SettlementId            int         `json:"settlementId"            dc:"结算模式ID"`
	SettlementRate          float64     `json:"settlementRate"          dc:"门店结算比例"`
	TotalOrderNum           int         `json:"totalOrderNum"           dc:"预约单总数量（包含退款）"`
	TotalOrderAmount        float64     `json:"totalOrderAmount"        dc:"预约单总金额（包含退款）"`
	PayOrderNum             int         `json:"payOrderNum"             dc:"预约单支付数量（不包含退款）"`
	PayOrderAmount          float64     `json:"payOrderAmount"          dc:"预约单支付金额（不包含退款）"`
	WaitConfirmOrderNum     int         `json:"waitConfirmOrderNum"     dc:"预约单待确认数量（已支付待确认）"`
	SettlementOrderNum      int         `json:"settlementOrderNum"      dc:"预约单已结算数量"`
	SettlementOrderAmount   float64     `json:"settlementOrderAmount"   dc:"已结算预约单订单金额"`
	TotalSettlementAmount   float64     `json:"totalSettlementAmount"   dc:"已结算金额"`
	VerifyCode              string      `json:"verifyCode"            dc:"核销码"`
	CanOrder                int         `json:"canOrder"                dc:"是否开放预定  1开放  2关闭"`
	DepositRate             float64     `json:"depositRate"             dc:"定金比例"`
	Sort                    int         `json:"sort"                    dc:"排序 越大越靠前"`
	CreateAt                *gtime.Time `json:"createAt"                dc:"创建时间"`
	UpdateAt                *gtime.Time `json:"updateAt"                dc:"更新时间"`
	CuisineList             []*struct {
		gmeta.Meta    `orm:"table:hg_food_restaurant_cuisine"`
		RestaurantId  int64 `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		CuisineId     int   `json:"cuisineId"        orm:"cuisine_id"           description:"菜系ID"`
		CuisineDetail *struct {
			gmeta.Meta `orm:"table:hg_food_cuisine"`
			*entity.FoodCuisine
		} `json:"cuisineDetail" orm:"with:id=cuisine_id"  dc:"菜系信息"`
	} `json:"cuisineList" orm:"with:restaurant_id=id" dc:"菜系列表"`
	LabelList []*struct {
		gmeta.Meta   `orm:"table:hg_food_restaurant_label"`
		RestaurantId int64 `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		LabelId      int   `json:"labelId"        orm:"label_id"           description:"标签ID"`
		LabelDetail  *struct {
			gmeta.Meta `orm:"table:hg_food_label"`
			*entity.FoodLabel
		} `json:"labelDetail" orm:"with:id=label_id"  dc:"标签信息"`
	} `json:"labelList" orm:"with:restaurant_id=id" dc:"标签列表"`
	CooperateTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_food_cooperate_type"`
		*entity.FoodCooperateType
	} `json:"cooperateTypeDetail" orm:"with:id=cooperate_type_id"  dc:"合作类型"`
	PAreaDetail *struct {
		gmeta.Meta `orm:"table:hg_food_area"`
		*entity.FoodArea
	} `json:"pAreaDetail" orm:"with:id=area_pid"  dc:"地址信息"`
	AreaDetail *struct {
		gmeta.Meta `orm:"table:hg_food_area"`
		*entity.FoodArea
	} `json:"areaDetail" orm:"with:id=area_id"  dc:"地址信息"`
	SettlementDetail *struct {
		gmeta.Meta `orm:"table:hg_food_settlement"`
		*entity.FoodSettlement
	} `json:"settlementDetail" orm:"with:id=settlement_id"  dc:"结算模式信息"`
	GoodsList []*struct {
		gmeta.Meta   `orm:"table:hg_food_goods"`
		RestaurantId int64  `json:"restaurantId"  orm:"restaurant_id"     description:"餐厅ID"`
		GoodsName    string `json:"goodsName"    orm:"goods_name"    description:"套餐名称"`
		GoodsState   int    `json:"goodsState"   orm:"goods_state"   description:"状态（1.正常2下架）"`
	} `json:"goodsList" orm:"with:restaurant_id=id" dc:"套餐列表"`
}

type FoodRestaurantAllListModel struct {
	Id                  int    `json:"id"                        dc:"优惠券ID"`
	Name                string `json:"name"                    dc:"名称"`
	OpenStatus          string `json:"openStatus"              dc:"营业状态"`
	CooperateTypeId     int    `json:"cooperateTypeId"         dc:"合作类型ID"`
	AreaPid             int    `json:"areaPid"                 dc:"地区省级ID"`
	AreaId              int    `json:"areaId"                  dc:"地区市级ID"`
	DetailAddress       string `json:"detailAddress"                 dc:"详细地址"`
	CooperateTypeDetail *struct {
		gmeta.Meta `orm:"table:hg_food_cooperate_type"`
		*entity.FoodCooperateType
	} `json:"cooperateTypeDetail" orm:"with:id=cooperate_type_id"  dc:"合作类型"`
	PAreaDetail *struct {
		gmeta.Meta `orm:"table:hg_food_area"`
		*entity.FoodArea
	} `json:"pAreaDetail" orm:"with:id=area_pid"  dc:"地址信息"`
	AreaDetail *struct {
		gmeta.Meta `orm:"table:hg_food_area"`
		*entity.FoodArea
	} `json:"areaDetail" orm:"with:id=area_id"  dc:"地址信息"`
}

// FoodRestaurantStatusInp 更新餐厅状态
type FoodRestaurantStatusInp struct {
	Id     int    `json:"id" v:"required#id不能为空" dc:"id"`
	Status string `json:"status" dc:"状态"`
	Desc   string `json:"desc" dc:"备注"`
}

func (in *FoodRestaurantStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.Status) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type FoodRestaurantStatusModel struct{}

// FoodRestaurantRestVerifyCodeInp 重置核销码
type FoodRestaurantRestVerifyCodeInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodRestaurantRestVerifyCodeInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	return
}

type FoodRestaurantRestVerifyCodeModel struct{}

// FoodRestaurantSwitchInp 更新预定状态
type FoodRestaurantSwitchInp struct {
	Id       int `json:"id" v:"required#id不能为空" dc:"id"`
	CanOrder int `json:"canOrder" dc:"状态"`
}

func (in *FoodRestaurantSwitchInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.CanOrder) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type FoodRestaurantSwitchModel struct{}

// FoodRestaurantSortInp 编辑餐厅排序
type FoodRestaurantSortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *FoodRestaurantSortInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodRestaurantSortModel struct{}
