package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/**
测试goroutine    类似于多进程   并发访问
 */
func main(){
	var start=time.Now()
	//线程安全的管道
	var ch=make(chan string)

	//使用goroutine
	for _,url:=range os.Args[1:]{
		//go关键字 开启一个go线程
		go fetch(url,ch)
	}

	//处理管道chan的结果
	fmt.Println("耗时\t字节数\turl")
	for range os.Args[1:]{
		fmt.Println(<-ch)//从管道接收内容
	}

	//整个并行过程经历的秒数
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

//并行的 核心提取函数 接收管道
func fetch(url string,ch chan <- string){
	start:=time.Now();
	resp,err:=http.Get(url)
	//处理错误
	if err!=nil{
		ch <- fmt.Sprint(err)
		return
	}

	//计算响应文本大小
	nBytes,err:=io.Copy(ioutil.Discard,resp.Body)
	//关闭流
	resp.Body.Close()
	//copy发生错误
	if err!=nil{
		ch <- fmt.Sprintf("while reading %s: %v",url,err)
		return
	}

	//该url提取耗时
	secs:=time.Since(start).Seconds()

	//输入管道中
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s",secs,nBytes,url)
}