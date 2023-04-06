package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/entity"
	"time"
)

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// this struct for jwt
type LoginDoRes struct {
	//Info interface{} `json:"info" dc:"用户信息"`
	////Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// this struct for gtoken
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}

type LoginRefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type LoginRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LoginLogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LoginLogoutRes struct {
}
