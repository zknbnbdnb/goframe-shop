package model

import (
	"goframe-shop/internal/model/entity"
)

// AdminCreateUpdateBase 创建/修改内容基类
type AdminCreateUpdateBase struct {
	Name     string // 用户名
	Password string // 密码
	RoleIds  string // 角色ID
	UserSalt string // 用户盐
	IsAdmin  int    // 是否是管理员
}

// AdminCreateInput 创建内容
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// AdminCreateOutput 创建内容返回结果
type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

// AdminUpdateInput 修改内容
type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

// AdminGetListInput 获取内容列表
type AdminGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// AdminGetListOutput 查询列表结果
type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// AdminSearchInput 搜索列表
type AdminSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// AdminSearchOutput 搜索列表结果
type AdminSearchOutput struct {
	List  []AdminSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type AdminGetListOutputItem struct {
	entity.AdminInfo
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}
