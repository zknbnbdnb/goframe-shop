package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" tags:"role" desc:"添加角色"`
	Name   string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id" dc:"角色ID"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" tags:"role" desc:"跟新角色"`
	Id     uint   `json:"id" v:"required#ID必填" dc:"角色ID"`
	Name   string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleUpdateRes struct {
	Id uint `json:"id" dc:"角色ID"`
}

// Delete 删除
type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta              `path:"/backend/role/list" method:"get" tags:"角色" summary:"角色列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type RoleGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
