package model

import "goframe-shop/internal/model/entity"

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 图片链接
	Link   string // 跳转链接
	Sort   int    // 排序
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

// RotationUpdateInput 修改内容
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// RotationSearchInput 搜索列表
type RotationSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationSearchOutput 搜索列表结果
type RotationSearchOutput struct {
	List  []RotationSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type RotationGetListOutputItem struct {
	entity.RotationInfo
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}
