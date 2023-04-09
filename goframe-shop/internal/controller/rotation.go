package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	data := model.RotationCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Rotation().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.RotationRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return
}

func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	data := model.RotationUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Rotation().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.RotationUpdateRes{Id: req.Id}, nil
}

func (a *cRotation) List(ctx context.Context, req *backend.RotationGetListCommonReq) (res *backend.RotationGetListCommonRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, err
}

func (a *cRotation) ListFronted(ctx context.Context, req *frontend.RotationGetListCommonReq) (res *frontend.RotationGetListCommonRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.RotationGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, err
}
