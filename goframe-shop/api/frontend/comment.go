package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CommentAddReq struct {
	g.Meta   `path:"/comment/add" method:"post" summary:"添加评论" tags:"前台评论"`
	UserId   uint   `json:"user_id"    description:"用户id"`
	ObjectId int    `json:"objectId"  description:"对象id" v:"required#对象id不能为空"`
	Type     uint8  `json:"type"      description:"评论类型：1商品 2文章" v:"in:1,2#评论类型只能为1或2"` //数据校验,范围约束
	ParentId uint   `json:"paren_id"  description:"父级评论id"`
	Content  string `json:"content"   description:"评论内容" v:"required#评论内容不能为空"`
}

type CommentAddRes struct {
	Id uint `json:"id"`
}

type CommentDeleteReq struct {
	g.Meta `path:"/comment/delete" method:"post" summary:"删除评论" tags:"前台评论"`
	Id     uint `json:"id" v:"required#id不能为空"`
	//UserId   uint  `json:"user_id"    description:"用户id"`
	//ObjectId int   `json:"objectId"  description:"对象id"`
	//Type     uint8 `json:"type"      description:"评论类型：1商品 2文章"`
	//评论只需要更具id删除即可,其他模块比如收藏和点赞,需要根据用户id,对象id,类型来删除
}

type CommentDeleteRes struct {
	Id uint `json:"id"`
}

type CommentListReq struct {
	g.Meta `path:"/comment/list" method:"get" summary:"评论列表" tags:"前台评论"`
	Type   uint8 `json:"type" description:"评论类型：1商品 2文章 0全部" v:"in:0,1,2#评论类型只能为0,1或2"`
	CommonPaginationReq
}

type CommentListRes struct {
	Page  int         `json:"page" description:"分页号码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"总数"`
	List  interface{} `json:"list" description:"评论列表"`
}

type CommentListItem struct {
	Id       uint         `json:"id" v:"required#id不能为空"`
	UserId   uint         `json:"user_id"    description:"用户id"`
	User     UserInfoBase `json:"user" description:"用户信息"`
	ObjectId int          `json:"objectId"  description:"对象id"`
	Type     uint8        `json:"type"      description:"评论类型：1商品 2文章"`
	Goods    interface{}  `json:"goods"`
	Article  interface{}  `json:"article"`
}

type CommentBase struct {
	Id        int          `json:"id"        description:""`
	ParentId  int          `json:"parent_id"  description:"父级评论id"`
	UserId    int          `json:"user_id"    description:""`
	User      UserInfoBase `json:"user" dc:"用户信息"`
	ObjectId  int          `json:"object_id"  description:""`
	Type      int          `json:"type"      description:"评论类型：1商品 2文章"`
	Content   string       `json:"content"   description:"评论内容"`
	CreatedAt *gtime.Time  `json:"created_at" description:""`
}
