package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"gomall/internal/middleware"
	"gomall/user/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth("user"))
				group.Bind(controller.Index)
				group.Bind(controller.Address)
				group.Bind(controller.Order)
				group.Bind(controller.Profile)
			})
			s.Run()
			return nil
		},
	}
)
