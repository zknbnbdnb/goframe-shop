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
	RotationId int `json:"rotationId"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}
