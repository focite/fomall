package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Login = cLogin{}

type (
	cLogin struct{}

	loginFormReq struct {
		g.Meta `path:"/login" tags:"认证模块" method:"get" summary:"登录表单"`
	}

	loginReq struct {
		g.Meta   `path:"/login" tags:"认证模块" method:"post" summary:"登录操作"`
		Username string `v:"required" dc:"用户名/手机号码/电子邮箱"`
		Password string `v:"required" dc:"用户登录密码"`
		Captcha  string `v:"required" dc:"图片验证码"`
	}
)

func (a *cLogin) ShowForm(ctx context.Context, req *loginFormReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("login.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cLogin) Login(ctx context.Context, req *loginReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("login.html")
	if err != nil {
		return nil, err
	}
	return
}
