package model

import (
	"goframe-shop/internal/model/entity"
)

type PraiseAddInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
}

type PraiseAddOutput struct {
	Id uint
}

type PraiseDeleteInput struct {
	Id       uint
	UserId   uint
	ObjectId int
	Type     uint8
}

type PraiseDeleteOutput struct {
	Id uint
}

type PraiseListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 1是商品,2是文章 0是全部
}

// PraiseListOutput 查询列表结果
type PraiseListOutput struct {
	List  []PraiseListOutputItem `json:"list" description:"列表"`
	Page  int                    `json:"page" description:"分页码"`
	Size  int                    `json:"size" description:"分页数量"`
	Total int                    `json:"total" description:"数据总数"`
}

type PraiseListOutputItem struct {
	entity.PraiseInfo
	Goods   GoodsItem   `json:"goods" orm:"with:id=object_id"`   // 通过orm:"with:id=object_id"指定外键静态关联
	Article ArticleItem `json:"article" orm:"with:id=object_id"` // 通过orm:"with:id=object_id"指定外键静态关联
}

type PraiseCheckInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
}

type PraiseCheckOutput struct {
	IsCollect bool
}
