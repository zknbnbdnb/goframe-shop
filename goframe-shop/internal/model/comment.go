package model

import (
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/model/entity"
)

type CommentAddInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
	ParentId uint
	Content  string
}

type CommentAddOutput struct {
	Id uint
}

type CommentDeleteInput struct {
	Id uint
}

type CommentDeleteOutput struct {
	Id uint
}

type CommentListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 1是商品,2是文章 0是全部
}

// CommentListOutput 查询列表结果
type CommentListOutput struct {
	List  []CommentListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type CommentListOutputItem struct {
	entity.CommentInfo
	Goods   GoodsItem   `json:"goods" orm:"with:id=object_id"`   // 通过orm:"with:id=object_id"指定外键静态关联
	Article ArticleItem `json:"article" orm:"with:id=object_id"` // 通过orm:"with:id=object_id"指定外键静态关联
}

type CommentCheckInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
}

type CommentCheckOutput struct {
	IsCollect bool
}

type CommentBase struct {
	do.CommentInfo
	User UserInfoBase `json:"user" orm:"with:id=user_id"` // 通过orm:"with:id=user_id"指定外键静态关联
}
