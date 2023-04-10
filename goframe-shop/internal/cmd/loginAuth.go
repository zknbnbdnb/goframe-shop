package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/utility"
	"goframe-shop/utility/response"
	"strconv"
)

// StartBackendGToken 管理后台登录
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:        consts.CacheModeRedis, // 缓存模式，1为内存模式，2为redis模式
		ServerName:       consts.BackendServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFuncBackend,      // 小写只在内部调用
		LoginAfterFunc:   loginAfterFuncBackend, // 小写只在内部调用
		LogoutPath:       "/logout",
		AuthPaths:        g.SliceStr{"/admin/info"},                                 // 需要拦截的路径
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不需要拦截的路径
		AuthAfterFunc:    authAfterFunc,                                             // 小写只在内部调用
		MultiLogin:       consts.BackendMultiLogin,
	}
	err = gfAdminToken.Start()
	return gfAdminToken, err
}

// StartFrontendGToken 管理前台登录
func StartFrontendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis, // 缓存模式，1为内存模式，2为redis模式
		ServerName:      consts.BackendServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFuncFrontend,      // 小写只在内部调用
		LoginAfterFunc:  loginAfterFuncFrontend, // 小写只在内部调用
		LogoutPath:      "/logout",
		//AuthPaths:        g.SliceStr{"/admin/info"},                                 // 需要拦截的路径
		//AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不需要拦截的路径
		AuthAfterFunc: authAfterFunc, // 小写只在内部调用
		MultiLogin:    consts.FrontendMultiLogin,
	}
	//err = gfAdminToken.Start() 不使用全局启动
	return gfAdminToken, err
}

// loginFuncBackend 后台登录鉴权函数
func loginFuncBackend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	// 唯一标识，扩展参数user data
	return consts.GTokenBackendPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// loginFuncFrontend 前台登录鉴权函数
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	// 唯一标识，扩展参数user data
	return consts.GTokenFrontendPrefix + strconv.Itoa(userInfo.Id), userInfo
}

// loginAfterFuncBackend 后台登录后置函数
func loginAfterFuncBackend(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenBackendPrefix)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

// loginAfterFuncFrontend 前台登录后置函数
func loginAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenFrontendPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := &frontend.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn,
			Name:     userInfo.Name,
			Avatar:   userInfo.Avatar,
			Sex:      userInfo.Sex,
			Sign:     userInfo.Sign,
			Status:   userInfo.Status,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
