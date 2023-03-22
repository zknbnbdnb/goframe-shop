package model

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 图片链接
	Link   string // 跳转链接
	Sort   int    // 排序
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
