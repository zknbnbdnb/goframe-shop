package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// User 角色管理
var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data) // 当你很明确的知道要转什么类型的时候就不用scan了，用scan会损失一部分性能
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{Id: out.Id}, nil
}

func (c *cUser) Info(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	res = &frontend.UserInfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sex = gconv.Int(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Int(ctx.Value(consts.CtxUserStatus))
	return
}

func (c *cUser) UpdatePassword(ctx context.Context, req *frontend.UpdatePasswordReq) (res *frontend.UpdatePasswordRes, err error) {
	data := model.UpdatePasswordInput{}

	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.UpdatePasswordRes{Id: out.Id}, nil
}
