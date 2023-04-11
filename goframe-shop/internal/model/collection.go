package model

type CollectionAddInput struct {
	UserId   uint
	ObjectId int
	Type     int
}

type CollectionAddOutput struct {
	Id uint
}

type CollectionDeleteInput struct {
	Id       uint
	UserId   uint
	ObjectId int
	Type     int
}

type CollectionDeleteOutput struct {
	Id uint
}
