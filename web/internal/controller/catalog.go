package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Catalog = cCatalog{}

type (
	cCatalog struct{}

	cCatalogIndexReq struct {
		g.Meta `path:"/catalog" tags:"Web.Catalog" method:"get" summary:"You first page"`
	}
)

func (a *cCatalog) Index(ctx context.Context, req *cCatalogIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("catalog.html")
	if err != nil {
		return nil, err
	}
	return
}
