package cmd

import (
	"context"

	"gomall/auth/internal/controller"
	"gomall/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.CSRF())
				group.Bind(controller.Login)
				group.Bind(controller.Register)
				group.Bind(controller.Forget)
				group.Bind(controller.Reset)
			})
			s.Run()
			return nil
		},
	}
)
