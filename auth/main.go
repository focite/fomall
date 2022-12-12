package main

import (
	"gomall/auth/internal/cmd"
	_ "gomall/internal/service/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
