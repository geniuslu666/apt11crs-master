// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsProperty is the golang structure of table hg_pms_property for DAO operations like Where/Data.
type PmsProperty struct {
	g.Meta               `orm:"table:hg_pms_property, do:true"`
	Id                   interface{} // 主键
	Uid                  interface{} // 在API合作伙伴系统中的物业ID
	Icon                 interface{} // 图标
	GroupIds             interface{} // 开放显示的会员分组IDS
	Name                 interface{} // 物业名称   多语言
	Style                interface{} // 物业类型
	Cover                interface{} // 封面
	Currency             interface{} // 物业的默认货币
	Language             interface{} // 物业的默认自动消息语言
	TimeZone             interface{} // 标准时区名称
	Address              interface{} // 地址描述
	Lat                  interface{} // 纬度
	Lng                  interface{} // 经度
	AddressDetail        interface{} // 百度详细地址
	GgAddressDetail      interface{} // google详细地址
	GgLat                interface{} // 谷歌纬度
	GgLng                interface{} // 谷歌经度
	TagList              interface{} // 标签多语言
	RoomDes              interface{} // 房间描述
	Surroundings         interface{} // 周边环境
	Description          interface{} // 物业描述  多语言
	BusStation           interface{} // 公交站
	Subway               interface{} // 地铁站
	ContactName          interface{} // 联系人
	Phone                interface{} // 联系方式
	ContactEmail         interface{} // 邮箱
	MinDaysNotice        interface{} // 短租模式最小预定区间
	MaxDaysNotice        interface{} // 最大预定区间
	MinutesAfterCheckout interface{} // 退房后 分钟
	MinutesBeforeCheckin interface{} // 入住前 分钟
	BookingLeadTimeLabel interface{} // 预约期限
	TurnoverDays         interface{} // 周转天数
	LinePrice            interface{} // 单价
	Price                interface{} // 单价
	CheckinAt            interface{} // 入住时间
	CheckoutAt           interface{} // 退房时间
	CancelPolicy         interface{} // 取消政策 多语言
	GalleryImages        interface{} // 画廊图片
	GalleryCover         interface{} // 画廊封面
	RequiredBook         interface{} //
	Close                interface{} // 1、开启该物业  2、关闭该物业
	LeaseClose           interface{} // 1、开启物业短租  2、关闭物业短租
	BookingClose         interface{} // 1、开启预订模式  2、关闭预订模式
	Sort                 interface{} // 排序(越大越靠前)
	CheckInGuide         interface{} // 入住指南
	SpaCanOrder          interface{} // 是否开放预定按摩服务  1开放  2关闭
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
	DeletedAt            *gtime.Time // 删除时间
	BatchReservation     interface{} // 是否开启多房型预定 （Y 开启 N  关闭）
	RegionId             interface{} // 地区
	AccessPass           interface{} // 门禁密码
}
