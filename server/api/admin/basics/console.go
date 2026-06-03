package basics

import "github.com/gogf/gf/v2/frame/g"

type ConsoleStatReq struct {
	g.Meta `path:"/console/stat" method:"get" tags:"ADMIN" summary:"控制台_综合数据统计"`
}

type ConsoleStatRes struct {
	Visits struct {
		DayVisits float64 `json:"dayVisits"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"visits"`
	Saleroom struct {
		WeekSaleroom float64 `json:"weekSaleroom"`
		Amount       float64 `json:"amount"`
		Degree       float64 `json:"degree"`
	} `json:"saleroom"`
	OrderLarge struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"orderLarge"`
	Volume struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"volume"`
}
