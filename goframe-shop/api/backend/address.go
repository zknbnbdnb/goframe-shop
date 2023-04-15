package backend

import "github.com/gogf/gf/v2/frame/g"

type AddressAddUpdateBase struct {
	ParentId int    `json:"parent_id" description:"父级id"`
	Name     string `json:"name" description:"名称"`
	Status   uint8  `json:"status" description:"状态"`
}

// AddAddressReq 新增地址
type AddAddressReq struct {
	g.Meta `path:"/address/add" method:"post" tags:"城市地址" summary:"新增地址"`
	AddressAddUpdateBase
}
type AddAddressRes struct {
	Id int `json:"id" dc:"id"`
}

// UpdateAddressReq 更新地址
type UpdateAddressReq struct {
	g.Meta `path:"/address/update" method:"post" tags:"城市地址" summary:"更新地址"`
	Id     int `json:"id" dc:"id"`
	AddressAddUpdateBase
}
type UpdateAddressRes struct{}

// DeleteAddressReq 删除地址
type DeleteAddressReq struct {
	g.Meta `path:"/address/delete" method:"delete" tags:"城市地址" summary:"删除地址"`
	Id     int `json:"id" dc:"id"`
}

type DeleteAddressRes struct{}

// PageAddressReq 获取地址
type PageAddressReq struct {
	g.Meta `path:"/address/page" method:"get" tags:"城市地址" summary:"获取分页地址"`
	CommonPaginationReq
}
type PageAddressRes struct {
	CommonPaginationRes
}

type CityAddressListReq struct {
	g.Meta `path:"/address/list" tags:"客户端收货地址" method:"post" summary:"客户端省市县区接口"`
}

type CityAddressListRes struct {
	List interface{} `json:"list" description:"列表"`
}
