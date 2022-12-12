package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Register = cRegister{}

type (
	cRegister struct{}

	registerFormReq struct {
		g.Meta `path:"/register" tags:"认证模块" method:"get" summary:"注册表单"`
	}

	registerReq struct {
		g.Meta    `path:"/register" tags:"认证模块" method:"post" summary:"注册操作"`
		Email     string `v:"required" dc:"用户电子邮箱"`
		Password  string `v:"required" dc:"用户登录密码"`
		Password2 string `v:"required" dc:"重复登录密码"`
		Captcha   string `v:"required" dc:"图片验证码"`
	}
)

func (a *cRegister) ShowForm(ctx context.Context, req *registerFormReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("register.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cRegister) Register(ctx context.Context, req *registerReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("register.html")
	if err != nil {
		return nil, err
	}
	return
}
