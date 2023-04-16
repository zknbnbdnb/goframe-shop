package model

import "goframe-shop/internal/model/entity"

// ArticleCreateUpdateBase 创建/修改内容基类
type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    string
	PicUrl  string
	IsAdmin int
	Praise  int
	Detail  string
}

// ArticleCreateInput 创建内容
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
}

// ArticleCreateOutput 创建内容返回结果
type ArticleCreateOutput struct {
	Id uint
}

// ArticleUpdateInput 修改内容
type ArticleUpdateInput struct {
	Id uint
	ArticleCreateUpdateBase
}

type ArticleUpdateOutput struct{}

// ArticleGetListInput 获取内容列表
type ArticleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	ArticleUserAction
}

// ArticleGetListOutput 查询列表结果
type ArticleGetListOutput struct {
	List  []ArticleGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type ArticleGetListOutputItem struct {
	entity.ArticleInfo // todo 字段长直接使用这个
}

type ArticleDetailInput struct {
	Id uint
}

type ArticleDetailOutput struct {
	entity.ArticleInfo
}

type ArticleDeleteInput struct {
	Id uint
	ArticleUserAction
}

type ArticleDeleteOutput struct{}

type ArticleUserAction struct {
	UserId  int // 用户
	IsAdmin int // 是否是管理员
}
