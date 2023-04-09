package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Role 角色管理
var Role = cRole{}

type cRole struct{}

func (c *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	data := model.RoleCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Role().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.RoleId}, nil
}

// AddPermission 角色添加权限
func (c *cRole) AddPermission(ctx context.Context, req *backend.RoleAddPermissionReq) (res *backend.RoleAddPermissionRes, err error) {
	//out, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
	//	RoleId:       req.RoleId,
	//	PermissionId: req.PermissionId,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//return &backend.RoleAddPermissionRes{Id: out.Id}, nil
	data := model.RoleAddPermissionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Role().AddPermission(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.RoleAddPermissionRes{Id: out.Id}, nil
}

func (c *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

// DeletePermission 删除角色权限
func (c *cRole) DeletePermission(ctx context.Context, req *backend.RoleDeletePermissionReq) (res *backend.RoleDeletePermissionRes, err error) {
	//err = service.Role().DeletePermission(ctx, model.RoleDeletePermissionInput{
	//	RoleId:       req.RoleId,
	//	PermissionId: req.PermissionId,
	//})
	//return
	data := model.RoleDeletePermissionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Role().DeletePermission(ctx, data)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	data := model.RoleUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Role().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.RoleUpdateRes{Id: req.Id}, nil
}

func (c *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, err
}
