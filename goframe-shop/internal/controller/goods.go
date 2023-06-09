package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Goods 内容管理
var Goods = cGoods{}

type cGoods struct{}

func (c *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
	data := model.GoodsCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsRes{GoodsId: out.GoodsId}, nil
}

func (c *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

func (c *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	data := model.GoodsUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsUpdateRes{Id: req.Id}, nil
}

func (c *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.GoodsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *cGoods) Detail(ctx context.Context, req *frontend.GoodsDetailReq) (res *frontend.GoodsDetailRes, err error) {
	data := model.GoodsDetailInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	tmp, err := service.Goods().Detail(ctx, data)
	err = gconv.Struct(tmp, &res)
	if err != nil {
		return res, err
	}
	return
}
