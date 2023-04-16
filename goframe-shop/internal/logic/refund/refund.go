package refund

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sRefund struct{}

func init() {
	service.RegisterRefund(New())
}

func New() *sRefund {
	return &sRefund{}
}

func (s *sRefund) Add(ctx context.Context, in model.AddRefundInput) (out model.AddRefundOutput, err error) {
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	in.Number = utility.GetRefundNum()
	in.Status = consts.RefundStatusWait
	lastInsertId, err := dao.RefundInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddRefundOutput{Id: uint(lastInsertId)}, err
}

func (s *sRefund) List(ctx context.Context, in model.ListRefundInput) (out model.ListRefundOutput, err error) {
	// 1. 获得*gdb.Model对象，方便后续调用
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	m := dao.RefundInfo.Ctx(ctx).Where(dao.RefundInfo.Columns().UserId, userId)
	// 2. 实例化响应结构体
	out = model.ListRefundOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	// 5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.ListRefundOutputItem, 0, in.Size)
	// 6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sRefund) Detail(ctx context.Context, in model.DetailRefundInput) (out model.DetailRefundOutput, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.RefundInfo.Ctx(ctx).
		Where(dao.RefundInfo.Columns().UserId, userId).
		WithAll().
		WherePri(in.Id).
		Scan(&out)
	if err != nil {
		return out, err
	}
	return
}
