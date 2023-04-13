package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) Add(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return out, err
	}
	return &model.CommentAddOutput{Id: uint(id)}, nil
}

func (s *sComment) Delete(ctx context.Context, in model.CommentDeleteInput) (out *model.CommentDeleteOutput, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: gconv.Uint(ctx.Value(consts.CtxUserId)),
	}
	if _, err = dao.CommentInfo.Ctx(ctx).Where(condition).Delete(); err != nil {
		return out, err
	}
	return &model.CommentDeleteOutput{Id: in.Id}, nil
}

func (s *sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	m := dao.CommentInfo.Ctx(ctx)
	err = gconv.Struct(in, &out)
	if err != nil {
		return out, err
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	if in.Type != 0 {
		listModel = listModel.Ctx(ctx).Where(dao.CommentInfo.Columns().Type, in.Type)
	}

	out.Total, err = listModel.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if in.Type == consts.CommentTypeGoods {
		err = listModel.With(model.GoodsItem{}).Scan(&out.List)
		if err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
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

func CommentCount(ctx context.Context, objectId uint, commentType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     commentType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

func CommentCheck(ctx context.Context, in model.CommentCheckInput) (out *model.CommentCheckOutput, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return out, err
	}
	if count > 0 {
		return &model.CommentCheckOutput{IsCollect: true}, nil
	} else {
		return &model.CommentCheckOutput{IsCollect: false}, nil
	}
}
