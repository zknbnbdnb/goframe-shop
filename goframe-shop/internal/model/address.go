package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/api/backend"
)

type AddressBase struct {
	ParentId int
	Name     string
	Status   uint8
}

type AddAddressInput struct {
	AddressBase
}

type AddAddressOutput struct {
	Id int
}

type DeleteAddressInput struct {
	Id int
}

type DeleteAddressOutput struct{}

type UpdateAddressInput struct {
	Id int
	AddressBase
}

type UpdateAddressOutput struct{}

type PageAddressInput struct {
	backend.CommonPaginationReq
}

type PageAddressOutput struct {
	backend.CommonPaginationRes
}

type CityAddressListOutput struct {
	List []CityAddressListOutputItem
}

type CityAddressListOutputItem struct {
	g.Meta   `orm:"table:address_info"`
	Id       int
	ParentId int
	Name     string
	Status   int
	Children []CityAddressListOutputItem `orm:"with:parent_id=id"`
}
