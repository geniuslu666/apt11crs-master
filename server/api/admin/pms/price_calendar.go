package pms

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type PriceCalendarReq struct {
	g.Meta `path:"/pms/priceCalendar" method:"post" tags:"ADMIN_PMS" summary:"PMS价格日历_列表"`
	Puid   string `json:"puid" dc:"物业id"`
	Tuid   string `json:"ruid" dc:"房型id"`
	Date   string `json:"date" v:"required|date#请输入日期|日期格式错误" dc:"日期"`
	Page   int    `json:"page" v:"required#请输入页码" dc:"页码"`
	Limit  int    `json:"limit" v:"required#请输入页长" dc:"页长"`
}

type PriceCalendarRes struct {
	Availabilities map[string]map[string][]*entity.PmsAvailabilities
	RoomTypeList   []*struct {
		*entity.PmsRoomType
		Property *struct {
			gmeta.Meta `orm:"table:hg_pms_property"`
			*entity.PmsProperty
		} `json:"property" orm:"with:uid=puid"`
	}
	Count int
}
