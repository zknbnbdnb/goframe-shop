package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImgToCloudReq struct {
	g.Meta `path:"/backend/upload/cloud" method:"post" mime:"multipart/form-data" tags:"工具" dc:"上传图片到云"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"上传的文件"`
}

type UploadImgToCloudRes struct {
	Url string `json:"url" dc:"图片地址"`
}
