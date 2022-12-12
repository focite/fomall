package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Brand = cBrand{}

type (
	cBrand struct{}

	cBrandIndexReq struct {
		g.Meta `path:"/brand" tags:"Web.Catalog" method:"get" summary:"You first page"`
	}
)

func (a *cBrand) Index(ctx context.Context, req *cBrandIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("brand.html")
	if err != nil {
		return nil, err
	}
	return
}
