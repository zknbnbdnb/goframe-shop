package frontend

import "github.com/gogf/gf/v2/frame/g"

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" method:"post" summary:"添加收藏" tags:"前台收藏"`
	ObjectId int   `json:"objectId"  description:"对象id" v:"required#对象id不能为空"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2#收藏类型只能为1或2"` //数据校验,范围约束
}

type CollectionAddRes struct {
	Id uint `json:"id"`
}

type CollectionDeleteReq struct {
	g.Meta   `path:"/collection/delete" method:"post" summary:"删除收藏" tags:"前台收藏"`
	Id       uint  `json:"id" v:"required#id不能为空"`
	UserId   uint  `json:"user_id"    description:"用户id"`
	ObjectId int   `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type CollectionDeleteRes struct {
	Id uint `json:"id"`
}

type CollectionListReq struct {
	g.Meta `path:"/collection/list" method:"get" summary:"收藏列表" tags:"前台收藏"`
	Type   uint8 `json:"type" description:"收藏类型：1商品 2文章 0全部" v:"in:0,1,2#收藏类型只能为0,1或2"`
	CommonPaginationReq
}

type CollectionListRes struct {
	Page  int         `json:"page" description:"分页号码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"总数"`
	List  interface{} `json:"list" description:"收藏列表"`
}

type CollectionListItem struct {
	Id       uint        `json:"id" v:"required#id不能为空"`
	UserId   uint        `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"objectId"  description:"对象id"`
	Type     uint8       `json:"type"      description:"收藏类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
