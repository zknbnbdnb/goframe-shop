package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionCreatUpdateBase struct {
	Name string `json:"name" v:"required#名称必填" dc:"权限名称"`
	Path string `json:"path"  dc:"权限路径"`
}

type PermissionReq struct {
	g.Meta `path:"/permission/add" method:"post" tags:"permission" desc:"添加权限"`
	PermissionCreatUpdateBase
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id" dc:"权限ID"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update" method:"post" tags:"permission" desc:"跟新权限"`
	Id     uint `json:"id" v:"required#ID必填" dc:"权限ID"`
	PermissionCreatUpdateBase
}

type PermissionUpdateRes struct {
	Id uint `json:"id" dc:"权限ID"`
}

// Delete 删除
type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete" method:"delete" tags:"权限" summary:"删除权限接口"`
	Id     uint `v:"min:1#请选择需要删除的权限" dc:"权限id"`
}
type PermissionDeleteRes struct{}

type PermissionGetListCommonReq struct {
	g.Meta              `path:"/permission/list" method:"get" tags:"权限" summary:"权限列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type PermissionGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
