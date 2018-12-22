package test

import (
	"fmt"
	"time"
)


/**
测试安装自己的包
 */
func Test(){
	fmt.Printf("Note: I come from \"test\" pkg!@time:%v",time.Now().Unix())
}
