package support

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

// GetResponseCookie 读取中间件设置的cookies，一般在最后调用。
func GetResponseCookie(r *ghttp.Request, key string) string {
	res := http.Response{Header: r.Response.Header()}
	for _, v := range res.Cookies() {
		if v.Name == key {
			return v.Value
		}
	}
	return ""
}
