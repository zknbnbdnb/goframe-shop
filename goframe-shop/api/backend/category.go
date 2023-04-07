package backend

import "github.com/gogf/gf/v2/frame/g"

type CategoryAddUpdateBase struct {
	ParentId uint   `json:"parent_id"  dc:"父级id"`
	Name     string `json:"name" v:"required#名称不能为空" dc:"名称"`
	PicUrl   string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Level    uint8  `json:"level"  dc:"等级"`
	Sort     uint8  `json:"sort" dc:"排序"` // 当非必需时，可以不写v:"required#排序不能为空"
}

// CategoryReq Create 创建
type CategoryReq struct {
	g.Meta `path:"/category/add" tags:"Category" method:"post" summary:"这是个注释"`
	CategoryAddUpdateBase
}
type CategoryRes struct {
	CategoryId int `json:"category_isd"`
}

// CategoryDeleteReq Delete 删除
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"商品分类" summary:"删除商品分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}
type CategoryDeleteRes struct {
	Id uint `json:"id"`
}

// CategoryUpdateReq Update 更新
type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" method:"post" tags:"商品分类" summary:"修改商品分类接口"`
	Id     uint8 `json:"id"         v:"min:1#请选择需要修改的商品分类" dc:"商品分类Id"`
	CategoryAddUpdateBase
}
type CategoryUpdateRes struct {
	Id uint8 `json:"id"`
}

// CategoryGetListCommonReq GetListCommon 获取列表
type CategoryGetListCommonReq struct {
	g.Meta              `path:"/category/list" method:"get" tags:"商品分类" summary:"商品分类列表接口"`
	Sort                int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq     // 翻页配置
}
type CategoryGetListCommonRes struct {
	// todo 以为要做前后端分离，使用不返回html
	List  interface{} `json:"list" description:"列表"` // 如果这里的list不做空接口设计那么每个api在从新设计的时候都要动态修改很不灵活
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
