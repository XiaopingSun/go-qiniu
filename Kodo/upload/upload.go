package main

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = "bjtWBQXrcxgo7HWwlC_bgHg81j352_GhgBGZPeOW"
	secretKey = "pCav6rTslxP2SIFg0XJmAw53D9PjWEcuYWVdUqAf"
	bucket    = "pursue"
)

func main() {
	formUpload()
}

// 表单
func formUpload() {
	localFile := "/Users/workspace_sun/Desktop/Document/2channel.mp4"
	key := "huidiaoceshi.mp4"
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
		CallbackURL:      "http://node.ijemy.com/qncback",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}

	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println("uploadToken:", upToken)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	////设置代理
	//proxyURL := "http://localhost:8888"
	//proxyURI, _ := url.Parse(proxyURL)

	////绑定网卡
	//nicIP := "100.100.33.138"
	//dialer := &net.Dialer{
	//	LocalAddr: &net.TCPAddr{
	//		IP: net.ParseIP(nicIP),
	//	},
	//}

	////构建代理client对象
	//client := http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyURI),
	//		Dial:  dialer.Dial,
	//	},
	//}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploaderEx(&cfg, nil)
	ret := storage.PutRet{}
	//// 可选配置
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}
	//putExtra.NoCrc32Check = true
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

