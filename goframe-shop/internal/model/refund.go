package model

import "goframe-shop/internal/model/entity"

type AddRefundInput struct {
	Number  string `json:"number"    description:"售后订单号"`
	OrderId int    `json:"orderId"   description:"订单id"`
	GoodsId int    `json:"goodsId"   description:"要售后的商品id"`
	Reason  string `json:"reason"    description:"退款原因"`
	Status  int    `json:"status"    description:"状态 1待处理 2同意退款 3拒绝退款"`
	UserId  int    `json:"userId"    description:"用户id"`
}

type AddRefundOutput struct {
	Id uint `json:"id"`
}

type ListRefundInput struct {
	Page   int // 分页号码
	Size   int // 分页数量，最大50
	UserId int
}

type ListRefundOutput struct {
	List  []ListRefundOutputItem
	Page  int
	Size  int
	Total int
}

type ListRefundOutputItem struct {
	entity.RefundInfo
}

type DetailRefundInput struct {
	Id uint
}

type DetailRefundOutput struct {
	entity.RefundInfo
}
