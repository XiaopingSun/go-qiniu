package main

import (
"fmt"
"os"

"github.com/qiniu/go-sdk/v7/auth"
"github.com/qiniu/go-sdk/v7/storage"
"strings"
)

var (
	accessKey = "VP6AUlErFhr5RsDldqBSUy671lWe27_bwSSZqaqp"
	secretKey = "BCR--MFdzy7EmJvEzbDhjyaP-E6pGFBFdvV_MV-e"
	bucket    = "langu-log"
)

func main() {
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	//列举所有文件
	prefix, delimiter, marker := "", "", ""
	entries, err := bucketManager.ListBucket(bucket, prefix, delimiter, marker)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListBucket: %v\n", err)
		os.Exit(1)
	}
	for listItem := range entries {
		fmt.Println(listItem.Marker)
		fmt.Println(listItem.Item)
		fmt.Println(listItem.Dir)
		fmt.Println(strings.Repeat("-", 30))
	}
}
