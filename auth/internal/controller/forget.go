package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Forget = cForget{}

type (
	cForget struct{}

	forgetFormReq struct {
		g.Meta `path:"/forget" tags:"认证模块" method:"get" summary:"忘记密码表单"`
	}

	forgetReq struct {
		g.Meta  `path:"/forget" tags:"认证模块" method:"post" summary:"找回密码"`
		Email   string `v:"required" dc:"用户的电子邮箱"`
		Captcha string `v:"required" dc:"图片验证码"`
	}
)

func (a *cForget) ShowForm(ctx context.Context, req *forgetFormReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("forget.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cForget) Forget(ctx context.Context, req *forgetReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("forget.html")
	if err != nil {
		return nil, err
	}
	return
}
