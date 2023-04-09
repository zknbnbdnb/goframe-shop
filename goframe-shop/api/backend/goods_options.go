package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsOptionsAddUpdateBase struct {
	GoodsId uint   `json:"goods_id" v:"required#商品id不能为空" dc:"商品id"`
	PicUrl  string `json:"pic_url" v:"required#图片地址不能为空" dc:"图片地址"`
	Name    string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Price   int    `json:"price" v:"required#价格不能为空" dc:"价格"`
	Stock   int    `json:"stock" v:"required#库存不能为空" dc:"库存"`
}

// GoodsOptionsReq Create 创建
type GoodsOptionsReq struct {
	g.Meta `path:"/goods/options/add" tags:"GoodsOptions" method:"post" summary:"这是个注释"`
	GoodsOptionsAddUpdateBase
}
type GoodsOptionsRes struct {
	GoodsOptionsId int `json:"coupon_id"`
}

// GoodsOptionsDeleteReq Delete 删除
type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/goods/options/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type GoodsOptionsDeleteRes struct {
	Id uint `json:"id"`
}

// GoodsOptionsUpdateReq Update 更新
type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/goods/options/update" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的用户优惠券" dc:"用户优惠券Id"`
	GoodsOptionsAddUpdateBase
}
type GoodsOptionsUpdateRes struct {
	Id uint8 `json:"id"`
}

// GoodsOptionsGetListCommonReq GetListCommon 获取列表
type GoodsOptionsGetListCommonReq struct {
	g.Meta              `path:"/goods/options/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type GoodsOptionsGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
