package main

import (
	_ "demo/internal/packed"

	_ "demo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
