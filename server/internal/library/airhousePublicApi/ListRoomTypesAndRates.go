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
	ListRoomTypesAndRates = "/room_type_and_rate_plans"
)

func GetListRoomTypesAndRates(ctx context.Context, params *ListRoomTypesAndRatesRequest) (response *ListRoomTypesAndRatesJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, ListRoomTypesAndRates, params); err != nil {
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

type ListRoomTypesAndRatesRequest struct {
	PropertyId  string `json:"property_id"`
	PropertyUid string `json:"property_uid"`
	RoomTypeId  string `json:"room_type_id"`
	RoomTypeUid string `json:"room_type_uid"`
}

type ListRoomTypesAndRatesJSONDataResponse struct {
	ObjectType string                      `json:"object_type"`
	Data       []ListRoomTypesAndRatesData `json:"data"`
}

type AdditionalGuestAmounts struct {
	AgeCategories []string `json:"age_categories"`
	Amount        int      `json:"amount"`
}
type DefaultBaseRate struct {
	BaseRate          int `json:"base_rate"`
	MondayBaseRate    int `json:"monday_base_rate"`
	TuesdayBaseRate   int `json:"tuesday_base_rate"`
	WednesdayBaseRate int `json:"wednesday_base_rate"`
	ThursdayBaseRate  int `json:"thursday_base_rate"`
	FridayBaseRate    int `json:"friday_base_rate"`
	SaturdayBaseRate  int `json:"saturday_base_rate"`
	SundayBaseRate    int `json:"sunday_base_rate"`
}
type PerDayPricing struct {
	OccupantsForBaseRate   int                      `json:"occupants_for_base_rate"`
	AdditionalGuestAmounts []AdditionalGuestAmounts `json:"additional_guest_amounts"`
	DefaultBaseRate        DefaultBaseRate          `json:"default_base_rate"`
	IsTierRatesEnabled     bool                     `json:"is_tier_rates_enabled"`
}
type ServiceFeesPerStay struct {
	FeeType       string `json:"fee_type"`
	AmountPerStay int    `json:"amount_per_stay"`
}
type ListRoomTypesAndRatesData struct {
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
