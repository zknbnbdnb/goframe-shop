package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponAddUpdateBase struct {
	Name       string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Price      int    `json:"price" v:"required#优惠券金额不能为空" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" v:"required#商品id不能为空" dc:"支持优惠券商品id,多个用逗号隔开"`
	CategoryId int    `json:"category_id" v:"required#分类id不能为空" dc:"支持优惠券商品分类id"`
}

// CouponReq Create 创建
type CouponReq struct {
	g.Meta `path:"/coupon/add" tags:"Coupon" method:"post" summary:"这是个注释"`
	CouponAddUpdateBase
}
type CouponRes struct {
	CouponId int `json:"coupon_id"`
}

// CouponDeleteReq Delete 删除
type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"商品分类" summary:"删除商品分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}
type CouponDeleteRes struct {
	Id uint `json:"id"`
}

// CouponUpdateReq Update 更新
type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"post" tags:"商品分类" summary:"修改商品分类接口"`
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的商品分类" dc:"商品分类Id"`
	CouponAddUpdateBase
}
type CouponUpdateRes struct {
	Id uint8 `json:"id"`
}

// CouponGetListCommonReq GetListCommon 获取列表
type CouponGetListCommonReq struct {
	g.Meta              `path:"/coupon/list" method:"get" tags:"商品分类" summary:"商品分类列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type CouponGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
