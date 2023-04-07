package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponAddUpdateBase struct {
	UserId   int   `json:"user_id" v:"required#用户id不能为空" dc:"用户id"`
	CouponId int   `json:"coupon_id" v:"required#优惠券id不能为空" dc:"优惠券id"`
	Status   uint8 `json:"status" v:"required#状态不能为空" dc:"优惠券状态"`
}

// UserCouponReq Create 创建
type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" tags:"UserCoupon" method:"post" summary:"这是个注释"`
	UserCouponAddUpdateBase
}
type UserCouponRes struct {
	UserCouponId int `json:"coupon_id"`
}

// UserCouponDeleteReq Delete 删除
type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type UserCouponDeleteRes struct {
	Id uint `json:"id"`
}

// UserCouponUpdateReq Update 更新
type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的用户优惠券" dc:"用户优惠券Id"`
	UserCouponAddUpdateBase
}
type UserCouponUpdateRes struct {
	Id uint8 `json:"id"`
}

// UserCouponGetListCommonReq GetListCommon 获取列表
type UserCouponGetListCommonReq struct {
	g.Meta              `path:"/user/coupon/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type UserCouponGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
