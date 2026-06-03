package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	RetrieveStay = "/stays/"
)

// GetRetrieveStay 检索订单信息
func GetRetrieveStay(ctx context.Context, id string) (response *RetrieveStayJSONDataRequest, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, RetrieveStay+id, nil); err != nil {
		return
	}
	if httpStatus != http.StatusOK {
		err = gerror.New(responseStr)
		return
	}
	response = new(RetrieveStayJSONDataRequest)
	if err = json.Unmarshal([]byte(responseStr), &response); err != nil {
		return
	}
	if g.IsEmpty(response.Data) {
		err = gerror.New(responseStr)
		return
	}
	return
}

type RetrieveStayJSONDataRequest struct {
	ObjectType string   `json:"object_type"`
	Data       StayData `json:"data"`
}
