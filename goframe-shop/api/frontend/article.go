package frontend

import "github.com/gogf/gf/v2/frame/g"

type ArticleAddUpdateBase struct {
	Title  string `json:"title" v:"required#标题不能为空" dc:"标题"`
	Desc   string `json:"desc" v:"required#描述不能为空" dc:"描述"`
	PicUrl string `json:"pic_url" v:"required#图片地址不能为空" dc:"图片地址"`
	Detail string `json:"detail" v:"required#详情不能为空" dc:"详情"`
}

// ArticleReq Create 创建
type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"Article" method:"post" summary:"前端文章"`
	ArticleAddUpdateBase
}
type ArticleRes struct {
	Id uint `json:"id"`
}

// ArticleDeleteReq Delete 删除
type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type ArticleDeleteRes struct {
	Id uint `json:"id"`
}

// ArticleUpdateReq Update 更新
type ArticleUpdateReq struct {
	g.Meta `path:"/article/update" method:"post" tags:"用户优惠券" summary:"修改用户优惠券接口"`
	Id     uint `json:"id"         v:"min:1#请选择需要修改的用户优惠券" dc:"用户优惠券Id"`
	ArticleAddUpdateBase
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}

type ArticleDetailReq struct {
	g.Meta `path:"/article/detail" method:"get" tags:"前端文章" summary:"文章详情"`
	Id     uint `json:"id"`
}

type ArticleDetailRes struct {
	Id        int    `json:"id"        description:""`
	UserId    int    `json:"userId"    description:"作者id"`
	Title     string `json:"title"     description:"标题"`
	Desc      string `json:"desc"      description:"摘要"`
	PicUrl    string `json:"picUrl"    description:"封面图"`
	IsAdmin   int    `json:"isAdmin"   description:"1后台管理员发布 2前台用户发布"`
	Praise    int    `json:"praise"    description:"点赞数"`
	Detail    string `json:"detail"    description:"文章详情"`
	CreatedAt string `json:"createdAt" description:""`
	UpdatedAt string `json:"updatedAt" description:""`
}

// ArticleGetListCommonReq GetListCommon 获取列表
type ArticleGetListCommonReq struct {
	g.Meta              `path:"/article/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type ArticleGetMyListReq struct {
	g.Meta `path:"/article/my" method:"get" tags:"前端文章" summary:"我的文章接口"`
	CommonPaginationReq
}
