package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/entity"
)

type CollectionAddInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
}

type CollectionAddOutput struct {
	Id uint
}

type CollectionDeleteInput struct {
	Id       uint
	UserId   uint
	ObjectId int
	Type     uint8
}

type CollectionDeleteOutput struct {
	Id uint
}

type CollectionListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 // 1是商品,2是文章 0是全部
}

// CollectionListOutput 查询列表结果
type CollectionListOutput struct {
	List  []CollectionListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type CollectionListOutputItem struct {
	entity.CollectionInfo
	Goods   GoodsItem   `json:"goods" orm:"with:id=object_id"`   // 通过orm:"with:id=object_id"指定外键静态关联
	Article ArticleItem `json:"article" orm:"with:id=object_id"` // 通过orm:"with:id=object_id"指定外键静态关联
}

type GoodsItem struct {
	g.Meta `orm:"table:goods_info"`
	Id     uint   `json:"id" description:"商品id"`
	Name   string `json:"name" description:"商品名称"`
	PicUrl string `json:"pic_url" description:"商品图片"`
	Price  int    `json:"price" description:"商品价格"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     uint   `json:"id" description:"文章id"`
	Title  string `json:"title" description:"文章标题"`
	PicUrl string `json:"pic_url" description:"文章图片"`
	Desc   string `json:"desc" description:"文章描述"`
}

type CollectionCheckInput struct {
	UserId   uint
	ObjectId int
	Type     uint8
}

type CollectionCheckOutput struct {
	IsCollect bool
}
