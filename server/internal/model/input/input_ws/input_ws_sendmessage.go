package input_ws

type ChatSendMessageIpt struct {
	Source    string `json:"source" v:"in:Member,Driver,System,Technician,DriverLeader,TechnicianLeader,TechnicianISP#请上送身份标识|身份标识错误" d:"Member"`
	Message   string `json:"message" v:"required#请上送发送的消息内容"`
	Type      string `json:"type" v:"required|in:text,img#请上送消息类型"`
	OrderSn   string `json:"orderSn" v:"required#请上送接送机订单"`
	OrderType string `json:"orderType" v:"required|in:car,spa#请上送订单类型|订单类型错误"`
}

type ChatMemberIds struct {
	MemberId int64  `json:"memberId"`
	MemberNo string `json:"memberNo"`
	Source   string `json:"source"`
	SourceId int64  `json:"sourceId"`
	Name     string `json:"name"`
}

type OrderUnreadListInput struct {
	OrderSn string `json:"orderSn" v:"required#order_number_miss" dc:"订单号"`
}

type OrderUnreadListModel struct {
	OrderSn     string `json:"orderSn" dc:"订单号"`
	UnReadCount int    `json:"unReadCount" dc:"未读数"`
}

type UpdateLatestMsgReadInput struct {
	OrderSn  string `json:"orderSn" v:"required#order_number_miss" dc:"订单号"`
	ImId     int64  `json:"imId"  dc:"消息ID"`
	MemberId int64  `json:"MemberId"  dc:"会员ID"`
}
