package main

import (
	"fmt"
	"unsafe"
)

//声明全局变量
var (
	num1 int
	num2 int
)

var aa,bb,cc int32 //全局常量可声明不使用

//枚举
const (
	con1=1
	con2=2
)

//这是第一go程序
func main() {
	/** 打印实例 */
	fmt.Println("****************")

	//变量声明 并打印
	var i int = 5
	fmt.Print(i, "\n")

	//自动推断类型
	var u = 1
	fmt.Println(u)

	var a, b int
	a, b = 1, 2
	fmt.Println(a, b)

	//:声明写法
	e, f, g := 4, 5, 6
	fmt.Println(e, f, g)

	//字符串声明
	var str string = "你好，世界"
	fmt.Println(str)

	str2 := "xixi,haha"
	fmt.Println(str2)

	var str3 = "hello, world!"
	var str4 string
	fmt.Println(str3, "**"+str4+"**")

	//bool形声明
	bo := true
	var bo1 bool //默认false
	fmt.Println(bo, bo1)

	//混合声明
	var integer, strs = 111, "嘻嘻"
	fmt.Println(integer, strs)

	//引用公共变量
	num1 += 1
	fmt.Println(num1, num2)

	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量
	_,_,res:=add(1,1)
	fmt.Println("抛弃的：",res)

	//常量定义 ,可以不使用
	const c1,c2,c3 = 1,'c',"string"
	fmt.Println(c1,c2,c3,len(c3),"sizeOf:",unsafe.Sizeof(c3))

	//iota使用,行计数器
	const(
		i1=iota
		i2
		i3
		i4
		s="---------"
		i5=iota
	)
	fmt.Println(i1,i2,i3,i4,s,i5)

	//枚举常量示例2
	const (
		j1 = 1<<iota
		j2 = 3<<iota
		j3 //默认同上一个运算方式
		j4 = iota
		j5
	)
	fmt.Println(j1,j2,j3,j4,j5)


	as := "hello"
	fmt.Println("hello的sizeOf为",unsafe.Sizeof(as))


	var aq int = 4
	var bq int32
	var cq float32
	var ptr *int
	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", aq )
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", bq)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", cq )

	/*  & 和 * 运算符实例 */
	ptr = &a    /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);
	fmt.Printf("*ptr 为 %d\n", *ptr);

	//for循环写法
	for i = 0; i < 10; i++ {
		fmt.Print(i)
	}
}

//定义有多个返回值的函数
func add(n1 int, n2 int) (int, int, int) {
	return n1,n2,n1+n2
}
