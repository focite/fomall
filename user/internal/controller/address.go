package controller

import (
	"context"

	"gomall/internal/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Address = cAddress{}

type (
	cAddress struct{}

	cAddressIndexReq struct {
		g.Meta `path:"/address" tags:"用户-收货地址" method:"get" summary:"地址列表"`
	}

	cAddressCreateReq struct {
		g.Meta `path:"/address/create" tags:"用户-收货地址" method:"get" summary:"创建地址"`
	}

	cAddressStoreReq struct {
		g.Meta `path:"/address" tags:"用户-收货地址" method:"post" summary:"保存地址"`
	}

	cAddressShowReq struct {
		g.Meta `path:"/address/edit" tags:"用户-收货地址" method:"get" summary:"编辑地址"`
	}

	cAddressUpdateReq struct {
		g.Meta `path:"/address" tags:"用户-收货地址" method:"put" summary:"更新地址"`
	}

	cAddressDestroyReq struct {
		g.Meta `path:"/address" tags:"用户-收货地址" method:"delete" summary:"删除地址"`
	}
)

func (a *cAddress) Index(ctx context.Context, req *cAddressIndexReq) (res *response.Res, err error) {
	err = g.RequestFromCtx(ctx).Response.WriteTpl("address.html")
	if err != nil {
		return nil, err
	}
	return
}

func (a *cAddress) Create(ctx context.Context, req *cAddressCreateReq) (res *response.Res, err error) {
	return
}

func (a *cAddress) Store(ctx context.Context, req *cAddressStoreReq) (res *response.Res, err error) {
	return
}

func (a *cAddress) Show(ctx context.Context, req *cAddressShowReq) (res *response.Res, err error) {
	return
}

func (a *cAddress) Update(ctx context.Context, req *cAddressUpdateReq) (res *response.Res, err error) {
	return
}

func (a *cAddress) Destroy(ctx context.Context, req *cAddressDestroyReq) (res *response.Res, err error) {
	return
}
