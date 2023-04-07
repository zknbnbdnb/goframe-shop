package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 表结构在position_info中定义
// 虽然这个表不满足3NF,但是考虑手工位经常查且后期会做cache的情况下,就不拆分了做冗余设计

// PositionReq Create 创建
type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"Position" method:"post" summary:"这是个注释"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"       v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name"       v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"       v:"required#商品id不能为空" dc:"商品id"`
	Sort      int    `json:"sort"       dc:"排序"` // 当非必需时，可以不写v:"required#排序不能为空"
}
type PositionRes struct {
	PositionId int `json:"positionId"`
}

// PositionDeleteReq Delete 删除
type PositionDeleteReq struct {
	g.Meta `path:"/position/delete" method:"delete" tags:"手工位图" summary:"删除手工位图接口"`
	Id     uint `v:"min:1#请选择需要删除的手工位图" dc:"手工位图id"`
}
type PositionDeleteRes struct {
	Id uint `json:"id"`
}

// PositionUpdateReq Update 更新
type PositionUpdateReq struct {
	g.Meta    `path:"/position/update/{Id}" method:"post" tags:"手工位图" summary:"修改手工位图接口"`
	Id        uint   `json:"id"         v:"min:1#请选择需要修改的手工位图" dc:"手工位图Id"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"       v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name"       v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"       v:"required#商品id不能为空" dc:"商品id"`
	Sort      int    `json:"sort"       dc:"排序"` // 当非必需时，可以不写v:"required#排序不能为空"
}
type PositionUpdateRes struct {
	Id uint8 `json:"id"`
}

// PositionGetListCommonReq GetListCommon 获取列表
type PositionGetListCommonReq struct {
	g.Meta              `path:"/position/list" method:"get" tags:"手工位图" summary:"手工位图列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type PositionGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
