package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Category 内容管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryReq) (res *backend.CategoryRes, err error) {
	data := model.CategoryCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Category().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.CategoryRes{CategoryId: out.CategoryId}, nil
}

func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	data := model.CategoryUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Category().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.CategoryUpdateRes{Id: req.Id}, nil
}

func (a *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CategoryGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
