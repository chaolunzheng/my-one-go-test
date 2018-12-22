package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)
//锁变量
var lock sync.Mutex
var count int

/**
web 服务器测试
 */
func main(){
	//这里/代表匹配所有请求
	http.HandleFunc("/",respHandler1)
	http.HandleFunc("/count",respHandler2)
	http.HandleFunc("/http",respHandler3)

	//匿名函数处理器示例
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		//实现了同样的接口

		//解析请求参数
		rs:=r.Form.Get("r")
		ri,_:= strconv.Atoi(rs)
		//返回一个gif到浏览器响应
		lissajous(w,ri)
	})
	//绑定服务器地址和端口
	log.Fatal(http.ListenAndServe("win10:8080",nil))
}

//处理器1
func respHandler1(w http.ResponseWriter,r *http.Request){
	//同步开始 对共享变量进行同步 控制每次只有一个线程来访问
	lock.Lock()
	count++
	//同步结束
	lock.Unlock()
	fmt.Fprintf(w,"url -> %s",r.URL.Path)
	fmt.Printf("来自%s的请求...\n",r.Host)
}
//处理器2
func respHandler2(w http.ResponseWriter,r *http.Request){
	lock.Lock()
	fmt.Fprintf(w,"一共访问了：%d 次。\n",count)
	//同步结束
	lock.Unlock()

	fmt.Printf("来自%s的请求...\n",r.Host)
}

//打印出请求形式
func respHandler3(w http.ResponseWriter,r * http.Request){
	//打印出协议
	fmt.Fprintf(w,"%s\t%s\t%s\n",r.Method,r.URL,r.Proto)
	//打印header
	for k,v:=range r.Header{
		fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
	}

	//打印Host信息
	fmt.Fprintf(w,"Host=%q\n",r.Host)
	//打印远程连接地址
	fmt.Fprintf(w,"RemoteAddr=%q\n",r.RemoteAddr)

	//判断错误
	if err:=r.ParseForm(); err!=nil{
		log.Print(err)
	}

	//打印形式
	for k,v:=range r.Form{
		fmt.Fprintf(w,"Form[%q]=%q\n",k,v)
		if k=="num"{
			tem, err :=strconv.Atoi(v[0])
			if err==nil{
				fmt.Println("获取参数：",tem*10)
			}
		}

	}
}

func lissajous(out io.Writer,r int) {
	var palette = []color.Color{color.White, color.Black}//切片

	//枚举 编译时就确定下来的
	var cycle =5.0//循环
	const(

		res=0.001//尖角
		size=100//画布大小
		nframes=64//动图数
		delay=8//
	)
	if r!=0{
		cycle=float64(r)
	}
	freq:=rand.Float64()*3.0//y的震荡频率
	anim:=gif.GIF{LoopCount:nframes}//结构体
	phase:=0.0//相位差
	for i:=0;i<nframes;i++{
		//方形画布位置
		rect:=image.Rect(0,0,2*size+1,2*size+1)
		img:=image.NewPaletted(rect,palette)
		for t:=0.0;t<cycle*2*math.Pi;t+=res{
			//函数
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),1)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}