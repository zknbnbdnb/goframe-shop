package upload

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
)

func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	// 定义上传目录
	dirPath := "/tmp/"
	name, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}
	// 定义本地文件路径
	localFile := dirPath + name
	// 获取七牛云配置参数
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	// 对接七牛云的sdk
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 生成token
	mac := qbox.NewMac(accessKey, secretKey)
	// 上传token
	upToken := putPolicy.UploadToken(mac)
	// 七牛云上传配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	// 上传结果的结构体
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	// 七牛云表单上传
	key := name
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)
	// 删除临时文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	return
}
