package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file" method:"post" mime:"multipart/form-data" tags:"工具" dc:"文件上传"`
	File   *ghttp.UploadFile `json:"file"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名"`
	Url  string `json:"url" dc:"文件地址"`
}
