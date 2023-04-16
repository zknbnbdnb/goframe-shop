package frontend

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (c *cArticle) Create(ctx context.Context, req *frontend.ArticleReq) (res *frontend.ArticleRes, err error) {
	data := model.ArticleCreateInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleRes{Id: uint(out.Id)}, nil
}

func (c *cArticle) Update(ctx context.Context, req *frontend.ArticleUpdateReq) (res *frontend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return res, err
	}

	// 获取当前用户id
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	data.IsAdmin = consts.ArticleIsUser
	err = service.Article().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleUpdateRes{Id: uint(req.Id)}, nil
}

func (c *cArticle) Detail(ctx context.Context, req *frontend.ArticleDetailReq) (res *frontend.ArticleDetailRes, err error) {
	out, err := service.Article().Detail(ctx, model.ArticleDetailInput{Id: req.Id})
	if err != nil || out == nil {
		return nil, err
	}
	err = gconv.Scan(out, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cArticle) Delete(ctx context.Context, req *frontend.ArticleDeleteReq) (res *frontend.ArticleDeleteRes, err error) {
	data := model.ArticleDeleteInput{
		Id: req.Id,
		ArticleUserAction: model.ArticleUserAction{
			UserId:  gconv.Int(ctx.Value(consts.CtxAdminId)),
			IsAdmin: consts.ArticleIsUser,
		},
	}
	err = service.Article().Delete(ctx, data)
	if err != nil {
		return res, err
	}
	return
}

func (c *cArticle) List(ctx context.Context, req *frontend.ArticleGetListCommonReq) (res *frontend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &frontend.ArticleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cArticle) MyList(ctx context.Context, req *frontend.ArticleGetMyListReq) (res *frontend.ArticleGetListCommonRes, err error) {
	list, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
		ArticleUserAction: model.ArticleUserAction{
			UserId:  gconv.Int(ctx.Value(consts.CtxAdminId)),
			IsAdmin: consts.ArticleIsUser,
		},
	})
	if err != nil {
		return nil, err
	}

	return &frontend.ArticleGetListCommonRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}, err
}
