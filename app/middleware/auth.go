package middleware

import (
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Auth 认证中间件
func Auth(guard string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		if auth, _ := r.Session.Get(guard + "_auth"); auth == nil {
			callback := r.Request.RequestURI
			if callback != "" {
				callback = "?callback=" + gurl.Encode(callback)
			}
			r.Response.RedirectTo("/" + guard + "/login" + callback)
		}
		r.Middleware.Next()
	}
}
