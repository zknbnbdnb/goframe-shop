package praise

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) Add(ctx context.Context, in model.PraiseAddInput) (out *model.PraiseAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return out, err
	}
	return &model.PraiseAddOutput{Id: uint(id)}, nil
}

func (s *sPraise) Delete(ctx context.Context, in model.PraiseDeleteInput) (out model.PraiseDeleteOutput, err error) {
	if in.Id != 0 {
		_, err = dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return out, err
		}
		return model.PraiseDeleteOutput{Id: in.Id}, nil
	} else {
		in.UserId = gconv.Uint(ctx.Value("userId"))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete() // todo 如果不加上空值忽略，那么id=0会成为查询条件但是默认的是空值所以会无法删除任何数据
		if err != nil {
			return out, err
		}
		return model.PraiseDeleteOutput{Id: gconv.Uint(id)}, nil
	}
}

func (s *sPraise) GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	m := dao.PraiseInfo.Ctx(ctx)
	err = gconv.Struct(in, &out)
	if err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	if in.Type != 0 {
		listModel = listModel.Ctx(ctx).Where(dao.PraiseInfo.Columns().Type, in.Type)
	}

	out.Total, err = listModel.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if in.Type == consts.PraiseTypeGoods {
		err = listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
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

func PraiseCount(ctx context.Context, objectId uint, praiseType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     praiseType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

func PraiseCheck(ctx context.Context, in model.PraiseCheckInput) (out *model.PraiseCheckOutput, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return out, err
	}
	if count > 0 {
		return &model.PraiseCheckOutput{IsCollect: true}, nil
	} else {
		return &model.PraiseCheckOutput{IsCollect: false}, nil
	}
}
