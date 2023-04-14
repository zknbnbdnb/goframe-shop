package model

import "github.com/gogf/gf/v2/frame/g"

type RegisterInput struct {
	Name         string
	Avatar       string
	Password     string
	UserSalt     string
	Sex          int
	Status       int
	Sign         string
	SecretAnswer string
}

type RegisterOutput struct {
	Id uint
}

type LoginInput struct {
	Name     string
	Password string
}

type LoginOutput struct {
}

type UpdatePasswordInput struct {
	SecretAnswer string
	Password     string
	UserSalt     string
}

type UpdatePasswordOutput struct {
	Id uint
}

type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    int    `json:"sex"`
	Sign   string `json:"sign"`
	Status int    `json:"status"`
}
