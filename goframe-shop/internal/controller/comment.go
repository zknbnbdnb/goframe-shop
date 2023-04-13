package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Comment = cComment{}

type cComment struct{}

// Add 添加收藏
func (c *cComment) Add(ctx context.Context, req *frontend.CommentAddReq) (res *frontend.CommentAddRes, err error) {
	data := model.CommentAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentAddRes{Id: out.Id}, nil
}

func (c *cComment) Delete(ctx context.Context, req *frontend.CommentDeleteReq) (res *frontend.CommentDeleteRes, err error) {
	data := model.CommentDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().Delete(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentDeleteRes{Id: out.Id}, nil
}

func (c *cComment) List(ctx context.Context, req *frontend.CommentListReq) (res *frontend.CommentListRes, err error) {
	data := model.CommentListInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().GetList(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentListRes{List: out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total}, nil
}
