package user_coupon

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

func (s *sUserCoupon) Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	lastInsertID, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCouponCreateOutput{UserCouponId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sUserCoupon) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
		dao.UserCouponInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) error {
	_, err := dao.UserCouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.UserCouponInfo.Columns().Id).
		Where(dao.UserCouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	m := dao.UserCouponInfo.Ctx(ctx)
	if err = gconv.Struct(in, &out); err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
