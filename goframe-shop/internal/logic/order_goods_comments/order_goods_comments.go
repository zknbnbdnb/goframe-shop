package order_goods_comments

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sOrderGoodsComments struct {
}

func init() {
	service.RegisterOrderGoodsComments(New())
}

func New() *sOrderGoodsComments {
	return &sOrderGoodsComments{}
}

func (s *sOrderGoodsComments) Add(ctx context.Context, in model.AddOrderGoodsCommentsInput) (out model.AddOrderGoodsCommentsOutput, err error) {
	id, err := dao.OrderGoodsCommentsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddOrderGoodsCommentsOutput{
		Id: gconv.Uint(id),
	}, nil
}

func (s *sOrderGoodsComments) Delete(ctx context.Context, in model.DeleteOrderGoodsCommentsInput) (out model.DeleteOrderGoodsCommentsOutput, err error) {
	_, err = dao.OrderGoodsCommentsInfo.Ctx(ctx).WherePri(in.Id).Delete()
	if err != nil {
		return out, err
	}
	return model.DeleteOrderGoodsCommentsOutput{Id: in.Id}, nil
}
