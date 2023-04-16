package model

import "github.com/gogf/gf/v2/os/gtime"

// RoleCreateUpdateBase 创建/修改内容基类
type RoleCreateUpdateBase struct {
	Name string // 用户名
	Desc string // 密码
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建内容返回结果
type RoleCreateOutput struct {
	RoleId uint `json:"role_id"`
}

// RoleUpdateInput 修改内容
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint
}

// RoleGetListInput 获取内容列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// RoleGetListOutput 查询列表结果
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

// RoleSearchInput 搜索列表
type RoleSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RoleSearchOutput 搜索列表结果
type RoleSearchOutput struct {
	List  []RoleSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int         `json:"stats"` // 搜索统计
	Page  int                    `json:"page"`  // 分页码
	Size  int                    `json:"size"`  // 分页数量
	Total int                    `json:"total"` // 数据总数
}

type RoleGetListOutputItem struct {
	Id        uint        `json:"id"`         // 自增ID
	Name      string      `json:"name"`       // 用户名
	Desc      string      `json:"desc"`       //  描述
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type RoleSearchOutputItem struct {
	RoleGetListOutputItem
}

type RoleAddPermissionInput struct {
	RoleId       uint `json:"role_id" v:"min:1#请选择需要添加权限的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要添加权限" dc:"权限id"`
}

type RoleAddPermissionOutput struct {
	Id uint `json:"id"`
}

type RoleDeletePermissionInput struct {
	RoleId       uint `json:"role_id" v:"min:1#请选择需要添加权限的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要添加权限" dc:"权限id"`
}

type RoleDeletePermissionOutput struct{}
