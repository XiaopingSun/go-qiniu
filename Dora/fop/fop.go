
package main

import (
	//"encoding/base64"
	"fmt"
	//"os"
	"strings"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = "bjtWBQXrcxgo7HWwlC_bgHg81j352_GhgBGZPeOW"
	secretKey = "pCav6rTslxP2SIFg0XJmAw53D9PjWEcuYWVdUqAf"
	bucket    = "pursue"
	pipeline = "ts-test-pipleline"
)

func main() {
	mac := auth.New(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	operationManager := storage.NewOperationManager(mac, &cfg)
	key := "2channel.mp4"
	saveBucket := bucket
	// 处理指令集合
	fopAvthumb := fmt.Sprintf("avthumb/mp4/s/480x320/vb/500k|saveas/%s",
		storage.EncodedEntry(saveBucket, "dora_test.mp4"))
	//fopVframe := fmt.Sprintf("vframe/jpg/offset/10|saveas/%s",
	//	storage.EncodedEntry(saveBucket, "pfop_test_qiniu.jpg"))
	//fopVsample := fmt.Sprintf("vsample/jpg/interval/20/pattern/%s",
	//	base64.URLEncoding.EncodeToString([]byte("pfop_test_$(count).jpg")))
	fopBatch := []string{fopAvthumb}
	fops := strings.Join(fopBatch, ";")
	// 强制重新执行数据处理任务
	force := true
	// 数据处理指令全部完成之后，通知该地址
	notifyURL := "http://node.ijemy.com/qncback"
	// 数据处理的私有队列，必须指定以保障处理速度
	persistentId, err := operationManager.Pfop(bucket, key, fops, pipeline, notifyURL, force)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(persistentId)
}