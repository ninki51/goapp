package main

import (
	"regexp"
	"fmt"
)

const email  = "邮箱 2556792125@qq.com"

func main(){
	re := regexp.MustCompile(`[a-zA-z0-9]+@.+\..+`)
	match := re.FindString(email)
	fmt.Println(match)
}
