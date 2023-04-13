package collection

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) Add(ctx context.Context, in model.CollectionAddInput) (out *model.CollectionAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return out, err
	}
	return &model.CollectionAddOutput{Id: uint(id)}, nil
}

func (s *sCollection) Delete(ctx context.Context, in model.CollectionDeleteInput) (out model.CollectionDeleteOutput, err error) {
	if in.Id != 0 {
		_, err = dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return out, err
		}
		return model.CollectionDeleteOutput{Id: in.Id}, nil
	} else {
		in.UserId = gconv.Uint(ctx.Value("userId"))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty().Where(in).Delete() // todo 如果不加上空值忽略，那么id=0会成为查询条件但是默认的是空值所以会无法删除任何数据
		if err != nil {
			return out, err
		}
		return model.CollectionDeleteOutput{Id: gconv.Uint(id)}, nil
	}
}

func (s *sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	m := dao.CollectionInfo.Ctx(ctx)
	err = gconv.Struct(in, &out)
	if err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	if in.Type != 0 {
		listModel = listModel.Ctx(ctx).Where(dao.CollectionInfo.Columns().Type, in.Type)
	}

	/*
		//var list []*entity.CollectionInfo
		//if err := listModel.WithAll().Scan(&list); err != nil {
		//	return out, err
		//}
		//// 没有数据
		//if len(list) == 0 {
		//	return out, nil
		//} //这段代码和上段代码就是做一个事情,查询列表长度是否为空,冗余了所以要优化
	*/
	out.Total, err = listModel.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if in.Type == consts.CollectionTypeGoods {
		err = listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else if in.Type == consts.CollectionTypeArticle {
		err = listModel.With(model.ArticleItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else {
		err = listModel.WithAll().Scan(&out.List)
		if err != nil {
			return out, err
		}
	}
	return
}

func CollectionCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CollectionInfo.Columns().ObjectId: objectId,
		dao.CollectionInfo.Columns().Type:     collectionType,
	}
	count, err = dao.CollectionInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

func CollectionCheck(ctx context.Context, in model.CollectionCheckInput) (out *model.CollectionCheckOutput, err error) {
	condition := g.Map{
		dao.CollectionInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		dao.CollectionInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CollectionInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return out, err
	}
	if count > 0 {
		return &model.CollectionCheckOutput{IsCollect: true}, nil
	} else {
		return &model.CollectionCheckOutput{IsCollect: false}, nil
	}
}
