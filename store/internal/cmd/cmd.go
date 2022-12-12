package cmd

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gomall/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/store", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth("store"))

			})
			s.Run()
			return nil
		},
	}
)
