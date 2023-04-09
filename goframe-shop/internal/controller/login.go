package controller

import (
	"context"
	"goframe-shop/api/backend"
	"goframe-shop/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

// todo this login method is for jwt
//func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
//	return
//}

func (c *cLogin) RefreshToken(ctx context.Context, req *backend.LoginRefreshTokenReq) (res *backend.LoginRefreshTokenRes, err error) {
	res = &backend.LoginRefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LoginLogoutReq) (res *backend.LoginLogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
