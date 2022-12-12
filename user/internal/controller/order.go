package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Order = cOrder{}

type (
	cOrder struct{}

	cOrderIndexReq struct {
		g.Meta `path:"/order" tags:"用户-订单" method:"get" summary:"You first user page"`
	}
)

func (a *cOrder) Index(ctx context.Context, req *cOrderIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("order.html")
	if err != nil {
		return nil, err
	}
	return
}
