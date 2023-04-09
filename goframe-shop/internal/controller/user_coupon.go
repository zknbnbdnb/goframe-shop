package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// UserCoupon 内容管理
var UserCoupon = cUserCoupon{}

type cUserCoupon struct{}

func (a *cUserCoupon) Create(ctx context.Context, req *backend.UserCouponReq) (res *backend.UserCouponRes, err error) {
	data := model.UserCouponCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.UserCoupon().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.UserCouponRes{UserCouponId: out.UserCouponId}, nil
}

func (a *cUserCoupon) Delete(ctx context.Context, req *backend.UserCouponDeleteReq) (res *backend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, req.Id)
	return
}

func (a *cUserCoupon) Update(ctx context.Context, req *backend.UserCouponUpdateReq) (res *backend.UserCouponUpdateRes, err error) {
	data := model.UserCouponUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	err = service.UserCoupon().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.UserCouponUpdateRes{Id: req.Id}, nil
}

func (a *cUserCoupon) List(ctx context.Context, req *backend.UserCouponGetListCommonReq) (res *backend.UserCouponGetListCommonRes, err error) {
	getListRes, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.UserCouponGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
