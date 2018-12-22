package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

/**
图像开发包
 */
//包级别的 全局变量
var palette = []color.Color{color.White, color.Black}//切片
//颜色枚举
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	//标准输出
	lissajous(os.Stdout)

}
//图形函数
func lissajous(out io.Writer) {
	//枚举 编译时就确定下来的
	const(
		cycle =5//循环
		res=0.001//尖角
		size=100//画布大小
		nframes=64//动图数
		delay=8//
	)
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
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}
