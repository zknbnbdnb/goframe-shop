package model

import (
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/model/entity"
)

type OrderListInput struct {
	Page           int    // 分页号码
	Size           int    // 分页数量，最大50
	Number         string `json:"number"           description:"订单编号"`
	UserId         int    `json:"userId"           description:"用户id"`
	PayType        int    `json:"payType"          description:"支付方式 1微信 2支付宝 3云闪付"`
	PayAtGte       string `json:"PayAtGte"            description:"支付时间>="`
	PayAtLte       string `json:"PayAtLte"            description:"支付时间<="`
	Status         int    `json:"status"           description:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价"`
	ConsigneePhone string `json:"consigneePhone"   description:"收货人手机号"`
	PriceGte       int    `json:"PriceGte"            description:"订单金额>="`
	PriceLte       int    `json:"PriceLte"            description:"订单金额<="`
	DateGte        string `json:"DateGte"            description:"创建时间>="`
	DateLte        string `json:"DateLte"            description:"创建时间<="`
}

type OrderListOutput struct {
	List  []OrderListOutputItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type OrderListOutputItem struct {
	entity.OrderInfo
}

type OrderDetailInput struct {
	Id uint
}

type OrderDetailOutput struct {
	do.OrderInfo
	GoodsInfo []*do.OrderGoodsInfo `orm:"with:order_id=id"`
}

type OrderAddInput struct {
	UserId           uint
	Number           string
	Remark           string `description:"备注"`
	Price            int    `description:"订单金额 单位分"`
	CouponPrice      int    `description:"优惠券金额 单位分"`
	ActualPrice      int    `description:"实际支付金额 单位分"`
	ConsigneeName    string `description:"收货人姓名"`
	ConsigneePhone   string `description:"收货人手机号"`
	ConsigneeAddress string `description:"收货人详细地址"`
	//商品订单维度
	OrderAddGoodsInfos []*OrderAddGoodsInfo
}

type OrderAddGoodsInfo struct {
	Id             int
	OrderId        int
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Remark         string
	Price          int
	CouponPrice    int
	ActualPrice    int
}

type OrderAddOutput struct {
	Id uint
}
