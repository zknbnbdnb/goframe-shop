package model

import (
	"goframe-shop/internal/model/entity"
)

// CouponCreateUpdateBase 创建/修改内容基类
type CouponCreateUpdateBase struct {
	Name       string // 名称
	Price      int    // 优惠券金额
	GoodsIds   string // 可用优惠券的商品id
	CategoryId int    // 可用优惠券的分类id
}

// CouponCreateInput 创建内容
type CouponCreateInput struct {
	CouponCreateUpdateBase
}

// CouponCreateOutput 创建内容返回结果
type CouponCreateOutput struct {
	CouponId int `json:"coupon_id"`
}

// CouponUpdateInput 修改内容
type CouponUpdateInput struct {
	Id uint
	CouponCreateUpdateBase
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

// CouponSearchInput 搜索列表
type CouponSearchInput struct {
	Key      string // 关键字
	Type     string // 内容模型
	CouponId uint   // 栏目ID
	Page     int    // 分页号码
	Size     int    // 分页数量，最大50
	Sort     int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponSearchOutput 搜索列表结果
type CouponSearchOutput struct {
	List  []CouponSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int           `json:"stats"` // 搜索统计
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

type CouponGetListOutputItem struct {
	entity.CouponInfo
}

type CouponSearchOutputItem struct {
	CouponGetListOutputItem
}
