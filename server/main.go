package main

import (
	"APT/internal/global"
	_ "APT/internal/logic"
	_ "APT/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"APT/internal/cmd"
)

func main() {
	var ctx = gctx.GetInitCtx()
	global.Init(ctx)
	cmd.Main.Run(gctx.GetInitCtx())
}
