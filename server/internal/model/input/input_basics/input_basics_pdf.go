package input_basics

type PdfLssIpt struct {
	UserName     string `json:"userName" dc:"用户姓名"`
	OrderSn      string `json:"orderSn"  dc:"订单号"`
	OrderAmount  int    `json:"orderAmount" dc:"订单金额"`
	StartTime    string `json:"startTime" dc:"开始日期"`
	EndTime      string `json:"endTime" dc:"结束日期"`
	CreateTime   string `json:"createTime" dc:"创建时间"`
	PropertyName string `json:"propertyName" dc:"物业名称"`
}
