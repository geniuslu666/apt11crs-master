package airhousePublicApi

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/net/context"
	"net/http"
)

var (
	ListStay = "/stays"
)

// GetListStay 获取订单列表
func GetListStay(ctx context.Context, params *ListStayJSONDataRequest) (response *ListStayJSONDataResponse, err error) {
	var (
		responseStr string
		httpStatus  int
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, ListStay, params); err != nil {
		return
	}
	if httpStatus != http.StatusOK {
		err = gerror.New(responseStr)
		return
	}
	if err = json.Unmarshal([]byte(responseStr), &response); err != nil {
		return
	}
	return
}

type ListStayJSONDataRequest struct {
	After             interface{} `json:"after"`              //返回项后面跟着特定的游标
	UID               interface{} `json:"uid"`                //过滤出具有特定供应商UID的对象
	FromCheckinAt     interface{} `json:"from_checkin_at"`    //如果任何房间预订的入住日期在该日期期间，则返回项目。包容
	FromCheckoutAt    interface{} `json:"from_checkout_at"`   //如果任何房间预订的退房日期在该日期期间，则返回项目。包容
	FromStaysAt       interface{} `json:"from_stays_at"`      //如果在该日期期间有任何预订的房间留在酒店，则返回物品。包容
	PropertyID        interface{} `json:"property_id"`        //属性的UUID
	PropertyUID       interface{} `json:"property_uid"`       //属性的UID
	ReservationNumber interface{} `json:"reservation_number"` //根据预订号筛选住宿
	ToCheckinAt       interface{} `json:"to_checkin_at"`      //如果任何房间预订的入住日期在该日期期间，则返回项目。包容
	ToCheckoutAt      interface{} `json:"to_checkout_at"`     //如果任何房间预订的退房日期在该日期期间，则返回项目。包容
	ToStaysAt         interface{} `json:"to_stays_at"`        //如果在该日期期间有任何预订的房间留在酒店，则返回物品。包容
}

type ListStayJSONDataResponse struct {
	ObjectType string     `json:"object_type"`
	Paginator  Paginator  `json:"paginator"`
	Data       []StayData `json:"data"`
}
