package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Collection = cCollection{}

type cCollection struct{}

// Add 添加收藏
func (c *cCollection) Add(ctx context.Context, req *frontend.CollectionAddReq) (res *frontend.CollectionAddRes, err error) {
	data := model.CollectionAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionAddRes{Id: out.Id}, nil
}

func (c *cCollection) Delete(ctx context.Context, req *frontend.CollectionDeleteReq) (res *frontend.CollectionDeleteRes, err error) {
	data := model.CollectionDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().Delete(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionDeleteRes{Id: out.Id}, nil
}

func (c *cCollection) List(ctx context.Context, req *frontend.CollectionListReq) (res *frontend.CollectionListRes, err error) {
	//getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
	//	Page: req.Page,
	//	Size: req.Size,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &backend.CategoryGetListCommonRes{List: getListRes.List,
	//	Page:  getListRes.Page,
	//	Size:  getListRes.Size,
	//	Total: getListRes.Total}, nil

	data := model.CollectionListInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().GetList(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionListRes{List: out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total}, nil
}
