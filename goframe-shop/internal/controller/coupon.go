package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Coupon 内容管理
var Coupon = cCoupon{}

type cCoupon struct{}

func (a *cCoupon) Create(ctx context.Context, req *backend.CouponReq) (res *backend.CouponRes, err error) {
	data := model.CouponCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.Coupon().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.CouponRes{CouponId: out.CouponId}, nil
}

func (a *cCoupon) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.Id)
	return
}

func (a *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	data := model.CouponUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.Coupon().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.CouponUpdateRes{Id: req.Id}, nil
}

func (a *cCoupon) List(ctx context.Context, req *backend.CouponGetListCommonReq) (res *backend.CouponGetListCommonRes, err error) {
	getListRes, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CouponGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
