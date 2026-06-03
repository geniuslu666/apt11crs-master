package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	RetrieveRoomReservation = "/room_reservations/"
)

// GetRetrieveRoomReservation 检索预定信息
func GetRetrieveRoomReservation(ctx context.Context, id string) (response RetrieveRoomReservationJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, RetrieveRoomReservation+id, nil); err != nil {
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

type RetrieveRoomReservationJSONDataResponse struct {
	ObjectType string           `json:"object_type"`
	Data       RoomReservations `json:"data"`
}
