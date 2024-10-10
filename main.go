package main

import (
	_ "k8sctl/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"k8sctl/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
