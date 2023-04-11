package frontend

import "github.com/gogf/gf/v2/frame/g"

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" method:"post" summary:"添加收藏" tags:"前台收藏"`
	UserId   uint `json:"user_id"    description:"用户id" v:"required#用户id不能为空"`
	ObjectId int  `json:"objectId"  description:"对象id" v:"required#对象id不能为空"`
	Type     int  `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2#收藏类型只能为1或2"` //数据校验,范围约束
}

type CollectionAddRes struct {
	Id uint `json:"id"`
}

type CollectionDeleteReq struct {
	g.Meta   `path:"/collection/delete" method:"post" summary:"删除收藏" tags:"前台收藏"`
	Id       uint `json:"id" v:"required#id不能为空"`
	UserId   uint `json:"user_id"    description:"用户id"`
	ObjectId int  `json:"objectId"  description:"对象id"`
	Type     int  `json:"type"      description:"收藏类型：1商品 2文章"`
}

type CollectionDeleteRes struct {
	Id uint `json:"id"`
}
