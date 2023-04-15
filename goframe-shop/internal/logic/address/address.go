package address

import (
	"context"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sAddress struct{}

func init() {
	service.RegisterAddress(New())
}

func New() *sAddress {
	return &sAddress{}
}

// todo
func (s *sAddress) Add(ctx context.Context, in model.AddAddressInput) (out *model.AddAddressOutput, err error) {
	id, err := dao.AddressInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.AddAddressOutput{Id: int(id)}, err
}

func (s *sAddress) Update(ctx context.Context, in model.UpdateAddressInput) (err error) {
	_, err = dao.AddressInfo.Ctx(ctx).Data(in).FieldsEx(in.Id).Where(dao.AddressInfo.Columns().Id, in.Id).Update()
	// 跟新操作用到了FieldsEx()方法，这个方法是排除某个字段的意思
	// 即排除了已经在Data()方法中的字段，这样就可以实现跟新某个字段的操作
	if err != nil {
		return err
	}
	return
}

func (s *sAddress) Delete(ctx context.Context, in model.DeleteAddressInput) (err error) {
	_, err = dao.AddressInfo.Ctx(ctx).Where(dao.AddressInfo.Columns().Id, in.Id).Delete()
	if err != nil {
		return err
	}
	return err
}

func (s *sAddress) Page(ctx context.Context, in model.PageAddressInput) (out *model.PageAddressOutput, err error) {
	// 1.获得*gbd.Model对象,方便后续调用
	m := dao.AddressInfo.Ctx(ctx)
	// 2.实例化响应结构体
	out = &model.PageAddressOutput{}
	out.Page, out.Size = in.Page, in.Size
	// 3.分页查询
	listModel := m.Page(in.Page, in.Size)
	// 4.查询总数
	out.Total, err = listModel.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	// 5.延迟初始化list切片,确定有数据在按期望大小初始化切片容量
	out.List = make([]entity.AddressInfo, 0, in.Size)
	// 6.查询数据
	var list []entity.AddressInfo
	err = listModel.Scan(&list)
	if err != nil {
		return out, err
	}
	out.List = list
	return
}

// 客户端获取省市县地址
func (s *sAddress) GetCityList(ctx context.Context) (out *model.CityAddressListOutput, err error) {
	out = &model.CityAddressListOutput{}
	err = dao.AddressInfo.Ctx(ctx).Where(dao.AddressInfo.Columns().ParentId, consts.ProvincePid).WithAll().Scan(&out.List)
	if err != nil {
		return out, err
	}
	return
}
