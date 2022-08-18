package web

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/wanganlin/goframe/app/request/web"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (a *cIndex) Index(ctx context.Context, req *web.HelloReq) (res *web.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteTpl("index.html")
	return
}
