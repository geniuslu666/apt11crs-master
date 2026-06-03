package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"net/http"
	"time"
)

var (
	RetrieveAFolio = "/folios/"
)

func GetRetrieveAFolio(ctx context.Context, id string) (response *RetrieveAFolioJSONDataResponse, err error) {
	var (
		responseStr string
		httpStatus  int
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, RetrieveAFolio+id, nil); err != nil {
		return
	}
	if httpStatus != http.StatusOK {
		err = gerror.New(responseStr)
		return
	}
	if err = json.Unmarshal([]byte(responseStr), &response); err != nil {
		return
	}
	return
}

type RetrieveAFolioJSONDataResponse struct {
	ObjectType string    `json:"object_type"`
	Data       FolioData `json:"data"`
}

type Payments struct {
	ID                 string          `json:"id"`
	ObjectType         string          `json:"object_type"`
	UID                string          `json:"uid"`
	OutTradeNo         string          `json:"out_trade_no"`
	RoomReservation    RoomReservation `json:"room_reservation"`
	Folio              Folio           `json:"folio"`
	Date               string          `json:"date"`
	PaymentSourceType  string          `json:"payment_source_type"`
	PaymentAccountType string          `json:"payment_account_type"`
	Description        string          `json:"description"`
	Amount             float64         `json:"amount"`
	RefundAmount       float64         `json:"refund_amount"`
	State              string          `json:"state"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}
type FolioData struct {
	ID                 string          `json:"id"`
	ObjectType         string          `json:"object_type"`
	UID                string          `json:"uid"`
	RoomReservation    RoomReservation `json:"room_reservation"`
	Charges            []Charges       `json:"charges"`
	Payments           []Payments      `json:"payments"`
	TotalChargeAmount  float64         `json:"total_charge_amount"`
	TotalPaymentAmount float64         `json:"total_payment_amount"`
	OutstandingBalance float64         `json:"outstanding_balance"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}
