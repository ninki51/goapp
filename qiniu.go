package main

import (
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"os"
)

var (
	accessKey = os.Getenv("your key")
	secretKey = os.Getenv("your secret")
	bucket    = os.Getenv("your 空间")
)

func qiniuToken() string{
	// 设置上传凭证有效期
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	putPolicy.Expires = 7200 //示例2小时有效期

	upToken := putPolicy.UploadToken(mac)
	//fmt.Println(upToken)

	//json.NewEncoder(getCors(w, r)).Encode(res)
	return upToken
}

func main(){
	token := qiniuToken()
	fmt.Println(token)
}
