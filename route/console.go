package route

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/wanganlin/goframe/app/controller/console"
	"github.com/wanganlin/goframe/app/middleware"
)

func Console(s *ghttp.Server) {
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CSRF())
		group.Bind(console.Auth)
		group.Middleware(middleware.Auth("admin"))
		group.ALLMap(g.Map{
			"/":     console.Index,
			"/user": console.User,
		})
	})
}
