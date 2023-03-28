package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 表结构在admin_info中定义

// Create 创建
type AdminReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"这是个注释"`
	Name     string `json:"name" v:"required#用户名名称不能为空" dc:"用户名名称"`
	Password string `json:"pass" v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids" dc:"角色id"`
	IsAdmin  int    `json:"is_admin" dc:"是否是超级管理员"`
}
type AdminRes struct {
	AdminId int `json:"admin_id"`
}

// Delete 删除
type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员接口"`
	Id     uint `v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}
type AdminDeleteRes struct{}

// Update 更新
type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       uint   `json:"id"         v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name" v:"required#用户名名称不能为空" dc:"用户名名称"`
	Password string `json:"pass" v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids" dc:"角色id"`
	IsAdmin  int    `json:"is_admin" dc:"是否是超级管理员"`
}
type AdminUpdateRes struct {
	Id uint `json:"id"`
}

type AdminGetListCommonReq struct {
	g.Meta              `path:"/backend/admin/list" method:"get" tags:"管理员" summary:"管理员列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type AdminGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}