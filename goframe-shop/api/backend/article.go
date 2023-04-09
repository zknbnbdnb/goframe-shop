package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleAddUpdateBase struct {
	Title   string `json:"title" v:"required#标题不能为空" dc:"标题"`
	Desc    string `json:"desc" v:"required#描述不能为空" dc:"描述"`
	PicUrl  string `json:"pic_url" v:"required#图片地址不能为空" dc:"图片地址"`
	IsAdmin int    `json:"is_admin" v:"required#是否管理员不能为空" dc:"是否管理员" d:"1"` // 默认为1（2为前台用户，1为后台管理员）
	Praise  int    `json:"praise" v:"required#点赞数不能为空" dc:"点赞数"`
	Detail  string `json:"detail" v:"required#详情不能为空" dc:"详情"`
}

// ArticleReq Create 创建
type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"Article" method:"post" summary:"这是个注释"`
	ArticleAddUpdateBase
}
type ArticleRes struct {
	ArticleId int `json:"coupon_id"`
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
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的用户优惠券" dc:"用户优惠券Id"`
	ArticleAddUpdateBase
}
type ArticleUpdateRes struct {
	Id uint8 `json:"id"`
}

// ArticleGetListCommonReq GetListCommon 获取列表
type ArticleGetListCommonReq struct {
	g.Meta              `path:"/article/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表接口"`
	CommonPaginationReq // 翻页配置
}
type ArticleGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
