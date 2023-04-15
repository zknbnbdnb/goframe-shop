package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponAddUpdateBase struct {
	Name       string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Price      int    `json:"price" v:"required#优惠券金额不能为空" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" v:"required#商品id不能为空" dc:"支持优惠券商品id,多个用逗号隔开"`
	CategoryId int    `json:"category_id" v:"required#分类id不能为空" dc:"支持优惠券优惠券id"`
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
	g.Meta `path:"/coupon/delete" method:"delete" tags:"优惠券" summary:"删除优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的优惠券" dc:"优惠券id"`
}
type CouponDeleteRes struct {
	Id uint `json:"id"`
}

// CouponUpdateReq Update 更新
type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"post" tags:"优惠券" summary:"修改优惠券接口"`
	Id     uint `json:"id"         v:"min:1#请选择需要修改的优惠券" dc:"优惠券Id"`
	CouponAddUpdateBase
}
type CouponUpdateRes struct{}

// CouponGetListCommonReq GetListCommon 获取列表
type CouponGetListCommonReq struct {
	g.Meta              `path:"/coupon/list" method:"get" tags:"优惠券" summary:"优惠券列表接口"`
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
