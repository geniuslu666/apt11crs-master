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
	ListProperties = "/properties"
)

func GetListProperties(ctx context.Context, size int, Uid interface{}) (response ListPropertiesJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, ListProperties, &ListPropertiesJSONDataRequest{
		Page: 1,
		Size: size,
		Uid:  Uid,
	}); err != nil {
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

type ListPropertiesJSONDataRequest struct {
	Page int         `json:"page"`
	Size int         `json:"size"`
	Uid  interface{} `json:"uid"`
}

type ListPropertiesJSONDataResponse struct {
	ObjectType string           `json:"object_type"`
	Paginator  Paginator        `json:"paginator"`
	Data       []PropertiesData `json:"data"`
}

type PropertiesData struct {
	ID            string      `json:"id"`
	ObjectType    string      `json:"object_type"`
	UID           string      `json:"uid"`
	Name          string      `json:"name"`
	PropertyType  string      `json:"property_type"`
	Currency      string      `json:"currency"`
	Language      string      `json:"language"`
	TimeZone      string      `json:"time_zone"`
	BookingWindow int         `json:"booking_window"`
	RoomTypes     []RoomTypes `json:"room_types"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
