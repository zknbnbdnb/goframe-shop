package permission

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// Create 创建
func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	// 插入数据和返回ID
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		dao.PermissionInfo.Columns().Id: id,
	}).Unscoped().Delete() // 这里是软删除，如果要实现硬删除可以在Delete()前面加上Unscoped()
	return err
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	// 跟新密码
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	m := dao.PermissionInfo.Ctx(ctx)
	if err = gconv.Struct(in, &out); err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
