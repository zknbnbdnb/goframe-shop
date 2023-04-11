package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta       `path:"/register" method:"post" tags:"前台注册" summary:"用户注册"`
	Name         string `json:"name"          description:"用户名" v:"required#用户名不能为空"`
	Avatar       string `json:"avatar"        description:"头像"`
	Password     string `json:"password"      description:"密码" v:"password"` //官方文档中的密码验证规则
	UserSalt     string `json:"user_salt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"           description:"1男 2女"`
	Status       int    `json:"status"        description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"          description:"个性签名"`
	SecretAnswer string `json:"secret_answer" description:"密保问题的答案"`
}

type RegisterRes struct {
	Id uint `json:"id"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"前台登录" summary:"用户登录"`
	Name     string `json:"name"          description:"用户名" v:"required#用户名不能为空"`
	Password string `json:"password"      description:"密码" v:"password"`
}

type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserInfoBase
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"前台用户信息" summary:"用户信息"`
}

type UserInfoRes struct {
	UserInfoBase
}

type UserInfoBase struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    int    `json:"sex"`
	Sign   string `json:"sign"`
	Status int    `json:"status"`
}
