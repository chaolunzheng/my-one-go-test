package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
控制流、条件跳转语句、goroutine的测试
 */

func main() {
	var heads, tails int

	switch coinFlip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("")
	}

	fmt.Println(heads)

	//测试管道队列,用于在线程间传输值
	//var ch=make(chan string)


	go printf()
	//阻塞主线程 否则主线程执行完就结束了 goroutine将来不及执行
	//select {}
	time.Sleep(1*time.Second)

	//-----------我是分割线--------
	fmt.Println("goroutine之间的通信:")
	//chan相当于一个先进先出的队列
	ch:=make(chan string,3)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("numXpu:",runtime.NumCPU())
	for i := 0; i < 3; i++ {
		//匿名函数
		go func(i int) {
			//往ch写值，相当于通信
			ch <- fmt.Sprintf("这是字符串s%d\n",i)
		}(i)//加个括号表示立刻调用
	}
	//获取chan的值
	for i:=0;i<3 ;i++  {
		b:=<-ch
		fmt.Println(b)
	}
	close(ch)
}
func printf() {
	fmt.Println("go go go~")
}

//硬币翻转函数
func coinFlip() string{
	return "heads"
}