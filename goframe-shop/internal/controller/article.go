package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleReq) (res *backend.ArticleRes, err error) {
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
	return &backend.ArticleRes{ArticleId: out.ArticleId}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data) // todo 当字段较长就可以使用gconv来进行转换
	if err != nil {
		return nil, err
	}
	// 获取当前用户id
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	err = service.Article().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleUpdateRes{Id: req.Id}, nil
}

func (a *cArticle) List(ctx context.Context, req *backend.ArticleGetListCommonReq) (res *backend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.ArticleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
