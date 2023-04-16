package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// GoodsOptions 内容管理
var GoodsOptions = cGoodsOptions{}

type cGoodsOptions struct{}

func (c *cGoodsOptions) Create(ctx context.Context, req *backend.GoodsOptionsReq) (res *backend.GoodsOptionsRes, err error) {
	data := model.GoodsOptionsCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.GoodsOptions().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsRes{Id: out.Id}, nil
}

func (c *cGoodsOptions) Delete(ctx context.Context, req *backend.GoodsOptionsDeleteReq) (res *backend.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().Delete(ctx, req.Id)
	return
}

func (c *cGoodsOptions) Update(ctx context.Context, req *backend.GoodsOptionsUpdateReq) (res *backend.GoodsOptionsUpdateRes, err error) {
	data := model.GoodsOptionsUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.GoodsOptions().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsUpdateRes{Id: req.Id}, nil
}

func (c *cGoodsOptions) List(ctx context.Context, req *backend.GoodsOptionsGetListCommonReq) (res *backend.GoodsOptionsGetListCommonRes, err error) {
	getListRes, err := service.GoodsOptions().GetList(ctx, model.GoodsOptionsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.GoodsOptionsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
