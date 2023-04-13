package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/model/entity"
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
	// todo 以为要做前后端分离，使用不返回html
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
	entity.GoodsInfo
	Options   interface{}
	Comments  interface{}
	IsComment bool
}
