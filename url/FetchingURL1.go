package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
请求url
 */
func main(){
	const protocol="http://"
	var start,end int64

	for _,url:=range os.Args[1:]{
		start = time.Now().Unix()
		//检测是否有协议头
		if !strings.HasPrefix(url,protocol){
			url=protocol+url
		}
		//通过net包下的http的get请求方法抓取网页
		resp,err:=http.Get(url)
		//发生错误
		if err!=nil{
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s--->状态码：%s\n",url,resp.Status)
		//TODO 方式1
		//解析网页的响应体
		b,err:=ioutil.ReadAll(resp.Body)
		//关闭响应体
		resp.Body.Close()
		if err!=nil{
			fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)

		//TODO 方式2
		//快速拷贝函数 des,src
		io.Copy(os.Stdout,resp.Body)

		end = time.Now().Unix()
		fmt.Println("总耗时：",(end-start))
	}
}
