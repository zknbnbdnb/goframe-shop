package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}
type LoginDoRes struct {
	//Info interface{} `json:"info" dc:"用户信息"`
	////Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LoginRefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"post"`
}

type LoginRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LoginLogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post"`
}

type LoginLogoutRes struct {
}
