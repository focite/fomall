package middleware

import (
	"net/http"
	"time"

	"github.com/gogf/csrf/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CSRF 跨站点请求中间件
func CSRF() func(r *ghttp.Request) {
	return csrf.NewWithCfg(csrf.Config{
		Cookie: &http.Cookie{
			Name:     "_csrf",
			Path:     "/",
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		},
		ExpireTime:      time.Minute * 20,
		TokenLength:     32,
		TokenRequestKey: "X-CSRF-Token",
	})
}
