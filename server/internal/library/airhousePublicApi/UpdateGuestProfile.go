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
	GuestProfile = "/guest_profiles"
)

type PutGuestProfileJSONDataResponse struct {
	ObjectType string `json:"object_type"`
	Data       struct {
		ID            string    `json:"id"`
		ObjectType    string    `json:"object_type"`
		UID           string    `json:"uid"`
		FirstName     string    `json:"first_name"`
		LastName      string    `json:"last_name"`
		FirstNameKana string    `json:"first_name_kana"`
		LastNameKana  string    `json:"last_name_kana"`
		FullName      string    `json:"full_name"`
		Language      string    `json:"language"`
		Email         string    `json:"email"`
		Phone         string    `json:"phone"`
		Nationality   string    `json:"nationality"`
		Address       string    `json:"address"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	} `json:"data"`
}

func PutGuestProfile(ctx context.Context, id string, params map[string]interface{}) (response *PutGuestProfileJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiPut(ctx, GuestProfile+"/"+id, params); err != nil {
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
