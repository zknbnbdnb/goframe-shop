package model

type AddOrderGoodsCommentsInput struct {
	OrderId        uint
	GoodsId        uint
	GoodsOptionsId uint
	ParentId       uint
	Content        string
}

type AddOrderGoodsCommentsOutput struct {
	Id uint
}

type DeleteOrderGoodsCommentsInput struct {
	Id uint
}

type DeleteOrderGoodsCommentsOutput struct {
	Id uint
}
