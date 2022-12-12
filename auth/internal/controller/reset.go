package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Reset = cReset{}

type (
	cReset struct{}

	resetFormReq struct {
		g.Meta `path:"/reset" tags:"认证模块" method:"get" summary:"重置密码表单"`
	}

	resetReq struct {
		g.Meta    `path:"/reset" tags:"认证模块" method:"post" summary:"重置密码操作"`
		Token     string `v:"required" dc:"找回邮件中的token密钥"`
		Password  string `v:"required" dc:"新的登录密码"`
		Password2 string `v:"required" dc:"重复新密码"`
	}
)

func (a *cReset) ShowForm(ctx context.Context, req *resetFormReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("reset.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cReset) Reset(ctx context.Context, req *resetReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("reset.html")
	if err != nil {
		return nil, err
	}
	return
}
