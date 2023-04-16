package frontend

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Refund 管理
var Refund = cRefund{}

type cRefund struct{}

func (c *cRefund) Add(ctx context.Context, req *frontend.AddRefundReq) (res *frontend.AddRefundRes, err error) {
	refundAddInput := model.AddRefundInput{}
	err = gconv.Struct(req, &refundAddInput)
	if err != nil {
		return nil, err
	}
	out, err := service.Refund().Add(ctx, refundAddInput)
	if err != nil {
		return nil, err
	}
	return &frontend.AddRefundRes{
		Id: out.Id,
	}, nil
}

func (c *cRefund) List(ctx context.Context, req *frontend.ListRefundReq) (res *frontend.ListRefundRes, err error) {
	listInput := model.ListRefundInput{}
	err = gconv.Struct(req, &listInput)
	if err != nil {
		return nil, err
	}
	orderListOutput, err := service.Refund().List(ctx, listInput)
	if err != nil {
		return nil, err
	}
	return &frontend.ListRefundRes{
		List:  orderListOutput.List,
		Page:  orderListOutput.Page,
		Size:  orderListOutput.Size,
		Total: orderListOutput.Total,
	}, nil
}

func (c *cRefund) Detail(ctx context.Context, req *frontend.DetailRefundReq) (res *frontend.DetailRefundRes, err error) {
	detail, err := service.Refund().Detail(ctx, model.DetailRefundInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &frontend.DetailRefundRes{}
	err = gconv.Struct(detail, res)
	if err != nil {
		return nil, err
	}
	return
}
