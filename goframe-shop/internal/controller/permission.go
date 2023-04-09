package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Permission 角色管理
var Permission = cPermission{}

type cPermission struct{}

func (c *cPermission) Create(ctx context.Context, req *backend.PermissionReq) (res *backend.PermissionRes, err error) {
	data := model.PermissionCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Permission().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.PermissionRes{PermissionId: out.PermissionId}, nil
}

func (c *cPermission) Delete(ctx context.Context, req *backend.PermissionDeleteReq) (res *backend.PermissionDeleteRes, err error) {
	err = service.Permission().Delete(ctx, req.Id)
	return
}

func (c *cPermission) Update(ctx context.Context, req *backend.PermissionUpdateReq) (res *backend.PermissionUpdateRes, err error) {
	data := model.PermissionUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Permission().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.PermissionUpdateRes{Id: req.Id}, nil
}

func (c *cPermission) List(ctx context.Context, req *backend.PermissionGetListCommonReq) (res *backend.PermissionGetListCommonRes, err error) {
	getListRes, err := service.Permission().GetList(ctx, model.PermissionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PermissionGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, err
}
