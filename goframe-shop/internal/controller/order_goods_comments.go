package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var OrderGoodsComments = cOrderGoodsComments{}

type cOrderGoodsComments struct{}

func (c *cOrderGoodsComments) Add(ctx context.Context, req *frontend.AddOrderGoodsCommentsReq) (res *frontend.AddOrderGoodsCommentsRes, err error) {
	in := model.AddOrderGoodsCommentsInput{}
	err = gconv.Struct(req, &in)
	if err != nil {
		return nil, err
	}
	add, err := service.OrderGoodsComments().Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return &frontend.AddOrderGoodsCommentsRes{
		Id: add.Id,
	}, err
}

func (c *cOrderGoodsComments) Delete(ctx context.Context, req *frontend.DelOrderGoodsCommentsReq) (res *frontend.DelOrderGoodsCommentsRes, err error) {
	out, err := service.OrderGoodsComments().Delete(ctx, model.DeleteOrderGoodsCommentsInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &frontend.DelOrderGoodsCommentsRes{Id: out.Id}, err
}
