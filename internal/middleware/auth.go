package middleware

import (
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"gomall/internal/consts"
	"net/http"
)

func Auth(guard string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		if auth, _ := r.Session.Contains(consts.AuthName + guard); auth == false {
			callback := r.Request.RequestURI
			if callback != "" {
				callback = "?callback=" + gurl.Encode(callback)
			}
			if r.IsAjaxRequest() {
				r.Response.WriteJsonExit(g.Map{
					"Code":    http.StatusForbidden,
					"Message": "Forbidden",
					"Data":    nil,
				})
			} else {
				r.Response.RedirectTo("/login?callback=" + callback)
			}
		}
		r.Middleware.Next()
	}
}
