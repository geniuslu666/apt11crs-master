package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	CreateStay = "/stays"
)

// CreateStayPost 创建stay订单
func CreateStayPost(ctx context.Context, params *CreateStayJSONDataRequest) (response *CreateStayJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	params.ReservationSourceCode = "manually_created"
	params.ReservationSourceName = "AppAirHost"
	if responseStr, err, httpStatus = CurlAirHostPublicApiPost(ctx, CreateStay, params); err != nil {
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

type CreateStayJSONDataRequest struct {
	ObjectType            string              `json:"object_type"`
	UID                   string              `json:"uid"`
	Property              Property            `json:"property"`
	ReservationNumber     string              `json:"reservation_number"`
	ReservationSourceCode string              `json:"reservation_source_code"`
	ReservationSourceName string              `json:"reservation_source_name"`
	Booker                Booker              `json:"booker"`
	RoomReservations      []*RoomReservations `json:"room_reservations"`
}

type CreateStayJSONDataResponse struct {
	ObjectType string   `json:"object_type"`
	Data       StayData `json:"data"`
}
