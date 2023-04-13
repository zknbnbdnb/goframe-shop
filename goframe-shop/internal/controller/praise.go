package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Praise = cPraise{}

type cPraise struct{}

// Add 添加收藏
func (c *cPraise) Add(ctx context.Context, req *frontend.PraiseAddReq) (res *frontend.PraiseAddRes, err error) {
	data := model.PraiseAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseAddRes{Id: out.Id}, nil
}

func (c *cPraise) Delete(ctx context.Context, req *frontend.PraiseDeleteReq) (res *frontend.PraiseDeleteRes, err error) {
	data := model.PraiseDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().Delete(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseDeleteRes{Id: out.Id}, nil
}

func (c *cPraise) List(ctx context.Context, req *frontend.PraiseListReq) (res *frontend.PraiseListRes, err error) {
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

	data := model.PraiseListInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().GetList(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseListRes{List: out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total}, nil
}
