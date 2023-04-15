package controller

import (
	"context"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Address 角色管理
var Address = cAddress{}

type cAddress struct{}

func (c *cAddress) Add(ctx context.Context, req *backend.AddAddressReq) (res *backend.AddAddressRes, err error) {
	out, err := service.Address().Add(ctx, model.AddAddressInput{
		AddressBase: model.AddressBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AddAddressRes{
		Id: out.Id,
	}, err
}

func (c *cAddress) Update(ctx context.Context, req *backend.UpdateAddressReq) (res *backend.UpdateAddressRes, err error) {
	err = service.Address().Update(ctx, model.UpdateAddressInput{
		Id: req.Id,
		AddressBase: model.AddressBase{
			ParentId: req.ParentId,
			Name:     req.Name,
			Status:   req.Status,
		},
	})
	if err != nil {
		return res, err
	}
	return
}

func (c *cAddress) Delete(ctx context.Context, req *backend.DeleteAddressReq) (res *backend.DeleteAddressRes, err error) {
	err = service.Address().Delete(ctx, model.DeleteAddressInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	return
}

func (C *cAddress) Page(ctx context.Context, req *backend.PageAddressReq) (res *backend.PageAddressRes, err error) {
	out, err := service.Address().Page(ctx, model.PageAddressInput{
		CommonPaginationReq: backend.CommonPaginationReq{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		return res, err
	}
	return &backend.PageAddressRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  out.Page,
			Total: out.Total,
			Page:  out.Page,
			Size:  out.Size,
		},
	}, err
}

func (c *cAddress) CityList(ctx context.Context, req *backend.CityAddressListReq) (res *backend.CityAddressListRes, err error) {
	out, err := service.Address().GetCityList(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.CityAddressListRes{List: out.List}, err
}
