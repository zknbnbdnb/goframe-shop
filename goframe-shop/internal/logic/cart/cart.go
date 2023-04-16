package cart

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sCart struct{}

func init() {
	service.RegisterCart(New())
}

func New() *sCart {
	return &sCart{}
}

func (s *sCart) Add(ctx context.Context, in model.AddCartInput) (out model.AddCartOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	//获取当前用户id
	condition := g.Map{
		dao.CartInfo.Columns().UserId:         in.UserId,
		dao.CartInfo.Columns().GoodsOptionsId: in.GoodsOptionsId,
	}
	count, err := dao.CartInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return out, err
	}
	//存在则更新 否则新增 即判断是否要增加购物车的商品数量还是新增一条购物车记录
	if count == 0 {
		id, err := dao.CartInfo.Ctx(ctx).Data(in).InsertAndGetId()
		if err != nil {
			return out, err
		}
		return model.AddCartOutput{Id: uint(id)}, nil
	}
	var cart = entity.CartInfo{}
	err = dao.CartInfo.Ctx(ctx).Where(condition).Scan(&cart)
	if err != nil {
		return out, err
	}
	_, err = dao.CartInfo.Ctx(ctx).Data(dao.CartInfo.Columns().Count, cart.Count+in.Count).WherePri(cart.Id).Update()
	if err != nil {
		return out, err
	}
	return model.AddCartOutput{Id: uint(cart.Id)}, nil
}

func (s *sCart) Delete(ctx context.Context, in model.DeleteCartInput) (out model.DeleteCartOutput, err error) {
	_, err = dao.CartInfo.Ctx(ctx).WherePri(in.Id).Delete()
	if err != nil {
		return out, err
	}
	return model.DeleteCartOutput{Id: uint(in.Id)}, nil
}

func (s *sCart) List(ctx context.Context, in model.ListCartInput) (out *model.ListCartOutput, err error) {
	// 1.获得*gbd.Model对象,方便后续调用
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	m := dao.CartInfo.Ctx(ctx).Where(dao.CartInfo.Columns().UserId, userId)
	// 2.实例化响应结构体
	out = &model.ListCartOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 3.分页查询
	listModel := m.Page(in.Page, in.Size)
	// 4.查询总数
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	// 5.延迟初始化list切片,确定有数据在按期望大小初始化切片容量
	out.List = make([]model.ListCartBase, 0, in.Size)
	// 6.把查询的结果赋值给响应结构体
	err = listModel.WithAll().Scan(&out.List)
	if err != nil {
		return nil, err
	}
	return out, err
}
