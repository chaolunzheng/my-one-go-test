package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
发现重复的行
 */
func main() {
	//定义一个哈希表
	counts := make(map[string]int)
	//定义一个键盘输入流
	input := bufio.NewScanner(os.Stdin)
	//input.Scan() 返回布尔值
	for input.Scan() {
		//对重复行计数
		tem:=input.Text()
		//遇到quit结束输入计算结果
		if tem=="quit"{
			break
		}
		counts[tem]++
	}
	//打印出重复行数目和重复行内容
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("内容：%s,重复了：%d次。\n", line, n)
		}else{
			fmt.Printf("内容：%s,没有重复。\n", line)
		}
	}

	switch a:=10;{
	case a>100:
		fmt.Println(">10")
	default:
		{
			fmt.Println("default")
		}
	}


}
