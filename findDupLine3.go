package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
发现重复的行 v3
 */
func main() {
	//定义一个哈希表
	counts := make(map[string]int)

	//直接获取命令行参数
	for _,file:= range os.Args[1:]{
		//使用io工具获取整个文件内容
		data,err:=ioutil.ReadFile(file)
		//如果发生错误
		if err!=nil{
			//读取文件 发生了错误 直接进行下一个文件读取
			fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			continue
		}
		//使用\n切分整个文件，按行读取----data使用了强制类型转化
		for _,line:=range strings.Split(string(data),"\n"){
			//计数
			counts[line]++
		}
	}

	//打印出重复行数目和重复行内容
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("内容：%s,重复了..：%d次。\n", line, n)
		}else{
			fmt.Printf("内容：%s,没有重复。\n", line)
		}
	}
}
