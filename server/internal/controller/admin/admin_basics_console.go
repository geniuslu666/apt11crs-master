package admin

import (
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) ConsoleStat(ctx context.Context, req *basics.ConsoleStatReq) (res *basics.ConsoleStatRes, err error) {
	res = new(basics.ConsoleStatRes)

	// 此处均为模拟数据，可以根据实际业务情况替换成真实数据

	res.Visits.DayVisits = 12010
	res.Visits.Rise = 13501
	res.Visits.Decline = 10502
	res.Visits.Amount = 10403

	res.Saleroom.WeekSaleroom = 20501
	res.Saleroom.Amount = 21002
	res.Saleroom.Degree = 83.66

	res.OrderLarge.WeekLarge = 39901
	res.OrderLarge.Rise = 31012
	res.OrderLarge.Decline = 30603
	res.OrderLarge.Amount = 36084

	res.Volume.WeekLarge = 40021
	res.Volume.Rise = 40202
	res.Volume.Decline = 45003
	res.Volume.Amount = 49004
	return
}
