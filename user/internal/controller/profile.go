package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Profile = cProfile{}

type (
	cProfile struct{}

	cProfileIndexReq struct {
		g.Meta `path:"/profile" tags:"用户-个人资料" method:"get" summary:"个人资料表单"`
	}

	cProfileUpdateReq struct {
		g.Meta `path:"/profile" tags:"用户-个人资料" method:"post" summary:"保存个人资料"`
	}
)

func (a *cProfile) ShowForm(ctx context.Context, req *cProfileIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("profile.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cProfile) Update(ctx context.Context, req *cProfileUpdateReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("profile.html")
	if err != nil {
		return nil, err
	}
	return
}
