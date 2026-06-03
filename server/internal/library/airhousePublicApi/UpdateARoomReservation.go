package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	UpdateARoomReservation = "/room_reservations/"
)

// UpdateRoomReservationPost 更新入住订单信息
func UpdateRoomReservationPost(ctx context.Context, id string, params map[string]interface{}) (response *RetrieveRoomReservationJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiPut(ctx, UpdateARoomReservation+id, params); err != nil {
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
