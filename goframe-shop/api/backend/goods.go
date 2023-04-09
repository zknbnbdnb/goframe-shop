package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsAddUpdateBase struct {
	PicUrl           string `json:"pic_url" v:"required#图片地址不能为空" dc:"图片地址"`
	Name             string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Price            int    `json:"price" v:"required#价格不能为空" dc:"价格"`
	Level1CategoryId int    `json:"level1_category_id" v:"required#一级分类不能为空" dc:"一级分类"`
	Level2CategoryId int    `json:"level2_category_id" v:"required#二级分类不能为空" dc:"二级分类"`
	Level3CategoryId int    `json:"level3_category_id" v:"required#三级分类不能为空" dc:"三级分类"`
	Brand            string `json:"brand" v:"required#品牌不能为空" dc:"品牌"`
	Stock            int    `json:"stock" v:"required#库存不能为空" dc:"库存"`
	Sale             int    `json:"sale" v:"required#销量不能为空" dc:"销量"`
	Tags             string `json:"tags" v:"required#标签不能为空" dc:"标签"`
	DetailInfo       string `json:"detail_info" v:"required#详情不能为空" dc:"详情"`
}

// GoodsReq Create 创建
type GoodsReq struct {
	g.Meta `path:"/goods/add" tags:"Goods" method:"post" summary:"这是个注释"`
	GoodsAddUpdateBase
}
type GoodsRes struct {
	GoodsId int `json:"coupon_id"`
}

// GoodsDeleteReq Delete 删除
type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type GoodsDeleteRes struct {
	Id uint `json:"id"`
}

// GoodsUpdateReq Update 更新
type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的用户优惠券" dc:"用户优惠券Id"`
	GoodsAddUpdateBase
}
type GoodsUpdateRes struct {
	Id uint8 `json:"id"`
}

// GoodsGetListCommonReq GetListCommon 获取列表
type GoodsGetListCommonReq struct {
	g.Meta              `path:"/goods/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type GoodsGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
