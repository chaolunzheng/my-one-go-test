package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
发现重复的行 v2
 */
func main() {
	//定义一个哈希表
	counts := make(map[string]int)
	//命令行 输入一系列文件名字
	files := os.Args[1:]
	if len(files) == 0 {
		//没有文件就处理键盘标准输入流
		countLines(os.Stdin, counts)
	} else {
		//迭代文件名，抛弃索引
		for _, fileName := range files {
			//打开文件
			file, err := os.Open(fileName)
			//检查文件打开是否错误
			if err != nil {
				//读取文件 发生了错误 直接进行下一个文件读取
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			//处理文件流
			countLines(file, counts)
			//关闭流 方法名字大写表示public的方法，小写表示private的方法
			file.Close()
		}
	}
	//打印出重复行数目和重复行内容
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("内容：%s,重复了：%d次。\n", line, n)
		} else {
			fmt.Printf("内容：%s,没有重复。\n", line)
		}
	}
}

//处理没有文件输入的情况，从标准输入读取
func countLines(fio *os.File, counts map[string]int) {
	//定义一个键盘输入流
	input := bufio.NewScanner(fio)
	//input.Scan() 返回布尔值
	for input.Scan() {
		//对重复行计数
		tem := input.Text()
		//遇到quit结束输入计算结果
		if tem == "quit" {
			break
		}
		//counts是一个引用，指向map本身，故改变了map
		counts[tem]++
	}
}
