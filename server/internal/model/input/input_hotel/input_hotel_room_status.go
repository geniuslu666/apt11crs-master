package input_hotel

type RoomStatus struct {
	StartDate   string `json:"startDate" v:"required|date#开始日期必填|开始日期格式错误" dc:"开始时间"`
	EndDate     string `json:"endDate" v:"required|date#结束日期必填|结束日期格式错误" dc:"结束时间"`
	PUid        string `json:"puid" dc:"物业ID"`
	RoomTypeUid string `json:"roomTypeUid" dc:"房型ID"`
	RoomUtilUid string `json:"roomUtilUid" dc:"房间ID"`
}
