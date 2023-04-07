package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryCreateUpdateBase 创建/修改内容基类
type CategoryCreateUpdateBase struct {
	ParentId uint   // 父级ID
	Name     string // 名称
	PicUrl   string // 图片链接
	Level    uint8  // 等级
	Sort     uint8  // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
}

// CategoryCreateInput 创建内容
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
}

// CategoryCreateOutput 创建内容返回结果
type CategoryCreateOutput struct {
	CategoryId int `json:"category_id"`
}

// CategoryUpdateInput 修改内容
type CategoryUpdateInput struct {
	Id uint
	CategoryCreateUpdateBase
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// CategorySearchInput 搜索列表
type CategorySearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategorySearchOutput 搜索列表结果
type CategorySearchOutput struct {
	List  []CategorySearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type CategoryGetListOutputItem struct {
	//Category *CategoryListItem `json:"category"`
	//Category *CategoryListCategoryItem `json:"category"`
	//User     *CategoryListUserItem      `json:"user"`
	Id        uint        `json:"id"`         // 自增ID
	ParentId  int         `json:"parent_id"`  // 父级ID
	Name      string      `json:"link"`       // 跳转链接
	PicUrl    string      `json:"pic_url"`    // 图片链接
	Level     uint8       `json:"level"`      // 等级
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CategorySearchOutputItem struct {
	CategoryGetListOutputItem
}
