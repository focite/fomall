package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"gomall/auth/api/v1"
	"gomall/internal/consts"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!"+consts.AppName)
	return
}
