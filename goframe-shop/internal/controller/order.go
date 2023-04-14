package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Order 订单管理
var Order = cOrder{}

type cOrder struct{}

// Add 添加订单
func (c *cOrder) Add(ctx context.Context, req *frontend.OrderAddReq) (res *frontend.OrderAddRes, err error) {
	orderAddInput := model.OrderAddInput{}
	//注意：这里要用scan 而不是struct, 因为里面有数组,只有纯粹的结构体才能用struct
	err = gconv.Scan(req, &orderAddInput)
	if err != nil {
		return nil, err
	}
	addRes, err := service.Order().Add(ctx, orderAddInput)
	if err != nil {
		return nil, err
	}

	return &frontend.OrderAddRes{
		Id: addRes.Id,
	}, nil
}
