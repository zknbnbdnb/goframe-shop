package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" tags:"role" desc:"添加角色"`
	Name   string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleRes struct {
	RoleId uint `json:"role_id" dc:"角色ID"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" tags:"role" desc:"跟新角色"`
	Id     uint   `json:"id" v:"required#ID必填" dc:"角色ID"`
	Name   string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleUpdateRes struct {
	Id uint `json:"id" dc:"角色ID"`
}

// RoleDeleteReq Delete 删除
type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta              `path:"/role/list" method:"get" tags:"角色" summary:"角色列表接口"`
	CommonPaginationReq // 翻页配置
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type RoleAddPermissionReq struct {
	g.Meta       `path:"/role/add/permission" method:"post" tags:"角色" summary:"角色添加权限接口"`
	RoleId       uint `json:"role_id" v:"min:1#请选择需要添加权限的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要添加权限" dc:"权限id"`
}

type RoleAddPermissionRes struct {
	Id uint `json:"id" dc:"角色id"`
}

type RoleDeletePermissionReq struct {
	g.Meta       `path:"/role/delete/permission" method:"delete" tags:"角色" summary:"角色删除权限接口"`
	RoleId       uint `json:"role_id" v:"min:1#请选择需要删除权限的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要删除权限" dc:"权限id"`
}

type RoleDeletePermissionRes struct{}
