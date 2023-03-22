package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"这是个注释"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"       v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"       dc:"排序"` // 当非必需时，可以不写v:"required#排序不能为空"
}
type RotationRes struct {
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotationId"`
}
