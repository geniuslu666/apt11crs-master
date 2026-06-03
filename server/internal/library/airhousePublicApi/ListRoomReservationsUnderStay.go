package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	ListRoomReservationsUnderStay = "/room_reservations"
)

// GetListRoomReservationsUnderStay 获取预定列表
func GetListRoomReservationsUnderStay(ctx context.Context, params ListRoomReservationsUnderStayJSONDataRequest) (response ListRoomReservationsUnderStayJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, ListRoomReservationsUnderStay, params); err != nil {
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

type ListRoomReservationsUnderStayJSONDataRequest struct {
	Uid     string `json:"uid"`
	StayId  string `json:"stay_id"`
	StayUid string `json:"stay_uid"`
}

type ListRoomReservationsUnderStayJSONDataResponse struct {
	ObjectType string             `json:"object_type"`
	Data       []RoomReservations `json:"data"`
}
