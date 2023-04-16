package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddCartInput struct {
	UserId         uint
	GoodsOptionsId uint
	Count          int
}

type AddCartOutput struct {
	Id uint `json:"id"`
}

type DeleteCartInput struct {
	Id uint
}

type DeleteCartOutput struct {
	Id uint `json:"id"`
}

type ListCartInput struct {
	Page int
	Size int
}

type ListCartOutput struct {
	Page  int
	Size  int
	Total int
	List  []ListCartBase
}

type ListCartBase struct {
	g.Meta         `orm:"table:cart_info"`
	Id             int
	UserId         int
	GoodsOptionsId int
	Count          int
	Ops            ListCartOpsBase `json:"ops" orm:"with:id=goods_options_id"`
	CreatedAt      *gtime.Time
	UpdatedAt      *gtime.Time
	DeletedAt      *gtime.Time
}

type ListCartOpsBase struct {
	g.Meta    `orm:"table:goods_options_info"`
	Id        int
	GoodsId   int
	PicUrl    string
	Name      string
	Price     int
	Stock     int
	Goods     ListCartGoodsBase `json:"goods" orm:"with:id=goods_id"`
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
	DeletedAt *gtime.Time
}

type ListCartGoodsBase struct {
	g.Meta           `orm:"table:goods_info"`
	Id               int
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	CouponId         int
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
	CreatedAt        *gtime.Time
	UpdatedAt        *gtime.Time
	DeletedAt        *gtime.Time
}
