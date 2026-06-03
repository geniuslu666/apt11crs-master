package model_gateway

type GatewayCode struct {
	ResCode     int         `json:"resCode"`
	ResMessage  string      `json:"resMessage"`
	ResData     interface{} `json:"resData"`
	ResJsonData interface{} `json:"resJsonData"`
}
