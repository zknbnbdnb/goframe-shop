package rotation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Delete() // 这里是软删除，如果要实现硬删除可以在Delete()前面加上Unscoped()
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	m := dao.RotationInfo.Ctx(ctx)

	err = gconv.Scan(in, &out)
	if err != nil {
		return nil, err
	}
	//out = &model.RotationGetListOutput{
	//	Page: in.Page,
	//	Size: in.Size,
	//	List: []model.RotationGetListOutputItem{},
	//}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.Order(dao.RotationInfo.Columns().Sort)

	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
