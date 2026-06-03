package input_hotel

type CheckHotelInventoryInp struct {
	StartDate string `json:"startDate" dc:"入住时间"`
	EndDate   string `json:"endDate" dc:"退房时间"`
	PUID      string `json:"houseId" dc:"物业ID"`
	TUID      string `json:"roomTypeId" dc:"房型ID"`
	Num       int    `json:"roomNum" dc:"预定数量"`
}
