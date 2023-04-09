package model

import "goframe-shop/internal/model/entity"

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	GoodsId int `json:"coupon_id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	Id uint
	GoodsCreateUpdateBase
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// GoodsSearchInput 搜索列表
type GoodsSearchInput struct {
	Key     string // 关键字
	Type    string // 内容模型
	GoodsId uint   // 栏目ID
	Page    int    // 分页号码
	Size    int    // 分页数量，最大50
}

// GoodsSearchOutput 搜索列表结果
type GoodsSearchOutput struct {
	List  []GoodsSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type GoodsGetListOutputItem struct {
	entity.GoodsInfo // todo 字段长直接使用这个
}

type GoodsSearchOutputItem struct {
	GoodsGetListOutputItem
}
