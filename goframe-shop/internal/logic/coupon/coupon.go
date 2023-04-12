package coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

func (s *sCoupon) Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error) {
	lastInsertID, err := dao.CouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CouponCreateOutput{CouponId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sCoupon) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.CouponInfo.Ctx(ctx).Where(g.Map{
		dao.CouponInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) error {
	_, err := dao.CouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CouponInfo.Columns().Id).
		Where(dao.CouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	m := dao.CouponInfo.Ctx(ctx)
	if err = gconv.Struct(in, &out); err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)

	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
