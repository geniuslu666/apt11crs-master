package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
	"time"
)

var (
	Availabilities = "/availabilities"
)

func GetAvailabilities(ctx context.Context, params *AvailabilitiesJSONDataRequest) (response *AvailabilitiesJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, Availabilities, params); err != nil {
		return
	}
	if httpStatus != http.StatusOK {
		err = gerror.New(responseStr)
		return
	}
	if err = json.Unmarshal([]byte(responseStr), &response); err != nil {
		return
	}
	if g.IsEmpty(response.Data) {
		err = gerror.New(responseStr)
		return
	}
	return
}

type AvailabilitiesJSONDataRequest struct {
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	RoomTypeId string `json:"room_type_id"`
}
type AvailabilitiesJSONDataResponse struct {
	ObjectType string `json:"object_type"`
	Data       Data   `json:"data"`
}
type RoomTypeAndRatePlans struct {
	ObjectType            string      `json:"object_type"`
	RoomType              RoomType    `json:"room_type"`
	RoomTypeName          string      `json:"room_type_name"`
	RoomTypeAccommodates  int         `json:"room_type_accommodates"`
	RoomTypeCapacity      int         `json:"room_type_capacity"`
	RoomTypeMinBaseRate   int         `json:"room_type_min_base_rate"`
	Currency              string      `json:"currency"`
	Connected             bool        `json:"connected"`
	ConnectedRoomType     interface{} `json:"connected_room_type"`
	ConnectedRoomTypeName interface{} `json:"connected_room_type_name"`
	RatePlans             []RatePlans `json:"rate_plans"`
	UpdatedAt             time.Time   `json:"updated_at"`
}
type Rates struct {
	RatePlan          RatePlan    `json:"rate_plan"`
	ConnectedRatePlan interface{} `json:"connected_rate_plan"`
	StartDate         string      `json:"start_date"`
	EndDate           string      `json:"end_date"`
	MinLos            int         `json:"min_los"`
	MaxLos            int         `json:"max_los"`
	Closed            string      `json:"closed"`
	PerDayBaseRate    int         `json:"per_day_base_rate"`
}
type Inventories struct {
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	InventoryCount int    `json:"inventory_count"`
}
type Blocks struct {
	StartDate string      `json:"start_date"`
	EndDate   string      `json:"end_date"`
	RoomUnit  RoomUnit    `json:"room_unit"`
	Memo      interface{} `json:"memo"`
}
type Data struct {
	ObjectType           string               `json:"object_type"`
	BookingWindow        int                  `json:"booking_window"`
	BookingLeadTime      int                  `json:"booking_lead_time"`
	RoomTypeAndRatePlans RoomTypeAndRatePlans `json:"room_type_and_rate_plans"`
	Rates                []Rates              `json:"rates"`
	Inventories          []Inventories        `json:"inventories"`
	Blocks               []Blocks             `json:"blocks"`
	RatesMemos           []interface{}        `json:"rates_memos"`
}
