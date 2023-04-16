package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*
	使用restful风格的api设计
	restful协议: 一些预先定义的函数，目的是能够让应用程序或开发人员能够具有访问指定网络资源的能力，而无需关心访问的远吗以及内部的工作机制细节。
*/

// GoodsGetListCommonReq GetListCommon 获取列表
type GoodsGetListCommonReq struct {
	g.Meta              `path:"/goods/list" method:"post" tags:"商品" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"post" tags:"商品" summary:"商品详情接口"`
	Id     uint `json:"id" v:"required#商品id不能为空"`
}

type GoodsDetailRes struct {
	GoodsInfoBase
	Options  []GoodsOptionsBase `json:"options"` //规格 sku
	Comments []CommentBase      `json:"comments"`
}

type GoodsInfoBase struct {
	Id               int         `json:"id"               description:""`
	PicUrl           string      `json:"pic_url"           description:"图片"`
	Name             string      `json:"name"             description:"商品名称"`
	Price            int         `json:"price"            description:"价格 单位分"`
	Level1CategoryId int         `json:"level1_category_id" description:"1级分类id"`
	Level2CategoryId int         `json:"level2_category_id" description:"2级分类id"`
	Level3CategoryId int         `json:"level3_category_id" description:"3级分类id"`
	Brand            string      `json:"brand"            description:"品牌"`
	Stock            int         `json:"stock"            description:"库存"`
	Sale             int         `json:"sale"             description:"销量"`
	Tags             string      `json:"tags"             description:"标签"`
	DetailInfo       string      `json:"detail_info"       description:"商品详情"`
	CreatedAt        *gtime.Time `json:"created_at"        description:""`
}

type GoodsOptionsBase struct {
	Id        int         `json:"id"        description:""`
	GoodsId   int         `json:"goods_id"   description:"商品id"`
	PicUrl    string      `json:"pic_url"    description:"图片"`
	Name      string      `json:"name"      description:"商品名称"`
	Price     int         `json:"price"     description:"价格 单位分"`
	Stock     int         `json:"stock"     description:"库存"`
	CreatedAt *gtime.Time `json:"created_at" description:""`
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
