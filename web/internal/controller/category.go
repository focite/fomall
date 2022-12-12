package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Category = cCategory{}

type (
	cCategory struct{}

	cCategoryIndexReq struct {
		g.Meta `path:"/category" tags:"Web.Category" method:"get" summary:"You first page"`
	}
)

func (a *cCategory) Index(ctx context.Context, req *cCategoryIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("category.html")
	if err != nil {
		return nil, err
	}
	return
}
