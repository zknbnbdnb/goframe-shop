package model

import "github.com/gogf/gf/v2/os/gtime"

// UserCouponCreateUpdateBase 创建/修改内容基类
type UserCouponCreateUpdateBase struct {
	UserId   int   // 用户id
	CouponId int   // 优惠券id
	Status   uint8 // 优惠券状态 排序类型(0:未使用, 1:使用, 2:过期)
}

// UserCouponCreateInput 创建内容
type UserCouponCreateInput struct {
	UserCouponCreateUpdateBase
}

// UserCouponCreateOutput 创建内容返回结果
type UserCouponCreateOutput struct {
	UserCouponId int `json:"coupon_id"`
}

// UserCouponUpdateInput 修改内容
type UserCouponUpdateInput struct {
	Id uint
	UserCouponCreateUpdateBase
}

// UserCouponGetListInput 获取内容列表
type UserCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

// UserCouponSearchInput 搜索列表
type UserCouponSearchInput struct {
	Key          string // 关键字
	Type         string // 内容模型
	UserCouponId uint   // 栏目ID
	Page         int    // 分页号码
	Size         int    // 分页数量，最大50
}

// UserCouponSearchOutput 搜索列表结果
type UserCouponSearchOutput struct {
	List  []UserCouponSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int               `json:"stats"` // 搜索统计
	Page  int                          `json:"page"`  // 分页码
	Size  int                          `json:"size"`  // 分页数量
	Total int                          `json:"total"` // 数据总数
}

type UserCouponGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	UserId    int         // 用户id
	CouponId  int         // 优惠券id
	Status    uint8       // 优惠券状态
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type UserCouponSearchOutputItem struct {
	UserCouponGetListOutputItem
}
