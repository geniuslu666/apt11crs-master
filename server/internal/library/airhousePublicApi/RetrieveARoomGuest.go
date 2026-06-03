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
	RetrieveARoomGuest = "/room_guests"
)

type RetrieveARoomGuestJSONDataResponse struct {
	ObjectType string `json:"object_type"`
	Data       []struct {
		ID              string      `json:"id"`
		ObjectType      string      `json:"object_type"`
		UID             interface{} `json:"uid"`
		RoomReservation struct {
			ID  string `json:"id"`
			UID string `json:"uid"`
		} `json:"room_reservation"`
		IsMainGuest       bool        `json:"is_main_guest"`
		FirstName         string      `json:"first_name"`
		LastName          string      `json:"last_name"`
		FirstNameKana     string      `json:"first_name_kana"`
		LastNameKana      string      `json:"last_name_kana"`
		FullName          string      `json:"full_name"`
		Gender            interface{} `json:"gender"`
		Dob               interface{} `json:"dob"`
		Nationality       string      `json:"nationality"`
		VisaNo            interface{} `json:"visa_no"`
		Occupation        interface{} `json:"occupation"`
		PostalCode        interface{} `json:"postal_code"`
		Address           interface{} `json:"address"`
		Email             string      `json:"email"`
		Phone             string      `json:"phone"`
		LastPortEmbark    interface{} `json:"last_port_embark"`
		NextPortDisembark interface{} `json:"next_port_disembark"`
		Photo             interface{} `json:"photo"`
		Selfie            interface{} `json:"selfie"`
		Signature         interface{} `json:"signature"`
		IdentPhoto        interface{} `json:"ident_photo"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
	} `json:"data"`
}

func GetRoomGuest(ctx context.Context, id string) (response *RetrieveARoomGuestJSONDataResponse, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, RetrieveARoomGuest+"?room_reservation_uid="+id, nil); err != nil {
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
