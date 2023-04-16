package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/model/entity"
)

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

type GoodsGetListOutputItem struct {
	entity.GoodsInfo // todo 字段长直接使用这个
}

// GoodsDetailInput 获取内容详情
type GoodsDetailInput struct {
	Id uint
}

// GoodsDetailOutput 获取内容详情返回结果
type GoodsDetailOutput struct {
	do.GoodsInfo
	Options  []*do.GoodsOptionsInfo `orm:"with:goods_id=id"` //规格 sku
	Comments []*CommentBase         `orm:"with:object_id=id, where:type=1"`
}

type BaseGoodsColumns struct {
	g.Meta     `orm:"table:goods_info"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Brand      string `json:"brand"`
	Tags       string `json:"tags"`
	PicUrl     string `json:"pic_url"`
	DetailInfo string `json:"detail_info"`
}
