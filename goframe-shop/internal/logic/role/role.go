package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 创建
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	// 插入数据和返回ID
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: uint(lastInsertID)}, err
}

// AddPermission 角色添加权限
func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error) {
	id, err := dao.RolePermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.RoleAddPermissionOutput{}, err
	}
	return model.RoleAddPermissionOutput{Id: uint(id)}, err
}

// Delete 删除
func (s *sRole) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
		dao.RoleInfo.Columns().Id: id,
	}).Unscoped().Delete() // 这里是软删除，如果要实现硬删除可以在Delete()前面加上Unscoped()
	return err
}

// DeletePermission 删除角色权限
func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).Where(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	if err != nil {
		return err
	}
	return err
}

// Update 修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	// 跟新密码
	_, err := dao.RoleInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	m := dao.RoleInfo.Ctx(ctx)
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
