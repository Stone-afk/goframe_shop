package main

import (
	_ "goframe_shop/internal/packed"

	_ "goframe_shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe_shop/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
