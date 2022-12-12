package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Index = cIndex{}

type (
	cIndex struct{}

	indexReq struct {
		g.Meta `path:"/" tags:"Web.Index" method:"get" summary:"You first page"`
	}
)

func (a *cIndex) Index(ctx context.Context, req *indexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("index.html")
	if err != nil {
		return nil, err
	}
	return
}
