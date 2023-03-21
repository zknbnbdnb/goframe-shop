package main

import (
	_ "goframe-shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
