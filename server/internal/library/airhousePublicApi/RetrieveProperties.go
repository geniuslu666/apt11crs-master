package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

var (
	RetrieveProperties = "/properties"
)

func GetRetrieveProperties(ctx context.Context, Uid string) (response RetrievePropertiesJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, RetrieveProperties+"/"+Uid, nil); err != nil {
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

type RetrievePropertiesJSONDataRequest struct {
	Page int         `json:"page"`
	Size int         `json:"size"`
	Uid  interface{} `json:"uid"`
}

type RetrievePropertiesJSONDataResponse struct {
	ObjectType string         `json:"object_type"`
	Paginator  Paginator      `json:"paginator"`
	Data       PropertiesData `json:"data"`
}
