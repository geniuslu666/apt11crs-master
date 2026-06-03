package main

import (
	"APT/internal/library/MobilePush"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	var (
		ctx context.Context
	)
	_ = MobilePush.SendPush(ctx,
		"welcome to live one member",
		"welcome to live one member",
		g.SliceStr{"88800045"}, g.MapStrStr{
			"path": "/service_detail",
			"data": gjson.New(g.Map{
				"type":    "C",
				"orderid": "22222222",
				"name":    "你好",
				"phone":   "15605284028",
				"member":  "Member",
				"status":  "1",
			}).String(),
		})
}
