package frontend

import "github.com/gogf/gf/v2/frame/g"

type PraiseAddReq struct {
	g.Meta   `path:"/praise/add" method:"post" summary:"添加点赞" tags:"前台点赞"`
	ObjectId int   `json:"objectId"  description:"对象id" v:"required#对象id不能为空"`
	Type     uint8 `json:"type"      description:"点赞类型：1商品 2文章" v:"in:1,2#点赞类型只能为1或2"` //数据校验,范围约束
}

type PraiseAddRes struct {
	Id uint `json:"id"`
}

type PraiseDeleteReq struct {
	g.Meta   `path:"/praise/delete" method:"post" summary:"删除点赞" tags:"前台点赞"`
	Id       uint  `json:"id" v:"required#id不能为空"`
	UserId   uint  `json:"user_id"    description:"用户id"`
	ObjectId int   `json:"objectId"  description:"对象id"`
	Type     uint8 `json:"type"      description:"点赞类型：1商品 2文章"`
}

type PraiseDeleteRes struct {
	Id uint `json:"id"`
}

type PraiseListReq struct {
	g.Meta `path:"/praise/list" method:"get" summary:"点赞列表" tags:"前台点赞"`
	Type   uint8 `json:"type" description:"点赞类型：1商品 2文章 0全部" v:"in:0,1,2#点赞类型只能为0,1或2"`
	CommonPaginationReq
}

type PraiseListRes struct {
	Page  int         `json:"page" description:"分页号码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"总数"`
	List  interface{} `json:"list" description:"点赞列表"`
}

type PraiseListItem struct {
	Id       uint        `json:"id" v:"required#id不能为空"`
	UserId   uint        `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"objectId"  description:"对象id"`
	Type     uint8       `json:"type"      description:"点赞类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
